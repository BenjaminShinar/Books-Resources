<!--
// cSpell:ignore chkconfig Gbps
 -->

[main](README.md)

## Section 4 - EC2

<details>
<summary>
Elastic Compute Cloud, Elastic Block Storage
</summary>

### EC2 Basics

<details>
<summary>
First look at EC2.
</summary>

#### EC2 101

> Amazon Elastic Compute Cloud (Amazon EC2) is a web service that provides resizable compute capacity in the cloud. Amazon EC2 reduces the time required to to obtain and boot new server instances to minutes, allowing you to quickly scale capacity, both up ad down, as your computing requirements change.

pricing:

1. **On Demand**: allows you to pay a fixed rate by the our (or by the second) with no commitment.
   - users who want the low cost and flexibility of amazon ec2 without up-front payment or long-term commitment.
   - applications with short term, spiky or unpredictable workloads that cannot be interrupted.
   - applications that are being developed or tested.
1. **Reserved**: capacity reservation, significant discount on the hourly charge for an instance. contract terms are either one or three years terms.
   - applications with steady state or predictable usage
   - applications that require reserved capacity
   - users are able to make upfront payments to reduce their total computing costs even further.
   - Comes in 3 types: (costs are compared to on demand pricing)
     - Standard Reserved instances: up to 75% discount, based on upfront payment and contract duration.
     - Convertible Reserved instances: up to 54% discount. we can change the attributes of the instances as we want as long as we keep the value. we can get more ram if needed.
     - Scheduled Reserved instances: available to launch at a time window, matchin capacity to a predictable schedule (hourly,daily, weekly,monthly) without paying for the off-time.
1. **Spot**: enables you to bid whatever price you want for instance capacity, even greater savings if your applications have a flexible start and end times. price varies based on supply and demand.
   - applications with flexible start and end time
   - application that are only feasible at very low compute prices.
   - users with urgent computing needs for large amount of additional capcity (and are willing to pay immediately)
   - if the spot instance is terminated bt Amazon, we will not pay for the partial hour of usage, but if we terminate it ourselves, we will pay for the charges in that hour time frame.
1. **Dedicated** Hosts: physical EC2 server dedicated for your use. can reduce costs by allowing use of existing server-bound software licenses (like oracle), and might be required by regulations.
   - useful for regulatory requirements that may not support multi-tenant virtualization.
   - great for licensing which does not support multi-tenancy or cloud deployments.
   - can be purchased on-damand (hourly)
   - can be purchased as a reservation for up to 70% off the on-demand price.

available instances types. we don't need to know everything for now. the number doesn't really matter, it's just the generation.
| Family | Specialty | Use Case |
| ------ | ------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| F1 | Field Programmable Gate Array | genomics research, financial analytics, real time video processing, big data |
| I3 | High Speed Storage | NoSQL databases, Data Warehousing |
| G3 | Graphics Intensive | Video Encoding, 3D application streaming |
| H1 | High Disk Throughput | Map-Reduce based workloads, distributed file systems such as HDFS and MapR-FS |
| T3 | Lowest Cost, General Purpose | Web servers, small databases |
| D2 | Dense Storage | Fileservers, Data Warehousing, Hadoop |
| R5 | Memory optimized | Memory intensive applications/ databases |
| M5 | General Purpose | Application servers |
| C5 | Compute Optimized | CPU intensive applications/ databases |
| P3 | Graphics / General Purpose GPU | Machine learning, Bitcoin minning |
| X1 | Memory Optimized | SAP HANA/ Apache Spark |
| Z1D | High compute capacity and a high memory footprint | Ideal for Electronic Design Automation (EDA) and certain relation databases workloads with high per-core licensing costs |
| A1 | ARM-based workloads | Scale-out workloads such as web servers |
| U-6tb1 | Bare Metal | Bare metal capabilities that eliminate virtualization overhead. |

> F - For FPGA \
> I - For IOPS \
> G - Graphics \
> H - High disk throughput \
> T - Cheap general purpose (think T2.micro) \
> D - For Density \
> R - For RAM \
> M - Main choice for general purpose apps \
> C - FOr Compute \
> P - Graphics (think Pictures) \
> X - Extreme memory \
> Z - Extreme memory and CPU \
> A - ARM based workloads \
> U - Bare metal
>
> FIGHT DR MC.PXZ AU - fight dr mcPixie in Australia

Summary:

> - Amazon Elastic Compute Cloud (EC2) is a web service that provides resizable compute capacity in the cloud. Amazon EC2 reduces the time required to obtain and boot new server instances to minutes, allwing to quickly scale capacity, both up and down, as your computing requirement change.
> - Pricing types:
> - On demand
> - Reserved
> - Spot
> - Dedicated

#### Let's Get Our Hands Dirty With EC2

we start by getting into the aws console, we choose the region, and under compute, we choose <kbd>EC2</kbd>. in the dashboard, we can see
which resources are running.

we press <kbd>Launch instance</kbd> to start a new EC2 machine. we begin by choosing an AMI - Amazon Machine Image, the ami is the virtual machine, we start by choosing the basic _Amazon linux 2 AMI_. not we choose instance types, and we choose the **t2.micro** instance, because it's part of the free tier.\

now we configure the instance details:

- Number of instaces - 1 is enough for now, we don't need auto scaling
- Purchasing option - we don't want spot instances
- networking
  - network - which VPC
  - subnet - select availability zone
  - auto-assign public ip - enable/disable
- Placement
  - Placement Group - for high prefromance (later)
  - Capacity reservation - (later)
- IAM role - what role to give the ec2
- System stuff
  - Shutdown behavior (stop, terminate)
  - Termination protection (prevent us from accidentally eliminating)
  - Detailed Monitoring with CloudWatch
  - Tenancy (shared, dedicated)
  - Elastic Interface
- T2/T3 unlimited

under the **advanced details** box we can add bootstrap commands as **User Data**

now we move to <kbd>Add Storage</kbd>, we can choose the storage for the root volume (either ssd or magnetic),and for the additional volumes there are even more options. we can choose the size, encryption and prevent it from being deleted on termination. we can add <kbd>Tags</kbd> as we want.

the next step is to configure **security groups**. a security group is a set of firewall rules that control traffic to and from our ec2 instance. we can add specific rules. we can control type, protocol,port ranges, etc. security groups can be shared. the source of `0.0.0.0/0` allows all ip addresses to acccess the instance, so it's not secured at all.

CIDR - Classless Inter-Domain Routing

we start with a security groups with two rules

| Type | Protocol | Port Range | Source                 | Description |
| ---- | -------- | ---------- | ---------------------- | ----------- |
| SSH  | TCP      | 22         | custom 0.0.0.0/0       | ssh access  |
| HTTP | TCP      | 80         | custom 0.0.0.0/0, ::/0 |             |

we can choose to add only our IP (<kbd>My Ip</kbd>), but it would be problematic if our ip changes.

next we choose key-pair for accessing. private and public keys. the public key goes on the ec2 instance, and we keep the private key on our local machine. we download the new key. now we launch the instance and wait for it to start.

we can see all the information for each of our ec2 machines. we can take the public ip address so we could ssh into the machine once it's up. private key files have the _.pem_ suffix.

if we want to ssh, we can use an terminal client from the local machine, but we can also click <kbd>Connect</kbd> and choose "EC2 Instance Connect (browser-based-ssh connection)" with whatever user name we want. now we have a tab with a terminal to the ec2 machine.

in a terminal we can use the private key and the ssh client to connect to the public ip address

```sh
mkdir ssh # create folder
mv myKey.pem SSH # move key into folder
CHMOD 400 ssh/myKey.pem # change permissions of key
ssh ec2-user@<public io address> -i ssh/myKey.pem
#confirm, see that the prompt changed and we are inside the ec2 machine
$ sudo su # change permissions
```

windows users can use an chrome extension to make our browser into a ssh client, **Secure Shell App**. we can install it on our browser and use it as a terminal. in the configuration screen, we need to configure all sorts of stuff, including a _.pub_ (public key). which we need to import together with the private key (without the _.pem_ extension).

```sh
cd ssh #go to the folder
ssh-keygen -y -f myKey.pem > myKey.pub # generate public key
mv myKey.pem myKey # remove extension
```

inside the ssh terminal (connected to the machine), we can start running commands

```sh
yum update -y #get updates
yum install httpd -y #install apache
cd /var/www/html
nano index.html #create a file
# <html><h1>Hello World</h1></html>
# exit nano wit kb
service httpd start # start apache
chkconfig on # make sure it's on for the next time
```

and now we go to the ip address in the browser and we see the html page we created.

back to EC2 dashboard. there several tabs, such as **description**, **status check** and **monitoring**. we can see detailed information, perfrom system status checks and instance status checks, and view monitoring for all sorts of metrics. later we will create metrics of our own.

for each instance, we can click the <kbd>Actions</kbd> button and choose actions like connecting, stop/terminate/start/reboot the instance, if we set up _termination protection_ we won't be able to terminate it.

in the sidebar, we can see spot requests and purchase reserved instances and pay for them.

we launch another ec2 instance, and in the storage step, there the volume types:

- General purpose ssd (gp2)
- Provisioned iops ssd (io1)
- Magnetic (standard)
- Cold HDD (sc1)
- Throughput Optimized HDD (st1)

we can also choose to encrypt the storage device, in the past we couldn't encrypt the root device, but this changed. _delete on termination_ isn't enabled by default for the non-root storage volumes.

> - **Terminatin protection** is off by deafult for ec2 instances
> - On an EBS-backed instance, the default action for the root EBS volume is to be **deleted when the instance is terminated**.
> - EBS root volumes of your default AMI can be encrypted. this can be done by a third party tool or with the aws console or using the cli.

#### Security Groups Basics

in the aws console. we gran the public ip address and view our html (which is provided by the apache server). we opened the web page with the deafult port of 80.

let's look at our security group, we can see the inbound and outbound rules.

in bound rules:

| type | protocol | port range | source    | notes |
| ---- | -------- | ---------- | --------- | ----- |
| HTTP | TCP      | 80         | 0.0.0.0/0 | Ip-V4 |
| HTTP | TCP      | 80         | ::/0      | Ip-V6 |
| SSH  | TCP      | 22         | 0.0.0.0/0 | Ip-V4 |

if we delete the port 80 rules, we can try to access the page again, and we see that this takes effect immediately and the page is not responding anymore.

security groups are stateful, any inbound rules has a coresponding outbound rule. we can't only have inbound port 80 without being able to respond to that request. this will be different for VPC, which are stateless and require setting inbound rules and outbound separately.

also, in security groups, there is no way to block ports or ips, this will come up again in the vpc section.

if we choose types like MS-SQL or MY-SQL, we will get the correct port. we can have more than one security group assigned to an EC2 instance.

> - All inbound traffic is blocked by default.
> - All outbound traffic is allowed.
> - Changes to security groups take effect immediately.
> - You can have any number of EC2 instances within a security group.
> - You can have multiple security groups attached to an EC2 instance.
> - Security groups are statefull.
>   - if you create an inbound rule allowing traffic in, that traffic is automatically allowed back out again.
> - You cannot block specific IP addresses using security Groups, instead use Network Access Control Lists.
>   You can specifically allow rules, but not deny rules.

</details>

### EBS Basics

<details>
<summary>
EBS - Elastic Block Storage
</summary>

#### EBS 101

EBS - elastic Block Store

> Amazon Elastic Block Store (EBS) provides persistent block storage volumes for use with amazon EC2 instances in the AWS cloud. Each Amazon EBS volume is automatically replicated within it's avaliability zone to protect you from componenet failure, offering high availability and durability.

| Option                | Api Name | Type | Volume Size      | max IOPS/volume | Use Cases                                     | Notes                                                                 |
| --------------------- | -------- | ---- | ---------------- | --------------- | --------------------------------------------- | --------------------------------------------------------------------- |
| General purpose       | gp2      | SSD  | 1 Gib - 16 TiB   | 16000           | Most work loads                               | General purpose, balance between cost and performance                 |
| Provision IOPS        | io1      | SSD  | 4 Gib - 16 TiB   | 64000           | Databases                                     | Highest performance, mission-critical applications                    |
| Throughtout optimized | st1      | HDD  | 500 Gib - 16 TiB | 500             | Big data, Data warehouses                     | Low cost hdd, for frequently accessed, throughput intensive workloads |
| Cold Hard Disk Drive  | sc1      | HDD  | 500 Gib - 16 TiB | 250             | File servers                                  | Lowest cost HDD, les frequently accessed workloads                    |
| Magnetic              | standard | HDD  | 1 Gib - 1 TiB    | 40-200          | Workloads where data is infrequently accessed | Previous generation HDD                                               |

#### Volumes & Snapshots

in our aws console, we look at our EC2 instance and choose the volume option in the side bar. we can see where ec2 instance is and where the EBS volume is. they will be in the availability zone. we can move them from one avalability zone to another. if we terminate the instance, the ebs volume will also go away (unless we decide to uncheck the <kbd>Delete on Termination</kbd>).

if we add more volumes to our EC2 instance, we can modify the volumes like changing the size. this won't happen immediately. we might also need to extend the OS file system to tell the machine to see the new allocated space.

only the root volume has something in the snapshot column. we can change the storage type (from gp2 to io1). this will take a few minutes. we can move the volume from one availability zone to another. (it's more a a copy than a move, actually)

we click <kbd>actions</kbd>, select <kbd>Create snapshot</kbd> and wait until it's ready. now, in the snapshots section, we click <kbd>actions</kbd>, select <kbd>Create Image</kbd>.

there are two types of virtualization (For now):

- paravirtual (pv)
- hardware virtual machine (hpv)

some ami configurations support both, usually HPV is safer and works with most ami images. once the image is created, we can use it to launch a new ec2 instance in a different subnet (different avalability zone).
we can also move images between regions, and use the same image ami in a different region.

non root volumes which were attached to the ec2 instance don't get deleted when the machine is termindated.

> - Volumes exist on EBS. think of EBS as a virtual hard disk.
> - Snapshots exists on S3, think of snapshost as photographs of the disk.
> - Snapshots are point in time copies of volumes.
> - Snapshots are incremental, only the additional changes since the last snapshot are stored on S3.
> - If this is the first snapshot of a volume, this might take time to create.
> - To create a snapshot for Amazon EBS volumes that serve as root devices, you should stop the instance before taking the snapshot.
> - It is possible to take a snapshot while instance is running
> - You can create AMI's from snapshots.
> - You can change EBS volume sized on the fly, including changing the size and the storage type.
> - Volumes will always be in the same avalability zone as the EC2 instance.
> - To move an EC2 volume from one AZ to another, take a snapshot of it, crate an AMI from the snapshot and then use the AMI to launch the EC2 in a new AZ.
> - we can copy the AMI between regions, and then we can start the new EC2 instance in the new region.

#### AMI Types (EBS vs Instance Store)

two types of AMI: EBS and Instance store.

when we select amis, we choose based on:

- Region
- Operating system
- Architecture (32-bit, 64-bit)
- Launch permissions
- Storage for the root device
  - EBS backed volumes
  - Instance Store (**Ephemeral Storage**)

so far we used EBS backed volumes.

> All AMIs are categorized as either backed by Amazon EBS or backed by instance store.
>
> EBS Volumes: the root device for an instance launched from the ami is an EBS volume created from an **Amazon EBS snapshot.**
>
> Instance Store Volumes: the root device for an instance launched from the ami is an EBS volume created from a **template stored in amazon S3**.

we go the console, launch one instance from the ebs backed volume as usual. for a different instance, we launch an instance from an instance store, so we select <kbd>Community AMIs</kbd>, and filter based on root device type and find the default instance store ami. for this image, we are limited to which machine instance type we can use. and for those instances we cannot change the volume storage type. we won't be able to see the instance store volume in the volumes list, and we cannot stop it, we can only terminate or delete it.

for instance based on EBS, stopping and starting again will load it on a different hypervisor. for instance store volumes, we cannot do it, the data is ephemeral.

> - Instance store volumes are sometimes called ephemeral storage.
> - Instance store volumes cannot be be stopped, if the underlying host fails, you will lose the data..
> - You can reboot both, and you will not lose your data.
> - By default, both ROOT volumes will be deleted on termination, however, with EBS volumes, you can tell AWS to keep the root device volume.

#### ENI vs ENA vs EFA [SAA-C02]

> - ENI - Elastic Network Interface - virtual network card.
> - EN - Enhanced Networking - uses single root I/O virtualization (SR-IOV) for better performan.
>   - ENA - Elastic Network Adaptor - a way to enable EN.
> - EFA - Elastic Fabric Adaptor - attach EC2 to accelerate High Performance Computing (HPC) and machine learning capabilities.

An ENI is a virtual network card on the EC2, it allows:

- A primary private IPv4 address from the IPv4 address range of the the VPC.
- One or more secondary private IPv4 addresses from the IPv4 address range of the the VPC.
- One public IPv4 address.
- One or more public IPv6 addresses.
- One or more Security groups.
- A MAC address.
- A source/destination check flag.
- A description.

> Scenarios for Network interfaces:
>
> - Create a management network.
> - Use network and security appliances.
> - Create a dual-homed instance with workloads/roles on distinct subnets.
> - Create a low-budget, high-avalability solution.

when ENI aren't enough, we can use Enhanced networking, SR-IOV is a method of device virtualization that provides better I/O performance and lower CPU utilization.\
we get higher bandwidth, high packet per seconds rate (PPS) and lower latency. there are no additional charges. we should use this where we want good network performance.\
not all instance types support EN. it can be enabled by using an **Elastic Network Adaptor** for network speeds of up to 100 Gbps or with _intel 82599 virtual function_ interace, which is older and supports speeds of up to 10 Gbps.

an ELastic Fabric Adapter can be attached to an EC2 instance to acclarate High Performace Computing, it has lower and more consistent latency and higher throughput than TCP transport. it can also use OS-bypass, which allows the machine learning applications to communicate directly with the EFA device without going through the operating system kernel. this is currently not supported in windows, only linux.

> Possible scenario questions:
>
> **ENI**\
> For basic networking. Perhaps you need a separate management network to your production network or a separate logging network and you need to do this at low cost. in this scenario use multiple ENIs for each network.
>
> **Enhanced Network**\
> For when you need speeds between 10 Gbps and 100Gbps. anywhere you need reliable, high throughput.
>
> **Elastic Fabric Adaptor**\
> For when you need to accelerate High Performance Computing (HPC) and machine learning application, or if you need to do an OS bypass. if you see a scenario question mentioning HPC or ML and asking what network adaptor you want, choose EFA.

#### Encrypted Root Device Volumes & Snapshots

The root device volume is basically the disk space with the operating system, in the past, we would need to create and snapshot and encrypt that snapshot, but now we can provision enctypred root device volumes directly.

in the console, we open the <kbd>EC2</kbd> and create a t2.micro machine, and in the storage page, we can add the encryption at creation. but if we don't encrypt it now, we can do so later. we add it to proper security group and start it.

now, under the **Elastic Block Store** option in the side bar, we can se that this volume isn't encrypted. <kbd>Actions</kbd>, <kbd>Create Snapshot</kbd>, and once it's live, we can click it, <kbd>Actions</kbd>, <kbd>Copy</kbd> and check the **Encryption** box. once it's done coyping, we click <kbd>Actions</kbd> again, and choose <kbd>Create Image</kbd> to create an encrypted AMI.

we could use this ami to create a new EC2 instance.

(this process can come up in the exam, but it isn't as popular as it once was).

> - Snapshots of enctyped volumes are encrypted automatically.
> - Volumes restoed from enctyped snapshots are encrypted automatically.
> - You can share snapshots, but only if they are un-encrypted.
> - These snapshots can be shared with other AWS accounts or made public.
> - It is now possible to encrypt root device volumes upon creation of the EC2 instance.
> - The process to make an un-encrypted machine into an encrypted root device one is:
>   - Create a snapshot of the un-encrypted root device volume.
>   - Create a copy of the snapshot and select the encrypt option.
>   - Create and AMI from the encrypted snapshot
>   - Use the AMI to create a new EC2 machine.

</details>

### Spot Instances & Spot Fleets [SAA-C02]

> **Amazon EC2 Spot Instances** let you take advantage of unused EC2 capacity in the AWS Cloud. Spot Instances are available at up to 90% discount compared to On-Demand prices. You can use **Spot Instances** for various stateless, fault tolerant or flexible applications, such as big data, containerized workloads, CI/CD, web servers, high-performance computing (HPC) and other test and development workloads.

to use spot instances,we decide on the maximum **spot price**, which is the price we are willing to pay. as long as the asking price is below this amount, the EC2 instance will be provisioned to us.

The hourly spot price varies depending on capacity and region, if the price goes above the maximum spot price we marked, we have **two minutes** to choose whether to stop or terminate the instances.

we can prevent spot instances from being terminated by activating **spot block**. this means that even if the price goes above the spot price, we will still continue with the job, this can be done in blocks between one and six hours. in the aws console, we can see the price fluctuations.

Spot instances aren't suitable for:

- Persistent workloads
- Critical jobs
- Databases

however,if our application can cope with sudden stops and terminations, it might be suitable.

the process:

1. we start by creating a request, this consists of the maximum spot price, the desired number of instances, launch specifications, the type of the request (one-time | persistent), and the valid from and valid until times.
2. if the current price is above the spot price, the request will fail. if not, it will launch the instance.
3. if the prices rises above the spot prices and we chose the the one-time request type, then our instance is stopped and terminated.
4. if the request type was persistent, then the instance is stopped, but once the price falls below the spot price, the instance will begin again.

there is a flow of how the request works in the amazon documentation.

> **Spot Fleets**\
> A Spot Fleet is a collection of Spot Instances and, optionally, On-Demand Instances.
>
> The Spot Fleet attempts to launch the number of Spot Instances and On-Demand Instance to meet the target capacity you specified in the Spot Fleet requests. The request for the Spot Instances is fulfilled if there is available capacity and the **maximum price you specified in the request exceeds the current Spot price**.\
> The Spot Fleet also attempts to maintain its target capacity fleet if your Spot Instances are interrupted.

Spot fleets will try and match the target capacity with you price restanits

- Different launce pools, like EC2 instance types, operating systems and availability zones.
- We can have multiple pools, and the fleet will choose the best way to implemenate the request based on the chosen strategy.
  - **capacityOptimized** - The Spot instances come from the pool with optimal capacity for the number of instances launching.
  - **diversified** - The Spot instances are distributed across all pools.
  - **lowestPrice** - The Spot instances come from the pool with the lowest price, default strategy.
    - **InstancePoolsToUseCount** - The Spot instances are distributed across the number of spot instance pools you specify, this is only valid when used in combination with **lowestPrice**.
- Spot fleets will stop launching instaces once the desired capacity is reached of the the price is above the price threshold.

> - Spot instances save up to 90% of the cost of On-Demand instances.
> - Useful for anytype of computing where you don't need persistent storage
> - You can use **Spot block** to stop running spot instances from terminating.
> - A Spot Fleet is a collection of Spot instances and, (optionally), On-Demand instances.

### EC2 Hibernate [SAA-C02]

for EC2 instances, we can **stop** or **terminate** them.

- **Stop** - the data is kept on the disk (EBS) and will remain there until the EC2 instance is started again.
- **Terminate** - by default, the root device volume will also be terminated. (there is an option to keep the root device)

when we start the ec2 instance

- Operating system boots up
- User data script is run (_boostreap scripts_)
- The application starts

> When you hibernate and EC2 instance, the operating system is told to perform hibernation (suspend-to-disk). Hibernation **saves the contents** from the instacne memory (RAM) to your amazon EBS root volume. We persist the instance Amazon EBS root volume and any attached EBS data volumes.
>
> When we start the instace out of hibernation
>
> - The Amazon EBS root is restored to its' previous state.
> - The RAM contnets are reloaded.
> - The processes that were previously running on the instance are resumed.
> - Previously attached data volumes are **reattached and the instance retains its instance ID**.

with hibernate, the instance boots much faster, the operating system doesn't need to reboot because the in-memory (RAM) state is preserved. this is usefull for long-running process and services that take a time to initialize.

in the AWS Console, <kbd>EC2</kbd>, we start with a normal AMI, and in the configuration step, there is an _Stop - Hibernate behavior_ checkbox.

- root volume must be large enough.
- root volume must be an encrypted EBS volume.

we ssh into the machine, and run the `uptime` command. then we choose <kbd>Actions</kbd> for the instance, <kbd>Instance State-> Stop - Hibernate</kbd>. and then we start it again by clicking <kbd>Actions</kbd> and choosing <kbd>Instance State -> Start</kbd>. we ssh into it again, and run the `uptime` again, and we see that it doesn't count as a re-boot.

> - EC2 Hibernate preserve the in-memory RAM on persistent storage (EBS).
> - Much faster to bot up because you do no't need to reload the operating system
> - Instance RAM must be **less than 150GB**.
> - Instance families include C3, C4, C5, M3,M4,M5, R3,R4,R5. (M - general purpose, C - compute optimized, R - memory optimized).
> - Available for WIndows, Amazon linux 2 AMI, and Ubuntu.
> - Instances cant be hibernated for more than **60 days**.
> - Available for **On-Demand Instances** and **Reserved Instances**.

### CloudWatch 101

> AWS CloudWatch is a Monitoring service to monitor your AWS resources, as well as the application that you run on AWS.

it monitors performacne, things such as:

- Compute
  - EC2 Instances
  - AutoScaling Groups
  - Elastic Load Balancers
  - Route53 Health Checks
- Storage & Contnt Delivery
  - EBS Volumes
  - Storage Gateways
  - CloudFront

Host level metrics:

- CPU
- Network
- Disk
- Status Check
- Hypervisor

not to be confused with Cloud Trail

> AWS CloudTrail increases visibility into your yser and resource activity by recording AWS management console actions and API calls. you can Identify which users and account called AWS, the source IP address from which the calls were made, and the the calls occurred.

CloudWatch knows how many EC2 machines are running, cloud trail knows if they were provisioned by the same user.

> - CloudWatch is used for monitoring performance.
> - CloudWatch can monitor most of AWS as well as your applications that run on AWS.
> - CloudWatch with EC2 will monitor event every 5 minutes by deafult.
> - You can have 1 minute intervals by turning on detailied monitoring.
> - You can create CloudWatch alarms which trigger notifications.
> - CloudWatch is all about performace, Cloud Trail is all about audting.

#### CloudWatch Lab

in the AWS console. we provision an EC2 instance. in the configuration step, we select the _Enable CloudWatch detailed monitoring (additional charges apply)_ checkbox. the rest is default.

when we look at the instance, we have the **monitoring tab**, and for now we have the basic metrics. in this demo, we will max out the cpu of the instance, and we will want to receive an alert about it.

to create the alert, we find the <kbd>CloudWatch</kbd> service, (under the _Management and Governance_ group), then we select **Alarms** and click <kbd>Create New Alarm</kbd>, now we can view the metrics for our running services, we find the EC2 instance, and choose the _Per-Instance Metrics_, and then we find the suitable metric for the EC2 instance we care about. the metric is _CPU Utilization_. so we click <kbd>Select metric</kbd> and we now need to configure it.

we give the alarm a name and a description, and choose that the threshold and window (how many data points must pass this threshold to trigger the alarm).
we can also choose how to treat missing data, and we can control the action to take when the alarm is triggered.

- send a notification (email)
- do some auto scaling stuff
- do an EC2 action

now we want to make the alarm trigger, so we ssh into the machine and start running an endless loop

```sh
ssh ec2-user@12.123.12.123 -i key.pem
$ sudo su
$ while true; do echo; done #infinite loop
```

now we should get an alarm in the email, and we will see the alarm state in the EC2 Instance monitoring tab.

we can also create dashboards to monitor alarms (global or regional), we can also monitor logs or monitor aws events.

> - Standard Monitoring : 5 minutes.
> - Detailed Monitoring : 1 minute.
> - Dashboards - Create dashboards to see what is happening with your aws environment
> - Alarms - Allows you to set Alarms that notify you when particular thresholds are hit.
> - Events - CloudWatch Events Helps you to respond to state changes in your aws Resources.
> - Logs - CloudWatch Logs Helps you to aggregate, monitor and store logs.
> - CloudWatch monitors performance.
> - CloudTrail monitor API calls in the AWS Platform.

### The AWS Command Line

AWS command line tool. interact directly with AWS from the terminal without using the aws console.

in order to use the AWS CLI, we need an user with programattic access. so we can create a user in the IAM. this user will have an **Access Key Id** and **Secret Access Key**. we can only get the key once, but we can make it inactive and get another one.

the command line tool is global.

if we create a new ssh key, we need to put it in the corret place and set the proper permissions

```sh
mv Key.pem ssh #move to ssh folder
chmod 400 ssh/Key.pem # change permissions
```

we can either run the aws cli commands from an EC2 instance or from the terminal locally.

```sh
aws configure # fill in the access key Id, secret Access Key, default region and output format.
aws s3 ls #list buckets
aws s3 mb s3://testbucket #make bucket

cd ~
ls -a # list hidden files
cd .aws
ls
#config, credentials
cat credentials
```

if we run the aws from the EC2 instances, then someone can pull the keys from the configuration folder, so it's more secure to use roles instead.

> - You can interact with aws from anywhere in the world just by using the command line interface (CLI).
> - You will need to set up access in IAM.
> - The commands themselves aren't in the exam, but some basic commands will be useful to know for real life situations.

#### Using IAM Roles With EC2

rather than using the secret key directly, we can use roles instead.

in the aws console, we go to <kbd>IAM</kbd>, under the roles option, we can see the roles we have for this account, including the one we created in the S3 section for cross region replication. for a new role, we click <kbd>Create Role</kbd>, we choose the service that will use the role, such as the EC2 entity. and then we attach policies to the role (like administrator policy), and in the EC2 instances, we can ssh back to it.

we can delete the _.aws_ folder (with the credentials!)

```sh
rm -rf .aws #delete
aws s3 ls #won't work
```

in the EC2 service, we choose <kbd>Actions</kbd>, and then <kbd>Instance Settings - Attach/Replace IAM Role</kbd>. we attach the admin access role which we created before,

now we can can run all commands from the EC2 instance without having the credentails stored in it.

> - Roles are more secure than storing your access key and secret access key on individual EC2 instances.
> - Roles are easier to manage.
> - Roles can be assigned to an EC2 instance after it is created by using either the web management console or the command line.
> - Roles are universal - you can use them in any region.

### Using Boot Strap Scripts

running commands in the EC2 instance once it boots up, like getting packages, updating software and running a program.

in the aws web console, we provision a new EC2 instance, in the configuration step we give it the role with admin access. this is because we want it to create a S3 bucket when it starts. then we click on _advanced details_ and there is a box called **User Date**. here we can type bash scripts which will run as the instance starts.

all scripts start with `#!/bin/bash`

in this example script we install the apache service to make the instance act as a website

```sh
#!/bin/bash
yum update -y
yum install httpd -y
service httpd start
chkconfig httpd on
cd /var/www/html
echo "<html><h1>Hello world</h1></html>" > index.html

aws s3 mb s3://someBucketName # make bucket

aws s3 cp index.html s3://someBucketName #copy file to bucket
```

### EC2 Instance Meta Data

ssh back into the ec2 instance, then curl to get the user-date

the 169.254 ip range is special

```sh
sudo su
curl http://169.254.169.254/latest/user-data > bootstrap.txt
curl http://169.254.169.254/latest/meta-data/ #see the metadata
curl http://169.254.169.254/latest/meta-data/local-ipv4 #see the local ip address
curl http://169.254.169.254/latest/meta-data/public-ipv4 #see the public ip address
```

### Elastic File System [SAA-C02]

> Amazon Elastic File System (EFS) is a file storage service for Amazon EC2 instances. Amazon EFS is easy to use and provides a siple interface that allows you to create and configure file systems quickly and easily. With Amazon EFS, storage capacity is elastic, growing and shrking automatically as you add and remove files, so your applications have the storage they need, when they need it.

EBS is mounted to a single EC2 instance, and can't be shared. EFS can be shared across the EC2 instances, and it can grow as needed, there is no wasted space like EBS volume, and we can use the same files in many instances.

in the aws management console, we choose services <kbd>EFS</kbd> under storage. and we click <kbd>Create File system</kbd>, we can have lifecycle policies, storage classes like S3, thorughput mode, performance mode and encryption.

in the EC2 service, we provision two instances, and give them the following User-data bash script

```sh
#!/bin/bash
yum update -y
yum install httpd -y
service httpd start
chkconfig httpd on
yum install -y amazon-efs-utils
```

in the security group page, we want to open the inbound rules to add the access through the NFS port (protocol TCP, port 2049) and we give the security group as the source.

back in the EFS page, we see that there are mount target states which are available. and with the EC2 instances, we grab the public ips and ssh into the EC2 instaces.

```sh
sudo ec2-user@<ip> -i key.pem
sudo su
cd /var/www/html # folder exists because of apache
cd ..

```

in the efs service, we want to grab commands to mount the efs.

```sh
mkdir efs
sudo mount -t efs fs-<identifier>:/ efs
sudo mount -t efs -o tls fs-<identifier>:/ efs #with encryption

```

but instead of mounting to efs, we want to mount to the /var/www/html folder, so in the ec2 instances

```sh
mount -t efs -o tls fs-<identifier>:/ /var/www/html #this should take a few seconds, no longer than a minute
```

now we can create a file in one ec2 instaces at the mounted folder, and it will be availbe in both instances. the file is shared in both instaces, so it's easier to update.

> - EFS supports the Network file System version 4 (NFSv4) protocol.
> - You only pat for the storage you use (no pre-provisioning required).
> - Can scale up to petabytes.
> - Can support thousads of concurrent NFS connections.
> - Data is stored across multiple Availability Zones within a region.
> - Read after Write Consistency.

### FSx for Windows & FSx for Lustre [SAA-C02]

> Amazon FSx for Windows File Server provides a fully managed native Microsoft WIndows file System so you can easily move your windows-based applications that require file storage to aws. Amazon FSx is built on Windows Server.
>
> **Windows FSx**
>
> - A managed windows server that runs Windows Server Message Block (SMB)- based file services.
> - Designed for Windows and Windows Applications.
> - Supports AD (active directory) users, access control lists groups and security policies, along with Distributed File System (DFS) namespaces and replication.
>
>   **Lustre FSx**
>
> - Designed specifically for fast processing of workloads such as machine learning, hight performance computing, video processing, financial modeling and electronic design automation (EDA).
> - Lets you launch and run a file system that provide sub-milliseconds access to your data and allows you to read and write data at speeds of up to hundreds of gigabytes per second of throughput and million of IOPS.
>
>   **EFS**
>
> - A managed NAS filer for EC2 instances based on Network file System (NFS) version 4.
> - One of the first network file sharing protocols native to Unix and Linux.

EFS is not Message blocks based.

> Amazon FSx for Luster is a fully managed file system that is optimized for compute intensive workloads, such as high performance computing (HPC), machine learning, media data processing workflows, and electronic desgin automation (EDA).
>
> With Amazon FSx, you launch and run a Lustre file system that can process massive data sets up to hundreds of gigabytes per second of thoughtput, million of IOPS (io operations per second), and sub-millisecond latencies.

summary points:

> - EFS: When you need distributed, highly resilient sotrage for linux instances and linux based applications.
> - Amazon FSx For Windows: when you need centralized storage for Windows-bae application such as Sharepoint, Microsoft SQL Server, workspaces, IIS web Server or any other native microsoft application.
> - Amazon FSx For Lustre: when you need high speed, high-capacity distributed storage. This will be for applications that do High Performance Compute (HPC), financial modling, etc... Remember that FSx for Lustre can store data directly on S3.

### EC2 Placement Groups

EC2 Placement groups are ways to place your EC2 machines in the physical world. it's how we set up the EC2 machine in terms of real computers and servers rack.

**Clustered**

> A cluster placement group is a grouping of instances within a single availability zone. Cluster Placement groups are recommended for applications that need a low network latency, high network throughput, or both.\
>  Only certain instances can be launched into a clustered placement group.
>
> **Spread**
> A spread placement group is a group of instances that are each place on distinct underlying hardware.\
>  Spread Placement groups are recommended for applications that have a small nubmer of critical instances that should be kept separate from each other.
>
> **Partitioned**
> When using partition placement groups, Amazon EC2 divides each group into logical segments called partitins. Amazon EC2 ensures that each partition within a placement group has it's own set of racks. each rack has its own network and power source.\
>  No two partitions within a placement group share the same rack, allowing you to isolate the impact of hardware failure within your application.

| Type        | Arrangement                                                 | Use case                                      |
| ----------- | ----------------------------------------------------------- | --------------------------------------------- |
| clustered   | put everything as close together as possible.               | low network latency, high network throughput. |
| partitioned | divide into groups (partitions), separate those partitions. | individual critical EC2 instances.            |
| spread      | instances are separate.                                     | HDFS,HBASE, Cassandra.                        |

> - a Clustered placement group must be in the same avalability zone.
> - The name for the placement group must be unique within the AWS account.
> - only certain types of instances can be launched in a placement group (Compute Optimized, GPU, Memort Optimized, Storage optimized).
> - AWS recommend homogenous instances within clustered placement groups.
> - Placement groups can't be merged.
> - it's possbile to move an existing (stopped) instance into a placement group.

### HPC On AWS [SAA-C02]

High performance computed.

used for industries such as genomics, finance and financial risk modeling, machine learning, weather prediction and autonomous drivers.

steps needed:

1. Data transfer
2. Compute and networking
3. Storage
4. Orchestration and automation

**data Transfer:**

- Snowball, Snowmobile (terabyte and petabyte)
- AWS dataSync to store on S3, EFS, FSx, etc..
- Direct Connect

> AWS Direct Connect is a cloud service solution that makes it easy to establish a dedicated connection from your premises to AWS. Using AWS direct connect, you can establish private connectivity between AWS and your data center, office, or colocation environment - which, in many cases, can reduce your network costs, increase bandwidth throughput, and provide a more consistent network experience than internet based connections.

**Compute and network Services:**

- EC2 instances tha are GPU or CPU optimized.
- EC2 Fleets (Spot Instances or Spot Fleets).
- Placement groups (cluster placement group).
- Enhanced networking:
  - Elastic Network Adapters.
  - Elastic Fabric Adapters.

> Enhanced networking uses **Single Root I/O virtualization (SR-IOV)** to provide high performance networking capabilities on supporeted instance types. SR-IOV is a method of device virtualization that provides Higher I/O performance and lower CPU utilization when compared to traditional virtualized network interfaces.\
> Enhanced networking provides higher bandwidth, higher packer per second (PPS) performance, and consistently lower-inter-instance latencies. _There is no additional charge for using enhanced networking._

has two different flavours:

- **Elastic Network Adapator (ENA)**, 100 GBPS.
- **Intel 82599 Virtual Function (VF)** interface, typically used for legacy instances.10 GBPS.

**Elastic Fabric Adapter**

- network device which we can attach to the amazon EC2 instance
- lower latency, higher throughput,
- uses **OS-bypass**, enables the HPC and machine learning applications to interact directly with the EFA device without the operating system kernel. linux only, not windows.

**Storage:**

- Instance attached storage:
  - EBS: scale up to 64,000 IOPS
  - Instance Store: scale to million of IOPS.
- Network storage:
  - Amazon S3: distributed object based storage, not a file system.
  - Amazon EFS: Scale IOPS based on total size, or use provisioned IOPS
  - Amazon FSx for Luster: HPC-optimized distributed file system, millions of IOPS, backed by S3.

**orchastration:**

AWS BATCH

> AWS Batch enables developers, scientists and enginners to easily and efficiently run hundreds of thousand of batch computing jobs on AWS. AWS batch support multi-node parallel jobs, which allows you to run a single job that spans multiple EC2 instances. you can easily schedule jobs and launch EC2 instances according to your needs.

AWS ParallelCluster

> Open source cluster management tool that makes it easy for you to deploy and manage HPC clusters on AWS. ParallelCluster uses a simple text file to model and provision all the resources needed for your HPC application in an automated and secure manner. Automatic creation of VPC, subnet, cluster types and instance types.

### AWS WAF [SAA-C02]

**AWS WAF - web Application firewall**

> AWS WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests that are forwarded to amazon CloudFront, an application load balancer of API gatway. AWS WAF also lets you control access to your content.

with http/https it happens at the application level, or layer 7, it can see the query string ana parameters.

so, with this query:

> `http://acloud.guru?id=1001&name=ryan`\
> there are two parameters
>
> - id, value 1001
> - name, value "ryan"

you can configure conditions for which ip addresses are allowed, what query parameters are needed, etc. the request will either be allowed to pass through or be blocked with a http 403 status code.

at the most basic level, there are 3 different behaviors for AWS WAF:

- allow all requests, except for the ones specified.
- block all requests, except for the ones specified.
- count the requests that match some specified properties (passive mode)

possible conditions:

> - IP address
> - Country that request originate from
> - Values in the request header
> - Strings in the request header
> - Lengths of the requests
> - Presence of sql code (possible sql injection)
> - Presence of scripts (cross site scripting)

in the exam, this can come up in questions about blocking malicious attacks, we could also use network ACLS.

### EC2 Summary

EC2 - virtual machine on the cloud. allows to provision computing power quickly and efficiently.

pricing modes:

- on demand
- reserved - provision in advance
- spot - bid with a price
- dedicated host - a dedicated physical machine.

EBS - Elastic block storage, termination is off by default. attached volumes aren't deleted automatically. the root volume can be encrypted now even at creation.

security groups:

- all inbound traffic is blocked by default, outbound traffic is allowed. security groups are stateful, opening a port for inbound opens it for outbound. we can specify allowed rules, but not blocing rules. for those we need ACL.

storage can be SSD or HDD, we can get different kinds of drives for different tasks. volumes exists on EBS, snapshots are on S3, snapshots are "points in time" of a volume, they are incremental, so the first one is heavy, but the next ones are just the deltas.

we can create AMI from both volumes and snapshot. Snapshot must exist in the same Avalability zone, but we can create instance from an AMI in the other avalability zone. theres also a process to move AMI from region to region.

in the past there was a need to do a process to encrypt the root volume. today it's not so needed.

Instance store volumes are ephemeral, if the instance fails, the data is lost. we can't stop the instance store volume. EBS backed stopped instances can be stopped, so we can reboot it and keep the data.

Cloud Watch to monitor performance, the default time interval is 5 minutes, but we can pay extra for detailed 1 minute intervals. Cloud trail audits AWS actions (api calls, user logins,etc...)

AWS CLI programmatic access, roles to use EC2 instances as the CLI instance without giving it the secret key. roles are universal. bootstrap scripts run when the EC2 instance starts.

- `curl http://169.254.169.254/latest/user-data/`
- `curl http://169.254.169.254/latest/meta-data/`

EFS - elastic file system, pay for what you need, shared accross instances. read after write consistency.

Placement groups:

- clustered
- spread
- partitioned

### Quiz 3: EC2 Quiz

> - Which of the following features only relate to Spread Placement Groups?
>   ANSWER: _Spread placement groups have a specific limitation that you can only have a maximum of 7 running instances per Availability Zone and therefore this is the only correct option. Deploying instances in a single Availability Zone is unique to Cluster Placement Groups only and therefore is not correct. The last two remaining options are common to all placement group types and so are not specific to Spread Placement Groups._
> - If an Amazon EBS volume is an additional partition (not the root volume), can I detach it without stopping the instance?
> - _ANSWER: YES, although it will take some time_
> - Individual instances are provisioned at?.
> - _ANSWER: In Avalability zones._
> - Can you attach an EBS volume to more than one EC2 instance at the same time?
> - _ANSWER: YES_

</details>

##

[next](Section_5_Databases.md)\
[main](README.md)
