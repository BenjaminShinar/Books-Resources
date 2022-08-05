<!--
// cSpell:ignore
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
- On premises access
  - [AWS Direct Connect](https://digitalcloud.training/aws-direct-connect/) - connect on-premises data center to cloud.
  - [Storage Gateway](https://digitalcloud.training/aws-storage-gateway/) - hybrid storage between the cloud and the on-premises environment.
  - [Migration service](https://digitalcloud.training/aws-migration-services/) - database, server, syncing.
- [Application integration Services](https://digitalcloud.training/aws-application-integration-services/) - services which communicate between other services
- [VPC](https://digitalcloud.training/amazon-vpc/) - logical partition of resources in the region
- [CloudWatch](https://digitalcloud.training/amazon-cloudwatch/) - monitoring tool
- [Kinesis](https://digitalcloud.training/amazon-kinesis/) - high performance real time data entry point.
- Managerial Services
  - [IAM](https://digitalcloud.training/aws-iam/) - identity access, policies
  - [AWS Organization](https://digitalcloud.training/aws-organizations/) - manage many AWS accounts in the same organization


Open Questions:
1. When is aws VPN used?
2. Are ELB in a single AZ or multiple?
3. What are these?
   - Pilot Light
    - Warm Standby
    - Multi-site
    - Backup-Restore
4. Penetration Testing on AWS
5. When is AWS Shield Advanced Used?
6. What is AWS Config?
7. S3 Vault Lock - **unchangeable policy** to S3 glacier. stronger than access policy, used for compliance. for example, allows to prevent deletion of objects.
8. Support Plan features
9. What is OpsWork - automated configuration of EC2 instances. uses 'Chef' and 'Puppet'.
10. Transit Gateway ?
11. Global Accelerator
12. MFA Delete? - yes this exists. it protects S3 objects from accidental deletes
13. Lambda edge
14. FSx -
15. Active MQ - standalone message broker queue, 
16. AMAZON MQ - message broker queue, but for existing (not aws native) applications.
17. Resource Access manager?
18. SCP - service control policy - under AWS organization.
19. OAI?
20. DynamoDAX - Caching Layer on top of dynamoDB
21. Direct connect(DX)
    1.  VIF?
22. Private link?
23. Internet Gateway
24. CloudWatch custom metrics (is there a cloudWatch agent?)
25. VPC endpoint?
26. AWS Athena - serverless, optimized, managed way to query **S3 objects**, pay-as-you-use, ad hoc.



## Tips:
Storage transitions: move from instance Store to EBS (persistnancy), and from EBS to EFS (multiple attachments, only multi-attached nitro EBS can be attached to multiple instances), and then maybe to FSx (fully managed?).


Gateways:
- file
- volume
- tape


Route53:
- geolocation is more specific than geoProximity.

internet Gateway - vpc