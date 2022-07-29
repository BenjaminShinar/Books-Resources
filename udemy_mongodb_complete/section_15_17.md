<!--
// cSpell:ignore ntrwp signup
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
<details>
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

connecting to mongoDB, in the cluster page in atlas, we click <kbd>Connect</kbd>, choose the driver, and copy a connection string value.

now in the app.js and products.js files, we can see the dummy date.

we first connect the app to the database when we start the application. just to make sure we are doing things right

connection_string is a combination of
- mongodb+srv://
- user
- password
- :@cluster.mongodb.net
- /database
- ?retryWrites=true


```js
const mongodb = require('mongodb').MongoClient;


/*all the code*/

mongodb.connect(connection_string)
.then(client => { 
    //function that executes when connection is successful
    console.log('Connected!');
    client.close();
})
.catch(err=>{
    console.log(err);
});
```

### Storing Products in the Database

now we want to store new products in the database.

we take the same code from the app file, and go to the products.js file

```js
const mongodb = require('mongodb'); //top level
const MongoClient = mongodb.MongoClient;

// Add new product
// Requires logged in user
router.post('', (req, res, next) => {
  const newProduct = {
    name: req.body.name,
    description: req.body.description,
    price: parseFloat(req.body.price), // store this as 128bit decimal in MongoDB
    image: req.body.image
  };
  //console.log(newProduct);

  MongoClient.connect(connection_string)
.then(client => { 
    //function that executes when connection is successful
    client.db.collection('products').insertOne(newProduct);
    client.close();
})
.catch(err=>{
    console.log(err);
});

  res.status(201).json({ message: 'Product added', productId: 'DUMMY' });
});

```
### Storing the Price as 128bit Decimal

we need to store data as a decimal, so we look at the documentation.

we import the type from mongo library and use the from string constructor, and we also add `then` and `catch` blocks.

```js
const mongodb = require('mongodb'); //top level
const MongoClient = mongodb.MongoClient;
const Decimal128 = mongodb.Decimal128;

// Add new product
// Requires logged in user
router.post('', (req, res, next) => {
  const newProduct = {
    name: req.body.name,
    description: req.body.description,
    price: Decimal128.fromString(req.body.price.toString()), // store this as 128bit decimal in MongoDB
    image: req.body.image
  };
  //console.log(newProduct);
    MongoClient.connect(connection_string)
.then(client => { 
    //function that executes when connection is successful
    client
    .db()
    .collection('products')
    .insertOne(newProduct)
    .then(result=> 
    {
      console.log(result);
      client.close();
        res.status(201).json({ message: 'Product added', productId: result.insertedId });
    })
    .catch(err=> 
    {
      console.log(err);
      client.close();
      res.status(500).json({ message: "an error occurred"});
    });
})
.catch(err=>{
    console.log(err);
});
});
```

now our code should work, in the app we can enter any product data (with price in decimal) and it should be store in the db.


### Fetching Data From the Database

now we modify the 'get' method. we comment out the existing code and copy in the same code from the insertion method, which we will now modify.

a `find` method returns a cursor, so we need cursor object function, and we need to store the returning values.

```js
// Get list of products products
router.get('/', (req, res, next) => {
    MongoClient.connect(connection_string)
.then(client => { 
    //function that executes when connection is successful
    const products = [];
    client
    .db()
    .collection('products')
    .find()
    .forEach(productDoc => {
      //console.log(productDoc);
      //return productDoc;
      productDoc.price = productDoc.price.toString();
      products.push(productDoc);
    })
    .then(result=> 
    {
      console.log(result);
      client.close();
        res.status(201).json(products);
    })
    .catch(err=> 
    {
      console.log(err);
      client.close();
      res.status(500).json({ message: "an error occurred"});
    });
  })
  .catch(err=>{
      console.log(err);
  })
});
```

we shouldn't store images in the database, only the path. 

### Creating a More Realistic Setup

our code works, but looks really messy and repeats itself, so let's clean it up.

we create a new file "db.s"
```js
const mongodb = require('mongodb'); //top level
const mongoClient = mongodb.MongoClient;

const mongoDbUrl = "connectionString";
let _db;
const initDb = callback=>{
  if (_db) {
    console.log('Database is already initiliazed');
    return callback(null, _db);
  }
  mongoClient.connect(mongoDbUrl)
  .then(client=>{
    _db = client;
    callback(null, _db);
  })
  .catch(err=>{
    callback(err);
  });
};

const getDb = () => {
  if(!_db)
  {
    throw Error('Database not initialized');
  }
  return _db;
}

module.exports= {
  initDb,
  getDb
};
```

in the app.js file, we import the new file module and use it.

```js
db.initDb((err,db)=>{
  if (err) {
    console.log(err);
  }
  else
  {
    app.listen(3100);
  }
});
```

and in the products.js file, we can get the database object without connecting each time. we need to clean up some code, but that's just trial and error.

```js
const db = require('../db');
```

a note: mongoDb provides us a connection pool, a shared connection that can handle multiple requests.

### Getting a Single Product

in the products.js file. we modify the current code like earlier, and we don't forget to modify the price value back to string.
```js

const ObjectId = mongodb.ObjectId;

// Get single product
router.get('/:id', (req, res, next) => {
  db.getDb()
  .db()
  .collection('products')
  .findOne({_id: new ObjectId(req.params.id)})
  .then(productDoc=> {
    productDoc.price = productDoc.price.toString();
    res.status(200).json(productDoc);
  })
  .catch(err=>
  {
    console.log(err);
    res.status(500).json({ message: "an error occurred"});
  });
});
```

### Editing & Deleting Products

as before, we modify the code to allow updating.

we need to make sure we update correctly, not just to add a field
```js
// Edit existing product
// Requires logged in user
router.patch('/:id', (req, res, next) => {
  const updatedProduct = {
    name: req.body.name,
    description: req.body.description,
    price: Decimal12.fromString(req.body.price.toString()), // store this as 128bit decimal in MongoDB
    image: req.body.image
  };
  db.getDb()
  .db()
  .collection('products')
  .updateOne({_id: new ObjectId(req.params.id)},{$set:updatedProduct}) //update, not add field
  .then(updatedProduct=> {
  res.status(200).json({ message: 'Product updated', productId: req.params.id });

  })
  .catch(err=>
  {
    console.log(err);
    res.status(500).json({ message: "an error occurred"});
  });
});
```
deleting is also simple

```js
// Delete a product
// Requires logged in user
router.delete('/:id', (req, res, next) => {

  db.getDb()
  .db()
  .collection('products')
  .deleteOne({_id: new ObjectId(req.params.id)})
  .then(updatedProduct=> {
    res.status(200).json({ message: 'Product deleted' });
  })
  .catch(err=> {
    console.log(err);
    res.status(500).json({ message: "an error occurred"});
  });
});
```

we also need to update the react code code.

```js
import React, { Component } from 'react';
import axios from 'axios';

import './Product.css';

class ProductPage extends Component {
  state = { isLoading: true, product: null };

  componentDidMount() {
    this.fetchData();
  };

  productDeleteHandler = productId => {
    axios
      .delete('http://localhost:3100/products/' + productId)
      .then(result => {
        console.log(result);
        this.fetchData();
      })
      .catch(err => {
        this.props.onError(
          'Deleting the product failed. Please try again later'
        );
        console.log(err);
      });
  };

  fetchData = ()=> {
      axios
      .get('http://localhost:3100/products/' + this.props.match.params.id)
      .then(productResponse => {
        this.setState({ isLoading: false, product: productResponse.data });
      })
      .catch(err => {
        this.setState({ isLoading: false });
        console.log(err);
        this.props.onError('Loading the product failed. Please try again later');
      });
  };
```


### Implementing Pagination

next is adding pagination, we need some more products data to make this viable.

back in the `get` method, we now work on the cursor objects, we skip some pages and limits another.

```js
// Get list of products products
router.get('/', (req, res, next) => {

    const queryPage =req.query.page;
    const querySize = 2;
    const products = [];
    db.getDb()
    .db()
    .collection('products')
    .find()
    .sort({price:-1}).
    .skip((queryPage-1)*querySize)
    .limit(querySize)
    .forEach(productDoc => {
      //console.log(productDoc);
      //return productDoc;
      productDoc.price = productDoc.price.toString();
      products.push(productDoc);
    })
    .then(result=> 
    {
      console.log(result);
        res.status(201).json(products);
    })
    .catch(err=> 
    {
      console.log(err);

      res.status(500).json({ message: "an error occurred"});
    });
  });
```

we also need some react code changes to add the page number.
```js
fetchData = ()=> {
      axios
      .get('http://localhost:3100/products?page=1' + this.props.match.params.id)
      .then(productResponse => {
        this.setState({ isLoading: false, product: productResponse.data });
      })
      .catch(err => {
        this.setState({ isLoading: false });
        console.log(err);
        this.props.onError('Loading the product failed. Please try again later');
      });
  };
```
### Adding an Index

in the earlier section, we sorted on the price fields, if we do this kind of operation often, we should create an index

```js
db.products.createIndex({price:1})
```

### Signing Users Up

we want to add authentication to our application, which is unrelated to the mongoDB authentication.

the users data is stored in the database, but is completely separated.

in the 'auth.js' file, we have *sign-up* and *log-in* routes

we start with the sign-up, we hash the password before pushing it to the database.

we use a token based approach, which isn't covered by this course.
```js
const db = require('../db');

router.post('/signup', (req, res, next) => {
  const email = req.body.email;
  const pw = req.body.password;
  // Hash password before storing it in database => Encryption at Rest
  bcrypt
    .hash(pw, 12)
    .then(hashedPW => {
      // Store hashedPW in database
      console.log(hashedPW);

      db.getDb().db()
      .collection("users")
      .insertOne(
        {email:email,
        password:hashedPW}
      )
      .then(result=>{
        console.log(result);
        const token = createToken();
        res.status(201)
          .json({ token: token, user: { email: email } });
      })
      .catch(err=>{
      console.log(err);
      res.status(500).json({ message: 'Inserting the user failed.' });
      });
    })
    .catch(err => {
      console.log(err);
      res.status(500).json({ message: 'Creating the user failed.' });
    });
  // Add user to database
});
```
with this code changed, we can try creating a user from the front-end app. we use a dummy email and password, and we can see how the password is hashed in the database.

### Adding an Index to Make the Email Unique
the problem is that we can create two users with the same email, so we need to ensure that it's unique.

this is done by adding an index to make the email unique.

```js
db.users.createIndex({email:1},{unique:true})
```
this will prevent us from using the same twice.

### Adding User Sign In

the final part is having the user log-in, we query the database, compare the hashed value, and return a token.\
because this it javascript code, we need to make sure everything is either a promise or an error.

```js
router.post('/login', (req, res, next) => {
  const email = req.body.email;
  const pw = req.body.password;
    
  db.getDb().db().collection('users').findOne({email:email})
  .then(userDoc=>{
    // Check if user login is valid  
    // If yes, create token and return it to client  
    return bcrypt.compare(pw,userDoc.password)
  })
  .then((result)=>  {
    console.log(result); // boolean value
    if (!result) throw Error; // throw error
    const token = createToken();
      res.status(200).json({ token: token, user: { email: email } });
  })
  .catch(err=>{
    res.status(401)
    .json({ message: 'Authentication failed, invalid username or password.' });
  });
});
```

### Wrap Up

we created a demo application that interacts with mongo by using a driver. it is very similar to using the shell, so the main concern is handling the connection to the database, and dividing the responsibilities in a matter which makes sense.

[YouTube Series on Building Restful API](https://academind.com/tutorials/building-a-restful-api-with-nodejs)

</details>

##
[main](README.md)