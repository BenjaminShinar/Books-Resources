<!--
// cSpell:ignore Unmanaged PaaS IaaS SaaS
 -->

# AZ-900 Microsoft Azure Fundamentals

udemy course [Learn the fundamentals of Azure, and get certified, with this complete beginner's AZ-900 course, includes practice test!](https://www.udemy.com/course/az900-azure/). by *Scott Duffy*.

> Learn the fundamentals of Azure, and get certified, with this complete beginner's AZ-900 course, includes practice test!

[exam requirements](https://learn.microsoft.com/en-us/certifications/exams/az-900)


## Introduction
<details>
<summary>
Course introduction.
</summary>

> Foundational level knowledge of cloud services and how those services are provided with microsoft azure.

for people without technical background, and for people who want a certificate.

the exam covers three topics
- Describe cloud concepts
- Describe Azure architecture and services
- Describe azure management and governance


> "What is *the cloud?*"\
> the common answer is "someone else's computer".

When we say "cloud", it usually means that we rent computing resources from some outside source, this is opposed to running the servers on premises or contracting some company to host your machines.

> Computing resources:
> - Windows and Linux Servers
> - Unlimited File Storage
> - Databases
> - Queues
> - Content Delivery Network
> - Batch Processing Jobs
> - Big Data (Hadoop)
> - Media services
> - Machine learning
> - Chat Bots
> - Cognitive Services

and many more services, some are basic and some are advanced.


### A Quick Look Into Azure
showing how to create a virtual machine in Azure. there are many options, but we can ignore them and use the default, and things will work. we can choose the image, the instance type (compute power).\
for windows machine, we need a admin user, for linux machine we use SSH keys. if we use windows servers, we can provide an existing license information and reduce our costs. we can set backups, automatic shutdown hours, and more stuff. each machine must exist in a network.

### AZ-900 Exam Requirements
The AZ-900 exam is the basic exam for azure.

> Azure Fundamentals exam is an opportunity to prove knowledge of cloud concepts, Azure services, Azure workloads, security and privacy in Azure, as well as Azure pricing and support. Candidates should be familiar with the general technology concepts, including concepts of networking, storage, compute, application support, and application development.

this is a one-time certificate, it isn't limited like other certificates.

### FAQs

</details>

## Cloud Computing
<details>
<summary>
Understanding the basics of cloud computing.
</summary>

cloud - pay only for what you use.

### What is Cloud Computing
describing cloud computing. all kinds of resources, such as virtual machine, virtual app (a wrapper over a virtual machine), a function app (serverless computing), logic app (connecting application together). there are also different categories of resources (not just computing), such as AI and machine learning, data analytics, block chain, databases, devops, databases, we can find many more options under the marketplace...

### Shared Responsibility Model

layers:
- Building Security
- Physical Network Security
- Physical Computer Security
- Operating System Patches
- Network and Firewall Settings
- Application Settings
- Authentication Platform
- User Accounts
- Devices
- Data

the responsibilities for each of the layers is divided between the customer and the cloud vendor. different services have different separation lines, some services give the user more control and some are more managed. in IaaS (infrastructure as a service), azure takes care of the physical layers, but the user controls updating the machines. in PaaS (platform as a Service), the vendor controls more aspects of the operating system and the networking, in SaaS (software as a service), the resource includes the application code, and the user just provides that data and interacts with it.

### Public Cloud, Private Cloud, Hybrid Cloud

Public cloud - computing services offered by a vendor over the public internet, anyone can purchase them. the vendor owns the hardware, the network and the infrastructure.\
Private cloud - computing services accessible either by the public internet or internal network, but the hardware belongs to the customer. azure has **Azure Stack**, which is a way to run cloud services on-premises. the customer provides the hardware.\
Sovereign clouds - government clouds,like us-gov or china.\
Hybrid cloud - combining a private and public cloud, leveraging cloud services together with on-premises operations. sliding scale of how much work in done locally and how much is done on the cloud. can be used for different workloads or for scaling.

### Cloud Pricing

Pricing models, this can be complicated. when we run on-premises, we can calculate the cost by adding the hardware costs, networking, utilities and paying for employees to run it. when using a hosting service, the payment is usually decided up front.\
for the cloud, it's different, the price is usually decided by at several metrics.As in example, for CosmosDB: operations, consumed storage, access and backup.

there are a number of free services (or free until a threshold)
- Virtual Network
- Private ip Address
- Azure Migrate
- Inbound Internet Traffic
- 5GB of Outbound Internet Traffic
- Azure policy
- Azure AD
- 1 million executions of Azure Functions
- Azure App services (some)

other services are charged by time
- Virtual Machines
- App services
- Databases
- Load Balancer
- managed Storages
- public ip Address

and some are charged by storage
- Database Storages
- Backups
- Unmanaged Disks
- Network traffic (between regions)
- Outbound traffic(above 5GB monthly)

some services charge by operation (usually fractions of a penny)
- Unmanaged Storage (reads, writes, deletes)
- Database Queries
- Messaging


some are charged by execution
- Azure functions (above the threshold)
- Serverless Databases
- Messaging Services
- Logic Apps

there are also regional differences, some regions cost more to operate in.


### Live Demo: Pricing Calculator

estimating the cost of azure resources, changing the settings and the region to see a monthly estimate of how much provisioning them will cost the user. this can be saved, shared, and updated as requirements become more clear over time.

</details>

## Benefits of Cloud Computing
<!-- <details> -->
<summary>

</summary>

- high availability
- reliability and predictability
- security and governance
- manageability

> Benefits
> - cost saving (both real and accounting)
> - availability and scalability
> - reliability and predictability
> - security and governance
> - manageability
> - global reach
> - range or ready on-demand services
> - range of tools


### Cost Savings Benefit of Cloud Computing

CapEx (capital expenditure) and OpEx (operation expenditure).

economics of scale - microsoft pays much less for hardware. total cost of ownership - electricity, internet, cooling, employees. we can also avoid consuming resources by scaling down, and we don't need to over buy machines in expectation of growth. with the cloud, we can provision what we need, and then get more if the need arises.

### High Availability, Scalability and Elasticity

high availability - is there a downtime? is the service responsive when it's needed?

we can achieve high availability with scalability: handling a growth of users or work. we can added capacity to the application to handle the increased demand. azure has the concept of elasticity - if we can detect we are reaching the max capacity, then we can provision more compute power automatically, and when the demand drops, we release the resources, as they are not needed anymore.

### Reliability and Predictability
### Security, Governance and Monitoring


</details>

## Cloud Service Types
<details>
<summary>
IaaS, PaaS, SaaS
</summary>

IaaS, PaaS, SaaS. not having to buy resources upfront, allowing for flexible use, it's also possible to have a cost saving plan.

IaaS - the lower level building blocks, usually there is a "real-world" alternative.
- computing - azure virtual machine
- storage - azure storage, (blobs, files, queues, tables), can also be configured as a data lake.
- networking - virtual networking, (they don't cost anything on their own), bandwidth (ingress and egress)

PaaS - a service layer on top of Iaas, middleware, development tools, ci-cd features... freedom from worrying about the VM
- azure app services
- managed storage
- azure sql database
- networking - Azure Front Door, load balancer, firewall

SaaS - ready-to-use applications 
- cloud apps
- office tools - office35

</details>

## Core Architectural Components
<details>
<summary>
Separation of resources- physical and logical
</summary>

### Regions, Region Pairs, Sovereign Regions

Regions, more than 60, some are more accessible than others. Azure has the most regions among the cloud vendors. Regions are divided into pairs, those pairs have the fastest network connection between them, when there are roll-outs or fallouts, one pair is done before the other, regions are usually paired inside the same nation (not always).\
Sovereign regions (US government, China) are special regions that have some special requirements, and aren't available for all.

### Availability Zones and Data Centers
availability zones reside inside regions, they are data centers that are the physical servers of the region. those data centers are connected to one another, but if one fails, the others should still work. not all regions have availability zones (two or more data centers).\
A data center has a separate power and networking infrastructure.

### Resources, Resource Groups, Subscriptions, Management Groups

Logical separation of resources, doesn't have to follow the physical grouping.
- management groups
- subscriptions
- resource groups
- resources

subscription is a unit of billing, users can have access to more than one subscription, and that access can be modified by the role. management groups can be nested, and control policies can be applied to management groups.
</details>

## Compute And Networking Services
<details>
<summary>
More in-depth description of Azure Services.
</summary>


### Azure Compute Services

- Virtual machines
- VM scale Sets
- App services
- Container instances
- Azure Kubernetes services
- Windows Virtual Desktop

> "executing code in the cloud"

We can take a machine running locally, and just push it onto the cloud, then it's loaded into a sever into the data center, which is a very powerful machine which virtualized the requested machine. the VM can have different number of cores (CPU), RAM and so on...\
VM Scale Sets are VM that run the same code and have a load balancer in front of them that directs requests to the VMs, the number of machines can grow and reduce based on demand.\
App-services (web-apps) are published apps in azure, there is no vm to manage. there is aks (azure kubernetes) and native container instances (ACI) - which are more fitting for small scale containers.\
Virtual desktop are just as the name suggests, they are a virtualization of a desktop computer, with a GUI, file system... etc.


### Azure Networking Services
- Virtual Networks (like VPC)
- VPN gateways
- Peering
- ExpressRoute - high speed **private** connection to Azure. not running over the internet.



- connection services
- protection services(ddos, firewall, security groups, private link)
- delivery services
- monitoring services - Azure monitor

subnets, vpn (connecting two networks as if they are the same one), DNS (domain name service resolution)
### Network Peering
virtual networks have address space, and can be divided into subnets, there are 5 reserved Ips for each subnet. communicating across virtual networks is done with "peering". this allows traffic to travel between the networks, and resource addresses can be resolved even if they reside the other network. Peering could be bi-directional or uni-directional. this can be done in the same region, or across regions, but global peering has more costs.

### Public and Private Endpoints

Azure resources can have different access, we can global public access, select public access (based on virtual networks and ip addresses) or private endpoint (specifically created)

</details>

## Azure Compute Demo
<details>
<summary>
Demos of creating Azure compute services
</summary>

### LIVE DEMO: Creating a Virtual Machine (VM)

We need a resource group. choosing the region, availability zone (or allowing azure to choose), OS image (server, desktop), or getting an image with the software already installed on it. for windows machines we need to create a user and password.

### LIVE DEMO: Connecting to a Virtual Machine

we can see public and private IPs, we can connect to it directly with RDP (remote-desktop-connection for windows) or SSH, we can make the windows server into a web server, (we might need to open up some ports). even when we close down the machine, we still pay for the storage of the disks, but that is a much lower cost.\
We can resize the machine into a stronger machine. and we can add more disk space to it if needed.

### LIVE DEMO: Creating Azure App Services / Web Apps

### LIVE DEMO: Azure App Services In Action

### LIVE DEMO: Creating Azure Functions

### LIVE DEMO: Kubernetes and Azure Container Instances



</details>

## Azure Storage
<details>
<summary>

</summary>
</details>

## Identity, Access and Security
<details>
<summary>

</summary>
</details>

## Cost Management
<details>
<summary>

</summary>
</details>

## Governance and Compliance
<details>
<summary>

</summary>
</details>

## Tools for Managing and Deploying Resources
<details>
<summary>

</summary>
</details>

## Monitoring Tools
<details>
<summary>

</summary>
</details>

## Practice Tests
<details>
<summary>

</summary>
</details>

## Takeaways
<details>
<summary>

</summary>
</details>


