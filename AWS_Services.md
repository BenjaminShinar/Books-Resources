<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href=".markdown-style.css">

# AWS Services

## Compute

- EC2 - Elastic Cloud Compute: Computing Instances
- EKS - Elastic Kubernetes Service - Managed Kubernetes Cluster
- ECS - Elastic Container Service - AWS propriety Containers
- Lambda - Serverless Computer

## Storage

- S3 (Simple Storage Service): Object Storage
- Snowball: physical data transfer
- FSx (windows and Lustre) - native windows managed file system. Lustre is for compute intensive data.
- EBS
- Instance Storage

## DataBases

- Aurora - highly available MySQL and PostgresSQL compatible Relational Database.
- DocumentDB - another NoSQL solution.
- DynamoDB - AWS serverless NoSQL solution.
- DataSync: Synchronize data between AWS and on-premises.
- DMS - Database Migration
- ElasticCache - in memory "database". redis, memcached

## Queries

- Athena - Serverless query service to analyze data stored in Amazon S3 with SQL syntax. can also work on other data sources with Data Source Connectors
- PartiQL - Query DynamoDB with SQL-like language
- RedShift - AWS data warehouse solution (BI).
- AppSync - Fully-Managed Serverless GraphQL API Service for Real-Time Data Queries
- OpenSearch - additional database for better searching capabilities replaces ElasticSearch

## Orchestration

- SES (Simple Email Service) - send emails
- SNS (Simple Notification Service): send emails
- SQS (Simple Queue Service)
- Step Functions
- SWF (Simple Work Flow service)
- - EventBridge - Central Events in AWS
- Kinesis - Streaming Data platform (continuously generated data from multiple source, not video streaming)
- - MSK - Managed Streaming for Apache Kafka

## CI-CD

- CodeCommit - Store code, a git-compatible repository
- CodePipeline - automate building to Elastic Beanstalk
- CodeBuild - building and testing
- CodeDeploy - deploy to EC2 instances
- CodeStar - manage at ci-cd services in single location
- CodeArtifact - store, publish and share software artifacts
- CodeGuru - code review by machine learning
- Cloud9 - in-browser IDE

## Deployment and infrastructure

- Elastic Beanstalk - one click solution to provision resource, like CloudFormation for dummies. easily deploy EC2 based applications (web-server, producer consumer) with load balancer.
- CloudFormation - templates for provisioning aws resources.
- CoPilot - provisions and manages ECS services.
- EC2 Image Builder - automate the creation of AMI images or Docker Images

## Management

- IAM (Identity Access Management)
- CloudWatch - Monitors Resource usage and other metrics, used for billing alarms. Alarms,Events, Logs and Dashboards.
- CloudWatch Evidently - Safely validate new features by serving them to a specified % of your users. feature flags and A/B testing.
- CloutTrail: Monitors AWS actions (from the console or the API), which users and accounts did what
- Cognito - Web Identity Federation service
- CognitoSync - allow marinating an offline copy of data on the user device.
- Managed Microsoft AD - Active Directory by AWS
- RAM - Resources Access Manger - share resources across account
- SSM - Parameter Store and Secrets Manager
- Service Quotes - see the AWS quotas at one place and request to increase them.
- License Manager - manage software licenses

## Networking and Security

- Internet Gateway - in the VPC, uses route table
- Elastic Load Balancer - classic, application, network
- Route53 - DNS
- API Gateway - publish and maintain APIs
- ACM (Amazon Certificate Manager) - Used to provide in-flight encryption for websites (HTTPS) SSL/TLS Certificates.
- KMS - Key Management Service
- Amazon Shield - DDos Protection

## Machine Learning

- Quicksight - data-driven unified business intelligence (BI) at hyper-scale.
- Macie - data security and data privacy service that uses machine learning and pattern matching to discover and protect your sensitive data.
- Bedrock - Foundation model as a service
- SageMaker - training machine learning models from scratch
- SageMaker JumpStart - re-training models with new custom data
- HealthScribe - Generative AI tool for transcribing medical conversations.
- CodeWhisperer - Coding Tool
- Lex - chatbot
- Kendra - AI based search
- Amazon Q Business/Developer - ai assistant
- Textract - extract text from scanned documents and handwritten text

## Other

- API Amplify - Full-Stack Web and Mobile Apps
- AppConfig - Configure, validate, and deploy dynamic configurations
- Elastic Transcoder - transcode (convert) media files between different formats.
- Glue - Transform Data
- AWSConfig - Monitor and prevent resources according to policy
