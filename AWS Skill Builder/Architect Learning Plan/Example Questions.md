<!--
ignore these words in spell check for this file
// cSpell:ignore Postgre Memecahed
 -->

### Test Axioms

<details>
<summary>
Things worth remembering, rule-of-thumb recomendations
</summary>

Resilient:

- Expect "Single AZ" will never be the right answer for resiliency.
- Using AWS manages services should always be preferred.
- Fault Tolerant and High Availability are not the same thing.
- Expect that everything will fail at some point and design accordingly.

Performant:

- if data is unstructured amazon S3 is generally the storage solution
- Use caching strategically to improve performance
- Know when and why to use Auto Scaling
- Choose the instance and database type that makes the most sense for your workload and performance need.

Security:

- Lock down the root user.
- Security Group (statefull) _allow_. Network ACL (stateless) _allow and explicitly deny_.
- Perfer IAM roles over access keys.

Cost Optimization:

- If you know it's going to be on, reserve it.
- Any unused CPU time is a waste of money.
- Use the most cost effective data storage service and class.
- Deterimine the most cost effective EC2 pricing model and instance type for each workload.

Operation Excellence:

- IAM roles are easier and safer than keys and passwords.
- Monitor metrics across the system.
- Automate responses to metrics where appropriate.
- Provide alerts for anomalous conditions.

</details>

### Sample Questions

<!-- <details> -->
<summary>
Sample Questions
</summary>

> A database is running on an EC2 instance, The database software has a backup feature that requires block storage. What Storage Option would be the lowest cost option for the backup data?
>
> - Amazon Glacier
> - EBS Cold HDD Volume
> - Amazon S3
> - EBS throughput optimized HDD volume

The answer is **EBS Cold HDD Volume**, we want _block storage_, which elimates the S3 (and glacier) options, and _low cost_, which pushes towards cold storage rather than the costly IOPS.

> Which of the following AWS services faclitate the implementation of loosely coupled architecture (select two)?
>
> - AWS CloudFormation
> - Amazon Simple Queue Service
> - AWS CloudTrail
> - Elastic Load Balancing
> - Amazon Elastic MapReduce

The answers Are SQS and ELB, sqs provides a message queue that holds messages (even if the consumer or the provider fails), while ELB decouples ip address from instances.\
Aws CloudFormation creates services, CloudTrail is a logging(auditing) tool, and MapReduce is a managed Hadoop service (big data).

> Your web service has a performance SLA to respond to 99% of requests in < 1 second. Under normal and heavy operations, distributing requests over four instances meets performance requirements. What architecture ensures cost efficient high availability of your service if an availability zone becomes unreachable?
>
> - Deploy the serice on four servers in a single Avalability Zone.
> - Deploy the serice on six servers in a single Avalability Zone.
> - Deploy the serice on four servers across two Avalability Zones.
> - Deploy the serice on eight servers across two Avalability Zones.

~~The Answer is 8 servers across two availability zones, so if one fails, we still have 4 servers running.~~\
Actually four servers in two AZ, because the the question mentions cost effectiveness, and high availability

> Your web service has a performance SLA to respond to 99% of requests in < 1 second. Under normal and heavy operations, distributing requests over four instances meets performance requirements. What architecture ensures cost efficient fault tolerant operation of your service if an availability zone becomes unreachable?
>
> - Deploy the serice on four servers in a single Avalability Zone.
> - Deploy the serice on six servers in a single Avalability Zone.
> - Deploy the serice on four servers across two Avalability Zones.
> - Deploy the serice on eight servers across two Avalability Zones.

This time the question focuses on fault tolerance, which is a higher bar - it matches the SLA.

> You are planning to use CloudFormation to deploy a linux EC2 instance in two different regions using the same base Amazon Machine Image (AMI). How can you do this using CloudFormation?
>
> - use two different cloudFormation templates since cloudFormation templates are region specific.
> - use mappings to specify the base AMI since AMI ids are different in each region.
> - use parameters to specify the base AMI since AMI ids are different in each region.
> - AMI Id's are identical across regions

The answer this time is about region and resources. parameters are user input, so it's ruled out. and AMI ids are different across regions.

> How can I access the output of print statements from Lambda?
>
> - SSH into Lambda and look at system Logs
> - Lambda Writes all output to Amazon S3
> - CloudWatch Logs
> - Print statements are ignored in lambda

The answer is CloudWatch logs. Lambdas are created on the fly, so there's no way to ssh into them.

> You are running an EC2 Instance which uses EBS for storing it's data. You take an EBS snapshot everyday. when the system crashes it takes you 10 minutes to bring it up again from the snapshot. What is your RTO and RPO going to be?
>
> - RTO will be 1 day,RPO will be 10 minutes
> - RTO will be 10 minutes,RPO will be 1 day
> - RTO and RTO will be 10 minutes
> - RTO and RTO will be 1 day

This question assumes we know RTO and RPO.

- RTO - recovery time objective - time to restore from backup.
- RPO - recovery point objective - how much data is lost when we restore from backup.

therefore, the answer is RTO - 10 minutes, RPO is one day.

> In what ways does Amazon S3 object differ from block and file storage (select 3)?
>
> - Amazon S3 allows storing an unlimited number of objects
> - object are immutable - the only way to change a single byte is to replace the object
> - object are replicated across availability Zones
> - object are replicated across all regions

The answers are:

- unlimited storage (as opposed to block storage)
- immutability - that's what object are
- replicated across AZ - by default. it's a manged service.

> Which of the following two are features of amazon EBS (elastic block storage) (select 2)?
>
> - Data stored on Amazon EBS is automatically replicated within an Avalability Zone.
> - Amazon EBS data is automatically backed up to tape
> - Amazon EBS volumes can be encrypted
> - Data on Amazon EBS volumes is lost when the attached instance is stopped.

EBS is not instance store, so it's not ephemeral, and can persist. backing up to tape isn't a thing. so the answers are encryption, and replication.

> Which Amazon RDS services engines support read replicas?
>
> - Microsoft SQL server and Oracle
> - MYSQL, MariaDB, PostgreSQL, Aurora
> - Aurora, Microsoft SQL server, Oracle
> - MySQL and PostgreSQL

The answer is MYSQL, MariaDB, PostgreSQL, and Aurora. meaning the microsoft SQL server and Oracle don't support.

> Which AWS datbase service is best suited for non relational databases
>
> - Amazon Redshift
> - Amazon Relation Database Service
> - Amazon Glacier
> - Amazon DynamoDB

The answer is DynamoDB, of course. Glacier is S3 storage class, and redshift is a structural BI solution.

> Which of the following Objects are good candidates to store in cache (select 3)
>
> - Session state
> - Shopping Cart
> - Product catalog
> - Bank account Balance

Product catalog is a definite answer, ~~shopping Cart might be, and maybe also bank account balance?~~ **not bank account balance.**

> Which of the following cache engine are supported by Amazon ElasticCache (select 2)?
>
> - MySQL
> - MemcacheD
> - Redis
> - Couchbase

the answers are memcached and Redis.

> Which services work together to enable auto scaling of EC2 instaces?
>
> - Auto Scaling and Elastic Load balancer
> - Auto Scaling and Cloud Watch
> - Auto Scaling, CloudWatch, Elastic Load balancer
> - CloudWatch and Elastic Load balancer
> - Auto Scaling

The answer is probably all three services, we need the Auto Scaling to perform the action, Cloud Watch for the alerts, and elastic load balancer, because otherwise the new instance won't do anything.

> What is the template the Auto Scaling uses to launch a fully configured instance automatically?
>
> - AMI ID
> - Instance type
> - Key pair
> - Launch configuration
> - User Data

the answer is **Launch Configuration**, it can use all the others inside it, which can be used by it to create the new instances.

> A radio station runs a contest where every day at noon they make an announcement that generates an immediate spike in traffic that requires 8 EC2 instances to process. All other times the webside requires 2 EC2 instances.\
> Which is the most cost effective way to meet these requirements?
>
> - Create an Auto scaling Group with a minimum Capacity of 2 and scale up based on CPU utilization.
> - Create an Auto scaling Group with a minimum Capacity of 8 at all times.
> - Create an Auto scaling Group with a minimum Capacity of 2 and set a schedule to scale up at 11:40 am.
> - Create an Auto scaling Group with a minimum Capacity of 2 and scale up based on memory utilization.

we first elimate the option of always having 8 instances (cost effective). next, the schedule option uses the "Scale Up" terminology, so it's probably a diversion. so either CPU utilization or memory utilization. ~~lets go with memory utilization.~~ this wrong, because,memory utilizations is not metric cloud watch can monitor. and also there's a delay in creating instances, so we need to be prepared before hand.

> An application Runs on EC@ instances in an Auto Scaling Group. the application runs optimally on 9 EC2 instances and must have at least 6 running instances to maintain minimally acceptable performance for a short period of time. which is the most cost-effective auto scaling group configuration that meets the requirements?
>
> - A desired capacity of 9 instances across 2 Avalability Zones.
> - A desired capacity of 9 instances across 3 Avalability Zones.
> - A desired capacity of 12 instances across 2 Avalability Zones.
> - A desired capacity of 1 instances across 1 Avalability Zones.

The answer is 9 instances across 3 AZ, if one AZ fails, we still have 6 instances running, and the autoScaling group will launch the instances at the remaning AZ if needed.

> Which of the following ar characteristics of the auto scaling service on AWS (select 3)?
>
> - Sends traffic to healthy instances.
> - Responds to changing conditions by adding or termination Amazon EC2 instances.
> - Collects and tracks metrics and sets alarms.
> - Deliever push notifications.
> - Launches instances from a specified Amazon Machine Image (AMI).
> - Enforces a minimum number of running Amazon EC instances.

the answer are responding, enforcing a minimum, and launching instances. collecting data and metrics in cloudWatch, sending traffic is ELB, notifications are SNS.

> The Web tier for an application is running on 6 EC2 instances spread across 2 availability Zones. the data tier is a mySQL database running on an EC2 instances. what changes will increase the avalability of the Application (select 2)?
>
> - migrate the MySQl database to a multi-AZ RDS MySQL database instance.
> - increase the instance size of the web tier EC2 instance.
> - Launch the web tier Ec2 instances in an auto sacling group.
> - Turn on CloudTrail in the AWS accout of the application.
> - Turn on cross zone load balancing on the classic Load Balancer.

CloudTrail is user action auding, it has nothing to do with avalability. the instance size of the web-tier doesn't effect availability either. using a multi-AZ database will help if the AZ falls, and an auto-scaling group for the web-tier applications will keep them at the correct amount of instances.
cross zone load balancing won't really help.

> Your AWS account administrator has left your company today. The administrator had access to the root user and a personal IAM administrator account. with these accounts, he generated IAM users and keys.\
> Which of the following should you do today to protect your AWS infrastructure (select 3)?
>
> - Change the password and add MFA to the root user
> - Put an Ip restriction root user logins
> - Rotate keys and change passwords for IAM users
> - Delete all IAM users
> - Delete rhe administrator IAM user
> - Relaunch all EC2 instances with new roles

The answer is first to change root password (and add MFA), also delete the IAM user which he used. launching EC2 roles aren't connected to identities, so there's no point in creating new ones. IAM users are separeate from the user which created them. rotating keys and password for IAM users can help(if he still has access to them). there probably isn't a way to restrict root user access.

> You have deployed an instance running a webserver in a subnet in your VPC. When you try to connect to it through a browser using HTTP over the internet the connection times out.\
> Which of these steps could fix the problem (select 3)?
>
> - Check that the VPC contains an internet gateway and the subnets route table is routing 0.0.0.0/0 to the internet gateway.
> - Check that the VPC contains a Virtual private gateway and the subnets route table is routing 0.0.0.0/0 to the virtual private gateway.
> - Check that the security group allows inbound access on port 80
> - Check that the security group allows outbound access on port 80
> - Check that the network ACL allows inbound Access on port 80

Security group are stateful, so they don't have separate rules for outbound data. ~~probably we care about the Virtual private gateway (incoming data)~~, and the port 80 for ACL.\
Virtual private gateway is about VPN and connecting to outside data centers.

> Which of the following Actions an be controlled with IAM policies (select 3)?
>
> - Creating Tables in a MySQL RDS database
> - Configuring a VPC security group
> - Logging into a .NET application
> - Creating an Orcale RDS database
> - Creating an Amazon S3 bucket

Logging into .Net applications have nothing to do with IAM, it's not an AWS resource. creating resources can be conrolled by AWS policies, so buckets and RDS databases. configuring VPCs is probably also an AWS action, while creating tables inside a MySQL table is governed by the authentication and authorization model of the database.

> You want to create a group of EC2 instances in an application tier subnet that accepts HTTP trafficonly from instances in the web tier (a group of instances in a different subnet sharing a web-tier security group), which of the following will achieve this?
>
> - adding a load balancer in front of the web-tier instances
> - association each instance in the application tier with a security group that allows inbound HTTP traffic from the web-tier security group
> - Adding an ACL to the application tier subnet that allows inbound HTTP traffic from the ip range of the web tier subnet
> - Changing the routing table of the web tier subnet to direct traffic to the application tier instances based on ip address.

A load balancer in the web-tier has nothing to do with this situation.

> You Are asked to make a PDF file publicly available on the web. This file will be downloaded by customers using their browsers millions of times.\
> Which option will be the most cost effective?
>
> - Store the file in S3 standard
> - Store the file in S3 standard-IA
> - Store the file in Glacier
> - Store the file on EFS

The question is between S3 standard and EFS, (IA and Glacier are S3 pricing tiers which are better suited for archive). EFS is a file system (for shared storage), it's not meant to share files (it needs to be mounted), so S3 is the answer.

> To monitor CPU utilization on your RDS instance you set up a CloudWatch alarm with a threshold of 70% over 3 periods of 5 minutes. if CPU utilization goes up to 80% for 10 minutes, how many alarms will you receive?
>
> - 0 Zero
> - 1 One
> - 2 Two
> - 3 Three

zero alarms, cloud watch needs to see 3 occurrences of the metric before triggering the alarm, in this case, there were only 2 periods.

> You are responsible for a web application running on EC2 instances. you want to track the number of 404 errors that users see in the application.\
> Which of the following options can you use?
>
> - Use VPC FlowLogs
> - Use CloudWatch Metrics
> - Use CloudWatch Logs to get the webserver logs from EC2 instance
> - Web application on AWS never have 404 errors

~~probably flow logs?~~ get the cloudWatch logs from EC2 instances and extract from that. VPC flowLogs capture layer 3 and 4 logs, not layer 7.

> You have written an application that needs access to a particular bucket in S3. The application will run on an EC2 instance.\
> What should you do to give the application access to the bucket securely?
>
> - Store your access key and secret access key on the EC2 instance in a file called 'secrets'
> - Attach an IAM role to the EC2 instance with a policy that grants it access to the bucket in S3
> - store your access key and secret key on the EC2 instance in "$Home/.aws/credentials"
> - Use S3 bucket polices to make the bucket public

obviously, we shouldn't store the access key on the virtual machine, and we don't want to make the bucket public to everybody. therefore, we should use IAM roles.

</details>
