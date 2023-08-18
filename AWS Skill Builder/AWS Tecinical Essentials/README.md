<!--
ignore these words in spell check for this file
// cSpell:ignore Seph
 -->

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

hosting the application with Amazon EC2, using the default VPC and other defaults. under the <kbd>EC2</kbd> service, we <kbd>Launch an instance</kbd> - a single virtual machine. we give it a name and use the linux AMI, the free tier, and under the network settings we click <kbd>edit</kbd> and choose the default VPC without subnet preferences. we need to <kbd>Add Security group</kbd> to allow HTTP and HTTPS traffic to reach the instance, and under the advanced details, we choose the role we created earlier as the instance profile. under the user date, we add the script that downloads the source code for the app and launches it with flask. after clicking <kbd>Launch instance</kbd>, we can wait a few minutes, and then use the public IP address to navigate into it in the browser.

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

we focus on the network VPC and it's components. when we created an EC2 machine, we put it into a vpc. we used a default vpc that AWS creates for us. the default VPC has inbound access from the internet, so it's dangerous to use it all the time. some services don't require VPCs, but it's still important to learn about it.

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
All the resources in the vpc are isolated from others, so to allow Internet connectivity, we need an **Internet Gateway** (attached to the vpc). <kbd>Create Internet Gateway</kbd>, then <kbd>Attach to VPC</kbd>. if we want to limit internet access and only allow access from a specific location, we can use **Virtual Private Gateway** (VPN). ideally we would want to have duplicated resources in a different Availability Zone for high availability.

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

##### Gateways:

Internet gateways act like a modem - they connect the VPC to the external internet. Virtual private gateways are VPN, they connect only to a matching customer gateway (a device or software) and provide encrypted communication between the two sides. There s also **AWS Direct Connect**, which is a physical connection between the on-premises data center and AWS metacenter, which has high speed and doesn't travel through the public internet at all.

</details>

#### Amazon VPC Routing

<details>
<summary>
Routing External Traffic to the subnets.
</summary>

> A route table contains a set of rules, called routes, that determine where network traffic from your subnet or gateway is directed.

(video)

once the traffic got to the internet gateway, we need to route it to the correct subnet and resources. this is done via the **Route Table**, it contains rules (routes) that move traffic around. a route table can at the vpc level or the subnet level. aws creates a default main route-table, which only allows connections between the local vpcs. in the route table <kbd>Routes</kbd> table, we can see the routing to the subnets (10.1.0.0/16). the connection from the internet gateway to the subnets determines if they are public or private, if a route exists, then it's a public subnet, otherwise it's private.\
in our app we create custom route tables, we click <kbd>Create Route Table</kbd>, give it a name and a vpc, and the <kbd>edit routes</kbd> and  <kbd>add routes</kbd>, and select the target type from the drop down.

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

any new VPC is isolated from internet access, because it doesn't have an internet gateway associated with the route table. but once we connect the subnet to the external internet, we need to add security. the two options available for us are **Access Control Lists (ACL)** and **Security Groups**. \
An ACL is like a firewall on the subnet level. the control inbound and outbound traffic. we can specify which access types are allowed (HTTP, HTTPS, SSH, etc..) by creating inbound rules and specifying the source. in ACL, we also need to open the corresponding outbound rules. this is because ACL are considered *stateless*.\
Security groups are created at the EC2 level, and they are mandatory (every EC2 has them). the default Security group behavior is to block inbound requests and allow outbound. if we want to allow requests from the internet to reach us, we need to open inbound rules(such as port 80 and 443 for http and https). security groups are *stateful*, so if there is an inbound rule that allows traffic, outbound traffic to the same destination will be allowed as well.
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
//TODO: add Summary
</summary>

Storage Types
File Storage with Amazon EFS and Amazon FSx
Block Storage with Amazon EC2 Instance Store and Amazon EBS
Object Storage with Amazon S3
Choosing the Right Storage Service
Demonstration: Creating an Amazon S3 Bucket
Module 4 Knowledge Check

</details>

### MODULE 5: Databases on AWS

<details>
<summary>
//TODO: add Summary
</summary>

Introduction to Databases on AWS
Amazon RDS
Purpose-Built Databases
Amazon DynamoDB
Choosing the Right Database Service
Demonstration: Implementing and Managing Amazon DynamoDB
Module 5 Knowledge Check

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
