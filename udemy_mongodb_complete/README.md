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

[configuration files documentation](https://www.mongodb.com/docs/manual/reference/configuration-options/)


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
- `db.myCollection.drop()`
- `db.customers.aggregate([$lookup:{from: "books",localField: "favBooks",foreignField:"_id",as: "favBookData"}])` - merge documents.
- `db.runCommand({colMod:"posts",validator:{}},validationLevel:"warn"})` - update validation schema and validation action
- `use <db>` - switch to a database
	<samp>
	switched to db shop
	</samp>
- `db.products.insertOne()`
- `db.products.find()`
- `.pretty()`

### special type of objects
- point: geospacial data
```
{
	"location": 
	{
		type: "Point",
		coordinates: [56.12,43.09]
	}
}
```

run mongodb as background service in windows
```cmd
net start MongoDB
net stop MongoDB
```

to quit 
```sh
use admin
db.shutdownServer()
```

### CLI flags


when running `mongod` - [documentation](https://www.mongodb.com/docs/manual/reference/program/mongod/)

- `--port` - port to run the service on
- `--quiet` - reduce verbosity, less output.
- `--logpath <path to file>` - where the logs are stored.
- `--dbpath <folder>` - where the data is actually stored.
- `--repair` - try and fix corruptions in database
- `--directoryperdb` - group databases into sub folders
- `--fork` - only for mac and linux. run as a background process, as a service. must have a log path (can't log to the terminal).
- `--storageEngine <engine>` - default is wiredTiger
- `--config` (or `-f`) - pass a mongod configuration file

when running `mongosh` or `mongo` - [documentation](https://www.mongodb.com/docs/manual/reference/program/mongo/)

- `mongo --help` - get help for shell
- `--nodb` - don't run with a database, just a js shell
- `--quiet` - less verbose output
- `--verbose` - more verbose
- `--port` - which port to connect, default is local host 27017
- `--host` - which host to connect
- `-u` - authentication, username
- `-p` - authentication, password


### Mongo Shell Commands
command | action
----|----
`db.help()` | help on db methods
`db.mycoll.help()`| help on collection methods
`db.stats()`| help on collection methods
`sh.help()` | sharding helpers
`rs.help()` | replica set helpers
`db.shutdownServer()` | shut down
`show dbs`| list databases
`show collections`| list collections in current database
`show users`| list users
`show profiles`|list profiles
`help` | general help
`help admin` | administrative help
`help connect` | connecting to a db help
`help keys` | key shortcuts
`help misc` | misc things to know
`help mr` | mapreduce

### Data Types
DataType | Notes | Example
---|---|---
Text | always quotes | "Max"
Boolean | true of false | true
Integer | int32 | 55, `NumberInt(11)`
NumberLong | int64 | 1000000000, `NumberLong(1000000000)`
double | floating point| 12.25
NumberDecimal | High precision | 12.99, `NumberDecimal(11.95)`
ObjectId | automatically generated, has a timestamp internally | ObjectId("text")
ISODate | date | ISODate("2018-09-09")
Timestamp| date time |Timestamp(11421532)
Embedded Documents | nesting | {"a":{}}
Array | list of values| {"b":[]}

</details>

