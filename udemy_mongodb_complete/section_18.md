<!--
// cSpell:ignore ntrwp signup
-->

[main](README.md)

## Section 18 - Introducing Stitch (MongoDB Realm)
<details>
<summary>
Beyond Data Storage. Serverless platform for building applications.
</summary>

Note: Stitch was renamed to **MongoDB realm**.

Stitch is a different way to work against the database. a set of services that can help us in creating a modern website or application, which helps us take care of the rest API parts without writing boilerplate code.

### What is Stitch?

Stitch is a serverless platform for building applications. it gives us a set of services out of the box, which allows us to focus on the core bussiness logic of our app and takes care of managing the RESTful api parts.

it is built on top the serverless managed mongo database Atlas.

it provides application authentication features out of-the-box. Stitch allows us to connect to the DB from the client code, but it doesn't expose credentials, so we can give fine-grained permissions to each user, without putting the database at risk.

Stitch also allows us to react to database events, and give us the ability to execute code and functionality at the cloud (similar to aws lambda or firebase).

- Stitch QueryAnywhere
- MongoDB Mobile - a local database that can sync with the cloud one
- Stitch Triggers
- Stitch Functions
- Stitch Services - intergrate with other services, such as AWS S3.

Stitch allows us to remove a lot of the backend code that was simply there to connect to the database.

### Preparations

we get to Stitch from the Atlas MongoDB console. we click the <kbd>Stitch Apps</kbd> option (*probably renamed to Realm*). and then <kbd>Create New Application</kbd>.

we give the app a name and choose which mongoDB cluster to connect with.

this section of the course uses a slightly differnet code than before, under the "Stitch_app" folder. inside the folder we run some commands.

```sh
npm install
npm start
```

now the new app will load.

### Start Using Stitch

Once the app is created in atlas, we can start using it. we can manage differnet features of Stitch from the main page.

- Clients - determines which kind of clients we are using, we need the correct SDK.
  - Webpage - which will use javascript from the browser
  - Android app - java
  - IOS app - swift
- Configuration
- Rules
- Triggers
- Services
- Users - application users
- Values
- Functions
- Logs - auditing
- Push notifications - for mobile apps.

From the 'client' option, we take the npm install command and run it

(*note: this has been renamed to mongodb-realm*)
```sh
npm install --save mongodb-stitch-browser-sdk@"^4.0.8" 
```
we restart the app, and from now we can work without the backend code.

### Adding Stitch to our App & Initializing It

we open the "app.js" file, and we first set the "isAuth" field to *true*. so now we run at an authenticated mode.

we also import the `Stitch` object and add a constructor which runs the `initializeDefaultAppClient` with the app client id, which we get from the Stitch console. this is first command which we must run.
```js
//... imports
import {Stitch} from 'mongodb-stitch-browser-sdk';

class App extends Component {
  state = {
    isAuth: true,
    authMode: 'login',
    error: null
  };

  constructor() {
    super();
    Stitch.initializeDefaultAppClient('app-client-id');
  }
  //... more code
}
```

we also import the stitch library and 'RemoteMongoClient' at the "products.js" file. and we use it when we fetch data, it replaces the axios code.\
We create a mongoDB connection by calling `Stitch.getAppClient().defaultAppClient(RemoteMongoClient.factory,'mongodb-atlas')`. we must use the "mongodb-atlas" string. we can then use the same driver functionality as before.\
There are many react code stuff which we need to look out for.

```js
import React, { Component } from 'react';
import axios from 'axios';
import {Stitch, RemoteMongoClient} from 'mongodb-stitch-browser-sdk';

import Products from '../../components/Products/Products';

class ProductsPage extends Component {
  state = { isLoading: true, products: [] };
  componentDidMount() {
    this.fetchData();
  }


  fetchData = () => {
    const mongodb = Stitch.defaultAppClient().getServiceClient(RemoteMongoClient.factory,'mongodb-atlas');
    
    mongodb
    .db('shop').collection('products')
    .find().asArray()
    .then(products=>{
        console.log(products);
        this.setState({products:products,isLoading:false});
    })
    .catch(err=>{
        this.setState({isLoading:false});
        this.props.onError(
            'Fetching the products failed. Please try again later'
        );
        console.log(err);
    })
  };
  //more code
}
```

because we removed the axios code, we might get some errors, so we comment out those parts for now. now our code fails on authentication problems, so we next add that.

### Adding Authentication

in the Stitch web console. we go to the <kbd>Users</kbd> section, under the <kbd>Providers</kbd> tab, here we can control authentication, we first switch on the option to log-in anonymously.

in the "app.js" file, we add an import to get `AnonymousCredintails`, and we use it in the constructor function.

```js
//... imports
import {Stitch,AnonymousCredintails} from 'mongodb-stitch-browser-sdk';

class App extends Component {
  state = {
    isAuth: true,
    authMode: 'login',
    error: null
  };

  constructor() {
    super();
    const client = Stitch.initializeDefaultAppClient('app-client-id');
    client.auth().loginWithCredential(new AnonymousCredintails());
  }
  //... more code
}
```

we still get an error, because our authenticated anonymous user isn't allowed to do anything.

### Sending Data Access Rules

back in the stitch console, under <kbd>Rules</kbd>, we first add a rules collection, we set the database and mongodb collection, we can use a template, or create our own rules under the <kbd>Permissions</kbd> tab. we can set fine grained control for fields and users. so we first set read access to all users.\
we can also set a schema for the data under the <kbd>Schema</kbd> tab.

we hav to fix our code again, because we need to parse the Decimal128 price (and the object ID) back to a string value, which we previously did in the backend.

### CRUD operations
<details>
<summary>
updating the rest API
</summary>

#### Fetching & Converting Data
we need to transform the mongoDB objects into something which we can render in the front-end.


```js
  fetchData = () => {
    const mongodb = Stitch.defaultAppClient().getServiceClient(RemoteMongoClient.factory,'mongodb-atlas');
    
    mongodb
    .db('shop').collection('products')
    .find().asArray()
    .then(products=>{
        console.log(products);
        const transfromedProducts = products.map( product=> {
            product._id = product._id.toString();
            product.price = product.price.toString();
            return product; 
        });
        this.setState({products:transfromedProducts,isLoading:false});
    })
    .catch(err=>{
        this.setState({isLoading:false});
        this.props.onError(
            'Fetching the products failed. Please try again later'
        );
        console.log(err);
    })
  };
```

we don't have the images anymore, because we removed the backend.

#### Deleting Products

we next modify the delete handler to use Stitch. we need to transform the string id into ObjectId. we need the bson package. 

```sh
npm install --save bson
```
we import and use it, and we can now delete products
```js
import BSON from 'bson';

//... more code

  productDeleteHandler = productId => {
    const mongodb = Stitch.defaultAppClient().getServiceClient(RemoteMongoClient.factory,'mongodb-atlas');
    
    mongodb
    .db('shop').collection('products')
    .deleteOne({_id:new BSON.ObjectId(productID)})
    .then(result=>{
        this.fetchData();
    })
    .catch(err=>{
        //this.setState({isLoading:false});
        this.props.onError(
            'Deleting the product failed. Please try again later'
        );
        console.log(err);
    })
  };
```
#### Finding a Single Product

just like deleting, when we fetch a single object, we need to convert the id. in the "product.js" file
we add the imports. (and remove 'axios')

```js
import {Stitch, RemoteMongoClient} from 'mongodb-stitch-browser-sdk';
import BSON from 'bson';


import React, { Component } from 'react';
import axios from 'axios';

import './Product.css';

class ProductPage extends Component {
  state = { isLoading: true, product: null };

  componentDidMount() {
    const mongodb = Stitch.defaultAppClient().getServiceClient(RemoteMongoClient.factory,'mongodb-atlas');
    
    mongodb
    .db('shop').collection('products')
    .find({_id:new BSON.ObjectId(this.props.match.params.id)})
    asArray()
    .then(productResponse=>{
      const product = productResponse[0];
      product._id=product._id.toString();
      product.price=product.price.toString();
        this.setState({isLoading:false, product: product})
    })
    .catch(err=>{
        this.setState({isLoading:false});
        this.props.onError(
            'Loading the product failed. Please try again later'
        );
        console.log(err);
    })
  }
  // more code
}
```
#### Adding Products

now we want to add new products, under the "EditProduct.js" file. we take the same imports, and we play around with the types some more.

```js
class ProductEditPage extends Component {
  state = {
    isLoading: true,
    title: '',
    price: '',
    imageUrl: '',
    description: ''
  };

  //more code
  editProductHandler = event => {

    const mongodb = Stitch.defaultAppClient().getServiceClient(RemoteMongoClient.factory,'mongodb-atlas');
    let requests;
    event.preventDefault();
    if (
      this.state.title.trim() === '' ||
      this.state.price.trim() === '' ||
      this.state.imageUrl.trim() === '' ||
      this.state.description.trim() === ''
    ) {
      return;
    }
    
    this.setState({ isLoading: true });
    const productData = {
      name: this.state.title,
      price: BSON.Decimal128.fromString(this.state.price.toString()),
      image: this.state.imageUrl,
      description: this.state.description
    };
    let request;
    if (this.props.match.params.mode === 'edit') {
      request = axios.patch(
        'http://localhost:3100/products/' + this.props.match.params.id,
        productData
      );
    } else {
      requests = mongodb
      .db('shop').collection('products')
      .insertOne(productData);
    }
    request
      .then(result => {
        this.setState({ isLoading: false });
        this.props.history.replace('/products');
      })
      .catch(err => {
        this.setState({ isLoading: false });
        console.log(err);
        this.props.onError(
          'Editing or adding the product failed. Please try again later'
        );
      });
  };
  // more code
}
```

but we get an error, because we aren't authorized to add products yet!

we need to go back to the stitch console and add 'write' permissions.

#### Updating Products

same as inserting... we need to first get the product from the DB, then we use the id from the component as a mongodb Selector.

</details>

### Authentication
<details>
<summary>
Switch to using real users
</summary>


#### Switching to User Email & Password Authentication

stitch gives us built in user authentication. rather than use acronymous authentication, we use password-email authentication. we just change some field.

#### Adding User Sign Up & Confirmation


we remove the anonymousAuthentication imports, and we import instead

```js
import {Stitch, UserPasswordAuthProviderClient} from 'mongodb-stitch-browser-sdk';
```

in auth handler. we change the code a bit
```js
//... imports
import {Stitch,UserPasswordAuthProviderClient} from 'mongodb-stitch-browser-sdk';

class App extends Component {
  state = {
    isAuth: true,
    authMode: 'login',
    error: null
  };

  constructor() {
    super();
    this.client = Stitch.initializeDefaultAppClient('app-client-id');
    //this.client.auth().loginWithCredential(new AnonymousCredintails()); 
  }
  //... more code
    authHandler = (event, authData) => {
    event.preventDefault();
    if (authData.email.trim() === '' || authData.password.trim() === '') {
      return;
    }
    const emailPassClient = this.client.auth.getProviderClient(UserPasswordAuthProviderClient.factory)
    emailPassClient.registerWithEmail(authData.email, authData.password);
    .then(res=>{
      console.log(res);
    })
    .catch(err=>{
      this.errorHandler("An error occurred");
      console.log(res);
      this.setState({isAuth:false});
    })
    }
}
```

we need to try this with a real email, as Stitch is really creating users.


in the "confirmAccount.js" file, we add the Stitch import
```js
import React, { Component } from 'react';
import {Stitch, UserPasswordAuthProviderClient} from 'mongodb-stitch-browser-sdk';

import './ConfirmAccount.css';

class AuthPage extends Component {
  componentDidMount() {
    const queryUrl = window.location.search;
    const queryParams = new URLSearchParams(queryUrl);
    const token = queryParams.get('token');
    const tokenId = queryParams.get('tokenId');
    console.log('Account confirmed');
    const emailPassClient = Stitch.defaultAppClient.auth.getProviderClient(UserPasswordAuthProviderClient.factory);
    emailPassClient.confirmUser(token, tokenId);
    .then(()=>{
      this.props.history.replace('/')l //redirect
    })
    .catch(err=>{
      console.log (err);
    })
    
  }
  //more code
}
```
#### Adding User Login

now that we have users signing up, we want to log-in the users. Stitch manages the sessions and toekn


in the app.js file

```js

import {Stitch, UserPasswordAuthProviderClient,UserPasswordCredential} from 'mongodb-stitch-browser-sdk';


  authHandler = (event, authData) => {
    event.preventDefault();
    if (authData.email.trim() === '' || authData.password.trim() === '') {
      return;
    }
    const emailPassClient = this.client.auth.getProviderClient(UserPasswordAuthProviderClient.factory);
    
    let request;

    if (this.state.authMode === 'login') {
      const credential = new UserPasswordCredential(auth.data.email, auth.data.password);
      request = this.client.auth.loginWIthCredeintails(credentials);
    } else {
         request= emailPassClient.registerWithEmail(authData.email, authData.password);
    }
    request
      .then(results => {
          console.log(token);
          if (result) {
            this.setState({ isAuth: true });

          }
        })
      .catch(err => {
        this.errorHandler("an error Occurred");
        console.log(err);
        this.setState({ isAuth: false });
      });
  };
```
#### Rules & Real Users


making it so that only logged-in users can interact with products.

in the stitch Console, under the permissions tab, we can modify the rules again, in this example we don't take advantage of everything stitch has to offer.

</details>

### Functions & Triggers

functions and triggers are important features of Stitch. we can write our custom functions in javascript code. they can access databases, write logs to database and anything else.

once we write the function, we can access it from the frontend code via the Stitch client Object. each function call is logged in the backend. the function code isn't visible from the client side code.

we can also execute function based on triggers, which we can set to events in the backend, like database actions. this helps us write backend code, even without having a 'proper' backend service running.

</details>

##
[main](README.md)
