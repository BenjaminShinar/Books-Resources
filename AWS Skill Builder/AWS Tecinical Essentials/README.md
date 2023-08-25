<!--
ignore these words in spell check for this file
// cSpell:ignore Seph
 -->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## AWS Technical Essentials

[AWS Technical Essentials](https://explore.skillbuilder.aws/learn/course/1851/aws-technical-essentials). instructors: _Morgan Willis_, _Seph R._

> AWS Technical Essentials introduces you to essential Amazon Web Services and common solutions. The course covers the fundamental AWS concepts related to compute, database, storage, networking, monitoring, and security. You will start working in AWS through hands-on course experiences.\
> The course covers the concepts necessary to increase your understanding of AWS services, so that you can make informed decisions about solutions that meet business requirements. Throughout the course, you will gain information on how to build, compare, and apply highly available, fault tolerant, scalable, and cost-effective cloud solutions.

### MODULE 1: Introduction to Amazon Web Services

<details>
<summary>
Cloud and AWS introduction.
</summary>

#### What Is AWS?

<details>
<summary>
Basic Cloud Deployment models and advantages.
</summary>

types of deployment models:

- On-premises - having physical hardware dedicated which belongs to the company, either in the company physical location or in a data center nearby. this requires paying to obtain that hardware and managing it.
- Cloud - getting the IT resources via the internet from a cloud provider (such as AWS), this means there is no need for the company to maintain its' own data centers and purchase hardware.
- Hybrid - a combination of both, connecting hardware that resides on-premises with resources on the cloud, can be part of a transition period or be caused by regulatory needs.

Running workloads On-Premises requires owning the hardware, configuring it and deploying the software, which can be time consuming and expensive. cloud resources can be provisioned on demand with pay-per-use pricing so there are less barriers to experiment, and the process can be completed more quickly. Using cloud compute resources also allows to focus on the unique parts of the project. the compute resources, the databases, backups and networking portions usually remain the same across projects and across businesses, so it makes sense to get them from an external source rather than spend special effort on setting them up.

Cloud computing offers six advantages:

- Pay-As-You-Go Model - getting resources from the cloud provider is priced based on usage, there is no need to buy expensive hardware that might not be fully utilized.
- Benefit from Massive Economics of Scale - Cloud Vendors are responsible to getting the hardware, and because they buying units in massive scales, they can get better prices than an individual company can. this allows them to offer lower prices in the "pay as you go" models.
- Stop Guessing Capacity - rather than overbuying hardware for an on-premises data center in order to be ready for demand spikes, cloud resources can be acquired immediately when extra capacity is needed, and can be released when the demand goes down.
- Increase Speed and Agility - Since resources are accessed over the internet, there is little waiting time between provisioning the resources and using them. there is no need to wait for shipping and installation like an on-premises data center.
- Realize Cost Savings - since the cloud vendor is responsible to hosting and maintaining the hardware, there are less costs of racking, powering, cooling (and other physical operations), and the company can focus on the core product.
- Go Global in Minutes - Cloud Vendors have data centers across the world, so if there is a need to deploy in a different region, it is easy to deploy the same stack at a different data center and get lower networking latency.

</details>

#### AWS Global Infrastructure

<details>
<summary>
AWS has dataCenters across the world, and they can be accessed from everywhere.
</summary>

(video)

redundancy for disaster recovery across Availability Zones and across Regions

choosing the aws region:

- compliance - the most important factor, are there regulatory controls dictating which region can be used?
- latency - the closer the region to the users, the better. closer proximity means higher speed.
- pricing - pricing can vary between regions (sometimes due to taxation structures).
- service availability - some services might not be available in all regions, especially when the service is new.

the global edge network (edge locations and region edge caches) - caching data closer to the end-user. **Amazon Cloud Front** can be used to cache data.

_Regions_ are geographical regions, named after where they are and given an aws code. regions are independent from one another, and there is no automatic data transfer between them without the user setting it up. Regions consist of two or more _Availability Zones_, which are data centers spread across the region. they are remote from one another to avoid having them go down together.

Services can belong to different scopes: _Availability Zone_, _Region_ or the _Global_ level. global scoped services (such as **IAM**) are the same for all regions, other service might require you to choose a region and an Availability Zone, this might be done to set data durability and availability. For other services, AWS itself manages the placement in Availability Zones, and you only set the region.

A service should be highly available and resilient, which usually means that it should be deployed in more than one Availability Zone, so it can easily recover and provide service if one of aws data centers goes down.

</details>

#### Interacting with AWS

<details>
<summary>
Different Ways to work with AWS services.
</summary>

we can interact with the AWS resources in different way, via the management console (website), with the cli tool (scripting) or with software development kits that integrate with other software.

- AWS Management Console - website
- AWS CLI - command line interface
- AWS SDk - software development kit - for each programming language

the aws **Cloud Shell** is a service that provides CLI access from the web management console. the cli has the format of `aws <service> <commands>`.

</details>

#### Security and the AWS Shared Responsibility Model

<details>
<summary>
Security in the cloud.
</summary>

(video)

Security Concerns are divided between AWS the User. AWS is always responsible to the physical security, but depending on the service, AWS might be in charge of more or less parts of the security. for an **EC2** machine, aws is in charge of the underlying hardware and the virtualization, but the user is in charge of the installed OS, patches and controlling access to the resource. AWS offers many options to get security features on services, and usually, the user is responsible to the data security and controlling access to it.

The more "managed" services are, the more aws is responsible for them, S3 is a very managed solution with less user controlled security, while EC2 is a non-managed solution where the user has to manage much more.

> A key concept is that customers maintain complete control of their data and are responsible for managing the security related to their content. For example, you are responsible for the following:
>
> - Choosing a Region for AWS resources in accordance with data sovereignty regulations
> - Implementing data-protection mechanisms, such as encryption and scheduled backups
> - Using access control to limit who can access your data and AWS resources

</details>

#### Protecting the AWS Root User

<details>
<summary>
Protecting The Root User with MultiFactor authentication
</summary>

(video)

the root user has unlimited permissions and powers. the root user is accessed with an email and password. since the root user has all the permissions, we need to protect it, as we don't want it to be compromised. to prevent this, we should enable Multi-Factor-Authentication to as a second layer. MFA device can be physical or virtual.\
An additional recommendation is to avoid using the root user for everyday tasks, and instead use IAM roles with limited powers.

Besides the email and password combination, there is another set of credentials, the key id and secret for programmatic access. they are used to connect to the IAM role via the CLI or other APIs.

> The root user has complete access to all AWS services and resources in your account, including your billing and personal information. Therefore, you should securely lock away the credentials associated with the root user and not use the root user for everyday tasks. Visit the links at the end of this lesson to learn more about when to use the AWS root user.\
> To ensure the safety of the root user, follow these best practices:
>
> - Choose a strong password for the root user.
> - Enable multi-factor authentication (MFA) for the root user.
> - Never share your root user password or access keys with anyone.
> - Disable or delete the access keys associated with the root user.
> - Create an Identity and Access Management (IAM) user for administrative tasks or everyday tasks.

Multi Factor authentication relies on having a combination of identification methods -

1. Something You Know - such as user name and password
2. Something You Have - such as a one-time pass-code from an authentication device
3. Something You Are - like a fingerprint or face scanning identification.

AWS mostly works with "Something You Have", this can be a security stick that's been certified by FIDO, or a pass-code generated by either a physical or virtual MFA device (like an application on the phone).

</details>

#### AWS Identity and Access Management

<details>
<summary>
IAM: Users, Groups, Policies, Roles and Identity providers
</summary>

(video)

access control, api permissions are required even when the resources are in the same account or VPN. different people need different access and permissions.

Aws IAM (Identity and Access Management) controls the log-in credentials and permissions, and the signing of API calls to resources. it Doesn't control application identity.

- Authentication - user is who they say they are. "who is this user?"
- Authorization - user can do the actions want to do. "what can this user do?"

we use IAM policies to grant or deny permissions to perform actions, and we attach those policies to AWS identities. Policies are json-based documents:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "ec2:*",
      "Resource": "",
      "Condition": {}
    }
  ]
}
```

we can also attach IAM policies to a group, a group is collection of users, which makes it easier to assign policies without managing individual users. this is the best practice to follow.

another best practice is to create an administrator user from the root user, and use that user for all administrative actions, leaving the root user unused. we can apply IAM policies to that new user, but not the root user. unlike human users, applications should use roles to access resource in AWS.

IAM is the centralized view service of the uses, roles and permissions in the aws account.

- Global - not specific to any one region - all IAM configurations are viewed in a unified way.
- Integrated with AWS services - by default.
- Shared Access - providing permissions to other users to act in the account without sharing the password or key.
- Multi Factor Authentication - supports for extra security.
- Identity Federation - Allows for integration with other identity providers (such as corporate users) to gain temporary access to the AWS account.
- Free to use - No additional charge for IAM services.

##### IAM User

the IAM user represents a person or a service that interacts with the resources in the accounts, any action the user performs is billed to the account. a User can authenticate to the management console (website) or with programmatic access with key and secret. Credentials are permanent until removed or rotated.\
Permissions can be granted directly to an IAM user, but as the number of users increases, this becomes harder to manage.

##### IAM Group

An IAM group is a collection of users, all the users in the group "inherit" permissions from the group, so it's easier to manage on a large scale. onboarding a new employee simply requires creating a new user in the "developers" group, rather than assigning specific permissions to each new user. and moving a user to a different role is as simple as removing it from one group and assigning them to another.\
Users can belong to several groups at once, but groups can not be nested.

##### IAM Policies

Permissions are managed through policies, which are attached to IAM identities (users, groups, roles), any action is checked against the policy document and is executed only if the permissions match. the policy document has:

- version - defines this as an IAM policy, mush be "2012-10-17". should always be placed in the documents.
- Statement - the collection of permissions
- Sid - textual description of the statement
- Effect - does the statement "allow" this behavior or "deny" it?
- Action - which actions the effect refers to, use asterisks `*` as wild card for all actions
- Resource - which resource the statement refers to, defined by the ARN or with wild card.

##### Roles

(video)

role-based-access

policies can be applied to user and groups, and also to roles. an IAM role is a temporary identity that can assumed in order to gain access to AWS credentials (and resources). Each api request to AWS is signed, so if we don't want to create a user for each component of the application, we should use IAM roles for them. roles also have policies attached , but they don't use static credentials, they have temporary log-in credentials the are dynamically created.

to create a role, we go to the <kbd>IAM</kbd> service, choose <kbd>Roles</kbd>, <kbd>Create Role</kbd> and choose the trusted entity of the EC2 machine. here we can attach permissions (policies) for the role, in our case we search for S3 and select the full access policy. we filter for "dynamoDb" and do the same. we give the role a name, description and tags. roles are used to communicate between aws services.

we can also use external Identity Providers to assume roles, this can be done when we have an existing identity providers, this way we create _Federated users_, which is done via the IAM identity Center.

##### IAM Best Practices

- Lock down the AWS root user - the most powerful role in the AWS account, should be protected. can not be limited.
  - Don't share credentials
  - Activate MFA on the root account
  - Consider deleting the root user access keys
- Follow the principle of least privilege - security standard to only grant the permissions needed to perform the desired task, and nothing more.
- Use IAM appropriately - IAM is only to provide access to AWS resources, not for applications level authentications.
- Use IAM Roles when possible - Roles are more efficient than users. they are easier to manage and use dynamic credentials that are temporary, rather than static credentials such as passwords and access keys.
- Consider using an identity provider - When there are many users who need access, it might be easier to have a single access point to all AWS accounts via an identity provider. this makes managing users easier, and provides a consolidated, single-source-of-truth access to AWS.
- Regularly review and remove unused users, roles and other credentials - remove unused identities to make monitoring easier.

</details>

#### Demonstration: Implementing Security with IAM

<details>
<summary>
Working with AWS for our demo application
</summary>

(video)

creating an IAM Role, and users, under the <kbd>IAM</kbd> service, we create a role for the EC2 service. other options are:

- aws account - different account
- web identity - federated users
- SAML 2.0 - active directory
- custom trust policies - other

we select the AWS service, and then grant policies. there are aws managed policies which are pre-written by aws, and there are user custom policies, which can be more granular and specific. but the managed policies are a good place to start.

we can review the document and see the new _Principal_ field, which determines who can assume the policy (rather than what it can do).

we also <kbd>Create a User</kbd>, and we allow the user to access the web management console, and we require it change the password at the first log-in. we also create a group and give it full EC2 permissions. in the user role, we can <kbd>Create Access key</kbd> (but we delete those immediately)

</details>

#### Hosting the Employee Directory Application on AWS

<details>
<summary>
Demo of launching an EC2 machine.
</summary>

(video)

hosting the application with Amazon EC2, using the default VPC and other defaults. under the <cloud>EC2</cloud> service, we <kbd>Launch an instance</kbd> - a single virtual machine. we give it a name and use the linux AMI, the free tier, and under the network settings we click <kbd>edit</kbd> and choose the default VPC without subnet preferences. we need to <kbd>Add Security group</kbd> to allow HTTP and HTTPS traffic to reach the instance, and under the advanced details, we choose the role we created earlier as the instance profile. under the user date, we add the script that downloads the source code for the app and launches it with flask. after clicking <kbd>Launch instance</kbd>, we can wait a few minutes, and then use the public IP address to navigate into it in the browser.

</details>

#### Module 1 Knowledge Check

<details>
<summary>
Recap questions
</summary>

> - Q: What are the four main factors that you should take into consideration when choosing a Region?
> - A: Latency, price, service availability, and compliance
> - Q: Which of the following best describes the relationship between Regions, Availability Zones, and data centers?
> - A: Regions are a grouping of Availability Zones. Availability Zones are one or more discrete data centers.
> - Q: Which of the following is a benefit of cloud computing?
> - A: Pay as you go.

</details>

</details>

### MODULE 2: AWS Compute

<details>
<summary>
The compute options of AWS: EC2, containers and serverless.
</summary>

#### Compute as a Service

<details>
<summary>
Understanding Compute
</summary>

(video)

every application requires computing power, using on-premises compute requires choosing the correct hardware, buying it, shipping it, and then setting it up. once the servers are acquired, it's hard to get more, and harder to replace them if they are no longer needed. with cloud vendors, AWS has taken care of buying, maintaining and setting up the servers, so there is no need to wait before starting to use them.\
EC2 is the most basic compute service, but it's not the only one. most of the lessons will be spent on EC2, but there are also serverless compute and containers, which might be the preferable option.

Servers are the first building block for applications, they usually handle HTTP requests and send responses to the the client, but we consider any api-based communication as belonging to client-server model, even if they don't use HTTP requests. The clients sends a request, and the server handlers it.

Servers are computers connected to the internet, which run a server configuration such as:

- Windows - Internet Information Service (IIS)
- Linux - Apache HTTP server, Nginx, Apache Tomcat.

servers can be run on virtual machine instances (EC2), on containers or as serverless compute. EC2 is the fundamental option, where the AWS hypervisor creates a virtualized machine with the required OS and then we have a "virtual computer" that runs inside a physical computer.

</details>

#### Getting Started with Amazon EC2

<details>
<summary>
EC2 Basics
</summary>

> When architect-ing any application for high availability, consider using at least two EC2 instances in two separate Availability Zones.

(video)

EC2 machines are the most flexible and "controllable" compute options, they are billed by second or hour, and can be shut down to reduce costs. EC2 machines can run a variety of Operating systems, based on the AMI - Amazon Machine Image.\
A single AMI can create multiple machines, and it contains the OS, device mapping, launch permissions, and sometimes pre-installed services. AMIs are provided by AWS, by the community through the marketplace, and can be created for custom needs. in addition to the AMI, the ec2 also has the instance type and size, which determines the computing power and memory available to it. Instance types are grouped by the use case, for example, the G-family is designed to handle graphic intensive, while the M-family is general purpose. after choosing the type, the size (starting from nano and going up to many "EXTRA") determines the compute powers (number of cores). this means that the hardware can be changed, so there is no need to over-provision hardware when starting out. the machine can be changed based on need.\
This also makes trying new options easier, and even to scale up and scale out for short periods of increased demand.

##### Amazon EC2

> Amazon EC2 is a web service that provides secure, resizable compute capacity in the cloud. With this service, you can provision virtual servers called EC2 instances.\
> With Amazon EC2, you can do the following:
>
> - Provision and launch one or more EC2 instances in minutes.
> - Stop or shut down EC2 instances when you finish running a workload.
> - Pay by the hour or second for each instance type (minimum of 60 seconds).

EC2 instances have hardware specification (CPU, memory, network, storage) and logical configurations (networking location, firewall rules, authentication).

##### Amazon Machine Image (AMI)

> An AMI includes the operating system, storage mapping, architecture type, launch permissions, and any additional preinstalled software applications.

EC2 machines are instances of the AMI definition, when an EC2 machine is created, the AMI is copied into the root device volume, and then the machine is booted from it. AMIs are re-usable, and can configure tha starting sequence, the programs that are needed, and can be used again and again to spin up more EC2 instances.\
AMIS

- Quick start AMIs - commonly used AMIs created by AWS.
- Market Place - AMIs created by third-party verified vendors.
- Account AMIs - created from running EC2 instances.
- Community AMIs - created and shared by the AWS user community.
- Custom Image - built with the EC2 Image builder.

AMIs have unique identifiers, each starting with "ami-" and then a unique hash. AMIs are region bound, and the "same" AMI has different identifiers for each region.

The EC2 machine can be configured with different resources, effecting the power and pricing of it, the instance type notation is one letter for the family, then the generation number, any additional data a period and then the instance size.

| Instance family       | Description                                                                                                                                                                                                                                                                                                                | Use Cases                                                                                                                                                                                                                                                                |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| General purpose       | General purpose instances provide a balance of compute, memory, and networking resources, and can be used for a variety of workloads.                                                                                                                                                                                      | Ideal for applications that use these resources in equal proportions, such as web servers and code repositories                                                                                                                                                          |
| Compute optimized     | Compute optimized instances are ideal for compute-bound applications that benefit from high-performance processors.                                                                                                                                                                                                        | Well-suited for batch processing workloads, media transcoding, high performance web servers, high performance computing (HPC), scientific modeling, dedicated gaming servers and ad server engines, machine learning inference, and other compute intensive applications |
| Memory optimized      | Memory optimized instances are designed to deliver fast performance for workloads that process large datasets in memory.                                                                                                                                                                                                   | Memory-intensive applications, such as high-performance databases, distributed web-scale in-memory caches, mid-size in-memory databases, real-time big-data analytics, and other enterprise applications                                                                 |
| Accelerated computing | Accelerated computing instances use hardware accelerators or co-processors to perform functions such as floating-point number calculations, graphics processing, or data pattern matching more efficiently than is possible in software running on CPUs.                                                                   | Machine learning, HPC, computational fluid dynamics, computational finance, seismic analysis, speech recognition, autonomous vehicles, and drug discovery                                                                                                                |
| Storage optimized     | Storage optimized instances are designed for workloads that require high sequential read and write access to large datasets on local storage. They are optimized to deliver tens of thousands of low-latency random I/O operations per second (IOPS) to applications that replicate their data across different instances. | NoSQL databases (Cassandra, MongoDB and Redis), in-memory databases, scale-out transactional databases, data warehousing, Elasticsearch, and analytics                                                                                                                   |
| HPC optimized         | High performance computing (HPC) instances are purpose built to offer the best price performance for running HPC workloads at scale on AWS.                                                                                                                                                                                | Ideal for applications that benefit from high-performance processors, such as large, complex simulations and deep learning workloads                                                                                                                                     |

Instances are also created in specific locations, meaning which Availability Zone hosts them. applications should be designed for speed and high availability.

</details>

#### Amazon EC2 Instance Lifecycle

<details>
<summary>
Virtual Machine life cycle and pricing.
</summary>

> An EC2 instance transitions between different states from the moment you create it until its termination.

(video)

Provisioning and removing instances, scaling the compute power as needed, so the application always has enough compute resources, but without paying extra for unused compute power.

EC2 instances have life cycle (states)

1. it starts from an AMI in the **PENDING** state
2. moving to **RUNNING** state
3. rebooting the instance temporarily to the **REBOOTING** state before returning th **RUNNING**
4. stopping the instance, moving to **STOPPING** and **STOPPED** states, which require the **PENDING**.
5. pausing (stop-and hibernate), which saves the machine to memory, so after **STOPING** and **STOPPED**, it can return to **RUNNING** immediately.
6. terminating the instance **SHUTTING DOWN** and **TERMINATED** which removes all data from it.

there is a service called _termination protection_ which can recover data from terminated instances for a limited period of time.

EC2 machines are charged for the **RUNNING** and **STOPPING** state.

> 1. When you launch an instance, it enters the **pending** state. When an instance is pending, billing has not started. At this stage, the instance is preparing to enter the running state. Pending is where AWS performs all actions needed to set up an instance, such as copying the AMI content to the root device and allocating the necessary networking components.
> 2. When your instance is **running**, it's ready to use. This is also the stage where billing begins. As soon as an instance is running, you can take other actions on the instance, such as reboot, terminate, stop, and stop-hibernate.
> 3. When you **reboot** an instance, it’s different than performing a stop action and then a start action. Rebooting an instance is equivalent to rebooting an operating system. The instance keeps its public DNS name (IPv4) and private and public IPv4 addresses. An IPv6 address (if applicable) remains on the same host computer and maintains its public and private IP address, in addition to any data on its instance store volumes.
> 4. When you stop your instance, it enters the **stopping** and then **stopped** state. This is similar to when you shut down your laptop. You can stop and start an instance if it has an Amazon Elastic Block Store (Amazon EBS) volume as its root device. When you stop and start an instance, your instance can be placed on a new underlying physical server. Your instance retains its private IPv4 addresses and if your instance has an IPv6 address, it retains its IPv6 address. When you put the instance into stop-hibernate, the instance enters the stopped state, but saves the last information or content into memory, so that the start process is faster.
> 5. When you terminate an instance, the instance stores are erased, and you lose both the public IP address and private IP address of the machine. Termination of an instance means that you can no longer access the machine. As soon as the status of an instance changes to **shutting down** or **terminated**, you stop incurring charges for that instance.

when in the stop-and-hibernate state, the state of the machine is saved to the EBS (elastic block storage) volume, so it's quicker to recover. however, it also incurs costs. not all instances can hibernate.

##### Pricing

pricing depends on the instance compute power and other factors.

- **On-Demand Instances** are EC2 machine that are provisioned directly, without up-front payment or commitments. they are used for flexibility, for short term workloads that can not be interrupted, and for testing and developing.
- **Spot Instances** have flexible start and end time, and can run when demand to EC2 machines is low, and therefore have lower costs. they are used for workloads which can happen at flexible times, can be stopped and resumed (or are stateless). the user can set a limit on spending, and bid for the price in which they want to provision a machine, and once the price falls to that level, then the instance is provisioned. pricing is set by AWS based on capacity.
- **Saving Plans** are long term contracts with AWS to provision resources over a long period of time (one year or 3 years), this provides a significant discount over "on-demand" instances, if the workloads have consistent and steady usage pattern and the user can make those up-front payment.
- **Reserved Instances** also have 1 year or 3 year plans, but with different payment options (all upfront, partial upfront, no upfront). they provided a cheaper alternative to "on-demand" instances when there is an expectation for steady workloads:
  - _Standard Reserved Instances_
  - _Convertible Reserved Instances_ - allow for changing the instance type to a stronger machine at a discounted price
  - _Scheduled Reserved Instance_ - reserving only for a window of time.
- **Dedicated Host** - a physical EC2 server that is only used by the client, and can help with cost reduction by incorporating software licenses and it might be part of compliance requirements. can be purchased on demand (hourly) or reserved for a discount.

</details>

#### Demonstration: Launching the Employee Directory Application on Amazon EC2

<details>
<summary>
Demo of Creating the EC2 web server.
</summary>

(video)

launching the EC2 on the default VPC. <kbd>Launch Instance</kbd>, provide name, choose AMI template, instance type and size (some belong to the free tier, while some have GPU), we can use a key-pair for ssh connection. under the network settings, we use the default the VPC and subnets, with the gateway for public internet access. with choose to auto-assign a public ip address.\
Under firewall options, we can set the Security group or create a new one, we can then add storage if needed. under advanced details, we can set the IAM instance profile with the role, and we paste the provided bash script into the "User data" section to have it run when the machine is booted. the script contains installation of packages and running the program itself.

</details>

#### Container Services

<details>
<summary>
Understanding Containers and container orchestration.
</summary>

(video)

some applications work better with container compute, rather than EC2 compute. for containers workloads (docker), it might be better to use container based compute services, such as ECS and EKS, which provide built-in container orchestration. this allows to scale hosts and scale the containers themselves. containers can run on top of EC2 machines or on AWS Fargate.

Containers are a standardized method of isolating processes and running them, which allows for portability and flexibility across machines as all the needed requirements are part of the container image. containers are similar in some regards to virtual machines, but they are more lightweight (and spin up faster), as they don't each maintain a copy of the operating system.\
In AWS containers can be run on EC2 machines, this could be done manually, but it becomes hard to manage as the number of machines scales. for theses cases, it's recommended to use an orchestrating service. the orchestrator handles:

> - How to place your containers on your instances
> - What happens if your container fails
> - What happens if your instance fails
> - How to monitor deployments of your containers

##### Managing containers with Amazon ECS

ECS is an amazon service that manages containers end-to-end. the containers are defined as "tasks" and can rn either on EC2 machines or on fargate instances.

- cluster - logical grouping of services, tasks and capacity providers in a region.
- Service - one or more identical tasks. checks and replaces unhealthy tasks.
- Task - one or more containers, specify compute, networking, IAM and configurations.

when using an EC2 to run ECS, the ECS agent is installed and manages the compute instance. this can be done for Linux and Windows machines. ECS cluster can manage launching and stopping containers, scaling containers, placing the container across the cluster, and assigning permissions. Tasks are defined as json documents with the image, containers and resources required for the container.

##### Using Kubernetes with Amazon EKS

EKS is the AWS service for running kubernetes clusters. Kubernetes is an open-source platform for managing containers, and it is popular and supported by many cloud vendors.

</details>

#### Introduction to Serverless

<details>
<summary>
Why go Serverless?
</summary>

> Spend time on the things that differentiate your application, rather than spending time on ensuring availability, scaling, and managing servers.

(video)

using virtual machines and containers requires setting up the infrastructure for maintenance, patching and high availability. while this is easier than running on-premises servers, and it does give a greater level of control, it might be too much for some applications.\
AWS serverless services hide away the server running the application, and don't allow access or knowledge of it. instead, all the management is done by AWS, and the user takes care of the application, rather than the infrastructure running it. this ties into the shared responsibility model, as we move into serverless options, AWS handles more and more of the actions.

Running a server on EC2 still requires the user to handle the OS, security patching, networking, storage and scaling. the serverless model abstracts away those issues and lets AWS handle them, while the user can focus more on the application and the value it provides. Serverless also has built-in scaling, high availability and fault tolerance.

</details>

#### Serverless with AWS Fargate

<details>
<summary>
Run Containers on Managed Compute Clusters
</summary>

> AWS Fargate scales and manages the infrastructure, so developers can work on what they do best, application development.

(video)

AWS fargate are compute instances which can be used instead of EC2 instances. instead of setting up EC2 machines to hold the containers, the containers are deployed to a managed cluster which does everything for you. Fargate supports spot and reserved instances for reduced costs. Fargate clusters have built on scaling options and use "per-per-use" modeling to avoid paying for under-utilized EC2 machines

</details>

#### Serverless with AWS Lambda

<details>
<summary>
Event based computing with Lambda functions.
</summary>

(video)

AWS lambda are serverless compute code that are activated by triggers and execute code in response to it. when the trigger is detected, the code runs in a managed environment. if the demand increases, then more AWS lambdas are deployed, all with the same code in an isolated environment.\
Lambdas are limited for 15 minutes, so they aren't intended for long running code, but they can be used for easy, scalable processing.\
In our example, we create a lambda to resize photos when they are uploaded to the S3 Bucket. we go the the <kbd>Lambda</kbd> service, and choose where the code comes from (from scratch, from a blueprint, or from a container). we can use different runtime (programming languages), and we assign the trigger source to start the lambda, in our case, we care about creating new files in the S3 bucket at a specific location, and we want to output it to another location. we also set up the IAM role for permissions. we can test this by adding an image to the S3 bucket.

Lambda code runs without provisioning and running a server. it can be used for many cases, with all kinds of runtime environments and programming languages. scaling and high availability are built-in into the lambda service.

- Function - the code that is invoked.
- Trigger - which event trigger the code.
- Events - The data that is passed to the code.
- Application environment - the isolated compute runtime and resource in which the code is executed.
- Deployment Package - either a zip file or a container image, when the lambda isn't taken from a blueprint or written from scratch.
- Runtime - Language specific runtime environment for each programming language, or even a custom environment.
- Lambda function Handler - the entry point of the lambda, the function in the code that meets the event.

Lambda billing is done on a granular bases, rounded up to the nearest millisecond of duration.

</details>

#### Choosing the Right Compute Service

<details>
<summary>
How to choose AWS services based on requirements.
</summary>

(video)

choosing the right service for the use case. practice questions.

</details>

#### Module 2 Knowledge Check

<details>
<summary>
Recap Questions.
</summary>

> - Q: What does an Amazon EC2 instance type indicate?
> - A: Instance family and instance size
> - Q: Which of the following is true about serverless?
> - A: You never pay for idle resources.

</details>

</details>

### MODULE 3: AWS Networking

<details>
<summary>
Networking and VPCs.
</summary>

#### Introduction to Networking

<details>
<summary>
Networking is how you connect computers around the world and allow them to communicate with one another.
</summary>

(video)

we focus on the network <cloud>VPC</cloud> and it's components. when we created an <cloud>EC2</cloud> machine, we put it into a vpc. we used a default vpc that AWS creates for us. the default VPC has inbound access from the internet, so it's dangerous to use it all the time. some services don't require VPCs, but it's still important to learn about it.

> Networking defined \
> Networking is how you connect computers around the world and allow them to communicate with one another. In this course, you’ve already seen a few examples of networking. One is the AWS Global Infrastructure. AWS has built a network of resources using data centers, Availability Zones, and Regions.

networking with analogy to letters and postal services, we need the letter (contents), but we also need the "from" and "to" addresses. in the digital world, this process is called _routing_, and it uses ip addresses.

**ipv4 notations** groups 32 bits into four groups of eight, and separates them by dots. so we have four units of numbers between zero and 255. the **CIDR notation** is a way to express ranges of ip addresses. it uses fixed and flexible parts. for example. _192.168.1.0/24_ the number after the slash is the number of **fixed** bits (masked bits), and the rest are flexible, this means that there are 8 flexible bits in this. so the smaller the number, the larger the range.
aws supports ranges from `/28` (which has 16 ip options : $2^{32-28}=16$) and up to `/16`, ($2^{32-16}=65,536$).

</details>

#### Amazon VPC

<details>
<summary>
A virtual private cloud (VPC) is an isolated network that you create in the AWS Cloud, similar to a traditional network in a data center.
</summary>

> To maintain redundancy and fault tolerance, create at least two subnets configured in two Availability Zones.

(video)
VPC - virtual private cloud, a boundary around applications and resources, from the internet and from other aws resources. VPC go inside an Availability Zone and have an IP range (CIDR), we will use oregon and 10.1.0.0/16. under the services bar, we choose <kbd>VPC</kbd> and then <kbd>create VPC</kbd>, we choose the name and the ip range. inside the vpc, we divide the resources into logical groups called _subnets_. we can separate subnets and have one public and open to internet access, and one private. The subnet lives inside a VPC, is attached to a specific Availability Zone, and has an ip range that is a subset of the VPC range (such as 10.1.1.0/24 for the public subnet, and 10.1.3.0/24 for the private subnet). inside the vpc page, we choose <kbd>Create Subnet</kbd>, choose the name, the vpc, the Availability Zone and provide the CIDR range.\
All the resources in the vpc are isolated from others, so to allow Internet connectivity, we need an <cloud>Internet Gateway</cloud> (attached to the vpc). <kbd>Create Internet Gateway</kbd>, then <kbd>Attach to VPC</kbd>. if we want to limit internet access and only allow access from a specific location, we can use <cloud>Virtual Private Gateway</cloud> (VPN). ideally we would want to have duplicated resources in a different Availability Zone for high availability.

If VPCs are like local networks, the subnets are similar to virtual local networks, we use subnets to isolate the access to resources inside the vpc, public subnets are connected to outside AWS, while private subnets only communicate with resources in the AWS VPC.

AWS has five reserved ip addresses

| IP address | Reserved For              |
| ---------- | ------------------------- |
| 10.0.0.0   | Network Address           |
| 10.0.0.1   | VPC local router          |
| 10.0.0.2   | DNS server                |
| 10.0.0.3   | Future use                |
| 10.0.2.255 | Network broadcast address |

> The five reserved IP addresses can impact how you design your network. A common starting place for those who are new to the cloud is to create a VPC with an IP range of /16 and create subnets with an IP range of /24. This provides a large amount of IP addresses to work with at both the VPC and subnet levels.

##### Gateways

Internet gateways act like a modem - they connect the VPC to the external internet. Virtual private gateways are VPN, they connect only to a matching customer gateway (a device or software) and provide encrypted communication between the two sides. There s also <cloud>AWS Direct Connect</cloud>, which is a physical connection between the on-premises data center and AWS metacenter, which has high speed and doesn't travel through the public internet at all.

</details>

#### Amazon VPC Routing

<details>
<summary>
Routing External Traffic to the subnets.
</summary>

> A route table contains a set of rules, called routes, that determine where network traffic from your subnet or gateway is directed.

(video)

once the traffic got to the internet gateway, we need to route it to the correct subnet and resources. this is done via the <cloud>Route Table</cloud>, it contains rules (routes) that move traffic around. a route table can at the vpc level or the subnet level. aws creates a default main route-table, which only allows connections between the local vpcs. in the route table <kbd>Routes</kbd> table, we can see the routing to the subnets (10.1.0.0/16). the connection from the internet gateway to the subnets determines if they are public or private, if a route exists, then it's a public subnet, otherwise it's private.\
in our app we create custom route tables, we click <kbd>Create Route Table</kbd>, give it a name and a vpc, and the <kbd>edit routes</kbd> and <kbd>add routes</kbd>, and select the target type from the drop down.

| Destination | Target           |
| ----------- | ---------------- |
| 10.1.0.0/16 | local            |
| 0.0.0.0/0   | internet gateway |

we next need to associate the table with the subnets, so we click <kbd>Edit subnet Associations</kbd> and choose the public subnets.

> Main route table:\
> When you create a VPC, AWS creates a route table called the main route table. A route table contains a set of rules, called routes, that are used to determine where network traffic is directed. AWS assumes that when you create a new VPC with subnets, you want traffic to flow between them. Therefore, the default configuration of the main route table is to allow traffic between all subnets in the local network. \
> The following rules apply to the main route table:
>
> - You cannot delete the main route table.
> - You cannot set a gateway route table as the main route table.
> - You can replace the main route table with a custom subnet route table.
> - You can add, remove, and modify routes in the main route table.
> - You can explicitly associate a subnet with the main route table, even if it's already implicitly associated.

in addition to the main route table, we can also specify more granular behavior by using a custom route table. associating a route table with a subnet will replace the main one.

</details>

#### Amazon VPC Security

<details>
<summary>
Security With Access Control Lists and Security Groups.
</summary>

> Cloud security at AWS is the highest priority. You benefit from a data center and network architecture that is built to meet the requirements of the most security-sensitive organizations.

(video)

any new VPC is isolated from internet access, because it doesn't have an internet gateway associated with the route table. but once we connect the subnet to the external internet, we need to add security. the two options available for us are <cloud>Access Control Lists (ACL)</cloud> and <cloud>Security Groups</cloud>. \
An ACL is like a firewall on the subnet level. the control inbound and outbound traffic. we can specify which access types are allowed (HTTP, HTTPS, SSH, etc..) by creating inbound rules and specifying the source. in ACL, we also need to open the corresponding outbound rules. this is because ACL are considered _stateless_.\
Security groups are created at the EC2 level, and they are mandatory (every EC2 has them). the default Security group behavior is to block inbound requests and allow outbound. if we want to allow requests from the internet to reach us, we need to open inbound rules(such as port 80 and 443 for http and https). security groups are _stateful_, so if there is an inbound rule that allows traffic, outbound traffic to the same destination will be allowed as well.
Security groups only have "allow" rules, while ACL have both "allow" and "deny" rules, and can control inbound and outbound traffic separately.

Networks ACLs have default inbound and outbound rules which allow all traffic (internal and external) to flow, we can customize the rules as we like to allow access from each kind of protocol based on the ip source, and deny all other traffic.

EC2 security groups also have defaults, for security groups, the default behavior is to deny inbound access and allow all outbound, but since they are stateful, then any connection that was initially established by the EC2 machine will allow inbound traffic. Security groups only have "Allow" rules, and not "Deny" rules.

</details>

#### Demonstration: Relaunching the Employee Directory Application in Amazon EC2

<details>
<summary>
Demo.
</summary>

(video)

Creating a vpc, 4 subnets, route table, internet gateway. in the VPC dashboard.
We start by clicking <kbd>Create a VPC</kbd>, choose VPC only. choose cidr range of 10.1.0.0/16. next we choose <kbd>Create subnet</kbd>, select the new vpc, choose name, Availability Zone and cidr range (10.1.1.0/24, 10.1.2.0/24), we do this twice - two subnets for each Availability Zone - one private and one public. the cidr ranges should not overlap with one another.\
Next we <kbd>create internet gateway</kbd>, click <kbd>Attach to VPC</kbd> and select our new vpc. Internet gateways have one-to-one relationship, an internet gateway can only be attached a to a single VPC. we next <kbd>Create Route Table</kbd> for the public subnets, and we attach it to VPC. we next click <kbd>Edit routes</kbd> and <kbd>Add route</kbd> and choose the destination as 0.0.0.0/0 (every ip address) and the target as our internet gateway. now we have two routes, the default route inside the vpc and the route to the internet. we next need to associate the route table with the subnets so we click <kbd>Edit subnet association</kbd> and choose the public subnets.. subnets without explicit associations use the main route table.\
We navigate back to the EC2 dashboard, we select the existing machine, and under <kbd>Actions</kbd>, we select <kbd>Image and Template</kbd> and <kbd>Launch more like this</kbd>. this will populate a wizard with the same values as the original, and here we can modify the VPC and the subnets, and set "auto-assign public IP" to true. now we need to choose a different security group, because the security groups are attached to the VPC. so we <kbd>Create a security group</kbd> and allow HTTP and HTTPS access. if we can access it from the public ip address, then it means we did things correctly.

</details>

#### Module 3 Knowledge Check

<details>
<summary>
Recap Questions.
</summary>

> - Q: Which of the following can a route table be attached to?
> - A: Subnets (or entire VPC).
> - Q: Which of the following is true for a security group's **default** setting?
> - A: It blocks all inbound traffic and allows all outbound traffic.
> - Q: A network access control list (network ACL) filters traffic at the Amazon EC2 instance level.
> - A: False (subnet level).

</details>

</details>

### MODULE 4: AWS Storage

<details>
<summary>
AWS storage services - File, Block and Object Storage.
</summary>

> With cloud computing, you can create, delete, and modify storage solutions within a matter of minutes.

#### Storage Types

<details>
<summary>
Basic Storage Types.
</summary>

(video)
Our demo application requires storage of different kinds, we have files database, static data and structured data. for the the static data (in this case, photos), we can use <cloud>Block Storage</cloud> or <cloud>Object Storage</cloud>. block storage splits the file into blocks, while object storage keeps the data in it's entirety. if we want to change the data, it's easy to do with block storage, but with object storage this requires rewriting the entire object. because of that, object storage often follows the "WORM" pattern - write once, read many. in our case, static data that doesn't change like the photos are more fitting to object storage, and application data and files are better fit for block storage.

> AWS storage services are grouped into three categories: <cloud>file storage</cloud>, <cloud>block storage</cloud>, and <cloud>object storage</cloud>. In file storage, data is stored as files in a hierarchy. In block storage, data is stored in fixed-size blocks. And in object storage, data is stored as objects in buckets.

when compared to traditional storage options, adding block storage is like adding a direct attached storage to the network, while file storage are often supported with <cloud>NAS</cloud> (network attached storage) servers. the benefit of the cloud is that there is no need to purchase the storage device and install it, everything is handled in a click of a button.

##### File Storage

<cloud>File Storage</cloud> is the same as what we tend to use in our personal computer, files are organized into files and folder in an hierarchical order. using cloud file storage is ideal for centerline access to files that are shared and managed by multiple users, with file locking protocols built-in. other use cases are:

###### Web serving

> Cloud file storage solutions follow common file-level protocols, file naming conventions, and permissions that developers are familiar with. Therefore, file storage can be integrated into web applications.

###### Analytics

> Many analytics workloads interact with data through a file interface and rely on features such as file lock or writing to portions of a file. Cloud-based file storage supports common file-level protocols and has the ability to scale capacity and performance. Therefore, file storage can be conveniently integrated into analytics workflows.

###### Media and entertainment

> Many businesses use a hybrid cloud deployment and need standardized access using file system protocols (NFS or SMB) or concurrent protocol access. Cloud file storage follows existing file system semantics. Therefore, storage of rich media content for processing and collaboration can be integrated for content production, digital supply chains, media streaming, broadcast play-out, analytics, and archive.

###### Home directories

> Businesses wanting to take advantage of the scalability and cost benefits of the cloud are extending access to home directories for many of their users. Cloud file storage systems adhere to common file-level protocols and standard permissions models. Therefore, customers can lift and shift applications that need this capability to the cloud.

##### Block Storage

<cloud>Block Storage</cloud> splits data into fixed-size chunks of data with individual addresses, rather than as object or files. this means that individual blocks can be retrieved directly, each block has metadata associated with it.

> Because block storage is optimized for low-latency operations, it is a preferred storage choice for high-performance enterprise workloads and transactional, mission-critical, and I/O-intensive applications.

###### Transactional workloads

> Organizations that process time-sensitive and mission-critical transactions store such workloads into a low-latency, high-capacity, and fault-tolerant database. Block storage allows developers to set up a robust, scalable, and highly efficient transactional database. Because each block is a self-contained unit, the database performs optimally, even when the stored data grows.

###### Containers

> Developers use block storage to store containerized applications on the cloud. Containers are software packages that contain the application and its resource files for deployment in any computing environment. Like containers, block storage is equally flexible, scalable, and efficient. With block storage, developers can migrate the containers seamlessly between servers, locations, and operating environments.

###### Virtual machines

> Block storage supports popular virtual machine (VM) hypervisors. Users can install the operating system, file system, and other computing resources on a block storage volume. They do so by formatting the block storage volume and turning it into a VM file system. So they can readily increase or decrease the virtual drive size and transfer the virtualized storage from one host to another

##### Object Storage

<cloud>Object Storage</cloud> is similar to file storage, but it uses a flat structure rather than hierarchical nested structure. object data is stored in "buckets", and it has support for data of larger size than regular file storage. this is useful for storing large and unstructured data.

###### Data archiving

> Cloud object storage is excellent for long-term data retention. You can cost-effectively archive large amounts of rich media content and retain mandated regulatory data for extended periods of time. You can also use cloud object storage to replace on-premises tape and disk archive infrastructure. This storage solution provides enhanced data durability, immediate retrieval times, better security and compliance, and greater data accessibility.

###### Backup and recovery

> You can configure object storage systems to replicate content so that if a physical device fails, duplicate object storage devices become available. This ensures that your systems and applications continue to run without interruption. You can also replicate data across multiple data centers and geographical regions.

###### Rich media

> With object storage, you can accelerate applications and reduce the cost of storing rich media files such as videos, digital images, and music. By using storage classes and replication features, you can create cost-effective, globally replicated architecture to deliver media to distributed users.

</details>

#### File Storage with Amazon EFS and Amazon FSx

<details>
<summary>
File Storage Options.
</summary>

<cloud>Amazon Elastic File System (EFS)</cloud> is aws service for file systems, it offers and endlessly growing, automatically scaling file system that can be mounted both on aws instances and to on-premises machine, thousands of concurrent instances can interact with the same file server concurrently. Payment is based on storage, with either the standard storage class for high availability across multiple Availability Zone or with a single zone storage that is cheaper.

AWS also supports <cloud>Amazon FSx</cloud>, a fully managed service that supports different kinds of file systems:

- Lustre - Fast, cost effective, integrates with AWS services directly, handles large workloads.
- NetApp ONTAP (<cloud>FSxN</cloud>) - rich management system that is accessible via standard API from windows, linux and macOS.
- OpenZFS (<cloud>FSxZ</cloud>) - low latency, low cost, works great for small file workloads.
- Windows File Server - a drop-in replacement for windows servers, fully managed and uses the <cloud>Service Message Block (SMB)</cloud> protocol.

</details>

#### Block Storage with Amazon EC2 Instance Store and Amazon EBS

<details>
<summary>
Block Storage, instance Storage and EBS.
</summary>

> The unique characteristics of block storage make it the preferred option for transactional, mission-critical, and I/O-intensive applications

(video)
every EC2 instance has a block storage, it can either be the boot storage or just as storage. the internal storage is called <cloud>Instance store</cloud>, and the external store is <cloud>Elastic Block Storage</cloud>.\
The instance storage is directly attached to the physical server, which makes it faster. however, it also means that it has the same lifecycle as the virtual machine, this means the data is _ephemeral_. for data that should outlive the instance, block storage is used. \
These drives are separate from the EC2 instance. an EC2 machine can interact with many block volumes, but only one machine can interact with any block storage. some types of of block storage can be attached to multiple instances, this is called <cloud>EBS Multi-Attach</cloud>. The EBS volume is separate from the EC2 instance, it has persistent storage that isn't tied to the lifecycle of the machine. so if there is need for persistent workloads, then EBS is usually the preferred option.\
There are also different types of EBS volumes, divided between SSD and HDD volumes, there is also a need to backup data, this is done via <cloud>EBS snapshots</cloud>, which are incremental and are stored redundantly.

##### Amazon EC2 Instance Store

every EC2 instance has a temporary block level storage attached to it. this storage is usually physically attached to the rack, so it has lower latency. however, it also means that it is "ephemeral", and will disappear when the machine goes down. this is useful for data that can be replicated to the machine (such as hadoop workloads) to take advantage of the lower latency, as well as temporary storage of frequently changing data such as buffer and caches.

##### Elastic Block Storage

<cloud>Elastic Block Storage</cloud> are block level storage that are independent from the instance. they are usually network attached, and act like external drives. they are _detachable_ and can be reattached to other EC2 instances, they have a _distinct_ lifetime from the machine that uses it, and they are size limited - they cannot endlessly grow.\
Most of the EBS volumes can only be attached to one EC2 machine, but there are a few <cloud>EBS multi-attached</cloud> types that can be attached to multiple EC2 machines in the same Availability Zone.

EBS storage can also be scaled in some cases, even when provisioning a specific size of EBS volume, the size can sometimes be increased. the current limit is 64Tb drives. beyond that, an EC2 machine can have more than one EBS volume attached to it, further increasing the available storage.

> Amazon EBS is useful when you must retrieve data quickly and have data persist long term. Volumes are commonly used in the following scenarios.

- Operating Systems - storing the boot and root volume, this is usually done for amazon AMIs, commonly called <cloud>EBS backed AMIs</cloud>.
- Databases - a storage layers for database running on EC2 that need consistency and low-latency performance.
- Enterprise applications - high availability and durability are important when running a production workload.
- Big analytics volumes - large amounts of data that needs to be attached and removed dynamically.

EBS volumes are also divided into types: either HDD or SSD based. there are even further division for each of those types.
in general, ssd devices are faster for frequent read/write operations with small I.O size, while HDD devices are more fit for large continues workloads with higher throughput. HDD devices can also be cheaper, if the "Cold HDD Volume" is chosen.

> - balanced: Provides a balance of price and performance for a wide variety of transactional workloads
> - high-performance - Provides high-performance SSD designed for latency-sensitive transactional workloads

| name              | Volume type                                         | volume Size    | Max IOOS | Max throughput | Multi attach support |
| ----------------- | --------------------------------------------------- | -------------- | -------- | -------------- | -------------------- |
| gp3               | balanced                                            | 1 GiB-16 TiB   | 16,000   | 1,000 MiB/s    | Not Supported        |
| gp2               | balanced                                            | 1 GiB-16 TiB   | 16,000   | 250 MiB/s      | Not Supported        |
| io2 Block Express | performance                                         | 4 GiB-64 TiB   | 256,000  | 4,000 MiB/s    | Supported            |
| io2               | performance                                         | 4GiB-16 TiB    | 64,000   | 1,000 MiB/s    | Supported            |
| io1               | performance                                         | 4GiB-16 TiB    | 64,000   | 1,000 MiB/s    | Supported            |
| st1               | low-cost, frequently accessed, throughput-intensive | 125 GiB-16 TiB | 500      | 500 Mib/s      | Not Supported        |
| sc1               | lowest cost                                         | 125 GiB-16 TiB | 250      | 250 Mib/s      | Not Supported        |

Elastic block storage has benefits for:

- High Availability - automatically replicated in it's Availability Zone.
- Data Persistence - Storage outlives the EC2 machine (not ephemeral).
- Data Encryption - all volumes support encryption.
- Flexibility - attach and detach from EC2 machines without stopping the instance.
- Backups - Any EBS volume can be backed-up.

<cloud>EBS snapshots</cloud> are the way that AWS supports back ups, each snapshot contains only the changes since the previous snapshot, so each one is incremental and doesn't replicate the entire size of the volume. backups are managed by AWS inside S3 buckets, and can be used to create EBS volumes based on the snapshot.

</details>

#### Object Storage with Amazon S3

<details>
<summary>
Object Storage in S3 Buckets - Storage tiers, Access Control, Life Cycle and Versioning.
</summary>

> Object storage is built for the cloud and delivers virtually unlimited scalability, high durability, and cost effectiveness.

(video)
EBS volumes have size limitations, and not all can be connected to any number of ec2 machines (multi attach). aws has a dedicated storage service that is independent from compute. <cloud>S3</cloud> is such a service. it has infinite size, and can store objects up to 5Tb in size. it isn't connected to any EC2 instance. S3 is object storage, using flat data without hierarchy. it is also distributed across many data centers, and it supports high level of availability and durability.\
All objects are stored in <cloud>S3 bucket</cloud>. we can navigate to the S3 service and select <kbd>Create Bucket</kbd>. buckets are region specific. they have a **globally unique** name (which is dns compliant - no white space, no special characters). once the bucket is created, we can <kbd>Upload Files</kbd> to add objects to it. each object has a unique url that is the combination of the the bucket unique name and the object unique key.\
Access to objects is private by default, they can be explicitly opened to public access. if we with to do so, we need to:

- make the bucket allow public access under the "permissions" tab.
- edit the <kbd>Object Ownership</kbd> and allow using <cloud>Access Control Lists</cloud>.
- once this is done, we can choose a specific item, under <kbd>Actions</kbd> select <kbd>Make public using ACL</kbd> to make it public and allow access through the URL address.

we usually want to have granular access to resources, this can be done with <cloud>IAM</cloud> roles and policies, or with <cloud>S3 Bucket Policies</cloud>. bucket policies use the same json format language, but bucket policies are only attached to buckets, and specify which actions are allowed or denied on the bucket.

S3 stands for <cloud>Simple Storage Service</cloud>, and is one of aws core services. it stores objects in a flat structure. each bucket is inside a region, and is replicated across multiple Availability Zones. bucket names must be unique all across all AWS, and must follow these rules

> - Bucket names must be between 3 (min) and 63 (max) characters long.
> - Bucket names can consist only of lowercase letters, numbers, dots (.), and hyphens (-).
> - Bucket names must begin and end with a letter or number.
> - Buckets must not be formatted as an IP address.
> - A bucket name cannot be used by another AWS account in the same partition until the bucket is deleted.

objects inside the bucket also must have a unique name, this is the object key. the key can contain slashes and can be formatted to support a hierarchical structure by using prefixes to represent folder.

Common S3 use cases are:

> - Backup and storage - Amazon S3 is a natural place to back up files because it is highly redundant. As mentioned in the last lesson, AWS stores your EBS snapshots in Amazon S3 to take advantage of its high availability.
> - Media hosting - Because you can store unlimited objects, and each individual object can be up to 5 TB, Amazon S3 is an ideal location to host video, photo, and music uploads.
>   Software delivery - You can use Amazon S3 to host your software applications that customers can download.
> - Data lakes - Amazon S3 is an optimal foundation for a data lake because of its virtually unlimited scalability. You can increase storage from gigabytes to petabytes of content, paying only for what you use.
> - Static websites - You can configure your S3 bucket to host a static website of HTML, CSS, and client-side scripts.
> - Static content - Because of the limitless scaling, the support for large files, and the fact that you can access any object over the web at any time, Amazon S3 is the perfect place to store static content.

Everything inside a S3 bucket is private by default, only available to the user which created the bucket. this can be changed to allow public access to everybody or to allow limited access.

S3 access can be configured with the same <cloud>IAM policies</cloud> as other resources, by attaching the permissions to users, groups and roles. this is useful as a centralized way to manage S3 bucket. Alternatively, <cloud>S3 policies</cloud> can be used on buckets directly to manage access. this can be done as a simple way for creating Cross-account access, or if a policy exceeds the size limit for an <cloud>IAM policy</cloud>.

All S3 objects use encryption at test and at transit with no additional costs.

##### S3 Storage Classes

Objects in S3 have <cloud>Storage Class</cloud>, this determines access patterns, costs, redundancy and how the objects are backed up. if no storage class is specified, then the <cloud>S3 standard is used</cloud>. there are standard tiers which allow for immediate access and <cloud>Glacier</cloud> tiers which are low cost backup tiers, which are for data which isn't expected to be accessed. Glacier tier either take longer to access the data or cost much more to do so immediately. there are also <cloud>S3 Outposts</cloud> which store data at aws outposts on-premises using the S3 API.

> - <cloud>S3 Standard</cloud> - This is considered general-purpose storage for cloud applications, dynamic websites, content distribution, mobile and gaming applications, and big data analytics.
> - <cloud>S3 Intelligent-Tiering</cloud> - This tier is useful if your data has unknown or changing access patters. S3 Intelligent-Tiering stores objects in three tiers: a frequent access tier, an infrequent access tier, and an archive instance access tier. Amazon S3 monitors access patterns of your data and automatically moves your data to the most cost-effective storage tier based on frequency of access.
> - <cloud>S3 Standard-Infrequent Access (S3 Standard-IA)</cloud> - This tier is for data that is accessed less frequently but requires rapid access when needed. S3 Standard-IA offers the high durability, high throughput, and low latency of S3 Standard, with a low per-GB storage price and per-GB retrieval fee. This storage tier is ideal if you want to store long-term backups, disaster recovery files, and so on.
> - <cloud>S3 One Zone-Infrequent Access (S3 One Zone-IA)</cloud> - Unlike other S3 storage classes that store data in a minimum of three Availability Zones, S3 One Zone-IA stores data in a single Availability Zone, which makes it less expensive than S3 Standard-IA. S3 One Zone-IA is ideal for customers who want a lower-cost option for infrequently accessed data, but do not require the availability and resilience of S3 Standard or S3 Standard-IA. It's a good choice for storing secondary backup copies of on-premises data or easily re-creatable data.
> - <cloud>S3 Glacier Instant Retrieval</cloud> - Use S3 Glacier Instant Retrieval for archiving data that is **rarely accessed and requires millisecond retrieval**. Data stored in this storage class offers a cost savings of up to 68 percent compared to the S3 Standard-IA storage class, with the same latency and throughput performance.
> - <cloud>S3 Glacier Flexible Retrieval</cloud> - S3 Glacier Flexible Retrieval offers low-cost storage for archived data that is accessed 1-2 times per year. With S3 Glacier Flexible Retrieval, your data can be accessed in as little as 1-5 minutes using an **expedited retrieval**. You can also request free **bulk retrievals** in up to 5-12 hours. It is an ideal solution for backup, disaster recovery, offsite data storage needs, and for when some data occasionally must be retrieved in minutes.
> - <cloud>S3 Glacier Deep Archive</cloud> - S3 Glacier Deep Archive is the lowest-cost Amazon S3 storage class. It supports long-term retention and digital preservation for data that might be accessed once or twice a year. Data stored in the S3 Glacier Deep Archive storage class has a default retrieval time of 12 hours. It is designed for customers that retain data sets for 7-10 years or longer, to meet regulatory compliance requirements. Examples include those in highly regulated industries, such as the financial services, healthcare, and public sectors.
> - <cloud>S3 on Outposts</cloud> - Amazon S3 on Outposts delivers object storage to your on-premises AWS Outposts environment using S3 API's and features. For workloads that require satisfying local data residency requirements or need to keep data close to on premises applications for performance reasons, the S3 Outposts storage class is the ideal option.

##### S3 Versioning

As mentioned before, S3 objects are identified by their unique names. versioning is a way to track changes to object keys even as the objects change. without versioning, uploading an object with the same key overwrites the existing one. with versioning, each key maintains the previous versions of the objects (by attaching a unique id) and this protects objects from deletions and changes. all versioned objects still take up space in the bucket, and therefore incur costs.

Versioning is applied to the bucket on the whole, and all objects are effected by it. Buckets can be in one of three states:

1. UnVersioned (default) - no object in the bucket have a version.
2. Versioning enabled - versioning is applied to the bucket, and all new object will have a version.
3. Versioning suspended - versioning was enabled and then stopped. new object will not have a version, but existing objects will keep their versions.

##### S3 Life Cycle Management

Similar to the <cloud>S3 Intelligent-Tiering</cloud> storage class, lifecycle can also be managed by the user. it is possible to define when objects should be moved between storage classes and when they should expire and be deleted permanently. this can be done to save cost, or due to regulatory reasons.

</details>

#### Choosing the Right Storage Service

<details>
<summary>
How to choose AWS storage services based on requirements.
</summary>

(video)
Choosing the right service for the use case. practice questions. Whether to use file, block (<cloud>EBS</cloud> and <cloud>Instance Store</cloud>) and object storage, and the sub-types of each, based on the requirements.

1. <cloud>Lambda</cloud> means that there is no block storage attached, large files implies <cloud>S3</cloud>. regulations also implies S3 with long storage such as <cloud>glacier</cloud>.
2. Service layer on <cloud>EC2</cloud> instance, frequent access and updates. needs to be durable and fast. this implies <cloud>Elastic Block Storage</cloud>. instance store is ruled out because it's ephemeral and we need something persistent.
3. Writing temporary data to disk and performing calculations on it with speed as the important factor. this implies <cloud>Instance Store</cloud>, the data is temporary so we don't care about persistent. instance store also saves on costs.
4. A shared platform with customizations - ~~this is probably <cloud>EBS snapshots</cloud> or <cloud>AMI</cloud>. we have a base image that we want to re-use again and again.~~ the actual answer is <cloud>Elastic File System</cloud>. we want to mount this to multiple instances of EC2 when they are booted.

> <cloud>Amazon EC2 instance store</cloud>\
> Instance store is ephemeral block storage. This is pre-configured storage that exists on the same physical server that hosts the EC2 instance and cannot be detached from Amazon EC2. _You can think of it as a built-in drive for your EC2 instance_.\
> Instance store is generally well suited for temporary storage of information that is constantly changing, such as buffers, caches, and scratch data. It is not meant for data that is persistent or long lasting. If you need persistent long-term block storage that can be detached from Amazon EC2 and provide you more management flexibility, such as increasing volume size or creating snapshots, you should use Amazon EBS.
>
> <cloud>Amazon EBS</cloud>\
> Amazon EBS is meant for data that changes frequently and must persist through instance stops, terminations, or hardware failures. Amazon EBS has two types of volumes: _SSD-backed_ volumes and _HDD-backed_ volumes.\
> The performance of SSD-backed volumes depends on the IOPs and is ideal for transactional workloads, such as databases and boot volumes. \
> The performance of HDD-backed volumes depends on megabytes per second (MBps) and is ideal for throughput-intensive workloads, such as big data, data warehouses, log processing, and sequential data I/O.\
> Here are a few important features of Amazon EBS that you need to know when comparing it to other services.
>
> - It is block storage.
> - You pay for what you provision (you have to provision storage in advance).
> - EBS volumes are replicated across multiple servers in a single Availability Zone.
> - Most EBS volumes can only be attached to a single EC2 instance at a time.
>
> <cloud>Amazon S3</cloud>
> If your data doesn’t change often, Amazon S3 might be a cost-effective and scalable storage solution for you. Amazon S3 is ideal for storing static web content and media, backups and archiving, and data for analytics. It can also host entire static websites with custom domain names.\
> Here are a few important features of Amazon S3 to know about when comparing it to other services:
>
> - It is object storage.
> - You pay for what you use (you don’t have to provision storage in advance).
> - Amazon S3 replicates your objects across multiple Availability Zones in a Region.
> - Amazon S3 is not storage attached to compute
>
> <cloud>Amazon EFS</cloud>\
> Amazon EFS provides highly optimized file storage for a broad range of workloads and applications. It is the only cloud-native shared file system with fully automatic lifecycle management. Amazon EFS file systems can automatically scale from gigabytes to petabytes of data without needing to provision storage. Tens, hundreds, or even thousands of compute instances can access an Amazon EFS file system at the same time. \
> Amazon EFS Standard storage classes are ideal for workloads that require the highest levels of durability and availability. EFS One Zone storage classes are ideal for workloads such as development, build, and staging environments.\
> Here are a few important features of Amazon EFS to know about when comparing it to other services:
>
> - It is file storage.
> - Amazon EFS is elastic, and automatically scales up or down as you add or remove files. And you pay only for what you use.
> - Amazon EFS is highly available and designed to be highly durable. All files and directories are redundantly stored within and across multiple Availability Zones.
> - Amazon EFS offers native lifecycle management of your files and a range of storage classes to choose from.
>
> <cloud>Amazon FSx</cloud>\
> Amazon FSx provides native compatibility with third-party file systems. You can choose from NetApp ONTAP, OpenZFS, Windows File Server, and Lustre. With Amazon FSx, you don't need to worry about managing file servers and storage. This is because Amazon FSx automates time consuming administration task such as hardware provisioning, software configuration, patching, and backups. This frees you up to focus on your applications, end users, and business.\
> Amazon FSx file systems offer feature sets, performance profiles, and data management capabilities that support a wide variety of use cases and workloads. Examples include machine learning, analytics, high performance computing (HPC) applications, and media and entertainment.
>
> | File System                        | Description                                                                                 |
> | ---------------------------------- | ------------------------------------------------------------------------------------------- |
> | Amazon FSx for NETAPP ONTAP        | Fully managed shared storage built on the NetApp popular ONTAP file system                  |
> | Amazon FSx for OpenZFS             | Fully managed shared storage built on the popular OpenZFS file system                       |
> | Amazon FSx for Windows File Server | Fully managed shared storage built on Windows Server                                        |
> | Amazon FSx for Lustre              | Fully managed shared storage built on the world's most popular high-performance file system |

</details>

#### Demonstration: Creating an Amazon S3 Bucket

<details>
<summary>
Demo of Creating an S3 bucket and using it in the EC2 machine.
</summary>

(video)

In the Portal, we go to the <cloud>S3</cloud> service and <kbd>create bucket</kbd>, we provide a name and a region. we use the rest of the defaults as they are. once the bucket is created, we test that we can upload an object into it with <kbd>upload</kbd> and choose some files. we don't make this publicly accessible, rather, we want to change the <cloud>bucket policy</cloud> from the permissions tab, we edit the policy like an <cloud>IAM</cloud> policy - we define the statement with effect (allow or deny) with an action (in this case, all S3 actions) and the resource (the bucket arn). we also provide a principal to control who can use this policy.\
Next we want to modify our <cloud>EC2</cloud> application to use the bucket, we choose our stopped instance and under <kbd>actions</kbd>, we choose the <kbd>launch more like this</kbd> option to use the same settings. under the "advanced details", we can assign the correct role, and under the user data script, we can add the bucket name which we created.

</details>

#### Module 4 Knowledge Check

<details>
<summary>
Recap Questions.
</summary>

> - Q: Which of the following is a typical use case for Amazon S3?
> - A: Object storage for media hosting
> - Q: A company that works with customers around the globe in multiple Regions hosts a static website in an Amazon S3 bucket. The company has decided that they want to reduce latency and increase data transfer speed by storing cache. Which solution should they choose to make their content more accessible?
> - A: Configure <cloud>Amazon CloudFront</cloud> to deliver the content in the S3 bucket.
> - Q: Which of the following storage services is recommended if a customer needs a storage layer for a high-transaction relational database on an Amazon EC2 instance?
> - A: <cloud>Amazon Elastic Block Store (Amazon EBS)</cloud>

</details>

</details>

### MODULE 5: Databases on AWS

<details>
<summary>
Databases on AWS - RDS, DynamoDB and others.
</summary>

> A high-performing database is crucial to any organization. Databases support the internal operations of companies and store interactions with customers and suppliers.

#### Introduction to Databases on AWS

<details>
<summary>
Relation Database introduction, managed and un-managed options.
</summary>

(video)

we want to store our demo app data in a database, we chose to use relational database for that. RDBMS stands for relational database management system. we can install databases on EC2 instances, which is great for migration from on-premises. but it still requires managing the machine, installing and updating the software. a different option is to have AWS take care of this and use a specialized database service.

There are several kinds of databases. the first and most basic are relation databases, they store data in tables, rows and columns. data in one table can be "linked" to data in another table. these link form the relationship between the data. relational databases require strict schemas of data detailing the tables, the columns types and hoe they relate to one another, changing the schema after the fact is difficult.

- MySQL
- PostgresSQL
- Oracle
- Microsoft SQL Server
- <cloud>Amazon Aurora</cloud>

Communication with Relational databases is done via **SQL** (structure query language), which allows for querying data and "jointing" multiple tables.\
Using relational databases provides benefits like the ability for complex and detailed SQL queries, reduced memory needs since there is no redundant data store. it is also highly popular and supported, and allows for transactional integrity. common use cases for relational databases are when using application that have a fixed schema hat doesn't change often, and when persistent storage is required and there is a need for high data integrity.

Databases on the cloud can be managed or un-managed, this choice shifts the responsibilities of managing the database between the cloud vendor and the user. **un-managed** databases run on <cloud>EC2</cloud> machines in the cloud just like any other software, and the user is responsible to manage the instance, install and update the database software, and apply any patches to it. this is all before doing the actual work of managing the data and creating schemas and queries. **managed databases** push some of the work onto the cloud vendor, this includes the machine OS, configuring it for high-availability, scaling, backups, and other tasks. This leaves the user with the smaller task of managing the data itself.

</details>

#### Amazon RDS

<details>
<summary>
AWS Relational Databases.
</summary>

> With Amazon Relational Database Service (Amazon RDS), you can focus on tasks that differentiate your application instead of infrastructure-related tasks, like provisioning, patching, scaling, and restoring.

(video)

<cloud>Amazon RDS</cloud> is a service that allows us to create and manage the databases. for our example, we will use the most simple and default options. so we click <kbd>Create Database</kbd>, and then <kbd>Easy create</kbd> to accept the default behavior and the "best practices" it supports. we next choose the database engine, AWS supports some common engines, but also has <cloud>Amazon Aurora</cloud>.\
Amazon Aurora is a RDS engine that was designed to take advantage of the cloud capabilities, it can be up to five times more efficient than other engines in some cases. however, in our demo APP we can use the simpler "MySQL" engine. we next choose the instance size and type, and give the instance name and choose the database administrator user name and password. this creation can take a few minutes.\
A RDS instance is placed inside a subnet in a VPC, which means it goes into an Availability Zone. the best practice is to always deploy applications across Availability Zones for high availability. this can be achieved with aws <cloud>Multi-AZ deployment</cloud>, which sets up a second instance with data replication. it also controls the fail-over transfer between the instances when the primary one fails. the endpoint doesn't change, so there is no need for code changes.

Amazon RDS supports popular RDBMS engines. commercial and open source, and has it's own unique cloud native engine.

- Commercial engines:
  - Oracle
  - SQL Server
- Open Source:
  - MySQL
  - PostgresSQL
  - MariaDB
- Cloud Native
  - Aurora

managed databases still run on virtual machine instances, so the user is required to choose a type of machine and the size of the instance, for most engines, storage is handled by <cloud>AWS EBS (elastic block storage)</cloud>, so that's another thing the user needs to choose, but not manage directly. The storage can be "general purpose SSD" (for most cases), "Provision IOPS SSD" (i.o intensive workloads) and "magnetic" (backwards compatibility, not recommended to use anymore).\
A RDS instance resides in a VPC subnet, so it's limited to a single Availability Zone. the subnet should be private (not have a route to internet gateway), and only be reachable from the application backend. in addition, the subnet should be restricted with <cloud>Access Control Lists (ACL)</cloud> and <cloud>Security Groups</cloud>.

##### Backups

Databases should be backed up on a regular basis, AWS provides automated backups for the DB instance, and has retention period of up to 35 days. this means that at any time, it's possible to perform a "point in time recovery" and re-create a RDS instance at the same state it was when the backup was done. There are also manual backups, which are user-initiated and aren't deleted automatically. this can be because of regulatory requirements.

##### High Availability

Databases should be deployed with high availability in mind, which means redundancy across Availability Zones. <cloud>Multi-AZ</cloud>
deployment creates a primary and secondary RDS instances with replication between them. only the primary database answers queries, and the secondary database is on "standby" mode until something happens to the primary one and it's then "promoted". data sent to the database goes through a DNS name, so when a failover happens, the name in the DNS is updated to point to the new instance.

##### Security

RES security is handled on different layers, this includes IAM security to handle resource configuration, creation and deletions. Security Groups and Access control lists for network security of traffic coming in and out of the RDS. data backups and snapshots can be encrypted at rest
and communication with the instance should be done over a secure layer, such as SSL or TLS.

</details>

#### Purpose-Built Databases

<details>
<summary>
Other types of databases.
</summary>

> AWS offers more than 15 purpose-built engines to support diverse data models, including relational, key-value, document, in-memory, graph, time series, wide column, and ledger databases.

(video)

The important thing is to choose the right database for the data and the application, no one database fits all needs. AWS has different databases for all sorts of purposes. for our demo app, we might not even need the capabilities of the relational database, since we don't have "joins" or aggregate actions, all we need is a lookup table. in addition to that, the RDS is charged based on the runtime, and since it's not going to see much traffic, it's not the ideal payment model (and we have virtually no usage of it outside of working hours).\
To avoid these issues, we can use <cloud>Amazon DynamoDB</cloud> instead, this is a NoSQL database which works for key-value pair and documents. it's highly scalable and charges based on usage (number of actions and the size of the data). There are also other databases, such as <cloud>Amazon DocumentDB</cloud> which is great for a CRM application, and can acts as a storage for catalogues and user profiles. if we had a social network, we could go with <cloud>Amazon Neptune</cloud>, a graph database for social networks, recommendations (also good for fraud detection). <cloud>Amazon QLDB (Quantum Ledger DataBase)</cloud> is an immutable ledger database (for banking, shipping, finance) which doesn't allow entires to be removed or modified.\
Every Database fits for a different use-case, and using the purpose build options removes some of the complexity of learning a new database.

##### AWS DynamoDB

A fully managed NoSQL database with fast, consistence performance. it works great for high scale applications and serverless workloads.

##### AWS ElasticCache

A fully managed in-memory cache with two engines: <cloud>Redis</cloud> and <cloud>Memcached</cloud>. it takes care of instance failover, backups and software upgrades.\
<cloud>Amazon MemoryDB for Redis</cloud> is a Redis-compatible AWS service with ultra-fast performance.

##### AWS DocumentDB

This document Database is MongoDB compatible, it stores rich documents and allows for queries and aggregation. it's suited for content management, profile management, and web and mobile applications.

##### AWS KeySpaces

This service is an <cloud>Apache Cassandra</cloud> compatible database for high-scale and high performance applications. it uses the same query language and drivers.

##### AWS Neptune

A fully managed Graph database for data with high connectivity and relations.

##### AWS TimeStream

Serverless, scalable database for time-series based data, good for Internet of Things data with large number of events.

##### AWS Quantum Ledger Database

Cryptographically verifiable immutable ledger database.

</details>

#### Amazon DynamoDB

<details>
<summary>
Digging Deeper into DynamoDB.
</summary>

> With Amazon DynamoDB, you have a fully managed service that handles the operations work.

(video)

Amazon DynamoDB is a serverless database (no need to manage the instances running it). the tables in dynamoDB don't need to have relations to one another, they are "stand-alone" tables, each containing items (like rows). AWS handles storing the data and scaling it. the data is stored redundantly across Availability Zones and mirrors the data across drives. it is also highly responsive and very scalable. it doesn't require a rigid schema and doesn't have constraints. elements in the table aren't required to have the same structure, and each can have different attributes. queries on dynamoDB tables are simpler and focus on a single table. this makes the response time quicker and faster to scale.\
For our demo application, we will replace the RDS database for a dynamoDB table. in our web portal, we click <kbd>Create Table</kbd>, give the table a name and choose a partition key and an optional sorting key. we can <kbd>Explore Table Items</kbd> and see the items in the table.

DynamoDB is a fully managed, scalable and high performant database. AWS handles the scaling, replication and configuration of the data storage. data is separated inside tables, and is stored by AWS in SSD devices with high availability and data durability built-in.\
A table stores Items (similar to rows), where each item has attributes (like as fields or columns). one attribute is the primary Key that manages how data is stored. additional indexes can be created to increase flexibility when querying.

##### DynamoDB Use Cases

> You might want to consider using DynamoDB in the following circumstances:
>
> - You are experiencing scalability problems with other traditional database systems.
> - You are actively engaged in developing an application or service.
> - You are working with an OLTP workload.
> - You care deploying a mission-critical application that must be highly available at all times without manual intervention.
> - You require a high level of data durability, regardless of your backup-and-restore strategy.

concrete cases can be for software applications that support user-content metadata and cache which requires high concurrency and a large number of concurrent requests (millions per second). one example can be storing media metadata, like analysis and interactive content. this is beneficial due to the lower latency and multi-region replication.\
Other uses are storing user data for a game platform, and for online shopping with inventory tracking when there is high traffic.

##### DynamoDB Security

DynamoDb handles security through IAM, encryption at test using <cloud>AWS Key Management Service</cloud>. network is protected via the AWS global network security procedures (not inside a VPC), and is backed up across multiple facilities in the region. Requests to DynamoDB with a managed key can be monitored by <cloud>AWS Cloud Trail</cloud> and can then be exported to an S3 bucket. IAM roles use temporary access keys and should be configured to use only the least privileged permissions that they need.

</details>

#### Choosing the Right Database Service

<details>
<summary>
How to choose AWS Database services based on requirements.
</summary>

> Choose the database service that is the best fit for the job to help you optimize scale, performance, and costs when designing applications.

| AWS Service(s)                                                 | Database Type | Use Cases                                                                                          |
| -------------------------------------------------------------- | ------------- | -------------------------------------------------------------------------------------------------- |
| Amazon RDS, Aurora, Amazon Redshift                            | Relational    | Traditional applications, ERP, CRM, e-commerce                                                     |
| DynamoDB                                                       | Key-value     | High-traffic web applications, e-commerce systems, gaming applications                             |
| Amazon ElastiCache for Memcached, Amazon ElastiCache for Redis | In-memory     | Caching, session management, gaming leaderBoards, geoSpatial applications                          |
| Amazon DocumentDB                                              | Document      | Content management, catalogs, user profiles                                                        |
| Amazon KeySpaces                                               | Wide column   | High-scale industrial applications for equipment maintenance, fleet management, route optimization |
| Neptune                                                        | Graph         | Fraud detection, social networking, recommendation engines                                         |
| TimeStream                                                     | Time series   | IoT applications, Development Operations (DevOps), industrial telemetry                            |
| Amazon QLDB                                                    | Ledger        | Systems of record, supply chain, registrations, banking transactions                               |

Most production grade application combine different databases, rather than use one database for all needs. This is part of the "micro-services" architecture model.

</details>

#### Demonstration: Implementing and Managing Amazon DynamoDB

<details>
<summary>
Demo of creating A dynamoDB table.
</summary>

(video)

in the <cloud>EC2</cloud> services, we select one machine, <kbd>Actions</kbd>, and launch a clone of the instance by selection <kbd>More like this</kbd> and adjusting the settings (we want a public IP). when the instance is ready, we can copy the public ip and navigate into it to see that it's ready.\
In the <cloud>DynamoDB</cloud> service, we can <kbd>Create Table</kbd>, give it the names "Employees" and a partition key "id", we ignore the sort key option for now. and then click <kbd>Create Table</kbd> to create it. now in our application, we can add employees and see that it's been updated in the directory and in the S3 bucket and the dynamoDB table.

</details>

#### Module 5 Knowledge Check

<details>
<summary>
Recap
</summary>

> - Q: With Amazon RDS, you can scale components of the service. What does this mean?
> - A: You can increase or decrease specific database configurations independently. (Scaling components of the service means you can alter memory, processor size, allocated storage, or IOPS individually without modifying other configurations you set in your database.)
> - Q: An organization needs a fully managed database service to build an application that requires high concurrency and connections for millions of users and millions of requests per second. Which AWS database service should the organization use?
> - A: <cloud>Dynamodb</cloud>

</details>

</details>

### MODULE 6: Monitoring, Load Balancing, and Scaling

<details>
<summary>
//TODO: add Summary
</summary>

Monitoring
Amazon CloudWatch
Solution Optimization
Traffic Routing with Elastic Load Balancing
Amazon EC2 Auto Scaling
Demonstration: Making the Employee Directory Application Highly Available
Employee Directory Application Redesign
Module 6 Knowledge Check

</details>
