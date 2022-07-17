<!--
// cSpell:ignore Schwarzmüller 
 -->

# Master MongoDB Development for Web & Mobile Apps. CRUD Operations, Indexes, Aggregation Framework - All about MongoDB!

 
udemy course [Master MongoDB Development for Web & Mobile Apps. CRUD Operations, Indexes, Aggregation Framework - All about MongoDB!](https://www.udemy.com/course/mongodb-the-complete-developers-guidey/) by *Maximilian Schwarzmüller*. 

- Introduction
- Understaning The Bascis and CRUD Operations
- Schemas & Relations: How to Structure Documents
- Exploring The Shell & The Server
- Using The MongoDB Comapass To Explore Data Visually
- Diving into Create Operations
- Read Operations - A Closer Look
- Update Operations
- Understanding Delete Operations
- Working With Indexes
- Working With GeoSpatial Data
- Understanding the Aggregation Framework
- Working With Numeric Data
- MongoDB & Security
- Performance, Fault Tolerancy & Deployment
- Transactions
- From Shell To Driver
- Introducing Stitch
- Takeaways
  - 

## Takeaways
<!-- <details> -->
<summary>
Things worth remembering
</summary>

[configuration files documentation](https://www.mongodb.com/docs/manual/reference/configuration-options/)


default port is 27017

- the `{$set:{}}` is used inside update commands.
- we can't use `.pretty()` after `findOne`.
- matching a value greater than a threshold `db.flightData.find({distance:{$gt:10000}})`
- `update` doesn't care if we forget the `{$set:{}}` part, it will replace the entire document.
- the **_id** field in always included in projections, unless excluded with `{_id:0}`.
- Nested Documents Limits:
  - up 100 levels of nesting.
  - max size of the document is 16MB.

- `db.dropDatabase()`
- `db.myCollection.drop()`
- `db.customers.aggregate([$lookup:{from: "books",localField: "favBooks",foreignField:"_id",as: "favBookData"}])` - merge documents.
- `db.runCommand({colMod:"posts",validator:{}},validationLevel:"warn"})` - update validation schema and validation action
- `use <db>` - switch to a database
	<samp>
	switched to db shop
	</samp>
- `db.products.insertOne()`
- `db.products.find()`
- `.pretty()`
- `mongoimport <path/to/file.json>` - import a file into a database. [documentation](https://docs.mongodb.com/manual/reference/program/mongoimport/index.html)
  - `-d` - database to use
  - `-c` - collection to use
  - `--JsonArray` - when we have an array of elements, not just one document
  - `--drop` - if collection exists, drop it (clear contents) before importing, otherwise it's an append operation
- the `$` sign in an update refers to the first element matched by `$elemMatch`.
- the `$[]` syntax in an update refers to all elements in the array.
- the `$[<el>]` syntax in an update to target other elements in the array based on different conditions

### Indexes and Explain

- `db.mycoll.getIndexes()` - view all indexes.
- `db.mycoll.createIndex({"field.to.index":1})` - create an index based on the field, either ascending or descending.
  - `db.mycoll.createIndex({"field1":1,"field2":-1})` - compounded index, the order matters!
  - `db.mycoll.createIndex({"field1":1},{unique:true})` - make an index that enforces uniqueness.
  - `db.mycoll.createIndex({"field1":1},{partialFilterExpression:{"field":{$gt:value}}})` - partial filter, when we know we have a sub segment of relevant values in the field which are used most of the time, and other which are rarely used.
  - `db.mycoll.createIndex({createdAt:1},{expireAfterSeconds:10})` - self destroying documents, will be removed after Time to Live expires.
  - `db.mycoll.createIndex({field:1},{background:true})` - background index, doesn't lock up the collection.
  - `db.mycoll.createIndex({textField:"text"})` - text field index, no stop words, one per collection.
  - `db.mycoll.createIndex({textField:"text"},{default_langague:"langauge"})` - text field based on a differnet language.
  - `db.places.createIndex({locationField:"2dsphere"})` - an geospatial index.
- `db.mycoll.dropIndex({"field.to.index":1})` - remove an index.
- `db.mycoll.explain().find({})` - provide a detailed explainnation of how the operation was performed.
  - `db.mycoll.explain("executionStats").find({})` - more verbose explainnation about the execution.


### Find Operators

- `db.collection.find({"key":{$gt:1000}})` - find based on a criteria.
- `db.collection.find({"array":{$elemMatch:{"key":"value"}}})` - find a document where the array contains an element with the properties given.
- `db.collection.find({"array":"element"})` - find a document where the array **contains** the element.
- `db.collection.find({"array":["element"]})` - find a document where the array **has only the element**.
- `db.collection.find({key:{$in:["value1","value2"]}})` - all document where the key is one of the values.
- `db.collection.find({key:{$nin:["value1","value2"]}})` - all document where the key is not one of the values.
- `db.collection.find({$or:[{"criteria1":"value1"},{"criteria2:{$gt:1}}]})` -  match one of the filters.
- `db.collection.find({$nor:[{"criteria1":"value1"},{"criteria2:{$gt:1}}]})` -  match none of the filters.
- `db.collection.find({$and:[{"criteria1":"value1"},{"criteria2:{$gt:1}}]})` -  match all of the filters. required if we want conditions on the same field.
- `db.collection.find({key: {$not:{$eq:value}}})` - inverse a query.
- `db.collection.find({key: {$exists:true, $ne:null}})` - a field exists and has a non-null value.
- `db.collection.find({"field.key": {$type:"number"}})` - a field has a certain type.
- `db.collection.find({"field": {$regex:/pattern/}})` - match a regex pattern.
- `db.collection.find({$expr:{$gt:["$field1","$field2"]}})` - find documents where the fields matchs an expression (boolean).
- `db.collection.find({$expr: {$gt:[{$cond:{if:{$le:[$field",value]],then:"A", else:"B"}},"$value2"]}})` - create conditional value.
- `db.collection.find({"arrayField.innerField":"value"})` - match an internal element of the array.
- `db.collection.find({arrayField:{$size:2}})` - exact match of size (can't compare and use operators here).
- `db.collection.find({arrayField:{$all:["value1","value2"]}})` - match documents which contain all the required values, without caring about order of if there are additional elements.
- `db.collection.find({arrayField:{$elemMatch:{"innerField1":"value1","innerField2":"value2"}}})` - match documents which have an elements in the array that matches all the required conditions.

#### Text Index Search

- `db.collection.find({$text:{$search:"value"}})` - search using the text index of the collection.
- `db.collection.find({$text:{$search:"value1 value2"}})` - match any of the values.
- `db.collection.find({$text:{$search:"\"value1 value2\""}})` - exact match of both values.
- `db.collection.find({$text:{$search:"value1 value2"}},{score:{$meta:"textScore}})` - text search match score.
- `db.collection.find({$text:{$search:"value1 -value2"}})` - exclude words from search
- `db.collection.find({$text:{$search:"value1", $language:"english"}})` - determine search language
- `db.collection.find({$text:{$search:"value1", $caseSensitive:true}})` - force case sensitive search.
- 
### Update Operators

- `db.collection.updateOne({},{})` - update the first matching document.
- `db.collection.updateMany({},{})` - update all matching documents
- `db.collection.updateOne({},{$set:{field:value}})` - set the value of a field, not effecting other fields
- `db.collection.updateOne({},{$inc:{field:value}})` - change a value from it's current value, either increment of decrement (by passing a negative value).
- `db.collection.updateOne({},{$mul:{field:factor}})` - change a value from it's current value by a numeric factor. One being neutral and value lower than one making it smaller.
- `db.collection.updateOne({},{$min:{field:value}})` - the value will be the minimum value between the existing value and the new value.
- `db.collection.updateOne({},{$max:{field:value}})` - the value will be the maximum value between the existing value and the new value.
- `db.collection.updateOne({},{$unset:{fieldName:""}})` - remove the field from the document, the `""` is a common value, but it doesn't matter what we pass.
- `db.collection.updateMany({},{$rename:{oldName:newName}})` - change the name of the field. doesn't add the field to documents which didn't have it.
- `db.collection.updateMany({array:{$elemMatch:{}}},{$set:{"array.$.field":value}})` - change only the array element which was matched. 
- `db.collection.updateMany({},{$set:{"array.$[el].field":value}},{arrayFilters:[{"el.field":value}]})` - target additional elements in the array based on a criteria.
- `db.collection.updateOne({},{$addToSet:{arrayField:{field1:value1,field2:value2}}})` - add unique element to array, doesn't create duplications.
- `db.collection.updateOne({},{$push:{arrayField:{$each:[{field1:value1,field2:value2},{field1:value1,field2:value2}]}})` - add multiple elements to array.
- `db.collection.updateOne({},{$pop:{arrayField:1}})` - remove last element from array. 
- `db.collection.updateOne({},{$pop:{arrayField:-1}})` - remove first element from 
- `db.collection.updateOne({},{$pull:{arrayField:{criteriaField:value}}})` - remove elements from array based on conditions.

### Operators
operator syntax | name | context | notes |
---|-------|--------|-------
`$eq`, `$neq` | logical operators | find |
`$gt`, `$gte` | logical operators | find |
`$lt`, `$lte` | logical operators | find |
`$in`, `$nin` | logical operators | find |
`$or` | logical operators | find |
`$nor` | logical operators | find |
`$and` | logical operators | find |
`$not` | logical operators | find |
`$exists` | test existence | find |
`$type` | test type of field | find |
`$expr` | test type of field | find |
`$jsonSchema` | | find |
`$mod` | | find |
`$text` | | find |
`$regex` | | find |
`$near`, `$geometry` | geospatial data | find |
`$minDistance`, `$maxDistance` | geospatial data | find |
`$geoWithin`, `$geoIntersect` | geospatial data | find |
`$centerSphere` | geospatial data | find | coordinates of the center and radians distance
`$cond`,`if`,`then`,`else` | | find |
`$` | field name specifier | find
`$substract` | | find
`$size` | array size | find
`$all` | match all elements in array | find
`$elemMatch` | match element in array | find |
`$` | | projection
`$slice` | | projection
`$meta` | | projection
`$set` | set field | update | new value to field
`$unset` | unset field | update | remove field from document
`$rename` | rename field | update | give new name to field
`$inc` | increment field value | update | 
`$min` |  | update | 
`$max` |  | update | 
`$mul` |  | update | 
`$push` |  | update | 
`$addToSet` |  | update | 
`$pop` |  | update | 
`$pull` |  | update | 
`$each` |  | update | 
`$` |  | update | 
`$[]` |  | update | 
`$[<identifier>]` |  | update | 

### Cursor Object

- `const cursor = db.colllection.find()` - create cursor.
- `cursor.hasNext()` - check if it was exhausted 
- `cursor.next()` - get next batch, same as `it`
- `cursor.count()` - count elements.
- `cursor.sort({field1:1, field2:-1})` - sort elements.
- `cursor.skip(50)` - move the cursor forward.
- `cursor.limit(5)` - change the number of elements in each fetch.


### Aggregation Steps
- `{$match:{field:value}}` - like `find` operations.
- `{$group: {_id:{key:"$valueFromDocument"},newField: {$sum: 1}}}}` - group into a new document with a aggregation operator:
  - `{$sum:1}` - sum - each value is 1 - basically count.
  - `{$avf:"$fieldToAverage"}` - average of field.
- `{$sort:{key:-1}}` - sort documents at any stage.


### Additional Options Arguments

inserts:
- ordered inserts - do insert for all or stop on the first error.
- write concern - how data is inserted (`insertMany`)
	- *w* - number of instances to write to (default 1, zero means no validation).
	- *j* - write to journal, (default undefined/ false)
	- *wtimeout* - time to wait until for response.

```js
db.mycoll.insertMany([{_id:1,name:"a"},{_id:2, name:"b"}],
{
	unordered:false,
	writeConcern: {
		w: 1,
		j: true,
		wtimeout: 200
	}
})
```

updates:
- upsert: update or insert, if document doesn't exists. create it.
```js
db.mycoll.updateOne({field:value},{$set:{field1:value1,field2:value2}},
{
	upsert:true
})
```
- array filters: target other elements in the array
```js
db.mycoll.updateMany({field:value},{$set:"array.$[identifier].field":value},{arrayFilters:[{"identifier.field":value}]})
```

  - `` - partial filter, when we know we have a sub segment of relevant values in the field which are used most of the time, and other which are rarely used.
   
index
- unique
- partialFilerExpression
- unique + partialFilerExpression to allow multiple null values
- expireAfterSecond
- text index
  - default language
  - weights
- background index
```js
db.mycoll.createIndex({"field1":1},{unique:true})
db.mycoll.createIndex({"field1":1},{partialFilterExpression:{"field":{$gt:value}}})
db.mycoll.createIndex({"field1":1},{unique:1,partialFilterExpression:{"field1":{$exists:true}}})
db.mycoll.createIndex({createdAt:1},{expireAfterSeconds:10})
db.mycoll.createIndex({textField:"text"})
db.mycoll.createIndex({textField1:"text",textField2:"text"})
db.mycoll.createIndex({textField:"text"},{default_language:"english"})
db.mycoll.createIndex({textField1:"text",textField2:"text"},{weights:{textField1:1, textField2:3}})
db.places.createIndex({locationKey:"2dsphere"})
db.mycoll.createIndex({field:1},{background:true}})
```

### Special types of objects

- point: geospacial data
```js
{
	"location": 
	{
		"type": "Point",
		"coordinates": [56.12,43.09],
	}
},
{
  "locationField":
  {
    "type": "Polygon",
    "coordinates": [[[],[],[],[],[]]]
  }
}
```
### 

run mongodb as background service in windows
```cmd
net start MongoDB
net stop MongoDB
```

to quit 
```sh
use admin
db.shutdownServer()
```

### CLI flags

when running `mongod` - [documentation](https://www.mongodb.com/docs/manual/reference/program/mongod/)

- `--port` - port to run the service on
- `--quiet` - reduce verbosity, less output.
- `--logpath <path to file>` - where the logs are stored.
- `--dbpath <folder>` - where the data is actually stored.
- `--repair` - try and fix corruptions in database
- `--directoryperdb` - group databases into sub folders
- `--fork` - only for mac and linux. run as a background process, as a service. must have a log path (can't log to the terminal).
- `--storageEngine <engine>` - default is wiredTiger
- `--config` (or `-f`) - pass a mongod configuration file

when running `mongosh` or `mongo` - [documentation](https://www.mongodb.com/docs/manual/reference/program/mongo/)

- `mongo --help` - get help for shell
- `--nodb` - don't run with a database, just a js shell
- `--quiet` - less verbose output
- `--verbose` - more verbose
- `--port` - which port to connect, default is local host 27017
- `--host` - which host to connect
- `-u` - authentication, username
- `-p` - authentication, password


### Mongo Shell Commands
command | action
----|----
`db.help()` | help on db methods
`db.mycoll.help()`| help on collection methods
`db.stats()`| help on collection methods
`sh.help()` | sharding helpers
`rs.help()` | replica set helpers
`db.dropDatabase()` | clear current database
`db.shutdownServer()` | shut down
`show dbs`| list databases
`show collections`| list collections in current database
`show users`| list users
`show profiles`|list profiles
`help` | general help
`help admin` | administrative help
`help connect` | connecting to a db help
`help keys` | key shortcuts
`help misc` | misc things to know
`help mr` | mapreduce

### Data Types
DataType | Notes | Example
---|---|---
Text | always quotes | "Max"
Boolean | true of false | true
double | floating point, default| 12.25
Integer | int32 | 55, `NumberInt(11)`
NumberLong | int64, use "" for large numbers | 1000000000, `NumberLong("1000000000")`
NumberDecimal | High precision, use the "" to be precise | 12.99, `NumberDecimal("11.95")`
ObjectId | automatically generated, has a timestamp internally | ObjectId("text")
ISODate | date | ISODate("2018-09-09")
Timestamp| date time |Timestamp(11421532)
Embedded Documents | nesting | {"a":{}}
Array | list of values| {"b":[]}
GeoJson | geospatial data | `{type:"point", coordinates:[]}`

</details>
