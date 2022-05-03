<!--
// cSpell:ignore
-->

[main](README.md)

## Section 3 - Schemas & Relations: How to Structure Documents
<!-- <details> -->
<summary>
</summary>

when we start a project, we first decide on how our data is modeled, and what types of relations exists. we use document schemas and data types to describe the data in our documents. we also have ways to model relations between entities in the database. in addition, we will also want to validate incoming data, and make sure it fits our schema.

### Why Do We Use Schemas?

In theory, mongoDB doesn't enforce any schema, we can have any kind of documents inside a collection. documents have a unique id, but there aren't any other requirements, each document can have entirely different fields.\
However, in the real world use-cases, we would want the documents to have some common fields (a schema).

### Structuring Documents

in SQL, all tables have a schema, and all entries are structured the same way. in mongoDB, we can have complete freedom, where each entry can have a different schema. mongoDB also allows us to find a middle ground, having a structure, and also having the possibility to have extra data.

we can use the `null` value to say that a field exists, but there is no value for it.it makes the structure a bit more clear. the mongoDB approach is to omit the fields without a value.

### Data Types - An Overview

DataType | Notes | Example
---|---|---
Text | always quotes | "Max"
Boolean | true of false | true
Integer | int32 | 55
NumberLong | int64 | 1000000000
NumberDecimal | High precision | 12.99
ObjectId | automatically generated, has a timestamp | ObjectId("text")
ISODate | date | ISODate("2018-09-09")
Timestamp| date time |Timestamp(11421532)
Embedded Documents | nesting | {"a":{}}
Array | list of values| {"b":[]}

the shell always uses floats to represent numbers. this is because the shell is based on JavaScript, which doesn't have diffrent numeric types.

```sh
use companyData # switch to other DB
db.companies.insertOne({name:"Fresh Apples Inc",isStartup:true, employess: 33, funding:12345678901234567890,details:{"ceo":"Mark Super"}, tags:["super","perfect"]}, foundingDate: new Date(), insertedAt: new Timestamp())
db.companies.findOne()
```
`new Data()` and `new Timestamp()` are built in shell functions that construct the current date and time (as a timeStamp format, epoch time),

but when we look at the outcome, the value of *funding* isn't what we entered, it got truncated.

```sh
db.numbers.insertOne({a:1})
db.stats()
db.numbers.deleteMany({})
db.numbers.insertOne({a:NumberInt(1)})
db.stats()
typeof db.numbers.findOne().a
```

there are also different ways of how data is stroed in BSON, see documentation.

### How to Derive your Data Structure - Requirements

some guidelines

Guiding Question | Examples | Actions
---|---|---
"Which Data does my application need or generates?" | User Information, Product Information, Orders | Defines the Fields you'll need (and how they relate)
"Where do I need my data?" | Welcome Page, Products List Page, Orders Page | Defines your required collections + field groupings
"Which kind of Data or Information do I want to display?" | Welcome Page, Product Names, Product Page | Defines which queries you'll need
"How often do I fetch my Data?" | every page, every second, on demand | Defines whether you should optimize for easy fetching
"How often do I write or change my Data" | Orders-> often, Product Data-> rarely | Defines whether you should optimize for easy writing

MongoDb core principal is to design the structure to fit the usage, so there won't be many joins and lookup and other operations across collections.

### Understanding Relations

we usually have multiple collections inside our database, the documents are usually related to one another,

- Nested/Embedded Documents
- References

an address can be stored as part of the customer data
```json
{
  "userName":"max",
  "age":25,
  "address":{
    "street": "second street",
    "city": "new York"
  }
}
```
but we can also store a reference (identifier, foreign key) in one document to a document in another collection, usually in cases where the nested data is shared across many documents. 
```json
// users
[
 {
   "userName":"A",
   "favoriteBooks":["id1",]
 },  
 {
   "userName":"B",
   "favoriteBooks":["id2","id3"]
 },
 {
   "userName":"C",
   "favoriteBooks":["id1","id3"]
 }  
  
]
// books
[
  {
  "_id":"id1",
  "name":"lord of the rings",
  "author": "tolkien",
  "publication":1820
  },
]
```
### One To One Relations 
if we have a one to one relation, such as a patient in a hospital and a summary of their case. each patient has a unique case summary. 

as a reference
```sh
db.patients.insertOne({name:"Max", age:29, diseaseSummary:"summary-max-1"})
db.diseaseSummaries.insertOne({_id:"summary-max-1",details:["as","s"]})

db.patients.findOne({name:"Max"})
var diseaseId = db.patients.findOne({name:"Max"}).diseaseSummary
db.diseaseSummaries.findOne({_id:diseaseId})
db.patients.deleteMany({})
```
but as an embedded document, this would be a simple call
```
db.patients.insertOne({name:"Max", age:29, diseaseSummary:{details:["cold","39 celsius"]}})
db.patients.findOne({name:"Max"})
```

however, there are cases where one-to-one relations work better with references, in our example, a person has a car, one car per person, one person per car.

```sh
db.drivers.insertOne({name:"Max", age:29, car:{model:"bmw",licenseId:12345}})
db.drivers.findOne({name:"Max"})
```
but maybe we don't really care about exploring the relationship between the persons and the cars, maybe we just analyze the cars in some cases, and the drivers in other cases, but we don't really care about who drives which car in terms of analyzing. in these cases, having an embedded document just forces us to use more projection on our data and bloats our queries.\
we could store the drivers and the cars apart from one another, and use references from one to the other in the rare cases that we do care about joining the data together.

### One To Many 

one to many - like one question with many answers, we can store references to objects in a different collection.

```sh
use support
db.questionThreads.insertOne({creator:"max",question "how does this work?",answers:["q1a1","q1a2"]})
db.answers.insertMany([{_id:"q1a1",answer:"aa"},{_id:"q1a2",answer:"ab"}])
```

alternately,we could store the answers inside the question objects. in this use case, embedding them makes sense.

```sh
use support
db.questionThreads.insertOne({creator:"max",question "how does this work?",answers:[{,answer:"aa"},{,answer:"ab"}]})
```

a different case might be the population of a city, we can store all the citizens of a city inside the city document, but that would mean storing millions of complete records inside the city object, even storing all the references(ids) can be too much. it'll be easier to have to two different collections, and query the citizens collection when needed.

we don't want to embed too much data, we remember that documents have a size limit!

### Many To Many 

a many to many example is many customers, each buying many products. we usually do this with references, and we might even have a relationship collection, which is the SQL way to do this.

```sh
use shop
db.products.insertOne({_id:"productA"})
db.customers.insertOne({_id:"customer1"})
db.orders.insertOne({productId:"productA",customerId:"customer1"})
```
but the mongo way is to use only two collections, and store the id as a referece
```sh
db.customers.updateOne({_id:"customer1"},{$set:{orders:[{productId:"productA",quantity:3}]}})
```
or to store it as a nested objects. but this might be a source of data duplications, and update to the nested documents will also have to be replicated. this might not be relevent if future changes don't affect existing copies, but this depends on the use case.

### Summarizing Relations
### Using `lookUp()` for Merging Reference Relations
### Planning the Example Exercise
### Implementing the Example Exercise
### Understanding Schema Validation
### Adding Collection Document Validation
### Changing the Validation Action
### Wrap Up
### Useful Resources & Links


</details>



##

[main](README.md)


