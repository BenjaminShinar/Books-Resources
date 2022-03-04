# AWS Certified Solutions Architect - Associate 2020

Udemy course [AWS Certified Solutions Architect - Associate 2020](https://www.udemy.com/course/draft/362328/). by _Ryan Kroonenburg_

## Section 1 - Introduction

<details>
<summary>
Course introduction.
</summary>

### Exam Blueprint

what do we need to know to pass the exam (2020 version)

- 130 minutes
- 60 question
- grades are between 100-1000, passing score is 720.
- qualification is valid for 2 years
- questions are scenario based - they aren't supposed to be tricky or memorization based.

we can look up the details in the Amazon website. it costs money. we need a certification account, we can then book an exam, get training, see the previous scores, etc.

### Why Should I learn AWS?

why learn and get certified in AWS?

(Ryan telling his own story), describing **A Cloud Guru** and **Linux Academy**.

aws Consulting partner qualification has tiers, select,advanced, premiers, in order to reach a certain tier, employees of the company need aws certifications, such as _practitioner_,_Associate_ and _professional_ and other specialized certificates.

each Tier of certifications has different certification

- Practitioner Tier
  - Certified Cloud Practitioner
- Associate Tier
  - Certified Solutions Architect Associate
  - Certified Developer Associate
  - Certified Sysops Administrator Associate
- Professional Tier
  - Certified Solutions Architect Professional
  - Certified Devops Professional
- Specialty Tier
  - Advance Networking
  - Database
  - Data Analysis
  - Machine Learning
  - Security
  - Alexa Skill BUilder

Ryan says some are easier than others, but it depends on the person. the aws platform grows each year.

</details>

## Section 2 - AWS - 10,000 Foot Overview

<details>
<summary>
Getting Acquainted with AWS
</summary>

### The History Of AWS

> "Invention requires two things:
>
> 1. The ability to try a lot of experiments
> 2. Not having to live wit the collateral damage of failed experiments"\
>    ~ (Andy Jassy, ceo of AWS)

aws started with SQS, and first marketed to developer and small companies, as it was easier to provision resources from amazon rather than buy them upfront.

Certification started in 2013,

re:invent is the aws conference, a lot of new stuff is announced then.

### AWS - 10,000 Foot Overview

there are tons of Aws Services, each year there are more and more, the services are grouped by concepts:

- Compute: EC2, Lambda
- Storage: S3, EFS
- Databases: RDS, DynamoDb
- Migration and Transfer: Snowball
- Network and Content delivery: Vpc, Cloud front
- Developer tools
- Robotics
- Block chain
- Satellite
- Management and Governance
- Media Services
- Machine Learning
- Analytics
- Security, Identity and Compliance
- Mobile
- AR and VR (augmented and virtual reality)
- Application Integration
- AWS Cost Management
- Customer Engagement
- Bussiness Application
- Desktop and App Streaming
- IOT (internet of thins
- Game Development

there are regions and availability zones. As of the time of the course, there are 24 regions and 72 availability zones. avalability zones are based on data-center. a datacenter is simply a location (one or more buildings) with tons of servers. A region consists of availability zones. there are also **edge locations**, which are end points for aws caching content, like this is used for CloudFront. edge locations aren't regions.

to pass the solution architert exam, one would need to know:

- **AWS Global infrastructure**
- **Compute**
- **Storage**
- **Databases**
- Migration and Transfer
- **Network and Content delivery**
- Management and Governance
- Machine Learning
- Analytics
- **Security, Identity and Compliance**
- Desktop and App Streaming

### How To Sign Up To AWS

Signing up into AWS and getting the free tier features.

<kbd>Create aws Account</kbd>\
use a personnel account, we need to provide credit information, even if we use a free account. choose the basic plan for support (free), we can personalize the account, and eventually sign into the console.

</details>

## Section 5 - Databases on AWS

## Section 6 - Advanced IAM

## Section 7 - Route 53

## Section 8 - VPCs

## Section 9 - HA Architecture

## Section 10 - Applications

## Section 11 - Security

## Section 12 - Serverless

## Takeaways

<!-- <details> -->
<summary>
Notes to Self
</summary>

key concepts:

- Region
- Availability zone - Amazon Data Center
- Edge Location
- Serverless
- Route53 - amazon's DNS

humans who need to interact with the aws services are **users**, they can be part of a **user group**. services that interact with other services have **roles**. Roles, users, and user groups have **policies** attached to them.

Aws services:

- IAM - Identity Access Management:
- Cloud Watch: for billing alarms
- SNS- Simple Notification Service: send emails
- S3 - Simple Storage Service: Object Storage
- EC2 - Elastic Cloud Compute
- SQS - :
- DataSync: Synchronize data between AWS and on-premises.
- Snowball: physical data transfer

Acronyms

| Shorthand | Long name                           | Usage                                             | notes                            |
| --------- | ----------------------------------- | ------------------------------------------------- | -------------------------------- |
| ACL       | Access Control List                 | S3 buckets                                        |
| ARN       | Amazon Resource Name                | amazon identifier                                 |
| AWS       | Amazon Web Service                  | the amazon cloud eco-system                       |
| CDN       | Content Delivery Network            |                                                   | Cloud Front                      |
| EC2       | Elastic Compute Cloud               |
| ECS       | Elastic Container Services          |
| EFS       | Elastic File System                 |
| IA        | Infrequent Access                   |                                                   | S3 IA tiers                      |
| IAM       | Identity Access Management          | users, roles, policies, access                    | always global                    |
| KMS       | Key Management Service              |
| KMS       | Key Management Service              |                                                   | SSE-KMS                          |
| MFA       | Multi Factor Authentication         |
| OU        | Organizational Unit                 |
| S3        | Simple Storage Service              | storage                                           |
| SCP       | Service Control Policies            | Manage access on accounts within AWS organization |
| SLA       | Service Level Agreement             | What AWS promises (rather than advertises)        |
| SNS       |
| SNS       | Simple Notification Service         | Push notification services                        | used in the billing alarm        |
| SQS       |
| SSE       | Server Side Encryption              |                                                   | SSE-S3                           |
| SWF       |
| VPC       | Virtual Private Cloud               |
| WORM      | Write Once, Read Many               |
| RTC       | Replication Time Control            |                                                   | S3 buckets replication           |
| TTL       | Time to Live                        | Objects at cached location life time              | CDN, CloudFront                  |
| WAF       | Web Application Firewall            |
| OAI       | Origin Access Identification        | Authentication                                    |
| SAN       | Storage Area Network                |
| EBS       | Elastic Block Store                 |
| PII       | Personally Identifiable Information |
| FPGA      | Field Programmable Gate Array       |
| EDA       | Electronic Design Automation        |
| AMI       | Amazon Machine Image                | EC2 machine images                                | linux flavours, windowes, etc... |
| CIDR      | Classless Inter-Domain Routing      |                                                   | `0.0.0.0/0` allows all access    |
| SSH       | Secure Shell                        |
| IOPS      | Input-Output Per Second             |                                                   | a metric for hard disk           |
| ENI       | Elastic Network Interface           | virtual network card                              |
| EN        | Enhanced Networking                 |
| SR-IOV    | Single Root I/O Virtualization      |
| EFA       | Elastic Fabric Adaptor              |
| HPC       | High Performance Computing          |
| MAC       | Media Access Control (address)      |
| PPS       | Packets Per Second                  |                                                   | networking metric                |
| ENA       | Elastic Network Adaptor             | enable enhanced networking                        |

</details>

ENI - Elastic Network Interface - virtual network card
EN - Enhanced Networking - uses single root I/O virtualization (SR-IOV) for better performan.
EFA - Elastic Fabric Adaptor - attach EC2 to accelerate High Performance Computing (HPC) and machine learning capabilities
