<!--
// cSpell:ignore Schwarzmüller 
 -->

# Master MongoDB Development for Web & Mobile Apps. CRUD Operations, Indexes, Aggregation Framework - All about MongoDB!

 
udemy course [Master MongoDB Development for Web & Mobile Apps. CRUD Operations, Indexes, Aggregation Framework - All about MongoDB!](https://www.udemy.com/course/mongodb-the-complete-developers-guidey/) by *Maximilian Schwarzmüller*. 



## Takeaways
<details>
<summary>
Things worth remembering
</summary>

default port is 27017

- the `{$set:{}}` is used inside update commands.
- we can't use `.pretty()` after `findOne`.
- matching a value greater than a threshold `db.flightData.find({distance:{$gt:10000}})`
- `update` doesn't care if we forget the `{$set:{}}` part, it will replace the entire document.
- the **_id** field in always included in projections, unless excluded with `{_id:0}`.
- Nested Documents Limits:
  - up 100 levels of nesting.
  - max size of the document is 16MB.
- `db.patients.find({"history":{$elemMatch:{"disease":"cold"}}}).pretty()`
- `db.dropDatabase()`
- `show dbs` - list database
- `use <db>` - switch to a database
	<samp>
	switched to db shop
	</samp>
- `show collections` - list collections in database
- `db.products.insertOne()`
- `db.products.find()`
- `.pretty()`

command | action
----|----
`db.help()` | help on db methods
`db.mycoll.help()`| help on collection methods
`sh.help()` | sharding helpers
`rs.help()` | replica set helpers
`help admin` | administrative help
`help connect` | connecting to a db help
`help keys` | key shortcuts
`help misc` | misc things to know
`help mr` | mapreduce
</details>

