<!--
// cSpell:ignore
-->

[main](README.md)

## Section 3 - Schemas & Relations: How to Structure Documents
<details>
<summary>
Data Types, Relations, Schema Validation
</summary>

when we start a project, we first decide on how our data is modeled, and what types of relations exists. we use document schemas and data types to describe the data in our documents. we also have ways to model relations between entities in the database. in addition, we will also want to validate incoming data, and make sure it fits our schema.

### Why Do We Use Schemas?

In theory, mongoDB doesn't enforce any schema, we can have any kind of documents inside a collection. documents have a unique id, but there aren't any other requirements, each document can have entirely different fields.\
However, in the real world use-cases, we would want the documents to have some common fields (a schema).

### Structuring Documents

in SQL, all tables have a schema, and all entries are structured the same way. in mongoDB, we can have complete freedom, where each entry can have a different schema. mongoDB also allows us to find a middle ground, having a structure, and also having the possibility to have extra data.

we can use the `null` value to say that a field exists, but there is no value for it.it makes the structure a bit more clear. the mongoDB approach is to omit the fields without a value.

### Data Types - An Overview

| DataType           | Notes                                    | Example               |
| ------------------ | ---------------------------------------- | --------------------- |
| Text               | always quotes                            | "Max"                 |
| Boolean            | true of false                            | true                  |
| Integer            | int32                                    | 55                    |
| NumberLong         | int64                                    | 1000000000            |
| NumberDecimal      | High precision                           | 12.99                 |
| ObjectId           | automatically generated, has a timestamp | ObjectId("text")      |
| ISODate            | date                                     | ISODate("2018-09-09") |
| Timestamp          | date time                                | Timestamp(11421532)   |
| Embedded Documents | nesting                                  | {"a":{}}              |
| Array              | list of values                           | {"b":[]}              |

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

| Guiding Question                                          | Examples                                      | Actions                                               |
| --------------------------------------------------------- | --------------------------------------------- | ----------------------------------------------------- |
| "Which Data does my application need or generates?"       | User Information, Product Information, Orders | Defines the Fields you'll need (and how they relate)  |
| "Where do I need my data?"                                | Welcome Page, Products List Page, Orders Page | Defines your required collections + field groupings   |
| "Which kind of Data or Information do I want to display?" | Welcome Page, Product Names, Product Page     | Defines which queries you'll need                     |
| "How often do I fetch my Data?"                           | every page, every second, on demand           | Defines whether you should optimize for easy fetching |
| "How often do I write or change my Data"                  | Orders-> often, Product Data-> rarely         | Defines whether you should optimize for easy writing  |

MongoDb core principal is to design the structure to fit the usage, so there won't be many joins and lookup and other operations across collections.

### Understanding Relations
<details>
<summary>
One to One, One to Many, Many to Many. Nested documents vs references.
</summary>


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
#### One To One Relations 
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

#### One To Many 

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

#### Many To Many 

a many to many example is many customers, each buying many products. we usually do this with references, and we might even have a relationship collection, which is the SQL way to do this.

```sh
use shop
db.products.insertOne({_id:"productA"})
db.customers.insertOne({_id:"customer1"})
db.orders.insertOne({productId:"productA",customerId:"customer1"})
```
but the mongo way is to use only two collections, and store the id as a reference
```sh
db.customers.updateOne({_id:"customer1"},{$set:{orders:[{productId:"productA",quantity:3}]}})
```
or to store it as a nested objects. but this might be a source of data duplications, and update to the nested documents will also have to be replicated. this might not be relevent if future changes don't affect existing copies, but this depends on the use case.

in some cases it's better to have many-to-many relationship as a reference. imagine that we have book and authors.

this is how an embedded data will look
```sh
use bookRegistry
db.books.insertOne({name: "my book", authors:[{name:"max", dob:"2000-01-13"},{name:"bob",dob:"1995-04-15"} ]})
db.books.find().pretty()
db.authors.insertMany([{name:"max", dob:"2000-01-13",address:{}},{name:"box",address:{},role:"editor",dob:"1995-04-15"}])
```

having a snapshot of the data is ok when the data isn't meant to change, but if we want the data to always be up to date, we would need to update all documents. this is worse if we have a high frequency of updates.

so the better approach is to use references. we might need to run some join commands, but it will be more efficient and have less errors than the nested documents approach.

```sh
use bookRegistry
db.books.insertOne({name: "my book", authors:["id1","id2" ]})
db.books.find().pretty()
db.authors.insertMany([{_id:"id1",name:"max", dob:"2000-01-13",address:{}},{_id:"id2",name:"box",address:{},role:"editor",dob:"1995-04-15"}])
```
#### Summarizing Relations

the correct relation depends on the type of data, the frequency of the updates, and the common use case

>**Nested/Embedded Documents**
>- Group data together locally.
>- Great for data that belongs together and is not really overlapping with other data.
>- Avoid super deep nesting (100+ levels) or extremely long arrays (16mb size limit per document).
>**References**
>- Split data across collections.
>- Great for related but shared data, as well as for data which is used in relations and standalone.
>- Allow you to overcome nesting and size limits (by creating new documents).

</details>

### Using `lookUp()` for Merging Reference Relations

when we have a relation that uses references, we can join the documents together by using the `$lookup` operator. it allows us to merge documents in one command.

```sh
db.customers.aggregate([$lookup:{
  from: "books",
  localField: "favBooks",
  foreignField: "_id",
  as: "favBookData"
}])
```
we need four values, which are passed as a document.

- *from* - which collection to relate to
- *localField* - how the value is called in the current document
- *foreignField* - how the value is called in the related documents
- *as* - the name of the new field which will be displayed
with our previous command.

```sh
db.books.aggregate([$lookup:{
  from:"authors", 
  localField:"authors",
  foreignField:"_id",
  as:"creators"
}])
```
### Example Exercise

we will start with an example project.

we have a user

users can:
- create Post
- edit post
- delete posts
- delete posts
- fetch post (single)
- comment on a post

the user will communicate with an app server, which holds the code itself, as well as the mongoDB driver, which connects to the MongoDB server, and eventually, the Database.

the data entities, and which data do they have
1. users
   1. _id
   2. name
   3. age
   4. email
2. posts
   1. _id
   2. title
   3. text
   4. tags
3. comments
   1. _id
   2. text

we also define the relationships.

- a user can create and delete posts
- a user can comment on a post
- a post has multiple comments, each belonging to different user.


we could model this as one collection, everthing is under the posts.
```json
// posts
{
  "creator":{}, //user
  "title":"",
  "tags":[],
  "text":"",
  "comments":
  [
    {
      "user":{},
      "text":"" //comment text

    }
  ]
}
```

in some sense, this is okay, nesting the comments seems right, as comments belong to a post, and are usually very tightly coupled with a post. but nesting the user might not be the smart idea. each user has many posts (and comments), so nesting the users inside the posts and comments might not be efficient.\
it's probably better to have a users collection and posts collections, but not a comments collection.

```sh
use blog
db.users.insertMany([{_id:"user1",name:"max",email:"a@b.com"},
{_id:"user2",name:"bob",email:"b@a.com"}])
db.posts.insertOne({title:"my post",text:"first post", tags:["a","b","c"]}, creator: "user1", comments:
[
  {text:"First!", author:"user2"},
  {text:"spam!", author:"user1",edits:4}
])
db.posts.findOne()
```

### Understanding Schema Validation

mongodb is flexable,we can have different documents and structures in the same collection. but sometimes we want to set the schema in place. we want to make sure all documents folllow a certain form, having data with a certain name and of certain type.

using schema validation, we can control what kind of documents exist inside the collection, so we have a tighter control over it. if a document doesn't fit the schema, it can't be added.

- Validation level - "Which documents get validated"
  - strict - All inserts and updates
  - moderate - All inserts and update to correct documents.
- Validation Action - "What happens if validation fails"
  - error - throw error and deny insert/update
  - warn - log warning but proceed

### Adding Collection Document Validation

using the posts example.

we can add a schema when we create a collection explicitly, we pass a document, which one of its' key is a **validator**.

```sh
db.posts.drop()
dp.createCollection("posts",{validator: {$jsonSchema:{
  bsonType:"object",
  required:[] ,
  properties:{}
}
}})
```

we can put this in a javascript file instead, to make it easier to read and understand

```js
dp.createCollection("posts",{validator: {$jsonSchema:{
  bsonType:"object",
  required:["title","text","creator","comments"],
  properties:
  {
    title: {
      bsonType: "string",
      description: "must be a string type and is required"
    },
    text: {
      bsonType: "string",
      description: "must be a string type and is required"
    },
    creator: {
      bsonType: "objectId",
      description: "must be an objectId type and is required"
    },
    comments: {
      bsonType: "array",
      descrption: "must be an array"
      required:['text','author'],
      items: {
        bsonType: "object",
        properties:
        {
          text: {
            bsonType:"string"
          },
          author:{
            bsonType: "objectid"
          }
        }
      }
    }

  } 
}
}})
```

now if we try to add document which doesn't match the schema, we will get an error

### Changing the Validation Action

when we have an existing database, we don't want to drop and recreate it, so we can modify it instead,
```sh
db.runCommand({colMod:"posts",validator:{}},validationLevel:"warn"})
```
### Wrap Up

What do we consider?
> - in which format will you fetch your Data?
> - How often will you fetch and change your data?
> - how much data will you save (and how big is it)?
> - how is your data related?
> - will duplicates hurt you (=> many updates)?
> - will you hit data storage limits?
>
> Modelling Schemas
> - Schemas should be modelled based on your application needs
> - important factors are: read and write frequency, relations, amount (and size) of data.
> 
> Modelling  Relations
> - Two options: embedded document or references.
> - Use embedded documents if you got one-to-oe or one-to-many relationships and no app or data size reason to split.
> - Use references if data amount/size or applications needs require it. or for many-to-many relations.
> - Exceptions are always possible. Keep your app requirement in mind.
> 
> Schema Validation
> - You can define rules to validate inserts and updates before writing to the database.
> - Choose your validation level and action based on your application requirements

</details>

## Section 4: Exploring The Shell and The Server
<details>
<summary>
Getting better with the mongo shell and mongo server, what we can configure and how to work with them.
</summary>

### Finding Available Options

we can always learn about options to configure in the documentations, we can look at the mongo shell section.

in the shell

```sh
mongod # start a mongo server
mongo # connect a client using legacy shell
mongosh # connect a client using modern shell
```
we can see the available flags with the `mongod --help` commnad

### Setting "dbpath" & "logpath"

- `--quiet` - reduce verbosity, less output.
- `--logpath <path to file>` - where the logs are stored.
- `--dbpath <folder>` - where the data is actually stored.

we can use this to have different databases and collections, and we can store the logs in a file, for later use.

### Exploring the MongoDB Options

- `--repair` - try and fix corruptions in database
- `--directoryperdb` - group databases into sub folders
- `--storageEngine <engine>` - default is wiredTiger


### MongoDB as a Background Service
- `--fork` - only for mac and linux. run as a background process, as a service. must have a log path (can't log to the terminal).


in windows we can run mongo server as a service, if we checked the option when we installed

```cmd
net start MongoDB
net stop MongoDB
```

to quit 
```sh
use admin
db.shutdownServer()
```
### Using a Config File

rather than run those option manually, we can create a config file with the *cfg* extenesion and store them there.

```yaml
storage:
  dbPath: "path/to/db"
systemLog:
  destination: file
  path: "path/to/log/file.log"
```
we can pass the  path to to the file with `--config` (or `-f`) option.

### Shell Options & Help

### Useful Resources & Links


</details>


## Section 5 : Using MongoDB Compass To Explore Data Visually
<details>
<summary>
MongoDB Compass allows us to work with the database visually
</summary>

we download the mongodb compass tool from the official site, we connect to a host, we must have a mongodb server running.

we launch the program, and we get a broweser GUI to connect to the server, and we can manage data from it, like creating databases, collections, inserting documents, and so on.

we can also do some data exploration like aggregations, see indexes, etc...

</details>

## Section 6: Diving Into Create Options
<details>
<summary>
Closer look at creating and importing documents.
</summary>

ways of creating documents
- `insertOne()` - one document
- `insertMany([])` - multiple documents
- `insert()` - deprecated, both a single and multiple documents
- `mongoimport -d <database> -c <collection> --drop --jsonArray`
- (update methods which can also create documents)

### Understanding `insert()` Methods

```sh
use contactData
db.persons.insertOne({name:"Max",age:30, hobbies:["sports","cooking"]}) #
db.persons.insertMany([
                  {name:"dan",age:28, hobbies:["cars"]},
                  {name:"anna", age:29, hobbies:["cats"]}
                  ]) #objects in an array
db.persons.insert({name:"phil",age:45}) # one document, doesn't return the object id.
db.persons.insert([{name:"josh",age:77},{name:"hans",age:46}]) #many documents, a different output
```

### Working with Ordered Inserts

it's ok to have our own ids,
```sh
db.hobbies.insertMany(
  [
    {name:"Sports",_id::"sports"},
    {name:"Cars",_id::"cars"},
    {name:"Cooking",_id::"cooking"},
    {name:"Cats",_id::"cats"}
  ]
)
```
but if we repeat the command, we might get an error for duplicated keys, but the first elements is still inserted, everything before the error happens. this is the default behavior. **Ordered Insert** means that each insert is handled separately. until the error which stops the execution of the command
```sh
db.hobbies.insertMany(
  [
    {name:"Horses",_id::"horses"}
    {name:"Sports",_id::"sports"},
    {name:"Yoga",_id::"yoga"}
  ]
)

db.hobbies.find({name:"Horses"}) #match
db.hobbies.find({name:"Yoga"}) #not found
```

we might not want this behavior, we can change it by passing a 2nd argument to the `insertMany()` command,we pass a document which specifies that we want all the documents to be inserted, even after an error.

```sh
db.hobbies.insertMany(
  [
    {name:"Horses",_id::"horses"}
    {name:"Sports",_id::"sports"},
    {name:"Yoga",_id::"yoga"} # still inserted
  ],
  {ordered:false}
)


db.hobbies.find({name:"Yoga"}) # found, even if there was an error before
```

### Understanding the **writeConcern**

another option to specify is the *writeConcern*.

when we issue an insert command, we pass it from the client, to the server and eventually the storage engine, which has data in memory and on disk. (also journal)

the writeConcern document determines how the insertion behaves. it can require the data to be written to a number of instances (if we have more than one), and to also be written to the journal. the journal is also a file on disk, which records actions that need to be taken.
The Journal is easier to manage than writing directly to the datbase files, as the journal doesn't have any indexes, positions and other stuff to handle. still slower than keeping the data in memory only. but it lets us recover in case there was a crash.

there is also a write timeout option, *wtimeout*.

the default is to write to one instance, and not write at the journal. setting the number of instances to zero means we don't get an acknowledgment. the write will still happen.
```js
{
  writeConcern: {
    w: 1,
    j: undefined,
    //wtimeout: 200
  }
}
```

and lets look at this in practice
```sh
db.persons.insertOne({name:"chris",age:25},{writeConcern:{w:1,j:false}})
```

### What is Atomicity?

mongoDB guarantees atomic operation for documents, a document is always either inserted/changed as requested, or fails and doesn't change at all. this isn't true for `insertMany` operations.


### Assignemt 2: Time to Practice - Create Operations

> 1. Insert multipleCompanies (comany data of your choice) into a collection - both with `insertOne()` and `insertMany()`.
> 2. Delibratry insert duplicate ID data and "fix" failing additions with unordered inserts.
> 3. Write Data for a new company with both jouraniling being guranteed and not being guaranteed.

```sh
use test
db.companies.insertOne({_id:"mcr",name:"microsoft"})
db.companies.insertMany([{_id:"ggl",name:"google"},{_id:"amzn",name:"amazon"}])

db.companies.insertMany([{_id:"nflx",name:"netflix"},{_id:"amzn",name:"amazon"},{_id:"dpx",name:"dropbox"}],{ordered:true})
db.companies.findOne({name:"netflix"}) #match
db.companies.findOne({name:"dropbox"}) #no match

db.companies.insertMany([{_id:"intl",name:"intel"},{_id:"amzn",name:"amazon"},{_id:"dpx",name:"dropbox"}],{ordered:false})
db.companies.findOne({name:"intel"}) #match
db.companies.findOne({name:"dropbox"}) #match

db.companies.insertOne({_id:"dl",name:"dell"},{writeConcern:{w:1,j:false}}) #no journaling
db.companies.insertOne({_id:"cpx",name:"checkpoint"},{writeConcern:{w:1,j:true}}) #journaling
```

### Importing Data

we can import a data file json into a db with `mongodbimport` command. we have flags for specifying the database, the collection, how the file is structured (one document or an array), and whether to append or replace the collection

```sh
cd resources
mongoimport tv-shows.json -d movieData -c movies --jsonArray --drop
```

### Wrap Up

>- `insertOne()`,`insertMany([])`
>   - You can insert documents with `insertOne()` (one document at a time) or `insertMany()` (multiple documents).
>   - `insert()` also exists, but it's not reccomended to use  it anymore, it also doesn't return the inserted ids.
> - Ordered Inserts
>   - By default, when using `insertMany()`, inserts are ordered - that means that the inserting process stops if an error occurs.
>   - You can change this by switching to "unordered inserts" - your inserting process will then continue, even if errors occurs.
>   - In both cases, no successfull inserts (before the error) will be rolled back.
> - Write Concern
>   - Data should be stored and you can control the "level of guarantee" of that to happen with the *writeConcern* option.
>   - Choose the option value based on your app requirements.
</details>
 
##

[main](README.md)


