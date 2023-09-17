<!--
ignore these words in spell check for this file
// cSpell:ignore elbv2 Neumann cgroups pictShare Kubelet eksctl
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Developer Learning Plan

[Developer Learning Plan](https://explore.skillbuilder.aws/learn/lp/84/developer-learning-plan)

> A Learning Plan pulls together training content for a particular role or solution, and organizes those assets from foundational to advanced.   Use Learning Plans as a starting point to discover training that matters to you.\
> This Learning Plan is designed to help Developers who want to learn how to develop modern applications on AWS. The digital training included in this Learning Plan will expose you to developing with serverless and container technologies, as well as the foundation of DevOps on AWS. This Learning Plan can also help prepare you for the AWS Certified Developer - Associate certification exam.

## Introduction Elastic BeanStalk

<details>
<summary>
Deploy Servers, Databases and Load Balancers in an integrated way.
</summary>

> AWS <cloud>Elastic Beanstalk</cloud> provides you with a platform enabling you to quickly deploy your applications in the cloud. This course will briefly discuss the different components of the AWS Elastic Beanstalk solution, and perform a demonstration of the service.

as developers, we want to get our application to the cloud quickly. <cloud>Elastic Beanstalk</cloud> belongs to the *Platform as a service* family of AWS features. it reduces management complexity and allows re-using existing code (some languages), it also allows for some control over the running infrastructure, such as the instance type, database and auto scaling.

supported platforms:

- Package Builder
- Single and Multi Container, pre-configured Docker
- Go
- Java SE, Java with Tomcat
- .Net on windows Server with IIS
- Node.js
- PHP
- Python
- Ruby

<cloud>EBS</cloud> allows for versioning and updating code and reusing the same deployment again and again.

(DEMO)

we can have a web server environment or use a worker environment, a web server can have a domain. we select the platform (such as python) and the application code in zip or rar format. we can modify the instance type, security group, notifications, monitoring, load balance and monitoring, we can also set scaling for high availability.

</details>

## Getting Started With .Net on AWS

<details>
<summary>
AWS Basics with .NET tools
</summary>

> In this course, you will learn the basics of deploying, managing, and securing .NET applications with Amazon Web Services (AWS). You will learn about AWS services and tools specifically designed for .NET applications. Finally, the course will walk you through a hands-on example of deploying a .NET application to the AWS Cloud.

### Introduction

<details>
<summary>
Dot.NET versions
</summary>

(video)

AWS services for .Net application, SDKs. support for framework and Core (open source, cross platform), choosing between .NET Framework and .NET Core. if possible, choose the CORE option, as it's the long-term support version and it can be cross platform (not just windows).

</details>

### AWS Services

<details>
<summary>
Going Over AWS Services: Compute, Storage, Identity and Monitoring.
</summary>

(video)

AWS Compute services: instances, containers and serverless. if we want to run .NET code on instances, we can use use <cloud>EC2</cloud> and deploy with <cloud>Elastic Beanstalk</cloud>.if we wish to reduce infrastructure management we can choose to use container services such as <cloud>Elastic Container Service</cloud> and <cloud>Elastic Kubernetes Service</cloud>. we can go one step further and use serverless services such as <cloud>Lambda</cloud> for event driven computing and <cloud>AWS Fargate</cloud> to run containerized workloads.\
next we have Storage and Database services, <cloud>S3</cloud> is an object storage service, <cloud>AWS RDS</cloud> is a fully managed relational databases services with several engines to choose from: <cloud>Amazon Aurora</cloud>, PostgresSQL, MySQl, MariaDB, Oracle Database, Microsoft SQL Server. for non-relational options, there is <cloud>DynamoDB</cloud> as a key-value database with high speed and low-latency. there are additional purpose built specialized databases as well.\
There are Identity Services (<cloud>IAM - Identity Access Management</cloud>) to manage permissions, control access and create user and roles to granularly limit who can access which resource and what actions are possible. there is also the option of <cloud>AWS Directory Service</cloud> to integrate with on-premises active-directory services. <cloud>Amazon Cognito</cloud> allows sign-in through other identity providers and using SAML to log-in into AWS. it also supports smart auditing of user behavior.\
There are also Monitoring and Auditing services, <cloud>CloudWatch</cloud> for monitoring performance and operational information such as logs and <cloud>CloudTrail</cloud> to record API behavior and actions to know which operations are used and by who.

moving from EC2 to other options (containers, serverless) decreases the amount of infrastructure overhead and offloads it to AWS.
</details>

### Developer Tools

<details>
<summary>
SDKs, IDE extensions, CLI libraries and other development tools.
</summary>

(video)

tools for developers, such as AWS SDK for .NET, Visual Studio, Visual Studio Code and JetBrains IDE extensions. Powershell tools and CLI scripts. extensions to the .NET Core CLI tool for easier deployment of AWS services.

- AWS Toolkit for <cloud>Azure</cloud> devops allows deploying AWS resources from Azure dev and release pipelines
- AWS Cloud Development Kit for creating infrastructure as code
  
</details>

### Practice Activity

<details>
<summary>
Deploy a sample .NET web application to the AWS Cloud using the AWS Toolkit for Visual Studio.
</summary>

> Tasks
>
> 1. Set up your AWS and Microsoft Visual Studio environment.
> 2. Create an Amazon RDS database instance.
> 3. Deploy a sample .NET application to the AWS Cloud.
> 4. Clean up your resources.

we start by creating an IAM user for API access.permissions are managed from IAM, so we want to create a user, give it appropriate permissions, create and use credentials for programmatic access. using <cloud>IAM policies</cloud>.\
Next we set up the AWS profile in microsoft visual studio. we install the AWS toolkit, and create a credential profile, there is an option called <kbd>AWS Explorer</kbd> that has wizards to help us.\
This Demo uses a sample application for URL redirection.

our next step is creating a database. in the <cloud>AWS Explorer</cloud> we can select the profile and the region. and then select <kbd>Launch Instance</kbd> on the <cloud>RDS</cloud> to start the wizards, we can use microsoft SQL server Express Edition for this application. we choose the instance class and select the database admin user and password. we need to select the <cloud>VPC</cloud> the instance resides in, the subnets, Availability Zone and security group. we can choose the port and additional settings, and then we can manage backup options, and eventually launch the resource.\
Once the instance is launched, we can select it and choose <kbd>Create SQL Server Database</kbd> and copy the address. next, run the `Update-Database` command in the nuget cli to create the database schema, and now we can run the application. we could also have used <cloud>Amazon Aurora</cloud> as a high-performance cloud optimized database, or <cloud>DynamoDB</cloud> as a NoSQL low latency serverless option.

so far, the application run on the local computer with a cloud based database, so we want to have it run on the cloud instead.\
<cloud>Elastic Beanstalk</cloud> is an easy to do so. in the visual studio extension, we can choose <kbd>Publish to AWS Elastic Beanstalk</kbd> to open a wizard, we choose to create a new application environment, give it a unique URL, and we choose the instance to run the application, just like before we need the VPC, subnet and security group. we also might need a key-pair to log-in into the machine, but we choose not to use it. we want our application to have multiple instances for high availability, so it needs an <cloud>Application Load Balancer</cloud>, and select the RDS we previously created. permissions are given by IAM roles, so we can use the default roles. once we're finished, we can click <kbd>Deploy</kbd> and wait for it to be done, and copy the URL and navigate to it in the browser.

After everything is done, we can delete our resources through the aws explorer.

</details>

### Assessment

<details>
<summary>
Recap Questions
</summary>

- Q: Why is .NET Core the future of the .NET platform?
- A: .NET Core is the recommended path for moving to .NET 5.
- Q: What is an advantage of hosting .NET applications on AWS?
- A:Provides fully featured services with deep functionalities
- Q: An application developer wants to host a .NET application in a containerized environment but does not want to manage the infrastructure. Which AWS compute service fits their needs?
- A: AWS Fargate
- Q: Which AWS service automatically handles capacity provisioning, load balancing, automatic scaling, and application health monitoring of compute instances running your .NET application?
- A: AWS Elastic Beanstalk
- Q: Which database option is most suitable for a developer who needs a key-value NoSQL database for their .NET application?
- A: Amazon DynamoDB
- Q: Which database engine is supported by Amazon Relational Database Service (Amazon RDS)?
- A: Microsoft SQL Server
- Q: Which service is used to enable fine-grained access control for users accessing your .NET applications and AWS resources?
- A: AWS Identity and Access Management (IAM)
- Q: An application developer recently deployed their .NET application to the AWS Cloud. They now want to provide users with a sign-in using social identity providers such as Google, Amazon, or Facebook. Which AWS service should they use?
- A: Amazon Cognito
- Q: Which of these is an AWS extension for a .NET developer tool?
- A: AWS Toolkit for Visual Studio
- Q: Which developer tool enables .NET developers to provision their cloud infrastructure through .NET code?
- A: AWS Cloud Development Kit (AWS CDK)

</details>

### Conclusion

<details>
<summary>
Benefits of choosing AWS
</summary>

(video)

reason to choose AWS:

1. functionality - wide selection of services that can fit general and specific use-cases.
2. community of customers and partners - millions of active customers, many developers and support options.
3. security -  satisfies security requirement for military, finance and government standards.
4. pace of innovation - continually adding services and technology, giving more options and making development easier.

</details>

</details>

## Containers

<!-- <details> -->
<summary>
Several Courses focusing on containers.
</summary>

### Introduction To Containers

<details>
<summary>
Basic introduction for containers and why they are used.
</summary>

> This is an introductory course designed for participants with little-to-no previous knowledge of containers. It will teach you the history and concepts behind containerization, provide an introduction to specific technologies used within the container ecosystem and discuss the importance of containers in microservice architectures.

A physical container is a standardized form for delivering goods, it has a set measurements, all containers can be stacked and shipped at the same way with the same machinery, and each transporting vehicle has a known capacity for the number and weight of containers it can carry.\
A digital container is a standardized unit of software the is meant to run on any hardware running a container platform. it is a method of abstraction that uses virtualization, a container should include everything it needs to run the application. a single server can have multiple containers, they can be connected or stand-alone.

Containers differ from other forms of virtualization, when using a bare-metal machine (such as a local computer or a server), then all programs share the hardware and the operating system, and they have the same libraries, so if they have conflicting library requirements, things can go wrong.\
the next level is virtual machine, where the virtualization platform creates a "guest OS" with the required libraries, so each unit of software creates a new OS, which is heavy and wastes resources.\
Container images share the host operating system kernel, they are composed of layers for reusability, and they are light-weight, faster to create and can share libraries and still have unique versions. they are also very portable, since everything is packaged into the image.

there are earlier implementations of the idea of virtualization, but it is now very popular, partly due to the rise of the docker containerization platform. docker has many benefits:

> - *Portable* runtime application environment
> - Package application and dependencies in a *single, immutable artifact*
> - Ability to run different application versions with different dependencies *simultaneously*
> - *Faster* development and deployment cycles
> - Better *Resource utilization and efficiency*

Images are read-only templates to create a container, so an container is an instance of an image. Images are based on one another, starting from the OS Image and building on it.\
Images can be created by <cloud>DockerFiles</cloud>, each line in the file adds a layer to the image. an instruction can run a script, copy files, and expose ports.\
Each container has a thin read-write layer that is unique to it, but can share other layers with other containers running the same image.

Containers go together with micro-service architecture. this approach is meant to increase development cycles. in a traditional architecture, all parts of the application run in the same server, so a spike in demand for one component requires scaling everything up. it also creates tightly coupled components which are hard to change.\
In contrast, using micro-services can split components into their own re-usable applications, that are simple to update, deploy and scale, and they can be more portable and run on different platforms.

Characteristics of microservices

- Decentralized, evolutionary design
- Smart endpoint, dumb pipes
- Independent products, not projects
- Designed for failure
- Disposability
- Development and production parity

</details>

### Introduction To AWS Fargate

<details>
<summary>
Fully Managed Infrastructure and Scaling for Containers
</summary>

> This is an introductory course to AWS Fargate, a new AWS service for deploying and managing containers. In this course, we cover how AWS Fargate makes it easier for you to run applications using containers and we walk through an example architecture of AWS Fargate and Amazon ECS so you can better understand how the service works.

<cloud>Fargate</cloud> is an AWS service for containers. it removes the need to manage <cloud>EC2</cloud> instances in terms of provisioning and scaling. it works with <cloud>ECS</cloud> and <cloud>EKS</cloud>, has pay-per-usage payment model. it helps developers focus on the application they run, and not the instances that run them.\
AWS takes care of creating the EC2 instances and managing them to have the correct agents (ECS agent, docker agent) and the correct AMI. they are also responsible for cluster management and engine placement. The Developer is in charge of creating the task and the ecs orchestration.

Fargate is easier to use, and can be used through the normal tools and APIs. there are only a few small differences. Fargate task run in VPC, and works with load balancers, but can be still monitored like <cloud>EC2</cloud> machines. there is no SSH access to the tasks.

Fargate can fit any container use case, such as:

> - Long running services
> - Highly available workloads
> - Monolithic app portability
> - Batch Jobs and microservices

there are cases where EC2 launch types are better, such as spot and reserved instances payment modes.

- ECS - aws native container service, which works with other AWS services.
- EKS - aws Kubernetes clusters, using the open source platform.
- ECR - container image registry to store images

</details>

### Deep Dive On Aws Fargate - Building Serverless Containers At Scale

<details>
<summary>
More Detailed view of Fargate
</summary>

> Containers allow you to craft sophisticated cloud-native applications, but how do you manage scale? In this course you will learn how to better launch and manage your large-scale containerized workloads with AWS Fargate.

#### Main Motivation For Fargate

developers love using containers, but there are more layers to using container than just creating them. there are many additional tasks when using containers in a real environment, such as orchestration, monitoring and resource managements. This was the original reason for the creation of the <cloud>ECS</cloud>. but even with those tools, there is still a need to manage more than just the containers. the complexity increases the more EC2 instances there are.

AWS Fargate was developed to reduce the overhead for using containers:

> - Managed by AWS - no EC2 instances to provision, scale or manage.
> - Elastic - scale up & down seamlessly, pay only for what you use.
> - Integrated - with the AWS ecosystem: <cloud>VPC</cloud> networking, <cloud>Elastic Load Balancing</cloud>, <cloud>IAM permissions</cloud>, <cloud>Amazon CloudWatch</cloud> and more.

#### AWS Fargate Components

A Task definition defines the application in term of image, resource requirements, etc... a single task is a running instantiation of a task definition, and it it set to use Fargate as a launch type. A service maintains and manages the running copies of the task, and is integrated with the load balancer to check and replace unhealthy tasks as needed. They reside inside a cluster, which is the boundary for infrastructure and IAM permissions.

a task definition is immutable, and changes to a definition create new versions. a task can have up to ten container definitions, which will all run on the same host. the definition has task level resources (CPU and memory), and can optionally define different slices to the containers. the resources determine the costs.

tasks are run inside VPC (so inside subnets), each has an <cloud>Elastic Network interface</cloud>, which means it has private ip address in that network. we can always assign a public ip address to the <cloud>ENI</cloud> if needed.

Storage for task is achieved by <cloud>EBS</cloud> and use ephemeral storage: writable layer and volume storage. all containers inside the task share a 10GB writable layers, which includes the image layers. the writes are not visible across containers. sharing data is done via volume, each task has 4GB volume storage that can be mapped in volume mounts in the task definition. this storage is not available after the task stops, so it's no persistent.

Since this is an AWS based service, it uses IAM permissions.

- cluster permissions control who can launch and describe tasks in the cluster.
- application permissions (IAM roles as EC2 machines) to allow the application container to access AWS resources securely.
- Housekeeping Permissions allow aws to manage the tasks properly, like pull ECR images, push logs to to <cloud>Cloud Watch</cloud>, create network interfaces and register and remove target from the load balancer. there **execution roles** to pull images and push logs, and there are **service linked roles** which are aws managed.

We can see the docker logs in <cloud>CloudWatch</cloud>, so we can have all of the logs in a central place, we can see the performance metrics of the running instances just like normal <cloud>EC2</cloud> machines. we can now see task metrics, which allows us to run custom scaling events.

Task metadata is available from inside the within the task, so monitoring tools can use it:

- Task Level
  - `169.254.170.2/v2/metadata` - metadata json for task
  - `169.254.170.2/v2/stats` - docker stats json for all containers in the the task
- Container Level
  - `169.254.170.2/v2/metadata/<container-id>` - metadata json for container
  - `169.254.170.2/v2/stats/<container-id>` - docker stats json for a container

A new option is **Managed Service Discovery**. this is done inside <cloud>Route 53</cloud> and provides a dns compatible addresses. this is managed by aws without running custom code.

Fargate allows us to run containers without having to worry about EC2 instances. it is usually a good idea to use Fargate over EC2, unless there is a special reason not to.

#### Demo

first part is Creating a task definition and running it, making sure it has auto scaling and outside connectivity.

```sh
aws ecs create-cluster --cluster-name pictShare
# edit the task definition json
nano task_definition.json
aws ecs register-task-definition --cli-input-json file://task_definition.json --query 'taskDefinition.taskDefinitionArn'
# edit service json
nano service.json
aws ecs create-service --cli-input-json file://service.json
# add scalability
aws application-autoscaling register-scalable-target --resource-id service/pictShare/pictShare --service-namespace ecs --scalable-dimension ecs:service:DesiredCount --min-capacity 1 --max-capacity 20 --role-an <ecsServiceAutoScalingRole>
# edit scaling policy
nano scale_out.json
aws application-autoscaling put-scaling-policy --cli-input-json file://scale_out.json
# checking connection
url=$(aws elbv2 describe-load-balancers --load-balancer-arns <arn> | jq 'LoadBalancers[].DNSName' | sed -e 's/"%//')
echo $url
curl -I $url
```

the next part setting up a CI-CD pipeline to update the cluster when code changes.

```sh
aws codecommit create-repository --repository-name pictShare
# this is one command!
aws codebuild create-project --name "pictShare" --description "Build project for pictShare" \ 
--source type="CODEPIPELINE" \
-- service-role=<CodeBuildExecutionRoleArn> \
--environment type="LINUX_CONTAINER", image="aws/codebuild/docker:17.09.0", computeType="BUILD_GENERAL1_SMALL" \
environmentVariables="[{name=AWS_DEFAULT_REGION, value='<region>'}, {name=AWS_ACCOUNT_ID, value=<accountId>}, {name=REPOSITORY_URL,value=<arn>}]" \
--artifacts type="CODEPIPELINE"
# edit pipeline, three stages: source, build, release, and the artifact store.
nano edit pipeline_structure.json
aws codepipeline create-pipeline --pipeline file://pipeline_structure.json
# edit event that triggers pipeline
nano event_put_rule.json
aws events put-rule --cli-input-json file://event_put_rule.json
# attach rule to target
aws events put-target --rule CodeCommitRulePictShare \
--targets Id=1,Arn=<pipeline_arn>,RoleArn=<CloudWatchCodePipelineTriggerRoleArn>
```

now we can change the application, commit the change and look at the web console and see how the stages are being executed. the change is recognized as the source, then the image is built to ECR, and lastly it's deployed to ECS where it replaces the existing tasks.

</details>

### Deep Dive on Container Security

<details>
<summary>
Some information about Linux containers and namespaces. what is shared and what is not.
</summary>

> Security should be the first concern for any project – maintaining the confidentiality, integrity and availability of your architecture. Containers present a unique middle ground between full instance management and pure services.

security in Linux containers, without focusing on any specific implementation or platform.

The risks are: Confidentiality, integrity, availability.

- Segregation(Confidentiality)
  - Container to Container
  - Process to Process
  - Container to outside
- Access
  - Who/When/Where
  - Logging
  - Start/Stop
  - Content
- Resource Usage

The Von Neumann Computer model, input, output, cpu (control unit, Registers, ALU - arithmetics logic unit), and with external memory. there are system libraries and application code which the user interacts with, and there are many other components that are part of the kernel, and the libraries interact with them through system calls.

When we have many applications, we want to run them together and have efficient CPU and memory usage. we focus on the applications and the security boundaries.\
Linux namespace are hierarchical, process can share some namespaces and have some unique namespace. in the PID(process) namespace, each process has a global pid and a local one. the first process in the namespace has zero process id, so it is the strongest process in that namespace. all namespaces still live on the same memory management.

We get into a namespace with the `clone(2)` command, and we can still `fork(2)` inside it. there are some problems with having an ssh client inside a container.

We manage cpu and memory in control groups (cgroups), they use policy based-scheduling. its sometimes possible to have CPU affinity for a namespace, but it's not always enforced. memory limitation is also difficult. Pages (dirty, used, empty) are also a global topic. Context migration are when the process is moved to another CPU, both context migrations and context switches have heavy memory costs.

The network namespace puts interfaces into namespaces, processes in the same net space can talk to the interface. Routing, forwarding, filters and bridging still happen in the kernel. in <cloud>AWS VPC</cloud>, <cloud>ENI</cloud> devices are mapped onto instances, which are then mapped to network namespaces that the application uses.

The Mount namespace controls the filesystem, it maps paths between the local namespace up to the root file system. The user namespace maps users from inside the namespace to outside users, but it's not recommended to use for managing users.

> - Linux Containers, as of today, sit on a shared Kernel
> - They sit on a shared platform,
> - They can influence each other quite easy.
> - Even if process-to-process isolation tight, it's just one layer.
> - Networking is always a discussion.

</details>

### Amazon Elastic Container Service (ECS) Primer

<details>
<summary>
Basic ECS overview.
</summary>

> This course goes beyond the basic concepts and benefits of containerization and teaches you more about the Amazon Elastic Container Service (ECS). You will learn about the implementation of containers on AWS using ECS and complementary services, such as the Amazon Elastic Container Registry (ECR). You will also learn about common microservices scenarios.

The goals of the course are:

- Familiarity with the <cloud>ECS (Elastic Container Service)</cloud>
- Understanding the difference between <cloud>EC2</cloud> and <cloud>Fargate</cloud> launch types
- Integrating <cloud>ECS</cloud> with other services
- Enforcing security on <cloud>ECS tasks</cloud>

#### What is Amazon ECS

review of containers and microservice architecture.

containers are a form of virtualization, happening at the operating system level, more lightweight than virtualizing complete operating systems. containers use Images, which are immutable "blueprints" to create containers from.\
Containers are strongly associated with microservice architecture, as they provide small scale applications with clearly defined apis. this fits well with the goals of having decoupling, agile and quick development.

#### Ecs Scalability and Micro Architecture

a host can easily run one or two containers on itself, but when there are tens of hosts and thousands of containers, things become messy. this gets worse with production environments, which need to be resilient, highly available, and support CI-CD for rapid development. this turns into a cluster management problem, and requires container management (orchestration) service.\
Theses service control health check, load balancing, monitoring, logging, networking, and replacing containers as needed. there are several tools, such as <cloud>ECS</cloud>, <cloud>Docker Swarm</cloud> or<cloud>Kubernetes</cloud>. <cloud>ECS</cloud> is highly available, high performance, AWS native orchestrator tool, it's highly integrated with other AWS services, and can schedule with its' own schedule or use a custom one.

there are two types of configurations: services and tasks. there are also two types of launch profiles, <cloud>EC2</cloud> and <cloud>Fargate</cloud>.\
Fargate launch type is closer to serverless architecture, as AWS manages provisioning the compute resources, EC2 launch types are better when the running instance is important and when usage requirements are known. they can also be mixed together.

#### ECS Components

<cloud>ECS tasks</cloud> are the smallest unit in ECS, they have a set of containers, and run once (or at intervals). for long running applications, services provide ability to scale out and scale in, and are aware of the Availability Zone they are in, so they support high availability spreads and can have a load balancer to manage traffic.

Tasks (either standalone or in a service) are defined in <cloud>Task Definition</cloud> files. those are the blueprints for creating tasks. they contain the name, memory resources, mounting points and what containers are running in the task (with which image).
task definitions also define which launch type is used: EC2 or Fargate.

With EC2 launch types, Tasks are hosted on EC2 instances, which run a docker agent and an ecs agent, those agent send telemetry to the ECS back-plane. with Fargate launch type, AWS manages the instances directly, saving the need for configuring the cluster.

#### Task Placements

when we use EC2 launch type, the task scheduler should place the task onto one of the instances. this is chosen based on a few filters:

- Cluster Constrains - CPU, memory, networking requirements
- Custom Constraints - location (Availability Zone), instance type, <cloud>AMI</cloud>
- Placement Strategies - best effort

those constrains are defined in the task definition, the placement strategies are best-effort, so a task can run on a instance even if it doesn't fit the placement strategy.

- Random
- BinPack - least available instance in terms of CPU and memory, trying to max out utilization.
- Spread - spread across instances based on some metric, such as Availability Zone.

there are also placement constraints, **bindings**, which can prevent a placement. they are not "best-effort".

- distinctInstance - only one task allowed on an instance (like Kubernetes daemonSet)
- memberOf - based on an expression (such as instance type or Availability Zone)

services also use placement strategies and constraints. and they have the "distinctInstance" option.

(examples of placing instances)

#### ECS Integration With Other AWS Services

Service | Purpose
---|---
<cloud>ECR</cloud> | Container Images
<cloud>SQS</cloud> | Decoupling
<cloud>SNS</cloud> | Decoupling
<cloud>ELB</cloud> | Load Balancing
<cloud>Route53</cloud>| DNS
<cloud>IAM</cloud>| Authentication and Authorization
<cloud>Secret Manager</cloud> | passwords and other secrets
<cloud>API Gateway</cloud> | exposing services
<cloud>Code Pipeline</cloud> | CI-CD
<cloud>CloudWatch</cloud> | Monitoring and logging

<cloud>ECR</cloud> is an cloud based AWS native image registry, highly available, secure, with at-rest encryption and fully integrated with <cloud>IAM</cloud> and the <cloud>ECS</cloud>.

ECS are compatible with DNS and can register themselves at route53 and expose themselves to other services.

example of ci-cd with <cloud>AWS CodeCommit</cloud>, <cloud>CodePipeline</cloud>, <cloud>CodeBuild</cloud>, <cloud>ECR</cloud>, and <cloud>CloudFormation</cloud>.

when there are new versions, it's possible to use "blue-green" deployment, with "green" being the new version, and "blue" being the old. both are running at the same time, and the load balancer directing traffic at them. this lessens the risk of deploying changes. ECS can also <cloud>autoscaling groups</cloud> and policies to scale up and down instances based on demand.

#### Security Enforcement on ECS

Each task has it's own IAM role, which gives it the specific permissions it needs, following the principle of least privilege. Tasks can retrieve secret from the <cloud>Parameter Store</cloud> and the <cloud>Secret Manager</cloud>, this is done again with <cloud>IAM roles</cloud>.

there are two additional scheduling strategies:

- replicas - always have a consistent number of tasks running.
- daemon (EC2 only) - always have the task running once on each of the EC2 instances.

</details>

### Amazon Eks Primer

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

> Kubernetes is a powerful container orchestration system that is the backbone of many microservices architectures, but it has a steep learning curve and is complex to manage. With Amazon Elastic Kubernetes Service (Amazon EKS), you can run Kubernetes on Amazon Web Services (AWS) without needing to install, operate, and maintain your own Kubernetes control plane.

#### Introduction To EKS Primer

<details>
<summary>
Basic EKS and Kubernetes Concepts
</summary>

> Amazon Elastic Kubernetes Service (<cloud>Amazon EKS</cloud>) is a managed container orchestration service that facilitates deploying, managing, and scaling Kubernetes applications in the AWS Cloud or on premises. Amazon EKS helps you provide highly available and secure clusters. Amazon EKS also helps you automate key tasks such as patching, node provisioning, and updates.

(video)

containers are light-weight virtualization, packaging runtime and software together, without the overhead of the entire operating systems. when used at scale, containers require a managing and orchestrating tool.

- scheduling and placement
- automatic scaling the number of containers
- removing unhealthy containers and replacing them with new ones
- integration with the cloud and other services, such as networking and persistent storage
- centralized security and monitoring

EKS benefits:

- managed Kubernetes across multiple Availability Zones, reducing points of failure
- tight integration with other AWS services, such as IAM and load balancing
- part of the Kubernetes community - works with existing plug-ins and configurations, portable and easy to migrate.

##### Kubernetes Objects

a review of Kubernetes objects and concepts:

- Cluster - a set of worker machines (nodes), a cluster has a least one node, and a cluster has a <cloud>control plane</cloud> that runs services and manages the cluster.
- Node - a physical or virtual machine that has workloads. managed by the control plane.
- Pod - a group of one or more containers, the basic building block of Kubernetes.
- Volumes - data storage:
  - Ephemeral volume - data storage that is tied to the life time of the pod, persists across pod restarts, but when a pod ceases to exist, it's also removed.
  - Persistent volume - data storage that has independent lifecycle and is not tied to any pod. can be backed up by another storage subsystem that is outside of the cluster node.
- Service - a logical collection of pods and a way to access them. tracks the number of available pods.
- Namespace -  a logical separation between workloads, can be useful to separate teams and projects that use the same cluster.
- ReplicaSet - ensuring a number of pod replicas are running at the same time
- Deployment - owns replicaSets and pods, manages the desired state.
- ConfigMap - api object that stores non-confidential data.
- Secret - storing confidential data.

Pods are deployed and removed according to the scheduler. it checks the resources required by the pods and then finds nodes to run them on. running through a set of filters:

- volume filters - volume requirements and constraints.
- resource filters - networking, cpu, memory.
- topology - constraints set at the node or pod level
- prioritization - (other stuff).

- **Control plane**: Control plane nodes manage the worker nodes and the pods in the cluster.
  - Controller Manager - background threaded that detect and respond to cluster events.
  - Cloud Controller - interacts with the underlying cloud provider.
  - Scheduler - selects where to place newly created workloads.
  - Api Server - exposes the Kubernetes api, frontend for the control plane. scales horizontally.
  - Etcd - key value dictionary that is the core persistence layer. stores critical cluster and state data.
- **Data plane**: Worker nodes host the pods that are the components of the application workload.
  - Kube-Proxy - maintains networking rules, performs connections and request forwarding if needed.
  - Container Runtime - can be Docker, Containerd, or something else.
  - Kubelet - the primary agent that runs on the worker nodes and manages their health.
  - Pods - a group of one ore more containers, can interact with other pods. containers in a pod always exist on the same node, and are scheduled together.

there are also custom resources which can be defined as CRD and controlled with custom controllers (<cloud>operators</cloud>). there is a command line tool <cloud>kubectl</cloud> to manage Kubernetes cluster and resources.

##### Amazon EKS and the Control Plate

in a self-managed Kubernetes clusters, the cluster owner is responsible for all the components of the control plane and the worker nodes.

> Amazon EKS provides a scalable, highly available control plane. Amazon EKS automatically manages the availability and scalability of the Kubernetes API servers and the etcd persistence layer for each cluster.

EKS has a least two API servers and three etcd nodes across three Availability Zones. unhealthy control planes are automatically replaced, which reduces the operational burden for running a cluster. the user still has to provision worker nodes (EC2 machines) to run the applications on. there are two control plane tools available in AWS: the general <cloud>aws cli</cloud> that works with many other aws resources, and the specialized <cloud>eksctl</cloud> that wraps over <cloud>CloudFormation</cloud> objects to provision resources.

to be clear: there are two api servers in question: one is the amazon EKS API and one is the Kubernetes API.

##### Amazon EKS and the Data Plate

While AWS manages the control plane nodes, the user is in charge of the worker nodes that run the applications, but even here, the level of responsability can be changed.

- Self Managed nodes: only the control plane is managed by aws
- Managed nodes:  aws takes more control, allows for easier provisioning, managing, updating and scaling. but the resources used are always visible.
- Fargate: offloading thr resource creation and management to AWS entirely, giving up control of the data plane and allowing aws to provision and manage the worker nodes natively. requires creating a <cloud>Fargate profile</cloud> on the cluster with proper selectors. in this case, every pod has a unique host(no two pods share a host), and there is no visibility into the host (via ssh or otherwise).

##### Quiz: Choosing The Correct API

- Q: Creating a Cluster
- A: Amazon EKS
- Q: Delete a managed node group
- A: Amazon EKS
- Q: Create a deployment
- A: Kubernetes
- Q: Get the fargate profile
- A: Amazon EKS
- Q: Get all the namespace
- A: Kubernetes

</details>

### Configuring Amazon Eks

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

</details>

### Separator

</details>
