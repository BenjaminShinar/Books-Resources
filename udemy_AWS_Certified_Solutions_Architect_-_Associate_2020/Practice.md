<!--
// cSpell:ignore Postgre Magento
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

#### Question 4

Question

> What is the purpose of an Egress-Only Internet Gateway? [Select 2]

Answer

> - Allows VPC based IPv6 traffic to communicate to the internet.
> - Preven IPv6 based interned resources initiating a connection into a VPC.

Explanation

> The purpose of an "Egress-Only Internet Gateway" is to allow IPv6 based traffic within a VPC to access the Internet, whilst denying any Internet based resources the possibility of initiating a connection back into the VPC.

To read up on:

- Egress Only
- VPC

#### Question 7

Question

> What does the common term 'Serverless' mean according to AWS [Select 2]

Answer

> - A Native Cloud Architecture that allows customers to shift more operational responsability to AWS.
> - The Ability to run applications and services without thinking about servers or capacity provisioning.

Explanation

> 'Serverless' computing is not about eliminating servers, but shifting most of the responsibility for infrastructure and operation of the infrastructure to a vendor so that you can focus more on the business services, not how to manage the infrastructure that they run on. Billing does tend to be based on simple units, but the choice of services, intended usage pattern (RIs), and amount of capacity needed also influences the pricing.

To read up on:

- Serverless principals

#### Question 10

Question

> Following advice from your consultant, you have configured your VPC to use Dedicated hosting tenancy. A subsequent change to your application has rendered the performance gains from dedicated tenancy superfluous, and you would now like to recoup some of these greater costs. How do you revert to Default hosting tenancy?​

Answer

> Use the AWS CLI to modify the instance placement attribute of each instance and the vpc tenancy attribute of the VPC.

Explanation

> Once a VPC is set to Dedicated hosting, it can be changed back to default hosting via the CLI, SDK or API. Note that this will not change hosting settings for existing instances, only future ones. Existing instances can be changed via CLI, SDK or API but need to be in a stopped state to do so

To read up on:

- Dedicated Hostings

#### Question 13

Question

> The risk with spot instances is that you are not guaranteed use of the resource for as long as you might want. Which of the following are scenarios under which AWS might execute a forced shutdown? [Select 4]

Answer

> - AWS sends a notification of termination but you do not receive it within the 120 seconds and the instance is shutdown.
> - AWS sends a notification of termination and you receive it 120 seconds before the intended forced shutdown.
> - AWS sends a notification of termination and you receive it 120 seconds before the intended forced shutdown, but AWS do not action the shutdown.
> - AWS sends a notification of termination and you receive it 120 seconds before the intended forced shutdown, but the normal lease expired before the forces shutdown.

Explanation

> notification of spot termination

To read up on:

- Spot and on-Demand instances

#### Question 15

Question

> AWS S3 has four different URLs styles that it can be used to access content in S3. The Virtual Hosted Style URL, the Path-Style Access URL, the Static web site URL, and the Legacy Global Endpoint URL. Which of these represents a correct formatting of the Path-Style Access URL style.

Answer

> "https://s3.us-west-2.amazonaws.com/my-bucket/slowpuppy.tar"

Explanation

> Virtual style puts your bucket name 1st, s3 2nd, and the region 3rd. Path style puts s3 1st and your bucket as a sub domain. Legacy Global endpoint has no region. S3 static hosting can be your own URL or your bucket name 1st, s3-website 2nd, followed by the region. AWS are in the process of phasing out Path style, and support for Legacy Global Endpoint format is limited and discouraged. However it is still useful to be able to recognize them should they show up in logs.
>
> - https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html
>
> - https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
>
> - https://docs.aws.amazon.com/general/latest/gr/s3.html
>
> - https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
>
> - https://docs.aws.amazon.com/AmazonS3/latest/dev/HostingWebsiteOnS3Setup.html
>
> - https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/

To read up on:

- S3 buckets name styles

#### Question 17

Question

> You have a database-style application that frequently has multiple reads and writes across the data set. Which of the following AWS storage services are capable of hosting this application? [Select 2]

Answer

> - Elastic File Service (EFS)
> - Elastic Block Store (EBS)

Explanation

> You would either user EBS or EFS. S3 is for object storage, not applications; and Glacier is for data archiving.

To read up on:

- storage services types and usages

#### Question 19

Question

> With which AWS orchestration service can you implement Chef recipes?

Answer

> Opsworks

To read up on:

- OpsWorks

#### Question 20

Question

> A product manager walks into your office and advises that the simple single node MySQL RDS instance that has been used for a pilot needs to be upgraded for production. She also advises that they may need to alter the size of the instance once they see how many people use the system during peak periods. The key concern is that there can not be any outages of more than a few seconds during the go-live period. Which of the following might you recommend, [Select 2]

Answer

> - Convert the RDS Instance to a multi-AZ implementation.
> - Consider replacing it with Aurora before going live.

Explanation

> There are two issues to be addressed in this question. Minimizing outages, whether due to required maintenance or unplanned failures. Plus the possibility of needing to scale up or down.
> Read-replicas can help you with high read loads, but are not intended to be a solution to system outages.
> Multi-AZ implementations will increase availability because in the event of a instance outage one of the instances in another AZs will pick up the load with minimal delay.
> Aurora provided the same capability with potentially higher availability and faster response.

To read up on:

- HA database

#### Question 25

Question

> You work for a famous bakery that is deploying a hybrid cloud approach. Their legacy IBM AS400 servers will remain on-premise within their own datacenter. However, they will need to be able to communicate to the AWS environment over a site-to-site VPN connection. What do you need to do to establish the VPN connection?

Answer

> Set an ASN for the virtual private gateway

Explanation

> The termination IP address on the AWS side is not at the gateway. It is defined as part of the AWS VPN configuration process. Direct Connect could be a carrier, but is not a VPN its self.

To read up on:

- Virtual private gateway

#### Question 27

Question

> Which AWS services allow you to natively run Docker containers? [Select 3]

Answer

> - Elastic Beanstalk
> - Fargate
> - ECS

Explanation

> Although it is possible to run Docker containers on all of the above AWS services, only ECS, Elastic Beanstalk and Fargate allow containers to run natively. EC2 instances can run Docker containers, but Docker has to be installed separately before a container can be deployed.

To read up on:

- Elastic BeanStalk

#### Question 34

Question

> Your Security Manager has hired a security contractor to audit your firewall implementation. When the consultant asks for the login details for the firewall appliance, which of the following might you do? [Select 2]

Answer

> - Create an IAM User with a policy that can read security groups and NACL setting.
> - Explain that AWS implements network security differently and that there is no such thing as a firewall appliance. you might then suggest that the consultant take the **A Cloud Guru** AWS CSA-A course in preperation for the audit.

Explanation

> AWS has removed the Firewall appliance from the hub of the network and implemented the firewall functionality as stateful Security Groups, and stateless subnet NACLs. This is not a new concept in networking, but rarely implemented at this scale.
> In this case an IAM role by itself will not be enough to gain access to the AWS infrastructure - an IAM user will also be required.

To read up on:

- Security

#### Question 35

Question

> Your company likes the idea of storing files on AWS. However, low-latency service of the last few days of files is important to customer service. Which Storage Gateway configuration would you use to achieve both of these ends? [Select 2]

Answer

> - Gateway-Cached
> - File Gateways

Explanation

> Gateway-Cached and File Gateway volumes retain a copy of frequently accessed data subsets locally. Cached volumes offer a substantial cost savings on primary storage and minimize the need to scale your storage on-premises. Note that AWS recently changed the naming. You should know both forms for the exam.

To read up on:

- Gateway types
- File gateways

#### Question 38

Question

> You work for a busy real estate company, and you need to protect your data stored on S3 from accidental deletion. Which of the following actions might you take to achieve this? [Select 2]

Answer

> - Enable protected access using Multi-Factor Authentication (MFA).
> - Enable Versioning on the bucket. If a file is accidentally deleted, delete the delete marker.

Explanation

> The best answers are to allow versioning on the bucket and to protect the objects by enabling protected access using Multi-Factor Authentication

To read up on:

- S3 policies

#### Question 39

Question

> You are reviewing Change Control requests and you note that there is a proposed change designed to reduce errors due to SQS Eventual Consistency by updating the "DelaySeconds" attribute. What does this mean?

Answer

> When a new message is added to SQS queue, it will be hidden from consumer instances for a fixed period.

Explanation

> Poor timing of SQS processes can significantly impact the cost effectiveness of the solution.

To read up on:

- SQS attributes

#### Question 42

Question

> You have been engaged by a company to design and lead the migration to an AWS environment. An argument has broken out about how to meet future Backup & Archive requirements and how to transition. The Security Manager and CTO are concerned about backup continuity and the ability to continue to access old archives. The Senior engineer is adamant that there is no way to retain the old backup solution in the AWS environment, and that they will lose access to all the current archives. What information can you share that will satisfy both parties in a cost-effective manner? [Select 2]

Answer

> - Meet with both parties and brief them on the AWS storage VTL solution. Explain that it can initially be installed in the on-oremuses environment utilizing the existing enterprise backup product to start the transition without losing access to the existing backups and archives, over the duration of the migration, most (if not all) all the backup cycles will be replaced bt the new VTL & VTS Tapes.
> - Suggest that during transition, a second AWS storage gateway VTL solution could be commissioned in the customers new VPPC and integrated with existing VTS. at the same time, the existing Enterprise backup solution could be used to perform tape-to-tape copies to migrate the archives from tapes to VTL/VTS virtual tape.

Explanation

> Any migration project needs to consider how to manage legacy data and data formats. This includes backup and archives. A 3rd party archive service is viable, but would be an ongoing expense. Storage Gateway can be used to efficiently move data into AWS. Old tapes could either be restored to the Storage Gateway volume, or migrated to Virtual tapes inside AWS using Tape Gateway.

To read up on:

- Tape gateway
- migration

#### Question 43

Question

> To establish a successful site-to-site VPN connection from your on-premise network to an AWS Virtual Private Cloud, which of the following might be combined & configured? [Select 4]

Answer

> - An on premises Customer Gateway
> - A private subnet in your VPC
> - A virtual private Gateway
> - A VPC with hardware VPN Access

Explanation

> There are a number of ways to set up a VPN. Based on the options provided, AWS have a standard solution that makes use of a VPC with; a private subnet, Hardware VPN Access, a VPG, and an on-premise Customer Gateway.

To read up on:

- VPN

#### Question 45

Question

> Lambda pricing is based on which of these measurements? [Select 2]

Answer

> - The amount of memory assigned
> - Duration of execution billed in fractions of seconds

Explanation

> Lambda billing is based on both The MB of RAM reserved and the execution duration in 100ms units.

To read up on:

- Lambda pricing

#### Question 50

Question

> When copying an AMI, which of the following types of information must be manually copied to the new instance? [Select 3]

Answer

> - S3 Bucket permissions
> - User Defined tags
> - Launch permissions

Explanation

> Launch permissions, S3 bucket permissions, and user-defined tags must be copied manually to an instance based on an AMI. User data is part of the AMI, itself, and does not need to be copied manually.

To read up on:

- AMI

#### Question 51

Question

> Which of the following are true for Security Groups? [Select 3]

Answer

> - Security Groups evaluate all rules before deciding whether to allow traffic.
> - Security Groups Support "allow" rules only
> - Security Groups operate at the instance level

Explanation

> Security Groups operate at the instance level, they support "allow" rules only, and they evaluate all rules before deciding whether to allow traffic.

To read up on:

- Security Groups vs ACL

#### Question 54

Question

> When you create a custom VPC, which of the following are created automatically? [Select 3]

Answer

> - Route Table
> - Access Control lists
> - Security group

Explanation

> When you create a custom VPC, a default Security Group, Access control List, and Route Table are created automatically. You must create your own subnets, Internet Gateway, and NAT Gateway (if you need one.)

To read up on:

- VPC

#### Question 55

Question

> Which of the following services can invoke Lambda function directly? [Select 3]

Answer

> - Elastic Load Balancer
> - Kinesis Data Firehose
> - API Gateway

Explanation

> API Gateway, Elastic Load Balancer, and Kinesis Data Firehose are all valid ways to directly trigger lambda.

To read up on:

- triggering lambdas

#### Question 56

Question

> You work for a construction company that has their production environment in AWS. The production environment consists of 3 identical web servers that are launched from a standard Amazon Linux AMI using Auto Scaling. The web servers are launched into the same public subnet and belong to the same security group. They also sit behind the same ELB. You decide to do some testing: you launch a 4th EC2 instance in to the same subnet and same security group. Annoyingly, your 4th instance does not appear to have internet connectivity. What could be the cause of this?

Answer

> You have not assigned an elastic IP address to this instance.

To read up on:

- Networking

#### Question 58

Question

> You've been commissioned to develop a high-availability application with a stateless web tier. Identify the most cost-effective means of reaching this end.

Answer

> Use an Elastic Load Balancer, a multi-AZ deployment of an auto-scaling group of EC2 Spot instances (primary) running in tandem with an auto Scaling group of EC2 on-demand instances (secondary), DynamoDb.

Explanation

> With proper scripting and scaling policies, the On-demand instances behind the Spot instances will deliver the most cost-effective solution because the on-demand will only spin up if the spot instances are not available.
> DynamoDB is a regional service, there is no need to explicitly create a multi-AZ deployment. RDS could be used, but DynamoDB lends itself better to supporting stateless web/app installations.

To read up on:

- Autoscaling, spot instances

#### Question 59

Question

> You need to store some easily-replaceable objects on S3. With quick retrieval times and cost-effectiveness in mind, which S3 storage class should you consider?

Answer

> S3 - OneZone-IA

Explanation

> S3 - OneZone-IA has replaced RRS as the recommended storage for when you want cheaper storage for infrequently accessed objects. It has the same durability but less availability. Plus there can be cost implications if you use it frequently or use it for short lived storage. Glacier is cheaper, but has a long retrieval time, and there is no such thing as S3 - Provisioned IOPS.

#### Question 60

Question

> ​Your company has just purchased another company. As part of the merger, your team has been instructed to cross-connect the corporate networks. You run all your confidential corporate services in a VPC and use Route 53 for your Internal DNS. The merged company has all their confidential corporate services and Internal DNS on-premises. After establishing a Direct Connect service between your VPC and their on-premise network, and confirming all the routing, firewalls, and authentication, you find that while you can resolve names against their DNS, the services in the other company are unable to resolve names of your AWS services. Why might this be happening?​

Answer

> By design, the AWS DNS service does not respond to request originating outside the VPC.

Explanation

> Route 53 has a security feature that prevents internal DNS from being read by external sources. The work around is to create a EC2 hosted DNS instance that does zone transfers from the internal DNS, and allows itself to be queried by external servers.

To read up on:

- Route53 and DNS

#### Question 61

Question

> A VPN connection consists of which of the following components? [Select 2]

Answer

> - Virtual Private Gateway
> - Customer Gateway

Explanation

> The correct answers are "Customer Gateway" and "Virtual Private Gateway".
> When connecting a VPN between AWS and a third party site, the Customer Gateway is created within AWS, but it contains information about the third party site e.g. the external IP address and type of routing.
> The Virtual Private Gateway has the information regarding the AWS side of the VPN and connects a specified VPC to the VPN
> "Direct Connect Gateway" and "Cross Connect" are both Direct Connect related terminology and have nothing to do with VPNs.

To read up on:

- Customer gatways

#### Question 64

Question

> You are a solutions architect working for a cosmetics company. Your company has a busy Magento online store that consists of a two-tier architecture. The webservers are behind an Auto Scaling Group and the database is on a Large MySQL instance. Your store is having a Black Friday sale at the end of the week, and having reviewed the performance for the last sale you expect the site to start running very slowly during the peak load. You investigate and you determine that the database was struggling to keep up with the number of reads that the store was generating. How can you successfully scale this environment out so as to increase the speed of the site? [Select 2]

Answer

> - Place the RDS instance behind an ElasticCache instance, then update the connection string.
> - Migrate the database from MySQL to Aurora for better performance, then update the connection string.

Explanation

> Adding a read replica on its own won't solve your problem, you would need to alter the code for Magento to use the read replica (which was not in the offered options). Multi-AZ is a reliability technique not a performance technique. The best answer available is to migrate the database to Aurora which has superior Read performance due to its design. Implementing ElastiCache, is relatively easy and will also offload some of the Read traffic.

To read up on:

- HA RDS

#### Question 65

Question

> On Friday morning your marketing manager calls an urgent meeting to celebrate that they have secured a deal to run a coordinated national promotion on TV, radio, and social media over the next 10 days. They anticipate a 500x increase on site visits and trial registrations. After the meeting you throw some ideas around with your team about how to ensure that your current 1 server web site will survive. Which of these best embody the AWS design strategy for this situation. [Select 2]

Answer

> - Create a duplicate sign-up page that stores registration details in DynamoDB for asynchronous processing using SQS and lambda.
> - Work with your web design team to create some web pages with embedded javascript to emulate your 5 most popular information web pages and sign-up web pages.

Explanation

> A 500x increase is beyond the scope of a well designed single server system to absorb unless it is already hugely overspecialized to accommodate this sort of burst load. An AWS solution for this situation might include S3 static web pages with client side scripting to meet high demand of information pages. Plus use of a noSQL database to collect customer registration for asynchronous processing, and SQS backed by scalable compute to keep up with the requests.
> Lightsail does provide a scalable provisioned service solutions, but these still need to be designed an planned by you and so offer no significant advantage in this situation. A standby server is a good idea, but will not help with the anticipated 500x load increase.

To read up on:

- HA

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

##

#### Question XXXX

Question

>

Answer

>

Explanation

>

To read up on:

- OPsworks?
- ASN
- ProvisionedThroughputExceededException
- Change Control Requests
- NAT source destination check
- SQS - delay seconds
- VPC vs VPG vs VCG
- LightSail?

[main](README.md)
