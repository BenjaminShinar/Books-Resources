# AWS Certified Solutions Architect - Associate 2020

Udemy course [AWS Certified Solutions Architect - Associate 2020](https://www.udemy.com/course/draft/362328/). by _Ryan Kroonenburg_

- [Section 1 - Introduction](Section_1_2_intro.md)
- [Section 2 - 10,000 Foot Overview](Section_1_2_intro.md)
- [Section 3 - IAM & S3](Section_3_IAM_S3.md)
- [Section 4 - EC2](Section_4_EC2.md)
- [Section 5 - Databases on AWS](Section_5_Databases.md)
- [Section 6 - Advanced IAM](Section_6_IAM.md)
- [Section 7 - Route 53](Section_7.md)
- [Section 8 - VPCs](Section_8_VPC.md)
- [Section 9 - HA Architecture](Section_9.md)
- [Section 10 - Applications](Section_10_Apps.md)
- [Section 11 - Security](Section_11_Security.md)
- [Section 12 - Serverless](Section_12_Serverless.md)

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
- SQS :
- FSx (windows and Lustre) - native windows managed file system. Lustre is for compute intensive data.
- RedShift - AWS data warehouse solution.
- ElastiCache - in memory "database".
- DynamoDB - AWS NoSQL solution.
- Aurora - Serverless RDS by Amazon.

CloudWatch monitors resource usage, such as the number of EC2 instances we have, EBS volumes used, redirections from load balancers, etc... . CloudTrail Monitors aws actions such as API calls or other aws actions, the focus is how the users and services interact with AWS itself, not the metrics of the services themselves.

To use the AWS CLI we need an AWS user with Programatic Access: Access Key Id, Secret Access Key.

### Acronyms

| Shorthand | Long name                           | Usage                                             | notes                                         |
| --------- | ----------------------------------- | ------------------------------------------------- | --------------------------------------------- |
| ACL       | Access Control List                 | S3 buckets                                        |
| AD        | Active Directory                    |                                                   | non-aws way to manage users in other systems  |
| AMI       | Amazon Machine Image                | EC2 machine images                                | linux flavours, windowes, etc...              |
| ARN       | Amazon Resource Name                | amazon identifier                                 |
| AWS       | Amazon Web Service                  | the amazon cloud eco-system                       |
| BI        | Business Intelligence               |
| CDN       | Content Delivery Network            |                                                   | Cloud Front                                   |
| CIDR      | Classless Inter-Domain Routing      | security groups                                   | `0.0.0.0/0` allows all access                 |
| DAX       | DynamoDB Accelerator                | in memory-cache                                   |
| DFS       | Distributed File System             |
| DMS       | Database Migration Service          | migrate from one database to another              |
| EBS       | Elastic Block Store                 |
| EC2       | Elastic Compute Cloud               | Virtual machine                                   |
| ECS       | Elastic Container Services          |
| EDA       | Electronic Design Automation        |
| EFA       | Elastic Fabric Adaptor              |
| EFS       | Elastic File System                 |
| EN        | Enhanced Networking                 |
| ENA       | Elastic Network Adaptor             | enable enhanced networking                        |
| ENI       | Elastic Network Interface           | virtual network card                              |
| FPGA      | Field Programmable Gate Array       |
| HPC       | High Performance Computing          |
| IA        | Infrequent Access                   |                                                   | S3 IA tiers                                   |
| IAM       | Identity Access Management          | users, roles, policies, access                    | always global                                 |
| IOPS      | Input-Output Per Second             |                                                   | a metric for hard disk                        |
| KMS       | Key Management Service              |                                                   | SSE-KMS                                       |
| MAC       | Media Access Control (address)      |
| MFA       | Multi Factor Authentication         |
| NFS       | Network File System                 |                                                   | used in EFS                                   |
| OAI       | Origin Access Identification        | Authentication                                    |
| OU        | Organizational Unit                 |
| PII       | Personally Identifiable Information |
| PPS       | Packets Per Second                  |                                                   | networking metric                             |
| RDS       | Relational Database                 |                                                   | mariadb, Aurora                               |
| RTC       | Replication Time Control            |                                                   | S3 buckets replication                        |
| S3        | Simple Storage Service              | storage                                           | globally unique names                         |
| SAN       | Storage Area Network                |
| SCP       | Service Control Policies            | Manage access on accounts within AWS organization |
| SLA       | Service Level Agreement             | What AWS promises (rather than advertises)        |
| SMB       | Server Message Block                |                                                   | FSx Windows                                   |
| SNS       | Simple Notification Service         | Push notification services                        | used in the billing alarm                     |
| SQS       |
| SR-IOV    | Single Root I/O Virtualization      |
| SSE       | Server Side Encryption              |                                                   | SSE-S3                                        |
| SSH       | Secure Shell                        | access to other machine                           |
| SWF       |
| TTL       | Time to Live                        | Objects at cached location life time              | CDN, CloudFront                               |
| VPC       | Virtual Private Cloud               |
| WAF       | Web Application Firewall            |
| WORM      | Write Once, Read Many               |
| WORM      | Write Once, Read Many               |
| OTAP      | Online Analytics Processing         |                                                   | BI operation                                  |
| OTLP      | Online Transaction Processing       |                                                   | BI operation                                  |
| PITR      | Point in Time Recovery              | DynamoDB backup                                   | protects against accidental writes or deletes |

</details>
