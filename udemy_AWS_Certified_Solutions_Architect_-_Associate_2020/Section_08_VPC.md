<!--
// cSpell:ignore
 -->

[main](README.md)

## Section 8 - VPC: Virtual Private Cloud

<details>
<summary>
Virtual Private Cloud.
</summary>

we can imagine vpc as virtual data center, stored in the cloud.

> Amazon Virtual Private Cloud lets you provision a logically isolated section of the Amazon Web services (AWS) Cloud where you can launch AWS resources in a virtual network that you define.\
> You have complete control over your virtual networking environment, including selection of your own IP address range, creation of subnets, and configuration of route tables and network gateways.\
> You can easily customize the network configuration for your amazon virtual private cloud, for example, you can create a public-facing subnet for your weservers that has access to the interenet, and place your backend systems such as databases or application servers in a private-facing subnet with no internet access.\
> You can leverage multiple layers of security, including security groups and network access control lists, to hel control access to Amazon EC2 instances in each subnet.\
> Additionally, you can create a Hardware Virtual Private Network (VPN) connection between your corporate datacenter and your VPC and leverage the AWS cloud as an extension of your coporate datacenter.

vpc diagram

- aws-region
- VPC
- internet gateway
- virtual private gateway
- Router
- route tables
- network ACLs
- security groups
- Subnets (SN)
- EC2 instances

network ACT are stateless. they have both _Allow_ and _Deny_ rules. we can have private and public subnets

**bastion** - ec2 instance on a public subnet which connects to a instance in private subnet.

there are 3 ip prefixes for subnets. these were assigned by the internet authority. these addresses are only availbe for private subnetworks.
Common name | lowest address | highest address |
---|---|--- |
10/8 | 10.0.0.0| 10.255.255.255 | not available in AWS, we can use 10/16 instead.
172.16/12 | 172.16.0.0| 172.31.255.255 |
192.168/16 | 192.168.0.0| 192.168.255.255 | this is the most common one.

in the web site [cidr.xyz](https://cidr.xyz/) is "AN INTERACTIVE IP ADDRESS AND CIDR RANGE VISUALIZER". this is helpful, but not required for this solution architect certification.

> what can we do with a vpc?
>
> - Launch instances into a subnet of your choosing.
> - Assign custom IP addresses ranges in each subnet.
> - Configure route tables between subnets.
> - Create internet gateway and attach it to our VPC.
> - Much better security control over your AWS resources.
> - Instance security groups.
> - Subnet network Access Control Lists (ACLs).

so far, we've been using the default VPC. the default VPC is user friendly, all subnets in the default VPC have a route out to the internet, and each EC2 instance has both a public and a private IP address.

VPC Peering

> - Allows you to connect one VPC with another via a direct network > route using private IP addresses.
> - Instance behave as if they were on the private network.
> - You can peer VPCs with other AWS accounts as well as with other > VPCs in the same account.
> - Peering is in a star configuration. 1 central and several > non-central. **NO TRANSITIVE PEERING!** each peering requires a new connection.
> - You can peer VPCs across regions.

Summery:

> - Think of VPC as a logical datacenter in AWS.
> - VPCs consists of
>   - VPG - virtual private gateways
>   - Route tables
>   - Network ACLs - access control lists
>   - Subnets
>   - Security groups
> - 1 subnet = 1 avalability zone. no limit on how many subnets inside an AZ. but each subnet can only be inside AZ.
> - Security groups are **stateful**.
> - Network ACLs are **stateless**.
> - No transitive peering.

### Build A Custom VPC

in the AWS console. under <kbd>Networking and Content Delivery</kbd> serives, we click <kbd>VPC</kbd> and we can see the VPC dashboard. we can <kbd>Launch VPC Wizard</kbd>, but we instead click <kbd>Your VPCs</kbd>, we can look at the used subnets and see the CIDR blocks for each subnet. there is also a default route table and internet gateways.

when we <kbd>Create VPC</kbd>, we use the `10.0.0.0/16` as an IPv4 CIDR block, we can get IPv6 address, and under _Tenancy_ we can decide if want dedicated hardware allocated (this increases costs). once this is created, we can look at the dashboards and see what was created:

- [x] route table
- [x] default network ACL
- [x] default security group
- [ ] subnets aren't created by defaults
- [ ] no internet gateways

to create a subnet, we click <kbd>Subnets</kbd>, then <kbd>Create subnet</kbd>. we decide on which VPC we want the subnets, which AZ to use, and `10.0.1.0/24` in the IPv4 CIDR block. we can give subnet a name tag.\
we create another subnet, in a different AZ, with different CIDR address blocks.

note: avalability zone names are randomized for each account. us-east-2a in one account is not the same AZ as us-east-2a in a different account.

we will want to enable _"auto assign public ip"_ for one of the subnets, and we can also see in the dashboard that the _"Available IPs"_ shows a slightly smaller number than what it should - 251 instead of 256 (the `/24` prefix means we have 8 bits for addresses, which should be 256). the reason is the subnet has reserved 5 vpc addresses.

- 10.0.0.0: network address
- 10.0.0.1: VPC router.
- 10.0.0.2: reserved. ip address for the DNS of the VPC.
- 10.0.0.3: reserved.
- 10.0.0.255: network broadcast. not supported in VPC, but reserved.

to make one subnet publicly accessable, we select one, click <kbd>Actions</kbd>, then <kbd>Modify auto-assign IP settings</kbd> and check the box.

next is to add an internet gateways,so under <kbd>Internet Gateways</kbd>, we click <kbd>Create Internet Gateway</kbd>, and simply give it a name. this new gateway is in a **Detached** state, so we click <kbd>Actions</kbd> -> <kbd>Attach To VPC</kbd>, select our new vpc and click <kbd>Attach</kbd>.

**Only One Internet Gateway can attached to a VPC**

now, moving to <kbd>Route Table</kbd>, we select the newly created (by deafult) route table, and move to the <kbd>Routes</kbd> tab and we see connection to the outside networks (main route table), we can visit the <kbd>Subnet Associations</kbd> tab and see associated subnets. we don't want the main route table to be public, so we click <kbd>Create Route Table</kbd>, give it a name and the VPC.\
we then choose the route table, click <kbd>Edit Routes</kbd> and add `0.0.0.0/0` to open a route out to the internet. the target is _internet gateway_ and we select the internet gateway which we created. for IPv6 routes, we use `::/0` to open access. now we associate the one of the subnet to the new route table.

we now provision ec2 instances, one in the public subnet and one in the private subnet.

when we configure the EC2 instances, select our new VPC under the <kbd>Network</kbd> field and the required subnet under the <kbd>Subnet</kbd> field. we create a new security group with ssh and http open. security groups are tied to the VPC. we can also see that auto assigning public ips is determined by the subnet.

one EC2 machine will be a webserver, and one will be a database instance.

even if we have the private ip address of the database instance, we can't ssh into it yet. so we create a new security group for it, under the VPC, and we add inbound rules. we use ICMP, HTTP, HTTPS, SSH and NTSQL/Aurora all with the same source `10.0.0.1/24`, which is the cidr block.

we select the EC2 instance, <kbd>Actions</kbd>-> <kbd>Networking</kbd> -> <kbd>Change Security Group</kbd> and select the new security group. now we can use the private ip from the webserver which we are already logged into.

for this example, we copy the private key into the ec2 machine and ssh into the private machine.
inside the private subnet, we can't get outside to the internet, but we want to update our software.

Summary:

> - When creating a VPC we get:
>   - [x] a deafult route table
>   - [x] network Access Control Lists
>   - [x] default security group
>   - [ ] we don't get a subnet
>   - [ ] we don't get an internet gateway
> - Amazon reserves 5 ip address within the subnets.
> - Only internet gateway per VPC.
> - Security Groups don't span Across VPCs.

### Network Address Translation (NAT)

NAT instances and NAT Gateways, a way to communicate with the internet gateway. we usually use NAT gateways, as NAT instances are getting phased out. A NAT instance is a EC2 instance, while NAT gateway is an highly aviable gatway.

**NAT Instance**

under the EC2 service, we create and EC2 Instance, select the community AMIs and search for "nat". we use the default values, and select the VPC and the public subnet.

we want to disable Source/Destintation Checks on the NAT instance, so we click <kbd>Actions</kbd>, <kbd>Networking</kbd> and disable the <kbd>Source Destination Check</kbd>. next, we want to allow the ec2 instances to to talk to the NAT instance with the route table.

so under the <kbd>VPC</kbd> service, we select the <kbd>Route Table</kbd> and <kbd>Edit Routes</kbd> of the public table, we set the destination of `0.0.0.0/0` and use the NAT instance as the target.

we ssh into the public EC2 machine, then SSH into the private EC2 machine, and now when we run `yum update`, the call goes through the NAT instance and we get a response.

the problem is that the NAT instace is a single EC2 instance, it can fail or be overwhelmed, this is a single point of failure. so instead, we can create a NAT gateway.

**NAT Gateway**

under the <kbd>NAT Gateways</kbd> service, we click <kbd>Create NAT Gateway</kbd>, select the public subnet, and create an elastic IP address.

now, we go to the route table, add the `0.0.0.0/0` route and target the gateway. this might take a few minutes to start.

now we can run more commands from the ec2 instance and get to the outside world.

Summary:

> Nat Instances:
>
> - When creating a NAT instance, disable source/destination check on the instance.
> - NAT instance must be in a public subnet.
> - There must be a route of of the private subnet to the NAT instance.
> - The amount of traffic that instances can support depends on the instance size. if you are bottlenecking, increase the instance size.
> - you can create high availability using Autoscaling Groups, multiple subnets in different AZs, and a script to automate failovers, but it's a pain.
> - always behind a security group.
>
> Nat Gateways:
>
> - Redundant inside the availability zone.
> - preferred over NAT instances
> - Starts at 5GPbs and scales up
> - No need to patch
> - Not associated with security groups
> - Remember to update the route table
> - NAT gateways belong to the AZ, if the AZ is done, then every service which uses the gateway loses access. to solve this, add a NAT gateway to each AZ and configure all resources to use the gatway from their own AZ.

### Access Control Lists (ACL)

NAT access Control Lists vs Security groups.

in the console, Under <kbd>VPC</kbd>, we look at the <kbd>Network ACLs</kbd>, we look at the <kbd>inbound rules</kbd> tab and the <kbd>outbound rules</kbd>. we click <kbd>create network ACL </kbd>, and by default all the inbound and outbound rules are blocking data.

we move the subnet (subnet association) into the network ACL which we created. so now we need to edit the inbound rules. we allow port 80, 443 and 22 (http, https, ssh). for outbound we allow port 80 and 443, but we also specify the range of `1024-65535` as _ephemeral ports_. they are allocated per session, a client goes through one of the main ports (80,443) and is then directed to an ephemeral port. we should also allow them as _inbound rules_.

rules are evaluated by order, and the best practice is to give the numbers in increments of 100. we can also have _deny_ rules, like blocking a port from a specific ip. the deny rule must be before (lower number) than the allowing rule. ACLs are evaluated before security groups.

summary:

> - Your VPC automatically comes with a default network ACL, and by default it allows all outbound and inbound traffic.
> - You can create custom network ACLs, by default, each custom network ACL denies all inbound and outbound traffic until you add rules.
> - Each subnet in your VPS must be associated with a network ACL. If you don't explicitly associate a subnet with a network ACL, the subnet is automatically associated with the default network ACL.
> - Blockin Ip addresses using Network ACL, but not with security groups.
> - you can associate a network ACL with multiple ACL, but a subnet can be associated with only one network ACL.
> - Network ACLs contain a numbered list of rules that is evaluated in order, starting with the lowest numbered rule.
> - Network ACLs have separe inbound and outbound rules, and each rule can either allow or deny traffic.
> - Network ACLs are stateless, responses to allowed inbound traffic are subject to the rules for outbound traffic (and vice versa.).

### Custom VPCs and ELBs

we need to discuss elastic load balancers in short.

in the console, <kbd>services</kbd>, <kbd>EC2</kbd>, select <kbd>Load Balancing</kbd> from the side bar, then <kbd>Create Load Balancer</kbd>, we will see three types:

- Application Load Balancer (http, https)
- Network Load Balancer (tcp, tls)
- Classic Load Balancer (previous generation - http, https, tcp)

we create an application Load Balancer, we can choose a scheme to be either _internet facing_ or _internal_. and the IP address type to be _ipv4_ or _dualstack_.\
we choose the custom vpc we created, and if we click a subnet with no internet gateway, we will see a warning. so we choose the subnet that has the gateway configured, but we also need to choose at least two public subnets, so we can't do this right now.

### VPC Flow Logs

> VPC Flow Logs is a feature that enables you to capture information about the IP traffic going to and from network interfaces in you VPC.\
> Flow log data is stored using Amazon CloudWatch Logs, after you've create a flow log, you can view and retrieve its' data in Amazon Cloud Watch Logs.

Flow logs can be created at three levels.

- VPC
- Subnet
- Network Interface Level

in the console <kbd>Services</kbd>, under **network content & delievery** we select <kbd>VPC</kbd>, select one of our VPCs, <kbd>Actions</kbd> and choose <kbd>Create Flow Log</kbd>.

we have filters (the type of traffic to log: all, accepted and rejected), we can the the data to cloud Watch logs or to an S3 bucket. the select a log destination, which we need to create.

<kbd>CloudWatch</kbd> (under **management and governance**), and <kbd>Create Log Group</kbd>.

we also need an IAM role, AWS provides an option to quickly create a role with the nesscary permissions.

we can now watch the logs under <kbd>Cloud Watch</kbd>, this might take some time to start, and now we can see the traffic of our network.

> - You cannot enable flow logs for VPCs that are peered with your VPC unless the peer VPC is in your account.
> - You can Tag Flow logs
> - After you've create a flow log, you cannot change its configuration.
>   - (e.g) You cannot associate a different IAM role with the flow log.
> - Not all IP traffic is monitored.
>   - [ ] Traffic generated by instances when they contact the Amazon DNS server are not monitored.
>   - [x] If you use your own DNS server, then all traffic to that DNS server is logged.
>   - [ ] Traffic generated by a Windows Instance for Amazon Windows Licensee activation is not monitored.
>   - [ ] Traffic to and from `169.254.169.254` for instance metadate is not monitored.
>   - [ ] DHCP traffic is not monitored
>   - [ ] traffic to the reserved IP adddress for the default VPC router is not monitored.

### Bastions

> A bastion host is a special purpose computer on a network specifically designed and configured to withstand attacks. the computer generally hosts a single application, for example, a proxy server, and all other servies are removed or limited to reduce the thread to the computer.\
> It is hardned in this manner primarily due to its location and purpose, which is either on the outside of a firewall or in a demilitarized zone (DMZ) and usually involves access from untrusted networks or computers.

a bastion host will be accessible from outside. it is the point of attack for access, it is what we use to ssh into instances on private subnets.

> - A NAT gateway or NAT instances is used to provide internet traffic to EC2 instances in a private subnet.
> - A Bastion is used to securely administer EC2 instances (using SSH or RDP). Bastions are called jump boxed in Australia.
> - You cannot use a NAT gateway as a Bastion host.

### Direct Connect

> AWS Direct Connect is a cloud service solution that makes it easy to establish a dedicated network connection from your premises to AWS. Using AWS Direct Connect, you can establish private connectivity between AWS and you datacenter, office or colocation environment, which in many cases can reduce you network costs, increase bandwidth throughput, and provide a more consistent network experience than internet-based connections.

using a dedicated line to connect to AWS. Direct Connect locations(DX) have an AWS cage (with a DX router) and a matching Customer/Partner cage with a Router, we connect the customer datacenter to the router in the DX cage, and the two routers a are connected with **X-connect** (cross connect), which are then connected to AWS using the **AWS Backbone Network**.

> - Direct Connect Directly connects your data center to AWS.
> - Useful for high throughput workloads (lots of network traffic)
> - Useful when a stable and reliable connection is require.

Steps:

1. Create a **Public Virtual Interface** in the direct connect conole.
2. Go to the <kbd>VPC</kbd> console and then <kbd>VPN</kbd> connections. create <kbd>Customer Gateway</kbd>.
3. Create a <kbd>Virtual Private Gateway</kbd>.
4. Attach the **Virtual Private Gateway** to the desired VPC.
5. Select <kbd>VPN</kbd> connections and create a new VPN connections.
6. Select the **Virtual Private Gateway** and the **Customer Gateway**.
7. Once the VPN is available, set up the **VPN** on the customer gateway or firewall.

### Global Accelerator [SAA-C02]

> AWS Global Accelerator is a service in which you create accelerators to improve avalability and performance of your applications for local and global users.\
> Global accelerators direct traffic to optimal endpoints over the AWS global network. This improves the availability and performace of your internet applications that are used by a global audience.

there are two default static IP addresses associated with the accelerator, but it is possible to bring your own ip addresses.

in the console, under <kbd>AWS Globabl Accelerator</kbd>, we can see a diagram. The user connects to an AWS edge location,which is connected to AWS global accelerator. then to the **Endpoint group**, which are associated with the AWS region, and then the traffic is directed to an application endpoint.

Global Accelerator concentrates the traffic into the Amazon Web Backbone network, so things are much faster.

AWS Global Accelerator Components:

- Static IP addresses
- Accelerator
- DNS Name
- Network Zone
- Listener
- Endpoint Group
- Endpoint

the accelerator direct traffic, each accelerator contains one or more listeners. the accelerator has a DNS name which points to the static IP address. we can set up another DNS record to route traffic from our custom domain name into the DNS name or the ip addresses.

**Network Zones**

> A Network Zone services the static IP address for the accelerator from a unique IP subnet. similar to an AWS avalability Zone, a network zone is an **isolated unit** with it's own set of physical infrastructure.\
> When you configure an accelerator, by default, Global Accelerator allocates two IPv4 ip addresses for it. If one Ip address from a network zone becomes cantabile (due to ip address blocking by client network or network disruptions), the client applications can retry on the healthy static IP address from the other isolated network zone.

**Listener**

> A listener processes inbound connection from clients to Global Accelerator, based on the port (or port range) that you configure. Global Accelerator supports both TCP and UDP protocols.\
> Each listener has one ore more endpoint groups associated with it, and traffic is forwarded to endpoints in one of the groups.\
> You can associate endpoint groups with listeners by specifying the regions that you want to distribute traffic to. Traffic is distributed to optimal endpoints within the endpoint groups associated with a listener.

**EndPoint Group**

> Each endpoint group is associated with a specific AWS Region. Endpoint Groups include one or more endpoints in the region.\
> You can increase or reduce the percentage of traffic that would be otherwise directed to an endpoint group by adjusting a setting called <kbd>Traffic Dial</kbd>.

The _traffic dial_ allows for easy performance testing or blue/green deployment testing for new releases across different AWS regions.

**Endpoint**

> Endpoints can be _network load balancers_,_Application Load Balancers_, EC2 instances or _Elastic IP_ addresses. An Application Load Balancer can be internet-facing or internal.\
> Traffic is routed to endpoints based on configuration options that you choose, such as endpoint weights. For each endpoints, you can configure weights, which are number that control the proportion of traffic to route to each endpoint. this can be useful for doing a performance testing within a region.

- Accelerator ...\* many ip addresses.
- Accelerator ...\* many Listerns.
- Listener ...\* many Endpoint Groups.
- EndPoint Group ...\* many Endpoints.

in the console, we need to create an endpoint. we create an EC2 instance, using the default configuration.

we find the <kbd>Global Accelerator</kbd> service under **Networking & Content Delievery**, then we click <kbd>Create Accelerator</kbd>, we choose the IPv4 address types. we choose listerns based on Port, Protocol,and client affinity (for statefull applications). then we choose endpoint Groups,and start adding endpoints to each group.\
Once we create the accelerator, we can see the static ip addresses we got. we can edit the configurations and listeners, and then we can disable it a and delete it.

> - AWS Global Accelerator is a service in which you create accelerators to improve availability and performance of your applications for local and global users.
> - Your are assigned two static IP addresses (or alternatively, you can bring your own)
> - You can control traffic using _traffic dials_. this is done within the endpoint group.

### VPC End Points [SAA-C02]

> A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by _PrivateLink_ without requireing an _internet gateway_, _NAT device_, _VPN connection_ or _AWS Direct Connect_ connection. Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.\
> Endpoints are virtual devices. they are horizonally scaled, redundant and highly available VPC components that allow communication between instances in your VPC and services without imposing avalability risks or bandwidth constraints on your network traffic.

two types:

- Interface endpoints
- Gateway endpoints

> An interface endpoint is an _elastic network interface_ with a private IP address that serves as an entry point for traffic destined to a supported service.\
> The current services are suppoeted:
>
> - Amazon API Gateway
> - AWS CloudFormation
> - Amazon CloudWatch
>   - Amazon CloudWatch Events
>   - Amazon CloudWatch Logs
> - AWS CodeBuild
> - AWS Config
> - Amazon EC2 API
> - Elastic Load Balancing API
> - AWS Key Management Service (KMS)
> - Amazon Kinesis Data Streams
> - Amazon SageMaker
>   - Amazon SageMaker Runtime
>   - Amazon SageMaker Notebook Instance
> - AWS Secret Manager
> - AWS Security Token Service
> - AWS Service Catalog
> - Amazon SNS
> - Amazon SQS
> - AWS Systems Manager
> - Endpoint services hosted by other AWS accounts
> - Supported AWS marketplace Partner services
> - (more probably added)

Gateway endpoints currently support:

- Amazon S3
- DynamoDB

to create a gateway endpoint, in <kbd>IAM</kbd> we make sure we have a role with the _S3 full access_ policies for EC2 instances. next in <kbd>EC2</kbd>, we add the role to the instances (the **"MyDBserver** machine). next we SSH into the ec2 instance.

```sh
ssh ec2-user@18.216.240.182 -i MyNewKP.pem # ssh into bastion
sudo su
ls # make sure private key still exists
ssh ec2-user@10.0.2.235 -i MyPvKey.pem # ssh into the private ec2 machine
sudo su
aws s3 ls #list buckets
echo "test" > test.txt
ls
aws s3 cp test.txt s3://bucket-name
```

now we want to remove the route from the gateway, so under <kbd>VPC</kbd>, we select <kbd>Route Tables</kbd>, and we edit out the route to the NAT gateway.

now back in the shell, we can't list the buckets or do any connection to the outside internet.

now in the <kbd>Endpoints</kbd> section, we <kbd>Create Endpoint</kbd> by choosing **AWS services** and then matching the s3 gateways, we select the subnet, the route table and create the VPC endpoint. it might take some time for the route table to update.

then in the shell, we now need to specify the region

```sh
aws s3 ls --region us-east-2
```

> Two types of VPC endpoints
>
> - Interface Endpoints
>   - too many services to count
> - Gateways Endpoints
>   - Amazon S3
>   - DynamoDb

### VPC Private Link [SAA-CO2]

opening up a service from one VPC to another VPC.

- we can open it up to the internet
  - Security considerations
  - A lot to manage - firewalls, ACL, etc...
  - everything in the public subnet is public.
- use VPC peering
  - create and manage many different peering relationships
  - the whole network will be accessbile

so there is a third option, **VPC Private Link**

> - The best way to expose a service VPC to tens, hundreds or thousands of customer VPCs.
> - Doesn't require VPC Peering, Route Tables, NAT, IGW etc...
> - Requires a Network Load Balancer (NLB) on the service VPC and an Elastic Network Interface (ENI) on the customer VPC.

### Transit Gateway [SAA-CO2]

over the years, architecture tends to grow more and more complicated, many services, VPCs, networks, subnets, and other stuff.

the AWS Transit gateway is a single connection point that simplifies connection.

> - Allows you have transitive peering between thousands of VPCs and on-premises data centers.
> - Works on a **hub-and-spoke** model.
> - Works on a regional bassis, but you can have it across multiple regions.
> - You can use it across multiple AWS accounts using RAM (Resource Access Manager).
> - You can use route tables to limit how VPCs talk to one another.
> - Works with Direct Connect as well as VPN Connections.
> - Supports **IP multicast** (which isn't supported by other AWS services).

### VPN Cloud Hub [SAA-CO2]

allows users to connect to a single VPN hub, a single point of contact.

> - If you have multiple sites, each with its own VPN connection, you can use AWS VPN CloudHub to connect those sites together.
> - **hun-and-spoke** model.
> - Low cost, easy to manage.
> - operates over ht public internet, bu all traffic between the customer gateway and the AWS VPN CloudHUb is encrypted.

### Networking Costs on AWS [SAA-CO2]

Different scenarios about cost optimizations.

traffic coming in the VPC is free.
internal traffic is free if it's in the same Region and using a private IP, otherwise there are costs. going outside the AZ (to the internet) will cost more.
connecting between VPCs in different regions costs more.

> - Use private IP addresses over Public IP addresses to save costs, this then utilizes the AWS backbone network.
> - If you want to cut all network costs, group your EC2 instances in the same avalability zone and use private IP addresses. this will be cost-free, but make sure to keep in mind single point of failure issues.

### Summary

build a VPC from memory - subnets, private and public ips

Think of a VPC as a logical datacenter in AWS.

- Consists of IGW (internet gateways) or virtual private gateways, Route Tables, Network Access Control Lists, Subnets and Security Groups.
- 1 subnet = 1 Avalability Zone.
- Security Groups as _statefull_ (in bound and outbound are together).
- Network Access Control Lists are _stateless_.
- **NO Transitive Peering!**

**Creating a VPC**

- When creating a VPC we get
  - [x] default route Table
  - [x] Network Access Control Lists (NACL)
  - [x] default security group
  - [ ] not creating subnets
  - [ ] not creating default internet gateway
- Availability zone names are randomized across accounts. the US-EAST-1A zone in one account won't be the same data center as US-EAST-1A in a different account.
- Amazon reserves 5 Ip address within the subnet.
  - `10.0.0.0`: network address
  - `10.0.0.1`: VPC router.
  - `10.0.0.2`: reserved. ip address for the DNS of the VPC.
  - `10.0.0.3`: reserved.
  - `10.0.0.255`: network broadcast. not supported in VPC, but reserved.
- only one Internet Gateway per VPC.
- Security Groups can't span VPC

**NAT Instances vs NAT Gateways**

- NAT: Network Address Translation - connect to the internet.
- when creating a NAT instance, disable _source/destination check_ on the instance.
- NAT instances must be in a public subnet.
- There must be a route (in the route table) out of the private subnet to the NAT instance, in order for it to work.
- The amount of traffic that NAT instances can support depends on the instance size. if you're experiencing bottlenecks, increase the instance size.
- You can create high avaliability using autoscaling groups, multiple subnets in different AZs and a script to automate failover. but it's easier to use a NAT gateway.
- NAT instances are always behind a security group.
- NAT Gateways are redundant inside the avalability zone.
- NAT Gateways are the preferred enterprise solution.
- NAT Gateways start at 5GB per second and scales (45GB).
- No need to worry about patching NAT Gateways
- NAT Gateways aren't associate with security groups.
- Requires updating the Route tables
- no need to disable _source/destination check_ for NAT Gateways.

> if you have resources in multiple avalability zones and they share one NAT gateway, in the event that the NAT Gateway AZ is down, resources in the other AZ lost internet access.\
> To create and AZ independent architecture, create a NAT Gateway in each AZ and configure your routing to ensure that resource that use the NAT getway which in the same AZ as they are.

**Network Access Control List (NACL)**

- Your Vpc automatically comes with a default network ACL, and by default it allows all outbound and inbound traffic.
- You can create custom network ACLS. by default, the new NACLs deny all inbound and outbound traffic until you add rules.
- Each subnet in the VPC must be associated with a network ACL. any subnet which wasn't explicitly associated with an NACL is associated with teh default network ACL
- You can block address using NACL, but not with security groups.
- You can associate a newwork ACL with multiple subnets, but a subnet can only be associated with one network ACL at a time.
- Network ACLs have separate inbound and outbound rules, each rule can either allow or deny traffic.
- Network ACLs contain a numbered lists of rules that is evaluated in order, starting with the lowest numbered rule.
  - **deny rules** should be before (lower number) than **allow rules**.
- network ACLS are _stateless_. responses to allowed inbound traffic are subject to the same rules as other outbound traffic.

**ELBs and VPC**

- you need a minimum of two public subnets to deploy an internet facing load balancer.

**VPC Flow Logs**

- you cannot enable flow logs for VPCs that are peered with you VPC unless the peer VPC is in your account.
- You can now tag flow logs
- after a flow logs is created, you cannot change the configuration (like chaning IAM role).
- not all IP traffic is monitored
  - [ ] Traffic to to Amazon DNS server is not monitored.
  - [x] Traffic to custom DNS server is monitored
  - [ ] Traffic generated by a windows instance for Amazon Windows license activation is not monitored.
  - [ ] Traffic to and from `169.254.169.254` (instance metadata) is not monitored.
  - [ ] DHCP traffic is not monitored
  - [ ] Traffic to the reserved Ip addresses of the subnet is not monitored.

**Bastion Hosts**

- A NAT Gateway or Nat Instance is used to provide internet traffic to EC2 instances in a private subnet.
- A bastion is used to securely administer EC2 instances (using SSH or RDP).
- A NAT gateway can't be used as a bastion Host.

**Direct Connet**

- directly connect your data center to AWS
- when you need a reliable/secure connections
- steps:
  - create a virtual interface in the direct connect console. this is a PUBLIC Virtual Interface.
  - go to the VPC console and the VPN Connections. Create a Customer Gateway.
  - Create a Virtual Private Gateway.
  - Attach teh Virtual Private Gateway to the desired VPC.
  - Select VPN connections and create new VPC connection.
  - Select the Virtual Private Gateway and the Customer Gateway.
  - Once the VPN is available, setup the VPN on the customer gateway or firewall.

**Global Accelerator**

- AWS Global Accelerator is a service in which you create accelerators to improve avalability and performance of your application for local and global users.
- You are assigned two static IP addresses by default, but you can bring your own.
- you can control traffic using traffic dials. this is done within the endpoint group.
- you can control weighting to individual end points using weights.
- an endpoint can be an ELB, EC2 instance, etc...

**VPC Endpoints**

> A Vpc Endpoint enables you to privately connect your VPC to suppoeted AWS services and VPC endpoint services powered bt PrivateLink without requeuing an internet gateway,NAT device, VPN connection or AWS direct Connect connection. Instance in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service doesn't leave the Amazon Network.\
> Endpoints are virtual devices, they are horizonally scaled, redundant and highly available VPC components that allow communication between instances in your VPC and services without imposting availability risks or bandwidth constraints on your network traffic.

- interface endpoints - any services
- gateway endpoints - s3, dynamoDB

**AWS Private Link**
peer multiple VPC, using **network load balances** and **ENI**

**Transit Gateway**

- hub and spoke model
- simply network topography
- transitive peering
- use route table to limit how vpc connect
- support IP multicast

**VPN Cloud Hub**
Simplify VPN connections

### Quiz 6: VPCs Quiz

> -"Having just created a new VPC and launching an instance into its public subnet, you realize that you have forgotten to assign a public IP to the instance during creation. What is the simplest way to make your instance reachable from the outside world?" _ANSWER: Create an Elastic IP address and associate it with your instance_
> -"Are you permitted to conduct your own vulnerability scans on your own VPC without alerting AWS first?" _ANSWER: Depends_
>
> - "By default, instances in new subnets in a custom VPC can communicate with each other across Availability Zones?" _ANSWER: True_
> - "How many internet gateways can I attach to my custom VPC?" _ANSWER: 1_
> - "When I create a new security group, all outbound traffic is allowed by default?" _ANSWER: TRUE_

</details>

##

[next](Section_09_HA.md)\
[main](README.md)
