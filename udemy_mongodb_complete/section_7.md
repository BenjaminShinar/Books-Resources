<!--
// cSpell:ignore
-->

[main](README.md)

## Section 7 - Read Operations: A Closer Look

<!-- <details> -->
<summary>
Accessing the Required Data Efficiently
</summary>

we can filter which documents we get, what structure they have (projections) and even transform the data.

we will work with the movies database which we worked with in the previous module.

### Methods, Filters & Operators

when we do a command, we have a specific structure of the syntax, we start by selecting a database, then a collection, a method, and inside the method we pass data.\
`<db>.<collection>.<method>(<filter>)`

for the `find` method, we pass a **filter** as data. the filter can be simple or complex, and can use operators such as `$gt`.

### Operators - An Overview

different types of operators
- Read Operators
  - Query Selectors
  - Projection Operators
- Update Operators
  - Fields
  - Arrays
- Query Modifiers - Deprecated
- Aggregation - complex transfors
  - Pipeline Stages
  - Pipeline Operators

Type | Purpose | Changes Data | Example
---|---|---|---
Query Operators | Locate data | No | `$eq`
Projection Operators | Modify data presentation | No | `$`
Update Operators | Modify and Add additional data | Yes | `$inc`

### Query Selectors & Projection Operators

Query Selectors Types:
- Comparison
- Logical
- Elements
- Evaluation
- Array
- Comments
- Geospatial

Projection Operators:
- `$`
- `$elemMatch`
- `$meta`
- `$slice`

### Understanding `findOne()` & `find()`

lets import the data again
`mongoimport tv-shows.json -d movieData -c movies --jsonArray --drop`

```sh
use MovieData
db.movies.findOne({})
db.movies.find({}).pretty()
#lets add a filter
db.movies.findOne({name:"The Last Ship"})
db.movies.find({runtime:60}).pretty()
```
by defaults, filter use equality.

### Working with Comparison Operators

playing with comparison operators
```sh
# equality same
db.movies.find({runtime:60}).pretty()
# also equality, explicit
db.movies.find({runtime:{$eq:60}}).pretty()
# not equals
db.movies.find({runtime:{$ne:60}}).pretty()
# lower than
db.movies.find({runtime:{$lt:40}}).count()
# lower than or equals
db.movies.find({runtime:{$lte:40}}).count()
```

### Querying Embedded Fields & Arrays

when we have embedded fields (objects and arrays), we can also query them. we do this by specifying the path, in this case we must use quotation marks, other wise the dot is not recognized

```sh
db.movies.find({"rating.average":{$gt:7}})
```

we can also query the elements of an array, by default, mongo searchs for the existence of the element inside the array, it doesn't have to the only element. if we want to search for an exact match (an array with only the single element), we can specify an array as the searched element.

```sh
db.movies.find({"genres": "Drama"}).pretty() # all documents where the array contains "Drama"
db.movies.find({"genres": ["Drama"]}).pretty() # all documents where the array contains only "Drama"
```

#### Understanding `$in` and `$nin`

`$in` and `$nin` have a slightly different behavior, they allows us to match different cases. the arguments are passed in as an array, and we can match to one of them or document which don't match any.
```sh
db.movies.find({runtime: {$in:[30,42]}}) # all documents where the runtime is 30 or 42
db.movies.find({runtime: {$nin:[30,42]}}) # all documents where the runtime is not 30 or 42
```

### Logical Operators

#### `$or` and `$nor`

matching elements based on combined criteria, multiple conditions, we start with the `$or` operator, and pass the filters as an array. we can also use `$nor`, to get documents which don' match any of the criteria.

```sh
db.movies.find({"rating.average":{$lt:5}}).count() #count matching elements
db.movies.find({"rating.average":{$gt:9.3}}).count() #count matching elements
db.movies.find({$or:[{"rating.average":{$lt:5}},{"rating.average":{$gt:9.3}}]}).count() #count matching elements
db.movies.find({$nor:[{"rating.average":{$lt:5}},{"rating.average":{$gt:9.3}}]}).count() #count matching elements
```

#### Understanding the `$and` Operator

find documents who match all of the conditions. this isn't required in basic cases, because we can put everything inside the regular documents. but it's used in some cases, as some mongo drivers don't support documents with a repeated field name and will only use the second defintion. this is very dangerous.

```sh
db.movies.find({$and:[{"rating.average":{$lt:5}},{genres:"Drama"}]}).count()
db.movies.find({"rating.average":{$lt:5},genres:"Drama"}).count()
#match on same field
db.movies.find({genres:"Horror"}).count() # check
db.movies.find({genres:"Drama",genres:"Horror"}).count() # same value not good!
db.movies.find({$and:[{genres:"Drama"},{genres:"Horror"}]) # this works!
```
#### Using `$not`
the `$not` operator inverts the result of a query, in many cases we can use `$neq`.

```sh
db.movies.find({runtime: {$not: {$eq:60}}}).count()
db.movies.find({runtime: {$neq:60}}).count()
```

### Diving Into Element Operators

this operators match on fields, rather than values. we can check if a fields exits, and check that it has a type or a valid value.

```sh
db.user.find({age:{$exists:true}}).pretty() # documents where the field exits
db.user.find({age:{$exists:true, $gt:30}}).pretty() # documents where the field exits and matches a criteria.
db.user.find({age:{$exists:true, $gt:30}}).pretty() # documents where the field exits and matches a criteria.
db.user.find({age:{$exists:true,$ne:null}}) #field exists and is not null
```

#### Working with `$type`
we can match for a specific data type for the field we query.
```sh
db.user.find({phone:{$type: "number"}}) #documents where the field is a number (double or integer)
db.user.find({phone:{$type: "integer"}}) #documents where the field is an integer
```


### Understanding Evaluation Operators - `$regex` and `$expr`

Evaluation operators
- `$expr` - aggregation expressions within the query language.
- `$jsonSchema` - validate document against the given JSON schema
- `$mod` - modules division.
- `$regex` - regular expression.
- `$text` - perform text search
- ~~`$where` - match documents against a javascripts expression~~ - **deprecated**
 
if we want to search for a sub string inside a text field, we can use `$regex`, or the `$text` index operator, if we have it defined. regex expressions don't have quotes. and they are surrounded by `/` marks.

```sh
db.movies.find({summary: {$regex: /musical/}})
```

`$expr` allows us to match fields inside the queried document with itself.

in this example, we want find documents where the "start" field is larger the the "end" field. we pass the operator and the fields as names, we pass the fields name with `$` symbol. We can also have more complex queries, for this we use `$cond`,`if`,`then` and `else`. we can choose which value to use as from a conditional computation.
```sh
use financialData
db.sales.insertMany([{start:10,end:12},{start:12,end:7},{start:7, end:25}])
db.sales.find({$expr: {$gt:["$start","$end"]}})
db.sales.find({$expr: {$gt:
[ 
  {
    $cond:{
      if:{$gte:["$end",10]},
      then:{$subtract: ["$end","$start"]},
      else:"$end"}
  },
  5
]
  }})
```

this behavior leads us into the aggregation pipeline syntax.
### Assignment 3: Time to Practice - Read Operations

### Diving Deeper Into Querying Arrays

#### Using Array Query Selectors - `$size`

#### Using Array Query Selectors - `$all`

#### Using Array Query Selectors - `$elemMatch`

### Assignment 4: Time to Practice - Array Query Selectors

### Understanding Cursors

### Applying Cursors

### Sorting Cursor Results

### Skipping & Limiting Cursor Results

### Using Projection to Shape our Results

### Using Projection in Arrays

### Understanding `$slice`

### Useful Resources & Links

</details>