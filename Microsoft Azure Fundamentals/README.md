<!--
// cSpell:ignore PAAS
 -->

# Microsoft Azure Fundamentals

microsoft course 

[az-900 exam](https://learn.microsoft.com/en-us/certifications/exams/az-900)

each learning path corresponds to one domain of the exam.

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

## Core Azure concepts
<!-- <details> -->
<summary>
Azure fundamentals is a six-part series that teaches you basic cloud concepts, provides a streamlined overview of many Azure services, and guides you with hands-on exercises to deploy your very first services for free.
</summary>

https://learn.microsoft.com/en-us/training/paths/az-900-describe-cloud-concepts/

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

####  Different Types of Cloud Models
####  Cloud Benefits and Considerations
####  Different Cloud Services

</details>

### Core Azure Architectural Components
<details>
<summary>
In this module, you'll examine the various concepts, resources, and terminology that are necessary to work with Azure architecture. For example, you'll learn about Azure subscriptions and management groups, resource groups and Azure Resource Manager, as well as Azure regions and availability zones.
</summary>
</details>

</details>

## Takeaways
<!-- <details> -->
<summary>

</summary>

- Azure Web Apps
- Azure Functions
- Container Instance
- Kubernetes Services
- Cosmos DB - noSQL
- Azure Portal


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

</details>


