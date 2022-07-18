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
#### Evolve
  
### Security Pillar
### Reliability Pillar
### Performance Efficiency
### Cost Pillar

##
</details>
