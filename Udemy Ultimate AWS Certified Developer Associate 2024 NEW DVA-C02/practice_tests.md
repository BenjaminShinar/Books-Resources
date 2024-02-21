<!--
ignore these words in spell check for this file
// cSpell:ignore NACL Alexa RCU xLarge
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Developer Exam Preperations

Getting ready for the exam.

Udemy [Course Practice Exams | AWS Certified Developer Associate 2023](https://checkpoint.udemy.com/course/aws-certified-developer-associate-practice-tests-dva-c01)



## Course Practice Exam

<details>
<summary>
Practice Exam from the Course
</summary>

### Question 01

Question

> You are designing a high-performance application that requires *millions of connections*. You have several EC2 instances running Apache2 web servers and the application will require capturing the user’s source IP address and source port without the use of X-Forwarded-For.\
> Which of the following options will meet your needs?

Answer

> Network Load Balancer

 A Network Load Balancer functions at the fourth layer of the Open Systems Interconnection (OSI) model. It can handle millions of requests per second. After the load balancer receives a connection request, it selects a target from the target group for the default rule. It attempts to open a TCP connection to the selected target on the port specified in the listener configuration. Incoming connections remain unmodified, so application software need not support X-Forwarded-For.
### Question 02

Question

> A security company is requiring all developers to perform server-side encryption with customer-provided encryption keys when performing operations in AWS S3. Developers should write software with C# using the AWS SDK and implement the requirement in the PUT, GET, Head, and Copy operations.
> Which of the following encryption methods meets this requirement?

Answer

> SSE-C

 You have the following options for protecting data at rest in Amazon S3:
> - Server-Side Encryption – Request Amazon S3 to encrypt your object before saving it on disks in its data centers and then decrypt it when you download the objects.
> - Client-Side Encryption – Encrypt data client-side and upload the encrypted data to Amazon S3. In this case, you manage the encryption process, the encryption keys, and related tools.
> 
> For the given use-case, the company wants to manage the encryption keys via its custom application and let S3 manage the encryption, therefore you must use Server-Side Encryption with Customer-Provided Keys (SSE-C).\
> Using server-side encryption with customer-provided encryption keys (SSE-C) allows you to set your encryption keys. With the encryption key you provide as part of your request, Amazon S3 manages both the encryption, as it writes to disks, and decryption, when you access your objects.


### Question 03

Question

> An IT company uses AWS CloudFormation templates to provision their AWS infrastructure for Amazon EC2, Amazon VPC, and Amazon S3 resources. Using cross-stack referencing, a developer creates a stack called NetworkStack which will export the subnetId that can be used when creating EC2 instances in another stack.\
> To use the exported value in another stack, which of the following functions must be used?

Answer

> !ImportValue



 > The intrinsic function Fn::ImportValue returns the value of an output exported by another stack. You typically use this function to create cross-stack references.
> Incorrect options:
> - !Ref - Returns the value of the specified parameter or resource.
> - !GetAtt - Returns the value of an attribute from a resource in the template.
> - !Sub - Substitutes variables in an input string with values that you specify.
### Question 04

Question

> A development team at a social media company uses AWS Lambda for its serverless stack on AWS Cloud. For a new deployment, the Team Lead wants to send only a certain portion of the traffic to the new Lambda version. In case the deployment goes wrong, the solution should also support the ability to roll back to a previous version of the Lambda function, with MINIMUM downtime for the application.\
> As a Developer Associate, which of the following options would you recommend to address this use-case?

Answer

> Set up the application to use an alias that points to the current version. Deploy the new version of the code and configure the alias to send 10% of the users to this new version. If the deployment goes wrong, reset the alias to point all traffic to the current version


### Question 05

Question

> A developer from your team has configured the load balancer to route traffic equally between instances or across Availability Zones. However, Elastic Load Balancing (ELB) routes more traffic to one instance or Availability Zone than the others.\
> Why is this happening and how can it be fixed? (Select two)

Answer

> - Sticky sessions are enabled for the load balancer
> - Instances of a specific capacity type aren’t equally distributed across Availability Zones


### Question 06

Question

> A developer is configuring a bucket policy that denies upload object permission to any requests that do not include the `x-amz-server-side-encryption` header requesting server-side encryption with SSE-KMS for an Amazon S3 bucket - examplebucket.\
> Which of the following policies is the right fit for the given requirement?

Answer

``` json
{ "Version":"2012-10-17", "Id":"PutObjectPolicy", "Statement":[{ "Sid":"DenyUnEncryptedObjectUploads", "Effect":"Deny", "Principal":"", "Action":"s3:PutObject", "Resource":"arn:aws:s3:::examplebucket/", "Condition":{ "StringNotEquals":{ "s3:x-amz-server-side-encryption":"aws:kms" } } } ] }
```


### Question 07

Question

> You are responsible for an application that runs on multiple Amazon EC2 instances. In front of the instances is an Internet-facing load balancer that takes requests from clients over the internet and distributes them to the EC2 instances.
> A health check is configured to ping the index.html page found in the root directory for the health status. When accessing the website via the internet visitors of the website receive timeout errors.\
> What should be checked first to resolve the issue?

Answer

> Security Groups


### Question 08

Question

> A development team wants to build an application using serverless architecture. The team plans to use AWS Lambda functions extensively to achieve this goal.\
> The developers of the team work on different programming languages like Python, .NET and Javascript. The team wants to model the cloud infrastructure using any of these programming languages.\
> Which AWS service/tool should the team use for the given use-case?

Answer

> AWS Cloud Development kit


### Question 09

Question

> A .NET developer team works with many ASP.NET web applications that use EC2 instances to host them on IIS. The deployment process needs to be configured so that multiple versions of the application can run in AWS Elastic Beanstalk. One version would be used for development, testing, and another version for load testing.\
> Which of the following methods do you recommend?

Answer

> Define a dev environment with a single instance and a 'load test' environment that has settings close to production environment.

 AWS Elastic Beanstalk makes it easy to create new environments for your application. You can create and manage separate environments for development, testing, and production use, and you can deploy any version of your application to any environment. Environments can be long-running or temporary. When you terminate an environment, you can save its configuration to recreate it later.\
> It is common practice to have many environments for the same application. You can deploy multiple environments when you need to run multiple versions of an application. So for the given use-case, you can set up 'dev' and 'load test' environment.
### Question 10

Question

> A company has a workload that requires 14,000 consistent IOPS for data that must be durable and secure. The compliance standards of the company state that the data should be secure at every stage of its lifecycle on all of the EBS volumes they use
> Which of the following statements are true regarding data security on EBS?

Answer

> EBS volumes support both in-flight encryption and encryption at rest using KMS


### Question 11

Question

> A firm runs its technology operations on a fleet of Amazon EC2 instances. The firm needs a certain software to be available on the instances to support their daily workflows. The developer team has been told to use the user data feature of EC2 instances.
> Which of the following are true about the user data EC2 configuration? ( Select two)

Answer

> - By default, scripts entered as user data are executed with root user privileges.
> - By default, user data runs only during the boot cycle when you first launch an instance


### Question 12

Question

> You have a Java-based application running on EC2 instances loaded with AWS CodeDeploy agents. You are considering different options for deployment, one is the flexibility that allows for incremental deployment of your new application versions and replaces existing versions in the EC2 instances. The other option is a strategy in which an Auto Scaling group is used to perform a deployment.\
> Which of the following options will allow you to deploy in this manner? (Select two)

Answer

> - In-place Deployment
> - Blue/green Deployment




### Question 13

Question

> Your company leverages Amazon CloudFront to provide content via the internet to customers with low latency. Aside from latency, security is another concern and you are looking for help in enforcing end-to-end connections using HTTPS so that content is protected.\
> Which of the following options is available for HTTPS in AWS CloudFront?


Answer

> Between clients and CloudFront as well as between CloudFront and backend



### Question 14

Question

> You are planning to build a fleet of EBS-optimized EC2 instances to handle the load of your new application. Due to security compliance, your organization wants any secret strings used in the application to be encrypted to prevent exposing values as clear text.
> The solution requires that decryption events be audited and API calls to be simple. How can this be achieved? (select two)

Answer

> - Store the secret as SecureString in SSM Parameter Store.
> - Audit using CloudTrail


### Question 15

Question

> You are storing your video files in a separate S3 bucket than your main static website in an S3 bucket. When accessing the video URLs directly the users can view the videos on the browser, but they can't play the videos while visiting the main website.\
> What is the root cause of this problem?

Answer

> Enable CORS


### Question 16

Question

> A development team has configured inbound traffic for the relevant ports in both the Security Group of the EC2 instance as well as the Network Access Control List (NACL) of the subnet for the EC2 instance. The team is, however, unable to connect to the service running on the Amazon EC2 instance.\
> As a developer associate, which of the following will you recommend to fix this issue?

Answer

> Security Groups are stateful, so allowing inbound traffic to the necessary ports enables the connection. Network ACLs are stateless, so you must allow both inbound and outbound traffic

### Question 17

Question

> A development team has deployed a REST API in Amazon API Gateway to two different stages - a test stage and a prod stage. The test stage is used as a test build and the prod stage as a stable build. After the updates have passed the test, the team wishes to promote the test stage to the prod stage.\
> Which of the following represents the optimal solution for this use-case?

Answer

> Update stage variable value from the stage name of test to that of prod





### Question 18

Question

> As an AWS certified developer associate, you are working on an AWS CloudFormation template that will create resources for a company's cloud infrastructure. Your template is composed of three stacks which are Stack-A, Stack-B, and Stack-C. Stack-A will provision a VPC, a security group, and subnets for public web applications that will be referenced in Stack-B and Stack-C.\
> After running the stacks you decide to delete them, in which order should you do it?

Answer

> Stack B, then Stack C, then Stack A


### Question 19

Question

> The Development team at a media company is working on securing their databases.\
> Which of the following AWS database engines can be configured with IAM Database Authentication? (Select two)

Answer

> - RDS MySQL
> - RDS PostGreSQL



### Question 20

Question

> A large firm stores its static data assets on Amazon S3 buckets. Each service line of the firm has its own AWS account. For a business use case, the Finance department needs to give access to their S3 bucket's data to the Human Resources department.
> Which of the below options is NOT feasible for cross-account access of S3 bucket objects?

Answer

> Use IAM roles and resource-based policies delegate access across accounts within different partitions via programmatic access only.

AWS Delegation are between partitions (china, gov, general aws).



### Question 21

Question

> When running a Rolling deployment in Elastic Beanstalk environment, only two batches completed the deployment successfully, while rest of the batches failed to deploy the updated version. Following this, the development team terminated the instances from the failed deployment.\
> What will be the status of these failed instances post termination?

Answer

> Elastic Beanstalk will replace them with instances running the application version from the most recent successful deployment



### Question 22

Question

> While defining a business workflow as state machine on AWS Step Functions, a developer has configured several states.\
> Which of the following would you identify as the state that represents a single unit of work performed by a state machine?

Answer

``` json
"HelloWorld": {
  "Type": "Task",
  "Resource": "arn:aws:lambda:us-east-1:123456789012:function:HelloFunction",
  "Next": "AfterHelloWorldState",
  "Comment": "Run the HelloWorld Lambda function"
}
```



### Question 23

Question

> A Developer at a company is working on a CloudFormation template to set up resources. Resources will be defined using code and provisioned based on certain conditions defined in the `Conditions` section.\
> Which section of a CloudFormation template cannot be associated with `Condition`?

Answer

> Parameters



### Question 24

Question

> An e-commerce company has multiple EC2 instances operating in a private subnet which is part of a custom VPC. These instances are running an image processing application that needs to access images stored on S3. Once each image is processed, the status of the corresponding record needs to be marked as completed in a DynamoDB table.\
> How would you go about providing private access to these AWS resources which are not part of this custom VPC?

Answer

> Create a separate gateway endpoint for S3 and DynamoDB each. Add two new target entries for these two gateway endpoints in the route table of the custom VPC



### Question 25

Question

> You team maintains a public API Gateway that is accessed by clients from another domain. Usage has been consistent for the last few months but recently it has more than doubled. As a result, your costs have gone up and would like to prevent other unauthorized domains from accessing your API.\
> Which of the following actions should you take?

Answer

> Restrict access by using CORS 



### Question 26

Question

> A company's e-commerce application becomes slow when traffic spikes. The application has a three-tier architecture (web, application and database tier) that uses synchronous transactions. The development team at the company has identified certain bottlenecks in the application tier and it is looking for a long term solution to improve the application's performance.\
> As a developer associate, which of the following solutions would you suggest to meet the required application response times while accounting for any traffic spikes?

Answer

> Leverage horizontal scaling for the web and application tiers by using Auto Scaling groups and Application Load Balancer

### Question 27

Question

> You are working for a shipping company that is automating the creation of ECS clusters with an Auto Scaling Group using an AWS CloudFormation template that accepts cluster name as its parameters. Initially, you launch the template with input value 'MainCluster', which deployed five instances across two availability zones. The second time, you launch the template with an input value 'SecondCluster'.\
> However, the instances created in the second run were also launched in 'MainCluster' even after specifying a different cluster name.\
> What is the root cause of this issue?

Answer

>  The cluster name Parameter has not been updated in the file /etc/ecs/ecs.config during bootstrap



### Question 28

Question

> A financial services company has developed a REST API which is deployed in an Auto Scaling Group behind an Application Load Balancer. The API stores the data payload in DynamoDB and the static content is served through S3. Upon analyzing the usage pattern, it's found that 80% of the read requests are shared across all users.\
> As a Developer Associate, how can you improve the application performance while optimizing the cost with the least development effort?

Answer

> Enable DynamoDB Accelerator (DAX) for DynamoDB and CloudFront for S3





### Question 29

Question

> AWS CloudFormation helps model and provision all the cloud infrastructure resources needed for your business.\
> Which of the following services rely on CloudFormation to provision resources (Select two)?

Answer

> - AWS Elastic Beanstalk
> - AWS Serverless Application Model (AWS SAM)



### Question 30

Question

> Your company has been hired to build a resilient mobile voting app for an upcoming music award show that expects to have 5 to 20 million viewers. The mobile voting app will be marketed heavily months in advance so you are expected to handle millions of messages in the system. You are configuring Amazon Simple Queue Service (SQS) queues for your architecture that should receive messages from 20 KB to 200 KB.\
> Is it possible to send these messages to SQS?

Answer

> Yes, the max message size is 256KB



### Question 31

Question

> As a developer, you are working on creating an application using AWS Cloud Development Kit (CDK).\
> Which of the following represents the correct order of steps to be followed for creating an app using AWS CDK?

Answer

> Create the app from a template provided by AWS CDK -> Add code to the app to create resources within stacks -> Build the app (optional) -> Synthesize one or more stacks in the app -> Deploy stack(s) to your AWS account



### Question 32

Question

> An Auto Scaling group has a maximum capacity of 3, a current capacity of 2, and a scaling policy that adds 3 instances.\
> When executing this scaling policy, what is the expected outcome?


Answer

> Amazon EC2 Auto Scaling adds only 1 instance to the group



### Question 33

Question

> Amazon Simple Queue Service (SQS) has a set of APIs for various actions supported by the service.\
> As a developer associate, which of the following would you identify as correct regarding the CreateQueue API? (Select two)

Answer

> - You can't change the queue type after you create it
> - The visibility timeout value for the queue is in seconds, which defaults to 30 seconds

### Question 34

Question

> ECS Fargate container tasks are usually spread across Availability Zones (AZs) and the underlying workloads need persistent cross-AZ shared access to the data volumes configured for the container tasks.\
> Which of the following solutions is the best choice for these workloads?

Answer

> Amazon EFS volumes



### Question 35

Question

> A Developer is configuring Amazon EC2 Auto Scaling group to scale dynamically.\
> Which metric below is NOT part of Target Tracking Scaling Policy?



Answer

> `ApproximateNumberOfMessagesVisible` 



### Question 36

Question

> You are working with a t2.small instance bastion host that has the AWS CLI installed to help manage all the AWS services installed on it. You would like to know the security group and the instance id of the current instance.
> Which of the following will help you fetch the needed information?

Answer

> Query the metadata at http://169.254.169.254/latest/meta-data



### Question 37

Question

> A company follows collaborative development practices. The engineering manager wants to isolate the development effort by setting up simulations of API components owned by various development teams.\
> Which API integration type is best suited for this requirement?


Answer

> Mock



### Question 38

Question

> You were assigned to a project that requires the use of the AWS CLI to build a project with AWS CodeBuild. Your project's root directory includes the buildspec.yml file to run build commands and would like your build artifacts to be automatically encrypted at the end.\
> How should you configure CodeBuild to accomplish this?


Answer

> Specify a KMS key to use





### Question 39

Question

> A company uses Amazon Simple Email Service (SES) to cost-effectively send subscription emails to the customers. Intermittently, the SES service throws the error: Throttling – Maximum sending rate exceeded.\
> As a developer associate, which of the following would you recommend to fix this issue?

Answer

> Use Exponential Backoff technique to introduce delay in time before attempting to execute the operation again



### Question 40

Question

> Your company has embraced cloud-native microservices architectures. New applications must be dockerized and stored in a registry service offered by AWS. The architecture should support dynamic port mapping and support multiple tasks from a single service on the same container instance. All services should run on the same EC2 instance.\
> Which of the following options offers the best-fit solution for the given use-case?

Answer

> Application Load Balancer + ECS



### Question 41

Question

> A company ingests real-time data into its on-premises data center and subsequently a daily data feed is compressed into a single file and uploaded on Amazon S3 for backup. The typical compressed file size is around 2 GB.\
> Which of the following is the fastest way to upload the daily compressed file into S3?

Answer

> Upload the compressed file using multipart upload with S3 transfer acceleration





### Question 42

Question

> A developer is defining the signers that can create signed URLs for their Amazon CloudFront distributions.\
> Which of the following statements should the developer consider while defining the signers? (Select two)

Answer

> - When you create a signer, the public key is with CloudFront and private key is used to sign a portion of URL
> - When you use the root user to manage CloudFront key pairs, you can only have up to two active CloudFront key pairs per AWS account



### Question 43

Question

> A company has recently launched a new gaming application that the users are adopting rapidly. The company uses RDS MySQL as the database. The development team wants an urgent solution to this issue where the rapidly increasing workload might exceed the available database storage.\
> As a developer associate, which of the following solutions would you recommend so that it requires minimum development effort to address this requirement?

Answer

> Enable storage auto-scaling for RDS MySQL


### Question 44

Question

> You are getting ready for an event to show off your Alexa skill written in JavaScript. As you are testing your voice activation commands you find that some intents are not invoking as they should and you are struggling to figure out what is happening. You included the following code console.log(JSON.stringify(this.event)) in hopes of getting more details about the request to your Alexa skill.\
> You would like the logs stored in an Amazon Simple Storage Service (S3) bucket named MyAlexaLog. How do you achieve this?

Answer

> Use CloudWatch integration feature with S3





### Question 45

Question

> A developer is testing Amazon Simple Queue Service (SQS) queues in a development environment. The queue along with all its contents has to be deleted after testing.\
> Which SQS API should be used for this requirement?

Answer

> DeleteQueue 



### Question 46

Question

> A popular mobile app retrieves data from an AWS DynamoDB table that was provisioned with read-capacity units (RCU’s) that are evenly shared across four partitions. One of those partitions is receiving more traffic than the other partitions, causing hot partition issues.\
> What technology will allow you to reduce the read traffic on your AWS DynamoDB table with minimal effort?

Answer

> DynamoDB DAX





### Question 47

Question

> A banking application needs to send real-time alerts and notifications based on any updates from the backend services. The company wants to avoid implementing complex polling mechanisms for these notifications.\
> Which of the following types of APIs supported by the Amazon API Gateway is the right fit?

Answer

> WebSocket APIs





### Question 48

Question

> A company would like to migrate the existing application code from a GitHub repository to AWS CodeCommit.\
> As an AWS Certified Developer Associate, which of the following would you recommend for migrating the cloned repository to CodeCommit over HTTPS?

Answer

> Use Git credentials generated from IAM



### Question 49

Question

> A pharmaceutical company uses Amazon EC2 instances for application hosting and Amazon CloudFront for content delivery. A new research paper with critical findings has to be shared with a research team that is spread across the world.\
> Which of the following represents the most optimal solution to address this requirement without compromising the security of the content?

Answer

> Use CloudFront signed URL feature to control access to the file





### Question 50

Question

> An e-commerce company uses AWS CloudFormation to implement Infrastructure as Code for the entire organization. Maintaining resources as stacks with CloudFormation has greatly reduced the management effort needed to manage and maintain the resources. However, a few teams have been complaining of failing stack updates owing to out-of-band fixes running on the stack resources.\
> Which of the following is the best solution that can help in keeping the CloudFormation stack and its resources in sync with each other?

Answer

> Use Drift Detection feature of CloudFormation





### Question 51

Question

> A company wants to automate its order fulfillment and inventory tracking workflow. Starting from order creation to updating inventory to shipment, the entire process has to be tracked, managed and updated automatically.\
> Which of the following would you recommend as the most optimal solution for this requirement?

Answer

> Use AWS Step Functions to coordinate and manage the components of order management and inventory tracking workflow



### Question 52

Question

> DevOps engineers are developing an order processing system where notifications are sent to a department whenever an order is placed for a product. The system also pushes identical notifications of the new order to a processing module that would allow EC2 instances to handle the fulfillment of the order.\
> In the case of processing errors, the messages should be allowed to be re-processed at a later stage.The order processing system should be able to scale transparently without the need for any manual or programmatic provisioning of resources.\
> Which of the following solutions can be used to address this use-case?

Answer

> SNS + SQS





### Question 53

Question

> A company has built its technology stack on AWS serverless architecture for managing all its business functions. To expedite development for a new business requirement, the company is looking at using pre-built serverless applications.\
> Which AWS service represents the easiest solution to address this use-case?

Answer

> AWS Serverless Application Repository (SAR)



### Question 54

Question

> What steps can a developer take to optimize the performance of a CPU-bound AWS Lambda function and ensure fast response time?

Answer

> Increase the function's memory





### Question 55

Question

> You work as a developer doing contract work for the government on AWS gov cloud. Your applications use Amazon Simple Queue Service (SQS) for its message queue service. Due to recent hacking attempts, security measures have become stricter and require you to store data in encrypted queues.\
> Which of the following steps can you take to meet your requirements without making changes to the existing code?

Answer

> Enable SQS KMS encryption





### Question 56

Question

> You have migrated an on-premise SQL Server database to an Amazon Relational Database Service (RDS) database attached to a VPC inside a private subnet. Also, the related Java application, hosted on-premise, has been moved to an Amazon Lambda function.\
> Which of the following should you implement to connect AWS Lambda function to its RDS instance?

Answer

> Configure Lambda to connect to VPC with private subnet and Security Group needed to access RDS



### Question 57

Question

> A developer has created a new Application Load Balancer but has not registered any targets with the target groups.\
> Which of the following errors would be generated by the Load Balancer?

Answer

> HTTP 503: Service unavailable





### Question 58

Question

> After a test deployment in ElasticBeanstalk environment, a developer noticed that all accumulated Amazon EC2 burst balances were lost.\
> Which of the following options can lead to this behavior?

Answer

> The deployment was either run with immutable updates or in traffic splitting mode



### Question 59

Question

> A website serves static content from an Amazon Simple Storage Service (Amazon S3) bucket and dynamic content from an application load balancer. The user base is spread across the world and latency should be minimized for a better user experience.\
> Which technology/service can help access the static and dynamic content while keeping the data latency low?

Answer

> Configure CloudFront with multiple origins to serve both static and dynamic content at low latency to global users



### Question 60

Question

> A serverless application built on AWS processes customer orders 24/7 using an AWS Lambda function and communicates with an external vendor's HTTP API for payment processing. The development team wants to notify the support team in near real-time using an existing Amazon Simple Notification Service (Amazon SNS) topic, but only when the external API error rate exceeds 5% of the total transactions processed in an hour.\
> As an AWS Certified Developer Associate, which option will you suggest as the most efficient solution?

Answer

> Configure and push high-resolution custom metrics to CloudWatch that record the failures of the external payment processing API calls. Create a CloudWatch alarm that sends a notification via the existing SNS topic when the error rate exceeds the specified rate



### Question 61

Question

> The development team at an e-commerce company is preparing for the upcoming Thanksgiving sale. The product manager wants the development team to implement appropriate caching strategy on Amazon ElastiCache to withstand traffic spikes on the website during the sale. A key requirement is to facilitate consistent updates to the product prices and product description, so that the cache never goes out of sync with the backend.\
> As a Developer Associate, which of the following solutions would you recommend for the given use-case?

Answer

> Use a caching strategy to write to the backend first and then invalidate the cache





### Question 62

Question

> A company has more than 100 million members worldwide enjoying 125 million hours of TV shows and movies each day. The company uses AWS for nearly all its computing and storage needs, which use more than 10,000 server instances on AWS. This results in an extremely complex and dynamic networking environment where applications are constantly communicating inside AWS and across the Internet. Monitoring and optimizing its network is critical for the company.\
> The company needs a solution for ingesting and analyzing the multiple terabytes of real-time data its network generates daily in the form of flow logs. Which technology/service should the company use to ingest this data economically and has the flexibility to direct this data to other downstream systems?

Answer

> Amazon Kinesis Data Streams


### Question 63

Question

> A company developed an app-based service for citizens to book transportation rides in the local community. The platform is running on AWS EC2 instances and uses Amazon Relational Database Service (RDS) for storing transportation data.\
> A new feature has been requested where receipts would be emailed to customers with PDF attachments retrieved from Amazon Simple Storage Service (S3).\
> Which of the following options will provide EC2 instances with the right permissions to upload files to Amazon S3 and generate S3 Signed URL?

Answer

> Create an IAM Role for EC2



### Question 64

Question

> A junior developer has been asked to configure access to an Amazon EC2 instance hosting a web application. The developer has configured a new security group to permit incoming HTTP traffic from 0.0.0.0/0 and retained any default outbound rules.\
> A custom Network Access Control List (NACL) connected with the instance's subnet is configured to permit incoming HTTP traffic from 0.0.0.0/0 and retained any default outbound rules.\
> Which of the following solutions would you suggest if the EC2 instance needs to accept and respond to requests from the internet?

Answer

> An outbound rule must be added to the Network ACL (NACL) to allow the response to be sent to the client on the ephemeral port range



### Question 65

Question

> A junior developer working on ECS instances terminated a container instance in Amazon Elastic Container Service (Amazon ECS) as per instructions from the team lead. But the container instance continues to appear as a resource in the ECS cluster.\
> As a Developer Associate, which of the following solutions would you recommend to fix this behavior?

Answer

> You terminated the container instance while it was in STOPPED state, that lead to this synchronization issues

</details>

## Exam Set 1 
<details>

### Question 01

Question

> A development team lead is responsible for managing access for her IAM principals. At the start of the cycle, she has granted excess privileges to users to keep them motivated for trying new things. She now wants to ensure that the team has only the minimum permissions required to finish their work.\
> Which of the following will help her identify unused IAM roles and remove them without disrupting any service?

Answer

> Access Advisor feature on IAM console



### Question 02

Question

> As a developer, you are working on creating an application using AWS Cloud Development Kit (CDK).\
> Which of the following represents the correct order of steps to be followed for creating an app using AWS CDK?

Answer

> Create the app from a template provided by AWS CDK -> Add code to the app to create resources within stacks -> Build the app (optional) -> Synthesize one or more stacks in the app -> Deploy stack(s) to your AWS account



### Question 03

Question

> An IT company is configuring Auto Scaling for its Amazon EC2 instances spread across different AZs and Regions.\
> Which of the following scenarios are NOT correct about EC2 Auto Scaling? (Select two)

Answer

> - Auto Scaling groups that span across multiple Regions need to be enabled for all the Regions specified 
> - An Auto Scaling group can contain EC2 instances in only one Availability Zone of a Region


### Question 04

Question

> A developer has been asked to create a web application to be deployed on EC2 instances. The developer just wants to focus on writing application code without worrying about server provisioning, configuration and deployment.\
> As a Developer Associate, which AWS service would you recommend for the given use-case?

Answer

> Elastic Beanstalk





### Question 05

Question

> Amazon Simple Queue Service (SQS) has a set of APIs for various actions supported by the service.\
> As a developer associate, which of the following would you identify as correct regarding the CreateQueue API? (Select two)

Answer

> - You can't change the queue type after you create it
> - The visibility timeout value for the queue is in seconds, which defaults to 30 seconds



### Question 06

Question

> A developer is testing Amazon Simple Queue Service (SQS) queues in a development environment. The queue along with all its contents has to be deleted after testing.\
> Which SQS API should be used for this requirement?

Answer

> DeleteQueue 



### Question 07

Question

> When running a Rolling deployment in Elastic Beanstalk environment, only two batches completed the deployment successfully, while rest of the batches failed to deploy the updated version. Following this, the development team terminated the instances from the failed deployment.\
> What will be the status of these failed instances post termination?

Answer

> Elastic Beanstalk will replace them with instances running the application version from the most recent successful deployment



### Question 08

Question

> A developer has been asked to create an application that can be deployed across a fleet of EC2 instances. The configuration must allow for full control over the deployment steps using the blue-green deployment.\
> Which service will help you achieve that?

Answer

> CodeDeploy



### Question 09

Question

> You are a developer working on AWS Lambda functions that are invoked via REST API's using Amazon API Gateway. Currently, when a GET request is invoked by the consumer, the entire data-set returned by the Lambda function is visible. Your team lead asked you to format the data response.\
> Which feature of the API Gateway can be used to solve this issue?

Answer

> Use API Gateway Mapping Templates



### Question 10

Question

> The development team has just configured and attached the IAM policy needed to access AWS Billing and Cost Management for all users under the Finance department. But, the users are unable to see AWS Billing and Cost Management service in the AWS console.\
> What could be the reason for this issue?

Answer

> You need to activate IAM user access to the Billing and Cost Management console for all the users who need access



### Question 11 

Question

> A SaaS company runs a HealthCare web application that is used worldwide by users. There have been requests by mobile developers to expose public APIs for the application-specific functionality. You decide to make the APIs available to mobile developers as product offerings.\
> Which of the following options will allow you to do that?

Answer

> Use API Gateway Usage Plans





### Question 12

Question

> A company is creating a gaming application that will be deployed on mobile devices. The application will send data to a Lambda function-based RESTful API. The application will assign each API request a unique identifier.\
> The volume of API requests from the application can randomly vary at any given time of day. During request throttling, the application might need to retry requests. The API must be able to address duplicate requests without inconsistencies or data loss.\
> Which of the following would you recommend to handle these requirements?

Answer

> Persist the unique identifier for each request in a DynamoDB table. Change the Lambda function to check the table for the identifier before processing the request



### Question 13

Question

> A company uses Elastic Beanstalk to manage its IT infrastructure on AWS Cloud and it would like to deploy the new application version to the EC2 instances. When the deployment is executed, some instances should serve requests with the old application version, while other instances should serve requests using the new application version until the deployment is completed.\
> Which deployment meets this requirement without incurring additional costs?

Answer

> Rolling



### Question 14

Question

> A gaming company wants to store information about all the games that the company has released. Each game has a name, version number, and category (such as sports, puzzles, strategy, etc). The game information also can include additional properties about the supported platforms and technical specifications. This additional information is inconsistent across games.\
> You have been hired as an AWS Certified Developer Associate to build a solution that addresses the following use cases:
> - For a given name and version number, get all details about the game that has that name and version number.
> - For a given name, get all details about all games that have that name.
> - For a given category, get all details about all games in that category.
>
> What will you recommend as the most efficient solution?

Answer

> Set up an Amazon DynamoDB table with a primary key that consists of the name as the partition key and the version number as the sort key. Create a global secondary index that has the category as the partition key and the name as the sort key



### Question 15

Question

> An organization has hosted its EC2 instances in two AZs. AZ1 has two instances and AZ2 has 8 instances. The Elastic Load Balancer managing the instances in the two AZs has cross-zone load balancing enabled in its configuration.\
> What percentage traffic will each of the instances in AZ1 receive?

Answer

> 10 - Cross-zone load Balancing in enabled



### Question 16

Question

> A development team has configured inbound traffic for the relevant ports in both the Security Group of the EC2 instance as well as the Network Access Control List (NACL) of the subnet for the EC2 instance. The team is, however, unable to connect to the service running on the Amazon EC2 instance.\
> As a developer associate, which of the following will you recommend to fix this issue?

Answer

> Security Groups are stateful, so allowing inbound traffic to the necessary ports enables the connection. Network ACLs are stateless, so you must allow both inbound and outbound traffic



### Question 17

Question

> You are a developer in a manufacturing company that has several servers on-site. The company decides to move new development to the cloud using serverless technology. You decide to use the AWS Serverless Application Model (AWS SAM) and work with an AWS SAM template file to represent your serverless architecture.\
> Which of the following is NOT a valid serverless resource type?

Answer

> AWS::Serverless::UserPool





### Question 18

Question

> A retail company is migrating its on-premises database to Amazon RDS for PostgreSQL. The company has read-heavy workloads. The development team at the company is looking at refactoring the code to achieve optimum read performance for SQL queries.\
> Which solution will address this requirement with the least current as well as future development effort?


Answer

> Set up Amazon RDS with one or more read replicas. Refactor the application code so that the queries use the endpoint for the read replicas



### Question 19

Question

> A development team wants to build an application using serverless architecture. The team plans to use AWS Lambda functions extensively to achieve this goal. The developers of the team work on different programming languages like Python, .NET and Javascript. The team wants to model the cloud infrastructure using any of these programming languages.\
> Which AWS service/tool should the team use for the given use-case?

Answer

> AWS Cloud Development Kit (CDK) 



### Question 20

Question

> The manager at an IT company wants to set up member access to user-specific folders in an Amazon S3 bucket - bucket-a. So, user x can only access files in his folder - bucket-a/user/user-x/ and user y can only access files in her folder - bucket-a/user/user-y/ and so on.\
> As a Developer Associate, which of the following IAM constructs would you recommend so that the policy snippet can be made generic for all team members and the manager does not need to create separate IAM policy for each team member?

Answer

> IAM policy variables





### Question 21

Question

> To enable HTTPS connections for his web application deployed on the AWS Cloud, a developer is in the process of creating server certificate.\
> Which AWS entities can be used to deploy SSL/TLS server certificates? (Select two)

Answer

> - AWS Certificate Manager
> - IAM 





### Question 22

Question

> You have deployed a Java application to an EC2 instance where it uses the X-Ray SDK. When testing from your personal computer, the application sends data to X-Ray but when the application runs from within EC2, the application fails to send data to X-Ray.\
> Which of the following does NOT help with debugging the issue?

Answer

> X-Ray sampling



### Question 23

Question

> After a test deployment in ElasticBeanstalk environment, a developer noticed that all accumulated Amazon EC2 burst balances were lost.\
> Which of the following options can lead to this behavior?

Answer

> The deployment was either run with immutable updates or in traffic splitting mode



### Question 24

Question

> An application is hosted by a 3rd party and exposed at yourapp.3rdparty.com. You would like to have your users access your application using www.mydomain.com, which you own and manage under Route 53.\
> What Route 53 record should you create?

Answer

> CNAME



### Question 25

Question

> You have chosen AWS Elastic Beanstalk to upload your application code and allow it to handle details such as provisioning resources and monitoring.\
> When creating configuration files for AWS Elastic Beanstalk which naming convention should you follow?

Answer

> `.ebextensions/<mysettings>.config`



### Question 26

Question

> A developer has an application that stores data in an Amazon S3 bucket. The application uses an HTTP API to store and retrieve objects. When the PutObject API operation adds objects to the S3 bucket the developer must encrypt these objects at rest by using server-side encryption with Amazon S3-managed keys (SSE-S3).\
> Which solution will guarantee that any upload request without the mandated encryption is not processed?


Answer

> Invoke the PutObject API operation and set the `x-amz-server-side-encryption` header as `AES256`. Use an S3 bucket policy to deny permission to upload an object unless the request has this header



### Question 27

Question

> A data analytics company processes Internet-of-Things (IoT) data using Amazon Kinesis. The development team has noticed that the IoT data feed into Kinesis experiences periodic spikes. The PutRecords API call occasionally fails and the logs show that the failed call returns the response shown below:
```json
HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
    "FailedRecordCount": 2,
    "Records": [
        {
            "SequenceNumber": "49543463076548007577105092703039560359975228518395012686",
            "ShardId": "shardId-000000000000"
        },
        {
            "ErrorCode": "ProvisionedThroughputExceededException",
            "ErrorMessage": "Rate exceeded for shard shardId-000000000001 in stream exampleStreamName under account 111111111111."
        },
        {
            "ErrorCode": "InternalFailure",
            "ErrorMessage": "Internal service failure."
        }
    ]
}
```
> As an AWS Certified Developer Associate, which of the following options would you recommend to address this use case? (Select two)

Answer

> - Use an error retry and exponential backoff mechanism
> - Decrease the frequency or size of your requests



### Question 28

Question

> Your global organization has an IT infrastructure that is deployed using CloudFormation on AWS Cloud. One employee, in us-east-1 Region, has created a stack 'Application1' and made an exported output with the name 'ELBDNSName'. Another employee has created a stack for a different application 'Application2' in us-east-2 Region and also exported an output with the name 'ELBDNSName'.\
> The first employee wanted to deploy the CloudFormation stack 'Application1' in us-east-2, but it got an error. What is the cause of the error?

Answer

> Exported Output Values in CloudFormation must have unique names within a single Region



### Question 29

Question

> The development team at a company creates serverless solutions using AWS Lambda. Functions are invoked by clients via AWS API Gateway which anyone can access. The team lead would like to control access using a 3rd party authorization mechanism.\
> As a Developer Associate, which of the following options would you recommend for the given use-case?

Answer

> Lambda Authorizer



### Question 30

Question

> Which of the following best describes how KMS Encryption works?

Answer

> KMS stores the CMK, and receives data from the clients, which it encrypts and sends back


### Question 31

Question

> A company has built its technology stack on AWS serverless architecture for managing all its business functions. To expedite development for a new business requirement, the company is looking at using pre-built serverless applications.\
> Which AWS service represents the easiest solution to address this use-case?



Answer

> AWS Serverless Application Repository (SAR)



### Question 32

Question

> ECS Fargate container tasks are usually spread across Availability Zones (AZs) and the underlying workloads need persistent cross-AZ shared access to the data volumes configured for the container tasks.\
> Which of the following solutions is the best choice for these workloads?

Answer

> Amazon EFS volumes



### Question 33

Question

> A company wants to improve the performance of its popular API service that offers unauthenticated read access to daily updated statistical information via Amazon API Gateway and AWS Lambda.\
> What measures can the company take?

Answer

> Enable API caching in API Gateway



### Question 34

Question

> An E-commerce business, has its applications built on a fleet of Amazon EC2 instances, spread across various Regions and AZs. The technical team has suggested using Elastic Load Balancers for better architectural design.\
> What characteristics of an Elastic Load Balancer make it a winning choice? (Select two)

Answer

> - Separate public traffic from private traffic 
> - Build a highly available system



### Question 35

Question

> A Developer has been entrusted with the job of securing certain S3 buckets that are shared by a large team of users. Last time, a bucket policy was changed, the bucket was erroneously available for everyone, outside the organization too.\
> Which feature/service will help the developer identify similar security issues with minimum effort

Answer

> IAM Access Analyzer



### Question 36

Question

> You are running workloads on AWS and have embedded RDS database connection strings within each web server hosting your applications. After failing a security audit, you are looking at a different approach to store your secrets securely and automatically rotate the database credentials.\
> Which AWS service can you use to address this use-case?

Answer

> Secrets Manager



### Question 37

Question

> A development team at a social media company uses AWS Lambda for its serverless stack on AWS Cloud. For a new deployment, the Team Lead wants to send only a certain portion of the traffic to the new Lambda version. In case the deployment goes wrong, the solution should also support the ability to roll back to a previous version of the Lambda function, with MIMINUM downtime for the application.\
> As a Developer Associate, which of the following options would you recommend to address this use-case?

Answer

> Set up the application to use an alias that points to the current version. Deploy the new version of the code and configure the alias to send 10% of the users to this new version. If the deployment goes wrong, reset the alias to point all traffic to the current version



### Question 38

Question

> Your company has configured AWS Organizations to manage multiple AWS accounts. Within each AWS account, there are many CloudFormation scripts running. Your manager has requested that each script output the account number of the account the script was executed in.\
> Which Pseudo parameter will you use to get this information?

Answer

> AWS::AccountId



### Question 39

Question

> CodeCommit is a managed version control service that hosts private Git repositories in the AWS cloud.\
> Which of the following credential types is NOT supported by IAM for CodeCommit?

Answer

> IAM username and password



### Question 40

Question

> A company wants to provide beta access to some developers on its development team for a new version of the company's Amazon API Gateway REST API, without causing any disturbance to the existing customers who are using the API via a frontend UI and Amazon Cognito authentication. The new version has new endpoints and backward-incompatible interface changes, and the company's development team is responsible for its maintenance.\
> Which of the following will satisfy these requirements in the MOST operationally efficient manner?

Answer

> Create a development stage on the API Gateway API and then have the developers point the endpoints to the development stage



### Question 41

Question

> A cybersecurity firm wants to run their applications on single-tenant hardware to meet security guidelines.\
> Which of the following is the MOST cost-effective way of isolating their Amazon EC2 instances to a single tenant?

Answer

> Dedicated Instances - Dedicated Instances are Amazon EC2 instances that run in a virtual private cloud (VPC) on hardware that's dedicated to a single customer. Dedicated Instances that belong to different AWS accounts are physically isolated at a hardware level, even if those accounts are linked to a single-payer account. However, Dedicated Instances may share hardware with other instances from the same AWS account that are not Dedicated Instances.\
> A Dedicated Host is also a physical server that's dedicated for your use. With a Dedicated Host, you have visibility and control over how instances are placed on the server.



### Question 42

Question

> An e-commerce company has developed an API that is hosted on Amazon ECS. Variable traffic spikes on the application are causing order processing to take too long. The application processes orders using Amazon SQS queues. The ApproximateNumberOfMessagesVisible metric spikes at very high values throughout the day which triggers the CloudWatch alarm. Other ECS metrics for the API containers are well within limits.\
> As a Developer Associate, which of the following will you recommend for improving performance while keeping costs low?

Answer

> Use backlog per instance metric with target tracking scaling policy



### Question 43

Question

> You are a developer for a web application written in .NET which uses the AWS SDK. You need to implement an authentication mechanism that returns a JWT (JSON Web Token).\
> Which AWS service will help you with token handling and management?

Answer

> Cognito User Pools



### Question 44

Question

> A media company has created a video streaming application and it would like their Brazilian users to be served by the company's Brazilian servers. Other users around the globe should not be able to access the servers through DNS queries.\
> Which Route 53 routing policy meets this requirement?

Answer

> Geolocation



### Question 45

Question

> Your company has stored all application secrets in SSM Parameter Store. The audit team has requested to get a report to better understand when and who has issued API calls against SSM Parameter Store.\
> Which of the following options can be used to produce your report?

Answer

> Use AWS CloudTrail to get a record of actions taken by a user



### Question 46

Question

> As part of his development work, an AWS Certified Developer Associate is creating policies and attaching them to IAM identities. After creating necessary Identity-based policies, he is now creating Resource-based policies.\
> Which is the only resource-based policy that the IAM service supports?

Answer

> Trust policy



### Question 47

Question

> A developer is configuring a bucket policy that denies upload object permission to any requests that do not include the x-amz-server-side-encryption header requesting server-side encryption with SSE-KMS for an Amazon S3 bucket - examplebucket.\
> Which of the following policies is the right fit for the given requirement?

Answer

```json
{ "Version":"2012-10-17", "Id":"PutObjectPolicy", "Statement":[{ "Sid":"DenyUnEncryptedObjectUploads", "Effect":"Deny", "Principal":"", "Action":"s3:PutObject", "Resource":"arn:aws:s3:::examplebucket/", "Condition":{ "StringNotEquals":{ "s3:x-amz-server-side-encryption":"aws:kms" } } } ] }
```




### Question 48

Question

> You are creating a Cloud Formation template to deploy your CMS application running on an EC2 instance within your AWS account. Since the application will be deployed across multiple regions, you need to create a map of all the possible values for the base AMI.\
> How will you invoke the !FindInMap function to fulfill this use case?

Answer

> `!FindInMap [ MapName, TopLevelKey, SecondLevelKey ]`



### Question 49

Question

> A multi-national company has multiple business units with each unit having its own AWS account. The development team at the company would like to debug and trace data across accounts and visualize it in a centralized account.\
> As a Developer Associate, which of the following solutions would you suggest for the given use-case?

Answer

> X-Ray



### Question 50

Question

> A multi-national company has just moved to AWS Cloud and it has configured forecast-based AWS Budgets alerts for cost management. However, no alerts have been received even though the account and the budgets have been created almost three weeks ago.\
> What could be the issue with the AWS Budgets configuration?

Answer

> AWS requires approximately 5 weeks of usage data to generate budget forecasts



### Question 51

Question

> You have created a Java application that uses RDS for its main data storage and ElastiCache for user session storage. The application needs to be deployed using Elastic Beanstalk and every new deployment should allow the application servers to reuse the RDS database. On the other hand, user session data stored in ElastiCache can be re-created for every deployment.\
> Which of the following configurations will allow you to achieve this? (Select two)

Answer

> - ElastiCache defined in .ebextensions/
> - RDS database defined externally and referenced through environment variables 



### Question 52

Question

> As an AWS Certified Developer Associate, you have configured the AWS CLI on your workstation. Your default region is us-east-1 and your IAM user has permissions to operate commands on services such as EC2, S3 and RDS in any region. You would like to execute a command to stop an EC2 instance in the us-east-2 region.\
> What of the following is the MOST optimal solution to address this use-case?

Answer

> Use the `--region` parameter



### Question 53

Question

> You are storing bids information on your betting application and you would like to automatically expire DynamoDB table data after one week.\
> What should you use?


Answer

> Use TTL





### Question 54

Question

> You have created an Elastic Load Balancer that has marked all the EC2 instances in the target group as unhealthy. Surprisingly, when you enter the IP address of the EC2 instances in your web browser, you can access your website.\
> What could be the reason your instances are being marked as unhealthy? (Select two)


Answer

> - The security group of the EC2 instance does not allow for traffic from the security group of the Application Load Balancer
> - The route for the health check is misconfigured



### Question 55

Question

> A development team wants to deploy an AWS Lambda function that requires significant CPU utilization.\
> As a Developer Associate, which of the following would you suggest for reducing the average runtime of the function?

Answer

> Deploy the function with its memory allocation set to the maximum amount -

### Question 56

Question

> A global e-commerce company wants to perform geographic load testing of its order processing API. The company must deploy resources to multiple AWS Regions to support the load testing of the API.\
> How can the company address these requirements without additional application code?


Answer

> Set up an AWS CloudFormation template that defines the load test resources. Leverage the AWS CLI create-stack-set command to create a stack set in the desired Regions

### Question 57

Question

> An organization has offices across multiple locations and the technology team has configured an Application Load Balancer across targets in multiple Availability Zones. The team wants to analyze the incoming requests for latencies and the client's IP address patterns.\
> Which feature of the Load Balancer will help collect the required information?

Answer

> ALB access logs
>
> 
### Question 58

Question

> A startup with newly created AWS account is testing different EC2 instances. They have used Burstable performance instance - T2.micro - for 35 seconds and stopped the instance.\
> At the end of the month, what is the instance usage duration that the company is charged for?

Answer

> 0 - free tier

### Question 59

Question

> A firm runs its technology operations on a fleet of Amazon EC2 instances. The firm needs a certain software to be available on the instances to support their daily workflows. The developer team has been told to use the user data feature of EC2 instances.\
> Which of the following are true about the user data EC2 configuration? ( Select two)

Answer

> - By default, scripts entered as user data are executed with root user privileges
> - By default, user data runs only during the boot cycle when you first launch an instance

### Question 60

Question

> As an AWS Certified Developer Associate, you are given a document written in YAML that represents the architecture of a serverless application. The first line of the document contains Transform: 'AWS::Serverless-2016-10-31'.\
> What does the Transform section in the document represent?

Answer

> Presence of 'Transform' section indicates it is a Serverless Application Model (SAM) template

### Question 61

Question

> Which of the following security credentials can only be created by the AWS Account root user?

Answer

> CloudFront Key Pairs

### Question 62

Question

> You're a developer working on a large scale order processing application. After developing the features, you commit your code to AWS CodeCommit and begin building the project with AWS CodeBuild before it gets deployed to the server. The build is taking too long and the error points to an issue resolving dependencies from a third-party. You would like to prevent a build running this long in the future for similar underlying reasons.\
> Which of the following options represents the best solution to address this use-case?

Answer

> Enable CodeBuild timeouts

### Question 63

Question

> A Developer at a company is working on a CloudFormation template to set up resources. Resources will be defined using code and provisioned based on certain conditions defined in the Conditions section.\
> Which section of a CloudFormation template cannot be associated with Condition?

Answer

> Parameters

### Question 64

Question

> As an AWS Certified Developer Associate, you have been asked to create an AWS Elastic Beanstalk environment to handle deployment for an application that has high traffic and high availability needs. You need to deploy the new version using Beanstalk while making sure that performance and availability are not affected.\
> Which of the following is the MOST optimal way to do this while keeping the solution cost-effective?

Answer

> Deploy using 'Rolling with additional batch' deployment policy

### Question 65

Question

> The Technical Lead of your team has reviewed a CloudFormation YAML template written by a new recruit and specified that an invalid section has been added to the template.\
> Which of the following represents an invalid section of the CloudFormation template?

Answer

> 'Dependencies' section of the template







</details>

## Exam Set 2
<details>

### Question 01

Question

> An e-commerce company uses AWS CloudFormation to implement Infrastructure as Code for the entire organization. Maintaining resources as stacks with CloudFormation has greatly reduced the management effort needed to manage and maintain the resources. However, a few teams have been complaining of failing stack updates owing to out-of-band fixes running on the stack resources.\
> Which of the following is the best solution that can help in keeping the CloudFormation stack and its resources in sync with each other?

Answer

> Use Drift Detection feature of CloudFormation

### Question 02

Question

> A serverless application built on AWS processes customer orders 24/7 using an AWS Lambda function and communicates with an external vendor's HTTP API for payment processing. The development team wants to notify the support team in near real-time using an existing Amazon Simple Notification Service (Amazon SNS) topic, but only when the external API error rate exceeds 5% of the total transactions processed in an hour.\
> As an AWS Certified Developer Associate, which option will you suggest as the most efficient solution?

Answer

> Configure and push high-resolution custom metrics to CloudWatch that record the failures of the external payment processing API calls. Create a CloudWatch alarm that sends a notification via the existing SNS topic when the error rate exceeds the specified rate

### Question 03

Question

> The app development team at a social gaming mobile app wants to simplify the user sign up process for the app. The team is looking for a fully managed scalable solution for user management in anticipation of the rapid growth that the app foresees.\
> As a Developer Associate, which of the following solutions would you suggest so that it requires the LEAST amount of development effort?

Answer

> Use Cognito User pools to facilitate sign up and user management for the mobile app

### Question 04

Question

> You create an Auto Scaling group to work with an Application Load Balancer. The scaling group is configured with a minimum size value of 5, a maximum value of 20, and the desired capacity value of 10. One of the 10 EC2 instances has been reported as unhealthy.\
> Which of the following actions will take place?

Answer

> The ASG will terminate the EC2 Instance

### Question 05

Question

> An e-commerce company manages a microservices application that receives orders from various partners through a customized API for each partner exposed via Amazon API Gateway. The orders are processed by a shared Lambda function.\
> How can the company notify each partner regarding the status of their respective orders in the most efficient manner, without affecting other partners' orders? Also, the solution should be scalable to accommodate new partners with minimal code changes required.

Answer 

> Set up an SNS topic and subscribe each partner to the SNS topic. Modify the Lambda function to publish messages with specific attributes to the SNS topic and apply the appropriate filter policy to the topic subscription.

### Question 06

Question

> A developer is looking at establishing access control for an API that connects to a Lambda function downstream.\
> Which of the following represents a mechanism that CANNOT be used for authenticating with the API Gateway?

Answer

> AWS Security Token Service (STS).

### Question 07

Question

> You are a developer working on a web application written in Java and would like to use AWS Elastic Beanstalk for deployment because it would handle deployment, capacity provisioning, load balancing, auto-scaling, and application health monitoring.\
> In the past, you connected to your provisioned instances through SSH to issue configuration commands. Now, you would like a configuration mechanism that automatically applies settings for you.\
> Which of the following options would help do this?

Answer

> Include config files in .ebextensions/ at the root of your source code



### Question 08

Question

> The development team at an analytics company is using SQS queues for decoupling the various components of application architecture. As the consumers need additional time to process SQS messages, the development team wants to postpone the delivery of new messages to the queue for a few seconds.\
> As a Developer Associate, which of the following solutions would you recommend to the development team?

Answer

> Use delay queues to postpone the delivery of new messages to the queue for a few seconds

### Question 09

Question

> As a senior architect, you are responsible for the development, support, maintenance, and implementation of all database applications written using NoSQL technology. A new project demands a throughput requirement of 10 strongly consistent reads per second of 6KB in size each.\
> How many read capacity units will you need when configuring your DynamoDB table?

Answer

> 20
> - calculation $\frac{10}{1}* \lceil{\frac{6}{4}\rceil} = 10 * 2 = 20$
> - strong consistent read $items* \lceil{\frac{itemSize}{4}\rceil}$
> - eventually consistent read $\frac{items}{2}* \lceil{\frac{itemSize}{4}\rceil}$
> - writes $items* \lceil{\frac{itemSize}{1}\rceil}$

### Question 10

Question

> A social gaming application supports the transfer of gift vouchers between users. When a user hits a certain milestone on the leaderboard, they earn a gift voucher that can be redeemed or transferred to another user.\
> The development team wants to ensure that this transfer is captured in the database such that the records for both users are either written successfully with the new gift vouchers or the status quo is maintained.\
> Which of the following solutions represent the best-fit options to meet the requirements for the given use-case? (Select two)

Answer

> - Use the DynamoDB transactional read and write APIs on the table items as a single, all-or-nothing operation
> - Complete both operations on RDS MySQL in a single transaction block



### Question 11

Question

> A developer wants to securely store and retrieve various types of variables, such as remote API authentication information, API URL, and related credentials across different environments of an application deployed on Amazon Elastic Container Service (Amazon ECS).\
> What would be the best approach that needs minimal modifications in the application code?

Answer

> Configure the application to fetch the variables and credentials from AWS Systems Manager Parameter Store by leveraging hierarchical unique paths in Parameter Store for each variable in each environment

### Question 12

Question

> A diagnostic lab stores its data on DynamoDB. The lab wants to backup a particular DynamoDB table data on Amazon S3, so it can download the S3 backup locally for some operational use.\
> Which of the following options is NOT feasible?

Answer

> Use the DynamoDB on-demand backup capability to write to Amazon S3 and download locally.\
> This option is not feasible for the given use-case. DynamoDB has two built-in backup methods (On-demand, Point-in-time recovery) that write to Amazon S3, but you will not have access to the S3 buckets that are used for these backups.

### Question 13

Question

> A CRM application is hosted on Amazon EC2 instances with the database tier using DynamoDB. The customers have raised privacy and security concerns regarding sending and receiving data across the public internet.\
> As a developer associate, which of the following would you suggest as an optimal solution for providing communication between EC2 instances and DynamoDB without using the public internet?

Answer

> Configure VPC endpoints for DynamoDB that will provide required internal access without using public internet

### Question 14

Question

> A company has created an Amazon S3 bucket that holds customer data. The team lead has just enabled access logging to this bucket. The bucket size has grown substantially after starting access logging. Since no new files have been added to the bucket, the perplexed team lead is looking for an answer.\
> Which of the following reasons explains this behavior?

Answer

> S3 access logging is pointing to the same bucket and is responsible for the substantial growth of bucket size

### Question 15

Question

> A developer in your company was just promoted to Team Lead and will be in charge of code deployment on EC2 instances via AWS CodeCommit and AWS CodeDeploy. Per the new requirements, the deployment process should be able to change permissions for deployed files as well as verify the deployment success.\
> Which of the following actions should the new Developer take?

Answer

> Define an appspec.yml file in the root directory.\
> An AppSpec file must be a YAML-formatted file named appspec.yml and it must be placed in the root of the directory structure of an application's source code.\
> The AppSpec file is used to:
> - Map the source files in your application revision to their destinations on the instance.
> - Specify custom permissions for deployed files.
> - Specify scripts to be run on each instance at various stages of the deployment process.
> - During deployment, the CodeDeploy agent looks up the name of the current event in the hooks section of the AppSpec file. If the event is not found, the CodeDeploy agent moves on to the next step. If the event is found, the CodeDeploy agent retrieves the list of scripts to execute. The scripts are run sequentially, in the order in which they appear in the file. The status of each script is logged in the CodeDeploy agent log file on the instance.

### Question 16

Question

> What steps can a developer take to optimize the performance of a CPU-bound AWS Lambda function and ensure fast response time?

Answer

> Increase the function's memory



### Question 17

Question

> A developer wants to package the code and dependencies for the application-specific Lambda functions as container images to be hosted on Amazon Elastic Container Registry (ECR).\
> Which of the following options are correct for the given requirement? (Select two)



Answer

> - To deploy a container image to Lambda, the container image must implement the Lambda Runtime API
> - AWS Lambda service does not support Lambda functions that use multi-architecture container images

### Question 18

Question

> The development team at a multi-national retail company wants to support trusted third-party authenticated users from the supplier organizations to create and update records in specific DynamoDB tables in the company's AWS account.\
> As a Developer Associate, which of the following solutions would you suggest for the given use-case?

Answer

> Use Cognito Identity pools to enable trusted third-party authenticated users to access DynamoDB

### Question 19

Question

> As an AWS certified developer associate, you are working on an AWS CloudFormation template that will create resources for a company's cloud infrastructure. Your template is composed of three stacks which are Stack-A, Stack-B, and Stack-C. Stack-A will provision a VPC, a security group, and subnets for public web applications that will be referenced in Stack-B and Stack-C.\
> After running the stacks you decide to delete them, in which order should you do it?

Answer

> Stack B, then Stack C, then Stack A



### Question 20

Question

> A pharmaceutical company uses Amazon EC2 instances for application hosting and Amazon CloudFront for content delivery. A new research paper with critical findings has to be shared with a research team that is spread across the world.

Which of the following represents the most optimal solution to address this requirement without compromising the security of the content?


Answer

> Use CloudFront signed URL feature to control access to the file



### Question 21

Question

> The development team at a retail company is gearing up for the upcoming Thanksgiving sale and wants to make sure that the application's serverless backend running via Lambda functions does not hit latency bottlenecks as a result of the traffic spike.

As a Developer Associate, which of the following solutions would you recommend to address this use-case?

Answer

> Configure Application Auto Scaling to manage Lambda provisioned concurrency on a schedule

### Question 22

Question

> While troubleshooting, a developer realized that the Amazon EC2 instance is unable to connect to the Internet using the Internet Gateway.\
> Which conditions should be met for Internet connectivity to be established? (Select two)

Answer

> - The network ACLs associated with the subnet must have rules to allow inbound and outbound traffic
> - The route table in the instance’s subnet should have a route to an Internet Gateway

### Question 23

Question

> Your company has embraced cloud-native microservices architectures. New applications must be dockerized and stored in a registry service offered by AWS. The architecture should support dynamic port mapping and support multiple tasks from a single service on the same container instance. All services should run on the same EC2 instance.\
> Which of the following options offers the best-fit solution for the given use-case?

Answer

> Application Load Balancer + ECS



### Question 24

Question

> A company runs its flagship application on a fleet of Amazon EC2 instances. After misplacing a couple of private keys from the SSH key pairs, they have decided to re-use their SSH key pairs for the different instances across AWS Regions.\
> As a Developer Associate, which of the following would you recommend to address this use-case?

Answer

> Generate a public SSH key from a private SSH key. Then, import the key into each of your AWS Regions.\
> Here is the correct way of reusing SSH keys in your AWS Regions:
> 1. Generate a public SSH key (.pub) file from the private SSH key (.pem) file.
> 1. Set the AWS Region you wish to import to.
> 1. Import the public SSH key into the new Region.

### Question 25

Question

> A business has their test environment built on Amazon EC2 configured on General purpose SSD volume.\
> At which gp2 volume size will their test environment hit the max IOPS?

Answer

> 5.3 TiB - General Purpose SSD (gp2) volumes offer cost-effective storage that is ideal for a broad range of workloads.\
> These volumes deliver single-digit millisecond latencies and the ability to burst to 3,000 IOPS for extended periods of time. Between a minimum of 100 IOPS (at 33.33 GiB and below) and a maximum of 16,000 IOPS (at 5,334 GiB and above), baseline performance scales linearly at 3 IOPS per GiB of volume size.

### Question 26

Question

> A developer needs to automate software package deployment to both Amazon EC2 instances and virtual servers running on-premises, as part of continuous integration and delivery that the business has adopted.\
> Which AWS service should he use to accomplish this task?

Answer

> AWS CodeDeploy

### Question 27

Question

> A developer working with EC2 Windows instance has installed Kinesis Agent for Windows to stream JSON-formatted log files to Amazon Simple Storage Service (S3) via Amazon Kinesis Data Firehose. The developer wants to understand the sink type capabilities of Kinesis Firehose.\
> Which of the following sink types is NOT supported by Kinesis Firehose.

Answer

> Amazon ElastiCache with Amazon S3 as backup.

### Question 28

Question

> A developer working with EC2 Windows instance has installed Kinesis Agent for Windows to stream JSON-formatted log files to Amazon Simple Storage Service (S3) via Amazon Kinesis Data Firehose. The developer wants to understand the sink type capabilities of Kinesis Firehose.\
> Which of the following sink types is NOT supported by Kinesis Firehose.

Answer

> Leverage Cognito user pools to manage user accounts and set up an Amazon Cognito user pool authorizer in API Gateway to control access to the API. Set up a Lambda function to store the images in Amazon S3 and save the image object's S3 key as part of the photo details in a DynamoDB table. Have the Lambda function retrieve previously uploaded images by querying DynamoDB for the S3 key.\
> A user pool is a user directory in Amazon Cognito. With a user pool, your users can sign in to your web or mobile app through Amazon Cognito. Your users can also sign in through social identity providers like Google, Facebook, Amazon, or Apple, and SAML identity providers. Whether your users sign in directly or through a third party, all members of the user pool have a directory profile that you can access through a Software Development Kit (SDK).

### Question 28

Question

> Consider an application that enables users to store their mobile phone images in the cloud and supports tens of thousands of users. The application should utilize an Amazon API Gateway REST API that leverages AWS Lambda functions for photo processing while storing photo details in Amazon DynamoDB.\
> The application should allow users to create an account, upload images, and retrieve previously uploaded images, with images ranging in size from 500 KB to 5 MB.\
> How will you design the application with the least operational overhead?

Answer

> DependentParameter



### Question 29

Question

> Your team lead has asked you to learn AWS CloudFormation to create a collection of related AWS resources and provision them in an orderly fashion. You decide to provide AWS-specific parameter types to catch invalid values.\
> When specifying parameters which of the following is not a valid Parameter type?

Answer

> Create an IAM role with S3 access in Account B and set Account A as a trusted entity. Create another role (instance profile) in Account A and attach it to the EC2 instances in Account A and add an inline policy to this role to assume the role from Account B

### Question 30

Question

> The development team at a HealthCare company has deployed EC2 instances in AWS Account A. These instances need to access patient data with Personally Identifiable Information (PII) on multiple S3 buckets in another AWS Account B.\
> As a Developer Associate, which of the following solutions would you recommend for the given use-case?

Answer

> EBS volumes are AZ locked



### Question 31

Question

> A developer with access to the AWS Management Console terminated an instance in the us-east-1a availability zone. The attached EBS volume remained and is now available for attachment to other instances. Your colleague launches a new Linux EC2 instance in the us-east-1e availability zone and is attempting to attach the EBS volume. Your colleague informs you that it is not possible and need your help.\
> Which of the following explanations would you provide to them?

Answer

> EBS volumes are AZ locked


### Question 32

Question

> After a code review, a developer has been asked to make his publicly accessible S3 buckets private, and enable access to objects with a time-bound constraint.\
> Which of the following options will address the given use-case?

Answer

> Share pre-signed URLs with resources that need access

### Question 33

Question

> A company needs a version control system for their fast development lifecycle with incremental changes, version control, and support to existing Git tools.\
> Which AWS service will meet these requirements?

Answer

> AWS CodeCommit

### Question 34

Question

> While defining a business workflow as state machine on AWS Step Functions, a developer has configured several states.\
> Which of the following would you identify as the state that represents a single unit of work performed by a state machine?

Answer

```json
"HelloWorld": {
  "Type": "Task",
  "Resource": "arn:aws:lambda:us-east-1:123456789012:function:HelloFunction",
  "Next": "AfterHelloWorldState",
  "Comment": "Run the HelloWorld Lambda function"
}
```

### Question 35 

Question

> You have launched several AWS Lambda functions written in Java. A new requirement was given that over 1MB of data should be passed to the functions and should be encrypted and decrypted at runtime.\
> Which of the following methods is suitable to address the given use-case?

Answer

> Use Envelope Encryption and reference the data as file within the code

### Question 36

Question

> A development team is working on an AWS Lambda function that accesses DynamoDB. The Lambda function must do an upsert, that is, it must retrieve an item and update some of its attributes or create the item if it does not exist.\
> Which of the following represents the solution with MINIMUM IAM permissions that can be used for the Lambda function to achieve this functionality?



Answer

> `dynamodb:UpdateItem`, `dynamodb:GetItem`

### Question 37

Question

> A developer is defining the signers that can create signed URLs for their Amazon CloudFront distributions.\
> Which of the following statements should the developer consider while defining the signers? (Select two)


Answer

> - When you create a signer, the public key is with CloudFront and private key is used to sign a portion of URL
> - When you use the root user to manage CloudFront key pairs, you can only have up to two active CloudFront key pairs per AWS account

### Question 38

Question

> The development team at an e-commerce company completed the last deployment for their application at a reduced capacity because of the deployment policy. The application took a performance hit because of the traffic spike due to an on-going sale.\
> Which of the following represents the BEST deployment option for the upcoming application version such that it maintains at least the FULL capacity of the application and MINIMAL impact of failed deployment?

Answer

> Deploy the new application version using 'Immutable' deployment policy



### Question 39

Question

> The technology team at an investment bank uses DynamoDB to facilitate high-frequency trading where multiple trades can try and update an item at the same time.\
> Which of the following actions would make sure that only the last updated value of any item is used in the application?

Answer

> Use ConsistentRead = true while doing GetItem operation for any item

### Question 40

Question

> An Auto Scaling group has a maximum capacity of 3, a current capacity of 2, and a scaling policy that adds 3 instances.\
> When executing this scaling policy, what is the expected outcome?

Answer

> Amazon EC2 Auto Scaling adds only 1 instance to the group

### Question 41

Question

> As a Developer, you are given a document written in YAML that represents the architecture of a serverless application. The first line of the document contains Transform: 'AWS::Serverless-2016-10-31'.\
> What does the Transform section in the document represent?

Answer

> Presence of Transform section indicates it is a Serverless Application Model (SAM) template

### Question 42

Question

> An Accounting firm extensively uses Amazon EBS volumes for persistent storage of application data of Amazon EC2 instances. The volumes are encrypted to protect the critical data of the clients. As part of managing the security credentials, the project manager has come across a policy snippet that looks like the following:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Allow for use of this Key",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::111122223333:role/UserRole"
            },
            "Action": [
                "kms:GenerateDataKeyWithoutPlaintext",
                "kms:Decrypt"
            ],
            "Resource": "*"
        },
        {
            "Sid": "Allow for EC2 Use",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::111122223333:role/UserRole"
            },
            "Action": [
                "kms:CreateGrant",
                "kms:ListGrants",
                "kms:RevokeGrant"
            ],
            "Resource": "*",
            "Condition": {
                "StringEquals": {
                "kms:ViaService": "ec2.us-west-2.amazonaws.com"
                }
            }
        }
    ]
}
```
> Which of the following options are correct regarding the policy?



Answer

> The first statement provides a specified IAM principal the ability to generate a data key and decrypt that data key from the CMK when necessary

### Question 43

Question

> A startup has been experimenting with DynamoDB in its new test environment. The development team has discovered that some of the write operations have been overwriting existing items that have the specified primary key. This has messed up their data, leading to data discrepancies.\
> Which DynamoDB write option should be selected to prevent this kind of overwriting?

Answer

> Conditional writes

### Question 44

Question

> As an AWS Certified Developer Associate, you have been hired to work with the development team at a company to create a REST API using the serverless architecture.\
> Which of the following solutions will you choose to move the company to the serverless architecture paradigm?

Answer

> API Gateway exposing Lambda Functionality



### Question 45

Question

> A company has a cloud system in AWS with components that send and receive messages using SQS queues. While reviewing the system you see that it processes a lot of information and would like to be aware of any limits of the system.\
> Which of the following represents the maximum number of messages that can be stored in an SQS queue?

Answer

> no limit

### Question 46

Question

> As a Senior Developer, you are tasked with creating several API Gateway powered APIs along with your team of developers. The developers are working on the API in the development environment, but they find the changes made to the APIs are not reflected when the API is called.\
> As a Developer Associate, which of the following solutions would you recommend for this use-case?

Answer

> Redeploy the API to an existing stage or to a new stage

### Question 47

Question

> A company is using a Border Gateway Protocol (BGP) based AWS VPN connection to connect from its on-premises data center to Amazon EC2 instances in the company’s account. The development team can access an EC2 instance in subnet A but is unable to access an EC2 instance in subnet B in the same VPC.\
> Which logs can be used to verify whether the traffic is reaching subnet B?

Answer

> VPC Flow Logs\
> VPC Flow Logs is a feature that enables you to capture information about the IP traffic going to and from network interfaces in your VPC. Flow log data can be published to Amazon CloudWatch Logs or Amazon S3. After you've created a flow log, you can retrieve and view its data in the chosen destination.



### Question 48

Question

> Other than the Resources section, which of the following sections in a Serverless Application Model (SAM) Template is mandatory?

Answer

> Transform

### Question 49

Question

> A company wants to automate its order fulfillment and inventory tracking workflow. Starting from order creation to updating inventory to shipment, the entire process has to be tracked, managed and updated automatically.\
> Which of the following would you recommend as the most optimal solution for this requirement?


Answer

> Use AWS Step Functions to coordinate and manage the components of order management and inventory tracking workflow

### Question 50

Question

> A development team is building a game where players can buy items with virtual coins. For every virtual coin bought by a user, both the players table as well as the items table in DynamodDB need to be updated simultaneously using an all-or-nothing operation.\
> As a developer associate, how will you implement this functionality?


Answer

>  Use `TransactWriteItems` API of DynamoDB Transactions

### Question 51

Question

> You have created a continuous delivery service model with automated steps using AWS CodePipeline. Your pipeline uses your code, maintained in a CodeCommit repository, AWS CodeBuild, and AWS Elastic Beanstalk to automatically deploy your code every time there is a code change. However, the deployment to Elastic Beanstalk is taking a very long time due to resolving dependencies on all of your 100 target EC2 instances.\
> Which of the following actions should you take to improve performance with limited code changes?

Answer

> Bundle the dependencies in the source code during the build stage of CodeBuild



### Question 52

Question

> A company wants to share information with a third party via an HTTP API endpoint managed by the third party. The company has the necessary API key to access the endpoint and the integration of the API key with the company's application code must not impact the application's performance.\
> What is the most secure approach?

Answer

> Keep the API credentials in AWS Secrets Manager and use the credentials to make the API call by fetching the API credentials at runtime by using the AWS SDK

### Question 53

Question

> You are a development team lead setting permissions for other IAM users with limited permissions. On the AWS Management Console, you created a dev group where new developers will be added, and on your workstation, you configured a developer profile. You would like to test that this user cannot terminate instances.\
> Which of the following options would you execute?

Answer

> Use the AWS CLI --dry-run option\
> The --dry-run option checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is DryRunOperation, otherwise, it is UnauthorizedOperation.

### Question 54

Question

> A pharmaceutical company runs their database workloads on Provisioned IOPS SSD (io1) volumes.\
> As a Developer Associate, which of the following options would you identify as an INVALID configuration for io1 EBS volume types?

Answer

> 200 GiB size volume with 15000 IOPS\
> This is an invalid configuration. The maximum ratio of provisioned IOPS to requested volume size (in GiB) is 50:1. So, for a 200 GiB volume size, max IOPS possible is 200*50 = 10000 IOPS.

### Question 55

Question

> A company uses Amazon Simple Email Service (SES) to cost-effectively send susbscription emails to the customers. Intermittently, the SES service throws the error: Throttling – Maximum sending rate exceeded.\
> As a developer associate, which of the following would you recommend to fix this issue?

Answer

> Use Exponential Backoff technique to introduce delay in time before attempting to execute the operation again

### Question 56

Question

> A Developer is configuring Amazon EC2 Auto Scaling group to scale dynamically.\
> Which metric below is NOT part of Target Tracking Scaling Policy?

Answer

> ApproximateNumberOfMessagesVisible

### Question 57

Question

> A junior developer has been asked to configure access to an Amazon EC2 instance hosting a web application. The developer has configured a new security group to permit incoming HTTP traffic from 0.0.0.0/0 and retained any default outbound rules. A custom Network Access Control List (NACL) connected with the instance's subnet is configured to permit incoming HTTP traffic from 0.0.0.0/0 and retained any default outbound rules.\
> Which of the following solutions would you suggest if the EC2 instance needs to accept and respond to requests from the internet?

Answer

> An outbound rule must be added to the Network ACL (NACL) to allow the response to be sent to the client on the ephemeral port range

### Question 58

Question

> A company uses AWS CodeDeploy to deploy applications from GitHub to EC2 instances running Amazon Linux. The deployment process uses a file called appspec.yml for specifying deployment hooks. A final lifecycle event should be specified to verify the deployment success.\
> Which of the following hook events should be used to verify the success of the deployment?

Answer

> ValidateService

### Question 59

Question

> A business has purchased one m4.xlarge Reserved Instance but it has used three m4.xlarge instances concurrently for an hour.\
> As a Developer, explain how the instances are charged?



Answer

> One instance is charged at one hour of Reserved Instance usage and the other two instances are charged at two hours of On-Demand usage

### Question 60

Question

> An application running on EC2 instances processes messages from an SQS queue. However, sometimes the messages are not processed and they end up in errors. These messages need to be isolated for further processing and troubleshooting.\
> Which of the following options will help achieve this?


Answer

> Implement a Dead-Letter Queue

### Question 61

Question

> Recently in your organization, the AWS X-Ray SDK was bundled into each Lambda function to record outgoing calls for tracing purposes. When your team leader goes to the X-Ray service in the AWS Management Console to get an overview of the information collected, they discover that no data is available.\
> What is the most likely reason for this issue?

Answer

> Fix the IAM Role

### Question 62

Question

> A data analytics company is processing real-time Internet-of-Things (IoT) data via Kinesis Producer Library (KPL) and sending the data to a Kinesis Data Streams driven application. The application has halted data processing because of a ProvisionedThroughputExceeded exception.\
> Which of the following actions would help in addressing this issue? (Select two)

Answer

> - Configure the data producer to retry with an exponential backoff
> - Increase the number of shards within your data streams to provide enough capacity

### Question 63

Question

> A university has created a student portal that is accessible through a smartphone app and web application. The smartphone app is available in both Android and IOS and the web application works on most major browsers. Students will be able to do group study online and create forum questions. All changes made via smartphone devices should be available even when offline and should synchronize with other devices.\
> Which of the following AWS services will meet these requirements?

Answer

> Cognito Sync



### Question 64

Question

> As a Team Lead, you are expected to generate a report of the code builds for every week to report internally and to the client. This report consists of the number of code builds performed for a week, the percentage success and failure, and overall time spent on these builds by the team members. You also need to retrieve the CodeBuild logs for failed builds and analyze them in Athena.\
> Which of the following options will help achieve this?

Answer

> Enable S3 and CloudWatch Logs integration 

### Question 65

Question

> A media publishing company is using Amazon EC2 instances for running their business-critical applications. Their IT team is looking at reserving capacity apart from savings plans for the critical instances.\
> As a Developer Associate, which of the following reserved instance types you would select to provide capacity reservations?

Answer

> Zonal Reserved Instances\
> A zonal Reserved Instance provides a capacity reservation in the specified Availability Zone. Capacity Reservations enable you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration. This gives you the ability to create and manage Capacity Reservations independently from the billing discounts offered by Savings Plans or regional Reserved Instances

</details>

I stopped keeping all the questions, It takes away from my learning time.

## Takeaways

1. can EBS volumes be encrypted in-flight? *YES*
2. When is "Pilot-light" Deployment used? *THIS IS DISASTER RECOVERY*
3. Which RDS databases can integrate with IAM roles? *IAM database authentication works with MySQL and PostgreSQL engines for Aurora as well as MySQL, MariaDB and RDS PostgreSQL engines for RDS.*
4. max SQS message size? *256KB*
5. ECS persistent Cross-AZ data volume? *EFS filesystem*
6. CodeBuild encryption options? *CAN USE KMS NATIVELY BUILT IN*
7. CloudFront key pairs *PRIVATE KEY SIGNS, PUBLIC KEY VERIFIES*
8. CloudWatch integration (S3 or Kinesis)? *S3*
9. Which Caching option has the least amount of work involved for DynamoDb? *Dynamo DAX Doesn't require code changes at all*
10. 50X errors? when is bad gateway and when is Service unavailable? *503 when thet are no target, 504 timeout, 502 SSL problems, etc...*
11. EC2 bursts?
12. When to use Kinesis Data Streams and when to use Firehose? *Data streams is real-Time ingesting events, Firehose is near real time and is used to bring data into analyics tools*
13. VPC Gateway endpoints when? *only for DynamoDb and S3*
14. *DLQ for sqs fifo doesn't have to be FIFO itself.*
15. Amazon Inspector vs Trusted Advisor?
    > - Access Advisor feature on IAM console- To help identify the unused roles, IAM reports the last-used timestamp that represents when a role was last used to make an AWS request. Your security team can use this information to identify, analyze, and then confidently remove unused roles. This helps improve the security posture of your AWS environments. Additionally, by removing unused roles, you can simplify your monitoring and auditing efforts by focusing only on roles that are in use.
    > - AWS Trusted Advisor - AWS Trusted Advisor is an online tool that provides you real-time guidance to help you provision your resources following AWS best practices on cost optimization, security, fault tolerance, service limits, and performance improvement.
    > - IAM Access Analyzer - AWS IAM Access Analyzer helps you identify the resources in your organization and accounts, such as Amazon S3 buckets or IAM roles, that are shared with an external entity. This lets you identify unintended access to your resources and data, which is a security risk.
    > - Amazon Inspector - Amazon Inspector is an automated security assessment service that helps improve the security and compliance of applications deployed on AWS. Amazon Inspector automatically assesses applications for exposure, vulnerabilities, and deviations from best practices.
16. CodePipeline or CodeDeploy?
    *CodeDeploy - EC2, Lambda, on-premises. CodePipeline is everything - build, test, deploy*
18. Cross Account X-Ray? *Possible*
19. Cross Region CloudFormation Template? *THOSE ARE STACK-SETS*
20. *No Billing Access by default for users, must be enabled*.
21. Cognito User Pools vs Identity Pools?
    1.  *USER POOLS are custom users*
    2.  *Identity Pools can interact with AWS*
22. how to authenticate against API gateway?
    1.  *cognito user pools*
    2.  *IAM roles/user (sigv4)*
    3.  *lambda authorizer*
23. RCU, WCU calculations
    1. *strong consistent read* $items* \lceil{\frac{itemSize}{4}\rceil}$
    2. *eventually consistent read* $\frac{items}{2}* \lceil{\frac{itemSize}{4}\rceil}$
    3. *writes* $items* \lceil{\frac{itemSize}{1}\rceil}$
24. 1.  MAX iops for gp2?
    1.  100 - 16,000 iops, 33.33 - 5333 GB. burst of 3000 iops
25. Are EBS volumes region locked or AZ locked? *THEY ARE AZ LOCKED - since EC2 machines are AZ scoped*
26. CloudFormation parameter types
    1. String – A literal string
    2. Number – An integer or float
    3. List\<Number> – An array of integers or floats
    4. CommaDelimitedList – An array of literal strings that are separated by commas
    5. AWS::EC2::KeyPair::KeyName – An Amazon EC2 key pair name
    6. AWS::EC2::SecurityGroup::Id – A security group ID
    7. AWS::EC2::Subnet::Id – A subnet ID
    8. AWS::EC2::VPC::Id – A VPC ID
    9. List\<AWS::EC2::VPC::Id> – An array of VPC IDs
    10. List\<AWS::EC2::SecurityGroup::Id> – An array of security group IDs
    11. List\<AWS::EC2::Subnet::Id> – An array of subnet IDs
27. Step function Activities? step function state machine? - *step function activities are for code that's running somewhere else*
28. *S3 buckets have regions, and they have cross region replications*
29. *immutable deployment is for beanstalk, codeDeploy has blue/*green*
30. *S3 access analyzer finds buckets with public access*
31. *no limit on number of environment variables, the total size of all variables can't exceed 4kb*
32. *kinesis data streams are more flexible that firehose and cost less. firehose is easier and scales more, but limited with it's consumers*
33. *detailed monitoring is 1 minute instead of 5, it doesn't change the metrics collected. custom metrics can be standard or high-resolution (1,5,10,30, 60 seconds)*
34. *S3 query string usually means pre-signed URL*
35. *using the S3 bucket requires being able to access the specified key. so KMS access in the IAM*
36. *parameters in the parameter store don't have resource policies, they have parameter policies. secrets in the secret manager have resource policies and other options*
37. *rds backups are in the same region (multi-az), the read-replicas cab be at different regions and be promoted.*
38. *cross zone load balancing is enabled by default for application load balancers. we can turn it off if we have different types of machines per Availability Zone.*
39. *to change how much data we send from x-ray, we enable x-ray sampling*
40. *CodeBuild is serverless, no provisioning machines, not auto-scaling groups*
41. *memcached is the simpler option than redis. redis cluster mode can store much more data as cache, can scale horizontally, support other opertations such as sort and rank*
42. is RAM use one of the standard metrics? - *NO*
43. auditing bucket objects with cloud trail? - *the object owner can see the auditing logs, not the bucket owner*
44. *rollbacks from codeDeploy are new deployments with a new deployment id*
45. *use `AWS STS decode-authorization-message` to decode cli authorization messages. permissions are still needed*
46. *ALB host means the address host (anything before the top level domain), path is anything after the "/" symbol*
48. ALB Target Types - *instance (ec2 instanceId), ip, lambda, cannot target publicly routeable ip address*
49. uploading files to s3 with encryption, which headers when using HTTP and HTTPS?
    1.  *"'x-amz-server-side-encryption': 'aws:kms'" - HTTP SSE-KMS*
    2.  *"'x-amz-server-side-encryption': 'AES256'" - HTTP SSE-S3*
    3.  *don't send any SSE-C (customer provided key) in HTTP - the connections isn't secured*
50. Access Control List - allow responses when behind a load balancer *Add a rule to the Network ACLs to allow outbound traffic on ports 1024 - 65535* - 
51. *redis clusters nodes must be in thh same region, redis imposes some limitations on RDS operations*
52. is there a special SSM audit trail? - *NO! only CloudTrail*
53. when to use the Kinesis produce library and when the kinesis agent? *kinesis agent can monitor some files and sends them to kinesis stream, the KPL writes to the streams, but does not monitor files*
54. what elastic Beanstalk configurations are available?
    1.  *single container - docker image, no ecs, includes a proxy server* 
    2.  *multi-container - ECS with several container* 
    3.  *custom platform - custom behavior, use a custom AMI, etc...*
55. *the `--page-size` option for the cli increases the number of api calls, but less pressure on each of them.*
56. caching strategies
    1.  *lazy loading - when requested, save the data to cache*
    2.  *write through - when updated, save the data to cache, never stale*
57. *cross account x-ray - the IAM role gets the location where to send the data, so if we assume a role from another account, we can send traces there*
58. *API gateway can be authenticated with*
    1.  *custom authorizers lambda*
    2.  *Cognito user pools*
    3.  *IAM sig4*
    4.  *STS is not supported!*
59. *x-ray on ecs - requires the deamon agent as side car or daemonset, and a proper IAM role. no damonsets on fargate*
60. *to direct to a login page on S3 hosted website from cloudFront, we need two cacheing behaviors on the same origin.*
61. *S3 versioning is for the entire bucket, replication can be per-path*
62. *codeBuild is serverless by itself, it can put artifacts and encrypt them with a given KMS key*
63. *api gateway responses*
    1.  *500 - WAF, DNS, no internet access for the VPC*
    2.  *501 - transfer encoding*
    3.  *502 - problem with handshake, TLS, message response, SSL*
    4.  *503 - no registered targets*
    5.  *504 - no subnet permissions, no response*
64. *if we want a stable Ip, we use elasticIP, that's it*
65. *CloudFront origins*
66. *S3 strong consistency?*

