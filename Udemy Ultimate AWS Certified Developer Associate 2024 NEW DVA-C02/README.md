<!--
// cSpell:ignore boto xlarge POSIX Proto AWSELB AWSALBTG AWSALBAPP NAPTR NACL DSSE ONTAP ebextensions Flink Distro PITR
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css"> 

# AWS Certified Developer Associate 2024 NEW DVA-C02

Udemy course [Full Practice Exam with Explanations included! PASS the Amazon Web Services Certified Developer Certification DVA-C02](https://www.udemy.com/course//aws-certified-developer-associate-dva-c01/). by _Stephane Maarek_.


[sample code and slides](https://courses.datacumulus.com/downloads/certified-developer-k92/)

## Getting Started With AWS
<details>
<summary>
Getting Started.
</summary>

Started as an internal service in 2002, then went public over the tears, offering more services and at more locations. used by many leading companies. allows to build sophisticated and scalable applications, it can be used in many ways, some are general for data centers, and some are unique to the cloud.

### Regions and Availability Zones

[The aws infrastructure](https://infrastructure.aws/)

- AWS Regions
- AWS Avalability Zones
- AWS Data Centers
- AWS Edge Locations/ Points of Presence

regions have a name, it's a cluster of data centers in a geographical region. most aws services are region-scoped. we choose a region based on several factors:

- Compliance with data governance and legal requirements
- Proximity to customer for reduced latency
- Available services - not all new features are suppoerte in all regions
- Pricing varies region to region

Regions are divide into Availability zones, usuall 3 zones in a reason. each AZ is one or more data-centers, AZ are separated from one another, so they shouldn't go down if there is power outage. they are linked between them with high speed connection.

Edge locations are points of presence that help with getting lower latency and content caching.

### Tour of the Console and Services

global scoped services:
- IAM
- Route53 (DNS)
- CloudFront (content Delivery network)
- WAF (web application firewall)

Most AWS services are region-scoped.

In the web console, we can choose the region. we usually work with a region that is close to us. but we can choose any region we want. if we choose a global service, then our selected region shows "global".\
In the main page we see the recently used services, and some other stuff. we can navigate services by category or by lexical order, or search them directly in the search bar.

AWS is constantly changing the layout, so some services look different than what is shown in the course.

### Budget Setup

Setting an alarm to avoid spending, we go to the billing and cost management service (as a root user) and allow our administration users to see their billing pages.

in our other accounts, we can now see charges by services and regions. this shows us what we are paying for. it's good to look at this every month to see that we don't forget anything. we can also look at the "free-tier" dashboard and check if we are going overboard and using more than what the free-tier limit is.

we can also set up an alarm that uses the pre-built "Zero-Spent budget", which will alarm us if we spend even a single cent.  we can also set a monthly budget alarm and be alerted when get close to the threshold.

</details>

## IAM & AWS ClI
<details>
<summary>
Identity And Access Management.
</summary>

This a global service.

when we create a AWS account, we create a root account, we should never share this account, and limit how much we use it. this is the strongest, unlimited user. we should create an administrator user for our management tasks, and create users for our people who might need to use our AWS account.

### Users and Groups

users are people within the organization, which could be grouped. groups only contain users, they can't be nested inside each other. user don't have to belong to a group, and they can belong to multiple groups.

we create users and groups to handle permissions. permissions are stored in IAM policies, json documents that specify what the user can and can't do, and on which services. we don't give more permission than needed, we follow the "least privilege principle".

to create a user, we go to the <cloud>IAM</cloud> dashboard and click <kbd>Create user</kbd>. we will start by creating an admin user to use instead of the root user. we choose the user name and select that it can use the web console. for now we create the IAM user in the old way, but we can also create a user in the identity center. we write our password, create a group "admins" with the administrator policy. we can add tags if we wish (most resources can be tagged).

users can have policies attached directly or through groups.

an account can have an alias, which makes signing into the console easier, this account alias must be globally unique (across all aws account in the world).

### Policies

as stated before, IAM policies manage permission to aws resources.

if we attach a policy to a group, then all member of the group get those permissions. a user can get policies from multiple groups.

the structure is a json file with statements, each statements has
- "Sid" - optional identifier
- "Effect" - "allow" or "deny"
- "Principal" - account/user/role/service to which this policy applies to
- "Action" - what can be done
- "Resources" - on which resources the actions are applied
- "condition" - when this policy is in effect. also optional

we can play with users and remove the admin user from the administrator group, and then that user won't even be able to see it's own permission. we can assign a different policy such as IAMReadOnlyAccess and control what that user can do.

### Multifactor Authentication



in IAM, we can create a password policy, and enforce certain standards on the passwords. we can also require users to change the password after some time(password expiration), and prevent password re-use.

but another mechanism is the MFA (Multi Factor Authentication), this combines the password with another layers, such as a device (phone application).

- virtual MFA device (authenticator application)
- Universal 2nd Factor security key (U2F) - physical device
- Hardware key MFA device - can be provided by a third part

in a demo, we define the password policy, either use the default or customize them.

for the root account, we can set the MFA, we follow the wizard and connect the device we want. now when we want to log-in into the root user, we have to provide the authentication from the device.

### Accessing AWS - Access Keys, CLI, SDK
we can access AWS in different ways:

- AWS management Console (website) - password and MFA
- AWS Command Line Interace - protected by access keys
- AWS Software Developer Kit - for code, uses access keys

Access Keys are secret, just like passwords, and they need to be protected, never share them.

The CLI is the command line interface, for use in text based terminals. the cli commands usuall y look like `aws <service> <command>`. the SDK is langauge specific API (library) that we use in our applications to communicate with AWS services. for this course, we will occasionally use the python sdk, called "boto".

we can install the AWS CLI on our windows machine, we follow the wizard and use the MSI installer. we upgrade the version by downloading the most recent installer. for Mac, we download the ".pkg" installer, for Linux we download the file using `curl`, unzip it and run the install script.

if we want to connect to our cloud, we need access keys, we get the by creating access keys from our users. (there are better ways, but this is also important).

with these keys, we can configure our CLI profile, and start playing with it. the access keys use the same permissions as the console user.

There is an easy alternative for using access keys, we can use CloudShell instead. this is a built-in terminal that has aws installed by default, and uses the same permissions as the user that's logged in. the environment is persistent so we can keep some scripts there if we wish, and download/upload files to it.

### Roles

Roles are like users, they also have permissions and policies, but we give them to aws services (resources), such EC2 machines. we give them roles to control what they can do.

when we create roles, we usually use the **AWS service role**, and then we choose the service we want to attach the role on (which controls the "principal" section), and then we select what the role can do. we eventually can assign the role to instaces (such as EC2 machine), which will be done later.

### Security Tools

IAM credentials report - account-level report, last access, password rotation, etc...

Access Advisor - user level, see what the user can do and what permission they have and don't use.

We can do both in the IAM console.

best practices:

> - Don't use the root account except for AWS account setup
> - One physical user = One AWS user
> - Assign users to groups and assign permissions to groups
> - Create a strong password policy
> - Use and enforce the use of Multi Factor Authentication (MFA)
> - Create and use Roles for giving permissions to AWS services
> - Use Access Keys for Programmatic Access (CLI / SDK)
> - Audit permissions of your account using IAM Credentials Report & IAM Access Advisor
> - Never share IAM users & Access Keys

Shared Responsibility model - separate what the vendor is responsible to and what the user (you) is responsible to in terms of security. this changes for each service, so for IAM service, there are things that the user must do to secure the account.

> AWS: 
>
> - Infrastructure (global network security)
> - Configuration and vulnerability analysis
> - Compliance validation
> 
> You:
> 
> - Users, Groups, Roles, Policies Management and monitoring
> - Enable MFA on all accounts
> - Rotate all your keys often
> - Use IAM tools to apply appropriate permissions
> - Analyze access patterns &
review permissions 

### Summary

AM Section - Summary
> - Users: mapped to a physical user, has a password for AWS Console
> - Groups: contains users only
> - Policies: JSON document that outlines permissions for users or groups
> - Roles: for EC2 instances or AWS services
> - Security: MFA + Password Policy
> - AWS CLI: manage your AWS services using the command-line
> - AWS SDK: manage your AWS services using a programming language
> - Access Keys: access AWS using the CLI or SDK
> - Audit: IAM Credential Reports & IAM Access Adviso

</details>

## EC2 Fundamentals
<details>
<summary>
EC2 - Elastic Compute Cloud
</summary>

IAAS - Infrastructure as a Service

> - Renting virtual machines (EC2)
> - Storing data on virtual drives (EBS)
> - Distributing load across machines (ELB)
> - Scaling the services using an auto-scaling group (ASG)

the basic compute service, a virtual machine running on the cloud.

### EC2 Basics

when we create an EC2 instance, we choose the configuration.

- Operating System - linux, windows, Max
- Compute - Cpu Cores
- Random Access Memory
- Storage
  - hardware attached (EC2 Instance Store)
  - network attached (Elastic Block System, Elastic File System)
- Networking - speed, public IP address
- Firewall rules - Security Groups
- Bootstrap script - EC2 User Data

the user script controls what the machine does on startup, it's a way to automate boot tasks, such as installing updates, software, downliding and starting programs. all of this is done with the machine root user.

there are many instances types, they have names such as "t2.micro", "m5.large" and all kinds of others, each instance type has different configuration, which should fit different use cases.

the "t2.micro" is part of the free-tier, and that's what we use.

we will create the first EC2 machine from the web console, it will be a web server, and we will give it a user data script.

in the EC2 service, we select <kbd>Launch Instance</kbd>, we give the instance a name and tags, we choose the base image, and select the Amazon Linux AMI for now. we could create our own AMI if we need. we select the instance type that is free-tier eligible.\
We next create a key-pair to connect to the ec2 machine, so we need to create one. we use the RSA type and choose the ".pem" option for mac, linux and windows 10 and above, ".ppk" is for Putty in older windows machines.\
We modify the network settings, and give us a public ip, we want to have network access in SSH from anywhere, and HTTP access form anywhere.\
We move to storage option, and configure the root volume storage (nothing to do here). we keep the rest of the option as default, and in the user script option, we paste a script that start an httpd server.

```sh
#!/bin/bash
# Use this for your user data (script from top to bottom)
# install httpd (Linux 2 version)
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
echo "<h1>Hello World from $(hostname -f)</h1>" > /var/www/html/index.html
```

we can launch the instance now, and when it's done, we have a running instance. we  see some details about the instance. we use the public ip address to navigate to our machine in the browser (we might need to specify the http protocol, depending on the browser).

we can stop the instance as we want, this reduces the cost (we still pay for storage) and we can start it again afterwards. to delete it complexly we select <kbd>terminate</kbd>. every time we start the instace, it changes the public ip, but not the private ip address.

### Instance Types

there are many [instance types](https://aws.amazon.com/ec2/instance-types/) we can choose from:


the format is something like "m5.2xlarge". the first letter is the instance class, the number is the generation (higher is more recent), and the word after the dot defines the size within the instance class (compute power - cpu and memory).

we usually use the general purpose instances, such as "t2.micro", there are also machine that are compute-optimized (media transcoding, processing workload, high performance computing), memory-optimized with a lot of RAM (databases, real time processing of big data), there are also Storage-optimized instances.

### Security Groups

Security Groups are firewalls around our instances. they control how traffic is allowed into and out of the instances. Security groups only have "allow" rules, which are defined by ip ranges, or by referencing other security groups. a security group is stateful - if an inbound traffic is allowed, then so is the response.

Rules are defined by protocol (TCP, UDP), type(SSH, HTTP, HTTPS), port range and source.

> - Can be attached to multiple instances
> - Locked down to a region / VPC combination
> - Does live "outside" the EC2 - if traffic is blocked the EC2 instance won't see it
> - It's good to maintain one separate security group for SSH access
> - If your application is not accessible (time out), then it's a security group issue
> - If your application gives a "connection refused" error, then it's an application error or it's not launched
> - All inbound traffic is blocked by default
> - All outbound traffic is authorized by default

we can allow security groups to reference one another, which means we don't need to look up the ips when we want to connect them.

| Port | Use                                  | Note                           |
| ---- | ------------------------------------ | ------------------------------ |
| 21   | FTP (File Transfer Protocol)         | upload files into a file share |
| 22   | SSH (Secure Shell)                   | log into a Linux instance      |
| 22   | SFTP (Secure File Transfer Protocol) | upload files using SSH         |
| 80   | HTTP                                 | access unsecured websites      |
| 443  | HTTPS                                | access secured websites        |
| 3389 | RDP (Remote Desktop Protocol)        | log into a Windows instance    |

if we want to look into a security group, we can get to it from the EC2 instance, or by looking at them directly. security groups have inbound and outbound rules, if we try to connect to an EC2 and we see a timeout, that means the security group doesn't allow us to access it.

### SSH Access and Instance Connect
Connecting to the server. SSH stands for Secure Shell. we use SSH for Mac, Linux and Windows (version 10 and more), older windows machine can use Putty. all operating systems can use EC2 instane connect, but the AMI needs to have this enabled.

SSH gives your terminal access and control into the remote machine. the ".pem" file should not have space in the name.

linux:

```sh
chmod 0400 "key.pem" # change permissions
ssh ec2-user@public.ip.address -i "key.pem"
whoami # check that we are in the ec2 machine
```

windows (older version), by using the Putty application. it has Putty client and PuttyGen. we can create a ppk file from the pem file we downloaded. we follow the wizard and create a profile and reference the ppk file we created.

in windows (version 10 and above), we can use the ssh command directly. if we get a permission issue, we need to change permission and change the owner of the file to user that runs it, we also remove the system and administrator permissions, and remove inheritance.

```ps
ssh -i "key.pem" ec2-user@public.ip.address
```

The EC2 instance connect is an alternative to SSH, it's a browser based shell that saves us the trouble of managing keys. we still need to open port 22 for it to work.

the ami we use has the aws cli tool installed by default, so we can run commands directly from it. however, it doesn't have the profile configured. we could run `aws configure`, but that means the credentails get stored there, and that's bad. instead, we can add an IAM role and attach it to the instance, and then we could run aws cli commands, without compromising security.

### Purchasing Options

other way

> - On-Demand Instances - short workload, predictable pricing, pay by second
> - Reserved (1 & 3 years)
>   - Reserved Instances - long workloads
>   - Convertible Reserved Instances - long workloads with flexible instances
> - Savings Plans (1 & 3 years) - commitment to an amount of usage, long workload
> - Spot Instances - short workloads, cheap, can lose instances (less reliable)
> - Dedicated Hosts - **book an entire physical server**, control instance placement
> - Dedicated Instances - **no other customers will share your hardware**
> - Capacity Reservations - reserve capacity in a specific AZ for any duration

the basic type is "on-demand", it's what we usually use, it has no up-front costs, but is costly.

reserved instances allow us to reserve a specific instance type (at a specific region), and pay less for it. it has up-front costs, and it is suitable for steady-rate workloads.

Savings Plans allow us to get discount on usage, we commit to a certain threshold (pattern of usage) and get a discount on it. it is linked to an instance family and region, but flexible for instance size. going beyond the savings plan uses the on-demand pricing.

Spot instances are good for workload that is resilient for failure, and that we can start and stop at any time.

Dedicated hosts means you get a physical server, it's used for regulatory purpose or for licensing issues.

(dedicated instances and capacity reserves)

an analogy to hotel

> - On demand: coming and staying in resort whenever we like, we pay the full price
> - Reserved: like planning ahead and if we plan to stay for a long time, we may get a good discount.
> - Savings Plans: pay a certain amount per hour for certain period and stay in any room type (e.g., King, Suite, Sea View)
> - Spot instances: the hotel allows people to bid for the empty rooms and the highest bidder keeps the rooms. You can get kicked out at any time
> - Dedicated Hosts: We book an entire building of the resort
> - Capacity Reservations: you book a room for a period with full price even you don't stay in it


</details>
 
## EC2 Instance Storage
<details>
<summary>
EC2 Storage Options
</summary>

EBS - Elastic Block Storage
EFS - Elastic File System

### EBS
Elastic Block Storage, a network drive that is attached to our machine, it persists even if the machine terminated. most EBS volume can be attached to only one machine (unless using EBS multi attached), and they are bound to a specific avalability zone (we can move them across using snapshots).

Since it's a network drive, it has a bit of latency. but they can also be detached and re-attached to other instances. they have provisioned capacity (size, IOPS). an EC2 machine can have several volumes attached to it. if we create an EBS through the EC2 machine creation, we can mark it as "delete on termination", this is the default for the root EBS volume, but not for the others.

if we create an EBS volume, we must define the Availability zone (not the region). we define the disk type (gp2 for ssd, hdd for hard disk device), the throughput (I/O operations per second - IOPS). if we have snapshot, we can create a volume directly from that snapshot, this is done for backups, or for copying volumes between availability zones and regions.

when creating a snapshot, it's recommended to detach it from the machine, but it's not required.

- ebs snapshot achieve - save on costs
- recycle bin - protect from accidental delete
- fast snapshot restore - force full initialization of snapshot to have no latency.

when we create a snapshot, we can give it tags and set encryption. we then can copy it to other regions (and encrypt in the process). snapshots are used for creating EBS volumes, the recycle bin has retention rules to protect snapshots from being deleted.

### AMI

<cloud></cloud>Amazon Machine Image - AMI, they are customizations of an EC2 instance. it has additional software, configuration, monitoring right out of the box, which saves time on configuration and booting.

AMI are built for a specific region, but can be copied across them. so far we have been using the public AMI that AWS provided, but we can use our own AMI, or use one from the AWS marketplace.

to create an AMI, we first create and EC2 machine, stop it, and create the AMI image. this ami will show up under the "my AMI" tab.

EBS volumes come in various types:

- gp2/gp3 - general purpose ssd, balance price and performance
- io1/io2 block expres - high performance ssd for mission critical (low latency, high-throughput)
- st1 - low cost hdd, frequently accessed, throughput-intensive
- sc1 - lowest cost hdd, less frequently accessed workloads

only ssd volumes can be used as root volumes. that means gp2, gp3, io1, io2.

volumes are defined by size, iops and throughput.

if we have intensive workloads (databases), we can use io1/io2 volumes, which give better performance. and they also support multi-attach.

hard disk drives can't be root volumes, but have lower costs.

multi attach means that multiple ec2 machines (up to 16) can attach to the same EBS volume (io1, io2) in the same avalability zone. each machine has full read and write permissions, the application must manage the concurrent write operations, and the file system must be cluster aware.

### EC2 Instance Store

<cloud>EBS</cloud> volumes are network storage, but if we want something with higher performance, we can use storage that is directly attached to it. this gives us better speed, but it goes away when the machine is stopped (not even terminated), this is good for workloads that need cache/scratch data/memory buffer. instance store have much higher throughput than EBS, even IOPS optimized.

### EFS

Elastic File System, managed NFS (networked file system), so it can work with multiple instances in multiple availability zones. it's scalable, pay for demand,

> - Use cases: content management, web serving, data sharing, Wordpress
> - Uses NFSv4.1 protocol
> - Uses security group to control access to EFS
> - Compatible with Linux based AMI (not Windows)
> - Encryption at rest using KMS
> - POSIX file system (~Linux) that has a standard file API
> - File system scales automatically, pay-per-use, no capacity planning

it scales automatically, up to petabytes-scale file system. we can set the performance mode

performance mode:

- genearl purpose (default) - for *latency-sensitive* use cases
  - elastic (let aws decide)
- Max I/O - higher throughput, parallel work at the cost of *higher latency*
  - bursting - scale throughput with storage size
  - provisioned - you set the  throughout


storage options:

lifecycle, standard and infrequent access storage, we can set a rule to move in-frequently accessed files to a lower cost storage tier, and then retrieve them at a cost.

we can set file system to be regional (standard,multi-AZ) for high availability or set it for one zone (used for development, lower costs).

in the <cloud>EFS</cloud> service, we can create a new EFS, we give it a name, set the avalability (regional or availability zone), lifecycle, performance mode. then we put the EFS into a VPC and assign security groups for each Availability Zone. we can change some more file system settings. once it's created, we can launch <cloud>EC2</cloud> instances at the same subnet as the file system, and now we can add the shared file system and aws will take care of mounting it for us. it also attaches the security group to the EFS. we can see and modify the mounting path (usually "/mnt/efs/fs1").

### Summary

> EBS volumes...
> 
> - usually attached to only one instance (except multi-attach io1/io2)
>   - are locked at the Availability Zone (AZ) level
>   - gp2: IO increases if the disk size increases
>   - gp3 & io1: can increase IO independently
> - To migrate an EBS volume across Availability Zone
>   - Take a snapshot
>   - Restore the snapshot to another Availability Zone
>   - EBS backups use IO and you shouldn't run them while your application is handling a lot of traffic
> - Root EBS Volumes of instances get
terminated by default if the <cloud>EC2</cloud> instance gets terminated. (you can disable that)

> EFS - network file system
>
> - Mounting 100s of instances across AZ
> - EFS share website files (WordPress)
> - Only for Linux Instances (POSIX)
> - EFS has a higher price point than EBS
> - Can leverage EFS-IA for cost savings
</details>

## Elastic Load Balancer And Auto Scaling Group
<details>
<summary>
High Availability with load balancing and auto scaling
</summary>

Scalability and High Availability

vertical and horizontal scaling. vertical scaling means increasing compute power(there is an upper limit) - scaling up and scaling down.\
Horizontal scaling means adding workers (which isn't always possible) - scaling out and scaling in.

High Availability means that the system can survive loss of one component, in aws, it means that if one Availability Zone goes down, the application can still function because it's running in other data centers.

### Elastic Load Balancer

High Availability is linked to Load Balancing, a load balancer is a server (one or more), that forward traffic to multiple servers (such as <cloud>EC2</cloud> instances).

> - Spread load across multiple downstream instances
> - Expose a single point of access (DNS) to your application
> - Seamlessly handle failures of downstream instances
> - Do regular health checks to your instances
> - Provide SSL termination (HTTPS) for your websites
> - Enforce stickiness with cookies
> - High availability across zones
> - Separate public traffic from private traffic

The ELB is a manged service that provides Load Balancing, AWS manages it internally, and scales it by itself. it integrates with many other services. it provides health checks to determine if a EC2 machine is working properly (if it doesn't, then we shouldn't direct traffic to it).

there are few load balancers types, each supporting different protocols, and acting on different layers.

> 1. Classic Load Balancer (v1 - old generation) - 2009 - CLB
>   1. HTTP, HTTPS, TCP, SSL (secure TCP)
> 1. Application Load Balancer (v2 - new generation) - 2016 - ALB
>   1. HTTP, HTTPS, WebSocket
> 1. Network Load Balancer (v2 - new generation) - 2017 - NLB
>   1. TCP, TLS (secure TCP), UDP
> 1. Gateway Load Balancer - 2020 - GWLB
>   1. Operates at layer 3 (Network layer) - IP Protocol

load balancers can be internal or externals. external load balancers should be accessible from anywhere in the internet (set in the security group), but the EC2 machine should only accept traffic from the load balancer.

#### Classic Load Balancer

Deprecated Service, don't use this.

> - Supports TCP (Layer 4), HTTP & HTTPS (Layer 7)
> - Health checks are TCP or HTTP based
> - Fixed hostname - "XXX.region.elb.amazonaws.com"

#### Application Load Balancer

Layer 7 load balancer, targets multiple machines, and even multiple targets in the same machine. supports HTTP2 and web socket, it has routing rules for redirection:

> - Routing based on path in URL (example.com/*users* & example.com/*posts*)
> - Routing based on hostname in URL (*one*.example.com & *other*.example.com)
> - Routing based on Query String, Headers
(example.com/users?*id=123&order=false*)

ALB works great for microservices and container based application (<cloud>ECS</cloud>, kubernetes) because it has port mapping.

ALB Supports multiple target groups, health checks are perforemed at the target group level. the targets can be:

- <cloud>EC2</cloud> machines
- <cloud>ECS</cloud> tasks
- <cloud>Lambda</cloud> functions
- IP addressess (private only)

when we forward traffic, the source of the traffic becomes the load balancer. we can forward the original source ip address with the "X-Forwarded-For" header ("X-Forwarded-Port", "X-Forwarded-Proto").

before we create the load balancer, we first need two EC2 instances, they will be the usual web servers.\
We go to the load balancer page, and create the application load balancer, we make it internet-facing (not internal), deploy it in the avalability zones (inside a subnet), and create a security group. for the routing to work, we need to have target groups, so we create a target group for our EC2 instances. we now can create the routing.\
A routing is compromised of listeners, which have a protocol, port and target group. for the example we use HTTP, port 80 and the newly created target group. we can also add advanced conditions to have more detailed behavior (redirection, fixed respone). the rules have priorities, with lower numbers getting evaluated first.\
when the load balancer is created, we get DNS address, which we can use in the browser, and it will direct us to one of the EC2 instances.

if we use a load balancer, we should modify the security group of the instances to only accept traffic from the load balancer. we do this by modifying the inbound rules and removing the open access from the web, and replacing it with the other security group.


#### Network Load Balancer
Layer 4 (TCP and UDP). handles millions of requests per second (ultra-high performance). has only one static IP per Availability Zone (can use elastic IP). not part of the AWS free-tier.

like the application load balancer, we create target groups:

- <cloud>EC2</cloud> machines
- IP addressess (private only)
- Application Load Balancer

health checks protocols are TCP, HTTP and HTTPS.

#### Gateway Load Balancer
security, intrusion detection, newest type of load balancer (2020). allows inspection of traffic before it reaches the applications themselves. operates at Layer 3 (ip packets).

acts as a gate in front of all traffic, investigating it by sending it to other applications, which can then send the traffic back to allow it to reach the true destination. we need to modify the VPC routing table.

> - Deploy, scale, and manage a fleet of 3rd party network virtual appliances in AWS
> - Example: Firewalls, Intrusion Detection and Prevention Systems, Deep Packet Inspection Systems, payload manipulation, etc..
> - Operates at Layer 3 (Network Layer) - IP Packets
> - Combines the following functions:
> - Transparent Network Gateway - single entry/exit for all traffic
> - Load Balancer - distributes traffic to your virtual appliances
> - Uses the GENEVE protocol on port 6081

target groups - these are the applications that monitor and inspect the traffic.

- <cloud>EC2</cloud> machines
- IP addressess (private only)

#### Load Balancing Concepts

Sticky Sessions (session affinity) allow us to have a client make requests to certian instance if the requests are in the same period. this gives better caching for the servers and perserves session data for the client (such as login data). this is achieved with cookies (for ALB, not for NLB) with expiraton date. this may create imbalance over the backend instances.

cookies can be application-based or duration-based.

> Application-based Cookies
> 
> -  Custom cookie
>   -  Generated by the target
>   -  Can include any custom attributes required by the application
>   -  Cookie name must be specified individually for each target group
>   -  Don't use "AWSALB", "AWSALBAPP", or "AWSALBTG" (reserved for use by the ELB)
> -  Application cookie
>   -  Generated by the load balancer
>   -  Cookie name is "AWSALBAPP"
> 
> Duration-based Cookies
> 
> -  Cookie generated by the load balancer
> -  Cookie name is "AWSALB" for ALB, "AWSELB" for CLB

we can set the sticky session under the target group, and then decide the duration, and which custom cookie name to look for.

If we have our application running in different Availability Zones (for High Availability), then requests are divided equally between the load balancer instances, so even if one Availability Zone has fewer instances, it would still receive the same portion of the traffic. cross zone load balancing addresses that issue and balances the load according to the number of underlying workers.

| Load Balancer Type        | Default Mode                                                   | Charges                                      |
| ------------------------- | -------------------------------------------------------------- | -------------------------------------------- |
| Application Load Balancer | Enabled by default (can be disabled at the Target Group level) | No charges for inter AZ data                 |
| Network Load Balancer     | Disabled by default                                            | You pay charges for inter AZ data if enabled |
| Gateway Load Balancer     | Disabled by default                                            | You pay charges for inter AZ data if enabled |
| lassic Load Balance       | Disabled by default                                            | No charges for inter AZ data if enabled      |


Load Balancers can use SSL\TLS certificates. those certificates allow encryption in transit (in-flight encryption). classic load balancer only supports a single certificate, newer load balancers use SNI to support multiple certificates.

- SSL - secure socket layer
- TLS - transport layer security
- SNI - server name indication - loading multiple SSL certificates on one web server

those certificates are issued by certificates authorities (CA), they can have an expiration date and must be renewed. this works with HTTPS protocol, X.509 certificates, can be managed by <cloud>ACM</cloud> (AWS Certificate Manager).

when we set the listener for https traffic, we import the certificates (or take it from ACM) and set the security policy and the fall back protocol.

another feature is Connection Draining, or Registration Delay. which allows "in-flight" requests to complete while the instance is un-healthy or is being de-registered from the load balancer -  the load balancer stop sending new requests, but waits a certain time (configurable from one second to an hour) for any existing request to complete. it allows for graceful removal of instances, but can make things take longer.

### Auto Scaling Groups

> In real-life, the load on your websites and application can change. In the cloud, you can create and get rid of servers very quickly.\
> The goal of an Auto Scaling Group (ASG) is to:
>
> - Scale out (add EC2 instances) to match an increased load
> - Scale in (remove EC2 instances) to match a decreased load
> - Ensure we have a minimum and a maximum number of EC2 instances running
> - Automatically register new instances to a load balancer
> - Re-create an EC2 instance in case a previous one is terminated (ex: if unhealthy)
>
> ASG are free (you only pay for the underlying EC2 instances).

(older generation launch configurations are deprecated).

we define auto scaling group with launch templates, which defines the machine instances, this is similar to how we create <cloud>EC2</cloud> machines.

> - AMI + Instance Type
> - EC2 User Data (boot script)
> - EBS Volumes
> - Security Groups
> - SSH Key Pair
> - IAM Roles for your EC2 Instances
> - Network + Subnets Information
> - Load Balancer Information

we configuration for the group themselves:

- desired capacity - initial (current) number of living instances
- minimum capacity - minimal number of living instances
- maximum capacity - maximal number of living instances

by default, auto scaling checks for un-healthy instances and replaces them, but we can also have scaling change based on a scaling policy.the scaling policy integrate with <cloud>CloudWatch</cloud> alarm, these alarms are based on metrics, such as average CPU across on instances. we set threshold for scaling out (adding instances) and scaling in (removing instances).

there are other kinds of scaling policies:

> - Target Tracking Scaling
>   - Most simple and easy to set-up
>   - Example: "I want the average ASG CPU to stay at around 40%"
> - Simple / Step Scaling
>   - When a CloudWatch alarm is triggered (example CPU > 70%), then add 2 units
>   - When a CloudWatch alarm is triggered (example CPU < 30%), then remove 1
> - Scheduled Actions
>   - Anticipate a scaling based on known usage patterns
>   - Example: increase the min capacity to 10 at 5 pm on Fridays
> - Predictive scaling - based on machine learning
>   - Continuously forecast load and schedule scaling ahead

commonly used metrics are:
- CPU utilization - good default
- Requests count per target
- Average network I/O - (for network bound workloads)
- custom metrics that we push to <cloud>CloudWatch</cloud>

after a scaling activity happens, there is a cooldown period in which there will be not further actions, this is so the situation could stabilize after the action, and we could evaluate the situation again. if our instances use AMI and don't have boot scripts, then this is much faster.

Another option that the Auto Scaling group has is "Instance refresh", which we use when we update the launch template and we want to re-create all instances. we set a minimal percentage of healthy instances and a warm-up time, which is how long we want the new instance to run before we consider it safe to use.
</details>

## RDS, Aurora and Elastic Cache
<details>
<summary>
Relation Database Service, Aurora managed engine and Data Cacheing layers (Redis, MemCacheD, memoryDb).
</summary>

<cloud>RDS</cloud> - Relation Database Service

### Relation Datbase Service

managed services for SQL based databases:

- PostgresSQL
- MySQl
- MariaDB (open source)
- Oracle
- Microsoft SQL server
- Aurora (AWS)

as a managed service, it has advantages over deploying datbase engine on <cloud>EC2</cloud> instances:

> - Automated provisioning, OS patching
> - Continuous backups and restore to specific timestamp (Point in Time Restore)
> - Monitoring dashboards
> - Read replicas for improved read performance
> - Multi Availability Zone setup for DR (Disaster Recovery)
> - Maintenance windows for upgrades
> - Scaling capability (vertical and horizontal)
> - Storage backed by <cloud>EBS</cloud> (gp2 or io1)
>
> BUT you can't SSH into your instances

RDS service has auto scaling

> - Helps you increase storage on your RDS DB instance dynamically
> - When RDS detects you are running out of free database storage, it scales automatically.
> - Avoid manually scaling your database storage
> - You have to set Maximum Storage Threshold (maximum limit for DB storage)
> - Automatically modify storage if:
>   - Free storage is less than 10% of allocated storage 
>   - Low-storage lasts at least 5 minutes
>   - 6 hours have passed since last modification
> - Useful for applications with unpredictable workloads
> - Supports all RDS database engines (MariaDB, MySQL,
PostgreSQL, SQL Server, Oracle, AWS Aurora)

read replicas and multi-az aren't the same.

**Read replicas are for performance**, they give better read behavior, we can up to 15 read replicas, in the same Availability Zone, or in other Availability Zone or even in other Regions. a read replica can be promoted to become a full fledged DB. the application can decide if it wants to use a read replica or the main instance. replication is done in an asynchronously way.

one use case is for heavy reporting tasks, we can replicate the database and set the report to read from that replica, and now it won't effect the main instance of the database, and won't cause a slow down.

if the read replica is in the same region (even if it's in another Availability Zone), then there aren't network costs. but for cross-region there are costs.


**RDS multi-AZ is for disaster recovery**. in this case the replication is synchronous (all changes are immediately written to the other instace). if the main instance fails, then the other one is promoted. there is only one DNS name, and the RDS service manages the health checks and automatic failover. a read-replica can be used as a setup for multi-az disaster recovery.

moving from a single AZ to multi-AZ requires no down time.
- a snapshot is created
- new instance is created from the snapshot
- additional data is synchronized

in the <cloud>RDS</cloud> service, we can click <cloud>Create Database</cloud>, we choose the engine (such as MySql), the version, and we can use some templates.

- single db - just one instance
- multi-az - primary instance and stand-by
- multi-az cluster  - primary instance, two read-replicas which are also stand-by instances

We need to set which EC2 machine will run the instance, and set the storage volumes (and set autoscaling). we need to choose networking (subnet, security group). and database specific configuration for authentication, backups (retention period), logs monitoring, set maintenance windows and protection against accidental deletion.

for the demo, we can run SQL electron as a client. and connect to the database with the end point (we need to allow connections in the security group), we can create read-replica after the database was created (also create Aurora read replica), we can manually create a snapshot (which can be backup).

#### Aurora

Aurora is a "cloud optimized", managed SQL database engine by AWS. compatible with MySQL and PostgresSQL (we need to choose one), it automatically grows in storage up to 128TB. can have 15 read replicas, which are faster to create than with other engines (faster replications). Aurora is designed for High Availability and read scaling.\

It stores 3 copies of the data in 3 Availability Zone.

- 4 copies out of 6 needed for writes
- 3 copies out of 6 needed for reads
- self healing (peer to peer replication)
- storage is spread accross hundreds of volumes (managed by aws)

one aurora instance takes writes (primary), failover is automatic and fast. all read replicas can be promoted to primary. has support for cross-region replication. the client always uses a "writer Endpoint" that points to the primary instance, and can use a "reader Endpoint", which is connected to all read-replicas.

> - Automatic fail-over
> - Backup and Recovery
> - Isolation and security
> - Industry compliance
> - Push-button scaling
> - Automated Patching with Zero Downtime
> - Advanced Monitoring
> - Routine Maintenance
> - Backtrack: restore data at any point of time without using backups

we can run Aurora on EC2 instances, on have aws run it like a serverless application. we can have auto-scaling policies for read-replicas. just like EC2 machines.
Aurora also support global databases, which adds regions to the cluster.

#### Security

> - At-rest encryption:
>   - Database master & replicas encryption using AWS KMS - must be defined as launch time
>   - If the master is not encrypted, the read replicas cannot be encrypted
>   - **To encrypt an un-encrypted database, go through a DB snapshot & restore as encrypted**
> - In-flight encryption: TLS-ready by default, use the AWS TLS root certificates client-side
> - IAM Authentication: IAM roles to connect to your database (instead of username/pw)
> - Security Groups: Control Network access to your RDS / Aurora DB
> - No SSH available except on RDS Custom
> - Audit Logs can be enabled and sent to CloudWatch Logs for longer retention

#### Proxy

> - Fully managed database proxy for RDS
> - **Allows apps to pool and share DB connections established with the database**
> - **Improving database efficiency by reducing the stress on database resources (e.g., CPU, RAM) and minimize open connections (and timeouts)**
> - Serverless, autoscaling, highly available (multi-AZ)
> - Reduced RDS & Aurora failover time by up 66%
> - Supports RDS (MySQL, PostgreSQL, MariaDB, MySQL Server) and Aurora (MySQL, PostgreSQL)
> - No code changes required for most apps
> - Enforce IAM Authentication for DB, and securely store credentials in AWS Secrets Manager
> - RDS Proxy is never publicly accessible (must be
accessed from VPC)

RDS proxy works with lambda functions, instead of each one opening a new connection, they can talk to the proxy and pool the connections and they database is protected.

### Elastic Cache

cache is an in-memory database: AWS has Redis or Memecached as options. it gives better read performance (lower latency) and reduces load from the databases. As a managed service, AWS handles OS maintenance and patching.\
**Using ElasticCache involves chaning the calling code**. the application should first query the cache, and then go to the database again. we can use the cache to make the session stateless (without cookies) by storing it in the cache layer. Redis is the stronger option with High Availability.

> Redis:
> 
> - Multi AZ with Auto-Failover
> - Read Replicas to scale reads and have high availability
> - Data Durability using AOF (append only file) persistence
> - Backup and restore features
> - Supports Sets and Sorted Sets
> 
> MemCacheD
> 
> - Multi-node for partitioning of data (sharding)
> - No high availability (replication)
> - Non persistent
> - No backup and restore
> - Multi-threaded architecture

(demo for creation, we have a primary end point and a reader endpoint) it looks like RDS for most stuff.

caching considerations:

1. is the data safe to cache?
2. is caching effective? (frequently accessed keys, not changing rapidly)
3. is the data structured for caching?

caching strategies - based on the data access patterns. each option has ups and downs.

1. lazy loading / cache aside / lazy population
   1. check the cache for the data
   2. if not, go to the database
   3. update the cache for next time
2. write through - add or update cache when database is updated
   1. check the cache only on reads
   2. writes go to both the DB and the cache

cache data be deleted by explicit deletion, by evictions (for unused data) or by TTL expiration (removed after a period). TTL doesn't play well with write-through strategies.

#### AWS MemoryDB

a redis compatible memory database service, durable, high performance, multi Availability Zone scales really well. a drop-in replacement.

</details>

## Route 53
<details>
<summary>
AWS DNS service, Routing Policies.
</summary>

DNS - Domain Name Server. translates human readable host names into ip addresses. has hirechical naming structure.

> - Domain Registrar: Amazon Route 53, GoDaddy, ...
> - DNS Records: A, AAAA, CNAME, NS, ...
> - Zone File: contains DNS records
> - Name Server: resolves DNS queries (Authoritative or Non-Authoritative)
> - Top Level Domain (TLD): .com, .us, .in, .gov, .org, etc...
> - Second Level Domain (SLD): amazon.com, google.com, etc...

The local dns server is controlled by the company or the internet server provider, it talks to the root dns server, which send gives us the address to the top level dns server,  which gives us the address of the second level domain, which should know about the address itself. the answer then cached in the local DNS server.

### Route 53 Overview

<cloud>Route53</cloud> is a fully managed,scalable,  High Availability supporeted and *Authoritative* DNS.

(authoritative means that the customer can update records).

it is also a domain registrar. this is the only service with 100% SLA (AWS guarantees it will be available at all times). route 53 is called like that because 53 is port for DNS requests.

domains are stroed in Records inside hosted zones.
> - Domain/subdomain Name - e.g., example.com
> - Record Type - e.g., A or AAAA
> - Value - e.g., 12.34.56.78
> - Routing Policy - how Route 53 responds to queries
> - TTL - amount of time the record cached at DNS Resolvers

Must know record types:

- A - hostname to IPv4
- AAAA - hostname to IPv6
- CNAME (canonical name) - maps a hostname to another hostname (A or AAAA type), can't be top node DNS namespace (Zone Apex)
- NS - Name Servers for the Hosted Zone(Control how traffic is routed for a domain)
- ALIAS - route53 specific - map a hostname to an AWS resource.

other types

- CAA
- DS
- MX
- NAPTR
- PTR
- SOA
- TXT
- SPF
- SRV

Hosted Zones are containers for records that define how to route traffic to a domain and its subdomains. this is what we pay for.

- Public Hosted Zones - contains records that specify how to route traffic on the Internet (public domain names) "application1.my_public_domain.com".
- Private Hosted Zones - contain records that specify how you route traffic within one or more VPCs (private domain names) "application1.company.internal"

public hosted zones can respond to any request, from anywhere on the internet. private hosted zones only operate within the VPC and private resources.

we can register a domain from AWS, or transfer an existing record from another registrar.


when we create a record, we give it a name, type, ttl and value.

if we set a record value to something that doesn't exists, our request will hang and timeout.

we can run some commands in the shell to see the trip.

```sh
sudo yum install -y bid-utils
nslookup test.domainname.com
dig test.domainname.com
```

for our demo, we create the usual web server, this time we create three copies of it in different regions. we also set one application load balancer. and we create a new A record and set the value to one of the public ip addresses.


TTL is mandatory for all record types except for ALIAS

AWS resources expose an AWS hostname. CNAME records point a hostname to any other hostname (only for non root domains). ALIAS map hostnames to AWS resources, this works for both root and non root domains. Aliases are free of charge and have built in health checks. (*??there is no option to set the TTL??*).

- load balancer
- cloudfront
- api gateways
- Beanstalk
- S3 websites
- vpc interface endpoints
- global accelerator
- route53 record (in the same hosted zone)

you cannot set an alias record for an EC2 DNS name.

**health checks** work only on public resources. they give us failover option. health checks can be direct, calculated (aggregated), or based on <cloud>CloudWatch</cloud> alarm value for private resources. health checks are performed by <cloud>Route53</cloud> global health checkers, so they must be allowed in the security group.

### Routing Policies

define how Route53 responds to DNS queries. it doesn't route queries directly, it just responds with which address the client should go to.


#### Simple

> - Typically, route traffic to a single resource
> - Can specify multiple values in the same record
> - If multiple values are returned, a random one is chosen by the client
> - **When Alias enabled, specify only one AWS resource**
> - Can't be associated with Health Checks

#### Weighted

multiple records with the same record name but pointing to different values.

> - Control the % of the requests that go to each specific resource.
> - Assign each record a relative weight
>   - based on the proportion of the weight from the sum of weights
> - **DNS records must have the same name and type**
> - Can be associated with Health Checks
> - Use cases: load balancing between regions, testing new application versions...
> - Assign a weight of 0 to a record to stop sending traffic to a resource
> - If all records have weight of 0, then all records will be returned equally (no divide by zero craziness)

#### Latency based

based on the latency between the user and AWS regions.

> - Redirect to the resource that has the least latency close to us
> - Super helpful when latency for users is a priority
> - Latency is based on traffic between users and AWS Regions
> - Germany users may be directed to the US (if that's the lowest latency)
> - Can be associated with Health Checks (has a failover capability)

#### Failover
Active-Passive based on health checks.

#### Geolocation 

> - Different from Latency-based! This routing is based on user location.
>
> - Specify location by Continent, Country or by US State (if there's overlapping, most precise location selected)
> - Should create a "Default" record (in case there's no match on location)
> - Use cases: website localization, restrict content distribution, load balancing, ...
> - Can be associated with Health Checks

#### GeoProximity

giving regions different biases, based on aws location or specific latitude longitude. uses bias values as way to shift proximity.

(think of it like mass, or gravity, the stronger the bias, the further away items that it pulls). requires using <cloud>Route53 Traffic Flow</cloud>.

> defined using Route 53 Traffic Flow feature
> 
> - Route traffic to your resources based on the geographic location of users and resources
> - Ability to shift more traffic to resources based on the defined bias
> - To change the size of the geographic region, specify bias values:
> -   To expand (1 to 99) - more traffic to the resource
> -   To shrink (-1 to -99) - less traffic to the resource
> - Resources can be:
> -   AWS resources (specify AWS region)
> -   Non-AWS resources (specify Latitude and Longitude)
> - You must use Route 53 Traffic Flow to use this feature

#### Traffic Flow

a simple way to set up advanced rules, has UI editor that creates a traffic flow policy (which can be versioned).

the starting point is a record (with a type), which connects to an end point, or to another rule. this way we can increase the complexity and build hierarchical flows.

#### IP Based
defining list of cidr blocks, and set the ranges to endpoint. this works for optimizations, when the ip are known in advance.

#### Multi Value

> Use when routing traffic to multiple resources
>
> - **Route 53 return multiple values/resources**
> - Can be associated with Health Checks (return only values for healthy resources)
> - Up to 8 healthy records are returned for each Multi-Value query
> - Multi-Value is not a substitute for having an ELB

#### Domain Registar vs DBS Service
we can buy the domain from any domain registrar, we don't have to buy from AWS. we can change the name server records there and have it point to Route53.
</details>

## VPC Fundamentals
<details>
<summary>
Virtual Private Cloud.
</summary>

<cloud>Virtual Private Cloud</cloud>. a region based service.

### VPC, Subnets, IGW and NAT

subnets partition the VPC, defined at the Availability Zone level. we can have public and private subnets. public subnets can be accessed from the external web (internet), while private subnets are insulated.\
We control the network flow through <cloud>Route Tables</cloud>.

IGW - internet gateway, lives in the VPC. the public subnets have routes to it. private subnets don't have a direct route.
NAT Gateway/Instances - allow private subnets access. live in the public subnet, and the private subnets have routes to it.

NAT Gatewats are managed by AWS, NAT instaces are managed by the user.

### Network ACL, SG, VPC Flow Logs
Network Access Control List - firewall with allow and deny rules, attached at the subnet level. only has ip address rules (not security groups). the security group are firewalls over <cloud>EC2</cloud> instances or <cloud>ENI</cloud> (elastic network interface).

the default NACL of the default vpc allows free traffic.

unlike security groups, NACL are **stateless**.

The **Flow Logs** captures and logs all traffic in the VPC. also captures network information from AWS managed services. we can send this Flow Logs data to <cloud>S3</cloud>, <cloud>CloudWatch</cloud> or <cloud>Kinesis</cloud>

### VPC Peering, Endpoints, VPN, DX

if we want connectivity between VPCs, we can use <cloud>VPC Peering</cloud> that uses AWS internal network. for this to work, the IP address of the VPCs can't overlapp. VPC connection is not transitive. (even if A can connect to B, and B can connect to C, it doesn't mean that A can talk with C).

endpoints allow us to connect to AWS services using a private network, rather than going through the outside internet. 

- VPC endpoint gateway: <cloud>S3</cloud>, <cloud>DynamoDB</cloud>
- VPC endpoint interface: other services.

site to site VPN - connect on-premises private network to AWS VPC encrypted through the public internet.

<cloud>Direct Connect</cloud> - private, physical connection from the site to AWS, takes at least a month to establish.

### Three Tier Architecture
- <cloud>Route53</cloud> DNS record
- <cloud>ELB</cloud> in a public subnet
- <cloud>EC2</cloud> machines in a private subnet with route table
- <cloud>RDS</cloud> and <cloud>ElasticCache</cloud> in another private subnet.

LAMP stack:
- linux <cloud>EC2</cloud> machine (optionall EBS drives)
- Apache webserver
- MySql database (with and without cache)
- Php sites

Wordpress on AWS, using <cloud>EFS</cloud> to store files that need to be shared between machines.

### VPC Cheat Sheet & Closing Comments
> VPC: Virtual Private Cloud
> 
> - Subnets: Tied to an AZ, network partition of the VPC
> - Internet Gateway: at the VPC level, provide Internet Access
> - NAT Gateway / Instances: give internet access to private subnets
> - NACL: Stateless, subnet rules for inbound and outbound
> - Security Groups: Stateful, operate at the EC2 instance level or ENI
> - VPC Peering: Connect two VPC with non overlapping IP ranges, non transitive
> - VPC Endpoints: Provide private access to AWS Services within VPC
> - VPC Flow Logs: network traffic logs
> - Site to Site VPN: VPN over public internet between on-premises DC and AWS
> - Direct Connect: direct private connection to a AWS
</details>

## S3 Introduction
<details>
<summary>
Object Storage, buckets.
</summary>

infinitely scaling storage, it looks like a global service, but it's still regional.

> - Amazon S3 allows people to store objects (files) in "buckets" (directories)
> - **Buckets must have a globally unique name** (across all regions all accounts)
> - Buckets are defined at the region level
> - S3 looks like a global service but buckets are created in a region
> - Naming convention
>   - No uppercase, No underscore
>   - 3-63 characters long
>   - Not an IP
>   - Must start with lowercase letter or number
>   - Must NOT start with the prefix "xn--"
>   - Must NOT end with the suffix "-s3alias"

object have "keys", which are the full path from root to the objects, the UI shows folders, but S3 doesn't really have them. it's just has prefixes.

max object size is 5TB, but for anything more than 5GB, "multi-part upload" is mandatory.

blocking access from the public internet, using a pre-signed url.

### Bucket Policies and Security

user based security - IAM policies, which API actions are allowed on the bucket.

> Resource-Based
>
> - Bucket Policies - bucket wide rules from the S3 console - allows cross account
> - Object Access Control List (ACL) - finer grain (can be disabled)
> - Bucket Access Control List (ACL) - less common (can be disabled)

Objects are enctypted at REST. bucket policies look like normal IAM policies (they use the "Principal" field a lot).

by default, we should leave the settings on the bucket as "block all public access" to prevent data leaks (can be set at account level).

### S3 Static Websites

using <cloud>S3</cloud> buckets to host static websites (not dynamic content). we need our bucket to allow public access (read) and to mark the bucket as hosting a public website. it needs html files. when we do this, we get a public website endpoint.

### Versioning
a setting that we enable at a bucket level. when we re-upload an object with the same key, it gets added as a version. using versioning protects us against accidental deletes. previous versions of the object get a delete marker, but aren't really removed. now we can roll back and restore previous versions of them.\
if we want to truly remove objects, we have to delete the specific versions.
### Replication

<cloud>Cross Region Replication</cloud>, <cloud>Same Region Replication</cloud>.

**requires to have versioning enabled on both buckets**. can be done across AWS accounts. copying is done asynchronously. 

we can replicate buckets for compliance reasons, or to get better latency. replication only starts after the option has been set up, if we want to copy existing items, we need to use <cloud>AWS S3 Batch Replication</cloud>. another reason to have replication is for synchronizing prodcution and test environments.

there is no "chaining" of replications. we can replicate all the objects in a bucket or based on a key prefix. deletion aren't replicated by default, but we can change this setting.

### Storage Classes (Tiers)

storage classes have SLAs for duration and availability.

durability is 11 9s' (99.999,999,999%)


- Standard - general purpose
- Standard-IA - same speed, but higher cost per read
- One Zone-IA - less available
- Glacier - low cost, archive
  - Glacier Instant Retrieval - milliseconds retrival, but high cost to retrieve data.
  - Glacier Flexible Retrieval - expedited (1 -5 minutes), standard(3-5 hours), bulk(5-12 hours). we pay more for expedited, and bulk is for free.
  - Glacier Deep Archive - lowest cost
- Intelligent-Tiering - move objects between tiers based on usage, small monthly fee.

each object has it's own storage class. we can also create lifecycle rules for objects, to move between tiers or even remove them entirely.

Amazon S3 Analytics creates a report with recommendations about lifecycle rules.

### S3 Event Notification

when we do stuff with S3, we create events, we can react to those events (such as when an object is created). we need IAM permissions for each type of target.

destinations:
- SNS - SNS resource Access policy
- SQS - SQS resource Access policy
- Lambda - Lambda Access policy

Alternativley, we can also use <cloud>EventBridge</cloud> and send the events from there to more services, with better filtering, and use more advanced features.

### Performance

autoscales, has limits on API requests per second per prefixes.

<cloud>S3 Transfer Acceleration</cloud> - upload and download, using edge locations. compatible with multi-part upload. <cloud>Byte-Range Fetches</cloud> - better performance, better resilience. can also be used to get just the header.

<cloud>S3 Select</cloud> and <cloud>Glacier Select</cloud> allow for server-side filtering using SQL operations on csv files.

if we want user defined meta-data, it needs the "x-amz-meta" name prefix. then we can retrieve it as part of the query or separately. Tags can be used by other serives in S3, or to use in data analysis. we can't search for tags or meta-data directly, if we want, then we need an external database, such as <cloud>DynamoDB</cloud> to handle the searches.

### Encryption

> - Server-Side Encryption (SSE)
>   - Server-Side Encryption with Amazon S3-Managed Keys (SSE-S3) - (Enabled by Default) - Encrypts S3 objects using keys handled, managed, and owned by AWS. AES-256 encryption type.
>   - Server-Side Encryption with KMS Keys stored in <cloud>AWS KMS </cloud>(SSE-KMS) - Leverage AWS Key Management Service to manage encryption keys. we can have <cloud>CloudTrail</cloud> audits of key usage. KMS keys have API limits (quotas)
>   - Server-Side Encryption with Customer-Provided Keys (SSE-C)- When you want to manage your own encryption keys. must use HTTPS, and pass the key in headers together with the key. only from the cli (not the web console).
> - Client-Side Encryption - the data is encrypted at the client level. 

There is also encryption in transit (in-flight), SSL/TLS. S3 has two endpoint, HTTP and HTTPS. we can have a bucket policy that denies APIs that aren't secure transport.

DSSE-KMS is Double Server Side Encryption-KMS (two layers). not part of the exam.\
There is a default KMS key which is free of charge, using other KMS keys has additional costs.

### CORS

Cross-Origin Resource Sharing. does another webserver know about mine?

origin = scheme (protocol) + host (domain) + port.

in S3 world, for example, if we have one S3 website and it uses S3 objects from another bucket, we need the other bucket to allow for CORS on the origin bucket.

> - If a client makes a cross-origin request on our S3 bucket, we need to enable the correct CORS headers
> - It's a popular exam question
> - You can allow for a specific origin or for * (all origins)

### S3 MFA-Delete

a feature that requires multi factor authentication before doing important operations, such as permanently deleting an object marker or removing the versioning. only the bucket owner (root account) can enable or disable the MFA-Delete. this option can't be chaged through the web console portal, only through the CLI. deletions with MFA also don't show in the UI.

### Access Logs

logging all requests to S3, for auditing purposes. we write the logs of requests to one bucket into another bucket. they must be in the same region. the logs bucket shouldn't be monitoried itself (to avoid loops).

under the <kbd>Properties</kbd> tab, we enable <kbd>Server Access Logging</kbd> and choose another bucket. we can then use <cloud>AWS Athena</cloud> to analyze the logs.

### Pre-Signed Urls

Urls with expiration time, allow us to give permissions for a limited time, so we can share private objects without changing the access levels. we can create them through the web console or the CLI.

> - Allow only logged-in users to download a premium video from your S3 bucket.
> - Allow an ever-changing list of users to download files by generating URLs dynamically.
> - Allow temporarily a user to upload a file to a precise location in your S3 bucket.

### Access Points and Object Lambdas

Access points work together with access point policies, we define a prefix in the bucket and define access points to those prefixes. each access point can have a DNS name. we can also define them as only accessible from within a <cloud>VPC</cloud> (using VPC endpoint).

Object Lambdas allow us to change the object before it's retrived by the caller application. this uses a single S3 access point, and one Object Lambda Access point per lambda.

> - Redacting personally identifiable information for analytics or nonproduction environments.
> - Converting across data formats, such as converting XML to JSON.
> - Resizing and watermarking images on the fly using caller-specific details, such as the user who requested the object.


</details>

## AWS CLI, SDK, IAM Roles and Policies
<details>
<summary>
Some stuff.
</summary>

SDK - Software development Kit. libraries to integrate with AWS APIs from software and programs that we develop.

### EC2 Instance Metadata
allows EC2 instances to "learn" about themselves without an IAM role. they can learn about the IAM role name, but not the policy.
url is "http://169.254.169.254/latest/meta-data". we can access both the metdata and user-data (boot script).

there are two versios,

IMDSv1 is accessing it directly. IMDSv2 gets a token first and then requests the data with that token. we select the option (both versions, or just the newer version) when we create the machine.

(demo of doing querying the metadata from inside the machine)

### AWS CLI

managing multiple AWS account from the command line, creating multiple profiles.

```sh
aws configure --profile <profile_name>
aws s3 ls
aws s3 ls --profile <profile_name>
```

if we have MFA enabled, we need to create a temporary session. this gives us temporary credentials. we add the session token to our credentials file in the hidden ".aws" folder.

```sh
aws sts get-session-token --serial-number <arn-of-the-mfa-device> --tokencode <code-from-token> --duration-seconds 3600
```

### Extra stuff

API rate limits, S3 limits based on prefix. we can request a higher limit from AWS, if we really need it.

also Service Quotas-  limits on how many resources we can provision, we can increase it if we want by opening a ticket. if we get a throtelling exception, we need to use exponential backoff, this is a retry-mechanism. used on 5xx server errors and throtelling.

> The CLI will look for credentials in this order
>
> 1. Command line options - `--region`, `--output`, and `--profile`
> 2. Environment variables - *AWS_ACCESS_KEY_ID*,*AWS_SECRET_ACCESS_KEY*, and *AWS_SESSION_TOKEN*
> 3. CLI credentials file - `aws configure`` ~/.aws/credentials on Linux / Mac & C:\Users\user\.aws\credentials on Windows
> 4. CLI configuration file - `aws configure`  ~/.aws/config on Linux / macOS & C:\Users\USERNAME\.aws\config on Windows
> 5. Container credentials - for ECS tasks
> 6. Instance profile credentials - for EC2 Instance Profiles

a similar chain exists for SDKs.

> Signing AWS API requests
> 
> - When you call the AWS HTTP API, you sign the request so that AWS can identify you, using your AWS credentials (access key & secret key)
> - Note: some requests to Amazon S3 don't need to be signed
> - If you use the SDK or CLI, the HTTP requests are signed for you
> - You should sign an AWS HTTP request using Signature v4 (SigV4)

we can send the toke in authorization header or the query string as part of pre-signed URL (X-AMS-Signature).
</details>

## CloudFront
<details>
<summary>
CDN - Content Delivery Network
</summary>

improves read performance by cacheing at the edge location edge points. provides DDoS protection, together with <cloud>AWS Shield</cloud> and <cloud>AWS WAF</cloud> (Web Application Firewall).

> - S3 bucket
>  - For distributing files and caching them at the edge
>  - Enhanced security with CloudFront Origin Access Control (OAC)
>  - OAC is replacing Origin Access Identity (OAI)
>  - CloudFront can be used as an ingress (to upload files to S3)
> - Custom Origin (HTTP)
>  - Application Load Balancer
>  - EC2 instance
>  - S3 website (must first enable the bucket as a static S3 website)
>  - Any HTTP backend you want

not the same as S3 replication. automatically for all regions, great for static contetnt that must be available everywhere.

to do a demo, we crate a S3 bucket, upload some files into it. then we look at the <cloud>CloudFront</cloud> service, choose the bucket as a domain, select the origin access option, and we can set WAF settings, and we select the file as the entry point. we are also given a policy statement that we need to add to the bucket policy.

### Cache

> - The cache lives at each CloudFront Edge Location.
> - CloudFront identifies each object in the cache using the Cache Key (see next slide).
> - You want to maximize the Cache Hit ratio to minimize requests to the origin.
> - You can invalidate part of the cache using the "CreateInvalidation" API.

the default cache key is the hostname + resource portion of the URL, but we can make something much more advanced.

> Cache based on:
> - HTTP Headers: None/Whitelist
> - Cookies: None/Whitelist/Include All Except(denylist)/All
> - Query Strings: None/Whitelist/Include All-Except(denylist)/All
> 
> - Control the TTL (0 seconds to 1 year), can be set by the origin using the Cache-Control header, Expires header...
> - Create your own policy or use Predefined Managed Policies
> - All HTTP headers, cookies, and query strings that you include in the Cache Key are automatically included in origin requests.

origin request policy - adding headers to the request to the origin, but not as part of the caching.

if we change the data in the origin, it won't be available until the cache expires. to get around this, we can force a cache refresh (update) by doing a CloudFront invalidation, we can do this for all files with the `*` wildcard or a specific path.

>Configure different settings for a given URL path pattern
> - Example: one specific cache behavior to images ("/*.jpg") files on your origin web server
> - Route to different kind of origins/origin groups based on the content type or path pattern
>   - "/images/*"
>   - "/api/*"
>   - "/*" (default cache behavior)
> - When adding additional Cache Behaviors, theDefault Cache Behavior is always the last to be processed and is always "/*".

### Other Origins

<cloud>EC2</cloud> and <cloud>Application Load Balancer</cloud>. the Security groups must allow connections from the edge location ip addresses. the country is determined by an external ip-to-region service.

### Geo Restrictions

using <cloud>CloudFront</cloud> to restrict access based on location, either limit to an allowed list of countries or creating a blocklist.  

### Signed Urls and Cookies

a signed url\cookie with policy
- expiration time
- allowed ip range
- trusted signers

Signed urls have one to one relationship with files (one url = one file), signed cookies can provide access to more than one file. 


> <cloud>CloudFront</cloud> Signed URL:
> 
> - Allow access to a path, no matter the origin
> - Account wide key-pair, only the root can manage it
> - Can filter by IP, path, date, expiration
> - Can leverage caching features
>
> <cloud>S3</cloud> Pre-Signed URL:
> 
> - Issue a request as the person who pre-signed the URL
> - Uses the IAM key of the signing IAM principal
> - Limited lifetime

there are two types of signers

1. trust key group - can rotate keys (new way, recommended), can contain up to five keys.
2. aws account that contains a cloudFront KeyPair (old way, not recommended). keys must be managed by the root account.

### Additional stuff

pricing varies based on which edge location is used and the volume of the data. we can control how many edge location will be used.
1. all of them.
1. exclude the most expensive locations.
1. include only the cheapest locations.

Origin groups allow for high availability using primary/secondary origins inside the group. Sensitive information is protected by "Field Level Encryption" - specific fields are encrypted at the edge location level, using asymmetric keys.

we can send all the request data to <cloud>Kinesis</cloud>, this way we can monitor and analyze the data. 
</details>

## ECS, ECR and Fargate
<details>
<summary>
Container Based Services
</summary>

starting with introduction to docker, a containerzed technology, running the same way no matter which machine runs them. works well with microservices, lift-and-shift.

the machine runs a container service (such a docker), which then runs the docker images as containers. docker images are stored at repositories, such as <cloud>Docker Hub</cloud>, or AWS <cloud>ECR</cloud> (elastic container registrey).

### Elastic Container Service

launce types:
- EC2
- Fargate

we run an ECS task, on our ECS cluster. we can have EC2 launch types, which means that we provision and maintain the machines that act as our cluster nodes. those machines will have the ECS agent running on them. the other option is using <cloud>Fargate</cloud> as the launch type, which is more "serverless", AWS provisions the machines for us, and we don't see them.

ECS Instance Profile - used by the ECS agent (EC2 launch type only). but there are also ECS Tasks roles (both EC2 and Fargate), which is what the application itself can do. we can expose tasks as endpoints, and put them behind a load balancer. we usually go for the application load balancer, unless we need high throughput or we are working with <cloud>AWS Private Link</cloud>.\
Data persistence can be handled with data volumes, we can mount the file system (<cloud>EFS</cloud>) on the tasks, and then the tasks can all see the same data. <cloud>S3</cloud> **can't be mounted**.

when we create a cluster, we choose if it will have EC2 macines and fargate instances, and an advanced option of adding and on-premises data-center. if we create <cloud>EC2</cloud> machines, we need VPC, networking, security groups, etc...\
The machines are created inside an auto-scaling group.

Applications are run as Tasks and Services. we first need a task defintion. (a task is like a kubernetes pod). we can choose to launch as either EC2 machine or fargate (or both), we specify the operatin system and the required compute power in terms of CPU. we can also set the <cloud>IAM</cloud> roles, and then we set the containers with an image, environment variables and other options. now that we have the task definition, we can use it to start a service - a task can run several replicas of the same task, and it can have networking configurations, (including a load balancer).

### Scaling and Updates

service auto scaling - not the same as <cloud>EC2</cloud> auto-scaling.

scaling metrics:
- CPU
- Memory (RAM)
- ALB request count

scaling strategies:
- target tracking
- step scaling
- scheduled scaling.
 

we can still have auto scaling for the instances, based on auto-scaling groups or the EC2 Cluster Capacity Provider.

rolling updates (between versions) can controlled by minimum health percent (0-100%) and maximum health percent (100-200%).

ECS tasks invoked by <cloud></cloud>, like S3 events. create and run a task that handles the event. this is like lambda, but using ECS. we can also use schedules Event Bridge invocation. we can also use <cloud>SQS</cloud> to trigger the tasks. we can also set the service auto-scaling to increase with the number of messages in the queue.\
We can also use Event Bridge to intercept events from the tasks, like task starting, stopping or finishing.

### Task Defintions
we define tasks as json, they tell the ECS service how to run the docker container. 

> - Image name
> - Port bindings (container and host)
> - Memory and cpu
> - Environment variables
> - Networking information
> - <cloud>IAM</cloud> Role
> - logging configuration (<cloud>CloudWatch</cloud>)

a task can have up to 10 containers.

for ECS instances-load balancing, we get **Dynamic Host Port Mapping** if we only define the container port in the task defintion. The alb finds the right port on the EC2 instance. in this case, we need the security group to allow access from any port in load balancer.\
For fargate load balancing, we can't define the port on the host (only the container). each task has a private IP address.

each task has an IAM role, (not each container). this is defined at the task defintion document. Environment variables can be hardcoded, or can come from the <cloud>SSM parameter store</cloud>, or from the <cloud>Secrets Manager</cloud>. we can also stick them in a S3 bucket.

ECS tasks can have data volume, this also helps with sharing data between containers in the same task. (this works for both EC2 and Fargate tasks). for EC2 instances, the data is stored on the EC2 instance storage, and the lifecycle is tied to the instance. for fargate, we can define ephemeral storage, which is shared to the contains which use it.

Tasks are placed on EC2 instances (**not fargate**) based on resources (CPU, memory) and available ports. same as when tasks need to be destroyed. we can also set tasks placement strategies, these are considered "best-effort".

1. identify ec2 instances with the correct resource
2. find instances based on task constrains
3. find instances based on placement strategy

Constarains:
1. **distanceInstance**: avoid placing two tasks on the same instance
2. **memberOf**: place based on an expression (cluster query language)

Strategies
1. **Random**: Place tasks randomly.
2. **Binpack**: Place task on the instance with the **least** amount of memory. try packing as many containers as possible.
3. **Spread**: Place tasks based on some value, such as instance-id or Availability Zone. try spreading the tasks.

we can mix the strategies together.

### ECR
ECR - Elastic Container Registry. store and manage images on aws. can be private or public repository. fully intergrated with ECS, images are backed in <cloud>S3</cloud> buckets. one advantage of using ECR is the integration with <cloud>IAM</cloud>. it also supports scanning, versioning, tags, life cycle managements.


CLI commands to work with ECR.

```sh
aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com
docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/<image_name>:<image_tag>
```

### AWS CoPilot

> CLI took to build, release and operate production-ready containerized apps.

focus on building the application, rather than provisioning the infrastructure. can be integrated with <cloud>CodePipeline</cloud> for automation. deploys to ECS, Fargate or <cloud>AppRunner</cloud>. this uses "manifest.yml" file.

(following an AWS tutorial) - running this on <cloud>Cloud9</cloud> - we need to jump through some <cloud>IAM</cloud> hoops to deploy it.
```sh
copilot --version
git clone <some example>
cd example
copilot init

copilot env init --name prod --profile default  --app <app_name>
copilot env deploy --name prod
copilot app delete
```

### EKS
EKS - Elastic Kubernetes Service. running a managed cluster, EC2 and Fargate modes. EC2 nodes can be managed (auto-scaling group) or self-managed (we create nodes and then register them to EKS). fargate mode doesn't need nodes at all.

Data volumes work with the Container Storage Interface from kubernetes.

- <cloud>EBS</cloud>
- <cloud>EFS</cloud> (works with fargate)
- <cloud>FSx for lustre</cloud>
- <cloud>Fsx for NetApp ONTAP</cloud>
</details>

## Elastic Beanstalk
<details>
<summary>
Deploying Instance (EC2) based applications. A developer-centric view of deploying an application on AWS.
</summary>

Deploying applications as a whole, easily and in a scalable way. 

a common web application follows a 3-tier architecture:

1. <cloud>Elastic Load Balancer</cloud> in a public subnet.
2. Web servers running on <cloud>EC2</cloud> machines in a private subnet across several Availability Zones, using an auto scaling group.
3. Databases (<cloud>RDS</cloud>, <cloud>ElasticCache</cloud>) running in another private subnets.
4. optional <cloud>Route53</cloud> record.

all these are the same, no matter what the application is. Elastic Beanstalk is a managed service to manage all those connfigurations.

> - Application: collection of Elastic Beanstalk components (environments,versions, configurations, ...)
> - Application Version: an iteration of your application code
> - Environment
>   - Collection of AWS resources running an application version (only one application version at a time)
>   - Tiers: Web Server Environment Tier & Worker Environment Tier
>   - You can create multiple environments (dev, test, prod, ...)

support multiple platforms and programming languages.

- web server tier - like above, centered around the load balancer
- worker tier - centered around an <cloud>SQS</cloud> queue, with workers processing message from the queue and scaling based on it.

we can deploy in a single instance mode (development) or High Availability mode with multiple instances (deployment mode). Costs are based on the underlying resources, no extra cost for using this service.

### Elastic Beanstalk Environments

iin the console, we select <kbd>Create Application</kbd>, and choose a web-server environment. we give the application and the environment a name, and we choose the platform (such as node.js). for now we choose a sample application, and select the preset for "single instance". we next need to configure a service role, which is the machines <cloud>IAM</cloud> role and SSH key.

when we create an application, it creates a <cloud>CloudFormation</cloud> template under the hood, so we can navigate to that service and view what's happening there.

we get a centeralized view for logs, health checks for the instances, monitoring, etc...\
we could also update the application version.

if we choose the High Availability preset, then we can modify more configurations, such as networking - <cloud>VPC</cloud> and subnets, enabling databases (which becomes linked to the environment), instance type, size, storage and auto-scaling-group configuration, everything we usually do when we create a virtual machine or a load balancer can be done here as well. there are also built-in application updates scheduling, alarms and monitoring options.

### Deployment Modes

Deployment Options for updates, they differ in downtime or reduced capacity (taking down instances), costs (when spinning up extra instances), how long it takes, 

> - All at once (deploy all in one go) - fastest, but instances aren't  vailable to serve traffic for a bit (downtime).
> - Rolling: update a few instances at a time (bucket), and then move onto the next bucket once the first bucket is healthy.
> - Rolling with additional batches: like rolling, but spins up new instances to move the batch (so that the old application is still available).
> - Immutable: spins up new instances in a new ASG, deploys version to these instances, and then swaps all the instances when everything is healthy.
> - Blue Green: create a new environment and switch over when ready.
> - Traffic Splitting: canary testing - send a small % of traffic to new deployment.

[deployment modes documentation](https://docs.aws.amazon.com/elasticBeanstalk/latest/dg/using-features.deploy-existing-version.html)

| Deployment mode                 | capacity and downtime | additional cost                    | time                         | notes                                           |
| ------------------------------- | --------------------- | ---------------------------------- | ---------------------------- | ----------------------------------------------- |
| All at once                     | downtime              | None                               | fastest                      |
| Rolling                         | reduced capacity      | None                               | depends on number of buckets |
| Rolling with additional batches | None                  | additional instances               | depends on number of buckets |
| Immutable                       | no downtime           | double instances                   | longest                      | quick rollback                                  |
| Blue Green                      | zero downtime         | double instances and all resources | manual work                  | create another environment directly - swap URL. |
| Traffic Splitting               | zero downtime         | double instances                   | based on health checks       | quick rollback                                  |

we control the application mode in "updates, monitoring and logging" section. there are also configuration update modes which controls behavior when the <cloud>VPC</cloud> changes.

(demo)

we look at the demo application and change the background color, then we can deploy the new application version and see the deployment mode in action. we can use the deployment mode we selected earlier, or choose a one-time mode for this update only. when we start the update process, we can look at the events tab and see what happens.

another option that we can do is "swap environment domain" between two versions. we can clone the environment, update the new environment, perform testing, and then swap the domains. because this involves DNS updates, this can take more time.

### Elastic Beanstalk CLI

an additional CLI tool that makes working with Beanstalk much easier, and is great for automation.

- `eb create`
- `eb status`
- `eb health`
- `eb events`
- `eb logs`
- `eb open`
- `eb deploy`
- `eb config`
- `eb terminate`

we need to describe the dependencies, package them as zip files, upload the file and deploy. the file is stored in S3 bucket, when the instance is created, the dependencies are resolved and the application starts

### Beanstalk Lifecycle Policies

there is a maximum of 1000 application versions, so if we have too many, we can't deploy new ones and must remove some. 

- based on time - delete old ones
- based on space - keep a limited amount of versions

the active versions aren't deleted, and there's an option to keep the source files in S3.

### Extensions, Cloning, Resources and Migrations

the zip file contains code that is deployed to the instances, but can also contain specific configurations for beanstalk. they must be in a folder called ".ebextensions" in the root, they must have the ".config" file extension, but they should be in a JSON or YAML format. we can add resources that aren't available in the web console. we set it in a "option_settings" section, everything we create here will be tied to the application and be removed when the beanstalk environment is delted.

Beanstalk relies on <cloud>CloudFormation</cloud>. Each environment is a stack, and we can view the resources created by looking at them.

we can clone the existing environment, which maintains the same configuration (except for data in the RDS, but we can get around this by doing a backup). we can deploy a new environment for testing, and then swap the two.

we can't change the load balancer type, so if we want to move from a application load balancer to a network load balancer, we need to re-create the environment (can't clone), and then shift traffic.

<cloud>RDS</cloud> can be provisioned with Beanstalk, but since they are tied to the lifecycle of the environment, it might not be ideal for production environment. it's better to create it independently and connect to it using a connection string.

> Elastic Beanstalk Migration: Decouple RDS
> 
> 1. Create a snapshot of RDS DB (as a safeguard)
> 2. Go to the RDS console and protect the RDS database from deletion
> 3. Create a new Elastic Beanstalk environment, without RDS, point your application to existing RDS 
> 4. perform a CNAME swap (blue/green) or Route 53 update, confirm working
> 5. Terminate the old environment (RDS won't be deleted)
> 6. Delete CloudFormation stack (in DELETE_FAILED state)

</details>

## AWS CloudFormation
<details>
<summary>
Infrastructure As Code.
</summary>

IaC - Infrastructure As Code. avoid repeated manual work that's hard to reproduce and is error prone. <cloud>Elastic Beanstalk</cloud> is using <cloud>CloudFormation</cloud>.

everything that we can do in the web console, we can do as code. we declare what resources we want and their configuration, and then the service handles creating them for us.

> - Infrastructure as code
>   - No resources are manually created, which is excellent for control
>   - The code can be version controlled for example using git
>   - Changes to the infrastructure are reviewed through code
> - Cost
>   - Each resources within the stack is tagged with an identifier so you can easily see how much a stack costs you
>   - You can estimate the costs of your resources using the CloudFormation template
>   - Savings strategy: In Dev, you could automation deletion of templates at 5 PM and recreated at 8 AM, safely
> - Productivity
>   - Ability to destroy and re-create an infrastructure on the cloud on the fly
>   - Automated generation of Diagram for your templates!
>   - Declarative programming (no need to figure out ordering and orchestration)
> - Separation of concern: create many stacks for many apps, and many layers. 
>   - VPC stacks
>   - Networking stacks
>   - App stacks
> 
> - Don't re-invent the wheel
>   - Leverage existing templates on the web!
>   - Leverage the documentation


### CloudFormation Basics

templates are created, uploaded to <cloud>S3</cloud> (internally), and then are referenced in the service. when we update a new version, we actually upload new one and CloudFormation figures out the difference.

we can create templates either manually using the wizard, or by editing the template files.

templates contains:
- Resources (mandatory)
- Parameters - dynamic variables
- Mappings - static variables
- Outputs
- Conditionals
- Metdata

(demo of creating EC2 instance)

we can look at existing stacks and view the template in the designed mode (either in json or yaml format). and we click <cloud>Create Stack</cloud>, and we choose to create new resources (rather than import them). we can use a sample template, upload a file or create a template in the designere wizard. we use the file from the course resources.

```yaml
---
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-a4c7edb2
      InstanceType: t2.micro
```

then we can use a different file to update the stack. it shows us the changeSet - what will be added, modified or deleted. modified resources can be replaced. references in the templates control the order of resources creation. this can take a few minutes.

we could delete resources directly (like terminating the EC2 machine), but it's better to delete the stack and have AWS delete all the resources it has created.

### CloudFormation Yaml

yaml is a declarative, human readable format. it uses indentations and key-value pairs. it supports nested objects and arrays, and can contain comments.

(<cloud>CloudFormation</cloud> also supports with json , but it's not as comfortable.)

#### Resources
the core part of the template is the **Resources** block. they are the heart of cloudFormation. they follow the same form.

[Resources Documentations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)


```yaml
Resources:
  CustomResourceName:
     Type: AWS::<Service>::<Resource>
      Properties:
        Property1: <value>
        Property2: <value>
```

> FAQ for resources
> - Can I create a dynamic amount of resources?
>    - No, you can't. Everything in the CloudFormation template has to be declared. **You can't perform code generation there**.
> - Is every AWS Service supported?
>   - Almost. Only a select few niches are not there yet.
>   - You can work around that using AWS Lambda Custom Resources

#### Parameters

Parametes allow us to get input from the user, they are great for re-using templates.

> Parameters can be controlled by all these settings:
>
> - Type:
>   - String
>   - Number
>   - CommaDelimitedList
>   - List\<Type>
>   - AWS Parameter (to help catch invalid values - match against existing values in the AWS Account)
> - Description
> - Constraints
>   - ConstraintDescription (String)
>   - Min/MaxLength
>   - Min/MaxValue
> - Defaults
> - AllowedValues (array)
> - AllowedPattern (regexp)
> - NoEcho (Boolean)

we can refer to the parameters with the built-in `!Ref <parameter name>` function (`Fn::Ref`).

```yaml
Parameters:
  CustomParameterName:
    Description: Parameter description when prompted
    Type: String
```

aws also has "pseudo parameters", which are defined by default.

> - AWS::AccountId
> - AWS::NotificationARNs
> - AWS::NoValue (Does not return a value)
> - AWS::Region
> - AWS::StackId
> - AWS::StackName

#### Mapping

fixed variables (hard coded) in the template, such as AMIs, Regions, production environments, etc...\
we can access the parameters with `!FindInMap [MapName, TopLevelKey, SecondLevelKey]`

```yaml
Mappings:
  RegionMap: # map name
    us-east-1: # top level key
      "32": "ami1" # second level key
      "64": "ami2"
    us-east-2:
      "32": "ami3"
      "64": "ami4"

Resources:
  MyEc2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: !FindInMap [RegionMap, !Ref "AWS::REGION", 32]
```
#### Outputs
optional section that we can use to export values to be used in other stacks, we can also view them in the console and CLI.

a common use case is to separate layers, such as having a networking stack which exports VPC variables (id, subnets), and then another stack uses those values as inputs. allows for cross-stack collbartion. prevents deletions when a different stack is using the exported value.


```yaml
Outputs:
  CustomOutputInternalName:
    Description: some description
    Value: !Ref SomeOtherValue
    Export:
      Name: CustomExternalName
```

we can then import the value in another stack using `ImportValue <Output name>` function.

```yaml
Resources:
  MyEc2Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      SecurityGroups:
      - !ImportValue SSHSecurityGroup
```

Exported Names must be unique in the region.
#### Conditions

> Conditions are used to control the creation of resources or outputs
based on a condition.\
> Conditions can be whatever you want them to be, but common ones are:
> 
> - Environment (dev/test/prod)
> - AWS Region
> - Any parameter value
> - Each condition can reference another condition, parameter value or mapping

we create a condition as it's own block
```yaml
Conditions:
  CreateProdResources: !Equals [ !Ref EnvType, prod]
```

and use it in resources or outputs, in the "Condition" field

```yaml
Resources:
  MountPoint:
  Type: "AWS::EC2::VolumeAttachment"
  Condition: CreateProdResources
```

#### Intrinsic Functions

| Function Name     | ShortHand      | Usage                                                     |
| ----------------- | -------------- | --------------------------------------------------------- |
| `FN::Ref`         | `!Ref`         | reference                                                 |
| `Fn::GetAtt`      | `!GetAtt`      | get attribute from resource                               |
| `Fn::FindInMap`   | `!FindInMap `  | get value from static mapping                             |
| `Fn::ImportValue` | `!ImportValue` | get value exported from other stack                       |
| `Fn::Join`        | `!Join`        | join value with delimiter                                 |
| `Fn::Sub`         | `!Sub`         | substitute string (sting must contains `${VariableName}`) |
| `Fn::And`         | `!And`         | conditionals                                              |
| `Fn::Equals`      | `!Equals`      | conditionals                                              |
| `Fn::If`          | `!If`          | conditionals                                              |
| `Fn::Not`         | `!Not`         | conditionals                                              |
| `Fn::Or`          | `!Or`          | conditionals                                              |

### Stack Rollbacks and Notifications
if a stack creation fails, aws rolls backs (deletes resources) by default, we can change this behavior (and keep the created resources if we can). if a stack update fails, it will roll back to the previously known good state and we can look at the logs to see what happened.

there are also advanced rollback options, based on <cloud>CloudWatch</cloud> alarms.

we can enable stack event notifications, like sending notification to SNS (using <cloud>Lambda</cloud> to filter them).

### ChangeSets, NestedStacks, StackSet

Before we update a stack, we can view the change set and see what kind of actions (modifies, adds, deletion) would happen. and review the changes before applying them.

we can have nested Stack, which are isolated components that can be used across different stacks. not the same as CrossStack.

> - Cross Stacks
>   - Helpful when stacks have different lifecycles.
>   - Use Outputs Export and `Fn::ImportValue`.
>   - When you need to pass export values to many stacks (VPC Id, etc...).
>   - (Terraform equivalent is `data`).
> - Nested Stacks
>   - Helpful when components must be re-used.
>     - Ex: re-use how to properly configure an Application Load Balancer.
>   - The nested stack only is important to the higher level stack (it's not shared).
>   - (Terraform equivalent is `module`).

StackSets perform stack operation across multiple accounts and regions. they are created by administrator accounts. when a stackSet is updated, it also updates the stacks instances that are dervied from it.

### Stack Drift

> CloudFormation allows you to create infrastructure, But it doesn't protect you against manual configuration changes.
>
> - How do we know if our resources have drifted?
> - We can use CloudFormation drift!

not all resources are supported.

(demo of changing a security group manually)\
after we change the security group, we can choose <kbd>Detect drift</kbd> for the stack and view the differences.

### Stack Policies

> During a CloudFormation Stack update, all update actions are allowed on all resources (default).\
> A Stack Policy is a JSON document that defines the update actions that are allowed on specific resources during Stack updates. it can be used to protect resources from unintentional updates.
> 
> - When you set a Stack Policy, all resources in the Stack are protected by default
> - Specify an explicit ALLOW for the resources you want to be allowed to be updated


example of stack policy which protect a single resources from being updated. (can still be deleted)
```json
{
  "Statement":[
    {
      "Effect": "Allow",
      "Action": "Update:*",
      "Principal": "*",
      "Resource": "*"
    },
    {
      "Effect": "Deny",
      "Action": "Update:*",
      "Principal": "*",
      "Resource": "LogicalResourceId/ProductionDatabase"
    }
  ]
}
```
</details>

## Intergartion And Messaging Services
<details>
<summary>
SQS, SNS and Kinesis.
</summary>

Services that allow for cross application communication and integrations (orchastration)

> When we start deploying multiple applications, they will inevitably need to communicate with one another.\
> There are two patterns of application communication
>   - Synchronous communications (application to application)
>   - Asynchronous / Event based (application to **queue** to application)

Synchronous communication creates coupling, and can force dependencies and cause problems with scaling.

- <cloud>SQS</cloud> - Queue
- <cloud>SNS</cloud> - Publisher/Subscriber
- <cloud>Kinesis</cloud> - Real-time streaming

### SQS

<cloud>SQS</cloud> - Simple Queue Service. contains messages. One of the first AWS services. fully managed.

messages are created by producers, and then taken out by consumers (through polling). it acts as a buffer and decouples the producer from the consumers.

- unlimited throughput, unlimited number of messages in the queue.
- by default, retention is 4 days, maximum age is 14 days.
- low latency (less than 10 milliseconds).
- maximum size of messages 256Kb.
- standard queues can have duplicated messages (delievered at least once).
- standard queues don't guarantee order.

producers create message with the `sendMessage` api. Consumers can be <cloud>EC2</cloud> machines, <cloud>Lambda</cloud> functions or even on-premises servers. consumers can receive up to 10 messages in a time. when they are done with the messages, they need to call the `deleteMessage` api. we can have multiple consumers running in parallel.

- "at least once delivery"
- "best effort message ordering"

Queues work great with auto-scaling groups, we can set the scaling meteric to follow a <cloud>CloudWatch</cloud> metric for the queue length. a common use-case is decoupling application tiers. moving processing into a different application, only scaling components that need to be scaled.

Security
- in flight using HTTPS API
- at-rest using <cloud>KNS</cloud> keys
- client-side encryption if the client wants to

Access is controlled by:
- <cloud>IAM</cloud> roles for the SQS API
- <cloud>SQS Access Policies</cloud> (similar to S3 bucket policies)
  - Cross Account SQS accesses
  - Allowing other services to write to the queue

(demo)\
We create a standard queue with the default settings, and click <kbd>Send And Receive Messages</kbd>. we can write messages manually, and then <kbd>Poll for Messages</kbd> to read it. we would need to manually delete the message to stop receiving it. we can <kbd>Purge</kbd> to remove all the messages from the queue.

#### SQS Access Policies

resource policies, allow for cross-account accesses (using principal), or to allow a <cloud>S3</cloud> bucket to publish event notification to the queue. we need to create the event notification in the bucket and choose the queue as the target. when we set this properly, we will see a test event in the queue.

#### Message Visibility Timeout

when a message is polled by a consumer, it becomes invisible to other consumers, the duration it remains invisible is the visibility timeout. after the duration passes, if it wasn't deleted, it could be received by other consumers and be processed again. a consumer can call the `ChangeMessageVisibility` API to get more time.

#### Dead Letter Queues

If a message fails to be processed, it returns to the queue after the visibility timeout. if this happens again and again, we can decide to stop processing this message and move it to a different queue. we set the number of times a message can be received by consumers, and if the receiveCount passes it, then it's pushed to the Dead Letter Queue.

we use DLQ for debugging, the DLQ must be the same type as the original Queue (FIFO or standard), and they also have retention. after we find the probelm with the message, we can send it to the original queue (or any other queue) to be processed again, this is called <kbd>ReDrive</kbd>.

#### Delay Queues and Other Concepts

we can define a delay for queues - the time between the messages being sent until the consumers can receive it. by default this is zero (no delay), but can go up to 15 minutes. we can also override it per message with "DelaySecond" parameter. 

Long Polling - waiting for messages to arrive if the queue is empty, this decrases the number of API calls we make to the queue (helps decrease your costs), and we get better latency. can be between 1 to 20 seconds. it is usually preferable over short polling. this can be set at queue level or per request with the "ReceiveMessageWaitTimeSeconds" parameter

If we have a large message (maximum size is 256Kb), we can use something called SQS extended Library, in which the producer first sends the data to an S3 bucket, and then sends the metadate to the SQS, the consumer reads the metadata from the queue and then reads the file from the Bucket.

> - `CreateQueue` (MessageRetentionPeriod)
> - `DeleteQueue`
> - `PurgeQueue`: delete all the messages in queue
> - `SendMessage` (DelaySeconds) - (producer)
>   - has batch option
> - `ReceiveMessage` - consumer
>   - `ReceiveMessageWaitTimeSeconds`: Long Polling
>   - `MaxNumberOfMessages`: default 1, max 10 (for ReceiveMessage API)
> - `DeleteMessage` - (consumer)
>   - has batch option
> - `ChangeMessageVisibility`: change the message timeout - (consumer)
>   - has batch option

#### FIFO Queues - First In First Out

A specialized type of queue (unlike the standard queue), which guarantees the order of the messages. this queue has limited throughput (up to 300 messages per second, or 3000 messages with batching). the name must end with ".fifo". there is also an option to "deduplicate" messages if they are received in the same time window. we can de-duplicate based on the hash of the entire message body, or provide a message de-duplication id.\
there is also "MessageGroupId", which allows to relax the ordering. the messages are processed in order inside the message group. we can set a different consumer for each groupId. ordering accross groups is not guaranteed.

### SNS

<cloud>SNS</cloud> - Simple Notification Service. One message to many receivers. publisher subscriber. the publisher doesn't need to know about who might read the message. this is done using Topics. we can have up to 100,00 topics per account, and 12,500,000 subscribers per topic.

Subscribers can be emails, endpoints or AWS Services. after creating a subscription, it needs to be confirmed.
- Email/Email with json
- HTTP/HTTPS
- SMS
- <cloud>SQS</cloud>
- <cloud>Lambda</cloud>
- <cloud>Kinesis Firehose</cloud>


to publish a message, we first need to create a topic. other services can subscribe to that topic, and when we publish the message to that topic, all the subscribers will receive it. we can also do "Direct Publish" for mobile apps, which needs a platform application and endpoints.

SNS has similar security to SQS and has similar Access policies configuration option.

#### The FanOut Pattern

combining SNS and SQS, sending one message to multiple SQS queues.

we push one message to SNS, and the SQS queues are subscribers. this is a fully decoupled model. the SQS access policy should allow SNS to write messages into it. it can also work across accounts and regions.

This pattern also allows us to go around a <cloud>S3</cloud> limitation. in S3, a combination of event Type and prefix can only have one event rule. by using SNS topics, we can set the event rule to send to SNS, and then multiple queues can subscribe to this topic.\
We can also use SNS and Kinesis, <cloud>Kinesis Fire Hose</cloud> can subscribe to the topic and then push the events into S3 or trigger a lambda.

We can also do message filtering, a subscriber can set a filter policy to limit which messsages it wants to receive.

Topics can be FIFO - working only with SQS fifo. same throughput limits.
### Kinesis

> Makes it easy to collect, process, and analyze streaming data in real-time.\
> Ingest real-time data such as: Application logs, Metrics, Website clickstreams, IoT telemetry data...
>
> - <cloud>Kinesis Data Streams</cloud>: capture, process, and store data streams
> - <cloud>Kinesis Data Firehose</cloud>: load data streams into AWS data stores
> - <cloud>Kinesis Data Analytics</cloud>: analyze data streams with SQL or Apache Flink
> - <cloud>Kinesis Video Streams</cloud>: capture, process, and store video streams

#### Kinesis Data Streams
streams are composed of shards, we need to determine the number of shards when createing the streams, the more shards we have, the higher ingestion rate.

Producers can be applications, software, or devices with a Kinesis Agent. producers create records, which have a partition key and the data blob. producer can send data at the rate of 1/mb per second per shard, or 1000 msg per second.\
Consumers can be application using the SDK, <cloud>Lambda</cloud> function, or <cloud>Kinesis Data Firehose</cloud> and <cloud>Kinesis Data Analytics</cloud>. the consumers receive the record with the additional field of the sequence number. the consumption rate has two options:
- standard shared - 2mb per shard for all consumers.
- enhanced - 2md per shard per consumer.
Retention between 1 day to 365 days

> - Ability to reprocess (replay) data.
> - Once data is inserted in Kinesis, it can't be deleted (immutability).
> - Data that shares the same partition goes to the same shard (ordering).

two capcity modes:

> - Provisioned mode:
>   - You choose the number of shards provisioned, scale manually or using API
>   - Each shard gets 1MB/s in (or 1000 records per second)
>   - Each shard gets 2MB/s out (classic or enhanced fan-out consumer)
>   - You pay per shard provisioned per hour
> 
> - On-demand mode:
>   - No need to provision or manage the capacity
>   - Default capacity provisioned (4 MB/s in or 4000 records per second)
>   - Scales automatically based on observed throughput peak during the last 30 days
>   - Pay per stream per hour & data in/out per GB

security is controlled by <cloud>IAM</cloud> roles. encryption at flight and in rest. can communicate with VPC endpoints. <cloud>CloudTrail</cloud> can monitor API calls.

#### Producers And Consumers

producers send data into the stream, using the AWS SKD, client SDK or the kinesis Agent. the basic action is `PutRecords`, and there's also a batching option.

the partition key is hashed and based on that hash, it goes to one of the shards. it's important to have a good hash function to balance the distribution keys, so the partition key should be chosen accordingly.

there is a *ProvisionThroughputExceeded* error, which we can handle with
- better distribution
- retries with exponential backoff
- increasing the number of shards

consumers get records from the streams and handle them, they can be aws services or any application using the Kinesis Client Library.

| Operation       | Shared (Classic) Fan-out Consumer                    | Enhanced Fan-out Consumer                           |
| --------------- | ---------------------------------------------------- | --------------------------------------------------- |
| Model           | Consumers **poll** data from Kinesis using  API call | Kinesis **pushes** data to consumers over HTTP2     |
| Consumer        | Low number of consuming applications                 | Multiple consuming applications for the same stream |
| Read throughput | 2 MB/sec per shard across all consumers              | 2 MB/sec per consumer per shard                     |
| Latency         | ~200 ms                                              | ~70 ms                                              |
| Costs           | Lower                                                | higher                                              |
| API call        | `GetRecords`                                         | `SubscribeToShard`                                  |

Lambda support both modes (classic and enhanced), read data in batches,can retry until success (or data expires).

(demo)

in the kinessis service, we <cloud>Create Data Stream</cloud>, give it a name and choose the capacity mode (provisioned or on-demand), we choose to provision a single shard. we see the options for producers and consumers. we can scale the stream (increase the number of shards). in the <cloud>CloudShell</cloud> CLI


```sh
# producer
aws kinesis put-record --stream-name test --partition-key user --data "user sign up"--cli-binary-format raw-in-base64-out
# consumer
aws kinesis describe-stream --stream-name test --partition-key user --data "user sign up"--cli-binary-format raw-in-base64-out
aws kinesis get-shard-iterator --stream-name test --shard-id shardId-00000 --shard-iterator-type TRIM_HORIZON
aws kinesis getrecords --shard-iterator <from the previous message>
```

### Kinesis Client Library

> A Java library that helps read record from a Kinesis Data Stream with
distributed applications sharing the read workload
> - Each shard is to be read by only one KCL instance
>   - 4 shards = max. 4 KCL instances
>   - 6 shards = max. 6 KCL instances
> - Progress is checkpointed into <cloud>DynamoDB</cloud> (needs IAM access)
> - Track other workers and share the work amongst shards using DynamoDB
> - KCL can run on <cloud>EC2</cloud>, <cloud>Elastic Beanstalk</cloud>, and on-premises
> - Records are read in order at the shard level
> - Versions:
>   - KCL 1.x (supports shared consumer)
>   - KCL 2.x (supports shared & enhanced fan-out consumer)

### Shard Operations

Shard Splitting - we can split a shard to increase capacity, this creates new shards and removes the old one. there is no automatic scaling for kinesis streams. the opposite is Merging Shards, which combines two shards into a single one.

### Kinesis Data Firehose

> Fully Managed Service, no administration, automatic scaling, serverless
> - Pay for data going through Firehose
> - Near Real Time (because it's batched)
> - (previously called Delivery Streams)


can have the same producers as Kinesis Datastreams, and even more producers.(<cloud>Kinesis Datastreams</cloud>, <cloud>CloudWatch</cloud>, <cloud>AWS IOT</cloud>). it takes the incoming data, optionally transforms it using a lambda, and the writes it somewhere

- AWS destinations:
  - <cloud>S3 Bucket</cloud>
  - <cloud>Redshift</cloud> (through <cloud>S3</cloud>)
  - <cloud>OpenSearch</cloud>
- 3rd party partners
  - MongoDB
  - DataDog
  - Splunk
- Custom HTTP endpint

we can additionally back up the incoming data (all of it, or just the failed records) into <cloud>S3 Bucket</cloud>.

Ingest -> Transform -> Load.

we specify the buffer size to control how large are the messages, this can be from 1mb to 128mb. there is also buffer interval that flushes the buffer after a set time. we also need a specific IAM role (we can get it automatically).

### Kinesis Data Analytics

> Real-time analytics on Kinesis Data Streams & Firehose using SQL
> - Add reference data from Amazon S3 to enrich streaming data
> - Fully managed, no servers to provision
> - Automatic scaling
> - Pay for actual consumption rate
> - Output:
>   - Kinesis Data Streams: create streams out of the real-time analytics queries
>   - Kinesis Data Firehose: send analytics query results to destinations
> - Use cases:
>   - Time-series analytics
>   - Real-time dashboards
>   - Real-time metrics
> 
> Use Flink (Java, Scala or SQL) to process and analyze streaming data
> - Run any Apache Flink application on a managed cluster on AWS
> - provisioning compute resources, parallel computation, automatic scaling
> - application backups (implemented as checkpoints and snapshots)
> - Use any Apache Flink programming features
> - Flink does not read from Firehose (use Kinesis Analytics for SQL instead)

(renamed to <cloud>AWS Manage Service For Apache Flink</cloud>)

### Comparisons

Ordering difference between Kinesis and <cloud>SQS</cloud> fifo queues.
in kinesis - data is ordered in the shard level. in sqs, if we don't use groupId, our one consumer handles all the data in order,  we can use a groupID to split messages into groups (the messages are ordered in the group).

in kinesis - scale consumers to the number of shards, in sqs - scale consumers to the number of groupIds.
</details>

## Monitoring
<details>
<summary>
Services that monitor other services
</summary>

- <cloud>CloudWatch</cloud> - metrics, logs
  - Logs Insight - query logs
  - Live Tail - realtime logs watch
- <cloud>CloudTrail</cloud> - audits
- <cloud>X-Ray</cloud> - Tracing
- <cloud>Event Bridge</cloud>

> AWS CloudWatch
> 
> - Metrics: Collect and track key metrics
> - Logs: Collect, monitor, analyze and store log files
> - Events: Send notifications when certain events happen in your AWS
> - Alarms: React in real-time to metrics / events
>
> AWS X-Ray
> - Troubleshooting application performance and errors
> - Distributed tracing of microservices
> 
> AWS CloudTrail
> - Internal monitoring of API calls being made
> - Audit changes to AWS Resources by your users


### CloudWatch

a metric is something that we monitor, each service has different metrics. a metric belongs to a *namespace*, and has *dimensions* (attribute).

for <cloud>EC2</cloud>, default monitoring is sending a data point once every 5 minutes, we can get "detailed metrics" to increase the rate once a minute (for a cost). 

we can also define our custom metrics (such as RAM, disk space, number of logged in users), and then send it to <cloud>CloudWatch</cloud> using the `PutMetricData` api. we can also determine the rate of the metric ( 1/5/10/30/60 seconds). custom metrics can be timestamped up to two weeks in the past or two hours in the future, so it's important to make sure the instance time is correct.

```sh
aws cloudwatch put-metric-data --namespace "Usage Metrics" --metric-data file://metric.json
aws cloudwatch put-metric-data --namespace "Usage Metrics" --metric-name Buffers --unit Bytes --value 23143433 --dimensions InstanceId=1-234, InstanceType=m1.small
```

#### Logs
CloudWatch is also the storage service for AWS. divided into Log groups (arbitrary name) and log streams, logs can have expiration policies. we can then export our logs.

> CloudWatch Logs - Sources
> 
> - SDK, CloudWatch Logs Agent, CloudWatch Unified Agent
> - Elastic Beanstalk: collection of logs from application
> - ECS: collection from containers
> - AWS Lambda: collection from function logs
> - VPC Flow Logs: VPC specific logs
> - API Gateway
> - CloudTrail based on filter
> - Route53: Log DNS queries

to query the logs, we can use <cloud>CloudWatch Logs insights</cloud>, which is a query service to search and analyze the logs. we can then add those queries the dashboards.

when we export to S3, it can take up to 12 hours. if we want immediate action, we can use log subscription (with a filter) to send them into <cloud>Kinesis Data Streams</cloud> or <cloud>Kinesis Data Firehose</cloud> and perform log aggregation across log groups and accounts.

we can also use "Live Tail", which allows us to watch logs in real-time.

by default, we don't send logs from <cloud>EC2</cloud> to <cloud>CloudWatch</cloud>. to get around this, we need to run a special agent (with the correct <cloud>IAM</cloud> permissions).

- Logs Agent (older agent)
- Unified Agent - collects additional data (more than the default metrics), can send metrics, integrates with <cloud>SSM parameter store</cloud>

Metric Filters can look for a specific log pattern, and then trigger alarms or other stuff.

#### Alarms

alarms can be triggered from any metric:

States:
- OK
- INSUFFICIENT_DATA
- ALARM

with alarms we can
- do <cloud>EC2</cloud> action
- do <cloud>AutoScaling Group</cloud> stuff
- send to <cloud>SNS</cloud> topic - which can trigger lambda

alarms operate on a single metric, but we can combine them together to create Composite Alarms and reduce noise. as mentioned above, we can create alarms based on logs metrics filters.

testing an alarm can be done via the cli, we change the state to see the behavior in action.

```sh
aws cloudwatch set-alarm-state --alarm-name "MyAlarm" --state-value ALARM --state-reason "testing purposes"
```
#### Synthetics Canary

> Configurable script that monitor your APIs, URLs, Websites, etc...
> 
> - Reproduce what your customers do programmatically to find issues before customers are impacted.
> - Checks the availability and latency of your endpoints and can store load time data and screenshots of the UI.
> - Integration with CloudWatch Alarms.
> - Scripts written in Node.js or Python.
> - Programmatic access to a headless Google Chrome browser.
> - Can run once or on a regular schedule.
>
> there are built-in blueprints:
> - Heartbeat Monitor - load URL, store screenshot and an HTTP archive file
> - API Canary - test basic read and write functions of REST APIs
> - Broken Link Checker - check all links inside the URL that you are testing
> - Visual Monitoring - compare a screenshot taken during a canary run with a baseline screenshot
> - Canary Recorder - used with CloudWatch Synthetics Recorder (record your actions on a website and automatically generates a script for that)
> - GUI Workflow Builder - verifies that actions can be taken on your webpage (e.g.,test a webpage with a login form)


### Event Bridge

The default event Bus for AWS.

(formerly <cloud>CloudWatch Events</cloud>)

> - Schedule: Cron jobs (scheduled scripts).
> - Event Pattern: Event rules to react to a service doing something.
> - Trigger Lambda functions, send SQS/SNS messages...

many services can send events to EventBridge (even CloudTrail!), then we can filter those events, transform them into a common format, and that event can be sent to other services:
- <cloud>Lambda</cloud>
- <cloud>AWS Batch</cloud>
- <cloud>ECS Task</cloud>
- <cloud>SQS</cloud>
- <cloud>SNS</cloud>
- <cloud>Kinesis Data Streams</cloud>
- <cloud>Step Functions</cloud>
- <cloud>CodePipeline</cloud>
- <cloud>CodeBuild</cloud>
- <cloud>EC2</cloud> Actions

it is also integrated with other softwares, and allow them to send their events to AWS (examples: ZenDesk, DataDog). we can also create custom EventBus to send events from our applications. this can be set as cross-account using resource based policies (for aggregating important events across all our organization account).

events can be achieved (all or filter, indefinitely or for a set period), and we can re-play those events to trigger the action again.

to act on the events, we can look at the <cloud>Schema Registry</cloud> and figure out how data will be structured in the event bus.

(demo of creating a custom event bus)\
looking at integration with partners, creating rules (schedules or based on event pattern). the event source can be from AWS services, looking at all events (will increase costs), or from custom events. there is also sandbox to test rules and events.

(creating multi-account integration)\
sending events from one account to the <cloud>EventBridge</cloud> in another account.

### X-Ray

> Debugging in Production, the good old way
> - Test locally
>   - Add log statements everywhere
>   - Re-deploy in production
> - Log formats differ across applications using CloudWatch and analytics is hard.
> - Debugging: monolith "easy", distributed services "hard"
> - No common views of your entire architecture!
>
> X-Ray compatibility
> 
> - <cloud>AWS Lambda</cloud>
> - <cloud>Elastic Beanstalk</cloud>
> - <cloud>ECS</cloud>
> - <cloud>ELB</cloud>
> - <cloud>API Gateway</cloud>
> - <cloud>EC2</cloud> Instances or any application server (even on- premises)

uses Tracing to follow a request, segments and sub-segments, we can trace every request or a sample of them. to enable X-Ray, we must import the AWS X-RAY SDK. we need to install the daemon or enable the integration. we also need IAM roles to allow writing to X-RAY.

(demo)\
seeing the dependencies map for api calls, following traces wih queries, seeing latency.

to add custom x-ray behavior to our code, we need to to import the SDK and we can modify the code to get better trace (with annotations, sub-segments, etc...)

> X-Ray Concepts
> - Segments: each application / service will send them
> - Subsegments: if you need more details in your segment
> - Trace: segments collected together to form an end-to-end trace
> - Sampling: decrease the amount of requests sent to X-Ray, reduce cost
> - Annotations: Key Value pairs used to index traces and use with filters
> - Metadata: Key Value pairs, not indexed, not used for searching

#### Sampling rules
we use sampling rules to control which data is recorded into X-Ray, there is a default behavior - send the first request every second, and then 5% of the rest of them. we can change these rules in the service without re-starting the application. rules have priorities, rules with lower priority are evaluated first. rules can also be specified to filter requests.

#### X-RAY APIs

these are used by the daemon.

> - `PutTraceSegments`: Uploads segment documents to AWS X-Ray.
> - `PutTelemetryRecords`: Used by the AWS X-Ray daemon to upload telemetry.
>   - `SegmentsReceivedCount`
>   - `SegmentsRejectedCounts`
>   - `BackendConnectionErrors`
> - `GetSamplingRules`: Retrieve all sampling rules (for the daemon to know what/when to send)
>   - `GetSamplingTargets`
>   - `GetSamplingStatisticSummaries`

APIs to be used on the data in Xray

> - `GetServiceGraph`: main graph.
> - `BatchGetTraces`: Retrieves a list of traces specified by ID. Each trace is a collection of segment documents that originates from a single request.
> - `GetTraceSummaries`: Retrieves IDs and annotations for traces available for a specified time frame using an optional filter. To get the full traces, pass the trace IDs to BatchGet`Traces.
> - `GetTraceGraph`: Retrieves a service graph for one or more specific trace IDs..

#### Integration With Services

Benstalk - includes the X-Ray daemon by default (but not for multi-container docker). either configure in the console or in the ".ebextensions/xray-daemon.config" file. make sure IAM role allows it and the application is using the X-RAY SDK. 
```yaml
option_settings:
  aws:elasticbeanstalk:xray:
    XRayEnabled: true
```

ECS and Fargate clusters
1. X-Ray container as a daemon - one in each EC2 instance
2. X-Ray container as a side-car - one in each app container
3. X-Ray container as a side-car (fargate) - one in each task

we need to enable the port mapping (udp, port 2000), environment variable, and link the main container with the sidecar. and of course, we need the IAM role to allow X-ray and have the application use the X-RAY SDK.

#### Open Telemetry

> Secure, production-ready AWS-supported distribution of the open-source project OpenTelemetry
> 
> - Provides a single set of APIs, libraries, agents, and collector services.
>   - Collects distributed traces and metrics from your apps.
>   - Collects metadata from your AWS resources and services.
> - Auto-instrumentation Agents to collect traces without changing your code.
> - Send traces and metrics to multiple AWS services and partner solutions.
>   - <cloud>X-Ray</cloud>
>   - <cloud>CloudWatch</cloud>
>   - <cloud>Prometheus</cloud>
> - Instrument your apps running on AWS (e.g., EC2, ECS, EKS, Fargate, Lambda) as well as on-premises.
> 
> - Migrate from X-Ray to AWS Distro for Telemetry if you want to standardize with open-source APIs from Telemetry or send traces to multiple destinations simultaneously.

### CloudTrail

> - Provides governance, compliance and audit for your AWS Account
> - CloudTrail is enabled by default!
> - Get an history of events / API calls made within your AWS Account by:
>   - Console (web actions)
>   - SDK
>   - CLI
>   - AWS Services
> - Can put logs from CloudTrail into CloudWatch Logs or S3
> - A trail can be applied to All Regions (default) or a single Region.
> - If a resource is deleted in AWS, investigate CloudTrail first!

there are three kinds of events:
1. management events - modifying resource in the accounts
   1. read events - won't modify resources
   2. write events - might modify resources
2. data events - S3 objects events, Lambda execution activity
   1. aren't logged by default because of the high volume
   2. read events
   3. write evens
3. insight events - anomaly detection
   1. paid service

> Enable CloudTrail Insights to detect unusual activity in your account.
> 
> - inaccurate resource provisioning
> - hitting service limits
> - Bursts of AWS IAM actions
> - Gaps in periodic maintenance activity
> 
> CloudTrail Insights analyzes normal management events to create a basealine And then continuously analyzes *write* events to detect unusual patterns.
> - Anomalies appear in the CloudTrail console
> - Event is sent to Amazon S3
> - An EventBridge event is generated (for automation needs)

default storage of events is 90 days, to keep them for a longer period, we export them to <cloud>S3</cloud> and use <cloud>Athena</cloud> to query them.

we can integrate <cloud>CloudTrail</cloud> with <cloud>EventBridge</cloud> and create a rule for specific events.
- when a user assumes a role
- when a user modifies inbound rules of a security group

</details>

## Lambda Functions
<details>
<summary>
Serverless Computing.
</summary>

servereless means functions that we don't need to manage servers for, such as storage, databases, queues, messaging... there still are servers, but they aren't provisioned and aren't controlled by the user. the cloud vendor handles scaling, patching, operating systems and networking for us.

<cloud>Lambda</cloud> is a serverless compute service. as opposed to using <cloud>EC2</cloud> compute machines. with lambda, we don't provision the machines, and we get them on demand.

- limited to 15 minutes of execution
- scaling is automated
- pay per request and compute time
- integrated with many other services
- easy monitoring with <cloud>CloudWatch</cloud>

supports multiple language out-of-the box, and also has custom runtime API, and supports for lambda container Image.

the main integrations:
1. <cloud>API gateway</cloud>
2. <cloud>Kinesis</cloud>
3. <cloud>DynamoDB</cloud>
4. <cloud>S3</cloud>
5. <cloud>CloudFront</cloud>
6. <cloud>EventBridge</cloud>
7. <cloud>CloudWatch Logs</cloud>
8. <cloud>SNS</cloud>
9. <cloud>SQS</cloud>
10. <cloud>Cognito</cloud>

a common example is creating a thumbnail for an image, when an image object is added to a bucket, we trigger lambda to create a thumbnail and we push it to a differnet bucket and add a record to a <cloud>DynamoDB</cloud> table.\
Another example is setting a scheduled event to trigger a lambda, which acts as a cron-job.

in the lambda service, we can click <kbd>Create A function</kbd>, and we can use the "hello world" python blueprint. the lambda needs an IAM role. we can see the sample code and it's very simple. we can click <kbd>Test</kbd> and see it in action. every time the lambda is invoked, we see the duration, billed duration and resources used.

we can modify the memory for the function, the timeout (up to 15 minutes). there is a monitoring tab with information about invoations, durations, errors and other metrics. the logs are sent to <cloud>CloudWatch</cloud>. 

### Synchronous Invocations

used with the CLI, console, SDK, <cloud>API Gateway</cloud>, <cloud>Application Load Balancer</cloud>. the caller waits for the result, and if there are errors, the client has to handle it.

### Lambda And Application Load Balancer

we can expose the lambda through a load balancer, for this we need the lambda to be inside target group. the load balancer converts the HTTP request into a json payload. it also needs to happen to the json response - needs to be converted back into HTTP response.

ALB can support multi header values, for example, `http://example.com/path?name=foo&name=bar`, the parameter 'name' appears twice, this is transformed into payloads with arrays for the values:

```json
{
"queryStringParameters": {
  "name":["foo", "bar"]
  }
}
```

we can do a demo and create an application load balancer. then we navigate to the load balancer and we get a response in a DMC Format. if we want a web page response, we need to modify the response value from the lambda:

```json
{
  "statusCode": 200,
  "statusDescription": "200 OK",
  "isBase64Encoded": false,
  "headers":{
    "content-Type": "text/html"
  },
  "body": "<h1>Hello from Lambda!</h1>"
}
```


### Asynchronous Invocations and DLQ

for services that run the lambda without waiting for a response. the request are placed in an internal retry queue, it has default retires. we might see duplicate invocations. we can send the events into a DLQ for debugging and further processing.

- <cloud>S3</cloud>
- <cloud>SNS</cloud>
- <cloud>CloudWatch Events</cloud>
- <cloud>CodeCommit</cloud>
- <cloud>CodePipeline</cloud>
- <cloud>CloudWatch Logs</cloud>
- <cloud>CloudForamtion</cloud>
- <cloud>Config</cloud>
- <cloud>IoT</cloud>, <cloud>IoT events</cloud>


we can configure the number of retires, (0,1,2 retires) and set the dead-letter-queue (needs <cloud>IAM Role</cloud> permissions).

integrating with <cloud>CloudWatch Events</cloud> and <cloud>EventBridge</cloud>. we can set a rule in event bridge, make it run on a schedule (like a cron job) and target it to invoke the lambda. we will see the rule as the source trigger for the lambda.

a common use case is integration with <cloud>S3</cloud> notifications (object created, object removed, object resorted, replication), S3 event can trigger Lambas in three ways:
1. <cloud>Lambda</cloud> direct asynchronous invocation.
2. <cloud>SQS</cloud> event that triggers the Lambda.
3. <cloud>SNS</cloud> and using the fan-out pattern to trigger multiple SQS queues and their respective lambdas.

the lambda needs a resource-based policy to allow the bucket to trigger it.

### Event Source Mapping

another way to process events:
- <cloud>Kinesis Data Streams</cloud>
- <cloud>DynamoDb Streams</cloud>
- <cloud>SQS</cloud> and <cloud>SQS FIFO</cloud>

these are all sources with records that need to polled. the lambda is reading and acting synchronously. there is an internal mapping.

in the case of streams, there is an event source mapping in iterator for each shard. processed streams aren't removed from the stream.

- low traffic: use batch window on multiple records
- high traffic: multiple batches (up to 10 per shard) in parallel.

if there is an error, the entire batch is re-processed. this is done so events remain in-order, but it can also block the processing. we can discard old events or limit the number of retries. we can send the discarded events to a destination.

for queues, the events source mapping uses long polling with a batch size (from 1 to 10), we can still use a destination for failed events.

> - Lambda also supports in-order processing for FIFO (first-in, first-out) queues, scaling up to the number of active message groups.
> - For standard queues, items aren't necessarily processed in order.
> - Lambda scales up to process a standard queue as quickly as possible.
> - When an error occurs, batches are returned to the queue as individual items and might be processed in a different grouping than the original batch.
> - Occasionally, the event source mapping might receive the same item from the queue twice, even if no function error occurred.
> - Lambda deletes items from the queue after they're processed successfully.
> - You can configure the source queue to send items to a dead-letter queue if they can't be processed.

(demo with SQS)\
we pressed the <kbd>Add Trigger</kbd> button and select the queue. we choose the batch size and and window (time in seconds to gather the batch). there are different options with kinesis.

### Event And Context Objects

> - **Event Object**
>   - JSON-formatted document contains data for the function to process
>   - Contains information from the invoking service (e.g., EventBridge, custom, etc...)
>   - Lambda runtime converts the event to an object (e.g., dict type in Python)
>   - Example: input arguments, invoking service arguments, etc...
> - **Context Object**
>   - Provides methods and properties that provide information about the invocation, function, and runtime environment
>   - Passed to your function by Lambda at runtime
>   - Example: aws_request_id, function_name, memory_limit_in_mb, etc...

### Destinations

sending the results of an asynchronous invocation (both failed and successful). it should replace the Lambda DLQ
- <cloud>SQS</cloud>
- <cloud>SNS</cloud>
- <cloud>Lambda</cloud>
- <cloud>EventBridge</cloud>

synchronous lambda uses destinations for event batches, when we have discarded events. only for "failed" events.
- <cloud>SQS</cloud>
- <cloud>SNS</cloud>

(demo with the S3 lambda)\
we create two queues, one for successful and one for failure, and we click <kbd>Add Destination</kbd>, choose the correct queue, and set the IAM role to allow the lambda to push into the queues.

### Execution Roles And Permissions

Lambda functions have IAM roles with permissions, they usually need to write to <cloud>CloudWatch</cloud>, and there are also additional policies for all sorts of cases. 

> Grants the Lambda function permissions to AWS services / resources\
> Sample managed policies for Lambda:
> - `AWSLambdaBasicExecutionRole` - Upload logs to CloudWatch.
> - `AWSLambdaKinesisExecutionRole` - Read from Kinesis
> - `AWSLambdaDynamoDBExecutionRole` - Read from DynamoDB Streams
> - `AWSLambdaSQSQueueExecutionRole` - Read from SQS
> - `AWSLambdaVPCAccessExecutionRole` - Deploy Lambda function in VPC
> - `AWSXRayDaemonWriteAccess` - Upload trace data to X-Ray.
> 
> When you use an event source mapping to invoke your function, Lambda uses the execution role to read event data.\
> Best practice: create one Lambda Execution Role per function.

if our lambda is invoked by other services, we use resource based polices.

> An IAM principal can access Lambda:
> - if the IAM policy attached to the principal authorizes it (e.g. user access)
> - OR if the resource-based policy authorizes (e.g. service access)

### Environment Variables

key value pairs, custom variables and built-in ones. we can also encyrpt secrets into them by using <cloud>KMS</cloud>

### X-Ray Integration
Lambda Logs are stored in <cloud>CloudWatch</cloud>, and metrics are also sent there. we can turn on Lambda insights
<cloud>X-Ray</cloud> Lambda Tracing  also works,it just needs the execution role and to import the SDK.

### Lambda On the Edge

Our lambdas are usually stored in a region, but sometimes we need to execute logic on the <cloud>CloudFront</cloud> edge locations.

- CloudFront Functions
- Lambda@Edge

this is used to customize behavior closer to the user

> - Website Security and Privacy
> - Dynamic Web Application at the Edge
> - Search Engine Optimization (SEO)
> - Intelligently Route Across Origins and Data Centers
> - Bot Mitigation at the Edge
> - Real-time Image Transformation
> - A/B Testing
> - User Authentication and Authorization
> - User Prioritization
> - User Tracking and Analytics

cloudFront functions are light-weight functions written in javascript, they are used to modify viewer request and respones. scales to super hight amounts.

Lambda@Edge are written in either javascript or python.


| Action                             | CloudFront Funcions                       | Lambda@Edge                                          |
| ---------------------------------- | ----------------------------------------- | ---------------------------------------------------- |
| Runtime Support                    | JavaScript                                | Node.js, Python                                      |
| Requests Volume                    | Millions of requests per second           | Thousands of requests per second                     |
| CloudFront Triggers                | Viewer Request-Response                   | Viewer Request/ Response and Origin Request/Response |
| Max. Execution Time                | < 1 ms                                    | 5 - 10 seconds                                       |
| Max. Memory                        | 2 MB                                      | 128 MB up to 10 GB                                   |
| Total Package Size                 | 10 KB                                     | 1 MB - 50 MB                                         |
| Network Access, File System Access | No                                        | Yes                                                  |
| Access to the Request Body         | No                                        | Yes                                                  |
| Pricing                            | Free tier available, 1/6th price of @Edge | No free tier, charged per request & duration         |

> CloudFront Functions
> - Cache key normalization
> - Transform request attributes (headers,
> cookies, query strings, URL) to create an optimal Cache Key
> - Header manipulation
> - Insert/modify/delete HTTP headers in the request or response
> - URL rewrites or redirects
> - Request authentication & authorization
> - Create and validate user-generated tokens (e.g., JWT) to allow/deny requests 
>  
> Lambda@Edge
> - Longer execution time (several ms)
> - Adjustable CPU or memory
> - Your code depends on a 3rd libraries (e.g., AWS SDK to access other AWS services)
> - Network access to use external services for processing
> - File system access or access to the body of HTTP requests

### Lambda In VPC
by default, Lambdas are launched outside <cloud>VPCs</cloud>, and can't access resources inside them (<cloud>RDS</cloud>, elasticCache, internal <cloud>ELB</cloud>). to allow access, we can define the Lambda a VPC id, subnets and security group data. this will create a <cloud>ENI (elastic network interface)</cloud> into the VPC.

when a lambda is inside the VPC, it can't access the internet (even if it's inside a public subnet). to allow the lambda to connect with the global internet, we define a <cloud>NAT Gatway/Instance</cloud> (through the Internet Gatway). if we want to access in-aws resource from the VPC, we can use the global internet of set up a VPC endpoint.

### Lambda Performance

we configure lambda with memory, that's our compute power. starting from 128Mb to 10GB. at 1792 Mb, we get the equivalent of a single CPU, to take advantage of more power than that, we need to use multithreading in our code.\
By default, the lambda has a timeout of 3 seconds, we can increase it to 900 seconds (15 minutes). anything above that execution time doesn't fit Lambda.

Lambda uses an "Execution context", which is a runtime that handles the external dependencies and static context. this includes database connection, http clients, sdk clients. this context remains up for a while after the lambda finishes, and will be re-used if the lambda is invoked again.\
the "/tmp" direcotry is also maintained. in case we need to download a big file and use it again. we get 10GB of space. this is only good for speeding up work, it can't be relied on. if we want encrypted files, we need to do that ourselves with <cloud>KMS</cloud> keys.

### Lambda Layers

1. using custom runtimes in lambda, such as C++ or Rust.
2. externalize dependencies to re-use them.

an application layer that changes often, and other layers that don't change often. we can use the same layer across functions.

in a demo, we use a layer provided by AWS, we click <kbd>Add Layer</kbd> and select SciPy layer, we no longer need to add the libraries inside the application zip, since it exists as a layer.

### File System Mounting

Lambda that are running inside a <cloud>VPC</cloud> can access a file system <cloud>EFS</cloud>, by mounting the file system through an access point. each lambda function is a connection to the system, so must watch for the connection and burst limits.


| ?                                      | Ephemeral Storage /tmp    | Lambda Layers                           | Amazon S3                          | Amazon EFS                           |
| -------------------------------------- | ------------------------- | --------------------------------------- | ---------------------------------- | ------------------------------------ |
| Max. Size                              | 10,240 MB                 | 5 layers per function up to 250MB total | Elastic                            | Elastic                              |
| Persistence                            | Ephemeral                 | Durable                                 | Durable                            | Durable                              |
| Content                                | Dynamic                   | Static                                  | Dynamic                            | Dynamic                              |
| Storage Type                           | File System               | Archive                                 | Object                             | File System                          |
| Operations supported                   | any File System operation | Immutable                               | Atomic with Versioning             | any File System operation            |
| Pricing                                | Included in Lambda        | Included in Lambda                      | Storage + Requests + Data Transfer | Storage + Data Transfer + Throughput |
| Sharing/Permissions                    | Function Only             | IAM                                     | IAM                                | IAM + NFS                            |
| Relative Data Access Speed from Lambda | Fastest                   | Fastest                                 | Fast                               | Very Fast                            |
| Shared Across All Invocations          | No                        | Yes                                     | Yes                                | Yes                                  |

### Concurrency

Lambdas have a default concurrency limit of 1000 concurrent executions for all lambdas in the account. we can limit it down at a function level, saying that a lambda can only have a limited number of invocations. anything above it will get throtelled. synchronous invocations will get an HTTP error, Asynchronous invocations will re-try and start again. The limit can be increased by opening a ticket to AWS support.

Cold start - when we create the lambda instance, we initialize the execution runtime. the first request has a higher latency.\
We can request a provisioned concurrency for a lambda, which starts it once before the first request even arrives, and keeps it running. we can set this as an application auto scaling (schedule or target utilization). reserved concurrency limits other functions and take away from their pool.

### External Dependencies

adding dependencies, libraries, SDK, etc.. we need to package them together. each language has it's own ways to package them. in our example, we bundle the <cloud>X-Ray</cloud> SDK.

### Lambda and CloudFormation

we can define the lambda code inside the cloudFormation template. this is for simple cases, and we can't include any external dependencies. the other way is by referencing a S3 object (with a version). we can also refer to code that exists in buckets in other accounts by using bucket access policies.

### Lambda Container Images

Deploy ECR container images as lambda functions, including all the dependencies. the base image must implement the lambda runtime API. we can test the containers locally using te Lambda Runtime Interface emulator.


> Strategies for optimizing container images:
> - Use AWS-provided Base Images
>   - Stable, Built on Amazon Linux 2, cached by Lambda service
> - Use Multi-Stage Builds
>   - Build your code in larger preliminary images, copy only the artifacts you need in your final container image, discard the preliminary steps
> - Build from Stable to Frequently Changing
>   - Make your most frequently occurring changes as late in your Dockerfile as possible
> - Use a Single Repository for Functions with Large Layers
>   -  ECR compares each layer of a container image when it is pushed to avoid uploading and storing duplicates

### Lambda Versions and Aliases
so far, we've worked on the mutable version of the lambda, which was immutable. if we publish the lambda, it becomes immutable, and we can no longer change it (code or environment variables), and will have an ARN of its' own.\
We can also give a lambda version an alias, which can be something like "prod", "dev", "test". this alias is mutable, and can act as the constant name for the user to talk to. aliases allow canary deployment (A/B testing as usual). aliases can't reference other aliases.

(demo of creating an alias)\
under <kbd>Actions</kbd>, we choose <kbd>Publish</kbd> and it gets the new version number when we publish it again. under the <kbd>Aliases</kbd> tab, we can create new alias configuration, it can point to one version and direct all traffic, or point to two versions and distribute traffic to them based on weights.

### Lambda with CodeDeploy

we can automate the traffic shift with <cloud>CodeDeploy</cloud>, this is also part of the <cloud>SAM</cloud> (serverless application model) framework.
- linear - grow traffic by percent every time period
- canary - try a percent, wait for a period, then move everything
- allAtOnce - move everything right now

we can create pre and post traffic hooks to trigger rollbacks if needed,

### Lambda Function URL

exposing the lambda function as an http address, without an api browser or stuff like that. a constant address that won't change. "https://\<url-id>.lambda-url.\<region>.on.aws".\
This http is accessible only through the public internet. we might need to use CORS configuration, and be based on resource-based policies. this can point to the latest version or an alias, but not a specific function version.

Security is handled with resource-based policy, which is always in effect. we can also define authentication types.

> - AuthType `NONE` - allow public and unauthenticated access
>   - Resource-based Policy is always in effect (must grant public access)
> - AuthType `AWS_IAM` - IAM is used to authenticate and authorize requests
>   - Both Principal's Identity-based Policy & Resource-based Policy are evaluated
>   - Principal must have lambda:InvokeFunctionUrl permissions
>   - Same account - Identity-based Policy OR Resource-based Policy as ALLOW
>   - Cross account - Identity-based Policy AND Resource Based Policy as ALLOW

### CodeGuru Integration
> Gain insights into runtime performance of your Lambda
functions using <cloud>CodeGuru</cloud> Profiler. CodeGuru creates a Profiler Group for your Lambda function.
> - Supported for Java and Python runtimes
> - Activate from AWS Lambda Console
> - When activated, Lambda adds:
> -   CodeGuru Profiler layer to your function
> -   Environment variables to your function
> -   `AmazonCodeGuruProfilerAgentAccess` policy to your function

### Lambda Limits and Best Practices
Per region limits

> Execution:
> - Memory allocation: 128 MB - 10GB (1 MB increments)
> - Maximum execution time: 900 seconds (15 minutes)
> - Environment variables (4 KB)
> - Disk capacity in the "function container" (in /tmp): 512 MB to 10GB
> - Concurrency executions: 1000 (can be increased)
> 
> Deployment:
> - Lambda function deployment size (compressed .zip): 50 MB
> - Size of uncompressed deployment (code + dependencies): 250 MB
> - Can use the /tmp directory to load other files at startup
> - Size of environment variables: 4 KB

AWS Lambda Best Practices
> - Perform heavy-duty work outside of your function handler
>   - Connect to databases outside of your function handler
>   - Initialize the AWS SDK outside of your function handler
>   - Pull in dependencies or datasets outside of your function handler
> - Use environment variables for:
>   - Database Connection Strings, S3 bucket, etc… don't put these values in your code
>   - Passwords, sensitive values… they can be encrypted using KMS
> - Minimize your deployment package size to its runtime necessities.
>   - Break down the function if need be
>   - Remember the AWS Lambda limits
>   - Use Layers where necessary
> - Avoid using recursive code, never have a Lambda function call itself

</details>


## DynamoDB
<details>
<summary>
Serverless NoSQL database.
</summary>

managed by AWS, scales automatically, NoSQL database, handles better with horizontal scaling.

> - Fully managed, highly available with replication across multiple AZs.
> - NoSQL database - not a relational database.
> - Scales to massive workloads, distributed database.
> - Millions of requests per seconds, trillions of row, 100s of TB of storage.
> - Fast and consistent in performance (low latency on retrieval).
> - Integrated with IAM for security, authorization and administration.
> - Enables event driven programming with DynamoDB Streams.
> - Low cost and auto-scaling capabilities.
> - Standard & Infrequent Access (IA) Table Class.
>
> DynamoDB is made of Tables
> 
> - Each table has a Primary Key (must be decided at creation time).
> - Each table can have an infinite number of items (= rows).
> - Each item has attributes (can be added over time - can be null).
> - Maximum size of an item is 400KB.
> - Data types supported are:
>   - Scalar Types - String, Number, Binary, Boolean, Null.
>   - Document Types - List, Map.
>   - Set Types - String Set, Number Set, Binary Set.


primary keys can be the partition key hash alone, or the combinnation of partition key and sort key. whatever option, the key should be diverse so it creates decent distribution.

demo of creating a table. inside the <cloud>DynamoDb</cloud> page, we click <kbd>Create Table</kbd>, and define the partition key and sort-key (optional), we can customize some settings:
1. Table class - standard or IA (infrequently accessed)
2. Capacity Calculator
3. Read/Write Capacity Setting, autoscaling
4. additional Indexes
5. Encrypting (default key, <cloud>KMS</cloud>, user-supplied key)

we can also click <kbd>View Items</kbd>, and then <kbd>Create Item</kbd>. the item must have the partition key defined, and then have any additional attribute. the primary key must be unique (either the partition key alone or the key + sort key).

### WCU and RCU

- RCU - read capcity unit (throughput)
- WCU - write capcity unit (throughput)

> Provisioned Mode (default)
> - You specify the number of reads/writes per second
> - You need to plan capacity beforehand
> - Pay for provisioned read & write capacity units
> 
> On-Demand Mode
> - Read/writes automatically scale up/down with your workloads
> - No capacity planning needed
> - Pay for what you use, more expensive
>
> You can switch between different modes once every 24 hours

we exceed our provisioned capacity by tapping into the "burstCapacity", if we finish it as well, then we will get throughputExceeded errors, and we should start doing exponential retries. two kinds of read modes - strongly consistent read and eventually consistent read.\
Because DynamoBD is a managed service, it has built-in replication. so if we do a read operation immediately after a write operation, it's possible that the read will go to an instance that wasn't updated yet with the new information, and we will receive "stale" data because of the replication. a strongly consistent read will wait until there are no 'in-transit' replications, and then read the data. we can do this by setting "ConsistentRead" Parameter to `True` in the api calls.



> One Write Capacity Unit (WCU) represents one write per second for an
item up to 1 KB in size.\
> If the items are larger than 1 KB, more WCUs are consumed
> - Example 1: we write 10 items per second, with item size 2 KB. We need $10 * (\frac{2_{KB}}{1_{KB}}) = 20_{WCU}$.
> - Example 2: we write 6 items per second, with item size 4.5 KB. We need $6 * (\frac{5_{KB}}{1_{KB}}) = 30_{WCU}$.
> - Example 3: we write 120 items per minute, with item size 2 KB. We need $\frac{120}{60} * (\frac{2_{KB}}{1_{KB}}) = 4_{WCU}$.
> 
> One Read Capacity Unit (RCU) represents one Strongly Consistent Read per
second, or two Eventually Consistent Reads per second, for an item up to 4
KB in size.\
> If the items are larger than 4 KB, more RCUs are consumed .
> - Example 1: 10 Strongly Consistent Reads per second, with item size 4 KB. we need $10 * (\frac{4_{KB}}{4_{KB}}) = 20_{WCU}$.
> - Example 2: 16 *Eventually Consistent* Reads per second, with item size 12 KB. we need $(\frac{16}{2}) * (\frac{12_{KB}}{4_{KB}}) = 24_{WCU}$.
> - Example 3: 10 Strongly Consistent Reads per second, with item size 6 KB. we need $10 * (\frac{8_{KB}}{4_{KB}}) = 20_{WCU}$ (we must round up 6 KB to 8 KB).

Data is stored inside partitions, Partition Keys go through a hashing algorithm to know to which partition they go to. the number of total partitions is based on the capacity and size.

To compute the number of partitions:
- by capacity: $(\frac{RCUs_{Total}}{3000})+(\frac{WCUs_{Total}}{1000})$
- by size - the more data we have, the more partition: $\frac{Total Size}{10_{GB}}$
- the final number is the higher value of the two, with a built in ceiling.

WCU and RCU are spread evenly across partitions. if we exceed RCU or WCU, we will get throrelling. if we have a hot key or a hot partition, then we might get an error reading from on partition but not when using another partition.\
The solution can be exponential retries, create a better partition key distribution, or switch to <cloud>DynamoDB DAX</cloud>.

On-Demand mode scales up and down based on workloads demand, unlimited capacity (no throtelling), the calculations are the same, but the units are "Read Request Unit" and "Write Request Unit" (RRU, WRU), and are about x2.5 times more expensive.

in our table, we can change the capacity modes after it's been created. we can choose the simple "On-demand" mode or the "Provisioned" mode. if we use the provisioned mode, we can use the capacity calculator to calculate the estimated cost. we can also set a minimum and maximum capacity, based on a target utilization value (read and writes).

### Basic Operations APIs


| API              | Operation                | Capacity | Notes                                   |
| ---------------- | ------------------------ | -------- | --------------------------------------- |
| `PutItem`        | Create\ Update           | WCU      |                                         |
| `Update`         | Create\Update            | WCU      | can be used for atomic counters         |
| `GetItem`        | Read                     | RCU      |                                         |
| `Query`          | Read                     | RCU      | one partitionKey only                   |
| `Scan`           | Read                     | RCU      | inefficient, costly                     |
| `DeleteItem`     | Delete                   | WCU      |                                         |
| `DeleteTable`    | Delete                   |          | Delete everything                       |
| `BatchWriteItem` | Write/Delete (no update) | WCU      | some items can partially fail           |
| `BatchGetItem`   | Read                     | RCU      | one or more tables, some items can fail |


> `PutItem`
> 
> - Creates a new item or fully replace an old item (same Primary Key).
> - Consumes WCUs.
> 
> `UpdateItem`
> 
> - Edits an existing item's attributes or adds a new item if it doesn't exist.
> - Can be used to implement Atomic Counters - a numeric attribute that's unconditionally incremented.
> 
> Conditional Writes
> - Accept a write/update/delete only if conditions are met, otherwise returns an error.
> - Helps with concurrent access to items.
> - No performance impact.
> 
> `GetItem`
> - Read based on Primary key.
> - Primary Key can be HASH or HASH+RANGE.
> - Eventually Consistent Read (default).
> - Option to use Strongly Consistent Reads (more RCU - might take longer).
> - ProjectionExpression can be specified to retrieve only certain attributes.
>
> `Query`- returns items based on:
> - KeyConditionExpression:
>   - Partition Key value (must be `=` operator) - required.
>   - Sort Key value (`=`, `<`, `<=`, `>`, `>=`, `Between`, `Begins with`) - optional.
> - FilterExpression:
>   - Additional filtering after the Query operation (before data returned to you).
>   - Use only with non-key attributes (does not allow HASH or RANGE attributes).
> - Ability to do pagination on the results
> - Can query table, a Local Secondary Index, or a Global Secondary Index
>
> `Scan`- the entire table and then filter out data (inefficient)
> Returns up to 1 MB of data - use pagination to keep on reading
> - Consumes a lot of RCU.
> - Limit impact using Limit or reduce the size of the result and pause.
> - For faster performance, use Parallel Scan.
>   - Multiple workers scan multiple data segments at the same time.
>   - Increases the throughput and RCU consumed.
>   - Limit the impact of parallel scans just like you would for Scans.
> - Can use ProjectionExpression & FilterExpression (**no changes to RCU**).
>
> `DeleteItem`
> - Delete an individual item.
> - Ability to perform a conditional delete.
> 
> `DeleteTable`
> - Delete a whole table and all its items.
> - Much quicker deletion than calling `DeleteItem` on all items.
>
> `BatchWriteItem`
> - Up to 25 PutItem and/or DeleteItem in one call.
> - Up to 16 MB of data written, up to 400 KB of data per item.
> - Can't update items (use UpdateItem).
> - UnprocessedItems for failed write operations (exponential backoff or add WCU).
> 
> `BatchGetItem`
> - Return items from one or more tables.
> - Up to 100 items, up to 16 MB of data.
> - Items are retrieved in parallel to minimize latency.
> - UnprocessedKeys for failed read operations (exponential backoff or add RCU).

also has a query language:

> DynamoDB - `PartiQL`
> -  SQL-compatible query language for DynamoDB
> -  Allows you to select, insert, update, and delete data in DynamoDB using SQL
> - (**no JOIN operations**)
> -  Run queries across multiple DynamoDB tables
> -  Run PartiQL queries from:
>   -  AWS Management Console
>   -  NoSQL Workbench for DynamoDB
>   -  DynamoDB APIs
>   -  AWS CLI
>   -  AWS SDK

```SQL
SELECT OrderID, Total
FROM Orders
WHERE OrderID in [1,2,3]
ORDER BY OrderID DESC
```

in the console, we can create items, update and delete them, or do batch delete, filters are client-side, they don't effect the RCU charges. the query conditions (partition + sort key) can reduce the RCU costs.


#### WriteTypes

- Concurrent Write - two operations at the same time, one overwrites the other.
- Conditional Write - one succeeded, one failed. do the write only if some condition applies. (optimistic locking)
- Atomic Write - both succeed. update value with operations.
- Batch Write - send both at the same - saves up on WCU and complexity

#### Conditional Writes / Optimistic Locking

> - DynamoDB has a feature called "Conditional Writes"
> - A strategy to ensure an item hasn't changed before you update/delete it.
> - Each item has an attribute that acts as a version number

we check and modify this attribute when we do a write, so if two operation happen, only one will be able to change the data

> For `PutItem`, `UpdateItem`, `DeleteItem`, and `BatchWriteItem`:\
> You can specify a Condition expression to determine which items should be modified:
> 
> - `attribute_exists` - delete if attribute exists
> - `attribute_not_exists` - delete if attribute doesn't exist
> - `attribute_type` - attribute is of certain
> - `contains` (for string) - check for substring
> - `begins_with` (for string) - check for prefix
> - `IN` - check if value is inside predetermined list
> - `size` (string length) - size of string
> 
> Note: Filter Expression filters the results of read queries, while Condition
Expressions are for write operations.

if we don't want to overwrite elements, we can use `attribute_not_exists(partition_key)`, so the conditional write will make sure the item isn't overwritten if it exists.

example of updating an item, but only if it fulfills a condition.

```json
{
  ":discount":{"N":"150"},
  ":limit":{"N":"500"},
  ":lo":{"N":"500"},
  ":hi":{"N":"500"},
  ":cat1":{"S":"Cars"},
  ":cat":{"S":"Shirt"},
  ":v_sub":{"S":"http://"}
}
```

```sh
# create item
aws dynamodb update-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --update-expression "SET Price = Price - :discount" --condition-expression "Price > :limit" --expression-attribute-values file://values.json
# conditional delete item
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "attribute_not_exists(Price)"
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "attribute_exists(ProductReviews.OneStart)"
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "(ProductCategory IN (:cat1, :cat2)) and (Price between :lo and :hi)" --expression-attribute-values file://values.json
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "begins_with(Pictures.FrontView, :v_sub)" --expression-attribute-values file://values.json
```
### Indexes

- LSI - Local Secondary Index
- GSI - Global Secondary Index

LSI acts as alternative **sort key** for the table, same partition key. up to 5 LSI per table, **must be created when the table is created**. the new sort key must be a scalar attribute, can have projections (`KEYS_ONLY`, `INCLUDE_ALL`).
> - Uses the WCUs and RCUs of the main table.
> - No special throttling considerations.

GSI acts as an alternative **primary key**, usefull speeding up queries on non-key attributes. must provision RCU and WCU for the new index. can be added and modified after creation. this is basically replication of the table with different partitions.
>- If the writes are throttled on the GSI, then the main table will be throttled!
>  - Even if the WCU on the main tables are fine
>- Choose your GSI partition key carefully!
>- Assign your WCU capacity carefully!

(demo of creating local and global index, then querying them)

### PartiQL

> Use a SQL-like syntax to manipulate DynamoDB tables
> Supports some (but not all) statements:
>   - INSERT
>   - UPDATE
>   - SELECT
>   - DELETE
> 
> It supports Batch operations

in the console, we can choose <kbd>PartiQL editor</kbd>, we have a built in statements to start from.

### DAX

DynamoDB Accelerator

> -  Fully-managed, highly available, seamless in-memory cache for DynamoDB
> -  Microseconds latency for cached reads & queries
> -  Doesn't require application logic modification (compatible with existing DynamoDB APIs)
> -  Solves the "Hot Key" problem (too many reads)
> -  5 minutes TTL for cache (default)
> -  Up to 10 nodes in the cluster
> -  Multi-AZ (3 nodes minimum recommended for production)
> -  Secure (Encryption at rest with <cloud>KMS</cloud>, <cloud>VPC</cloud>, <cloud>IAM</cloud>, <cloud>CloudTrail</cloud>, etc...)

can be used in combinnation with <cloud>ElasticCache</cloud>, use DAX for simple queries on individual objects, and use <cloud>ElasticCache</cloud> to store aggregate results.

we need to choose the cluster type and number of nodes, we put them into VPCs and subgroup and security group (we need to open a specific port). an endpoint is created which we can direct our applications at.

### Streams

> Ordered stream of item-level modifications (create/update/delete) in a table.
> - Stream records can be:
>   - Sent to <cloud>Kinesis Data Streams</cloud>
>   - Read by <cloud>Kinesis Client Library applications</cloud>
>   - Read by <cloud>AWS Lambda</cloud>
> - Data Retention for up to 24 hours
> 
> - Use cases:
>   - react to changes in real-time (welcome email to users)
>   - Analytics
>   - Insert into derivative tables
>   - Insert into OpenSearch Service
>   - Implement cross-region replication

all operation are stored in the dynamoDB streams, we can add a processing layer if needed.

> Ability to choose the information that will be written to the stream:
>   - `KEYS_ONLY` - only the key attributes of the modified item
>   - `NEW_IMAGE` - the entire item, as it appears after it was modified
>   - `OLD_IMAGE` - the entire item, as it appeared before it was modified
>   - `NEW_AND_OLD_IMAGES` - both the new and the old images of the item
> DynamoDB Streams are made of shards, just like Kinesis Data Streams You don't provision shards, this is automated by AWS. Records are not retroactively populated in a stream after enabling it.


to work with lambda, we define the event source mapping, give it the appropriate permissions, and then it starts polling from the stream. we have a batch size, as usual.

### TTL Configuration
TTL is Time-to-live, automatically delete the item based on the timestamp, expired items don't consume WCU to delete, and deletions will happen within 48 hours since the expiration time. deleted items are removed from the indexes, and it also adds the delete operation into the stream (so we could monitor it like a regular item). the TTL must be a numeric attribute.

### DynamoDB CLI

> - `--projection-expression`: one or more attributes to retrieve
> - `--filter-expression:` filter items before returned to you
> 
> General AWS CLI Pagination options (e.g., <cloud>DynamoDb</cloud>, <cloud>S3</cloud>, etc...)
> 
> - `--page-size`: specify that AWS CLI retrieves the full list of items but with a larger number of API calls instead of one API call (default: 1000 items).
> - `--max-items`: max number of items to show in the CLI (returns *NextToken*)
> - `--starting-token`: specify the last *NextToken* to retrieve the next set of items

### Transactions

either everyworks, or nothing works and everything failed.

> Coordinated, **all-or-nothing operations** (add/update/delete) to multiple items across one or more tables.
> - Provides Atomicity, Consistency, Isolation, and Durability (ACID)
> - Read Modes - Eventual Consistency, Strong Consistency, Transactional
> - Write Modes - Standard, Transactional
> - **Consumes 2x WCUs & RCUs**
> - DynamoDB performs 2 operations for every item (prepare & commit)
> - Two operations:
>   - `TransactGetItems` - one or more `GetItem` operations
>   - `TransactWriteItems` - one or more `PutItem`, `UpdateItem`, and `DeleteItem` operations
> 
> - Use cases: financial transactions, managing orders, multiplayer games, etc...

### Session State Table
using DynamoDB to store session state as cache, common use-case for <cloud>DynamoDB</cloud> and <cloud>ElasticCache</cloud>.

> - vs. <cloud>ElastiCache</cloud> - ElastiCache is in-memory, but DynamoDB is serverless, Both are key/value stores.
> - vs. <cloud>EFS</cloud> -  EFS must be attached to <cloud>EC2</cloud> instances as a network drive.
> - vs. <cloud>EBS</cloud> & Instance Store - EBS & Instance Store can only be used for local caching, not shared caching.
> - vs. <cloud>S3</cloud> - S3 is higher latency, and not meant for small objects.

### Partitioning Strategies

if we don't have a good partition key to act as a distributed key, we might get hot-key and bad distribution. we can get around this by adding a suffix, either randomly or by calculation.

### Patterns with S3
we shouldn't store large objects in <cloud>DynamoDB</cloud>, instead, we store the large objects in a bucket and save the URL into the record. the client will read the record and then pull the big object from <cloud>S3</cloud>.\
We can also use <cloud>DynamoDB</cloud> to index <cloud>S3</cloud> data. when an object is created in the bucket, we trigger a lambda and store the data in dynamodb with some relevant data.

### Operations

Table cleanup:
- Scan + DeleteItem - slow, expensive
- Drop Table + recreate it

Copying tables
- use <cloud>Date Pipeline</cloud> (uses <cloud>EMR</cloud> internally)
- backup and restore data - takes more time.
- scan + putItem - write your own code

### Security Considerations and Other

> Security
> - VPC Endpoints available to access DynamoDB without using the Internet
> - Access fully controlled by IAM
> - Encryption at rest using AWS KMS and in-transit using SSL/TLS
> 
> Backup and Restore feature available
> - Point-in-time Recovery (PITR) like RDS
> - No performance impact
> 
> Global Tables 
> - Multi-region, multi-active, fully replicated, high performance
> 
> DynamoDB Local
> - Develop and test apps locally without accessing the DynamoDB web service (without Internet)
> 
> AWS <cloud>Database Migration Service</cloud> (AWS DMS) can be used to migrate to DynamoDB (from MongoDB, Oracle, MySQL, S3, ...)

if clients need to interact with the database directly, we can give them temporary credentails through an identity provider. we use the Cognito Identity Pool and have a specific IAM Role policy that limits access based on condition (leading key or attribute).
</details>

## API Gateway
<details>
<summary>
Build, Deploy And Manage APIs. Exposing our application as REST API.
</summary>

in our serverless world, we want the client to invoke the Lambda operations, rather than using the lambda directly or a public IP address with a load balancer, we can use <cloud>API Gateway</cloud>.

> - AWS Lambda + API Gateway: No infrastructure to manage
> - Support for the WebSocket Protocol
> - Handle API versioning (v1, v2,...)
> - Handle different environments (dev, test, prod,...)
> - Handle security (Authentication and Authorization)
> - Create API keys, handle request throttling
> - Swagger / Open API import to quickly define APIs
> - Transform and validate requests and responses
> - Generate SDK and API specifications
> - Cache API responses

gateways can integrate with
- lambda function
- any http endpoint
- aws service - connect with any AWS API

even if we don't need the REST api behavior, we can make use of the additional features of authentication, throtelling, etc...

EndPoint Types:
- Edge optimized (default) - routing through <cloud>CloudFront</cloud> edge locations.
- Regional - for clients in the same region, for manual configurations.
- Private - can only be accessed from your VPC using a VPC endpoint (<cloud>ENI</cloud>).

Security:
- authentication through
  - IAM roles
  - <cloud>Cognito</cloud> - for external users
  - custom Authorizers
- HTTPS security
  - using <cloud>AWS Certificate Manager</cloud>
  - combine with CNAME or A-Alias in <cloud>Route53</cloud>

(demo)\
we can choose from HTTP API, REST API, REST API private and WebSocket APIs. we select the REST API version and start from a new API. we use the Regional API type. then we define methods with <kbd>Create Methods</kbd> and we define the type and the integration. we use a lambda function for our demo. we need an <cloud>IAM</cloud> role that can invoke it (resource based access polices). API gateway has it's our timeout, regardless of the lambda duration. we also turn on the Lambda Proxy integration (to restructure the http request to an event). we can run test request and set the parameters and headers ourselves. we can also <kbd>Create Resources</kbd> with method on the specific resources.\
when we're ready, we can deploy the API, which gives us an invoke URL, which we can use in the broweser.

### Stages And Deployments
changes aren't effective until we run a deployment. changes are deployed to "stages", which can have any name, each stage has configuration parametes, and can be rolled back for a history of deployments in this stage. each stage has its' own URL. this allows us to run multiple version and the same time.

stages also have variables, we can change them without deploying the API again. those variables can be used for ARN, HTTP endpoint or parameter mapping templates. they are part of the context object in the lambda.

we create a lambda function with several aliases. and we create new API method, and in the lambda ARN, we add `:${stageVariable.LambdaAlias}`. we might need to change IAM roles. once we deploy the gateway, we can manage the mapping and add the stageVariables for each stage. we can add other configurations, set up specific log behavior, detailed metrics and X-Ray tracing for each stage.

#### Canary Deployment

like a blue/green deployment, we direct some traffic to two stages. we can then promote the canary deployment to 100%.

### Integration Types and Mapping.

> - Integration Type `MOCK`
>   - API Gateway returns a response without sending the request to the backend.
> 
> - Integration Type `HTTP` / `AWS` (Lambda & AWS Services)
>   - you must configure both the integration request and integration response.
>   - Setup data mapping using mapping templates for the request & response.
> 
> - Integration Type `AWS_PROXY` (Lambda Proxy):
>   - incoming request from the client is the input to Lambda.
>   - The function is responsible for the logic of request / response.
>   - No mapping template, headers, query string parameters are passed as arguments.
> 
> - Integration Type `HTTP_PROXY`
>   - No mapping template.
>   - The HTTP request is passed to the backend.
>   - The HTTP response from the backend is forwarded by API Gateway.
>   - Possibility to add HTTP Headers if need be (ex: API key).
>
>
> Mapping templates can be used to modify request / responses
> - Rename / Modify query string parameters
> - Modify body content
> - Add headers
> - Uses <cloud>Velocity Template Language</cloud> (VTL): (for loop, if-else)
> - Filter output results (remove unnecessary data)
> - Content-Type can be set to "application/json" or "application/xml"

a common example is transforming between json and XML for SOAP apis(which are xml based). we demo this by creating a method and not enabling lambda proxy. we can click on "integration response", and create a mapping template. this is a json document (with extra steps) that modifies the response before returning it.

```json
{
  "my-key": "my-value",
  "renamed-key": $input.json('$.example')
}
```

### Open API
API gateway can be integrated with the OPEN API specifications. we can import or export specification to that format. we can then generate clientSide SDK.\
we can also use Request Validation to make sure the request conforms to the specification, then we can validate it our the api without going to the backend. this can be defined for all methods of just some of them.

### Gateway Caching

Caching reduces the number of calls made to
the backend.

> - Default TTL (time to live) is 300 seconds (min: 0s, max: 3600s)
> - Caches are defined per stage
> - Possible to override cache settings per method
> - Cache encryption option
> - Cache capacity between 0.5GB to 237GB
> - Cache is expensive, makes sense in production, may not make sense in dev / test
> - Able to flush the entire cache (invalidate it) immediately
- Clients can invalidate the cache with header: `Cache-Control: max-age=0` (with proper IAM authorization)
- If you don't impose an InvalidateCache policy (or choose the Require authorization check box in the console), any client can invalidate the API cache.

### Usage Plans

we can control access to our APIs with usage plans and API Keys.

> Usage Plan:
> 
> - who can access one or more deployed API stages and methods
> - how much and how fast they can access them
> - uses API keys to identify API clients and meter access
> - configure throttling limits and quota limits that are enforced on individual client
> 
> API Keys:
> 
> - alphanumeric string values to distribute to your customers
>   - Ex: WBjHxNtoAb4WPKBC7cGm64CBibIb24b4jt8jJHo9
> - Can use with usage plans to control access
> - Throttling limits are applied to the API keys
> - Quotas limits is the overall number of maximum requests
> 
> To configure a usage plan
> 1. Create one or more APIs, configure the methods to require an API key, and deploy the APIs to stages.
> 2. Generate or import API keys to distribute to application developers (your customers) who will be using your API.
> 3. Create the usage plan with the desired throttle and quota limits.
> 4. Associate API stages and API keys with the usage plan.
> 
> Callers of the API must supply an assigned API key in the `x-api-key` header in requests to the API.

### Monitoring Gateways

we can intgrate the gateway with <cloud>CloudWatch</cloud>, enabled at the stage level, and override with per API settings. the logs contain information about requests and responses.\
we can enable <cloud>X-Ray</cloud> to have the requests traced (so we could follow them from the gateway to the lambda).\
There are <cloud>CloudWatch metrics</cloud>, such as `CacheHitCount` and `CacheMissCount` to track the efficacy of the cacheing, we can follow the overhead latency with the `Latency` and `IntegrationLatency` values. we can track the `4XXError` and `5XXError` metrics to count errors of varus types.\
The maximum time for the API to await a response is 29 seconds.\

Throtelling:
- Account Limit : by default we have 10000 request per second across all APIs.
  - can be increased with request to AWS.
- Stage and Method Limit - limit the number of request for a API so it wouldn't affect other APIs.
- Usage Plans - throttle per customer

### CORS

CORS - Cross Origin Resource Sharing.

> CORS must be enabled when you receive API calls from another domain.\
> The `OPTIONS` pre-flight request must contain the following headers:
>  - Access-Control-Allow-Methods
>  - Access-Control-Allow-Headers
>  - Access-Control-Allow-Origin
> 
> CORS can be enabled through the console

### Authentication and Authorization

<cloud>IAM</cloud> permissions: attach to user or role and allow the role to invoke the the API. combines with "sig4" to validate the IAM credentials in the headers. both Authentication and Authorization are handled by IAM.\
Resource Policies can set who can interact with the Gateway, and can be used for Cross Account Access.

><cloud>Cognito User Pools</cloud>
> - Cognito fully manages user lifecycle, token expires automatically
> - API gateway verifies identity automatically from AWS Cognito
> - No custom implementation required
> - Authentication: Cognito User Pools
> - Authorization: API Gateway methods
>
> Lambda Authorizer (formerly Custom Authorizers)
> - Token-based authorizer (bearer token) - ex JWT (JSON Web Token) or Oauth
> - A request parameter-based Lambda authorizer (headers, query string, stage var)
> - Lambda must return an IAM policy for the user, result policy is cached
> - Authentication: External service
> - Authorization: Lambda function

for each method, we can look at the "method request" and see the authentication. we can also use resource policies for the entire gateway. we can also create a custom authorizer (lambda) or use a coginto user pool.

### Gatway Types

different types of gateways:
- HTTP API
- REST API
- WebSocket API

> HTTP APIs
> - low-latency, cost-effective AWS Lambda proxy, HTTP proxy APIs and private integration (no data mapping)
> - support OIDC and OAuth 2.0 authorization, and built-in support for CORS
> - No usage plans and API keys
> 
> REST APIs
> - All features (except Native OpenID Connect / OAuth 2.0)
>
> WebSockets:\
> Two-way interactive communication
between a user's browser and a server
> - Server can *push* information to the client
> - This enables *stateful* application use cases
> - WebSocket APIs are often used in realtime applications such as:
>   - chat applications.
>   - collaboration platforms.
>   - multiplayer games.
>   - financial trading platforms.
> - Works with AWS Services (<cloud>Lambda</cloud>, <cloud>DynamoDB</cloud>) or HTTP endpoints.

uses different lambda methods for each event (or message):
- `onConnect` - first request, establishing connection
- `onDisconnect` - when connection is closed, best-effort behavior. not guarantted to happen
- `default` - fallback
- other messages can be defined for each api.

the webSocket url has this form: "wss://[some-unique-id].execute-api.[region].amazonaws.com/[stage-name]". each message in the connection is called "a frame", and re-uses the connectionId. the client can listen to messages from the backend via a request to a callaback url "wss://[some-unique-id].execute-api.[region].amazonaws.com/[stage-name]/@connection/[connection-id]". the client will listen to this and respond.
- `POST` - Send a message from the server to the connected WS client
- `GET` - Gets the larest connection status of the connected WS Client
- `DELETE` - Disconnect the connected client from the WS connection.

routing for websockets are based on fields in the message. if there is no match, then `$default` is route is used.

| <>                              | HTTP                          | REST       | Websocket             |
| ------------------------------- | ----------------------------- | ---------- | --------------------- |
| <cloud>Lambda Authorize</cloud> | YES                           | YES        |                       |
| <cloud>IAM</cloud>              | YES                           | YES        |                       |
| Resource Policies               | NO                            | YES        |                       |
| <cloud>Cognito</cloud>          | YES                           | YES        |                       |
| Native OpenId/OAth 2.0/ JWT     | YES                           | NO         |                       |
| Data Mapping                    | NO                            | YES        |
| Usage Plans and API Keys        | No                            | YES        |
| Use case                        | low cost proxies, integration | normal use | realtime applications |

</details> 

## AWS CI/CD: CodeCommit, CodeBuild and CodeDeploy
<details>
<summary>
Devops tools and flows.
</summary>

CI/CD - Continues Integration, Continues Deployment

so far we did mostly manual steps, but we want to do those actions in an automated way, and have them recorded into code.


- <cloud>CodeCommit</cloud> - Store code
- <cloud>CodePipeline</cloud> - automate building to <cloud>Elastic Beanstalk</cloud>
- <cloud>CodeBuild</cloud> - building and testing
- <cloud>CodeDeploy</cloud> - deploy to <cloud>EC2</cloud> instances
- <cloud>CodeStar</cloud> - manage at a single location
- <cloud>CodeArtifact</cloud> - store, publish and share software artifacts
- <cloud>CodeGuru</cloud> - code review by machine learning
- <cloud>Cloud9</cloud> - in-browser IDE

### CodeCommit

code repository, like github, gitlab or bitbucket. source code version control. rollbacks. provides the same features as github, but it lives on the AWS cloud and is protected by AWS. integrates seamlessly with industry standard tooling.\
Can use <cloud>IAM</cloud> as authentication and authorization, uses <cloud>KMS</cloud> for encryption.

(demo)\
we create a repository and we get the connection options. we can upload files and add a commit message.
- branches
- tags
- pull requests
- commit visualizers
- comparisons

we can set up notification rules, which will trigger based on events, and we can set target (<cloud>SNS</cloud> and Slack Chatbot). triggers can happen on specific code events, and can send to <cloud>SNS</cloud> or <cloud>Lambda</cloud>.

We can set approval rules and templates for pull requests.

in <cloud>IAM</cloud>, we can make Access Keys for our repository.
- SSH keys for <cloud>CodeCommit</cloud>
- Git Credentails keys for <cloud>CodeCommit</cloud>
(if we have existing SSH keys, we can upload them here).

### CodePipeline

> Visual Workflow to orchestrate your CI/CD
> - Source - <cloud>CodeCommit</cloud>, <cloud>ECR</cloud>, <cloud>S3</cloud>, Bitbucket, GitHub
> - Build - <cloud>CodeBuild</cloud>, Jenkins, CloudBees, TeamCity
> - Test - <cloud>CodeBuild</cloud>, AWS Device Farm, 3rd party tools, etc...
> - Deploy - <cloud>CodeDeploy</cloud>, <cloud>Elastic Beanstalk</cloud>, <cloud>CloudFormation</cloud>, <cloud>ECS</cloud>, <cloud>S3</cloud>, ...
> - Invoke - <cloud>Lambda</cloud>, <cloud>Step Functions</cloud>
> 
> Consists of stages:
> - Each stage can have sequential actions and/or parallel actions
> - Example: Build -> Test -> Deploy -> Load Testing
> - Manual approval can be defined at any stage.

Each pipeline creates artifcats, they are stored in <cloud>S3</cloud> and are passed to the next stage.

> - For CodePipeline Pipeline/Action/Stage Execution State Changes
> - Use <cloud>CloudWatch</cloud> Events (Amazon <cloud>EventBridge</cloud>). Example:
>   - You can create events for failed pipelines
>   - You can create events for cancelled stages
> - If CodePipeline fails a stage, your  pipeline stops, and you can get information in the console
>- If pipeline can't perform an action, make sure the "IAM Service Role" attached does have enough IAM permissions (IAM Policy).
>- AWS <cloud>CloudTrail</cloud> can be used to audit AWS API calls.

(demo)\
before we do anything, we create an <cloud>ELastic Beanstalk</cloud> environment, we will use the CI/CD tools to deploy into it.\ in <cloud>CodePipeline</cloud> we first <kbd>Create Pipeline</kbd>, it needs a service Role with appropriate permissions. we can specify the artifact storeage, either the default <cloud>S3 Bucket</cloud> or a different one, and we set the source to be the <cloud>CodeCommit</cloud> repository. we next set the detection rule to use <cloud>CloudWatch Events</cloud>. we can set the build provider (<cloud>CodeBuild</cloud>, Jenkins) and define a deploy stage, can be:

- AWS <cloud>CloudFromation</cloud>
- AWS <cloud>CodeDeploy</cloud>
- AWS <cloud>Elastic Beanstalk</cloud>
- AWS <cloud>Service Catalog</cloud>
- AWS <cloud>ECS</cloud>, ECS (blue-green)
- AWS <cloud>S3</cloud>

when the pipeline is created, it does the first run automatically. for our example, it will upload the code to change the beanstalk appplication.\
We can add stages, and define action groups, they interact with action providers (either aws services or manual), action groups can be sequential or parallel.\
now we can commit our code to <cloud>CodeCommit</cloud>, and see how the pipeline starts and goes through all of the stages.

### CodeBuild

> A fully managed continuous integration (CI) service.
> - Continuous scaling (no servers to manage or provision – no build queue)
> - Compile source code, run tests, produce software packages, etc...
> - Alternative to other build tools (e.g., Jenkins)
> - Charged per minute for compute resources (time it takes to complete the builds)
> - Leverages Docker under the hood for reproducible builds
> - Use prepackaged Docker images or create your own custom Docker image
> - Security:
>   - Integration with KMS for encryption of build artifacts
>   - IAM for CodeBuild permissions, and VPC for network security
>   - AWS CloudTrail for API calls logging

- Source - repository (github, bitbucket, <cloud>CodeCommit</cloud>)
- Build instructions - in the "buildspec.yml" file
- output logs
- <cloud>CloudWatch Logs</cloud> and Alarms

supported environments:
- Java
- Ruby
- Python
- GO
- Node.js
- Android
- .Net Core
- PHP
- Docker - extended custom environment

the instructions for the build are inside the "buildspec.yaml" file. this can take some time, but we can add caching if needed. artifacts are stored in S3 bucket.

(demo)\
rather than simply deploying, we can use <cloud>CodeBuild</cloud> to run tests. we can set the build to run inside a VPC, or give it access to an <cloud>Elastic File System</cloud>. we can also add the build stage to our <cloud>CodePipeline</cloud> from before, this is an action group.

### CodeDeploy

> - Deployment service that automates application deployment
> - Deploy new applications versions to EC2 Instances, On-premises servers, Lambda functions, ECS Services
> - Automated Rollback capability in case of failed deployments, or trigger CloudWatch Alarm
> - Gradual deployment control
> - A file named "appspec.yml" defines how the
deployment happens
>

Deploys to EC2 machine, Lambda Function or ECS clusters.

> Can deploy to EC2 Instances & on-premises servers
> - Perform in-place deployments or blue/green deployments
> - Must run the <cloud>CodeDeploy</cloud> Agent on the target instances. (The CodeDeploy Agent must be running on the EC2 instances as a prerequisites)
>   - It can be installed and updated
automatically if you're using Systems Manager.
>   - The EC2 Instances must have sufficient permissions to access Amazon S3 to get deployment bundles
> - Define deployment speed
>   - AllAtOnce: most downtime
>   - HalfAtATime: reduced capacity by 50%
>   - OneAtATime: slowest, lowest vailability impact
>   - Custom: define your %
> 
> CodeDeploy can help you automate traffic shift for Lambda aliases.
> - Feature is integrated within the SAM framework
> - Linear: grow traffic every N minutes
until 100%
>   - LambdaLinear10PercentEvery3Minutes
>   - LambdaLinear10PercentEvery10Minutes
> - Canary: try X percent then 100%
>   - LambdaCanary10Percent5Minutes
>   - LambdaCanary10Percent30Minutes
> - AllAtOnce: immediate
> 
> CodeDeploy can help you automate
the deployment of a new ECS Task
Definition
> - **Only Blue/Green Deployments**
> - Linear: grow traffic every N minutes
until 100%
>   - ECSLinear10PercentEvery3Minutes
>   - ECSLinear10PercentEvery10Minutes
> - Canary: try X percent then 100%
>   - ECSCanary10Percent5Minutes
>   - ECSCanary10Percent30Minutes
> - AllAtOnce: immediate

for our demo, we create <cloud>IAM</cloud> roles, one for the <cloud>CodeDeploy</cloud> service to manage <cloud>EC2</cloud> instances, and one for instances to pull from <cloud>S3</cloud>. we also install the agent on the EC2 instances. we create a deployment group from it, we need to tag our instance with a proper tag, so we would be able to identify which machines belong to the tag group.\
we set the deployment setting from one of the built-ins, or create a configuration of our own. we can also work with the load balancer.

we look at the "appspec.yml" to see how the deployment behaves.

> CodeDeploy – Deployment to EC2
> - Define how to deploy the application using appspec.yml + Deployment Strategy
> - Will do In-place update to your fleet of EC2 instances
> - Can use hooks to verify the deployment after each deployment phase
>
> CodeDeploy – Deploy to an ASG
> - In-place Deployment 
>   - Updates existing EC2 instances
>   - Newly created EC2 instances by an ASG will also get automated-deployments
> - Blue/Green Deployment
>   - A new Auto-Scaling Group is created (settings are copied)
>   - Choose how long to keep the old EC2 instances (old ASG)
>   - Must be using an ELB

we can do redeploys and rollbacks, either manually or automatically (based on health checks, deployment failure)

### CodeStar

<cloud>CodeStar</cloud> is being replaced by <cloud>CodeCatalyst</cloud>.

> An integrated solution that groups: GitHub, <cloud>CodeCommit</cloud>, <cloud>CodeBuild</cloud>, <cloud>CodeDeploy</cloud>, cloud>CloudFormation</cloud>, <cloud>CodePipeline</cloud>, <cloud>CloudWatch</cloud>, ...
> - Quickly create "CI/CD-ready" projects for <cloud>EC2</cloud>, <cloud>Lambda</cloud>, <cloud>Elastic Beanstalk</cloud>
> - Supported languages: 
>   - C#
>   - Go
>   - HTML 5
>   - Java
>   - Node.js
>   - PHP
>   - Python
>   - Ruby
> - Issue tracking integration with JIRA / GitHub Issues
> 
> Ability to integrate with <cloud>Cloud9</cloud> to obtain a web IDE (not all regions)
> - One dashboard to view all your components
> - Free service, pay only for the underlying usage of other services
> - Limited Customization

our demo will use python on <cloud>Elastic Beanstalk</cloud> as the starting template, looking at <cloud>CodeCommit</cloud> repository, we set the <cloud>EC2</cloud> configuration. the project is controlled via <cloud>CloudFormation</cloud> which handles creating the resources for us. we can see everything that was created for us. we can add an issue tracker (jira) or invite other users to our project.

### CodeArtifact

> Software packages depend on each other to be built (also called code dependencies), and new ones are created.\
> Storing and retrieving these dependencies is called *artifact management*. Traditionally you need to setup your own artifact management system.\
> <cloud>CodeArtifact</cloud> is a secure, scalable, and cost-effective artifact
management for software development
> - Works with common dependency management tools such as Maven, Gradle, npm, yarn, twine, pip, and NuGet
> - Developers and CodeBuild can then retrieve dependencies straight from CodeArtifact

Artifacts exists inside the VPC, we can use it as a proxy for public artifact repositories (which gives us some security, ana also gives us caching). <cloud>CodeBuild</cloud> should also use this option. we can have <cloud>EventBridge</cloud> notifications when a package is created/modified/deleted, and we can set up a response, such as re-building an application that has had a dependency change.

we can configure resource policies for cross account access, which can be either all packages or none of them.

(demo)\
we define an artifact repository that will track an external source (upstream repository). we need to define a domain.
we connect to the artifactory through the cloudShell console, and install the packages that we want on our repository.

we can have policies for repositories or domains.

### CodeGuru

> An ML-powered service for automated code reviews and application performance recommendations.\
> Provides two functionalities
> - CodeGuru Reviewer: automated code reviews for static code analysis (development)
> - CodeGuru Profiler: visibility/recommendations about application performance during runtime (production).
>
> CloudGuru Reviewer:\
> Identify critical issues, security vulnerabilities, and hard-to-find bugs
> Example: common coding best practices,
resource leaks, security detection, input
validation
> - Uses Machine Learning and automated
reasoning
> - Hard-learned lessons across millions of
code reviews on 1000s of open-source and Amazon repositories
> - Supports Java and Python
> - Integrates with GitHub, Bitbucket, and
AWS <cloud>CodeCommit</cloud>
>
> CloudGuru Profiler:\
> Helps understand the runtime behavior of your application
> - Example: identify if your application is consuming excessive CPU capacity on a logging routine
> - Features:
>   - Identify and remove code inefficiencies
>   - Improve application performance (e.g., reduce CPU utilization)
>   - Decrease compute costs
>   - Provides heap summary (identify which objects using up memory)
>   - Anomaly Detection
> - Support applications running on AWS or on-premise
> - Minimal overhead on application

the profiler works with an agent, which has some configurations we can control.
- MaxStackDepth
- MemoryUsageLimitPercent
- MinimumTimeForReportingInMilliseconds
- ReportingIntervalInMilliseconds
- SamplingIntervalInMilliseconds

### Cloud9

> Cloud-based Integrated Development
Environment (IDE).
> - Code editor, debugger, terminal in a browser
> - Work on your projects from anywhere with
an Internet connection
> - Prepackaged with essential tools for popular programming languages (JavaScript, Python, PHP, ...)
> - Share your development environment with your team (pair programming)
> - Fully integrated with AWS SAM & Lambda to easily build serverless applications

demo of using <cloud>Cloud9</cloud> machines. comes with the <cloud>AWS Explorer</cloud> to list all of our tools. also integrated with <cloud>CloudWhisperer</cloud>.
</details>

## SAM - Serverless Application Model
<details>
<summary>
Serverless Application Model - easy deployment
</summary>
A short cut for <cloud>CloudFormation</cloud>, setting up connections between serverless services.

> - SAM = Serverless Application Model
> - Framework for developing and deploying serverless applications
> - All the configuration is YAML code
> - Generate complex CloudFormation from simple SAM YAML file
> - Supports anything from CloudFormation: Outputs, Mappings, Parameters, Resources...
> - Only two commands to deploy to AWS
> - SAM can use CodeDeploy to deploy Lambda functions
> - SAM can help you to run Lambda, API Gateway, DynamoDB locally

SAM uses "Recipes", it's yaml files with the top level field of "Transform: 'AWS::Serverless-2016-10-31'".

Types:
- `AWS::Serverless::Function` - <cloud>Lambda</cloud>
- `AWS::Serverless::API` - <cloud>ApiGateway</cloud>
- `AWS::Serverless::SimpleTable` - <cloud>DynamoDb</cloud>

cli commands:
- `aws cloudformation package`
- `aws cloudformation deploy`
- `sam init`
- `sam build` (local)
- `sam package`
- `sam deploy`
- `sam deploy --guided`

we transform our recipe into a <cloud>CloudFormation</cloud> template with `sam build`, we package (zip and upload to <cloud>S3</cloud>) it with `sam package`, and then deploy with `sam deploy`.

we can also use the CLI tool and combine it with <cloud>AWS Toolkit</cloud> for local debugging. it's integrated with the common IDEs.

we start by installing the sam cli tool, which is not the same as the general `aws` cli tool.

### Our First SAM Project - Lambda Function

we can start with the default sam project by running `sam init`. but instead, we begin with an empty folder and start creating files. one python file and one template file. the python file has a lambda handler, and the template file defines the function name and its' properties (memory size, timeout, policies).


```sh
# create an s3 bucket
aws s3 mb s3://<unique-bucket-name>

# package cloudformation
aws cloudformation package  --s3-bucket <unique-bucket-name> --template-file template.yaml --output-template-file gen/template-generated.yaml
# sam package ... 

# deploy 
aws cloudformation deploy --template-file gen/template-generated.yaml --stack-name <some-stack-name> --capabilities CAPABILITY_IAM
```

this will start the flow to create he lambda.

### API Gateway and Table

we can update our template to also deploy an API gateway and a <cloud>DynamoDB</cloud> table. we modify our lambda code to respond with a json, and we update the template. we just add an `Events` field on our Lambda resource and set the path and method (`GET`).\
we run the same commands as before to transform the template, upload it to an S3 bucket, and run the deploy command. when we add the API part, it creates many more resources for us.

adding a dynamoDB isn't much more complicated, we modify the function again to use the "boto3" client, and add a resource for the table. we also set the environment variables and the <cloud>IAM</cloud> role and policies. this can all be done in a few lines!

### SAM Policies

> List of templates to apply permissions to
your Lambda Functions.\
> Full list available here: https://docs.aws.amazon.com/serverlessapplicationmodel/latest/developerguide/serverlesspolicy-templates.html#serverless-policytemplate-table
>
> Important examples:
> - S3ReadPolicy: Gives read only permissions to
objects in S3
> - SQSPollerPolicy: Allows to poll an SQS queue
> - DynamoDBCrudPolicy: CRUD = create read
update delete

### SAM Code Deploy

integrating <cloud>CodeDeploy</cloud> and <cloud>SAM</cloud>.

> SAM framework natively uses CodeDeploy to update Lambda functions
> - Traffic Shifting feature
> - Pre and Post traffic hooks features to validate deployment (before the traffic shift starts and after it ends)
> - Easy & automated rollback using <cloud>CloudWatch</cloud> Alarms.

the template has some new fields, such as "AutoPublishAlias" and "DeploymentPreference" (which has alarms and hooks).

(demo of deploying)\
Going over the template file and the traffic shift in codeDeploy.
### SAM Local Capabilities

> Locally start AWS Lambda. `sam local start-lambda`\
> Starts a local endpoint that emulates AWS Lambda. Can run automated tests against this local endpoint.
>
> Locally Invoke Lambda Function. `sam local invoke`.\
> Invoke Lambda function with payload once and quit after invocation completes
> - Helpful for generating test cases
> - If the function make API calls to AWS, make sure you are using the correct `--profile` option
>
> Locally Start an API Gateway Endpoint `sam local start-api`.\
>  Starts a local HTTP server that hosts all your functions. Changes to functions are automatically reloaded.
> 
> Generate AWS Events for Lambda Functions `sam local generate-event`.\
> Generate sample payloads for event sources
> - <cloud>S3</cloud>
> - <cloud>API Gateway</cloud>
> - <cloud>SNS</cloud>
> - <cloud>Kinesis</cloud>
> - <cloud>DynamoDB</cloud>

### Summary
> - SAM is built on CloudFormation
> - SAM requires the Transform and Resources sections
> - Commands to know:
>   - `sam build`: fetch dependencies and create local deployment artifacts
>   - `sam package`: package and upload to Amazon S3, generate CF template
>   - `sam deploy`: deploy to CloudFormation
> - SAM Policy templates for easy IAM policy definition
> - SAM is integrated with CodeDeploy to do deploy to Lambda aliases.
> 
</details>

## Cloud Development Kit
<details>
<summary>
Infrastructure as code.
</summary>

defining the cloud infra with a programming language. libraries that support creating AWS resources. the componenets are called "constructs", and it's compiled into a cloud formation template.

we can deploy the resource together with the code. we use the cli `cdk synth` to make the template.

| Command                      | Description                                        |
| ---------------------------- | -------------------------------------------------- |
| `npm install -g aws-cdk-lib` | Install the CDK CLI and libraries                  |
| `cdk init app`               | Create a new CDK project from a specified template |
| `cdk synth`                  | Synthesizes and prints the CloudFormation template |
| `cdk ls`                     | (list something)                                   |
| `cdk bootstrap`              | Deploys the CDK Toolkit staging Stack              |
| `cdk deploy`                 | Deploy the Stack(s)                                |
| `cdk diff`                   | View differences of local CDK and deployed Stack   |
`cdk destroy` Destroy the Stack(s)

### Hands on Demo

we install the cdk from npm, and we initialize a new app with `cdk init` and tell which language we want to use. we copy the file from the resources and look at it. it's a javascript file that defines the resources (constructs).\
`cdk bootstrap` provides the cdk access to our account - we only need to this once per region per account. `cdk synth` creates the cloud formation template. `cdk deploy` creates the stack, and `cdk destroy` removes it. we need to empty the <cloud>S3 bucket</cloud> before removing the stack.

### Constructs

> CDK Construct is a component that encapsulates everything CDK needs to create the final CloudFormation stack.\
> Can represent a single AWS resource (e.g., S3 bucket) or multiple related resources (e.g., worker queue with compute).
> 
> AWS Construct Library
> - A collection of Constructs included in AWS CDK which contains Constructs for every AWS resource
> - Contains 3 different levels of Constructs available (L1, L2, L3)
> - Construct Hub – contains additional Constructs from AWS, 3rd parties, and open-source CDK community
>
> Layer 1 Constructs (L1)
> 
> Can be called CFN Resources which represents all resources directly
available in CloudFormation.\
> Constructs are periodically generated from CloudFormation Resource Specification.
> - Construct names start with Cfn (e.g., `CfnBucket`)
> - You must explicitly configure all resource properties.
>
> Layer 2 Constructs (L2)
> 
> Represents AWS resources but with a higher level (intent-based API).\
> Similar functionality as L1 but with convenient defaults and boilerplate. You don't need to know all the details about the resource properties.\
>  Provide methods that make it simpler to work with the resource (e.g., `bucket.addLifeCycleRule()`)
>
> Layer 3 Constructs (L3)
> 
> Can be called Patterns, which represents multiple related resources.\
> Helps you complete common tasks in AWS
> 
>Examples:
> - `aws-apigateway.LambdaRestApi` represents an API Gateway backed by a Lambda function
> - `aws-ecs-patterns.applicationLoadBalancerFargateService` which represents an architecture that includes a Fargate cluster with Application Load Balancer.

### Unit Testing

for unit testing, we can get the CDK Assertion Module to make sure the code we write does what we thing it does. like checking if the template creates a lambda of a specific runtime. we can also check this against a defined <cloud>cloudFormation</cloud> (either in string form or from an actual resource).

</details>

## Amazon Cognito
<details>
<summary>
Give users an identity to interact with our web or mobile application.
</summary>

> Give users an identity to interact with our web or mobile application.
> 
> Cognito User Pools:
> - Sign in functionality for app users
> - Integrate with API Gateway & Application Load Balancer
> 
> Cognito Identity Pools (Federated Identity):
> 
> - Provide AWS credentials to users so they can access AWS resources directly
> - Integrate with Cognito User Pools as an identity provider
>
> Cognito vs IAM: “hundreds of users", "mobile users", “authenticate with SAML"

### Cognito User Pools

> Create a serverless database of user for your web & mobile apps.
> - Simple login: Username (or email) / password combination
> - Password reset
> - Email & Phone Number Verification
> - Multi-factor authentication (MFA)
> - Federated Identities: users from Facebook, Google, SAML...
> - Feature: block users if their credentials are compromised elsewhere
> - Login sends back a JSON Web Token (JWT)

when the user authenticate against <cloud>Cognito</cloud>, they get back a JWT. this token is base64 encoded and contains sections
- header
- payload - user information
- signature - needs to be verified before trusting the payload

intgrates with <cloud>API Gateway</cloud> and <cloud>Application Load Balancer</cloud>, the users pass the token in the request.

we select the requirements for passwords, expiration, MFA, account recovery, registration and sign-in options. we can use the cognito email address for up to 50 emails a day, or use <cloud>SES</cloud>. there is a built-in UI to log in, and we choose a domain if we have one. for this, we must have an <cloud>ACM</cloud> certificate in us-east-1 region.\
we can connect our application to some clients. the callback url is where the users are sent after the log-in. if we want, we can customize the UI.\
we can trigger flow via <cloud>Lambda</cloud> to user events like sign-up, log-in, migration...

there is a feature called "adaptive authentication", which assigns a risk score to each sign-in attempt and can require MFA. this is based on device, location, address.
### Application Load Balancer - Authenticate Users

> Your Application Load Balancer can securely authenticate users.\
> Offload the work of authenticating users to your load balancer, Your applications can focus on their business logic.
> 
> Authenticate users through:
> - Identity Provider (IdP): OpenID Connect (OIDC) compliant
> - Cognito User Pools:
> - Social IdPs, such as Amazon, Facebook, or Google
> - Corporate identities using SAML, LDAP, or Microsoft AD
> 
> Must use an **HTTPS** listener to set `authenticate-oidc` & `authenticate-cognito` rules.\
> `OnUnauthenticatedRequest` – authenticate (default), deny, allow.

there are some step needed to use OIDC authentication.

### Cognito Identity Pools (Federated Identity)

> Get identities for "users" so they obtain temporary AWS credentials.
> 
>  Your identity pool (e.g identity source) can include:
> - Public Providers (Login with Amazon, Facebook, Google, Apple)
> - Users in an Amazon Cognito user pool
> - OpenID Connect Providers & SAML Identity Providers
> - Developer Authenticated Identities (custom login server)
> - Cognito Identity Pools allow for unauthenticated (guest) access
> 
> Users can then access AWS services directly or through API Gateway. The IAM policies applied to the credentials are defined in Cognito, They can be customized based on the user_id for fine grained control.

`${cognito-identity.amazonaws.com:sub}`

the users assume a role once the authenticate in AWS.

### Summary

> - Cognito User Pools (for authentication = identity verification)
> - Database of users for your web and mobile application
> - Allows to federate logins through Public Social, OIDC, SAML...
> - Can customize the hosted UI for authentication (including the logo)
> - Has triggers with AWS Lambda during the authentication flow
> - Adapt the sign-in experience to different risk levels (MFA, adaptive authentication, etc…)
> 
> - Cognito Identity Pools (for authorization = access control)
> - Obtain AWS credentials for your users
> - Users can login through Public Social, OIDC, SAML & Cognito User Pools
> - Users can be unauthenticated (guests)
> - Users are mapped to IAM roles & policies, can leverage policy variables
> - CUP + CIP = authentication + authorization


</details>

## Other Services
<details>
<summary>
Additional services
</summary>

### Step Function

model a workflow as a state-machine, written in json, visualized as a graphical entity. workflows can start manually or from an <cloud>Event Bridge</cloud>.

Has Tasks and States.

Tasks:

> Invoke one AWS service:
>   - Can invoke a <cloud>Lambda</cloud> function
>   - Run an <cloud>AWS Batch</cloud> job
>   - Run an <cloud>ECS task</cloud> and wait for it to complete
>   - Insert an item from <cloud>DynamoDB</cloud>
>   - Publish message to <cloud>SNS</cloud>, <cloud>SQS</cloud>
>   - Launch another <cloud>Step Function</cloud> workflow...
> 
> Run one Activity
> - <cloud>EC2</cloud>, <cloud>Amazon ECS</cloud>, on-premises
> - Activities poll the Step functions for work
> - Activities send results back to Step Functions

States:
> - Choice State - Test for a condition to send to a branch (or default branch)
> - Fail or Succeed State - Stop execution with failure or success
> - Pass State - Simply pass its input to its output or inject some fixed data, without performing work.
> - Wait State - Provide a delay for a certain amount of time or until a specified time/date.
> - Map State - Dynamically iterate steps.
> - Parallel State - Begin parallel branches of execution.

demo of creating a state machine, using the "hello world" example. there is a visual editor to create a flow. we can also use some predefined templates. for a lambda function step, we can define how the output (result) of the function controls how we move to the next state, and define how it behaves in case of errors.

step functions should control how errors are handled, either if there is a problem inside one of the states, or from the step functions flows. we can either "retry" running the task, or "catch" it and move the an failure path.

> Predefined error codes:
> - States.ALL : matches any error name
> - States.Timeout: Task ran longer than TimeoutSeconds or no heartbeat received
> - States.TaskFailed: execution failure
> - States.Permissions: insufficient privileges to execute code

we can define how many times we retry, on which errors and with what backoff rate. if we run out of retries, we move into the catch state. the catch is defined similarly (based on the error), but it directs to other states.\
**When we handle the error logic in the state machine flow, we make the application logic much simpler.**



> Step Functions – Wait for Task Token
> - Allows you to pause Step Functions during a Task until a Task Token is returned
> - Task might wait for other AWS services, human approval, 3rd party integration, call legacy systems...
> - Append `.waitForTaskToken` to the Resource field to tell Step Functions to wait for the Task Token to be returned
> - Task will pause until it receives that Task Token back with a `SendTaskSuccess` or `SendTaskFailure` API call.
>
> Step Functions – Activity Tasks
> 
> Enables you to have the Task work performed by an Activity Worker
> - Activity Worker apps can be running on <cloud>EC2</cloud>, <cloud>Lambda</cloud>,
mobile device...
> -  Activity Worker poll for a Task using `GetActivityTask` API.
> -  After Activity Worker completes its work, it sends a response of its success/failure using `SendTaskSuccess` or `SendTaskFailure`
> 
> - To keep the Task active:
>   - Configure how long a task can wait by setting `TimeoutSeconds`
>   - Periodically send a heartbeat from your Activity Worker using `SendTaskHeartBeat` within the time you set in `HeartBeatSeconds`
>   - By configuring a long `TimeoutSeconds` and actively sending a heartbeat, Activity Task can wait up to 1 year.

we can use standard step functions, or express model (asynchronous "at-least once", or synchronous "at-most once").

| Metric            | Standard                                           | Express                                                   |
| ----------------- | -------------------------------------------------- | --------------------------------------------------------- |
| Max. Duration     | Up to 1 year                                       | Up to 5 minutes                                           |
| Execution Model   | Exactly-once Execution                             | "at-least once" and "at-most once"
| Execution Rate    | Over 2000 / second                                 | Over 100,000 / second                                     |
| Execution History | Up to 90 days or using <cloud>CloudWatch</cloud>   | <cloud>CloudWatch</cloud> Logs                                           |
| Pricing           | # of State Transitions                             | # of executions, duration, and memory consumption         |
| Use cases         | Non-idempotent actions (e.g., processing Payments) | IoT data ingestion, streaming data, mobile app backends.. |

### AppSync
### AWS Amplify
</details>

## Take Away
<details>
<summary>
Stuff to remember
</summary>


shell commands

```sh
aws --version # check cli version
aws configure # configure user
aws iam list-users # show all users in account

# MFA
aws sts get-session-token --serial-number <arn-of-the-mfa-device> --tokencode <code-from-token> --duration-seconds 3600

# S3
aws s3 ls --profile <profile_name>

# ECR
aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com
docker push <aws_account_id>.dkr.ecr.<region>.amazonaws.com/<image_name>:<image_tag>

# CloudWatch
aws cloudwatch put-metric-data --namespace "Usage Metrics" --metric-data file://metric.json
aws cloudwatch put-metric-data --namespace "Usage Metrics" --metric-name Buffers --unit Bytes --value 23143433 --dimensions InstanceId=1-234, InstanceType=m1.small
aws cloudwatch set-alarm-state --alarm-name "MyAlarm" --state-value ALARM --state-reason "testing purposes"

# Kinesis
## producer
aws kinesis put-record --stream-name test --partition-key user --data "user sign up"--cli-binary-format raw-in-base64-out
## consumer
aws kinesis describe-stream --stream-name test --partition-key user --data "user sign up"--cli-binary-format raw-in-base64-out
aws kinesis get-shard-iterator --stream-name test --shard-id shardId-00000 --shard-iterator-type TRIM_HORIZON
aws kinesis getrecords --shard-iterator <from the previous message>

# Lambda
aws lambda list-functions
aws lambda invoke --function-name hello-world  --cli-binary-format raw-in-base64-out --payload '{"key1": "value1", "key2": "value2"}'
aws lambda invoke --function-name hello-world  --cli-binary-format raw-in-base64-out --payload '{"key1": "value1", "key2": "value2"}' --invocation-type Event response.json

# DynamoDB
aws dynamodb update-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --update-expression "SET Price = Price - :discount" --condition-expression "Price > :limit" --expression-attribute-values file://values.json
## conditional delete item
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "attribute_not_exists(Price)"
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "attribute_exists(ProductReviews.OneStart)"
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "(ProductCategory IN (:cat1, :cat2)) and (Price between :lo and :hi)" --expression-attribute-values file://values.json
aws dynamodb delete-item --table-name ProductionCatalog --key '{ "Id": {"N": "456" } }' --condition-expression "begins_with(Pictures.FrontView, :v_sub)" --expression-attribute-values file://values.json
## scans with expressions
aws dynamodb scan --table-name UserPost --projection-expression "user_id, content"
aws dynamodb scan --table-name UserPost --filter-expression "user_id= :u" --expression-attribute-values '{":u":{"S":"john123"}}'
aws dynamodb scan --table-name UserPost --filter-expression "user_id= :u" --expression-attribute-values '{":u":{"S":"john123"}}'
aws dynamodb scan --table-name UserPost --page-size 1
aws dynamodb scan --table-name UserPost --max-items 100
aws dynamodb scan --table-name UserPost --max-items 100 --starting-token <nextToken>

# sam
sam init
sam package
sam build
sam deploy --guided
sam local start-lambda
sam local invoke
sam local start-api
sam local generate-event
```

1. if we want really high IOPS (More than 250,000), we have to use Instance Store (can't use ELB).
2. only NLB can have an elastic IP address.
3. "A Read Replica in a different AWS Region than the source database can be used as a standby database and promoted to become the new production database in case of a regional disruption. So, we'll have a highly available (because of Multi-AZ) RDS DB Instance in the destination AWS Region with both read and write available."
4. "What is the maximum number of Read Replicas you can add in an ElastiCache Redis Cluster with Cluster-Mode Disabled?" - **5** 
5. "Elastic Beanstalk application versions can be deployed to **Many "Environments**
6. "Your deployments on Elastic Beanstalk have been painfully slow. After checking the logs, you realize that this is due to the fact that your application dependencies are resolved on each instance each time you deploy. What can you do to speed up the deployment process with minimal impact?" - **Resolve the dependencies beforehand and package them in the zip file uploaded to Elastic Beanstalk**.
7. CloudFormation Stacks
   1. Cross Stack - different lifecycles, import and export
   2. Nested Stack - same lifecycle, reusable components
   3. StackSet - a master stack in the administrator account, controls stacks in multiple regions, accounts.
8. "If you set an alarm on a high-resolution metric, you can specify a high-resolution alarm with a period of 10 seconds or 30 seconds, or you can set a regular alarm with a period of any multiple of 60 seconds." - **If you set an alarm on a high-resolution metric, you can specify a high-resolution alarm with a period of 10 seconds or 30 seconds, or you can set a regular alarm with a period of any multiple of 60 seconds.**
9.  `PartiQL` - SQL-compatible query language for DynamoDB.
10. <cloud>DynamoDB DAX</cloud> - built-in accelarator, in memory-cache.

<cloud>CoPilot</cloud> manages ECS applications, while <cloud>Beanstalk</cloud> manages instance-based applications.
</details>