# AWS Certified Solutions Architect - Associate 2020

Udemy course [AWS Certified Solutions Architect - Associate 2020](https://www.udemy.com/course/draft/362328/). by _Ryan Kroonenburg_

- Section 1 - Introduction
- Section 2 - 10,000 Foot Overview
- Section 3 - IAM & S3
- Section 4 - EC2
- Section 5 - Databases on AWS
- Section 6 - Advanced IAM
- Section 7 - Route 53
- Section 8 - VPCs
- Section 9 - HA Architecture
- Section 10 - Applications
- Section 11 - Security
- Section 12 - Serverless

## Section 5 - Databases on AWS

## Section 6 - Advanced IAM

## Section 7 - Route 53

## Section 8 - VPCs

## Section 9 - HA Architecture

## Section 10 - Applications

## Section 11 - Security

## Section 12 - Serverless

[Aws This Week](https://acloud.guru/aws-this-week): weekly AWS content.

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

CloudWatch monitors resource usage, such as the number of EC2 instances we have, EBS volumes used, redirections from load balancers, etc... . CloudTrail Monitors aws actions such as API calls or other aws actions, the focus is how the users and services interact with AWS itself, not the metrics of the services themselves.

To use the AWS CLI we need an AWS user with Programatic Access: Access Key Id, Secret Access Key.

### Acronyms

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
| NFS       | Network File System                 |                                                   | used in EFS                      |
| SMB       | Server Message Block                |                                                   | FSx Windows                      |
| AD        | Active Directory                    |
| DFS       | Distributed File System             |

</details>
