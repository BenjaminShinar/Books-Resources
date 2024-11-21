<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# AWS Database Day 2024-11-11

learning about different kinds of AWA databases

relational databases, noSQL databases - document, key value, graph.

RDS - mariadb, mysql, microsoft SQL server, postgresSQL, Aurora (mySQL and postgresSQL)

database replication, database resiliency (snapshots, backups, disaster recovery), synchronous and asynchronous replication

authentication - <cloud>IAM</cloud>,

<cloud>DynamoDB</cloud> - NoSQL database in AWS. scalable, requires thinking about data modeling. the main thing is Table, they aren't inside a database. table contains item, each item must have the **primary key** (partition key and optional sort key), but the rest of the attributes can be different for each item.

additional indexes: 
- Local secondary index (LSI) - another sort key using the same partition key. 
- Global secondary index (GSI) - change both partition and sort key.

query - by partition key and sort key (exact or range)

<cloud>ElasticCache</cloud> - in-memory cache database.
- redis
- memcached
- valKey (new)

<cloud>DocumentDB</cloud> - mongo compatible document database.

<cloud>Neptune</cloud> - Graph database, modeling relations, usually networks, (objects are nodes, the relations are the connections), edges and vertices.

## Workshop

catalog.workshops.aws/join, a137-02f73b-61

After logging in to the site for the first time, you will see the home page for the Bookstore. The arrows point on specific features that we've implemented inside of this application that use various different database services:

- A Shopping Cart, driven by a key/value store (Amazon DynamoDB).
- A Best Sellers List, driven by an in-memory sorted set (Amazon Elasticache - Redis).
- Order Processing, driven by relational database (Amazon Aurora PostgreSQL).
- Product Search, driven by a document search index store (Amazon OpenSearch).
- A Product Catalog, driven by a key/value store (Amazon DynamoDB).
- Recommendations, driven by a graph database (Amazon Neptune).
