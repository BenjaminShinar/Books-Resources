<!--
// cSpell:ignore chkconfig Gbps
 -->

[main](README.md)

## Section 4 - EC2

<!-- <details> -->
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

<!-- <details> -->
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

</details>

### Spot Instances & Spot Fleets [SAA-C02]

### EC2 Hibernate [SAA-C02]

### AWS This Week

### CloudWatch 101

### CloudWatch Lab

### The AWS Command Line

### Using IAM Roles With EC2

### Using Boot Strap Scripts

### EC2 Instance Meta Data

### Elastic File System [SAA-C02]

### FSX for Windows & FSX for Lustre [SAA-C02]

### EC2 Placement Groups

### HPC On AWS [SAA-C02]

### AWS WAF [SAA-C02]

### EC2 Summary

### Quiz 3: EC2 Quiz

</details>

##

[next](section_4_EC2.md)\
[main](README.md)
