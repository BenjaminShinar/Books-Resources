<!--
// cSpell:ignore
-->

[main](README.md)

## Section 12 - Understanding The Aggregation Framework
<details>
<summary>
Retriving Data Efficiently & In a Structured Way.
</summary>

A more powerful way of finding data, allowing for more complex transformation. there are cases where we can't model the data in the collection in a way that satisfies all of our requirements, for these cases, we can transform the data.

### What is the Aggregation Framework?

the aggergarion framework is an alternative to find methods, it starts with a collection, and then we have a pipeline of operations which result in a different collection of output documents.

each stage of the pipeline feeds the next stage.

### Getting Started with the Aggregation Pipeline

```sh
mongoimport resources/aggregation/persons.json -d analytics -c persons --jsonArray
```

```js
use analytics
show collections
db.persons.findOne()
```

instead, we will use `aggregate`, this method takes an array - the series of steps, every step is a document. aggregation can use indexes to get better performance, and it retruns a cursor object.

the first step is `{$match}`, which is a filtering operation, just like find.
```js
db.persons.aggregate([
    {$match: {gender:"female"}}
]).pretty()
```

### Understanding the Group Stage

we now move to the `{$group}` stagem which groups our data according to a field, we will group all the females in the collection based on their state, and then count them.

the `{$group}` stage always takes a documents as an argument, with *_id* as a property, the value of this field is a document by itself, we give the grouping variable a name, and tell it where to point to- the `$` sign is important to tell the engine to interpet the string as  we should also provide an aggregation operator to the grouping step. in our case we use `$sum`.

```js
db.persons.aggregate([
    {$match: {gender:"female"}},
    {$group: {_id: {state: "$location.state"}, totalPersons: {$sum: 1}}},
]).pretty()
```

### Diving Deeper Into the Group Stage

with the aggregation framework, we can sort at any stage.

```js
db.persons.aggregate([
    {$match: {gender:"female"}},
    {$group: {_id: {state: "$location.state"}, totalPersons: {$sum: 1}}},
    {$sort: {totalPersons:-1}}
]).pretty()
```

### Assignment 7: Time to Practice - The Aggregation Framework

build a pipeline.
1. people older than 50
2. group by gender, count how many per gender, and what the avarage age is.
3. order the output by total persons.

```js
db.persons.aggregate([
    {$match: {"dob.age":{$gt:50}}},
    {$group: {_id: {gender: "$gender"}, totalPersons: {$sum: 1}, avgAge:{$avg:"$dob.age"}}},
    {$sort: {totalPersons:-1}}
]).pretty()
```

### Working with `$project`

The `$project` stage allows us to transform documents into other forms.

we also use the `$concat` operator to paste together fields.  we can also pass operators to each field, like `$toUpper`, `$substrCP` (which takes a string, start postion,number of elements), and then the `$subtract` with `$strLenCP`.
```js
db.persons.aggregate([
    {$project:{_id:0, gender:1,fullName:{$concat:["$name.first"," ","$name.last"]}}}
]).pretty()

db.persons.aggregate([
    {$project:{_id:0, gender:1,fullName:{$concat:[
        {$toUpper:"$name.first"},
        " ",
        {$toUpper:{$substrCP:["$name.last",0,1]}},
        {$substrCP:["$name.last",1,{$subtract:[{$strLenCP:"$name.last"},1]}]}
    ]}}}
]).pretty()

```

### Turning the Location Into a geoJSON Object

we will also tranform the location field into a geoJson object. for this we need to convert the string values into numeric data with `$convert`(input, to, onError, onNull).
```js
db.persons.aggregate([
    {$project:{_id:0,email:1,date:"$dob.date", 
    loc:{type:"Point",coordinates:
    [
        {$convert: {input:"$location.coordinates.longitude",to:"double",onError:0.0,onNull:0.0}},
        {$convert: {input:"$location.coordinates.latitude", to:"double",onError:0.0,onNull:0.0}}
    ]
    }}}
]).pretty()
```

### Transforming the Birthdate

we also want to bring out the birth data fields to another level, we will also like to convert it. we can also use many `$project` steps, if we want to make things easier to read.

```js
db.persons.aggregate([
    {$project:{
        _id:0,
        email:1,
        date:"$dob.date", 
        birthdate: {$convert: {input:"$dob.date", to: "date"}}
        age: "$dob.age"}}
]).pretty()
```

### Using Shortcuts for Transformations

if we want a simple conversion, without specifying the on.error and the on.null values, we could use `$toDate`, `$toLong`, etc...

```js
db.persons.aggregate([
    {$project:{
        _id:0,
        email:1,
        date:"$dob.date", 
        birthdate: {$toDate: "$dob.date"}
        age: "$dob.age"}}
]).pretty()
```

### Understanding the `$isoWeekYear` Operator

after we created all kinds of fields, we can use them as grouping variables, 
```js
db.persons.aggregate([
    {$project:{
        _id:0,
        birthdate: {$toDate: "$dob.date"}}},
        {$group: {_id: {birthYear: {$isoWeekYear: "$birthdate"}}, numPersons: {$sum:1}}},
        {$sort: {numPersons:-1}}
]).pretty()
```

### `$group` vs `$project`

`$group` operations transform many documents into one, based on some criteria, then we create a value based on those documents, such as sum, count, average orn array of values.\
`$project` operations operate on the document itself each documents producing a new document, we use then to include/exclude fields and to transform fields within the document.

### Array Aggregation Stages 
<details>
<summary>
Special Aggregations on Arrays
</summary>

using the "array-data" file. storing it in the 'friends' collection.

#### Pushing Elements Into Newly Created Arrays

we can special things with arrays. we use the `$push` operator to elements into an array. in this example, we push the existing arrays into an array, creating an array of arrays
```js
db.friends.aggregate([
    {$group: {_id: {age: "$age"}, allHobbies: {$push: "$hobbies" }}}
]).pretty()
```

#### Understanding the `$unwind` Stage

if we want to operate on elements of an array, rathen than the array itself, we can use the `$unwind` stage. in the base form, the `$unwind` operator creates one document for each value of the array, it creates many documents out of a single one.

```js
db.friends.aggregate([
    {$unwind: "$hobbies"}
    {$group: {_id: {age: "$age"}, allHobbies: {$push: "$hobbies" }}}
]).pretty()
```

#### Eliminating Duplicate Values

an alternative to `$push` is `addToSet`, which elimates duplications.

```js
db.friends.aggregate([
    {$unwind: "$hobbies"}
    {$group: {_id: {age: "$age"}, allHobbies: {$addToSet: "$hobbies" }}}
]).pretty()
```

#### Using Projections with Arrays

if we want to take a subset of an arrays, we can use the `$slice` operator.

```js
db.friends.aggregate([
    {$project: {_id: 0, examScore: {$slice:["$examScores",1]}}} // first element
]).pretty()
db.friends.aggregate([
    {$project: {_id: 0, examScore: {$slice:["$examScores",-2]}}} //last two element
]).pretty()
db.friends.aggregate([
    {$project: {_id: 0, examScore: {$slice:["$examScores",2,1]}}} // start at the 2nd element, and take one.
]).pretty()
```

#### Getting the Length of an Array

getting the size (length) of an array.

```js
db.friends.aggregate([
    {$project: {_id: 0, numScores: {$size: "$examScores"}}}
]).pretty()
```

#### Using the `$filter` Operator

we can get a subset of an array based on a condition, we use the `$filter` operation, we have a temporary value which we define by `as` to give it a temporary name, and we then use the `$cond` document, in which we use the double dollar syntax (`$$`) to refer to the local element.

```js
db.friends.aggregate([
    {$project: {_id: 0, examScores: {$filter: {input: "$examScores", as: "sc", cond: {$gt: ["$$sc.score",60]}}}}}
]).pretty()
```

#### Applying Multiple Operations to our Array

lets take the highest score for each person!
```js
db.friends.aggregate([
    {$unwind: "$examScores"},
    {$project: {_id:1, name:1, age:1, score:"$examScores.score"}}
    {$sort: {score:-1}},
    {$group: {_id: "$_id",name:{$first:name},maxScore: {$max:"$score"}}},
    {$sort:{maxScore:-1}}
]).pretty()
```

</details>

### Understanding `$bucket`

we can use the `$bucket` operator to create bins of data points based on some criteria. this allows us to get an idea of how the data is distributed
```js
db.persons.aggregate([
{$bucket: {groupBy: "$dob.age", boundaries:[0,18,30,50,80,120], output:
{
    //names: {$push: "$name.first"},
    average: {$abg: "$dob.age"},
    numPersons: {$sum:1}
}}}
]).pretty()
```

an alternative is to tell mongoDB to create the buckets by itself, without us defining the boundaries.
```js
db.persons.aggregate([
{$bucketAuto: {
    groupBy: "$dob.age",
    buckets: 5,
    output: {
    average: {$abg: "$dob.age"},
    numPersons: {$sum:1}}
}}
]).pretty()
```

### Diving Into Additional Stages

`$limit` stage to pull a number of results.
```js
db.persons.aggregate([
    {$project:{
        _id:0,
        name:{$concat: ["$name.first", " ", "$name.last"]},
        birthDate: {$toDate: "$dob.date"}
    }},
    {$sort: {birthDate: -1}},
    {$limit: 10}
]).pretty()
```

if we want the next bunch, we can use `$skip`, but unlike the **find** operator, now the order of the operations matter.
```js
db.persons.aggregate([
    {match: {gender: "male"}},
    {$project:{
        _id:0,
        name:{$concat: ["$name.first", " ", "$name.last"]},
        birthDate: {$toDate: "$dob.date"}
    }},
    {$sort: {birthDate: -1}},
    {$skip: 10},
    {$limit: 10}
]).pretty()
```

we need to be carefull with how we write the stages of the pipeline, this could effect performance.

### Writing Pipeline Results Into a New Collection
if we have a complex pipeline operation (or a pipeline which produces geoData), we can write it into a different collection (where we could query it as any other collection). we add this as as `$out` stage.

(this is important for geoData because it usually requires indices, which we only have in the first stage of the pipeline, but not afterwards)
```js
db.persons.aggregate([
    {$project:{_id:0, name:{$concat: ["$name.first", " ", "$name.last"]},
    loc:{type:"Point",coordinates:
    [
        {$convert: {input:"$location.coordinates.longitude",to:"double",onError:0.0,onNull:0.0}},
        {$convert: {input:"$location.coordinates.latitude", to:"double",onError:0.0,onNull:0.0}}
    ]
    }}},
    {$out: "transformedCollection"}
]).pretty()
```
### Working with the `$geoNear` Stage

with our new collections, lets work on with `$geoNear` pipeline stage. **it has to be the first stage in the pipeline.**

```js
db.transformedCollection.createIndex({loc:"2dsphere"})
db.transformedCollection.aggregate([
    {$geoNear: {
        near: {type: "Point", coordinates:[-18.4,-4.8]},
        maxDistance: 100000.
        num:10, //limit
        query: {age:{$gt:30}}, //other filters,
        distanceField: "distance" // new field name
    }}
]).pretty()
```
### Wrap Up


> Stage And Operators
> - There are plenty of available stages and operators you can choose from.
> - Stages define the different steps your data is funneled through.
> - Each stage recives the output of the last stage as input. 
>   - The first stage takes the original data, so it can use the indexes.
> - Operators can be used inside of stage to tranform, limit or re-calculate date.
> Important Stage
> - The most important stages are `$match`, `$group`, `$project`, `$sort` and `$unwind` - you will work with these a lot.
> - While the are some common behaviors between `find()` filters + projections and `$match` + `$project`, the aggregation stages are generally more flexible.




[documentation](https://www.mongodb.com/docs/manual/core/aggregation-pipeline-optimization/)

</details>



## Section 13 - Working with Numeric Data

<details>
<summary>
Differnet Numeric Types
</summary>

> More Complex Than You Might Think.

### Number Types - An Overview


Integers (int-32bit, long int-64bit) - full numbers.\
Double (double-64bit, high precision double-128bit). the default type for numeric data is double.

in regular double numbers, the decimal value is approximated,while with 128-bit doubles, we have high precision (34 decimal points).

we can use integers if we know the number will never be fractional, and we would like to save some memory.

**note: when using the mongo from the shell, all numbers are doubles, because that's how javascript works.**

### Understanding Programming Language Defaults

when using the mongo shell, numbers are double by default, this also happens with the nodeJs driver. 
```js
let x = 12 // actually 12.0
let y = 12.0 // double 64-bit
```
this depends on the language, python uses integers by default, so the two values won't be the same.

### Working with `int32`

int32 - a 32 bits (four bytes) integer number.

```js
db.people.insertOne({age:14})
db.people.stats()
```
the size of the objects is 35.
```js
db.people.deleteMany({})
db.people.insertOne({age:NumberInt(15)})
//db.people.insertOne({age:NumberInt("15")}) //also works
db.people.stats()
```

now the size is 31, a bit smaller than before.


### Working with `int64`

```js
db.companies.insertOne({valuation: NumberInt("5000000000")})
db.companies.findOne()
db.companies.insertOne({valuation: NumberInt(2147483647)}) // max value
db.companies.insertOne({valuation: NumberInt(2147483648)}) // over flow to minimum number
db.companies.insertOne({valuation: 2147483648}) // double type, can be stored
db.companies.find().pretty()
```

we don't an error, we get a different number, there is an over/under flow of numbers.

we can use `NumberLong` instead
```js
db.companies.insertOne({valuation: NumberLong(2147483648)}) // valid long
db.companies.insertOne({valuation: NumberLong(9223372036854775807)}) // larger than max long
```

we should wrap the number in quation marks, because otherwise the shell won't be able to handle the number.
```js
db.companies.insertOne({valuation: NumberLong("9223372036854775807")}) // larger than max long
```

### Doing Maths with Floats `int32`s & `int64`s
(some warning about not storing numeric data as text, for obvious reasons)

```js
db.accounts.insertOne({num: NumberInt(10)});
db.accounts.updateOne({},{$inc: {num: 10}});
db.accounts.findOne()
```

even though we started with Int32, because we add a double to it, the type changed to double. so if we want to keep the type, we need to make sure we update it with an integer32 value.\

the same happens with long integers

```js
db.companies.deleteMany({})
db.companies.insertOne({value: NumberLong("123456789123456789")})
db.companies.updateOne({},{$inc: {value: 1}})
db.companies.findOne()
```
we add a 1.0 to our long integer number, which converted it to a double and then it was out of the valid ranges for the double64 type.


### What's Wrong with Normal Doubles?

we can use int32 and int64 (`NumberInt` and `NumberLong`) as query operators, just like other numbers.

the normal double is a floating point, so we get some weird results of mathematical operations.
```js
db.science.insertOne({a:0.3, b:0.1})
db.science.find().pretty()
db.science.aggregate([$project:{result: {$subtract: ["$a","$b"]}}])
```
the result will be some weird number, and sometimes it's fine (if we're just displaying the data), but if we are using them for more calculations, our data might drift away and away from the true result.

### Working with Decimal 128bit

`NumberDecimal` is the builder for double128 bit. we should pass the value as a string, to avoid the original problem of having floating point issues. 
```js
db.science.insertOne({a:NumberDecimal("0.3"), b:NumberDecimal("0.1")})
db.science.find().pretty()
db.science.aggregate([$project:{result: {$subtract: ["$a","$b"]}}])
```

now the value is as expected.

however, like before, if we try to modify the data, it will default back into a normal double. so we should be using the NumberDecimal instead.

```js
db.science.updateOne({},{$inc:{a:0.1}}) // imprecision
db.science.updateOne({},{$inc:{a:NumberDecimal("0.1")}}) // correct
db.science.find().pretty()
```

of course, using double64 does take a larger amount of memory.

### Wrap Up

when we hav monetary data, we should be careful with our numbers, there is the old 'scaled approach', which uses integer numbers by scaling up the numbers with a factor. this is like using 100 cents to represent a dollar, and 150 cents instead of 1.5$ dollars.

</details>


## Section 14 - MongoDB & Security

<details>
<summary>
Lock Down Your Data
</summary>

Security should always matter. even if it's usually the role of the database manager rather than the developer.

Security Checklist
- Authentication and Authorization - the database will user-aware.
- Transport Encryption - Data sent between the app and the database is enctyped to avoid someone spoofing the data when it's passed.
- Encryption at Rest - the data inside the database is encrypted, not just plain text files.
- Auditing - track changes and actions
- Server & Network Config and Setup - security of the server/instance holding the database.
- Backups & Software update.

in this module, we focus on authentication/authorization, Transport encryption, and encryption at Rest.


### Understanding Role Based Access Control

* authentication - identify valid users of the database.
* authorization - identify what these users may actually do in the database.

who can connect to the database, and what they can do in the database, which actions? which resources?

users can be actual people, or applications. like a database analyst, or the website that fetches data.

RBAC - Role Based Access Control

each mongoDB server has a special "Admin Database", in addition to whatever collections we use to store the data.

if we have authentication enabled, a user will have to login, and then the allowed operations are determined by privileges: a privilege is a combination of a resource and actions. resources are collections or databases, and actions are verbs that operate on the data. we usually store the privileges in a Role, and then we assign users roles as needed.

granting the minimal needed privileges protects us from malicious actors and from accidental. so it's the favorable approach in the industry.

Roles allow us to seprate between different types of database users, we have administrative roles, a developer role (which is the application actions), and we also have a role for a data scientists or analyst.

### Creating a User

`createUser()` the command to create a user, must have at least one role.
the user document is created on a database, which is authenticates against, it can authenticate against one database, and still have access to other databases. there is also `updateUser()`

we run the mongod server and tell it to require authentication
```sh
sudo mongod --auth #// require authentication
```

`db.auth(username, password)` -  authenticate from inside the shell.\
`mongo -u <username> -p <password>` - authenticate from the commandline, when connecting the client to the server.

when we connect to a database that doesn't have any users, this is a special case, where we can create a special to start our process. this is called **localhost exception**.

```js
use admin
db.createUser({user: "max", pwd:"max", roles:["userAdminAnyDatabase"]})
db.auth('max','max')
```

### Built-In Roles - An Overview

mongoDB ships with some built-in roles, which should cover most use-cases.

- Database User
  - read
  - readWrite
- Database Admin
  - dbAdmin
  - userAdmin
  - dbOwner
- All Database Roles
  - readAnyDatbase
  - readWriteAnyDatbase
  - userAdminAnyDatbase
  - dbAdminAnyDatbase
- Cluster Admin
  - clusterManager
  - clusterMonitor
  - hostManager
  - clusterAdmint
- Backup / Restore
  - backup
  - restore
- Superuser
  - dbOwner (admin)
  - userAdmin (admin)
  - userAdminAnyDatabase
  - root

### Assigning Roles to Users & Databases

we login again to the mongod client, so we need to specify the collection we are authenticating against.

```sh
mongo -u <user> -p <password> --authenticationDatabase admin
```

we can now add users to a specific database. we need to log out of the previous user before we start running commands from the new user.

```js
use shop
db.createUser({user:'dev', pwf:'aaaaa', roles["readWrite"]})
db.logout() 
db.auth('dev','aaaa')
```

because we created the user on the shop database, the roles are scoped by default to the specific database.

### Updating & Extending Roles to Other Databases

we want to add another role to the user, so the user could work on more than one database. we must be logged in as a user with enough permissions to create users.
when we log out of a user, we should do so from the database that the user is attached to.
```js
db.lougout()
use admin
db.auth("max","max")
use shop
db.updateUser('dev',{roles: ["readWrite", {role:"readWrite",db:"blog"}]})
db.getUser("dev")
```
### Assignment 8: Time to Practice - Security

> 1. new mongodb environment - delete everything in mongo (drop all user and database)
> 2. create three users - remember the localthost exception
>    1. Database Admin - works on database, create collection, create indexes.
>    2. User Admin - manage admin
>    3. Developer - Read and write Data in "Customer" and "Sales" Databases.

//clean database

```js
use admin
db.createUser({user:"testAdmin",pwd:"admin", roles:["userAdminAnyDatabase"]})
db.auth('testAdmin','admin')
use dbTest
// db admin
db.createUser({user:"dbTestAdmin",pwd:"testDbAdmin", roles:["dbAdmin"]})
// user admin
db.createUser({user:"userTestAdmin",pwd:"testUserAdmin", roles:["userAdmin"]})
use Sales
db.createUser({user:"devTest",pwd:"test", roles:[{role:"readWrite",db:"Sales"},{role:"readWrite",db:"Customers"}]})
```

> solutuion
>
> starting a mongo server
> ```sh
> mongod --auth
> ```
> adding roles
> ```js
> use admin
> db.createUser({user: 'max', pwd:'max', roles:["userAdminAnyDatabase"]})
> db.auth('max', 'max')
> db.createUser({user:'globalAdmin', pwd:'admin',roles:["dbAdminAnyDatabase"]})
> dv.createUser({user:'dev',pwd:'dev',roles:[{role:"readWrite", db: "customers"},{role:"readWrite",db:"sales"}]})
> 
> db.logout()
> ```
> verifying that we can connect to the newly created users
> ```sh
> mongo -u max -p max --authenticationDatabase admin
> mongo -u globalAdmin -p admin --authenticationDatabase admin
> mongo -u dev -p dev --authenticationDatabase admin
> ```

### Adding SSL Transport Encryption

[creating a self-signed ssl certificate](https://stackoverflow.com/questions/10175812/how-to-generate-a-self-signed-ssl-certificate-using-openssl)

securing the data which is being transferred between mongo and mongoDB. we want to encrypt the data while it's being transmitted.

mongo uses TLS/SSL, making use of a private and public keys.

for linux and mac, we can simply run this command. for windows we can use an executable to install the openssl library for windows, and run the same commands.
```sh

cd /etc/ssl
# create the certificate
openssl req -newkey ras:2048 -new -x509 -days 365 -nodes -out mongodb-cert.crt -keyout mongodb-cert.key

#make a .pem file - linux
cat mongodb-cert.key mongodb-cert.cr > mongodb.pem
#make a .pem file - windows
type mongodb-cert.key mongodb-cert.cr > mongodb.pem
```

we get prompted to fill in some data, most of it doesn't matter, but when we are asked about 
> "Common Name (e.g. server FQDN or YOUR name)[]:"

we should fill out **"localhost"**, or the address of the webserver (in production).

we then crete the .pem file, which we use to encrypt the data at Transit. so we start the mongodb server, we can add an SSL certificate, and also use a certificate authority file.

```sh
mongod --sslMode requireSSL --sslPemKeyFile mongodb.pem
```

when we connect with the mongo client, we also pass the pem file
```sh
mongo --ssl -sslCAFile mongodb.pem --host localhost
```

### Encryption at REST

Encryption at Rest means that the data is encrypted at the database, this can either be encrypting the files themselves, or encrypt specific data in the collections (like user passwords, social security numbers).

mongoDB enterprise version come with built in options to encrypt files.

### Wrap Up

> Users And Roles:
> - MongoDb users a Role Based Access Control  Approach (RBAC).
> - You create users on databases and you then log > in with your credentials (against those databases).
> - Users have no right by default, you need to add roles to allow certain operations.
> - Permission are granted by roles ("Privileges") are only granted for the database the user was added to, unless you explicitly grant access to other databases.
> - You can use "AnyDatabase" roles for cross-databae access.
> 
> Encryption:
> - You can encrypt data during transportation and at rest.
> - During transportation, you use TLS/SSL to encrypt data.
> - For production, you should use SSL certificates issued by a certificate authority (Not self-sigend certificates)
> - For encryption at rest, you can encrypt both the files that hold your data (made simple with "MongoDB Enterprise") and the values inside your documents.

[builtIn Roles](https://www.mongodb.com/docs/manual/reference/built-in-roles/)


</details>


##
[main](README.md)