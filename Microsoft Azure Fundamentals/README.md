<!--
// cSpell:ignore yamlc Intune Postgre nsight
 -->

# Microsoft Azure Fundamentals

[Microsoft Azure Fundamentals:](https://learn.microsoft.com/en-us/training/paths/az-900-describe-cloud-concepts/), [az-900 exam](https://learn.microsoft.com/en-us/certifications/exams/az-900), [Azure Global infrastructure](https://infrastructuremap.microsoft.com/), [Azure Cli Documentation](https://learn.microsoft.com/en-us/cli/azure/?view=azure-cli-latest)

each learning path corresponds to one domain of the exam. a learning path is composed of several modules.

| AZ-900 Domain Area                                              | Weight |
| --------------------------------------------------------------- | ------ |
| Describe cloud concepts                                         | 20-25% |
| Describe core Azure services                                    | 15-20% |
| Describe core solutions and management tools on Azure           | 10-15% |
| Describe general security and network security features         | 10-15% |
| Describe identity, governance, privacy, and compliance features | 20-25% |
| Describe Microsoft cost management and Service Level Agreements | 10-15% |

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

[Core Azure Concepts](https://learn.microsoft.com/en-us/training/paths/az-900-describe-cloud-concepts/)

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
>
> - Disc
> - Blob
> - File
> - Archive
>   These services all share several common characteristics:\
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
>
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

#### Different Types of Cloud Models

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

#### Cloud Benefits and Considerations

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
>
> - No upfront costs.
> - No need to purchase and manage costly infrastructure that users might not use to its fullest.
> - The ability to pay for additional resources when they are needed.
> - The ability to stop paying for resources that are no longer needed.

#### Different Cloud Services

Cloud services can be divided into different service models, based on how the responsability of managing them is distributed between the provider and the user. on one end, there is IaaS (Infrastructure-as-a-Service), which gives the user the most control over the resources, while at the other end there is SaaS (Software-as-a-Service), which is the most managed services where the provider takes care of most of the work, between them there is PaaS (Platform-as-a-Service).

| Acronym | Full Name                   | Provider responsability                                  | User responsability                                | Example                |
| ------- | --------------------------- | -------------------------------------------------------- | -------------------------------------------------- | ---------------------- |
| Iaas    | Infrastructure-as-a-Service | hardware                                                 | perating system maintenance, network configuration | Azure virtual Machines |
| PaaS    | Platform-as-a-Service       | Virtual machines, networking                             | Applications                                       | Azure App Services     |
| SaaS    | Software-as-a-Service       | Virtual machines, networking, data storage, applications | Application data                                   | Microsoft Office 365   |

in IaaS, you simply rent the infrastructure instead of buying it, everything else is like on premises data center, the customer has the most control over the resources, and can tailor them to its' needs. in PaaS, the cloud provder has more responsability, which means the client can focus on the application development. in SaaS, the client only provides the data to the application, but the application (software) is provided by the cloud host. so this is the easiest option to use, and it provides access to the most up-to-date software, but it limits what kinds of applications the client can run.

| Model                       | Storage          | networking       | Compute          | virtual machine  | operating System | runtime          | Applications     | Data & Access    |
| --------------------------- | ---------------- | ---------------- | ---------------- | ---------------- | ---------------- | ---------------- | ---------------- | ---------------- |
| On-Premises (private cloud) | Customer Managed | Customer Managed | Customer Managed | Customer Managed | Customer Managed | Customer Managed | Customer Managed | Customer Managed |
| IaaS                        | Cloud Provided   | Cloud Provided   | Cloud Provided   | Customer Managed | Customer Managed | Customer Managed | Customer Managed | Customer Managed |
| PaaS                        | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Customer Managed | Customer Managed |
| SaaS                        | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Cloud Provided   | Customer Managed |

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

There are also _special Azure regions_, which are used for legal and compliance purposes, some are US-based and some China based.

Each region is comprised of availability zones, an availability zone is a data center. if one availability zone goes down, the others aren't effected. availability zones are connected to one another through a private, high speed, fiber-optic network. we can run applications in different availability zones for high availability or in the same availability zone for high performance.

services can be zonal, zonal-redundant, and non-regional.

- a zonal service is pinned to an specific availability zone. such as a virtual machine or a managed disc.
- a zonal-redundant service is replicated automatically across availability zones. such examples might be a database.
- a non-regional service is not tied to a specific region, and should be available even if the entire region goes down.

in addition to that, Azure has the concept of _Geography_, which contains _region pairs_. a geography is a conceptual location (such as asia, europe, us), and a region pair is a pairing of two regions, which are at a sufficient distance from one another but (mostly) reside in the same state. region pairs are preferred for cross region data redundancy, and azure keeps each region in the region pair on a different maintenance schedule, and will prioritize fixing one of the regions if a large scale disaster happens. this means that even if a disaster happens or an azure update goes horribly wrong, the data from one region can be preserved in the other region.

#### Resources and Azure Resource Manager

> - Resource: A manageable item that's available through Azure. Virtual machines (VMs), storage accounts, web apps, databases, and virtual networks are examples of resources.
> - Resource group: A container that holds related resources for an Azure solution. The resource group includes resources that you want to manage as a group. You decide which resources belong in a resource group based on what makes the most sense for your organization.

All resources resign in a resource group, and a resource can be a member of only one resource group. resources can be moved between resource groups (under some conditions), there is not nesting of resource groups. Resource groups allow for a logical grouping of resources based on the organization needs. Resource groups also control **lifecycle** - when a resource group is removed (deleted), all the resources inside it are also removed. in addition, we also use resource groups as an **authorization** scope, when we give role based access control (RBAC) permissions we can allow access to resources in a specific resource group.

> **Azure Resource Manager**:\
> Azure Resource Manager is the deployment and management service for Azure. It provides a management layer that enables you to create, update, and delete resources in your Azure account. You use management features like access control, locks, and tags to secure and organize your resources after deployment.

all azure actions (via the web console, apis, sdk, etc...) are processed via the Azure Resource Manager, so the results are always consistent no matter how the user interacts with azure.

> With Resource Manager, you can:
>
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

an azure Account has subscriptions, a subscription can have a _billing models_ and an _access-management/control_ policy (model). a billing model groups subscriptions into a different reports or invoices, which are required for managing costs. the access management defines the boundaries of which resources can be provisioned.

common ways to organize subscriptions is by geographical environment (which also helps with compliance), by the organizational structure (which helps with managing permissions) and by billing purposes (which helps track costs).

subscriptions also have limits, such as maximum number of Azure ExpressRoute circuits (10 per subscription), so a new subscription can allow for more of those limited resources.

**Billing profiles** are organized under a single **billing account**, a profile is matches an invoice and has payment method. inside each billing profile, we can designate different subscriptions a invoice sections.

**Azure management groups** provide a level of scope above that of subscriptions, and are used to manage access, policies and compliance in a single location. management groups can be nested (contain other management groups) and they contain the subscriptions. permissions are managed at the management group and are inherited by the nested groups and the subscriptions. all subscriptions in a management group must trust the same Azure AD tenant.

management group policies can limit what can be done by the subscriptions inside it, and helps manage all the permissions for different users.

> Important facts about management groups
>
> - 10,000 management groups can be supported in a single directory.
> - A management group tree can support up to six levels of depth. This limit doesn't include the root level or the subscription level.
> - Each management group and subscription can support only one parent.
> - Each management group can have many children.
> - All subscriptions and management groups are within a single hierarchy in each directory.

</details>

</details>

## Azure Services

<details>
<summary>
Basic Azure Architecture, Services and Security
</summary>

[Core Azure Services](https://learn.microsoft.com/en-us/training/paths/azure-fundamentals-describe-azure-architecture-services/)

### Architectural Components of Azure

<details>
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

we activate the sandbox environment in the azure account. it is limited to daily activations and each activation has a time limit.

Azure cloud shell

```ps
Get-Date
az version
bash #switch to bash
```

```sh
date
az upgrade
pwsh # switch to powershell
```

we can enter the interactive mode with `az interactive`, which allows us to run azure commands without the `az` command, and has autocompleteion.

```sh
version # az version
upgrade # az upgrade
exit #exit az interactive mode
```

#### Azure Physical Infrastructure

physical infrastucure and management infrastructure

the azure physical infrastructure is based on the azure data centers, which are are located all around the world

[Azure Global infrastructure](https://infrastructuremap.microsoft.com/). data centers are grouped into azure availability zones and regions, and then into geographies.

resources can be _zonal_ - pinned to a specific zone, _zone-redundant_ - automatically (or at request) duplicated across avalability zones, and _non-regional_ - which are always available, even in case of a region wide outage.

An _avalability zone_ is s datacenter (sometiemes more) which is physically separated from than other data centers, it has it's own power, cooling, and networks.\
Avalability zones allows for high availability and reddany in case of failure, and we can duplicate data across avalability zones (at a cost).

A _region_ contains at least (but usually more) availability zones (data centers), which are networked toegether. if one availability zone is down, the others continue to operate. azure also has the concept of _region pairs_, which is a pairing of two regions in the same geography. this allows for more redundancy, and azure keeps the two regions on different maintenance scheduling for extra safety.\
(some regions don't have a region pair or only have a one way backup).\
lastly, there are _sovereign regions_, such as us-gov and china. us-gov regions are operated by employees with us security clearnance, while china regions are operated by a chinese company, rather than microsoft.

#### Azure Management Infrastructure

the management infrastructure describes how azure resources are managed logically.

a _resource_ is anything that the user can deploy onto the cloud, it is the building block of azure - virtual machines, virtual networks, databases, cognitive services, everything.

_Resource groups_ are grouping of resources, each resource must be placed inside a resource group (some can be moved between them). we can apply actions to a resource group and it will applied to all the resources inside it - delete a resource group or give access. resource groups cannot be nested.

_subscriptions_ are azure unit of management, billing and scale. subscriptions contain resource groups, and are used to interact with resources and the azure portal. subscriptions act both as billing boundaries (separate invoices per subscription) and as access control boundaries (applying policies on a subscription to determine which operations are allowed).

Azure _management groups_ are the layer that manages subscriptions, this is done for enterprise level policies. management groups can be nested, and they contain subscriptions. we use management groups to apply governance policies and to provide access to multiple subscriptions via azure RBAC (role based access control).

#### Exercise - Create an Azure Resource

in this Exercise, we create a virtual machine.

in the portal, we click <kbd>Create a resource</kbd>, then choose <kbd>Virtual Machine</kbd> and <kbd>Create</kbd>.

we fill the settings with values.

```sh
az group list
az vm list --resource-group <group-name>
```

</details>

### Compute and Networking Services

<details>
<summary>
This module focuses on some of the computer services and networking services available within Azure.
</summary>

- compute options:
  - Virtual machines
  - containers
  - Azure functions
- networking
  - azure virtual networks
  - azure DNS
  - Azure express Route

#### Azure Virtual Machines

Azure virtual machine are a form of IaaS, we get a virtualized machine, which we can control and customize: the operating system, the software and the host configuration.

we can re-use the same virtual machine configuration by creating an Image out of it, and then use the image as a template.

the resources used for vms are:

- computing power: cpu cores, ram
- storage: hard disk drives, ssd
- netwokring: virtual networks, public ip address, ports

**Scaling: Scale Sets and Availability Sets**

_scale sets_ allow us to create and manage identical groups of virtual machines, all running the same software and the same configurations. the scale set vms all run on the same routing parameters, and can be monitored to scale up or down based on schedule or demand. they sit behind a load balancer.

_avaliability sets_ are a tool to deploy multiple vms in a way that ensures high availability. they are deployed in a way that prevents downtime. there is no additional costs for using avalability sets.

- update domain - separate machines so that they aren't updated at the same cycle, and there is a gap between updating each group.
- fault domain - separate machines by power source and network switch.

> Examples of when to use VMs:\
> Some common examples or use cases for virtual machines include:
>
> - During testing and development. VMs provide a quick and easy way to create different OS and application configurations. Test and development personnel can then easily delete the VMs when they no longer need them.
> - When running applications in the cloud. The ability to run certain applications in the public cloud as opposed to creating a traditional infrastructure to run them can provide substantial economic benefits. For example, an application might need to handle fluctuations in demand. Shutting down VMs when you don't need them or quickly starting them up to meet a sudden increase in demand means you pay only for the resources you use.
> - When extending your datacenter to the cloud: An organization can extend the capabilities of its own on-premises network by creating a virtual network in Azure and adding VMs to that virtual network. Applications like SharePoint can then run on an Azure VM instead of running locally. This arrangement makes it easier or less expensive to deploy than in an on-premises environment.
> - During disaster recovery: As with running certain types of applications in the cloud and extending an on-premises network to the cloud, you can get significant cost savings by using an IaaS-based approach to disaster recovery. If a primary datacenter fails, you can create VMs running on Azure to run your critical applications and then shut them down when the primary datacenter becomes operational again

"lift and shift" - moving from physical server to the cloud.

#### Exercise - Create an Azure Virtual Machine

> In this exercise, you create an Azure virtual machine (VM) and install Nginx, a popular web server.
>
> 1. Use the following Azure CLI commands to create a Linux VM and install Nginx. After your VM is created, you'll use the Custom Script Extension to install Nginx. The Custom Script Extension is an easy way to download and run scripts on your Azure VMs. It's just one of the many ways you can configure the system after your VM is up and running.
>
> ```sh
> az vm create \
>   --resource-group learn-85829cc9-09c8-47e6-9c14-519ca17cdc77 \
>   --name my-vm \
>   --image UbuntuLTS \
>   --admin-username azureuser \
>   --generate-ssh-keys
> ```
>
> 2. Run the following az vm extension set command to configure Nginx on your VM:
>
> ```sh
> az vm extension set \
>   --resource-group learn-85829cc9-09c8-47e6-9c14-519ca17cdc77 \
>   --vm-name my-vm \
>   --name customScript \
>   --publisher Microsoft.Azure.Extensions \
>   --version 2.1 \
>   --settings '{"fileUris":["https://raw.githubusercontent.com/MicrosoftDocs/mslearn-welcome-to-azure/master/configure-nginx.sh"]}' \
>   --protected-settings '{"commandToExecute": "./configure-nginx.sh"}'
> ```

we created a virtual machine ("my-vm") from an ubuntu Image, set an admin to the machine and created ssh keys. then we downloaded a script onto it and run it, this script installs nginx on the vm.

#### Azure Virtual Desktop

Azure virtual desktop is a virtual machine that runs an windows machine on the cloud, which can be used as any windows computer, not just as a server. we can separate the environment and the data from the hardware, the user can access the same desktop computer, no matter which device it is running.

this also allows for providing stronger machines without buying stronger hardware. and it also allows for better security, as all the important data is on the cloud, and not on the users machines.

#### Azure Containers

Azure containers are a way to run multiple instances of an application on a single host, rather than running multiple hosts. with containers, we don't manage the operating system. containers are designed to be a light-weight solution that responds better and faster to changes in demand.

vm - an abstraction layer for cpu, memory and storage, and operating system. but only one operating system per machine. containers bundle a single app and the dependencies, then it is run on a host machine in a container runtime, and many containers can run in a single host machine. because containers are smaller, they can scale up more easily, and can be orchestrated with orchestration services.

VM virtualize hardware, while containers virtualize the OS and runtime.

Azure container instances are a form of PaaS, we can split a website into different parts (frontend, backend, database) and run each in a different container, thus providing greater flexability.

#### Azure Functions

> Azure Functions is an event-driven, serverless compute option that doesn’t require maintaining virtual machines or containers. If you build an app using VMs or containers, those resources have to be “running” in order for your app to function. With Azure Functions, an event wakes the function, alleviating the need to keep resources provisioned when there are no events.

Azure functions are great for event driven operations, such as responding to an api. they can be scaled automatically. azure functions are stateless by default, but can also be stateful and maintain the context.

Serverless computing is the idea that we separate the server management (instaling os, patching, updating) from the developers.

- no infrastructure management
- scalability
- pay for what is used - not paying for resources which are not used.

#### Application Hosting Options

(hosting - making applications accessable from the web, like a website)

at the basic level, we can host applications on either virtual machines or on containers, but there is also the option of sing Azure App Service.

Azure App Service takes care of managing the infrastructure, provides automatic scaling and high avalability. it can integrate with github, azure devops or other code repository services for continues deployment model.\
we can run web apps (website), api apps (REST api) and webJobs, as well as mobile apps for ios and android.

#### Azure Virtual Networking

virtual networks, together with virtual subnets, allow azure resources to communicate with one another, with the outside web and with on-premises computers.

> Azure virtual networks provide the following key networking capabilities:
>
> - Isolation and segmentation
> - Internet communications
> - Communicate between Azure resources
> - Communicate with on-premises resources
> - Route network traffic
> - Filter network traffic
> - Connect virtual networks

public endpoints allow access to resources from anywhere in the world (public ip address), while private end point exists only within the virtual network and only have a private ip address, accessible only from inside the the address space of the containing virtual network.

Isolation and segmentation - private ip address which exist only inside the virtual network, and name resolution that can either be external or internal to the virtual network.

Internet communication - access to incoming traffic, either directly to the resource or via a load balancer.

Communication between azure resources - not only compute resources (virtual machines), but also with azure services such as databases, storage, and others.

Communication with on-premises - linking the azure cloud resources with local resources which reside in the data center, using VPN (point to site, site to site or with azure expressRoute as a private, high speed, dedicated connection.

Routing network traffic - control how traffic is routed in the virtual network, using route tables, gateways and other servies.

Filtering network traffic - using inbound and outbound security rules, and running a firewall.

Connecting virtual network - allowing separate virtual networks to connect to one another with network peering.

#### Exercise - Configure Network Access

opening our nginx server to connections from the outside web.

at first, we get the ip address, and see that we cannot access it with `curl` or from the browser.

> Task 1: Access your web server
>
> 1. Run the following az vm list-ip-addresses command to get your VM's IP address and store the result as a Bash variable
>
> ```sh
> IPADDRESS="$(az vm list-ip-addresses \
>   --resource-group learn-97546a8d-0942-4ed1-b361-766bbf499022 \
>   --name my-vm \
>   --query "[].virtualMachine.network.publicIpAddresses[*].ipAddress" \
>   --output tsv)"
> ```
>
> 2. Run the following curl command to download the home page:
>
> ```sh
> curl --connect-timeout 5 http://$IPADDRESS
> ```
>
> The --connect-timeout argument specifies to allow up to five seconds for the connection to occur. After five seconds, you see an error message that states that the connection timed out. This message means that the VM was not accessible within the timeout period. 3. as an optional step, try to access the web server from a browser:\
> Run the following to print your VM's IP address to the console:
>
> ```sh
> echo $IPADDRESS
> ```
>
> Copy the IP address that you see to the clipboard.\
> Open a new browser tab and go to your web server. After a few moments, you see that the connection isn't happening.\
> If you wait for the browser to time out, you'll see something like this:

next, we get the virtual network security group running on the resource group, we then list the rules on it, and we see that it has a single rule, allowing for ssh access via port 22.

> Task 2: List the current network security group rules
>
> 1. Run the following az network nsg list command to list the network security groups that are associated with your VM:
>
> ```sh
> az network nsg list \
> --resource-group learn-ed8fb9df-3281-4fe3-a6fb-40a617edb947 \
> --query '[].name' \
> --output tsv
> ```
>
> Every VM on Azure is associated with at least one network security group. In this case, Azure created an NSG for you called my-vmNSG. 2. Run the following az network nsg rule list command to list the rules associated with the NSG named my-vmNSG:
>
> ```sh
> az network nsg rule list \
> --resource-group learn-ed8fb9df-3281-4fe3-a6fb-40a617edb947 \
> --nsg-name my-vmNSG
> ```
>
> 3.  Run the az network nsg rule list command a second time. This time, use the --query argument to retrieve only the name, priority, affected ports, and access (Allow or Deny) for each rule. The --output argument formats the output as a table so that it's easy to read.
>
> ```sh
> az network nsg rule list \
> --resource-group learn-ed8fb9df-3281-4fe3-a6fb-40a617edb947 \
> --nsg-name my-vmNSG \
> --query '[].{Name:name, Priority:priority, Port:destinationPortRange, Access:access}' \
> --output table
> ```
>
> You see the default rule, default-allow-ssh. This rule allows inbound connections over port 22 (SSH). SSH (Secure Shell) is a protocol that's used on Linux to allow administrators to access the system remotely. The priority of this rule is 1000. Rules are processed in priority order, with lower numbers processed before higher numbers.\
> By default, a Linux VM's NSG allows network access only on port 22. This enables administrators to access the system. You need to also allow inbound connections on port 80, which allows access over HTTP.

now that we know that only ssh access is allowed, we create a rule that allows web traffic via port 80, once created, we list the rules again and see it exists.

> Task 3: Create the network security rule\
> Here, you create a network security rule that allows inbound access on port 80 (HTTP).
>
> 1. Run the following az network nsg rule create command to create a rule called allow-http that allows inbound access on port 80:\
>    For learning purposes, here you set the priority to 100. In this case, the priority doesn't matter. You would need to consider the priority if you had overlapping port ranges.
>
> ```sh
> az network nsg rule create \
> --resource-group > learn-ed8fb9df-3281-4fe3-a6fb-40a617edb947 \
> --nsg-name my-vmN> SG \
> --name allow-http \
> --protocol tcp \
> --priority 100 \
> --destination-port-range 80 \
> --access Allow
> ```
>
> 2. To verify the configuration, run az network nsg rule list to see the updated list of rules:\
>    You see this both the default-allow-ssh rule and your new rule, allow-http:
>
> ```sh
> az network nsg rule list \
> --resource-group learn-ed8fb9df-3281-4fe3-a6fb-40a617edb947 \
> --nsg-name my-vmNSG \
> --query '[].{Name:name, Priority:priority, Port:destinationPortRange, Access:access}' \
> --output table
> ```

now that the rule exists, we can try to access the resource again, and this time we can.

> Task 4: Access your web server again\
> Now that you've configured network access to port 80, let's try to access the web server a second time.
>
> 1. Run the same curl command that you ran earlier:
>
> ```sh
> curl --connect-timeout 5 http://$IPADDRESS
> ```
>
> 2. As an optional step, refresh your browser tab that points to your web server.
>
> Nice work. In practice, you can create a standalone network security group that includes the inbound and outbound network access rules you need. If you have multiple VMs that serve the same purpose, you can assign that NSG to each VM at the time you create it. This technique enables you to control network access to multiple VMs under a single, central set of rules.

#### Azure Virtual Private Networks

VPN - virtual private network

> A virtual private network (VPN) uses an encrypted tunnel within another network. VPNs are typically deployed to connect two or more trusted private networks to one another over an untrusted network (typically the public internet). Traffic is encrypted while traveling over the untrusted network to prevent eavesdropping or other attacks. VPNs can enable networks to safely and securely share sensitive information.

VPN Gateway is one type of a virtual network gateway, it is deployed in the a dedicated subnet of the virtual network and enables:

- connection from on premises datacenter through site-to-site connection
- connections from individual devices through point-to-site connections
- connecting other virtual network thorugh network-to-network connection

there can only be one vpn gateway for each virtual network, but a vpn gateway can connect to multiple locations.\
There are two ways to deploy a vpb gateway, _policy based_ or _route based_.

- policy based gateways encrypt data based on the ip address (which is statically specified)
- route based gateways use IPSec tunnels for routing, the routing can be static or dynamic, and it is the preferred approach and is more resilient to topology changes.
  - Connections between virtual networks
  - Point-to-site connections
  - Multisite connections
  - Coexistence with an Azure ExpressRoute gateway

we might also want to maximize the resiliency of the vpn gateway, so it isn't effected by outages. there are different ways to ensure high availability

- active/standby: two vpn gateways are created, but only one is active, when a failover or maintenance occurs, the standby is promoted and takes care of connections
- active/active: support of BGP routing protocol, two active vpn gateways, each with a unique public ip address. two separate tunnels to the vpn gateways.
- ExpressRoute Failover: combining expressRoute as the primary connection, and vpn gateway as the backup.
- Zone Redundant gatways - spreading the vpn gateways across multiple availability zones, and using gateway SKU and standard public ip address (and not basic public ip address)

#### Azure ExpressRoute

ExpressRoute connects the on-premises networks into the cloud with a expressRoute circuit connectivity provider. the connection can be through vpn, ethernet or any other connection. Express Routes has built in redundancy for high avilability

ExpresRoute can be used for global connectivity, if two datacenters are connected to the microsoft network, they can communicate over this network without going over to the public internet.\
Using Border Gateway Protocol (BGP) allows for dynamic routing.

> ExpressRoute supports four models that you can use to connect your on-premises network to the Microsoft cloud:
>
> - CloudExchange colocation
> - Point-to-point Ethernet connection
> - Any-to-any connection
> - Directly from ExpressRoute sites

#### Azure DNS

> Azure DNS is a hosting service for DNS domains that provides name resolution by using Microsoft Azure infrastructure. By hosting your domains in Azure, you can manage your DNS records using the same credentials, APIs, tools, and billing as your other Azure services.

we can use Azure to manage DNS records, this means we get the benefits of azure in terms of relability, performance (from anywhere in the world) and ease of use with consolidating management at the azure platform. we can use azure RBAC to control access who can do what with the DNS records, monitor actions with audit logs and prevent modifications with policies.

Azure also supports alias record sets and custom domain names.

</details>

### Storage Services

<details>
<summary>
This module introduces you to storage in Azure, including things such as different types of storage and how a distributed infrastructure can make your data more resilient.
</summary>

#### Azure Storage Accounts

- object store - blob, unstructured data
- disk store - for virtual machine and applications (ssd and hdd)
- file store - file shares in the cloud
- messaging store - queue, communicating between applications
- noSQL store - table storage, key-value pairs, semi-structured data

tiers:

- hot - frequently accessed data
- cold - infrequently accessed, at least 30 days
- archive - rarely accessed, at least 120 days

Azure storage accounts have unique namespaces, and are accessible from anywhere in the world over HTTP and HTTPS. storage accounts have different redundancy options.

| Type                        | Supported services                                                                        | Redundancy Options                   | Usage                                                                                                                                                                                                                                        |
| --------------------------- | ----------------------------------------------------------------------------------------- | ------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Standard general-purpose v2 | Blob Storage (including Data Lake Storage), Queue Storage, Table Storage, and Azure Files | LRS, GRS, RA-GRS, ZRS, GZRS, RA-GZRS | Standard storage account type for blobs, file shares, queues, and tables. Recommended for most scenarios using Azure Storage. If you want support for network file system (NFS) in Azure Files, use the premium file shares account type.    |
| Premium block blobs         | Blob Storage (including Data Lake Storage)                                                | LRS, ZRS                             | Premium storage account type for block blobs and append blobs. Recommended for scenarios with high transaction rates or that use smaller objects or require consistently low storage latency.                                                |
| Premium file shares         | Azure Files                                                                               | LRS, ZRS                             | Premium storage account type for file shares only. Recommended for enterprise or high-performance scale applications. Use this account type if you want a storage account that supports both Server Message Block (SMB) and NFS file shares. |
| Premium page blobs          | Page blobs only                                                                           | LRS                                  | Premium storage account type for page blobs only.                                                                                                                                                                                            |

because of the unique namespace of each storage account, it can also have a unique endpoint address. each storage type has it's own endpoint format.

| Storage service        | Endpoint format                                            |
| ---------------------- | ---------------------------------------------------------- |
| Blob Storage           | https://\<storage-account-name>.**blob**.core.windows.net  |
| Data Lake Storage Gen2 | https://\<storage-account-name>.**dfs**.core.windows.net   |
| Azure Files            | https://\<storage-account-name>.**file**.core.windows.net  |
| Queue Storage          | https://\<storage-account-name>.**queue**.core.windows.net |
| Table Storage          | https://\<storage-account-name>.**table**.core.windows.net |

#### Azure Storage Redundancy

Redundancy means storing multiple copies of the data, protecting it from events ch as hardware failure, network crushes and power outages.

- primary region data replication
  - Locally redundant storage (LRS)
  - Zone-redundant storage (ZRS)
- cross region data replication
  - Geo-redundant storage (GRS)
  - Geo-zone-redundant storage (GZRS)
- read-access to the replicated data in another region
  - Read-access geo-redundant storage (RA-GRS)
  - Read-access geo-zone-redundant storage (RA-GZRS)

Data is always replicated three times in the primary region, the data can be replicated in the same avilability zone, which protects against rack failure (LRS - local redundancy storage), or across different availability zones (ZRS - Zone-redundant storage), which protects against disasters to the data center.

if we wish to have better durability, we can also replicate the data to another region. the data is replicated to a single data center in the other zone (LRS). we can combine LRS and GRS with geo redundancy, controling how the data is replicated in the primary zone.\
When we have the data replicated in a secondary region, the default behavior is to only use it as backup (disaster recovery), but we can also configure it to allow read access from applications (even when the primary region is running properly). the data might not be immediately up-to-date.

| Type    | Primary Region | Secondary Region | Read Access? |
| ------- | -------------- | ---------------- | ------------ |
| GRS     | LRS            | LRS              | No           |
| GZRS    | ZRS            | LRS              | No           |
| RA-GRS  | LRS            | LRS              | Yes          |
| RA-GZRS | ZRS            | LRS              | Yes          |

#### Azure Storage Services

> - Azure Blobs: A massively scalable object store for text and binary data. Also includes support for big data analytics through Data Lake Storage Gen2.
> - Azure Files: Managed file shares for cloud or on-premises deployments.
> - Azure Queues: A messaging store for reliable messaging between application components.
> - Azure Disks: Block-level storage volumes for Azure VMs.

Azure storage benfits:

- Durable and highly available
- Secure
- Scalable
- Managed
- Accessible

##### Blob Storage

Storing unstructured data, such as text, photos, video or binary data. can store any format of files, and files of massive size. blob storage supports simultaneous uploads and downloads, and is often used for backup data or for data for analysis for other azure services. there are many ways to access data in blob storage, with direct URL access, REST api, cli tools (powershell, bash) and SDKs for multiple programing languages

> Blob storage is ideal for:
>
> - Serving images or documents directly to a browser.
> - Storing files for distributed access.
> - Streaming video and audio.
> - Storing data for backup and restore, disaster recovery, and archiving.
> - Storing data for analysis by an on-premises or Azure-hosted service.

blob storage is unlimited in size, so only a small portion of the data is frequently accessed, and sometimes blob storage is used for archiving and backup purposes, because of that, it's possible to save costs of storage by changing how it's stored. with hot storage tier, data is available immediately, with cooler storage, accessing the data takes longer, and per-access costs are higher, but the price for storage drops. Archive storage doesn't allow immediate access (the data is stored offline and must be "rehydrated"), but the costs are much lower.

Access Tiers:

- Hot - frequently accessed
- Cold (cool) - infrequently accessed, stored for at least 30 days
- Archive - rarely accessed, stored for at least 180 days

> - Only the hot and cool access tiers can be set at the account level. The archive access tier isn't available at the account level.
> - Hot, cool, and archive tiers can be set at the blob level, during or after upload.

##### Azure Files

A fully managed file share service, it can be accessed with standard protocols such as Server Message Block (SMB) and Network File System (NFS).

- shared access - both from azure service and the on-premises users
- fully managed - no need to manage the server hardware or OS
- scripting and tooling - support for powershell
- resiliency - highly available, durable
- familiar programmability - using the same system IO apis, so code doesn't break

##### Queue storage

A message storage service that stores messages, they can later be accessed from other services or using APIs. the maximal size of a message is 64k. we use this service to create backlog of work which can be processed asynchronously, they are commonly used together with azure functions.

##### Disk storage

Block storage to be used by azure virtual machines. this is the storage disk that is attached to them.

#### Exercise - Create a Storage Blob

Creating a new storage account

> 1. Sign in to the Azure portal at https://portal.azure.com
> 1. Select <kbd>Create a resource.</kbd>.
> 1. Under Categories, select <kbd>Storage</kbd>.
> 1. Under Storage Account, select <kbd>Create</kbd>.
> 1. On the Basics tab of the Create storage account blade, fill in the following information. Leave the defaults for everything else.
> 1. Select <kbd>Review + Create</kbd> to review your storage account settings and allow Azure to validate the configuration.
> 1. Once validated, select <kbd>Create</kbd>. Wait for the notification that the account was successfully created.
> 1. Select Go to resource.

working with blob storage and uploading an image.

> 1. Under Data storage, select <kbd>Containers</kbd>.
> 1. Select <kbd>+ Container</kbd> and complete the information.
> 1. Select <kbd>Create</kbd>.
> 1. Back in the Azure portal select the container you created, then select <kbd>Upload</kbd>.
> 1. Browse for the image file you want to upload. Select it and then <kbd>Upload</kbd>.
> 1. Select the Blob (file) you just uploaded. You should be on the properties tab.
> 1. Copy the URL from the URL field and paste it into a new tab. You should receive an error message similar to the following.

Changing the access level of the blob - making it viewable

> 1. Go back to the Azure portal
> 1. Select Change access level
> 1. Set the Public access level to Blob (anonymous read access for blobs only)
> 1. Select OK
> 1. Refresh the tab where you attempted to access the file earlier.

#### Azure Data Migration Options

Getting data from and into Azure

**Azure Migrate**\
Azure migrate is a service that enables migration from an on-premises environment. we can discover servers and machines (both physical and virtual) running with _azure migrate: discovery_,and migrate them with _azure migrate: server migration_. for databases, we can use the migration assistant tool to analyze our SQL server databases for potential issues which might prevent migration, and once solver, use _azure database migration service_ to migrate them to the cloud. there is also _azure app migration assistant_ to help with migrating on-premises websites.

**Azure Data Box**\
Azure data box is a physical device that can store a large amount of data, which can then be used to migrate into an Azure cloud. we get the device by ordering it from azure, and we connect it to our network or one of the computers in the network. once connected, we can use the Data box as storage and populate it with data, and it is shipped back to azure, and the data is transferred to the cloud.\
The device is physically hardened, shipped with tracking, and is wiped clean after each use to ensure the data is protected.

Databoxes are used when there are network problems, when the volume of data is too large to transfer over the internet, and when it's important to migrate large amounts. it can also be used for disaster recovery with data from the cloud being transferred to the on-premises store. it can also be used for migrating between cloud vendors.

#### Azure File Movement Options

Besides transfering entire services, it is also possible to work with individual files.

_AzCopy_ is a command line utility that copies blobs or files from or to the storage account. it can also do one-way synchronization.

_Azure Storage Explorer_ is a web interface app the provides a graphical way to navigate and manage the stored data.

_Azure File Sync_ is a tool the synchronizes a local windows server with the cloud storage, so it can act as a local cache for the stored data, and we can deploy it anywhere we want.

</details>

### Identity, Access, and Security

<details>
<summary>
This module covers some of the authorization and authentication methods available with Azure.
</summary>
</details>

#### Azure Directory Services

Azure AD is a directory service tha apples for both the cloud and the applications in it, and can be used for the on-premises machines.\
Azure AD is a cloud based identity and access management service, which runs on a windows server.

Azure AD can also provide auditing functionality, such as detecting unknown devices trying to connect from unknown locations.

Azure AD is for:

> - **IT administrators.** Administrators can use Azure AD to control access to applications and resources based on their business requirements.
> - **App developers**. Developers can use Azure AD to provide a standards-based approach for adding functionality to applications that they build, such as adding SSO functionality to an app or enabling an app to work with a user's existing credentials.
> - **Users**. Users can manage their identities and take maintenance actions like self-service password reset.
> - **Online service subscribers**. Microsoft 365, Microsoft Office 365, Azure, and Microsoft Dynamics CRM Online subscribers are already using Azure AD to authenticate into their account.

Azure AD provides:

> - **Authentication**: This includes verifying identity to access applications and resources. It also includes providing functionality such as self-service password reset, multifactor authentication, a custom list of banned passwords, and smart lockout services.
> - **Single sign-on**: Single sign-on (SSO) enables you to remember only one username and one password to access multiple applications. A single identity is tied to a user, which simplifies the security model.As users change roles or leave an organization, access modifications are tied to that identity, which greatly reduces the effort needed to change or disable accounts.
> - **Application management**: You can manage your cloud and on-premises apps by using Azure AD. Features like Application Proxy, SaaS apps, the My Apps portal, and single sign-on provide a better user experience.
> - **Device management**: Along with accounts for individual people, Azure AD supports the registration of devices. Registration enables devices to be managed through tools like Microsoft **Intune**. It also allows for device-based Conditional Access policies to restrict access attempts to only those coming from known devices, regardless of the requesting user account.

Azure AD can be used as an identity management for on-premises machines, this requires connecting the local Active directory with Azure AD, and it cab be done with **Azure AD Connect**.

**Azure Active Directory Domain Services** is a managed domain service, which allows for directrory services to run on the cloud without having an on-premises machine to run them. this includes legacy applications and current services. it also integrates with the Azure AD tenant.

AD DS creates a managed domain (a unique namespace) with windows server domain controller machines deployed to the selected region. data is synchronized to a degree.

#### Azure Authentication Methods

authentication is establishing the identity of the user (person/service/device), the user provides some sort of credentials as proof of identity. there are different forms of authentication, such as standard passwords, single sign-on (SSO), multifactor authentication and passwordles.\
Tradionally, there is a tradeoff between the security of the password and the convenience of using it.

##### Single Sign-on

single sign on allows a user to enter the credentials once, and use the same authorization across multiple applications. this requires that all the applications trust the initial authenticator.\
A SSO approach provides benefits to the user, as he only needs to remember a single password, it is also beneficial for the organization, as all access is mediated through the SSO, and there it is easier to disable and remove access if the need arises (such as when the user leaves the organization).

##### Multifactor Authentication

Multifactor authentication is the process of requiring an additional form of identification during the sign on. it provides an additional layer of security if the first layer (user password) was compromised.\
MFA usually takes form as an additional code sent to the mobile device of the user, or as face scan anf fingerprint data.

Azure AD MFA services provides the ability to authenticate with MFA, such as an application notification or a phone call.

##### Passwordless Authentication

Passwordless authentications is another form of identifying the user, it relies on having the device be connected to the identity, which lessens the threat of password leak. A second form of identification is still required, but it can be simplified (fingerprint, pin code) as it will only be valid for the specific device.

Azure AD has three forms of passwordless authentication:

- Windows Hello for bussiness - designated credentials tied to a specific PC, often using biometric and PIN (code) as the 2nd layer. Has built-in support for SSO and PKI (public key infrastructure) so it is a solid choice for large organizations.
- Microsoft Authenticator App - an application that can be configured to provide MFA for multiple accounts and services.
- FIDO2 security keys - a passwordless standard for authentication, often replacing password credentials with some kind of physical key (which is easier to use, harder to steal) such as a usb key, or a device with NFC/bluetooth.

#### Azure External Identities

an external Identity refers to a user (person/device/service) outside of the organization. common examples of external identities are those of vendors, suppliers, collabarators, partners, etc... they require some form of access to the organization resources, but it should be separate from the identity of those inside.

External Users can occasionally "bring the own identity" by using an identity provider (like google or facebook) or their own corporate identity. External users can become "guest users", or have a two-way connection with the organization, gaining access to some resources according to policy.

#### Azure Conditional Access

> Conditional Access is a tool that Azure Active Directory uses to allow (or deny) access to resources based on identity signals. These signals include who the user is, where the user is, and what device the user is requesting access from.

conditional access allows for granular control over identification and authentication, access attempts are classified according to signals (such as user, device, location, time), and if they are determined to ba unusual then a stronger form of authentication is required.

#### Azure Role-Based Access Control

following the principal of least privilege, RBAC is a way to define permissions which uses roles, which are then attached to user of user groups. with those roles, user can have access policies which define which resources are available to them.

azure RBAC uses scopes and roles. scope are the resources in question, and roles are the permitted actions.\
scopes:

- management group
- subscription
- resource group
- resource

roles:

- reader
- contributor/custom
- owner

Rbac is hierarchical, permissions at a parent scope (such as subscription) apply to all chile scopes (resource groups, resources). it is possbile to have multiple roles apply at once, with the permissions combining.

#### Zero Trust Model

Zero trust is a security model that assumes the worst case scenario for all access requests, and treats each request as if it came from an uncontrolled network. this is opposed to traditional models of security which differentiated between requests originating from inside the network and were considered safe by default

zero trust means:

> - Verify explicitly - Always authenticate and authorize based on all available data points.
> - Use least privilege access - Limit user access with Just-In-Time and Just-Enough-Access (JIT/JEA), risk-based adaptive policies, and data protection.
> - Assume breach - Minimize blast radius and segment access. Verify end-to-end encryption. Use analytics to get visibility, drive threat detection, and improve defenses.

#### Defense-In-Depth

> A defense-in-depth strategy uses a series of mechanisms to slow the advance of an attack that aims at acquiring unauthorized access to data.

each layer is protected, and no single layer allows complete access even if it's breached. if a layer is breached, then a security team is alerted to take action.

> - The **physical security** layer is the first line of defense to protect computing hardware in the datacenter.
> - The **identity and access** layer controls access to infrastructure and change control.
> - The **perimeter** layer uses distributed denial of service (DDoS) protection to filter large-scale attacks before they can cause a denial of service for users.
> - The **network** layer limits communication between resources through segmentation and access controls.
> - The **compute** layer secures access to virtual machines.
> - The **application** layer helps ensure that applications are secure and free of security vulnerabilities.
> - The **data** layer controls access to business and customer data that you need to protect.

##### Physical Security

This layer is about access to the physical hardware and the datacenters. controlled entirely by microsoft. preventing physical access means that other layers can't be bypassed.

##### Identity & Access

This layer is about access, who had it and to which resources. sign-in events should be logged and audited.

> - Control access to infrastructure and change control.
> - Use single sign-on (SSO) and multifactor authentication.
> - Audit events and changes.

##### Perimeter

Protecting against network based attacks on the resources, such as DDoS protection to avoid losing the availability of the system.

> - Use DDoS protection to filter large-scale attacks before they can affect the availability of a system for users.
> - Use perimeter firewalls to identify and alert on malicious attacks against your network.

##### Network

limiting network connectivity between and across resources, making sure that only the required connectivity is possible, thus minimizing the "blast radius" of a breach.

> - Limit communication between resources.
> - Deny by default.
> - Restrict inbound internet access and limit outbound access where appropriate.
> - Implement secure connectivity to on-premises networks.

##### Compute

Making sure the compute resources are secure - using secure operating systems, patched to the updated protected version, and cannot be accessed by outside.

> - Secure access to virtual machines.
> - Implement endpoint protection on devices and keep systems patched and current.

##### Application

Securing the application, either the one being developed or being used. keeping secrets and access keys outside the reach of application code.

> - Ensure that applications are secure and free of vulnerabilities.
> - Store sensitive application secrets in a secure storage medium.
> - Make security a design requirement for all application development.

##### Data

Protecting the data being used by the applications, storing personal and customer data in a secure way.

> - Stored in a database.
> - Stored on disk inside virtual machines.
> - Stored in software as a service (SaaS) applications, such as Office 365.
> - Managed through cloud storage.

#### Microsoft Defender for Cloud

> Defender for Cloud is a monitoring tool for security posture management and threat protection. It monitors your cloud, on-premises, hybrid, and multicloud environments to provide guidance and notifications aimed at strengthening your security posture.

A package of tools to harden resources, track security issues, protect against cyber attacks. this tool is integrated into azure. For on-premises datacenters, agents can be deployed to collect data. hybrid and environments can make use of Azure Arc. Protection cab be extended to resources running in other clouds as well.

Microsoft Defender for Cloud has capabilities to detect vulnerabilities and suggest security improvements of azure native resources:

- Azure PaaS services - anomaly detection, threat detection(Azure App Service, Azure SQL, Azure Storage Account)
- Azure data services - classify and protect data stored in databases(,Azure SQL, Azure Storage Account)
- Network - limit exposure to brute force, close VM ports, secured access policies based on user, ip addresses, ports. using Just-In-Time VM access.

Multicloud environments (using resources from other vendors, such as google GCP or amazon AWS) can be protected by Mircosoft Defender for cloud. this includes agentless protection, compliance assessments, container threat detection and protection for compute resources (windows and EC2 instances).

##### Assess, Secure, and Defend

> Defender for Cloud fills three vital needs as you manage the security of your resources and workloads in the cloud and on-premises:
>
> - Continuously assess – Know your security posture. Identify and track vulnerabilities.
> - Secure – Harden resources and services with Azure Security Benchmark.
> - Defend – Detect and resolve threats to resources, workloads, and services.

Continually assesing the state of the environment (virtual machines, container registries, SQL servers) for vulnerabilities, creating automatic reports with actionable suggestions.

Security in the cloud is multilayered, security policies are built on top of azure policies, can range from affecting a subscription, a management group or the entire tennat.

When resources are added or changed, Defender for Cloud can detect if they are configured according to best practices guidelines detailed in Azure Security Benchmark. and it provides a simple way to configure the resources to address problems. There is also a dashboard that displays how secured is the system.

When a security event is detected, Defender for Cloud can trigger a security alert, which summarizes the issue and details how it can be addressed. Multiple security events are grouped together to form a "cyber kill-chain", which details how an attack travels through recourses and how it affected the system.

</details>

## Solutions and Management Tools on Azure

<!-- <details> -->
<summary>

</summary>

[Azure Management and Reverences](https://learn.microsoft.com/en-us/training/paths/describe-azure-management-governance/)

### Cost Management

<details>
<summary>
This module explores methods to estimate, track, and manage costs in Azure.
</summary>

> factors that impact costs in Azure and tools to help you both predict potential costs and monitor and control costs.

#### Factors That Can Affect Costs In Azure

Using cloud infrastrcutere is about shifting cost from CapEx to OpEx. rather than spend money on investing in hardware, we can pay for (rent) infrastructure as we need it.

- Total Cost of Ownership (TCO) calculator to asses the cost of azure services as opposed to running them on premises
- Pricing calculator - which azure services fit the budget
- Azure advisor - monitoring costs and fining ways to save
- Spending Limits - prevent over-spending

There are many factors that impact the cost of running servies in Azure

##### Resource Type

obviously, each resource has a different cost, and that cost is impacted by the region it is running in. the cost is effected by resource specific parameters, such as sotrage account (blob storage) performance and access tiers, or virtual machine compute power.

##### Consumption

Azure has a "pay-as-you-go" billing structure, so the more a resource is used, the higher the cost is. some resources can be "reserved", and it that case, the customer gets a discount. reserving resources is done for a period of time (usually one or three years), and if the usage exceeds the reserved amount, normal pricing applies to the extra workloads.

##### Maintenance

While scaling up is a major advnatage of using cloud infrastructure, it is also a source of extra costs. every time a resource is provisoned, it might provision additional resources with it (VM requires networking and storage). It is up to the user to periodically review the resources and make sure they are deprovisioned when no longer needed.

##### Geography

Depending on where the resource is located, the cost of using it may vary. the location of the resource also effects traffic costs - both from outside the network and inside the cloud. costs are lower when residing in the azure region and when the physical distance is smaller.

##### Network Traffic

some traffic has costs, usually data going into Azure is free, and data leaving azure is charged. this is modified by billing zones.

##### Subscription type

Subscriptions can have different costs, allownances, access to services and free credit.

##### Azure Marketplace

Azure marketplace allows to purchase azure based solution from third-party vendors, such as pre-installed servers or backup services. In these cases, there is a cost attached to both the resources used and the vendor might have a billing structure as well.

#### Pricing and Total Cost of Ownership Calculators

Two Calculators that can give an estimate of costs.

The _Pricing Calculator_ is designed to give an estimate of how much it would cost to provision resources.

The _TCO Calculator_ (total cost of ownership) is deigned to help compare the costs of running workloads on premises against running them on the cloud. the user enter his current configuration of hardware and an estimate of labor/IT costs and is presented with a coresponding cost estimate of how much it would cost to run the workload in azure.

#### Exercise - Estimate Workload Costs by Using The Pricing Calculator

[Pricing Calculator](https://azure.microsoft.com/en-us/pricing/calculator/)

setting up two virtual machines, an application gateway and SQL database.

#### Exercise - Compare Workload Costs Using The TCO Calculator

[Tco Calculator](https://azure.microsoft.com/en-us/pricing/tco/calculator/)

cost of migrating to the cloud. virtual machines and a network
#### Cost Management Tool

Because Azure is a global cloud provider, and it supports rapid scaling, it is possible that resources will be provisioned and the user might not remember them until they appear in the invoice (and even then, the user might not recognize them as being unnecessary).

The Cost Management tool allows to quickly view the costs of provisioned resources, to filter them by region, billing cycle, resource type, etc...

with this information, azure can create forecasts of costs and then the user can see if there is an expected spike. in addition, the user can create special alerts to monitor spending.\
**Budget Alerts** can be used to notify when costs or usage exceeds a threshold or for some other conditions, alerts can be configured to send an email.\
**Credit Alerts** are used to track the consumption of azure credit monetary commitments. this is relvent for organization with Enterprise Agreements.\
**Department Spending Quota Alerts** are used to notify when a spending quota threshold is reached for a department.

Azure also has **Budgets**, which are spending limits based on a condition (subscription, resource group, service type,etc...) which can be configured and enforced.

#### The Purpose of Tags

Resources in Azure are orgainzed by the subscription (management group), the resource group by the location the resource belongs to. Tags allow for extra information that is fully customizable. resources of specific tags can be grouped together for billing and cost reports, they can be used to classify resources according to project, security level, compliance to government regulation and to track which department manages the resource.

Tags can be enforced by azure policies, both in terms of requiring tags and for applying other policies. for example, a policy might dictate that every resource must have an owner tag and an impact tag, and the level of this tag limits which resources can be issued.

</details>

### Governance and Compliance

<!-- <details> -->
<summary>
This module introduces you to tools that can help with governance and compliance within Azure.
</summary>

#### Azure Blueprints

> What happens when your cloud starts to grow beyond just one subscription or environment? How can you scale the configuration of features? How can you enforce settings and policies in new subscriptions?

Azure Blueprint standardizes the creation of subscriptions and environments. it allows for rapid deploy of identical instances which are already configured.

a component in a blue print is called _Artifact_, artifacts can have configurable parameters, which can be limited according to policies. for example, an artifact can be a VM, and the parameter being the compute type.

Blueprint have versions, so they can be updated with incremental improvements. Each resource created by a blueprint maintains the connection to the blueprint (and the version), which helpts with tracking and auditing deployments.

#### Azure Policy

#### Resource Locks

#### Exercise - Configure a resource lock

#### Service Trust portal

</details>

### Managing and Deploying Resource

<details>
<summary>
This module covers tools that help you manage your Azure and on-premises resources.
</summary>

</details>

### Montiroing Tools

<details>
<summary>
This module covers tools that you can use to monitor your Azure environment.

</summary>

</details>

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

<details>
<summary>
Stuff Worth remembering
</summary>

**services:**

- Azure Web Apps - scalable host websites
- Azure Functions - event driven actions.
- Container Instance
- Kubernetes Services
- Cosmos DB - noSQL
- Azure Portal
- Azure Resource Manager
- Azure AD - Active directory
- Border Gateway Protocol (BGP)
- Azure Route Server
- Azure ExpressRoute
- AzCopy - CLI tool to copy files and blobs to storage account
- Azure Storage Explorer - web app interface to storage account.
- Azure File Sync - synchronize windows server with storage account
- Microsoft Defender for Cloud - monitoring tool for security posture management and threat protection.
- Azure Arc - Microsoft defender extension for on-premises resources
- Azure Security Benchmark - security and compliance
  guidelines for azure.
  **misc:**

- resource groups cannot be nested.
- management groups can be nested.
- blob storage access
  - Only the hot and cool access tiers can be set at the account level. The archive access tier isn't available at the account level.
  - Hot, cool, and archive tiers can be set at the blob level, during or after upload.

### Azure Cli

<details>
<summary>
Azure command line and cloud shell commands
</summary>

[all cli commands reference](https://learn.microsoft.com/en-us/cli/azure/reference-index?view=azure-cli-latest).

all commands begin with `az`, unless inside interactive mode. we exit interactive mode with `exit`.

- Azure CLI - `az <command>` - commands which aren't specific to any service
  - `az version` - azure cli version
  - `az upgrade` - upgrade azure cli version
  - `az interactive` - enter interactive mode
    - `exit` - exit interactive mode
- Azure Resource Groups - `az group`
  - `az group list` - list resource group
- Azure Virtual Machines - `az vm`
  - `az vm list` - list virtual machines in the default resource group
    - `-g, --resource-group <group-name>` - list in a specifc resource group
  - `az vm create` - create virtual machine
    - `--resource-group`
    - `--name` - the name of the virtual machine
    - `--image` - operating system URN/custom image name/Id
    - `--admin-username` - Username for the VM. Default value is current username of OS. If the default value is system reserved, then default value will be set to azureuser.
    - `--generate-ssh-keys` - Generate SSH public and private key files if missing.
  - `az vm extension set` - add post deployment application to to vm
    - `--resource-group`
    - `--vm-name <vm-name>` - virtual machine name
    - `--name` - extension name
    - `--publisher`- extension publisher
    - `--version` - extension version
    - `--settings` - pass data in json format or path to json
    - `--protected-seetings` - pass sensitive information, json format or path to json
  - `az vm list-ip-addresses` - list ip addresses for virtual machine
    - `--resource-group`, `--query`, `--output`
    - `--vm-name <vm-name>` - virtual machine name
- Azure Network - `az network`
  - `az network nsg` - network security groups
    - `az network nsg list` - list network security group
      - `--resource-group`, `--query`, `--output`
    - `az network nsg rule list` - list rules for a specific network security group
      - `--resource-group`
      - `--nsg-name <nsg-name>` - network security group name
    - `az network nsg rule create` - create network security group
      - `--resource-group`
      - `--nsg-name <nsg-name>` - network security group name
      - `--name <rule-name>` name of the nsg rule
      - `--protocol <*|tcp|udp|icmp|esp|ah>` - protocol to apply this rule to
      - `--priority <numeric priority>` - number between 100 (highest priority) and 4096 (lowest), unique for each rule
      - `--destination-port-range <list of port number or range>` - which port this rule applies to (default 80)
      - `--access <Allow|Deny>` - rule type
- General flags
  - `-g, --resource-group <group-name>` - resource group name - `--query "<query>"` - query to show specific data (drill down). [query format](https://learn.microsoft.com/en-us/cli/azure/query-azure-cl) - `--output <json|jsonc|yaml|yamlc|tsv|table|none>` - [output format](https://learn.microsoft.com/en-us/cli/azure/format-output-azure-cli)
  </details>

### Azure Services

<details>
<summary>
Table of Services in Azure
</summary>

| Service name                                           | Service function                                                                                                                                                                                                                                              | Section    |
| ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| Azure Virtual Machines                                 | Windows or Linux virtual machines (VMs) hosted in Azure.                                                                                                                                                                                                      | Compute    |
| Azure Virtual Machine Scale Sets                       | Scaling for Windows or Linux VMs hosted in Azure.                                                                                                                                                                                                             | Compute    |
| Azure Kubernetes Service                               | Cluster management for VMs that run containerized services.                                                                                                                                                                                                   | Compute    |
| Azure Service Fabric                                   | Distributed systems platform that runs in Azure or on-premises.                                                                                                                                                                                               | Compute    |
| Azure Batch                                            | Managed service for parallel and high-performance computing applications.                                                                                                                                                                                     | Compute    |
| Azure Container Instances                              | Containerized apps run on Azure without provisioning servers or VMs.                                                                                                                                                                                          | Compute    |
| Azure Functions                                        | An event-driven, serverless compute service.                                                                                                                                                                                                                  | Compute    |
| Azure Virtual Network                                  | Connects VMs to incoming virtual private network (VPN) connections.                                                                                                                                                                                           | Networking |
| Azure Load Balancer                                    | Balances inbound and outbound connections to applications or service endpoints.                                                                                                                                                                               | Networking |
| Azure Application Gateway                              | Optimizes app server farm delivery while increasing application security.                                                                                                                                                                                     | Networking |
| Azure VPN Gateway                                      | Accesses Azure Virtual Networks through high-performance VPN gateways.                                                                                                                                                                                        | Networking |
| Azure DNS                                              | Provides ultra-fast DNS responses and ultra-high domain availability.                                                                                                                                                                                         | Networking |
| Azure Content Delivery Network                         |
| Delivers high-bandwidth content to customers globally. | Networking                                                                                                                                                                                                                                                    |
| Azure DDoS Protection                                  | Protects Azure-hosted applications from distributed denial of service (DDOS) attacks.                                                                                                                                                                         | Networking |
| Azure Traffic Manager                                  | Distributes network traffic across Azure regions worldwide.                                                                                                                                                                                                   | Networking |
| Azure ExpressRoute                                     | Connects to Azure over high-bandwidth dedicated secure connections.                                                                                                                                                                                           | Networking |
| Azure Network Watcher                                  | Monitors and diagnoses network issues by using scenario-based analysis.                                                                                                                                                                                       | Networking |
| Azure Firewall                                         | Implements high-security, high-availability firewall with unlimited scalability.                                                                                                                                                                              | Networking |
| Azure Virtual WAN                                      | Creates a unified wide area network (WAN) that connects local and remote sites.                                                                                                                                                                               | Networking |
| Azure Blob storage                                     | Storage service for very large objects, such as video files or bitmaps.                                                                                                                                                                                       | Storage    |
| Azure File storage                                     | File shares that can be accessed and managed like a file server.                                                                                                                                                                                              | Storage    |
| Azure Queue storage                                    | A data store for queuing and reliably delivering messages between applications.                                                                                                                                                                               | Storage    |
| Azure Table storage                                    | Table storage is a service that stores non-relational structured data (also known as structured NoSQL data) in the cloud, providing a key/attribute store with a schemaless design.                                                                           | Storage    |
| Azure Cosmos DB                                        | Globally distributed database that supports NoSQL options.                                                                                                                                                                                                    | Databases  |
| Azure SQL Database                                     | Fully managed relational database with auto-scale, integral intelligence, and robust security.                                                                                                                                                                | Databases  |
| Azure Database for MySQL                               | Fully managed and scalable MySQL relational database with high availability and security.                                                                                                                                                                     | Databases  |
| Azure Database for PostgreSQL                          | Fully managed and scalable PostgreSQL relational database with high availability and security.                                                                                                                                                                | Databases  |
| SQL Server on Azure Virtual Machines                   | Service that hosts enterprise SQL Server apps in the cloud.                                                                                                                                                                                                   | Databases  |
| Azure Synapse Analytics                                | Fully managed data warehouse with integral security at every level of scale at no extra cost.                                                                                                                                                                 | Databases  |
| Azure Database Migration Service                       | Service that migrates databases to the cloud with no application code changes.                                                                                                                                                                                | Databases  |
| Azure Cache for Redis                                  | Fully managed service caches frequently used and static data to reduce data and application latency.                                                                                                                                                          | Databases  |
| Azure Database for MariaDB                             | Fully managed and scalable MariaDB relational database with high availability and security.                                                                                                                                                                   | Databases  |
| Azure App Service                                      | Quickly create powerful cloud web-based apps.                                                                                                                                                                                                                 | Web        |
| Azure Notification Hubs                                | Send push notifications to any platform from any back end.                                                                                                                                                                                                    | Web        |
| Azure API Management                                   | Publish APIs to developers, partners, and employees securely and at scale.                                                                                                                                                                                    | Web        |
| Azure Cognitive Search                                 | Deploy this fully managed search as a service.                                                                                                                                                                                                                | Web        |
| Web Apps feature of Azure App Service                  | Create and deploy mission-critical web apps at scale.                                                                                                                                                                                                         | Web        |
| Azure SignalR Service                                  | Add real-time web functionalities easily.                                                                                                                                                                                                                     | Web        |
| IoT Central                                            | Fully managed global IoT software as a service (SaaS) solution that makes it easy to connect, monitor, and manage IoT assets at scale.                                                                                                                        | IoT        |
| Azure IoT Hub                                          | Messaging hub that provides secure communications between and monitoring of millions of IoT devices.                                                                                                                                                          | IoT        |
| IoT Edge                                               | Fully managed service that allows data analysis models to be pushed directly onto IoT devices, which allows them to react quickly to state changes without needing to consult cloud-based AI models.                                                          | IoT        |
| Azure Synapse Analytics                                | Run analytics at a massive scale by using a cloud-based enterprise data warehouse that takes advantage of massively parallel processing to run complex queries quickly across petabytes of data.                                                              | Big Data   |
| Azure HDInsight                                        | Process massive amounts of data with managed clusters of Hadoop clusters in the cloud.                                                                                                                                                                        | Big Data   |
| Azure Databricks                                       | Integrate this collaborative Apache Spark-based analytics service with other big data services in Azure.                                                                                                                                                      | Big Data   |
| Azure Machine Learning Service                         | Cloud-based environment you can use to develop, train, test, deploy, manage, and track machine learning models. It can auto-generate a model and auto-tune it for you. It will let you start training on your local machine, and then scale out to the cloud. | AI         |
| Azure ML Studio                                        | Collaborative visual workspace where you can build, test, and deploy machine learning solutions by using prebuilt machine learning algorithms and data-handling modules.                                                                                      | AI         |
| Vision                                                 | Use image-processing algorithms to smartly identify, caption, index, and moderate your pictures and videos.                                                                                                                                                   | AI         |
| Speech                                                 | Convert spoken audio into text, use voice for verification, or add speaker recognition to your app.                                                                                                                                                           | AI         |
| Knowledge mapping                                      | Map complex information and data to solve tasks such as intelligent recommendations and semantic search.                                                                                                                                                      | AI         |
| Bing Search                                            | Add Bing Search APIs to your apps and harness the ability to comb billions of webpages, images, videos, and news with a single API call.                                                                                                                      | AI         |
| Natural Language processing                            | Allow your apps to process natural language with prebuilt scripts, evaluate sentiment, and learn how to recognize what users want.                                                                                                                            | AI         |
| Azure DevOps                                           | Use development collaboration tools such as high-performance pipelines, free private Git repositories, configurable Kanban boards, and extensive automated and cloud-based load testing. Formerly known as Visual Studio Team Services.                       | DevOps     |
| Azure DevTest Labs                                     | Quickly create on-demand Windows and Linux environments to test or demo applications directly from deployment pipeline                                                                                                                                        | DevOps     |

</details>

### Acronyms

<details>
<summary>
Acronyms worth remembering
</summary>

| Acronym | Full Name                              | Notes                                                                      | Domain             |
| ------- | -------------------------------------- | -------------------------------------------------------------------------- | ------------------ |
| CapEx   | Capital Expenditure                    | up-front spending of money on physical infrastructure                      | Finance            |
| OpEx    | Operational Expenditure                | spending money on services or products now, and being billed for them now. | Finance            |
| SLA     | Service Level Agreement                | what the company legally guarantees                                        | ?                  |
| ARM     | Azure Resource Manager                 | deployment and management service layer for Azure                          | Azure              |
| NSG     | Network Security Group                 |                                                                            | Virtual Network    |
| VPN     | Virtual Private Network                |                                                                            | Networking         |
| BGP     | Border Gateway Protocol                |                                                                            | Networking         |
| LRS     | Locally redundant storage              | Replication in the same AZ                                                 | Storage            |
| ZRS     | Zone-redundant storage                 | Replication in the different AZ in the same region                         | Storage Redundancy |
| GRS     | Geo-redundant storage                  | LRS + another copy in another region                                       | Storage Redundancy |
| GZRS    | Geo-zone-redundant storage             | ZRS + another copy in another region                                       | Storage Redundancy |
| RA-GRS  | Read-access geo-redundant storage      | GRS + ability for application to read from the backup region               | Storage Redundancy |
| RA-GZRS | Read-access geo-zone-redundant storage | GZRS + ability for application to read from the backup region              | Storage Redundancy |
| SMB     | Server Message Block                   | linux, macOs, windows                                                      | File Storage       |
| NFS     | Network File syStem                    | linux and MacOs                                                            | File Storage       |
| AD DS   | Active Directory Domain Services       |                                                                            | Identity           |
| LDAP    | lightweight directory access protocol  |                                                                            | Identity           |
| FIDO    | Fast Identity Online                   | passwordless authentication standard                                       | Identity           |
| PKI     | Public Key Infrastructure              |                                                                            | Identity           |
| CSPM    | Cloud Security Posture Management      |                                                                            | Security           |
| JIT     | Just-In-Time                           | least privilege access                                                     | Identity/Security  |
| JEA     | Just-Enough-Access                     | least privilege access                                                     | Identity/Security  |
| TCO     | Total Cost of Ownership                | TCO calculator                                                             | Pricing            |

</details>

### Azure and AWS

<details>
<summary>
which services in azure are comparable to aws services
</summary>

| Service Purpose         | Azure Name         | AWS Name                    | Notes                                     |
| ----------------------- | ------------------ | --------------------------- | ----------------------------------------- |
| Virtual Machine         | Azure VM           | EC2                         | compute                                   |
| General Storage         | Blob storage       | S3                          | object / blob storage                     |
| General Storage Tiers   | Hot, Cool, Archive | S3 Standard, IA, Glacier    | storage tiers for costs and retrival time |
| Event-Driven Serverless | Azure Functions    | Lambda                      |
| Message Queue           | Queue storage      | SQS                         | asynchronous message handlng              |
| VM disk volume          | Azure Disks        | EBS - elastic block storage | storage volumes for virtual compute       |
| File Storage            | Azure Files        | EFS - elastic file system   |
| Physical data transfer  | Azure Data Box     | AWS SnowBall                | Data migration                            |
| Resource Creation       | Azure Blueprints   | AWS CloudFormation          | create resources from templates           |

</details>

###

</details>
