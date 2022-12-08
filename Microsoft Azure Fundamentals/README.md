<!--
// cSpell:ignore PAAS
 -->

# Microsoft Azure Fundamentals

[Microsoft Azure Fundamentals:](https://learn.microsoft.com/en-us/training/paths/az-900-describe-cloud-concepts/), [az-900 exam](https://learn.microsoft.com/en-us/certifications/exams/az-900)

each learning path corresponds to one domain of the exam. a learning path is composed of several modules.

AZ-900 Domain Area | Weight
---|---
Describe cloud concepts |20-25%
Describe core Azure services | 15-20%
Describe core solutions and management tools on Azure |10-15%
Describe general security and network security features | 10-15%
Describe identity, governance, privacy, and compliance features| 20-25%
Describe Microsoft cost management and Service Level Agreements|10-15%

> New to the cloud? Azure fundamentals is a six-part series that teaches you basic cloud concepts, provides a streamlined overview of many Azure services, and guides you with hands-on exercises to deploy your very first services for free.\
> After completing this learning path, you'll be able to:
>
> - Understand the benefits of cloud computing in Azure and how it can save you time and money
> - Explain cloud concepts such as high availability, scalability, elasticity, agility, and disaster recovery
> - Describe core Azure architecture components such as subscriptions, management groups, resources and resource groups
> - Summarize geographic distribution concepts such as Azure regions, region pairs, and availability zones

## Azure Concepts
<details>
<summary>
Basic cloud concepts, Azure concepts and base components.
</summary>

### Introduction to Azure Fundamentals
<details>
<summary>
This module introduces you to the basics of cloud computing and Azure, and how to get started with Azure's subscriptions and accounts.
</summary>

#### What is Cloud Computing?

cloud service: servers, sotrage, database, networking, software, analytics, intelligence.

pay only for the services used, and have them being managed. we can also easily scale if the requirements change.

by using the cloud, we can reduce the development cycle, focus on the core business and not the infrastructure, and integrate with other cloud services.

#### What is Azure?

Azure is microsoft's cloud platform:

- IAAS - infrastructure as a service
- PAAS - platform as a service
- SAAS - software as a service

compute, storage, event driven azure functions, many types of relation databases. integration with on premises data centers (hybrid environment).

azure uses virtualization, using a hypervisor on top of hardware to abstract running many instances of different machines.\
in their data centers, each rack runs multiple machines with this hypervisor, in addition, each rack has a **Fabric Controller**, connected to the **orchestrator**, which are the managemet and networking points of the rack. the user interacts with the orchestrator, and through it to the fabric controller.

The **Azure Portal** is the web portal, a simple web console that can do everything we need, such as creating resources, and do operations on the resources. we can create custom dashboards and monitor our applications.\
The portal is replicated in each azure data center, so it's designed for resiliency and high availability.

The azure marketplace offers us to shop for services which were built for running on azure, and were verified to run on azure.

#### Tour of Azure Services

The commonly used categories:
- Compute
- Networking
- Sotrage
- Mobile
- Databases
- Web
- Internet of Things (IoT)
- Big Data
- AI (Artificial intelligence)
- Dev Ops

> Compute:\ 
> Compute services are often one of the primary reasons why companies move to the Azure platform. Azure provides a range of options for hosting applications and services. Here are some examples of compute services in Azure.
> 
> Networking:\
> Linking compute resources and providing access to applications is the key function of Azure networking. Networking functionality in Azure includes a range of options to connect the outside world to services and features in the global Azure datacenters.
>
> Storage:\
> Azure provides four main types of storage services.
> - Disc
> - Blob
> - File
> - Archive
> These services all share several common characteristics:\
> - Durable and highly available with redundancy and replication.
> - Secure through automatic encryption and role-based access control.
> - Scalable with virtually unlimited storage.
> - Managed, handling maintenance and any critical problems for you.
> - Accessible from anywhere in the world over HTTP or HTTPS.
> 
> Mobile:\
> With Azure, developers can create mobile back-end services for iOS, Android, and Windows apps quickly and easily. Features that used to take time and increase project risks, such as adding corporate sign-in and then connecting to on-premises resources such as SAP, Oracle, SQL Server, and SharePoint, are now simple to include.\
> 
> Other features of this service include:
> - Offline data synchronization.
> - Connectivity to on-premises data.
> - Broadcasting push notifications.
> - Autoscaling to match business needs.
> 
> Databases:\
> Azure provides multiple database services to store a wide variety of data types and volumes. And with global connectivity, this data is available to users instantly.
> 
> Web:\
> Having a great web experience is critical in today's business world. Azure includes first-class support to build and host web apps and HTTP-based web services. The following Azure services are focused on web hosting.
> 
> IoT:\
> People are able to access more information than ever before. Personal digital assistants led to smartphones, and now there are smart watches, smart thermostats, and even smart refrigerators. Personal computers used to be the norm. Now the internet allows any item that's online-capable to access valuable information. This ability for devices to garner and then relay information for data analysis is referred to as IoT.\
> Many services can assist and drive end-to-end solutions for IoT on Azure.
> Big data:\
> Data comes in all formats and sizes. When we talk about big data, we're referring to large volumes of data. Data from weather systems, communications systems, genomic research, imaging platforms, and many other scenarios generate hundreds of gigabytes of data. This amount of data makes it hard to analyze and make decisions. It's often so large that traditional forms of processing and analysis are no longer appropriate.\
> Open-source cluster technologies have been developed to deal with these large data sets. Azure supports a broad range of technologies and services to provide big data and analytic solutions.
> 
> AI:\
> AI, in the context of cloud computing, is based around a broad range of services, the core of which is machine learning. Machine learning is a data science technique that allows computers to use existing data to forecast future behaviors, outcomes, and trends. Using machine learning, computers learn without being explicitly programmed.\
> Forecasts or predictions from machine learning can make apps and devices smarter. For example, when you shop online, machine learning helps recommend other products you might like based on what you've purchased. Or when your credit card is swiped, machine learning compares the transaction to a database of transactions and helps detect fraud. And when your robot vacuum cleaner vacuums a room, machine learning helps it decide whether the job is done.\
> A closely related set of products are the cognitive services. You can use these prebuilt APIs in your applications to solve complex problems.
> 
> DevOps:\
> DevOps brings together people, processes, and technology by automating software delivery to provide continuous value to your users. With Azure DevOps, you can create build and release pipelines that provide continuous integration, delivery, and deployment for your applications. You can integrate repositories and application tests, perform application monitoring, and work with build artifacts. You can also work with and backlog items for tracking, automate infrastructure deployment, and integrate a range of third-party tools and services such as Jenkins and Chef. All of these functions and many more are closely integrated with Azure to allow for consistent, repeatable deployments for your applications to provide streamlined build and release processes.

#### Get Started with Azure Accounts

the hierarchy starts with an **Azure account**, which can have **Subscriptions**, which in turn have **Resource groups**, and under them, there are **resources**. subscriptions can grouped into invoices groups, and we can set multiple billing profiles in the same billing account.

The free Azure account has:
> - Free access to popular Azure products for 12 months.
> - A credit to spend for the first 30 days.
> - Access to more than 40 products that are always free.

The free student account has:
> - Free access to certain Azure services for 12 months.
> - A credit to use in the first 12 months.
> - Free access to certain software developer tools.

when learning, we can use the sandbox.

#### Introducing the Case Study
for the training, we will work on a fictional cloud environment that belongs to the fake company "tailwind traders". the company runs a retail website.\
They have an on premises data center with data and video. currently, the IT team runs everything, but in the course of the fundamental course, we will see which parts of their work can be offloaded to azure.

</details>

### Azure Fundamental Concepts
<details>
<summary>
This module introduces you to the basics of cloud computing and Azure, and how to get started with Azure's subscriptions and accounts.
</summary>

Azure (and other cloud computing services) allow for increase in availability, scalability, and help make deployment faster.\
We can 

####  Different Types of Cloud Models
there are three configurations of cloud deployment models: public, private and hybrid.\

in a public cloud, the computing machines and storage are located in the datacenters of the hosting company, there is no capital (monterey) barrier to start or to expand, and all access is done via the internet.\
A private cloud is exclusively used by the purchasing company, and it can be hosted either on premises or at a 3rd part site. all hardware is purchased by the company, and maintained by them, this gives the company complete control over security.\
a hybrid cloud deployment combines the two, allowing the company to control which servies run on the public cloud an which run on the private cloud (i.e. for legal reasons and compliance with regulations).

> **Public cloud:**\
> Services are offered over the public internet and available to anyone who wants to purchase them. Cloud resources, such as servers and storage, are owned and operated by a third-party cloud service provider, and delivered over the internet.
> 
> **Private cloud:**\
> A private cloud consists of computing resources used exclusively by users from one business or organization. A private cloud can be physically located at your organization's on-site (on-premises) datacenter, or it can be hosted by a third-party service provider.
> 
> **Hybrid cloud:**\
> A hybrid cloud is a computing environment that combines a public cloud and a private cloud by allowing data and applications to be shared between them.

####  Cloud Benefits and Considerations

Advantages of cloud computing (azure):

> - High availability: Depending on the service-level agreement (SLA) that you choose, your cloud-based apps can provide a continuous user experience with no apparent downtime, even when things go wrong.
> - Scalability: Apps in the cloud can scale vertically and horizontally:
>   - Scale vertically to increase compute capacity by adding RAM or CPUs to a virtual machine.
>   - Scaling horizontally increases compute capacity by adding instances of resources, such as adding VMs to the configuration.
> - Elasticity: You can configure cloud-based apps to take advantage of autoscaling, so your apps always have the resources they need.
> - Agility: Deploy and configure cloud-based resources quickly as your app requirements change.
> - Geo-distribution: You can deploy apps and data to regional datacenters around the globe, thereby ensuring that your customers always have the best performance in their region.
> - Disaster recovery: By taking advantage of cloud-based backup services, data replication, and geo-distribution, you can deploy your apps with the confidence that comes from knowing that your data is safe in the event of disaster.

when we consider the costs, there are two types of costs to consider, Capital Expenditure (CapEx) and Operational Expenditure (CapOx).
> - Capital Expenditure (CapEx) is the up-front spending of money on physical infrastructure, and then deducting that up-front expense over time. The up-front cost from CapEx has a value that reduces over time.
> - Operational Expenditure (OpEx) is spending money on services or products now, and being billed for them now. You can deduct this expense in the same year you spend it. There is no up-front cost, as you pay for a service or product as you use it.

Buying infrastucure is considered CapEx, it is paid for upfront, in goes into the accounting books as an asset, the value of which decreases over time. Purchasing services from the cloud provider is considered OpEx, it effects the operations balance sheet of the company, as well as the net-profit and taxable income.

> Cloud Computing is consumption-based model, which means that end users only pay for the resources that they use. Whatever they use is what they pay for.\
> A consumption-based model has many benefits, including:
> - No upfront costs.
> - No need to purchase and manage costly infrastructure that users might not use to its fullest.
> - The ability to pay for additional resources when they are needed.
> - The ability to stop paying for resources that are no longer needed. 

####  Different Cloud Services

Cloud services can be divided into different service models, based on how the responsability of managing them is distributed between the provider and the user. on one end, there is IaaS (Infrastructure-as-a-Service), which gives the user the most control over the resources, while at the other end there is SaaS (Software-as-a-Service), which is the most managed services where the provider takes care of most of the work, between them there is PaaS (Platform-as-a-Service).

Acronym |Full Name | Provider responsability | User  responsability | Example
---|---|---|---|---|
Iaas | Infrastructure-as-a-Service |hardware| perating system maintenance, network configuration  |Azure virtual Machines
PaaS| Platform-as-a-Service |Virtual machines, networking| Applications | Azure App Services
SaaS | Software-as-a-Service|Virtual machines, networking, data storage, applications| Application data| Microsoft Office 365 

in IaaS, you simply rent the infrastructure instead of buying it, everything else is like on premises data center, the customer has the most control over the resources, and can tailor them to its' needs. in PaaS, the cloud provder has more responsability, which means the client can focus on the application development. in SaaS, the client only provides the data to the application, but the application (software) is provided by the cloud host. so this is the easiest option to use, and it provides access to the most up-to-date software, but it limits what kinds of applications the client can run.

Model | Storage | networking| Compute | virtual machine| operating System | runtime | Applications|Data & Access
---|---|---|---|---|---|---|---|---|
On-Premises (private cloud) | Customer Managed |Customer Managed|Customer Managed|Customer Managed|Customer Managed|Customer Managed|Customer Managed|Customer Managed
IaaS | Cloud Provided |Cloud Provided|Cloud Provided|Customer Managed|Customer Managed|Customer Managed|Customer Managed|Customer Managed
PaaS | Cloud Provided |Cloud Provided|Cloud Provided|Cloud Provided|Cloud Provided|Cloud Provided|Customer Managed|Customer Managed
SaaS | Cloud Provided |Cloud Provided|Cloud Provided|Cloud Provided|Cloud Provided|Cloud Provided|Cloud Provided|Customer Managed


**Serverless computing** is like PaaS in a way, as it allows the user to build applications on their own, and not rely on the software the cloud provider supports. in this case, the cloud automatically provisions and manages the resources behind the scenes, which makes serverless computing very scalable. serverless computing tends to be event-driven.

</details>

### Core Azure Architectural Components
<details>
<summary>
In this module, you'll examine the various concepts, resources, and terminology that are necessary to work with Azure architecture. For example, you'll learn about Azure subscriptions and management groups, resource groups and Azure Resource Manager, as well as Azure regions and availability zones.
</summary>

getting an understanding of azure's terminology.

the hierarchy begins at the management group, which has subscriptions, which have resource groups, which contain azure resources.

> - Management groups: These groups help you manage access, policy, and compliance for multiple subscriptions. All subscriptions in a management group automatically inherit the conditions applied to the management group.
> - Subscriptions: A subscription groups together user accounts and the resources that have been created by those user accounts. For each subscription, there are limits or quotas on the amount of resources that you can create and use. Organizations can use subscriptions to manage costs and the resources that are created by users, teams, or projects.
> - Resource groups: Resources are combined into resource groups, which act as a logical container into which Azure resources like web apps, databases, and storage accounts are deployed and managed.
> - Resources: Resources are instances of services that you create, like virtual machines, storage, or SQL databases.

#### Regions, Availability Zones, and Region Pairs

> A region is a geographical area on the planet that contains at least one but potentially multiple datacenters that are nearby and networked together with a low-latency network. Azure intelligently assigns and controls the resources within each region to ensure workloads are appropriately balanced.

There are also *special Azure regions*, which are used for legal and compliance purposes, some are US-based and some China based.

Each region is comprised of availability zones, an availability zone is a data center. if one availability zone goes down, the others aren't effected. availability zones are connected to one another through a private, high speed, fiber-optic network. we can run applications in different availability zones for high availability or in the same availability zone for high performance.

services can be zonal, zonal-redundant, and non-regional.
- a zonal service is pinned to an specific availability zone. such as a virtual machine or a managed disc.
- a zonal-redundant service is replicated automatically across availability zones. such examples might be a database.
- a non-regional service is not tied to a specific region, and should be available even if the entire region goes down.

in addition to that, Azure has the concept of *Geography*, which contains *region pairs*. a geography is a conceptual location (such as asia, europe, us), and a region pair is a pairing of two regions, which are at a sufficient distance from one another but (mostly) reside in the same state. region pairs are preferred for cross region data redundancy, and azure keeps each region in the region pair on a different maintenance schedule, and will prioritize fixing one of the regions if a large scale disaster happens. this means that even if a disaster happens or an azure update goes horribly wrong, the data from one region can be preserved in the other region.

#### Resources and Azure Resource Manager

> - Resource: A manageable item that's available through Azure. Virtual machines (VMs), storage accounts, web apps, databases, and virtual networks are examples of resources.
> - Resource group: A container that holds related resources for an Azure solution. The resource group includes resources that you want to manage as a group. You decide which resources belong in a resource group based on what makes the most sense for your organization.

All resources resign in a resource group, and a resource can be a member of only one resource group. resources can be moved between resource groups (under some conditions), there is not nesting of resource groups. Resource groups allow for a logical grouping of resources based on the organization needs. Resource groups also control **lifecycle** - when a resource group is removed (deleted), all the resources inside it are also removed. in addition, we also use resource groups as an **authorization** scope, when we give role based access control (RBAC) permissions we can allow access to resources in a specific resource group.

> **Azure Resource Manager**:\
> Azure Resource Manager is the deployment and management service for Azure. It provides a management layer that enables you to create, update, and delete resources in your Azure account. You use management features like access control, locks, and tags to secure and organize your resources after deployment.

all azure actions (via the web console, apis, sdk, etc...) are processed via the Azure Resource Manager, so the results are always consistent no matter how the user interacts with azure.

> With Resource Manager, you can:
> - Manage your infrastructure through declarative templates rather than scripts. A Resource Manager template is a JSON file that defines what you want to deploy to Azure.
> - Deploy, manage, and monitor all the resources for your solution as a group, rather than handling these resources individually.
> - Redeploy your solution throughout the development life cycle and have confidence your resources are deployed in a consistent state.
> - Define the dependencies between resources so they're deployed in the correct order.
> - Apply access control to all services because RBAC is natively integrated into the management platform.
> - Apply tags to resources to logically organize all the resources in your subscription.
> - Clarify your organization's billing by viewing costs for a group of resources that share the same tag.

it also allows us to protect resources from accidental deletion by applying policies and resource locks. we can also use Azure policies to ensure the same tags are used.

#### Subscriptions and Management Groups

> Azure subscriptions:\
> Using Azure requires an Azure subscription. A subscription provides you with authenticated and authorized access to Azure products and services. It also allows you to provision resources. An Azure subscription is a logical unit of Azure services that links to an Azure account, which is an identity in Azure Active Directory (Azure AD) or in a directory that Azure AD trusts.

an azure Account has subscriptions, a subscription can have a *billing models* and an *access-management/control* policy (model). a billing model groups subscriptions into a different reports or invoices, which are required for managing costs. the access management defines the boundaries of which resources can be provisioned.

common ways to organize subscriptions is by geographical environment (which also helps with compliance), by the organizational structure (which helps with managing permissions) and by billing purposes (which helps track costs).

subscriptions also have limits, such as maximum number of Azure ExpressRoute circuits (10 per subscription), so a new subscription can allow for more of those limited resources.

**Billing profiles** are organized under a single **billing account**, a profile is matches an invoice and has payment method. inside each billing profile, we can designate different subscriptions a invoice sections.

**Azure management groups** provide a level of scope above that of subscriptions, and are used to manage access, policies and compliance in a single location. management groups can be nested (contain other management groups) and they contain the subscriptions. permissions are managed at the management group and are inherited by the nested groups and the subscriptions. all subscriptions in a management group must trust the same Azure AD tenant.

management group policies can limit what can be done by the subscriptions inside it, and helps manage all the permissions for different users.

> Important facts about management groups
> - 10,000 management groups can be supported in a single directory.
> - A management group tree can support up to six levels of depth. This limit doesn't include the root level or the subscription level.
> - Each management group and subscription can support only one parent.
> - Each management group can have many children.
> - All subscriptions and management groups are within a single hierarchy in each directory.
</details>

</details>



## Azure Services
<!-- <details> -->
<summary>

</summary>

### Architectural Components of Azure
<!-- <details> -->
<summary>
This module explains the basic infrastructure components of Microsoft Azure. You'll learn about the physical infrastructure, how resources are managed, and have a chance to create an Azure resource.
</summary>

#### What is Microsoft Azure

Azure is microsoft's cloud computing solution, which provides IaaS, PaaS and SaaS options, as well as scalability. security and high availability options.

azure can provide simple virtual machines, databases, and even AI and machine learning services.

#### Get Started with Azure Accounts

the top level of azure is the azure account (which hold subscriptions), for free accounts, there are two levels:
- free account (free credit, access to free services, one year access to popular paid services)
- free student account (more time to use the credit, access to some more services).

the microsoft learn sandbox is a temporary subscription in the azure account, which cleans up all the created resources after the session is complete, and provides the resources for free.

#### Exercise - Explore the Learn Sandbox
#### Describe Azure Physical Infrastructure
#### Describe Azure Management Infrastructure
#### Exercise - Create an Azure Resource

</details>

### Compute and Networking Services
<details>
<summary>
This module focuses on some of the computer services and networking services available within Azure.
</summary>
</details>

### Storage Services
<details>
<summary>
This module introduces you to storage in Azure, including things such as different types of storage and how a distributed infrastructure can make your data more resilient.
</summary>
</details>

### Identity, Access, and Security
<details>
<summary>
This module covers some of the authorization and authentication methods available with Azure.
</summary>
</details>


</details>

## Solutions and Management Tools on Azure
<details>
<summary>

</summary>
</details>

## General Security and Network Security Features
<details>
<summary>

</summary>
</details>

## Identity, Governance, Privacy, and Compliance Features
<details>
<summary>

</summary>
</details>

## Microsoft Cost Management and Service Level Agreements
<details>
<summary>

</summary>
</details>


## Takeaways
<!-- <details> -->
<summary>

</summary>

- Azure Web Apps - scalable host websites
- Azure Functions - event driven actions.
- Container Instance
- Kubernetes Services
- Cosmos DB - noSQL
- Azure Portal
- Azure Resource Manager
- Azure AD - Active directory

### Azure Services
<details>
<summary>
Table of Services in Azure
</summary>

Service name | Service function | Section
---|--- |---
Azure Virtual Machines |Windows or Linux virtual machines (VMs) hosted in Azure. | Compute
Azure Virtual Machine Scale Sets | Scaling for Windows or Linux VMs hosted in Azure. | Compute
Azure Kubernetes Service | Cluster management for VMs that run containerized services. | Compute
Azure Service Fabric | Distributed systems platform that runs in Azure or on-premises. | Compute
Azure Batch | Managed service for parallel and high-performance computing applications.| Compute
Azure Container Instances |  Containerized apps run on Azure without provisioning servers or VMs. | Compute
Azure Functions | An event-driven, serverless compute service. | Compute
Azure Virtual Network |Connects VMs to incoming virtual private network (VPN) connections. | Networking
Azure Load Balancer |Balances inbound and outbound connections to applications or service endpoints. | Networking
Azure Application Gateway | Optimizes app server farm delivery while increasing application security. | Networking
Azure VPN Gateway | Accesses Azure Virtual Networks through high-performance VPN gateways.  | Networking
Azure DNS | Provides ultra-fast DNS responses and ultra-high domain availability. | Networking
Azure Content Delivery Network |
Delivers high-bandwidth content to customers globally.  | Networking
Azure DDoS Protection | Protects Azure-hosted applications from distributed denial of service (DDOS) attacks. | Networking
Azure Traffic Manager| Distributes network traffic across Azure regions worldwide. | Networking
Azure ExpressRoute| Connects to Azure over high-bandwidth dedicated secure connections. | Networking
Azure Network Watcher| Monitors and diagnoses network issues by using scenario-based analysis. | Networking
Azure Firewall| Implements high-security, high-availability firewall with unlimited scalability. | Networking
Azure Virtual WAN|Creates a unified wide area network (WAN) that connects local and remote sites. | Networking
Azure Blob storage | Storage service for very large objects, such as video files or bitmaps. | Storage
Azure File storage | File shares that can be accessed and managed like a file server. | Storage
Azure Queue storage |A data store for queuing and reliably delivering messages between applications. | Storage
Azure Table storage |Table storage is a service that stores non-relational structured data (also known as structured NoSQL data) in the cloud, providing a key/attribute store with a schemaless design. | Storage
Azure Cosmos DB | Globally distributed database that supports NoSQL options. | Databases
Azure SQL Database | Fully managed relational database with auto-scale, integral intelligence, and robust security.| Databases
Azure Database for MySQL|Fully managed and scalable MySQL relational database with high availability and security.| Databases
Azure Database for PostgreSQL|Fully managed and scalable PostgreSQL relational database with high availability and security.| Databases
SQL Server on Azure Virtual Machines |Service that hosts enterprise SQL Server apps in the cloud.| Databases
Azure Synapse Analytics | Fully managed data warehouse with integral security at every level of scale at no extra cost.| Databases
Azure Database Migration Service|Service that migrates databases to the cloud with no application code changes.| Databases
Azure Cache for Redis |Fully managed service caches frequently used and static data to reduce data and application latency.| Databases
Azure Database for MariaDB | Fully managed and scalable MariaDB relational database with high availability and security. | Databases
Azure App Service | Quickly create powerful cloud web-based apps. | Web
Azure Notification Hubs | Send push notifications to any platform from any back end.| Web
Azure API Management | Publish APIs to developers, partners, and employees securely and at scale.| Web
Azure Cognitive Search |  Deploy this fully managed search as a service.| Web
Web Apps feature of Azure App Service | Create and deploy mission-critical web apps at scale.| Web
Azure SignalR Service|Add real-time web functionalities easily.| Web
IoT Central | Fully managed global IoT software as a service (SaaS) solution that makes it easy to connect, monitor, and manage IoT assets at scale. | IoT
Azure IoT Hub| Messaging hub that provides secure communications between and monitoring of millions of IoT devices.| IoT
IoT Edge|Fully managed service that allows data analysis models to be pushed directly onto IoT devices, which allows them to react quickly to state changes without needing to consult cloud-based AI models.| IoT
Azure Synapse Analytics|Run analytics at a massive scale by using a cloud-based enterprise data warehouse that takes advantage of massively parallel processing to run complex queries quickly across petabytes of data.| Big Data
Azure HDInsight|Process massive amounts of data with managed clusters of Hadoop clusters in the cloud.| Big Data
Azure Databricks |Integrate this collaborative Apache Spark-based analytics service with other big data services in Azure.| Big Data
Azure Machine Learning Service |Cloud-based environment you can use to develop, train, test, deploy, manage, and track machine learning models. It can auto-generate a model and auto-tune it for you. It will let you start training on your local machine, and then scale out to the cloud. | AI
Azure ML Studio | Collaborative visual workspace where you can build, test, and deploy machine learning solutions by using prebuilt machine learning algorithms and data-handling modules.| AI
Vision|Use image-processing algorithms to smartly identify, caption, index, and moderate your pictures and videos.| AI
Speech |Convert spoken audio into text, use voice for verification, or add speaker recognition to your app.| AI
Knowledge mapping|Map complex information and data to solve tasks such as intelligent recommendations and semantic search.| AI
Bing Search |Add Bing Search APIs to your apps and harness the ability to comb billions of webpages, images, videos, and news with a single API call.| AI
Natural Language processing| Allow your apps to process natural language with prebuilt scripts, evaluate sentiment, and learn how to recognize what users want.| AI
Azure DevOps | Use development collaboration tools such as high-performance pipelines, free private Git repositories, configurable Kanban boards, and extensive automated and cloud-based load testing. Formerly known as Visual Studio Team Services. | DevOps
Azure DevTest Labs | Quickly create on-demand Windows and Linux environments to test or demo applications directly from deployment pipeline | DevOps
</details>

### Acronyms
<details>
<summary>
Acronyms worth remembering
</summary>

Acronym | Full Name | Notes | Domain 
---|---|---
CapEx | Capital Expenditure | up-front spending of money on physical infrastructure | Finance
OpEx | Operational Expenditure |  spending money on services or products now, and being billed for them now.| Finance
SLA | Service Level Agreement | what the company legally guarantees | ?
ARM | Azure Resource Manager | deployment and management service layer for Azure | Azure
</details>

###

</details>


