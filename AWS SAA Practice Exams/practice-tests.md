<!--
// cSpell:ignore DFSR Referer SATA CIFS OCSP Mbps Gbps Caffe2 SDLC DFSN RTMP OIDC 2xlarge NACL
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

> **Explanation:**\
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

> **Explanation:**\
> You can use a Lambda function to process Amazon Simple Notification Service notifications. Amazon SNS supports Lambda functions as a target for messages sent to a topic. This solution decouples the Amazon EC2 application from Lambda and ensures the Lambda function is invoked.

### Q23
> A new application will run across multiple Amazon ECS tasks. Front-end application logic will process data and then pass that data to a back-end ECS task to perform further processing and write the data to a datastore. The Architect would like to reduce-interdependencies so failures do no impact other components.\
> Which solution should the Architect use?


**front pushes to SQS, backend polls**

> **Explanation:**\
> This is a good use case for Amazon SQS. SQS is a service that is used for decoupling applications, thus reducing interdependencies, through a message bus. The front-end application can place messages on the queue and the back-end can then poll the queue for new messages. Please remember that Amazon SQS is pull-based (polling) not push-based (use SNS for push-based).


### Q24

> A company wishes to restrict access to their Amazon DynamoDB table to specific, private source IP addresses from their VPC.\
> What should be done to secure access to the table?
> - Create an AWS VPN Connection the amazon DynamoDB endpoint
> - Create gateway VPC endpoint and add an entry to the route table
> - Create an interface VPC endpoint in the VPC with an elastic network interface
> - Create the amazon dynamoDB table in the VPC

~~maybe **vpc endpoint?**~~

> **Explanation:**\
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

> **Explanation:**\
> A custom origin can point to an on-premises server and CloudFront is able to cache content for dynamic websites. CloudFront can provide performance optimizations for custom origins even if they are running on on-premises servers. These include persistent TCP connections to the origin, SSL enhancements such as Session tickets and OCSP stapling.
> Additionally, connections are routed from the nearest Edge Location to the user across the AWS global network. If the on-premises server is connected via a Direct Connect (DX) link this can further improve performance.

### Q26
> A company plans to make an Amazon EC2 Linux instance unavailable outside of business hours to save costs. The instance is backed by an Amazon EBS volume. There is a requirement that the contents of the instance’s memory must be preserved when it is made unavailable.\
> How can a solutions architect meet these requirements?

**Hibernation**

> **Explanation:**\
> When you hibernate an instance, Amazon EC2 signals the operating system to perform hibernation (suspend-to-disk). Hibernation saves the contents from the instance memory (RAM) to your Amazon Elastic Block Store (Amazon EBS) root volume. Amazon EC2 persists the instance's EBS root volume and any attached EBS data volumes. When you start your instance:
> - The EBS root volume is restored to its previous state
> - The RAM contents are reloaded
> - The processes that were previously running on the instance are resumed
> - Previously attached data volumes are reattached and the instance retains its instance ID


### Q27

> The database tier of a web application is running on a Windows server on-premises. The database is a Microsoft SQL Server database. The application owner would like to migrate the database to an Amazon RDS instance.\
> How can the migration be executed with minimal administrative effort and downtime?

use **Database migration service**

> **Explanation:**\
> You can directly migrate Microsoft SQL Server from an on-premises server into Amazon RDS using the Microsoft SQL Server database engine. This can be achieved using the native Microsoft SQL Server tools, or using AWS DMS as depicted below:
> ![Q27 database migration](https://img-c.udemycdn.com/redactor/raw/2020-05-17_07-59-54-f1a0a1c8024b0af89fa52ac80a9bc55c.JPG)


### Q28
> An insurance company has a web application that serves users in the United Kingdom and Australia. The application includes a database tier using a MySQL database hosted in eu-west-2. The web tier runs from eu-west-2 and ap-southeast-2. Amazon Route 53 geoproximity routing is used to direct users to the closest web tier. It has been noted that Australian users receive slow response times to queries.\
> Which changes should be made to the database tier to improve performance?


maybe **move to aurora**?

> **Explanation:**\
> The issue here is latency with read queries being directed from Australia to UK which is great physical distance. A solution is required for improving read performance in Australia.\
> An Aurora global database consists of one primary AWS Region where your data is mastered, and up to five read-only, secondary AWS Regions. Aurora replicates data to the secondary AWS Regions with typical latency of under a second. You issue write operations directly to the primary DB instance in the primary AWS Region.\
> This solution will provide better performance for users in the Australia Region for queries. Writes must still take place in the UK Region but read performance will be greatly improved.
> 
> ![Q28 Aurora replicates](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-05-41-0e623951d5ed62017a22f93bb14c6c67.jpg)

### Q29
> A web application runs in public and private subnets. The application architecture consists of a web tier and database tier running on Amazon EC2 instances. Both tiers run in a single Availability Zone (AZ).\
> Which combination of steps should a solutions architect take to provide high availability for this architecture? (Select TWO.)

**multi az**

> **Explanation:**\
> To add high availability to this architecture both the web tier and database tier require changes. For the web tier an Auto Scaling group across multiple AZs with an ALB will ensure there are always instances running and traffic is being distributed to them.\
> The database tier should be migrated from the EC2 instances to Amazon RDS to take advantage of a managed database with Multi-AZ functionality. This will ensure that if there is an issue preventing access to the primary database a secondary database can take over.
### Q30
> A financial services company has a web application with an application tier running in the U.S and Europe. The database tier consists of a MySQL database running on Amazon EC2 in us-west-1. Users are directed to the closest application tier using Route 53 latency-based routing. The users in Europe have reported poor performance when running queries.\
> Which changes should a Solutions Architect make to the database tier to improve performance?

~~Read replicas in europe?~~


> **Explanation:**\
> Amazon Aurora Global Database is designed for globally distributed applications, allowing a single Amazon Aurora database to span multiple AWS regions. It replicates your data with no impact on database performance, enables fast local reads with low latency in each region, and provides disaster recovery from region-wide outages.\
> A global database can be configured in the European region and then the application tier in Europe will need to be configured to use the local database for reads/queries. The diagram below depicts an Aurora Global Database deployment.
> 
> ![Q30 Aurora Global Database](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-40-32-e096373484158f919b4e661ab7c1aa8d.jpg)


### Q31
> A manufacturing company captures data from machines running at customer sites. Currently, thousands of machines send data every 5 minutes, and this is expected to grow to hundreds of thousands of machines in the near future. The data is logged with the intent to be analyzed in the future as needed.\
> What is the SIMPLEST method to store this streaming data at scale?

**Kinesis firehose**

> **Explanation:**\
> Kinesis Data Firehose is the easiest way to load streaming data into data stores and analytics tools. It captures, transforms, and loads streaming data and you can deliver the data to “destinations” including Amazon S3 buckets for later analysis.


### Q32
> An Amazon RDS Read Replica is being deployed in a separate region. The master database is not encrypted but all data in the new region must be encrypted.\
> How can this be achieved?

~~KNS keys when creating the read replicas?~~

> **Explanation:**\
> You cannot create an encrypted Read Replica from an unencrypted master DB instance. You also cannot enable encryption after launch time for the master DB instance. Therefore, you must create a new master DB by taking a snapshot of the existing DB, encrypting it, and then creating the new DB from the snapshot. You can then create the encrypted cross-region Read Replica of the master DB.


### Q33
> A video production company is planning to move some of its workloads to the AWS Cloud. The company will require around 5 TB of storage for video processing with the maximum possible I/O performance. They also require over 400 TB of extremely durable storage for storing video files and 800 TB of storage for long-term archival.\
> Which combinations of services should a Solutions Architect use to meet these requirements?

Instance store for performance, S3 for storage, Glacier for archiving.

> **Explanation:**\
> The best I/O performance can be achieved by using instance store volumes for the video processing. This is safe to use for use cases where the data can be recreated from the source files so this is a good use case.
> For storing data durably Amazon S3 is a good fit as it provides 99.999999999% of durability. For archival the video files can then be moved to Amazon S3 Glacier which is a low cost storage option that is ideal for long-term archival.


### Q34
> A developer created an application that uses Amazon EC2 and an Amazon RDS MySQL database instance. The developer stored the database user name and password in a configuration file on the root EBS volume of the EC2 application instance. A Solutions Architect has been asked to design a more secure solution.\
> What should the Solutions Architect do to achieve this requirement?

IAM permissions on the EC2 instance.

> **Explanation:**\
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

> **Explanation:**\
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

> **Explanation:**\
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

> **Explanation:**\
> Scaling based on a schedule allows you to set your own scaling schedule for predictable load changes. To configure your Auto Scaling group to scale based on a schedule, you create a scheduled action. This is ideal for situations where you know when and for how long you are going to need the additional capacity.


### Q38
> A legacy tightly-coupled High Performance Computing (HPC) application will be migrated to AWS.\
> Which network adapter type should be used?
> - Elastic network adaptor (ENA)
> - Elastic Fabric adaptor (EFA)
> - Elastic Ip address
> - Elastic Network interface (ENI)

EFA

> **Explanation:**\
> An Elastic Fabric Adapter is an AWS Elastic Network Adapter (ENA) with added capabilities. The EFA lets you apply the scale, flexibility, and elasticity of the AWS Cloud to tightly-coupled HPC apps. It is ideal for tightly coupled app as it uses the Message Passing Interface (MPI).


### Q39
> A company's web application is using multiple Amazon EC2 Linux instances and storing data on Amazon EBS volumes. The company is looking for a solution to increase the resiliency of the application in case of a failure.\
> What should a solutions architect do to meet these requirements?

~~EC2 instances in each AZ with EBS volumes~~

> **Explanation:**\
> To increase the resiliency of the application the solutions architect can use Auto Scaling groups to launch and terminate instances across multiple availability zones based on demand. An application load balancer (ALB) can be used to direct traffic to the web application running on the EC2 instances.
> Lastly, the Amazon Elastic File System (EFS) can assist with increasing the resilience of the application by providing a shared file system that can be mounted by multiple EC2 instances from multiple availability zones.


### Q40
> A company's application is running on Amazon EC2 instances in a single Region. In the event of a disaster, a solutions architect needs to ensure that the resources can also be deployed to a second Region.\
> Which combination of actions should the solutions architect take to accomplish this? (Select TWO.)

make an AMI, copy to other region, launch there

> **Explanation:**\
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

> **Explanation:**\
> In this case AWS Lambda can perform the computation and store the data in an Amazon DynamoDB table. Lambda can scale concurrent executions to meet demand easily and DynamoDB is built for key-value data storage requirements and is also serverless and easily scalable. This is therefore a cost effective solution for unpredictable workloads.


### Q42
> A company has two accounts for perform testing and each account has a single VPC: VPC-TEST1 and VPC-TEST2. The operations team require a method of securely copying files between Amazon EC2 instances in these VPCs. The connectivity should not have any single points of failure or bandwidth constraints.\
> Which solution should a Solutions Architect recommend?

VPC peering

> **Explanation:**\
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

> **Explanation:**\
> S3 standard is the best choice in this scenario for a short term storage solution. In this case the size and number of logs is unknown and it would be difficult to fully assess the access patterns at this stage. Therefore, using S3 standard is best as it is cost-effective, provides immediate access, and there are no retrieval fees or minimum capacity charge per object.


### Q44
> A company is working with a strategic partner that has an application that must be able to send messages to one of the company’s Amazon SQS queues. The partner company has its own AWS account.\
> How can a Solutions Architect provide least privilege access to the partner?

update the SQS policy and grant permissions 

> **Explanation:**\
> Amazon SQS supports resource-based policies. The best way to grant the permissions using the principle of least privilege is to use a resource-based policy attached to the SQS queue that grants the partner company’s AWS account the sqs:SendMessage privilege.
> The following policy is an example of how this could be configured:
> 
> ![Q44 permission policies](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-02-25_10-46-16-3d75509bc08efb886557569410905bfe.jpg)


### Q45
> A company runs an application in a factory that has a small rack of physical compute resources. The application stores data on a network attached storage (NAS) device using the NFS protocol. The company requires a daily offsite backup of the application data.\
> Which solution can a Solutions Architect recommend to meet this requirement?

~~AWS Storgate gateway volume to replicate data to S3~~

> **Explanation:**\
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

> **Explanation:**\
> To protect the distribution rights of the content and ensure that users are directed to the appropriate AWS Region based on the location of the user, the geolocation routing policy can be used with Amazon Route 53.
> Geolocation routing lets you choose the resources that serve your traffic based on the geographic location of your users, meaning the location that DNS queries originate from.
> When you use geolocation routing, you can localize your content and present some or all of your website in the language of your users. You can also use geolocation routing to restrict distribution of content to only the locations in which you have distribution rights.


### Q47

> Amazon EC2 instances in a development environment run between 9am and 5pm Monday-Friday. Production instances run 24/7.\
> Which pricing models should be used? (choose 2)

reserved instances for production, on-demand capacity reserations for development.

> **Explanation:**\
> Capacity reservations have no commitment and can be created and canceled as needed. This is ideal for the development environment as it will ensure the capacity is available. There is no price advantage but none of the other options provide a price advantage whilst also ensuring capacity is available.\
> Reserved instances are a good choice for workloads that run continuously. This is a good option for the production environment.


### Q48
> An organization want to share regular updates about their charitable work using static webpages. The pages are expected to generate a large amount of views from around the world. The files are stored in an Amazon S3 bucket. A solutions architect has been asked to design an efficient and effective solution.\
> Which action should the solutions architect take to accomplish this?

CloudFront with S3 as origin.

> **Explanation:**\
> Amazon CloudFront can be used to cache the files in edge locations around the world and this will improve the performance of the webpages.\
> To serve a static website hosted on Amazon S3, you can deploy a CloudFront distribution using one of these configurations:
> - Using a REST API endpoint as the origin with access restricted by an origin access identity (OAI)
> - Using a website endpoint as the origin with anonymous (public) access allowed
> - Using a website endpoint as the origin with access restricted by a Referer header


### Q49
> An application running on an Amazon ECS container instance using the EC2 launch type needs permissions to write data to Amazon DynamoDB.\
> How can you assign these permissions only to the specific ECS task that is running the application?

IAM policy attached to the ~~container Instance~~

> **Explanation:**\
> To specify permissions for a specific task on Amazon ECS you should use IAM Roles for Tasks. The permissions policy can be applied to tasks when creating the task definition, or by using an IAM task role override using the AWS CLI or SDKs. The taskRoleArn parameter is used to specify the policy.

### Q50
> An Amazon VPC contains several Amazon EC2 instances. The instances need to make API calls to Amazon DynamoDB. A solutions architect needs to ensure that the API calls do not traverse the internet.
> How can this be accomplished? (Select TWO.)

~~ENI endpoint~~, route table entry

> **Explanation:**\
> Amazon DynamoDB and Amazon S3 support gateway endpoints, not interface endpoints. With a gateway endpoint you create the endpoint in the VPC, attach a policy allowing access to the service, and then specify the route table to create a route table entry in.
> 
> ![Q50 API calls not traversing the internet](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-00-45-ac665c89acb1641afb831f1eb795210e.jpg)


### Q51
> An application is being created that will use Amazon EC2 instances to generate and store data. Another set of EC2 instances will then analyze and modify the data. Storage requirements will be significant and will continue to grow over time. The application architects require a storage solution.\
> Which actions would meet these needs?

using EFS and all EC2 instances mount it.

> **Explanation:**\
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

> **Explanation:**\
> The S3 Intelligent-Tiering storage class is designed to optimize costs by automatically moving data to the most cost-effective access tier, without performance impact or operational overhead.
> It works by storing objects in two access tiers: one tier that is optimized for frequent access and another lower-cost tier that is optimized for infrequent access. This is an ideal use case for intelligent-tiering as the access patterns for the log files are not known.


### Q53
> A company is investigating methods to reduce the expenses associated with on-premises backup infrastructure. The Solutions Architect wants to reduce costs by eliminating the use of physical backup tapes. It is a requirement that existing backup applications and workflows should continue to function.\
> What should the Solutions Architect recommend?

AWS Storage gateway with virtual tape library (VTL)

> **Explanation:**\
> The AWS Storage Gateway Tape Gateway enables you to replace using physical tapes on premises with virtual tapes in AWS without changing existing backup workflows. Tape Gateway emulates physical tape libraries, removes the cost and complexity of managing physical tape infrastructure, and provides more durability than physical tapes.
> 
> ![Q53 Tape Gateway](https://img-c.udemycdn.com/redactor/raw/test_question_description/2020-11-27_09-45-15-153ae046c9838c5d880c4f0fce858218.jpg)


### Q54
> A company uses an Amazon RDS MySQL database instance to store customer order data. The security team have requested that SSL/TLS encryption in transit must be used for encrypting connections to the database from application servers. The data in the database is currently encrypted at rest using an AWS KMS key.
> How can a Solutions Architect enable encryption in transit?

~~use self-signed certificates~~

> **Explanation:**\
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

> **Explanation:**\
> Amazon FSx for Windows File Server provides fully managed, highly reliable file storage that is accessible over the industry-standard Server Message Block (SMB) protocol.\
> Amazon FSx is built on Windows Server and provides a rich set of administrative features that include end-user file restore, user quotas, and Access Control Lists (ACLs).\
> Additionally, Amazon FSX for Windows File Server supports Distributed File System Replication (DFSR) in Single-AZ deployments as can be seen in the feature comparison table below.
> 
> ![Q55 Fsx deployments](https://img-c.udemycdn.com/redactor/raw/2020-05-17_14-56-37-c7a7eff7f9da52f8dadafb160c3ac4c0.JPG)


### Q56
> Your company shares some HR videos stored in an Amazon S3 bucket via CloudFront. You need to restrict access to the private content so users coming from specific IP addresses can access the videos and ensure direct access via the Amazon S3 bucket is not possible.\
> How can this be achieved?

Cloudfront with signedURL, origin access policy then restricting S3 bucket.

> **Explanation:**\
> A signed URL includes additional information, for example, an expiration date and time, that gives you more control over access to your content. You can also specify the IP address or range of IP addresses of the users who can access your content.\
> If you use CloudFront signed URLs (or signed cookies) to limit access to files in your Amazon S3 bucket, you may also want to prevent users from directly accessing your S3 files by using Amazon S3 URLs. To achieve this you can create an origin access identity (OAI), which is a special CloudFront user, and associate the OAI with your distribution.\
> You can then change the permissions either on your Amazon S3 bucket or on the files in your bucket so that only the origin access identity has read permission (or read and download permission).
> ![Q56 S3 origin and cloudfront](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-44-34-2bf62dad00d3baac5aa2bf88745d793e.jpg)


### Q57
> An eCommerce company runs an application on Amazon EC2 instances in public and private subnets. The web application runs in a public subnet and the database runs in a private subnet. Both the public and private subnets are in a single Availability Zone.\
> Which combination of steps should a solutions architect take to provide high availability for this architecture? (Select TWO.)

EC2 autoscaling in multiple AZ, private and public subnets and multi-region RDS.

> **Explanation:**\
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

> **Explanation:**\
> AWS DataSync can be used to move large amounts of data online between on-premises storage and Amazon S3 or Amazon Elastic File System (Amazon EFS). DataSync eliminates or automatically handles many of these tasks, including scripting copy jobs, scheduling and monitoring transfers, validating data, and optimizing network utilization.\
> The source datastore can use file servers that communicate using the Server Message Block (SMB) protocol.

### Q59
> A company runs an application in an on-premises data center that collects environmental data from production machinery. The data consists of JSON files stored on network attached storage (NAS) and around 5 TB of data is collected each day. The company must upload this data to Amazon S3 where it can be processed by an analytics application. The data must be transferred securely.\
> Which solution offers the MOST reliable and time-efficient data transfer?

AWS DataSync over DirectConnect

> **Explanation:**\
> The most reliable and time-efficient solution that keeps the data secure is to use AWS DataSync and synchronize the data from the NAS device directly to Amazon S3. This should take place over an AWS Direct Connect connection to ensure reliability, speed, and security.\
> AWS DataSync can copy data between Network File System (NFS) shares, Server Message Block (SMB) shares, self-managed object storage, AWS Snowcone, Amazon Simple Storage Service (Amazon S3) buckets, Amazon Elastic File System (Amazon EFS) file systems, and Amazon FSx for Windows File Server file systems.


### Q60
> A Solutions Architect has been tasked with re-deploying an application running on AWS to enable high availability. The application processes messages that are received in an ActiveMQ queue running on a single Amazon EC2 instance. Messages are then processed by a consumer application running on Amazon EC2. After processing the messages the consumer application writes results to a MySQL database running on Amazon EC2.\
> Which architecture offers the highest availability and low operational complexity?

AmazonMQ accross two AZ, autosacling in multi-AZ and RDS in multi-AZ.

> **Explanation:**\
> The correct answer offers the highest availability as it includes Amazon MQ active/standby brokers across two AZs, an Auto Scaling group across two AZ,s and a Multi-AZ Amazon RDS MySQL database deployment.\
> This architecture not only offers the highest availability it is also operationally simple as it maximizes the usage of managed services.


### Q61
> A solutions architect is creating a system that will run analytics on financial data for several hours a night, 5 days a week. The analysis is expected to run for the same duration and cannot be interrupted once it is started. The system will be required for a minimum of 1 year.\
> What should the solutions architect configure to ensure the EC2 instances are available when they are needed?

~~Regional Reserved instances~~

> **Explanation:**\
> On-Demand Capacity Reservations enable you to reserve compute capacity for your Amazon EC2 instances in a specific Availability Zone for any duration. This gives you the ability to create and manage Capacity Reservations independently from the billing discounts offered by Savings Plans or Regional Reserved Instances.\
> By creating Capacity Reservations, you ensure that you always have access to EC2 capacity when you need it, for as long as you need it. You can create Capacity Reservations at any time, without entering a one-year or three-year term commitment, and the capacity is available immediately.\
> The table below shows the difference between capacity reservations and other options:
> 
> ![Q61 EC2 reservations](https://img-c.udemycdn.com/redactor/raw/test_question_description/2022-01-10_20-24-52-810be510ec31f342fd622453a1ebec94.png)


### Q62
> A company runs a large batch processing job at the end of every quarter. The processing job runs for 5 days and uses 15 Amazon EC2 instances. The processing must run uninterrupted for 5 hours per day. The company is investigating ways to reduce the cost of the batch processing job.\
> Which pricing model should the company choose?

~~Spot Instances~~ 

> **Explanation:**\
> Each EC2 instance runs for 5 hours a day for 5 days per quarter or 20 days per year. This is time duration is insufficient to warrant reserved instances as these require a commitment of a minimum of 1 year and the discounts would not outweigh the costs of having the reservations unused for a large percentage of time.\
> In this case, there are no options presented that can reduce the cost and therefore on-demand instances should be used.


### Q63
> An AWS Organization has an OU with multiple member accounts in it. The company needs to restrict the ability to launch only specific Amazon EC2 instance types.\
> How can this policy be applied across the accounts with the least effort?

~~use Resources Access Manager~~

> **Explanation:**\
> To apply the restrictions across multiple member accounts you must use a Service Control Policy (SCP) in the AWS Organization. The way you would do this is to create a deny rule that applies to anything that does not equal the specific instance type you want to allow.\
> The following architecture could be used to achieve this goal:
> 
> ![Q63 Service control policy](https://img-c.udemycdn.com/redactor/raw/2020-05-21_00-48-06-e73551726ea3baf15e8e687947948cba.jpg)


### Q64
> A company hosts a multiplayer game on AWS. The application uses Amazon EC2 instances in a single Availability Zone and users connect over Layer 4. Solutions Architect has been tasked with making the architecture highly available and also more cost-effective.\
> How can the solutions architect best meet these requirements? (Select TWO.)

Network Load Balancer, AutoScaling group in multiple AZ.

> **Explanation:**\
> The solutions architect must enable high availability for the architecture and ensure it is cost-effective. To enable high availability an Amazon EC2 Auto Scaling group should be created to add and remove instances across multiple availability zones.\
> In order to distribute the traffic to the instances the architecture should use a Network Load Balancer which operates at Layer 4. This architecture will also be cost-effective as the Auto Scaling group will ensure the right number of instances are running based on demand.


### Q65
> A company uses Docker containers for many application workloads in an on-premise data center. The company is planning to deploy containers to AWS and the chief architect has mandated that the same configuration and administrative tools must be used across all containerized environments. The company also wishes to remain cloud agnostic to safeguard mitigate the impact of future changes in cloud strategy.\
> How can a Solutions Architect design a managed solution that will align with open-source software?

use EKS

> **Explanation:**\
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

## Exam 3
<details>
<summary>
65 questions
</summary>

### Q01
> An application is deployed on multiple AWS regions and accessed from around the world. The application exposes static public IP addresses. Some users are experiencing poor performance when accessing the application over the Internet.\
> What should a solutions architect recommend to reduce internet latency?

~~use route53~~

> **Explanation:**\
> AWS Global Accelerator is a service in which you create accelerators to improve availability and performance of your applications for local and global users. Global Accelerator directs traffic to optimal endpoints over the AWS global network. This improves the availability and performance of your internet applications that are used by a global audience. Global Accelerator is a global service that supports endpoints in multiple AWS Regions, which are listed in the AWS Region Table.\
> By default, Global Accelerator provides you with two static IP addresses that you associate with your accelerator. (Or, instead of using the IP addresses that Global Accelerator provides, you can configure these entry points to be IPv4 addresses from your own IP address ranges that you bring to Global Accelerator.)\
> The static IP addresses are anycast from the AWS edge network and distribute incoming application traffic across multiple endpoint resources in multiple AWS Regions, which increases the availability of your applications. Endpoints can be Network Load Balancers, Application Load Balancers, EC2 instances, or Elastic IP addresses that are located in one AWS Region or multiple Regions.
> 
> ![Q1 global accelerator](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-57-03-6258b515dd5b68a3f0c87a51702006a3.jpg)


### Q02
> A company has experienced malicious traffic from some suspicious IP addresses. The security team discovered the requests are from different IP addresses under the same CIDR range.\
> What should a solutions architect recommend to the team?

inbound deny rule, acl, low number

> **Explanation:**\
> You can only create deny rules with network ACLs, it is not possible with security groups. Network ACLs process rules in order from the lowest numbered rules to the highest until they reach and allow or deny. The following table describes some of the differences between security groups and network ACLs:\
> Therefore, the solutions architect should add a deny rule in the inbound table of the network ACL with a lower rule number than other rules.
> 
> ![Q2 security groups vs ACL](https://img-c.udemycdn.com/redactor/raw/2020-05-21_02-02-30-acd31eb94885375c9562ead5a41a3639.jpg)


### Q03
> An application allows users to upload and download files. Files older than 2 years will be accessed less frequently. A solutions architect needs to ensure that the application can scale to any number of files while maintaining high availability and durability.\
> Which scalable solutions should the solutions architect recommend?

S3 with lifecycle policy.

> **Explanation:**\
> S3 Standard-IA is for data that is accessed less frequently, but requires rapid access when needed. S3 Standard-IA offers the high durability, high throughput, and low latency of S3 Standard, with a low per GB storage price and per GB retrieval fee. This combination of low cost and high performance make S3 Standard-IA ideal for long-term storage, backups, and as a data store for disaster recovery files.


### Q04
> An application has multiple components for receiving requests that must be processed and subsequently processing the requests. The company requires a solution for decoupling the application components. The application receives around 10,000 requests per day and requests can take up to 2 days to process. Requests that fail to process must be retained.\
> Which solution meets these requirements most efficiently?

SQS, dead letter queue

> **Explanation:**\
> The Amazon Simple Queue Service (SQS) is ideal for decoupling the application components. Standard queues can support up to 120,000 in flight messages and messages can be retained for up to 14 days in the queue.\
> To ensure the retention of requests (messages) that fail to process, a dead-letter queue can be configured. Messages that fail to process are sent to the dead-letter queue (based on the redrive policy) and can be subsequently dealt with.


### Q05
> A company needs to migrate a large quantity of data from an on-premises environment to Amazon S3. The company is connected via an AWS Direct Connect (DX) connection. The company requires a fully managed solution that will keep the data private and automate and accelerate the replication of the data to AWS storage services.\
> Which solution should a Solutions Architect recommend?

Datasync, probably vpc endpoint

> **Explanation:**\
> AWS DataSync can be used to automate and accelerate the replication of data to AWS storage services. Note that Storage Gateway is used for hybrid scenarios where servers need local access to data with various options for storing and synchronizing the data to AWS storage services. Storage Gateway does not accelerate replication of data.\
> To deploy DataSync an agent must be installed. Then a task must be configured to replicated data to AWS. The task requires a connection to a service endpoint. To keep the data private and send it across the DX connection, a VPC endpoint should be used.
> 
> ![Q5 data sync](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-04-32-b583aab06b8f4559436be3906fdb3054.jpg)


### Q06
> A Solutions Architect is designing a mobile application that will capture receipt images to track expenses. The Architect wants to store the images on Amazon S3. However, uploading the images through the web server will create too much traffic.\
> What is the MOST efficient method to store images from a mobile application on Amazon S3?

using presigned url

> **Explanation:**\
> Uploading using a pre-signed URL allows you to upload the object without having any AWS security credentials/permissions. Pre-signed URLs can be generated programmatically and anyone who receives a valid pre-signed URL can then programmatically upload an object. This solution bypasses the web server avoiding any performance bottlenecks.


### Q07
> An Architect needs to find a way to automatically and repeatably create many member accounts within an AWS Organization. The accounts also need to be moved into an OU and have VPCs and subnets created.\
> What is the best way to achieve this?

~~aws organization cli~~

> **Explanation:**\
> The best solution is to use a combination of scripts and AWS CloudFormation. You will also leverage the AWS Organizations API. This solution can provide all of the requirements.


### Q08
> A Solutions Architect needs a solution for hosting a website that will be used by a development team. The website contents will consist of HTML, CSS, client-side JavaScript, and images.\
> Which solution is MOST cost-effective?

this is a static website (no backend dynamic content), so S3 bucket? but maybe it's a development site, so paying having and AWS fargat will be cheaper?

> **Explanation:**\
> Amazon S3 can be used for hosting static websites and cannot be used for dynamic content. In this case the content is purely static with client-side code running. Therefore, an S3 static website will be the most cost-effective solution for hosting this website.


### Q09
> An application runs on Amazon EC2 instances across multiple Availability Zones. The instances run in an Amazon EC2 Auto Scaling group behind an Application Load Balancer. The application performs best when the CPU utilization of the EC2 instances is at or near 40%.\
> What should a solutions architect do to maintain the desired performance across all instances in the group?

~~simple sacling maybe?~~ target tracking scaling.

> **Explanation:**\
> With target tracking scaling policies, you select a scaling metric and set a target value. Amazon EC2 Auto Scaling creates and manages the CloudWatch alarms that trigger the scaling policy and calculates the scaling adjustment based on the metric and the target value.\
> The scaling policy adds or removes capacity as required to keep the metric at, or close to, the specified target value. In addition to keeping the metric close to the target value, a target tracking scaling policy also adjusts to the changes in the metric due to a changing load pattern.


### Q10
> A security officer requires that access to company financial reports is logged. The reports are stored in an Amazon S3 bucket. Additionally, any modifications to the log files must be detected.\
> Which actions should a solutions architect take?

cloud trail,data events?

> **Explanation:**\
> Amazon CloudTrail can be used to log activity on the reports. The key difference between the two answers that include CloudTrail is that one references data events whereas the other references management events.\
> Data events provide visibility into the resource operations performed on or within a resource. These are also known as data plane operations. Data events are often high-volume activities.\
> Example data events include:
> - Amazon S3 object-level API activity (for example, GetObject, DeleteObject, and PutObject API operations).
> -  AWS Lambda function execution activity (the Invoke API).
>
> Management events provide visibility into management operations that are performed on resources in your AWS account. These are also known as control plane operations. Example management events include:
> -  Configuring security (for example, IAM AttachRolePolicy API operations)
> - Registering devices (for example, Amazon EC2 CreateDefaultVpc API operations).
>
> Therefore, to log data about access to the S3 objects the solutions architect should log read and write data events.\
> Log file validation can also be enabled on the destination bucket:
> 
> ![q10 data event](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-53-27-925a2b271e999a9b3ac4717b47ed69a3.jpg)
> 
> ![q10 log validation](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-53-38-ab1c8f32e2aec2ce99d13ff85ff2a881.jpg)


### Q11
> A systems administrator of a company wants to detect and remediate the compromise of services such as Amazon EC2 instances and Amazon S3 buckets.\
> Which AWS service can the administrator use to protect the company against attacks?

~~probably inspector~~

> **Explanation:**\
> Amazon GuardDuty gives you access to built-in detection techniques that are developed and optimized for the cloud. The detection algorithms are maintained and continuously improved upon by AWS Security. The primary detection categories include reconnaissance, instance compromise, account compromise, and bucket compromise.
> Amazon GuardDuty offers HTTPS APIs, CLI tools, and Amazon CloudWatch Events to support automated security responses to security findings. For example, you can automate the response workflow by using CloudWatch Events as an event source to trigger an AWS Lambda function.


### Q12
> A gaming company collects real-time data and stores it in an on-premises database system. The company are migrating to AWS and need better performance for the database. A solutions architect has been asked to recommend an in-memory database that supports data replication.\
> Which database should a solutions architect recommend?

in memory means elasticCache, i think Redis has better performance

> **Explanation:**\
> Amazon ElastiCache is an in-memory database. With ElastiCache Memcached there is no data replication or high availability. As you can see in the diagram, each node is a separate partition of data:
> 
> ![Q12 memCached](https://img-c.udemycdn.com/redactor/raw/2020-05-21_02-01-20-6b625ae592ca124e03dd8f3ac8c2a94d.jpg)
> 
> Therefore, the Redis engine must be used which does support both data replication and clustering. The following diagram shows a Redis architecture with cluster mode enabled:
> 
> ![Q12 Redis](https://img-c.udemycdn.com/redactor/raw/2020-05-21_02-01-34-c19e5fe437e8f141a5fe0a655e0990bd.jpg)


### Q13
> An online store uses an Amazon Aurora database. The database is deployed as a Multi-AZ deployment. Recently, metrics have shown that database read requests are high and causing performance issues which result in latency for write requests.\
> What should the solutions architect do to separate the read requests from the write requests?

~~create read-replicas?~~

> **Explanation:**\
> Aurora Replicas are independent endpoints in an Aurora DB cluster, best used for scaling read operations and increasing availability. Up to 15 Aurora Replicas can be distributed across the Availability Zones that a DB cluster spans within an AWS Region.\
> The DB cluster volume is made up of multiple copies of the data for the DB cluster. However, the data in the cluster volume is represented as a single, logical volume to the primary instance and to Aurora Replicas in the DB cluster.\
> As well as providing scaling for reads, Aurora Replicas are also targets for multi-AZ. In this case the solutions architect can update the application to read from the Aurora Replica
> ![Q13 Aurora replicas](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-56-02-d92494e987b1b8e80c6782e078540b97.jpg)


### Q14
> A company hosts statistical data in an Amazon S3 bucket that users around the world download from their website using a URL that resolves to a domain name. The company needs to provide low latency access to users and plans to use Amazon Route 53 for hosting DNS records.\
> Which solution meets these requirements?

~~probably route53 CNAME record, S3 bucket as cloud front origin~~

> **Explanation:**\
> This is a simple requirement for low latency access to the contents of an Amazon S3 bucket for global users. The best solution here is to use Amazon CloudFront to cache the content in Edge Locations around the world. This involves creating a web distribution that points to an S3 origin (the bucket) and then create an Alias record in Route 53 that resolves the applications URL to the CloudFront distribution endpoint.


### Q15
> A company requires a high-performance file system that can be mounted on Amazon EC2 Windows instances and Amazon EC2 Linux instances. Applications running on the EC2 instances perform separate processing of the same files and the solution must provide a file system that can be mounted by all instances simultaneously.
> Which solution meets these requirements?

~~maybe EFS general purpose for both?~~

> **Explanation:**\
> Amazon FSx for Windows File Server provides a fully managed native Microsoft Windows file system so you can easily move your Windows-based applications that require shared file storage to AWS. You can easily connect Linux instances to the file system by installing the cifs-utils package. The Linux instances can then mount an SMB/CIFS file system.
> 
> ![Q15 FSx windows file server](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-13-16-ac3ea6d1bde80de577cab4ac0352a586.jpg)


### Q16
> An application runs on Amazon EC2 Linux instances. The application generates log files which are written using standard API calls. A storage solution is required that can be used to store the files indefinitely and must allow concurrent access to all files.\
> Which storage service meets these requirements and is the MOST cost-effective?

not instance store (ephemeral), not EBS (only one), ~~probably EFS because of the standard API calls~~

> **Explanation:**\
> The application is writing the files using API calls which means it will be compatible with Amazon S3 which uses a REST API. S3 is a massively scalable key-based object store that is well-suited to allowing concurrent access to the files from many instances.\
> Amazon S3 will also be the most cost-effective choice. A rough calculation using the AWS pricing calculator shows the cost differences between 1TB of storage on EBS, EFS, and S3 Standard.
> 
> ![Q16 storage pricing](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-55-11-6ae407730f9451342d76fb514c903b8a.jpg)


### Q17
> A Solutions Architect needs to monitor application logs and receive a notification whenever a specific number of occurrences of certain HTTP status code errors occur.\
> Which tool should the Architect use?

probably cloudwatch ~~metrics?~~

> **Explanation:**\
> You can use CloudWatch Logs to monitor applications and systems using log data. For example, CloudWatch Logs can track the number of errors that occur in your application logs and send you a notification whenever the rate of errors exceeds a threshold you specify. This is the best tool for this requirement.

### Q18
> A solutions architect is designing a high performance computing (HPC) application using Amazon EC2 Linux instances. All EC2 instances need to communicate to each other with low latency and high throughput network performance.\
> Which EC2 solution BEST meets these requirements?

probably one AZ cluster placement group.

> **Explanation:**\
> When you launch a new EC2 instance, the EC2 service attempts to place the instance in such a way that all of your instances are spread out across underlying hardware to minimize correlated failures. You can use placement groups to influence the placement of a group of interdependent instances to meet the needs of your workload. Depending on the type of workload, you can create a placement group using one of the following placement strategies:
> - Cluster – packs instances close together inside an Availability Zone. This strategy enables workloads to achieve the low-latency network performance necessary for tightly-coupled node-to-node communication that is typical of HPC applications.
> - Partition – spreads your instances across logical partitions such that groups of instances in one partition do not share the underlying hardware with groups of instances in different partitions. This strategy is typically used by large distributed and replicated workloads, such as Hadoop, Cassandra, and Kafka.
> - Spread – strictly places a small group of instances across distinct underlying hardware to reduce correlated failures.
> 
> For this scenario, a cluster placement group should be used as this is the best option for providing low-latency network performance for a HPC application.


### Q19
> A production application runs on an Amazon RDS MySQL DB instance. A solutions architect is building a new reporting tool that will access the same data. The reporting tool must be highly available and not impact the performance of the production application.\
> How can this be achieved?

highly available means multiple AZ, 

> **Explanation:**\
> You can create a read replica as a Multi-AZ DB instance. Amazon RDS creates a standby of your replica in another Availability Zone for failover support for the replica. Creating your read replica as a Multi-AZ DB instance is independent of whether the source database is a Multi-AZ DB instance.


### Q20
> An application on Amazon Elastic Container Service (ECS) performs data processing in two parts. The second part takes much longer to complete.\
> How can an Architect decouple the data processing from the backend application component?

SQS

> **Explanation:**\
> Processing each part using a separate ECS task may not be essential but means you can separate the processing of the data. An Amazon Simple Queue Service (SQS) is used for decoupling applications. It is a message queue on which you place messages for processing by application components. In this case you can process each data processing part in separate ECS tasks and have them write an Amazon SQS queue. That way the backend can pick up the messages from the queue when they’re ready and there is no delay due to the second part not being complete.

### Q21
> A High Performance Computing (HPC) application needs storage that can provide 135,000 IOPS. The storage layer is replicated across all instances in a cluster.\
> What is the optimal storage solution that provides the required performance and is cost-effective?

replicated means that it's not shared, so it doesn't need to persistent.

> **Explanation:**\
> Instance stores offer very high performance and low latency. As long as you can afford to lose an instance, i.e. you are replicating your data, these can be a good solution for high performance/low latency requirements. Also, the cost of instance stores is included in the instance charges so it can also be more cost-effective than EBS Provisioned IOPS.


### Q22
> An application runs on a fleet of Amazon EC2 instances in an Amazon EC2 Auto Scaling group behind an Elastic Load Balancer. The operations team has determined that the application performs best when the CPU utilization of the EC2 instances is at or near 60%.\
> Which scaling configuration should a Solutions Architect use to optimize the applications performance?

~~again, maybe simple scaling?~~

> **Explanation:**\
> With target tracking scaling policies, you select a scaling metric and set a target value. Amazon EC2 Auto Scaling creates and manages the CloudWatch alarms that trigger the scaling policy and calculates the scaling adjustment based on the metric and the target value.\
> The scaling policy adds or removes capacity as required to keep the metric at, or close to, the specified target value. In addition to keeping the metric close to the target value, a target tracking scaling policy also adjusts to changes in the metric due to a changing load pattern.\
> The following diagram shows a target tracking policy set to keep the CPU utilization of the EC2 instances at or close to 60%.
> ![Q22 dynamic scaling - target tracking policy](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-24-12-01662811c941eb690ffe2f963ddac8c5.jpg)


### Q23
> A company has several AWS accounts that are used by developers for development, testing and pre-production environments. The company has received large bills for Amazon EC2 instances that are underutilized. A Solutions Architect has been tasked with restricting the ability to launch large EC2 instances in all accounts.\
> How can the Solutions Architect meet this requirement with the LEAST operational overhead?

probably AWS organization service control policy.

> **Explanation:**\
> Service control policies (SCPs) are a type of organization policy that you can use to manage permissions in your organization. SCPs offer central control over the maximum available permissions for all accounts in your organization. An SCP defines a guardrail, or sets limits, on the actions that the account's administrator can delegate to the IAM users and roles in the affected accounts.\
> In this case the Solutions Architect can use an SCP to define a restriction that denies the launch of large EC2 instances. The SCP can be applied to all accounts, and this will ensure that even those users with permissions to launch EC2 instances will be restricted to smaller EC2 instance types.


### Q24
> An e-commerce web application needs a highly scalable key-value database. Which AWS database service should be used?

dynamoDB

> **Explanation:**\
> A key-value database is a type of nonrelational (NoSQL) database that uses a simple key-value method to store data. A key-value database stores data as a collection of key-value pairs in which a key serves as a unique identifier. Amazon DynamoDB is a fully managed NoSQL database service that provides fast and predictable performance with seamless scalability – this is the best database for these requirements.


### Q25
> An organization is extending a secure development environment into AWS. They have already secured the VPC including removing the Internet Gateway and setting up a Direct Connect connection.\
> What else needs to be done to add encryption?

~~maybe border gateway protocol?~~

> **Explanation:**\
> A VPG is used to setup an AWS VPN which you can use in combination with Direct Connect to encrypt all data that traverses the Direct Connect link. This combination provides an IPsec-encrypted private connection that also reduces network costs, increases bandwidth throughput, and provides a more consistent network experience than internet-based VPN connections.


### Q26
> A company runs a financial application using an Amazon EC2 Auto Scaling group behind an Application Load Balancer (ALB). When running month-end reports on a specific day and time each month the application becomes unacceptably slow. Amazon CloudWatch metrics show the CPU utilization hitting 100%.\
> What should a solutions architect recommend to ensure the application is able to handle the workload and avoid downtime?

auto scaling with schedule

> **Explanation:**\
> Scheduled scaling allows you to set your own scaling schedule. In this case the scaling action can be scheduled to occur just prior to the time that the reports will be run each month. Scaling actions are performed automatically as a function of time and date. This will ensure that there are enough EC2 instances to serve the demand and prevent the application from slowing down.


### Q27
> Health related data in Amazon S3 needs to be frequently accessed for up to 90 days. After that time the data must be retained for compliance reasons for seven years and is rarely accessed.\
> Which storage classes should be used?

standard then deep archive

> **Explanation:**\
> In this case the data is frequently accessed so must be stored in standard for the first 90 days. After that the data is still to be kept for compliance reasons but is rarely accessed so is a good use case for DEEP_ARCHIVE.


### Q28
> A company operates a production web application that uses an Amazon RDS MySQL database. The database has automated, non-encrypted daily backups. To increase the security of the data, it has been recommended that encryption should be enabled for backups. Unencrypted backups will be destroyed after the first encrypted backup has been completed.\
> What should be done to enable encryption for future backups

create a snapshot, encrypt the snapshot, reload from it.

> **Explanation:**\
> Amazon RDS uses snapshots for backup. Snapshots are encrypted when created only if the database is encrypted and you can only select encryption for the database when you first create it. In this case the database, and hence the snapshots, ad unencrypted.\
> However, you can create an encrypted copy of a snapshot. You can restore using that snapshot which creates a new DB instance that has encryption enabled. From that point on encryption will be enabled for all snapshots.


### Q29
> A Solutions Architect created the following policy and associated to an AWS IAM group containing several administrative users:
> ```json
> {
>    "Version": "2012-10-17",
>     "Statement": [
>     {
>         "Effect": "Allow",
>         "Action": "ec2:TerminateInstances",
>         "Resource": "*",
>         "Condition": {
>             "IpAddress": {
>                 "aws:SourceIp": "10.1.2.0/24"
>             }
>         }
>     },
>     {
>         "Effect": "Deny",
>         "Action": "ec2:*",
>         "Resource": "*",
>         "Condition": {
>             "StringNotEquals": {
>                 "ec2:Region": "us-east-1"
>             }
>         }
>     }
>     ] 
> }
> ```
> What is the effect of this policy?

terminate ec2 instances from the ip,only in us-east-1

> **Explanation:**\
> The Condition element (or Condition block) lets you specify conditions for when a policy is in effect. The Condition element is optional. In the Condition element, you build expressions in which you use condition operators (equal, less than, etc.) to match the condition keys and values in the policy against keys and values in the request context.\
> In this policy statement the first block allows the "ec2:TerminateInstances" API action only if the IP address of the requester is within the "10.1.2.0/24" range. This is specified using the "aws:SourceIp" condition.\
> The second block denies all EC2 API actions with a conditional operator (StringNotEquals) that checks the Region the request is being made in ("ec2:Region"). If the Region is any value other than us-east-1 the request will be denied. If the Region the request is being made in is us-east-1 the request will not be denied.


### Q30
> A high-performance file system is required for a financial modelling application. The data set will be stored on Amazon S3 and the storage solution must have seamless integration so objects can be accessed as files.\
> Which storage solution should be used?

high performance means, so maybe lustre?

> **Explanation:**\
> Amazon FSx for Lustre provides a high-performance file system optimized for fast processing of workloads such as machine learning, high performance computing (HPC), video processing, financial modeling, and electronic design automation (EDA). Amazon FSx works natively with Amazon S3, letting you transparently access your S3 objects as files on Amazon FSx to run analyses for hours to months.
> 
> ![Q30 FSX lustre](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-50-43-b59732a7c4ed02f0d7d405bca509ede2.jpg)


### Q31
> A web application in a three-tier architecture runs on a fleet of Amazon EC2 instances. Performance issues have been reported and investigations point to insufficient swap space. The operations team requires monitoring to determine if this is correct.\
> What should a solutions architect recommend?

detailed monitoring is more granular time, not metrics? swap utilization isn't a default metric, so we need an agent?

> **Explanation:**\
> You can use the CloudWatch agent to collect both system metrics and log files from Amazon EC2 instances and on-premises servers. The agent supports both Windows Server and Linux, and enables you to select the metrics to be collected, including sub-resource metrics such as per-CPU core.\
> There is now a unified agent and previously there were monitoring scripts. Both of these tools can capture SwapUtilization metrics and send them to CloudWatch. This is the best way to get memory utilization metrics from Amazon EC2 instances.


### Q32
> A new application will be launched on an Amazon EC2 instance with an Elastic Block Store (EBS) volume. A solutions architect needs to determine the most cost-effective storage option. The application will have infrequent usage, with peaks of traffic for a couple of hours in the morning and evening. Disk I/O is variable with peaks of up to 3,000 IOPS.\
> Which solution should the solutions architect recommend?

~~maybe throughput optimized HDD1?~~

> **Explanation:**\
> General Purpose SSD (gp2) volumes offer cost-effective storage that is ideal for a broad range of workloads. These volumes deliver single-digit millisecond latencies and the ability to burst to 3,000 IOPS for extended periods of time.\
> Between a minimum of 100 IOPS (at 33.33 GiB and below) and a maximum of 16,000 IOPS (at 5,334 GiB and above), baseline performance scales linearly at 3 IOPS per GiB of volume size. AWS designs gp2 volumes to deliver their provisioned performance 99% of the time. A gp2 volume can range in size from 1 GiB to 16 TiB.\
> In this case the volume would have a baseline performance of 3 x 200 = 600 IOPS. The volume could also burst to 3,000 IOPS for extended periods. As the I/O varies, this should be suitable.


### Q33
> A Solutions Architect has been tasked with migrating 30 TB of data from an on-premises data center within 20 days. The company has an internet connection that is limited to 25 Mbps and the data transfer cannot use more than 50% of the connection speed.\
> What should a Solutions Architect do to meet these requirements?

snowball (one is enough, 80 terra)

> **Explanation:**\
> This is a simple case of working out roughly how long it will take to migrate the data using the 12.5 Mbps of bandwidth that is available for transfer and seeing which options are feasible. Transferring 30 TB of data across a 25 Mbps connection could take upwards of 200 days.\
> Therefore, we know that using the Internet connection will not meet the requirements and we can rule out any solution that will use the internet (all options except for Snowball). AWS Snowball is a physical device that is shipped to your office or data center. You can then load data onto it and ship it back to AWS where the data is uploaded to Amazon S3.\
> Snowball is the only solution that will achieve the data migration requirements within the 20-day period.


### Q34
> A web application is running on a fleet of Amazon EC2 instances using an Auto Scaling Group. It is desired that the CPU usage in the fleet is kept at 40%.\
> How should scaling be configured?

~~again, simple scaling?~~ or maybe target tracking?

> **Explanation:**\
> This is a perfect use case for a target tracking scaling policy. With target tracking scaling policies, you select a scaling metric and set a target value. In this case you can just set the target value to 40% average aggregate CPU utilization.


### Q35
> A security team wants to limit access to specific services or actions in all of the team's AWS accounts. All accounts belong to a large organization in AWS Organizations. The solution must be scalable and there must be a single point where permissions can be maintained.\
> What should a solutions architect do to accomplish this?

probably AWS organization with service control policy

> **Explanation:**\
> Service control policies (SCPs) offer central control over the maximum available permissions for all accounts in your organization, allowing you to ensure your accounts stay within your organization’s access control guidelines.\
> SCPs alone are not sufficient for allowing access in the accounts in your organization. Attaching an SCP to an AWS Organizations entity (root, OU, or account) defines a guardrail for what actions the principals can perform. You still need to attach identity-based or resource-based policies to principals or resources in your organization's accounts to actually grant permissions to them.
> 
> ![Q35 Service Control Policies](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-57-51-f9ea69c694961184b290662661f770a4.jpg)


### Q36
> A company has created a disaster recovery solution for an application that runs behind an Application Load Balancer (ALB). The DR solution consists of a second copy of the application running behind a second ALB in another Region. The Solutions Architect requires a method of automatically updating the DNS record to point to the ALB in the second Region.\
> What action should the Solutions Architect take?

route53 health check

> **Explanation:**\
> Amazon Route 53 health checks monitor the health and performance of your web applications, web servers, and other resources. Each health check that you create can monitor one of the following:
> - The health of a specified resource, such as a web server
> - The status of other health checks
> - The status of an Amazon CloudWatch alarm
> 
> Health checks can be used with other configurations such as a failover routing policy. In this case a failover routing policy will direct traffic to the ALB of the primary Region unless health checks fail at which time it will direct traffic to the secondary record for the DR ALB.


### Q37
> A company has deployed an application that consists of several microservices running on Amazon EC2 instances behind an Amazon API Gateway API. A Solutions Architect is concerned that the microservices are not designed to elastically scale when large increases in demand occur.\
> Which solution addresses this concern?

none of the answers are perfect, i guess that SQS is a way to mitigate issues where we can't use autoscaling groups...

> **Explanation:**\
> The individual microservices are not designed to scale. Therefore, the best way to ensure they are not overwhelmed by requests is to decouple the requests from the microservices. An Amazon SQS queue can be created, and the API Gateway can be configured to add incoming requests to the queue. The microservices can then pick up the requests from the queue when they are ready to process them.


### Q38
> A Kinesis consumer application is reading at a slower rate than expected. It has been identified that multiple consumer applications have total reads exceeding the per-shard limits.\
> How can this situation be resolved?

no idea.... more shards? ~~api throteling?~~

> **Explanation:**\
> One shard provides a capacity of 1MB/sec data input and 2MB/sec data output. One shard can support up to 1000 PUT records per second. The total capacity of the stream is the sum of the capacities of its shards.\
> In a case where multiple consumer applications have total reads exceeding the per-shard limits, you need to increase the number of shards in the Kinesis data stream.


### Q39
> A Solutions Architect must design a solution to allow many Amazon EC2 instances across multiple subnets to access a shared data store. The data must be accessed by all instances simultaneously and access should use the NFS protocol. The solution must also be highly scalable and easy to implement.\
> Which solution best meets these requirements?

NFS means EFS.

> **Explanation:**\
> The Amazon Elastic File System (EFS) is a perfect solution for this requirement. Amazon EFS filesystems are accessed using the NFS protocol and can be mounted by many instances across multiple subnets simultaneously. EFS filesystems are highly scalable and very easy to implement.
> 
> ![Q39 EFS overview](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-08-11-6d39438584e2d604d169e4d38f77f00e.jpg)


### Q40
> A company has deployed an API in a VPC behind an internal Network Load Balancer (NLB). An application that consumes the API as a client is deployed in a second account in private subnets.\
> Which architectural configurations will allow the API to be consumed without using the public Internet? (Select TWO.)

Direct connect (DX) is for on-premises access. so not this.\ 
maybe **VPC peering** and **privateLink**

> **Explanation:**\
> You can create your own application in your VPC and configure it as an AWS PrivateLink-powered service (referred to as an endpoint service). Other AWS principals can create a connection from their VPC to your endpoint service using an interface VPC endpoint. You are the service provider, and the AWS principals that create connections to your service are service consumers.\
> This configuration is powered by AWS PrivateLink and clients do not need to use an internet gateway, NAT device, VPN connection or AWS Direct Connect connection, nor do they require public IP addresses.\
> Another option is to use a VPC Peering connection. A VPC peering connection is a networking connection between two VPCs that enables you to route traffic between them using private IPv4 addresses or IPv6 addresses. Instances in either VPC can communicate with each other as if they are within the same network. You can create a VPC peering connection between your own VPCs, or with a VPC in another AWS account.
> 
> ![Q40 api gateway service](https://img-c.udemycdn.com/redactor/raw/2020-05-18_13-30-56-6e12fed09086b40d3a54c6810873d3c2.JPG)


### Q41
> A company is planning to use Amazon S3 to store documents uploaded by its customers. The images must be encrypted at rest in Amazon S3. The company does not want to spend time managing and rotating the keys, but it does want to control who can access those keys.\
> What should a solutions architect use to accomplish this?

I think SSE-S3 keys aren't managed (they are part of the upload/download process), so maybe SSE-KMS?

> **Explanation:**\
> SSE-KMS requires that AWS manage the data key but you manage the customer master key (CMK) in AWS KMS. You can choose a customer managed CMK or the AWS managed CMK for Amazon S3 in your account.\
> Customer managed CMKs are CMKs in your AWS account that you create, own, and manage. You have full control over these CMKs, including establishing and maintaining their key policies, IAM policies, and grants, enabling and disabling them, rotating their cryptographic material, adding tags, creating aliases that refer to the CMK, and scheduling the CMKs for deletion.\
> For this scenario, the solutions architect should use SSE-KMS with a customer managed CMK. That way KMS will manage the data key but the company can configure key policies defining who can access the keys.


### Q42
> A company runs a business-critical application in the us-east-1 Region. The application uses an Amazon Aurora MySQL database cluster which is 2 TB in size. A Solutions Architect needs to determine a disaster recovery strategy for failover to the us-west-2 Region. The strategy must provide a recovery time objective (RTO) of 10 minutes and a recovery point objective (RPO) of 5 minutes.\
> Which strategy will meet these requirements?

- RTO - time to recover.
- RPO - maximal time lost.

~~maybe multi-region aurora, route53 healthchecks?~~

> **Explanation:**\
> Amazon Aurora Global Database is designed for globally distributed applications, allowing a single Amazon Aurora database to span multiple AWS regions. It replicates your data with no impact on database performance, enables fast local reads with low latency in each region, and provides disaster recovery from region-wide outages.\
> If your primary region suffers a performance degradation or outage, you can promote one of the secondary regions to take read/write responsibilities. An Aurora cluster can recover in less than 1 minute even in the event of a complete regional outage.\
> This provides your application with an effective Recovery Point Objective (RPO) of 1 second and a Recovery Time Objective (RTO) of less than 1 minute, providing a strong foundation for a global business continuity plan.
> 
> ![Q2 aurora global database](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-22-11-2d4e0f2ec825327cf3bcbdbe7905c475.jpg)


### Q43
> An application runs on-premises and produces data that must be stored in a locally accessible file system that servers can mount using the NFS protocol. The data must be subsequently analyzed by Amazon EC2 instances in the AWS Cloud.\
> How can these requirements be met?

on premises, NFS - this means storage gateway with file gateway.

> **Explanation:**\
> The best solution for this requirement is to use an AWS Storage Gateway file gateway. This will provide a local NFS mount point for the data and a local cache. The data is then replicated to Amazon S3 where it can be analyzed by the Amazon EC2 instances in the AWS Cloud.
> ![Q43 file gatway](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-19-32-45d0ef2489b4ea2fb6fae5b15c844aa0.jpg)


### Q44
> A Solutions Architect is designing a solution for an application that requires very low latency between the client and the backend. The application uses the UDP protocol, and the backend is hosted on Amazon EC2 instances.\
> The solution must be highly available across multiple Regions and users around the world should be directed to the most appropriate Region based on performance.\
> How can the Solutions Architect meet these requirements?

WAF is security, so not this. maybe globabl accelerator? maybe route53 multivalue answers?

> **Explanation:**\
> An NLB is ideal for latency-sensitive applications and can listen on UDP for incoming requests. As Elastic Load Balancers are region-specific it is necessary to have an NLB in each Region in front of the EC2 instances.\
> To direct traffic based on optimal performance, AWS Global Accelerator can be used. GA will ensure traffic is routed across the AWS global network to the most optimal endpoint based on performance.
> ![Q44 global accelerator](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-37-03-06957cedd24132e515e0704262ea64fe.jpg)


### Q45
> A company is deploying a solution for sharing media files around the world using Amazon CloudFront with an Amazon S3 origin configured as a static website. The company requires that all traffic for the website must be inspected by AWS WAF.\
> Which solution meets these requirements?

probably AWS-WAF on the cloudfront, but not sure.

> **Explanation:**\
> The AWS Web Application Firewall (WAF) can be attached to an Amazon CloudFront distribution to enable protection from web exploits. In this case the distribution uses an S3 origin, and the question is stating that all traffic must be inspected by AWS WAF. This means we need to ensure that requests cannot circumvent AWS WAF and hit the S3 bucket directly.\
> This can be achieved by configuring an origin access identity (OAI) which is a special type of CloudFront user that is created within the distribution and configured in an S3 bucket policy. The policy will only allow requests that come from the OAI which means all requests must come via the distribution and cannot hit S3 directly.
> ![Q45 Origin Access Identity](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-32-57-0cec77f550d1e2e6100046094949925b.jpg)


### Q46
> A company has a Production VPC and a Pre-Production VPC. The Production VPC uses VPNs through a customer gateway to connect to a single device in an on-premises data center. The Pre-Production VPC uses a virtual private gateway attached to two AWS Direct Connect (DX) connections. Both VPCs are connected using a single VPC peering connection.\
> How can a Solutions Architect improve this architecture to remove any single point of failure?

- vpc to vpc - with peering - seems ok
- production vpc - customer gateway on-premises- single device (seems to be the issue).
- pre-production vpc - two DX connections with virtual private gateway - probably ok.

probably add another device to the production vpc.

> **Explanation:**\
> The only single point of failure in this architecture is the customer gateway device in the on-premises data center. A customer gateway device is the on-premises (client) side of the connection into the VPC. The customer gateway configuration is created within AWS, but the actual device is a physical or virtual device running in the on-premises data center. If this device is a single device, then if it fails the VPN connections will fail. The AWS side of the VPN link is the virtual private gateway, and this is a redundant device.


### Q47
> An application requires a MySQL database which will only be used several times a week for short periods. The database needs to provide automatic instantiation and scaling.\
> Which database service is most suitable?

probably aurora serverless? (aren't they all serverless)

> **Explanation:**\
> Amazon Aurora Serverless is an on-demand, auto-scaling configuration for Amazon Aurora. The database automatically starts up, shuts down, and scales capacity up or down based on application needs. This is an ideal database solution for infrequently-used applications.


### Q48
> An application runs on Amazon EC2 instances backed by Amazon EBS volumes and an Amazon RDS database. The application is highly sensitive and security compliance requirements mandate that all personally identifiable information (PII) be encrypted at rest.\
> Which solution should a Solutions Architect choose to this requirement?

~~maybe encrypted RDS with macie?~~

> **Explanation:**\
> The data must be encrypted at rest on both the EC2 instance’s attached EBS volumes and the RDS database. Both storage locations can be encrypted using AWS KMS keys. With RDS, KMS uses a customer master key (CMK) to encrypt the DB instance, all logs, backups, and snapshots.


### Q49
> A solutions architect is designing a microservices architecture. AWS Lambda will store data in an Amazon DynamoDB table named Orders. The solutions architect needs to apply an IAM policy to the Lambda function’s execution role to allow it to put, update, and delete items in the Orders table. No other actions should be allowed.\
> Which of the following code snippets should be included in the IAM policy to fulfill this requirement whilst providing the LEAST privileged access?

explicit table name in resource, explicit methods in action.

> **Explanation:**\
> The key requirements are to allow the Lambda function the put, update, and delete actions on a single table. Using the principle of least privilege the answer should not allow any other access.


### Q50
> A company requires a fully managed replacement for an on-premises storage service. The company’s employees often work remotely from various locations. The solution should also be easily accessible to systems connected to the on-premises environment.\
> Which solution meets these requirements?

~~maybe storage gateway (volume) with S3?~~

> **Explanation:**\
> Amazon FSx for Windows File Server (Amazon FSx) is a fully managed, highly available, and scalable file storage solution built on Windows Server that uses the Server Message Block (SMB) protocol. It allows for Microsoft Active Directory integration, data deduplication, and fully managed backups, among other critical enterprise features.
> An Amazon FSx file system can be created to host the file shares. Clients can then be connected to an AWS Client VPN endpoint and gateway to enable remote access. The protocol used in this solution will be SMB.


### Q51
> A company has some statistical data stored in an Amazon RDS database. The company wants to allow users to access this information using an API. A solutions architect must create a solution that allows sporadic access to the data, ranging from no requests to large bursts of traffic.\
> Which solution should the solutions architect suggest?

probably lambda

> **Explanation:**\
> This question is simply asking you to work out the best compute service for the stated requirements. The key requirements are that the compute service should be suitable for a workload that can range quite broadly in demand from no requests to large bursts of traffic.
> AWS Lambda is an ideal solution as you pay only when requests are made and it can easily scale to accommodate the large bursts in traffic. Lambda works well with both API Gateway and Amazon RDS.


### Q52
> An application consists of a web tier in a public subnet and a MySQL cluster hosted on Amazon EC2 instances in a private subnet. The MySQL instances must retrieve product data from a third-party provider over the internet. A Solutions Architect must determine a strategy to enable this access with maximum security and minimum operational overhead.\
> What should the Solutions Architect do to meet these requirements?

~~probably internet gateway~~ 

> **Explanation:**\
> The MySQL clusters instances need to access a service on the internet. The most secure method of enabling this access with low operational overhead is to create a NAT gateway. When deploying a NAT gateway, the gateway itself should be deployed in a public subnet whilst the route table in the private subnet must be updated to point traffic to the NAT gateway ID.\
> The configuration can be seen in the diagram below:
> 
> ![Q52 NAT gateway](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-02-45-d46799d3b11819240a38a0c28fcd0171.jpg)


### Q53
> A company has created an application that stores sales performance data in an Amazon DynamoDB table. A web application is being created to display the data.\
> A Solutions Architect must design the web application using managed services that require minimal operational maintenance.\
> Which architectures meet these requirements? (Select TWO.)

EC2 isn't managed.
API gateway REST API to lambdas. maybe route53?

> **Explanation:**\
> There are two architectures here that fulfill the requirement to create a web application that displays the data from the DynamoDB table.\
> The first one is to use an API Gateway REST API that invokes an AWS Lambda function. A Lambda proxy integration can be used, and this will proxy the API requests to the Lambda function which processes the request and accesses the DynamoDB table.\
> The second option is to use an API Gateway REST API to directly access the sales performance data. In this case a proxy for the DynamoDB query API can be created using a method in the REST API.


### Q54
> A company has created a duplicate of its environment in another AWS Region. The application is running in warm standby mode. There is an Application Load Balancer (ALB) in front of the application. Currently, failover is manual and requires updating a DNS alias record to point to the secondary ALB.\
> How can a solutions architect automate the failover process?

route53 health check?

> **Explanation:**\
> You can use Route 53 to check the health of your resources and only return healthy resources in response to DNS queries. There are three types of DNS failover configurations:
> 1. Active-passive: Route 53 actively returns a primary resource. In case of failure, Route 53 returns the backup resource. Configured using a failover policy.
> 1. Active-active: Route 53 actively returns more than one resource. In case of failure, Route 53 fails back to the healthy resource. Configured using any routing policy besides failover.
> 1. Combination: Multiple routing policies (such as latency-based, weighted, etc.) are combined into a tree to configure more complex DNS failover.
>
> In this case an alias already exists for the secondary ALB. Therefore, the solutions architect just needs to enable a failover configuration with an Amazon Route 53 health check.\
> The configuration would look something like this:
> ![Q54 Route53 health checks](https://img-c.udemycdn.com/redactor/raw/2020-05-21_02-03-25-80669d24056e76a9d92f8db887b73ccc.jpg)


### Q55
> A web app allows users to upload images for viewing online. The compute layer that processes the images is behind an Auto Scaling group.\
> The processing layer should be decoupled from the front end and the ASG needs to dynamically adjust based on the number of images being uploaded.\
> How can this be achieved?

SQS and auto scaling based on number of messages.

> **Explanation:**\
> The best solution is to use Amazon SQS to decouple the front end from the processing compute layer. To do this you can create a custom CloudWatch metric that measures the number of messages in the queue and then configure the ASG to scale using a target tracking policy that tracks a certain value.


### Q56
> A company’s staff connect from home office locations to administer applications using bastion hosts in a single AWS Region. The company requires a resilient bastion host architecture that requires minimal ongoing operational overhead.\
> How can a Solutions Architect best meet these requirements?

maybe autoscaling group, multi AZ and load balancer?

> **Explanation:**\
> Bastion hosts (aka “jump hosts”) are EC2 instances in public subnets that administrators and operations staff can connect to from the internet. From the bastion host they are then able to connect to other instances and applications within AWS by using internal routing within the VPC.\
> All answers use a Network Load Balancer which is acceptable for forwarding incoming connections to targets. The differences are in where the connections are forwarded to. The best option is to create an Auto Scaling group with EC2 instances in multiple Availability Zones. This creates a resilient architecture within a single AWS Region which is exactly what the question asks for.


### Q57
> An application generates unique files that are returned to customers after they submit requests to the application. The application uses an Amazon CloudFront distribution for sending the files to customers. The company wishes to reduce data transfer costs without modifying the application.\
> How can this be accomplished?

~~maybe S3 transfer acceleration?~~

> **Explanation:**\
> Lambda@Edge is a feature of Amazon CloudFront that lets you run code closer to users of your application, which improves performance and reduces latency. Lambda@Edge runs code in response to events generated by the Amazon CloudFront.\
> You simply upload your code to AWS Lambda, and it takes care of everything required to run and scale your code with high availability at an AWS location closest to your end user.\
> In this case Lambda@Edge can compress the files before they are sent to users which will reduce data egress costs.


### Q58
> An application runs on Amazon EC2 instances in a private subnet. The EC2 instances process data that is stored in an Amazon S3 bucket. The data is highly confidential and a private and secure connection is required between the EC2 instances and the S3 bucket.\
> Which solution meets these requirements?

probably vpc endpoint.

> **Explanation:**\
> A gateway VPC endpoint can be used to access an Amazon S3 bucket using private IP addresses. To further secure the solution an S3 bucket policy can be created that restricts access to the VPC endpoint so connections cannot be made to the bucket from other sources.


### Q59
> A company is deploying an analytics application on AWS Fargate. The application requires connected storage that offers concurrent access to files and high performance.\
> Which storage option should the solutions architect recommend?

~~high performance means lustre?~~

> **Explanation:**\
> The Amazon Elastic File System offers concurrent access to a shared file system and provides high performance. You can create file system policies for controlling access and then use an IAM role that is specified in the policy for access.

### Q60
> A company is creating a solution that must offer disaster recovery across multiple AWS Regions. The solution requires a relational database that can support a Recovery Point Objective (RPO) of 1 second and a Recovery Time Objective (RTO) of 1 minute.\
> Which AWS solution can achieve this?

maybe aurora global database?

> **Explanation:**\
> Aurora Global Database lets you easily scale database reads across the world and place your applications close to your users. Your applications enjoy quick data access regardless of the number and location of secondary regions, with typical cross-region replication latencies below 1 second.\
> If your primary region suffers a performance degradation or outage, you can promote one of the secondary regions to take read/write responsibilities. An Aurora cluster can recover in less than 1 minute even in the event of a complete regional outage. This provides your application with an effective Recovery Point Objective (RPO) of 1 second and a Recovery Time Objective (RTO) of less than 1 minute, providing a strong foundation for a global business continuity plan.


### Q61
> A company is deploying an application that produces data that must be processed in the order it is received. The company requires a solution for decoupling the event data from the processing layer. The solution must minimize operational overhead.\
> How can a Solutions Architect meet these requirements?

SQS FIFO

> **Explanation:**\
> Amazon SQS can be used to decouple this application using a FIFO queue. With a FIFO queue the order in which messages are sent and received is strictly preserved. You can configure an AWS Lambda function to poll the queue, or you can configure a Lambda function as a destination to asynchronously process messages from the queue.
> ![Q61 SQS types](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-05-18_05-17-48-037e094b09540ca2dd15db19aec25f8a.jpg)


### Q62
> A company is planning to migrate a large quantity of important data to Amazon S3. The data will be uploaded to a versioning enabled bucket in the us-west-1 Region. The solution needs to include replication of the data to another Region for disaster recovery purposes.\
> How should a solutions architect configure the replication?

create a S3 bucket with versioning, enable cross region replication

> **Explanation:**\
> Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can copy objects between different AWS Regions or within the same Region. Both source and destination buckets must have versioning enabled.


### Q63
> A company runs a containerized application on an Amazon Elastic Kubernetes Service (EKS) using a microservices architecture. The company requires a solution to collect, aggregate, and summarize metrics and logs.\
> The solution should provide a centralized dashboard for viewing information including CPU and memory utilization for EKS namespaces, services, and pods.\
> Which solution meets these requirements?

maybe cloudTrail container insights in EKS?

> **Explanation:**\
> Use CloudWatch Container Insights to collect, aggregate, and summarize metrics and logs from your containerized applications and microservices. Container Insights is available for Amazon Elastic Container Service (Amazon ECS), Amazon Elastic Kubernetes Service (Amazon EKS), and Kubernetes platforms on Amazon EC2.\
> With Container Insights for EKS you can see the top contributors by memory or CPU, or the most recently active resources. This is available when you select any of the following dashboards in the drop-down box near the top of the page:
> - ECS Services
> - ECS Tasks
> - EKS Namespaces
> - EKS Services
> - EKS Pods


### Q64
> You need to scale read operations for your Amazon Aurora DB within a region. To increase availability you also need to be able to failover if the primary instance fails.\
> What should you implement?

~~globabl database?~~

> **Explanation:**\
> Aurora Replicas are independent endpoints in an Aurora DB cluster, best used for scaling read operations and increasing availability. Up to 15 Aurora Replicas can be distributed across the Availability Zones that a DB cluster spans within an AWS Region. To increase availability, you can use Aurora Replicas as failover targets. That is, if the primary instance fails, an Aurora Replica is promoted to the primary instance.\
> The graphic below provides an overview of Aurora Replicas:
> ![Q64 aurora replicas](https://img-c.udemycdn.com/redactor/raw/2020-05-21_01-45-23-f3373b15a644aff433005158b5840cc7.jpg)


### Q65
> A Solutions Architect is designing an application that will run on Amazon EC2 instances. The application will use Amazon S3 for storing image files and an Amazon DynamoDB table for storing customer information. The security team require that traffic between the EC2 instances and AWS services must not traverse the public internet.\
> How can the Solutions Architect meet the security team’s requirements?

~~interface vpc endpoints?~~

> **Explanation:**\
> A VPC endpoint enables private connections between your VPC and supported AWS services and VPC endpoint services powered by AWS PrivateLink. A gateway endpoint is used for Amazon S3 and Amazon DynamoDB. You specify a gateway endpoint as a route table target for traffic that is destined for the supported AWS services.


</details>

## Exam 4

<details>
<summary>
65 questions
</summary>

### Q01
> A company has 200 TB of video files stored in an on-premises data center that must be moved to the AWS Cloud within the next four weeks. The company has around 50 Mbps of available bandwidth on an Internet connection for performing the transfer.\
> What is the MOST cost-effective solution for moving the data within the required timeframe?


probably several snowballs devices (3 will be enough), maybe a snowmobile.

> **Explanation:**\
> To move 200 TB of data over a 50 Mbps link would take over 300 days. Therefore, the solution must avoid the Internet link. The most cost-effective solution is to use multiple AWS Snowball devices to migrate the data to AWS.\
> Snowball devices are shipped to your data center where you can load the data and then ship it back to AWS. This avoids the Internet connection and utilizes local high-bandwidth network connections to load the data.


### Q02
> An application runs on Amazon EC2 instances. The application reads data from Amazon S3, performs processing on the data, and then writes the results to an Amazon DynamoDB table.\
> The application writes many temporary files during the data processing. The application requires a high-performance storage solution for the temporary files.\
> What would be the fastest storage option for this solution?

temporary data means ephemeral, so instance store. RAID 0 probably makes things a bit faster.

> **Explanation:**\
> As the data is only temporary it can be stored on an instance store volume which is a volume that is physically attached to the host computer on which the EC2 instance is running.\
> To increase aggregate IOPS, or to improve sequential disk throughput, multiple instance store volumes can be grouped together using RAID 0 (disk striping) software. This can improve the aggregate performance of the volume.


### Q03
> An application runs across a fleet of Amazon EC2 instances and uses a shared file system hosted on Amazon EFS. The file system is used for storing many files that are generated by the application. The files are only accessed for the first few days after creation but must be retained.\
> How can a Solutions Architect optimize storage costs for the application?

lifecycle policy on EFS? ~~move to S3 IA?~~

> **Explanation:**\
> The solution uses Amazon EFS, and the files are only accessed for a few days. To reduce storage costs the Solutions Architect can configure the AFTER_7_DAYS lifecycle policy to transition the files to the IA storage class 7 days after the files are last accessed.\
> You define when Amazon EFS transitions files an IA storage class by setting a lifecycle policy. A file system has one lifecycle policy that applies to the entire file system. If a file is not accessed for the period of time defined by the lifecycle policy that you choose, Amazon EFS transitions the file to the IA storage class that is applicable to your file system.

### Q04
> A Solutions Architect is considering the best approach to enabling Internet access for EC2 instances in a private subnet.\
> What advantages do NAT Gateways have over NAT Instances? (choose 2)

AWS managed, highly available

> **Explanation:**\
> NAT gateways are managed for you by AWS. NAT gateways are highly available in each AZ into which they are deployed. They are not associated with any security groups and can scale automatically up to 45Gbps\
> NAT instances are managed by you. They must be scaled manually and do not provide HA. NAT Instances can be used as bastion hosts and can be assigned to security groups.\
> ![Q4 Nat instance ans Nat Gatway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-48-49-c6531613d95f7e3a86c1cb11ae7798c8.png)


### Q05
> A Solutions Architect is designing a web-facing application. The application will run on Amazon EC2 instances behind Elastic Load Balancers in multiple regions in an active/passive configuration.\
> The website address the application runs on is example.com. AWS Route 53 will be used to perform DNS resolution for the application.\
> How should the Solutions Architect configure AWS Route 53 in this scenario based on AWS best practices? (choose 2)

active passive means failover policy. probably alias record.

> **Explanation:**\
> The failover routing policy is used for active/passive configurations. Alias records can be used to map the domain apex (example.com) to the Elastic Load Balancers.\
> Failover routing lets you route traffic to a resource when the resource is healthy or to a different resource when the first resource is unhealthy. The primary and secondary records can route traffic to anything from an Amazon S3 bucket that is configured as a website to a complex tree of records.


### Q06
> A company is deploying an Amazon ElastiCache for Redis cluster. To enhance security a password should be required to access the database.\
> What should the solutions architect use?

Probably Redis AUTH, this isn't an AWS thing.

> **Explanation:**\
> Redis authentication tokens enable Redis to require a token (password) before allowing clients to execute commands, thereby improving data security.\
> You can require that users enter a token on a token-protected Redis server. To do this, include the parameter --auth-token (API: AuthToken) with the correct token when you create your replication group or cluster. Also include it in all subsequent commands to the replication group or cluster.


### Q07
> A company is migrating an eCommerce application into the AWS Cloud. The application uses an SQL database, and the database will be migrated to Amazon RDS.\
> A Solutions Architect has been asked to recommend a method to attain sub-millisecond responses to common read requests.\
> What should the solutions architect recommend?

DynamoDB accelerator isn't for RDS, IOPS won't help much (EBS storage), common read-requests means caching, so ElasticCache.

> **Explanation:**\
> Amazon ElastiCache is a fully managed in-memory data store and cache service. ElastiCache can be used to cache requests to an Amazon RDS database through application configuration. This can greatly improve performance as ElastiCache can return responses to queries with sub-millisecond latency.


### Q08
> A company has over 200 TB of log files in an Amazon S3 bucket. The company must process the files using a Linux-based software application that will extract and summarize data from the log files and store the output in a separate Amazon S3 bucket.\
> The company needs to minimize data transfer charges associated with the processing of this data.\
> How can a Solutions Architect meet these requirements?

do everything within the same region, never let data leave AWS. not sure if EC2 access buckets in the same region without VPC end point, maybe the lambda cos is too large... guessing the EC2.\
maybe S3 requires S3 endpoint, so there's also that.

> **Explanation:**\
> The software application must be installed on a Linux operating system so we must use Amazon EC2 or an on-premises VM. To avoid data charges however, we must ensure that the data does not egress the AWS Region. The best solution to avoid the egress data charges is to use an Amazon EC2 instance in the same Region as the S3 bucket that contains the log files. The processed output files must also be stored in a bucket in the same Region to avoid any data going out from EC2 to another Region.


### Q09
> A Solutions Architect works for a systems integrator running a platform that stores medical records. The government security policy mandates that patient data that contains personally identifiable information (PII) must be encrypted at all times, both at rest and in transit. Amazon S3 is used to back up data into the AWS cloud.\
> How can the Solutions Architect ensure the medical records are properly secured? (choose 2)

SSE-S3 with AES-256, maybe prior to upload use personal keys?

> **Explanation:**\
> When data is stored in an encrypted state it is referred to as encrypted “at rest” and when it is encrypted as it is being transferred over a network it is referred to as encrypted “in transit”. You can securely upload/download your data to Amazon S3 via SSL endpoints using the HTTPS protocol (In Transit – SSL/TLS).\
> You have the option of encrypting the data locally before it is uploaded or uploading using SSL/TLS so it is secure in transit and encrypting on the Amazon S3 side using S3 managed keys. The S3 managed keys will be AES-256 (not AES-128) bit keys.\
> Uploading data using CloudFront with an EC2 origin or using an encrypted EBS volume attached to an EC2 instance is not a solution to this problem as your company wants to backup these records onto S3 (not EC2/EBS).


### Q10
> A company is testing a new web application that runs on Amazon EC2 instances. A Solutions Architect is performing load testing and must be able to analyze the performance of the web application with a granularity of 1 minute.\
> What should the Solutions Architect do to meet this requirement?

cloudWatch detailed monitoring. performance means metrics.

> **Explanation:**\
> By default, your instance is enabled for basic monitoring. You can optionally enable detailed monitoring. After you enable detailed monitoring, the Amazon EC2 console displays monitoring graphs with a 1-minute period for the instance.\
> The following describes the data interval and charge for basic and detailed monitoring for instances:
> 
> ![Q10 - cloudWatch Monitoring](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-06-02_23-23-28-82670aaf7831d936e2f4a75866c1a1a1.jpg)


### Q11
> Every time an item in an Amazon DynamoDB table is modified a record must be retained for compliance reasons.\
> What is the most efficient solution to recording this information?

~~cloudTrails, ? ~~

> **Explanation:**\
> Amazon DynamoDB Streams captures a time-ordered sequence of item-level modifications in any DynamoDB table and stores this information in a log for up to 24 hours. Applications can access this log and view the data items as they appeared before and after they were modified, in near-real time.\
> For example, in the diagram below a DynamoDB stream is being consumed by a Lambda function which processes the item data and records a record in CloudWatch Logs:
>
> ![Q11 DynamoDb Streams](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-55-43-04c370b6a009c1594a77bd12b9499c3d.png)


### Q12
> A Solutions Architect needs to design a solution for providing a shared file system for company users in the AWS Cloud. The solution must be fault tolerant and should integrate with the company’s Microsoft Active Directory for access control.\
> Which storage solution meets these requirements?

interfation with AD means FSx windows.

> **Explanation:**\
> Amazon FSx for Windows File Server provides fully managed Microsoft Windows file servers, backed by a fully native Windows file system. Multi-AZ file systems provide high availability and failover support across multiple Availability Zones by provisioning and maintaining a standby file server in a separate Availability Zone within an AWS Region.\
> Amazon FSx works with Microsoft Active Directory (AD) to integrate with your existing Microsoft Windows environments. Active Directory is the Microsoft directory service used to store information about objects on the network and make this information easy for administrators and users to find and use.


### Q13
> A company runs a legacy application that uses an Amazon RDS MySQL database without encryption. The security team has instructed a Solutions Architect to encrypt the database due to new compliance requirements.\
> How can the Solutions Architect encrypt all existing and new data in the database?

snapshot, copy encrypted, load from snapshot.

> **Explanation:**\
> This comes up on almost every exam so be prepared! The key fact to remember is that you cannot alter the encryption state of an RDS database after you have deployed it. You also cannot create encrypted replicas from unencrypted instances.\
> The only solution is to create a snapshot (which will be unencrypted) and subsequently create an encrypted copy of the snapshot. You can then create a new database instance from the encrypted snapshot. The new database will be encrypted and will have a new endpoint address.


### Q14
> A company requires that IAM users must rotate their access keys every 60 days. If an access key is found to older it must be removed. A Solutions Architect must create an automated solution that checks the age of access keys and removes any keys that exceed the maximum age defined.\
> Which solution meets these requirements?

probably EventBridge to AWS config.

> **Explanation:**\
> Amazon EventBridge uses the same underlying service and API as Amazon CloudWatch Events. You can use EventBridge to detect and react to changes in the status of AWS Config events. You can create a rule that runs whenever there is a state transition, or when there is a transition to one or more states that are of interest. Then, based on rules you create, Amazon EventBridge invokes one or more target actions when an event matches the values you specify in a rule. Depending on the type of event, you might want to send notifications, capture event information, take corrective action, initiate events, or take other actions.\
> The AWS Config rule can be configured using the “access-keys-rotated” managed rule which checks if the active access keys are rotated within the number of days specified in maxAccessKeyAge. The rule is NON_COMPLIANT if the access keys have not been rotated for more than maxAccessKeyAge number of days.\
> Amazon EventBridge can react to the change of state to NON_COMPLIANT and trigger an AWS Lambda function that invalidates and removes the access key.


### Q15
> An application that runs a computational fluid dynamics workload uses a tightly-coupled HPC architecture that uses the MPI protocol and runs across many nodes. A service-managed deployment is required to minimize operational overhead.\
> Which deployment option is MOST suitable for provisioning and managing the resources required for this use case?

~~probably cloudFormation for cluster placement.~~

> **Explanation:**\
> AWS Batch Multi-node parallel jobs enable you to run single jobs that span multiple Amazon EC2 instances. With AWS Batch multi-node parallel jobs, you can run large-scale, tightly coupled, high performance computing applications and distributed GPU model training without the need to launch, configure, and manage Amazon EC2 resources directly.\
> An AWS Batch multi-node parallel job is compatible with any framework that supports IP-based, internode communication, such as Apache MXNet, TensorFlow, Caffe2, or Message Passing Interface (MPI).\
> This is the most efficient approach to deploy the resources required and supports the application requirements most effectively.


### Q16
> A Solutions Architect must design a solution for providing single sign-on to existing staff in a company. The staff manage on-premise web applications and also need access to the AWS management console to manage resources in the AWS cloud.\
> Which combination of services are BEST suited to delivering these requirements?

probably SAML.

> **Explanation:**\
> Single sign-on using federation allows users to login to the AWS console without assigning IAM credentials. The AWS Security Token Service (STS) is a web service that enables you to request temporary, limited-privilege credentials for IAM users or for users that you authenticate (such as federated users from an on-premise directory).\
> Federation (typically Active Directory) uses SAML 2.0 for authentication and grants temporary access based on the users AD credentials. The user does not need to be a user in IAM.


### Q17
> A Solutions Architect is creating a URL that lets users who sign in to the organization’s network securely access the AWS Management Console. The URL will include a sign-in token that authenticates the user to AWS. Microsoft Active Directory Federation Services is being used as the identity provider (IdP).\
> Which of the steps below will the Solutions Architect need to include when developing the custom identity broker? (choose 2)

probably AWS federation stuff.

> **Explanation:**\
> The aim of this solution is to create a single sign-on solution that enables users signed in to the organization’s Active Directory service to be able to connect to AWS resources. When developing a custom identity broker you use the AWS STS service.\
> The AWS Security Token Service (STS) is a web service that enables you to request temporary, limited-privilege credentials for IAM users or for users that you authenticate (federated users). The steps performed by the custom identity broker to sign users into the AWS management console are:
> - Verify that the user is authenticated by your local identity system
> - Call the AWS Security Token Service (AWS STS) AssumeRole or GetFederationToken API operations to obtain temporary security credentials for the user
> - Call the AWS federation endpoint and supply the temporary security credentials to request a sign-in token
> - Construct a URL for the console that includes the token
> - Give the URL to the user or invoke the URL on the user’s behalf


### Q18
> An application running on Amazon EC2 requires an EBS volume for saving structured data. The application vendor suggests that the performance of the disk should be up to 3 IOPS per GB. The capacity is expected to grow to 2 TB.\
> Taking into account cost effectiveness, which EBS volume type should be used?

~~provisioned iops?~~

> **Explanation:**\
> SSD, General Purpose (GP2) provides enough IOPS to support this requirement and is the most economical option that does. Using Provisioned IOPS would be more expensive and the other two options do not provide an SLA for IOPS.\
> More information on the volume types:
> - SSD, General Purpose (GP2) provides 3 IOPS per GB up to 16,000 IOPS. Volume size is 1 GB to 16 TB.
> - Provisioned IOPS (Io1) provides the IOPS you assign up to 50 IOPS per GiB and up to 64,000 IOPS per volume. Volume size is 4 GB to 16TB.
> - Throughput Optimized HDD (ST1) provides up to 500 IOPS per volume but does not provide an SLA for IOPS.
> - Cold HDD (SC1) provides up to 250 IOPS per volume but does not provide an SLA for IOPS.

### Q19
> A company is planning to use an Amazon S3 bucket to store a large volume of customer transaction data. The data will be structured into a hierarchy of objects, and they require a solution for running complex queries as quickly as possible. The solution must minimize operational overhead.\
> Which solution meets these requirements?

Athena? ~~EMR? dunno~~

> **Explanation:**\
> Amazon Athena is an interactive query service that makes it easy to analyze data in Amazon S3 using standard SQL. Athena is serverless, so there is no infrastructure to setup or manage, and you can start analyzing data immediately. While Amazon Athena is ideal for quick, ad-hoc querying, it can also handle complex analysis, including large joins, window functions, and arrays.\
> Athena is the fastest way to query the data in Amazon S3 and offers the lowest operational overhead as it is a fully serverless solution.

### Q20
> To increase performance and redundancy for an application a company has decided to run multiple implementations in different AWS Regions behind network load balancers. The company currently advertise the application using two public IP addresses from separate /24 address ranges and would prefer not to change these. Users should be directed to the closest available application endpoint.\
> Which actions should a solutions architect take? (Select TWO.)

probably AWS globabl accelerator.

> **Explanation:**\
> AWS Global Accelerator uses static IP addresses as fixed entry points for your application. You can migrate up to two /24 IPv4 address ranges and choose which /32 IP addresses to use when you create your accelerator.\
> This solution ensures the company can continue using the same IP addresses and they are able to direct traffic to the application endpoint in the AWS Region closest to the end user. Traffic is sent over the AWS global network for consistent performance.
> 
> ![Q20 Global Accelerator](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-50-14-8353a5204e90912fe109176eef30c333.png)


### Q21
> A dynamic website runs on Amazon EC2 instances behind an Application Load Balancer (ALB). Users are distributed around the world, and many are reporting poor website performance. The company uses Amazon Route 53 for DNS.\
> Which set of actions will improve website performance while minimizing cost?

cloudFront with route53?

> **Explanation:**\
> The most cost-effective option for improving performance is to create an Amazon CloudFront distribution. CloudFront can be used to serve both static and dynamic content. This solution will ensure that wherever users are located they will experience improved performance due to the caching of content and the usage of the AWS global network.


### Q22
> Over 500 TB of data must be analyzed using standard SQL business intelligence tools. The dataset consists of a combination of structured data and unstructured data. The unstructured data is small and stored on Amazon S3.\
> Which AWS services are most suitable for performing analytics on the data?

~~Probably RDS mariaDB and Athena~~

> **Explanation:**\
> Amazon Redshift is an enterprise-level, petabyte scale, fully managed data warehousing service. An Amazon Redshift data warehouse is an enterprise-class relational database query and management system. Redshift supports client connections with many types of applications, including business intelligence (BI), reporting, data, and analytics tools.\
> Using Amazon Redshift Spectrum, you can efficiently query and retrieve structured and semistructured data from files in Amazon S3 without having to load the data into Amazon Redshift tables. Redshift Spectrum queries employ massive parallelism to execute very fast against large datasets.\
> Used together, RedShift and RedShift spectrum are suitable for running massive analytics jobs on both the structured (RedShift data warehouse) and unstructured (Amazon S3) data.
> 
> ![Q22 Redshift](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-51-58-8e2ddfee4c97eb8dc943cfb7b27307b9.png)


### Q23
> A company plans to provide developers with individual AWS accounts. The company will use AWS Organizations to provision the accounts. A Solutions Architect must implement secure auditing using AWS CloudTrail so that all events from all AWS accounts are logged.\
> The developers must not be able to use root-level permissions to alter the AWS CloudTrail configuration in any way or access the log files in the S3 bucket. The auditing solution and security controls must automatically apply to all new developer accounts that are created.\
> Which action should the Solutions Architect take?

~~Service control policy?~~

> **Explanation:**\
> You can create a CloudTrail trail in the management account with the organization trails option enabled and this will create the trail in all AWS accounts within the organization.\
> Member accounts can see the organization trail but can't modify or delete it. By default, member accounts don't have access to the log files for the organization trail in the Amazon S3 bucket.


### Q24
> A web application hosts static and dynamic content. The application runs on Amazon EC2 instances in an Auto Scaling group behind an Application Load Balancer (ALB). The database tier runs on an Amazon Aurora database. A Solutions Architect needs to make the application more resilient to periodic increases in request rates.\
> Which architecture should the Solutions Architect implement? (Select TWO.)

cloudfront and aurora replicas.

> **Explanation:**\
> Using an Amazon CloudFront distribution can help reduce the impact of increases in requests rates as content is cached at edge locations and delivered via the AWS global network. For the database layer, Aurora Replicas will assist with serving read requests which reduces the load on the main database instance.


### Q25
> Question 25:
A Solutions Architect would like to store a backup of an Amazon EBS volume on Amazon S3.\
> What is the easiest way of achieving this?

~~automatically backed up?~~

> **Explanation:**\
> Snapshots capture a point-in-time state of an instance. Snapshots of Amazon EBS volumes are stored on S3 by design so you only need to take a snapshot and it will automatically be stored on Amazon S3.


### Q26
> A Solutions Architect is designing an application that consists of AWS Lambda and Amazon RDS Aurora MySQL. The Lambda function must use database credentials to authenticate to MySQL and security policy mandates that these credentials must not be stored in the function code.\
> How can the Solutions Architect securely store the database credentials and make them available to the function?

SSM store.

> **Explanation:**\
> In this case the scenario requires that credentials are used for authenticating to MySQL. The credentials need to be securely stored outside of the function code. Systems Manager Parameter Store provides secure, hierarchical storage for configuration data management and secrets management.\
> You can easily reference the parameters from services including AWS Lambda as depicted in the diagram below:
> 
> ![Q26 parameter store](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-54-15-28bb13e5c530a63a9405297b225a7cac.png)


### Q27
> A company is using Amazon Aurora as the database for an online retail application. Data analysts run reports every fortnight that take a long time to process and cause performance degradation for the database. A Solutions Architect has reviewed performance metrics in Amazon CloudWatch and noticed that the ReadIOPS and CPUUtilization metrics are spiking when the reports run.\
> What is the MOST cost-effective solution to resolve the performance issues?

move to aurora replica?

> **Explanation:**\
> You can issue queries to the Aurora Replicas to scale the read operations for your application. You typically do so by connecting to the reader endpoint of the cluster. That way, Aurora can spread the load for read-only connections across as many Aurora Replicas as you have in the cluster.\
> This solution is the most cost-effective method of scaling the database for reads for the regular reporting job. The reporting job will be run against the read endpoint and will not cause performance issues for the main database.


### Q28
> A customer has a public-facing web application hosted on a single Amazon Elastic Compute Cloud (EC2) instance serving videos directly from an Amazon S3 bucket.\
> Which of the following will restrict third parties from directly accessing the video assets in the bucket?

~~IAM role?~~

> **Explanation:**\
> To allow read access to the S3 video assets from the public-facing web application, you can add a bucket policy that allows s3:GetObject permission with a condition, using the aws:referer key, that the get request must originate from specific webpages. This is a good answer as it fully satisfies the objective of ensuring the that EC2 instance can access the videos but direct access to the videos from other sources is prevented.


### Q29
> A Solutions Architect has been assigned the task of moving some sensitive documents into the AWS cloud. The security of the documents must be maintained.\
> Which AWS features can help ensure that the sensitive documents cannot be read even if they are compromised? (choose 2)

SSE-S3, EBS encryption

> **Explanation:**\
> It is not specified what types of documents are being moved into the cloud or what services they will be placed on. Therefore we can assume that options include S3 and EBS. To prevent the documents from being read if they are compromised we need to encrypt them.\
> Both of these services provide native encryption functionality to ensure security of the sensitive documents. With EBS you can use KMS-managed or customer-managed encryption keys. With S3 you can use client-side or server-side encryption.


### Q30
> A company is migrating an application that comprises a web tier and a MySQL database into the AWS Cloud. The web tier will run on EC2 instances, and the database tier will run on an Amazon RDS for MySQL DB instance. Customers access the application via the Internet using dynamic IP addresses.\
> How should the Solutions Architect configure the security groups to enable connectivity to the application?

web tier inbound 443 from anywhere, database tier inbound 3306 from the webtier.

> **Explanation:**\
> The customers are connecting from dynamic IP addresses so we must assume they will be changing regularly. Therefore, it is not possible to restrict access from the IP addresses of the customers. The security group for the web tier must allow 443 (HTTPS) from 0.0.0.0/0, which means any IP source IP address.\
> For the database tier, this can best be secured by restricting access to the web tier security group. The port required to be opened is 3306 for MySQL.


### Q31
> A Solutions Architect needs to work programmatically with IAM.\
> Which feature of IAM allows direct access to the IAM web service using HTTPS to call service actions and what is the method of authentication that must be used? (choose 2)

access key and secret access key, ~~IAM Role.~~

> **Explanation:**\
> AWS recommend that you use the AWS SDKs to make programmatic API calls to IAM. However, you can also use the IAM Query API to make direct calls to the IAM web service. An access key ID and secret access key must be used for authentication when using the Query API.


### Q32
> An application is being monitored using Amazon GuardDuty. A Solutions Architect needs to be notified by email of medium to high severity events.\
> How can this be achieved?

~~CloudWatch alarm based on guard Duty?~~ maybe SNS should be involved here...

> **Explanation:**\
> A CloudWatch Events rule can be used to set up automatic email notifications for Medium to High Severity findings to the email address of your choice. You simply create an Amazon SNS topic and then associate it with an Amazon CloudWatch events rule.

### Q33
> An application running AWS uses an Elastic Load Balancer (ELB) to distribute connections between EC2 instances. A Solutions Architect needs to record information on the requester, IP, and request type for connections made to the ELB. Additionally, the Architect will also need to perform some analysis on the log files.\
> Which AWS services and configuration options can be used to collect and then analyze the logs? (choose 2)

access logs on ELBm and EMR to analyze logs.

> **Explanation:**\
> The best way to deliver these requirements is to enable access logs on the ELB and then use EMR for analyzing the log files.\
> Access Logs on ELB are disabled by default. Information includes information about the clients (not included in CloudWatch metrics) such as the identity of the requester, IP, request type etc. Logs can be optionally stored and retained in S3.\
> Amazon EMR is a web service that enables businesses, researchers, data analysts, and developers to easily and cost-effectively process vast amounts of data. EMR utilizes a hosted Hadoop framework running on Amazon EC2 and Amazon S3.


### Q34
> A Solutions Architect is designing an application that will run on an Amazon EC2 instance. The application must asynchronously invoke an AWS Lambda function to analyze thousands of .CSV files. The services should be decoupled.\
> Which service can be used to decouple the compute services?

opsWorks is EC2 configuration, SNS is notifications, Kinesis is always acting, ~~so maybe SWF? ~~

> **Explanation:**\
> You can use a Lambda function to process Amazon Simple Notification Service notifications. Amazon SNS supports Lambda functions as a target for messages sent to a topic. This solution decouples the Amazon EC2 application from Lambda and ensures the Lambda function is invoked.


### Q35
> An Amazon VPC has been deployed with private and public subnets. A MySQL database server running on an Amazon EC2 instance will soon be launched.\
> According to AWS best practice, which subnet should the database server be launched into?

private subnet.

> **Explanation:**\
> AWS best practice is to deploy databases into private subnets wherever possible. You can then deploy your web front-ends into public subnets and configure these, or an additional application tier to write data to the database.


### Q36
> Three Amazon VPCs are used by a company in the same region. The company has two AWS Direct Connect connections to two separate company offices and wishes to share these with all three VPCs. A Solutions Architect has created an AWS Direct Connect gateway.\
> How can the required connectivity be configured?

~~probably transit virtual interface to each VPC?~~

> **Explanation:**\
> You can manage a single connection for multiple VPCs or VPNs that are in the same Region by associating a Direct Connect gateway to a transit gateway. The solution involves the following components:
> - A transit gateway that has VPC attachments.
> - A Direct Connect gateway.
> - An association between the Direct Connect gateway and the transit gateway.
> - A transit virtual interface that is attached to the Direct Connect gateway.
> 
> The following diagram depicts this configuration:
> 
> ![Q36 transit gateway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-51-04-96fea7bcb03eb5705e207f44e632a778.png)


### Q37
> A team of scientists are collecting environmental data to assess the impact of pollution in a small regional town. The scientists collect data from various sensors and cameras. The data must be immediately processed to validate its accuracy, but the scientists have limited local storage space on their laptops and intermittent and unreliable connectivity to their Amazon EC2 instances and S3 buckets.\
> What should a Solutions Architect recommend?

probably snowball edge, but the answer with Kinesis seems weird.

> **Explanation:**\
> AWS Snowball Edge is a type of Snowball device with on-board storage and compute power for select AWS capabilities. Snowball Edge can do local processing and edge-computing workloads in addition to transferring data between your local environment and the AWS Cloud.\
> You can run Amazon EC2 compute instances on a Snowball Edge device using the Amazon EC2 compatible endpoint, which supports a subset of the Amazon EC2 API operations. Data can subsequently be transferred to Amazon S3 for storage and additional processing.


### Q38
> An application is deployed using Amazon EC2 instances behind an Application Load Balancer running in an Auto Scaling group. The EC2 instances connect to an Amazon RDS database. When running performance testing on the application latency was experienced when performing queries on the database. The Amazon CloudWatch metrics for the EC2 instances do not show any performance issues.\
> How can a Solutions Architect resolve the application latency issues?

read replicas

> **Explanation:**\
> The latency is most likely due to the RDS database having insufficient resources to handle the load. This can be resolved by deploying a read replica and directing queries to the replica endpoint. This offloads the performance hit of the queries from the master database which will improve overall performance and reduce the latency associated with database queries.


### Q39
> The disk configuration for an Amazon EC2 instance must be finalized. The instance will be running an application that requires heavy read/write IOPS. A single volume is required that is 500 GiB in size and needs to support 20,000 IOPS.\
> What EBS volume type should be selected?

probably provisioned IOPS ssd?

> **Explanation:**\
> This is simply about understanding the performance characteristics of the different EBS volume types. The only EBS volume type that supports over 16,000 IOPS per volume is Provisioned IOPS SSD.\
> 
> SSD, General Purpose – gp2
> - Volume size 1 GiB – 16 TiB.
> - Max IOPS/volume 16,000.
>
> SSD, Provisioned IOPS – i01
> 
> - Volume size 4 GiB – 16 TiB.
> - Max IOPS/volume 64,000.
> - HDD, Throughput Optimized – (st1)
>  Volume size 500 GiB – 16 TiB.
> - Throughput measured in MB/s, and includes the ability to burst up to 250 MB/s per TB, with a baseline throughput of 40 MB/s per TB and a maximum throughput of 500 MB/s per volume.
> 
> HDD, Cold – (sc1)
> 
> - Volume size 500 GiB – 16 TiB.
> - Lowest cost storage – cannot be a boot volume.
> - These volumes can burst up to 80 MB/s per TB, with a baseline throughput of 12 MB/s per TB and a maximum throughput of 250 MB/s per volume.
> 
> HDD
> 
> - Magnetic – Standard – cheap
> - infrequently accessed storage 
> - lowest cost storage that can be a boot volume.


### Q40
> A company runs a legacy application on an Amazon EC2 Linux instance. The application code cannot be modified, and the system cannot run on more than one instance. A Solutions Architect must design a resilient solution that can improve the recovery time for the system.\
> What should the Solutions Architect recommend to meet these requirements?

~~maybe CloudWatch alarm? ~~

> **Explanation:**\
> A RAID array uses multiple EBS volumes to improve performance or redundancy. When fault tolerance is more important than I/O performance a RAID 1 array should be used which creates a mirror of your data for extra redundancy.\
> The following table summarizes the differences between RAID 0 and RAID 1:
> 
> ![Q40 Raid0 and Raid1](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-06-02_22-58-17-415cd9d4684946806d0e3e4daa6dfe99.jpg)


### Q41
> A mobile app uploads usage information to a database. Amazon Cognito is being used for authentication, authorization and user management and users sign-in with Facebook IDs.\
> In order to securely store data in DynamoDB, the design should use temporary AWS credentials. What feature of Amazon Cognito is used to obtain temporary credentials to access AWS services?

~~SAML identity providers?~~

> **Explanation:**\
> Amazon Cognito identity pools provide temporary AWS credentials for users who are guests (unauthenticated) and for users who have been authenticated and received a token. An identity pool is a store of user identity data specific to your account.\
> With an identity pool, users can obtain temporary AWS credentials to access AWS services, such as Amazon S3 and DynamoDB.
> 
> ![Q41 Identity Pool](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-47-55-07a54edd005318f0e25395aaf4ddb73a.png)


### Q42
> When using throttling controls with API Gateway what happens when request submissions exceed the steady-state request rate and burst limits?

~~probably error 500? or maybe buffering?~~ 429?

> **Explanation:**\
> You can throttle and monitor requests to protect your backend. Resiliency through throttling rules based on the number of requests per second for each HTTP method (GET, PUT). Throttling can be configured at multiple levels including Global and Service Call.\
> When request submissions exceed the steady-state request rate and burst limits, API Gateway fails the limit-exceeding requests and returns 429 Too Many Requests error responses to the client.


### Q43
> A Solutions Architect is re-architecting an application with decoupling. The application will send batches of up to 1000 messages per second that must be received in the correct order by the consumers.\
> Which action should the Solutions Architect take?

SQS fifo?

> **Explanation:**\
> Only FIFO queues guarantee the ordering of messages and therefore a standard queue would not work. The FIFO queue supports up to 3,000 messages per second with batching so this is a supported scenario.
> 
> ![Q43 SQS types](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-53-23-3f1e752167769131a330719cd1b24d10.png)


### Q44
> A company has two accounts in an AWS Organization. The accounts are: Prod1 and Prod2. An Amazon RDS database runs in the Prod1 account. Amazon EC2 instances run in the Prod2 account. The EC2 instances in the Prod2 account must access the RDS database.\
> How can a Solutions Architect meet this requirement MOST cost-effectively?

maybe VPC sharing?

> **Explanation:**\
> VPC sharing makes use of the AWS Resource Access Manager (AWS RAM) service. It enables the sharing of VPCs across accounts. In this model, the account that owns the VPC (owner) shares one or more subnets with other accounts (participants) that belong to the same organization from AWS Organizations.\
> This scenario could be implemented with Prod1 account as the VPC owner and the Prod2 account as a VPC participant. This would allow the central control of the shared resource whilst enabling the EC2 instances in Prod2 to access the database.


### Q45
> An application in a private subnet needs to query data in an Amazon DynamoDB table. Use of the DynamoDB public endpoints must be avoided.\
> What is the most EFFICIENT and secure method of enabling access to the table?

gateway vpc endpoint?

> **Explanation:**\
> A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by AWS PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection.\
> Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.\
> With a gateway endpoint you configure your route table to point to the endpoint. Amazon S3 and DynamoDB use gateway endpoints.\
> The table below helps you to understand the key differences between the two different types of VPC endpoint:
> ![Q45 interface and gateway endpoints](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-56-38-b09d4455515fb71733786fc582abbdfa.png)


### Q46
> A company is migrating a decoupled application to AWS. The application uses a message broker based on the MQTT protocol. The application will be migrated to Amazon EC2 instances and the solution for the message broker must not require rewriting application code.
> Which AWS service can be used for the migrated message broker?

AmazonMQ is probably the right choice.

> **Explanation:**\
> Amazon MQ is a managed message broker service for Apache ActiveMQ that makes it easy to set up and operate message brokers in the cloud. Connecting current applications to Amazon MQ is easy because it uses industry-standard APIs and protocols for messaging, including JMS, NMS, AMQP, STOMP, MQTT, and WebSocket. Using standards means that in most cases, there’s no need to rewrite any messaging code when you migrate to AWS.


### Q47
> A critical web application that runs on a fleet of Amazon EC2 Linux instances has experienced issues due to failing EC2 instances. The operations team have investigated and determined that insufficient swap space is a likely cause. The operations team require a method of monitoring the swap space on the EC2 instances.\
> What should a Solutions Architect recommend?

install cloudWatch agent, custom metric.

> **Explanation:**\
> The unified CloudWatch agent enables you to collect internal system-level metrics from Amazon EC2 instances across operating systems. The metrics can include in-guest metrics, in addition to the metrics for EC2 instances. The metrics that are collected include swap_free, swap_used, and swap_used_percent.


### Q48
> A financial services company provides users with downloadable reports in PDF format. The company requires a solution that can seamlessly scale to meet the demands of a growing, global user base. The solution must be cost-effective and minimize operational overhead.\
> Which combination of services should a Solutions Architect recommend to meet these requirements?

CloudFront and S3

> **Explanation:**\
> The most cost-effective option is to use Amazon S3 for storing the PDF files and Amazon CloudFront for caching the files around the world in edge locations. This combination of services will provide seamless scalability and is cost-effective. This is also a serverless solution so operational overhead is minimized.


### Q49
> A website uses web servers behind an Internet-facing Elastic Load Balancer.\
> What record set should be created to point the customer’s DNS zone apex record at the ELB?

??? maybe alias?

> **Explanation:**\
> An Alias record can be used for resolving apex or naked domain names (e.g. example.com). You can create an A record that is an Alias that uses the customer’s website zone apex domain name and map it to the ELB DNS name.


### Q50
> A large MongoDB database running on-premises must be migrated to Amazon DynamoDB within the next few weeks. The database is too large to migrate over the company’s limited internet bandwidth so an alternative solution must be used.\
> What should a Solutions Architect recommend?

snowbal, Database migration service. ~~no need for schema conversion.~~

> **Explanation:**\
> Larger data migrations with AWS DMS can include many terabytes of information. This process can be cumbersome due to network bandwidth limits or just the sheer amount of data. AWS DMS can use Snowball Edge and Amazon S3 to migrate large databases more quickly than by other methods.\
> When you're using an Edge device, the data migration process has the following stages:
> 1. You use the AWS Schema Conversion Tool (AWS SCT) to extract the data locally and move it to an Edge device.
> 2. You ship the Edge device or devices back to AWS.
> 3. After AWS receives your shipment, the Edge device automatically loads its data into an Amazon S3 bucket.
> 4. AWS DMS takes the files and migrates the data to the target data store. If you are using change data capture (CDC), those updates are written to the Amazon S3 bucket and then applied to the target data store.


### Q51
> A Solutions Architect needs to select a low-cost, short-term option for adding resilience to an AWS Direct Connect connection.\
> What is the MOST cost-effective solution to provide a backup for the Direct Connect connection?

is a second device a short term option? I think a vpn option is better

> **Explanation:**\
> This is the most cost-effective solution. With this option both the Direct Connect connection and IPSec VPN are active and being advertised using the Border Gateway Protocol (BGP).\
> The Direct Connect link will always be preferred unless it is unavailable.

### Q52
> An organization has a data lake on Amazon S3 and needs to find a solution for performing in-place queries of the data assets in the data lake. The requirement is to perform both data discovery and SQL querying, and complex queries from a large number of concurrent users using BI tools.\
> What is the BEST combination of AWS services to use in this situation? (choose 2)

athena for ad hoc. redshift spectrum for complex?

> **Explanation:**\
> Performing in-place queries on a data lake allows you to run sophisticated analytics queries directly on the data in S3 without having to load it into a data warehouse.\
> You can use both Athena and Redshift Spectrum against the same data assets. You would typically use Athena for ad hoc data discovery and SQL querying, and then use Redshift Spectrum for more complex queries and scenarios where a large number of data lake users want to run concurrent BI and reporting workloads.


### Q53
> A storage company creates and emails PDF statements to their customers at the end of each month. Customers must be able to download their statements from the company website for up to 30 days from when the statements were generated. When customers close their accounts, they are emailed a ZIP file that contains all the statements.\
> What is the MOST cost-effective storage solution for this situation?

lifecycle policy- 30 days in standard, then glacier.

> **Explanation:**\
> The most cost-effective option is to store the PDF files in S3 Standard for 30 days where they can be easily downloaded by customers. Then, transition the objects to Amazon S3 Glacier which will reduce the storage costs. When a customer closes their account, the objects can be retrieved from S3 Glacier and provided to the customer as a ZIP file.\
> Be cautious of subtle changes to the answer options in questions like these as you may see several variations of similar questions on the exam. Also, be aware of the supported transitions (below) and minimum storage durations.
> 
> ![Q53 Life cycle](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-06-02_23-11-44-ac0d2161c900ffdfb1f1bd955d008285.jpg)


### Q54
> A HR application stores employment records on Amazon S3. Regulations mandate the records are retained for seven years. Once created the records are accessed infrequently for the first three months and then must be available within 10 minutes if required thereafter.\
> Which lifecycle action meets the requirements whilst MINIMIZING cost?

S3-IA for three months, and then S3 glacier, because we can do expidated queries if needed.

> **Explanation:**\
> The most cost-effective solution is to first store the data in S3 Standard-IA where it will be infrequently accessed for the first three months. Then, after three months expires, transition the data to S3 Glacier where it can be stored at lower cost for the remainder of the seven year period. Expedited retrieval can bring retrieval times down to 1-5 minutes.


### Q55
> A Solutions Architect is designing a three-tier web application that includes an Auto Scaling group of Amazon EC2 Instances running behind an Elastic Load Balancer. The security team requires that all web servers must be accessible only through the Elastic Load Balancer and that none of the web servers are directly accessible from the Internet.\
> How should the Architect meet these requirements?

security groups are only 'allow' rules,

> **Explanation:**\
> The web servers must be kept private so they will be not have public IP addresses. The ELB is Internet-facing so it will be publicly accessible via it’s DNS address (and corresponding public IP).\
> To restrict web servers to be accessible only through the ELB you can configure the web tier security group to allow only traffic from the ELB. You would normally do this by adding the ELBs security group to the rule on the web tier security group.


### Q56
> An application will gather data from a website hosted on an EC2 instance and write the data to an S3 bucket. The application will use API calls to interact with the EC2 instance and S3 bucket.\
> Which Amazon S3 access control method will be the MOST operationally efficient? (choose 2)

bucket policy ~~and IAM policy.~~

> **Explanation:**\
> Policies are documents that define permissions and can be applied to users, groups and roles. Policy documents are written in JSON (key value pair that consists of an attribute and a value).\
> Within an IAM policy you can grant either programmatic access or AWS Management Console access to Amazon S3 resources.

### Q57
> An application stores encrypted data in Amazon S3 buckets. A Solutions Architect needs to be able to query the encrypted data using SQL queries and write the encrypted results back the S3 bucket. As the data is sensitive fine-grained control must be implemented over access to the S3 bucket.\
> What combination of services represent the BEST options support these requirements? (choose 2)

~~ACL and~~ athena?

> **Explanation:**\
> Athena allows you to easily query encrypted data stored in Amazon S3 and write encrypted results back to your S3 bucket. Both, server-side encryption and client-side encryption are supported.\
> AWS IAM policies can be used to grant IAM users’ with fine-grained control to Amazon S3 buckets.


### Q58
> An Amazon S3 bucket is going to be used by a company to store sensitive data. A Solutions Architect needs to ensure that all objects uploaded to an Amazon S3 bucket are encrypted.\
> How can this be achieved?

~~probably aws:SecureTransport header~~

> **Explanation:**\
> To encrypt an object at the time of upload, you need to add a header called x-amz-server-side-encryption to the request to tell S3 to encrypt the object using SSE-C, SSE-S3, or SSE-KMS.\
> To enforce object encryption, create an S3 bucket policy that denies any S3 Put request that does not include the x-amz-server-side-encryption header. There are two possible values for the x-amz-server-side-encryption header: AES256, which tells S3 to use S3-managed keys, and aws:kms, which tells S3 to use AWS KMS–managed keys.\
> The example policy below denies Put requests that do not have the correct encryption header set:
> 
> ![Q58 Put Policy Json](https://img-c.udemycdn.com/redactor/raw/test_question_description/2021-06-02_22-55-40-e722cfe635c4da8e594c03abb6d181b5.jpg)


### Q59
> A highly elastic application consists of three tiers. The application tier runs in an Auto Scaling group and processes data and writes it to an Amazon RDS MySQL database. The Solutions Architect wants to restrict access to the database tier to only accept traffic from the instances in the application tier. However, instances in the application tier are being constantly launched and terminated.\
> How can the Solutions Architect configure secure access to the database tier?

security group allow from other security group?

> **Explanation:**\
> The best option is to configure the database security group to only allow traffic that originates from the application security group. You can also define the destination port as the database port. This setup will allow any instance that is launched and attached to this security group to connect to the database.


### Q60
> A company are finalizing their disaster recovery plan. A limited set of core services will be replicated to the DR site ready to seamlessly take over the in the event of a disaster. All other services will be switched off.\
> Which DR strategy is the company using?

~~warm standby?~~

> **Explanation:**\
> In this DR approach, you simply replicate part of your IT structure for a limited set of core services so that the AWS cloud environment seamlessly takes over in the event of a disaster.\
> A small part of your infrastructure is always running simultaneously syncing mutable data (as databases or documents), while other parts of your infrastructure are switched off and used only during testing.\
> Unlike a backup and recovery approach, you must ensure that your most critical core elements are already configured and running in AWS (the pilot light). When the time comes for recovery, you can rapidly provision a full-scale production environment around the critical core.


### Q61
> An eCommerce company has a very popular web application that receives a large amount of traffic. The application must store customer profile data and shopping cart information in a database. A Solutions Architect must design the database solution to support peak loads of several million requests per second and millisecond response times. Operational overhead must be minimized, and scaling should not cause downtime.\
> Which database solution should the Solutions Architect recommend?

probably DynamoDB

> **Explanation:**\
> Amazon DynamoDB is a non-relational database that is managed for you. It can scale without downtime and with minimal operational overhead. DynamoDB can support the request rates and response times required by this solution and is often used in eCommerce solutions and for session state use cases.


### Q62
> A company runs a streaming application on AWS that ingests data in near real-time and then processes the data. The data processing takes 30 minutes to complete. As the volume of data being ingested by the application has increased, high latency has occurred. A Solutions Architect needs to design a scalable and serverless solution to improve performance.\
> Which combination of steps should the Solutions Architect take? (Select TWO.)

EC2 isn't serverless. Lambdas are time limited. so maybe Firehose + Fargate?

> **Explanation:**\
> The application is a streaming application that ingests near real time data. This is a good fit for Amazon Kinesis Data Firehose which can ingest data and load it directly to a data store where it can be subsequently processed. We then need a serverless solution for processing the data. AWS Fargate is a serverless service that uses Amazon ECS for running Docker containers on AWS.\
> This solution will seamlessly scale for the data ingestion and processing. It is also fully serverless.


### Q63
> A Solutions Architect would like to implement a method of automating the creation, retention, and deletion of backups for the Amazon EBS volumes in an Amazon VPC.\
> What is the easiest way to automate these tasks using AWS tools?

~~scheduled jobs with 'create-backup'?~~

> **Explanation:**\
> You backup EBS volumes by taking snapshots. This can be automated via the AWS CLI command “create-snapshot”. However the question is asking for a way to automate not just the creation of the snapshot but the retention and deletion too.\
> The EBS Data Lifecycle Manager (DLM) can automate all of these actions for you and this can be performed centrally from within the management console.


### Q64
> A retail organization sends coupons out twice a week and this results in a predictable surge in sales traffic. The application runs on Amazon EC2 instances behind an Elastic Load Balancer. The organization is looking for ways lower costs while ensuring they meet the demands of their customers.\
> How can they achieve this goal?

~~maybe combine spot instances with on-demand?~~

> **Explanation:**\
> On-Demand Capacity Reservations enable you to reserve compute capacity for your Amazon EC2 instances in a specific Availability Zone for any duration. By creating Capacity Reservations, you ensure that you always have access to EC2 capacity when you need it, for as long as you need it. When used in combination with savings plans, you can also gain the advantages of cost reduction.


### Q65
> A Solutions Architect has setup a VPC with a public subnet and a VPN-only subnet. The public subnet is associated with a custom route table that has a route to an Internet Gateway. The VPN-only subnet is associated with the main route table and has a route to a virtual private gateway.\
> The Architect has created a new subnet in the VPC and launched an EC2 instance in it. However, the instance cannot connect to the Internet.\
> What is the MOST likely reason?

maybe automatically associated with the default one...

> **Explanation:**\
> When you create a new subnet, it is automatically associated with the main route table. Therefore, the EC2 instance will not have a route to the Internet. The Architect should associate the new subnet with the custom route table.
</details>


## Exam 5
<details>
<summary>
65 questions
</summary>

### Q01
> A client has requested a design for a fault tolerant database that can failover between AZs. You have decided to use RDS in a multi-AZ configuration.\
> What type of replication will the primary database use to replicate to the standby instance?

~~probably asynchronous~~, but who knows

> **Explanation:**\
> Multi-AZ RDS creates a replica in another AZ and synchronously replicates to it (DR only). Multi-AZ deployments for the MySQL, MariaDB, Oracle and PostgreSQL engines utilize synchronous physical replication. Multi-AZ deployments for the SQL Server engine use synchronous logical replication (SQL Server-native Mirroring technology).
> 
> ![Q1 replication](https://img-c.udemycdn.com/redactor/raw/2020-06-28_19-59-49-4aeaf1017517ceea6f5e5216d8b1e6de.png)


### Q02
> A Solutions Architect created a new subnet in an Amazon VPC and launched an Amazon EC2 instance into it. The Solutions Architect needs to directly access the EC2 instance from the Internet and cannot connect.\
> Which steps should be undertaken to troubleshoot the issue? (choose 2)

probably route table and public IP address. i think that NAT gateway isn't mandatory (NAT instance?)

> **Explanation:**\
> A public subnet is a subnet that's associated with a route table that has a route to an Internet gateway.\
> Public subnets are subnets that have:
> - “Auto-assign public IPv4 address” set to “Yes”.
> - The subnet route table has an attached Internet Gateway.
### Q03
> The development team in a media organization is moving their SDLC processes into the AWS Cloud.\
> Which AWS service can a Solutions Architect recommend that is primarily used for software version control?

probably CodeCommit

> **Explanation:**\
> AWS CodeCommit is a fully-managed source control service that hosts secure Git-based repositories. It makes it easy for teams to collaborate on code in a secure and highly scalable ecosystem. CodeCommit eliminates the need to operate your own source control system or worry about scaling its infrastructure. You can use CodeCommit to securely store anything from source code to binaries, and it works seamlessly with your existing Git tools.


### Q04
> Encrypted Amazon Elastic Block Store (EBS) volumes are attached to some Amazon EC2 instances.\
> Which statements are correct about using encryption with Amazon EBS volumes? (choose 2)

~~I think that it's only encrypted at rest~~, and maybe all types can be encrypted?

> **Explanation:**\
> Some facts about Amazon EBS encrypted volumes and snapshots:
> - All EBS types support encryption and all instance families now support encryption.
> - Not all instance types support encryption.
> - Data in transit between an instance and an encrypted volume is also encrypted (data is encrypted in trans.
> - You can have encrypted an unencrypted EBS volumes attached to an instance at the same time.
> - Snapshots of encrypted volumes are encrypted automatically.
> - EBS volumes restored from encrypted snapshots are encrypted automatically.

- EBS volumes created from encrypted snapshots are also encrypted.
### Q05
> A Solutions Architect just completed the implementation of a 2-tier web application for a client. The application uses Amazon EC2 instances, Amazon ELB and Auto Scaling across two subnets. After deployment the Solutions Architect noticed that only one subnet has EC2 instances running in it.\
> What might be the cause of this situation?

probably the auto scaling group wasn't configured properly?

> **Explanation:**\
> You can specify which subnets Auto Scaling will launch new instances into. Auto Scaling will try to distribute EC2 instances evenly across AZs. If only one subnet has EC2 instances running in it the first thing to check is that you have added all relevant subnets to the configuration.


### Q06
> A Solutions Architect needs to create a file system that can be concurrently accessed by multiple Amazon EC2 instances across multiple availability zones. The file system needs to support high throughput and the ability to burst. As the data that will be stored on the file system will be sensitive, it must be encrypted at rest and in transit.\
> Which storage solution should the Solutions Architect use for the shared file system?

file system means EFS

> **Explanation:**\
> EFS is a fully-managed service that makes it easy to set up and scale file storage in the Amazon Cloud. EFS file systems are mounted using the NFSv4.1 protocol. EFS is designed to burst to allow high throughput levels for periods of time. EFS also offers the ability to encrypt data at rest and in transit.


### Q07
> An application that is being installed on an Amazon EC2 instance requires a persistent block storage volume. The data must be encrypted at rest and regular volume-level backups must be automated.\
> Which solution options should be used?

persistent block storage means EBS

> **Explanation:**\
> For block storage the Solutions Architect should use either Amazon EBS or EC2 instance store. However, the instance store is non-persistent so EBS must be used. With EBS you can encrypt your volume and automate volume-level backups using snapshots that are run by Data Lifecycle Manager.


### Q08
> A company has deployed an API using Amazon API Gateway. There are many repeat requests and a solutions architect has been asked to implement measures to reduce request latency and the number of calls to the Amazon EC2 endpoint.\
> How can this be most easily achieved?

cacheing, ~~over a method?~~

> **Explanation:**\
> You can enable API caching in Amazon API Gateway to cache your endpoint's responses. With caching, you can reduce the number of calls made to your endpoint and also improve the latency of requests to your API.\
> When you enable caching for a stage, API Gateway caches responses from your endpoint for a specified time-to-live (TTL) period, in seconds. API Gateway then responds to the request by looking up the endpoint response from the cache instead of making a request to your endpoint. The default TTL value for API caching is 300 seconds. The maximum TTL value is 3600 seconds. TTL=0 means caching is disabled.
> 
> ![Q8 api gateway cacheing](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-14-50-ea00cfb6679c3d0035a50e9c28b4e682.png)


### Q09

> A Solutions Architect is designing an application for processing and extracting data from log files. The log files are generated by an application and the number and frequency of updates varies. The files are up to 1 GB in size and processing will take around 40 seconds for each file.\
> Which solution is the most cost-effective?

files are saved to S3, processing done by lambda.

> **Explanation:**\
> The question asks for the most cost-effective solution and therefor a serverless and automated solution will be the best choice.\
> AWS Lambda can run custom code in response to Amazon S3 bucket events. You upload your custom code to AWS Lambda and create a function. When Amazon S3 detects an event of a specific type (for example, an object created event), it can publish the event to AWS Lambda and invoke your function in Lambda. In response, AWS Lambda executes your function.
> ![Q9 lambda and S3 bucket](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-12-32-5693ddd5be6f40334bb8d981a4099fbd.png)


### Q10
> An organization is planning their disaster recovery solution. They plan to run a scaled down version of a fully functional environment. In a DR situation the recovery time must be minimized.\
> Which DR strategy should a Solutions Architect recommend?

scaled down version is "warm standby"

> **Explanation:**\
> The term warm standby is used to describe a DR scenario in which a scaled-down version of a fully functional environment is always running in the cloud. A warm standby solution extends the pilot light elements and preparation.\
> It further decreases the recovery time because some services are always running. By identifying your business-critical systems, you can fully duplicate these systems on AWS and have them always on.


### Q11
> A company runs an application on-premises that must consume a REST API running on Amazon API Gateway. The company has an AWS Direct Connect connection to their Amazon VPC. The solutions architect wants all API calls to use private addressing only and avoid the internet.\
> How can this be achieved?

private virtual interface, vpc endpoint?

> **Explanation:**\
> The requirements are to avoid the internet and use private IP addresses only. The best solution is to use a private virtual interface across the Direct Connect connection to connect to the VPC using private IP addresses. A VPC endpoint for Amazon API Gateway can be created and this will provide access to API Gateway using private IP addresses and avoids the internet completely.


### Q12
> A Solutions Architect regularly deploys and manages infrastructure services for customers on AWS. The SysOps team are facing challenges in tracking changes that are made to the infrastructure services and rolling back when problems occur.\
> How can a Solutions Architect BEST assist the SysOps team?

probably cloudFormation templates? but maybe codeDeploy?

> **Explanation:**\
> When you provision your infrastructure with AWS CloudFormation, the AWS CloudFormation template describes exactly what resources are provisioned and their settings. Because these templates are text files, you simply track differences in your templates to track changes to your infrastructure, similar to the way developers control revisions to source code.\
> For example, you can use a version control system with your templates so that you know exactly what changes were made, who made them, and when. If at any point you need to reverse changes to your infrastructure, you can use a previous version of your template.


### Q13
> A manager is concerned that the default service limits my soon be reached for several AWS services.\
> Which AWS tool can a Solutions Architect use to display current usage and limits?

~~cloudWatch? maybe dashboard?~~

> **Explanation:**\
> Trusted Advisor is an online resource to help you reduce cost, increase performance, and improve security by optimizing your AWS environment. Trusted Advisor provides real time guidance to help you provision your resources following AWS best practices.\
> AWS Trusted Advisor offers a Service Limits check (in the Performance category) that displays your usage and limits for some aspects of some services.


### Q14
> A Solutions Architect is building a small web application running on Amazon EC2 that will be serving static content. The user base is spread out globally and speed is important.\
> Which AWS service can deliver the best user experience cost-effectively and reduce the load on the web server?

cloudFront

> **Explanation:**\
> This is a good use case for Amazon CloudFront as the user base is spread out globally and CloudFront can cache the content closer to users and also reduce the load on the web server running on EC2.


### Q15
> A Solutions Architect has deployed a number of AWS resources using CloudFormation. Some changes must be made to a couple of resources within the stack. Due to recent failed updates, the Solutions Architect is a little concerned about the effects that implementing updates to the resources might have on other resources in the stack.\
> What is the easiest way to proceed cautiously?

~~maybe creating a new stack?~~
> **Explanation:**\
> AWS CloudFormation provides two methods for updating stacks: direct update or creating and executing change sets. When you directly update a stack, you submit changes and AWS CloudFormation immediately deploys them.\
> Use direct updates when you want to quickly deploy your updates. With change sets, you can preview the changes AWS CloudFormation will make to your stack, and then decide whether to apply those changes.


### Q16
> A company has several AWS accounts each with multiple Amazon VPCs. The company must establish routing between all private subnets. The architecture should be simple and allow transitive routing to occur.\
> How should the network connectivity be configured?

~~i think that transit gateway is for DX. vpc peering isn't transitive, so maybe hub-and-spoke?~~

> **Explanation:**\
> You can build a hub-and-spoke topology with AWS Transit Gateway that supports transitive routing. This simplifies the network topology and adds additional features over VPC peering. AWS Resource Access Manager can be used to share the connection with the other AWS accounts.
> 
> ![Q16 transit gateway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-10-15-06b3dc1b26269bff9b93a25c89727620.png)


### Q17
> A company has over 2000 users and is planning to migrate data into the AWS Cloud. Some of the data is user’s home folders on an existing file share and the plan is to move this data to Amazon S3. Each user will have a folder in a shared bucket under the folder structure: bucket/home/%username%.\
> What steps should a Solutions Architect take to ensure that each user can access their own home folder and no one else’s? (choose 2)

IAM policy (folder level) ~~and bucket policy~~

> **Explanation:**\
> The [AWS blog URL ](https://aws.amazon.com/blogs/security/writing-iam-policies-grant-access-to-user-specific-folders-in-an-amazon-s3-bucket/)explains how to construct an IAM policy for a similar scenario. Please refer to the article for detailed instructions.


### Q18
> An existing Auto Scaling group is running with eight Amazon EC2 instances. A Solutions Architect has attached an Elastic Load Balancer (ELB) to the Auto Scaling group by connecting a Target Group. The ELB is in the same region and already has ten EC2 instances running in the Target Group.\
> When attempting to attach the ELB the request immediately fails, what is the MOST likely cause?

~~probably that you can't attach existing EC2 instances to ASG, they must be created by it?~~

> **Explanation:**\
> You can attach one or more Target Groups to your ASG to include instances behind an ALB and the ELBs must be in the same region. Once you do this any EC2 instance existing or added by the ASG will be automatically registered with the ASG defined ELBs. If adding an instance to an ASG would result in exceeding the maximum capacity of the ASG the request will fail.


### Q19
> A Solutions Architect has deployed an API using Amazon API Gateway and created usage plans and API keys for several customers. Requests from one particular customer have been excessive and the solutions architect needs to limit the rate of requests. Other customers should not be affected.\
> How should the solutions architect proceed?

I'm guessing the 'per-client' throttling.

> **Explanation:**\
> Per-client throttling limits are applied to clients that use API keys associated with your usage policy as client identifier. This can be applied to the single customer that is issuing excessive API requests. This is the best option to ensure that only one customer is affected.\
> In the diagram below, per-client throttling limits are set in a usage plan:
> 
> ![Q19 API gateway throtelling](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-13-18-dca547aabda672e802edb7d89cd98c7c.png)


### Q20
> The database layer of an on-premises web application is being migrated to AWS. The database uses a multi-threaded, in-memory caching layer to improve performance for repeated queries.\
> Which service would be the most suitable replacement for the database cache?

~~Redis is the stronger option from the two caching engines.~~

> **Explanation:**\
> Amazon ElastiCache with the Memcached engine is an in-memory database that can be used as a database caching layer. The memached engine supports multiple cores and threads and large nodes.
> 
> ![Q20 elasticCache options](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-11-45-0d3bdca18464fbf41521677922bfcec8.png)


### Q21
> An application is running in a private subnet of an Amazon VPC and must have outbound internet access for downloading updates. The Solutions Architect does not want the application exposed to inbound connection attempts.\
> Which steps should be taken?

internet gateway ~~but without a NAT gateway?~~

> **Explanation:**\
> To enable outbound connectivity for instances in private subnets a NAT gateway can be created. The NAT gateway is created in a public subnet and a route must be created in the private subnet pointing to the NAT gateway for internet-bound traffic. An internet gateway must be attached to the VPC to facilitate outbound connections.\
> You cannot directly connect to an instance in a private subnet from the internet. You would need to use a bastion/jump host. Therefore, the application will not be exposed to inbound connection attempts.
> 
> ![Q21 internet gateway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-07-43-70afcd404db57b167e180b7ee7fee9d7.png)


### Q22
> An application regularly uploads files from an Amazon EC2 instance to an Amazon S3 bucket. The files can be a couple of gigabytes in size and sometimes the uploads are slower than desired.\
> What method can be used to increase throughput and reduce upload times?

probably multi-part upload

> **Explanation:**\
> Multipart upload can be used to speed up uploads to S3. Multipart upload uploads objects in parts independently, in parallel and in any order. It is performed using the S3 Multipart upload API and is recommended for objects of 100MB or larger. It can be used for objects from 5MB up to 5TB and must be used for objects larger than 5GB.


### Q23
> A Solutions Architect is designing the messaging and streaming layers of a serverless application. The messaging layer will manage communications between components and the streaming layer will manage real-time analysis and processing of streaming data.\
> The Architect needs to select the most appropriate AWS services for these functions.\
> Which services should be used for the messaging and streaming layers? (choose 2)

Kinesis for the streaming data, SNS for the communication and messaging.

> **Explanation:**\
> Amazon Kinesis makes it easy to collect, process, and analyze real-time streaming data. With Amazon Kinesis Analytics, you can run standard SQL or build entire streaming applications using SQL.\
> Amazon Simple Notification Service (Amazon SNS) provides a fully managed messaging service for pub/sub patterns using asynchronous event notifications and mobile push notifications for microservices, distributed systems, and serverless applications.


### Q24
> A customer runs an application on-premise that stores large media files. The data is mounted to different servers using either the SMB or NFS protocols. The customer is having issues with scaling the storage infrastructure on-premise and is looking for a way to offload the data set into the cloud whilst retaining a local cache for frequently accessed content.\
> Which of the following is the best solution?

I think that NFS means file gateway, but maybe SMB is volume (with cached)?

> **Explanation:**\
> File gateway provides a virtual on-premises file server, which enables you to store and retrieve files as objects in Amazon S3. It can be used for on-premises applications, and for Amazon EC2-resident applications that need file storage in S3 for object based workloads. Used for flat files only, stored directly on S3. File gateway offers SMB or NFS-based access to data in Amazon S3 with local caching.
> ![Q24 file gateway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-01-33-c38df2adc575ee7ed2b27d0d67d08e2f.png)


### Q25
> A Solutions Architect is designing a migration strategy for a company moving to the AWS Cloud. The company use a shared Microsoft filesystem that uses Distributed File System Namespaces (DFSN).\
> What will be the MOST suitable migration strategy for the filesystem?

probably FSx windows.

> **Explanation:**\
> The destination filesystem should be Amazon FSx for Windows File Server. This supports DFSN and is the most suitable storage solution for Microsoft filesystems. AWS DataSync supports migrating to the Amazon FSx and automates the process.


### Q26
> An application has been migrated from on-premises to an Amazon EC2 instance. The migration has failed due to an unknown dependency that the application must communicate with an on-premises server using private IP addresses.\
> Which action should a solutions architect take to quickly provision the necessary connectivity?

maybe virtual private gateway? DX takes time to establish.

> **Explanation:**\
> A virtual private gateway is a logical, fully redundant distributed edge routing function that sits at the edge of your VPC. You must create a VPG in your VPC before you can establish an AWS Managed site-to-site VPN connection. The other end of the connection is the customer gateway which must be established on the customer side of the connection.
> ![Q26 virtual private gateway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-08-32-b6eebdb5a8f23f26677a1be7da633443.png)


### Q27
> An application stack is being created which needs a message bus to decouple the application components from each other. The application will generate up to 300 messages per second without using batching. A Solutions Architect needs to ensure that a message is delivered only once and duplicates are not introduced into the queue. It is not necessary to maintain the order of the messages.\
> Which SQS queue type should be used?

~~SQS standard~~
> **Explanation:**\
> The key fact you need to consider here is that duplicate messages cannot be introduced into the queue. For this reason alone you must use a FIFO queue. The statement about it not being necessary to maintain the order of the messages is meant to confuse you, as that might lead you to think you can use a standard queue, but standard queues don’t guarantee that duplicates are not introduced into the queue.\
> FIFO (first-in-first-out) queues preserve the exact order in which messages are sent and received – note that this is not required in the question but exactly once processing is. FIFO queues provide exactly-once processing, which means that each message is delivered once and remains available until a consumer processes it and deletes it.


### Q28
> Amazon CloudWatch is being used to monitor the performance of AWS Lambda.
> Which metrics does Lambda track? (choose 2)

probably number of request and latency per request.

> **Explanation:**\
> AWS Lambda automatically monitors Lambda functions and reports metrics through Amazon CloudWatch. Lambda tracks the number of requests, the latency per request, and the number of requests resulting in an error. You can view the request rates and error rates using the AWS Lambda Console, the CloudWatch console, and other AWS resources.
### Q29
> There is new requirement for a database that will store a large number of records for an online store. You are evaluating the use of DynamoDB.\
> Which of the following are AWS best practices for DynamoDB? (choose 2)

400kb data limit. ~~maybe separate local secondary index?~~

> **Explanation:**\
> DynamoDB best practices include:
> - Keep item sizes small.
> - If you are storing serial data in DynamoDB that will require actions based on data/time use separate tables for days, weeks, months.
> - Store more frequently and less frequently accessed data in separate tables.
> - If possible compress larger attribute values.
> - Store objects larger than 400KB in S3 and use pointers (S3 Object ID) in DynamoDB.

### Q30
> An Amazon VPC contains a mixture of Amazon EC2 instances in production and non-production environments. A Solutions Architect needs to devise a way to segregate access permissions to different sets of users for instances in different environments.\
> How can this be achieved? (choose 2)

use tagging and IAM policy

> **Explanation:**\
> You can use the condition checking in IAM policies to look for a specific tag. IAM checks that the tag attached to the principal making the request matches the specified key name and value.


### Q31
> The application development team in a company have created a new application written in .NET. A Solutions Architect is looking for a way to easily deploy the application whilst maintaining full control of the underlying resources.\
> Which PaaS service provided by AWS would BEST suit this requirement?

i think that CloudFormation is for managing the cloud resources (like iam policies, subnets, etc..) so maybe beanstalk?

> **Explanation:**\
> AWS Elastic Beanstalk can be used to quickly deploy and manage applications in the AWS Cloud. Developers upload applications and Elastic Beanstalk handles the deployment details of capacity provisioning, load balancing, auto-scaling, and application health monitoring. It is considered to be a Platform as a Service (PaaS) solution and allows full control of the underlying resources.


### Q32
> A Solutions Architect needs to migrate an Oracle database running on RDS onto Amazon RedShift to improve performance and reduce cost.\
> Which combination of tasks using AWS services should be used to execute the migration? (choose 2)

probably schema conversion and database migration

> **Explanation:**\
> Convert the data warehouse schema and code from the Oracle database running on RDS using the AWS Schema Conversion Tool (AWS SCT) then migrate data from the Oracle database to Amazon Redshift using the AWS Database Migration Service (AWS DMS).


### Q33
> A new application runs on Amazon EC2 instances and uses API Gateway and AWS Lambda. The company is planning on running an advertising campaign that will likely result in significant hits to the application after each ad is run.\
> A Solutions Architect is concerned about the impact this may have on the application and would like to put in place some controls to limit the number of requests per second that hit the application.\
> What controls should the Solutions Architect implement?

maybe throtelling rules on the API? lambdas scale, but gateway can have too many requests.

> **Explanation:**\
> The key requirement is to limit the number of requests per second that hit the application. This can only be done by implementing throttling rules on the API Gateway. Throttling enables you to throttle the number of requests to your API which in turn means less traffic will be forwarded to your application server.


### Q34
> An Amazon ElastiCache for Redis cluster runs across multiple Availability Zones. A solutions architect is concerned about the security of sensitive data as it is replicated between nodes.\
> How can the solutions architect protect the sensitive data?

maybe in-transit encryption?


> **Explanation:**\
> Amazon ElastiCache in-transit encryption is an optional feature that allows you to increase the security of your data at its most vulnerable points—when it is in transit from one location to another. Because there is some processing needed to encrypt and decrypt the data at the endpoints, enabling in-transit encryption can have some performance impact. You should benchmark your data with and without in-transit encryption to determine the performance impact for your use cases.\
> ElastiCache in-transit encryption implements the following features:
> - Encrypted connections—both the server and client connections are Secure Socket Layer (SSL) encrypted.
> - Encrypted replication—data moving between a primary node and replica nodes is encrypted.
> - Server authentication—clients can authenticate that they are connecting to the right server.
> - Client authentication—using the Redis AUTH feature, the server can authenticate the clients.


### Q35
> An application analyzes images of people that are uploaded to an Amazon S3 bucket. The application determines demographic data which is then saved to a .CSV file in another S3 bucket. The data must be encrypted at rest and then queried using SQL. The solution should be fully serverless.\
> Which actions should a Solutions Architect take to encrypt and query the data?

redshift spectrum is for unstructured data on S3...  but probably KMS + athena

> **Explanation:**\
> Amazon Athena is an interactive query service that makes it easy to analyze data in Amazon S3 using standard SQL. Athena is serverless, so there is no infrastructure to manage, and you pay only for the queries that you run. Amazon Athena supports encrypted data for both the source data and query results, for example, using Amazon S3 with AWS KMS.


### Q36
> A three-tier web application that is deployed in an Amazon VPC has been experiencing heavy load on the database layer. The database layer uses an Amazon RDS MySQL instance in a multi-AZ configuration. Customers have been complaining about poor response times. During troubleshooting it has been noted that the database layer is experiencing high read contention during peak hours of the day.\
> What are two possible options that could be used to offload some of the read traffic from the database to resolve the performance issues? (choose 2)

add elasticCache or read replicas.

> **Explanation:**\
> Amazon ElastiCache is a web service that makes it easy to deploy and run Memcached or Redis protocol-compliant server nodes in the cloud. The in-memory caching provided by ElastiCache can be used to significantly improve latency and throughput for many read-heavy application workloads or compute-intensive workloads.\
> Read replicas are used for read heavy DBs and replication is asynchronous. They are for workload sharing and offloading and are created from a snapshot of the master instance.


### Q37
> A Solutions Architect is creating a design for a two-tier application with a MySQL RDS back-end. The performance requirements of the database tier are hard to quantify until the application is running and the Architect is concerned about right-sizing the database.\
> What methods of scaling are possible after the MySQL RDS database is deployed? (choose 2)

horizontal scaling for reads with read replicas. ~~maybe horizontal with multi-master RDS?~~ can the instance size be changed after deployment ~~(probably not...)~~

> **Explanation:**\
> To handle a higher load in your database, you can vertically scale up your master database with a simple push of a button. In addition to scaling your master database vertically, you can also improve the performance of a read-heavy database by using read replicas to horizontally scale your database.

### Q38
> An event in CloudTrail is the record of an activity in an AWS account.\
> What are the two types of events that can be logged in CloudTrail? (choose 2)

data events and managements events (control plane)

> **Explanation:**\
> Trails can be configured to log Data events and management events:
> - Data events: These events provide insight into the resource operations performed on or within a resource. These are also known as data plane operations.
> - Management events: Management events provide insight into management operations that are performed on resources in your AWS account. These are also known as control plane operations. Management events can also include non-API events that occur in your account.


### Q39
> A large quantity of data is stored on a NAS device on-premises and accessed using the SMB protocol. The company require a managed service for hosting the filesystem and a tool to automate the migration.\
> Which actions should a Solutions Architect take?

FSx with aws DataSync?

> **Explanation:**\
> Amazon FSx for Windows File Server provides fully managed, highly reliable, and scalable file storage that is accessible over the industry-standard Server Message Block (SMB) protocol. This is the most suitable destination for this use case.\
> AWS DataSync can be used to move large amounts of data online between on-premises storage and Amazon S3, Amazon EFS, or Amazon FSx for Windows File Server. The source datastore can be Server Message Block (SMB) file servers.
> 
> ![Q39 DataSync](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-11-07-d909de298264e2e376a88159790c0be1.png)


### Q40
> A large multinational retail company has a presence in AWS in multiple regions. The company has established a new office and needs to implement a high-bandwidth, low-latency connection to multiple VPCs in multiple regions within the same account. The VPCs each have unique CIDR ranges.\
> What would be the optimum solution design using AWS technology? (choose 2)

it's suggested to have more than one DX,~~so one to each region?~~ or maybe on DX and a gateway?

> **Explanation:**\
> The company should implement an AWS Direct Connect connection to the closest region. A Direct Connect gateway can then be used to create private virtual interfaces (VIFs) to each AWS region.\
> Direct Connect gateway provides a grouping of Virtual Private Gateways (VGWs) and Private Virtual Interfaces (VIFs) that belong to the same AWS account and enables you to interface with VPCs in any AWS Region (except AWS China Region).\
> You can share a private virtual interface to interface with more than one Virtual Private Cloud (VPC) reducing the number of BGP sessions required.

### Q41
> The AWS Acceptable Use Policy describes permitted and prohibited behavior on AWS and includes descriptions of prohibited security violations and network abuse.\
> According to the policy, what is AWS’s position on penetration testing?

~~**no penetration testing without approval!**~~

> **Explanation:**\
> AWS customers are welcome to carry out security assessments or penetration tests against their AWS infrastructure without prior approval for 8 services.


### Q42
> A Solutions Architect is designing the compute layer of a serverless application. The compute layer will manage requests from external systems, orchestrate serverless workflows, and execute the business logic.\
> The Architect needs to select the most appropriate AWS services for these functions. Which services should be used for the compute layer? (choose 2)

amazon gateway API and lambda for business logic, SWS step functions for workflows?

> **Explanation:**\
> With Amazon API Gateway, you can run a fully managed REST API that integrates with Lambda to execute your business logic and includes traffic management, authorization and access control, monitoring, and API versioning.\
> AWS Step Functions orchestrates serverless workflows including coordination, state, and function chaining as well as combining long-running executions not supported within Lambda execution limits by breaking into multiple steps or by calling workers running on Amazon Elastic Compute Cloud (Amazon EC2) instances or on-premises.


### Q43
> A new financial platform has been re-architected to use Docker containers in a micro-services architecture. The new architecture will be implemented on AWS and a Solutions Architect must recommend the solution configuration. For operational reasons, it will be necessary to access the operating system of the instances on which the containers run.\
> Which solution delivery option should the Architect select?

probably ECS with EC2

> **Explanation:**\
> Amazon Elastic Container Service (ECS) is a highly scalable, high performance container management service that supports Docker containers and allows you to easily run applications on a managed cluster of Amazon EC2 instances.\
> The EC2 Launch Type allows you to run containers on EC2 instances that you manage so you will be able to access the operating system instances.


### Q44
> A Solutions Architect must enable an application to download software updates from the internet. The application runs on a series of EC2 instances in an Auto Scaling group running in a private subnet. The solution must involve minimal ongoing systems management effort.\
> How should the Solutions Architect proceed?

NAT gateway is much more 'managed' than NAT instance.

> **Explanation:**\
> Both a NAT gateway or a NAT instance can be used for this use case. Both services enable internet access for instances in private subnets. However, the NAT instance runs on an EC2 instance you must launch, configure and manage and therefore involves more ongoing systems management effort.


### Q45
> A company has an eCommerce application that runs from multiple AWS Regions. Each region has a separate database running on Amazon EC2 instances. The company plans to consolidate the data to a columnar database and run analytics queries.\
> Which approach should the company take?

COPY into redshift warehouse and run the analytics there.

> **Explanation:**\
> Amazon Redshift is an enterprise-level, petabyte scale, fully managed data warehousing service. It uses columnar storage to improve the performance of complex queries.\
> You can use the COPY command to load data in parallel from one or more remote hosts, such Amazon EC2 instances or other computers. COPY connects to the remote hosts using SSH and executes commands on the remote hosts to generate text output.


### Q46
> An Amazon EC2 instance running a video on demand web application has been experiencing high CPU utilization. A Solutions Architect needs to take steps to reduce the impact on the EC2 instance and improve performance for consumers.\
> Which of the steps below would help?

~~maybe RTMP distribution?~~

> **Explanation:**\
> This is a good use case for CloudFront which is a content delivery network (CDN) that caches content to improve performance for users who are consuming the content. This will take the load off of the EC2 instances as CloudFront has a cached copy of the video files.\
> An origin is the origin of the files that the CDN will distribute. Origins can be either an S3 bucket, an EC2 instance, and Elastic Load Balancer, or Route 53 – can also be external (non-AWS).


### Q47
> A client has made some updates to their web application. The application uses an Auto Scaling Group to maintain a group of several EC2 instances. The application has been modified and a new AMI must be used for launching any new instances.\
> What does a Solutions Architect need to do to add the new AMI?

modify the asg to a new launch configuration

> **Explanation:**\
> A launch configuration is the template used to create new EC2 instances and includes parameters such as instance family, instance type, AMI, key pair and security groups.\
> You cannot edit a launch configuration once defined. In this case you can create a new launch configuration that uses the new AMI and any new instances that are launched by the ASG will use the new AMI.


### Q48
> A Solutions Architect is creating a multi-tier application that includes loosely-coupled, distributed application components and needs to determine a method of sending notifications instantaneously.\
> Using Amazon SNS which transport protocols are supported? (choose 2)

probably email and HTTPS?

> **Explanation:**\
> Note that the questions asks you which transport protocols are supported, NOT which subscribers – therefore AWS Lambda is not supported.\
> Amazon SNS supports notifications over multiple transport protocols:
> - HTTP/HTTPS – subscribers specify a URL as part of the subscription registration.
> - Email/Email-JSON – messages are sent to registered addresses as email (text-based or JSON-object).
> - SQS – users can specify an SQS standard queue as the endpoint.
> - SMS – messages are sent to registered phone numbers as SMS text messages.


### Q49
> An application running in an on-premise data center writes data to a MySQL database. A Solutions Architect is re-architecting the application and plans to move the database layer into the AWS cloud on Amazon RDS. The application layer will run in the on-premise data center.\
> What must be done to connect the application to the RDS database via the Internet? (choose 2)

public subnet for the RDS and a security group?

> **Explanation:**\
> When you create the RDS instance, you need to select the option to make it publicly accessible. A security group will need to be created and assigned to the RDS instance to allow access from the public IP address of your application (or firewall).


### Q50
> There has been an increase in traffic to an application that writes data to an Amazon DynamoDB database. Thousands of random tables reads occur per second and low-latency is required.\
> What can a Solutions Architect do to improve performance for the reads without negatively impacting the rest of the application?

DynamoDB accelerator - it's another cacheing layer.

> **Explanation:**\
> DAX is a DynamoDB-compatible caching service that enables you to benefit from fast in-memory performance for demanding applications. DAX addresses three core scenarios:
> 1. As an in-memory cache, DAX reduces the response times of eventually consistent read workloads by an order of magnitude from single-digit milliseconds to microseconds.
> 2. DAX reduces operational and application complexity by providing a managed service that is API-compatible with DynamoDB. Therefore, it requires only minimal functional changes to use with an existing application.
> 3. For read-heavy or bursty workloads, DAX provides increased throughput and potential operational cost savings by reducing the need to overprovision read capacity units. This is especially beneficial for applications that require repeated reads for individual keys.
> 
> DynamoDB accelerator is the best solution for caching the reads and delivering them at extremely low latency.


### Q51
> A company has multiple AWS accounts for several environments (Prod, Dev, Test etc.). A Solutions Architect would like to copy an Amazon EBS snapshot from DEV to PROD. The snapshot is from an EBS volume that was encrypted with a custom key.\
> What steps must be performed to share the encrypted EBS snapshot with the Prod account? (choose 2)

share the custom key, modify permissions?

> **Explanation:**\
> When an EBS volume is encrypted with a custom key you must share the custom key with the PROD account. You also need to modify the permissions on the snapshot to share it with the PROD account. The PROD account must copy the snapshot before they can then create volumes from the snapshot.\
> Note that you cannot share encrypted volumes created using a default CMK key and you cannot change the CMK key that is used to encrypt a volume.


### Q52
> A Solutions Architect is deploying a high performance computing (HPC) application on Amazon EC2 instances. The application requires extremely low inter-instance latency.\
> How should the instances be deployed for BEST performance?

EFA adapter and a cluster placement group.

> **Explanation:**\
> It is recommended to use either enhanced networking or an Elastic Fabric Adapter (EFA) for the nodes of an HPC application. This will assist with decreasing latency. Additionally, a cluster placement group packs instances close together inside an Availability Zone.\
> Using a cluster placement group enables workloads to achieve the low-latency network performance necessary for tightly-coupled node-to-node communication that is typical of HPC applications.\
> The table below helps you to understand the key differences between the different placement group options:
> 
> ![Q52 placement groups options](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-14-01-817a833ddcb4d86fdaff2f66faee59f2.png)


### Q53
> A new department will begin using AWS services an AWS account and a Solutions Architect needs to create an authentication and authorization strategy.
> Select the correct statements regarding IAM groups? (choose 2)

an IAM group is not an identity, but it can be used to assign permissions to users.

> **Explanation:**\
> An IAM group is a collection of IAM users. Groups let you specify permissions for multiple users, which can make it easier to manage the permissions for those users.\
> The following facts apply to IAM Groups:
> - Groups are collections of users and have policies attached to them.
> - A group is not an identity and cannot be identified as a principal in an IAM policy.
> - Use groups to assign permissions to users.
> - IAM groups cannot be used to group EC2 instances.
> - Only users and services can assume a role to take on permissions (not groups).


### Q54
> An operations team would like to be notified if an RDS database exceeds certain metric thresholds.\
> How can a Solutions Architect automate this process for the operations team?

CloudWatch and SNS to send an e-mail.

> **Explanation:**\
> You can create a CloudWatch alarm that watches a single CloudWatch metric or the result of a math expression based on CloudWatch metrics. The alarm performs one or more actions based on the value of the metric or expression relative to a threshold over a number of time periods.\
> The action can be an Amazon EC2 action, an Amazon EC2 Auto Scaling action, or a notification sent to an Amazon SNS topic. SNS can be configured to send an email notification.


### Q55
> A Solutions Architect is creating an application design with several components that will be publicly addressable. The Architect would like to use Alias records.\
> Using Route 53 Alias records what targets can you specify? (choose 2)

probably a cloudfront distribution, ~~maybe EFS file system?~~

> **Explanation:**\
> Alias records are used to map resource record sets in your hosted zone to Amazon Elastic Load Balancing load balancers, API Gateway custom regional APIs and edge-optimized APIs, CloudFront Distributions, AWS Elastic Beanstalk environments, Amazon S3 buckets that are configured as website endpoints, Amazon VPC interface endpoints, and to other records in the same Hosted Zone.


### Q56
> A Solutions Architect is conducting an audit and needs to query several properties of EC2 instances in a VPC.\
> What two methods are available for accessing and querying the properties of an EC2 instance such as instance ID, public keys and network interfaces? (choose 2)

querying 'latest/meta-data' from inside the EC2. maybe meta-data query tool?

> **Explanation:**\
> This information is stored in the instance metadata on the instance. You can access the instance metadata through a URI or by using the Instance Metadata Query tool.\
> The instance metadata is available at http://169.254.169.254/latest/meta-data.\
> The Instance Metadata Query tool allows you to query the instance metadata without having to type out the full URI or category names.


### Q57
> An application is running on EC2 instances in a private subnet of an Amazon VPC. A Solutions Architect would like to connect the application to Amazon API Gateway. For security reasons, it is necessary to ensure that no traffic traverses the Internet and to ensure all traffic uses private IP addresses only.\
> How can this be achieved?

interface VPC endpoint?

> **Explanation:**\
> An Interface endpoint uses AWS PrivateLink and is an elastic network interface (ENI) with a private IP address that serves as an entry point for traffic destined to a supported service. Using PrivateLink you can connect your VPC to supported AWS services, services hosted by other AWS accounts (VPC endpoint services), and supported AWS Marketplace partner services.


### Q58
> A Solutions Architect needs a storage solution for a fleet of Linux web application servers. The solution should provide a file system interface and be able to support millions of files.\
> Which AWS service should the Architect choose?

EFS

> **Explanation:**\
> The Amazon Elastic File System (EFS) is the only storage solution in the list that provides a file system interface. It also supports millions of files as requested.


### Q59
> A Solutions Architect is creating a solution for an application that must be deployed on Amazon EC2 hosts that are dedicated to the client. Instance placement must be automatic and billing should be per instance.\
> Which type of EC2 deployment model should be used?

~~Dedicated host?~~

> **Explanation:**\
> Dedicated Instances are Amazon EC2 instances that run in a VPC on hardware that’s dedicated to a single customer. Your Dedicated instances are physically isolated at the host hardware level from instances that belong to other AWS accounts. Dedicated instances allow automatic instance placement and billing is per instance.


### Q60
> An application you manage runs a number of components using a micro-services architecture. Several ECS container instances in your ECS cluster are displaying as disconnected. The ECS instances were created from the Amazon ECS-Optimized AMI.\
> What steps might you take to troubleshoot the issue? (choose 2) 

~~fargate~~ +  IAM instance profile permissions?

> **Explanation:**\
> The ECS container agent is included in the Amazon ECS optimized AMI and can also be installed on any EC2 instance that supports the ECS specification (only supported on EC2 instances). Therefore, you don’t need to verify that the agent is installed.\
> You need to verify that the installed agent is running and that the IAM instance profile has the necessary permissions applied.
> Troubleshooting steps for containers include:
> - Verify that the Docker daemon is running on the container instance.
> - Verify that the Docker Container daemon is running on the container instance.
> - Verify that the container agent is running on the container instance.
> - Verify that the IAM instance profile has the necessary permissions.


### Q61
> A Solutions Architect is writing some code that uses an AWS Lambda function and would like to enable the function to connect to an Amazon ElastiCache cluster within an Amazon VPC in the same AWS account.\
> What VPC-specific information must be included in the function to enable this configuration? (choose 2)

~~peering id + route table id?~~

> **Explanation:**\
> To enable your Lambda function to access resources inside your private VPC, you must provide additional VPC-specific configuration information that includes VPC subnet IDs and security group IDs. AWS Lambda uses this information to set up elastic network interfaces (ENIs) that enable your function.


### Q62
> A Solutions Architect is attempting to clean up unused EBS volumes and snapshots to save some space and cost.\
> How many of the most recent snapshots of an EBS volume need to be maintained to guarantee that you can recreate the full EBS volume from the snapshot?

the most recent, let AWS handle combining the snapshots.

> **Explanation:**\
> Snapshots capture a point-in-time state of an instance. If you make periodic snapshots of a volume, the snapshots are incremental, which means that only the blocks on the device that have changed after your last snapshot are saved in the new snapshot.\
> Even though snapshots are saved incrementally, the snapshot deletion process is designed so that you need to retain only the most recent snapshot in order to restore the volume.


### Q63
> A company runs an API on a Linux server in their on-premises data center. The company are planning to migrate the API to the AWS cloud. The company require a highly available, scalable and cost-effective solution.\
> What should a Solutions Architect recommend?

API gateway + lambda.

> **Explanation:**\
> The best option is to use a fully serverless solution. This will provide high availability, scalability and be cost-effective. The components for this would be Amazon API Gateway for hosting the API and AWS Lambda for running the backend.\
> As you can see in the image below, API Gateway can be the frontend for multiple backend services:
> 
> ![Q63 Api gateway](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-09-23-7dad959e9a620f37463b0048341769bc.png)


### Q64
> A Solutions Architect manages multiple Amazon RDS MySQL databases. To improve security, the Solutions Architect wants to enable secure user access with short-lived credentials.\
> How can these requirements be met?

~~maybe Security Token service?~~

> **Explanation:**\
> With MySQL, authentication is handled by AWSAuthenticationPlugin—an AWS-provided plugin that works seamlessly with IAM to authenticate your IAM users. Connect to the DB instance and issue the CREATE USER statement, as shown in the following example.\
> `CREATE USER jane_doe IDENTIFIED WITH AWSAuthenticationPlugin AS 'RDS';`\
> The IDENTIFIED WITH clause allows MySQL to use the AWSAuthenticationPlugin to authenticate the database account (jane_doe). The AS 'RDS' clause refers to the authentication method, and the specified database account should have the same name as the IAM user or role. In this example, both the database account and the IAM user or role are named jane_doe.


### Q65
> A Python application is currently running on Amazon ECS containers using the Fargate launch type. An ALB has been created with a Target Group that routes incoming connections to the ECS-based application. The application will be used by consumers who will authenticate using federated OIDC compliant Identity Providers such as Google and Facebook. The users must be securely authenticated on the front-end before they access the secured portions of the application.\
> How can this be configured using an ALB?

Use Cognito

> **Explanation:**\
> ALB supports authentication from OIDC compliant identity providers such as Google, Facebook and Amazon. It is implemented through an authentication action on a listener rule that integrates with Amazon Cognito to create user pools.\
> SAML can be used with Amazon Cognito but this is not the only option.


</details>

## Exam 6
<!-- <details> -->
<summary>
65 questions
</summary>


### Q01
> An Amazon EBS-backed EC2 instance has been launched. A requirement has come up for some high-performance ephemeral storage.\
> How can a Solutions Architect add a new instance store volume?

only when we launch the instance.

> **Explanation:**\
> You can specify the instance store volumes for your instance only when you launch an instance. You can’t attach instance store volumes to an instance after you’ve launched it.

### Q02
> A web application runs on a series of Amazon EC2 instances behind an Application Load Balancer (ALB). A Solutions Architect is updating the configuration with a health check and needs to select the protocol to use.\
> What options are available? (choose 2)

probably http and https.

> **Explanation:**\
An Application Load Balancer periodically sends requests to its registered targets to test their status. These tests are called health checks.\
> Each load balancer node routes requests only to the healthy targets in the enabled Availability Zones for the load balancer. Each load balancer node checks the health of each target, using the health check settings for the target groups with which the target is registered. After your target is registered, it must pass one health check to be considered healthy. After each health check is completed, the load balancer node closes the connection that was established for the health check.\
> If a target group contains only unhealthy registered targets, the load balancer nodes route requests across its unhealthy targets.\
> For an ALB the possible protocols are HTTP and HTTPS. The default is the HTTP protocol.


### Q03
> A distribution method is required for some static files. The requests will mainly be GET requests and a high volume of GETs is expected, often exceeding 2000 per second. The files are currently stored in an S3 bucket.
> According to AWS best practices, how can performance be optimized?

integrate CloudFront with S3 for caching

> **Explanation:**\
> Amazon S3 automatically scales to high request rates. For example, your application can achieve at least 3,500 PUT/POST/DELETE and 5,500 GET requests per second per prefix in a bucket. There are no limits to the number of prefixes in a bucket.\
> If your workload is mainly sending GET requests, in addition to the preceding guidelines, you should consider using Amazon CloudFront for performance optimization. By integrating CloudFront with Amazon S3, you can distribute content to your users with low latency and a high data transfer rate.


### Q04
> A Solutions Architect needs to run a PowerShell script on a fleet of Amazon EC2 instances running Microsoft Windows. The instances have already been launched in an Amazon VPC.\
> What tool can be run from the AWS Management Console that to execute the script on all target EC2 instances?

maybe RunCommand? ~~maybe CodeDeploy? OPS works?~~

> **Explanation:**\
> Run Command is designed to support a wide range of enterprise scenarios including installing software, running ad hoc scripts or Microsoft PowerShell commands, configuring Windows Update settings, and more.\
> Run Command can be used to implement configuration changes across Windows instances on a consistent yet ad hoc basis and is accessible from the AWS Management Console, the AWS Command Line Interface (CLI), the AWS Tools for Windows PowerShell, and the AWS SDKs.


### Q05
> An application uses an Amazon RDS database and Amazon EC2 instances in a web tier. The web tier instances must not be directly accessible from the internet to improve security.\
> How can a Solutions Architect meet these requirements?

instances in a private subnet, ELB?

> **Explanation:**\
> To prevent direct connectivity to the EC2 instances from the internet you can deploy your EC2 instances in a private subnet and have the ELB in a public subnet. To configure this you must enable a public subnet in the ELB that is in the same AZ as the private subnet.
> 
> ![Q5 ALB and private subnets](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-29-50-1c44f7977941f91bda9d48cd258911c6.png)


### Q06
> A Solutions Architect enabled Access Logs on an Application Load Balancer (ALB) and needs to process the log files using a hosted Hadoop service.
> What configuration changes and services can be leveraged to deliver this requirement?

move to S3 and use EMR.

> **Explanation:**\
> Access Logs can be enabled on ALB and configured to store data in an S3 bucket. Amazon EMR is a web service that enables businesses, researchers, data analysts, and developers to easily and cost-effectively process vast amounts of data. EMR utilizes a hosted Hadoop framework running on Amazon EC2 and Amazon S3.


### Q07
> A Solutions Architect is launching an Amazon EC2 instance with multiple attached volumes by modifying the block device mapping.\
> Which block device can be specified in a block device mapping to be used with an EC2 instance? (choose 2)

Instance Store and EBS.

> **Explanation:**\
> Each instance that you launch has an associated root device volume, either an Amazon EBS volume or an instance store volume.\
> You can use block device mapping to specify additional EBS volumes or instance store volumes to attach to an instance when it’s launched. You can also attach additional EBS volumes to a running instance.\
> You cannot use a block device mapping to specify a snapshot, EFS volume or S3 bucket.


### Q08
> A company has launched a multi-tier application architecture. The web tier and database tier run on Amazon EC2 instances in private subnets within the same Availability Zone.\
> Which combination of steps should a Solutions Architect take to add high availability to this architecture? (Select TWO.)

migrage to RDS multi-AZ, use EC2 autoscaling group across multiple AZ.

> **Explanation:**\
> The Solutions Architect can use Auto Scaling group across multiple AZs with an ALB in front to create an elastic and highly available architecture. Then, migrate the database to an Amazon RDS multi-AZ deployment to create HA for the database tier. This results in a fully redundant architecture that can withstand the failure of an availability zone.


### Q09
> A Solutions Architect has created an AWS Organization with several AWS accounts. Security policy requires that use of specific API actions are limited across all accounts. The Solutions Architect requires a method of centrally controlling these actions.\
> What is the SIMPLEST method of achieving the requirements?

service control policy.

> **Explanation:**\
> Service control policies (SCPs) offer central control over the maximum available permissions for all accounts in your organization, allowing you to ensure your accounts stay within your organization’s access control guidelines.\
> In the example below, a policy in OU1 restricts all users from launching EC2 instance types other than a t2.micro:
> 
> ![Q9 Service control policy](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-33-24-3942465af0c21392a215faf713e45486.png)


### Q10
> A customer is deploying services in a hybrid cloud model. The customer has mandated that data is transferred directly between cloud data centers, bypassing ISPs.\
> Which AWS service can be used to enable hybrid cloud connectivity?

Direct Connect?

> **Explanation:**\
> With AWS Direct Connect, you can connect to all your AWS resources in an AWS Region, transfer your business-critical data directly from your datacenter, office, or colocation environment into and from AWS, bypassing your Internet service provider and removing network congestion.
> 
> ![Q10 Direct connect](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-25-29-2c6aa3f61ef4f4462a7e9088ceb29acf.png)


### Q11
> The Solutions Architect in charge of a critical application must ensure the Amazon EC2 instances are able to be launched in another AWS Region in the event of a disaster.\
> What steps should the Solutions Architect take? (Select TWO.)

create AMIs, copy them, launch in another region.

> **Explanation:**\
You can create AMIs of the EC2 instances and then copy them across Regions. This provides a point-in-time copy of the state of the EC2 instance in the remote Region.\
> Once you’ve created AMIs of EC2 instances and copied them to the second Region, you can then launch the EC2 instances from the AMIs in that Region.\
> This is a good DR strategy as you have moved stateful EC2 instances to another Region.


### Q12
> A company is investigating ways to analyze and process large amounts of data in the cloud faster, without needing to load or transform the data in a data warehouse. The data resides in Amazon S3.\
> Which AWS services would allow the company to query the data in place? (choose 2)

redshift spectrum. S3 select?

> **Explanation:**\
> Amazon S3 Select is designed to help analyze and process data within an object in Amazon S3 buckets, faster and cheaper. It works by providing the ability to retrieve a subset of data from an object in Amazon S3 using simple SQL expressions\
> Amazon Redshift Spectrum allows you to directly run SQL queries against exabytes of unstructured data in Amazon S3. No loading or transformation is required.


### Q13
> A large quantity of data that is rarely accessed is being archived onto Amazon Glacier. Your CIO wants to understand the resilience of the service.\
> Which of the statements below is correct about Amazon Glacier storage?  (choose 2) 

9-11 durability (like all S3 buckets),~~replicated globably?~~ or one-region resilient?

> **Explanation:**\
> Glacier is designed for durability of 99.999999999% of objects across multiple Availability Zones. Data is resilient in the event of one entire Availability Zone destruction. Glacier supports SSL for data in transit and encryption of data at rest. Glacier is extremely low cost and is ideal for long-term archival.


### Q14
> One of the departments in a company has been generating a large amount of data on Amazon S3 and costs are increasing. Data older than 90 days is rarely accessed but must be retained for several years. If this data does need to be accessed at least 24 hours notice is provided.\
> How can a Solutions Architect optimize the costs associated with storage of this data whilst ensuring it is accessible if required?

lifecycle policy to move to glacier.

> **Explanation:**\
> To manage your objects so that they are stored cost effectively throughout their lifecycle, configure their lifecycle. A lifecycle configuration is a set of rules that define actions that Amazon S3 applies to a group of objects. Transition actions define when objects transition to another storage class.\
> For example, you might choose to transition objects to the STANDARD_IA storage class 30 days after you created them, or archive objects to the GLACIER storage class one year after creating them.\
> GLACIER retrieval times:
> - Standard retrieval is 3-5 hours which is well within the requirements here.
> - You can use Expedited retrievals to access data in 1 – 5 minutes.
> - You can use Bulk retrievals to access up to petabytes of data in approximately 5 – 12 hours.


### Q15
> A company has a fleet of Amazon EC2 instances behind an Elastic Load Balancer (ELB) that are a mixture of c4.2xlarge instance types and c5.large instances. The load on the CPUs on the c5.large instances has been very high, often hitting 100% utilization, whereas the c4.2xlarge instances have been performing well.\
> What should a Solutions Architect recommend to resolve the performance issues?

~~use weight routing policy?~~

> **Explanation:**\
> The 2xlarge instance type provides more CPUs. The best answer is to use this instance type for all instances as the CPU utilization has been lower.


### Q16
> A fleet of Amazon EC2 instances running Linux will be launched in an Amazon VPC. An application development framework and some custom software must be installed on the instances. The installation will be initiated using some scripts.\
> What feature enables a Solutions Architect to specify the scripts the software can be installed during the EC2 instance launch?

user-data is when we stick the the start-up scripts.

> **Explanation:**\
> When you launch an instance in Amazon EC2, you have the option of passing user data to the instance that can be used to perform common automated configuration tasks and even run scripts after the instance starts. You can pass two types of user data to Amazon EC2: shell scripts and cloud-init directives.\
> User data is data that is supplied by the user at instance launch in the form of a script and is limited to 16KB.

### Q17
> A Solutions Architect has created a new Network ACL in an Amazon VPC. No rules have been created.\
> Which of the statements below are correct regarding the default state of the Network ACL? (choose 2)

default inbound rule denying everything. ~~maybe default outbound rule allowing all?~~

> **Explanation:**\
> A VPC automatically comes with a default network ACL which allows all inbound/outbound traffic. A custom NACL denies all traffic both inbound and outbound by default.\
> Network ACL’s function at the subnet level and you can have permit and deny rules. Network ACLs have separate inbound and outbound rules and each rule can allow or deny traffic.\
> Network ACLs are stateless so responses are subject to the rules for the direction of traffic. NACLs only apply to traffic that is ingress or egress to the subnet not to traffic within the subnet.

### Q18
> An Amazon EC2 instance is generating very high packets-per-second and performance of the application stack is being impacted. A Solutions Architect needs to determine a resolution to the issue that results in improved performance.\
> Which action should the Architect take?

RAID-1 is for durability, multiple IP address won't help, so maybe enhanced networking?

> **Explanation:**\
> Enhanced networking provides higher bandwidth, higher packet-per-second (PPS) performance, and consistently lower inter-instance latencies. If your packets-per-second rate appears to have reached its ceiling, you should consider moving to enhanced networking because you have likely reached the upper thresholds of the VIF driver. It is only available for certain instance types and only supported in VPC. You must also launch an HVM AMI with the appropriate drivers.\
> AWS currently supports enhanced networking capabilities using SR-IOV. SR-IOV provides direct access to network adapters, provides higher performance (packets-per-second) and lower latency.


### Q19
> A large multi-national client has requested a design for a multi-region database. The master database will be in the EU (Frankfurt) region and databases will be located in 4 other regions to service local read traffic. The database should be a managed service including the replication.\
> The solution should be cost-effective and secure.\
> Which AWS service can deliver these requirements?

RDS with cross region read replicas.

> **Explanation:**\
> Amazon RDS Read replicas are used for read heavy databases and the replication is asynchronous. Read replicas are used for workload sharing and offloading. Read replicas can be in another region. This solution will enable better performance for users in the other AWS regions for database queries and is a managed service.


### Q20
> An application makes calls to a REST API running on Amazon EC2 instances behind an Application Load Balancer (ALB). Most API calls complete quickly. However, a single endpoint is making API calls that require much longer to complete and this is introducing overall latency into the system.\
> What steps can a Solutions Architect take to minimize the effects of the long-running API calls?

probably SQS.

> **Explanation:**\
> An Amazon Simple Queue Service (SQS) can be used to offload and decouple the long-running requests. They can then be processed asynchronously by separate EC2 instances. This is the best way to reduce the overall latency introduced by the long-running API call.


### Q21
> An on-premise data center will be connected to an Amazon VPC by a hardware VPN that has public and VPN-only subnets. The security team has requested that traffic hitting public subnets on AWS that’s destined to on-premise applications must be directed over the VPN to the corporate firewall.\
> How can this be achieved?

in the public subnet, configure routetable to the virtual private gateway.

> **Explanation:**\
> Route tables determine where network traffic is directed. In your route table, you must add a route for your remote network and specify the virtual private gateway as the target. This enables traffic from your VPC that’s destined for your remote network to route via the virtual private gateway and over one of the VPN tunnels. You can enable route propagation for your route table to automatically propagate your network routes to the table for you.


### Q22
> Three AWS accounts are owned by the same company but in different regions. Account Z has two AWS Direct Connect connections to two separate company offices. Accounts A and B require the ability to route across account Z’s Direct Connect connections to each company office. A Solutions Architect has created an AWS Direct Connect gateway in account Z.\
> How can the required connectivity be configured?

~~PrivateLink connections to ENI?~~

> **Explanation:**\
> You can associate an AWS Direct Connect gateway with either of the following gateways:
> - A transit gateway when you have multiple VPCs in the same Region.
> - A virtual private gateway.
> 
> In this case account Z owns the Direct Connect gateway so a VPG in accounts A and B must be associated with it to enable this configuration to work. After Account Z accepts the proposals, Account A and Account B can route traffic from their virtual private gateway to the Direct Connect gateway.
> 
> ![Q22 direct connect with many accounts](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-27-50-f7c361a3527b03f1a25861bfd2e7664b.png)


### Q23
> A Solutions Architect needs to upload a large (2GB) file to an S3 bucket.\
> What is the recommended way to upload a single large file to an S3 bucket?

multi-part upload.

> **Explanation:**\
> In general, when your object size reaches 100 MB, you should consider using multipart uploads instead of uploading the object in a single operation.


### Q24
> A Solutions Architect has created an AWS account and selected the Asia Pacific (Sydney) region. Within the default VPC there is a default security group.\
> What settings are configured within this security group by default? (choose 2)

~~default of default? maybe all is allowed?~~

> **Explanation:**\
> Default security groups have inbound allow rules (allowing traffic from within the group) whereas custom security groups do not have inbound allow rules (all inbound traffic is denied by default). All outbound traffic is allowed by default in custom and default security groups.


### Q25
> An organization in the agriculture sector is deploying sensors and smart devices around factory plants and fields. The devices will collect information and send it to cloud applications running on AWS.\
> Which AWS service will securely connect the devices to the cloud applications?

Glue? IotCore?

> **Explanation:**\
> AWS IoT Core is a managed cloud service that lets connected devices easily and securely interact with cloud applications and other devices. AWS IoT Core can support billions of devices and trillions of messages, and can process and route those messages to AWS endpoints and to other devices reliably and securely.
> 
> ![Q25 AWS IoT-Core](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-24-14-48d1a8665d7d70aa01818799e7168c7e.png)


### Q26
> A government agency is using CloudFront for a web application that receives personally identifiable information (PII) from citizens.\
> What feature of CloudFront applies an extra level of encryption at CloudFront edge locations to ensure the PII data is secured end-to-end?   

maybe field level encryption?

> **Explanation:**\
> With Amazon CloudFront, you can enforce secure end-to-end connections to origin servers by using HTTPS. Field-level encryption adds an additional layer of security that lets you protect specific data throughout system processing so that only certain applications can see it.\
> Field-level encryption allows you to enable your users to securely upload sensitive information to your web servers. The sensitive information provided by your users is encrypted at the edge, close to the user, and remains encrypted throughout your entire application stack. This encryption ensures that only applications that need the data—and have the credentials to decrypt it—are able to do so.
> 
> ![Q26 cloud front field level encryption](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-26-41-7ef0752935b0a796e94d9b3eead1f7ef.png)


### Q27
> A Solutions Architect has created a VPC and is in the process of formulating the subnet design. The VPC will be used to host a two-tier application that will include Internet facing web servers, and internal-only DB servers. Zonal redundancy is required.\
> How many subnets are required to support this requirement?

probably 4, if we use two AZ.

> **Explanation:**\
> Zonal redundancy indicates that the architecture should be split across multiple Availability Zones. Subnets are mapped 1:1 to AZs.\
> A public subnet should be used for the Internet-facing web servers and a separate private subnet should be used for the internal-only DB servers. Therefore you need 4 subnets – 2 (for redundancy) per public/private subnet.


### Q28
> A company needs to capture detailed information about all HTTP requests that are processed by their Internet facing Application Load Balancer (ALB). The company requires information on the requester, IP address, and request type for analyzing traffic patterns to better understand their customer base.\
> Which actions should a Solutions Architect recommend?

maybe Access logs?

> **Explanation:**\
> You can enable access logs on the ALB and this will provide the information required including requester, IP, and request type. Access logs are not enabled by default. You can optionally store and retain the log files on S3.


### Q29
> A company has multiple Amazon VPCs that are peered with each other. The company would like to use a single Elastic Load Balancer (ELB) to route traffic to multiple EC2 instances in peered VPCs within the same region.\
> How can this be achieved?

~~I think route53 is needed.~~

> **Explanation:**\
> With ALB and NLB IP addresses can be used to register:
> - Instances in a peered VPC.
> - AWS resources that are addressable by IP address and port.
> - On-premises resources linked to AWS through Direct Connect or a VPN connection.


### Q30
> A company runs a web-based application that uses Amazon EC2 instances for the web front-end and Amazon RDS for the database back-end. The web application writes transaction log files to an Amazon S3 bucket and the quantity of files is becoming quite large. It is acceptable to retain the most recent 60 days of log files and permanently delete the rest.\
> Which action can a Solutions Architect take to enable this to happen automatically?

S3 lifecycle policy.

> **Explanation:**\
> To manage your objects so that they are stored cost effectively throughout their lifecycle, configure their Amazon S3 Lifecycle. An S3 Lifecycle configuration is a set of rules that define actions that Amazon S3 applies to a group of objects. There are two types of actions:
> - Transition actions—Define when objects transition to another storage class. For example, you might choose to transition objects to the S3 Standard-IA storage class 30 days after you created them, or archive objects to the S3 Glacier storage class one year after creating them.
> - Expiration actions—Define when objects expire. Amazon S3 deletes expired objects on your behalf.


### Q31
> A customer has requested some advice on how to implement security measures in their Amazon VPC. The client has recently been the victim of some hacking attempts. The client wants to implement measures to mitigate further threats. The client has explained that the attacks always come from the same small block of IP addresses.\
> What would be a quick and easy measure to help prevent further attacks?

network ACL can have deny rules.

> **Explanation:**\
> With NACLs you can have permit and deny rules. Network ACLs contain a numbered list of rules that are evaluated in order from the lowest number until the explicit deny. Network ACLs have separate inbound and outbound rules and each rule can allow or deny traffic.
> 
> ![Q31 NACL and security groups](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-17-36-ac4cd4adb1c69a741c1683a12b4a2344.png)


### Q32
> An application runs on EC2 instances in a private subnet behind an Application Load Balancer in a public subnet. The application is highly available and distributed across multiple AZs. The EC2 instances must make API calls to an internet-based service.\
> How can the Solutions Architect enable highly available internet connectivity?

highly available internet probably means NAT gateway. I think we need to update the route tables for each private subnet,

> **Explanation:**\
> The only solution presented that actually works is to create a NAT gateway in the public subnet of each AZ. They must be created in the public subnet as they gain public IP addresses and use an internet gateway for internet access.\
> The route tables in the private subnets must then be configured with a route to the NAT gateway and then the EC2 instances will be able to access the internet (subject to security group configuration).

### Q33
> A Solutions Architect has created a new security group in an Amazon VPC. No rules have been created.\
> Which of the statements below are correct regarding the default state of the security group? (choose 2)

no inbound rules, so all traffic is denied. ~~outbound rule to the internet gateway?~~

> **Explanation:**\
> Custom security groups do not have inbound allow rules (all inbound traffic is denied by default) whereas default security groups do have inbound allow rules (allowing traffic from within the group). All outbound traffic is allowed by default in both custom and default security groups.\
> Security groups act like a stateful firewall at the instance level. Specifically security groups operate at the network interface level of an EC2 instance. You can only assign permit rules in a security group, you cannot assign deny rules and there is an implicit deny rule at the end of the security group. All rules are evaluated until a permit is encountered or continues until the implicit deny. You can create ingress and egress rules.


### Q34
> An application uses a MySQL database running on an Amazon EC2 instance. The application generates high I/O and constant writes to a single table on the database.\
> Which Amazon EBS volume type will provide the MOST consistent performance and low latency?

provisoned IOPS ssd?

> **Explanation:**\
> The Provisioned IOPS SSD (io1) volume type will offer the most consistent performance and can be configured with the amount of IOPS required by the application. It will also provide the lowest latency of the options presented.
> 
> ![Q34 SSD vs HDD](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-30-39-f006b5d6b4ffa97e2db382aff552857e.png)


### Q35
> The load on a MySQL database running on Amazon EC2 is increasing and performance has been impacted.\
> Which of the options below would help to increase storage performance? (choose 2)

Provisioned IOPS EBS volumes and Optimized EBS instances?

> **Explanation:**\
> EBS optimized instances provide dedicated capacity for Amazon EBS I/O. EBS optimized instances are designed for use with all EBS volume types.\
> Provisioned IOPS EBS volumes allow you to specify the amount of IOPS you require up to 50 IOPS per GB. Within this limitation you can therefore choose to select the IOPS required to improve the performance of your volume.\
> RAID can be used to increase IOPS, however RAID 1 does not. For example:
> - RAID 0 = 0 striping – data is written across multiple disks and increases performance but no redundancy.
> - RAID 1 = 1 mirroring – creates 2 copies of the data but does not increase performance, only redundancy.
>
> HDD, Cold – (SC1) provides the lowest cost storage and low performance


### Q36
> A company requires an Elastic Load Balancer (ELB) for an application they are planning to deploy on AWS. The application requires extremely high throughput and extremely low latencies. The connections will be made using the TCP protocol and the ELB must support load balancing to multiple ports on an instance.\
> Which ELB would should the company use?

network load balancer?

> **Explanation:**\
> The Network Load Balancer operates at the connection level (Layer 4), routing connections to targets – Amazon EC2 instances, containers and IP addresses based on IP protocol data. It is architected to handle millions of requests/sec, sudden volatile traffic patterns and provides extremely low latencies.\
> The NLB provides high throughput and extremely low latencies and is designed to handle traffic as it grows and can load balance millions of requests/second. NLB also supports load balancing to multiple ports on an instance.


### Q37
> An application is generating a large amount of clickstream events data that is being stored on S3. The business needs to understand customer behavior and want to run complex analytics queries against the data.\
> Which AWS service can be used for this requirement?

RedShift (spectrum?)

> **Explanation:**\
> Amazon Redshift is a fast, fully managed data warehouse that makes it simple and cost-effective to analyze all your data using standard SQL and existing Business Intelligence (BI) tools.\
> RedShift is used for running complex analytic queries against petabytes of structured data, using sophisticated query optimization, columnar storage on high-performance local disks, and massively parallel query execution.\
> With RedShift you can load data from Amazon S3 and perform analytics queries. RedShift Spectrum can analyze data directly in Amazon S3, but was not presented as an option.

### Q38
> An on-premises server runs a MySQL database and will be migrated to the AWS Cloud. The company require a managed solution that supports high availability and automatic failover in the event of the outage of an Availability Zone (AZ).\
> Which solution is the BEST fit for these requirements?

migrate directly to multi-AZ RDS

> **Explanation:**\
> The AWS DMS service can be used to directly migrate the MySQL database to an Amazon RDS Multi-AZ deployment. The entire process can be online and is managed for you. There is no need to perform schema translation between MySQL and RDS (assuming you choose the MySQL RDS engine).
> ![Q38 database migration](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-31-42-b6a87bc2209574bb6eabad34a39f6577.png)



### Q39
> A tool needs to analyze data stored in an Amazon S3 bucket. Processing the data takes a few seconds and results are then written to another S3 bucket. Less than 256 MB of memory is needed to run the process.\
> What would be the MOST cost-effective compute solutions for this use case?

lambda

> **Explanation:**\
> AWS Lambda lets you run code without provisioning or managing servers. You pay only for the compute time you consume. Lambda has a maximum execution time of 900 seconds and memory can be allocated up to 3008 MB. Therefore, the most cost-effective solution will be AWS Lambda.


### Q40
> A Solutions Architect is deploying a production application that will use several Amazon EC2 instances and run constantly on an ongoing basis. The application cannot be interrupted or restarted.\
> Which EC2 pricing model would be best for this workload?

reserved instances

> **Explanation:**\
> In this scenario for a stable process that will run constantly on an ongoing basis RIs will be the most affordable solution.\
> RIs provide you with a significant discount (up to 75%) compared to On-Demand instance pricing. You have the flexibility to change families, OS types, and tenancies while benefitting from RI pricing when you use Convertible RIs.


### Q41
> A security officer has requested that all data associated with a specific customer is encrypted. The data resides on Elastic Block Store (EBS) volumes.\
> Which of the following statements about using EBS encryption are correct? (choose 2)

~~all instances support encryption~~, transit is also encrypted.

> **Explanation:**\
> All EBS types and all instance families support encryption but not all instance types support encryption.\
> There is no direct way to change the encryption state of a volume. Data in transit between an instance and an encrypted volume is also encrypted.


### Q42
> Some data has become corrupted in an Amazon RDS database. A Solutions Architect plans to use point-in-time restore to recover the data to the last known good configuration.
> Which of the following statements is correct about restoring an RDS database to a specific point-in-time? (choose 2)

~~the restore overwrites the existing, and the custom security groups are applied.~~

> **Explanation:**\
> You can restore a DB instance to a specific point in time, creating a new DB instance. When you restore a DB instance to a point in time, the default DB security group is applied to the new DB instance. If you need custom DB security groups applied to your DB instance, you must apply them explicitly using the AWS Management Console, the AWS CLI modify-db-instance command, or the Amazon RDS API ModifyDBInstance operation after the DB instance is available.\
> Restored DBs will always be a new RDS instance with a new DNS endpoint and you can restore up to the last 5 minutes.


### Q43
> A company needs to ensure that they can failover between AWS Regions in the event of a disaster seamlessly with minimal downtime and data loss. The applications will run in an active-active configuration.\
> Which DR strategy should a Solutions Architect recommend?

active active means multi-site? 

> **Explanation:**\
> A multi-site solution runs on AWS as well as on your existing on-site infrastructure in an active- active configuration. The data replication method that you employ will be determined by the recovery point that you choose. This is either Recovery Time Objective (the maximum allowable downtime before degraded operations are restored) or Recovery Point Objective (the maximum allowable time window whereby you will accept the loss of transactions during the DR process).


### Q44
> A company is deploying a new two-tier web application that uses EC2 web servers and a DynamoDB database backend. An Internet facing ELB distributes connections between the web servers.\
> The Solutions Architect has created a security group for the web servers and needs to create a security group for the ELB.\
> What rules should be added? (choose 2)

inbound http\s rule with 0.0.0.0/0 to get all internet access. outbound http\s rule to vpc cidr?

> **Explanation:**\
> An inbound rule should be created for the relevant protocols (HTTP/HTTPS) and the source should be set to any address (0.0.0.0/0).\
> The outbound rule should forward the relevant protocols (HTTP/HTTPS) and the destination should be set to the web server security group.\
> Note that on the web server security group you’d want to add an Inbound rule allowing HTTP/HTTPS from the ELB security group.


### Q45
> A Solutions Architect has logged into an Amazon EC2 Linux instance using SSH and needs to determine a few pieces of information including what IAM role is assigned, the instance ID and the names of the security groups that are assigned to the instance.\
> From the options below, what would be the best source of this information?

metadata.

> **Explanation:**\
> Instance metadata is data about your instance that you can use to configure or manage the running instance. Instance metadata is divided into categories, for example, host name, events, and security groups.\
> Instance metadata is available at http://169.254.169.254/latest/meta-data.


### Q46
> A company runs an application on premises that stores a large quantity of semi-structured data using key-value pairs. The application code will be migrated to AWS Lambda and a highly scalable solution is required for storing the data.\
> Which datastore will be the best fit for these requirements?

key value pairs is dynamoDB.

> **Explanation:**\
> Amazon DynamoDB is a no-SQL database that stores data using key-value pairs. It is ideal for storing large amounts of semi-structured data and is also highly scalable. This is the best solution for storing this data based on the requirements in the scenario.


### Q47
> A Solutions Architect needs to capture information about the traffic that reaches an Amazon Elastic Load Balancer. The information should include the source, destination, and protocol.\
> What is the most secure and reliable method for gathering this data?

vpc flow logs for the network interface?

> **Explanation:**\
> You can use VPC Flow Logs to capture detailed information about the traffic going to and from your Elastic Load Balancer. Create a flow log for each network interface for your load balancer. There is one network interface per load balancer subnet.


### Q48
> A web application receives order processing information from customers and places the messages on an Amazon SQS queue. A fleet of Amazon EC2 instances are configured to pick up the messages, process them, and store the results in a DynamoDB table. The current configuration has been resulting in a large number of empty responses to `ReceiveMessage` API requests.\
> A Solutions Architect needs to eliminate empty responses to reduce operational overhead. How can this be done? 

long polling

> **Explanation:**\
> The correct answer is to use Long Polling which will eliminate empty responses by allowing Amazon SQS to wait until a message is available in a queue before sending a response.\
> The problem does not relate to the order in which the messages are processed in and there are no concerns over messages being delivered more than once so it doesn’t matter whether you use a FIFO or standard queue.
> 
> Long Polling:
> - Uses fewer requests and reduces cost.
> - Eliminates false empty responses by querying all servers.
> - SQS waits until a message is available in the queue before sending a response.
> 
> Short Polling:
> - Does not wait for messages to appear in the queue.
> - It queries only a subset of the available servers for messages (based on weighted random execution).
> - Short polling is the default.
> - ReceiveMessageWaitTime is set to 0.
> 
> ![Q48 Sqs queues](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-20-24-648b8e1ff1276f787c24226b61455dd3.png)


### Q49
> The application development team in a company have developed a Java application and saved the source code in a .war file. They would like to run the application on AWS resources and are looking for a service that can handle the provisioning and management of the underlying resources it will run on.\
> Which AWS service should a Solutions Architect recommend the Developers use to upload the Java source code file?

beanStalk? codeDeploy?

> **Explanation:**\
> AWS Elastic Beanstalk can be used to quickly deploy and manage applications in the AWS Cloud. Developers upload applications and Elastic Beanstalk handles the deployment details of capacity provisioning, load balancing, auto-scaling, and application health monitoring.\
> Elastic Beanstalk supports applications developed in Go, Java, .NET, Node.js, PHP, Python, and Ruby, as well as different platform configurations for each language. To use Elastic Beanstalk, you create an application, upload an application version in the form of an application source bundle (for example, a Java .war file) to Elastic Beanstalk, and then provide some information about the application.


### Q50
> A financial services company regularly runs an analysis of the day’s transaction costs, execution reporting, and market performance. The company currently uses third-party commercial software for provisioning, managing, monitoring, and scaling the computing jobs which utilize a large fleet of EC2 instances.\
> The company is seeking to reduce costs and utilize AWS services.\
> Which AWS service could be used in place of the third-party software?

AWS batch?

> **Explanation:**\
> AWS Batch eliminates the need to operate third-party commercial or open source batch processing solutions. There is no batch software or servers to install or manage. AWS Batch manages all the infrastructure for you, avoiding the complexities of provisioning, managing, monitoring, and scaling your batch computing jobs.


### Q51
> A development team needs to run up a few lab servers on a weekend for a new project. The servers will need to run uninterrupted for a few hours.\
> Which EC2 pricing option would be most suitable?

there isn't enough here to use reserved or spot, so On-demand should be used.

> **Explanation:**\
> On-Demand pricing ensures that instances will not be terminated and is the most economical option. Use on-demand for ad-hoc requirements where you cannot tolerate interruption.


### Q52
> A Solutions Architect is designing the system monitoring and deployment layers of a serverless application. The system monitoring layer will manage system visibility through recording logs and metrics and the deployment layer will deploy the application stack and manage workload changes through a release management process.\
>  The Architect needs to select the most appropriate AWS services for these functions.\
> Which services and frameworks should be used for the system monitoring and deployment layers? (choose 2)

CloudWatch for monitoring metrics and logs, SAM to package and deploy?

> **Explanation:**\
> AWS Serverless Application Model (AWS SAM) is an extension of AWS CloudFormation that is used to package, test, and deploy serverless applications.\
> With Amazon CloudWatch, you can access system metrics on all the AWS services you use, consolidate system and application level logs, and create business key performance indicators (KPIs) as custom metrics for your specific needs.


### Q53
> An application receives a high traffic load between 7:30am and 9:30am daily. The application uses an Auto Scaling group to maintain three instances most of the time but during the peak period it requires six instances.\
> How can a Solutions Architect configure Auto Scaling to perform a daily scale-out event at 7:30am and a scale-in event at 9:30am to account for the peak load?

scheduled policy.

> **Explanation:**\
> The following scaling policy options are available:
> - Simple – maintains a current number of instances, you can manually change the ASGs min/desired/max and attach/detach instances.
> - Scheduled – Used for predictable load changes, can be a single event or a recurring schedule
> - Dynamic (event based) – scale in response to an event/alarm.
> - Step – configure multiple scaling steps in response to multiple alarms.


### Q54
> An Amazon EC2 instance behind an Elastic Load Balancer (ELB) is in the process of being de-registered.\
> Which ELB feature is used to allow existing connections to close cleanly?

maybe connection draining?

> **Explanation:**\
>Connection draining is enabled by default and provides a period of time for existing connections to close cleanly. When connection draining is in action an CLB will be in the status “InService: Instance deregistration currently in progress”.


### Q55
> A legacy application is being migrated into AWS. The application has a large amount of data that is rarely accessed. When files are accessed they are retrieved sequentially. The application will be migrated onto an Amazon EC2 instance.\
> What is the LEAST expensive EBS volume type for this use case?

~~HDD throughput optimized?~~ or cold HDD?

> **Explanation:**\
> The cold HDD (sc1) EBS volume type is the lowest cost option that is suitable for this use case. The sc1 volume type is suitable for infrequently accessed data and use cases that are oriented towards throughput like sequential data access.
> 
> ![Q55 ssd vs hdd](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-29-08-7da5cb7bfcd2de70f1d5a02ddab9fb12.png)


### Q56
> An Amazon EC2 instance has been launched into an Amazon VPC. A Solutions Architect needs to ensure that instances have both a private and public DNS hostnames.\
> Assuming settings were not changed during creation of the VPC, how will DNS hostnames be assigned by default? (choose 2)

default vpc gets both, no-deafult don't get public.

> **Explanation:**\
> When you launch an instance into a default VPC, we provide the instance with public and private DNS hostnames that correspond to the public IPv4 and private IPv4 addresses for the instance.\
> When you launch an instance into a nondefault VPC, we provide the instance with a private DNS hostname and we might provide a public DNS hostname, depending on the DNS attributes you specify for the VPC and if your instance has a public IPv4 address.\
> All other statements are incorrect with default settings.


### Q57
> The database layer of an on-premises web application is being migrated to AWS. The database currently uses an in-memory cache. A Solutions Architect must deliver a solution that supports high availability and replication for the caching layer.\
> Which service should the Solutions Architect recommend?

redis?

> **Explanation:**\
> Amazon ElastiCache Redis is an in-memory database cache and supports high availability through replicas and multi-AZ. The table below compares ElastiCache Redis with Memcached:
> 
> ![Q57 memcached and redis](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-32-43-a842e78974d591a0b9d23fb5d088de78.png)


### Q58
> Several Amazon EC2 Spot instances are being used to process messages from an Amazon SQS queue and store results in an Amazon DynamoDB table. Shortly after picking up a message from the queue AWS terminated the Spot instance. The Spot instance had not finished processing the message.\
> What will happen to the message?

message will re-appear after the visibility timeout

> **Explanation:**\
> The visibility timeout is the amount of time a message is invisible in the queue after a reader picks up the message. If a job is processed within the visibility timeout the message will be deleted. If a job is not processed within the visibility timeout the message will become visible again (could be delivered twice). The maximum visibility timeout for an Amazon SQS message is 12 hours.
> 
> ![Q58 Sqs Visibility](https://img-c.udemycdn.com/redactor/raw/2020-06-28_20-16-54-ed49b8985fe33f8c30a3f5c24800aca9.png)


### Q59
> A Solutions Architect is designing the disk configuration for an Amazon EC2 instance. The instance needs to support a MapReduce process that requires high throughput for a large dataset with large I/O sizes.\
> Which Amazon EBS volume is the MOST cost-effective solution for these requirements?

~~general purpose ssd?~~

> **Explanation:**\
> EBS Throughput Optimized HDD is good for the following use cases (and is the most cost-effective option):\
> Frequently accessed, throughput intensive workloads with large datasets and large I/O sizes, such as MapReduce, Kafka, log processing, data warehouse, and ETL workloads.\
> Throughput is measured in MB/s, and includes the ability to burst up to 250 MB/s per TB, with a baseline throughput of 40 MB/s per TB and a maximum throughput of 500 MB/s per volume.


### Q60
> An Auto Scaling group of Amazon EC2 instances behind an Elastic Load Balancer (ELB) is running in an Amazon VPC. Health checks are configured on the ASG to use EC2 status checks. The ELB has determined that an EC2 instance is unhealthy and has removed it from service. A Solutions Architect noticed that the instance is still running and has not been terminated by EC2 Auto Scaling.\
> What would be an explanation for this behavior?

~~maybe cool-down timer?~~ maybe elb health-check in the ASG?

> **Explanation:**\
> If using an ELB it is best to enable ELB health checks as otherwise EC2 status checks may show an instance as being healthy that the ELB has determined is unhealthy. In this case the instance will be removed from service by the ELB but will not be terminated by Auto Scaling\
> More information on ASG health checks:
> - By default uses EC2 status checks.
> - Can also use ELB health checks and custom health checks.
> - ELB health checks are in addition to the EC2 status checks.
> - If any health check returns an unhealthy status the instance will be terminated.
> - With ELB an instance is marked as unhealthy if ELB reports it as OutOfService
> - A healthy instance enters the InService state.
> - If an instance is marked as unhealthy it will be scheduled for replacement.
> - If connection draining is enabled, Auto Scaling waits for in-flight requests to complete or timeout before terminating instances.
> - The health check grace period allows a period of time for a new instance to warm up before performing a health check (300 seconds by default).


### Q61
> A Solutions Architect created a new IAM user account for a temporary employee who recently joined the company.\
> The user does not have permissions to perform any actions, which statement is true about newly created users in IAM?

no permissions.

> **Explanation:**\
> Every IAM user starts with no permissions.. In other words, by default, users can do nothing, not even view their own access keys. To give a user permission to do something, you can add the permission to the user (that is, attach a policy to the user). Or you can add the user to a group that has the intended permission.


### Q62
> An Amazon Elastic File System (EFS) has been created to store data that will be accessed by a large number of Amazon EC2 instances. The data is sensitive and a Solutions Architect is creating a design for security measures to protect the data. It is required that network traffic is restricted correctly based on firewall rules and access from hosts is restricted by user or group.\
> How can this be achieved with Amazon EFS? (choose 2)

~~IAM control and Network ACL?~~

> **Explanation:**\
> You can control who can administer your file system using IAM. You can control access to files and directories with POSIX-compliant user and group-level permissions. POSIX permissions allows you to restrict access from hosts by user and group. EFS Security Groups act as a firewall, and the rules you add define the traffic flow.


### Q63
> A company is transitioning their web presence into the AWS cloud. As part of the migration the company will be running a web application both on-premises and in AWS for a period of time. During the period of co-existence the client would like 80% of the traffic to hit the AWS-based web servers and 20% to be directed to the on-premises web servers.\
> What method can a Solutions Architect use to distribute traffic as requested?

route53 with weighted policy.

> **Explanation:**\
> Route 53 weighted routing policy is similar to simple but you can specify a weight per IP address. You create records that have the same name and type and assign each record a relative weight which is a numerical value that favours one IP over another (values must total 100). To stop sending traffic to a resource you can change the weight of the record to 0.


### Q64
> A company runs a streaming media service and the content is stored on Amazon S3. The media catalog server pulls updated content from S3 and can issue over 1 million read operations per second for short periods. Latency must be kept under 5ms for these updates.\
> Which solution will provide the BEST performance for the media catalog updates?

~~use cloudfront and cache S3.~~

> **Explanation:**\
> Some applications, such as media catalog updates require high frequency reads, and consistent throughput. For such applications, customers often complement S3 with an in-memory cache, such as Amazon ElastiCache for Redis, to reduce the S3 retrieval cost and to improve performance.\
> ElastiCache for Redis is a fully managed, in-memory data store that provides sub-millisecond latency performance with high throughput. ElastiCache for Redis complements S3 in the following ways:
> - Redis stores data in-memory, so it provides sub-millisecond latency and supports incredibly high requests per second.
> - It supports key/value based operations that map well to S3 operations (for example, GET/SET => GET/PUT), making it easy to write code for both S3 and ElastiCache.
> - It can be implemented as an application side cache. This allows you to use S3 as your persistent store and benefit from its durability, availability, and low cost. Your applications decide what objects to cache, when to cache them, and how to cache them.
> 
> In this example the media catalog is pulling updates from S3 so the performance between these components is what needs to be improved. Therefore, using ElastiCache to cache the content will dramatically increase the performance.


### Q65
> An Amazon DynamoDB table has a variable load, ranging from sustained heavy usage some days, to only having small spikes on others. The load is 80% read and 20% write. The provisioned throughput capacity has been configured to account for the heavy load to ensure throttling does not occur.\
> What would be the most efficient solution to optimize cost?

~~probably DAX to make use of caching and get better costs.~~

> **Explanation:**\
> Amazon DynamoDB auto scaling uses the AWS Application Auto Scaling service to dynamically adjust provisioned throughput capacity on your behalf, in response to actual traffic patterns. This is the most efficient and cost-effective solution to optimizing for cost.

</details>

##