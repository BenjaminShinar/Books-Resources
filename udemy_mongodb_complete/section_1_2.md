<!--
// cSpell:ignore
-->

[main](README.md)

## Section 1 - Introduction
<details>
<summary>
First steps, installations, introduction.
</summary>

### What is MongoDB?

mongoDB is a database, the company behind it is also called MongoDB, the source of the name is "Humongous", meaning large and big. mongoDB is a server,which can contain databases, databases contain colletions, and documents contain documents. documents are Schemaless, we aren't required to have the same field in all documents! 

| Mongo      | SQL      |
| ---------- | -------- |
| Database   | Database |
| Collection | Table    |
| Document   | Row      |

a documents is a JSON data format, which means nested objects inside a documents, we can have arrays, primitives and strings. this structure means that we can have complex related data in a single document, rather than having many joins to combine data from different tables.

bson is a binary json document, which is how it's stored in memory. 

### The Key MongoDB Characteristics (and how they differ from SQL Databases)

mongodb is a no-sql solution, the data is not normalized, it doesn't enforce a schema, which in sql results in many tables, this leads to potential data mess, but also allows for flexability. another core feature is that the number of tables is reduced, data is stored together, there are much less joins and merges of data, which makes querying much faster.

the flexability makes developing and changing the documents easier,and also works great for read-write heavy operations.

### Understanding the MongoDB Ecosystem

other products by the MongoDb company
- the mongoDB database - the core product
    - self managed/enterprise edition.
    - community edition
    - Atlas - the Cloud Solution
    - mobile solution
  - compass - GUI for mongodb
  - BI Connectors
  - MongoDB charts
- Stitch - a serverless backend solution
  - serverless query API
  - serverless function
  - database triggers - listen to events on database
  - real-time sync - synchronize stitch with the database

### Installing

<details>
<summary>
Installing whatever we need.
</summary>

in [mongodb](www.mongodb.com) website,we can start grabing what we need.

we will use the community server of mongoDb, the enterprise edition allows some more features and security, but it doesn't change how we work with the database. we need to download a recent stable edition, we will also look at the documentation to read how to install the mongodb Server.

#### Installing MongoDB on Windows

we download the *msi* installer, and we walk through the installer. we choose the custom setup, and select the serve, client, tools and etc...

we make sure that we install MongoDB as a service, rather than running it manually, we might want to change the data and log directories. we can install Compass now, or install it manually later from.

once installed, we open the <kbd>services</kbd> from the <kbd>windows start</kbd> menu, and we will see MongoDB service running.

we can also use the cmd prompt to stop and start the server.

```cmd
net stop MongoDB 
net start MongoDB 
```

the next thing we need is a shell, a client to work against the database, this will be the `mongo.exe` program, which acts as a console to work with it.

there is also a new mongo shell, we can get it from the official website. it's called  **mongosh** (mongo shell).

the final tool is **mongoImport**, which lets us work with data from the course files


</details>

### Time To Get Started!

we can see all our databases with `show dbs`, create (or switch) a new database with `use <db name>`, create a collection by inserting a document `db.products.insertOne({name: "book", price: 12.99})`. the quotation marks are optional for first level fields.

### Shell vs Drivers

eventually we will want to use an application to connect to the database, this is done via a **driver**, the commands are very similar to what we would use in the mongo shell.

### MongoDB + Clients: The Big Picture

we have an application- frontend, backend, and we have a Database. the application (the backend) uses drivers to connet with the MongoDB server, which communicates with a Storage Engine, which eventually stores the data.

we use the mongoDB shell as a playground and a configuration manager.

the storage engine writes and reads data from files, but it also stores data in memory.

</details>

## Section 2 - Understaning The Bascis and CRUD Operations

<details>
<summary>
CRUD - Create, Read, Update, Delete. Basic Mongo db actions and concepts
</summary>

core concepts of crud operations, basics of collections and documents, basic datatypes used in mongodb.

in the real world, we would use a mongo driver to communicate with the mongo server from whatever langauge we are programming with.the syntax might change according to the language, but the commands are generally similar. but for this part of the course, we will use the shell to write commands and queries.



### Understanding Databases, Collections & Documents

the top level is the database server, which contains databases, a database can contain collections, and collections contain documents.

when we start storing data, the databases, collections, and documents are created implicitly.

start a mongo server
```sh
mongod #start mongo server, default port is 27017
mongod --port 27018 # start mongo server with a custom port.
```
and in another terminal
```sh
mongosh --port 27018 #start mongo shell / client at custom port
```

### Creating Databases & Collections

`show dbs` - display databases, we can switch to an existing database with `use db_name`, we can even switch to a db that doesn't exist!\
however, even if we switch to a new database, it won't be listed when we use `show dbs`, it isn't really created until we add data into it. 

we reference the current used database with `db`, and we can chain the commands with the dot `.` symbol.


a database should have collections, so we implicitly create a collection and the document.

a document is always created with curly braces, as it s simply a json - key pair values.

`db.flightData.insertOne()`

### Understanding JSON Data

assume that we have this json document
```json
[
  {
    "departureAirport": "MUC",
    "arrivalAirport": "SFO",
    "aircraft": "Airbus A380",
    "distance": 12000,
    "intercontinental": true
  },
  {
    "departureAirport": "LHR",
    "arrivalAirport": "TXL",
    "aircraft": "Airbus A320",
    "distance": 950,
    "intercontinental": false
  }
]

```
this has two objects inside an array/list. each object is separated with curly braces, and each object is composed of key-value pairs.

so lets insert one documents.
```sh
db.flightData.insertOne({
    "departureAirport": "MUC",
    "arrivalAirport": "SFO",
    "aircraft": "Airbus A380",
    "distance": 12000,
    "intercontinental": true
  })
```
we should see a response that looks like this

<samp>
{
  "acknowledged":true,
  "insertedId": ObjectId("############")
}
</samp>

to view the objects inside the collections, we can list them with the `find` query, without specifying anything. `db.flightData.find()`. we can also make the output look better by appending the `pretty` command at the end. `dn.flightData.find().pretty()`.


### Comparing JSON & BSON

actually, mongoDB uses BSON data, which is a binary form of json. the conversion is done by the mongoDB driver. this both helps with memory and also allows us to use additional types, which is what the *ObjectId("####")* really is. it isn't valid json, but in bson form, it can be parsed.

if we want, we can insert documents manually and we don't have to wrap the key name with quoteation marks, as long as it doesn't have white spaces.

```sh
db.flightData.insertOne({
    departureAirport: "TXL",
    arrivalAirport: "LHR"
  })
```

the id field is autogenerated. we don't have to use it directly, we can also add it ourselves. this must be a unique value.
```sh
db.flightData.insertOne({
    departureAirport: "TXL",
    arrivalAirport: "LHR",
    _id: "txl-lhr-1"
  })
```

if we try to add the ame key, we will get ann error about the duplicated key.



### Create, Read, Update, Delete (CRUD) & MongoDB

we interact with the database using crud operations, we can call them from the shell, the driver or some other way (such as BI connector). we would want to create, read, update and delete documents.

some common queries:
- Create:
  - `insertOne(data, options)`
  - `insertMany(data, options)`
- Read:
  - `find(filter, options)`
  - `findOne(filter, options)`
- Update:
  - `updateOne(filter, data, options)`
  - `updateMany(filter, data, options)`
  - `replaceOne(filter, data, options)`
- Delete:
  - `deleteOne(filter, options)`
  - `deleteMany(filter, options)`


### Finding, Inserting, Deleting & Updating Elements

now lets have some examples. we will find everything, delete by a filter, update a document and delete all. 

a filter is a document as well. 
```sh
db.flightData.find()
db.flightData.find().pretty()
# clear one - based on filter
db.flightData.deleteOne({departureAirpot:"TXL"})
db.flightData.find().pretty()
db.flightData.deleteOne({_id_:"txl-lhr-1"})
db.flightData.find().pretty()
db.flightData.deleteMany() # this fails!
db.flightData.updateOne({distance:12000},{marker:"delete"}) #also Error!
db.flightData.updateOne({distance:12000},{$set:{marker:"delete"}}) #this works
db.flightData.updateMany({},{$set:{marker:"toDelete"}}) #update all
#db.flightData.deleteMany({}) # this will work
db.flightData.deleteMany({marker:"toDelete"})
```
the dollar sign `$` is a special symbol that mongodb knows how to handle.

### Understanding `insertMany()`

before, we inserted documents one by one, but with with `insertMany`, we can add multiple documents at once, we pass them as an array.

```sh
db.flightData.insertMany([
  {
    "departureAirport": "MUC",
    "arrivalAirport": "SFO",
    "aircraft": "Airbus A380",
    "distance": 12000,
    "intercontinental": true
  },
  {
    "departureAirport": "LHR",
    "arrivalAirport": "TXL",
    "aircraft": "Airbus A320",
    "distance": 950,
    "intercontinental": false
  }
])
db.flightData.find().pretty()
```

### Diving Deeper Into Finding Data

so far, we used the `find` without any arguments to list all the data, but we can pass a filter (a document) as a condition to match some documents. this allows us to grab a subset of the data.

`db.flightData.find({intercontinental:true}).pretty()`

lets get more advanced, lets find all documents with a distance larger than some value (using the `$` symbol again).

`db.flightData.find({distance:{$gt:10000}}).pretty()`

### `update` vs `updateMany()`

as before, we used the `$set` command to add a field to a document.
```sh
db.flightData.updateOne({distance:1200},{$set:{delayed:true}}) 
```

if we use "update" and not "updateOne", things will still work. `update` is very similar to `updateMany`.
```sh
db.flightData.update({distance:1200},{$set:{delayed:false}}) 
```
however, if we remove the `{$set:{}}` part. we no longer see the error from before, but the entire document is changed.
```sh
db.flightData.update({distance:1200},{delayed:false}) # no error
db.flightOne.find().pretty()
```

if this is the behavior which we want, we should use `replaceOne()`, the `update` command can both update and replace the document, so it's very dangerous.

### Understanding `find()` & the Cursor Object

now we create another collection, and we add the passengers data.

```sh
db.passengers.insertMany([
  {
    "name": "Max Schwarzmueller",
    "age": 29
  },
  {
    "name": "Manu Lorenz",
    "age": 30
  },
  {
    "name": "Chris Hayton",
    "age": 35
  },
  {
    "name": "Sandeep Kumar",
    "age": 28
  },
  {
    "name": "Maria Jones",
    "age": 30
  },
  {
    "name": "Alexandra Maier",
    "age": 27
  },
  {
    "name": "Dr. Phil Evans",
    "age": 47
  },
  {
    "name": "Sandra Brugge",
    "age": 33
  },
  {
    "name": "Elisabeth Mayr",
    "age": 29
  },
  {
    "name": "Frank Cube",
    "age": 41
  },
  {
    "name": "Karandeep Alun",
    "age": 48
  },
  {
    "name": "Michaela Drayer",
    "age": 39
  },
  {
    "name": "Bernd Hoftstadt",
    "age": 22
  },
  {
    "name": "Scott Tolib",
    "age": 44
  },
  {
    "name": "Freddy Melver",
    "age": 41
  },
  {
    "name": "Alexis Bohed",
    "age": 35
  },
  {
    "name": "Melanie Palace",
    "age": 27
  },
  {
    "name": "Armin Glutch",
    "age": 35
  },
  {
    "name": "Klaus Arber",
    "age": 53
  },
  {
    "name": "Albert Twostone",
    "age": 68
  },
  {
    "name": "Gordon Black",
    "age": 38
  }
])
db.passengers.find().pretty()
```
but actually, we don't see all the the elements, and we see a message

<samp>
Type "it" for more
</samp>

we can use the `it` command to see the rest of the documents. the return value is actually a cursor, an iterator that points to somewhere on the results. it's then used to fetch more data afterwards.

we can force the return of all the results by turning them all into an array.

`db.passengers.find().toArray()`

we can also do something to each document in the collection

`db.passengers.find().forEach( (passangerData) => { printjson(passengerData) } )`

this cursor business is also the reason why `.pretty()` doesn't work with `findOne()`. the function exists for cursor objects, not for a single document. for the other operations (insert, update, delete) there isn't any cursor, as those operations don't fetch data. 

### Understanding Projection

Projections are a way to display a part of each document (only certain fields), and this is done in the mongo engine side, rather than getting all the data and modifying it in the client side. projections allow us to perfrom this on the server side, therefore reducing the load that is retrieved.

to perform a projection on a find query, we pass another document as the second argument, where we specify which fields we wish to display.

`db.passengers.find({},{name:1}).pretty()`

the id is a special field, it is always included, unless explicitly excluded. other fields are implicitly removed if not specified.

`db.passengers.find({},{name:1,_id:0}).pretty()`

### Embedded Documents & Arrays

the value of a field can be a document by itself, not just primitives, we can have nested documents (json objects,but without the id field,of course).

limits:
- up 100 levels of nesting
- max size of the document is 16MB

arrays are values which are in a list form.

lets update the data with a nested document.

`db.flightData.updateMany({},{$set:{status:{description:"on-time",lastUpdated:"1 hour ago"}}})`

lets add an array to one of the passengers. we simply use the square brackets.

`db.passengers.updateOne({name: "Albert Twostone"},{$set:{hobbies:["sports",cookies"]}})`


### Accessing Structured Data

mongoDb knows how to query arrays and find an element within them, if we have a nested search document string we need to wrap the entire path in quotation marks.


```sh
db.passengers.findOne({name:"Albert Twostone"}).hobbies #findOne returned one object
db.passengers.find({hobbies:"sports"}) #mongo db knows how to handle arrays
db.flightData.find({"status.descrption":"on-time"}) #nested drill down
```

### Assignments 1: Time to Practice - The Basics & CRUD Operations

patients data example
```json
{
  "firstName": "Max",
  "lastName": "Schwarzmueller",
  "age":29,
  "history":[
    {"disease":"cold", "treatment": "rest"}
  ]
}
```

tasks:
> 1. insert 3 patient into a new db and new collection, each patients has at least one history entry
> 2. update one patient with new age, name and history entry
> 3. find all patients who are older than some age.
> 4. delete all patients who got a cold as a disease.

solution:

```sh

use quiz1

db.patients.insertMany(
  [
    {
    "firstName": "Dan",
    "lastName": "Green",
    "age":20,
    "history":[
      {"disease":"cold", "treatment": "rest"}
      ]
    },
    {
    "firstName": "Ann",
    "lastName": "Field",
    "age":29,
    "history":[
      {"disease":"flue", "cause": "virus"}
      ]
    },
    {
    "firstName": "John",
    "lastName": "Smith",
    "age":35,
    "history":[
      {"disease":"pox", "cure": "none"}
      ]
    }
  ]
)

db.patients.updateOne({"firstName":"Ann"}, {$set:{"firstName":"Anna","age":32,"history":[{"disease":"danceMania","maxSteps":50},
{"disease":"discoFever","beats":[1,2,3]}
]}})

db.patients.find({"age":{$gt:30}}).pretty()
db.patients.find({"history":{"disease":"cold"}}).pretty()
#db.patients.find({"history":{$elemMatch:{"disease":"cold"}}}).pretty()
db.patients.deleteMany({"history":{$elemMatch:{"disease":"cold"}}})
db.dropDatabase()
```

in the solution he used a different syntax

`db.patients.deleteMany({"history.disease":"cold"})`

</details>



##

[main](README.md)


