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
6. When is AWS Shield Advanced Used?
7. What is AWS Config?
8. S3 Vault Lock - **unchangeable policy** to S3 glacier. stronger than access policy, used for compliance. for example, allows to prevent deletion of objects.
9. Support Plan features
10. What is OpsWork - automated configuration of EC2 instances. uses 'Chef' and 'Puppet'.
11. Transit Gateway - a way to connect multiple VPC, VPN and stuff together. like extra vpc peering, also in DX.
12. Global Accelerator
13. MFA Delete? - yes this exists. it protects S3 objects from accidental deletes
14. Lambda edge - running code closer to the user (inside cloudFront) to get better performance and less data costs.
15. FSx - FSx windows and FSx Lustre. windows is fore general use like a hyped up EFS, lustre is for HPC.
16. Active MQ - standalone message broker queue, 
17. AMAZON MQ - message broker queue, but for existing (not aws native) applications.
18. Resource Access manager? - share resources with other accounts
19. SCP - service control policy - under AWS organization.
20. OAI - origin access policy - sits on the objects (like S3 bucket) and controls who can access it.
21. DynamoDAX - Caching Layer on top of dynamoDB
22. Direct connect(DX)
    1.  VIF?
23. Private link? what is classic link? - private link is "endpoint service".
24. Internet Gateway
25. CloudWatch custom metrics (is there a cloudWatch agent?)
26. VPC endpoint?
27. AWS Athena - serverless, optimized, managed way to query **S3 objects**, pay-as-you-use, ad hoc.
28. Inspector - "Amazon Inspector is an automated vulnerability management service that continually scans AWS workloads for software vulnerabilities and unintended network exposure." - *Can I be attacked?*
29. GuardDuty - "Amazon GuardDuty is a threat detection service that continuously monitors your AWS accounts and workloads for malicious activity and delivers detailed security findings for visibility and remediation." - *Was I attacked?*
30. Macie - "Amazon Macie is a fully managed data security and data privacy service that uses machine learning and pattern matching to discover and protect your sensitive data in AWS." - can't be used in RDS?
31. Cognito - "Amazon Cognito lets you add user sign-up, sign-in, and access control to your web and mobile apps quickly and easily. Amazon Cognito scales to millions of users and supports sign-in with social identity providers, such as Apple, Facebook, Google, and Amazon, and enterprise identity providers via SAML 2.0 and OpenID Connect."
32. SSE-S3 keys, how are they different than other keys? when to use SSE-S3 and when to use KMS? -  SSE-S3 enctypes each object differently, KMS uses several keys, also charge you for it.
33. event bridge rules?
34. S3 transfer acceleration? - use edge locations for speed increase. massive uploads / downloads.
35. cross region replication vs cross region resource sharing (CORS)
36. CloudWatch in EKS?
37. Auto Scaling groups
   1. simple - wait for health checks, cooldown periods
   2. target - try to keep as close as possible to a metric
   3. step - fine control over how adjustments are made.
38. Raid 0
39. NAT gatway vs NAT interface
40. DynamoDB streams? - "DynamoDB Streams captures a time-ordered sequence of item-level modifications in any DynamoDB table and stores this information in a log for up to 24 hours." auditing of dynamoDB events.
41. Step functions - pipeline automation
42. CodeCommit - version control service
43. CodeStar - managing software development, integrates with IDS, build steps, etc
44. Cacheing over a method/stage? - API gateway caches over a stage.
45. CodeDeploy - "AWS CodeDeploy is a fully managed deployment service that automates software deployments to a variety of compute services such as Amazon EC2, AWS Fargate, AWS Lambda, and your on-premises servers. AWS CodeDeploy makes it easier for you to rapidly release new features, helps you avoid downtime during application deployment, and handles the complexity of updating your applications. You can use AWS CodeDeploy to automate software deployments, eliminating the need for error-prone manual operations. The service scales to match your deployment needs."
46. System Manager - "Amazon Systems Manager is a management service that helps you automatically collect software inventory, apply OS patches, create system images, and configure Windows and Linux operating systems. These capabilities help you define and track system configurations, prevent drift, and maintain software compliance of your EC2 and on-premises configurations. By providing a management approach that is designed for the scale and agility of the cloud but extends into your on-premises data center, Systems Manager makes it easier for you to seamlessly bridge your existing infrastructure with Amazon Web Services."
47. what does dashboard display?
48. Virtual Private Gateway?
49. What Are SQS auto scaling queues?
50. CloudWatch metrics per service
51. When is server migration service used?
52. MPLS - "Multiprotocol Label Switching, or MPLS, is a networking technology that routes traffic using the shortest path based on “labels,” rather than network addresses, to handle forwarding over private wide area networks."
53. RTMP - media on Cloudfront, for videos and stuff.
54. MetaData querying tool?
55. Dedicated Instance vs dedicated hosts - dedicated instances are seperated from other accounts at physical level (pay per instance). dedicated hosts is more 'stringent' and more 'separated', it's a real physical machine solely dedicated to you (per per host).
56. When to use cognito and when SAML\tokens?
57. what does EMR work on? files? databases? structured?
58. S3 select?
59. Amazon Neptune - graph database
60. AWS Lex - ""Amazon Lex" is incorrect. Amazon Lex is a service for building conversational interfaces into any application using voice and text."
61. AWS X-Ray - "AWS X-Ray lets you analyze and debug serverless applications by providing distributed tracing and service maps to easily identify performance bottlenecks by visualizing a request end-to-end."
62. IOT Core - "AWS IoT Core is a managed cloud service that lets connected devices easily and securely interact with cloud applications and other devices. AWS IoT Core can support billions of devices and trillions of messages, and can process and route those messages to AWS endpoints and to other devices reliably and securely."
63. Glue - "AWS Glue is a fully managed extract, transform, and load (ETL) service that makes it easy for customers to prepare and load their data for analytics."
 



## Tips:


[Penetration Testing on AWS](https://aws.amazon.com/security/penetration-testing/) - allowed on our resources, and only some of them. some actions aren't allowed.\
Permitted Services:
- Amazon EC2 instances, NAT Gateways, and Elastic Load Balancers
- Amazon RDS
- Amazon CloudFront
- Amazon Aurora
- Amazon API Gateways
- AWS Fargate
- AWS Lambda and Lambda Edge functions
- Amazon Lightsail resources
- Amazon Elastic Beanstalk environments


Service limits are seen in the trusted advisor.


transit gateway isn't only used in DX, it can connect many VPCs together (instead of peering), and just brings together VPC and other connection points like VPN or DX.

RTMP must be stored in S3.

RunCommand - installing, running cli commands, configuration of active servers.

### Storage transitions: 
move from instance Store to EBS (persistnancy), and from EBS to EFS (multiple attachments, only multi-attached nitro EBS can be attached to multiple instances), and then maybe to FSx (fully managed?, works with SMB, VPNs).

Lustre works with High performance and works natively with S3 objects. not integrated with fargate (EFS is)

EFS has lifecycle policies as well.

RAID 0 is for performance, RAID1 is for fault tolerance.

all EBS families support encryption

Ebs snapshots are incremental, but deletion will make sure the recent one contains all the data.



### internet Gateway - vpc
- private connection requires interface VPC endpoint (privateLink)
- interface endpoints are for resources like vpc, elb. gateway resources are for S3 and dynamoDB




### Compute

Batch jobs are run across several EC2 instances, for parallel jobs.

EMR - Elastic map reduce - same as hadoop - big data



#### Endpoints
limited access 
- edge optimized - from cloudfront
- regional endpoint - services in the same region.
- private endpoint - from the same vpc

### SQS
Fifo sqs ensures order and ensures that there a no duplicates.


### RDS - relational database
 [RDS](https://digitalcloud.training/amazon-rds/) - relational database

Multi-AZ uses synchronously replication.

| Action          | Multi-AZ Deployments                                      | Read Replicas                                                       |
| --------------- | --------------------------------------------------------- | ------------------------------------------------------------------- |
| Replication     | Synchronous Replication – highly durable                  | Asynchronous replication – highly scalable                          |
| Active engines  | Only database engine on primary instance is active        | All read replicas are accessible and can be used for read scaling   |
| backups         | Automated backups are taken from standby                  | No backups configured by default                                    |
| Availability    | Always span two availability zones within a single region | Can be within an Availability Zone, Cross-AZ, or Cross-Region       |
| versioning      | Database engine version upgrades happen on primary        | Database engine version upgrade is independent from source instance |
| fault tolerance | Automatic failover to standby when a problem is detected  | Can be manually promoted to a standalone database instance          |


there is a special aws authentication plugin to MySQL.



### Aurora  - fully managed serverless RDS
 [Aurora](https://digitalcloud.training/amazon-aurora/) - fully managed serverless RDS

Aurora can be instance based or serverless.


### DynamoDB - document database (like mongo)
[DynamoDB cheatsheet](https://digitalcloud.training/amazon-dynamodb/)

max object size is 400Kb

read/write units on a partiton key. max 3000RCU (reads) or 1000WCU (writes)

a best practices is to store more frequently and less frequently accessed data in separate tables.

also has auto scaling.


### Redshift - data warehouse for analytics
 [Redshift](https://digitalcloud.training/amazon-redshift/) 

"Amazon RedShift Spectrum is a feature of Amazon Redshift that enables you to run queries against exabytes of unstructured data in Amazon S3, with no loading or ETL required."
  
normal redshift is managed and serverless, but redshift spectrum isn't serverless, it requires a redshift cluster.

### ElasticCache  - Caching Layer
[ElasticCache](https://digitalcloud.training/amazon-elasticache/)


Redis has more features than memcached

Redis is usually stronger, except that it's not multithreaded! memCacheD supports multicore!
 
### Lambda
 [Lambda](https://digitalcloud.training/aws-lambda/) - serverless compute action


### EC2
 [EC2](https://digitalcloud.training/amazon-ec2/) - virtual machines
### Auto Scaling
 [EC2 Auto scaling](https://digitalcloud.training/amazon-ec2-auto-scaling/) - increased performance.
###
 [ECS and EKS](https://digitalcloud.training/amazon-ecs-and-eks/) - run workloads in containers

### EBS
 [EBS](https://digitalcloud.training/amazon-ebs/) - elastic block storage, persistent
### S3
 [S3 and Glacier](https://digitalcloud.training/amazon-s3-and-glacier/) - object store buckets
### Athena
 [Athena](https://digitalcloud.training/amazon-athena/) - Query S3.
### FSX
 [FSX](https://digitalcloud.training/amazon-fsx/) - fully managed file servers
### EFS
 [EFS](https://digitalcloud.training/amazon-efs/) - elastic file system. NFS protocol

### ELB - elastic load balancing
 [ELB CheatSheet](https://digitalcloud.training/aws-elastic-load-balancing-aws-elb/) - elastic load balancing - disturbing load across resources.
### CloudFront
 [CloudFront CheatSheet](https://digitalcloud.training/amazon-cloudfront/) - content distribution
###
 [Global Accelerator CheatSheet](https://digitalcloud.training/aws-global-accelerator/) - improved networking
###
 [WAF & Shield CheatSheet](https://digitalcloud.training/aws-waf-shield/) - protecting from web exploits and attack.


### Route53 - DNS Services
[Route53 CheatSheet](https://digitalcloud.training/amazon-route-53/) 

CNAME vs ALIAS? - CNAME is standard, ALIAS isn't, (A for address is)
  1.  A - maps record name to to ip address
  2.  CNAME - maps name to name, requires uniqueness
  3.  ALIAS - map name to name, doesn't require uniqueness
  4.  URL - redirecting, rather than resolving


- geolocation is more specific than geoProximity.

ALIAS records are used to map resource record sets in your hosted zone to:
- Amazon Elastic Load Balancing load balancers 
- API Gateway custom regional APIs and edge-optimized APIs
- CloudFront Distributions
- AWS Elastic Beanstalk environments
- Amazon S3 buckets that are configured as website endpoints
- Amazon VPC interface endpoints
- other records in the same Hosted Zone.



### Direct Connect DX
 [AWS Direct Connect CheatSheet](https://digitalcloud.training/aws-direct-connect/) - connect on-premises data center to cloud.

- Gateways:
  - file
  - volume
  - tape
- Encryption requires a VPN tunnel (VPG)
Transit gateway can connect one DX through a transit virtual interface to many VPC.

| New Name                   | Old Name                     | Interface | Use Case                                                                        |
| -------------------------- | ---------------------------- | --------- | --------------------------------------------------------------------------------- |
| File Gateway               | None                         | NFS, SMB  | Allow on-prem or EC2 instances to store objects in S3 via NFS or SMB mount points |
| Volume Gateway Stored Mode | Gateway-Stored Volumes       | iSCSI     | Asynchronous replication of on-prem data to S3                                    |
| Volume Gateway Cached Mode | Gateway-Cached Volumes       | iSCSI     | Primary data stored in S3 with frequently accessed data cached locally on-prem    |
| Tape Gateway               | Gateway-Virtual Tape Library | ISCSI     | Virtual media changer and tape library for use with existing backup software      |

###
 [Storage Gateway CheatSheet](https://digitalcloud.training/aws-storage-gateway/) - hybrid storage between the cloud and the on-premises environment.
 
###
 [Migration services CheatSheet](https://digitalcloud.training/aws-migration-services/) - database, server, syncing.
###
 [VPC CheatSheet](https://digitalcloud.training/amazon-vpc/) - logical partition of resources in the region

### API Gateway - manage Rest API
[API Gateway CheatSheet](https://digitalcloud.training/amazon-api-gateway)

Caching is per stage

throteling can be done per client or per method.


###
 [CloudWatch CheatSheet](https://digitalcloud.training/amazon-cloudwatch/) - monitoring tool
###
 [CloudTrail CheatSheet](https://digitalcloud.training/aws-cloudtrail/) - auditing
###
 [AWS Config CheatSheet](https://digitalcloud.training/aws-config/) - history of configuration changes - similar to cloud watch, but about changes, not actions.
### Beanstalk - deploy applications
[BeanStalk cheatSheet](https://digitalcloud.training/aws-elastic-beanstalk/)

because beanstalk creates entire application stacks (like a wordpress site), we can point an Alias record to them.

###
 [Kinesis CheatSheet](https://digitalcloud.training/amazon-kinesis/) - high performance real time data entry point.
###
 Managerial Services
###
 [IAM CheatSheet](https://digitalcloud.training/aws-iam/) - identity access, policies
###
 [AWS Organization CheatSheet](https://digitalcloud.training/aws-organizations/) - manage many AWS accounts in the same organization
###
 [KMS](https://digitalcloud.training/aws-kms/) - key management.
###
 [Application integration Services](https://digitalcloud.training/aws-application-integration-services/) - services which communicate between other services
###
 [Additional Services CheatSheet](https://digitalcloud.training/additional-aws-services/) - Glue, Data Pipeline, EMR