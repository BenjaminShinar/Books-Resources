<!--
// cSpell:ignore Rodrigues nginx linux
 -->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# AWS Developer Associate Training

<!-- <details> -->
<summary>
A bit old course, so this document will be brief
</summary>

udemy course [AWS Developer Associate Training](https://www.udemy.com/course/aws-developer-associate-training/). by _Alan Rodrigues_.

## Fundamentals

<details>
<summary>
AWS Basics: VPC and EC2 machines.
</summary>

before everything, we can create a free-tier aws account.

### Regions, AZ and VPC

Regions are separate geographic areas, each region has multiple Availability Zone. those Availability Zone are connected by low-latency high speed connection.

a <cloud>VPC</cloud> is a virtual private cloud, a locally isolated section of the cloud which host resources such as virtual servers, it's a collection of subnets (resources are hosted on those subnets).. the vpc has a CIDR block, and the address block is shared across those subnets. it can also have an <cloud>Internet Gateway</cloud> or be connected to an on-premises network with VPN.

There is no charge for a VPC itself. the size of the VPC can't be changed after creation. the account has a default VPC. a subnet always corresponds to a single Availability Zone.

Services and Costs vary between regions. especially for new services.

### Default VPC

the default VPC has a default subnet, and an internet gateway, prebuilt security group, Network Access Control List and DHCP options. it's easy to start with it and launch instances onto it.

VPC can attach the DNS host of instances inside it. subnets can be exposed to the outside world, and can auto-assign public ip addresses to instances. Internet Gateways are attached to a VPC.

### Elastic Compute Cloud

<cloud>EC2</cloud> is the web service to provision virtual machines. we choose the machine based on the <cloud>AMI</cloud> (amazon machine image), which can be either a linux or windows machine. we can configure storage as either part of the machine (instance store) or connected to it (<cloud>Elastic Block Storage</cloud>) that can outlive the instance.\
We also specify the instance type (compute power), storage type and volume (SSD, HD), and then we launch the machine inside a subnet in a VPC.

Security groups define which traffic can access the instance, it is 'allow-only'. any request will have the response. For EC2 machines, we can also define SSH keys so we could connect to them via the shell.

instances always have private Ip addresses, and can have a public Ip address.

### Building Windows Instances

In this lecture, we make our windows machine into a web-server using IIS. we connect to it through remote connection. we need to allow HTTP and HTTPS access in the security group (even if we have the public ip address).

### Building Linux Instances

In this lecture we make a linux machine a web-server using nginx. this time we connect via ssh.

### Elastic IP

even if we have a public ip address, it would change every time we reset the machine. elastic ips allow us to have a consistent ip, even if there is a reset, or if we want to direct traffic between instances (load balancing).

### Network Access Control Lists

Access control lists are like firewalls, they control access to and from the internet to the **entire subnet**. they are more versatile than security groups, as the can have both 'allow' and 'deny' rules. they are also stateless - if an inbound access is allowed, it doesn't mean that outbound access (even the response) is allowed. we can you <cloud>ACL</cloud> to block traffic from a set of ip addresses.

a rule has priority (lower = evaluated first), and sets the type, protocol, ports and sources (ip), then we define it as allowed or denied.

we do this for both inbound and outbound traffic.

### Security Groups

Security groups operate on instances (not on the subnet). they are stateful. they define "allowed traffic" only, without explicit deny rules.

### Elastic Block Storage

Storage for EC2 instances. there is the root EBS volume that is attached to the EC2 machine on creation, but we can later attach additional storage. block storage persists after the EC2 machine was shutdown, and we can store it as a snapshot for backup or to transfer it across regions. we can also encrypt the storage.\
We don't need to shutdown the instance to use the newly attached storage. we can change the size of the volume after it was created.

### Public and Private Subnets

A public subnet is connected to the internet through the gateway, while the private does not. this is defined in the route table. we can also use a <cloud>NAT gateway</cloud> to allow outbound traffic for the private subnet.

There are some default, depending on whether this is the default vpc or not.

### Creating a VPC

we can create additional VPCs, and we can use some  pre-configured options (subnets, NAT access). we can use NAT instance or NAT gateway.

### Quiz 1: Quiz on Fundamentals

</details>

## Simple Storage Service S3

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## DynamoDb

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## Security

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## Additional Services

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## Additional Topics

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

</details>
