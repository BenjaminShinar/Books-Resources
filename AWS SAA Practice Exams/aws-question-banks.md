<!--
// cSpell:ignore 
-->

https://wwww.udemy.com/course/aws-question-banks

## Quiz 1
<!-- <details> -->
<summary>
</summary>

### Q01
>A company is planning to store their archives in AWS.\
Which of the following storage mechanisms provided by AWS would provide an ideal and cost-efficient storage option for storing the Archive data?
> - Amazon S3
> - Amazon Glacier
> - Amazon EBS Volumes
> - Amazon EBS Snapshots

Snapshots are for restarting instances from a saved state, EBS volumes are storage for running instances, S3 is a general solution, while **Glacier** is for long term archives.

> **Explanation**:\
> Amazon Glacier is the best solution for storing archive data. Glacier is meant for cold storage wherein data is not frequently accessed. Since this is used for less frequently accessed data , the cost of this storage is also less than the normal S3 storage. Amazon S3 is meant for object storage that is accessed frequently. Amazon EBS volumes is storage that is attached to EC2 Instances. Amazon EBS Snapshots is used to take a backup of EBS Volumes.

### Q02
> Your company is planning on hosting an application in AWS. This application will allow users to upload videos which will be processed at a later point in time.\
> Which of the following would be the ideal place in AWS to store the videos uploaded by the users?
> - Amazon Glacier
> - Amazon EBS Volumes
> - Amazon S3
> - Amazon EBS Snapshots

Glacier fails because it's for long term objects, EBS volumes and snapshots don't allow easy outside access. **S3** is the beet option.

> **Explanation**:\
Amazon S3 is the ideal place for storage of objects such as files , videos and pictures. For each object you also get a URL which would allow access to the video. Amazon Glacier is normally meant for storing archive data. Amazon EBS volumes is storage that is attached to EC2 Instances. Amazon EBS Snapshots is used to take a backup of EBS Volumes.
### Q03
> Which of the following component of the Cloudfront service allows for content to be cached and then delivered to users across the globe?
> - AWS Regions
> - AWS Availability Zones
> - AWS Edge Locations
> - AWS Data Centers

**Edge locations** are the core part of the AWS CDNS (content delivery network system)

> **Explanation**:\
AWS Cloudfront consists of Edge locations located across the world which are used to cache content and deliver content to users across the world. Content is sent from the Origin to the edge location that is closes to the user. AWS Regions is a geographical location that is used to host a resource. AWS Availability Zones is a logical representation of one or more data centers. AWS Data Center is the physical location where the hardware is placed for hosting AWS resources. 

### Q04
> Your company is planning to host their application on AWS. This application is mission critical and can only afford very little downtime. The application must be available, even in the case of a disaster.\
> Which of the following parts of AWS should be used in the design of the application in AWS
> - AWS Regions
> - AWS Availability Zones
> - AWS Edge Locations
> - AWS Cloudfront

~~**Availability Zones** are inside an AWS region, we use them for highly available architecture.~~

*AZ are for HA, Regions are for DR*

> **Explanation**:\
When your application needs to be available even in the case of a disaster , then you need to ensure that your application is either built to be available across Regions or at least come up in another region if the primary region goes down. Always think of regions when it comes to Disaster recovery. Availability Zones is good when you want to achieve high availability for your application. But in case of a complete region goes down in case of a disaster , then you cannot recover by just deploying applications across availability zones. AWS Cloudfront and Edge locations is used to deliver content to users across the globe

### Q05
>Your company is planning on hosting their testing environment in AWS.\
>Which of the following instance types would be perfect for hosting the EC2 Instances in the testing environment?
> - Spot Instances
> - On-demand
> - Reserved Instances
> - Dedicated Hosts

~~Dedicated hosts are for legal compliance, Spot instances is for workloads. **Reserved** instances give discounts over on-demand.~~\
In this Question, A test environment won't be running for a really long time, so a on-demand is enough.


> **Explanation**:\
On-Demand instances are the most flexible. Here you can spin up and terminate resources as required. Reserved Instances are beneficial for production workloads where you know that you are going to be utilizing resources for a longer period of time. Spot Instances are good for scenarios where you have batch processing needs. In a test environment , let’s say you host a web server and choose a spot instance. If you loose the Spot Instance , then you would need to spin up the Instance all over again

### Q06
> Your company has a set of resources defined in AWS. They are looking at options on how to optimize their workloads on AWS. They are looking at the advantages of using the Trusted Advisor service.\
> For which of the following does the trusted Advisor provide Insights into?\
> Choose 3 answers from the options given below.
> - Performance
> - Cost Optimization
> - Security
> - Fault Tolerance
> - Edge Locations

The trusted Advisor helps with the AWS core pillars:
**Performance**, **Cost Optimization** and **Security**.

> **Explanation**:\
The Trusted Advisor service can give you insights on how to improve the performance of your workloads. Also it can tell you if you are not using capacity that you have provisioned in AWS. This can help from a cost aspect and reduce the amount of spending. Then it can also tell you how to better secure your environment. For example , it can tell you which security groups are left too open.

### Q07
> You are planning on hosting EC2 Instances in AWS. You need to ensure secure access to these EC2 Instances for Administrators.\
> Which of the following can be used to define rules for this sort of secure access?
> - VPC Groups
> - Flow Groups
> - EC2 Groups
> - Security Groups

out of the options, only **Security Groups** are a real thing.

> **Explanation**:\
Using Security Groups , you can define Rules that define the traffic which can flow inside of the EC2 Instance. So if you need to secure access for Administrators from only certain workstations, then you can use Security Groups to define those rules accordingly.

### Q08
> A company has defined a 3-tier application in AWS. The architecture consists of a web , application and database tier. Which of the following AWS Services can be used to monitor the metrics of this architecture. Also, there is a need to monitor all API activity related to this architecture.\
> How can you achieve this?\
> Choose 2 answers from the options given below.
> - AWS SQS
> - AWS CloudWatch
> - AWS CloudTrail
> - AWS VPC

**CloudWatch** monitors performance metrics, while **CloudTrail** audits actions on resources.

> **Explanation**:\
Cloudwatch can be used to monitor your AWS Resources. Here you can use metrics in Cloudwatch to look at different statistics such as the CPU utilization of resources. CloudTrail is an API monitoring tool that is used to monitor for all API activity. Amazon SQS is the queuing service provided by AWS. Amazon VPC is the virtual private cloud that is used to host your EC2 resources.

### Q09
>Your company is trying to establish a hybrid IT environment.\
>Which of the following can provide a dedicated connection to AWS from your on-premise environment?
>- AWS EC2
>- AWS Direct Connect
>- AWS VPN
>- AWS VPC

EC2 is compute, VPC is a logical partition, **Direct Connect** is a way for on-premises to interact with Cloud environment.

> **Explanation**:\
Direct Connect is a service that can help provide a dedicated pipeline from AWS to your on-premise environment. This connection helps in low latency and better bandwidth connections. AWS VPN is used to also connect your on-premise infrastructure to AWS. But this connection is over the internet and is not a dedicated connection. AWS VPC is used to define your EC2 Infrastructure in AWS. AWS EC2 is your cloud compute service.

### Q10
> Which of the following best describes the main feature of an Elastic Load balancer in AWS?\
> Choose an answer from the options below?
> - To evenly distribute traffic among multiple EC2 instaces in separate availability Zones.
> - To evenly distribute traffic among multiple EC2 instaces in a single availability Zones.
> - To evenly distribute traffic among multiple EC2 instaces in multiple Regions.
> - To evenly distribute traffic among multiple EC2 instaces in multiple counties.

counties aren't an AWS Term, ELB are in a single region (for cross regions we use Route53). ~~**probably Single AZ**. Not sure.~~\
ELB are in a single region, across all Avalability Zones.

> **Explanation**:\
Remember that the Elastic Load Balancer can distribute traffic across separate Availability Zones in a region. It cannot distribute traffic across regions. If you want to distribute traffic across regions , then you can use Route 53 for that purpose. Also the entire idea of the load balancer is to ensure that traffic gets directed to instances in multiple Availability Zones. So even if one Zone goes down for any reason ,the Elastic Load Balancer could still distribute traffic to Instances in the other Availability Zone.

### Q11
> Your business continuity team is deciding on a disaster recovery scenario.\
> Which of the following disaster recovery techniques has the highest amount of downtime?
> - Pilot Light
> - Warm Standby
> - Multi-site
> - Backup-Restore

(NEVER HEARD OF ANY OF THEM)

> **Explanation**:\
The Backup option takes time , and then the restoration of the entire environment will just take the same amount of time as the backup. Hence you just have a lot of downtime for the application if this mechanism is employed for disaster recovery. From the above the scenario which will take the least amount of time is the Multi-Site scenario where you already have another side ready and you just need to change the URL to point to the backup site.

### Q12
> Which of the following are advantages when considering the costing aspects of working with resources in AWS when compared to on-premise Infrastructure.
>  Choose 2 answers from the options given below?
> - Lower Upfront Costs
> - Lower number of resources
> - Lower Varying costs
> - Lower Security

lowered security isn't an advantage, ~~on-premises has lower cost variability,~~ while cloud environment has **lower upfront costs** ~~and we can use **lower number or resources** because a lot of AWS is managed.~~\
Cost variability is higher with on-premises because of unexpected changes (like hiring staff), while AWS pricing is more stable (for a consistent amount of usage).

> **Explanation**:\
In AWS , you have lower to almost no upfront costs for hosting resources. Also with on-premise you always have to consider maintenance , employee costing when considering the operating costing for resources. And this always starts varying with complexity in the costing model. But with AWS , the costing model is pretty consistent and only lowers with time.
### Q13
> What is true when it comes to conducting Penetration testing on the AWS Cloud?
> - It is allowed at any point in time
> - It is allowed at any point in time but only on your own resources
> - You need to get prior permission and authorization from AWS before conducting Penetration Testing
> - It is allowed at any point in time as long as you conduct it within AWS limits

~~(NO idea, guessing that we can **only try on our resources**)~~

> **Explanation**:\
AWS has very strict Security policies. Under no circumstance can a customer conduct a Penetration test on their underlying resources without prior permission and authorization.

### Q14
> Your company is planning on initiating their move to the AWS Cloud. They already have resources on their on-premise network which they need to move to AWS.\
> Which of the following is considered when moving infrastructure to AWS in terms of the cost?\
> Choose 2 answers from the options given below.
> - The underlying hardware costs
> - The compute capacity of the database servers
> - The number of servers
> - The number of security groups

we don't care about the security groups for costs. we probably care about the **Number of servers**. unclear if we care about the **compute capacity of the database server** (but the question is worded weirdly).

> **Explanation**:\
When it comes to the calculation of how much it would cost to go to the cloud, you need to consider the number of servers you want to migrate. You would be charged on the individual server basis. Next you also need to understand what the current hardware is assigned to the servers in your on-premise infrastructure. In AWS you need to choose a corresponding Instance Type for the capacity. And this also decides on the cost you pay per server.

### Q15
> You are planning to host a web application using Cloud Front and the AWS Application Load Balancer. You are worried about potential attacks from the Internet.\
> Which of the following services can help assist with DDoS attacks from the Internet.\
> Choose 2 answers from the options given below.
> - AWS WAF
> - AWS SQS
> - AWS Config
> - AWS Shield Advanced

~~WAF is Web application firewall (so probably not here), **SQS** helps us store messags for later consumption, which should help with DDOS.~~ Guessing **Shield Advanced**

> **Explanation**:\
The AWS Shield is a service that can help protect your web application against DDoS attacks. The WAF service is the Web Application Firewall service that can be used to also protect your web site against DDoS attacks.

### Q16
> You have an EC2 Instance that is going to host a database server which is going to experience a high number of Input/Output Operations.\
> Which of the following would be an ideal storage solution for the EC2 Instance?
> - Use Amazon S3 Connected to EC2
> - Use EBS Provision IOPS Volumes
> - Use EBS General purpose SSD Volumes
> - Use Amazon S3-IA access connected to EC2

S3 is a storage, it shouldn't be used as a DB, high IO consumption means which we want **EBS with high IOPS** (IO operation per second).

> **Explanation**:\
Provisioned IOPS volumes are best suited for database workloads which have a high expectation of Input Output operations. The technology used for these underlying volumes is specifically built for working with High Input and Output operations.

### Q17
> You have been instructed by your IT Supervisor to make use of serverless components on the AWS Cloud.\
> Which of the following would you consider to be part of your decision on which component to use?
> - AWS EC2
> - AWS RDS
> - AWS EMR
> - AWS Lambda

EC2 isn't serverless, only Aurora RDS is, EMR (Elastic Map reduce) can be serverless, but **Lambda** is pure serverless.

> **Explanation**:\
AWS Lambda is the serverless computational service that is provided by AWS. Amazon EC2 in a server based compute component. AWS RDS is not a fully serverless component. The Relational database is anyway created on the server. The same goes for AWS EMR , which is the Elastic Map Reduce Service provided by AWS.

### Q18
> Your development community continuously works on various applications and underlying technologies. They have always complained on the amount of time it takes to create a new development environment on the cloud.\
> Which of the following service can be used to quickly spin up development environments on the AWS Cloud?
> - AWS EC2
> - AWS Beanstalk
> - AWS API Gateway
> - AWS Lambda

EC2 and lambda are compute, API gateway is a Restful API endpoint, **Benstalk** allows us to spin up services.

> **Explanation**:\
The Elastic beanstalk service can be used to spin up Development environments on the AWS Cloud. There are several options available , you can create .Net , Java or even custom Docker based environments.

### Q19
> You want to quickly host a wordpress site on the AWS Cloud.\
> Which of the following can assist in this requirement?
> - Use the AWS RDS service to spin up the relevant database
> - Use the AWS Lambda service to spin up a wordpress application
> - Use the AWS marketplace and use a relvent Wordpress Instance
> - Use the AWS marketplace and use a relvent Wordpress AMI

RDS is database, Lambda is computer. the AWS marketplace houses **AMI**s, we use them to create instances.

> **Explanation**:\
The AWS Marketplace has custom built AMI's which can be used to spin up ready made solutions. This is the easiest way to create an environment from already available custom based solutions.
### Q20
> A company is storing their archival data in Amazon Glacier. They need some data from the vault in a time period of 10 minutes.\
> Which of the following needs to be done to fulfill this requirement?
> - Vault Lock
> - Expeideted Retrival
> - Bulk Retrival
> - Standard Retrival

VaultLock is a policy, bulk and standard retrival are slow for glacier, **Expeideted Retrival** is fast and expensive.

> **Explanation**:\
When you place items in Glacier , if you want to retrieve them the normal way , then you need to submit a job. And then it would normally take 3-5 hrs to have the ability to download the object. But if you want the object immediately , then you need to use the Expedited retrieval option.
### Q21
> You want to build a fault tolerant solution for your web application which is hosted on EC2 Instances.\
> Which of the below service can help fulfil this requirement?
> - AWS Elastic Load Balancer
> - AWS Shield
> - AWS WAF
> - AWS EC2

EC2 is compute, Shield and WAF are security, **ELB** helps with fault tolerance and aviability.

> **Explanation**:\
The Elastic Load Balancer can be used to build fault tolerant solution by diverging traffic to multiple Ec2 Instances. If one instance fails , and another one is in place , the Elastic Load Balancer will divert traffic to this EC2 Instance.

### Q22
> Your storage team is looking for options to store objects using the Simple Storage service. The files that are going to be stored are very critical and it needs to be ensured that the most durable option is chosen.\
> Which of the following storage classes would you use for this purpose?
> - Standard Storage
> - Infrequent Access
> - Reduced Redundancy
> - Glacier

~~**Glacier** probably has higher durability.maybe~~ all of them have 99.999_999_999%. 

> **Explanation**:\
Amazon S3 Standard Storage provides the maximum availability and durability of your data so if you have a decision to store critical objects or files , then choose this storage class.

### Q23
> Your current system consists of multiple distributed components. This system is currently hosted on your on-premise environment. The biggest challenge for this application is that whenever one component is changed , the entire system systems goes down with no way to keep the current inflight messages. You are planning on moving this system to the AWS Cloud.\
> Which of the following service can be used to decouple components of an application?
> - AWS EC2
> - AWS SQS
> - AWS Lambda 
> - AWS Config 

Decoupling usually means **SQS**, a way to store messages between services.

> **Explanation**:\
AWS SQS is a messaging system available in AWS. This service can be used to decouple components of an application. So you can push all messages to this service. This service is a highly available and scalable service. So even if components of your applications fail , the messages will still be present in the queue.

### Q24
> Your company is planning to purchase a support plan from AWS. They need to have Operational reviews, recommendations, and reporting for the resources hosted in their account.\
> Which cost efficient plan could they choose?
> - Basic
> - Developer
> - Standard
> - Enterprise

~~(Guessing that basic)~~

> **Explanation**:\
Only the Enterprise support plan has the option for Operational reviews, recommendations, and reporting

### Q25
> Your security team wants to have measures in place when resources start getting created in the AWS Cloud. They want to create privileged users who will certain levels of administrative access on the AWS Cloud.\
> Which of the following can be used to define users and groups in your AWS account?
> - AWS IAM
> - AWS SQS
> - AWS Config
> - AWS Users and Groups

**IAM** is the service which defined users, groups, roles, etc...

> **Explanation**:\
The AWS Identity and Access management feature can be used to define users and groups in your AWS account. Here you can create multiple users , assign permissions to those users based on what type of access is required.

### Q26
> You need to use AWS services that can be used to store files?  Which of the following can be used to fulfill this requirement?\
> Choose 2 answers from the options below.
> - Amazon Cloud Watch
> - Amazon Simple Storage Service (S3)
> - Amazon Elastic Block Store (EBS)
> - AWS Config
> - Amazon Athena

CloudWatch is monitoring, Athena is some sort of RDS advicing, **S3** can store objects (files), **EBS** is a block storage for EC2 machines.

> **Explanation**:\
Amazon S3 is object storage built to store and retrieve any amount of data from anywhere Amazon Elastic Block Store (Amazon EBS) provides persistent block storage volumes for use with Amazon EC2 instances in the AWS Cloud.

### Q27
> Which of the following service make use of AWS edge locations?
> - Amazon Virtual Private Cloud (VPC) 
> - Amazon CloudFront 
> - Amazon Elastic Cloud Compute(EC2) 
> - AWS Storage Gateway

**CloudFront** is the CDN service which uses edge locations. 

> **Explanation**:\
Amazon CloudFront is a web service that speeds up distribution of your static and dynamic web content. This is done via Edge locations which is used to cache content.

### Q28
> Which of the following is a benefit of Amazon Elastic Compute Cloud (Amazon EC2) over physical servers?
> - Automated backup
> - Paying only for what you use
> - The ability to choose hardware vendors
> - Root/ administrator access

we don't choose vendor, access is hardware on the cloud, backups are a feature, so only the option of **Paying for what you use** is valid.

> **Explanation**:\
The biggest advantage is the ability to only pay for what you use. So for On-demand servers , you have an hourly billing concept that is based on the usage of the EC2 Instance. Automated backup is not available in EC2. There are options for backup available , but you still need to choose them accordingly.

### Q29
> Which AWS service provides infrastructure security optimization recommendations?
> - AWS Price List Application programming interface
> - Reserved Instance
> - AWS Trusted Advisor
> - EC2 Spot Fleet

**Trusted advisor** helps compling to the "well-architecture model"

> **Explanation**:\
The AWS Trusted Advisor gives you security recommendations and can help you improve the overall way you utilize resources in your AWS Account.

### Q30
> You need to collect and track metrics for various AWS services.\
> Which of the following AWS service can help fulfil this requirement?
> - Amazon CloudWatch
> - Amazon CloudFront
> - Amazon CloudSearch
> - Amazon Machine Learning (ML)

**CloudWatch** tracks metrics. Cloud Front is CDN, Machine learning has nothing to do with this issue.

> **Explanation**:\
Amazon CloudWatch is a monitoring service for AWS cloud resources and the applications you run on AWS. You can use Amazon CloudWatch to collect and track metrics, collect and monitor log files, set alarms, and automatically react to changes in your AWS resources
### Q31
>A company needs to know which user was responsible for terminating several critical Amazon Elastic Compute Cloud (Amazon EC2) Instances.\
>Where can the customer find this information?
> - AWS Trusted advisor
> - EC2 instance usage report
> - Amazon CloudWatch
> - AWS Cloud Trail Logs

Cloudwatch monitors metrics, **Cloud Trail** audits actions

> **Explanation**:\
The AWS Cloud Trail service logs all API calls and can be used to track who terminated the instance. Every API call that gets triggered in AWS gets recorded in AWS Cloudtrail.

### Q32
> You want to register a new domain with AWS. Which of the below service would you use?
> - Route53
> - CloudFront
> - Elastic Load Balancing
> - Virtual Private Cloud

**Route53** is amazon DNS

> **Explanation**:\
Route 53 is the domain name system available in AWS. Route53 allows for registration of new domain names in AWS.
### Q33

> What is the value of having AWS Cloud services accessible through an Application Programming Interface (API)?
- Cloud resources can be managed programmatically
- AWS infrastructure will always be cost optimized
- All application testing is managed by AWS
- Customer --owned, on-premises infrastructure becomes programmable

API aren't magic! they just allow to **sends commands** without the management console.

> **Explanation**:\
With the advantage of having API access to resources , you can manage the AWS in a programmatic fashion. You don’t need to log into the console to manage resources.
### Q34


> Which of the following examples supports the cloud design principle "design for failure and nothing will fail"?
> - Adding an elastic load balancer in front of a single EC2 instance
> - Deploying an application in multiple AZ
> - Creating and deploying the most-cost effective option
> - Using CloudWatch alerts to monitor performance

**Multiple Availability Zones** means HA architecture which handles stress and occasional failures of other services without failing.

> **Explanation**:\
By deploying your solution to multiple Availability Zones , that means that you are deploying your solution to multiple data centers. And that means you are designing for failure in case any data center goes down.

### Q35
> Which service allows an administrator to create and modify AWS user permissions?
> - AWS Config
> - AWS Cloud Trail
> - AWS Key Management Service
> - AWS Identity and Access Management

**IAM** controls access and users.

> **Explanation**:\
AWS Identity and Access Management (IAM) is a web service that helps you securely control access to AWS resources.

### Q36
> Which of the below service can be used to build a fully managed petabyte-scale data warehouse on the AWS Cloud?
> - Amazon Redshift
> - Amazon DynamoDb
> - Amazon ElasticCache
> - Amazon Aurora

DynamoDb is document based, Aurora is RDS, Elastic Cache is a cacheing layer, **Redshift** is Big Data (date warehouse).

> **Explanation**:\
Amazon Redshift is a fully managed, petabyte-scale data warehouse service in the cloud. You can start with just a few hundred gigabytes of data and scale to a petabyte or more.

### Q37
> Which of the following is the responsibility of the AWS customer according to the Shared Security Model?
- Securing Edge Locations
- Managing AWS Identity and Access Management
- Monitoring physical device security
- implementing service organization control standards

Real World stuff is aws, **user and application stuff is the customer.**

> **Explanation**:\
The responsibility of managing the various permissions of users and the roles and permission is with the AWS customer. AWS is responsible for managing other controls such as the security of its physical data centres and edge locations.

### Q38
> You want to get more information on the costs incurred for your running EC2 Instances.\
> Where can you get the relevant information on this?
> - EC2 Dashboard
> - EC2 Cost and Usage reports
> - AWS Trusted advisor dashboard
> - AWS Cloud Trail logs in S3

EC2 dashboard is for machine state, trusted advisor can help with reducing costs, but the **Cost and Usage** report is what matters.

> **Explanation**:\
Cost Explorer is a free tool that you can use to view your costs. Here you can get the present view for the month on your spending. You can also get a forecast of your future possible spending and also see detailed breakdown of your current and past spending.

### Q39
> Which of the following entity has complete control over an AWS account?
> - AWS support team
> - AWS security team
> - AWS account owner
> - AWS technical account manager

the strongest user is the root user, the **account owner**.

> **Explanation**:\
Remember you as the account owner has complete control over the account and the resources defined in it.

### Q40
> One main design concept of developing application architectures is decoupling.\
> What is meant by this concept?
> - Enable data synchronization across the web application layer
> - create a tightly integrated application
> - Reduce inter dependencies so failures do not impact other components
> - have the ability to execute automated bootstrapping actions

Decoupling means that we **reduce dependencies** of components.

> **Explanation**:\
One of the most important design concepts for applications is to ensure that components of a system are lightly coupled , which means they don’t have too much of a dependency on each other. If you have a tightly coupled system , then failure in one component could cause a failure in other components.

### Q41
> Which of the following is a benefit of running an application across two Availability Zones?
> - Performance is improved over running in a single avaliability zone
> - It is more secure than running in a single avaliability zone
> - It incrases the availability of the application compared to running in a single avaliability zone
> - It significantly reduces the total cos of ownership versus running in a single avaliability zone

Avaliability zones are always about making the application **Highly Availbe**.

> **Explanation**:\
By running the application across multiple availability zones , means that the application is running off multiple data centres. It also means that if one data centre were to go down , the application would still be available since it has a component running in another Availability zone or data centre.

### Q42
> Which of the following security requirements are managed by AWS customers?\
> Select 2 answers from the options given below.
> - Password Policies
> - Physical security
> - User permissions
> - Disk Disposal
> - Hardware Patching

AWS manages the physical side, the customer manages **passwords** and **permissions**.

> **Explanation**:\
The hardware and physical security is maintained by AWS. It’s the Password policies and User permissions which can be maintained by the customer.


### Q43
> Which of the following is in line with the concept of Elasticity when it comes to the design of an application?
> - Create systems that scale to the required capacity based on changes in demand
> - Minimize storage requirement by reducing logging and auditing activities
> - Enable AWS to automatically select the most cost-effective services
> - Accelerate the design process because recovery from failure is automated

The only option that makes senses is that of promoting **Scaling**.

> **Explanation**:\
This is the concept of seeing how flexible your application is to demand. So if your demand is high , then your infrastructure should have the ability to scale accordingly. And if the demand goes down , the infrastructure should be scaled back down to ensure cost is managed effectively.

### Q44
> Which of the following workloads are best suited for AWS Spot Instances?
> - Workloads that are run only in the morning and stopped at night
> - Workloads that are critical and need Amazon EC2 instances with termination protection
> - Workloads where the avilability of the EC2 instances can be flexible
> - Workloads that need to run for long periods of time and can be interrupted at any time

the timing is not critical, so no scheduled jobs, spots can be lost, so not critical jobs. application availability isn't related, so only workloads which can be **interrupted and resumed**.

> **Explanation**:\
Spot Instances are a cost-effective choice if you can be flexible about when your applications run and if your applications can be interrupted.

### Q45
> Which AWS feature enables a user to manage services through a web-based user interface?
> - AWS Management Console
> - AWS API
> - AWS Software development kit (SDK)
> - Amazon Cloud Watch

API is a set of commands, SDK allows other programs to interact with the SDK, cloud watch is monitoring, so **Management console**.

> **Explanation**:\
The AWS Console is the single place where you can use the Web interface to manage all your AWS resources

### Q46
> Which tool can display the distribution of AWS spending? You need to also see the forecast of your expenses.
> - AWS Organization
> - AWS Dev Pay
> - AWS Trusted Advisor
> - AWS Cost Explorer

The **Cost Explorer** shows forecasts and where money is spent.

> **Explanation**:\
Cost Explorer is a free tool that you can use to view your costs. Here you can get the present view for the month on your spending. You can also get a forecast of your future possible spending and also see detailed breakdown of your current and past spending.

### Q47
> How can you add an extra layer of unauthorized access in the AWS Management Console for the current users who are defined in IAM?
> - Set up a secondary password
> - Apply multi-factor authentication (MFA)
> - Request root access privileges
> - Disable AWS console access

**MFA** gives us better protection.

> **Explanation**:\
AWS Multi-Factor Authentication (MFA) is a simple best practice that adds an extra layer of protection on top of your user name and password.

### Q48
> You are planning for a disaster recovery procedure for your AWS infrastructure. When you look towards building a disaster recovery plan.\
> Which of the following should be used to spin up your backup resources?
> - Subnet
> - Edge location
> - VPC
> - Region

???(Unclear question)

> **Explanation**:\
If you need to have a disaster recovery mechanism in place , then plan to place your resources in multiple regions
### Q49
> Which of the following is a factor when calculating Total Cost of Ownership (TCO) for the AWS Cloud?
> - The number of servers migrated to aws
> - The number of users migrated to aws
> - The number of passwords migrated to aws
> - The number of keys migrated to aws

Costs are based on **servers** (compute,storage and requests), users, passwords and keys don't matter.

> **Explanation**:\
Since EC2 Instances carry a charge when they are running, you need to factor in the number of servers that need to be migrated to AWS.

### Q50
> Which of the following is a fully managed NoSQL database on the AWS Cloud?
> - RDS
> - DynamoDb
> - Redshift
> - MongoDb

**DynamoDB** is the serverless noSql document-based database. 

> **Explanation**:\
Amazon DynamoDB is a fast and flexible NoSQL database service for all applications that need consistent, single-digit millisecond latency at any scale. It is a fully managed cloud database and supports both document and key-value store models

### Q51
> A company wants to start shifting their servers to the AWS Cloud.\
> Which of the following is used to house servers in the AWS cloud?
- AWS EC2
- AWS VPC
- AWS REGIONS
- AWS Availability Zones

~~servers are hosted as **EC2** machine~~\
But they are hosted at VPC...

> **Explanation**:\
The AWS VPC is the equivalent of a private data centre. This is carving out of the portion of the AWS Cloud. AWS EC2 is the equivalent of the virtual servers , but the underlying place where the EC2 servers are placed in the AWS Virtual Private Cloud.

### Q52
> A company wants to start shifting their servers to the AWS Cloud. Due to strict compliance regulations , they need to have their own independent hardware on the AWS Cloud. They cannot share the hardware with anyone else.\
> Which of the following types of servers on the AWS Cloud meets this demand?
> - Reserved Instances
> - Dedicated Instances
> - Dedicated Hosts
> - On-demand Instances

only **Dedicated hosts** relates to hardware.

> **Explanation**:\
Dedicated Hosts provide dedicated hardware that is fully reserved for the AWS customer. Dedicated Instances are also dedicated , but this only at the Instance level. The hardware is still shared with other AWS customers.

### Q53
> You want to migrate your existing MySQL database to the AWS Cloud.\
> Which of the following is a MySQL compatible database on the AWS Cloud?
> - AWS Aurora
> - AWS DynamoDB
> - AWS EC2
> - AWS VPC

**Aurora** is a serverless RDS engine.

> **Explanation**:\
Amazon Aurora is a MySQL and PostgreSQL compatible relational database engine that combines the speed and availability of high-end commercial databases with the simplicity and cost-effectiveness of open source databases.

### Q54
> An application needs to be migrated to the AWS Cloud. It needs to be ensure that this application is PCI compliant as per compliance needs. How can you accomplish this?\
> Choose 2 answers from the options given below.
> - Ensure to check which services are PCI-compliant
> - Ensure your application has the necessary checks for PCI-compliance
> - All AWS service are PCI-compliant, so any service can be used.
> - Once you deploy your application at AWS, it automatically becomes PCI-compliant

AWS isn't magic, we need to check our own **application**, and check which services are **compliant**.

> **Explanation**:\
Always check and ensure that the services which you are going to use in AWS is PCI Compliant. AWS has a comprehensive security program wherein they certify a lot of their services with many security and compliance programs. But you still need to confirm based on the service you are going to use. Also at the application level you need to guarantee that the necessary checks are in place to achieve PCI Compliance.

### Q55
> Your company is planning on purchasing a support plan. As part of the requirements drafting , it is mentioned that a Support Concierge is required as part of the support plan.\
> Which of the following support plans meets this demand?
> - Basic
> - Developer
> - Standard
> - Enterprise

guessing that **Enterprise** is the highest and get a human.

> **Explanation**:\
Only the Enterprise Plan has access to the Support Concierge , so you would need to choose this plan.

### Q56
> Which of the following service allows you to define Infrastructure as Code?
> - AWS EC2
> - AWS Lambda
> - AWS Cloud Formation
> - AWS Elastic Beanstalk

**Cloud Formation** is like terraform.

> **Explanation**:\
The AWS Cloudformation service is used to define templates which is equivalent to Infrastructure as code. Using this you can make templates run which in turn would create resources in the AWS Cloud.

### Q57
> You are hosting AWS resources in the Cloud. You want to be notified in case of issues in the underlying hardware that hosts the AWS Resources.\
> Which of the following service can help in this regard?
> - AWS personal health Dashboard
> - AWS CloudTrail
> - AWS VPC
> - AWS Consolidated billing

~~**Cloud Trail**~~\
CloudTrail is auditing, not health. 

> **Explanation**:\
The AWS Personal Health Dashboard displays the health of the underlying hardware resources. They also can be used to give you alerts in case there are any issues in the underlying hardware.
### Q58
> You are planning on hosting EC2 Instances on the AWS Cloud. You need to ensure scalability of the Instances based on demand.\
> Which of the below services can help for such a requirement.\
> Choose 2 answers from the options given below.
> - AWS Elastic load balancer
> - AWS Autoscaling
> - AWS VPC
> - AWS Elastic Beanstalk

The **Load Balancer** works together with the **scaled up groups** to distrbute work across many instances.

> **Explanation**:\
When you talk about scalability , look towards the combination of Elastic Load Balancer and Autoscaling. Autoscaling can be used to scale resources based on demand. And the Elastic Load Balancer can be used to distribute load to resources evenly.

### Q59
> Which of the following features of the AWS Relational database service provides the facility of High availability?
> - Snapshots
> - Multi AZ
> - Backup
> - Multi-VPC

snapshots and backups are distaster recovery, **Availability zones** are for highly available architecture.

> **Explanation**:\
Having snapshots and backup can guarantee safety up to a point. But if you need to have high available then you need to have the Multi-AZ feature in place. This feature will automatically create a backup database in another availability zone.
### Q60
> Which of the following are cost benefits when it comes to On-demand EC2 Instances?\
> Choose 2 answers from the options given below.
> - Pay per day for compute capacity
> - Pay per second for compute capacity
> - Payment of partial upfront costs
> - No upfront costs

Aws prices based on **compute time in seconds**, not days. unlike reserved instances, on-demand machine **have no upfront costs**.

> **Explanation**:\
When it comes to On-demand EC2 Instances , you don’t need to pay any upfront costs and only pay per second for the EC2 Instance.


### Q61
> When creating EC2 Instances, which of the following properties of EC2 instances contribute to the cost of the EC2 Instance?\
> Choose 2 right options.
> - Instance Type
> - Keys Assigned to the Instance
> - Type of Operating system
> - Tags assigned

we pay for **instance type** (machine power), and which **OS** it's running.

> **Explanation**:\
The Instance type defines different factors such as the CPU and Memory assigned to the Instance. So if you choose a higher configuration , you will pay more. Also it depends on the type of operating system. The charge for the OS is also a part of the total cost incurred.

### Q62
> Which of the following services is a fully managed NoSQL database on the AWS Cloud?
> - AWS DynamoDb
> - AWS Aurora
> - AWS MySQL
> - AWS Oracle

oracle and MySQL are proprietary engines, aurora is serverless RDS, **DynamoDB** is serverless noSQL.

> **Explanation**:\
AWS DynamoDB is a fully managed NoSQL database on the AWS Cloud.

### Q63
> Senior Management needs to understand what is the cost benefit of moving assets from their on-premise environment to the AWS Cloud.\
> Which of the following can help in this assessment?
> - AWS Trusted Advisor
> - TCO calculator
> - AWS inspector
> - AWS personnel health dashboard

**TCO**, because the rest look at cloud resources.

> **Explanation**:\
You can use the TCO Calculator to do the cost benefit analysis of hosting resources on the AWS Cloud.

### Q64
> Which of the following should be considered when you are planning on deploying an application to the AWS Cloud?
> - The amount of hardware you have to purchase
> - Inspection of the data centre
> - Putting resources in the region which is closest to the customer
> - Plan towards making a large capital investment

The **closer the servers** are to the user, the better.

> **Explanation**:\
You should identify regions which are closest to the customer and ensure the resources are deployed to these regions. This would give a better user experience and least latency for the users.

### Q65

> Which of the following storage devices is used to store data that is attached to EC2 Instance?
> - AWS Simple Storage Service (S3)
> - AWS Elastic Block Storage
> - AWS Glacier
> - AWS SQS

Glacier is a type of S3, SQS is a queue. **EBS **is a block storage device that can be attached to EC2 instances.

> **Explanation**:\
The AWS Elastic Block Storage is used to attach storage to EC2 Instances. You can attach multiple EBS volumes to an EC2 Instance.

</details>

## Quiz 2
<details>
<summary>

</summary>

### Q01
### Q02
### Q03
### Q04
### Q05
### Q06
### Q07
### Q08
### Q09
### Q10
### Q11
### Q12
### Q13
### Q14
### Q15
### Q16
### Q17
### Q18
### Q19
### Q20
### Q21
### Q22
### Q23
### Q24
### Q25
### Q26
### Q27
### Q28
### Q29
### Q30
### Q31
### Q32
### Q33
### Q34
### Q35
### Q36
### Q37
### Q38
### Q39
### Q40
### Q41
### Q42
### Q43
### Q44
### Q45
### Q46
### Q47
### Q48
### Q49
### Q50
### Q51
### Q52
### Q53
### Q54
### Q55
### Q56
### Q57
### Q58
### Q59
### Q60
### Q61
### Q62
### Q63
### Q64
### Q65

</details>

## February 2018 - Quiz 1
<details>
<summary>

</summary>

### Q01
### Q02
### Q03
### Q04
### Q05
### Q06
### Q07
### Q08
### Q09
### Q10
### Q11
### Q12
### Q13
### Q14
### Q15
### Q16
### Q17
### Q18
### Q19
### Q20
### Q21
### Q22
### Q23
### Q24
### Q25
### Q26
### Q27
### Q28
### Q29
### Q30
### Q31
### Q32
### Q33
### Q34
### Q35
### Q36
### Q37
### Q38
### Q39
### Q40
### Q41
### Q42
### Q43
### Q44
### Q45
### Q46
### Q47
### Q48
### Q49
### Q50
### Q51
### Q52
### Q53
### Q54
### Q55
### Q56
### Q57
### Q58
### Q59
### Q60
### Q61
### Q62
### Q63
### Q64
### Q65

</details>

## February 2018 - Quiz 2
<details>
<summary>

</summary>

### Q01
### Q02
### Q03
### Q04
### Q05
### Q06
### Q07
### Q08
### Q09
### Q10
### Q11
### Q12
### Q13
### Q14
### Q15
### Q16
### Q17
### Q18
### Q19
### Q20
### Q21
### Q22
### Q23
### Q24
### Q25
### Q26
### Q27
### Q28
### Q29
### Q30
### Q31
### Q32
### Q33
### Q34
### Q35
### Q36
### Q37
### Q38
### Q39
### Q40
### Q41
### Q42
### Q43
### Q44
### Q45
### Q46
### Q47
### Q48
### Q49
### Q50
### Q51
### Q52
### Q53
### Q54
### Q55
### Q56
### Q57
### Q58
### Q59
### Q60
### Q61
### Q62
### Q63
### Q64
### Q65

</details>

## February 2018 - Quiz 3
<details>
<summary>

</summary>

### Q01
### Q02
### Q03
### Q04
### Q05
### Q06
### Q07
### Q08
### Q09
### Q10
### Q11
### Q12
### Q13
### Q14
### Q15
### Q16
### Q17
### Q18
### Q19
### Q20
### Q21
### Q22
### Q23
### Q24
### Q25
### Q26
### Q27
### Q28
### Q29
### Q30
### Q31
### Q32
### Q33
### Q34
### Q35
### Q36
### Q37
### Q38
### Q39
### Q40
### Q41
### Q42
### Q43
### Q44
### Q45
### Q46
### Q47
### Q48
### Q49
### Q50
### Q51
### Q52
### Q53
### Q54
### Q55
### Q56
### Q57
### Q58
### Q59
### Q60
### Q61
### Q62
### Q63
### Q64
### Q65

</details>

## February 2018 - Quiz 4
<details>
<summary>

</summary>
</details>



