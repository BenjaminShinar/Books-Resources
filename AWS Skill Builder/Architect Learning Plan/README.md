<!--
ignore these words in spell check for this file
// cSpell:ignore Vogels
 -->




[Architect Learning Plan](https://explore.skillbuilder.aws/learn/lp/78/Architect%252520Learning%252520Plan)

# Architect Learning Plan

## Introduction to AWS Billing and Cost Management
<details>
<summary>
Understanding the tools available, and creating spending alerts.
</summary>

Billing and cost managemtent.
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
<!-- <details> -->
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

>"Relability is the ability of a workload to perform its intended function correctly and consistently when it's expected to. This includes the ability to operate and test the workload through its total lifecycle"
>
Resiliency:
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

### Cost Pillar

##
</details>
