<!--
// cSpell:ignore Postgre
 -->

[main](README.md)

## Practice and Sample Exams

### EXAM 1

<details>
<summary>
exam 1 from the course
</summary>

#### Quesion 1

Question:

> You are consulting to a mid-sized company with a predominantly Mac & Linux desktop environment. In passing they comment that they have over 30TB of unstructured Word and spreadsheet documents of which 85% of these documents don't get accessed again after about 35 days. They wish that they could find a quick and easy solution to have tiered storage to store these documents in a more cost-effective manner without impacting staff access. What options can you offer them?

Answers:

> - Migrate Document to EFS storage and make use of life-cycle using infrequent access storage
> - migrate document to File Gateway presented as NFS and make use of life-cycle using infrequent access storage.

Explanation:

> Trying to use S3 without File Gateway in front would be a major impact to the user environment. Using File Gateway is the recommended way to use S3 with shared document pools. Life-cycle management and Infrequent Access storage is available for both S3 and EFS. A restriction however is that 'Using Amazon EFS with Microsoft Windows is not supported'. File Gateway does not support iSCSI in the client side.

To read up on:

- File Gateway
- EFS, NFS, iSCSI

#### Question 2

Question

> By definition, a public subnet within a VPC is one that **\_\_\_\_**.

Answer

> Has a least one route in its routing table that uses an internet gatway (IGW)

Explanation

To read up on:

- Internet Gateway
- subnets
- Network Address Translation

#### Question 3

Question

> Your company has asked you to investigate the use of KMS for storing and managing keys in AWS. From the options listed below, what key management features are available in KMS?

Answer

> Import you own keys, disable and re-enable keys and define key management roles in iAM.

Explanation

> There are many features which are native to the KMS service. However, of the above, only import your own keys, disable and re-enable keys and define key management roles in IAM are valid. Importing keys into a custom key store and migrating keys from the default key store to a custom key store are not possible. Lastly operating as a private, native HSM is a function of CloudHSM and is not possible directly within KMS.

To read up on:

- KMS
- Custom Keys

#### Question 5

Question

> You work for a large software company in Seattle. They have their production environment provisioned on AWS inside a custom VPC. The VPC contains both a public and private subnet. The company tests their applications on custom EC2 instances inside a private subnet. There are approximately 500 instances, and they communicate to the outside world via a proxy server. At 3am every night, the EC2 instances pull down OS updates, which are usually 150MB or so. They then apply these updates and reboot: if the software has not downloaded within half an hour, then the update will attempt to download the following day. You notice that a number of EC2 instances are continually failing to download the updates in the allotted time. Which of the following answers might explain this failure? [Select 2]

Answer

> - The proxy server is on an inadequately sized EC2 instance and does not have sufficient network throughput to handle all updates simulationsly. you should increase the size or type of th EC2 instance for the proxy server/
> - The proxy server is in a private subnet and uses a NAT instance to connect to the internet. However, this instance is too small to handle the required network traffic. You should re-provision the NAT solution so that it's able to handle the throughput.

Explanation

> Network throughput is the obvious bottleneck. You are not told in this question whether the proxy server is in a public or private subnet. If it is in a public subnet, the proxy server instance size itself may not be large enough to cope with the current network throughput. If the proxy server is in a private subnet, then it must be using a NAT instance or NAT gateway to communicate out to the internet. If it is a NAT instance, this may also be inadequately provisioned in terms of size. You should therefore increase the size of the proxy server and/or the NAT solution.

To read up on:

- public and private subnets
- NAT gateways
- proxy server

#### Question 6

Question

> With SAML-enabled single sign-on, **\_\_\_\_**. [Select 2]

Answer

> - After the client browser post the SAML assertion, AWS sends the sign-in URL as a redirect, and the client browser is redirected to the Console.
>   The portal first verifies the user's identity in your organization, then generates a SAML authentication respones.

Explanation

> To see the process by which federated users are granted access to the AWS console, please follow the link, below.

To read up on:

- SAML
- Federation

#### Question 10

Question

> If you don't use one of the AWS SDKs, you can perform DynamoDB operations over HTTP using the POST request method. The POST method requires you to specify the operation in the header of the request and provide the data for the operation in JSON format in the body of the request. Which of the following are valid DynamoDB Headers attributes? [Select 4]

Answer

> - content-typ
> - x-amz-date
> - x-amz-target
> - host

Explanation

> When interacting with DynamoDB directly, there is a short list of header attributes that are required

To read up on:

- Dynamo DB required attributes

#### Question 14

Question

> A client is concerned that someone other than approved administrators is trying to gain access to the Linux web app instances in their VPC. She asks what sort of network access logging can be added. Which of the following might you recommend? [Select 3]

Answer

> - Set up a Flow Log for the group of instances and forward them to S3.
> - Set up a Flow Log for the group of instances and forward them to CloudWatch.
> - Make use of OS level logging tools such as iptables and log events to CloudWatch or S3.

Explanation

> Security and Auditing in AWS needs to be considered during the Design phase.

To read up on:

- Flow Log

#### Question 15

Question

> You are a solutions architect with a manufacturing company running several legacy applications. One of these applications needs to communicate with services that are currently hosted on-premise. The people who wrote this application have left the company, and there is no documentation describing how the application works. You need to ensure that this application can be hosted in a bespoke VPC, but remains able to communicate to the back-end services hosted on-premise. Which of the following answers will allow the application to communicate back to the on premise equipment without the need to reprogram the application? [Select 3]

Answer

> - You should ensure the VPC has an internet gateway attached to it. That way, you can establish site-to-site VPC with the on-premise environment.
> - You should configure the VPC subnet in which the application sits so that it does not have an IP address range that conflict with that of the on-premises VLAN in which the backend services sit.
> - You should configure AWS Direct Connect link between the VPC and the site with the on-premises solution.

Explanation

> You need to ensure that your application in your custom VPC can communicate back to the on-premise data center. You can do this by either using a site to site VPN or Direct Connect. It will be using an internal IP address range, so you must make sure that your internal IP addresses do not overlap.

To read up on:

- VPN
- Direct Connect
- Vlan
- Ip ranges

#### Question 16

Question

> AWS provides a number of security-related managed services. From the options below, select which AWS service is related to protecting your infrastructure from which security issue. [Select 4]

Answer

> - AWS WAF block IP addresses based on rules
> - AWS shield protects from Distributed Denial of service Attacks
> - Amazon Macie uses Machine Learning to protect sensitive data.
> - AWS WAF protects from Cross-site scripting attacks.

Explanation

> AWS provides various services to cope with many security related issues and because of this, there are a number of options which are correct. AWS Shield has two options listed above, but only one is correct. AWS Shield operates on layer 3 and 4 of the ISO network model and its primary purpose is to protect against DDoS attacks. It does not have any affect against SQL Injection attacks which are dealt with by AWS WAF. WAF also protects against Cross Site Scripting and can block traffic from IP addresses based on rules and therefore these options are also correct. Finally, Amazon Macie tackles a different problem related to Data Loss Prevention and protects sensitive data and so this answer is also correct.

To read up on:

- AWS WAF
- AWS Shield
- Amazon Macie
- network layers.

#### Question 17

Question

> Which of the following features only relate to Spread Placement Groups?

Answer

> - The placement group can only have 7 running instances per Avalability Zone.

Explanation

> Spread placement groups have a specific limitation that you can only have a maximum of 7 running instances per Availability Zone and therefore this is the only correct option. Deploying instances in a single Availability Zone is unique to Cluster Placement Groups only and therefore is not correct. The last two remaining options are common to all placement group types and so are not specific to Spread Placement Groups.

To read up on:

- Cluster Placement Groups

#### Question 19

Question

> You have an enterprise solution that operates Active-Active with facilities in Regions US-West and India. Due to growth in the Asian market you have been directed by the CTO to ensure that only traffic in Asia (between Turkey and Japan) is directed to the India Region. Which of these will deliver that result? [Select 2]

Answer

> - Route 53 - Geoproximity routing policy
> - Route 53 - Geolocation routing policy

Explanation

> The instruction from the CTO is clear that that the division is based on geography. Latency based routing will approximate geographic balance only when all routes and traffic evenly supported which is rarely the case due to infrastructure and day night variations. You cannot combine blacklisting and whitelisting in CloudFront. Weighted routing is randomized and will not respect Geo boundaries. Geolocation is based on national boundaries and will meet the needs well. Geoproximity is based on Latitude & Longitude and will also provide a good approximation with potentially less configuration.

To read up on:

- Route 53 Routing Policies.

#### Question 20

Question

> You have created a Direct Connect Link from your on premise data center to your Amazon VPC. The link is now active and routes are being advertised from the on-premise data center. You can connect to EC2 instances from your data center; however, you cannot connect to your on premise servers from your EC2 instances. Which of the following solutions would remedy this issue? [Select 2]

Answer

> - Edit your VPC subnet route table, adding a route back to the on-premises data center.
> - Enable route propagation your Virtual Private Gateway (VPG)

Explanation

> There is no route connecting your VPC back to the on premise data center. You need to add this route to the route table and then enable propagation on the Virtual Private Gateway.

To read up on:

- subnet route table
- Virtual private Gateway
- Direct Connect link

#### Question 21

Question

> You're building out a single-region application in us-west-2. However, disaster recovery is a strong consideration, and you need to build the application so that if us-west-2 becomes unavailable, you can fail-over to us-west-1. Your application relies exclusively on pre-built AMI's. In order to share those AMI's with the region you're using as a backup, which process would you follow?

Answer

> Copy the AMI from us-west-2, manually apply launch permissions, user-defined tags, and Amazon S3 bucket permission of the default AMI to the new instance, and launch the instance.

Explanation

> AWS does not copy launch permissions, user-defined tags, or Amazon S3 bucket permissions from the source AMI to the new AMI.

To read up on:

- AMI, launch permissions, etc...

#### Question 24

Question

> Which of the following DynamoDB features are chargeable, when using a single region? [Select 2]

Answer

> - Read and Write cpacity
> - Storage Data

Explanation

> There will always be a charge for provisioning read and write capacity and the storage of data within DynamoDB, therefore these two answers are correct. There is no charge for the transfer of data into DynamoDB, providing you stay within a single region (if you cross regions, you will be charged at both ends of the transfer.) There is no charge for the actual number of tables you can create in DynamoDB, providing the RCU and WCU are set to 0, however in practice you cannot set this to anything less than 1 so there always be a nominal fee associated with each table.

To read up on:

- Dynamo Db charge

#### Question 25

Question

> ​Your company has a policy of encrypting all data at rest. You host your production environment on EC2 in a bespoke VPC. Attached to your EC2 instances are multiple EBS volumes, and you must ensure this data is encrypted. Which of the following options will allow you to do this? [Select 3]

Answer

> - Use third party volume encryption tools
> - Encrypt the data using native encryption tool avaable in the operating system (such as windows BitLocker)
> - Encrypt your data inside your application, before storing it on EBS

Explanation

> EBS volumes can be encrypted, but they are not encrypted by default. SSL certificates will only be useful to encrypt data in transit, not data at rest.

To read up on:

- data encrption

#### Question 27

Question

> Your company has hired a young and enthusiastic accountant. After reviewing the AWS documentation and usage graphs, he announces that you are wasting vast amounts of money running your Windows servers for a full hour instead of spinning them up only when they are needed and down again as soon as they are idle for 1 minute. He cites the AWS claim that you only pay for what you use, and that as a senior engineer, you should be more conscious of wasting company money. How do you respond?

Answer

> You thank him for his concern, and advice him that he has misinterpreted the pricing documentation. Windows instances are billed by the full hour, and partial hours are billed as such. Additionally, storage charges are incurred even of the DB instance sits idle. Taking into Account productivy losses, stopping and restarting DB instances may actually result in additional costs. as such, your solution is find as it now stands.

Explanation

> The study of AWS Billing is a discipline unto itself. For more information, please see the AWS Cost Control Course on the A Cloud Guru platform.

To read up on:

- Bulling

#### Question 28

Question

> In addition to choosing the correct EBS volume type for your specific task, what else can be done to increase the performance of your volume? [Select 3]

Answer

> - Schedule snapshots of HDD based volumes for periods of low use
> - Stripe volumes together in RAID 0 configuration
> - Ensure that your EC2 instances are types that can be optimized for use with EBS.

Explanation

> There are a number of ways you can optimize performance above that of choosing the correct EBS type. One of the easiest options is to drive more I/O throughput than you can provision for a single EBS volume, by striping using RAID 0. You can join multiple gp2, io1, st1, or sc1 volumes together in a RAID 0 configuration to use the available bandwidth for these instances.\
> You can also choose an EC2 instance type that supports EBS optimization. This ensures that network traffic cannot contend with traffic between your instance and your EBS volumes.\
> The final option is to manage your snapshot times, and this only applies to HDD based EBS volumes. When you create a snapshot of a Throughput Optimized HDD (st1) or Cold HDD (sc1) volume, performance may drop as far as the volume's baseline value while the snapshot is in progress. This behavior is specific to these volume types. Therefore you should ensure that scheduled snapshots are carried at times of low usage. \
> The one option on the list which is entirely incorrect is the option that states "Never use HDD volumes, always ensure that SSDs are used" as the question first states "In addition to choosing the correct EBS volume type for your specific task". HDDs may well be suitable to certain tasks and therefore they shouldn't be discounted because they may not have the highest specification on paper.

To read up on:

- EBS
- RAID 0
- performance optimization

#### Question 29

Question

> You run a meme creation website that stores the original images in S3 and each meme's metadata in DynamoDB. You need to decide upon a low-cost storage option for the memes, themselves. If a meme object is unavailable or lost, a Lambda function will automatically recreate it but at a $10 licensing cost per creation. Which storage solution should you use to store the memes in the most cost-effective way?

Answer

> S3 - IA

Explanation

> The Question describes a situation where low cost OneZone-IA would be perfect. However it also says that there is a high license cost with each meme generation. The storage savings between IA and OneZone-IA are about $0.0025 this is small compared to the $10 for licensing. Therefore you may well be better to pay for full S3-IA.

#### Question 31

Question

> You are a consultant planning to deploy DynamoDB across three AZs. Your lead DBA is concerned about data consistency. Which of the following do you advise the lead DBA to do?

Answer

> To ask the development team to code for strongly consistent reads. As the conslutant, you will advise the CTO of the increased cost.

Explanation

> The term consistency has specific meaning in relationship to DynamoDB.

To read up on:

- Data consistency in dynamo DB
- DynamoDB pricing

#### Question 33

Question

> You work for a popular media outlet about to release a story that is expected to go viral. During load testing on the website, you discover that there is read contention on the database tier of your application. Your RDS instance consists of a MySQL database on an extra large instance. Which of the following approaches would be best to further scale this instance to meet the anticipated increase in traffic your viral story will generate? [Select 3]

Answer

> - Add an RDS read Replica for increased read performance
> - Provision a larger instance size with provisioned IOPS.
> - Use elastic Cache to cache the frequently read, static data.

Explanation

> You should consider; using ElastiCache, using RDS Read Replicas Scaling up may also resolve the contention, however it may be more expensive than offloading the read activities to cache or Read-Replicas. RDS Multi-AZ is for resilience only.

To read up on:

- High availability with RDS
- Caching

#### Question 35

Question

> The Customer Experience manager comes to see you about some odd behaviors with the ticketing system: messages presented to the support team are not arriving in the order in which they were generated. You know that this is due to the way that the underlying SQS standard queue service is being used to manage messages. Which of the following are correct explanations? [Select 2]

Answer

> - SQS uses multiple hosts, and each host holds only a portion of all the messages. When a staff member calls for their next message, the consumer process does not see all the host or all the messages. As such, messages are not necessarily delivered in the order in which they were generated.

Explanation

> With a Standard queue, delivery is "at-least-once", and FIFO delivery is not guaranteed. If FIFO delivery is required, A FIFO queue should be used.

To read up on:

- SQS

#### Question 36

Question

> How does AWS deliver high durability for DynamoDB?

Answer

> DynamoDB data is automatically replicated across multiple AZs.

Explanation

> Basic good DB architecture.

To read up on:

- DynamoDB HA

#### Question 38

Question

> Your company likes the idea of storing files on AWS. However, low-latency service of the majority of files is important to customer service. Which Storage Gateway configuration would you use to achieve both of these ends? [Select 2]

Answer

> - File Gateways
> - Gateway Cache

Explanation

> Gateway-Stored volumes store your primary data locally, while asynchronously backing up that data to AWS. Depending on the Cache allocated you can achieve the same with File Gateway

To read up on:

- Gateway configurations

#### Question 39

Question

> You work for a large media organization who has traditionally stored all their media on large SAN arrays. After evaluating AWS, they have decided to move their storage to the cloud. Staff will store their personal data on S3, and will have to use their Active Directory credentials in order to authenticate. These items will be stored in a single S3 bucket, and each staff member will have their own folder within that bucket named after their employee ID. Which of the following steps should you take in order to help set this up? [Select 3]

Answer

> - Create an IAM role
> - Create either a federation proxy or identity provider
> - Use AWS security token service to create temporary tokens.

Explanation

> You cannot tag individual folders within an S3 bucket. If you create an individual user for each staff member, there will be no way to keep their active directory credentials synched when they change their password. You should either create a federation proxy or identity provider and then use AWS security token service to create temporary tokens. You will then need to create the appropriate IAM role for which the users will assume when writing to the S3 bucket.

To read up on:

- IAM Federation
- Identity Providers

#### Question 41

Question

> This NAT instance allows individual EC2 instances in private subnets to communicate out to the internet without being directly accessible via the internet. As the company has grown over the last year, they are finding that the additional traffic through the NAT instance is causing serious performance degradation. What might you do to solve this problem?

Answer

> increase the class size of the NAT instance from an m4.medium to m4.xLarge

To read up on:

- NAT

#### Question 42

Question

> You have provisioned a custom VPC with a subnet that has a CIDR block of 10.0.3.0/28 address range. Inside this subnet, you have 2 webservers, 2 application servers, 2 database servers, and a NAT. You have configured an Autoscaling group on the two web servers to automatically scale when the CPU utilization goes above 90%. Several days later you notice that autoscaling is no longer deploying new instances into the subnet, despite the CPU utilization of all web servers being at 100%. Which of the following answers may offer an explanation? [Select 2]

Answer

> - AWS reserves both first four and the last ip address in each subnet CIDR block.
> - Your autoscaling group has provisioned too many EC2 instances and has exhausted the number of internal IP addresses available in the subnet.

Explanation

> A /28 subnet will only have 16 addresses available. AWS reserve both the first four and last IP addresses in each subnet’s CIDR block. It is likely that your autoscaling group has provisioned too many EC2 instances and you have run out of internal private IP addresses.

To read up on:

- CIR blocks
- Ip ranges

#### Question 43

Question

> You are a systems administrator and you need to monitor the health of your production environment. You decide to do this using CloudWatch. However, you notice that you cannot see the health of every important metric in the default dashboard. When monitoring the health of your EC2 instances, for which of the following metrics do you need to design a custom CloudWatch metric?

Answer

> Memory Usage

Explanation

> Remember under the shared security model that AWS can see the instance, but not inside the instance to what it is doing. AWS can see that you have Memory, but how much of the memory is being used cannot be seen by AWS. In the case of CPU AWS can see how much of CPU you are using, but cannot see what you are using if for.

To read up on:

- Cloud Watch

#### Question 46

Question

> You are running a media-rich website with a global audience from us-east-1 for a customer in the publishing industry. The website updates every 20 minutes. The web-tier of the site sits on three EC2 instances inside an Auto Scaling Group. The Auto Scaling group is configured to scale when CPU utilization of the instances is greater than 70%. The Auto Scaling group sits behind an Elastic Load Balancer, and your static content lives in S3 and is distributed globally by CloudFront. Your RDS database is already the largest instance size available. CloudWatch metrics show that your RDS instance usually has around 2GB of memory free, and an average CPU utilization of 75%. Currently, it is taking your users in Japan and Australia approximately 3 - 5 seconds to load your website, and you have been asked to help reduce these load-times. How might you improve your page load times? [Select 3]

Answer

> - Set up CloudFront with dynamic content support to enable the caching of reusable content from the media rich website
> - Use ElasticChache to cache the most commonly accessed DB queries.
> - Set up a clone of your production environment in the Asia Pacific region and configure latency based routing on route53.

Explanation

> Additional clones of your production environment, ElastiCache, and CloudFront can all help improve your site performance. Changing your autoscaling policies will not help improve performance times as it is much more likely that the performance issue is with the database back end rather than the front end. The Provisioned IOPS would also not help, as the bottleneck is with the memory, not the storage.

To read up on:

- CloudFront

#### Question 47

Question

> At the monthly product meeting, one of the Product Owners proposes an idea to address an immediate shortcoming of the product system: storing a copy of the customer price schedule in the customer record in the database. You know that you can store large text or binary objects in DynamoDB. You give a tentative OK to do a Minimal Viable Product test, but stipulate that it must comply with the size limitation on the Attribute Name & Value. Which is the correct limitation?

Answer

> The combined Value and Name must not exceed 400KB.

Explanation

> DynamoDB allows for the storage of large text and binary objects, but there is a limit of 400 KB.

To read up on:

- Dynamo DB limitation

#### Question 46

Question

> Which of the below are factors that have helped make public cloud so powerful? [Select 2]

Answer

> - Not having to deal with the collateral damage of failed experiments
> - The ability to try out new ideas and experiment without upfront commitment.

Explanation

> Public cloud allows organizations to try out new ideas, new approaches and experiment with little upfront commitment. If it doesn't work out, organizations have the ability to terminate the resources and stop paying for them.

#### Question 53

Question

> Which of the following RDS database engines have a limit to the number of databases that can run per instance? [Select 2]

Answer

> - SQL Server
> - Oracle

Explanation

> Both the Oracle and SQL Server database engines have limits to how many databases that can run per instance. Primarily, this is due to the underlying technology being proprietary and requiring specific licensing to operate. The database engines based on Open Source technology such as Aurora, MySQL, MariaDB or PostgreSQL have no such limits.

To read up on:

- RDS Database

#### Question 55

Question

> What is the underlying Hypervisor for EC2 ? [Select 2]

Answer

> - Xen
> - Nitro

Explanation

> Until very recently AWS exclusively used Xen Hypervisors, Recently they started making use of Nitro Hypervisors.

#### Question 56

Question

> You successfully configure VPC Peering between VPC-A and VPC-B. You then establish an IGW and a Direct-Connect connection in VPC-B. Can instances in VPC-A connect to your corporate office via the Direct-Connect service, and connect to the Internet via the IGW?

Answer

> VPC peering does not support edge to edge routing

Explanation

> VPC peering only routes traffic between source and destination VPCs. VPC peering does not support edge to edge routing

To read up on:

- VPC peering
- Direct Connect

#### Question 57

Question

> Your server logs are full of what appear to be application-layer attacks, so you deploy AWS Web Application Firewall. Which of the following conditions may you set when configuring AWS WAF? [Select 3]

Answer

> - Size Constraint Conditions
> - IP Match Conditions
> - String Match Conditions

To read up on:

- AWS WAF

#### Question 58

Question

> Which of the following data formats does Amazon Athena support? [Select 3]

Answer

> - Apache Parquet
> - Json
> - Apache ORC

Explanation

> Amazon Athena is an interactive query service that makes it easy to analyse data in Amazon S3, using standard SQL commands. It will work with a number of data formats including "JSON", "Apache Parquet", "Apache ORC" amongst others, but "XML" is not a format that is supported.

To read up on:

- Amazon Athena

#### Question 60

Question

> You have been engaged as a consultant by a company that generates utility bills and publishes them online. PDF images are generated, then stored on a high-performance RDS instance. Customarily, invoices are viewed by customers once per month. Recently, the number of customers has increased threefold, and the wait-time necessary to view invoices has increased unacceptably. The CTO is unwilling to alter the codebase more than necessary this quarter, but needs to return performance to an acceptable level before the end-of-the-month print run. Which of the following solutions would you feel comfortable proposing to the CTO and GM? [Select 2]

Answer

> - Evaluate the risks and benefits associated with an RDS instance upgrade
> - Create RDS Read-Replicas and additional Web/App instances across all the available AZs.

Explanation

> Caching content is not always effective. Sometimes, optimal solutions cannot be achieved; so you need to figure out the next best way to keep the show going.

#### Question 61

Question

> Which of the following provide the lowest cost EBS options? [Select 2]

Answer

> - throughput optimized (st1)
> - Cold (sc1)

Explanation

> Of all the EBS types, both current and of the previous generation, HDD based volumes will always be less expensive than SSD types. Therefore, of the options available in the question, the Cold (sc1) and Throughout Optimized (st1) types are HDD based and will be the lowest cost options.

To read up on:

- EBS types

#### Question 67

Question

> You work for a games development company that are re-architecting their production environment. They have decided to make all web servers stateless. Which of the following the AWS services will help them achieve this goal? [Select 3]

Answer

> - DynamoDB
> - RDS (aurora)
> - Elastic cache

Explanation

> An Elastic Load Balancer can help you deliver stateful services, but not stateless. Elastic Map Reduce is a data crunching services and is not related to servicing web traffic.

#### Question 71

Question

> You are leading a design team to implement an urgently needed collection and analysis project. You will be collecting data for an array of 50,000 anonymous data collectors which will be summarized each day and then rarely used again. The data will be pulled from collectors approximately once an hour. The Dev responsible for the DynamoDB design is concerned about how to design the Partition and Local keys to ensure efficient use of the DynamoDB tables. What advice would you provide. [Select 2]

Answer

> - Create a new table each day, and reconfigure the old table for infrequent us after the summation is complete.
> - Insert a calculated hash in front of the Dat/Time value in the partition key to force DynamoDB to use partitions in parallel.

Explanation

> There are two issues here: how to handle stale data to avoid paying for high provisioned throughput for infrequently used data, and how to design a partition key that will distribute IO from sequential data across partitions evenly to avoid performance bottlenecks.

To read up on:

- DynamoDB partition keys

#### Question 72

Question

> Your company has decided to set up a new AWS account for test and dev purposes. They already use AWS for production, but would like a new account dedicated for test and dev so as to not accidentally break the production environment. You launch an exact replica of your production environment using a CloudFormation template that your company uses in production. However, CloudFormation fails. You use the exact same CloudFormation template in production, so the failure is something to do with your new AWS account. The CloudFormation template is trying to launch 60 new EC2 instances in a single availability zone. After some research, you discover that the problem is **\_\_\_\_**.

Answer

> For all new AWS accounts, there is a soft limit of 20 EC2 instances per region. You should submit the limit increase form and retry the template after your limit has been increased.

To read up on:

- account limitations

#### Question 73

Question

> When editing permissions (policies and ACLs), to whom does the concept of the "Owner" refer?

Answer

> The "Owner" refers to the identity and email address used to create the AWS account.

Explanation

> The Owner concept comes into play especially when setting or locking down access to various objects.

#### Question 74

Question

> In AWS Route 53, which of the following are true? [Select 2]

Answer

> - Alias Records provide a Route53 specific extension to DNS functionality
> - Route53 allows you to create an alias record at the top node of a DNS namespace (zone apex)

Explanation

> Alias Records have special functions that are not present in other DNS servers. Their main function is to provide special functionality and integration into AWS services. Unlike CNAME records, they can also be used at the Zone Apex, where CNAME records cannot. Alias Records can also point to AWS Resources that are hosted in other accounts by manually entering the ARN

To read up on:

- Route 53

#### Question 75

Question

> How is the Public IP address managed in an instance session via the instance GUI/RDP or Terminal/SSH session?

Answer

> The public IP address is not managed on the instance, it is, instead, an alias applied as network address translation of the private IP address.

Explanation

> AWS networking is implemented differently from most conventional data centers.

To read up on:

- NAT

</details>

### EXAM 2

<details>
<summary>
exam 2 from the course
</summary>

</details>

### AWS Sample Question

<details>
<summary>
from the aws Website
</summary>

[aws sample questions](https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Sample-Questions.pdf)

1.  A customer relationship management (CRM) application runs on Amazon EC2 instances in multiple Availability Zones behind an Application Load Balancer. If one of these instances fails, what occurs?

    1. The load balancer will stop sending requests to the failed instance.
    1. The load balancer will terminate the failed instance.
    1. The load balancer will automatically replace the failed instance.
    1. The load balancer will return 504 Gateway Timeout errors until the instance is replaced.

2.  A company needs to perform asynchronous processing, and has Amazon SQS as part of a decoupled architecture. The company wants to ensure that the number of empty responses from polling requests are kept to a minimum. What should a solutions architect do to ensure that empty responses are reduced?

    1. Increase the maximum message retention period for the queue.
    2. Increase the maximum receives for the redrive policy for the queue.
    3. Increase the default visibility timeout for the queue.
    4. Increase the receive message wait time for the queue.

3.  A company currently stores data for on-premises applications on local drives. The chief technology officer wants to reduce hardware costs by storing the data in Amazon S3 but does not want to make modifications to the applications. To minimize latency, frequently accessed data should be available locally. What is a reliable and durable solution for a solutions architect to implement that will reduce the cost of local storage?

    1.  Deploy an SFTP client on a local server and transfer data to Amazon S3 using AWS Transfer for SFTP.
    2.  Deploy an AWS Storage Gateway volume gateway configured in cached volume mode.
    3.  Deploy an AWS DataSync agent on a local server and configure an S3 bucket as the destination.
    4.  Deploy an AWS Storage Gateway volume gateway configured in stored volume mode.

4.  A company runs a public-facing three-tier web application in a VPC across multiple Availability Zones. Amazon EC2 instances for the application tier running in private subnets need to download software patches from the internet. However, the instances cannot be directly accessible from the internet. Which actions should be taken to allow the instances to download the needed patches? (Select TWO.)

    1.  Configure a NAT gateway in a public subnet.
    2.  Define a custom route table with a route to the NAT gateway for internet traffic and associate it with the private subnets for the application tier.
    3.  Assign Elastic IP addresses to the application instances.
    4.  Define a custom route table with a route to the internet gateway for internet traffic and associate it with the private subnets for the application tier.
    5.  Configure a NAT instance in a private subnet.

5.  A solutions architect wants to design a solution to save costs for Amazon EC2 instances that do not need to run during a 2-week company shutdown. The applications running on the instances store data in instance memory (RAM) that must be present when the instances resume operation. Which approach should the solutions architect recommend to shut down and resume the instances?

    1.  Modify the application to store the data on instance store volumes. Reattach the volumes while restarting them.
    2.  Snapshot the instances before stopping them. Restore the snapshot after restarting the instances.
    3.  Run the applications on instances enabled for hibernation. Hibernate the instances before the shutdown.
    4.  Note the Availability Zone for each instance before stopping it. Restart the instances in the same Availability Zones after the shutdown.

6.  A company plans to run a monitoring application on an Amazon EC2 instance in a VPC. Connectionsare made to the instance using its private IPv4 address. A solutions architect needs to design a solution that will allow traffic to be quickly directed to a standby instance if the application fails and becomes unreachable. Which approach will meet these requirements?

    1.  Deploy an Application Load Balancer configured with a listener for the private IP address and register the primary instance with the load balancer. Upon failure, de-register the instance and register the secondary instance.
    2.  Configure a custom DHCP option set. Configure DHCP to assign the same private IP address to the secondary instance when the primary instance fails.
    3.  Attach a secondary elastic network interface (ENI) to the instance configured with the private IP address. Move the ENI to the standby instance if the primary instance becomes unreachable.
    4.  Associate an Elastic IP address with the network interface of the primary instance. Disassociate the Elastic IP from the primary instance upon failure and associate it with a secondary instance.

7.  An analytics company is planning to offer a site analytics service to its users. The service will require that the users’ webpages include a JavaScript script that makes authenticated GET requests to the company’s Amazon S3 bucket. What must a solutions architect do to ensure that the script will successfully execute?

    1.  Enable cross-origin resource sharing (CORS) on the S3 bucket.
    2.  Enable S3 versioning on the S3 bucket.
    3.  Provide the users with a signed URL for the script.
    4.  Configure a bucket policy to allow public execute privileges.

8.  A company’s security team requires that all data stored in the cloud be encrypted at rest at all times using encryption keys stored on-premises. Which encryption options meet these requirements? (Select TWO.)

    1.  Use Server-Side Encryption with Amazon S3 Managed Keys (SSE-S3).
    2.  Use Server-Side Encryption with AWS KMS Managed Keys (SSE-KMS).
    3.  Use Server-Side Encryption with Customer Provided Keys (SSE-C).
    4.  Use client-side encryption to provide at-rest encryption.
    5.  Use an AWS Lambda function triggered by Amazon S3 events to encrypt the data using the customer’s keys.

9.  A company needs to maintain access logs for a minimum of 5 years due to regulatory requirements. The data is rarely accessed once stored, but must be accessible with one day’s notice if it is needed. What is the MOST cost-effective data storage solution that meets these requirements?

    1.  Store the data in Amazon S3 Glacier Deep Archive storage and delete the objects after 5 years using a lifecycle rule.
    2.  Store the data in Amazon S3 Standard storage and transition to Amazon S3 Glacier after 30 days using a lifecycle rule.
    3.  Store the data in logs using Amazon CloudWatch Logs and set the retention period to 5 years.
    4.  Store the data in Amazon S3 Standard-Infrequent Access (S3 Standard-IA) storage and delete the objects after 5 years using a lifecycle rule.

10. A company uses Reserved Instances to run its data-processing workload. The nightly job typically takes 7 hours to run and must finish within a 10-hour time window. The company anticipates temporary increases in demand at the end of each month that will cause the job to run over the time limit with the capacity of the current resources. Once started, the processing job cannot be interrupted before completion. The company wants to implement a solution that would allow it to provide increased capacity as cost-effectively as possible. What should a solutions architect do to accomplish this?
    1. Deploy On-Demand Instances during periods of high demand.
    2. Create a second Amazon EC2 reservation for additional instances.
    3. Deploy Spot Instances during periods of high demand.
    4. Increase the instance size of the instances in the Amazon EC2 reservation to support the increased workload.

| Question | my answer                                                                                                                                                                                                        | correct answer                                                                                                                                                                                                                                                                                                                                                                                                                                                          | points? |
| -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- |
| 1        | a: The load balancer will stop sending requests to the failed instance.                                                                                                                                          | "A – An Application Load Balancer (ALB) sends requests to healthy instances only. An ALB performs periodic health checks on targets in a target group. An instance that fails health checks for a configurable number of consecutive times is considered unhealthy. The load balancer will no longer send requests to the instance until it passes another health check."                                                                                               | Yes     |
| 2        | d: Increase the receive message wait time for the queue.                                                                                                                                                         | "D – When the ReceiveMessageWaitTimeSeconds property of a queue is set to a value greater than zero, long polling is in effect. Long polling reduces the number of empty responses by allowing Amazon SQS to wait until message is available before sending a response to a ReceiveMessage request"                                                                                                                                                                     | Yes     |
| 3        | b: Deploy an AWS Storage Gateway volume gateway configured in cached volume mode                                                                                                                                 | "B – An AWS Storage Gateway volume gateway connects an on-premises software application with cloudbacked storage volumes that can be mounted as Internet Small Computer System Interface (iSCSI) devices from on-premises application servers. In cached volumes mode, all the data is stored in Amazon S3 and a copy of frequently accessed data is stored locally."                                                                                                   | Yes     |
| 4        | (d: Define a custom route table with a route to the internet gateway for internet traffic and associate it with the private subnets for the application tier), (e: Configure a NAT instance in a private subnet) | "A, B – A NAT gateway forwards traffic from the instances in the private subnet to the internet or other AWS services, and then sends the response back to the instances. After a NAT gateway is created, the route tables for private subnets must be updated to point internet traffic to the NAT gateway"                                                                                                                                                            | No      |
| 5        | b: Snapshot the instances before stopping them. Restore the snapshot after restarting the instances.                                                                                                             | "C – Hibernating an instance saves the contents of RAM to the Amazon EBS root volume. When the instance restarts, the RAM contents are reloaded."                                                                                                                                                                                                                                                                                                                       | No      |
| 6        | c: Attach a secondary elastic network interface (ENI) to the instance configured with the private IP address. Move the ENI to the standby instance if the primary instance becomes unreachable.                  | "C – A secondary ENI can be added to an instance. While primary ENIs cannot be detached from an instance, secondary ENIs can be detached and attached to a different instance."                                                                                                                                                                                                                                                                                         | Yes     |
| 7        | a: Enable cross-origin resource sharing (CORS) on the S3 bucket.                                                                                                                                                 | "A – Web browsers will block the execution of a script that originates from a server with a different domain name than the webpage. Amazon S3 can be configured with CORS to send HTTP headers that allow the script execution."                                                                                                                                                                                                                                        | Yes     |
| 8        | (b: Use Server-Side Encryption with AWS KMS Managed Keys (SSE-KMS)),(c: Use Server-Side Encryption with Customer Provided Keys (SSE-C))                                                                          | "C, D – Server-Side Encryption with Customer-Provided Keys (SSE-C) enables Amazon S3 to encrypt objects server side using an encryption key provided in the PUT request. The same key must be provided in GET requests for Amazon S3 to decrypt the object. Customers also have the option to encrypt data client side before uploading it to Amazon S3 and decrypting it after downloading it. AWS SDKs provide an S3 encryption client that streamlines the process." | No      |
| 9        | a : Store the data in Amazon S3 Glacier Deep Archive storage and delete the objects after 5 years using a lifecycle rule.                                                                                        | "A – Data can be stored directly in Amazon S3 Glacier Deep Archive. This is the cheapest S3 storage class. "                                                                                                                                                                                                                                                                                                                                                            | Yes     |
| 10       | a: Deploy On-Demand Instances during periods of high demand.                                                                                                                                                     | "A – While Spot Instances would be the least costly option, they are not suitable for jobs that cannot be interrupted or must complete within a certain time period. On-Demand Instances would be billed for the number of seconds they are running."                                                                                                                                                                                                                   | Yes     |

</details>

#### Question XXXX

Question

>

Answer

>

Explanation

>

To read up on:

##

[main](README.md)
