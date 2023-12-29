<!--
ignore these words in spell check for this file
// cSpell:ignore elbv2 Neumann cgroups pictShare Kubelet eksctl Karpenter kube-proxy kubeconfig kube-system Alexa omponent
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Developer Learning Plan

[Developer Learning Plan](https://explore.skillbuilder.aws/learn/lp/84/developer-learning-plan)

> A Learning Plan pulls together training content for a particular role or solution, and organizes those assets from foundational to advanced. Use Learning Plans as a starting point to discover training that matters to you.\
> This Learning Plan is designed to help Developers who want to learn how to develop modern applications on AWS. The digital training included in this Learning Plan will expose you to developing with serverless and container technologies, as well as the foundation of DevOps on AWS. This Learning Plan can also help prepare you for the AWS Certified Developer - Associate certification exam.

## Introduction Elastic BeanStalk

<details>
<summary>
Deploy Servers, Databases and Load Balancers in an integrated way.
</summary>

> AWS <cloud>Elastic Beanstalk</cloud> provides you with a platform enabling you to quickly deploy your applications in the cloud. This course will briefly discuss the different components of the AWS Elastic Beanstalk solution, and perform a demonstration of the service.

as developers, we want to get our application to the cloud quickly. <cloud>Elastic Beanstalk</cloud> belongs to the _Platform as a service_ family of AWS features. it reduces management complexity and allows re-using existing code (some languages), it also allows for some control over the running infrastructure, such as the instance type, database and auto scaling.

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
3. security - satisfies security requirement for military, finance and government standards.
4. pace of innovation - continually adding services and technology, giving more options and making development easier.

</details>

</details>

## Containers

<details>
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

> - _Portable_ runtime application environment
> - Package application and dependencies in a _single, immutable artifact_
> - Ability to run different application versions with different dependencies _simultaneously_
> - _Faster_ development and deployment cycles
> - Better _Resource utilization and efficiency_

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

| Service                       | Purpose                          |
| ----------------------------- | -------------------------------- |
| <cloud>ECR</cloud>            | Container Images                 |
| <cloud>SQS</cloud>            | Decoupling                       |
| <cloud>SNS</cloud>            | Decoupling                       |
| <cloud>ELB</cloud>            | Load Balancing                   |
| <cloud>Route53</cloud>        | DNS                              |
| <cloud>IAM</cloud>            | Authentication and Authorization |
| <cloud>Secret Manager</cloud> | passwords and other secrets      |
| <cloud>API Gateway</cloud>    | exposing services                |
| <cloud>Code Pipeline</cloud>  | CI-CD                            |
| <cloud>CloudWatch</cloud>     | Monitoring and logging           |

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

<details>
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
- Namespace - a logical separation between workloads, can be useful to separate teams and projects that use the same cluster.
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
- Managed nodes: aws takes more control, allows for easier provisioning, managing, updating and scaling. but the resources used are always visible.
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

#### Configuring Amazon EKS

<details>
<summary>
Setting Up an EKS cluster
</summary>

> There are three tasks to perform when building a cluster:
>
> 1. Secure your AWS environment.
> 1. Configure the <cloud>virtual private cloud (VPC)</cloud> networking for the cluster.
> 1. Create the Amazon <cloud>EKS</cloud> cluster.

##### Perpearing your AWS Environment

The AWS shared responsability model. aws protects the physical infrastructure, while customers protect the content and content security data. the level of separation changes between self-managed, managed, and fargate configurations.

- self managed - aws manages the control plane and the user manages the data plane (<cloud>IAM</cloud>, pod security, runtime security, network security, code security)
- managed node groups - aws manages building the optimized AMI OS,and the user has to deploy it.
- Fargate - aws is responsible for sacling node workers.

authentication - who somebody is? either a user or a service. authorization - what can that identity do? which permissions does it posses? EKS uses IAM for authentication, and the native RBAC for authorization.

- cluster IAM role - managing EKS and <cloud>EC2</cloud> machines.
- node IAM role - like pulling images from <cloud>ECR</cloud>
- RBAC user - mapped from an IAM role, by default this is too strong, so there should be granular control.

Networking is configured, an EKS must live inside a <cloud>VPC</cloud>. which can have public subnets, private subnets, or a mix.

##### Creating A Cluster

demo of creating a cluster, using the eksctl tool.

eksctl uses <cloud>AWS CloudFormation</cloud> stacks to create <cloud>EKS</cloud> clusters. it has some default actions which make it simple and easy to use.

- Creates IAM roles for the cluster and worker nodes.
- Creates a dedicated VPC with Classless Inter-Domain Routing (CIDR) 192.168.0.0/16.
- Creates a cluster and a nodegroup.
- Configures access to API endpoints.
- Installs CoreDNS.
- Writes a kubeconfig file for the cluster.

```sh
# download eks and install it
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin

# create a cluster from a config file

eksctl create cluster -f ./prod-cluster-config.yaml

# view nodes
kubectl get nodes
```

video demo of managing eks cluster with the web management console. seeing the cluster workloads on the cluster, must have three workloads which define aws functionality.

- aws-node
- coredns
- kube-proxy

seeing the compute type, networking, security groups, connecting with OIDC as an additional way to log in into the cluster. cluster logging is disabled by default, and could be enabled at any time.

##### Configuring Horizontal and Vertical Scaling

> - Horizontal scaling: A horizontally scalable system is one that can increase or decrease capacity by adding or removing compute resources. In this example, more pods are deployed when demand spikes (scale out). Pods are removed when demand drops (scale in).
> - Vertical scaling: A vertically scalable system increases performance by adding more resources to the compute resource, such as faster (or more) central processing units (CPUs), memory, or storage.

in Kubernetes, Horizontal scaling is much easier than it is in traditional server infrastructure, vertical scaling is still a possibility. the number of nodes can be adjusted by using a **Cluster Autoscaler**, which means attaching the worker nodes to an <cloud>EC2 Auto Scaling Group</cloud>. there is an alternative to autoscaling, called **Karpenter**, which is a node lifecycle management solution.

Kubernetes has an internal scaling option, which can be configured to either track cpu utilization or metrics from cloud watch.

```sh
kubectl autoscale deployment myApp --cpu-percent=50 --min=1 --max=10
```

Kubernetes also has a vertical scaling option, which can scale down pods with over-requested resources to allow new pods to scheduled.

(demo video)
deploying a web application and see auto scaling, the 'cluster-autoscaler' and 'metrics-server' run as deployments. we can simulate an increase in demand:

```sh
kubectl get pods -o wide -A
kubectl apply -f https://k8s.io/examples/application/php-apache.yaml
kubectl describe deployment php-apache
kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
# on a different terminal
watch "kubectl get nodes; echo; kubectl get hpa, pods -o wide"
# back to the original terminal
kubectl run -i --tty load-generator --rm --image=busybox --restart=Never -- /bin/sh/ -c "while sleep 0.01; do wget -q -O -http://php-apache; done"
```

##### Managing Communication in Amazon EKS

> To simplify inter-node communication, Amazon EKS integrates Amazon VPC networking into Kubernetes through the Amazon VPC <cloud>Container Network Interface (CNI)</cloud> plugin for Kubernetes. The Amazon VPC CNI plugin allows Kubernetes pods to have the same IP address inside the pod as they do on the Amazon VPC network.

a <cloud>VPC</cloud> is an isolated partition in a data center, it's created inside a single region, but spans across all Availability Zone in that region. a VPC has an address range (a CIDR block).\
EKS communication can be:

- between containers in the same pod
- between pods (either on the same node or not)
- ingress connections from outside he cluster

communication inside a pod uses the local host, this doesn't require a <cloud>NAT</cloud> (network address translation). comunication between pods uses Linux namespaces and their internal routing tables. communication from outside uses the <cloud>VPC CNI</cloud> plugin. it provides each pod with an IP address inside the VPC, even though they all live on the same host machine. this uses secondary ip addresses.

Services are Kubernetes native solution for communication with pods (which might disappear and be replaced at any moment). A Kubernetes Service provides a consistent IP address and acts as the entry point for a number of pods.

(video)

- Cluster ip - fixed ip address, only available internally
- NodePort - fixed port for nodes, externall
- LoadBalancer - connects with the cloud vendor load balancer
- ExternalName - maps an internal address into an external resource.

> Ingress:\
> With Kubernetes ingress objects, you can reduce the number of load balancers you use. An ingress object exposes HTTP and HTTPS routes from outside the cluster to your services and defines traffic rules.
>
> AWS Load Balancer Controller:\
> The AWS Load Balancer Controller is a controller that manages Elastic Load Balancing (ELB) for a Kubernetes cluster. The load balancers can be Application Load Balancers when you create a Kubernetes Ingress or Network Load Balancers when you create a Kubernetes service of type LoadBalancer.\
> An Application Load Balancer balances application traffic at Layer 7 (for example, HTTP or HTTPS) of the Open Systems Interconnection (OSI) model, while a Network Load Balancer balances network traffic at Layer 4 [for example, Transmission Control Protocol (TCP), User Datagram Protocol (UDP), and so forth]. Application Load Balancers can be used with pods that are deployed to nodes or to Fargate.\
> Application Load Balancers can be deployed to public or private subnets. Network Load Balancers can load balance network traffic to pods deployed to Amazon EC2 IP and instance targets or to Fargate IP targets.

</details>

#### Integrating Amazon EKS with Other Services

<details>
<summary>
Running your workloads on an Amazon EKS cluster provides the benefit of using other AWS service.
</summary>

##### Managing Storage in Amazon EKS

> Running your workloads on an Amazon EKS cluster provides the benefit of using other AWS services including several storage services. In this lesson, you explore how to manage your application workload storage requirements with Amazon Elastic Block Store (Amazon EBS) and Amazon Elastic File System (Amazon EFS)

###### Kubernetes Persistent Storage

> Application workloads requiring data persistence independent of the pod lifecycle require at least two Kubernetes objects, a <cloud>persistent volume (PV)</cloud> and a <cloud>persistent volume claim (PVC)</cloud>.

Persistent Volume storage are still stored on the cluster, but they aren't tied to any pod, and they have their own lifetime. there is also a <cloud>Storage Class</cloud> object which allows for scaling and management of the storage. a final object is the <cloud>Container Storage Interface (CSI)</cloud> driver that connects the cluster to an external storage provider. in EKS, there are <cloud>CSI</cloud> drivers for <cloud>Elastic Block Storage (EBS)</cloud> and <cloud>Elastic File System (EFS)</cloud>.

> When a cluster user submits a PVC with the requisite parameters, the Amazon EBS storage class calls on the EBS CSI driver to allocate storage per the PVC request. The EBS CSI driver makes the necessary AWS API calls to create an EBS volume and attach the volume to the designated cluster node. When attached, the persistent volume is allocated to the PVC. The Amazon EBS CSI driver can be configured to use the various functionality of Amazon EBS including volume resizing, creating volume snapshots, and so forth.
>
> A Kubernetes storage class backed by Amazon EFS will direct the Amazon EFS CSI driver to make calls to the appropriate AWS APIs to create an access point to a preexisting file system. When a PVC is created, a dynamically provisioned PV will use the access point for access to the Amazon EFS file system then bind to the PVC.

Fargate Nodes work well with <cloud>EFS</cloud>, and they don't require installing a <cloud>CSI</cloud>.

(video)

attaching <cloud>EBS</cloud> storage to a cluster.

```sh
kubectl get pv, all -A
# download IAM role
curl -o example-iam-policy.json https://raw..githubusercontent.com/kubernetes-sigs/aws-ebs-csi-driver/master/docs/example-iam-policy.json
aws iam create policy --policy-name EBC_CSI_Driver_Policy --policy-document file://example-iam-policy.json
eksctl create iamserviceaccount --name ebs-csi-controller-sa --namespace kube-system --cluster dev --attach-policy-arn <policy> --approve --override-existing-serviceaccounts

# helm
helm repo add aws-ebs-csi-driver https://kubernetes-sigs.github.io/aws-ebs-csi-driver
helm repo update
helm upgrade --instal aws-ebs-csi-driver aws-ebs-csi-driver/aws-ebs-csi-driver --namespace kube-system --set image.repository=<> --set controller.serviceAccount.create=false --set controller.serviceAccount.name=ebs-csi-contoller-sa
# verify that works properly
# not gonna write this down
```

##### Deploying Applications to Amazon EKS

we mostly use the `kubectl` cli tool to deploy applications, but this doesn't scale well. instead, there are other options. we can set up a ci-cd pipeline using aws services:

- <cloud>AWS CodePipeline</cloud>
- <cloud>AWS CodeCommit</cloud>
- <cloud>AWS CodeBuild</cloud>

we commit new code and trigger a pipeline, which runs tests and builds an image, this image is pushed into <cloud>ECR</cloud>. at the same time, a <cloud>Lambda</cloud> is triggered which invokes the Kubernetes api to update the application, which then pulls the new image from the repository.

we can also use other tools for ci-cd and integrate them with EKS. for example, a pipeline from github uses a webhook to trigger Jenkins, then the image is stored in a Harbor repository. This then activates a Spinnaker pipeline which can include creating a new Helm manifest and deploying it onto a cluster.

##### Gaining Observability

> Observability is the ability to analyze and view data or processes. It is achieved only after monitoring data (such as metrics) has been compiled. _Observability_ is a term often used interchangeably with _monitoring_, but they are two different concepts.

observability comes in several forms: metrics, logs and traces. metrics are structured data points that can be aggregated and visualized, such as health checks and resource utilization. logs are less structured, and are created by each container and can contain valuable data. Traces follow the path of requests across multiple services.

<cloud>AWS CloudWatch</cloud> has container insights for eks clusters, and can collect diagnostic information. the cloudwatch agent runs as a daemonset on every node and uses fluentBit (a lightweight version of FluentD) to collect logs. there are also open source tools for collecting logs and metrics, such as <cloud>Grafana</cloud>, <cloud>Prometheus</cloud>, <cloud>OpenSearch</cloud>, etc...

(video)

enabling control plane logging for components:

- API Server
- Audit
- Authenticator
- Controller Manager
- Scheduler

using the same load generator from the previous section to create logs. and then seeing those logs in <cloud>AWS CloudWatch</cloud> at the correct log group. we can choose <cloud>Container Insights</cloud> to see metrics about our clusters.

##### Deploying a Service Mesh with AWS App Mesh

> AWS App Mesh is a service mesh that provides application-level networking to help your services communicate with each other across multiple types of compute infrastructure. App Mesh gives end-to-end visibility and high availability for your applications

the more fragmented our microservices become, network communications become more important and complicated. there are challenges to development, visibility, security and scaling (traffic management).

> A better way to address managing service communication at scale is with a service mesh. A service mesh is a dedicated infrastructure layer in which you can abstract away the service-to-service communication. The abstraction is done with an array of lightweight network proxies deployed alongside the application code.\
> Developers can use a service mesh to focus on the business application instead of focusing on configuring the intelligence of the network. By removing the logic from the application code, you keep the services smaller and the logic more consistent.\
> A service mesh monitors all service-to-service traffic and abstracts its configuration. The mesh tracks all data on the wire, which you can use.

The service mesh is deployed on the cluster control plane, and then each pod has mesh proxy sidecar container controlling the traffic.

</details>

#### Maintaining Your Amazon EKS Cluster

<details>
<summary>
Updating EKS Clusters to New Kubernetes Versions.
</summary>

> Add-ons extend the operational capabilities of Amazon EKS clusters but are not specific to any one application. This includes software like observability agents or Kubernetes drivers that allow the cluster to interact with underlying AWS resources for networking, compute, and storage. Add-on software is typically built and maintained by the Kubernetes community, cloud providers like AWS, or third-party vendors.

##### Maintining Add-ons

> Amazon EKS automatically installs on every cluster the following add-ons:
>
> - Amazon VPC CNI
> - kube-proxy
> - CoreDNS

they have two flavours - managed and self-managed.

There are also third-party add-ons, but they can get out of date if a new Kubernetes version comes and breaks the API the add-ons use.

addons can viewed in the configuration tab of in the aws management console.

##### Maintining Upgrades

upgrading starts with deploying the new API Server nodes and checking the health of the nodes, if a problem is encountered then it rollback is performed. a possible cause for failure is not having enough free ip addresses in the subnet. aws only updates the control plane nodes, worker nodes must be upgraded by the user (but not fargate nodes).

for a managed node group, aws creates a new EC2 machine (based on the autoscaling group and launch template), it then removes one of the existing nodes. this is repeated until the upgrade is completed.

> There are two options for the update strategy:
>
> - Rolling update – This option respects PodDisruptionBudgets for your cluster. The update fails if Amazon EKS is unable to gracefully drain the pods that are running on this node group because of a PodDisruptionBudget issue.
> - Force update – This option does not respect PodDisruptionBudgets. It forces node restarts.

for self-managed node groups, the process needs to be done manually, either by creating a new node group and migrating to it or by updating the <cloud>CloudFormation</cloud> stack node group to se the new AMI.

(video)
demo of upgrading an <cloud>EKS</cloud> cluster version from 1.20 to 1.21. in the web console, there is a mesage to update a cluster. in the <kbd>configuration</kbd>, we can click <kbd>Update now</kbd> and select the new version. this will update the control plane only.\
Once done, we are prompted to upgrade the ami for the node groups. we proceed with the update by selecting each node group and choosing the same version as we updated the control plane to, we have the option of selecting which update strategy to use. We should also update the add-ons of the cluster in the <kbd>Add-Ons</kbd> tab.\
We can always see the log of the cluster upgrade in the <kbd>update history</kbd> tab (either for the cluster itself or the node group) or in the add-on page.

</details>

#### Managing Amazon EKS Costs

<details>
<summary>
EKS Costs.
</summary>

> Part of the Six Pillars of the AWS the Well-Architected Framework is cost optimization, which is summarized as running systems to deliver business value at the lowest price point.

The Bulk of the costs associated with an <cloud>EKS Cluster</cloud> come from the compute resources. the control Plane and Networking services make up a smaller share of the total costs.

Compute resources follow they same cost options as <cloud>EC2 machines</cloud>:

- On Demand - "pay as you use"
- Saving Plans and Reserved Instances - make a commitment for 1 or 3 years and receive a bulk discount
- Spot and Fargate Spot - for stateless or fault tolerant workloads that aren't time critical.

Using managed node groups also can decrease the costs.

</details>

</details>

</details>

## Serverless Computing

<!-- <details> -->
<summary>
Several Courses Focusing on Serverless.
</summary>

### Introduction to Serverless Development

<details>
<summary>
Introduction for Servereless for Developers.
</summary>

> Writing Code
>
> - System Architecture
> - Design Patterns
> - Frameworks and Libraries
>
> Managing Code
>
> - Tools
> - Developer Workflows
> - Test/Deployment Automation
> - Environment Management

(video)

moving to serverless has benfits in costs (both upfront and scaling), ease of scaling, and it helps with making development cycles faster and reaching the market quicker.

- <cloud>Lambda</cloud> - Compute
- <cloud>API Gateway</cloud> - API proxy
- <cloud>S3</cloud> - Storage
- <cloud>DynamoDB</cloud> - Database
- <cloud>SNS</cloud>, <cloud>SQS</cloud> - Interprocess Messaging
- <cloud>Step Functions</cloud> - Orchestration
- <cloud>Kinesis</cloud>, <cloud>Athena</cloud> - Analytics
- <cloud>Developer Tools</cloud> - Tools

we will focus mostly on <cloud>Lambda</cloud>, and event based compute service. it has an <cloud>IAM resource Policy</cloud> for what can trigger the code, and a different policy for what the code can do and interact with (based on the execution role).

#### Writing Lambda Functions

(video)

best practices for writing lambda code are the same as regular code. a code should have `Handler` method, which acts as an entry point. the typical layout divides the code into layers:

1. Handler - Function configuration, <cloud>Lambda</cloud> specific code and no business logic.
2. Controller - event processing, the core business logic.
3. Service - external integrations, connections to other services.

(video)

we can include external libraries and fraeworks, this will effect the speed in which the lambda operates.\
lambda has two "starts" - the "cold start" when the code is downloaded and the execution runtime starts, and the "warm start", when the event is triggered. the smaller the size of the library (less dependencies) the better. it might be better to have statically linked libraries.

#### MManaging Serverless Applications

(video)

The developer tool chain doesn't have to change when moving to serverless, but there is an additional step of packaging and deploying the application onto the lambda service. if the code is spread across multiple lambdas that need to interact, then that's an additional issue to consider. even a simple API can become a very complicated <cloud>CloudFormation</cloud> template. .because of that, it's suggested to use an application framework, such as <cloud>AWS SAM (Serverless Application Model)</cloud> which simplifies the process. it defines everyting in a smaller way. it also has a CLI tool to package the code (and upload to S3).

```sh
sam package
sam deploy
```

there are also other options besides <cloud>AWS SAM</cloud>, such as APEX, Zappa, Chalice, Sparta, Claudia.js, and more.

(video)

The Code should be orgainzed in the source code repository. a Repository can contain one function or all of them. the way to think about this is to divide the function into "services", each being one or more functions (and other services) that work together, and they should be in the same separate repository.

(video)

there should a local environment (IDE) to write and develop the code, but there should also be a AWS environment where the code is deployed and tested using real AWS resources. each developer can have it's own sandbox or have the whole team use a shared sandbox. it's better to separate production environment from the development account.

#### Testing and Debugging Serverless Applications

(video)

hierarchy of testing - locally, remotely, and having integration (regression, acceptance) tests. Unit tests should focus on the business logic (the controller layer), and have mockups for the other services. there are mocking options such as <cloud>LocalStack</cloud> and <cloud>DynamoDB local</cloud>, and there is an option to build a custom mocking solution (which is complicated). it's also possible that the local tests simply run against the cloud.

(video)

besides testing, we should also be able to debug the code, this is harder to do with <cloud>Lambda</cloud> functions. it's impossible to run a step-by-step debugger on an active instance. however, with the <cloud>SAM CLI</cloud> tool, it is possible to launch a docker container with the lambda code and then debug it.

</details>

### Getting Into The Serverless Mindset

<details>
<summary>
Overview of Serverless Architecture Mindset
</summary>

(video)

Key concepts of serverless architecture, and changing how we think. a "traditional" web service application would probably use and <cloud>EC2 machine</cloud> and some database. we can grow and add machines for increased demand and use a <cloud>Elastic Load Balancer</cloud>, and use an <cloud>EC2 Auto Scaling Group</cloud> to start and terminate machines as needed. we can even spread it across multiple Availability Zones for high availability. this model is fine, but it has a lot of over head to consider: machine and OS patching, scaling, health checks... this is all time consuming and detracts from focusing on the core business logic. and since this is similar for most applications, it's a lot easier to offload this burden onto the cloud Vendor.

#### How Can Serverless Help?

(video)

in a serverless world, the architecture changes. if we focus on the core components, we can determine that there are requests from clients, some business logic applied to it, and a backend / database. we could consider the requests as 'events' and the business logic as 'handling events', and that leads us to <cloud>Lambda</cloud> functions.\
When we say "serverless", what we means is:

- No Server Management
- Flexible Scaling
- Automating High Availability (Fault Tolerance)
- No Idle Capacity

the same benefits of moving to the cloud apply to serverless architecture.

#### Managed Services Connected by Functions

(video)

when we move from tractional architecture to serverless, there are other shifts that we should make. <cloud>Lambda</cloud> is a core component, but there are other services, such as <cloud>API Gateway</cloud>, databases, messaging, and storage.\
Multiple functions can be combined with <cloud>Step Functions</cloud>, deployment and pipelines can be handled by other services. moving to serverless allows the application to reduce itself and stop being a monolith, it makes the application clearer and easier to handle. we do this by thinking about "Event-orient architecture", where each event (not just api calls) can trigger a flow.

- Data Stores
  - <cloud>S3</cloud>
  - <cloud>DynamoDB</cloud>
  - <cloud>Kinesis</cloud>
  - <cloud>Cognito</cloud>
  - <cloud>RDS - Aurora</cloud>
- End Points
  - <cloud>Alexa</cloud>
  - <cloud>API Gateway</cloud>
  - <cloud>IOT Core</cloud>
  - <cloud>Step Functions</cloud>
- Repositores
  - <cloud>CloudFormation</cloud>
  - <cloud>CloudTrail</cloud>
  - <cloud>CloudWatch</cloud>
  - <cloud>CodeCommit</cloud>
- Event / Message Services
  - <cloud>SES</cloud>
  - <cloud>SQS</cloud>
  - <cloud>SNS</cloud>
  - Cron Events

#### From Hours to Minutes with Parallelization

(video)

another thing with serverless architecture is how it works with parallelism, since we don't have idle servers running, we can take advantage of this and focus on breaking down the tasks into parallelized work. a task that takes an hour on a single machine could be broken into 360 task each taking 10 seconds and be completed in less than a minute.\
one example is transcoding video to different formats. the work is split between many lambdas (each working on a small part of the video) and is then combined at the end. there is also "pyWren" which executes python code on <cloud>Lambda</cloud> functions.

#### More Environments, More Innovation

(video)

another benefit of serverless is that it allows creating and deploying more environment, like giving each developer a sandbox environment, or having an active environment for each feature. this makes automation more critical, and requires using an application framework such as <cloud>AWS SAM</cloud>.

(video)

serverless architecture allows more time to focus on providing unique value and frees time from infrastructure management.

</details>

### AWS Lambda Foundations

<details>
<summary>
Serverless and Lambda basics.
</summary>

> AWS Lambda is an event-driven, serverless compute service that lets you run code without provisioning or managing servers. This course focuses on what you need to start building Lambda functions and serverless applications. You learn how AWS Lambda works and how to write and configure Lambda functions. You explore deployment and testing considerations and finally end with a discussion on monitoring and troubleshooting Lambda functions.

#### Introduction to Serverless

Serverless applications are a further step in hiding away the infrastructure layer. there is no need to manage instances, operating systems and servers. the amount of operational tasks is reduced, allowing the customer to focus on core business development.

| Deployment and Operational tasks               | Traditional Environment | Serverless |
| ---------------------------------------------- | ----------------------- | ---------- |
| Configure an instance                          | YES                     | NO         |
| Update operating system (OS) YES               | NO                      |
| Install application platform                   | YES                     | NO         |
| Build and deploy apps                          | YES                     | YES        |
| Configure automatic scaling and load balancing | YES                     | NO         |
| Continuously secure and monitor instances      | YES                     | NO         |
| Monitor and maintain apps                      | YES                     | YES        |

Some AS Serverless services:

| Service                    | Type                   | Notes                        |
| -------------------------- | ---------------------- | ---------------------------- |
| <cloud>Lambda</cloud>      | Compute                |
| <cloud>Lambda@Edge</cloud> | Compute                | Compute on edge servers      |
| <cloud>S3</cloud>          | Storage                | Object Storage               |
| <cloud>DynamoDB</cloud>    | Storage                | Data Store, NoSQL            |
| <cloud>EventBridge</cloud> | Event Bus              |
| <cloud>SNS</cloud>         | Interprocess Messaging | Simple Notification Service  |
| <cloud>SQS</cloud>         | Interprocess Messaging | Simple Queue Service         |
| <cloud>API Gateway</cloud> | Api Integration        |                              |
| <cloud>AppSync</cloud>     | Api Integration        |                              |
| <cloud>CDK</cloud>         | Developer Tools        | Cloud Development Kit        |
| <cloud>SAM</cloud>         | Developer Tools        | Serverless Application Model |

<cloud>AWS Lambda</cloud> is an event driven, high-availability and automatically scaling compute service. it runs user code without requiring provisioning and management of instances. code is easily logged and monitored via <cloud>CloudWatch</cloud>. it supports integration with other aws services and has a flexible permissions model. it is highly available and scales with demand, but payment is based on usage.

Event Driven Architecture means that when changes happen, an event is published and other components can respond to it (consume it). events can be user initiated or generated from other aws services, such as changes in a database. Events are created by _Producers_, and then are sent to the lambda via a _Router_ (such as <cloud>EventBridge</cloud>) and eventually get acted on by _Consumers_.

Lambdas are stateless functions, or pieces of code designed to run in response to one of those events. they don't maintain state between invocations and don't rely on any data stored in the running instances. A lambda has:

- Access permissions - which services it is allowed to interact with.
- Triggering Events - which kind of events causes the lambda to run.
- Code - user provided code.
- Configurations - memory, timeout and lambda concurrency.

#### How AWS Lambda Works

Lambdas can be invoked synchronously, asynchronously and by "polling" depending on the kind of event.

Synchronous invocations returns the function response and data about the lambda invocation (version, runtime, etc...), there are no built-in "retries" in case the function fails.

The following AWS services invoke Lambda synchronously:

- <cloud>Amazon API Gateway</cloud>
- <cloud>Amazon Cognito</cloud>
- <cloud>AWS CloudFormation</cloud>
- <cloud>Amazon Alexa</cloud>
- <cloud>Amazon Lex</cloud>
- <cloud>Amazon CloudFront</cloud>

Asynchronous invocation model utilizes an event queue and is used when the response isn't immediately required. the event is queued and the invoking function (or user) continues without getting the response. With the asynchronous model, it's possible to send records of lambda invocation to a destination. this can be configured based on the result of the lambda or other conditions.

The following AWS services invoke Lambda asynchronously:

- <cloud>Amazon SNS </cloud>
- <cloud>Amazon S3</cloud>
- <cloud>Amazon EventBridge</cloud>

as part of Lambdas integration with other services, it can watch for changes in queue and streaming services, and poll for events that match a condition to trigger the lambda.

- <cloud>Amazon Kinesis</cloud>
- <cloud>Amazon SQS</cloud>
- <cloud>Amazon DynamoDB Streams</cloud>

When A lambda is invoked, it runs in a _Lambda Execution Environment_. this handles the resources (memory and cpu), manages the the lifecycle of the functionand provides external extensions. the lifecycle has three stages: Init, Invoke and Shutdown.

The Init stage includes creating the execution environment, downloading the code onto it with any decencies and running the functions' static code. the stage has three phases:

1. Extension init - starts all extensions
1. Runtime init - bootstraps the runtime
1. Function init - runs the function's static code

In the Invoke stage, the function handler code is run, and once completed, the lambda prepares for an addition invocation.

The Shutdown stage happens if the lambda did not receive any invocations after a period of time. the lambda shuts down the extensions and then removes the runtime environment.

> When you write your function code, do not assume that Lambda automatically reuses the execution environment for subsequent function invocations. Other factors may require Lambda to create a new execution environment, which can lead to unexpected results. Always test to optimize the functions and adjust the settings to meet your needs.

If we want our code to be performant, there are some ways to optimize it for Lambda invocations.
When a lambda is invoked, it can either go through a "cold" or "warm" start. a "cold start" happens when a new execution environment is required, in this case the lambda must download the code and initialize the runtime. a "warm start" happens when a lambda is invoked on an existing runtime. in this case, only the user code needs to be initialized (such as user packages and dependencies), and the code starts running faster. Billing begins at the "warm start" phase.

If the time needed for the "cold start" is important for the proper running of the application, it is possible to reduce it and make it predicable by using provisioned concurrency, which ensures that a known number of runtime exectution environments are active (passed the "cold start") at all time, so there is no suprise latency.

> Best practice: Write functions to take advantage of warm starts
>
> - Store and reference dependencies locally.
> - Limit re-initialization of variables.
> - Add code to check for and reuse existing connections.
> - Use tmp space as transient cache.
> - Check that background processes have completed.

#### AWS Lambda Function Permissions

The <cloud>Lambda</cloud> function has two sets of permissions, both fully controlled and integrated with <cloud>IAM</cloud>.

```json
{
  "Version": "2012-40-17",
  "Statement": [
    {
      "Sid": "Allow PutItem in table/test",
      "Effect": "Allow",
      "Action": "dynamodb:putItem",
      "Resource": "arn:aws:dynamodb:us-west-2:###:table/test"
    },
    {
      "Effect": "Allow",
      "Action": "sts:assumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      }
    }
  ]
}
```

- Resource policy - permissions to invoke the function
- Execution role - permissions for the lambda itself to act with

the execution role controls which other services the lambda can act upon, such as <cloud>DynamoDB</cloud>, <cloud>S3</cloud> and others. this role must include a trust policy for to allow the function to assume it.

The resource policy (also called function policy) controls who can invoke the lambda (users, roles, aws services, other aws accounts).

the resource based policy is easier to use than roles for most cases, but it does have a size limit.

> - Lambda resource-based (function) policy
>   - Associated with a "push" event source such as Amazon API Gateway
>   - Created when you add a trigger to a Lambda function
>   - Allows the event source to take the _lambda:InvokeFunction_ action
> - IAM execution role
>   - Role selected or created when you create a Lambda function
>   - IAM policy includes actions you can take with the resource
>   - Trust policy that allows Lambda to _AssumeRole_
>   - Creator must have permission for _iam:PassRole_

Policy management can be part of the <cloud>Sereveless Application Model</cloud>. the permissions are scoped to the resources used in the application, and are added in the <cloud>SAM</cloud> template.

Lambda can also be allowed to access resources inside <cloud>Virtual Private Cloud</cloud>. this is done by connecting the Lambda to the subnet and security groups in the <cloud>VPC</cloud> or through <cloud>AWS Private Link</cloud>.

#### Authoring AWS Lambda Functions

Lambda functions can be written in several programming languages, and has plugins for several IDEs.

- Node.js
- Python
- Java
- Go
- C#
- Ruby
- PowerShell

AWS continuously updates the supported versions for popular runtimes, with new versions being added and old versions removed (deprecated), the complete list is maintained in the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).

The Lambda execution begins at the function handler method, this method takes two arguments, the required event object and (optionally) a context object. the event object is mandatory, and each type of event creates a different event object. the object that represents a S3 event isn't the same as the object from an API gateway event. the context object contains information about the lambda function itself, such as the requestID, the specific runtime and which <cloud>CloudWatch</cloud> stream will contain the log statements from the lambda. other data depends on the specific runtime.

in terms of design, the best-practice suggests separating the business logic code from the lambda code, with the lambda specific code simply calling the actual code. this allows for portability and testing. another suggestion is writing modular functions - each lambda should only do one thing. Lambda functions are stateless, which means they shouldn't rely on any information saved in the context object (the runtime), every invocation can be run on a different runtime, so there is no consistent state. if data needs to be stored across invocations, it should be done in a different service:

> - <cloud>Amazon DynamoDB</cloud> is serverless and scales horizontally to handle your Lambda invocations. It also has single-millisecond latency, which makes it a great choice for storing state information.
> - <cloud>Amazon ElastiCache</cloud> may be less expensive than DynamoDB if you have to put your Lambda function in a VPC.
> - <cloud>Amazon S3</cloud> can be used as an inexpensive way to store state data if throughput is not critical and the type of state data you are saving will not change rapidly.

A final suggestion is to minimize the size and the number of dependencies inside the deployment package. This is done by minimizing the number of modules imported, correctly bundling packages and using lightweight frameworks if possible. A modular design that keeps the responsabilities of each lambda small and well defined is a great way to achieve this - if a function has a limited scope, then the amount of external dependencies is also limited. The size and complexity of a deployment package effects the start-up time of the lambda.

There are also best practice advice for the code itself:

- Include logging statements.
- Return a result from the lambda.
- Provide environment variables - separate configuration from the code itself. environment variables are encrypted using either the <cloud>Customer Master Key</cloud> or a user key from the <cloud>Key Management Service</cloud>.
- Manage Confidential data in the <cloud>Parameter Store</cloud> and retrieve it with the secrets manager. this can be combined with <cloud>AWS AppConfig</cloud>.
- Avoid Recursive code - **DON'T** call the lambda from itself. if this ever occurs, set the concurrent execution limit to zero to throttle requests while the code is being fixed (to avoid racking up costs).
- Gather metrics with <cloud>CloudWatch</cloud> - use the embedded metric format to automatically extracy metrics from logs.
- Reuse execution context - Take advantage of an existing execution context when you get a warm start by doing the following:
  > - Store dependencies locally.
  > - Limit re-initialization of variables.
  > - Reuse existing connections.
  > - Use tmp space as transient cache.
  > - Check that background processes have completed.

```js
console.log("Loading Function");
const aws = require("aws-sdk");
const s3 = new aws.S3({ apiVersion: "2006-03-01" });
exports.handler = async (event, context) => {
  // Get the object from the event and show its content type
  const bucket = event.Records[0].s3.bucket.name;
  const key = decodeURIComponent(
    event.Records[0].s3.object,
    key.replace(/\+/g, " ")
  );
  const params = {
    Bucket: bucket,
    Key: key,
  };
  try {
    const { ContentType } = await s3.getObject(params).promise();
    console.log("CONTENT TYPE:", ContentType);
    return ContentType;
  } catch (err) {
    console.log(err);
    const message = `Error getting object ${key} from bucket ${bucket}. Make sure they exists and your bucket is in the same region as this function.`;
    console.log(message);
    throw new Error(message);
  }
};
```

Lambda functions cab be built in several ways. it can be written directly in the consode editor using the <cloud>AWS Cloud9 IDE</cloud>, this allows for rapid testing. another advantage is that the console provides common blueprints and complete apps as built-in resources, we can also use a container image to deploy the lambda. larger code that rely on dependencies (other than the AWS SDK) can be uploaded as a deployment package - either a zip archive file or a container image in the <cloud>Elastic Container Registry</cloud>. there are also ways to automate lambda deployment with services such as <cloud>SAM</cloud>, <cloud>CodeBuild</cloud>, <cloud>CodeDeploy</cloud> and <cloud>CodePipeline</cloud>.

> <cloud>AWS Serverless Application Model</cloud> is an open-source framework for building serverless applications. It provides shorthand syntax to express functions, APIs, databases, and event source mappings.

<cloud>SAM</cloud> uses a yaml to define resource, permissions and mappings in a concise way (compared to the more detailed <cloud>CloudFormation</cloud> syntax) that creates a serverless application with lambdas, databases, apis and event mappings. it includes a number of prebuilt policies which follow the principle of least privleges. the SAM template is then expanded and built as a <cloud>CloudFormation stack</cloud>. \
The first line in the template tells AWS that this SAM template that needs to be transformed into a CloudFormation one. then it define a resource, such as a lambda function. the lambda has properties (code location, handler, runtime), policies (what it's allowed to do) and sets up a mapping to to an <cloud>API Gateway endpoint</cloud>.

There is a command line interactive tool for managing SAM applications, it launches a docker container to test, validate and debug the code. [SAM CLI](https://github.com/aws/serverless-application-model). the key commands are:

- `sam init` - initialize a serverless application
- `sam local` - run locally
- `sam validate` - validate a template
- `sam deploy` - deploy the serverless application
- `sam build` - builds the artifact

other aws services (<cloud>CodeBuild</cloud> and <cloud>CodeDeploy</cloud>) integrate with <cloud>SAM</cloud> to automate packaging, testing and safe deployment.

#### Configuring Your Lambda Functions

When creating A lambda, there are three configuration parameters:

- memory - max 10GB of memory allocated to a lambda, with CPU and other resources proportional to it.
- timeout - max of 900 seconds (15 minutes), time until the lambda quits on it own.
- concurrency - how many functions can run at the same time, set at region level for the account (default 1000).

Lambda is billed based on the memory configuration, the number of invocations and the total runtime. the duration is billed in 1 milliseconds increments. setting a timeout can reduce costs, and while higher memory allocations cost more, they can also reduce the runtime (and therefore cost less) - this can be checked with a tool called _aws-lambda-power-tunning_.

there are three concurrency types:

- Unreserved concurrency - concurrency that is not allocated to any specific set of functions. hard limit of minimal 100.
- Reserved concurrency - maximum number of concurrent instances for the function, no other function can use it. no additional costs.
- Provisioned capacity - number of instances which were "warm started" at any given time. provides high performance and low latency, incurs costs.

we might want to limit concurrent lambda invocations (reserved concurrency) to match it with downstream resource that doesn't scale as quickly as lambda functions.

Lambda also have concurrency Burst options, which increase the concurrency limit for all functions in the region based on demand. Concurrency metrics are sent to <cloud>CloudWatch</cloud>, so monitors and alerts can be created.

#### Deploying and Testing Serverless Applications

A traditional (server based) application uses existing machines to hold the application, while a serverless application customizes the resource and instances around itself using a blueprint. <cloud>CloudFormation</cloud> is AWS Infrastructure-as-Code service, a single template can create a stack, which is a collection of resources that are managed as a single unit. Traditional applications are easier to test locally, but serverless applications testing more closely resembles the "real thing".

##### Demo : Using AWS SAM and AWS Cloud9 to Create, Edit, and Deploy a Lambda Function

S3 event triggers a lambda which writes to dynamoDB. using the <cloud>Cloud9</cloud> service, we first click <kbd>Create Environment</kbd> and give it a name and description. the type is an <cloud>EC2</cloud> machine with a <cloud>IAM</cloud> role that can call other services. we can verify the version with the command `sam --version`. we select the <kbd>AWS Explorer</kbd> and choose the region and click on the <kbd>Lambda</kbd> symbol and we can either import, create or deploy a lambda. we click <kbd>Create Lambda SAM Application</kbd> and follow the wizrad and select the <kbd>Node.Js</kbd> runtime, and we use the basic "hello world" option. we can view the "launch.json" and see the launch configuration of the lambda. we can run the application to test it.\
if we wish to generate an s3 event, we can type in the console `sam local generate-event s3 put` and copy the payload and add it into launch file as the testing payload.\
we now add the <cloud>S3</cloud> bucket and the event mapping. we do this by editing the "template.yaml" file and adding it to the resources object, we also add the S3 event to the mapping so the object creation event will trigger the lambda.

```yaml
Resources:
  HelloWorldFunction:
    Type: AWS::Serverless::Function
    Properties:
      # other stuff
      Events:
        S3Event:
          Type: S3
          Properties:
            Bucket:
              Ref: HelloWorldFunctionBucket
            Events: s3:ObjectCreated:*

  HelloWorldFunctionBucket:
    Type: AWS::S3::Bucket
```

we can verify the template with the command `sam validate` in containing folder, this will warn us of errors and indentation problems.

we next create a bucket to store the artifacts, this is done with the explorer and select <kbd>Create S3 Bucket</kbd> and giving it a name. we can deploy the application, store the artifact in the new bucket, and then the resources are actually created in our aws environment. we can now upload a file to the application bucket ("HelloWorldFunctionBucket") and see that our lambda was invoked.

our next step is expanding the logic to have the lambda write to a <cloud>dynamoDb</cloud>. we create the table as a resource in the template file, and we grant the lambda permissions to write to the table.

```yaml
Resources:
  HelloWorldFunction:
    Type: AWS::Serverless::Function
    Properties:
      # other stuff
      Events:
        S3Event:
          Type: S3
          Properties:
            Bucket:
              Ref: HelloWorldFunctionBucket
            Events: s3:ObjectCreated:*
      Policies:
        - DynamoDBCrudPolicy:
            TableName: !Ref HelloWorldFunctionTable

  HelloWorldFunctionBucket:
    Type: AWS::S3::Bucket
  HelloWorldFunctionTable:
    Type: AWS::Serverless::SimpleTable
```

if we deploy the application now, we can see the newly created table. we can now modify the code to write the information

```js
console.log("Loading Function");
const aws = require("aws-sdk");
const ddb = new aws.DynamoDB({ apiVersion: "2012-08-10" });
exports.handler = async (event, context) => {
    // boiler plate code before here

    const s3ObjectKey = event.Records[0].s3.object.key;
    const s3TimeStamp = event.Records[0].eventTime;
    const params = {
      TableName: '' // copy the name
      Item: {
        'id': {S: s3ObjectKey},
        'timestap': {S: s3TimeStamp}
      }
    };
    await ddb.putItem(params, function(err, data) {
      if(err) {
        console.log("Error", err);
      } else {
        console.log("Success");
      }
    }).promise();
    return response;
};
```

we can now deploy again and test by uploading a file to the bucket and seeing it created in the dynamodb table.

---

a drawback of serverless applications is that they can go live without proper testing, this risk can be mitigated by employing versioning and rollbacks to ensure safe deployments. when we publish a lambda, it gets a unique incremental number, and is also published with the tag of "LATEST". we can use the unique number to always refer to that version in an <cloud>ARN</cloud>, and we can create specific aliases such as "DEV", "BETA" and "PRODUCTION" and have them refer to specific versions.

with aliasing in place, we can test using routing, sending part of the traffic to the existing version and part of it to a new version.

> You can point an alias to a maximum of two Lambda function versions. The versions must meet the following criteria:
>
> - Both versions must have the same runtime role.
> - Both versions must have the same dead-letter queue configuration, or no dead-letter queue configuration.
> - Both versions must be published. The alias cannot point to $LATEST.

we can also integrate this with <cloud>CodeDeploy</cloud>, which offers traffic shifting strategies:

> - Canary – Traffic is shifted in two increments. If the first increment is successful, the second is completed based on the time specified in the deployment.
> - Linear – With linear traffic shifting, traffic is slowly shifted in a predetermined percentage every X minutes based on how you have it configured.
> - All-at-once – Shifts all traffic from the original Lambda function to the updated Lambda function version at once.

<cloud>CodeDeploy</cloud> also has alarms and hooks for extra control. this options can be set up in the SAM template

```yaml
Resources:
  HelloWorldFunction:
    Type: AWS::Serverless::Function
    Properties:
      # other stuff
      Events:
        S3Event:
          Type: S3
          Properties:
            Bucket:
              Ref: HelloWorldFunctionBucket
            Events: s3:ObjectCreated:*
      Policies:
        - DynamoDBCrudPolicy:
            TableName: !Ref HelloWorldFunctionTable
      AutoPublishAlias: live
      DeploymentPreference:
        Type: Canary10Percent10Minutes
        Alarms:
          - !Ref AliasErrorMetricGreaterThanZeroAlarm
          - !Ref LaterstVersionErrorMetricGreaterThanZeroAlarm
        Hooks:
          PreTraffic: !Ref PreTrafficLambdaFunction
          PostTraffic: !Ref PostTrafficLambdaFunction
```

#### Monitoring and Troubleshooting

AWS Lambda integrates with monitoring services to monitor and troubleshoot functions. for <cloud>CloudWatch</cloud>, it automatically tracks and displays:

> - Invocations - The number of times your function code is run, including successful runs and runs that result in a function error. If the invocation request is throttled or otherwise resulted in an invocation error, invocations aren't recorded.
> - Duration - The amount of time that your function code spends processing an event. The billed duration for an invocation is the value of Duration rounded up to the nearest millisecond.
> - Errors -The number of invocations that result in a function error. Function errors include exceptions thrown by your code and exceptions thrown by the Lambda runtime. The runtime returns errors for issues such as timeouts and configuration errors.
> - Throttles - The number of times that a process failed because of concurrency limits. When all function instances are processing requests and no concurrency is available to scale up, Lambda rejects additional requests.
> - IteratorAge - Pertains to event source mappings that read from streams. The age of the last record in the event. The age is the amount of time between when the stream receives the record and when the event source mapping sends the event to the function.
> - DeadLetterErrors - For asynchronous invocation, this is the number of times Lambda attempts to send an event to a dead-letter queue but fails.
> - ConcurrentExecutions- The number of function instances that are processing events. \
>   You can also view metrics for the following:
>   - UnreservedConcurrentExecutions – The number of events that are being processed by functions that don't have reserved concurrency.
>   - ProvisionedConcurrentExecutions – The number of function instances that are processing events on provisioned concurrency. For each invocation of an alias or version with provisioned concurrency, Lambda emits the current count.

Another Option is using the <cloud>CloudWatch Lambda Insights</cloud> for system level metrics such as cold starts, worker shutdown and other diagnostic data. this is done by adding a lambda layers that collects the metrics and emits the data as a single log using the embedded metric formatting. Each lambda has it's own dashboard, and there is an aggregate dashboard for all functions.

<cloud>AWS X-Ray</cloud> is a stool that visualizes data flow between services and can identify performance bottlenecks and show timelines for cold starts and function initializations. <cloud>CloudTrail</cloud> audits api actions, and a dead-letter-queue (usually <cloud>SQS</cloud> or <cloud>SNS</cloud>) can be used to capture events that must be handled even if they failed at the first time.

#### Quiz: AWS Lambda

- Q: Which of these statements describe a resource policy? (Select THREE.)
- A: Can give Amazon S3 permission to initiate a Lambda function, Determines who has access to invoke the function, Can grant access to the Lambda function across AWS accounts.
- Q: Which monitoring tool provides the ability to visualize the components of an application and the flow of API calls?
- A: <cloud>AWS X-Ray</cloud>
- Q: Which capabilities are features of Lambda? (Select THREE.)
- A: Triggers Lambda functions on your behalf in response to events, Runs code without you provisioning or managing servers, Scales automatically.
- Q: What is the importance of the IAM execution role?
- A: Gives your function permissions to interact with other services.
- Q: What are the reasons for setting a concurrency limit (or reserve) on a function? (Select THREE.)
- A: Match the limit with a downstream resource, Manage costs, Regulate how long it takes to process a batch of events.
- Q: Which patterns are Lambda invocation models? (Select THREE.)
- A: Synchronous, Asynchronous, Polling.
- Q: What does an AWS Identity and Access Management (IAM) resource-based policy control?
- A: Permissions to Invoke the function
- Q: Match the terms on the left with the appropriate definition.
- A: Producer: Create events with all required information, Router:Ingests and filters events using rules, Consumer: Subscribe and are notified when events occur.
- Q: Which feature can a developer enable to create a copy of a function for testing?
- A: ~~Aliases~~ Versioning.

</details>

### Amazon API Gateway for Serverless Applications

<details>
<summary>
API Gateway, types, endpoints, access controls and usage limits.
</summary>

> In this course, you will learn how API Gateway lets you define and deploy application programming interfaces (APIs) at scale, and why it makes a great front door to your AWS Lambda functions and backend APIs. You will learn what you need to know to plan for, launch, and use API Gateway for your serverless applications and how to use it to decouple a monolithic application. You will learn to analyze API Gateway traffic and identify opportunities or improvements, validations, responses, and mapping.

#### Introduction to API Gateway

> APIs are mechanisms that facilitate two software components communicating with each other. APIs act as the front door for applications to access data, business logic, or functionality from backend services.

however, there are some common challenges for using apis:

> - Handling API calls in a serverless application
> - Working with multiple API versions and environments
> - Controlling access and authorization
> - Managing traffic spikes
> - Monitoring third-party access

we can manage these challenges by employing an API management tool, which acts as a front door to the APIs themselves, <cloud>API Gateway</cloud> is one such option. an API gateway is a service that handles creation, publishing, maintenance, monitoring and security of of apis at scale. this includes handling traffic management, Cross Origin Resource Sharing (CORS), authorization and access control, throtelling and api versioning.

we can divide the features of the api gateway based on categories:

- Developer Features:
  - Running Multiple versions of the API at the same time - either for AB testing or for different users.
  - Quick SDK generation - ability to generate SDKs for external software (using `get-sdk` from the command line tool)
  - Transform and Validate requests and responses - another layer over the data before it reaches the backend.
- Management Features:
  - Reducing latency - taking advantage of <cloud>CloudFront</cloud> edge locations
  - Throtelling traffic - disallow all traffic from reaching the backend and reduce traffic spikes when not authorized
  - Flexible authorization options - control access with multiple authorization options, such as <cloud>IAM</cloud>, <cloud>Amazon Cognito</cloud>, OAuth Token and others
  - Manage API keys for third art developers - another wat for fine grained control and tracking usage.

API gateways support different kinds of APIs

- HTTP - lower latency and lower costs, less management functionality
- RESTful - stateless, common verbs (`GET`, `PUT`, `POST`, `DELETE`), api management functionalities and API proxy
- WebSocket - bidirectional communication, used for real-time applications, maintain persistent connection with the client.

| REST API                                                                                        | HTTP API                                                                          | WebSocket API                                                                                                                         |
| ----------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| All-inclusive set of features needed to build, manage, and publish APIs at a single price point | Building modern APIs that are equipped with native OIDC and OAuth 2 authorization | Bidirectional communication that lets clients and services independently send messages to each other                                  |
| Building APIs that use certificates for backend authentication, AWS WAF, or resource policies   | Building proxy APIs for Lambda or any HTTP endpoint                               | Richer client-to-service interactions because services can push data to clients without requiring clients to make an explicit request |
| Workloads that need an edge-optimized or private API type                                       | APIs for latency-sensitive workloads                                              | APIs for real-time communication                                                                                                      |

> REST APIs are intended for APIs that require API proxy functionality and API management features in a single solution. HTTP APIs are optimized for building APIs that proxy to Lambda functions or HTTP backends, making them ideal for serverless workloads. HTTP APIs are a cheaper and faster alternative to REST APIs, but they do not currently support API management functionality. Unlike a REST API, which receives and responds to requests, a WebSocket API supports two-way communication between client apps and your backend. The backend can send callback messages to connected clients.

#### Designing WebSocket APIs

> In a WebSocket API, the client and server can send messages to each other at any time. With a WebSocket connection, your backend servers can push data to connected users and devices, avoiding the need to implement complex polling mechanisms.

a WebSocket API has a persistent state, and can handle messages from the same client, routing each message to a different aws service (<cloud>Lambda</cloud>, <cloud>DynamoDB</cloud> or an HTTP endpoint). the bi-directional nature also means that the server can send messages back to the client, so the client doesn't have to make explicit requests or "poll" for data.

for WebSocket API, costs are based only for when the APIs are in use, with three different aspects:

- number of messages, with each message being up to 128Kb in size.
- total connection minutes
- additional charges for using other AWS services or transferring data

##### Creating WebSocket APIs

to create an WebSocket API, we need to have at least one route, one integration and one "stage". the route is the expected value inside a message, which is evaluated by a route selection expression and is used to route the data properly. the integration is the selected service to handle the message, and the stage defines the path through which the deployment is available.

The client communicates withe the websocket using json messages. the handling of the message is based on the content of the message.routes have the dollar sign prefix. there are three pre-defined routes:

- "connect" - triggered at first connection
- "disconnect" - triggered when client disconnects (best effort)
- "default" - any non-defined route, or can be a fallback route, or to handle non-json messages.

other custom routes can evaluate specific key-value pairs and match them in the message json.

For each route, we specify an integration - which endpoint on the backend handles it. both the request and response can be transformed before and after the backend service handles them. if there is an integration response, then the communication is bi-directional, otherwise it's one-way communication.

WebSocket APIs have a selection expression, which evaluate the request or the response and match a key.

a connection with a WebSocket isn't established until the "connect" route integration is successfully invoked.

#### Designing REST APIs

> a REST API in API Gateway is a collection of resources and methods that are integrated with backend HTTP endpoints, Lambda functions, or other AWS services. API Gateway REST APIs use a request-response model, where a client sends a request to a service and the service responds back synchronously. This kind of model is suitable for many different kinds of applications that depend on synchronous communication.

##### REST APIs endpoints

The RESF API can connect to different endpoints. this can be changed (mostly) after creation.

> Regional endpoint:\
> The regional endpoint is designed to reduce latency when calls are made from the same AWS Region as the API. In this model, API Gateway does not deploy its own CloudFront distribution in front of your API. Instead, traffic destined for your API will be directed straight at the API endpoint in the Region where you’ve deployed it.\
> This endpoint type gives you lower latency for applications that are invoking your API from within the same Region (for example, an API that is going to be accessed from EC2 instances within the same Region).
>
> Edge-optimized endpoint:\
> The edge-optimized endpoint is designed to help you reduce client latency from anywhere on the internet. If you choose an edge-optimized endpoint, API Gateway will automatically configure a fully managed CloudFront distribution to provide lower latency access to your API.\
> This endpoint-type setup reduces your first hit latency for your API. An additional benefit of using a managed CloudFront distribution is that you don’t have to pay for or manage a CDN separately from API Gateway.
>
> Private endpoint:\
> The private endpoint is designed to expose APIs only inside your selected Amazon Virtual Private Cloud (Amazon VPC). This endpoint type is still managed by API Gateway, but requests are only routable and can only originate from within a single virtual private cloud (VPC) that you control.

##### REST API caching

> You can turn on API caching in API Gateway to cache your endpoint's responses. With caching, you can reduce the number of calls made to your endpoint and also improve the latency of requests to your API.

Caching allows to store a respons on the api gateway itself, rather than calling the backend service again and again. the response is stored for a specified time based on the TTL period. this allows for reduced latency and reduced load on the backend servers themselves.\
we can configure the cache overall size, the TTL (in seconds), configure encryption of the data inside the cache, and override configuration for specific methods.\
The relevant <cloud>CloudWatch</cloud> metrics are "CacheHitCount" and "CacheMissCount".

---

Pricing for REST APIs is based on

- number of requests
- data transfer out (standard aws costs), private API endpoints don't have data transfer costs, but <cloud>PrivateLink</cloud> charges apply
- additional costs for API cache, if used. hourly rate based on size.

#### Building and Deploying APIs with API Gateway

the base API invoke call is an URL

"restAPI_id.execute-API.region.amazon.com/stage/[resource or resource path]"

the id is generated by aws, the region being where the gateway is located, a stage represents specific version of the API, and the final part is the resource path. the url can be customized by using a custom domain name as the host, and having a mapping from the base path to the true URL.

##### Building The API Gateway

in the web console, we first choose the API type (HTTP, REST, Websocket), and then continue with the wizard. we select name the API and select the endpoint type. we can then choose <kbd>Actions</kbd> and <kbd>Create Resource</kbd> to add a resource url. this is a path relative to the base URL, and we can specify path paremeters here. it is also possible to specify a proxy resource, which has a special verb `ANY` . a lambda can also be a proxy option. for each resource, we can click <kbd>CREATE METHOD</kbd> and add options based on the HTTP verb. we can configure values for each method such as timeout and integration type.

Integration types determine how the data from the request is passed to the backend.

> - Lambda Function
>   - When you are using API Gateway as the gateway to a Lambda function, you’ll use the Lambda integration. This will result in requests being proxy-ed to Lambda with request details available to your function handler in the event parameter, supporting a streamlined integration setup. The setup can evolve with the backend without requiring you to tear down the existing setup.
>   - For integrations with Lambda functions, you will need to set an IAM role with required permissions for API Gateway to call the backend on your behalf.
> - HTTP Endpoint
>   - HTTP integration endpoints are useful for public web applications where you want clients to interact with the endpoint. This type of integration lets an API expose HTTP endpoints in the backend.
>   - When the proxy is not configured, you’ll need to configure both the integration request and the integration response, and set up necessary data mappings between the method request-response and the integration request-response.
>   - If the proxy option is used, you don’t set the integration request or the integration response. API Gateway passes the incoming request from the client to the HTTP endpoint and passes the outgoing response from the HTTP endpoint to the client.
> - AWS Service
>   - AWS Service is an integration type that lets an API expose AWS service actions. For example, you might drop a message directly into an Amazon Simple Queue Service (Amazon SQS) queue.
> - Mock
>   - Mock lets API Gateway return a response without sending the request further to the backend. This is a good idea for a health check endpoint to test your API. Anytime you want a hardcoded response to your API call, use a Mock integration.
> - VPC Link
>   - With VPC Link, you can connect to a Network Load Balancer to get something in your private VPC. For example, consider an endpoint on your EC2 instance that’s not public. API Gateway can’t access it unless you use the VPC link and you have to have a Network Load Balancer on your backend.
>   - For an API developer, a VPC Link is functionally equivalent to an integration endpoint.

with the endpoint being set, it's time to add more details, this can be custom header parameters, query strings or other transformations. it's also possible to test API methods - this actually runs the API, but it doesn't store cloudWatch logs for the API gateway.

##### Gateway Stages

> A stage is a snapshot of the API and represents a unique identifier for a version of a deployed API.\
> With stages, you can have multiple versions and roll back versions. Anytime you update anything about the API, you need to redeploy it to an existing stage or to a new stage that you create as part of the deploy action.

as mentioned before, stages can have different configurations (caching, throtelling, usage plans), and each stage can be exported to an sdk (or swagger or postman extension). stages can be based on environment (production, dev, beta, etc...), customer, or simple versioning. each stage can define different variables in it - namely the url and the lambdaFn (function name). using stages allows to have dynamic routing for the same API, so we can keep the number of APIs small, and set the endpoints differently.

##### Best Practices for API Gateway

> Use API Gateway stages with Lambda aliases\
> To highlight something that was mentioned in the previous example, Lambda and API Gateway are both designed to support flexible use of versions. You can do this by using aliases in Lambda and stages in API Gateway. When you couple that with stage variables, you don't have to hard-code components, which leads to having a smooth and safe deployment.
>
> - In Lambda, enable **versioning** and use **aliases** to reference.
> - In API Gateway, use **stages** for environments.
> - Point API Gateway **stage variables** at the Lambda **aliases**.
>
> Use Canary deployments\
> With Canary deployments, you can send a percentage of traffic to your "canary" while leaving the bulk of your traffic on a known good version of your API until the new version has been verified. API Gateway makes a base version available and updated versions of the API on the same stage. This way, you can introduce new features in the same environment for the base version.\
> To set up a Canary deployment through the console, select a stage and then select the Canary tab for that stage.
>
> Use <cloud>AWS SAM</cloud> to simplify deployments\
> One of the challenges of serverless deployments is the need to provide all the details of your deployment environment as part of your deployment package. The <cloud>AWS Serverless Application Model (AWS SAM)</cloud> is an open-source framework that you can use to build serverless applications on AWS.

SAM templates are transformed into <cloud>CloudFormation</cloud> templates, and are a simple, clean and straight forward way to define applications.

#### Managing API Access

> The next important step for you to consider after you have successfully designed and deployed your API is how you will manage access and authorization for the API. API Gateway provides you with multiple, customizable options for:
>
> - Authorizing an entity to access your APIs
> - Providing more granular control
> - Controlling the amount of access through throttling

| Option                    | Authentication | Authorization | Signature V4 | Cognito User Pools | Third-Party Auth | Multiple Header Support | Additional Costs                       |
| ------------------------- | -------------- | ------------- | ------------ | ------------------ | ---------------- | ----------------------- | -------------------------------------- |
| AWS IAM                   | Yes            | Yes           | Yes          | No                 | No               | No                      | None                                   |
| Lambda Authorizer Token   | Yes            | Yes           | No           | Yes                | Yes              | No                      | Pay per authorizer invoke              |
| Lambda Authorizer Request | Yes            | Yes           | No           | Yes                | Yes              | Yes                     | Pay per authorizer invoke              |
| Amazon Cognito            | Yes            | Yes           | No           | Yes                | No               | No                      | Pay based on your monthly active users |

> As shown in the comparison table, there are three main ways to authorize API calls to your API Gateway endpoints:
>
> 1. Use IAM and Signature version 4 (also known as Sig v4) to authenticate and authorize entities to access your APIs.
> 2. Use Lambda Authorizers, which you can use to support bearer token authentication strategies such as OAuth or SAML.
> 3. Use Amazon Cognito with user pools

Authorizing <cloud>IAM</cloud> is suited when using internal services or when there just a few customers. especially if they are already using IAM roles. with IAM, all requests are signed with a Sig v4 credentials, which you get from the IAM service and attach to the request authorization header. inside the gateway, the key is parsed and the user is checked for appropriate permissions. if it doesn't match, then the request is denied at the gateway level.

if the user already has an OAuth strategy, then Lambda Authorizers can be used. when a lambda uses Lambda Authorizers, the request is sent to an authorizer lambda, which checks it for a token or a header and returns a policy. this can then be cached in the Gateway to reduce subsequent calls to the authorizer lambda.\
we can create an authorizing lambda from the "api-gateway-authorizer" blueprint.the authorizing function can use either a token or a request. a token can be something such as OAuth or bearer, a request can contain more data and can be fine-grained to control which resources and actions are allowed per stage for the method.

<cloud>AWS Cognito</cloud> (and users pools) are another option, Coginto user pools are apis that can be integrated into the application to provide authentication, this works for mobile and web applications where authentication is handled inside the application. it also allows for fine grained access control and scopes.

##### Throttling and Usage Plans

As hinted above, we can use API Gateways to manage the volume of api calls. these limits can help by preventing a single customer from using all the backend resources, and can protect backend resources which aren't easily scalable from usage spikes that could crush them.

one way to do this is by including customer specific api-keys in the request header ("x-API-key"). with this key, it becomes possible to:

1. throttle calls - limit the number of calls per second (and burst)
2. set quota - limit the overall number of calls by day, week and month
3. track usage - monitor usage for each customer

the limits work by using a token bucket algorithm. this can be thought of as queue with a maximum size determined by the "burst" property, where each token is an element in the queue. there are default limits per account and region, but they can fine-tuned with usage plans.

the hierarchy is applied in the following order

1. per client, per method limits
2. per client limits
3. default and custom per-method limits
4. account level limits

##### IAM Permissions in API Gateway

there are two types of permissions used in the API gateway. the first type is `apigateway:*`, which controls who can manage the gateway (update, configure, delete and view) itself. the second are `execute-api:*` permissions, which control who can invoke the API. more granular controls can be configured via <cloud>Resource Polices</cloud> that are attached to the API and can limit access based on account, ip address range, vpc and vpc endpoint.

> Resource policy and authentication methods work together to grant access to your APIs. As illustrated below, methods for securing your APIs work in aggregate.
>
> - API Gateway resource policy only - Explicit allow is required on the inbound criteria of the caller. If not found, deny the caller.
> - Lambda Authorizer and resource policy - If the policy has explicit denials, the caller is denied access immediately. Otherwise, the Lambda Authorizer is called and returns a policy document that is evaluated with the resource policy.
> - IAM authentication and resource policy - If the user authenticates successfully with IAM, policies attached to the IAM user and resource policy are evaluated together.
>   - If the caller and API owner are from separate accounts, both the IAM user policies and the resource policy explicitly allow the caller to proceed.
>   - If the caller and the API owner are in the same account, either user policies or the resource policy must explicitly allow the caller to proceed.
> - Cognito authentication and resource policy - If API Gateway authenticates the caller from Cognito, the resource policy is evaluated independently.
>   - If there is an explicit allow, the caller proceeds.
>   - Otherwise, deny or neither allow nor deny will result in a deny.

#### Monitoring and Troubleshooting API Gateway

##### CloudWatch Metrics

> After your APIs are deployed, you can use CloudWatch Metrics to monitor performance of deployed APIs. API Gateway has seven default metrics out of the box.
>
> - Count: Total number of API requests in a period
> - Latency: Time between when API Gateway receives a request from a client and when it returns a response to the client; this includes the integration latency and other API Gateway overhead.
> - IntegrationLatency: Time between when API Gateway relays a request to the backend and when it receives a response from the backend.
> - 4xxError: Client-side errors captured in a specified period.
> - 5xxError: Server-side errors captured in a specified period.
> - CacheHitCount: Number of requests served from the API cache in a given period.
> - CacheMissCount: Number of requests served from the backend in a given period, when API caching is turned on.

the overhead of the gateway API is the difference between the **Latency** (complete round trip) and the **IntegrationLatency** metric (from the gateway to the backend service and back). looking at this value can help identify bottlenecks.\
$Latency - IntegrationLatency = Gateway Overhead$

##### CloudWatch Logs

API Gateways store two kinds of logs in cloudWatch: execution and access logging.

> - The first type is **execution logging**, which logs what’s happening on the roundtrip of a request. You can see all the details from when the request was made, the other request parameters, everything that happened between the requests, and what happened when API Gateway returned the results to the client that’s calling the service.
>   - Execution logs can be useful to troubleshoot APIs, but can result in logging sensitive data. Because of this, it is recommended you don't enable Log full requests/responses data for production APIs.
>   - In addition, there is a cost component associated with logging your APIs.
> - The second type is **access logging**, which provides details about who's invoking your API. This includes everything including IP address, the method used, the user protocol, and the agent that's invoking your API.
>   - Access logging is fully customizable using JSON formatting. If you need to, you can publish them to a third-party resource to help you analyze them.

A possible scenario to use execution logs is when there is a spike in 4xx errors, in which case we can open the logs and search for something such as "Key throttle limit exceeded" and configure better throttling limits.

##### Monitoring with X-Ray and CloudTrail

> - You can use <cloud>AWS X-Ray</cloud> to trace and analyze user requests as they travel through your Amazon API Gateway APIs to the underlying services. With X-Ray, you can understand how your application is performing to identify and troubleshoot the root cause of performance issues and errors.
>   - X-Ray gives you an end-to-end view of an entire request, so you can analyze latencies and errors in your APIs and their backend services.
>   - You can also configure sampling rules to tell X-Ray which requests to record, and at what sampling rates, according to criteria that you specify.
> - <cloud>AWS CloudTrail</cloud> captures all API calls for API Gateway as events, including calls from the API Gateway console and from code calls to your API Gateway APIs.
>   - Using the information collected by CloudTrail, you can determine the request that was made to API Gateway, the IP address from which the request was made, who made the request, when it was made, and additional details.
>   - You can view the most recent events in the CloudTrail console in Event history.

<cloud>AWS X-Ray</cloud> give a complete view over "round-trips" in AWS, so we can see information for the gateway portion and the lambda portion of the same requests in a single page (such as lambda cold and warm starts).

example of investigating an "access denied" error using <cloud>X-Ray</cloud>

#### Data Mapping and Request Validation

> In API Gateway, an API's method request can take a payload in a different format from the corresponding integration request payload as required by your backend and the reverse.\
> Mapping templates can be added to the integration request to transform the incoming request to the format required by the backend of the application or to transform the backend payload to the format required by the method response.

we can transform the request before it reaches the backend, or the response before sending it back to the customer, this is done via mapping. this is also where we can add stage variables to the request if needed.

The API gateway can also control error responses, such as having a custom response to the customer (rahter than forewaring lambda error response) or modifying it. this can be done based on the http response code.

API gateways can also take care of some basic request validations, rather than performing them in the backend.

#### Quiz: API Gateway Knowledge check

- Q: Which of these endpoint types are available for REST APIs for Amazon API Gateway?
- A: Regional, Private and Edge-optimized.
- Q: Which of these are **not** a use case for WebSocket APIs?
- A: Application that needs to cache using API Gateway
- Q: Which of these tools can be used to monitor and log your APIs?
- A: <cloud>CloudWatch</cloud>, <cloud>CloudTrail</cloud> and <cloud>X-Ray</cloud>
- Q: Which of these are WebSocket predefined routes in Amazon API Gateway?
- A: "connect", "disconnect" and "default" routes
- Q: Which of these accurately describe a stage in Amazon API Gateway?
- A: A snapshot of the API and a unique identifier for a version of a deployed API

---

> - You learned about the challenges of API management and how API Gateway helps alleviate these challenges.
> - You learned about the differences between REST and WebSocket APIs, and the different features API Gateway provides for each.
> - You learned to build and deploy APIs into API Gateway including the anatomy of these APIs, the integration types, and how to test them.
> - You learned about the different authorization and authentication options that can be used for managing API access and how to manage the volume of API calls that are processed through your API endpoint.
> - You reviewed the different CloudWatch Metrics for API Gateway and learned how to use these metrics to troubleshoot and monitor your APIs.
> - You learned how to transform data using mapping templates and how to handle errors in your gateway responses.

</details>

### Amazon DynamoDB for Serverless Architectures

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

#### Introduction to Amazon DynamoDB for Serverless Architectures

> DynamoDB is a serverless, fully managed NoSQL (non-relational) database service designed for Online Transactional Processing (OLTP) workloads.

(video)

<cloud>DynamoDb</cloud> is a NoSQL database, relational databases were created when storage was expensive, and they used rigid schemas to store a lot of data. queries then used relational connections between tables to answer questions about the data, but complex queries could require a lot of computing power. Non-relational databases, such as <cloud>DynamoDB</cloud> take advantage of the decreased cost of storage, and store related data locally to reduce CPU usage. this also allows for data to be stored without forcing it to a specific schema.\
<cloud>DynamoDb</cloud> itself is a serverless service, fully managed by AWS and easy to integrate with other AWS services.

We usually have two kinds of queries: transactional and analytical.

- Online Transactional Processing (OLTP)
- Online Analytical Processing (OLAP)

<cloud>DynamoDb</cloud> serves for transaction queries, with known request patterns. live operations, active connections, quick and precise with small objects. more in-depth queries (which don't demand the answer right now) are more fit for OLAP and different databases.

> - Flexible Schema
> - JSON document or key-value data structures
> - Supports event-driven programming
> - Accessible via AWS Management Console, CLI, and SDK
> - Availability, durability, and scalability built-in
> - Scales horizontally
> - Provides fine-grained access control
> - Integrates with other AWS services

DatabaseBases in AWS:

- <cloud>RDS</cloud> - relational, transactional queries, supports multiple engines(including <cloud>Amazon Aurora</cloud>)
- <cloud>DynamoDb</cloud> - non relational, transactional queries, document storage and key-value store models.
- <cloud>Redshift</cloud> - relational, analytical, fully managed date warehouse.
- <cloud>Neptune</cloud> - non-relation, analytical, highly connected datasets (graphs)
- <cloud>ElasticCache</cloud> - in memory data cache, supports redis and memcached engines.

#### How Amazon DynamoDB Works

(video)

data is stored in tables, a table stores items, each item has attributes, there is a required 'partiton key' attribute, and an optional 'sort key' attribute. the combination of the partition key and the sort key makes up the primary key, which must be unique. items can have additional attributes, but those are flexible and not required.\
When tables require scaling, they use a technique called "sharding", which stores items based on the hash of the partition key. each partition is stored on a different server (this happens automatically). attributes can be:

- numeric
- boolean (true or false)
- string (text)
- binary (base64 encoded)
- null
- sets of numbers, strings or binaries (unordered)

if the data uses json, we can store it as document model (map, list). partition and sort-key can't be mapped.

data is stored in several copies across the region (different Availability Zones), this replication allows for 99.99% reliability. there two forms of consistency: eventual consistency (the default) and strong consistency. strong consistency reads always read from the most recently updated replication, so it increases the load and can lead to some unexpected behavior. we can set the throughput by specifying Read and Write capacity units.

- 1 read capacity unit (RCU) = 1 item(4kb or less) read per second
- 1 write capacity unit (WCU) = 1 item(1kb or less) write per second. updating an item is the same as writing. Deleting is also a 'write' request.

dynamoDb also has burst capacity, which adapts for imbalanced workloads. strong consistency reads take twice as much RCU than eventual consistency reads.

basic requests:

- `PutItem` - write item to specified primary key
- `UpdateItem` - change attributes for item of specified primary key
- `BatchWriteItem` - write a bunch of items to specified primary keys
- `Delete` - remove item from specified primary key
- `GetItem` - retrieve item associated with a specified primary key
- `BatchGetItems` - retrieve items associated with a number of primary keys
- `Query` - for a specified partition key, retrieve item  matching sort key expressions (forward/reverse order)
- `Scan` - give me every item in my table

Queries are more cost efficient.

DynamoDb supports secondary indexes, which allow to query Data based on attributes other than the primary key. this also supports projection.

- Local Secondary index (LSI) - local to a specific partition key, share the base table capacity, this limit the data to 10GB per partition. can't be modified after the table is created.
- Global Secondary index (GSI) - not local to a partition key, can span across all items in the table, so it's an alternative partition of the original table. they have an independent provision throughput. but they don't offer strong consistent reads. they can be created and deleted at any time. they don't require unique primary keys.

we can define up to 5 globally secondary indexes and up to 5 locally secondary indexes for each table.

there also dynamoDb streams, which logs all operations on the table.

(demo video)

we click <kbd>Create Table</kbd> to create a new table with a partition and sort key (creating primary key). we can add items to the table, the items must have the primary key attributes, and we can add other attributes if we want. we can either <kbd>Query</kbd> or <kbd>Scan</kbd> the table, queries are faster and costless, but require passing a partition key. scans are slower and more costly, but can filter all the data in the table.

#### Operating Amazon DynamoDB

(video)

configuring the clients to handle errors. errors can be because the service is down, or because the provisioned throughput was exceeded (throttling). the AWS SDK has built-in retry logic, with other configurable options. also need to handle retires for batch requests.\
DynamoDb has auto scaling for provisioned throughputs, separably for read and write.\
there is also support for globally replacted tables, which offer durability and faster operations in multiple regions. using this requires some more configuration to direct clients to the same replica for reads and writes.\
Items can expire by using TTL - automatic deletion with time to live, this is an epoch field. expiring items can be moved to S3 by reading the dynamodb stream.\
we can control access to the table with <cloud>IAM</cloud> roles and limit permissions, we can also set a VPC endpoint.

DAX (DynamoDB Accelerator) is an in-memory acceleration way to reduce load and latency for eventually consistent reads. tables are backed up either by demand or with a rolling 35 days windows.

monitoring:

- aws error code returned from operations
- logging api calls with <cloud>CloudTrail</cloud>
- setting up <cloud>CloudWatch</cloud> alarms for performance metrics

**continue later**

#### Design Considerations

#### Serverless Architecture Patterns

#### DynamoDb Assessment

</details>

### Seperator

<!-- end of serveless -->

## CI-CD Services

<details>
<summary>
Courses Focusing on CI/CD and deployment in AWS.
</summary>

### Build and Deploy APIs with a Serverless CI/CD

<details>
<summary>
Why use serverless and how to create a pipeline.
</summary>

> Building an API engine, managing a CI/CD pipeline: achieving these DevOps goals historically took managing a number of instances with all the associated operational overhead.\
> You’ll start by understanding how APIs are currently managed with traditional methods, then learn the real-world, best practices of how serverless application methods (SAM) can streamline your operations.

(single video)

Agenda:

- API
- Traditional Applications
- Serverless Applications
- Develop, Debug and Deploy the SAM way
- Serverless CI/CD pipeline
- Lessons Learn and best practices

API - application program interface, the contract that the program exposes outside. we usually use it in a RESTful way. for example, a node.js (express js) project has an expressJS app that acts as a router to the endpoints, which then perform actions on the persistent data in some database. this is simple enough, but to get it to production level, it needs to be deployed, scaled, set up load balancing, etc...

> What is Serverless?
>
> - No server management
> - Flexible scaling
> - High availability by default
> - Pay as you go model

most data centers usually run in under-utilization, people pay for compute power that is only rarely used to it's fullest.

AWS serverless application use <cloud>api gateway</cloud> and <cloud>Lambda</cloud> functions. using an external gateway creates a distinction between the inside and outside of the business logic, and allows controlling the routing without effecting the functionality itself.

> What is AWS SAM?
>
> - <cloud>Serverless Application Model</cloud>.
> - Simplifies <cloud>CloudFormation</cloud> configuration to define serverless applications.
> - Support configurations of <cloud>Lambda</cloud>, <cloud>API Gateway</cloud>, <cloud>DynamoDb</cloud> and event sources.
> - SAM local - local testing

demo of deploying to AWS Lambda using SAM:

packaging and deploying the application using the `sam package` command. the SAM template is identified by the "Transform: AWS::Serverless-2016-10-31" field. it has sections for parametes, resources and outputs.

> Serverless Best Practices - Do's:
>
> - 12 factor applications.
> - Ephemeral (stateless, don't assume storage).
> - Instantiate expensive objects outside event handler (minimize warm start time).
> - Local dev/test for fast iteration.
> - End-to-End integration testing/profiling early on.
> - Profile your application for bottlenecks.
>
> Pitfalls to Avoid - Don't:
>
> - Don't architect differently for serverless.
> - Avoid significant logic in the API Gateway layer.
> - Don't Implement worklows in APIs (use step functions for this use-case instead)
>
> AWS Based CI/CD services
>
> - <cloud>CodeCommit</cloud> - fully managed git source control service.
> - <cloud>CodeBuild</cloud> - fully managed build service.
> - <cloud>CodeDeploy</cloud> - automates software deployment.
> - <cloud>CodePipeline</cloud> - orachestrates build, test and deployment.

unlike Jenkins, there is no server management with AWS CI/CD services.

> Demo of a ci/cd pipeline with four phases
>
> 1. Source: Checks out source code from codeCommit.
> 1. Build: Lint, runs unit tests and packages.
> 1. Staging: Deploy to staging, run integration tests.
> 1. Live: Push to production (gradual rollout).

CI/CD Best Practices

- Security
- Logging
- Notification
- Keep tests fast - fail fast
- Enforce all environment updates to be released via pipeline
- Build once, push through code pipelines
- Use Short lived feature branches
- Run all tests locally first
- Only build what has changed
- Version Control everything

</details>

### Getting Started with DevOps on AWS

<details>
<summary>
DevOps services in AWS
</summary>

> This beginner-level course is for technical learners in the development and operations domains who are interested in learning the basic concepts of DevOps on Amazon Web Services (AWS). Using discussions, interactive content, and demonstrations, you will learn about culture, practices, and tools used in a DevOps environment. You will also explore concepts for developing and delivering secure applications at high velocity on AWS.

#### Introduction to DevOps

<details>
<summary>
What DevOps is
</summary>

> DevOps is the combination of cultural philosophies, practices, and tools that increases an organization’s ability to deliver applications and services at high velocity: evolving and improving products at a faster pace than organizations using traditional software development and infrastructure management processes.\
> This speed enables organizations to better serve their customers and compete more effectively in the market.

DevOps comes from Developmet Operations, Development teams create the software, while Operation teams deliever and monitor it. so DevOps combines it together.

##### Problems with Traditional Development Practices

The traditional Model of development is the "waterfall" model, staring with design, code building, testing, deployment and the monitoring.

Transition from each stare requires extensive information passing, as each part is performed by another team. there are significant costs to fix issues that are discovered in later phases, and the whole process is fairly slow.

This is also true for monolithic applications, they are developed as a single unit, have high coupling and are developed using a single technology stack. this makes them big and costly to upgrade (each change requires re-testing everything), hard to understand (everything is connected), and it's very hard to switch technology or upgrade versions (all or nothing).

a lot of traditional devops is done manually, either by creating and configuring the test environments, packaging the deployment correctly, and running test on it. it's easy to forget something or make a mistake.

##### Why DevOps?

There are several benefits for using devops strategies:

- Agility
- Rapid Delivery
- Relability
- Scale
- Improved Collaboration
- Security

</details>

#### DevOps Methodlogy

<details>
<summary>
What DevOps use and care about
</summary>

> DevOps methodology increases collaboration through the entire service lifecycle, from product design through the development process to production operations. It brings people together to work and remove obstacles so they can efficiently accomplish their goals.\
> This module will dive deeper into the DevOps methodology: culture, processes, and tools.

##### DevOps Culture

Shifting to DevOps requires changes to the culture and mindest of developers.

> 1. Create a highly collaborative environment:\
>    DevOps brings together development and operations to break down silos, align goals, and deliver on common objectives. The whole team (development, testing, security, operations, and others) has end-to-end ownership for the software they release. They work together to optimize the productivity of developers and the reliability of operations. Teams learn from each other's experiences, listen to concerns and perspectives, and streamline their processes to achieve the required results.\
>    This increased visibility enables processes to be unified and continuously improved to deliver on business goals. The collaboration also creates a high-trust culture that values the efforts of each team member, and transfers knowledge and best practices across teams and the organization.
> 2. Automate when possible\
>    With DevOps, repeatable tasks are automated, enabling teams to focus on innovation. Automation provides the means to rapid development, testing, and deployment. Identify automation opportunities at every phase, such as code integrations, reviews, testing, security, deployments, and monitoring, using the right tools and services.\
>    For example, infrastructure-as-code (IaC) can be used for predefined or approved environments, and versioned so that repeatable and consistent environments are built. You can also define regulatory checks and incorporate them in test that continuously run as part of your release pipeline.
> 3. Focus on customer needs\
>    A customer first mindset is a key factor in driving development. For example, with feedback loops DevOps teams stay in-touch with their customer and develop software that meets the customer needs. With a microservices architecture, they are able to quickly switch direction and align their efforts to those needs.\
>    Streamlined processes and automation deliver requested updates faster and keep customer satisfaction high. Monitoring helps teams determine the success of their application and continuously align their customer focused efforts.
> 4. Develop small and release often\
>    Applications are no longer being developed as one monolithic system with rigid development, testing, and deployment practices. Application architectures are designed with smaller, loosely coupled components. Overarching policies (such as backward compatibility, or change management) are in place and provide governance to development efforts. Teams are organized to match the required system architecture. They have a sense of ownership for their efforts.\
>    Adopting modern development practices, such as small and frequent code releases, gives teams the agility they need to be responsive to customer needs and business objectives.
> 5. Include security at every phase\
>    To support continuous delivery, security must be iterative, incremental, automated, and in every phase of the application lifecycle, instead of something that is done before a release. Educate the development and operations teams to embed security into each step of the application lifecycle. This way, you can identify and resolve potential vulnerabilities before they become major issues and are more expensive to fix. \
>    For example, you can include security testing to scan for hard-coded access keys, or usage of restricted ports.
> 6. Continuously experiment and learn\
>    Inquiry, innovation, learning, and mentoring are encouraged and incorporated into DevOps processes. Teams are innovative and their progress is monitored. With innovation, failure will happen. Leadership accepts failure and teams are encouraged to see failure as a learning opportunity.\
>    For example, teams use DevOps tools to spin-up environments on demand, enabling them to experiment and innovate, perhaps on the use of new technology to support a customer requirement.
> 7. Continuously improve\
>    Thoughtful metrics help teams monitor their progress, evaluate their processes and tools, and work toward common goals and continuous improvement. For example, teams strive to improve development performance measures such as throughput.\
>    They also strive to increase stability and reduce the mean time to restore service. Using the right monitoring tools, you can set application benchmarks for usual behaviors, and continuously monitor for variations.

##### DevOps Practices

> DevOps culture leads to DevOps practices that are geared toward streamlining and improving the development lifecycle, to reliably deliver frequent updates, while maintaining stability.

1. Communication and collaboration - rapid, clear, transperant and frequent communication between teams.
1. Monitoring and observability - assessing how the system is performing at all times. using discrete information (such as error logs) and numeric metrics, using visualizations and centralized tools to make the data available and actionable.
1. Continuous integration (CI) - regularly mergin small changes into the main code branch (after running tests on it), finding bugs quickly and addressing them.
1. Continuous delivery/continuous deployment (CD) - regularly delivering the upgraded software to the client, rather than releasing "service packs" or yearly versions.
1. Microservices architecture - design around services with limited functionality, reduce coupling and complexity, allowing for changes and innovation in one service without effecting other components.
1. Infrastructure as code - storing configurations for testing environments and development machines in a repository, allowing for repeatability and scaling that isn't possible with manual operations.

The devops pipeline is similar to the waterfall pipeline, but it's more fluid, and since it's smaller in scale and involves less people, it's easier to go back and change something when encountering an error.

- Code
- Build
- Test
- Release
- Deploy
- Monitor

##### DevOps Tools

Devops use the cloud, as it's both a source of services that can facilitate ci/cd on the highest level, and as a way to provision infrastructure resources without buying them.

There are Tools that make development better, such as IDEs and SDKs. IDEs make writing code easier, with refactoring tools, auto completion and code suggestions. SDKs make interacting with external libraries (such as AWS cloud) easy, instead of writing API calls directly.

CI/CD tools allow automating the process from changes to the code in the source-code repository and taking it all the way to the finished product, that means building the artifacts, running tests, checking for security problems, and even releasing the code onto the cloud.

There are also tools that allow for acquiring and configuring infrastructure, which means that different environments (production, development, testing) can have consistent configurations, and any difference between them are also documented. provisioning infrastructure also allows for fine grained control of permissions, networking and dependencies.

Container services and Serverless services also allow developers to focus on the core product, rather than worry about the host environment. this can be similar to virtual machines (only lighter), or without any worries about the underlying system at all.

There are also tools for monitoring and observability, such as collecting logs, metrics, tracking requests, monitoring traffic, etc..

</details>

#### Amazon's DevOps Transformation

<details>
<summary>
Case study of How amazon itself moved to DevOps
</summary>

In the early 2000s, Amazon was a retail company with a monolith architecture. as a large company, it suffered from all the downsides of traditional devops processes. they changed it by giving small teams responsability over features, and gave them the power to control the entire pipeline, so there was no need for knowledge transfer between teams across stages, and they could eliminate redundancies.

</details>

#### AWS DevOps Tools

<details>
<summary>
AWS Tools for DevOps
</summary>

<cloud>AWS Cloud9</cloud> is a cloud based IDE to write, run and debug code.
<cloud>AWS Code Commit</cloud> is a source control service that stores the code on the cloud (similar to github). <cloud>AWS CodeBuild</cloud> automates compiling the source code, running tests and creating software packages (artifacts). <cloud>AWS CodeDeploy</cloud> automates the task of deploying the code onto compute resources (<cloud>EC2</cloud>, <cloud>Fargate</cloud>, <cloud>Lambda</cloud>, etc...). <cloud>AWS CodePipeline</cloud> combines them into a CI/CD pipeline that can run all the tasks in a repeatable way.

Once the deployment is done, <cloud>AWS CloudWatch</cloud> collects the logs from the services, and <cloud>AWS X-Ray</cloud> can trace how requests pass through the system.

</details>

#### AWS DevOps Demo

<details>
<summary>
Create and Control a CI/CD Pipeline
</summary>

> In this demo, you will see how AWS DevOps services are used to create and control a continuous integration and continuous delivery (CI/CD) pipeline. This release pipeline will automate deploying a working application to multiple Regions. The application being deployed is a simple web page. You will see how to control the workflow and speed up the pipeline by eliminating manual interventions.
>
> These steps have been completed before the demo:
>
> - The infrastructure required for the application to run was provisioned in two AWS Regions.
> - Along with the application source, an appspec.yml file and supporting scripts have been created and provided with the code.
> - A Lambda function has been created. When invoked, it checks some text on a webpage.
>
> During the demo, the pipeline will be created in the following steps:
>
> 1. The demo starts with a quick review of the provisioned infrastructure and the template file that was used to create them.
> 1. AWS CodeCommit is configured to hold the code, and is the continuous integration service that starts the pipeline with every code change.
> 1. AWS CodeDeploy is configured with specifics about each deployment Region. CodeDeploy is the continuous deployment service that installs the application on the infrastructure.
> 1. AWS CodePipeline uses the configured services to create pipelines.
>    1. First, a two stage pipeline is created that automatically deploys to Region 1.
>    1. Then, we build on the pipeline and add two new stages. This release pipeline will deploy the simple web application to two distinct AWS Regions. The pipeline still automatically deploys code changes to Region 1, but requires a manual approval before deploy to Region 2.
>    1. Finally, a Lambda function is used speed up the pipeline while maintaining control. It replaces the manual approval gate with automation.

(video)

we have a <cloud>CloudFormation</cloud> template that provions resources (<cloud>EC2</cloud> machines), we have a <cloud>CodeCommit</cloud> source code repository, <cloud>CodeDeploy</cloud> and <cloud>CodePipeline</cloud> services that automate the deployment.

we create a repository and clone it, then we create the application and commit the files into it. we use the <cloud>Cloud9</cloud> tool to write our code. we next create the deployment in the <cloud>CodeDeploy</cloud> service, and we set it to use the EC2 instances. we follow the wizard for <cloud>CodePipeline</cloud> and set the repository as our source, so we will run the pipeline when the code is changed. we also add an approval gate to require manual approval before deploying to a different region. we can replace the approval check with a lambda function

</details>

</details>

</details>

</details>
