# AWS Certified Solutions Architect - Associate 2020

Udemy course [AWS Certified Solutions Architect - Associate 2020](https://www.udemy.com/course/draft/362328/). by _Ryan Kroonenburg_

- [Section 1 - Introduction](Section_01_02_Intro.md#section-1---introduction)
- [Section 2 - 10,000 Foot Overview](Section_01_02_intro.md#section-2---aws---10000-foot-overview)
- [Section 3 - IAM & S3](Section_03_IAM_S3.md)
- [Section 4 - EC2](Section_04_EC2.md)
- [Section 5 - Databases on AWS](Section_05_Databases.md)
- [Section 6 - Advanced IAM](Section_06_07_Advanced_IAM_Route53.md#section-6---advanced-iam)
- [Section 7 - Route 53](Section_06_07_Advanced_IAM_Route53.md#section-7---route53)
- [Section 8 - VPCs](Section_08_VPC.md)
- [Section 9 - HA Architecture](Section_09_HA.md)
- [Section 10 - Applications](Section_10_11_Applications_Security.md#section-10---applications)
- [Section 11 - Security](Section_10_11_Applications_Security.md#section-11---security)
- [Section 12 - Serverless](Section_12_Serverless.md)
- [Practice and Sample Exams](Practice.md)

## Takeaways

<!-- <details> -->
<summary>
Notes to Self
</summary>

[Aws This Week](https://acloud.guru/aws-this-week): weekly AWS content.

key concepts:

- Region
- Availability zone - Amazon Data Center
- Edge Location
- Serverless
- Route53 - amazon's DNS

humans who need to interact with the aws services are **users**, they can be part of a **user group**. services that interact with other services have **roles**. Roles, users, and user groups have **policies** attached to them.

Aws services:

- IAM (Identity Access Management):
- Cloud Watch: Monitors Resource usage and other metrics, used for billing alarms. Alarms,Events,Logs and Dashboards.
- CloutTrail: Monitors AWS actions (from the console or the API), which users and accounts.
- SNS (Simple Notification Service): send emails
- S3 (Simple Storage Service): Object Storage
  - DataSync: Synchronize data between AWS and on-premises.
  - Snowball: physical data transfer
- EC2 (Elastic Cloud Compute): Computing Instances
- SQS (Simple Queue Service):
- SWF (Simple Work Flow service)
- FSx (windows and Lustre) - native windows managed file system. Lustre is for compute intensive data.
- RedShift - AWS data warehouse solution (BI).
- ElastiCache - in memory "database".
- DynamoDB - AWS NoSQL solution.
- Aurora - highly availbe MySQL and PostgresSQL compatible Relational Database.
- Aurora Serverless - Serverless RDS by Amazon.
- DMS - Database Migration
- Managed Microsoft AD - Active Directory by AWS
- RAM - Resources Access Manger - share resources across account
- Route53 - Dns
- CloudFormation - templates for provsioning aws resources.
- Beanstalk - one click solution to provision resource, like CloudFormation for dummies.
- Elastic Transcoder - transcode (convert) media files between different formats.
- Kinesis - Streaming Data platform (continuously generated data from multiple source, not video streaming)
- Cognito - Web Identity Federation service.

CloudWatch monitors resource usage, such as the number of EC2 instances we have, EBS volumes used, redirections from load balancers, etc... . CloudTrail Monitors aws actions such as API calls or other aws actions, the focus is how the users and services interact with AWS itself, not the metrics of the services themselves.

To use the AWS CLI we need an AWS user with Programatic Access: Access Key Id, Secret Access Key.

### Acronyms

| Shorthand | Long name                             | Usage and Related AWS service                     | Notes                                                                        |
| --------- | ------------------------------------- | ------------------------------------------------- | ---------------------------------------------------------------------------- |
| ACL       | Access Control List                   | **S3, VPC**                                       |
| AD        | Active Directory                      |                                                   | non-aws way to manage users in other systems                                 |
| AMI       | Amazon Machine Image                  | EC2                                               | linux flavours, windowes, etc...                                             |
| ARN       | Amazon Resource Name                  |                                                   | amazon identifier                                                            |
| AWS       | Amazon Web Service                    | the amazon cloud eco-system                       |
| BI        | Business Intelligence                 |
| CDN       | Content Delivery Network              | Cloud Front                                       |                                                                              |
| CIDR      | Classless Inter-Domain Routing        | security groups                                   | `0.0.0.0/0` and `::/0` allow all access                                      |
| CMK       | Customer Master Keys                  | KMS                                               |
| DAX       | DynamoDB Accelerator                  | Aws DynamoDB                                      | in memory-cache                                                              |
| DFS       | Distributed File System               |
| DMS       | Database Migration Service            | migrate from one database to another              |
| DX        | Direct Connect                        | **AWS Service**                                   | Connect customer DataCenter directly to AWS and bypass internet connections. |
| EBS       | Elastic Block Store                   |
| EC2       | Elastic Compute Cloud                 | **AWS Service**                                   | Virtual machine                                                              |
| ECS       | Elastic Container Services            | **AWS Service**                                   |
| EDA       | Electronic Design Automation          |
| EFA       | Elastic Fabric Adaptor                |
| EFS       | Elastic File System                   |
| ELB       | Elastic Load Balancer                 | **Aws Service**                                   |
| EMR       | Elastic Map Reduce                    | RDs                                               | Big data processing                                                          |
| EN        | Enhanced Networking                   |
| ENA       | Elastic Network Adaptor               | enable enhanced networking                        |
| ENI       | Elastic Network Interface             | virtual network card                              |
| FPGA      | Field Programmable Gate Array         |
| HDFS      | Hadoop Distributed File System        | Amazon EMR                                        |
| HPC       | High Performance Computing            |
| HSM       | Hardware Security Module              |                                                   | key managements                                                              |
| IA        | Infrequent Access                     | S3                                                | IA tiers                                                                     |
| IAM       | Identity Access Management            | **AWS Service**                                   | always global                                                                |
| IOPS      | Input-Output Per Second               |                                                   | a metric for hard disk                                                       |
| KMS       | Key Management Service                | **AWS Service**                                   | SSE-KMS                                                                      |
| MAC       | Media Access Control (address)        |                                                   | networking                                                                   |
| MFA       | Multi Factor Authentication           | IAM                                               |
| MMP       | Massively Parallel Processing         | Redshift                                          |
| NACL      | Network Access Control Lists          | **VPC**                                           | stateless                                                                    |
| NAT       | Network Address Translation           | **VPC**                                           |
| NFS       | Network File System                   |                                                   | FS                                                                           |
| OAI       | Origin Access Identification          | Authentication                                    |
| OTAP      | Online Analytics Processing           | RedShift                                          | BI operation                                                                 |
| OTLP      | Online Transaction Processing         |                                                   | BI operation                                                                 |
| OU        | Organizational Unit                   |
| PII       | Personally Identifiable Information   |
| PITR      | Point in Time Recovery                | DynamoDB backup                                   | protects against accidental writes or deletes                                |
| PPS       | Packets Per Second                    |                                                   | networking metric                                                            |
| RDS       | Relational Database                   |                                                   | mariadb, Aurora                                                              |
| RTC       | Replication Time Control              |                                                   | S3 buckets replication                                                       |
| S3        | Simple Storage Service                | storage                                           | globally unique names                                                        |
| SAM       | Serverless Application Model          | CloudFormation Extention, Serverless Applications | Framework to build serverless applications                                   |
| SAN       | Storage Area Network                  |
| SCP       | Service Control Policies              | IAM                                               | Manage access on accounts within AWS organization                            |
| SCT       | Schema Conversion Tool                | database migration                                |
| SLA       | Service Level Agreement               |                                                   | What AWS promises (rather than advertises)                                   |
| SMB       | Server Message Block                  |                                                   | FSx Windows                                                                  |
| SMS       | Server Migration Service              |
| SMPS      | System Manager Parameter Store        | KMS and others                                    | parameters storage                                                           |
| SNS       | Simple Notification Service           | Push notification services                        | used in the billing alarm                                                    |
| SQS       | Simple Queue Service                  | **AWS Service**                                   | pull-based, standard and fifo queues.                                        |
| SR-IOV    | Single Root I/O Virtualization        |
| SSE       | Server Side Encryption                | S3                                                | SSE-S3                                                                       |
| SSH       | Secure Shell                          | access to other machine                           |
| SWF       | Simple WorkFlow Service               | **AWS Service**                                   |
| TCO       | Total Cost Of Ownership               | Application Discovery Service                     | estimate                                                                     |
| TTL       | Time to Live                          | CDN, CloudFront                                   | Objects at cached location life time                                         |
| VPC       | Virtual Private Cloud                 | **AWS Service**                                   |
| WAF       | Web Application Firewall              |
| WORM      | Write Once, Read Many                 | S3                                                |
| LDAP      | Lightweight Directory Access Protocol |
| DNS       | Domain Name Service                   |
| DC        | Domain Controller                     | AWS Managed Microsoft AD                          | running windows Server                                                       |
| RAM       | Resource Access Manager               | **AWS Service**                                   | Share resource between AWS accounts                                          |
| SAML      | Security Assertion Markup Langauge    | IAM access                                        | single sign on                                                               |
| SOA       | Start of Authority                    | Route53, DNS                                      | record with information about domain access                                  |
| VPN       | Virtual Private Network               | VPC                                               |
| IGW       | Internet Gateways                     | VPC                                               |
| XSS       | Cross-Site Scripting                  |                                                   | security attack                                                              |
| JWT       | JSON web Token                        | AWS Cognito                                       | authenticaton form                                                           |
| ECC       | Elliptic-Curve Cryptography           | KMS                                               | A type of asymmetric Key                                                     |
| DEK       | Data Encryption Key                   | KMS                                               | for file larger than 4kb                                                     |
| HSM       | Hardware Security Module              |                                                   | higher level of security than KMS                                            |

</details>
