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

when we do a command, we have a specific structure of the syntax, we start by selectiog a daabase, then a collection, a method, and inside the method we pass data.\
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
by defaults, filter use equality
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

### Understanding `$in` and `$nin`

### `$or` and `$nor`

### Understanding the `$and` Operator

### Using `$not`

### Diving Into Element Operators

### Working with `$type`

### Understanding Evaluation Operators - `$regex`

### Understanding Evaluation Operators - `$expr`

### Assignment 3: Time to Practice - Read Operations

### Diving Deeper Into Querying Arrays

### Using Array Query Selectors - `$size`

### Using Array Query Selectors - `$all`

### Using Array Query Selectors - `$elemMatch`

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