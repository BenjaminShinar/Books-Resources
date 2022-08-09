<!--
// cSpell:ignore RTMP MPLS ISCSI
-->

Practice exams from Different Sources:

- https://wwww.udemy.com/course/aws-question-banks
- https://www.udemy.com/course/aws-certified-solutions-architect-associate-practice-tests-k

[Cheat Sheets:](https://digitalcloud.training/aws-cheat-sheets/)

- Database
  - [RDS](https://digitalcloud.training/amazon-rds/) - relational database
  - [Aurora](https://digitalcloud.training/amazon-aurora/) - fully managed serverless RDS
  - [DynamoDB](https://digitalcloud.training/amazon-dynamodb/) - document database (like mongo)
  - [Redshift](https://digitalcloud.training/amazon-redshift/) - data warehouse for analytics
  - [ElasticCache](https://digitalcloud.training/amazon-elasticache/) - Caching Layer
- Compute
  - [Lambda](https://digitalcloud.training/aws-lambda/) - serverless compute action
  - [EC2](https://digitalcloud.training/amazon-ec2/) - virtual machines
  - [EC2 Auto scaling](https://digitalcloud.training/amazon-ec2-auto-scaling/) - increased performance.
  - [ECS and EKS](https://digitalcloud.training/amazon-ecs-and-eks/) - run workloads in containers
- Storage
  - [EBS](https://digitalcloud.training/amazon-ebs/) - elastic block storage, persistent
  - [S3 and Glacier](https://digitalcloud.training/amazon-s3-and-glacier/) - object store buckets
  - [Athena](https://digitalcloud.training/amazon-athena/) - Query S3.
  - [FSX](https://digitalcloud.training/amazon-fsx/) - fully managed file servers
  - [EFS](https://digitalcloud.training/amazon-efs/) - elastic file system. NFS protocol
- Networking
  - [ELB](https://digitalcloud.training/aws-elastic-load-balancing-aws-elb/) - elastic load balancing - disturbing load across resources.
  - [CloudFront](https://digitalcloud.training/amazon-cloudfront/) - content distribution
  - [Global Accelerator](https://digitalcloud.training/aws-global-accelerator/) - improved networking
  - [WAF & Shield](https://digitalcloud.training/aws-waf-shield/) - protecting from web exploits and attack.
  - [Route53](https://digitalcloud.training/amazon-route-53/) - DNS Services
- On premises access
  - [AWS Direct Connect](https://digitalcloud.training/aws-direct-connect/) - connect on-premises data center to cloud.
  - [Storage Gateway](https://digitalcloud.training/aws-storage-gateway/) - hybrid storage between the cloud and the on-premises environment.
  - [Migration service](https://digitalcloud.training/aws-migration-services/) - database, server, syncing.
- [VPC](https://digitalcloud.training/amazon-vpc/) - logical partition of resources in the region
- [API Gateway](https://digitalcloud.training/amazon-api-gateway/) - manage Rest API
- [CloudWatch](https://digitalcloud.training/amazon-cloudwatch/) - monitoring tool
- [CloudTrail](https://digitalcloud.training/aws-cloudtrail/) - auditing
- [AWS Config](https://digitalcloud.training/aws-config/) - history of configuration changes - similar to cloud watch, but about changes, not actions.
- [BeanStalk cheatSheet](https://digitalcloud.training/aws-elastic-beanstalk/) - deploy applications
- [Kinesis](https://digitalcloud.training/amazon-kinesis/) - high performance real time data entry point.
- Managerial Services
  - [IAM](https://digitalcloud.training/aws-iam/) - identity access, policies
  - [AWS Organization](https://digitalcloud.training/aws-organizations/) - manage many AWS accounts in the same organization
- [KMS](https://digitalcloud.training/aws-kms/) - key management.
- [Application integration Services](https://digitalcloud.training/aws-application-integration-services/) - services which communicate between other services
- [Additional](https://digitalcloud.training/additional-aws-services/) - Glue, Data Pipeline, EMR


Open Questions:
1. When is aws VPN used?
2. Are ELB in a single AZ or multiple? - multiple AZ, but single region.
3. What are these - **DISASTER RECOVERY POLICIES**
   - Pilot Light - only core services are ready
   - Warm Standby - everything is ready, but scaled down
   - Multi-site - duplicates are active
   - Backup-Restore - no redundancy
4. Penetration Testing on AWS
5. When is AWS Shield Advanced Used?
6. What is AWS Config?
7. S3 Vault Lock - **unchangeable policy** to S3 glacier. stronger than access policy, used for compliance. for example, allows to prevent deletion of objects.
8. Support Plan features
9. What is OpsWork - automated configuration of EC2 instances. uses 'Chef' and 'Puppet'.
10. Transit Gateway ?
11. Global Accelerator
12. MFA Delete? - yes this exists. it protects S3 objects from accidental deletes
13. Lambda edge - running code closer to the user (inside cloudFront) to get better performance and less data costs.
14. FSx -
15. Active MQ - standalone message broker queue, 
16. AMAZON MQ - message broker queue, but for existing (not aws native) applications.
17. Resource Access manager? - share resources with other accounts
18. SCP - service control policy - under AWS organization.
19. OAI - origin access policy - sits on the objects (like S3 bucket) and controls who can access it.
20. DynamoDAX - Caching Layer on top of dynamoDB
21. Direct connect(DX)
    1.  VIF?
22. Private link? what is classic link? - private link is "endpoint service".
23. Internet Gateway
24. CloudWatch custom metrics (is there a cloudWatch agent?)
25. VPC endpoint?
26. AWS Athena - serverless, optimized, managed way to query **S3 objects**, pay-as-you-use, ad hoc.
27. Inspector - "Amazon Inspector is an automated vulnerability management service that continually scans AWS workloads for software vulnerabilities and unintended network exposure." - *Can I be attacked?*
28. GuardDuty - "Amazon GuardDuty is a threat detection service that continuously monitors your AWS accounts and workloads for malicious activity and delivers detailed security findings for visibility and remediation." - *Was I attacked?*
29. Macie - "Amazon Macie is a fully managed data security and data privacy service that uses machine learning and pattern matching to discover and protect your sensitive data in AWS." - can't be used in RDS?
30. Cognito - "Amazon Cognito lets you add user sign-up, sign-in, and access control to your web and mobile apps quickly and easily. Amazon Cognito scales to millions of users and supports sign-in with social identity providers, such as Apple, Facebook, Google, and Amazon, and enterprise identity providers via SAML 2.0 and OpenID Connect."
31. CNAME vs ALIAS?
32. SSE-S3 keys, how are they different than other keys? when to use SSE-S3 and when to use KMS?
33. event bridge rules?
34. S3 transfer acceleration?
35. cross region replication vs cross region resource sharing (CORS)
36. CloudWatch in EKS?
37. Auto Scaling groups
   1. simple - wait for health checks, cooldown periods
   2. target - try to keep as close as possible to a metric
   3. step - fine control over how adjustments are made.
38. Raid 0
39. NAT gatway vs NAT interface
40. DynamoDB streams? - "DynamoDB Streams captures a time-ordered sequence of item-level modifications in any DynamoDB table and stores this information in a log for up to 24 hours." auditing of dynamoDB events.
41. Redshift spectrum? - "Amazon RedShift Spectrum is a feature of Amazon Redshift that enables you to run queries against exabytes of unstructured data in Amazon S3, with no loading or ETL required."
42. Step functions?
43. RDS replication types?
    1.  synchronous
    2.  scheduled
    3.  asynchronous
    4.  continues
44. CodeCommit
45. CodeStar
46. Cacheing over a method/stage? - API gateway caches over a stage.
47. CodeDeploy - "AWS CodeDeploy is a fully managed deployment service that automates software deployments to a variety of compute services such as Amazon EC2, AWS Fargate, AWS Lambda, and your on-premises servers. AWS CodeDeploy makes it easier for you to rapidly release new features, helps you avoid downtime during application deployment, and handles the complexity of updating your applications. You can use AWS CodeDeploy to automate software deployments, eliminating the need for error-prone manual operations. The service scales to match your deployment needs."
48. System Manager - "Amazon Systems Manager is a management service that helps you automatically collect software inventory, apply OS patches, create system images, and configure Windows and Linux operating systems. These capabilities help you define and track system configurations, prevent drift, and maintain software compliance of your EC2 and on-premises configurations. By providing a management approach that is designed for the scale and agility of the cloud but extends into your on-premises data center, Systems Manager makes it easier for you to seamlessly bridge your existing infrastructure with Amazon Web Services."
49. what does dashboard display?
50. Virtual Private Gateway?
51. What Are SQS auto scaling queues?
52. CloudWatch metrics per service
53. When is server migration service used?
54. MPLS?
55. RTMP?
56. MetaData querying tool?
57. Dedicated Instance vs dedicated hosts - dedicated instances are seperated from other accounts at physical level (pay per instance). dedicated hosts is more 'stringent' and more 'separated', it's a real physical machine solely dedicated to you (per per host).
58. When to use cognito and when SAML\tokens?
 



## Tips:

Service limits are seen in the trusted advisor.


transit gateway isn't only used in DX, it can connect many VPCs together (instead of peering), and just brings together VPC and other connection points like VPN or DX.

RTMP must be stored in S3.

### Storage transitions: 
move from instance Store to EBS (persistnancy), and from EBS to EFS (multiple attachments, only multi-attached nitro EBS can be attached to multiple instances), and then maybe to FSx (fully managed?, works with SMB, VPNs).

Lustre works with High performance and works natively with S3 objects. not integrated with fargate (EFS is)

EFS has lifecycle policies as well.

RAID 0 is for performance, RAID1 is for fault tolerance.

all EBS families support encryption

Ebs snapshots are incremental, but deletion will make sure the recent one contains all the data.

### Redshift

normal redshift is managed and serverless, but redshift spectrum isn't serverless, it requires a redshift cluster.

### Direct Connect (DX)
- Gateways:
  - file
  - volume
  - tape
- Encryption requires a VPN tunnel (VPG)
Transit gateway can connect one DX through a transit virtual interface to many VPC.

New Name	| Old Name|	Interface	|Use Case
---|---|---|---
File Gateway	|None	|NFS, SMB	|Allow on-prem or EC2 instances to store objects in S3 via NFS or SMB mount points
Volume Gateway Stored Mode	|Gateway-Stored Volumes|	iSCSI	|Asynchronous replication of on-prem data to S3
Volume Gateway Cached Mode|	Gateway-Cached Volumes|	iSCSI	|Primary data stored in S3 with frequently accessed data cached locally on-prem
Tape Gateway|	Gateway-Virtual Tape Library|	ISCSI|	Virtual media changer and tape library for use with existing backup software

### Route53:
- geolocation is more specific than geoProximity.


Alias records are used to map resource record sets in your hosted zone to:
- Amazon Elastic Load Balancing load balancers 
- API Gateway custom regional APIs and edge-optimized APIs
- CloudFront Distributions
- AWS Elastic Beanstalk environments
- Amazon S3 buckets that are configured as website endpoints
- Amazon VPC interface endpoints
- other records in the same Hosted Zone.

### internet Gateway - vpc
- private connection requires interface VPC endpoint (privateLink)
- interface endpoints are for resources like vpc, elb. gateway resources are for S3 and dynamoDB

### elasticCache
- Redis has more features than memcached



### Compute

Batch jobs are run across several EC2 instances, for parallel jobs.

EMR - Elastic map reduce - same as hadoop - big data

### API Gateway
[CheatSheet](https://digitalcloud.training/amazon-api-gateway)

Caching is per stage

throteling can be done per client or per method.

#### Endpoints
limited access 
- edge optimized - from cloudfront
- regional endpoint - services in the same region.
- private endpoint - from the same vpc

### RDS

Aurora can be instance based or serverless.

Multi-AZ uses synchronously replication.

Action |Multi-AZ Deployments |	Read Replicas
---|---|---
Replication|Synchronous Replication – highly durable	|Asynchronous replication – highly scalable
Active engines |Only database engine on primary instance is active | All read replicas are accessible and can be used for read scaling
backups | Automated backups are taken from standby| No backups configured by default
Availability |Always span two availability zones within a single region |	Can be within an Availability Zone, Cross-AZ, or Cross-Region
versioning | Database engine version upgrades happen on primary|Database engine version upgrade is independent from source instance
fault tolerance |Automatic failover to standby when a problem is detected | Can be manually promoted to a standalone database instance


there is a special aws authentication plugin to MySQL.

### DynamoDB
[DynamoDB cheatsheet](https://digitalcloud.training/amazon-dynamodb/)

max object size is 400Kb

read/write units on a partiton key. max 3000RCU (reads) or 1000WCU (writes)

a best practices is to store more frequently and less frequently accessed data in separate tables

### Beanstalk
[BeanStalk cheatSheet](https://digitalcloud.training/aws-elastic-beanstalk/)

because beanstalk creates entire application stacks (like a wordpress site), we can point an Alias record to them.

### ElasticCache

Redis is usually stronger, except that it's not multithreaded! memCacheD supports multicore!

### SQS
Fifo sqs ensures order and ensures that there a no duplicates.