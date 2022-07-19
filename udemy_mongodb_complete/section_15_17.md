<!--
// cSpell:ignore ntrwp
-->

[main](README.md)

## Section 15 - Performance, Fault Tolerancy & Deployment
<details>
<summary>
Entering the Enterprise World.
</summary>

Topics that are the database manager/admin responsability, rather than the developers.

### What Influences Performance?
performance is effected by many factors, some of which are implemented by the developer:
- Efficient Queries/Operations
- Indexes
- Fitting Data Schema

but other factors are covered by the database administrator
- Hardware and Network
- Sharding
- Replica Sets

### Understanding Capped Collections

An explicitly created collection with a set size, and old data is deleted if new data is added once the limit is reached.

```js
use performance
db.createCollection("capped",{capped:true,size:10000, max:3})
db.capped.insertOne({name:"max"})
db.capped.insertOne({name:"anna"})
db.capped.insertOne({name:"dan"})
```

in a capped collection, the retrival order is the insertion order.

```js
db.capped.find().sort({natural:-1}).pretty() // reverse order
```

when we add another document
```js
db.capped.insertOne({name:"maria"})
db.capped.find().pretty()
```

capped collections are good for cases when we need to efficiently query a small amount of data, and we don't care if we lose some of it. this can be for caching, for rolling windows calculations, recent operations, etc...

### What are Replica Sets?

so far, we used a simple flow, from client (the shell), to the mongoDBserver, and to the primary node.

with replica sets, we have secondary nodes, which are filled by the primary Node in asynchronous way. if the primary node fails, a secondary node is promoted. this gives us fault tolerance.

in addition to that, replica sets also allow us to have better read performance, the mongodb server can distribute the requests to different nodes and give us better performance. write operations are still directed to the primary node.

### Understanding Sharding

Sharding - Horizontal Scaling.

sharding is a way to split the data, multiple computers running mongoDB, but with different data.

Data is distributed (not replicated) across shards, queries are directed to all shards.

the flow for sharding uses a new component, called the mongos router. it forewards operations to the correct shards, it uses a 'shard key' (partition key) which is a field in the document, this determines how the data is split across the shards.

when we have query, it might have the shard key field, if it does, then the request is forewarded to the correct shard. if not, then the request is broadcasted, and the router combines the responses before sending them foreward.

if we know we are using sharding, then we, as developers, should make sure that our queries contain the shard key.

### MongoDB Atlas
<details>
<summary>
A managed MongoDB service
</summary>

#### Deploying a MongoDB Server

getting the localhost mongod to a web server, there are a lot of configurations involved

- sharding
- replica sets
- secure user / auth setup
- protecting the web server and network
- regular back ups
- software update
- Encryption (transaction & at rest)

this is a lot, so we can use a managed solution - mongoDB Atlas.

#### Using MongoDB Atlas
we navigate to the website, choose the mongoDB atlas service, we need to sign up (no credit care required), and now we have to create a project.

a cluster is a mongoDB environment, shards, replica sets, nodes, etc...

we click <kbd>Create New Cluster</kbd>, 
- how to cluster is configured globably (have it distributed across the world).
-  underlying cloud provider (aws, gce, azure)
-  the region (some regions aren't available for free tier), we choose 
-  the cluster tier (how powerful is the machine running the server), the **M0** cluster is free tier. 
-  how much storage we have
-  which storage engine version we use (WiredTiger)
-  configure backups as needed (continues backup or snapshot)
-  sharding options (requires a strong machine), ow many shards
-  BI connection
-  encryption at rest
-  additional settings on indexes

once we're ready, we click <kbd>Build a New Cluster</kbd> to deploy it.

in the **security** tab we can configure authentication and users, and set the privileges of the users.

we can also set the ip wite ist (allowlist), we need to allow access from the ip address of the application (or the local application), there are some other security options as well.

#### Backups & Setting Alerts in MongoDB Atlas

we can create new alerts to notify us on some events, like user access, or when some metric exceeds a threshold.

there are options to view the cluster, to migrate it, etc..

#### Connecting to our Cluster

once the cluster is running, we would want to work against it. on the **overview** page,we click <kbd>connect</kbd> to see options on how to connect to it. we can see the ip white list, and we choose which method to use to connet to (we use connection from the shell from now), and the connection string of how to connect with it.

```sh
mongod "mongodb+srv://cluster0-ntrwp.mongodb.net/test" --username max
# enter the password
```
and now we are connected to the live database. and we can work against it just like how we did with the local mongodb.


</details>

### Wrap Up

> Performance & Fault Tolerancy
> - Consider Capped Collections for cases where you want to clear old data automatically.
> - Performance is all about having efficient queries/operations, fitting data formats and a best-practice MongoDB server config.
> - Replica sets provide a fault tolerancy (with automatic recovery) and improved read performance.
> - Sharding allows you to scale your MongoDb server horizonally.
> 
> Deployment & MongoDB Atlas
> - Deployment is a complex matter since it involves many tasks - some of them are not even directly related to MongoDB.
> - Unless you are an experience admin (or you got one), you should consider a managed solution like MongoDB Atlas.
> - Atlas is a managed service where you can configure a mongoDB environment and pay as a by-usage basis.

</details>


## Section 16 - Transactions
<details>
<summary>
Fail Together
</summary>

imagine that we have a replica set, and one of those instances fails to update/delete documents? this means we have an incomplete operations. this is where transactions come into play.

### What are Transactions?

imagine that we have two related collections users and posts, so when we want delete a user, we also want to delete the posts.

but what if deleting the user succeeds, but we fail to delete the posts? with transactions, we force the operations to work together, either they complete successfully, or one fails and we rollback to the previous database state.

### A Typical Usecase

this requires a mongodb server of version 4.0 and above. the video uses the Atlas database (not free tier)

```js
use blog
db.users.insertOne({name:"max"})
//copy the id
db.posts.insertOne({user: id, text:"a"})
db.posts.insertOne({user: id, text:"b"})
```

### How Does a Transaction Work?

a transaction uses a session. we use the session objet as a connector to the collections. we then start a transaction, do our changes, and then commit the transaction (or abort it).

```js
const session = db.getMongo().startSession()

session.startTransaction()
//
const usersC = session.getDatabase("blog").users
const postsC = session.getDatabase("blog").posts
usersCol.deleteOne({_id:id})
db.users.find().pretty() //still visible
postsC.deleteMany({user:id})

session.commitTransaction()
```

if the transaction fails, then the operations are rolled back.

transactions give us atomicity on an operational level, rather than just a document level.

</details>

## Section 17 - From Shell To Driver
<!-- <details> -->
<summary>
Writing Application Code
</summary>

moving from the shell to the application. also a nice project with mongoDB and nodeJs (react).

translating shell comands to driver commands, coneecting to the mongodb Server, doing CRUD operations.

### Splitting Work Between the Driver & the Shell

some work is done by the shell, mostly managing the database. some work is done by the applications with the driver, this is the operational work, the day-to-day queries.

Shell
- configure database
- create collections
- create indexes

Driver
- CRUD operations
- aggregation pipeline


### Preparing our Project

we will use the atlas cluster, we have a some users with "readWriteAnyDatabase" permissions, which will be used by the application.
we need to have our ip white listed so we could connect to it.

we will make a single-page react application, with a restfull API. the code is in the resources section.

we need to have nodeJS installed, and then run `npm install` to get the dependencies.

we can run `npm start` to start the frontend appplication. this should open a web page. we also need to run `npm start:server` to start the node rest api. both should be running at the same time.

for now, the data is dummy data, fetched locally, not from any database.

### Installing the Node.js Driver

installing the nodeJs mongoDB driver.

`npm install mongodb --save`

there is no react driver, we never connect from the client code to the database. this will expose the credentails to the user, which is awful.

### Connecting Node.js & the MongoDB Cluster


### Storing Products in the Database
### Storing the Price as 128bit Decimal
### Fetching Data From the Database
### Creating a More Realistic Setup
### Getting a Single Product
### Editing & Deleting Products
### Implementing Pagination
### Adding an Index
### Signing Users Up
### Adding an Index to Make the Email Unique
### Adding User Sign In
### Wrap Up

</details>

##
[main](README.md)