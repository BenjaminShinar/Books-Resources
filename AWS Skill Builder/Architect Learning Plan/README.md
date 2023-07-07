<!--
ignore these words in spell check for this file
// cSpell:ignore Vogels Postgre Memecahed
 -->

[Architect Learning Plan](https://explore.skillbuilder.aws/learn/lp/78/Architect%252520Learning%252520Plan)

# Architect Learning Plan

## Introduction to AWS Billing and Cost Management

<details>
<summary>
Understanding the tools available, and creating spending alerts.
</summary>

Billing and cost management.

- estimate and plat aws costs
- receive alerts if costs exceed or approach a threshold
- asses the biggest investment in aws resources.
- simplify accounting when working with multiple AWS accounts.

in the console, we have a high level dashboard, we can see the current month, or look at month-to-month comparison, we can also check if we are exceeding the aws free-tier.

we can get the billing invoices as pdf or csv, and if we have additional costs (like saving plan or support plan), they will also be there.

cost explorer is a BI tool that can show the breakdown of the costs (filters and views), we can check costs by services, regions, hours, etc.. it's also possible to create custom views.

there is also AWS current ussage report - AWS CUR, which is more detailed. we can store the generated data in S3 bucket, and then we can visualize it with external tools.

we have APIs which can expose the data externally.

### Monitoring Costs

**aws-busget**\
Allows us to create budgets, set up alerts, and track if we are on the way to exceed the budget. this can be customized for region, tags, and so on.
budgets can be created via the console, API, or from cloud formation templates.
aws budgets is free, it can be combined with SNS to send alert, or be integrated with various tools and other messaging services.

**AWS Cost Anomaly Detection**\
this feature uses machine learning to discover anomalous spending, and finding the root cause.

### Setting a spending alert with AWS budget

in the console, we choose the <kbd>Aws budget</kbd> service, click <kbd>Create Budget</kbd>, and under budget type we select <kbd>Cost-Budget</kbd>, we decide the amount, the scope, the details (name, period), and then we <kbd>Add an alert threshold</kbd>, which will trigger when a percentage is reached (and an email will be sent). we could also add sns alert and chatbot alert, then we click <kbd>Create Budget</kbd>. it will take about 24 hours for it to be populated.

###

cost explorer is proably the best tool to start with. there are tools that can help us monitor the costs, so we won't be surprised at the end of the month.

</details>

## AWS Shared Responsibility Model

<details>
<summary>
AWS and the customer share responsibility for security and compliance
</summary>

both AWS and the developer are responsible for security, this is done by dividing the layers, some are under aws responsability, and some one managed by the customer.

layers:

- Physical - metal,brick and mortar - aws responsability
- Network - the protocols that operate the VPCs, etc...
- Hypervisor - Xen based hypervisor, but custom builds.
- Guest os - if EC2 - then the user chooses the image, and from this point, the information is secured and only the customer can view.
- application - user
- user date - user

aws is audited by many companies.

</details>

## AWS Well-Architected

<details>
<summary>
Dive deep into the Well-Architected Framework
</summary>

### The AWS Well Architected Framework

strategies and best practices, measure your architecture against benchmark and address any shortcoming.

short video by Dr Werner Vogels (amazon cto)

- security
- performance
- relability
- cost effectiveness

> "Everything which used to be hardware is now software."

this removed many constraints of the past.

a well architecured framework is a way to ask questions about the workload and about how it structured. there are also desing principals and pillars

> "What is a workload?"\
> A workload is defined as a collection of interrelated applications, infrastructure, policy, governance and operations running on AWS that provide business or operation value.

in the traditional world, we had to guess which infrastructure the code will run on, and it was also hard to test on scale, hard to justify experimenting with other options due to the costs, and the architecture was settled the moment it was released to production, as it was very hard to change and switch over.

in the cloud, those constraints were removed

- no guessing of capacity
- testing at production scale
- experimenting made easier
- architecture can evolve
- data-driven architecture

Pillars of well architecture

1. Operation Excellence
2. Security
3. Relability
4. Performance Efficiency
5. Cost optimizations

stable, efficient and consistent architecture.

Operation Excellence - run and monitor systems that deliever business values.

- organization - how the organization structure enables development
- prepare - when are they ready to move
- operate - how to run the day-to-day procedures, how to identify changes and risks.
- evolve - continues improvement

Security - protecting information, system and assets.

- (IAM) Identity and access management - who can do what to which resource
- detection
- infrastructure protection
- data protection
- incident response - responding to security event

Reliability - recover from failure, meeting demands

- foundations
- workload architecture
- change management
- failure management

Performance Efficiency - using IT resources efficiently

- selection the right tools
- review - make changes if needed
- monitoring
- trade-offs

Cost optimization - achieve outcome at lowest price

- practice cloud financial managements
- expenditure and usage awareness
- cost effective resources
- manage demand and supply resources
- optimize over time - new features are continuously released, so we can take advantage of that.

each pillar has a set of questions, based on context, and some common best practices.

we can use these pillars to identify and asses how other organizations and teams are doing. we learn to think in a cloud-native way and to apply cloud design principle and to consider "what-ifs" and failure scenarios.

### Operational Excellence Pillar

features, design principals and best practices

> The ability to run and monitor system to delever business value and to continually improve supporting process and procedures.

in a traditional environment,

- manual changes
- batch changes - big releases
- not enough time to test
- reactive, not proactive, never enough time to learn
- stale documentation

in the cloud, those constraints are removed, and we can treat infrastructure change more similarly to how we treat software changes.

- perform operation as code
- make frequent small, reversible changes
- refine operations procedures frequently
- anticipate failure
- learn from all operational failures

#### Organization

common understanding, shared business goals and knowledge, understaning responsibility, dependencies and how teams interact. having an organization culture.

businesses exists to serve customer needs, operations exists to serve business needs, there are internal and external requirements which effect operation priorities, and there always are trade-offs.

#### Prepare

Design telemetry - making sure we have the correct information.

improve flow - excellence the rate of improvements, identifying issues before production

mitigate deployment risks - finding and identify problems, be able to recover and rollback

understand operationa readiness - know what we can do know and what we don't know.

using multiple environment, infrastructure as code, having environment increasingly similar to production, with configurations and security and scale increasing at each level.

#### Operate

understanding the health of the workload and operations health. we need to know the health of the workload, and also know the operations metrics. we need to make the metrics and the data visible. we should also have metrics for changes like failing deployments, to know if we are operating successfully.

- Events - observation of intrest
- Incident - an event that requires a respone
- Problem - an incident that either recurs or cannot be currently be resolved.

when we have an alert - some event that we care about, we should have a defined "playbook" - what happened, who is responsible, what steps can be taken, and what escalations are possible. we should also inform affected stakeholders when alerts are raised and when they are resolved and everything is back to normal.

#### Evolve

learn from experience, make improvements, share leaning and lessons across the team.

feedback loops - a way to identify areas for improvements, should come from the team, they can be periodically meetings to go over metrics and determine if a change is needed. they can also be used to recognize improvement.

#### Summary

- Understand business priority
- Design for operations
- evaluate operations readiness
- understand workload and operation health
- prepare for, and respond to, events
- learn from experience, share learning and make improvements.

### Security Pillar

> The ability to protect information, systems, and assets while deliever business value through risk assessments and mitigation strategies

security is applied all laters, having strong identity foundations, fine grained access control and be prepared to handle security events and automate it.

- aws accounts
- aws organization
- aws control tower

#### (IAM) Identity and Access Management

managing human and machine identities with IAM. granular permissions, sign-in mechanisms, centeralized identity provider. using secrets in a secure matter with aws services.

using the principle of least privilege, use groups and roles rather than user policy. limit public and cross-account access. continually reduce permissions.

**Amazon Cognito**

identity broker. supports multiple login providers, manage users/device, no matter which identity provider they use. provide outside identities with access to aws resources.

#### Detection Control

detecting and identifying security events.

lifecycle controls, internal auding, automated alerting and responses.

we can use AWS Config Service to run code when there is a change to aws resource, we can set up rules that validate the behavior, and can run code automatically.

#### Infrastructure Protection

systems and services in the workload are protected.

trust boundariessystem security configuration, policy enforcement points.

controling traffic at all layer - protect against external access from the internet, (amazon WAF - web application firewall), in a vpc, use subnets and security groups.

managed services also reduce maintenance tasks, such as provioning and patching.

#### data Protection

Identifying and classifying data, keep it protected at rest and at transit.

PII - personally identifiable information

amazon MACIE can help with classifying data.
Key management service, keep people away from the data, use dashboards rather than allow direct access, and use protection in transit.

integrating aws services with encryption options. even when users have access to data, prefer to have them operate through other venues to investigate data in normal use-cases.

#### Incident Response

how to respond to security events, have the correct tools, such as as a clean room, simulate security events and prepare your response.

we can use AWS cloudFormation to spin up an clean environment with the right tools and security, use ebs snapshot to investigate in an isolated environment.

#### Summary

- Protect information, system and assets
- keep root account credentials protected
- enctyp data at rest and at transit
- ensure only authorized and authenticated users are able to access your resources
- using detective controls to identify security breaches.

### Reliability Pillar

> "Relability is the ability of a workload to perform its intended function correctly and consistently when it's expected to. This includes the ability to operate and test the workload through its total lifecycle"
>
> Resiliency:

- recover from infrastructure or service disruptions
- dynamically acquire computing resources to meet demand
- mitigate disruptions

recover from failures

- Foundations - setup and cross project
- Workload architecture
- Change Management
- Failure Managements

in a traditional environment, we test during production and test performance, but we don't tend to test the failure situation, so when failure happens we handle it manually. and when needed, we write down instructions for the future, but we don't make this part of the flow. and since there are many points of failure, and failure is so time consuming, we tend to be over cautious and over-estimate our capacity needs, just to avoid the failure case.

in the cloud native, we can have recovery procedures, we can test the recovery procedures, and if needed, we can scale horizontally to increase capacity.

#### Foundations

Foundational requirements are beyond the scope of a single project or workload. this can be the bandwidth of the data server, or other constraints. in traditional environments, these requirements are constraints and can't be easily changed or even observed.

for cloud based architecture, there are service limits, which both protect the user from over provisioning, and prevents over usage of apis. some of these limits can be changed by the normal flow of work (scaling up the disk size), and some require active work against AWS to increase the limits per account.

network topology.

#### Workload Architecture

SOA - service oriented architecture, microservice oriented architecture. we want to have distributed systems with services and micro-services, which can operate with one another, without effecting one another.

we use loosly coupling, and we design to prevent failures, to limit the failures from effecting other components, and so on.

#### Change Management

knowing how changes effect the system, monitoring services and seeing audits with aws tools - like monitoring how many restarts occur, and keeping track of provisioned resources.

(AWS CloudWatch, AWS SNS).

one example is automatic scaling to meet demands for a service that has peaks and vallies in it's usage.

having a pipeline for changes is one kind of change management, this can include integration testing, immutable infrastructure (green-blue deployment) and a manual / checklist of how to deploy.

#### Failure Managements

failure will occur, we need to know that they happen, be able to anticipate them, and know how to respond.

we can use 'fault isolation' boundaries - keeping the problem contained.

- multiple availability zones and region
- bulkhead (partitons, shards, cells) architecture

we should have a backup and recovery (distaster recovery) chains, we want to be able to recover automatically, and we want to monitor how effective this DR process is, in terms of speed, reliability, etc.

we should also document how we identify failures, what we look for. we should have tests that inject 'chaos' and failures to the system to make sure it's resilient.

### Performance Efficiency Pillar

using resources effectively

in a traditional environment, we usually use the same technology stack, it was hard to get new resources, even if it's just for an experiment.

in the cloud, we can try new technologies quickly, and use server from all around the globe.

#### Selection

we need to select the appropriate resource types, and sometimes have multiple solutions for a same task.

Compute, Storage, Database and Networking resources.

benchmarking, load testing.

#### Review

things change fast in the cloud, so there are new things all the time, maybe in the time since we've made our decision about the resources, a new option was released?

using AWS CloudFormation to define architecture as code. so we can always try out new things.

#### Monitoring

monitor performance. use automation and alarms (when thresholds are exceeded)

CloudWatch, Kinesis, SQS, Lambda

#### Trade-offs

there are tradeoff, we can trade space (memory/storage) and usually get speed, and we can trade scale (increase cost) to get better performance in many cases.

Proximity and Caching

- CloudFront - content distributuion network
- ElasticCache - cache layer
- RDS - read replicas

### Cost Optimization Pillar

achieve business outcome at lowest price.

financial managements
expenditure and usage awareness
cost-effective resources
manage demand and supply
optimize over time

in a traditional environment, there is a usually centeralized coss, so it's hard to attribute cost to a specific module or change. someone has to be paid to maintain the servers, and it's hard to take advantage of economics of scale.

in the cloud, those restraints don't apply, we can pay only for what we use, aws takes care of the economics of scale. we can better analyze and attribute costs to the service that caused them, and there is no extra costs for maintaining the servers.

#### Cost Effective Resources

use the corret machine, sometimes it's better to use a stronger machine if we can get the job done faster than a weaker machine, and the overall costs will be lower.\
it's also adviced to use managed services, which elimintate this consideration entirely

#### Pricing model

- on-demand - pay as you use, no commitments
- saving plan - pay upfront, get a discount
- reserved instances -
- spot instances - for work that can be run at anytime

#### Managed Aws Services

using services that aws provides, so you don't have to provision aws ec2 machines.

- RDS - relational databases
- RedShift - data warehouse, BI
- CloudFormation - build from templates
- Elasticsearch - key-value search databases
- DynamoDB - document database
- ElasticCache - caching layers
- Elastic Beanstalk - build from recpies(?)
- WorkMail - email server

#### Manage Supply and Demand

match the workload demand with resources of demand. don't pay for extra.

- auto scaling
- buffering/queing to distribute workloads over time
- monitoring tools

#### Expenditure Awareness

understanding where our costs are.

- tag resources
- aws cost explorer

#### Optimizing Over Time

use new features as they become available, remove old and unused resources. staying up to date with changes. look at the "AWS blog" to see what's new.

### The Well-Architected Review

looking at a system and determing if it's "well architected". we can make this review possible by following a consistent approach.

we look at each of the pillars and ask questions about our system, this is not an audit, and it's designed to be pragmatic, down to earth. it's also a process, rather than a single review event.

earlier is better, be aware of the the decisions made and the decisions not made.

there is a framework whitepaper, also resources per pillar, and other online resources by aws.

APN - Amazon Partner Network

### AWS Well-Architected Tool

> The AWS Well-Architected Tool is an architecture review tool that provides customers and partners with a consistent approach to reviewing their architectures against current AWS best practices, and gives advice on how to architect workloads for the cloud.

it compares the five pillars against the AWS well architecture framework.

a workload is a set of components (aws resources, code, etc..) which are collected together to bring value. they can be in one account or distributed across several aws accounts.

we use milestones to keep records of how the system is holding up as a well-architected system over time.

in the console, as one of the management services, choose <kbd>AWS Well-Architected Tool</kbd>.

- Define Workload and Perfrom Reviews
- Locate Helpful Resources
- save milestones
- use dashboard
- generate PDF reports
- assign priories to pillars
- create improvement plan

#### Define Workload and Perfrom Reviews

click <kbd>Define workload</kbd>, give it a name and details, then click <kbd>Start Review</kbd>, and then follow the pillars and answer the questions, there are resources, an inline glossary for each question,

#### Improvement Plan

after we finish answering the questions, we can choose the <kbd>improvement plan </kbd> tab to see the high risk items and how to fix them.

#### Saving a Milestone

A milestore captures the current status of a workload review, it's a way to mark changes and keep records. we simply click <kbd>Save milestone</kbd>, and give it a name.

#### Generate Reports

we can get a report about each workload, we select a workload and click <kbd>Generate Report</kbd>, this creates a pdf file.

#### Dashboard

the dashboard shows us a overview of all our workloads, the status of the review process, the milestones, and the number of risk items.

</details>

## Exam Readiness: AWS Certified Solutions Architect – Associate (Digital)

<details>
<summary>
Going over the topics in the exam
</summary>

### The Exam Overview

preparing for the exam, covering the domains of the exam (which are analogue to the five pillars)

| Domain                         | Pillar                  |
| ------------------------------ | ----------------------- |
| Design Resilient               | Resilient               |
| Define Performant              | Performant              |
| Specify Secure Applications    | Secure                  |
| Design Cost Optimized          | Cost Optimized          |
| Define Operationally Excellent | Operationally Excellent |

the old version of the exam focused on aws services, while the newer version focuses on understanding the how services relate to the design.

the solutions architect associate is one of the most popular certifications, it's the first step towards the professional and specialized certifications.

65 questions, we can mark hard questions for review at the end, 130 questions, no penalty for guessing. some questions have one correct answer, and in other we need to select several answers.

> 1. Read both the question and the answer in full one time through.
> 2. Identify the features mention in the answers.
> 3. Identify text in the question that implies certain AWS features.
>    - required IOPS
>    - data retrival times
> 4. Pay Attention to qualifying clauses, these clauses might eliminate certain answers.
>    - "in the most cost-effective way"
>    - "will best fulfill"
> 5. Eliminate obviously wrong answers to narrow the selection of possible right answers.

### Design Resilient Architectures

Best practices:

> 1. Choose reliable/resilient storage - don't loose data.
> 2. Determine how to design decoupling mechanisms using AWS services - avoid cascading failures.
> 3. Determine how to design a multi-tier architecture solution - not everything scales the same.
> 4. Determine how to design a high availability and/or fault tolerant solutions - keep providing value even when something fails

EC2 Instance Store

- Ephemeral volumes
- Only certain EC2 instances
- Fixed Capacity
- Disk Type and capacity depends on EC2 instance type
- Application-level durability

we use this for caching and storing data we have somewhere else, so we get speed, but the data itself is durable, as it is stored elsewhere.

Elastic Block Store (EBS) - attachable storage

- Different types
- Encryption
- Snapshots
- Provisoned Capacity (iops)
- Independent life cycle than EC2 instance - stop and restore ec2 instances
- Multiple Volumes striped to create large volumes

there ssd and hdd storage. sdd is usually faster, but hdd are good for sequential data. if we have fewer operations but they operate on sequential data we can use HDD, but for random access we should use SDD. there is usually a general purpose option and a high performance option.

[ebs volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html) - tables are here.

IOPS - input output operations / second - higher is better. provisoned iops means that we can get more operatios at a cost.

ssd types:

- gp2 (general purpose)
- gp3 (general purpose, higher throughput)
- io1 (io intensive)
- io2 (io intensive, durable)
- io2 block express (higher iops and throughput)

hdd types

- sc1 (cold storage) - low cost
- st1 - throughput optimized

EFS Elastic File system

- File storage in the AWS Cloud - not block or object storage.
- Shared storage - multiple EC2 instances can access
- Petabyte scale file system - very high
- Elastic capacity - scales as needed
- Supports NFS v4.0 and 4.1 (NFSv4) protocol
- Compatible with linux-based AMIs for amazon EC2 - not supported for windows

an EFS has mount end points, each mount is attached to an EC2 at a mount target.

Amazon S3 - object bucket

- Consistency Model - distributed systems
- Storage Classes and Durability (standard, standard IA)
- Encryption (at rest) with SSE-S3, SSE-KMS, SSE-C
- Encryption (at transit) - https
- Versioning

strong consistency for new objects, eventual consistency for updates. Storage classes with different prices. different ways to encrypt data with S3 keys, KMS keys, or customer provided keys.

Glacier is a S3 tier, has vaults and archives. we can retrieve data in different tiers. (bulk - slower, standard, expedited - faster). encryption by default.

S3 buckets can have lifecycle policies, moving between tiers according to usage.

#### Decoupling Services

if one service fails, it doesn't fail the others. for tightly coupled services, the blast radius of failure is large and effects other systems, this means that it's harder to identify the root cause of a problem, and updates are also harder.

SQS is one way to decouple messages between services. we can also use a logging service decoupled from the application server. a Load balancer can also help with distributing, and sometimes a queue can be enough.

elastic Ip addresses allow us to have a consistent ip address that isn't dependant on the machine itself, if one machine fails, the elastic ip address is simply routed to the new instance, and the client doesn't need to know.

#### High Avalability

> "Everything Fails, All the time"

we need to design the system to handle failure, we need to accept failure of some resources as something that happens occasionally, so we decouple our resources, use automation tools, and have a highly available architecture that incorporates those failures and deals with them.

#### Cloud Formation

a service to create a deployment, a declarative programming Langauge for deploying resources. a template to create a stack of resources in a single operation (think terraform)

#### Lambda

Data is stored in databases, compute is done by code.

> - Fully managed compute service that runs stateless code (Node.js, Java, C#, Go and python) in response to an event or on a time-based interval.
> - Allows you to run code without managing infrastructure like EC2 instances and auto Scaling groups.

we pay per invocation and compute time. we get scalability and resiliency from aws.

### Design Performant Architectures

- Choose Performant storage and databases
- Apply Caching to improve performance
- Design Solutions for elasticity and scalability

#### S3 Buckets - object storage

EBS is block storage, the performance metric is volume size, iops and throughput.

offload static data to S3 - image files, html, etc... the webserver should focus on dynamic data. buckets are tied to a regional, but are accessible anywhere. the naming is globally unique.

S3 payment model\
pay only for what you use:

- Gbs per month
- Transfer out of region
- PUT, COPY, POST, LIST and GET requests.

Free of charge:

- Transfer into Amazon S3
- Transfer ouf ot Amazon S3 to amazon CloudFront (edge location) or the same region.

storage classes: general, IA, Glacier...

data files are usually accessed used in the time period they are created at, and after some time, they stop being relevent and are kept for archiving purposes only. that's why we use LifeCycle Polices.

#### Performant Storage On Database

- RDS - relational Database
- DynamoDB - noSQL document
- Redshit - SQL like syntax, but for analytical operations, not insertions

Use RDS:

- complex transactions or complex queries
- medium-to-high query/write rate
- No more than a single worker node/shard
- high durability

Don't use RDS:

- massive read writes rates (over 100k a second)
- Sharding
- Simple GET/PUT requests and queres
- RDBMS customization

dynamodb automatically shards data, and scales horizontally.

the RDS can scale up with a stronger machine, we can also use read replicas to get better read performance (some engines)

DynamoDB is a managed document storage database, we define the capacity for read/writes. there is a data size limit.

- RCU - read capacity unit
- WCU - write capcity unit

#### Caching

Applying Caching to improve performance, an easy way to get performance boosts. it can be done at many levels.

- CloudFront level - for static files, which can also be stored at S3 buckets, rather than the webserver
- Elastic Cache - over the database backend
  - Memecahed - simpler
  - Redis - higher complexity,

#### Cloud Front

Content delivery system, both static and dynamic.

origins:

- S3
- EC2
- ELB2
- HTTP servers

improved security:

- AWS shield (standard and advanced)
- AWS WAF (web application firewall)

even if we use dynamic data, we get a speed up because it operates over the amazon fibers, rather than the regular internet.

#### Auto Scaling

**vertical scaling (scale up, scale down)** - change specifications of instances (more memory, cpu).\
**horizontal scaling (scale out, scale in)** - change the number of instances (add and remove as needed).

the easy way to scale out is by using Auto Scaling, which can launch new instances, register them with load balancers, and launch them across avalability zones.

we can have CloudWatch monitor our metrics, and have this alarm trigger an auto scaling action.

**Components:**

Auto Scaling Launch Configuration

- specifies EC2 instance size and AMI name (id)

Auto Scaling Group

- References the launch configuration
- Specifies min, max and desired size of the auto scaling group
- May reference an ELB
- Health Check Type

Auto Scaling Policy

- Specifies how much to scale in or scale out
- One or more can be attached to auto scaling group

auto scaling uses CloudWatch to monitor metrics

- cpu
- network
- queue size
- custom metrics

the load balancer distributes requests between instances.

### Specify Secure Applications and Architectures

1. Determine how to secure application tiers.
2. Determine how to secure data.
3. Define the networking infrastructure for a single VPC application.

**Shared responsibility model:**\
the line between what Aws takes care of, in managed services, aws takes more responsibility.

**Principle of least privileges:**\
Users are allowed to do all that they need to do, and nothing more.

**AWS identities (IAM):**

- creating user, groups, roles and polices.
- define permissions to control which AWS resources users can access

IAM integrates with microsoft Active Directory and AWS Directory service using SAML identity federation

- IAM Users - user created within the account
- Roles - temporay identities used by EC2 instances, lambdas and external users.
- Federation - Users with Active Directory identities or other corporate credentials have role assigned in IAM
- Web Identity Federation - Users with web identities from amazon.com or other openID providers have role assign using Security Token Service (STS)

#### Amazon VPC - Virtual Private Cloud

- Organization: Subnets (private ip address ranges)
- Security:
  - security groups
  - access control list
- Network isolation:
  - internet gateways
  - virtual private gateways
  - NAT gateways
- Traffic Direction: Routes

> **Public subnets** are used to support inbound/outbound access the public internet, and the include a routing table entry to an internet gateway.\
> **Private subnets** do not have a routing table entry to an internet gateway. the aren't directly accessible from the public internet. To support restricted, outbound only public internet access, we typically use a "jump box" (NAT, Proxy, Bastion host).

- NAT - Network Address Translation
- ENI - Elastic network interface

|             | Security Groups                     | Access Control Lists                | Notes                              |
| ----------- | ----------------------------------- | ----------------------------------- | ---------------------------------- |
| Access type | Specify a port, protocol, sourec Ip | Specify a port, protocol, sourec Ip | SSH, http                          |
| Rules       | Explicit Allow only                 | Explicit Allow or Deny              |
| State       | Stateful                            | Stateless                           | stateful - request in, respone out |
| Application | Applied to ENIs                     | applied to Subnets                  |
| Association | Associated with a single VPC        | Associated with a single VPC        |
| Supports    | VPC and EC2 classic                 | VPC only                            |

we use security group to control traffic into, out of, and beteween resources. VPCs can span across Avalability Zones, and so can subnets.

VPC connections

- internet Gateway - connect to the internet
- Virtual private gateway - connect to VPN
- AWS direct Connect - dedicated pipe to on premises data centers
- VPC peering - connect to other VPCs
- NAT instance/gateways - allow internet traffic from private subnets

1. a vpc has a internet gateway.
2. a public subnet has a NAT instance or a NAT gateway with public IP
3. a private subnet with a private ip
4. routing table

NAT gateway is scalable, while a NAT instance is a single EC2 machine.

#### Securing the Data Tier

Data in transit

- in and out of AWS
  - SSL over Web
  - VPN for IPsec
  - IPsec over AWS Direct Connect
  - Import/Export/SnowBall
- between AWS
  - AWS API calls use HTTPS/SSL by default

Data at Rest

- Amazon S3
  - private by default, reqires AWS credentials to access
  - access over HTTP and HTTPS
  - audits access to all objects
  - suports ACL and policies at levels of granularity
    - Buckets
    - Prefixes (directory/folder)
    - Objects
- Amazon EBS

Server side encryption (SSE)

- Amazon S3 managed keys (SSE-S3)
- KMS managed keys (SSE-KMS)
- Customer provided keys (SSE-C)

Client Side encryption (CSE)

- KMS managed master encryption keys (CSE-KMS)
- Customer provided master encryption keys (CSE-C)

we have aws services for managing keys: KMS, and AWS CloudHSM (hardware encryption, stronger compliance).

### Design Cost Optimized Architectures

1. Determine how to design cost optimized storage
2. Determine how to design cost optimized compute

pay as you go,pay less when you reserve, pay even less per unit by using more (volume discount).

in aws, we pay for:

- compute
- storage
- data transfer

for compute, we pay for time, machine configuration, elastic ip addresses, montiroing,

EC2 pricing:

- ec2 instance family
- tenancy
- pricing option

instance storage is free, but it's ephemeral.

reserved instaces can give discounts

- standard
- convertable
- schedules

spot instances are a way to get an instance at a lower price, but it can be lost if the price goes higher (there are spot blocks).

#### Storage

pricing factors:

- Storage class - (S3, IA, Glacier)
- Storage amount
- Requests
- Data transfer

EBS:

- Volumes
- input/output operations per second (IOPS)
- Snapshots
- Data Transfer
- storage type: SSD, HDD

#### Serveless architecture

We don't pay for serverless services when we don't use them.

Lambda - S3 - DynamoDB - Amazon API Gateway

Cloudfront allows Cacheing, there is no extra cost for tranfering data between S3 and cloudFront. it also helps reducing compute time (and money).

CloudFront Pricing:

- Traffic distribution
- Requests
- Data transfer out

### Define Operationally-Excellent Architectures

1. Choose Design Features in solutions that enable operation excellence

automated system which adapts to changes

- prepare
- operate
- evolve

Best practices

- Perform operations with code.
- Annotate documentation.
- Make frequent, small, reversible changes.
- Refine operation procedures frequently.
- Anticipate failure.
- Learn from all operational failures.

AWS services which support operational excellence:

1. AWS Config - tracks resources
2. AWS CloudFormation - infrastructure as code
3. AWS CloudTrail - logs api call and actions (auditing)
4. VPC Flow Logs - monitor network traffic
5. AWS Inspector - EC2 vulnerabilities
6. AWS Trusted Advisor - checks if the account follows best practices
7. AWS CloudWatch - track metrics, triggers alarms, monitor resources

### Wrap up:

the five domains

</details>
