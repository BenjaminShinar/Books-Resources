<!--
// cSpell:ignore boto xlarge POSIX Proto AWSELB AWSALBTG AWSALBAPP NAPTR NACL DSSE
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
> - Does live "outside” the EC2 - if traffic is blocked the EC2 instance won't see it
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

> EBS volumes…
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

> - Domain Registrar: Amazon Route 53, GoDaddy, …
> - DNS Records: A, AAAA, CNAME, NS, …
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
- beanstalk
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
> - Should create a “Default” record (in case there's no match on location)
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

> - Amazon S3 allows people to store objects (files) in “buckets” (directories)
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
>   - Server-Side Encryption with Amazon S3-Managed Keys (SSE-S3) – (Enabled by Default) - Encrypts S3 objects using keys handled, managed, and owned by AWS. AES-256 encryption type.
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
> - It’s a popular exam question
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
> - Note: some requests to Amazon S3 don’t need to be signed
> - If you use the SDK or CLI, the HTTP requests are signed for you
> - You should sign an AWS HTTP request using Signature v4 (SigV4)

we can send the toke in authorization header or the query string as part of pre-signed URL (X-AMS-Signature).
</details>

## CloudFront
<!-- <details> -->
<summary>
Content Delivery Network
</summary>

### Cache
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
```

1. if we want really high IOPS (More than 250,000), we have to use Instance Store (can't use ELB).
3. only NLB can have an elastic IP address.
4. "A Read Replica in a different AWS Region than the source database can be used as a standby database and promoted to become the new production database in case of a regional disruption. So, we'll have a highly available (because of Multi-AZ) RDS DB Instance in the destination AWS Region with both read and write available."
5. "What is the maximum number of Read Replicas you can add in an ElastiCache Redis Cluster with Cluster-Mode Disabled?" - **5** 

</details>