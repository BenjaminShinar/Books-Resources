<!--
// cSpell:ignore Unmanaged PaaS IaaS SaaS AZCOPY GZRS Passwordless Spendings functionapp eastus Kusto
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
<details>
<summary>
MISSING STUFF HERE
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

### Resources and Resource Groups

resources exist inside resource groups. resource is a generic term to an azure service we have access to, such as VM, storage accounts or databases. resources have a name, which sometimes should be globally unique. resources usually (but not always) are tied to a specific region. most resources have costs.

each resource is tied to a subscription, which determines who is the billing address.

Resource groups are associated with a region, but they can contain resources from other regions as well, all the resources in a resource group should be related to one another.\
we can assign permission at the resource group level, but resource groups do not force security boundaries, a resource isn't limited to interact with resources in the same group.

### Subscriptions

subscriptions are the billing unit of Azure, a subscription has a payment method associated with it. users can work against different subscriptions, and can have different roles for each of them, there are different subscription plans, such as free, pay as you go, etc...

usually, subscriptions are used to divide separate business units, such as divisions or geographical regions. they can also act as security boundaries, such as denying developers access to the finance subscription, or vice-versa.

### Management Groups

management group contain one or more subscriptions, and can contain nested management groups inside it. the top level is the "root" management group. management groups have azure policies (blueprints) that enforce security or limitations on resources inside the subscriptions.

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

we can see public and private IPs, we can connect to it directly with RDP (remote-desktop-protocol) for windows or SSH for linux, we can make the windows server into a web server, (we might need to open up some ports). 

Even when we close down the machine, we still pay for the storage of the disks, but that is a much lower cost.\
We can resize the machine into a stronger machine. and we can add more disk space to it if needed.

when we created the VM, a resource group with some resources is created, so we delete the resource group and everything is deleted.

### LIVE DEMO: Creating Azure App Services / Web Apps

another service available to us is the **WebApp**, it must have globally unique name, we choose between a code app, a static web app (just html) and a docker container app. for code apps, we need to choose the <kbd>runtime stack</kbd> (Java, node, python, .NET, ruby, JS, etc...). depending on the runtime, we might be limited to windows or linux operating system.\
we also choose the <kbd>region</kbd>, and the hosting power - measured in ACU. we can also add more features such as backup, and we can get zone redundancy for high availability or get application insight.

### LIVE DEMO: Azure App Services In Action

when we deploy the webapp, we see the URL (not the IP), there is no option to "connect" to it. The webapp starts with a default configuration. we can look at the monitoring tools to see metrics for response time and other metrics.\
If we see high load on our app, we can decide to scale up (use a stronger machine) or scale out (add more machine). we can also add **auto-scaling** rules to make this process happen without manual action.

Another option is adding **Deployment Slots** and run additional versions of the app on different URLs, or we can decide to route traffic between different instances to perform A/B testing or preview new versions. there are options for CI-CD from repos such as github or azure Repos. under <kbd>Networking</kbd> we can set restrictions on who can access the app.

while there is no option to connect with RDP or SSH into the webapp, there is the <kbd>Console</kbd> option which provides a web based terminal access.

just like before, we close up by deleting the resource group.

### LIVE DEMO: Creating Azure Functions

Azure function can be serverless of use the PAAS model. we again need globally unique name, we can choose a <kbd>runtime stack</kbd> or containers, we can only choose one kind of language, and we can't mix between them. we have to connect to a storage account, and we choose the OS and <kbd>Plan</kbd>. there are 3 plans: 
- app service plan
- consumption (serverless) 
- premium functions

the plan dictates which features are available and how they scale up. there are less options to choose from than we saw in VMs or WebApps.

The url is not a website. if we want functionality, we need to choose <kbd>Functions</kbd>, then <kbd>Create</kbd> and set up the trigger. each "code" function should be really small, the demo code allows for an http request with a personalized message. pricing is based on consumption (number of execution). 

Functions are good for small pieces of code that don't require a complete server. there are also **durable functions** which are more advanced and can call other functions and wait for the response from them.

### LIVE DEMO: Kubernetes and Azure Container Instances

many services have container options, such as web-app for containers.

Container Instances is a simplified option to deploy containers, it's not very scalable, but it's quick. Kubernetes is fit for large scale deployments.

Azure Containers don't require a global unique name, they require an image (either from a registry or a quick start template), and we need to set the networking type, name, and which ports are open.

Microsoft Container Registry is a storage space for images, the image is a self contained code with all the required dependencies, and it's portable between many vendors, not just Azure.

</details>

## Azure Storage
<details>
<summary>
Storage Solutions: Blob, Disk, File and Tables
</summary>

storage is one of the foundation technologies of cloud services, as many other services rely on it in order to function.

- Container (blob) storage
- Disk Storage
- File Storage
- Storage Tiers

Azure storage account - Blobs, Tables, Queues, Files, Azure Data Lake. general purpose storage.

there are more sophisticated options, but storage account is the cheapest one.\
BLOB - Binary Large Object. any file can be stored in blob.

- access tiers: Hot, Cool, Archive - how fast they data can be retrieved with cost tradeoff
- performance tiers - Standard or Premium
- location
- redundancy/replication - storing back up of the data in other regions
- failover options - using the backup before azure comes back online

Disk Storage is for Azure Virtual Machines, managed disks act as hard disks for our VMs, they are optimized for that purpose, and have different pricing model.
### LIVE DEMO: Create an Unmanaged Storage Account

in Azure, we choose <kbd>Storage Account</kbd>, <kbd>Create</kbd>, give it a globally unique name, choose the region, performance and redundancy tiers.

the region affects the cost, some regions are cheaper than others. Premium performance uses SSD disks so it has lower latency, and can be further optimized:
- Block blobs - high transaction rates and lw storage latency.
- File Shares - high performance applications that need to scale
- Page blobs - random read and write operations

Redundancy:
- LRS - Locally Redundant Storage - basic redundancy, backup in the same Data-Center, but on a different rack.
- ZRS - Zone Redundant Storage - backup at a different data-center (availability zone) in the same region
- GRS - backup at a different region, useful for data recovery and failover
- GZRS - combine zone and geo backups for high Availability and Disaster Recovery

using premium performance effects which types of redundancy is available.

we can control access and use encryption, enable or disable public access, or use Azure Active Directory. Data Lakes allow for big data antics, but it requires enabling the option and structuring the folders in a certain way.

We set the **Access Tier** as being used for hot Access or Cool storage, this depends on our usage pattern, archive tier is even cheaper to store data, but retrieving data can take hours.\
File storage acts like file servers, and we can mount it on our computer directly.

we set the networking options and who can access it, and we can also enable file versioning immutability, which might be required by regulation.

### LIVE DEMO: Upload Files to a Storage Account

The storage account contains the four types of data:
- Containers - blob
- File share
- Queues
- Tables - database

blob containers are simply a bucket for data, we create containers and organize files in them, each container can have different access policy. when we upload a file we choose the storage tier. each object has URL, and accessing requires a security key.

we can create AccessKeys to access all files, or use Shared Access Signature to provide access to a single object for a limited time.

under lifecycle management, we can move object between storage tiers or delete it completely based on custom rules.

### LIVE DEMO: Azure Storage Explorer & Storage Browser

we shouldn't really interact with Azure Storage manually in the day to day scenarios, it should be done in a programmatic way from our applications, or use **Storage Explorer**, which is a simplified view of the storage account. we can use it the browser or download an application to connect to it locally. this service is deprecated and will eventually be removed from Azure.\
**Storage Browser** is in preview mode and is the recommended way to interact with the data.

### AZCOPY

if we want to copy data between containers, we don't have to manually download and reupload the data (which would take time and cost us money in networking bandwidth). the better way is to copy the data inside the Azure ecosystem, this is done with a CLI tool called **AZCOPY**, which we can download to our local machine or use directly from the azure cloud shell. we can also use it to download files.

1. we create SAS from the source container with read permissions. 
2. we create SAS from the destination container with write permissions.
3. inside the cloud shell console we run the command
```ps
azcopy ? # show help
azcopy copy '<source url>' '<destination url>'
```

this will also work across subscriptions, because we got the permissions through the Shared Access Signature.

### Azure File Sync
Many companies still have File servers, so Azure has cloud solution called "File Share". the metric is the IO/s operations per second. we interact with them using the SMB protocol, so we can mount them on our machine.

- Storage Sync Service - organize groups and end point
- File Sync Agent - synchronize between the on-premises file server and the cloud

### Azure Migrate
A tool for creating a migration process from the on-premises servers and workloads into the cloud. it discovers local machines and services, maps them into the correct Azure service, and notifies if something requires action (updating) before it can be migrated.

### Azure Data Box

When we have a large amount of data (TerraBytes of data), it's no longer practical to upload it over the internet. it's possible to get an AzureExpressRoute private connection, but that would take time to set up.

Azure provides a way to send data in a physical device.

Name | Storage | Notes
------|---|---
Data Box Disk | 8 TB | SSD disk
Data Box | 100 TB | Computer
Data Box Heavy | 1 PB | Wheeled Version

we can also use Azure Stack Edge (either physical or virtual) which also help.

the data is encrypted, and the disks are wiped clean after the upload

</details>

## Identity, Access and Security
<details>
<summary>
Identity and Security tools
</summary>

security and identification

### Identity and Azure Active Directory

Identity can refer to a person, application or a device, it's a digital representation. it usually requires a password, secret key or certificates to assume the identity.

traditional models of client-server would rely on all sorts of identification methods, some were handmade, some used cookies, and in many cases the were hacks and exploits:
- storing passwords in plain text
- using a simple, reversible hash algorithm such as MD5
- storing the "salt" (random element) along with the data
- not enforcing password change policies
- not enforcing password complexity policies

Azure has an identity management system based on Active Directory, which is very popular in the corporate world. in Azure it's called Azure Active Directory. but it's not the same.

Traditional AD does not work with Internet protocols.

Azure AD uses "identity as a service" model.  it removes the need to write code to handle users, passwords, password reset, etc...

- **Client** uses UserId and password to connect to the **Identity Provider** and receives a signed token.
- this signed token allows the **Client** to connect to the **Server**.
- The **Server** and the **Identity Provider** have a trust-key relationship.

protocols:
- SAML
- OpenID
- WS Federation

### Benefits of Azure Active Directory (Azure AD)

Security: Azure AD knows what they do, Has Support tools, it has many features which are widely used, and it provides a centralized administration control over all application. \
The company doesn't need to develop everything from scratch.

User has one User Id and one password - Single Sign-On.

Azure services integrate with AzureAD, even across azure accounts, which again reduces the amount of identity management.

### Authentication vs Authorization
- Authentication - the user proving who they are.
- Authorization - ensuring that a user is permitted to perform an action. what can the user do.

not every user should have administrator access and privileges. this creates problems, people can make mistakes and bad actors can wreck havoc.

### Azure AD Conditional Access

A common case of hacking is when the credentials of one user are exposed, and then the hacker uses them. in most cases, the hacker is not inside the office building or even the same country, and has very different usage patterns than the user who's identity was stolen.

**Conditional Access** assigns levels of trust based on different conditions, such as location, time, time since last log-in (if a user attempts to log in after months of inactivity), device. some access attempts are "routine" and some are not, so conditional access allows to enforce extra limitations in different cases.

### Multi-Factor Authentication (MFA or 2FA)

MultiFactor authentication requires more evidence to log in. it allows for more security in case one factor is compromised.
- something you know - password
- something you have - access to mobile phone, email account, authentication App.
- something you are - fingerprint scan, retina scan, photograph.

in azure, we can use SMS, email, authenticator app or phone call as an identity factor.

### PasswordLess

Passwordless is a new way to perform authentication without causing inconvenient and without compromising security 

X | Low Security |High Security
-----|----|-
InConnivent | &empty; | Passwords + 2FA
Connivent| Passwords | PasswordLess

Phone gestures are one type of PasswordLess authentication, it's tied to the device, so it can't be used elsewhere. face identification is another form, or asking you to verify who you are on a different device.

### Role-Based Access Control (RBAC)

this deals with authorization. Role based as opposed to "claim based".

the administrator creates roles that represent common tasks, and those roles get permissions according to what they need. The permissions can be granular.

Users are then assigned to roles, and they get those permissions, this reduces the complexity of handling permissions across many users.

in azure, there are usually some built-in roles which represent common operations. 
- reader - read only
- contributor - read and write
- owner - assign permission

we can always create custom permissions as needed.

### Zero-Trust Model of Security

in the traditional paradigm, there is a single security threshold (the firewall). in a zero trust model, there are multiple security threshold, and services don't trust each other, everything is guarded.

Zero Trust principles
- Verify Explicitly
- Use Least Privileged Access
- Assume Breach

- Just-in-Time - request permissions when needed, not holding them all the time.
- Just-Enough-Access - request and grant only the required permissions and nothing more

have security detection for suspicious behavior, auditing, segmentation (not everything can do everything). we should also ensure compliance for devices and make sure they are up to date. Data should be encrypted, segmented and monitored.

The network should also monitor for risky behavior.

### Defense in Depth

Having multiple layers of defense, protecting each layer of the stack and limiting blast radius.

- Data
- Application
- Compute
- Network
- Perimeter
- Identity and Access
- Physical

### Microsoft Defender for Cloud

A paid service for security posture management, it does analysis, gives a rating and recommendation for security, follow regulation, firewalls and other security products.


</details>

## Cost Management
<details>
<summary>
Understanding, Estimating and Monitoring Spendings
</summary>

### Factors that Affect Cost

each service has different cost calculations. some services are free:
- Resource groups
- Virtual network (up to 50)
- Load balancer (basic)
- Azure Active Directory (basic)
- Network security groups
- Free-tier web apps (up to 10)

other services have metrics that effect the costs.

the consumption model is based on innovations, such as Azure functions. the first million executions are free, and then it's another 20&cent; for each million executions. compared to the cost of running the cheapest virtual machine, which is about 20$ per month.

Pay per usage Services:
- Azure functions
- Logic Apps
- Storage (pay per GB)
- Outbound Bandwidth
- Cognitive Services API

Other services are billed by time, such as Virtual machines. so when the machine stops, so does the billing.

the pricing isn't stable, and each month can have a different cost. if stability is important, then there are options to reserve resources and in those cases the pay is the same even if they aren't used. (there might be discounts).

outbound Bandwidth is also charged, the first 5GB going out are free, in boundData is free, transferring data inside a region is free, but there are costs to transfer data between regions.

### Azure Pricing Calculator

The Pricing Calculator is a tool that helps us estimate the costs based on the resources we use. we set up the services and configuration, and we can see an estimate of how much using azure will cost.

we start by adding the services, and then choose which configurations we use and the usage metrics. we set the savings plan we plan to use. we can do this for each resource and create a estimate. 

### Total Cost of Ownership Calculator

The Total Cost of Ownership calculator focus on costs which aren't just paying azure for cloud resources.

The cost of running a Server on-premises isn't just the hardware itself, we need to take into account everything else: form electricity, cooling, networking, backup, maintenance labor... those costs are recurring.

The TCO calculator is used to estimate the relative cost of running the workloads on premises compared to running it on the cloud. we describe the workloads we use, and then set the assumptions for costs. the calculator will use those assumptions to compare between staying on-premises or migrating to the cloud.

### Azure Cost Management
The cost management service is a tool to view and analyze spending in azure. it tracks spending across time, regions, services, and it can allows to set a budget and have an alert if the costs exceed a threshold.
we can use those alerts to set automatic actions.

we can also create monthly costs reports.

### Resource Tags

Resource tags are meta-data on resources that we can set with custom information. they can help with billing and support issues, we can generate billing reports by tags, so we can track what resource belongs to which project, and who is the owner.

at any resource, we can click <kbd>Tags</kbd> and start adding key-value pairs.

</details>

## Governance and Compliance
<details>
<summary>
Governance and Compliance guidelines.
</summary>

Governance and Compliance are guidelines which must be followed for resources. we can set up rules (policies) that define how resources should behave. Azure has tools to enforce policies or to audit resources which don't follow them.

- Azure Blueprints
- Azure Policy
- Resource Locks
- Service Trust Portal

Blueprints are templates for subscription with roles and policies already defined. Policies are rules which act across azure resources, such as requiring all databases to have backup enabled, the policy can be used to enforce the behavior and block resources which don't follow the rule from being created, or evaluate resources according to the policy and report them.

There are built in policies, such as:
- requiring minimal SQL servers version
- limiting storage account SKUs
- limiting virtual machine SKUs
- limiting locations
- requiring specific tags
- disallowing certain resource types

Policies can be defined in json format. they can be defined per subscription or resource group.

Resource locks are a way to mark resources as being immutable (can't be modified) or to mark it as a resource which can not be deleted. locks should be managed by RBAC and only a limited number of roles should be able to add and remove them.

The microsoft Service Trust Portal holds all the documents about how azure is compliant with regulatory demands, and how it's security is tested and assessed.

</details>

## Tools for Managing and Deploying Resources
<details>
<summary>
Ways to work and manage resources in Azure
</summary>

Azure CLI
PowerSeel
Azure Portal
Azure CloudShell
Azure Mobile App

### Azure Portal and Command Line Tools

The Azure portal is the web management tool, it's the website and the GUI that we use.

Azure supports PowerShell and CLI (bash) SDK, which acts as a library for the scripting language. Those same commands are also available through the CloudShell in the portal, which comes in both PowerShell and Bash flavours. this allows us to work with scripts.

```sh
az functionapp list
az webapp list
```

using the cli on a local machine requires logging in, but using the cloudShell skips this step.

### LIVE DEMO: Create Resources Using Command Line

in the cloudShell, we create a resource group with a virtual machine

```sh
az group create --name newRg --location eastus
az vm create --resource-group newRg --name newVm --image Win2019DataCenter --public-ip-sku Standard --admin -user-name adminUser
# provide password and wait for the machine to load
az vm open-port --port 80 --resource-group newRg --name newVm
az group delete --name newRg
```

### Azure Arc

Azure arc allows us to see and manage on-premises resources and even on other cloud vendors.

it's mostly used for infrastructure, such as servers (compute machines), sql server databases, kubernetes clusters and so on.

Azure Stack refers to a private cloud, with HCI referring to edge machines which are outside of azure.

### ARM Templates

ARM stands for Azure Resource Manager, this is the core of azure, all services which run with azure (portal, cli, sdk) go thorough ARM. the commands are converted into an ARM template which describe the resource.

### LIVE DEMO: Generate ARM Templates in the Azure Portal

In the demo, we create a storage account, and instead of creating, we download the template. the template is both the template and the parameters, and we can then run them from our local machine.

there is also a resource called template spec, which can be stored and then we have template spec we can deploy later.

</details>

## Monitoring Tools

<details>
<summary>
Monitoring logs, metrics and resource health
</summary>

### Azure Advisor and Azure Service Health

Azure advisor analyzes the resource used, and creates recommendation based on usage profiling, such as downgrading a VM if it's under utilized. it also detects open ports such as RDP. it will also recommend us ways to decrease cost and improve resiliency.

The Azure Service Health is concerned the health of azure in general, rather than the specific Azure account. it gives us the ability to know if there are problems in azure and how they effect out resources. for example, we can get an alert if we have resources in regions that are experiencing outages or undergoing maintenance.

### Azure Diagnostics Settings

Diagnostic settings is a storage of the metrics from our resources, we can save them into a storage account, or send them to the log analytics workspace, an event hub or a 3rd party solution.

there is a "classic" diagnostic service which is being phased out. 

### Azure Monitor
An dashboard that gives an overview about our resources, for most accounts, we can set alerts based on all sorts of queries, we can look at metrics and save those reports to the dashboard. there is a small charge for each rule and for the notification itself.

there are built in reports for each type of resource, the queries are using the Kusto Query Language syntax. some queries can also create visualizations. we can gather them together into workbooks to create customized reports.

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

ACU - Azure Compute Unit - a relative measurement of compute power.
ACI - Azure Containers Instance

### Azure and AWS

<details>
<summary>
which services in azure are comparable to aws services
</summary>

| Service Purpose         | Azure Name         | AWS Name                    | Notes                                     |
| ----------------------- | ------------------ | --------------------------- | ----------------------------------------- |
| Virtual Machine         | Azure VM           | EC2                         | compute                                   |
| General Storage         | Blob storage       | S3                          | object / blob storage                     |
| General Storage Tiers   | Hot, Cool, Archive | S3 Standard, IA, Glacier    | storage tiers for costs and retrieval time |
| Event-Driven Serverless | Azure Functions    | Lambda                      |
| Message Queue           | Queue storage      | SQS                         | asynchronous message handling              |
| VM disk volume          | Azure Disks        | EBS - elastic block storage | storage volumes for virtual compute       |
| File Storage            | Azure Files        | EFS - elastic file system   |
| Physical data transfer  | Azure Data Box     | AWS SnowBall                | Data migration                            |
| Resource Creation       | Azure Blueprints   | AWS CloudFormation          | create resources from templates           |
Containers | Azure Container Instances | (not ECS) | simple containers
Kubernetes | AKS | EKS | managed kubernetes cluster
Object Storage | Blob Container | S3 Bucket | data storage
Accessing Object Storage | Shared Access Signature | PreSigned Url | granular, time limited access


</details>


