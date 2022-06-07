<!--
// cSpell:ignore chkconfig Postgre Sybase Flink Hudi
 -->

[main](README.md)

## Section 5 - Databases

<details>
<summary>
Databases on AWS
</summary>

### Databases 101

we are mostly used to relational databases, like a spreadsheet, tables, rows and columns.

if we liken a database to an excel file, then the file itself is the database, each sheet is a table, and the sheets have rows (observations) and columns (fields), and for each observation/field instersection, we have a value.

there are several relational databases available on aws:

- SQL server, by microsoft.
- Oracle.
- MySQL
- Postgres SQL
- Aurora (amazon)
- Maria DB

we will be using MySQL.

relational databases can have som key features

- Multi-AZ (availability zone) - for distaster recovery
- Read Replicas - for performance

ec2 machine connects to a datbase with a connection string, there is a primary and second database, aws controlls where the dns (connection string) points to. so if the primary database in AZ1 fails, the connection is directed to the secondary datbase in AZ2.

for read replicas, any write is replicated to another database, there isn't an automatic failover if the primary ones fails. but it has a usage when we have many reads and not so many writes, we can direct read-requests into a replicas.

non relational databases:

- collection : basically a table
- document : an observation
- key-value pairs : fields and values.

this is a document example, a json with key value pairs, values can nested inside on another, and each document could contain different fields.

```json
{
  "_id": "5126.....d77",
  "firstname": "John",
  "surname": "Smith",
  "Age": 23,
  "address": [
    {
      "street": "21 Jump Street",
      "suburb": "Richmond"
    }
  ]
}
```

> Data Warehousing is used for business intelligence, tools like Congos, Jaspersoft, SQL Server Reporting Services, Oracle Hyperion, Sap NetWeaver. it is used to pull in very large and complex data sets. Usually used by management to do queries on data (such as current performance vs target).
>
> - OTLP - Online Transaction Processing - "find a specific order"
> - OTAP - Online Analytics Processing - "find and compute some metric"

OTAP is much more intensive, so we would want different architecture for data warehouses doing OTAP queries and databases doing OLTP queries.
Amazon has a data warehouse solution "RedShift".

> ElastiCache is a web service that makes it easy to deploy, operate and scale in-memory cache in the cloud. The service improves the performance of web applications by allowing you to retrive information from fast, managed, in-memory caches, instead of relying entirely on slower disk-based databases.

there are two supported (open-sourced) in-memory caching engines:

1. Memcached
2. Redis

DynamoDB is amazon's noSQL solution.

### RDS

<details>
<summary>
Relational Database.
</summary>

#### Let's Create An RDS Instance

in the console, under services: databases: choose <kbd>RDS</kbd> and provision a database by clicking <kbd>Create Database</kbd>, we choose an engine, in this case MySQL, we choose a version and the Free tier template.

we give the instance a name, create a master user and password, we choose the database instance type and storage size, we select a connectivity and a VPC, we might want to make it publicly available, and choose the port and an avalability zone. under _Additional configurations_, we can set the initial Database name, set backup, enable termination protection and set the modification time windows.

datbases can take a few minutes to launch.

we now launch an EC2 instance, and under _advanced details_ we pass it a bootstrap script to download and install php, php-mysql and installing wordpress as a web server.

```sh
#!/bin/bash
yum install httpd php php-mysql -y
cd /var/www/html
wget https://wordpress.org/latest.tat.gz
tar -xzf latest.tar.gz
cp -r wordpress/* /var/www/html
rm -rf wordpress
rm -rf latest.tar.gz
chmod -R 755 wp-content
chown -R apache:apache wp-content
chkconfig httpd on
service httpd start
```

we give it a storage, tags, a security group and launch it.

we want the webserver to connect to the db instance, so in the new security group which we created for the rds, we click the <kbd>Inbound</kbd> tab, then <kbd>Edit</kbd>, and add a new rule (for the same port) with the source being the other security group.

we wait for both of them to be up, and look at the <kbd>Connectivity</kbd> tab and grab the endpoint value from it. now in the EC2 instance, ww get the public ip and navigate to it. if we did everyhing correctly then we should see a wordpress login page. we fill in the values, and under "database host" we paste the endpoint (instead of "localhost"). it then tells us that it us to manually write a file, so we ssh into the ec2 machine

```sh
ssh ec2-user@ip.ip.ip.ip -i key.pem
$ sudo su
$ cd /var/www/html
$ nano wp-config.php
# paste the contents from
```

again, if we did everything correctly,the wordpress screen should know change and ask for other information.

> - RDS runs on virtual machines. but no ones we can log into.
> - Patching of the RDS operating System and DB is Amazon's responsability
> - RDS is **NOT** Serverless
> - Aurora Serverless is serveless

#### RDS Backups, Multi-AZ & Read Replicas

two different types of backups

- Automated backup
- Database snapshot

> Automated Backups allow you to recover your database to any point in time within a "retention period". the retention period can e between one and 35 days. Automated backups will take a full daily snapshot and will also store transaction logs throughout the day. When you do a recovery, AWS will first choose the mode recent daily backup, and then apply transaction logs relevent to that day. This allows you to do a point in time recovery down to a second, within the retention period.

**Automated backups are enabled by deafult** and are stored in S3, you get an S3 bucket the size of the of the database storage. the backups are taken within a time windows, during this time storage I/O might be suspended and you might have more latency while the data is being backed-up. there is no additional charge for the S3 bucket.

> Database Snapshots are done manually (user initiated), they are stored even after the original RDS is deleted (unlike automated backups).

when we restore an RDS from a backup (automated or manual snapshot), it is restored as a new RDS instance with a new RDS endpoint.

we can have encryption at rest, supported for all six engined (MySQL, Oracle, SQL Server, PostgreSQL,MariaDB and Aurora). encryption is done with the AWS KMS (key management service). if we encyrpt the RDS instance, all the data stroed in it is enctypted, and also the backups, read replicas and snapshots.

when we have **multiple Availability zones**, all updates are mirrored to the standby instance in the other AZ, and if one AZ fails, amazon will automatically redirect the DNS to the other instances. this can also happen in cases of planned database maintenance, we don't need to change the connection string.

however, Multi-AZ is for disaster recovery only, it is not a way to improve performance. Multi-Az is available for most engines (MySQL, Oracle, SQL Server, PostgreSQL and MariaDB), Aurora has a different architecture of fault tolerance.

A **Read Replica** are asynchoursly copies of the primary database, they can help in cases of read-heavy database workloads. a read-replica can be promoted to a primary database. Read replicas are available for most engines (MySQL, Oracle, PostgreSQL and MariaDB).

> - Used for Scaling, not disaster recovery.
> - Must have automatic backups turned on in order to deploy a read-replica.
> - You can have up to 5 read replica copies of any database.
> - You can have read replicas of read replicas (but there is a latency issue).
> - Each read replica will have it's own DNS endpoint.
> - You can have read replica that have multi-AZ,
> - You can create read replicas of Multi-Az source databases.
> - Read replicas can be promoted to be their own databases. this breaks the replication.
> - You can have read Replicas in a second region.

#### RDS Backups, Multi-AZ & Read Replicas - Lab

in the aws web console:

Under the <kbd>RDS</kbd> service, we can perform <kbd>Actions</kbd> like taking snapshots create an Aurora replica. but for now, we click on <kbd>Modify</kbd> and check the "Multi-Az deployment" option. we get a warning that it can cause performance impact, so we get an option to schedule the change to the next maintenance window. now, in the <kbd>Actions</kbd> option, we can choose <kbd>Reboot</kbd> and select "with failover" so the other Zone becomes the primary instance.

To create read Replicas, we must have database backup turned on (retention period). again, modifying the database requires some down time. In the <kbd>Actions</kbd> menu, we click <kbd>Create read replicas</kbd>. we can now choose the destination region, choose if they should be publicly available, enctyped, have multi AZ availability, etc. we do need to give it a unique name.

once it's up, we can <kbd>Actions</kbd> and <kbd>Promote Read Replica</kbd> to make it a primary database.

</details>

### DynamoDB

<details>
<summary>
NoSql database.
</summary>

> Amazon DynamoDB is a fast and flexible NoSQL database service for all applications that need consistent, single-digit millisecond latency at any scale. It is a fully managed database and supports bort document and key-value data models. Its' flexible data model and reliable performance make it a great fit for mobile, web, gaming, ad-tech, IoT and many other applications.

the basics:

- stored on SSD storage (hence the speed)
- spread across 3 geographically distinct data centers (so it highly available and redundance)
- supprots:
  - Eventual consistent reads (deafult)
  - Strongly consistent reads

The One second rule:

> - **Eventual consistent read** - Consistency across all copies of data is usually reached within a second. Repeating a read after a short time should Return the updated data. (Best Read Performance).
> - **Strongly consistent reads** - returns a result that reflects all writes that received a successful response prior to the read.

#### Advanced DynamoDB [SAA-C02]

**DynamoDB Accelerator(DAX)**

> - Fully managed, highly availbe, in memory cache.
> - ten time performace improvement. Reduces Request times from millisecond to microsecond.
> - No need for developers to manage caching logic.
> - Compatible with existing API calls.

in traditional cache solutions, like memcached or Redis, the cache is seperated from the main database. Dax sits between the applications and the database, and it allows for writes and not only reads. it also supports failover.

**Transactions**

> - Multiple "all-or-nothing" operations
> - Financial transactions, fulfilling orders.
> - Under the hood, there are two reads/writes operations: _prepare_ and _commit_.
> - up to 25 items or 4mb.

**On Demand Capcity**

> - Pay-per-request pricing
> - Balance cost and performance
> - No minimum capacity
> - No charge for reads/write - only for storage and backups.
> - used for servies where you don't yet know if there is enough demand to warrent a full database.

**On Demand Backup and Restore**

> - Full backups at any time.
> - Zero impact on table performance or availability.
> - Consistent withing seconds, and retained until deleted.
> - Operates only within the same region as the source table.

**Point-in-Time Recovery**

> - Protects against accidental writes or deletes.
> - Restore to any point in the last 35 days.
> - Incremental backups.
> - **Not enables by deafult**.
> - Latest restorable point is usually five minutes ago.

**Streams**

> - Time-ordered sequence of item-level changes in a table.
> - Stored for 24 hours
> - Insrets, updates and deletes.
> - Stream records are operations on the database, they are stored in something called "Shard".
> - Combine with Lambda functions for functionality like stored procedures.

**Global Table**

> - Managed multi-master, multi-region replication.
> - Globally distributed applications.
> - Based on DynamoDB streams.
> - provides multi region redundancy for Distaster recovery or high availability.
> - No application rewrites needed.
> - Replication latency is under one second.

<kbd>Create Table</kbd>, in the <kbd>Global Tables</kbd> tab, we need to <kbd>Enable Streams</kbd>, and then <kbd>Add region</kbd> (not all regions are supported), and then <kbd>Create Replica</kbd>. now we can see the table being created in the other region. now we create a new item in the table and its immediately replicated.

**Database Migration Service**

Source Database (on-premises, EC2 or RDS): Aurora, S3, DB2, MariaDB, AzureDB, SQL Server, MongoDB, MySQL, Oracle, PostgreSQL, Sybase.

Target Database (on-premises, EC2, RDS,etc): Aurora, DocumentDB, DynamoDB, Kinesis, Redshift, S3, ElasticSearch, Kafka, MariaDb, SQL Server, MongoDb, MySQL, Oracle, PostgreSQL, Sybase.

the source database remains operational.

**Security**

> - fully enctyped at rest using KMS.
> - site-to-site VPN.
> - Direct Connect.
> - IAM policies and roles.
> - Fine-grained access.
> - CloudWatch and CloudTrail.
> - VPC endpoints.

</details>

### Redshift

> Amazon redshift is a fast and powerful, fully managed, petabyte scale data warehouse service in the cloud. customers can start small and scale.

it is used for BI (business intelligence). OLAP - online analytics processing, requires much more compute power than OLTP (online transaction processing).

there are single Node configuration modes (160GB) or multi node. multi nodes have a _leader node_ which manages client connections and recives queries, and up to 128 _Compute Node_ which do the work.

> Advanced Compression:\
> Columnar data stores can be compressed much more than row-based data stores because similar data is stored sequentially on dsk. Amazon Redshift employs multiple compression techniques and can often achieve significant compression relative to traditional relational data stores. in addition, Amazon Redshift doesn't require indexes or materialized viewas and so usesss less sapace than traditional relational database systems. Wen loading data into an empty table, Amazon redshift automatically samples your data and selects the most appropriate compression scheme.
>
> Massively Parallel Processing (MMP):\
> Amazon Redshift automatically distributes data and query load across all nodes, Amazon Redshift makes it easy to add nodes to you data warehouse and enables you to maintain fast query performance as your data warehouse grows.

backups are enable by default, starting from one day to 35. there will always be at least 3 copies of the data (original, replica, S3 backup), and an additional disaster recovery snapshot at a differnet region.

pricing is based on the compute Nodes hours (not for leader node hours), also charge for backups and data transfer inside the vpc.

- encrypted in transit with ssl
- encrypted in res using AES-256
- Amazon Redshift manages keys by itself, but we can use HSM (hardware security module) or AWS-KMS

currently (2020) only avalabile on 1 AZ, but we can have snapshot backup on a different Region.

### Aurora [SAA-C02]

Amazon property database. like Oracle or MySQL server

> Amazon Aurora is a MySQL and PostgreSQL compatible relational database engine that combines the speed and availability of high-end commercial databases with the simplicity and cost-effectiveness of open source databases.

- starts at 10GB, scales to 64TB in 10GB increments (storage autoscaling).
- Compute resources can scale up to 32v CPUs and 244GB memory.
- each availability zone has two copies of the data, with a minimum of 3 AZ, so at least six copies.

aurora is designed to handle the loss of two copies without effecting write avaliability, and handle the loss of three copies without effecting reads. it is also self-healing, so data-blocks and disks are scanned for errors and repaired automatically.

replicas:

- aurora replicas (up to 15)
- mysql read replicas (up to 5)
- postgres read replicas (one)

backups are always enabled, no effect on performance, we can take database snapshot and share them with other AWS accounts.

> Amazon Aurora Serverless is an on-demand, autoscaling configuration for the MySQL and PostgreSQL-compatible editions of Amazon Aurora. An Aurora Serverless DB cluster automatically starts up, shuts down and scales capacity up or down based on your application's needs.

it's a simple, cost-effective, option for infrequent, intermittent or unpredictable workloads. in serverless options, we only pay for operations (invocation), not for the time and storage.

in the AWS console, we want to create an aurora read replica of a mysql database. <kbd>Create Database</kbd>, <kbd>Create Aurora Read Replica</kbd>, the two node (read and write) have different end points, we can now <kbd>Promote Read Replica</kbd> which effectively migrates the original database to the aurora database, we can now take a snapshot and restore it into a database instance as it's own.

> - 2 copies of the data in each AZ, minimal 3 AZ, so minmal 6 copies.
> - we can share aurora snapshots with other AWS accounts.
> - Aurora, MySql and PostgreSQL replicas are availbe.
> - automated failover is only for Aurora replicas.
> - automated backup is turned on by default.
> - Aurora serverless if we want a simple, cost effective option for infrequent, intermittent or unpredictable workloads.

### ElastiCache

> ElasticCache is a web service that makes it easy to deploy, operate and scale an in-memory cache in the cloud. The service improves the performance of web applications bby allowing you to retrieve information from fast, managed, in-memory caches, instead of relying entirely on slower disk-based databases.

supports _Memcached_ and _Redis_ as engines.

| Requirement                         | Memcached | Redis |
| ----------------------------------- | --------- | ----- |
| Simple Cache to offload DB          | yes       | yes   |
| Horizontal Scaling                  | yes       | yes   |
| Multi-threaded Performance          | yes       | no    |
| Advanced data types                 | no        | yes   |
| Ranking/Sorting data sets           | no        | yes   |
| Publisher / Subscriber capabilities | no        | yes   |
| Persistence                         | no        | yes   |
| Multi Avalability Zone              | no        | yes   |
| Backup and Restore capabilities     | no        | yes   |

- Use ElasticCache to increase database and web application performance.
- Redis is AZ and supports backups.

### Database Migration Services (DMS) [SAA-C02]

> AWS Database Migration Service (DMS) is a cloud service that makes it easy to migrate relational databases, data warehouses, NoSQL databases and other types of data stores. you can use AWS DMS to migrate tour data into the AWS cloud, between on-premises instance (through AWS Cloud setups) or between combinations of cloud and on-premises setups.

AWS DMS is a server that runs replication services, it connects a source and target databases through endpoint and migrates the contents. we can use **AWS Schema Conversion Tool (SCT)** to to create some or all of the target tables, indexes, triggers, etc...

types of DMS migrations

- _homogenous migrations_ - from db of the same type (orcale -> orcale)
- _heterogenous migrations_ - source and target of different types (MySql -> aurora). requires a schema conversion.

sources:\
Orcale, Microsoft SQL server, MySQL, MariaDb, PostgreSQL, SAP, MongoDB, Db2, Azure, Amazon RDS (also Aurora), Amazon S3.\
Targets:
Orcale, Microsoft SQL server, MySQL, MariaDb, PostgreSQL, SAP, RDS, Redshift, DynamoDb, S3, **Elasticsearch**, **kinesis** data strems, Document DB.

### Caching Strategies on AWS [SAA-C02]

services with caching capabilities:

- Cloudfront (at edge location)
- API gateways
- ElasticCache (Memcached or Redis)
- DAX (dynamo DB accelerator)

scenarios:

1. Cloud front -> origin
2. Cloud front -> Api Gateway -> Lambda -> (elasticCache, RDS, DynamoDb)

the earlier we cache, the lower the latency, there is always a tradeoff between being up to date and being fast.

### EMR Overview [SAA-C02]

EMR - Elastic Map Reduce

> Amazon EMR (Elastic Map Reduce) is the industry leading cloud big data platform for processing vast amounts of data using open-source tools such as Apache Spark, Apache Hive, Apache HBase, Apache Flink, Apache Hudi and Presto. With EMR, you can run petabyte-scale analysis at less than half the cost of traditional on-premises solutions and over three times faster than standard Apache Spark.

it works with big-data that is inside Aws, it uses an EMR cluster of EC2 machines, where each node has a role, called _node type_. each role runs a diffrent software.

- Master Node - a node that manages the cluster. The master Node **tracks the status of tasks** and monitors the health of the cluster. Every cluster has a master node.
- Core Node - a node with software components that **runs tasks and stores data in the Hadoop Distributed File System (HDFS)** on the cluster. Multi-node cluster have at least one core node.
- Task Node - a node with software components that only runs task and **does not store data in HDFS**. Task Nodes are optional.

| Type   | Amount                        | Role                      |
| ------ | ----------------------------- | ------------------------- |
| Master | one                           | manage the cluster        |
| Core   | at least one (if not manager) | runs task and stores data |
| Task   | zero or more                  | only runs tasks           |

all the logs are stored in the master node "/mnt/var/log/", wo if we want the logs to persist, we configure archive to an S3 bucket (five-minutes intervals), this can be configured only when setting up the cluster.

### Databases Summary

> - RDS (OLTP)
>
>   - SQL
>   - MySQL
>   - PostgreSQL
>   - Oracle
>   - Aurora
>   - MariaDb
>
> - DynamoDB (NoSQL)
> - ReShift (OLAP)
> - ElasticCache
>
>   - Memcached
>   - Redis
>
> - RDS runs on virtual machines. you cannot log in to the operating systems (no ssh).
> - Patching of the RDS is amazon's responsibility.
> - RDS is not Serverless.
> - Aurora Serverless is the exception.
> - Backups can be automated or with snapshots.
> - Read Replicas
>   - Can be multi-AZ.
>   - Used to increase performance.
>   - must have backups turned on.
>   - can be in different regions.
>   - can be MySQL, PostgreSQL, MariaDb, Oracle, Aurora (not SQL server).
>   - can be promoted to master, this will break the read replicas.
> - Multi-Az
>   - Used for Disaster Recovery, not performance.
>   - you can force a failover from one AZ t another by rebooting the RDS instance.
> - Encryption is done with KMS
>   - read Replicas, backups and snapshos will also be encrypted.
> - DynamoDB
>   - Stored on SSD Storage.
>   - Spread Across 3 geographically distinct data centers.
>   - Eventual Consistent Reads (deafult)
>   - Strongly Consistet Reads
> - RedShift
>   - used for BI
>   - backups from 1 to 35 period
>   - mantains at least 3 copies of the data (original,replica, S3 )
>   - we can take a snapshot and Restore it.
> - Aurora
>   - two copies at each AZ, at least 3 Az, so 6 copies total
>   - snapshots can be shared with other AWS accounts
>   - replica types: aurora, MySQl, PostgreSQL.
>   - only aurora replicas can be used as automatic failover.
>   - automated backup by deafult.
>   - Aurora serverless for a simple, const effective solution for infrequent, intermittent or unpredictable workloads.
> - ElasticCache
>   - Use Elasticache to increase database and web application performance.
>   - Redis is multi-AZ.
>   - Redis supports backups and restores.
>   - Memcached supports horizontal scaling.

### Quiz 4: Databases on AWS Quiz

> - RDS Reserved instances are available for multi-AZ deployments. _ANSWER: TRUE_.
> - RDS Reserved instances are available for multi-AZ deployments. _ANSWER: TRUE_.
> - MySQL installations default to port number. _ANSWER: 3306_.
> - If you want your application to check RDS for an error, have it look for an **\_\_** node in the response from the Amazon RDS API. _ANSWER: "error"_
> - What data transfer charge is incurred when replicating data from your primary RDS instance to your secondary RDS instance?. _ANSWER: There is no extra charge._
> - If you are using Amazon RDS Provisioned IOPS storage with a Microsoft SQL Server database engine, what is the maximum size RDS volume you can have by default? _ANSWER: 16TB._
> - In RDS, changes to the backup window take effect **\_\_\_\_**. _ANSWER: Immediately._
> - When you add a rule to an RDS DB security group, you must specify a port number or protocol. _ANSWER: FALSE_

</details>

##

[next](Section_06_07_Advanced_IAM_Route53.md)\
[main](README.md)
