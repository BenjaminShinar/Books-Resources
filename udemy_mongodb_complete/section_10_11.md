<!--
// cSpell:ignore IXSCAN
-->

[main](README.md)

## Section 10 - Working with Indexes
<details>
<summary>
Retriving Data Efficiently.
</summary>

Indexes can speed up our operations (or make them much slower). we will understand what indexes are, which types of indexes exists, and how to use them.

### What Are Indexes & Why Do We Use Them?

indexes can speed up our operations - find, delete and update.

in the default behavior, when there are no indexes, find a matching document requires scanning over all the documents and checking for the field,if the collections is large, the search can take a while.
```js
db.products.find({seller:"Max"})
```

if there is an index, then the index is stored as an ordered list of the values (seller), with each value holding a pointer to the complete document. this allows to search for the matching values in an efficient way (because the index list is ordered), and then to retrieve the documents with a direct access.

however, we shouldn't use too many indexes, because indexes speed up find operations, but they have performance costs on inserts and updates (the ordered list needs to be sorted), and they do take up some space.

### Adding a Single Field Index

we will use the "persons.json" file for this lesson. we first import it.

`mongoimport persons.json -d contactData -c contacts --jsonArray --drop`

and now we take a look. there are 5000 documents, some of the fields are nested documents. we start with running a query to match all contacts older than 60.

```js
show dbs
use contactData
show collections
db.contacts.count()
db.contacts.findOne()
// people older than 60
db.contacts.find({"dob.age":{$gt:60}})
db.contacts.find({"dob.age":{$gt:60}}).count()
```
we can analyze a query by adding a method before the command and see what happend
```js
db.contacts.explain().find({"dob.age":{$gt:60}})
```
the result is the following json file

```json
{
    "queryPlanner" : {
            "plannerVersion" : 1,
            "namespace" : "contactData.contacts",
            "indexFilterSet" : false,
            "parsedQuery" : {
                    "dob.age" : {
                            "$gt" : 60
                    }
            },
            "winningPlan" : {
                    "stage" : "COLLSCAN",
                    "filter" : {
                            "dob.age" : {
                                    "$gt" : 60
                            }
                    },
                    "direction" : "forward"
            },
            "rejectedPlans" : [ ]
    },
    "serverInfo" : {
            "host" : "hostname-l14",
            "port" : 27017,
            "version" : "3.4.24",
            "gitVersion" : "865b4f6a96d0f5425e39a18337105f33e8db504d"
    },
    "ok" : 1
}
```
we don't care much about the server info,but the "queryPlanner.winningPlan" is very interesting. it can compare different plans and determine which one was selected, if we don't have an index, the only option is a full collection scan.

if we want a better explained version, we can pass an argument to the `explain()` function and get more verbose analysis.
```js
db.contacts.explain("executionStats").find({"dob.age":{$gt:60}})
```
we can see how long the execution took, and that we scanned all 5000 documents to match 1222 of them.
```json
"executionStats" : {
    "executionSuccess" : true,
    "nReturned" : 1222,
    "executionTimeMillis" : 2,
    "totalKeysExamined" : 0,
    "totalDocsExamined" : 5000,
    "executionStages" : {
            "stage" : "COLLSCAN",
            "filter" : {
                    "dob.age" : {
                            "$gt" : 60
                    }
            },
            "nReturned" : 1222,
            "executionTimeMillisEstimate" : 0,
            "works" : 5002,
            "advanced" : 1222,
            "needTime" : 3779,
            "needYield" : 0,
            "saveState" : 39,
            "restoreState" : 39,
            "isEOF" : 1,
            "invalidates" : 0,
            "direction" : "forward",
            "docsExamined" : 5000
    }
},
```
so lets add an index and see if things change, we pass the name of the field and the direction (1 for ascending and -1 for descending)

```js
db.contacts.createIndex({"dob.age":1})
db.contacts.explain("executionStats").find({"dob.age":{$gt:60}})
```
now the output is different, and we see that only 1222 documents where examined. now there is an input stage and a fetch stage.

 ### Indexes Behind the Scenes

> What does createIndex() do in detail?
> 
> Whilst we can't really see the index, you can think of the index as a simple list of values + > pointers to the original document.
> 
> Something like this (for the "age" field):
> 
> (29, "address in memory/ collection a1")
> 
> (30, "address in memory/ collection a2")
> 
> (33, "address in memory/ collection a3")
> 
> The documents in the collection would be at the "addresses" a1, a2 and a3. The order does not > have to match the order in the index (and most likely, it indeed won't).
> 
> The important thing is that the index items are ordered (ascending or descending - depending > on how you created the index). createIndex({age: 1}) creates an index with ascending sorting, > createIndex({age: -1}) creates one with descending sorting.
> 
> MongoDB is now able to quickly find a fitting document when you filter for its age as it has > a sorted list. Sorted lists are way quicker to search because you can skip entire ranges (and > don't have to look at every single document).
> 
> Additionally, sorting (via sort(...)) will also be sped up because you already have a sorted > list. Of course this is only true when sorting for the age.

### Understanding Index Restrictions

when we run the indexed query, we had a speed up. but if we run it again, with a different threshold for the age, we see that we examine all the documents, and our execution is slower this time.

```js
db.contacts.explain("executionStats").find({"dob.age":{$gt:5}})
db.contacts.dropIndex({"dob.age":1})
db.contacts.explain("executionStats").find({"dob.age":{$gt:5}})
```
if our query returns all (or most) of the documents, than having an index makes the operation slower. because it creates a extra step over the collection scan (probably an issue of memory locality).


### Creating Compound Indexes

we can create indexes on numbers and on text, but not on boolean values (that wouldn't speed things up that much)
```js
db.contacts.findOne()
db.contacts.explain("executionStats").find({gender:"male"})
db.contacts.createIndex({gender:1})
db.contacts.explain("executionStats").find({gender:"male"})
```

but we can combine indexes together, and create compounded index. the order of the field matters when we create the index, but not when we write the find command
```js
db.contacts.dropIndex({gender:1})
db.contacts.createIndex({"dob.age":1, gender:1})
db.contacts.explain().find({"dob.age":35, gender:"male"})
```

the index speeds up queries that use the start of the index. it works left to right.

- `{"dob.age":35,gender:"male"}` - speed up. match by index
- `{gender:"male","dob.age":35,}` - speedup, order doesn't matter.
- `{"dob.age":35}` - speed up. partial part of the index.
- `{"dob.age":35,gender:"male",nat:"US"}` - speed up, extra fields
- `{"dob.age":35, nat:"US"}` - speed up, partial match of the index, with extra fields
- `{gender:"male"}` - no speed up, the index list is sorted by age first, so it can't scan by gender.

because the index list is sorted over all fields, we can scan parts of the sorted lists, but only if we use the 'head' parts of the index.

### Using Indexes for Sorting

indexes also help us with sorting, because we already have a sorted list, so one part of the work was already done for us!

```js
db.contacts.explain().find({"dob.age":35}).sort({gender:-1})
```

it's also important to note that if we don't sort on an existing index, we can potentially timeout on the request, or run out of memory, the memory threshold is 32MB.

so sometimes we want an index to allow us to sort, even if we don't get a speed up for fetching.

### Understanding the Default Index

we can see all indexes for a collection.
`db.contacts.getIndexes()`

there is always a default index based on the _id field.

### Configuring Indexes

we might want unique index, like the _id field, this guarantees that we have unique field without duplicates.
```js
db.contacts.findOne()
db.contacts.createIndex({email:1},{unique:true})
```

### Understanding Partial Filters

if we know that some values of the index aren't used frequently, we can have a partial index, which ignores those values and is more performant. this means that we accept that if we do need those documents, we will have to perform a collection scan on them.

```js
db.contacts.createIndex({"dob.age":1},{partialFilterExpression:{"dob.age": {$gt:60}}})
db.contacts.createIndex({"dob.age":1},{partialFilterExpression:{gender:"male"}})
db.contacts.explain.find({"dob.age":{$gt:70}}).pretty()
db.contacts.explain.find({"dob.age":{$gt:70}, gender:"male"}).pretty()
```

in a partial filter, the size of the index list is smaller, and we don't need to change it if we add a document with the non-indexed value.

### Applying the Partial Index

we can have a use case of combing a partial index with unique index.

```js
db.users.insertMany([{name:"Max", email: "max@test.com"0},{name:"Manu"}])
db.users.createIndex({email:1},{unique:true})
db.users.insertOne({name:"Anna"}) // duplicate key - null vale is a value in the index
```

if we still want yo allow this behavior. we create a partial index that is created only when the field exists, so we can have multiple null values, but unique existing fields.

```js
db.users.dropIndex({email:1})
db.users.createIndex({email:1},{unique:true, partialFilterExpression:{email:{$exists:true}}})
```

### Understanding the Time-To-Live (TTL) Index

an index for self-destroying data, we decide when data is removed from the collection.

```js
db.sessions.insertOne({data: "asa",createdAt: new Date()})
db.sessions.find().pretty()
db.sessions.createIndex({createdAt:1},{expireAfterSeconds:10})
db.sessions.find().pretty()
db.sessions.insertOne({data: "x",createdAt: new Date()})
//now both will be deleted
db.sessions.find().pretty()
```

this triggers when we have a new insertion, so if we add the index after we created the elements, it won't delete them until we trigger it.

this works only for datetime fields, and is only for single field indexes.

### Query Diagnosis & Query Planning

> - `explain("queryPlanner")` - default. Show summary for executed Query + Winning Plan.
> - `explain("executionStats")` - Show **Detailed** Summary for executed query + Winning Plan + Possibly Rejected Plans.
> - `explain("allPlansExecution")` - Show **Detailed** Summary for executed query + Winning Plan + Winning Plan Decision Process.

things we should look at:
- Milliseconds Process Time (**IXSCAN** typically beats **COLLSCAN**)
- number of keys(in index) examind
- number of documents examined (should be as close as possible to number of indexes, or zero)
- number of documents returned

the case of Zero documents is relevent in a covered Query.

### Understanding Covered Queries

```js
db.customers.insertMany([
        {name:"Max", age:29,salary:3000},
        {name:"Manu", age:30,salary:4000}
])
db.customers.createIndex({name:1})
db.customers.explain("executionStats").find({name:"Max"})
```

the index holds a pointer to the document, but it also holds the key (which is the index), so if we just return the field in the index, te we get a covered query, where no documents were examined.

```js
db.customers.explain("executionStats").find({name:"Max"},{_id:0, name:1})
```

now the *stage* field is **"PROJECTION"**, and this query is very efficient and fast.

### How MongoDB Rejects a Plan

MongoDb can reject plans

```js
db.customers.getIndexes() // id and name
db.customers.CreateIndex({age:1,name:1}) // compound index - order matters
db.customers.explain().find({age:"Max", name:30}) // order doesn't matter
```

the rejected plans now contains the Index Scan on the name alone. mongo can determine which plan to use. it first takes all the indexes which might be useful in the query. it then deides between them by using a subset of the data and running all the plans. then it takes the fastest to run on the entire data set, it then keeps a cache about this query. the cache is cleared after sufficient amount of writes, if the index is rebuilt of if indexes are added or removed, and of course, when the mongoDB server is restarted.

```js
db.customers.explain("allPlansExecution").find({age:"Max", name:30})
```
now we see detailed metrics for the rejected plans.

### Using Multi-Key Indexes

indexes on on array fields,multi key indexes.

```js
db.contacts.dropCollection()
db.contacts.insertOne({name:"Max",hobbies:["Cooking", "Sports",addresses:[{street:"first"},{street:"second"}]]})
db.contacts.createIndex({hobbies:1})
db.contacts.explain("executionStats").find({hobbies:"Sports"})
```
each value of the arrays is an entry in the index list, so if a collection has a total of 100 values in arrays over 30 documents, then the size of the index list will be 100, one for each element.


```js
db.contacts.createIndex({addresses:1})
db.contacts.explain("executionStats").find({"addresses.street":"main street"})
```
this time a collection scan is used, because the indexes field is a nested document, so it indexes over the documents, not parts of it.

```js
db.contacts.explain("executionStats").find({addresses:{street:"main street"}}) // index Scan
```
we can create multikey field on elements of the arrays.
```js
db.contacts.createIndex({"addresses.street":1}) 
db.contacts.explain("executionStats").find({"addresses.street":"main street"}) //index Scan again
```

we can't have compound indexes made of more than one multi-key index.
```js
db.contacts.createIndex({addresses:1,hobbies:1}) // won't work
```

### Text Indexes
<details>
<summary>
Special Behaviors for text Indexes.
</summary>

we can create text indexes, which makes searching text much better performant. a text index is a basically a multi-key array index, where the text was tokensized, standardized and had the stop words removed (is, a, the, for) and the keywords should remain.

```js
db.prodcuts.insertMany([
        {title:"A book", description: "Awesome book about a young artist"},
        {title:"Red Shirt", description: "This t-shirt is read and prettry awesome"}
])

db.products.createIndex({description:1})//normal index
db.products.createIndex({description:"text"})//text index index
db.products.find({$text:{$search:"awesome"}}).pretty()
db.products.find({$text:{$search:"red book"}}).pretty()
```
we can have only one text index per collection, so we don't need to specify which field we are searching on (we can combine them).

when we search using the `$text:{$search:""}` syntax, the case doesn't matter, as the index words are stored in lower case, and if we search for multiple words, then it's treated as if we want matches to any of those words. if we want an exact match, we need to wrap the text in quotation marks, so we have to escape them.\
this is faster than using regular expressions.

```js
db.products.find({$text:{$search:"\"red book\""}}).pretty() // no match
db.products.find({$text:{$search:"\"awesome book\""}}).pretty() // match
```

#### Text Indexes & Sorting

when we search by text index, some results are "better" than others, those who match more keywords in the search phrase are considered more relevant, so we can reflect this.
```js
db.products.find({$text:{$search:"awesome t-shirt"}}).pretty() // match 2
db.products.find({$text:{$search:"awesome t-shirt"}},{score:{$meta:"textScore"}}).pretty() // show index matching score
db.products.find({$text:{$search:"awesome t-shirt"}},{score:{$meta:"textScore"}}).sort({score:{$meta:"textScore"}}).pretty() // explicit sort
```

now we see how much each of the documents scored when matching against the text query, and the results are also sorted by this score!

#### Creating Combined Text Indexes

we can only have a single text index per collection, but we can combine multiple fields if we need.
```js
db.products.createIndex({title:"text"}) // fails because we already have a text index
db.products.dropIndex({description:"text"}) // doesn't work
db.product.getIndexes() // take index name
db.products.dropIndex({description:"descrption_text"}) // this works
db.products.createIndex({title:"text", descrption:"text"}) // compound index
```
#### Using Text Indexes to Exclude Words
with text indexes, we can also exclude words from our search. this is done by prefixing the excluded word with the minus sign.

```js
db.products.find({$text:{$search: "awesome"}}).pretty() // match 2
db.products.find({$text:{$search: "awesome -t-shirt"}}).pretty() // match 1
```
### Setting the Default Language & Using Weights

we can change the default language, which defines how words are tokenized (which stop words are removed). we pass the argument to both the index and the search field. we can also set the query to be case sensitive.
we can also give weights for each field of the index, which determines how the score is calculated, maybe some fields are more important than other.

```js
db.products.getIndexes()
db.products.dropIndex("title_text_descrption_text")
db.products.createIndex({title:"text",descrption:"text"},{default_langague:"german", weights: {title:1, descrption:2}})
db.products.find($text:{$search:"red",$language:"german"})
db.products.find($text:{$search:"Red",$caseSensitive:true})
```

</details>

### Building Indexes

Indexes can be added as the Foreground and the Background. until now, we added the indexes in the foreground. adding in the foreground is faster, but locks up the collection from other queries. adding an index in the background is slower, but doesn't lock the collection.

we have a script that adds one million documents to a collection (this isn't an efficient code)
```sh
mongo resources/credit-rating.js
```

lets do some stuff
```js
show dbs
suse credit
show collections
db.ratings.count()
db.ratings.findOne()
db.ratings.explain("executionStats").find({age:{$gt:80}})
db.ratings.createIndex({age:1})
db.ratings.explain("executionStats").find({age:{$gt:80}})
db.ratings.dropIndex({age:1})
```
we can open a different shell and run a find command while we run the index creation command. the command will wait until the index creation is completed.

```js
//terminal 1
db.ratings.createIndex({age:1})
//terminal 2
db.ratings.find({score:{$gt:3}})
```

in small collections, this doesn't matter, but for big collections with complex indexes, we can't lock the collection while it's being indexed. so a background index might be better

```js
db.ratings.dropIndex({age:1})
db.ratings.createIndex({age:1},{background:true})
```

### Wrap Up

> What And Why?
> - Indexes allow you to retrive data more efficiently (if used correctly) because your queries only have to look at a subset of all documents.
> - You can use single-field, compound, multi-key(array) and text indexes.
> - Indexes don't come for free, they will slow down your writes.
> 
> Queries & Sorting
> - Indexes can be used for both queries and efficient sorting
> - Compound indexes can be used as a whole or in a "left-to-right" (prefix) manner. (e.g. only consider the "name" part of the "name-age" compound index)
> 
> Queries Diagnosis Planning
> - use `explain()` to understand how MongoDB will execute your queries.
>   - "queryPlanner"
>   - "executionStats"
>   - "allPlansExecution"
> - This allows to optimize both your queries and indexes.
> 
> Index Options
> - you can create TTL, unique or partial indexes.
> - for text indexes weights and default language can be assigned.

- [Partial Indexes](https://docs.mongodb.com/manual/core/index-partial/)
- [Supported Languages](https://docs.mongodb.com/manual/reference/text-search-languages/#text-search-languages)
- [Different Langauges in the same Index](https://docs.mongodb.com/manual/tutorial/specify-language-for-text-index/#create-a-text-index-for-a-collection-in-multiple-languages)



</details>

## Section 11 - Working with Geospatial Data
<details>
<summary>
Querying locations
</summary>


mongoDB can use Geospatial Data and query on it. 

### Adding GeoJSON Data

GeoJson is a special format of json.

lets add some places, we can get the coordinates from google maps, we need latitude and longitude. we then insert them as a special document, which has the *type* property, and the *coordinates* as an array([\<longitude>,\<latitude>])


```js
use awesomePlaces
db.places.insertOne({name:"California academy of sciences",locationKey:{type:"point", coordinates:{-122.4724356,37.7672544}}})
db.places.findOne()
```

the supported types are: point, line, polygon

### Running Geo Queries

maybe we want our application to tell us which elements in the collection is close to us, so lets take a location from google maps, we use `$near` and `$geometry`

```js
db.places.find({locationKey:{$near:{$geometry:{type:"Point", coordinates:[-122.471114,37.771104]}}}})
```
but this get us an error, saying that it can't find an index. so we must have a geospatial index in order to use `$near`.

### Adding a Geospatial Index to Track the Distance

like with text indexes, we have a special type of geospatial index, the "2dsphere" type. once we create it, we can use the `$near` document. but we can also define the distance that counts as near. this is a distance in meters: `$minDistance` and `$maxDistance`
```js
db.places.createIndex({locationKey:"2dsphere"})
db.places.find({locationKey:{$near:{$geometry:{type:"Point", coordinates:[-122.471114,37.771104]}}}})
db.places.find({locationKey:{$near:{$geometry:{type:"Point", coordinates:[-122.471114,37.771104]},$maxDistance:40,$minDistance:10}}})
```

### Finding Places Inside a Certain Area

finding points inside a specified data.
lets start by adding more documents and then we find location inside a polygon.
```js
db.places.insertOne({name: "Conservatory of Flowers",locationKey:{type:"Point",coordinates:[-122.4615748,37.7701756]}})
db.places.insertOne({name: "Golden Gate Tennis Court",locationKey:{type:"Point",coordinates:[-122.4593702,37.7705046]}})
db.places.insertOne({name: "Club Nop",locationKey:{type:"Point",coordinates:[-122.4389058,37.7747415]}})
```

we get the polyon coordinates by using a gogole maps, and store them in some javascript variables. we then use the `geoWithin` operator and pass a `$geometry` document with the type **"Polygon"**, and the coordinates as nested array of values, we need to pass the first point in the last position as well, so the polygon will be closed.
```js
const p1 = [-122.45470,37.77473]
const p2 = [-122.45303,37.76641]
const p3 = [-122.51026,37.76411]
const p4 = [-122.51088,37.77131]
db.places.find({locationKey:{$geoWithin:{$geometry:{type: "Polygon",coordinates:[[p1,p2,p3p,p4,p1]]}}}})
```

### Finding Out If a User Is Inside a Specific Area

the opposite case is trying to find if a user is in the vicinity of the area. like the previous query, but other way around. we use the `$geoIntersects` operator.

```js
db.areas.insertOne({name:"Golder Gate park", area:{type:"Polygon", coordinates:[[p1,p2,p3,p4,p1]]}})
db.areas.find()
db.areas.createIndex({area:"2dsphere"}) // must create an index
db.areas.find({area:{$geoIntesects: {$geometry:{type:"Point",coordinates:[-122.49089,37.769992]}}}}) // match
db.areas.find({area:{$geoIntesects: {$geometry:{type:"Point",coordinates:[-122.48446,37.77776]}}}}) // no match
```

### Finding Places Within a Certain Radius

Another geospatial action is finding location in a radius from some point. we use `$geoWithin` again, this time with `$centerSphere` operator, it takes an array of corridnates, and a radius value in *radians* units. the formula is the desired radius divided by the equatorial radius of the earth.


miles:\
$\frac{radius-in-miles}{3963.2}$\
kilometers:\
$\frac{radius-in-km}{6378.1}$

```js
db.places.find(locationKey:{$geoWithin: {$centerSphere:[-122.46203,37.77286],1/6738.1}})
```

the `$near` operator sorts the documents by proximity, while using `$geoWithin` and `$centerSphere` keeps tha original order of the documents.

### Assignment 6: Time to Practice - Geospatial Data

> 1. Pick 3 points in Google Maps and store them in a collection.
> 2. Pick a 4th point and find the nears points within a min a max distance.
> 3. Pick an area and see which points (that are stored in your collection) it contains.
> 4. Store at least one area in a different collection.
> 5. Pick a point a find out which areas in your collection contain that point.

```js
use geoSpatial
//1
db.places.insertMany([
        {name:"Museum of London",loc:{type:"Point",coordinates:[-0.1020399,51.5150566]}},
        {name:"SmithField Market",loc:{type:"Point",coordinates:[-0.1028778,51.5191299]}},
        {name:"Tower of London",loc:{type:"Point",coordinates:[-0.122178,51.5161625]}},
])
db.places.find()
//2
db.places.find({loc:{$near:{$geometry:{type:"Point",coordinates:[-0.1066466,51.5139945]},$minDistance:10, $maxDistance:1000}}}) // should fail without an index
db.places.createIndex({loc:"2dsphere"})
db.places.find({loc:{$near:{$geometry:{type:"Point",coordinates:[]},$minDistance:10, $maxDistance:1000}}}) // should work now
//3
db.places.find({loc:{$geoWithin:{$geometry:{type:"Polygon",coordinates:[[[-0.0945804,51.5187667],[-0.1028778,51.5191299],[-0.102359,51.5127805],[-0.0945804,51.5187667]]]}}}})
//4
db.areas.insertOne({name:"Some Cut of London",area:{type:"Polygon",coordinates:[[[-0.0945804,51.5187667],[-0.1028778,51.5191299],[-0.102359,51.5127805],[-0.0945804,51.5187667]]]}})

db.areas.createIndex({area:"2dsphere"}) // actually not needed
//5
db.areas.find({area:{$geoIntersects: {$geometry:{type:"Point",coordinates:[-0.097715,51.516767]}}}})
```


### Wrap Up

>Storing Geospatial Data:
>- You store geospatial data next to your other data in you documents.
>- Geospatial data has to follow the special GeoJSON format - and respect theytypes supported by MongoDB.
>- Don't forget that the coordinates are [longitude, latitude], not the other way around!
>
> GeoSpatial Indexes:
>- You can add an index to geospatial data: "2dsphere".
>- Some Operations (`$near`) require such an index.
>
> GeoSpatial Queries:
>- `$near`, `$geoWithin` and `$geoIntersects` get you very far.
>- Geospatial queires work with GeoJSON data.


[GeoJson documentation](https://www.mongodb.com/docs/manual/reference/geojson/)

</details>

##
[main](README.md)