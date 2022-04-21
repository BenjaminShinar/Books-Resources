<!--
// cSpell:ignore
 -->

[main](README.md)

## Section 8 - VPC: Virtual Private Cloud

<!-- <details> -->
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

### VPC Flow Logs

### Bastions

### Direct Connect

### Setting Up A VPN Over A Direct Connect Connection [SAA-C02]

### Global Accelerator [SAA-C02]

### VPC End Points [SAA-C02]

### VPC Private Link [SAA-CO2]

### Transit Gateway [SAA-CO2]

### VPN Hub [SAA-CO2]

### Networking Costs on AWS [SAA-CO2]

### Summary

### Quiz 6: VPCs Quiz

</details>

##

[next](Section_9_VPC.md)\
[main](README.md)
