<!--
// cSpell:ignore
-->

[main](README.md)

## Section 10 - Working with Indexes
<!-- <details> -->
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


### Understanding Partial Filters
### Applying the Partial Index
### Understanding the Time-To-Live (TTL) Index
### Query Diagnosis & Query Planning
### Understanding Covered Queries
### How MongoDB Rejects a Plan
### Using Multi-Key Indexes
### Understanding Text Indexes
### Text Indexes & Sorting
### Creating Combined Text Indexes
### Using Text Indexes to Exclude Words
### Setting the Default Language & Using Weights
### Building Indexes
### Wrap Up


</details>

## Section 11 - Working with Geospatial Data
<details>
<summary>

</summary>
</details>

##
[main](README.md)