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
> - 



[documentation](https://www.mongodb.com/docs/manual/core/aggregation-pipeline-optimization/)

</details>




##
[main](README.md)