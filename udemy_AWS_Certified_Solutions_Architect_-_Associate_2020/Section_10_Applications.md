<!--
// cSpell:ignore
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

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

### Reducing Security Threats [SAA-C02]

### Key Management Service (KMS) [SAA-C02]

### CloudHSM [SAA-C02]

### Parameter Store [SAA-C02]

</details>

##

[next](Section_12_Serverless.md)\
[main](README.md)
