<!--
// cSpell:ignore DFSR Referer SATA CIFS OCSP
-->

- https://www.udemy.com/course/aws-certified-solutions-architect-associate-practice-tests-k

## Exam 1
<details>
<summary>
65 questions
</summary>


### Q01

> A retail company with many stores and warehouses is implementing IoT sensors to gather monitoring data from devices in each location. The data will be sent to AWS in real time. A solutions architect must provide a solution for ensuring events are received in order for each device and ensure that data is saved for future processing.\
> Which solution would be MOST efficient?
> - Use Amazon Kinesis Data streams for real-time events with a shard for each device, use Amazon Kinesis Data Firehose to save data to Amazon EBS
> - Use Amazon SQS FIFO for real-time events with a shard for each device, use Amazon AWS Lambda for the SQS Queue to save data to Amazon EFS
> - Use Amazon Kinesis Data streams for real-time events with a partition key for each device, use Amazon Kinesis Data Firehose to save data to Amazon S3
> - Use Amazon SQS standard for real-time events with a shard for each device, Trigger Amazon AWS Lambda for the SQS Queue to save data to Amazon S3

**Kinesis** combined with **S3**

> **Explanation:**\
> Amazon Kinesis Data Streams collect and process data in real time. A Kinesis data stream is a set of shards. Each shard has a sequence of data records. Each data record has a sequence number that is assigned by Kinesis Data Streams. A shard is a uniquely identified sequence of data records in a stream.\
> A partition key is used to group data by shard within a stream. Kinesis Data Streams segregates the data records belonging to a stream into multiple shards. It uses the partition key that is associated with each data record to determine which shard a given data record belongs to.\
> For this scenario, the solutions architect can use a partition key for each device. This will ensure the records for that device are grouped by shard and the shard will ensure ordering. Amazon S3 is a valid destination for saving the data records.
> 
> ![Q1 kinesis](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-04-57-65202de89627ab9ac70ef6b89817c981.jpg)

### Q02
> A company runs a web application that serves weather updates. The application runs on a fleet of Amazon EC2 instances in a Multi-AZ Auto scaling group behind an Application Load Balancer (ALB). The instances store data in an Amazon Aurora database. A solutions architect needs to make the application more resilient to sporadic increases in request rates.\
> Which architecture should the solutions architect implement? (Select TWO.)
> - Add Aurora replicas
> - Add AWS WAF in front of the ALB
> - Add an AWS Globabl accelerator endpoint
> - Add an Amazon Cloudfront distribution in front of the ALB
> - Add an AWS Transit Gateway to the availability Zones


~~Maybe **WAF** to protect against DDOS?~~ maybe **CloudFront** for caching?
> **Explanation:**\
> The architecture is already highly resilient but the may be subject to performance degradation if there are sudden increases in request rates. To resolve this situation Amazon Aurora Read Replicas can be used to serve read traffic which offloads requests from the main database. On the frontend an Amazon CloudFront distribution can be placed in front of the ALB and this will cache content for better performance and also offloads requests from the backend.


### Q03

> A company has uploaded some highly critical data to an Amazon S3 bucket. Management are concerned about data availability and require that steps are taken to protect the data from accidental deletion. The data should still be accessible, and a user should be able to delete the data intentionally.\
> Which combination of steps should a solutions architect take to accomplish this? (Select TWO.)
> - Enable MFA Delete on the S3 bucket
> - Create a lifecycle policy for the objects in the S3 bucket
> - Enable default encryption on the S3 bucket
> - Create a bucket policy on the S3 bucket
> - Enable versioning on the S3 bucket


**versioning** will work, ~~maybe **bucket policy**~~

> **Explanation:**\
> Multi-factor authentication (MFA) delete adds an additional step before an object can be deleted from a versioning-enabled bucket.\
> With MFA delete the bucket owner must include the x-amz-mfa request header in requests to permanently delete an object version or change the versioning state of the bucket.
### Q04

> A Solutions Architect has deployed an application on several Amazon EC2 instances across three private subnets. The application must be made accessible to internet-based clients with the least amount of administrative effort.\
> How can the Solutions Architect make the application available on the internet?


~~maybe **NAT gateway**~~

> **Explanation:**\
> To make the application instances accessible on the internet the Solutions Architect needs to place them behind an internet-facing Elastic Load Balancer. The way you add instances in private subnets to a public facing ELB is to add public subnets in the same AZs as the private subnets to the ELB. You can then add the instances and to the ELB and they will become targets for load balancing.\
> An example of this architecture is shown below:
> 
> ![Q4 private and public subnets](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-09-08-4d552faa7e6a02f1057665490b64d36b.jpg)


### Q05
> Storage capacity has become an issue for a company that runs application servers on-premises. The servers are connected to a combination of block storage and NFS storage solutions. The company requires a solution that supports local caching without re-architecting its existing applications.\
> Which combination of changes can the company make to meet these requirements? (Select TWO.)

probably **storage gateway**

> **Explanation:**\
> In this scenario the company should use cloud storage to replace the existing storage solutions that are running out of capacity. The on-premises servers mount the existing storage using block protocols (iSCSI) and file protocols (NFS). As there is a requirement to avoid re-architecting existing applications these protocols must be used in the revised solution.\
> The AWS Storage Gateway volume gateway should be used to replace the block-based storage systems as it is mounted over iSCSI and the file gateway should be used to replace the NFS file systems as it uses NFS.

### Q06
> A recent security audit uncovered some poor deployment and configuration practices within your VPC. You need to ensure that applications are deployed in secure configurations.\
> How can this be achieved in the most operationally efficient manner?
> - Remove the ability for staff to deploy applications
> - use cloud formation with securely configured templates
> - manually check all application configurations before deployment
> - use aws inspector to apply secure configuration

probably **cloud formation**

> **Explanation:**\
> CloudFormation helps users to deploy resources in a consistent and orderly way. By ensuring the CloudFormation templates are created and administered with the right security configurations for your resources, you can then repeatedly deploy resources with secure settings and reduce the risk of human error.

### Q07
> A company has deployed a new website on Amazon EC2 instances behind an Application Load Balancer (ALB). Amazon Route 53 is used for the DNS service. The company has asked a Solutions Architect to create a backup website with support contact details that users will be directed to automatically if the primary website is down.\
> How should the Solutions Architect deploy this solution cost-effectively?
> - Create the backup website on EC2 and ALB in another Region and create an AWS accelerator point
> - Deploy the backup website on EC2 and ALB in another Region and use Route53 health checks for failover routing
> - Configure a static website using amazon S# and create a route53 weighted routing policy
> - Configure a static website using amazon S# and create a route53 failover routing policy

~~**Another region with route53 health checks**~~

> **Explanation:**\
> The most cost-effective solution is to create a static website using an Amazon S3 bucket and then use a failover routing policy in Amazon Route 53. With a failover routing policy users will be directed to the main website as long as it is responding to health checks successfully.\
> If the main website fails to respond to health checks (its down), Route 53 will begin to direct users to the backup website running on the Amazon S3 bucket. It’s important to set the TTL on the Route 53 records appropriately to ensure that users resolve the failover address within a short time.
> 
### Q08

> A website runs on Amazon EC2 instances in an Auto Scaling group behind an Application Load Balancer (ALB) which serves as an origin for an Amazon CloudFront distribution. An AWS WAF is being used to protect against SQL injection attacks. A review of security logs revealed an external malicious IP that needs to be blocked from accessing the website.\
> What should a solutions architect do to protect the application?


modify the **WAF** to protect against the Ip, cloudfront means that the ip is changed.

> **Explanation:**\
> A new version of the AWS Web Application Firewall was released in November 2019. With AWS WAF classic you create “IP match conditions”, whereas with AWS WAF (new version) you create “IP set match statements”. Look out for wording on the exam.\
> The IP match condition / IP set match statement inspects the IP address of a web request's origin against a set of IP addresses and address ranges. Use this to allow or block web requests based on the IP addresses that the requests originate from.\
> AWS WAF supports all IPv4 and IPv6 address ranges. An IP set can hold up to 10,000 IP addresses or IP address ranges to check.

### Q09
> A persistent database must be migrated from an on-premises server to an Amazon EC2 instances. The database requires 64,000 IOPS and, if possible, should be stored on a single Amazon EBS volume.\
> Which solution should a Solutions Architect recommend?

~~maybe **i3?**~~

> **Explanation:**\
> Amazon EC2 Nitro-based systems are not required for this solution but do offer advantages in performance that will help to maximize the usage of the EBS volume. For the data storage volume an i01 volume can support up to 64,000 IOPS so a single volume with sufficient capacity (50 IOPS per GiB) can be deliver the requirements.\
> The current list of EBS volume types is in the table below:
> 
> ![Q9 EBS Volumes](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-17-09-93913bab7527ec0862c9e21624ba7869.jpg)

### Q10
> A company offers an online product brochure that is delivered from a static website running on Amazon S3. The company’s customers are mainly in the United States, Canada, and Europe. The company is looking to cost-effectively reduce the latency for users in these regions.\
> What is the most cost-effective solution to these requirements?


~~**Cloudfront with some origin points** (not all)~~

> **Explanation:**\
> With Amazon CloudFront you can set the price class to determine where in the world the content will be cached. One of the price classes is “U.S, Canada and Europe” and this is where the company’s users are located. Choosing this price class will result in lower costs and better performance for the company’s users.

### Q11
> A company provides a REST-based interface to an application that allows a partner company to send data in near-real time. The application then processes the data that is received and stores it for later analysis. The application runs on Amazon EC2 instances.\
> The partner company has received many 503 Service Unavailable Errors when sending data to the application and the compute capacity reaches its limits and is unable to process requests when spikes in data volume occur.\
> Which design should a Solutions Architect implement to improve scalability?

**Kinesis + Lambda**

> **Explanation:**\
> Amazon Kinesis enables you to ingest, buffer, and process streaming data in real-time. Kinesis can handle any amount of streaming data and process data from hundreds of thousands of sources with very low latencies. This is an ideal solution for data ingestion.
> To ensure the compute layer can scale to process increasing workloads, the EC2 instances should be replaced by AWS Lambda functions. Lambda can scale seamlessly by running multiple executions in parallel.


### Q12

> A company runs an application on six web application servers in an Amazon EC2 Auto Scaling group in a single Availability Zone. The application is fronted by an Application Load Balancer (ALB). A Solutions Architect needs to modify the infrastructure to be highly available without making any modifications to the application.\
> Which architecture should the Solutions Architect choose to enable high availability?

**spread across Availability zones**

> **Explanation:**\
> The only thing that needs to be changed in this scenario to enable HA is to split the instances across multiple Availability Zones. The architecture already uses Auto Scaling and Elastic Load Balancing so there is plenty of resilience to failure. Once the instances are running across multiple AZs there will be AZ-level fault tolerance as well.

### Q13
> A new application is to be published in multiple regions around the world. The Architect needs to ensure only 2 IP addresses need to be whitelisted. The solution should intelligently route traffic for lowest latency and provide fast regional failover.\
> How can this be achieved?


~~**route53**?~~

> **Explanation:**\
> AWS Global Accelerator uses the vast, congestion-free AWS global network to route TCP and UDP traffic to a healthy application endpoint in the closest AWS Region to the user.
> This means it will intelligently route traffic to the closest point of presence (reducing latency). Seamless failover is ensured as AWS Global Accelerator uses anycast IP address which means the IP does not change when failing over between regions so there are no issues with client caches having incorrect entries that need to expire.
> 
> ![Q13 Globabl Accelerator](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-46-55-314c6149f4e1a3552921ab6fa213d8d8.jpg)


### Q14
> A solutions architect is designing the infrastructure to run an application on Amazon EC2 instances. The application requires high availability and must dynamically scale based on demand to be cost efficient.\
> What should the solutions architect do to meet these requirements?

**application load balancer, multiple AZ**

> **Explanation:**\
> The Amazon EC2-based application must be highly available and elastically scalable. Auto Scaling can provide the elasticity by dynamically launching and terminating instances based on demand. This can take place across availability zones for high availability.\
> Incoming connections can be distributed to the instances by using an Application Load Balancer (ALB).


### Q15
> A solutions architect is creating a document submission application for a school. The application will use an Amazon S3 bucket for storage. The solution must prevent accidental deletion of the documents and ensure that all versions of the documents are available. Users must be able to upload and modify the documents.\
> Which combination of actions should be taken to meet these requirements? (Select TWO.)
> - Enable MFA Delete on the bucket
> - Enable Versioning on the bucket
> - Set read only permissions on the bucket
> - Attach an IAM policy to the bucket
> - Encrypt the bucket using AWS SSE-S3

**Versioning**, maybe **MFA delete** 

> **Explanation:**\
> None of the options present a good solution for specifying permissions required to write and modify objects so that requirement needs to be taken care of separately. The other requirements are to prevent accidental deletion and the ensure that all versions of the document are available.
> The two solutions for these requirements are versioning and MFA delete. Versioning will retain a copy of each version of the document and multi-factor authentication delete (MFA delete) will prevent any accidental deletion as you need to supply a second factor when attempting a delete.


### Q16

> A company runs an application that uses an Amazon RDS PostgreSQL database. The database is currently not encrypted. A Solutions Architect has been instructed that due to new compliance requirements all existing and new data in the database must be encrypted. The database experiences high volumes of changes and no data can be lost.\
> How can the Solutions Architect enable encryption for the database without incurring any data loss?

~~**multi AZ mode with secondary and primary?**~~

> **Explanation:**\
> You cannot change the encryption status of an existing RDS DB instance. Encryption must be specified when creating the RDS DB instance. The best way to encrypt an existing database is to take a snapshot, encrypt a copy of the snapshot and restore the snapshot to a new RDS DB instance. This results in an encrypted database that is a new instance. Applications must be updated to use the new RDS DB endpoint.
> In this scenario as there is a high rate of change, the databases will be out of sync by the time the new copy is created and is functional. The best way to capture the changes between the source (unencrypted) and destination (encrypted) DB is to use AWS Database Migration Service (DMS) to synchronize the data.
> 
> ![Q16 move to encrypted EBS](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-43-00-2b10c6ed31c5f439b4d82b7c3dbf77e5.jpg)


### Q17
> An eCommerce application consists of three tiers. The web tier includes EC2 instances behind an Application Load balancer, the middle tier uses EC2 instances and an Amazon SQS queue to process orders, and the database tier consists of an Auto Scaling DynamoDB table. During busy periods customers have complained about delays in the processing of orders. A Solutions Architect has been tasked with reducing processing times.\
> Which action will be MOST effective in accomplishing this requirement?


~~probably **DynamoDB accelerator**~~, maybe autoscaling group?(won't it take too long?)

> **Explanation:**\
> The most likely cause of the processing delays is insufficient instances in the middle tier where the order processing takes place. The most effective solution to reduce processing times in this case is to scale based on the backlog per instance (number of messages in the SQS queue) as this reflects the amount of work that needs to be done.



### Q18
> A web application allows users to upload photos and add graphical elements to them. The application offers two tiers of service: free and paid. Photos uploaded by paid users should be processed before those submitted using the free tier. The photos are uploaded to an Amazon S3 bucket which uses an event notification to send the job information to Amazon SQS.\
> How should a Solutions Architect configure the Amazon SQS deployment to meet these requirements?

~~**short polling for the free, long polling to the paid**~~

> **Explanation:**\
> AWS recommend using separate queues when you need to provide prioritization of work. The logic can then be implemented at the application layer to prioritize the queue for the paid photos over the queue for the free photos.


### Q19
>A company requires that all AWS IAM user accounts have specific complexity requirements and minimum password length.\
How should a Solutions Architect accomplish this?

**password policy** for the entire account (if possible)

> **Explanation:**\
> The easiest way to enforce this requirement is to update the password policy that applies to the entire AWS account. When you create or change a password policy, most of the password policy settings are enforced the next time your users change their passwords. However, some of the settings are enforced immediately such as the password expiration period.


### Q20
> A company is deploying a fleet of Amazon EC2 instances running Linux across multiple Availability Zones within an AWS Region. The application requires a data storage solution that can be accessed by all of the EC2 instances simultaneously. The solution must be highly scalable and easy to implement. The storage must be mounted using the NFS protocol.\
> Which solution meets these requirements?

i think **NFS should means EFS**?

> **Explanation:**\
> Amazon EFS provides scalable file storage for use with Amazon EC2. You can use an EFS file system as a common data source for workloads and applications running on multiple instances. The EC2 instances can run in multiple AZs within a Region and the NFS protocol is used to mount the file system.\
> With EFS you can create mount targets in each AZ for lower latency. The application instances in each AZ will mount the file system using the local mount target.


### Q21

> A company delivers content to subscribers distributed globally from an application running on AWS. The application uses a fleet of Amazon EC2 instance in a private subnet behind an Application Load Balancer (ALB). Due to an update in copyright restrictions, it is necessary to block access for specific countries.\
> What is the EASIEST method to meet this requirement?

(maybe acl?, probably **CloudFront**)

> Explanation
> When a user requests your content, CloudFront typically serves the requested content regardless of where the user is located. If you need to prevent users in specific countries from accessing your content, you can use the CloudFront geo restriction feature to do one of the following:
> - Allow your users to access your content only if they're in one of the countries on a whitelist of approved countries.
> - Prevent your users from accessing your content if they're in one of the countries on a blacklist of banned countries.
> 
> For example, if a request comes from a country where, for copyright reasons, you are not authorized to distribute your content, you can use CloudFront geo restriction to block the request.\
> This is the easiest and most effective way to implement a geographic restriction for the delivery of content.

### Q22
> An application running on Amazon EC2 needs to asynchronously invoke an AWS Lambda function to perform data processing. The services should be decoupled.\
> Which service can be used to decouple the compute services?
> - AWS Config
> - AWS Step Functions
> - AWS MQ
> - AWS SNS

~~are **step functions a thing?**~~

> Explanation
> You can use a Lambda function to process Amazon Simple Notification Service notifications. Amazon SNS supports Lambda functions as a target for messages sent to a topic. This solution decouples the Amazon EC2 application from Lambda and ensures the Lambda function is invoked.

### Q23
> A new application will run across multiple Amazon ECS tasks. Front-end application logic will process data and then pass that data to a back-end ECS task to perform further processing and write the data to a datastore. The Architect would like to reduce-interdependencies so failures do no impact other components.\
> Which solution should the Architect use?


**front pushes to SQS, backend polls**

> Explanation
> This is a good use case for Amazon SQS. SQS is a service that is used for decoupling applications, thus reducing interdependencies, through a message bus. The front-end application can place messages on the queue and the back-end can then poll the queue for new messages. Please remember that Amazon SQS is pull-based (polling) not push-based (use SNS for push-based).


### Q24

> A company wishes to restrict access to their Amazon DynamoDB table to specific, private source IP addresses from their VPC.\
> What should be done to secure access to the table?
> - Create an AWS VPN Connection the amazon DynamoDB endpoint
> - Create gateway VPC endpoint and add an entry to the route table
> - Create an interface VPC endpoint in the VPC with an elastic network interface
> - Create the amazon dynamoDB table in the VPC

~~maybe **vpc endpoint?**~~

> Explanation
> There are two different types of VPC endpoint: interface endpoint, and gateway endpoint. 
> - With an interface endpoint you use an ENI in the VPC. 
> - With a gateway endpoint you configure your route table to point to the endpoint. 
> 
> Amazon S3 and DynamoDB use gateway endpoints. This solution means that all traffic will go through the VPC endpoint straight to DynamoDB using private IP addresses.
> 
> ![Q24 endpoint](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-48-44-654011ba439713ad1e049b259d7d5611.jpg)


### Q25
> A company runs a dynamic website that is hosted on an on-premises server in the United States. The company is expanding to Europe and is investigating how they can optimize the performance of the website for European users. The website’s backed must remain in the United States. The company requires a solution that can be implemented within a few days.\
> What should a Solutions Architect recommend?

~~maybe **Lambda Edge**~~

> Explanation
> A custom origin can point to an on-premises server and CloudFront is able to cache content for dynamic websites. CloudFront can provide performance optimizations for custom origins even if they are running on on-premises servers. These include persistent TCP connections to the origin, SSL enhancements such as Session tickets and OCSP stapling.
> Additionally, connections are routed from the nearest Edge Location to the user across the AWS global network. If the on-premises server is connected via a Direct Connect (DX) link this can further improve performance.

### Q26
> A company plans to make an Amazon EC2 Linux instance unavailable outside of business hours to save costs. The instance is backed by an Amazon EBS volume. There is a requirement that the contents of the instance’s memory must be preserved when it is made unavailable.\
> How can a solutions architect meet these requirements?

**Hibernation**

> Explanation
> When you hibernate an instance, Amazon EC2 signals the operating system to perform hibernation (suspend-to-disk). Hibernation saves the contents from the instance memory (RAM) to your Amazon Elastic Block Store (Amazon EBS) root volume. Amazon EC2 persists the instance's EBS root volume and any attached EBS data volumes. When you start your instance:
> - The EBS root volume is restored to its previous state
> - The RAM contents are reloaded
> - The processes that were previously running on the instance are resumed
> - Previously attached data volumes are reattached and the instance retains its instance ID


### Q27

> The database tier of a web application is running on a Windows server on-premises. The database is a Microsoft SQL Server database. The application owner would like to migrate the database to an Amazon RDS instance.\
> How can the migration be executed with minimal administrative effort and downtime?

use **Database migration service**

> Explanation
> You can directly migrate Microsoft SQL Server from an on-premises server into Amazon RDS using the Microsoft SQL Server database engine. This can be achieved using the native Microsoft SQL Server tools, or using AWS DMS as depicted below:
> ![Q27 database migration](https://img-c.udemycdn.com/redactor/raw/2020-05-17_07-59-54-f1a0a1c8024b0af89fa52ac80a9bc55c.JPG)


### Q28
> An insurance company has a web application that serves users in the United Kingdom and Australia. The application includes a database tier using a MySQL database hosted in eu-west-2. The web tier runs from eu-west-2 and ap-southeast-2. Amazon Route 53 geoproximity routing is used to direct users to the closest web tier. It has been noted that Australian users receive slow response times to queries.\
> Which changes should be made to the database tier to improve performance?


maybe **move to aurora**?

> Explanation
> The issue here is latency with read queries being directed from Australia to UK which is great physical distance. A solution is required for improving read performance in Australia.\
> An Aurora global database consists of one primary AWS Region where your data is mastered, and up to five read-only, secondary AWS Regions. Aurora replicates data to the secondary AWS Regions with typical latency of under a second. You issue write operations directly to the primary DB instance in the primary AWS Region.\
> This solution will provide better performance for users in the Australia Region for queries. Writes must still take place in the UK Region but read performance will be greatly improved.
> 
> ![Q28 Aurora replicates](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-05-41-0e623951d5ed62017a22f93bb14c6c67.jpg)

### Q29
> A web application runs in public and private subnets. The application architecture consists of a web tier and database tier running on Amazon EC2 instances. Both tiers run in a single Availability Zone (AZ).\
> Which combination of steps should a solutions architect take to provide high availability for this architecture? (Select TWO.)

**multi az**

> Explanation
> To add high availability to this architecture both the web tier and database tier require changes. For the web tier an Auto Scaling group across multiple AZs with an ALB will ensure there are always instances running and traffic is being distributed to them.\
> The database tier should be migrated from the EC2 instances to Amazon RDS to take advantage of a managed database with Multi-AZ functionality. This will ensure that if there is an issue preventing access to the primary database a secondary database can take over.
### Q30
> A financial services company has a web application with an application tier running in the U.S and Europe. The database tier consists of a MySQL database running on Amazon EC2 in us-west-1. Users are directed to the closest application tier using Route 53 latency-based routing. The users in Europe have reported poor performance when running queries.\
> Which changes should a Solutions Architect make to the database tier to improve performance?

~~Read replicas in europe?~~


> Explanation
> Amazon Aurora Global Database is designed for globally distributed applications, allowing a single Amazon Aurora database to span multiple AWS regions. It replicates your data with no impact on database performance, enables fast local reads with low latency in each region, and provides disaster recovery from region-wide outages.\
> A global database can be configured in the European region and then the application tier in Europe will need to be configured to use the local database for reads/queries. The diagram below depicts an Aurora Global Database deployment.
> 
> ![Q30 Aurora Global Database](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-40-32-e096373484158f919b4e661ab7c1aa8d.jpg)


### Q31
> A manufacturing company captures data from machines running at customer sites. Currently, thousands of machines send data every 5 minutes, and this is expected to grow to hundreds of thousands of machines in the near future. The data is logged with the intent to be analyzed in the future as needed.\
> What is the SIMPLEST method to store this streaming data at scale?

**Kinesis firehose**

> Explanation
> Kinesis Data Firehose is the easiest way to load streaming data into data stores and analytics tools. It captures, transforms, and loads streaming data and you can deliver the data to “destinations” including Amazon S3 buckets for later analysis.


### Q32
> An Amazon RDS Read Replica is being deployed in a separate region. The master database is not encrypted but all data in the new region must be encrypted.\
> How can this be achieved?

~~KNS keys when creating the read replicas?~~

> Explanation
> You cannot create an encrypted Read Replica from an unencrypted master DB instance. You also cannot enable encryption after launch time for the master DB instance. Therefore, you must create a new master DB by taking a snapshot of the existing DB, encrypting it, and then creating the new DB from the snapshot. You can then create the encrypted cross-region Read Replica of the master DB.


### Q33
> A video production company is planning to move some of its workloads to the AWS Cloud. The company will require around 5 TB of storage for video processing with the maximum possible I/O performance. They also require over 400 TB of extremely durable storage for storing video files and 800 TB of storage for long-term archival.\
> Which combinations of services should a Solutions Architect use to meet these requirements?

Instance store for performance, S3 for storage, Glacier for archiving.

> Explanation
> The best I/O performance can be achieved by using instance store volumes for the video processing. This is safe to use for use cases where the data can be recreated from the source files so this is a good use case.
> For storing data durably Amazon S3 is a good fit as it provides 99.999999999% of durability. For archival the video files can then be moved to Amazon S3 Glacier which is a low cost storage option that is ideal for long-term archival.


### Q34
> A developer created an application that uses Amazon EC2 and an Amazon RDS MySQL database instance. The developer stored the database user name and password in a configuration file on the root EBS volume of the EC2 application instance. A Solutions Architect has been asked to design a more secure solution.\
> What should the Solutions Architect do to achieve this requirement?

IAM permissions on the EC2 instance.

> Explanation
> The key problem here is having plain text credentials stored in a file. Even if you encrypt the volume there is still as security risk as the credentials are loaded by the application and passed to RDS.
> The best way to secure this solution is to get rid of the credentials completely by using an IAM role instead. The IAM role can be assigned permissions to the database instance and can be attached to the EC2 instance. The instance will then obtain temporary security credentials from AWS STS which is much more secure.


### Q35
> A Microsoft Windows file server farm uses Distributed File System Replication (DFSR) to synchronize data in an on-premises environment. The infrastructure is being migrated to the AWS Cloud.\
> Which service should the solutions architect use to replace the file server farm?
> - AWS storage gateway
> - AWS EFS
> - AWS FSx
> - AWS EBS

FSX

> Explanation
> Amazon FSx for Windows file server supports DFS namespaces and DFS replication. This is the best solution for replacing the on-premises infrastructure. Note the limitations for deployment:
> 
> ![Q35 deployment limitations](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_09-52-13-238b160c6f54f8650622a66a33fc34a4.jpg)


### Q36
> A company runs an application on an Amazon EC2 instance the requires 250 GB of storage space. The application is not used often and has small spikes in usage on weekday mornings and afternoons. The disk I/O can vary with peaks hitting a maximum of 3,000 IOPS. A Solutions Architect must recommend the most cost-effective storage solution that delivers the performance required?\
> Which configuration should the Solutions Architect recommend?
> - Amazon EBS Cold HDD(sc1)
> - Amazon EBS provisioned IOS SSD(i01)
> - Amazon EBS Throughput Optimized HDD(st1)
> - Amazon EBS General purpose SSD (gp2)

~~Throughput optimized HDD~~

> Explanation
> General Purpose SSD (gp2) volumes offer cost-effective storage that is ideal for a broad range of workloads. These volumes deliver single-digit millisecond latencies and the ability to burst to 3,000 IOPS for extended periods of time.\
> Between a minimum of 100 IOPS (at 33.33 GiB and below) and a maximum of 16,000 IOPS (at 5,334 GiB and above), baseline performance scales linearly at 3 IOPS per GiB of volume size. AWS designs gp2 volumes to deliver their provisioned performance 99% of the time. A gp2 volume can range in size from 1 GiB to 16 TiB.\
> In this configuration the volume will provide a baseline performance of 750 IOPS but will always be able to burst to the required 3,000 IOPS during periods of increased traffic.


### Q37
> An e-commerce application is hosted in AWS. The last time a new product was launched, the application experienced a performance issue due to an enormous spike in traffic. Management decided that capacity must be doubled this week after the product is launched.\
> What is the MOST efficient way for management to ensure that capacity requirements are met?
> - Add a step scaling policy
> - Add a simple scaling policy
> - Add Amazon EC2 spot instance
> - Add a scheduled scaling action

scheduled scaling

> Explanation
> Scaling based on a schedule allows you to set your own scaling schedule for predictable load changes. To configure your Auto Scaling group to scale based on a schedule, you create a scheduled action. This is ideal for situations where you know when and for how long you are going to need the additional capacity.


### Q38
> A legacy tightly-coupled High Performance Computing (HPC) application will be migrated to AWS.\
> Which network adapter type should be used?
> - Elastic network adaptor (ENA)
> - Elastic Fabric adaptor (EFA)
> - Elastic Ip address
> - Elastic Network interface (ENI)

EFA

> Explanation
> An Elastic Fabric Adapter is an AWS Elastic Network Adapter (ENA) with added capabilities. The EFA lets you apply the scale, flexibility, and elasticity of the AWS Cloud to tightly-coupled HPC apps. It is ideal for tightly coupled app as it uses the Message Passing Interface (MPI).


### Q39
> A company's web application is using multiple Amazon EC2 Linux instances and storing data on Amazon EBS volumes. The company is looking for a solution to increase the resiliency of the application in case of a failure.\
> What should a solutions architect do to meet these requirements?

~~EC2 instances in each AZ with EBS volumes~~

> Explanation
> To increase the resiliency of the application the solutions architect can use Auto Scaling groups to launch and terminate instances across multiple availability zones based on demand. An application load balancer (ALB) can be used to direct traffic to the web application running on the EC2 instances.
> Lastly, the Amazon Elastic File System (EFS) can assist with increasing the resilience of the application by providing a shared file system that can be mounted by multiple EC2 instances from multiple availability zones.


### Q40
> A company's application is running on Amazon EC2 instances in a single Region. In the event of a disaster, a solutions architect needs to ensure that the resources can also be deployed to a second Region.\
> Which combination of actions should the solutions architect take to accomplish this? (Select TWO.)

make an AMI, copy to other region, launch there

> Explanation
> You can copy an Amazon Machine Image (AMI) within or across AWS Regions using the AWS Management Console, the AWS Command Line Interface or SDKs, or the Amazon EC2 API, all of which support the CopyImage action.
> Using the copied AMI the solutions architect would then be able to launch an instance from the same EBS volume in the second Region.

### Q41
> A solutions architect is designing a new service that will use an Amazon API Gateway API on the frontend. The service will need to persist data in a backend database using key-value requests. Initially, the data requirements will be around 1 GB and future growth is unknown. Requests can range from 0 to over 800 requests per second.\
> Which combination of AWS services would meet these requirements? (Select TWO.)
> - AWS Lambda
> - AWS EC2 autos scaling
> - AWS DynamoDB
> - Amazon RDS
> - AWS fargate

**Lambda** and **Dynamo**

> Explanation
> In this case AWS Lambda can perform the computation and store the data in an Amazon DynamoDB table. Lambda can scale concurrent executions to meet demand easily and DynamoDB is built for key-value data storage requirements and is also serverless and easily scalable. This is therefore a cost effective solution for unpredictable workloads.


### Q42
> A company has two accounts for perform testing and each account has a single VPC: VPC-TEST1 and VPC-TEST2. The operations team require a method of securely copying files between Amazon EC2 instances in these VPCs. The connectivity should not have any single points of failure or bandwidth constraints.\
> Which solution should a Solutions Architect recommend?

VPC peering

> Explanation
> A VPC peering connection is a networking connection between two VPCs that enables you to route traffic between them using private IPv4 addresses or IPv6 addresses. Instances in either VPC can communicate with each other as if they are within the same network.
> You can create a VPC peering connection between your own VPCs, or with a VPC in another AWS account. The VPCs can be in different regions (also known as an inter-region VPC peering connection).
> ![Q42 VPC peering](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_09-50-39-d2cdb66079d7dfcc2c83b098f21eafa2.jpg)


### Q43
> A team are planning to run analytics jobs on log files each day and require a storage solution. The size and number of logs is unknown and data will persist for 24 hours only.\
> What is the MOST cost-effective solution?
> - Amazon S3 intelligent tiering
> - Amazon S3 glacier deep archive
> - Amazon S3 Standard
> - Amazon S3 one zone infrequent access

**S3 standard**

> Explanation
> S3 standard is the best choice in this scenario for a short term storage solution. In this case the size and number of logs is unknown and it would be difficult to fully assess the access patterns at this stage. Therefore, using S3 standard is best as it is cost-effective, provides immediate access, and there are no retrieval fees or minimum capacity charge per object.


### Q44
> A company is working with a strategic partner that has an application that must be able to send messages to one of the company’s Amazon SQS queues. The partner company has its own AWS account.\
> How can a Solutions Architect provide least privilege access to the partner?

update the SQS policy and grant permissions 

> Explanation
> Amazon SQS supports resource-based policies. The best way to grant the permissions using the principle of least privilege is to use a resource-based policy attached to the SQS queue that grants the partner company’s AWS account the sqs:SendMessage privilege.
> The following policy is an example of how this could be configured:
> 
> ![Q44 permission policies](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-46-16-3d75509bc08efb886557569410905bfe.jpg)


### Q45
> A company runs an application in a factory that has a small rack of physical compute resources. The application stores data on a network attached storage (NAS) device using the NFS protocol. The company requires a daily offsite backup of the application data.\
> Which solution can a Solutions Architect recommend to meet this requirement?

~~AWS Storgate gateway volume to replicate data to S3~~

> Explanation
> The AWS Storage Gateway Hardware Appliance is a physical, standalone, validated server configuration for on-premises deployments. It comes pre-loaded with Storage Gateway software, and provides all the required CPU, memory, network, and SSD cache resources for creating and configuring File Gateway, Volume Gateway, or Tape Gateway.
> A file gateway is the correct type of appliance to use for this use case as it is suitable for mounting via the NFS and SMB protocols.


### Q46
> A company hosts an application on Amazon EC2 instances behind Application Load Balancers in several AWS Regions. Distribution rights for the content require that users in different geographies must be served content from specific regions.\
> Which configuration meets these requirements?
> - Route53 with geoproximity
> - Route53 with geolocation
> - ALB with multi-region routing
> - CloudFronot with multiple origins and AWS WAF

route53 with geolocation

> Explanation
> To protect the distribution rights of the content and ensure that users are directed to the appropriate AWS Region based on the location of the user, the geolocation routing policy can be used with Amazon Route 53.
> Geolocation routing lets you choose the resources that serve your traffic based on the geographic location of your users, meaning the location that DNS queries originate from.
> When you use geolocation routing, you can localize your content and present some or all of your website in the language of your users. You can also use geolocation routing to restrict distribution of content to only the locations in which you have distribution rights.


### Q47

> Amazon EC2 instances in a development environment run between 9am and 5pm Monday-Friday. Production instances run 24/7.\
> Which pricing models should be used? (choose 2)

reserved instances for production, on-demand capacity reserations for development.

> Explanation
> Capacity reservations have no commitment and can be created and canceled as needed. This is ideal for the development environment as it will ensure the capacity is available. There is no price advantage but none of the other options provide a price advantage whilst also ensuring capacity is available.\
> Reserved instances are a good choice for workloads that run continuously. This is a good option for the production environment.


### Q48
> An organization want to share regular updates about their charitable work using static webpages. The pages are expected to generate a large amount of views from around the world. The files are stored in an Amazon S3 bucket. A solutions architect has been asked to design an efficient and effective solution.\
> Which action should the solutions architect take to accomplish this?

CloudFront with S3 as origin.

> Explanation
> Amazon CloudFront can be used to cache the files in edge locations around the world and this will improve the performance of the webpages.\
> To serve a static website hosted on Amazon S3, you can deploy a CloudFront distribution using one of these configurations:
> - Using a REST API endpoint as the origin with access restricted by an origin access identity (OAI)
> - Using a website endpoint as the origin with anonymous (public) access allowed
> - Using a website endpoint as the origin with access restricted by a Referer header


### Q49
> An application running on an Amazon ECS container instance using the EC2 launch type needs permissions to write data to Amazon DynamoDB.\
> How can you assign these permissions only to the specific ECS task that is running the application?

IAM policy attached to the ~~container Instance~~

> Explanation
> To specify permissions for a specific task on Amazon ECS you should use IAM Roles for Tasks. The permissions policy can be applied to tasks when creating the task definition, or by using an IAM task role override using the AWS CLI or SDKs. The taskRoleArn parameter is used to specify the policy.

### Q50
> An Amazon VPC contains several Amazon EC2 instances. The instances need to make API calls to Amazon DynamoDB. A solutions architect needs to ensure that the API calls do not traverse the internet.
> How can this be accomplished? (Select TWO.)

~~ENI endpoint~~, route table entry

> Explanation
> Amazon DynamoDB and Amazon S3 support gateway endpoints, not interface endpoints. With a gateway endpoint you create the endpoint in the VPC, attach a policy allowing access to the service, and then specify the route table to create a route table entry in.
> 
> ![Q50 API calls not traversing the internet](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-00-45-ac665c89acb1641afb831f1eb795210e.jpg)


### Q51
> An application is being created that will use Amazon EC2 instances to generate and store data. Another set of EC2 instances will then analyze and modify the data. Storage requirements will be significant and will continue to grow over time. The application architects require a storage solution.\
> Which actions would meet these needs?

using EFS and all EC2 instances mount it.

> Explanation
> Amazon Elastic File System (Amazon EFS) provides a simple, scalable, fully managed elastic NFS file system for use with AWS Cloud services and on-premises resources. It is built to scale on-demand to petabytes without disrupting applications, growing and shrinking automatically as you add and remove files, eliminating the need to provision and manage capacity to accommodate growth.\
> Amazon EFS supports the Network File System version 4 (NFSv4.1 and NFSv4.0) protocol. Multiple Amazon EC2 instances can access an Amazon EFS file system at the same time, providing a common data source for workloads and applications running on more than one instance or server.\
> For this scenario, EFS is a great choice as it will provide a scalable file system that can be mounted by multiple EC2 instances and accessed simultaneously.
> 
> ![Q51 attaching EFS](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-51-10-5bfba0da41e8553b6b332f2566f0f56e.jpg)


### Q52
> A solutions architect needs to backup some application log files from an online e-commerce store to Amazon S3. It is unknown how often the logs will be accessed or which logs will be accessed the most. The solutions architect must keep costs as low as possible by using the appropriate S3 storage class.\
> Which S3 storage class should be implemented to meet these requirements?
> - Amazon S3 infrequent access
> - Amazon S3 intelligent tiering
> - Amazon S3 one zone infrequent access
> - Amazon S3 glacier 

S3 intelligent tiering

> Explanation
> The S3 Intelligent-Tiering storage class is designed to optimize costs by automatically moving data to the most cost-effective access tier, without performance impact or operational overhead.
> It works by storing objects in two access tiers: one tier that is optimized for frequent access and another lower-cost tier that is optimized for infrequent access. This is an ideal use case for intelligent-tiering as the access patterns for the log files are not known.


### Q53
> A company is investigating methods to reduce the expenses associated with on-premises backup infrastructure. The Solutions Architect wants to reduce costs by eliminating the use of physical backup tapes. It is a requirement that existing backup applications and workflows should continue to function.\
> What should the Solutions Architect recommend?

AWS Storage gateway with virtual tape library (VTL)

> Explanation
> The AWS Storage Gateway Tape Gateway enables you to replace using physical tapes on premises with virtual tapes in AWS without changing existing backup workflows. Tape Gateway emulates physical tape libraries, removes the cost and complexity of managing physical tape infrastructure, and provides more durability than physical tapes.
> 
> ![Q53 Tape Gateway](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_09-45-15-153ae046c9838c5d880c4f0fce858218.jpg)


### Q54
> A company uses an Amazon RDS MySQL database instance to store customer order data. The security team have requested that SSL/TLS encryption in transit must be used for encrypting connections to the database from application servers. The data in the database is currently encrypted at rest using an AWS KMS key.
> How can a Solutions Architect enable encryption in transit?

~~use self-signed certificates~~

> Explanation
> Amazon RDS creates an SSL certificate and installs the certificate on the DB instance when Amazon RDS provisions the instance. These certificates are signed by a certificate authority. The SSL certificate includes the DB instance endpoint as the Common Name (CN) for the SSL certificate to guard against spoofing attacks.\
> You can download a root certificate from AWS that works for all Regions or you can download Region-specific intermediate certificates.


### Q55
> A company is migrating from an on-premises infrastructure to the AWS Cloud. One of the company's applications stores files on a Windows file server farm that uses Distributed File System Replication (DFSR) to keep data in sync. A solutions architect needs to replace the file server farm.\
> Which service should the solutions architect use?
> - S3
> - FSx
> - EFS
> - Storage gateway

FSx

> Explanation
> Amazon FSx for Windows File Server provides fully managed, highly reliable file storage that is accessible over the industry-standard Server Message Block (SMB) protocol.\
> Amazon FSx is built on Windows Server and provides a rich set of administrative features that include end-user file restore, user quotas, and Access Control Lists (ACLs).\
> Additionally, Amazon FSX for Windows File Server supports Distributed File System Replication (DFSR) in Single-AZ deployments as can be seen in the feature comparison table below.
> 
> ![Q55 Fsx deployments](https://img-c.udemycdn.com/redactor/raw/2020-05-17_14-56-37-c7a7eff7f9da52f8dadafb160c3ac4c0.JPG)


### Q56
> Your company shares some HR videos stored in an Amazon S3 bucket via CloudFront. You need to restrict access to the private content so users coming from specific IP addresses can access the videos and ensure direct access via the Amazon S3 bucket is not possible.\
> How can this be achieved?

Cloudfront with signedURL, origin access policy then restricting S3 bucket.

> Explanation
> A signed URL includes additional information, for example, an expiration date and time, that gives you more control over access to your content. You can also specify the IP address or range of IP addresses of the users who can access your content.\
> If you use CloudFront signed URLs (or signed cookies) to limit access to files in your Amazon S3 bucket, you may also want to prevent users from directly accessing your S3 files by using Amazon S3 URLs. To achieve this you can create an origin access identity (OAI), which is a special CloudFront user, and associate the OAI with your distribution.\
> You can then change the permissions either on your Amazon S3 bucket or on the files in your bucket so that only the origin access identity has read permission (or read and download permission).
> ![Q56 S3 origin and cloudfront](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-44-34-2bf62dad00d3baac5aa2bf88745d793e.jpg)


### Q57
> An eCommerce company runs an application on Amazon EC2 instances in public and private subnets. The web application runs in a public subnet and the database runs in a private subnet. Both the public and private subnets are in a single Availability Zone.\
> Which combination of steps should a solutions architect take to provide high availability for this architecture? (Select TWO.)

EC2 autoscaling in multiple AZ, private and public subnets and multi-region RDS.

> Explanation
> High availability can be achieved by using multiple Availability Zones within the same VPC. An EC2 Auto Scaling group can then be used to launch web application instances in multiple public subnets across multiple AZs and an ALB can be used to distribute incoming load.\
> The database solution can be made highly available by migrating from EC2 to Amazon RDS and using a Multi-AZ deployment model. This will provide the ability to failover to another AZ in the event of a failure of the primary database or the AZ in which it runs.


### Q58
> An organization has a large amount of data on Windows (SMB) file shares in their on-premises data center. The organization would like to move data into Amazon S3. They would like to automate the migration of data over their AWS Direct Connect link.\
> Which AWS service can assist them?
> - DataSync
> - Database migration Service
> - Database Cloud Formation
> - AWS snowball

DataSync

> Explanation
> AWS DataSync can be used to move large amounts of data online between on-premises storage and Amazon S3 or Amazon Elastic File System (Amazon EFS). DataSync eliminates or automatically handles many of these tasks, including scripting copy jobs, scheduling and monitoring transfers, validating data, and optimizing network utilization.\
> The source datastore can use file servers that communicate using the Server Message Block (SMB) protocol.

### Q59
> A company runs an application in an on-premises data center that collects environmental data from production machinery. The data consists of JSON files stored on network attached storage (NAS) and around 5 TB of data is collected each day. The company must upload this data to Amazon S3 where it can be processed by an analytics application. The data must be transferred securely.\
> Which solution offers the MOST reliable and time-efficient data transfer?

AWS DataSync over DirectConnect

> Explanation
> The most reliable and time-efficient solution that keeps the data secure is to use AWS DataSync and synchronize the data from the NAS device directly to Amazon S3. This should take place over an AWS Direct Connect connection to ensure reliability, speed, and security.\
> AWS DataSync can copy data between Network File System (NFS) shares, Server Message Block (SMB) shares, self-managed object storage, AWS Snowcone, Amazon Simple Storage Service (Amazon S3) buckets, Amazon Elastic File System (Amazon EFS) file systems, and Amazon FSx for Windows File Server file systems.


### Q60
> A Solutions Architect has been tasked with re-deploying an application running on AWS to enable high availability. The application processes messages that are received in an ActiveMQ queue running on a single Amazon EC2 instance. Messages are then processed by a consumer application running on Amazon EC2. After processing the messages the consumer application writes results to a MySQL database running on Amazon EC2.\
> Which architecture offers the highest availability and low operational complexity?

AmazonMQ accross two AZ, autosacling in multi-AZ and RDS in multi-AZ.

> Explanation
> The correct answer offers the highest availability as it includes Amazon MQ active/standby brokers across two AZs, an Auto Scaling group across two AZ,s and a Multi-AZ Amazon RDS MySQL database deployment.\
> This architecture not only offers the highest availability it is also operationally simple as it maximizes the usage of managed services.


### Q61
> A solutions architect is creating a system that will run analytics on financial data for several hours a night, 5 days a week. The analysis is expected to run for the same duration and cannot be interrupted once it is started. The system will be required for a minimum of 1 year.\
> What should the solutions architect configure to ensure the EC2 instances are available when they are needed?

~~Regional Reserved instances~~

> Explanation
> On-Demand Capacity Reservations enable you to reserve compute capacity for your Amazon EC2 instances in a specific Availability Zone for any duration. This gives you the ability to create and manage Capacity Reservations independently from the billing discounts offered by Savings Plans or Regional Reserved Instances.\
> By creating Capacity Reservations, you ensure that you always have access to EC2 capacity when you need it, for as long as you need it. You can create Capacity Reservations at any time, without entering a one-year or three-year term commitment, and the capacity is available immediately.\
> The table below shows the difference between capacity reservations and other options:
> 
> ![Q61 EC2 reservations](https://img-c.udemycdn.com/redactor/raw/test_question_description/2022-01-10_20-24-52-810be510ec31f342fd622453a1ebec94.png)


### Q62
> A company runs a large batch processing job at the end of every quarter. The processing job runs for 5 days and uses 15 Amazon EC2 instances. The processing must run uninterrupted for 5 hours per day. The company is investigating ways to reduce the cost of the batch processing job.\
> Which pricing model should the company choose?

~~Spot Instances~~ 

> Explanation
> Each EC2 instance runs for 5 hours a day for 5 days per quarter or 20 days per year. This is time duration is insufficient to warrant reserved instances as these require a commitment of a minimum of 1 year and the discounts would not outweigh the costs of having the reservations unused for a large percentage of time.\
> In this case, there are no options presented that can reduce the cost and therefore on-demand instances should be used.


### Q63
> An AWS Organization has an OU with multiple member accounts in it. The company needs to restrict the ability to launch only specific Amazon EC2 instance types.\
> How can this policy be applied across the accounts with the least effort?

~~use Resources Access Manager~~

> Explanation
> To apply the restrictions across multiple member accounts you must use a Service Control Policy (SCP) in the AWS Organization. The way you would do this is to create a deny rule that applies to anything that does not equal the specific instance type you want to allow.\
> The following architecture could be used to achieve this goal:
> 
> ![Q63 Service control policy](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-48-06-e73551726ea3baf15e8e687947948cba.jpg)


### Q64
> A company hosts a multiplayer game on AWS. The application uses Amazon EC2 instances in a single Availability Zone and users connect over Layer 4. Solutions Architect has been tasked with making the architecture highly available and also more cost-effective.\
> How can the solutions architect best meet these requirements? (Select TWO.)

Network Load Balancer, AutoScaling group in multiple AZ.

> Explanation
> The solutions architect must enable high availability for the architecture and ensure it is cost-effective. To enable high availability an Amazon EC2 Auto Scaling group should be created to add and remove instances across multiple availability zones.\
> In order to distribute the traffic to the instances the architecture should use a Network Load Balancer which operates at Layer 4. This architecture will also be cost-effective as the Auto Scaling group will ensure the right number of instances are running based on demand.


### Q65
> A company uses Docker containers for many application workloads in an on-premise data center. The company is planning to deploy containers to AWS and the chief architect has mandated that the same configuration and administrative tools must be used across all containerized environments. The company also wishes to remain cloud agnostic to safeguard mitigate the impact of future changes in cloud strategy.\
> How can a Solutions Architect design a managed solution that will align with open-source software?

use EKS

> Explanation
> Amazon EKS is a managed service that can be used to run Kubernetes on AWS. Kubernetes is an open-source system for automating the deployment, scaling, and management of containerized applications. Applications running on Amazon EKS are fully compatible with applications running on any standard Kubernetes environment, whether running in on-premises data centers or public clouds. This means that you can easily migrate any standard Kubernetes application to Amazon EKS without any code modification.\
> This solution ensures that the same open-source software is used for automating the deployment, scaling, and management of containerized applications both on-premises and in the AWS Cloud.
> 
> ![65 EKS](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_09-55-33-2abdd31ee10dd8c28e19ec75e2a29309.jpg)


</details>

## Exam 2
<details>
<summary>
65 questions
</summary>

### Q01
> An application uses Amazon EC2 instances and an Amazon RDS MySQL database. The database is not currently encrypted. A solutions architect needs to apply encryption to the database for all new and existing data.\
> How should this be accomplished?

snapshot, encrypted copy, restore from encrypted copy.

> **Explanation:**\
> There are some [limitations for encrypted Amazon RDS DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.html#Overview.Encryption.Limitations)s: you can't modify an existing unencrypted Amazon RDS DB instance to make the instance encrypted, and you can't create an encrypted read replica from an unencrypted instance.\
> However, you can use the Amazon RDS snapshot feature to encrypt an unencrypted snapshot that's taken from the RDS database that you want to encrypt. Restore a new RDS DB instance from the encrypted snapshot to deploy a new encrypted DB instance. Finally, switch your connections to the new DB instance.

### Q02
> A DynamoDB database you manage is randomly experiencing heavy read requests that are causing latency.\
> What is the simplest way to alleviate the performance issues?

DynamoDB is managed, I don't think there are read replicas for it or auto scaling. so maybe **DynamoDB DAX**.

> **Explanation:**\
> DynamoDB offers consistent single-digit millisecond latency. However, DynamoDB + DAX further increases performance with response times in microseconds for millions of requests per second for read-heavy workloads.\
> The DAX cache uses cluster nodes running on Amazon EC2 instances and sits in front of the DynamoDB table as you can see in the diagram below:
> 
> ![Q2 DynamoDB DAX](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-21-53-bb155b9024a83ff1b09b8bcaedfb50c2.jpg)

### Q03
> A website is running on Amazon EC2 instances and access is restricted to a limited set of IP ranges. A solutions architect is planning to migrate static content from the website to an Amazon S3 bucket configured as an origin for an Amazon CloudFront distribution. Access to the static content must be restricted to the same set of IP addresses.\
> Which combination of steps will meet these requirements? (Select TWO.)

AWS WAF web ACL on the cloudfront distribution, OAI on the S3 bucket.

> **Explanation:**\
> To prevent users from circumventing the controls implemented on CloudFront (using WAF or presigned URLs / signed cookies) you can use an origin access identity (OAI). An OAI is a special CloudFront user that you associate with a distribution.\
> The next step is to change the permissions either on your Amazon S3 bucket or on the files in your bucket so that only the origin access identity has read permission (or read and download permission). This can be implemented through a bucket policy.\
> To control access at the CloudFront layer the AWS Web Application Firewall (WAF) can be used. With WAF you must create an ACL that includes the IP restrictions required and then associate the web ACL with the CloudFront distribution.


### Q04

> A company has acquired another business and needs to migrate their 50TB of data into AWS within 1 month. They also require a secure, reliable and private connection to the AWS cloud.\
> How are these requirements best accomplished?

probably snowball, vpn then direct connect.

> **Explanation:**\
> AWS Direct Connect provides a secure, reliable and private connection. However, lead times are often longer than 1 month so it cannot be used to migrate data within the timeframes. Therefore, it is better to use AWS Snowball to move the data and order a Direct Connect connection to satisfy the other requirement later on. In the meantime the organization can use an AWS VPN for secure, private access to their VPC.


### Q05
> A multi-tier application runs with eight front-end web servers in an Amazon EC2 Auto Scaling group in a single Availability Zone behind an Application Load Balancer. A solutions architect needs to modify the infrastructure to be highly available without modifying the application.\
> Which architecture should the solutions architect choose that provides high availability?


use more availability zones.

> **Explanation:**\
> High availability can be enabled for this architecture quite simply by modifying the existing Auto Scaling group to use multiple availability zones. The ASG will automatically balance the load so you don’t actually need to specify the instances per AZ.\
> The architecture for the web tier will look like the one below:
> 
> ![Q5 High Availability](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-27-38-903c5a1a8a69212049571597cbc31b0f.jpg)




### Q06
> A solutions architect has been tasked with designing a highly resilient hybrid cloud architecture connecting an on-premises data center and AWS. The network should include AWS Direct Connect (DX).\
> Which DX configuration offers the HIGHEST resiliency?

~~probably DX and multiple public VIFs?~~

> **Explanation:**\
> The most resilient solution is to configure DX connections at multiple DX locations. This ensures that any issues impacting a single DX location do not affect availability of the network connectivity to AWS.\
> Take note of the following AWS recommendations for resiliency:
> "AWS recommends connecting from multiple data centers for physical location redundancy. When designing remote connections, consider using redundant hardware and telecommunications providers. Additionally, it is a best practice to use dynamically routed, active/active connections for automatic load balancing and failover across redundant network connections. Provision sufficient network capacity to ensure that the failure of one network connection does not overwhelm and degrade redundant connections."\
> The diagram below is an example of an architecture that offers high resiliency:
> 
> ![Q6 DX high resiliency](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_10-13-36-8f8a2d3dea9984331c8894ec400a4303.jpg)




### Q07

> An application running on Amazon ECS processes data and then writes objects to an Amazon S3 bucket. The application requires permissions to make the S3 API calls.\
> How can a Solutions Architect ensure the application has the required permissions?

update the taskRoleArn with IAM policy.

> **Explanation:**\
> With IAM roles for Amazon ECS tasks, you can specify an IAM role that can be used by the containers in a task. Applications must sign their AWS API requests with AWS credentials, and this feature provides a strategy for managing credentials for your applications to use, similar to the way that Amazon EC2 instance profiles provide credentials to EC2 instances.\
> You define the IAM role to use in your task definitions, or you can use a taskRoleArn override when running a task manually with the RunTask API operation.\
> Note that there are instances roles and task roles that you can assign in ECS when using the EC2 launch type. The task role is better when you need to assign permissions for just that specific task.
> 
> ![Q7 Task Roles](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-02-13-cca043ef28407cf8376847f12e461e32.jpg)


### Q08
> An AWS workload in a VPC is running a legacy database on an Amazon EC2 instance. Data is stored on a 2000GB Amazon EBS (gp2) volume. At peak load times, logs show excessive wait time.\
> What should be implemented to improve database performance using persistent storage?

persistent storage excludes instance store volumes, gp2 is already ssd. ~~so maybe **burstable performance**?~~

> **Explanation:**\
> The data is already on an SSD-backed volume (gp2), therefore to improve performance the best option is to migrate the data onto a provisioned IOPS SSD (io1) volume type which will provide improved I/O performance and therefore reduce wait times.

### Q09

> A company have 500 TB of data in an on-premises file share that needs to be moved to Amazon S3 Glacier. The migration must not saturate the company’s low-bandwidth internet connection and the migration must be completed within a few weeks.\
> What is the MOST cost-effective solution?


snowball into S3 with lifecycle policy.

> **Explanation:**\
> As the company’s internet link is low-bandwidth uploading directly to Amazon S3 (ready for transition to Glacier) would saturate the link. The best alternative is to use AWS Snowball appliances. The Snowball edge appliance can hold up to 80 TB of data so 7 devices would be required to migrate 500 TB of data.\
> Snowball moves data into AWS using a hardware device and the data is then copied into an Amazon S3 bucket of your choice. From there, lifecycle policies can transition the S3 objects to Amazon S3 Glacier.

### Q10

> A company is planning a migration for a high performance computing (HPC) application and associated data from an on-premises data center to the AWS Cloud. The company uses tiered storage on premises with hot high-performance parallel storage to support the application during periodic runs of the application, and more economical cold storage to hold the data when the application is not actively running.\
> Which combination of solutions should a solutions architect recommend to support the storage needs of the application? (Select TWO.)


Lustre is for hpc, probably S3 for cold storage

> **Explanation:**\
> Amazon FSx for Lustre provides a high-performance file system optimized for fast processing of workloads such as machine learning, high-performance computing (HPC), video processing, financial modeling, and electronic design automation (EDA).These workloads commonly require data to be presented via a fast and scalable file system interface, and typically have data sets stored on long-term data stores like Amazon S3.\
> Amazon FSx works natively with Amazon S3, making it easy to access your S3 data to run data processing workloads. Your S3 objects are presented as files in your file system, and you can write your results back to S3. This lets you run data processing workloads on FSx for Lustre and store your long-term data on S3 or on-premises data stores.\
> Therefore, the best combination for this scenario is to use S3 for cold data and FSx for Lustre for the parallel HPC job.
> 
> ![Q10 FSX lustre](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-29-14-9ca4fa2525b286cbb51131ad2546dea3.jpg)



### Q11
> An application stores transactional data in an Amazon S3 bucket. The data is analyzed for the first week and then must remain immediately available for occasional analysis.\
> What is the MOST cost-effective storage solution that meets the requirements?

~~S3 standard IA after 7 days~~

> **Explanation:**\
> The transition should be to Standard-IA rather than One Zone-IA. Though One Zone-IA would be cheaper, it also offers lower availability and the question states the objects “must remain immediately available”. Therefore the availability is a consideration.\
> Though there is no minimum duration when storing data in S3 Standard, you cannot transition to Standard IA within 30 days. This can be seen when trying to create a lifecycle rule.\
> Therefore, the best solution is to transition after 30 days. 
> 
> ![Q11 minimum lifecycle](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-12-38-75ade2c799e4d5cf4b201792e2ee34b6.jpg)

### Q12

> An organization is migrating data to the AWS cloud. An on-premises application uses Network File System shares and must access the data without code changes. The data is critical and is accessed frequently.\
> Which storage solution should a Solutions Architect recommend to maximize availability and durability?

NFS usually means EFS, but it's on premises, so maybe **File gatway**?

> **Explanation:**\
> The solution must use NFS file shares to access the migrated data without code modification. This means you can use either Amazon EFS or AWS Storage Gateway – File Gateway. Both of these can be mounted using NFS from on-premises applications.\
> However, EFS is the wrong answer as the solution asks to maximize availability and durability. The File Gateway backs off of Amazon S3 which has much higher availability and durability than EFS which is why it is the best solution for this scenario.

### Q13

> An application has been deployed on Amazon EC2 instances behind an Application Load Balancer (ALB). A Solutions Architect must improve the security posture of the application and minimize the impact of a DDoS attack on resources.\
> Which of the following solutions is MOST effective?

Athena is for datbases, not security, so probably WAF ACL with rates.

> **Explanation:**\
> A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action on IPs with rates that go over a limit. You set the limit as the number of requests per 5-minute time span.\
> You can use this type of rule to put a temporary block on requests from an IP address that's sending excessive requests. By default, AWS WAF aggregates requests based on the IP address from the web request origin, but you can configure the rule to use an IP address from an HTTP header, like X-Forwarded-For, instead.


### Q14

> A company is storing a large quantity of small files in an Amazon S3 bucket. An application running on an Amazon EC2 instance needs permissions to access and process the files in the S3 bucket.\
> Which action will MOST securely grant the EC2 instance access to the S3 bucket?


least privileged IAM role on the EC2

> **Explanation:**\
> IAM roles should be used in place of storing credentials on Amazon EC2 instances. This is the most secure way to provide permissions to EC2 as no credentials are stored and short-lived credentials are obtained using AWS STS. Additionally, the policy attached to the role should provide least privilege permissions.

### Q15

> A company’s Amazon EC2 instances were terminated or stopped, resulting in a loss of important data that was stored on attached EC2 instance stores. They want to avoid this happening in the future and need a solution that can scale as data volumes increase with the LEAST amount of management and configuration.\
> Which storage is most appropriate?

I want to say EBS, but least management is S3...

> **Explanation:**\
> Amazon EFS is a fully managed service that requires no changes to your existing applications and tools, providing access through a standard file system interface for seamless integration. It is built to scale on demand to petabytes without disrupting applications, growing and shrinking automatically as you add and remove files. This is an easy solution to implement and the option that requires the least management and configuration.\
> An instance store provides temporary block-level storage for an EC2 instance. If you terminate the instance you lose all data. The alternative is to use Elastic Block Store volumes which are also block-level storage devices but the data is persistent. However, EBS is not a fully managed solution and doesn’t grow automatically as your data requirements increase – you would need to increase the volume size and then extend your filesystem.


### Q16

> A company wants to migrate a legacy web application from an on-premises data center to AWS. The web application consists of a web tier, an application tier, and a MySQL database. The company does not want to manage instances or clusters.\
> Which combination of services should a solutions architect include in the overall architecture? (Select TWO.)

RDS + fargate

> **Explanation:**\
> Amazon RDS is a managed service and you do not need to manage the instances. This is an ideal backend for the application and you can run a MySQL database on RDS without any refactoring. For the application components these can run on Docker containers with AWS Fargate. Fargate is a serverless service for running containers on AWS.

### Q17

> An application launched on Amazon EC2 instances needs to publish personally identifiable information (PII) about customers using Amazon SNS. The application is launched in private subnets within an Amazon VPC.\
> Which is the MOST secure way to allow the application to access service endpoints in the same region?

~~maybe NAT gateway?~~

> **Explanation:**\
> To publish messages to Amazon SNS topics from an Amazon VPC, create an interface VPC endpoint. Then, you can publish messages to SNS topics while keeping the traffic within the network that you manage with the VPC. This is the most secure option as traffic does not need to traverse the Internet.

### Q18
> An application is running on Amazon EC2 behind an Elastic Load Balancer (ELB). Content is being published using Amazon CloudFront and you need to restrict the ability for users to circumvent CloudFront and access the content directly through the ELB.\
> How can you configure this solution?

~~**ACL on ELB**~~

> **Explanation:**\
> The only way to get this working is by using a VPC Security Group for the ELB that is configured to allow only the internal service IP ranges associated with CloudFront. As these are updated from time to time, you can use AWS Lambda to automatically update the addresses. This is done using a trigger that is triggered when AWS issues an SNS topic update when the addresses are changed.


### Q19

> A company's web application is using multiple Amazon EC2 Linux instances and storing data on Amazon EBS volumes. The company is looking for a solution to increase the resiliency of the application in case of a failure.\
> What should a solutions architect do to meet these requirements?
 
~~ec2 instances in availability zones, EBS on each instance~~

> **Explanation:**\
> To increase the resiliency of the application the solutions architect can use Auto Scaling groups to launch and terminate instances across multiple availability zones based on demand. An application load balancer (ALB) can be used to direct traffic to the web application running on the EC2 instances.\
> Lastly, the Amazon Elastic File System (EFS) can assist with increasing the resilience of the application by providing a shared file system that can be mounted by multiple EC2 instances from multiple availability zones.

### Q20

> A data-processing application runs on an i3.large EC2 instance with a single 100 GB EBS gp2 volume. The application stores temporary data in a small database (less than 30 GB) located on the EBS root volume. The application is struggling to process the data fast enough, and a Solutions Architect has determined that the I/O speed of the temporary database is the bottleneck.\
> What is the MOST cost-efficient way to improve the database response times?

use instance storage instead (temporary data)

> **Explanation:**\
> EC2 Instance Stores are high-speed ephemeral storage that is physically attached to the EC2 instance. The i3.large instance type comes with a single 475GB NVMe SSD instance store so it would be a good way to lower cost and improve performance by using the attached instance store. As the files are temporary, it can be assumed that ephemeral storage (which means the data is lost when the instance is stopped) is sufficient.

### Q21

> A Solutions Architect must design a solution that encrypts data in Amazon S3. Corporate policy mandates encryption keys be generated and managed on premises.\
> Which solution should the Architect use to meet the security requirements?

SSE-C -server side encryption with customer managed keys

> **Explanation:**\
> Server-side encryption is about protecting data at rest. Server-side encryption encrypts only the object data, not object metadata. Using server-side encryption with customer-provided encryption keys (SSE-C) allows you to set your own encryption keys. With the encryption key you provide as part of your request, Amazon S3 manages the encryption as it writes to disks and decryption when you access your objects. Therefore, you don't need to maintain any code to perform data encryption and decryption. The only thing you do is manage the encryption keys you provide.\
> When you upload an object, Amazon S3 uses the encryption key you provide to apply AES-256 encryption to your data and removes the encryption key from memory. When you retrieve an object, you must provide the same encryption key as part of your request. Amazon S3 first verifies that the encryption key you provided matches and then decrypts the object before returning the object data to you.

### Q22

> A group of business analysts perform read-only SQL queries on an Amazon RDS database. The queries have become quite numerous and the database has experienced some performance degradation. The queries must be run against the latest data. A Solutions Architect must solve the performance problems with minimal changes to the existing web application.\
> What should the Solutions Architect recommend?

~~redshift~~

> **Explanation:**\
> The performance issues can be easily resolved by offloading the SQL queries the business analysts are performing to a read replica. This ensures that data that is being queries is up to date and the existing web application does not require any modifications to take place.

### Q23

> A solutions architect is optimizing a website for real-time streaming and on-demand videos. The website’s users are located around the world and the solutions architect needs to optimize the performance for both the real-time and on-demand streaming.\
> Which service should the solutions architect choose?


~~maybe Global accelerator?~~

> **Explanation:**\
> Amazon CloudFront can be used to stream video to users across the globe using a wide variety of protocols that are layered on top of HTTP. This can include both on-demand video as well as real time streaming video.


### Q24

> A company has a file share on a Microsoft Windows Server in an on-premises data center. The server uses a local network attached storage (NAS) device to store several terabytes of files. The management team require a reduction in the data center footprint and to minimize storage costs by moving on-premises storage to AWS.\
> What should a Solutions Architect do to meet these requirements?

~~EFS with IPSec VPN~~

> **Explanation:**\
> An AWS Storage Gateway File Gateway provides your applications a file interface to seamlessly store files as objects in Amazon S3, and access them using industry standard file protocols. This removes the files from the on-premises NAS device and provides a method of directly mounting the file share for on-premises servers and clients.
> 
> ![Q24 storage gateway - file storage](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-04-55-d6065baf4b15649e450a4fb5cd195708.jpg)

### Q25
> An Amazon RDS PostgreSQL database is configured as Multi-AZ. A solutions architect needs to scale read performance and the solution must be configured for high availability.\
> What is the most cost-effective solution?

~~read replica in the same AZ as master~~

> **Explanation:**\
> You can create a read replica as a Multi-AZ DB instance. Amazon RDS creates a standby of your replica in another Availability Zone for failover support for the replica. Creating your read replica as a Multi-AZ DB instance is independent of whether the source database is a Multi-AZ DB instance.


### Q26

> A Solutions Architect needs to design a solution that will allow Website Developers to deploy static web content without managing server infrastructure. All web content must be accessed over HTTPS with a custom domain name. The solution should be scalable as the company continues to grow.\
> Which of the following will provide the MOST cost-effective solution?

~~S3 with static website~~

> **Explanation:**\
> You can create an Amazon CloudFront distribution that uses an S3 bucket as the origin. This will allow you to serve the static content using the HTTPS protocol.\
> To serve a static website hosted on Amazon S3, you can deploy a CloudFront distribution using one of these configurations:
> - Using a REST API endpoint as the origin with access restricted by an origin access identity (OAI).
> - Using a website endpoint as the origin with anonymous (public) access allowed.
> - Using a website endpoint as the origin with access restricted by a Referer header.


### Q27

> Some objects that are uploaded to Amazon S3 standard storage class are initially accessed frequently for a period of 30 days. Then, objects are infrequently accessed for up to 90 days. After that, the objects are no longer needed.\
> How should lifecycle management be configured?


~~STANDARD_IA after 30 days, after 90days move to glacier~~

> **Explanation:**\
> In this scenario we need to keep the objects in the STANDARD storage class for 30 days as the objects are being frequently accessed. We can configure a lifecycle action that then transitions the objects to INTELLIGENT_TIERING, STANDARD_IA, or ONEZONE_IA. After that we don’t need the objects so they can be expired.\
> All other options do not meet the stated requirements or are not supported lifecycle transitions. For example:
> - You cannot transition to REDUCED_REDUNDANCY from any storage class.
> - Transitioning from STANDARD_IA to ONEZONE_IA is possible but we do not want to keep the objects so it incurs unnecessary costs.
> - Transitioning to GLACIER is possible but again incurs unnecessary costs.


### Q28

> A company allows its developers to attach existing IAM policies to existing IAM roles to enable faster experimentation and agility. However, the security operations team is concerned that the developers could attach the existing administrator policy, which would allow the developers to circumvent any other security policies.\
> How should a solutions architect address this issue?

IAM permissions boundaries

> **Explanation:**\
> The permissions boundary for an IAM entity (user or role) sets the maximum permissions that the entity can have. This can change the effective permissions for that user or role. The effective permissions for an entity are the permissions that are granted by all the policies that affect the user or role. Within an account, the permissions for an entity can be affected by identity-based policies, resource-based policies, permissions boundaries, Organizations SCPs, or session policies.\
> Therefore, the solutions architect can set an IAM permissions boundary on the developer IAM role that explicitly denies attaching the administrator policy.
> ![Q28 permission Boundaries](https://img-c.udemycdn.com/redactor/raw/2020-05-18_13-16-47-d821741ed0b16e76b6ea87b337c0aeeb.JPG)


### Q29

> A Solutions Architect must select the most appropriate database service for two use cases. A team of data scientists perform complex queries on a data warehouse that take several hours to complete. Another team of scientists need to run fast, repeat queries and update dashboards for customer support staff.\
> Which solution delivers these requirements MOST cost-effectively?

~~redshift + elasticCache~~

> **Explanation:**\
> RedShift is a columnar data warehouse DB that is ideal for running long complex queries. RedShift can also improve performance for repeat queries by caching the result and returning the cached result when queries are re-run. Dashboard, visualization, and business intelligence (BI) tools that execute repeat queries see a significant boost in performance due to result caching.

### Q30

> A solutions architect is designing an application on AWS. The compute layer will run in parallel across EC2 instances. The compute layer should scale based on the number of jobs to be processed. The compute layer is stateless. The solutions architect must ensure that the application is loosely coupled and the job items are durably stored.\
> Which design should the solutions architect use?

**SQS and autosacling group based on sqs queue** (probably also lambda would work)


> **Explanation:**\
> In this case we need to find a durable and loosely coupled solution for storing jobs. Amazon SQS is ideal for this use case and can be configured to use dynamic scaling based on the number of jobs waiting in the queue.\
> To configure this scaling you can use the backlog per instance metric with the target value being the acceptable backlog per instance to maintain. You can calculate these numbers as follows:
> - Backlog per instance: To calculate your backlog per instance, start with the ApproximateNumberOfMessages queue attribute to determine the length of the SQS queue (number of messages available for retrieval from the queue). Divide that number by the fleet's running capacity, which for an Auto Scaling group is the number of instances in the InService state, to get the backlog per instance.
> - Acceptable backlog per instance: To calculate your target value, first determine what your application can accept in terms of latency. Then, take the acceptable latency value and divide it by the average time that an EC2 instance takes to process a message.
> 
> This solution will scale EC2 instances using Auto Scaling based on the number of jobs waiting in the SQS queue.

### Q31

> A solutions architect has created a new AWS account and must secure AWS account root user access.\
> Which combination of actions will accomplish this? (Select TWO.)

Strong password, MFA.

> **Explanation:**\
> There are several security best practices for securing the root user account:
> - Lock away root user access keys OR delete them if possible
> - Use a strong password
> - Enable multi-factor authentication (MFA)
> 
> The root user automatically has full privileges to the account and these privileges cannot be restricted so it is extremely important to follow best practice advice about securing the root user account.
### Q32

> A company runs an eCommerce application that uses an Amazon Aurora database. The database performs well except for short periods when monthly sales reports are run. A Solutions Architect has reviewed metrics in Amazon CloudWatch and found that the Read Ops and CPUUtilization metrics are spiking during the periods when the sales reports are run.\
> What is the MOST cost-effective solution to solve this performance issue?

aurora replica, use replica end point for reporting

> **Explanation:**\
> The simplest and most cost-effective option is to use an Aurora Replica. The replica can serve read operations which will mean the reporting application can run reports on the replica endpoint without causing any performance impact on the production database.

### Q33

> A website runs on Amazon EC2 instances behind an Application Load Balancer (ALB). The website has a mix of dynamic and static content. Customers around the world are reporting performance issues with the website.\
> Which set of actions will improve website performance for users worldwide?

~~use transit gateway?~~

> **Explanation:**\
> Amazon CloudFront is a content delivery network (CDN) that improves website performance by caching content at edge locations around the world. It can serve both dynamic and static content. This is the best solution for improving the performance of the website.
### Q34

> A solutions architect is designing a two-tier web application. The application consists of a public-facing web tier hosted on Amazon EC2 in public subnets. The database tier consists of Microsoft SQL Server running on Amazon EC2 in a private subnet. Security is a high priority for the company.\
> How should security groups be configured in this situation? (Select TWO.)

datbase tier: 1433 from web tier. web tier: inbound 443 from 0.0.0.0/0 (anywhere)

> **Explanation:**\
> In this scenario an inbound rule is required to allow traffic from any internet client to the web front end on SSL/TLS port 443. The source should therefore be set to 0.0.0.0/0 to allow any inbound traffic.\
> To secure the connection from the web frontend to the database tier, an outbound rule should be created from the public EC2 security group with a destination of the private EC2 security group. The port should be set to 1433 for MySQL. The private EC2 security group will also need to allow inbound traffic on 1433 from the public EC2 security group.
> 
> ![Q34 subnets security groups](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-31-06-6dcaf8d88c9ec27f837343a0ae2630f9.jpg)


### Q35

> A company runs a number of core enterprise applications in an on-premises data center. The data center is connected to an Amazon VPC using AWS Direct Connect. The company will be creating additional AWS accounts and these accounts will also need to be quickly, and cost-effectively connected to the on-premises data center in order to access the core applications.\
> What deployment changes should a Solutions Architect implement to meet these requirements with the LEAST operational overhead?

~~VPC endpoints in the direct connect VPC~~

> **Explanation:**\
> AWS Transit Gateway connects VPCs and on-premises networks through a central hub. With AWS Transit Gateway, you can quickly add Amazon VPCs, AWS accounts, VPN capacity, or AWS Direct Connect gateways to meet unexpected demand, without having to wrestle with complex connections or massive routing tables. This is the operationally least complex solution and is also cost-effective.
> 
> ![Q35 Transit Gateway](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_10-12-08-05293542d6a09dfb24fcf7e9a59f123e.jpg)

### Q36

> An application running video-editing software is using significant memory on an Amazon EC2 instance.How can a user track memory usage on the Amazon EC2 instance?

maybe cloud watch agent (is there such a thing?)

> **Explanation:**\
> There is no standard metric in CloudWatch for collecting EC2 memory usage. However, you can use the CloudWatch agent to collect both system metrics and log files from Amazon EC2 instances and on-premises servers. The metrics can be pushed to a CloudWatch custom metric.

### Q37

> A company uses a Microsoft Windows file share for storing documents and media files. Users access the share using Microsoft Windows clients and are authenticated using the company’s Active Directory. The chief information officer wants to move the data to AWS as they are approaching capacity limits. The existing user authentication and access management system should be used.\
> How can a Solutions Architect meet these requirements?


FSX for windows?

> **Explanation:**\
> Amazon FSx for Windows File Server makes it easy for you to launch and scale reliable, performant, and secure shared file storage for your applications and end users. With Amazon FSx, you can launch highly durable and available file systems that can span multiple availability zones (AZs) and can be accessed from up to thousands of compute instances using the industry-standard Server Message Block (SMB) protocol.\
> It provides a rich set of administrative and security features, and integrates with Microsoft Active Directory (AD). To serve a wide spectrum of workloads, Amazon FSx provides high levels of file system throughput and IOPS and consistent sub-millisecond latencies.\
> You can also mount FSx file systems from on-premises using a VPN or Direct Connect connection. This topology is depicted in the image below:
> 
> ![Q37 FSX windows file server](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-06-44-85272267dfd4d12b26b830545075dc49.jpg)

### Q38

> A web application that allows users to upload and share documents is running on a single Amazon EC2 instance with an Amazon EBS volume. To increase availability the architecture has been updated to use an Auto Scaling group of several instances across Availability Zones behind an Application Load Balancer. After the change users can only see a subset of the documents.\
> What is the BEST method for a solutions architect to modify the solution so users can see all documents?

copy to EFS

> **Explanation:**\
> The problem that is being described is that the users are uploading the documents to an individual EC2 instance with a local EBS volume. Therefore, as EBS volumes cannot be shared across AZs, the data is stored separately and the ALB will be distributing incoming connections to different instances / data sets.\
> The simple resolution is to implement a shared storage layer for the documents so that they can be stored in one place and seen by any user who connects no matter which instance they connect to.
> ![Q38 EFS](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-30-06-e730f4c928621d8e81593b31b92020f6.jpg)

### Q39

> You have created a file system using Amazon Elastic File System (EFS) which will hold home directories for users.\
> What else needs to be done to enable users to save files to the EFS file system?

subdirecotry for each user with permissions, then mount directly there.

> **Explanation:**\
> After creating a file system, by default, only the root user (UID 0) has read-write-execute permissions. For other users to modify the file system, the root user must explicitly grant them access.\
> One common use case is to create a “writable” subdirectory under this file system root for each user you create on the EC2 instance and mount it on the user’s home directory. All files and subdirectories the user creates in their home directory are then created on the Amazon EFS file system.

### Q40
> An automotive company plans to implement IoT sensors in manufacturing equipment that will send data to AWS in real time. The solution must receive events in an ordered manner from each asset and ensure that the data is saved for future processing.\
> Which solution would be MOST efficient?

kinesis, then save to S3

> **Explanation:**\
> Amazon Kinesis Data Streams is the ideal service for receiving streaming data. The Amazon Kinesis Client Library (KCL) delivers all records for a given partition key to the same record processor, making it easier to build multiple applications reading from the same Amazon Kinesis data stream. Therefore, a separate partition (rather than shard) should be used for each equipment asset.\
> Amazon Kinesis Firehose can be used to receive streaming data from Data Streams and then load the data into Amazon S3 for future processing.

### Q41
> A Solutions Architect is designing a web application that runs on Amazon EC2 instances behind an Elastic Load Balancer. All data in transit must be encrypted.\
> Which solution options meet the encryption requirement? (choose 2)

HTTPS listeners on ~~both NLB and~~ ALB

> **Explanation:**\
> You can passthrough encrypted traffic with an NLB and terminate the SSL on the EC2 instances, so this is a valid answer.\
> You can use a HTTPS listener with an ALB and install certificates on both the ALB and EC2 instances. This does not use passthrough, instead it will terminate the first SSL connection on the ALB and then re-encrypt the traffic and connect to the EC2 instances.

### Q42

> A large media site has multiple applications running on Amazon ECS. A Solutions Architect needs to use content metadata to route traffic to specific services.\
> What is the MOST efficient method to fulfil this requirement?

ALB with path based routing

> **Explanation:**\
> The ELB Application Load Balancer can route traffic based on data included in the request including the host name portion of the URL as well as the path in the URL. Creating a rule to route traffic based on information in the path will work for this solution and ALB works well with Amazon ECS.\
> The diagram below depicts a configuration where a listener directs traffic that comes in with /orders in the URL to the second target group and all other traffic to the first target group:
> 
> ![Q42 ALB](https://img-c.udemycdn.com/redactor/raw/2020-05-19_02-52-53-e64bfa586a07ebf08866a789f514b515.JPG)

### Q43
> An application upgrade caused some issues with stability. The application owner enabled logging and has generated a 5 GB log file in an Amazon S3 bucket. The log file must be securely shared with the application vendor to troubleshoot the issues.\
> What is the MOST secure way to share the log file?

presigned URLs.

> **Explanation:**\
> A presigned URL gives you access to the object identified in the URL. When you create a presigned URL, you must provide your security credentials and then specify a bucket name, an object key, an HTTP method (PUT for uploading objects), and an expiration date and time. The presigned URLs are valid only for the specified duration. That is, you must start the action before the expiration date and time.\
> This is the most secure way to provide the vendor with time-limited access to the log file in the S3 bucket.

### Q44
> A solutions architect is finalizing the architecture for a distributed database that will run across multiple Amazon EC2 instances. Data will be replicated across all instances so the loss of an instance will not cause loss of data. The database requires block storage with low latency and throughput that supports up to several million transactions per second per server.\
> Which storage solution should the solutions architect use?

instance storage


> **Explanation:**\
> An instance store provides temporary block-level storage for your instance. This storage is located on disks that are physically attached to the host computer. Instance store is ideal for temporary storage of information that changes frequently, such as buffers, caches, scratch data, and other temporary content, or for data that is replicated across a fleet of instances, such as a load-balanced pool of web servers.\
> Some instance types use NVMe or SATA-based solid state drives (SSD) to deliver high random I/O performance. This is a good option when you need storage with very low latency, but you don't need the data to persist when the instance terminates or you can take advantage of fault-tolerant architectures.\
> In this scenario the data is replicated and fault tolerant so the best option to provide the level of performance required is to use instance store volumes.
> 
> ![Q44 Instance Store](https://img-c.udemycdn.com/redactor/raw/2020-05-18_12-54-38-9844f13d68a10a257851c57f0f920791.JPG)

### Q45
> A company has divested a single business unit and needs to move the AWS account owned by the business unit to another AWS Organization.\
> How can this be achieved?

migrate with AWS organization Console


> **Explanation:**\
> Accounts can be migrated between organizations. To do this you must have root or IAM access to both the member and master accounts. Resources will remain under the control of the migrated account.


### Q46
> A web application is deployed in multiple regions behind an ELB Application Load Balancer. You need deterministic routing to the closest region and automatic failover. Traffic should traverse the AWS global network for consistent performance.\
> How can this be achieved?

~~Route 53 alias with latency based routing policies~~

> **Explanation:**\
> AWS Global Accelerator is a service that improves the availability and performance of applications with local or global users. You can configure the ALB as a target and Global Accelerator will automatically route users to the closest point of presence.\
> Failover is automatic and does not rely on any client side cache changes as the IP addresses for Global Accelerator are static anycast addresses. Global Accelerator also uses the AWS global network which ensures consistent performance.
> 
> ![Q46 route53 global accelerator](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-24-35-e7819bc00c77975464fae1690e0de51c.jpg)
### Q47

> An application is hosted on the U.S west coast. Users there have no problems, but users on the east coast are experiencing performance issues. The users have reported slow response times with the search bar autocomplete and display of account listings.\
> How can you improve the performance for users on the east coast?

ElasticCache in the us-east region

> **Explanation:**\
> ElastiCache can be deployed in the U.S east region to provide high-speed access to the content. ElastiCache Redis has a good use case for autocompletion (see links below).

### Q48

> An organization plans to deploy a higher performance computing (HPC) workload on AWS using Linux. The HPC workload will use many Amazon EC2 instances and will generate a large quantity of small output files that must be stored in persistent storage for future use.\
> A Solutions Architect must design a solution that will enable the EC2 instances to access data using native file system interfaces and to store output files in cost-effective long-term storage.\
> Which combination of AWS services meets these requirements?

FSX lustre and S3

> **Explanation:**\
> Amazon FSx for Lustre is ideal for high performance computing (HPC) workloads running on Linux instances. FSx for Lustre provides a native file system interface and works as any file system does with your Linux operating system.\
> When linked to an Amazon S3 bucket, FSx for Lustre transparently presents objects as files, allowing you to run your workload without managing data transfer from S3.\
> This solution provides all requirements as it enables Linux workloads to use the native file system interfaces and to use S3 for long-term and cost-effective storage of output files.
> 
> ![Q48 HPC](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_10-07-01-3904cc64f7ccc25ee4a538dcf782af68.jpg)


### Q49

> A company is planning to upload a large quantity of sensitive data to Amazon S3. The company’s security department require that the data is encrypted before it is uploaded.
> Which option meets these requirements?

client side encryption with ~~S3 managed encryption keys~~.

> **Explanation:**\
> The requirement is that the objects must be encrypted before they are uploaded. The only option presented that meets this requirement is to use client-side encryption. You then have two options for the keys you use to perform the encryption:
> - Use a customer master key (CMK) stored in AWS Key Management Service (AWS KMS).
> - Use a master key that you store within your application.
> 
> In this case the correct answer is to use an AWS KMS key. Note that you cannot use client-side encryption with keys managed by Amazon S3.
> 
> ![Q49 S3 encryption](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-00-48-9dad6f62c8e7c9414b51a788e346c78c.jpg)


### Q50

> A company runs an application on Amazon EC2 instances which requires access to sensitive data in an Amazon S3 bucket. All traffic between the EC2 instances and the S3 bucket must not traverse the internet and must use private IP addresses. Additionally, the bucket must only allow access from services in the VPC.\
> Which combination of actions should a Solutions Architect take to meet these requirements? (Select TWO.)


VPC endpoint to S3, bucket policy to restrict access to S3 endpoint.

> **Explanation:**\
> Private access to public services such as Amazon S3 can be achieved by creating a VPC endpoint in the VPC. For S3 this would be a gateway endpoint. The bucket policy can then be configured to restrict access to the S3 endpoint only which will ensure that only services originating from the VPC will be granted access.
> 
> ![Q50 VPC endpoint](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-15-18-8ca86e9211911c46e6ebba4877ad1dca.jpg)


### Q51

> An IoT sensor is being rolled out to thousands of a company’s existing customers. The sensors will stream high volumes of data each second to a central location. A solution must be designed to ingest and store the data for analytics. The solution must provide near-real time performance and millisecond responsiveness.\
> Which solution should a Solutions Architect recommend?

kinesis, lambda, ~~Redshift~~

> **Explanation:**\
> A Kinesis data stream is a set of shards. Each shard contains a sequence of data records. A consumer is an application that processes the data from a Kinesis data stream. You can map a Lambda function to a shared-throughput consumer (standard iterator), or to a dedicated-throughput consumer with enhanced fan-out.\
> Amazon DynamoDB is the best database for this use case as it supports near-real time performance and millisecond responsiveness.


### Q52

> A company requires a solution to allow customers to customize images that are stored in an online catalog. The image customization parameters will be sent in requests to Amazon API Gateway. The customized image will then be generated on-demand and can be accessed online.\
> The solutions architect requires a highly available solution. Which solution will be MOST cost-effective?

Lambda, store both original and manipulated, use cloudFront with S3 as origin

> **Explanation:**\
> All solutions presented are highly available. The key requirement that must be satisfied is that the solution should be cost-effective and you must choose the most cost-effective option.\
> Therefore, it’s best to eliminate services such as Amazon EC2 and ELB as these require ongoing costs even when they’re not used. Instead, a fully serverless solution should be used. AWS Lambda, Amazon S3 and CloudFront are the best services to use for these requirements.

### Q53

> A web application is being deployed on an Amazon ECS cluster using the Fargate launch type. The application is expected to receive a large volume of traffic initially. The company wishes to ensure that performance is good for the launch and that costs reduce as demand decreases.\
> What should a solutions architect recommend?

ECS service auto scaling and cloudWatch

> **Explanation:**\
> Amazon ECS uses the AWS Application Auto Scaling service to scales tasks. This is configured through Amazon ECS using Amazon ECS Service Auto Scaling.\
> A Target Tracking Scaling policy increases or decreases the number of tasks that your service runs based on a target value for a specific metric. For example, in the image below the tasks will be scaled when the average CPU breaches 80% (as reported by CloudWatch):

![Q53 - auto scaling](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_10-01-16-332b18084ef8decec084431b6cc2ec4c.jpg)

### Q54

> A company requires a solution for replicating data to AWS for disaster recovery. Currently, the company uses scripts to copy data from various sources to a Microsoft Windows file server in the on-premises data center. The company also requires that a small amount of recent files are accessible to administrators with low latency.\
> What should a Solutions Architect recommend to meet these requirements?

Storage Gateway file gateway

> **Explanation:**\
> The best solution here is to use an AWS Storage Gateway File Gateway virtual appliance in the on-premises data center. This can be accessed the same protocols as the existing Microsoft Windows File Server (SMB/CIFS). Therefore, the script simply needs to be updated to point to the gateway.\
> The file gateway will then store data on Amazon S3 and has a local cache for data that can be accessed at low latency. The file gateway provides an excellent method of enabling file protocol access to low cost S3 object storage.

![Q54 Filegatway](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-10-11-1743a212516d41e3c51c5f97fc6d09c5.jpg)


### Q55

> A company has refactored a legacy application to run as two microservices using Amazon ECS. The application processes data in two parts and the second part of the process takes longer than the first.\
> How can a solutions architect integrate the microservices and allow them to scale independently?

use SQS

> **Explanation:**\
> This is a good use case for Amazon SQS. The microservices must be decoupled so they can scale independently. An Amazon SQS queue will enable microservice 1 to add messages to the queue. Microservice 2 can then pick up the messages and process them. This ensures that if there’s a spike in traffic on the frontend, messages do not get lost due to the backend process not being ready to process them.


### Q56

> A company runs an internal browser-based application. The application runs on Amazon EC2 instances behind an Application Load Balancer. The instances run in an Amazon EC2 Auto Scaling group across multiple Availability Zones. The Auto Scaling group scales up to 20 instances during work hours, but scales down to 2 instances overnight. Staff are complaining that the application is very slow when the day begins, although it runs well by midmorning.\
> How should the scaling be changed to address the staff complaints and keep costs to a minimum?


~~set desired capacity to 20~~

> **Explanation:**\
> Though this sounds like a good use case for scheduled actions, both answers using scheduled actions will have 20 instances running regardless of actual demand. A better option to be more cost effective is to use a target tracking action that triggers at a lower CPU threshold.\
> With this solution the scaling will occur before the CPU utilization gets to a point where performance is affected. This will result in resolving the performance issues whilst minimizing costs. Using a reduced cooldown period will also more quickly terminate unneeded instances, further reducing costs.


### Q57

> A High Performance Computing (HPC) application will be migrated to AWS. The application requires low network latency and high throughput between nodes and will be deployed in a single AZ.\
> How should the application be deployed for best inter-node performance?

cluster placement group

> **Explanation:**\
> A cluster placement group provides low latency and high throughput for instances deployed in a single AZ. It is the best way to provide the performance required for this application.

### Q58

> A shared services VPC is being setup for use by several AWS accounts. An application needs to be securely shared from the shared services VPC. The solution should not allow consumers to connect to other instances in the VPC.\
> How can this be setup with the least administrative effort? (choose 2)

private link ~~and VPC peering~~

> **Explanation:**\
> VPCs can be shared among multiple AWS accounts. Resources can then be shared amongst those accounts. However, to restrict access so that consumers cannot connect to other instances in the VPC the best solution is to use PrivateLink to create an endpoint for the application. The endpoint type will be an interface endpoint and it uses an NLB in the shared services VPC.

![Q58 connecting VPC](https://img-c.udemycdn.com/redactor/raw/2020-05-25_01-23-21-8df694147b53b155434524cea44ac356.jpg)

### Q59

> A company runs several NFS file servers in an on-premises data center. The NFS servers must run periodic backups to Amazon S3 using automatic synchronization for small volumes of data.\
> Which solution meets these requirements and is MOST cost-effective?


DataSync agent to S3

> **Explanation:**\
> AWS DataSync is an online data transfer service that simplifies, automates, and accelerates copying large amounts of data between on-premises systems and AWS Storage services, as well as between AWS Storage services. DataSync can copy data between Network File System (NFS) shares, or Server Message Block (SMB) shares, self-managed object storage, AWS Snowcone, Amazon Simple Storage Service (Amazon S3) buckets, Amazon Elastic File System (Amazon EFS) file systems, and Amazon FSx for Windows File Server file systems.\
> This is the most cost-effective solution from the answer options available.


### Q60

> A highly sensitive application runs on Amazon EC2 instances using EBS volumes. The application stores data temporarily on Amazon EBS volumes during processing before saving results to an Amazon RDS database. The company’s security team mandate that the sensitive data must be encrypted at rest.\
> Which solution should a Solutions Architect recommend to meet this requirement?

encrypt EBS volumes and RDS with AWS KMS keys

> **Explanation:**\
> As the data is stored both in the EBS volumes (temporarily) and the RDS database, both the EBS and RDS volumes must be encrypted at rest. This can be achieved by enabling encryption at creation time of the volume and AWS KMS keys can be used to encrypt the data. This solution meets all requirements.


### Q61

> A website runs on a Microsoft Windows server in an on-premises data center. The web server is being migrated to Amazon EC2 Windows instances in multiple Availability Zones on AWS. The web server currently uses data stored in an on-premises network-attached storage (NAS) device.\
> Which replacement to the NAS file share is MOST resilient and durable?

~~migrate to amazon EBS~~

> **Explanation:**\
> Amazon FSx for Windows File Server provides fully managed, highly reliable file storage that is accessible over the industry-standard Server Message Block (SMB) protocol. It is built on Windows Server, delivering a wide range of administrative features such as user quotas, end-user file restore, and Microsoft Active Directory (AD) integration. It offers single-AZ and multi-AZ deployment options, fully managed backups, and encryption of data at rest and in transit.\
> This is the only solution presented that provides resilient storage for Windows instances.


![Q61 FSX](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-28-32-63d08e70f92de39652bd9053f98f90df.jpg)

### Q62

> A Solutions Architect must design a storage solution for incoming billing reports in CSV format. The data will be analyzed infrequently and discarded after 30 days.\
> Which combination of services will be MOST cost-effective in meeting these requirements?

write to S3 and use Athena

> **Explanation:**\
> Amazon S3 is great solution for storing objects such as this. You only pay for what you use and don’t need to worry about scaling as it will scale as much as you need it to.\
> Using Amazon Athena to analyze the data works well as it is a serverless service so it will be very cost-effective for use cases where the analysis is only happening infrequently.\
> You can also configure Amazon S3 to expire the objects after 30 days.

### Q63
> A company runs an application in an Amazon VPC that requires access to an Amazon Elastic Container Service (Amazon ECS) cluster that hosts an application in another VPC. The company’s security team requires that all traffic must not traverse the internet.\
> Which solution meets this requirement?


NLB in one VPC, AWS private link in another VPC

> **Explanation:**\
> The correct solution is to use AWS PrivateLink in a service provider model. In this configuration a network load balancer will be implemented in the service provider VPC (the one with the ECS cluster in this example), and a PrivateLink endpoint will be created in the consumer VPC (the one with the company’s application).

![Q63 private link](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_12-11-25-32f12f4c4d3d65de75af283af88992aa.jpg)

### Q64
> A company is deploying a new web application that will run on Amazon EC2 instances in an Auto Scaling group across multiple Availability Zones. The application requires a shared storage solution that offers strong consistency as the content will be regularly updated.\
> Which solution requires the LEAST amount of effort?

use EFS

> **Explanation:**\
> Amazon EFS is a fully-managed service that makes it easy to set up, scale, and cost-optimize file storage in the Amazon Cloud. EFS file systems are accessible to Amazon EC2 instances via a file system interface (using standard operating system file I/O APIs) and support full file system access semantics (such as strong consistency and file locking).\
> EFS is a good solution for when you need to attach a shared filesystem to multiple EC2 instances across multiple Availability Zones.


### Q65
> A web application has recently been launched on AWS. The architecture includes two tier with a web layer and a database layer. It has been identified that the web server layer may be vulnerable to cross-site scripting (XSS) attacks.\
> What should a solutions architect do to remediate the vulnerability?

use ALB and AWS WAF

> **Explanation:**\
> The AWS Web Application Firewall (WAF) is available on the Application Load Balancer (ALB). You can use AWS WAF directly on Application Load Balancers (both internal and external) in a VPC, to protect your websites and web services.\
> Attackers sometimes insert scripts into web requests in an effort to exploit vulnerabilities in web applications. You can create one or more cross-site scripting match conditions to identify the parts of web requests, such as the URI or the query string, that you want AWS WAF to inspect for possible malicious scripts.

</details>

##