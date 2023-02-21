<!--
// cSpell:ignore  Unmanaged
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
### Udemy Video Player
### FAQs




</details>

## Cloud Computing
<details>
<summary>
Understanding the basics of cloud computing.
</summary>

cloud - pay only for what you use.

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
</details>

## Benefits of Cloud Computing
<details>
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

</summary>
</details>

## Core Architectural Components
<details>
<summary>

</summary>
</details>

## Compute And Networking Services
<details>
<summary>

</summary>
</details>

## Compute Demo
<details>
<summary>

</summary>
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


