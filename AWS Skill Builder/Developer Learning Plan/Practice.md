<!--
ignore these words in spell check for this file
// cSpell:ignore elbv2 Neumann cgroups pictShare Kubelet eksctl Karpenter kube-proxy kubeconfig kube-system Alexa omponent
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Developer Exam Prep

## Exam 1

<details>
<summary>
not usefull at all
</summary>

[Exam Prep Official Practice Question Set: AWS Certified Developer - Associate (DVA-C02 - English)](https://explore.skillbuilder.aws/learn/course/13757/play/78188/overview-and-instructions-aws-certification-official-practice-question-sets)

### Q01

> A company needs to encrypt an existing Amazon RDS DB instance that is unencrypted. Downtime is acceptable for the project.
>
> Select and order the correct steps from the following list to provide the required encryption. (Select THREE.)
>
> - Enable encryption on the existing DB instance.
> - Create an encrypted read replica of the existing DB instance.
> - Create a snapshot of the existing DB instance.
> - Create an encrypted copy of the snapshot.
> - Restore the DB instance from the encrypted read replica.
> - Restore the DB instance from the encrypted snapshot

answer:

- Create a snapshot of the existing DB instance.
- Create an encrypted copy of the snapshot.
- Restore the DB instance from the encrypted snapshot.

### Q02

> Select and order the AWS infrastructure layers in the following list from LARGEST scope to SMALLEST scope. (Select THREE.)
>
> - Availability Zone
> - AWS Region
> - Subnet

Answer:

- AWS Region
- Availability Zone
- Subnet

### Q03

> A company needs to encrypt an existing Amazon RDS DB instance that is unencrypted. Downtime is acceptable for the project.\
> Which sequence of steps will provide the required encryption?
>
> 1. Create a snapshot of the existing DB instance. Create an encrypted copy of the snapshot. Restore the DB instance from the encrypted snapshot.
> 1. Enable encryption on the existing DB instance. Create an encrypted read replica of the existing DB instance. Restore the DB instance from the encrypted read replica.
> 1. Enable encryption on the existing DB instance. Create a snapshot of the existing DB instance. Restore the DB instance from the snapshot.
> 1. Enable automated backups on the existing DB instance. Enable encryption on the automated backup. Restore the DB instance from the encrypted backup.

answer:

Create a snapshot of the existing DB instance. Create an encrypted copy of the snapshot. Restore the DB instance from the encrypted snapshot.

### Q04

> A company needs to minimize its Amazon S3 storage costs.\
> Select the correct S3 storage class from the following list to meet this requirement for each use case.
>
> - S3 Glacier Flexible Retrieval
> - S3 Standard
>
> 1. Application logs that are retrieved daily for analysis | S3 Standard
> 1. Backups that are restored every 3 months and have a required retrieval time of up to 12 hours | _S3 Glacier Flexible Retrieval_
> 1. Data that is retrieved one time each year | _S3 Glacier Flexible Retrieval_
> 1. Video files that must be retrieved within milliseconds | _S3 Standard_

### Q05

> Select the correct category of workload from the following list for each AWS service.
>
> - Analytics
> - Machine learning
>
> for each select one
>
> - Amazon EMR | *Analytics*
> - Amazon Forecast | *Analytics*
> - Amazon Kendra | *Machine learning*
> - Amazon QuickSight | *Machine learning*
> - AWS Glue | *Analytics*
> - AWS Panorama | *Machine learning*

### Q06

> A company is migrating an e-commerce application to AWS. The application consists of web servers, application servers, relational databases, storage, and a cache. The company needs to design an architecture that provides resilience against failures.
>
> Which combination of actions will achieve fault tolerance for the web servers and application servers? (Select TWO.)
>
> 1. Configure Auto Scaling groups of Amazon EC2 instances across multiple Availability Zones.
> 1. Deploy Amazon EC2 instances in multiple subnets in one Availability Zone.
> 1. Implement load balancing for the Amazon EC2 instances.
> 1. Launch Amazon EC2 Spot Instances.
> 1. Launch large Amazon EC2 instances.

Answers:

- Configure Auto Scaling groups of Amazon EC2 instances across multiple Availability Zones.
- Implement load balancing for the Amazon EC2 instances.

### Q07

> A company is migrating an e-commerce application to AWS. The application consists of web servers, application servers, relational databases, storage, and a cache. The company needs to design an architecture that provides resilience against failures.
>
> Which action will ensure high availability for the databases?
>
> 1. Configure Amazon DynamoDB Accelerator (DAX).
> 1. Configure an Amazon RDS Multi-AZ DB instance deployment.
> 1. Deploy Amazon RDS DB instances in private subnets.
> 1. Deploy read replicas in multiple AWS Regions.

Answer:\
Configure an Amazon RDS Multi-AZ DB instance deployment.
