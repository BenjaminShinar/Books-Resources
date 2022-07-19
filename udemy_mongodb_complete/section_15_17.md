<!--
// cSpell:ignore
-->

[main](README.md)

## Section 15 - Performance, Fault Tolerancy & Deployment
<!-- <details> -->
<summary>
Entering the Enterprise World.
</summary>

Topics that are the database managers' responsability, rather than the developers.

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


### Deploying a MongoDB Server
### Using MongoDB Atlas
### Backups & Setting Alerts in MongoDB Atlas
### Connecting to our Cluster
### Wrap Up


</details>


## Section 16 - Transactions
<!-- <details> -->
<summary>

</summary>
</details>

##
[main](README.md)