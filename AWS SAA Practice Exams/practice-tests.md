
<!--
// cSpell:ignore DFSR
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



### Q02
> A company runs a web application that serves weather updates. The application runs on a fleet of Amazon EC2 instances in a Multi-AZ Auto scaling group behind an Application Load Balancer (ALB). The instances store data in an Amazon Aurora database. A solutions architect needs to make the application more resilient to sporadic increases in request rates.\
> Which architecture should the solutions architect implement? (Select TWO.)
> - Add Aurora replicas
> - Add AWS WAF in front of the ALB
> - Add an AWS Globabl accelerator endpoint
> - Add an Amazon Cloudfront distribution in front of the ALB
> - Add an AWS Transit Gateway to the availability Zones


Maybe **WAF** to protect against DDOS? maybe **CloudFront** for caching?

### Q03
> A company has uploaded some highly critical data to an Amazon S3 bucket. Management are concerned about data availability and require that steps are taken to protect the data from accidental deletion. The data should still be accessible, and a user should be able to delete the data intentionally.\
> Which combination of steps should a solutions architect take to accomplish this? (Select TWO.)
> - Enable MFA Delete on the S3 bucket
> - Create a lifecycle policy for the objects in the S3 bucket
> - Enable default encryption on the S3 bucket
> - Create a bucket policy on the S3 bucket
> - Enable versioning on the S3 bucket


**versioning** will work, maybe **bucket policy**

### Q04

> A Solutions Architect has deployed an application on several Amazon EC2 instances across three private subnets. The application must be made accessible to internet-based clients with the least amount of administrative effort.\
> How can the Solutions Architect make the application available on the internet?


maybe **NAT gateway**

### Q05
> Storage capacity has become an issue for a company that runs application servers on-premises. The servers are connected to a combination of block storage and NFS storage solutions. The company requires a solution that supports local caching without re-architecting its existing applications.\
> Which combination of changes can the company make to meet these requirements? (Select TWO.)

probably **storage gateway**

### Q06
> A recent security audit uncovered some poor deployment and configuration practices within your VPC. You need to ensure that applications are deployed in secure configurations.\
> How can this be achieved in the most operationally efficient manner?
> - Remove the ability for staff to deploy applications
> - use cloud formation with securely configured templates
> - manually check all application configurations before deployment
> - use aws inspector to apply secure configuration

probably **cloud formation**
### Q07
> A company has deployed a new website on Amazon EC2 instances behind an Application Load Balancer (ALB). Amazon Route 53 is used for the DNS service. The company has asked a Solutions Architect to create a backup website with support contact details that users will be directed to automatically if the primary website is down.\
> How should the Solutions Architect deploy this solution cost-effectively?
> - Create the backup website on EC2 and ALB in another Region and create an AWS accelerator point
> - Deploy the backup website on EC2 and ALB in another Region and use Route53 health checks for failover routing
> - Configure a static website using amazon S# and create a route53 weighted routing policy
> - Configure a static website using amazon S# and create a route53 failover routing policy

**Another region with route53 health checks**


### Q08

> A website runs on Amazon EC2 instances in an Auto Scaling group behind an Application Load Balancer (ALB) which serves as an origin for an Amazon CloudFront distribution. An AWS WAF is being used to protect against SQL injection attacks. A review of security logs revealed an external malicious IP that needs to be blocked from accessing the website.\
> What should a solutions architect do to protect the application?


modify the **WAF** to protect against the Ip, cloudfront means that the ip is changed.


### Q09
> A persistent database must be migrated from an on-premises server to an Amazon EC2 instances. The database requires 64,000 IOPS and, if possible, should be stored on a single Amazon EBS volume.\
> Which solution should a Solutions Architect recommend?

maybe **i3?**

### Q10
> A company offers an online product brochure that is delivered from a static website running on Amazon S3. The company’s customers are mainly in the United States, Canada, and Europe. The company is looking to cost-effectively reduce the latency for users in these regions.\
> What is the most cost-effective solution to these requirements?


**Cloudfront with some origin points** (not all)


### Q11
> A company provides a REST-based interface to an application that allows a partner company to send data in near-real time. The application then processes the data that is received and stores it for later analysis. The application runs on Amazon EC2 instances.\
> The partner company has received many 503 Service Unavailable Errors when sending data to the application and the compute capacity reaches its limits and is unable to process requests when spikes in data volume occur.\
> Which design should a Solutions Architect implement to improve scalability?

**Kinesis + Lambda**

### Q12

> A company runs an application on six web application servers in an Amazon EC2 Auto Scaling group in a single Availability Zone. The application is fronted by an Application Load Balancer (ALB). A Solutions Architect needs to modify the infrastructure to be highly available without making any modifications to the application.\
> Which architecture should the Solutions Architect choose to enable high availability?

**spread across Availability zones**



### Q13
> A new application is to be published in multiple regions around the world. The Architect needs to ensure only 2 IP addresses need to be whitelisted. The solution should intelligently route traffic for lowest latency and provide fast regional failover.\
> How can this be achieved?


**route53**?

### Q14
> A solutions architect is designing the infrastructure to run an application on Amazon EC2 instances. The application requires high availability and must dynamically scale based on demand to be cost efficient.\
> What should the solutions architect do to meet these requirements?

**application load balancer, multiple AZ**

### Q15
> A solutions architect is creating a document submission application for a school. The application will use an Amazon S3 bucket for storage. The solution must prevent accidental deletion of the documents and ensure that all versions of the documents are available. Users must be able to upload and modify the documents.\
> Which combination of actions should be taken to meet these requirements? (Select TWO.)
> - Enable MFA Delete on the bucket
> - Enable Versioning on the bucket
> - Set read only permissions on the bucket
> - Attach an IAM policy to the bucket
> - Encrypt the bucket using AWS SSE-S3

**Versioning**, maybe **MFA delete** 

### Q16

> A company runs an application that uses an Amazon RDS PostgreSQL database. The database is currently not encrypted. A Solutions Architect has been instructed that due to new compliance requirements all existing and new data in the database must be encrypted. The database experiences high volumes of changes and no data can be lost.\
> How can the Solutions Architect enable encryption for the database without incurring any data loss?

**multi AZ mode with secondary and primary?**

### Q17
> An eCommerce application consists of three tiers. The web tier includes EC2 instances behind an Application Load balancer, the middle tier uses EC2 instances and an Amazon SQS queue to process orders, and the database tier consists of an Auto Scaling DynamoDB table. During busy periods customers have complained about delays in the processing of orders. A Solutions Architect has been tasked with reducing processing times.\
> Which action will be MOST effective in accomplishing this requirement?


probably **DynamoDB accelerator**, maybe autoscaling group?(won't it take too long?)



### Q18
> A web application allows users to upload photos and add graphical elements to them. The application offers two tiers of service: free and paid. Photos uploaded by paid users should be processed before those submitted using the free tier. The photos are uploaded to an Amazon S3 bucket which uses an event notification to send the job information to Amazon SQS.\
> How should a Solutions Architect configure the Amazon SQS deployment to meet these requirements?

**short polling for the free, long polling to the paid**

### Q19
>A company requires that all AWS IAM user accounts have specific complexity requirements and minimum password length.\
How should a Solutions Architect accomplish this?

**password policy** for the entire account (if possible)

### Q20
> A company is deploying a fleet of Amazon EC2 instances running Linux across multiple Availability Zones within an AWS Region. The application requires a data storage solution that can be accessed by all of the EC2 instances simultaneously. The solution must be highly scalable and easy to implement. The storage must be mounted using the NFS protocol.\
> Which solution meets these requirements?

i think **NFS should means EFS**?

### Q21

> A company delivers content to subscribers distributed globally from an application running on AWS. The application uses a fleet of Amazon EC2 instance in a private subnet behind an Application Load Balancer (ALB). Due to an update in copyright restrictions, it is necessary to block access for specific countries.\
> What is the EASIEST method to meet this requirement?

(maybe acl?, probably **CloudFront**)


### Q22
> An application running on Amazon EC2 needs to asynchronously invoke an AWS Lambda function to perform data processing. The services should be decoupled.\
> Which service can be used to decouple the compute services?
> - AWS Config
> - AWS Step Functions
> - AWS MQ
> - AWS SNS

are **step functions a thing?**

### Q23
> A new application will run across multiple Amazon ECS tasks. Front-end application logic will process data and then pass that data to a back-end ECS task to perform further processing and write the data to a datastore. The Architect would like to reduce-interdependencies so failures do no impact other components.\
> Which solution should the Architect use?


**front pushes to SQS, backend polls**
### Q24

> A company wishes to restrict access to their Amazon DynamoDB table to specific, private source IP addresses from their VPC.\
> What should be done to secure access to the table?
> - Create an AWS VPN Connection the amazon DynamoDB endpoint
> - Create gateway VPC endpoint and add an entry to the route table
> - Create an interface VPC endpoint in the VPC with an elastic network interface
> - Create the amazon dynamoDB table in the VPC

maybe **vpc endpoint?**


### Q25
> A company runs a dynamic website that is hosted on an on-premises server in the United States. The company is expanding to Europe and is investigating how they can optimize the performance of the website for European users. The website’s backed must remain in the United States. The company requires a solution that can be implemented within a few days.\
> What should a Solutions Architect recommend?

maybe **Lambda Edge**

### Q26
> A company plans to make an Amazon EC2 Linux instance unavailable outside of business hours to save costs. The instance is backed by an Amazon EBS volume. There is a requirement that the contents of the instance’s memory must be preserved when it is made unavailable.\
> How can a solutions architect meet these requirements?

**Hibernation**


### Q27

> The database tier of a web application is running on a Windows server on-premises. The database is a Microsoft SQL Server database. The application owner would like to migrate the database to an Amazon RDS instance.\
> How can the migration be executed with minimal administrative effort and downtime?

use **Database migration service**

### Q28
> An insurance company has a web application that serves users in the United Kingdom and Australia. The application includes a database tier using a MySQL database hosted in eu-west-2. The web tier runs from eu-west-2 and ap-southeast-2. Amazon Route 53 geoproximity routing is used to direct users to the closest web tier. It has been noted that Australian users receive slow response times to queries.\
> Which changes should be made to the database tier to improve performance?


maybe **move to aurora**?

### Q29
> A web application runs in public and private subnets. The application architecture consists of a web tier and database tier running on Amazon EC2 instances. Both tiers run in a single Availability Zone (AZ).\
> Which combination of steps should a solutions architect take to provide high availability for this architecture? (Select TWO.)

**multi az**
### Q30
> A financial services company has a web application with an application tier running in the U.S and Europe. The database tier consists of a MySQL database running on Amazon EC2 in us-west-1. Users are directed to the closest application tier using Route 53 latency-based routing. The users in Europe have reported poor performance when running queries.\
> Which changes should a Solutions Architect make to the database tier to improve performance?


### Q31
> A manufacturing company captures data from machines running at customer sites. Currently, thousands of machines send data every 5 minutes, and this is expected to grow to hundreds of thousands of machines in the near future. The data is logged with the intent to be analyzed in the future as needed.\
> What is the SIMPLEST method to store this streaming data at scale?

**Kinesis firehose**


### Q32
> An Amazon RDS Read Replica is being deployed in a separate region. The master database is not encrypted but all data in the new region must be encrypted.\
> How can this be achieved?

### Q33
> A video production company is planning to move some of its workloads to the AWS Cloud. The company will require around 5 TB of storage for video processing with the maximum possible I/O performance. They also require over 400 TB of extremely durable storage for storing video files and 800 TB of storage for long-term archival.\
> Which combinations of services should a Solutions Architect use to meet these requirements?


### Q34
> A developer created an application that uses Amazon EC2 and an Amazon RDS MySQL database instance. The developer stored the database user name and password in a configuration file on the root EBS volume of the EC2 application instance. A Solutions Architect has been asked to design a more secure solution.\
> What should the Solutions Architect do to achieve this requirement?

### Q35
> A Microsoft Windows file server farm uses Distributed File System Replication (DFSR) to synchronize data in an on-premises environment. The infrastructure is being migrated to the AWS Cloud.\
> Which service should the solutions architect use to replace the file server farm?
> - AWS storage gateway
> - AWS EFS
> - AWS FSx
> - AWS EBS


### Q36
> A company runs an application on an Amazon EC2 instance the requires 250 GB of storage space. The application is not used often and has small spikes in usage on weekday mornings and afternoons. The disk I/O can vary with peaks hitting a maximum of 3,000 IOPS. A Solutions Architect must recommend the most cost-effective storage solution that delivers the performance required?\
> Which configuration should the Solutions Architect recommend?
> - Amazon EBS Cold HDD(sc1)
> - Amazon EBS provisioned IOS SSD(i01)
> - Amazon EBS Throughput Optimized HDD(st1)
> - Amazon EBS General purpose SSD (gp2)


### Q37
> An e-commerce application is hosted in AWS. The last time a new product was launched, the application experienced a performance issue due to an enormous spike in traffic. Management decided that capacity must be doubled this week after the product is launched.\
> What is the MOST efficient way for management to ensure that capacity requirements are met?
> - Add a step scaling policy
> - Add a simple scaling policy
> - Add Amazon EC2 spot instance
> - Add a scheduled scaling action



### Q38
> A legacy tightly-coupled High Performance Computing (HPC) application will be migrated to AWS.\
> Which network adapter type should be used?
> - Elastic network adaptor (ENA)
> - Elastic Fabric adaptor (EFA)
> - Elastic Ip address
> - Elastic Network interface (ENI)

### Q39
> A company's web application is using multiple Amazon EC2 Linux instances and storing data on Amazon EBS volumes. The company is looking for a solution to increase the resiliency of the application in case of a failure.\
> What should a solutions architect do to meet these requirements?
> 

### Q40
> A company's application is running on Amazon EC2 instances in a single Region. In the event of a disaster, a solutions architect needs to ensure that the resources can also be deployed to a second Region.\
> Which combination of actions should the solutions architect take to accomplish this? (Select TWO.)
> 

### Q41
> A solutions architect is designing a new service that will use an Amazon API Gateway API on the frontend. The service will need to persist data in a backend database using key-value requests. Initially, the data requirements will be around 1 GB and future growth is unknown. Requests can range from 0 to over 800 requests per second.\
> Which combination of AWS services would meet these requirements? (Select TWO.)
> - AWS Lambda
> - AWS EC2 autos scaling
> - AWS DynamoDB
> - Amazon RDS
> - AWS fargate

**Lambda** and **Dynamo**

### Q42
> A company has two accounts for perform testing and each account has a single VPC: VPC-TEST1 and VPC-TEST2. The operations team require a method of securely copying files between Amazon EC2 instances in these VPCs. The connectivity should not have any single points of failure or bandwidth constraints.\
> Which solution should a Solutions Architect recommend?


### Q43
> A team are planning to run analytics jobs on log files each day and require a storage solution. The size and number of logs is unknown and data will persist for 24 hours only.\
> What is the MOST cost-effective solution?
> - Amazon S3 intelligent tiering
> - Amazon S3 glacier deep archive
> - Amazon S3 Standard
> - Amazon S3 one zone infrequent access

**S3 standard**

### Q44
> A company is working with a strategic partner that has an application that must be able to send messages to one of the company’s Amazon SQS queues. The partner company has its own AWS account.\
> How can a Solutions Architect provide least privilege access to the partner?




### Q45
> A company runs an application in a factory that has a small rack of physical compute resources. The application stores data on a network attached storage (NAS) device using the NFS protocol. The company requires a daily offsite backup of the application data.\
> Which solution can a Solutions Architect recommend to meet this requirement?



### Q46
> A company hosts an application on Amazon EC2 instances behind Application Load Balancers in several AWS Regions. Distribution rights for the content require that users in different geographies must be served content from specific regions.\
> Which configuration meets these requirements?
> - Route53 with geoproximity
> - Route53 with geolocation
> - ALB with multi-region routing
> - CloudFronot with multiple origins and AWS WAF



### Q47

> Amazon EC2 instances in a development environment run between 9am and 5pm Monday-Friday. Production instances run 24/7.\
> Which pricing models should be used? (choose 2)



### Q48
> An organization want to share regular updates about their charitable work using static webpages. The pages are expected to generate a large amount of views from around the world. The files are stored in an Amazon S3 bucket. A solutions architect has been asked to design an efficient and effective solution.\
> Which action should the solutions architect take to accomplish this?

### Q49
> An application running on an Amazon ECS container instance using the EC2 launch type needs permissions to write data to Amazon DynamoDB.\
> How can you assign these permissions only to the specific ECS task that is running the application?


### Q50
> An Amazon VPC contains several Amazon EC2 instances. The instances need to make API calls to Amazon DynamoDB. A solutions architect needs to ensure that the API calls do not traverse the internet.
> How can this be accomplished? (Select TWO.)



### Q51
> An application is being created that will use Amazon EC2 instances to generate and store data. Another set of EC2 instances will then analyze and modify the data. Storage requirements will be significant and will continue to grow over time. The application architects require a storage solution.\
> Which actions would meet these needs?

### Q52
> A solutions architect needs to backup some application log files from an online e-commerce store to Amazon S3. It is unknown how often the logs will be accessed or which logs will be accessed the most. The solutions architect must keep costs as low as possible by using the appropriate S3 storage class.\
> Which S3 storage class should be implemented to meet these requirements?
> - Amazon S3 infrequent access
> - Amazon S3 intelligent tiering
> - Amazon S3 one zone infrequent access
> - Amazon S3 glacier 

### Q53
> A company is investigating methods to reduce the expenses associated with on-premises backup infrastructure. The Solutions Architect wants to reduce costs by eliminating the use of physical backup tapes. It is a requirement that existing backup applications and workflows should continue to function.\
> What should the Solutions Architect recommend?



### Q54
> A company uses an Amazon RDS MySQL database instance to store customer order data. The security team have requested that SSL/TLS encryption in transit must be used for encrypting connections to the database from application servers. The data in the database is currently encrypted at rest using an AWS KMS key.
> How can a Solutions Architect enable encryption in transit?




### Q55
> A company is migrating from an on-premises infrastructure to the AWS Cloud. One of the company's applications stores files on a Windows file server farm that uses Distributed File System Replication (DFSR) to keep data in sync. A solutions architect needs to replace the file server farm.\
> Which service should the solutions architect use?
> - S3
> - FSx
> - EFS
> - Storage gateway


### Q56
> Your company shares some HR videos stored in an Amazon S3 bucket via CloudFront. You need to restrict access to the private content so users coming from specific IP addresses can access the videos and ensure direct access via the Amazon S3 bucket is not possible.\
> How can this be achieved?



### Q57
> An eCommerce company runs an application on Amazon EC2 instances in public and private subnets. The web application runs in a public subnet and the database runs in a private subnet. Both the public and private subnets are in a single Availability Zone.\
> Which combination of steps should a solutions architect take to provide high availability for this architecture? (Select TWO.)

### Q58
> An organization has a large amount of data on Windows (SMB) file shares in their on-premises data center. The organization would like to move data into Amazon S3. They would like to automate the migration of data over their AWS Direct Connect link.\
> Which AWS service can assist them?
> - DataSync
> - Database migration Service
> - Database Cloud Formation
> - AWS snowball


### Q59
> A company runs an application in an on-premises data center that collects environmental data from production machinery. The data consists of JSON files stored on network attached storage (NAS) and around 5 TB of data is collected each day. The company must upload this data to Amazon S3 where it can be processed by an analytics application. The data must be transferred securely.\
> Which solution offers the MOST reliable and time-efficient data transfer?



### Q60
> A Solutions Architect has been tasked with re-deploying an application running on AWS to enable high availability. The application processes messages that are received in an ActiveMQ queue running on a single Amazon EC2 instance. Messages are then processed by a consumer application running on Amazon EC2. After processing the messages the consumer application writes results to a MySQL database running on Amazon EC2.\
> Which architecture offers the highest availability and low operational complexity?

### Q61
> A solutions architect is creating a system that will run analytics on financial data for several hours a night, 5 days a week. The analysis is expected to run for the same duration and cannot be interrupted once it is started. The system will be required for a minimum of 1 year.\
> What should the solutions architect configure to ensure the EC2 instances are available when they are needed?

### Q62
> A company runs a large batch processing job at the end of every quarter. The processing job runs for 5 days and uses 15 Amazon EC2 instances. The processing must run uninterrupted for 5 hours per day. The company is investigating ways to reduce the cost of the batch processing job.\
> Which pricing model should the company choose?
> 


### Q63
> An AWS Organization has an OU with multiple member accounts in it. The company needs to restrict the ability to launch only specific Amazon EC2 instance types.\
> How can this policy be applied across the accounts with the least effort?



### Q64
> A company hosts a multiplayer game on AWS. The application uses Amazon EC2 instances in a single Availability Zone and users connect over Layer 4. Solutions Architect has been tasked with making the architecture highly available and also more cost-effective.\
> How can the solutions architect best meet these requirements? (Select TWO.)
> 

### Q65
> A company uses Docker containers for many application workloads in an on-premise data center. The company is planning to deploy containers to AWS and the chief architect has mandated that the same configuration and administrative tools must be used across all containerized environments. The company also wishes to remain cloud agnostic to safeguard mitigate the impact of future changes in cloud strategy.\
> How can a Solutions Architect design a managed solution that will align with open-source software?






</details>