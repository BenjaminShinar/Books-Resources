<!--
// cSpell:ignore
-->

[main](README.md)

## Section 11 - Understanding The Aggregation Framework
<!-- <details> -->
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

with the aggreation framework, we can sort at any stage.

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

### Working with $project

### Turning the Location Into a geoJSON Object

### Transforming the Birthdate

### Using Shortcuts for Transformations

### Understanding the $isoWeekYear Operator

### group vs $project

### Pushing Elements Into Newly Created Arrays

### Understanding the $unwind Stage

### Eliminating Duplicate Values

### Using Projection with Arrays

### Getting the Length of an Array

### Using the $filter Operator

### Applying Multiple Operations to our Array

### Understanding $bucket

### Diving Into Additional Stages

### How MongoDB Optimizes Your Aggregation Pipelines

### Writing Pipeline Results Into a New Collection

### Working with the $geoNear Stage

### Wrap Up

</details>




##
[main](README.md)