<!--
// cSpell:ignore Udemy boto xlarge
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css"> 

# AZ-900 Ultimate AWS Certified Developer Associate 2024 NEW DVA-C02

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

> - Don’t use the root account except for AWS account setup
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

AM Section – Summary
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
> - Does live "outside” the EC2 – if traffic is blocked the EC2 instance won’t see it
> - It’s good to maintain one separate security group for SSH access
> - If your application is not accessible (time out), then it’s a security group issue
> - If your application gives a "connection refused" error, then it’s an application error or it’s not launched
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

> - On-Demand Instances – short workload, predictable pricing, pay by second
> - Reserved (1 & 3 years)
>   - Reserved Instances – long workloads
>   - Convertible Reserved Instances – long workloads with flexible instances
> - Savings Plans (1 & 3 years) – commitment to an amount of usage, long workload
> - Spot Instances – short workloads, cheap, can lose instances (less reliable)
> - Dedicated Hosts – **book an entire physical server**, control instance placement
> - Dedicated Instances – **no other customers will share your hardware**
> - Capacity Reservations – reserve capacity in a specific AZ for any duration

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
> - Capacity Reservations: you book a room for a period with full price even you don’t stay in it


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
Amazon Machine Image - AMI, they are customizations of an EC2 instance. it has additional software, configuration, monitoring right out of the box, which saves time on configuration and booting.

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

EBS volumes are network storage, but if we want something with higher performance, we can use storage that is directly attached to it. this gives us better speed, but it goes away when the machine is stopped (not even terminated), this is good for workloads that need cache/scratch data/memory buffer. instance store have much higher throughput than EBS, even IOPS optimized.

### EFS

### Summary

</details>

## Take Away

shell commands

```sh
aws --version # check cli version
aws configure # configure user
aws iam list-users # show all users in account
```