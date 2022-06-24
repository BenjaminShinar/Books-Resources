<!--
// cSpell:ignore fileb boto
-->

[main](README.md)

## Section 10 - Applications

<details>
<summary>
Amazon Services which are standalone applications.
</summary>

### SQS [SAA-C02] - Simple Queue Service

SQS - Simple Queue Service.

one of the first AWS services

> Amazon SQS is a web service that gives you access to a message queue that can be used to store messages while waiting for a computer to process them.\
> I't a distributed queue system that enables web service applications to quickly and reliably queue messages that one componenet in the application generates to be consumed by another component.\
> A queue is a temporary repository for messages that are awaiting processing.

example: meme generator website. we start by uploading a photo to S3, which triggers an lambda, which takes the image and grabs a text from the message in a queue. then an EC2 machine combines the text and the image and uploads the result to an S3 bucket. if the EC2 machine fails, the message remains in the queue.

another example, travel website. a user enters the information at the website, which is pushed to a SQS queue, which is pulled by EC2 machines that search for flights in those dates. the data is then pulled and returned to the website and shown to the user.

> Using Amazon SQS you can decouple the components of an application so they run independently, easing message management between components. Any component of a distributed application can store messages in a fail-safe queue.\
> Messages can contain up to 256 kb of text in any format, and any componenet can later retrieve the messages programmatically using the amazon SQS Api.

_(the key word is **decouple**, look for it in the exam)_

we can get larger messages, but then they are stored in s3, not sqs.

> The queue acts as a buffer between the component producing and saving data, and the componenet receiving the data for processing.\
> This means the queue resolves issues that arise if the producer is producing work faster than the consumer can process it, or if the the producer or consuer are only intermittently connected to the network.

two types of queue:

- standards queue: default, nearly unlimited transactions, each message is guranteed to be delivered at least once. and the orders might be delevered out of order.
- FIFO queue - first in, first out. the order is preserved, and a message is delivered exactly once. messages are available until a consumer processes it and deletes it, no duplicates.

  - support message groups - multiple ordered message groups within a single queue.
  - limited to 300 transaction per second (TPS).

Messages have a **visibility timeout**, which stars when a reader picks up the message, and lasts until the reader deletes it (completes the processing) or a period of time. if the period of time passes then the message become visible again and can be picked by another consumer, this is how we might result in processing the same message twice.

Short polling checks for messages and returns immediately with a result (even if the queue is empty). Long Polling returns only if there is a message, or until a timeout. we might decide our consumer should perform long polling if we know our queue will be mostly empty.

Exam tips

- SQS is pull-based, not pushed base.
- Messages are 256K in size
- Messages can be kept in the queue from 1 minute to 14 days, the default retention period is 4 days.
- Visibility period timeout - length of time that is provided to a consumer to process the message before another consumer picks up. the maximum is 12 hours.
- SQS guarantees that your message will be processed at least once.
- Short and Long polling. short - immediate return, long - wait and timeout.
- The term **Decoupling** is usually used in the context of SQS

### SWF - Simple Work Flow Service

mostly comes up in comparison to sqs.

> Amazon Simple Workflow service (SWF) is a web service that makes it east yo coordinate work across distributed application components. SWF enable applications for a range of use cases, including media processing, web application back-end,, business process workflows, and analtics pipelines, to be designed as coordination of tasks.\
> tasks represent invocation of various processing steps in an application which can be performed by executable code, web service calls, human actions and scripts.

it's a way of combining steps between different components, some of which can be manual tasks, like getting a delivery, which has a human element to it.

the terminology used is about tasks

| Operation        | SQS                                                            | SWF                                            |
| ---------------- | -------------------------------------------------------------- | ---------------------------------------------- |
| retention period | max 14 days                                                    | up to a year                                   |
| API              | Message oriented API                                           | Task oriented API                              |
| Duplications     | Duplicate messages are possbile, up to the user to handle them | Task is assigned once and never has duplicates |
| Tracking         | User implemented (if needed)                                   | Keeps track of all the task and events         |

> SWF Actors:
>
> - Workflow Starters - an application that can initiate (start) a workflow. (e-commerce website, mobile app).
> - Deciders - control the flow of activity tasks in a workflow execution. If something has finished (or failed) in a workflow, a Decider decides what to do next.
> - Activity Workers - Carry our the activity tasks.

### SNS - Simple Notification Service

> Amazon Simple Notification Service (SNS) is a web service that makes it easy to set up, operate and send notifications from the cloud.\
> It provides developes with a highly scalable, flexible and cost-effective capability to publish messages from an application and immediately deliver them to subscribers or other applications.

it allows push notifications to phone (apple, google, windows), it also provides notifications by SMS messages, emails. amazon SQS messages or any other http endpoint.

> SNS allows you to group multiple recipients using **topics**. A topic is an "access point" for allowing recipients to dynamically subscribe for identical copies of the same notification.\
> One topic can support deliveries to multiple endpoint types, for example, you can group together IOS, Android and SMS recipients. When you publish once to a topic, SNS delvers appropriately formatted copies of your message to each subscriber.

all messages are stored redundantly in multiple availability zones.

> SNS Benefits
>
> - instantaneous, push based delivery (no polling)
> - Simple APIS and easy intgreation with applications
> - Flexible messages delivery over multiple transport protocols.
> - Inexpensive, pay-as-you-go model with no up-front costs.
> - Web bases AWS Management Console offers the simplicity of a point and click-interface.

**SQS is pull based, while SNS is push based.**

### Elastic Transcoder

a media transcorer in the cloud. convert media filers from the original source format into different formats that will play on different devices (smartphone, tablets, PCs). There are presets for popular formats transcoding, so theres no need to guess which format works best on which device.

payment is based on computation time and resolution.

in example of using this is storing a video on S3, which triggers a lambda which sends the data to elastic transcoder, and the output is stored at a different S3 bucket after it's been converted to all the supported formats.

### API Gateway

> Amazon API Gateway is a fully managed service that makes it easy for developers to publish, maintain, monitor and secure APIs at any scale.

it's something that acts as a "front door" for applications to access data or functionality from our backend services (or other sources). it distributes traffic across targets, it's usually used to connects with lambda, but can also work with dynamoDB or EC2 machine.

API gateway can

> - Expose HTTPS endpoints to define a RESTful Api
> - Serverless-ly connect to services like Lambda and DynamoDb
> - Send each API endpoint to a different target
> - Run efficiently with low cost
> - Scales effortlessly
> - Track and control usage by API key
> - Throttle requests to prevent attacks
> - Connect to CloudWatch to log all requests for monitoring
> - Maintain multiple versions of your API

configuration:

1. we define an api (container
2. define resource and nested resources (URL paths)
3. for each resource
   1. we select supported HTTP methods (verbs: Get, Put, Post, Delete)
   2. set security
   3. choose target (Ec2, lambda, DynamoDb,etc...)
   4. set request and response transformations

we deploy at gateway to a stage it has a default gateway domain but can use custom domain, and now supports QS certificate Manager (SSL/TLS certs).

> Cacheing:\
> You can enable API caching in Amazon API gateway to cache your endpoint's response. With caching you can reduce the number of calls made to your endpoint and also improve the latency of the requests to your API.\
>  When you enable caching for a stage, API gateway caches responses from your end point for a specified time-to-live (TTL) period (in seconds). API gateway then responds to the requests by looking up the endpoint response from the chache instead of making a request to your endpoint.
>
> Same Origin Policy:\
> In computing, the same-origin policy is an important concept in the web application security model. under the policy, a web browser permists scripts contained in a first web page to access data in a second web page, but only if both web pages have the same origin (domain name).\
> This is done to prevent Cross-Site Scripting (XSS) attacks.
>
> - enforced by web browsers
> - some tools ignore it (postman, curl)

in aws,we always have different domain names, like S3, lambda. so we use CORS

> CORS is one way the server at the other end (not the client code in the browser) an relax the same-origin policy.\
> Cross origin resource sharing (CORS) is a mechanism that allows restricted resources (e.g. fonts) on a web page to be requested from another domain outside the domain from which the resource was served.

if we see the error message "origin policy cannot be read at the remote resource" we need to enable CORS.

Exam Tips

> - API gateway at a high level - entry point to connect to other targets (Lambda, EC2, dynamoDB)
> - Caching (TTL) to increase performance.
> - Low cost and scales automatically
> - You can throttle API gateway to precent attack
> - You can log results to cloud watch
> - If you are using javascript/AJAX that uses multiple domains with API gateway, ensure that you have enables CORS on API gateway.
> - CORS in enforced by the client.

### Kinesis [SAA-C02]

> Amazon Kinesis is a platform to send your streaming data to. Kinesis makes it easy to load and analyze streaming data, also providing the ability for you to build your own custom applications four your business needs.

> Streaming Data is data that is generated continiously by thousands of data sources, which typically send data record simultaneously, and in small sizes (oreder of kilobyes).

this can be purchases from online stores, stock prices, game date (as the gamer plays, what happens), social network data (twitter) or geospatial data (like uber.com) and IOT sensor data (internet of things). data that is generated all the time, in small sizes.

3 types of kinesis

- kinesis streams
- kinesis firehose
- kinesis analysis

Kinesis Streams

data producers stream the data to kinesis.

- data is stored in _shards_.
- default storage period is 24 hours, but can be configured to seven days.
- Consumers (EC2 machine) use the data in the shards to perform some analytics, and then they can store the result of of the work.

Shards:

| Action                  | Read           | Write                           |
| ----------------------- | -------------- | ------------------------------- |
| limits                  | 5 transactions | 1000 records                    |
| maximum size per second | 2MB            | 1 MB (including partition keys) |

the data capcity of the stream is a function of the number shards, the total capacity is the sum of the shards.

Kinesis Firehose:\
no persistent storage, data is analyzed as it comes, triggers a lambda to run on the data.

Kinesis analytics works with either kinesis streams or kinesis firehose. it analyzes the data and stores a result.

> Exam Tips
>
> - Know the difference between Kinesis Streams and Kinesis Firehose. choose the relvent one to the given scenario.
> - Understand what kinesis Analytics is.

### Web Identity Federation - Cognito

> Web Identity Federation lets you give your users access to AWS resources after they have successfully authenticated with a web-based identity provider like amazon, facebook or google.\
> Following successfull authentication, the user recives an authentication code from the WEB ID provider, which they can trade temporary for AWS security credentials.

Amazon Cognito is a Web Identity Federation service

- Sign-up and sign-in to your apps
- Access for guest users
- Acts as an identity broker between your application and Web ID providers, so you don't need to write any additional code.
- Synchronizes user data for multiple devices
- Recommended for all mobile application AWS services

a user logs in with facebook and this gives them an authentication token, which they send to cognito and get back application keys for aws services.

> Cognito brokers between the app and (facebook or google) to provide temporary credential which maps to an IAM role allowing access to the required resource.\
> No need for the application to embed or store AWS credentials locally on the device, so it gives users a seamless experience across all mobile devices.

User Pools and identity Pools:

> Cognito User pools are user directories used to manage sign-up and sign-in functionalities for mobile and we applications. Users can sign-in directly to the user poo, or by using Facebook, Amazon or Google.\
> Coginto acts as an identity broker between the identity provider and AWS. successfull authentication generates ad JSON web Token (JWT).
>
> Identity Pools provide temporary AWS credentials to access AWS services like S3 or DynamoDB.

example: a user logs in using facebook, which passes a token to the aws UserPool, and a JWT token is granted. with this token, AWS IdentityPool is accessed, and the user is given AWS credentials, which can be used to access aws resources.

> Cognito tracks the association between user identity and the various different devices they sign-in from.\
> In order to provide a seamless user experience for your application, Cognito uses Push Synchronization to push updates and synchronize user data across multiple devices. Cognito uses SNS to send a notification to all the devices associated with a given user identity whenever stored in the cloud changes.

it means that if a user changes something in one device, cognito sends a silent notification to all devices connected with to this user and updates the data there.

Exam Tips

- **Federation** allows users to authenticate with a web Identity Provider (google, facebook, amazon).
- The user authenticates first with the Web ID Provider and revices an authentication token, which is exchanged for a temporary AWS credentials allowing them to assume an IAM role.
- **Coginto** is an Identity Broker which handles interaction between tour applications the web ID provider (you don't need to write you own code to do this).
- User Pool is user based. it handels things like user registration, authenticaton and account recovery.
- Identity Pool authorize access to AWS resource - IAM Role.

### Summary

SQS - Simple Queue Service

- SQS is a way to de-couple your infrastructure
- SQS is pull based, not push based
- messages in SQS are 256 Kb
- retention time is from one minute to 14 days, with the default being 4 days
- Standard SQS and FIFO SQS
  - Standard does not guarantee order and messages can be delivered more than once.
  - FIFO maintans order and messages are delivered only once.
- Visibility timeout - the time where the message is reserved for a reader that took it until it's given to another worker, if the timeout is too short, it may cause duplicate processing.
  - max visibility timeout is 12 hours
- SQS guarantees that messages will be processed at least once.
- Long polling (waits for a message or a timeout) and Short polling(return immediately)

SWF vs SQS

- retnetion period of 14 days vs 1 year (SWF)
- SWF uses task oriented API, SQS uses message oriented terminology
- SWF ensures tasks are processed once and never duplicated
- SWF keeps track of all events, for SQS this needs to be implemented by the user.
- SWF Actors
  - Starters - initiate a workflow
  - Deciders - deicide on the control flow
  - Workers - carry out activities

SNS

- Push based delivery (instantaneous, no polling)
- Simple API, easy integration
- Flexible, multiple transport protocols
- Inexpensive, pay as you go
- Web based management with point-and-click interface

Elastic transcoder is a media transcoder in the cloud.

API Gateway

- gateway to aws resources
- caching capabilities
- scales automatically
- can be throttled
- results can be logged to cloud watch
- enable CORS if using AJAX

Kinessis:

- kinesis streams - data persistent, using shards.
- kinesis firehose - no persistency, consumed immediately
- kinsis analytics - analyze data

Cognito and Federation

- authenticate with Web identity provider
- A token is exchanged for temporary AWS credentials that allow the user to assume an IAM role.
- Cognito is an identity Broker which handels interaction between applications and the Web ID provider.
- user pool - user based: registration, authentication, account recovery
- identity pool - authorize access to IAM Roles and AWS resources.

### Quiz 8: Applications Quiz

> - What happens when you create a topic on Amazon SNS? _ANSWER: An amazon resource name is created._
> - In SWF, what does a "domain" refer to? _ANSWER: A collection of related workers_
> - What does Amazon SES stand for? _ANSWER: Simple Email Service_

</details>

## Section 11 - Security

<details>
<summary>
Different ways to reduce security risks.
</summary>

### Reducing Security Threats [SAA-C02]

there are some **Bad Actors**, clients (malicious software) that try to access the service in a way that they shouldn't. they might try to steal or scrape information, pretned to be someone else by sending fake data, and they can even try to shut down the service by flooding it with request (Denial of service).

if we know the ip of the bad actor, we can use Network Access Control Lists and create an Inbound rule to deny access to that ip. we can also use host-based firewall services.

if we have an application load balancer, we can put all the security operation on it, and terminate the connection at the load balancer level. this is one reason to have different security groups, one for the load balancer, and another for the EC2 service (which can only be accessed from the load balancer).

when using Network Load Balancers,the traffic passes through it to the EC2 service, and the ip is visible across the entire connection, so it should be terminated at the EC2 level.

A WAF (Web Application Firewall) can be attached to the load balancer and it monitors (blocks and filters) requests. it has presets for common attacks such as SQL injection.

A WAF uses layer 7 - so it is applications aware, when we wish to block ip addresses ranges, we use level 4 protection, meaning NACL.

we can also attach a WAF to a cloudFront edge, this is helpful because if we have a cloudfront connection, then the load balancer sees the ip from it, and not that of the original client.

### Key Management Service (KMS) [SAA-C02]

> - **Regional** secure key management, encryption and decryption.
> - Manages **Customer Master Keys** (CMKs)
> - Ideal for S3 object, database passwords and API keys stored in System Manager Parameter Store.
> - Encrypt and decrypt data up to **4 kb** in size.
> - Integrated with most AWS services.
> - Pay per API call.
> - Audit capability using ClouTrail - logs delivered to S3.
> - FIPS 140-2 Level - a type of security standard.
>   - Level 3 is supported in CloudHSM.

three types of CMKS:

> 1. Customer Managed CMK: allows key _rotation_. controlled via key policies and can be enabled/disabled.
> 1. AWS Managed CMK: free. used by deafult if you pick encryption in most AWS service. only that service can use them directly.
> 1. AWS owned CMK: used by AWS on a shared basis accross many accounts. uncommon.

| Type             | User View | User Manage | Dedicated To Account |
| ---------------- | --------- | ----------- | -------------------- |
| Customer Managed | Yes       | Yes         | Yes                  |
| AWS Managed CMK  | Yes       | No          | Yes                  |
| AWS Owned CMK    | No        | Yes         | No                   |

keys can be symmetric or asymmetric.

Symmetric:

- same key used for encrption an decryption.
- AES-256.
- Never leaves AWS un-encrypted.
- Must call the KMS APIs to use.
- AWS services which are integrated with KMS use symmetric CMKs.
- Encrypt, decrypt and re-encrypt data.
- Generate data keys, date key pairs and random byte strings.
- **Import** your own key material.

Asymmetric:

- Mathemetically related public/private key pair (ssh uses this).
  - the public key can be given to anyone.
- **RSA** and **Elliptic-Curve Cryptography** (ECC).
- Private key never leaves AWS un-encrypted.
- Must call the KMS APIs to use the private key.
- Download the public key and use outside AWS.
- Used outside AWS by users who can't call KMS APIs.
- AWS services which are integrated with KMS **do not support** asymmetric CMKs.
- used to Sign messages and verify signatures.

Default Key Policy - grant AWS account (root user) _full access_ to the CMK.

```json
{
  "Sid": "Enable IAM User Permissions",
  "Effect": "Allow",
  "Principal": {
    "AWS": "arn:aws:iam::111122223333:root"
  },
  "Action": "kms:*",
  "Resource": "*"
}
```

we can also give permissions to a specific role.

```json
{
  "Sid": "Allow use of the key",
  "Effect": "Allow",
  "Principal": {
    "AWS": "arn:aws:iam::111122223333:role/EncryptionApp"
  },
  "Action": [
    "kms:DescribeKey",
    "kms:GenerateDataKey",
    "kms:Encrypt",
    "kms:ReEncrypt",
    "kms:Decrypt"
  ],
  "Resource": "*"
}
```

CMKS are region specific, so if we want to move data between regions, it must be decrypted, moved, and enctypted again with a different key.

in the console, we select the **KMS** service, and then choose the <kbd>AWS Managed Keys</kbd> option, when we integrate a service with KMS, a key with an alias is created automatically.\
under <kbd>Customer Managed Keys</kbd>, we can see the keys which we created. we can rotate the key automatically.

we create an EC2 machine and give it the Encryption Role. we can also create an alias, which points to the key, this facilitates key rotation, as it allows changing keys without changing the code itself (we simply point the alias to a different key). the name of the alias must be prefixed with "alias/".

```sh
aws kms create-key --description "Demo CMK"
# we take the key id
aws kms create-alias --target-key-id <keyId> --alias-name "alias/demoKey"
aws kms list-keys
# using the key
echo "this is a secret message" > topsecret.txt
cat topsecret.txt
# now output a base64 blob
aws kms encrypt --key-id "alias/demoKey" --plaintext file://topsecret.txt --output text --query CiphertextBlob
# decode and push to a file
aws kms encrypt --key-id "alias/demoKey" --plaintext file://topsecret.txt --output text --query CiphertextBlob | base64 --decode > topsecret.txt.encrypted
cat topsecret.txt.encrypted
# decrypt, fileb is for binary
aws kms decrypt --ciphertext-blob fileb://topsecret.txt.encrypted --output text --query Plaintext | base64 --decode
```

we don't have to specify which key to use for decryption, as this data is part of the encryption metadata.

for files above 4 KB, we can use a Data Encryption Key (DEK). it uses envelope encryption. this allows us to reduce the amount the date we use.

```sh
aws kms generate-data-key --key-id "alias/demoKey" --key-spec AES_256
# we take the ciphertextBlob
```

### CloudHSM - Hardware Security Module [SAA-C02]

HSM Hardware Security Module. a way to manage private keys, validated control on the key.

> - **Dedicated** hardware security module.
> - **FIPS 140-2 Level 3**.
> - stronger than KMS level2.
> - Manage your own Keys.
> - **No access** to the AWS-managed component.
> - Runs within a VPC in your Account
> - Single tenant, multi-AZ, running on a cluster dedicated to the user.
> - Industry standard APIs- no **AWS APIs**.
> - _PKCS#11_, _Java Cryptography Extensions (JCE)_, > _Microsoft CryptoNG (CNG)_
> - If the key is lost, it's irretrievable (even by AWS), so keys must be kept safe.

the cluster runs in an existing or a new VPC, this cluster project ENI (elastic network interface) for the other vpc to communicate with. HSM isn't redundant by default, so we need to create HSMs in multiple AZ.

### Parameter Store [SAA-C02]

a way to manage secret and configuration in aws services. passwords, connection strings, access keys, and so on.

> - Component of the AWS System Manager (SSM)
> - Secure, **serverless** storage for configuration and secrets:
>   - Passwords
>   - Database connection strings
>   - License codes
>   - API keys
> - Values can be stored encrypted (KMS) or in plaintext
> - Allows for separation of data from source control.
> - Store parameters in **hierarchies**
> - Track versions of values
> - Set TTL to expire values such as passwords

parameters can be stored in hierarchies. so data can be restored from the root or from a leaf. this also allows to restrict access based on hierarchies or path.

to get all the parameters, we call the `GetParametersByPath` api and provide the path.

Parameter Store is integrated with CloudFormation

```yaml
Parmeters:
  LatestAmiId:
    Type: "AWS::SSM::Parameter::Value<AWS::EC2::Image:Id>"
    Default: "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"

Resources:
  Instance:
    Type: "AWS::EC2::Instance"
    Properties:
      ImageId: !Ref LatestAmiId
```

in the console

we will create a lambda function to access the Parameter Store, but we first need a role to access it.

under IAM, <kbd>policies</kbd>, we <kbd>Create Policy</kbd> and use this policy, we give it a name, like "lambda_parameter_store_policy".

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "ssm:GetParameter*",
        "ssm:GetParametersByPath*"
      ],
      "Resource": "*"
    }
  ]
}
```

now we create a role to use this policy, so <kbd>Role</kbd>, <kbd>Create Role</kbd>, and select a common use case "Lambda" and search for this policy, attach it, and give the role a name "lambda_parameter_store_role".

next we create the lambda itself. Under **Lambda**, <kbd>Create Function</kbd>, provide a name, choose a runtime (such as python), and give the lambda the role.

now we replace the boilerplate code with the following. we create an aws client (boto3).

```py
import json
import os

import boto3

client = boto3.client("ssm")
env = os.environ["ENV"]
app_config_path = os.environ["APP_CONFIG_PATH"]
full_config_path = "/" + env + "/" + app_config_path

def lambda_handler(event, context):
  print("Config Path: "+ full_config_path)
  param_details = client.get_parameters_by_path(Path=full_config_path, Recursive=True, WithDecryption=True)
  print(json.dumps(param_details, default=str))
```

we also need to set the environment variables for the lambda, so we click <kbd>Edit Environment variables</kbd>: "ENV":"prod", "APP_CONFIG_PATH":"acg". we will need them in the parameter store. all of our parameters will be under the path "prod/acg".

we also want a key, so in the **KMS** service, we select <kbd>customer managed keys</kbd>, then <kbd>Create a key</kbd>, a symmetric key, give it a name. we need to define the key administrator, which a user/role that owns the key. we also set up who can access the key.

under the **SSM** service, we select <kbd>Parameter store</kbd> and start creating parametesr, there are two tiers, depending on the number of parameters which we have (we choose standard key). we also select <kbd>SecureString</kbd>, and choose the KMS key we created. now we can start entering the value. the value can be up to 4096. we can create another parameters, this time a <kbd>StringList</kbd>, which isn't encrypted.

to see the lambda in action, we configure a test event, with an event name, we don't care about the values. we click <kbd>Test</kbd> and then we se the log with the json containing all the parameters which we created.

</details>

##

[next](Section_12_Serverless.md)\
[main](README.md)
