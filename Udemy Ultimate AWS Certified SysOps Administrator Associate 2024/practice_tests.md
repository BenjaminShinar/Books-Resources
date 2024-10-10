<!--
ignore these words in spell check for this file
// cSpell:ignore NACL
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# SysOps Exam Preperations

Getting ready for the exam.

Udemy [Practice Exams: AWS Certified SysOps Administrator Associate](https://checkpoint.udemy.com/course/practice-exams-aws-certified-sysops-administrator-associate/) and [AWS SysOps Administrator Associate SOA-C02 Practice Exam](https://checkpoint.udemy.com/course/aws-sysops-administrator-associate-practice-exams-soa-c02/learn), [AWS Certified SysOps Administrator Associate Practice Exams](https://checkpoint.udemy.com/course/aws-certified-sysops-administrator-associate-aws-practice-exam)




## Takeaways

1. Application Load balancers do not send data to X-Ray
1. <cloud>CloudFormation</cloud> StackSets allow you to roll out CloudFormation stacks over multiple AWS accounts and in multiple Regions with just a couple of clicks. When AWS launched StackSets, grouping accounts was primarily for billing purposes. Since the launch of AWS Organizations, you can centrally manage multiple AWS accounts across diverse business needs including billing, access control, compliance, security and resource sharing.
1. Q: what does the <cloud>Elastic Load Balancer</cloud> do when all targets are unhealthy?\
    *A: If all are unavailable, it sends to them anyway, it's best effort*
1. Q: Can we move data directly into <cloud>AWS Glacier</cloud>?\
    *A: seems so, but not from Snowball edge devices*
1. Q: <cloud>CloudFormation</cloud> Stack Policies?\
    *A: protects stacks against updates, we define what can be updated later*
1. Q: <cloud>EC2</cloud> bandwidth metrics?\
    *A: `Network In/Out`*
1. Q: is there such as thing as <cloud>ServiceLens</cloud>?\
    *A: Amazon CloudWatch ServiceLens is a new feature that enables you to visualize and analyze the health, performance, and availability of your applications in a single place.*
3. Q: <cloud>S3</cloud> origin access Identity?\
    *A: only allow access to the bucket from a specific source, not for static S3 websites with cloudFront*
4. Q: <cloud>Elastic BeanStalk</cloud> health check configuration?\
    *A: default is checking EC2 status, not ELB, if we want AutoScaling group to replace them, we need to set health checks on ELB*
5. Q: <cloud>SSM Patch manager</cloud>?\
    *A: patching EC2 and on-premises machines, linux, macOs and Windows, discover which patches are installed, choose what, when and from where to install patches*
6. Q: What happens with <cloud>EC2</cloud> termination protection?
7. Q: <cloud>AWS Config</cloud> vs <cloud>AWS Trusted Advisor</cloud>?\
    *A: Trusted advisor is AWS best practices - cost, security, fault toleranace, etc... AWS Config is compliance, built-in rules and custom rules*
8. Q: When is the <cloud>CloudWatch Agent</cloud> installed?
9. Q: Possible Errors with <cloud>RDS</cloud> replication
10. Q: Traffic between EC2 instances in different regions.
11. Q: <cloud>CloudFormation</cloud> StackSets\
    *A: deploy stacks across multiple accounts, either with self managed IAM role or through AWS organizations*
12. Q: CloudWatch agent with multiple configuration files
13. Q: when do we use <cloud>EC2</cloud> enhanced networking
14. Q: Problem with "attaching" <cloud>EBS</cloud> volumes
15. Q: <cloud>CloudFormation</cloud> ChangeSets?\
    *A: A preview of what's about to change when we deploy the new template, like `terraform plan`*
16. Q: Storage Gateway security?
17. Q: How to configure the unified CloudWatch Agent?\
    *A: it's possible to have multiple files, as long as they have different names (not just different paths) if they have the same name, they can't be appended together, and the files overwrite*
18. Q: Elastic Bean Stalk default Auto-scaling behavior?\
    *A: based on `NetworkOut` <cloud>EC2</cloud> metric for scale-out and scale-in*
19. Q: <cloud>EC2</cloud> Capacity Reservations?
20. Q: AutoScaling Group errors?
21. Q: Special SSM documents?\
    - *`AWS-ApplyPatchBaseline`*
    - *`AWSSupport-TroubleshootS3PublicRead`*
22. Q: Who doesn't send data to <cloud>X-Ray</cloud>?\
    *A: Application load Balancer*
24. Q: EC2 status Checks?\
    *A: status checks can't be disabled*
25. Q: <cloud>CloudWatch</cloud> Synthetics?\
    *A: Configurable script that monitor your APIs, URLs, Websites*
26. Q: Cross account Resource Access\
    *A: requires permissions in both accounts*
28. Q: Subnets communication?\
    *A: subnets inside the same VPC can communicate with one another by default, no need for special setups (security groups are still required)*
29. Q: creating AMIs\
    *A: AMIs can be created from EBS and from EC2 instances (`no-reboot` option to create AMI without shutdown)*
30. if all of <cloud>Aurora</cloud> replicas are in a downed AZ, we must create a new instance manually.
31. <cloud>S3</cloud> buckets can have policies which deny access to the root account. the root account can modify the policy.
32. Q: what does the "Client.InternalError: Client error on launch" mean for AutoScaling Groups?\
    *A:* attempt to launch an instance from an encrypted <cloud>EBS</cloud> volume without the proper CMK.
33. Automated <cloud>EBS</cloud> snapshots requires an <cloud>CloudWatch Event</cloud> or <cloud>Event Bridge</cloud> set-up.
34. S3 lifecycles will not prevent accidental deletions, it needs object lock settings.
35. <cloud>Directory Service</cloud> requires creating a security group with unrestricted access
36. <cloud>S3</cloud> server access logging - no extra charges beyond those of the data storage.
37. Using custom names in <cloud>CloudFormation</cloud> stacks can cause resources to not be created
38. public IPv4 addresses aren't retained after changing instances.
39. Gemalto MFA devices can't be reused accross companies.
40. `RAMUtilization` Metric must be manually inserted!
41. Auto scaling group termination suspended: can't terminate instances, but can grow to 110% of maximum size.
42. Termination Protection?
43. <cloud>GuardDuty</cloud> - Threat Detection
44. No such thing as IAM Usage Report!
45. <cloud>CloudWatch</cloud> dashboards are global, and can include metrics from different regions and accounts!
46. in place encryption without impacts?
47. Q: failing to delete an "Empty" <cloud>S3</cloud> bucket?\
    *A: versioning is enabled, delete markers still exist*
49. When are <cloud>EC2</cloud> instances replace by auto-scaling group?
50. Volume vs File Gateway?\
    *A: Volume gateway is gateway for <cloud>EBS</cloud>, File Gateway exposes S3 as filesystem.*
52. AWS Inspector?\
   *A: automatic security checks - network access, known vunrabilites (package CVEs)* 
53. Q: How to debug <cloud>CloudFormation</cloud> init scripts?\
    *A: setting `OnFailure=Do_Nothing` will keep the instance, so the logs could still be accessed. no need for special IAM role to signal back, but needs the correct network routes*
55. Sharing AMIs across accounts? - no such thing "just seeing", if it's visible, it's usable.
56. `InsufficientInstanceCapacity` error?\
    *A: AWS can't create the requested instance, wait or change instance or AZ*
57. Q: ELB metrics?\
    *- SurgeQueueLength: The total number of requests (HTTP listener) or connections (TCP listener) that are pending routing to a healthy instance. Help to scale out ASG. Max value is 1024*
    *- SpillOverCount: The total number of requests that were rejected because the surge queue is full.*
    *- ActiveConnectionCount: number of TCP connections *
1. Elastic Load Balancer can be pre-warmed
59. Q: Is there a grace period for Autoscaling groups?\
    *A: the grace period is the time until the first health check is performed, it does not matter for already running instances that become unhealthy*
60. > <cloud>AWS OpsWorks</cloud> is a configuration management service that provides managed instances of Chef and Puppet. Chef and Puppet are automation platforms that allow you to use code to automate the configurations of your servers. OpsWorks lets you use Chef and Puppet to automate how servers are configured, deployed and managed across your Amazon EC2 instances or on-premises compute environment
61. T2 unlimited means we can sustain high CPU performance over a long time.
62. If we want increased performance from fully utilized gp2 drives without increased costs, we can split it into smaller devices and mount them together on RAID-0. RAID-1 is for fault tolerance.
---
1. we can reach <cloud>S3</cloud> through <cloud>Direct Connect</cloud> by using dedicated/host connection, cross network connections and public virtual interface
2. policies interactions? - explicit deny wins and overrides anything else.
3. <cloud>S3</cloud> encryption headers?\
    the request must include the header,
    - `'x-amz-server-side-encryption': 'AES256'` - use default encryption
   - `'x-amz-server-side-encryption': 'aws:kms'` - use sse-kms
4. <cloud>VPC Flow Logs</cloud> `access error` happens when we don't have the permissions, trust relationship or the correct service as principal. we can't modify the flow logs configuration, we must delete and re-create it.
5. <cloud>RDS</cloud> replication and read-replicas - which is asynchronous?\
   - read replicas are synchros
   - replication is asynchronous - same AZ, different AZ in the same region or even different account
6. is <cloud>Aurora</cloud> integrated with <cloud>CloudWatch</cloud>? how to see logs, metrics, audits?\
    not on by default, needs to be enabled through the parameter group, differnet behavior for MySQL and PostgresSQL databases.
8. IOPs have a max ratio of 1:50. for every 1Gb, we can get up to 50 Iops. so for 100GB, max is 5000 IOPS.
9.  is there SSM audit? No such thing, we use <cloud>CloudTrail</cloud> for SSM and KMS.
10. Dedicated Hosts vs Dedicated Instance?\
    - Dedicated Hosts gives us a physical server, this is important when we have compliance requirements and we need to use existing software licensees which are bound to the hardware itself. this can be on-demand or reserved. this is the most expensive option in AWS. we can use **per-socket** stuff here\
    - Dedicate instances - no other accounts can run on the same hardware as these EC2 instances. but it can share the hardware with instances from the same account, and we have no control over instance placement.
11. <cloud>S3</cloud> replication failure notification?\
    we use S3 Replication Time Control to get replication metrics and notifications.
12. <cloud>CloudFormation</cloud> stack termination protections protects against accidental deletion of stacks, the protection must first be removed, otherwise deletion process will not happen (no failure message)
13. for encrypted connection, we need both Direct Connect and VPN
14. EC2 termination policies, `DisableApiTermination` can't be used with spot instances, but does not prevent auto-scaling group from terminating the instance. we can use "InstanceProtection" for that.
15. <cloud>CloudFront</cloud> communicates with <cloud>S3</cloud> using the HTTP protocol when configured as a website, but not for other cases. otherwise, it can use the same protocol as the original request.
16. <cloud>S3</cloud> storage lens?\
    - S3 Storage Lens is the first cloud storage analytics solution to provide a single view of **object storage usage and activity across** hundreds, or even thousands, of accounts in an organization, with drill-downs to generate insights at multiple aggregation levels.
    - S3 inventory creates a CSV (or apache compatible) file with metadata about the files in the bucket. this can include replication status.
17. Default NACL configuration (inbound, outbound) - by default the allow all traffic (inbound, outbound), they are optional above security groups. every subnet has an ACL.
18. Load balancers must have one or more listeners, a target can be registered to multiple target groups at the same time.
19. what can we do when terminating Spot Instances? what are the possible interuption behaviors?
    - can be terminated (default)
    - can be stopped
    - can hibernated
    - **can't be rebooted**
21. <cloud>SSM Parameter Store</cloud> and <cloud>KMS</cloud> "pending deletion"?\
    - KMS has pending deletion state, between 7 to 30 days period, default is 30.
    - Secret Manager also has something, as seen by the metrics:`StartSecretVersionDelete` ,`CancelSecretVersionDelete`, `EndSecretVersionDelete`. there is a 7 days recovery window.
22. Trust Policy is a resource based policy.
23. S3 SSE-C does not have audit trails for encryption key usage.
24. <cloud>SQS</cloud> metrics dimensions? - Queue name alone is enough, no need for queue Id (arn)
25. DataLifeCycle Manager is for <cloud>EBS</cloud>, <cloud>AWS Backup</cloud> is for most resources
26. Organization Trail? - can be seen by member accounts, but not modified.
27. Cloud Trail only tracks S3 bucket level actions by default, for object-level actions audit, we need to enable data events trail. 
28. Where in <cloud>CloudFront</cloud> templates can we and can't we use conditions?
    - can't be used on parameters, can be used on resources, outputs, etc...
29. what's required when using <cloud>S3</cloud> retention rules?
---
1. EFS metrics? - `ClientConnections` to measure how many clients.
2. Manage multiple <cloud>Snowball</cloud> devices? use <cloud>OpsHub</cloud>
3. Service Catalog Sharing? use imported portfolios
4. WAF and <cloud>CloudFront</cloud> we can put <cloud>WAF</cloud> around cloudfront,and use the manager to replicate it.
5. with a network load balancer, an AZ can't be disabled, they can be added, but not removed.
6. <cloud>ACM</cloud> must be loaded into the same region as the load balancer.
7. EFS Mount Helper?
   - Auto mounting on instance reboot
   - mounting with <cloud>IAM</cloud> authorization
8. Replacing Disk on Storage Gateway - need to create a new gateway to change cache, but not when changing buffer.
9.  AWS Control Tower? Manage Multiple AWS accounts and AWS environments
10. <cloud>System Manager</cloud> allows to control multiple EC2 instances
11. Redis instances can be manually promoted only when Multi-az and failover are disabled, replication across AZs is asynchronous. redis prefers to promote the instance with the least replication lag.
12. when a different account creates objects in an <cloud>S3 bucket</cloud>, the bucket owner doesn't have permissions on them.
13. Auroras are still inside a VPC
14. we can't install amazon issued certificates onto <cloud>EC2</cloud> machine, we need to use third party SSL. ACM can't be used for email encryption
15. when <cloud>EBS</cloud> volume is in `error` status, it means hardware failure.
16. for programmatic access and integrations, there is the <cloud>AWS Health API</cloud>.
17. when we get a domain for cloud front, we need an external certificate that covers the domain name.
18. Memecached is simpler to use than Redis
19. Alias Records work with domains that are shorter (top node of DNS namespace), unlike A records. so if the record is "example.com", we can use alias record, but not A record.
20. Ephemeral ports are for the range 1024 - 65535.
21. <cloud>S3 Glacier</cloud> gives us 10Gb of data retrival for free each month
22. public IP address are maintained when the <cloud>EC2</cloud> instance goes into hibernation.
23. <cloud>DynamoDB DAX</cloud> can't integrate with <cloud>RDS</cloud>
24. the `fsck` command can repair file-system
25. Para-virtual AMI aren't supported in all regions, they allow for better boot performance.
26. <cloud>AWS Config</cloud> will not send a new notification if the state of a resource didn't change.
27. cloudfront does "sticky-ness" using cookies. and they need to be forwarded
28. <cloud>CloudTrail</cloud> log files are encrypted using SSE-S3 key.
---
1. <cloud>Aurora</cloud> point in time recovery vs backtracking?
2. is there something called Compute Optimizer?\
    *A: yes, it gives cost cuting recomendations*
5. <cloud>CloudFormation</cloud> Single start of authority?
6. <cloud>EFS</cloud> mounting for latency?
7. possible issues with NAT gateways?
8. using certificates <cloud>ACM</cloud>?
9. which IP to use when configuring a customer gateway? - *the nat device IP*
10. <cloud>WAF</cloud> scopes?
11. <cloud>RDS</cloud> monitoring?
    - Enhanced monitoring - cpu, processes, etc...
    - Performance Insight - DB specific metrics
12. Limits on <cloud>EC2</cloud> placement groups?
    - up to seven instance per Availability Zone in spread placement group
13. can't set dns private name
---
1. <cloud>CloudFront</cloud> caching behavior control headers
2. <cloud>Route53 resolver endpoints</cloud> - inbound, outbound,
3. <cloud>EBS</cloud> fast snapshot?
4. encrypting existing <cloud>RDS</cloud>?
5. <cloud>Network Load Balancer</cloud> when to use? HTTP, HTTPs?\
    *A: only TCP (HTTP, HTTPS), high load with low latency*
6. spot instances?
7. Global Accelerator - make applications faster, doesn't care about protocol.
8. <cloud>Session Manager</cloud> doesn't need the instances to open port 22 or bastion stuff, but they need an <cloud>IAM Role</cloud> that allows session manager to control them. also needs an agent
9. we can stop and start EC2 instances so they can be relaunched on a different host
---
1. is there something special about orcale database cross-region backups?
2. rds re-sizing? - possible
3. <cloud>S3 Glacier</cloud> vault lock policies? - wait until after the amount of time has passed, and then validate.
4. secrets manager minimal rotation period?
5. billing alerts must be enabled from the management account in <cloud>AWS Organizations</cloud>
6. AWS organization need to enable stuff like tag policies before using them.
---
1. compliance mode is more strict than governance mode.
2. can we change <cloud>SQS</cloud> queue type in place (from standard to fifo)?
3. <cloud>EventBridge</cloud> archive and replay? - yes.
4. internal and external ALB?
5. patching windows and EC2 machine using he <cloud>System Manager</cloud>?
6. Launch template are new, launch configurations are deprecated.
7. can a single auto scaling group combine multiple stuff? both on-demand and spot?
8. EBS metrics only work directly for instance store, if we have EBS volumes, we need to push the metrics manually
9. split tunnel option on VPN client - only use tunnel for some traffic.
10. it's always the same price (or cheaper) to move instance to newer generations. AWS wants to get rid of the old ones.
11. stronger instance can mean faster recovery and failover.
12. encrypted AMIs should be copied across accounts
---
1. how to check and validate all security groups?
2. aurora disaster recovery?
---
1. can windows machine use regular <cloud>EFS</cloud>? **NO!**
2. do we need a separate gateway endpoint for each <cloud>S3</cloud> bucket? **NO!**
3. public vs private virtual interface?
   -  private virtual interface for resources inside a <cloud>VPC</cloud>
   -  public virtual interface for resources outside of VPC
4. when are alarms stopped?
5. Lambda inside VPC
---
1. kms key rotation with imported key
2. rds key rotation
3. blocking all HTTP traffic compliance
4. AWS single sign on process - *SSO integrated with permission sets, Directory services...*
5. blocking services from being used with AWS Organizations
6. WAF header match
7. Auto scaling group does not scale - *<cloud>Trusted Advisor</cloud> for service limits,  CloudTrail for permission*
8. block malicious Ips? - *use WAF (firewall)*
9. S3 bucket access restriction
10. traffic routing difference for mobile versions
11. securely accessing S3 from EC2
12. block IP in VPC
13. ensuring no internet access
14. patching OS vunrabilites on EC2 machines
15. Improve Aurora performance
16. Changing CloudWatch Alarm state - *as long as the threshold is reached, alarm state.*
17. Encryption in hybrid environment - *VPN is a must,* 
18. multi-tier subnets with ALB?
19. Scaling on SQS
20. managing instances with system manager - *needs agent and IAM role, tags are for organizing, never to allow or disallow something*
21. improve S3 performance
22. EC2 machine moving host
23. Flow Logs format
24. single IP for whitelisting? - *NAT gateway, ENI is not something to share*
25. CloudFormation template failure
26. cost recomendations
27. analyze s3 access logs
28. more flow logs stuff
29. ensureing encryption on EBS volumes
30. cloudwatch alarms with ALB
31. adding account to AWS organization
32. cost allocation tags in aws organizations
33. tracking reserved instances costs in consolidated billing
34. ensure no SSH access in all EC2
35. single IP for end users
36. ACM and CloudFront
37. enforce serverside S3 objects
38. disable access keys
39. cloudFormation template not receiving signals
40. glacier vault
41. adding dashboard for custom metrics
42. autoscaling group scheduled
43. no access from home network
44. creating a new version of application with cloudFormation - *use `AutoScalingReplacingUpdate` policy when updating*
45. accessing EC2 instances
46. storage gateway rebooting
47. EC2 instance terminated immediately reasons - *EBS is encrypted and no permissions for KMS key*
48. increase cloudFront performance
49. running a failed cloudFormation stack - *relaunch the template*
50. High Availability for web applications
51. detect high uses for NAT gateway
52. Session Manager to access custom AMI
53. manage patches - *like before, agent installed and IAM role*
54. track memory metrics
55. Route53 across account - *`ALIAS` record pointing to Fully qualified name in the other account*
56. S3 as static website and 503 errors
57. connecting to S3 from inside AWS
58. RDS performance
59. S3 object lock governance mode bypass
60. connecting to S3 from private subnets - *NAT gateway for general internet access, Gateway Endpoint for S3 specific access*
61. elastic cache memecached increase size
62. shared Responsibility model
63. best practices
64. s3 as static website with 403 errors
65. passing data across regions without the public internet
---
1. securely review in other account
2. cloud trail file integration - *use CLI commands*
3. quicker EC2 warmup - *when possible, use warm pools, as this guarntess better start up time, unlike increasing instance sizes, which might, but not necessary*
4. shared responsibility model with NAT instance
5. better EC2 performance
6. Sharing RDS snapshots - *Share KMS key*
7. monitoring ALB
8. EC2 get AMIs for review - *use AWS Config*
9. on-premises storage
10. prevent S3 deletion by AWS Organization
11. most cost effective EC2
12. shared responsibility model with RDS
13. Stop Cross-Site Scripting with which service? - *WAF*
14. number of objects in S3
15. IAM policy to restrict access from IP
16. HPC with high resiliency
17. detect changes on cloudFormation stack
18. S3 versioning
19. generate cloudTrail report for one user
20. ALB and the originating IP for each request - *ALB access logs*
21. Improve RDS performance
22. cloudFormation parameters
23. DNS resolution from on-premises to AWS - *inbound resolver*
24. check which version of software is installed for vunrabilites - *Amazon Inspector*
25. hybrid cluster - *when using virtual private gateway, think route propagation*
26. Route53 and static S3 website -*Record and bucket must have the same name*
27. securing cloudtrail logs - *Enable cloud trail integrity loads, use MFA to protect against deletes*
28. redundancy in web application
29. restrict internet access but allow SSH - *VPN and virtual private gateway also go together*
30. can't ping EC2 machine
31. vertical scaling
32. ami and cloud formation
33. ELB failing health checks
34. maintenance events visibility - *Health API to get specific stuff related to you, RSS is only general data about AWS*
35. cloudwatch metrics timestamp
36. EC2 system status check failed
37. shared responsability model for cross account access
38. stack fails on another region - *can be ami that doesn't exists, or entire service (or permissions)*
39. Protect against various IPs - *WAF and rate-based rules*
40. High Availability - *means at least two instances in each AZ*
41. deploy stuff with minimal effort but custom stuff
42. ELB 5xx errors
43. replicate EC2 with special AMI across regions - *use the `ec2 copy-image` command*
44. restoring data from EBS snapshots
45. scaling lambda
46. rebooting EC2 with high CPU after 3 minutes
47. Memecached with high CPU - *usually avoid vertical scaling, as it requires changes to the application but it can be done. we can have up to 40 nodes*
48. route53 health check records
49. increase s3 upload performance
50. Storage Gateway volume cache performance - *we can add another disk without stopping*
51. best cost plans (reserved instances, dedicated host, etc...)
52. EC2 state change alert (state manager or event bridge)
53. memcached performance - *add nodes*
54. local storage and cloud backup - *Storage gatways (not cache gateway)*
55. EBS volume for big data
56. ensure HTTPS
57. cloudFormation Rollback failed - *use <kbd>ContinueUpdateRollback</kbd> from console or `continue-update-rollback` api call*
58. route53 health check failed when looking at response body - *we can use the first 5120 bytes of the response*
59. NAT instance vs NAT gateway
60. Amazon Inspector possible issues?
61. ensure S3 logging on all buckets
62. CloudFormation StackSets with `OUTDATED` status - *tried to create a global resource which isn't unique, permissions or trust relation missing, service quota reached*
63. S3 bucket Policy
64. NAT instance troubleshoot - *source/destination checks*
65. IPv6 internet access - *use egress-only internet gateway*
---
1. IAM and Active Directory
2. EBS encryption, KMS key rotation
3. <cloud>EC2</cloud> capacity
4. Spot instances
5. cloud formation stack creation failed
6. where to check for layer 7 errors?
7. finding cost increase
8. restrict resources in aws organization
9. cloudfront location restrictions
10. elastic beanstalk environment creation failed
11. automate lambda invocation
12. High Availability RDS
13. s3 access without public internet
14. aws organization, AWS health for each account
15. EC2 autoscaling schedule
16. reusing stacks
17. restrict S3 access
18. imported portfolio actions
19. web application performance
20. networking on a new VCP
21. instance affected by maintenance
22. restrict services for accounts
23. too many connections on RDS
24. block traffic to specific IP from instance
25. reduce failover time
26. allow some account to work despite policies to block resources
27. route53 policies
28. shared responsability model - NAT gateway
29. get better performance with website
30. monitor componenets for metrics
31. auto scaling group capture instance before termination
32. finding networking issues
33. patching resource
34. migrating to RDS
35. reduce S3 load
36. maintenance windows overcoming
37. EC2 access SSM permissions
38. Routing to cloudFront
39. copy RDS to another account
40. retain logs before EC2 terminates
41. maintain consistent file access across many EC2 machines
42. EBS metrics monitoring
43. saving plans when using Fargate
44. manually set alarm state
45. find out who puts stuff in SQS
46. KMS key rotation
47. forward logs from EC2
48. networking security
49. prevent accidental EC2 termination
50. monitor S3 data events
51. aurora "inaccessible-encryption-credentials" error
52. Shared Responsability model on IAM
53. monitor EFS free space
54. find which port is accessed
55. Redis Cache
56. SCP policy to deny something
57. costs breakdown with tags
58. sticky sessions
59. possible RDS failure that lead to failover
60. move RDS to encrypted
61. redshit access
62. portfolio sharing
63. scaling on cpu
64. deploy cloudFormation to multiple accounts
65. sticky sessions
---
1. getting low latency
2. grouping lgs
3. flow logs format
4. cost saving - *use compute saving plans when lambda and fargate are mentioned*
5. migrating across AZ zones
6. `DiskReadyBytes` metric
7. monitoring service quotas
8. kibana - *think <cloud>ElasticSearch</cloud>*
9. dynamoDB replication - *global tables*
10. storing credentials
11. routing tales
12. internet access
13. capturing EBS metrics - *`VolumeReadBytes` and `VolumeWriteBytes`*
14. only private access to elasticSearch thorugh VPN - *we can't change endpoint from public to private or back, use IP cidr*
15. share KMS customer managed key - *do key stuff*
16. stackSet testing? - *create a new stack, stackSets apply for all accounts*
17. matching Availability Zones across account - *use zoneId, zoneName is fictional*
18. MFA when using CLI
19. only allow HTTPs on ALB - *private certificates are only trusted inside the AWS account, for outside, we need public SSL*
20. allow internet access from VPC
21. better EFS performance - *Burting throughput scales based on size of the file system, if we want the performance to always be ok, we need provisioned throughput*
22. restrict access to cached content from S3
23. certificate renews - *dns validation is automatic, email validation requires manual work*
24. cost allocation
25. dynamoDB replication
26. SAML web console access - *AWS federated use SSO saml*
27. consolidated billing
28. ensure all object uploaded to S3 are encrypted - *enable default encryption and use bucket policies, remember to check if in rest or in-transit encryption, not ACL*
29. share S3 bucket across organization policy
30. latency routing
31. ensure cloudTrail is enable
32. geolocation routing
33. IAM policy
34. protect from DDoS using WAF and rate limiting
35. track cost by user - *there is a `CreatedBy` tag*
36. Session Manager
37. access S3 bucket from private subnet
38. increasing performance for web application - *Global Accelerator service*
39. analyze layer7 results
40. RDS disaster recovery - *if geographic region is mentioned, cross-region replication*
41. Single Ip
42. CloudFront sticky session - *elb sticky sessions + cookie forwarding*
43. content restrictions
44. EC2 performance but not CPU - *if this is going on for a while, we are out of burst credits, and we need to allow unlimited bursting*
45. access denied S3 CloudFront - *no such thing as default S3 bucket policy that denies public access*
46. S3 object integrity - *CloudTrail integrity*
47. reboot based on metric and 2 minutes period
48. ELB routing to correct instance
49. health checks
50. scale EKS - *Cluster AutoScaler*
51. S3 private access limits
52. aws organizations features - *for tagging policies, we must have all features on! includeing consolidated billing, but that's not enough on it's own*
53. **InstanceLimitExceeded** error - *this is per region! service quota increase!*
54. improve cloudFront hit ratio
55. stop unused EC2 instances
56. increase aurora performance
57. **InstanceLimitExceeded** error - *request quota increase!*
58. ensure S3 are never deleted aws organizations
59. analyze access to ELB
60. faster S3 upload
61. RDS access fails
62. Cross site scripting attack
63. SSH fails to public EC2
64. read only IAM
65. connect to S3 buckets through gateway endpoint - *one endpoint is enough for all buckets*
