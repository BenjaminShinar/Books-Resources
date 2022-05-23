<!--
// cSpell:ignore
-->

[main](README.md)

## Section 7 - Read Operations: A Closer Look

<details>
<summary>
Accessing the Required Data Efficiently.
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

> 1. Import the attached data into a new database (e.g. boxOffice) and collection (e.g. movieStarts).
> 2. Search all movies that have a rating higher than 9.2 and a runtime lower than 100 minutes.
> 3. Search all movies that have a genre of "drama" or "action".
> 4. Search all movies where visitores exceeded excpectedVisitors.

importing the data
```sh
`mongoimport boxoffice.json -d boxOffice -c movieStarts --jsonArray --drop`
```

tasks
```sh
db.movieStarts.find({"meta.rating":{$gt:9.2},"meta.runtime":{$lt:100}}).pretty()
db.movieStarts.find({genre: {$in:["drama","action"]}}).pretty()
#doesn't work with shell 3.4
db.movieStarts.find({$expr:{$gt:["$visitors","$expectedVisitors"]}}).pretty()
```

cleaning up
```js
db.boxOffice.drop()
db.dropDatabase()
```
### Diving Deeper Into Querying Arrays
there special operator which help us with querying arrays.

we can look at our earlier 'users' collection. we used nested documents there, so simply searching for the value doesn't work. we need to search inside the document, without specifying the exact object structure.

we can use mongodb integrated array functionalities. it knows to match all the elements in the arrays and search all of them.

```sh
use users
# no matchs
db.users.find({hobbies: "Sports"}).pretty()
# no matches either, search for an exact match of document equality.
db.users.find({hobbies: {"title": "Sports"}}).pretty()
# this works
db.users.find({"hobbies.title": "Sports"}).pretty()
```

there are also dedicated query selectors, which work on arrays.

#### Using Array Query Selectors - `$size`

querying the size of an array.

```sh
db.users.insertOne({name: "Chris",hobbies: ["Sports","Cooking","Hiking"]})
db.users.find({hobbies: {$size: 3}})
```

mongo db currently doesn't support matching the size with an operator, like finding documents with more than a specified amount of elements.

#### Using Array Query Selectors - `$all`

we want to match documents who have the requested elements, but without caring about the order in which they appear. it will also match any document that contains the required elements, even if the document has additional elements.
```sh
use boxOffice
# matchs ["action", "thriller"]
db.movieStarts.find({"genre":["action","thriller"]}).pretty()
# matchs ["thriller","action"] - but not ["action", "thriller"]
db.movieStarts.find({"genre":["thriller","action"]}).pretty()
# matchs both documents as above, as well as the third document
db.movieStarts.find({"genre":{$all: ["thriller","action"]}}).pretty()
```

#### Using Array Query Selectors - `$elemMatch`

we want to find documents which have an elements that matches a criteria, and we want the document to have an element that matches all the criteria, rather than having one element which matches the first criteria, and maybe a different element matches the other condition.

```sh
use users
# oops! it can match the two conditions in different elements!
db.users.find({$and:[{"hobbies.title":"Sports},{"hobbies.frequency":{$gt:3}}]}).pretty()
# match all conditions on a single element inside the array
db.users.find({kids :{$elemMatch: {$gt:30,$lt:50}}})
db.users.find({hobbies: {$elemMatch: {title:"Sports",frequency: {$gte:3}}}})
```

### Assignment 4: Time to Practice - Array Query Selectors

> 1. Import the attached data into a new collection (e.g. exMovieStarts) in the boxOffice database.
> 2. Find all movies with exactly two genres.
> 3. Find all movies which aired in 2018.
> 4. Find all movies which have rating greater than 8 but lower than 10.

importing data
```sh
mongoimport boxoffice-extended.json -d boxOffice -c exMovieStarts --jsonArray --drop
```

tasks

```sh
db.exMovieStarts.find({genre:{$size:2}}).pretty()
db.exMovieStarts.find({"meta.aired":2018}).pretty()
#db.exMovieStarts.find({$and:[{"ratings":{$gt:9.5}},{"ratings":{$lt:10}}]}).pretty()
db.exMovieStarts.find({ratings:{$elemMatch:{$gt:8,$lt:10}}}).pretty()
```

### Understanding Cursors

the `find()` method returns a cursor, unlike the `findOne()` method, which returns a single document.
A cursor is a pointer object that stores a position in the database, and we an use it to fetch batches of objects. the shell has a default of 20 documents, which we can change, or use a different value when we connect to a database using a mongodb Driver.

#### Applying Cursors


```sh
use MovieDate
db.movies.find().count()
```

`.count()` is already a cursor function, `it` gets us the next batch, but in the driver it's usually called `.next()`, but we need to strore the cursor, otherwise it will re-run the same query.

```js
const dataCursor = db.movies.find()
dataCursor.next()
dataCursor.next()
```
we can also use arrow functions on the elements in the cursor. this will execute on all the remaining documents. we can also add conditions, but it's more efficient to add the filters to query.
```js
dataCursor.forEach(doc => {printjson(doc)})
dataCursor.hasNext() //false
```

#### Sorting Cursor Results

we can sort the elements in the cursor, either in ascending or descening order. we can sort by multiple fields, using the order which we pass
the field.

```sh
db.movies.find({}).sort({"rating.average":1}).pretty()
db.movies.find({}).sort({"rating.average":1,runtime:-1}).pretty()
```

#### Skipping & Limiting Cursor Results

we might want to skip results, like if we implement pagination, there is no raeson to fetch data we don't care about. we can also change the batch size per iteration.

```sh
db.movies.find({}).skip(15)
db.movies.find({}).limit(5)
```

the order of the cursor functions doesn't matter, skip,sort and limit will always execute in the same order
1. sort
2. skip
3. limit

so\
`db.movies.find({}).sort({runtime:-1}).skip(10).limit(5)`\
is the same as:\
`db.movies.find({}).limit(5).sort({runtime:-1}).skip(10)`

this won't be true for the aggregation pipeline.

### Using Projection to Shape our Results

shaping the returned data into a clean format, we want smaller (and more readable) results, and as a bonus, we get better performance. projection is the 2nd argument to the find method.

we can include fields with **1**, or exclude them with **0**. the id field is always returned, unless we explicitly exclude it.

we can also have embedded documents fields.

```sh
db.movies.find({},{name:1, genres:1, runtime:1, rating:1})
db.movies.find({},{name:1, rating:1,_id:0})
db.movies.find({},{name:1, "schedule.time":1})
```

### Using Projection in Arrays

we can use use special syntax to project only the elements of the array we care about, this returns the first match. it's simple when we match for one field, but not if we have a complex find.

we can also project fields that weren't in the find query!
```sh
db.movies.find({genres:"drama"},{"genres.$":1})
db.movies.find({genres: {$all:["drama","horror"]}},{"genres.$":1})
db.movies.find({genres:"drama"},{"genres":{$elemMatch:{$eq:"horror"}}).pretty()
```

### Understanding `$slice`

the `$slice` operator allows us to control how many elements we project, or ever which elements, by specifying how many elements to skip, and then how many to show
```sh
db.movies.find({"rating.average":{$gt:9}},{name:1,genres: {$slice:2}})
db.movies.find({"rating.average":{$gt:9}},{name:1,genres: {$slice:[1,2]}})
```

</details>

## Section 8 - Update Operations

<!-- <details> -->
<summary>

</summary>


### Updating Fields with `updateOne()`, `updateMany` and `$set`
### Updating Multiple Fields with `$set`
### Incrementing & Decrementing Values
### Using `$min`, `$max` and `$mul`
### Getting Rid of Fields
### Renaming Fields
### Understanding `upsert()`
### Assignment
### Updating Matched Array Elements
### Updating All Array Elements
### Finding & Updating Specific Fields
### Adding Elements to Arrays
### Removing Elements from Arrays
### Understanding `$addToSet`
### Wrap Up
### Useful Resources & Links


</details>


##

[main](README.md)
