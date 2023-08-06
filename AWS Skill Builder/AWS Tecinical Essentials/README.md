<!--
ignore these words in spell check for this file
// cSpell:ignore Seph
 -->

## AWS Technical Essentials

[AWS Technical Essentials](https://explore.skillbuilder.aws/learn/course/1851/aws-technical-essentials). instructors: *Morgan Willis*, *Seph R.*

> AWS Technical Essentials introduces you to essential Amazon Web Services and common solutions. The course covers the fundamental AWS concepts related to compute, database, storage, networking, monitoring, and security. You will start working in AWS through hands-on course experiences. The course covers the concepts necessary to increase your understanding of AWS services, so that you can make informed decisions about solutions that meet business requirements. Throughout the course, you will gain information on how to build, compare, and apply highly available, fault tolerant, scalable, and cost-effective cloud solutions.

### MODULE 1: Introduction to Amazon Web Services

<!-- <details> -->
<summary>
//TODO: add Summary
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

*Regions* are geographical regions, named after where they are and given an aws code. regions are independent from one another, and there is no automatic data transfer between them without the user setting it up. Regions consist of two or more *Availability Zones*, which are data centers spread across the region. they are remote from one another to avoid having them go down together.

Services can belong to different scopes: *Availability Zone*, *Region* or the *Global* level. global scoped services (such as **IAM**) are the same for all regions, other service might require you to choose a region and an Availability Zone, this might be done to set data durability and availability. For other services, AWS itself manages the placement in Availability Zones, and you only set the region.

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
> - Choosing a Region for AWS resources in accordance with data sovereignty regulations
> - Implementing data-protection mechanisms, such as encryption and scheduled backups
> - Using access control to limit who can access your data and AWS resources

</details>

#### Protecting the AWS Root User
#### AWS Identity and Access Management
#### Demonstration: Implementing Security with IAM
#### Hosting the Employee Directory Application on AWS
#### Module 1 Knowledge Check

</details>

MODULE 2: AWS Compute
Compute as a Service
Getting Started with Amazon EC2
Amazon EC2 Instance Lifecycle
Demonstration: Launching the Employee Directory Application on Amazon EC2
Container Services
Introduction to Serverless
Serverless with AWS Fargate
Serverless with AWS Lambda
Choosing the Right Compute Service
Module 2 Knowledge Check
MODULE 3: AWS Networking
Introduction to Networking
Amazon VPC
Amazon VPC Routing
Amazon VPC Security
Demonstration: Relaunching the Employee Directory Application in Amazon EC2
Module 3 Knowledge Check
MODULE 4: AWS Storage
Storage Types
File Storage with Amazon EFS and Amazon FSx
Block Storage with Amazon EC2 Instance Store and Amazon EBS
Object Storage with Amazon S3
Choosing the Right Storage Service
Demonstration: Creating an Amazon S3 Bucket
Module 4 Knowledge Check
MODULE 5: Databases on AWS
Introduction to Databases on AWS
Amazon RDS
Purpose-Built Databases
Amazon DynamoDB
Choosing the Right Database Service
Demonstration: Implementing and Managing Amazon DynamoDB
Module 5 Knowledge Check

MODULE 6: Monitoring, Load Balancing, and Scaling
Monitoring
Amazon CloudWatch
Solution Optimization
Traffic Routing with Elastic Load Balancing
Amazon EC2 Auto Scaling
Demonstration: Making the Employee Directory Application Highly Available
Employee Directory Application Redesign
Module 6 Knowledge Check