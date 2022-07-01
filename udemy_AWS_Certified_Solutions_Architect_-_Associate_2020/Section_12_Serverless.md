<!--
// cSpell:ignore xhttp Xeon
-->

[main](README.md)

## Section 12 - Serverless

<details>
<summary>
Serverless Services on AWS, no machine provisions required
</summary>

A brief history: in the past we had data centers, which was a physical machine that was dedicated you, then amazon launched EC2 (2006), which was a virtual machine which could be provisioned in minutes - infrastructure as a service. then we started with Platforms as a service, and later Containers, and after that we have Serverless - abstracting away the management, leaving only the code.

### Lambda Concepts

Lambda is the ultimate abstraction layer:

- Data Centers
- Hardware
- Assembly Code/Protocols
- High Level languages
- Operating Systems
- Application Layer / AWS APIs
- AWS Lambda

> AWS Lambda is a compute service where you can upload your code and create a lambda function. AWS Lambda takes care of provisioning and managing the servers that you use to run the code. You don't have to worry about operating systems, patching, scaling, etc...
>
> Can used in the following ways:
>
> - As an event-driven compute service where aws Lambda run your code in respose to events. These events could be changes to data in an Amazon S3 bucket or an Amazon DynamoDD table.
> - As a compute service to run your code in response to HTTP requests using Amazon API Gateway or API calls made using AWS SDKs.

lambdas can trigger other lambda, lambdas are run in isolation, it triggers in a different machine each time, so there is no problem of scale.

in traditional architecture, we use load balancers to distribute requests to web servers (virtual machines). in serverless architecture, we don't care about load balancing, lambdas are sprung up immediately.

in serverless architecture, there will always be an API gateway, a lambda, and probably a database that is serverless itself (such as aurora).

Lambda supports many languages:

- node.js
- java
- python
- c#
- Go
- powershell

lambdas are priced by the number of requests, the first million request per month are free, and then pay per million requests. there is also a billing based on duration, rounded up to the nearest 100 ms.

lambdas allows us to abstract away the servers, no more worrying about EC2 machines, patching Operating systems, provisioning them... we only worry about the code. lambda are also very scalable, and they are very cheap compared to virtual machines.

Lambdas are used by amazon Alexa (amazon Echo)

> Exam Tips:
>
> - Lambda scales out (not up) automatically.
> - Lambda functions are independent, 1 event = 1 function.
> - Lambda is serverless.
> - Know what services are serverless.
>   - RDS is not serverless, except aurora.
>   - DynamoDB is serverless
>   - S3 is serverless.
>   - EC2 is not serverless
> - Lambda function can trigger other lambda functions.
> - Architecture can get extremely complicateed, AWS X-ray allows you to debug what is happening.
> - Lambda can do things globally, you can use it to back up S3 buckets to other S3 buckets, etc...
> - Understand different triggers, what can trigger lambdas.

### Let's Build A Serverless Webpage

overview

- route 53
- s3 bucket
  - index.html
  - dynamic contnet
- api gateway
- lambda

we start by creating a lambda, so we click <kbd>Create Function</kbd>, we start from scratch, using the python runtime. we give the lambda a role, and use the template policy of "simple microservice permissions".\ now we can edit the differnet triggers under the <kbd>Designer</kbd> options, and we can modify the function code. we want the lambda to return a response with status code 200 (success).

```py
def lambda_handler(event, context):
    print("In lambda handler")

    resp = {
        "statusCode": 200,
        "headers": {
            "Access-Control-Allow-Origin": "*",
        },
        "body": "Some Text"
    }

    return resp
```

we can also define environment variables, tags and description, we can also limit the memory used, set a timeout, choose the vpc, and some other options.

next we trigger a trigger, we select <kbd>Create a new API</kbd> from the dropdown menu, and <kbd>AWS IAM</kbd> from the security dropdown menu. this creates an API gateway endpoint. we can click the api gateway to view the flow in the gateway, we can set the accepted requests. we click <kbd>Actions</kbd> and choose to allow a **Get** request. with integration Type "Lambda Function", and we fill out the checkbox called "Use Lambda Proxy integration". we select the region and the lambda which we created.\
we click <kbd>Action</kbd>, then <kbd>Deploy</kbd> to deploy the api. we can then perform the action from the api gateway page by clicking the url called "Invoke URL". we copy this address and we will put it in the html file that serves as our website.

```html
<html>
  <script>
    function myFunction() {
      var xhttp = new XMLHttpRequest();
      xhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
          document.getElementById("my-demo").innerHTML = this.responseText;
        }
      };
      xhttp.open("GET", "YOUR-API-GATEWAY-LINK-HERE", true);
      xhttp.send();
    }
  </script>
  <body>
    <div align="center">
      <br /><br /><br /><br />
      <h1>Hello <span id="my-demo">Cloud Gurus!</span></h1>
      <button onclick="myFunction()">Click me</button><br />
      <img
        src="https://s3.amazonaws.com/acloudguru-opsworkslab-donotdelete/ACG_Austin.JPG"
      />
    </div>
  </body>
</html>
```

now we configure the S3 bucket, we can also configure the route53.

under S3, we <kbd>Create New Bucket</kbd>, give it a name (which is in the same domain as our route53). by default, objects in S3 buckets are private, so we change the public access settings allow access.

in this bucket, we configure <kbd>Static Website Hosting</kbd> and enable it, pointing to "index.html" and "errors.html". we upload the html files to this bucket, select them, <kbd>Actions</kbd> and then <kbd>Make Public</kbd>. now we have abutton that we can click and it will change the text.

we can also set up the route53 domain to point to this bucket
under <kbd>hosted zones</kbd>, <kbd>create record Set</kbd>, choose type A record, and point it to the S3 bucket.

### Let's Build An Alexa Skill

building an alexa skill, using the polly service. we will encode mp3 files, store them in a bucket, and then use the alexa service.

Amazon echo isn't the same as alexa, when we talk to alexa, we actually talk to a lambda. we could use other hardware devices to talk to alexa. all of this is done by using developers.amazon.com, we need a different account for that.

> Skill building
>
> - Skill Service: AWS lambda
> - Skill Interface
>   - Invocation Name
>   - Intent Schema
>   - Slot Type
>   - Utterances

in thw aws Console, we create an s3 bucket, we want the bucket to public. with a bucket policy that everything in the bucket is public as well. this bucket will store out mp3 files, we copy the bucket name and ARN.

next, under the **machine learning** category, we choose the _Amazon Polly_ service. we can play with the text-to-speech option. we add some text, and then click <kbd>Synthesize to S3</kbd>, we fill in the form with the bucket name, an optional key prefix, and an sns topic arn for notifications, and then we click <kbd>Synthesize</kbd> to confirm. this creates a task, which we can see in the _S3 synthesis tasks_ view, with the task details and status.

we can see this file in the s3 bucket which we created.

under the <kbd>Lambda</kbd> service, we can create a lambda, it needs to be in a region where the alexa trigger is enabled, such as north-virginia. aws no longer recommends using blueprints for alexa lambda, so we choose <kbd>AWS Serverless Application Repository</kbd>, and then choose **alexa-skills-kit-nodejs-facskill**. we use default values, and <kbd>Deploy</kbd>.\
Once the lambda is deployed, we can open it, and we see that the trigger is a lambda skill, and we can also see the nodejs code. we can add more lines to the data array. we take the arn.

next, we go to developer.amazon.com, choose <kbd>Amazon Alexa</kbd>, then select <kbd>Alexa Skills</kbd>, and <kbd>Create Skill</kbd>. we give it name and choose the language (based on the alexa device which we use), we choose the model to be <kbd>Custom</kbd> and the backend to be <kbd>Self Hosted</kbd>.

there are templates, so we choose the <kbd>Fact Skill</kbd> template, we can change the <kbd>Skill Invocation Name</kbd>, and under the **endpoint** we enter the lambda arn as the default Region, and save. under the _intents_ option, we can modify the <kbd>GetNewFactIntent</kbd> with more innovations (phrases that the user can say). we save and build this model. this will take some time.

under the <kbd>Test</kbd> tab, we can enable testing in development, and use our computer microphone or write the innovation to test the lambda. if we change the data in the lambda code, it will change the facts we get read to us.

next, we go back to the S3 bucket,grab the mp3 object url, and in the lambda function, we edit the code

```js
const data = ['<audio src ="url"/>', "other text"];
```

now the mp3 file which we created is one of the 'facts' which alexa can read to us.

### Serverless Application Model (SAM) [SAA-C02]

Serverless Application Model (SAM) is an open source framework to build serverless applications easily.

> - CloudFormation Extension optimized for serverless applications.
> - Supports anything CloudFromation suports.
> - has new resource types:
>   - functions
>   - APIs
>   - tables
> - Run serverless appellations locally (using docker, great for unit testing).
> - Package and deploy using **CodeDeploy**.

A SAM template is a yaml like file, similar to what cloudFormation uses.

```yaml
# tells CloudFormation this is a SAM template
AWSTemplateForamtVersion: "2010-09-09"
Transform: AWS::serverless-2016-10-31
Description: Hello World SAM Template

# Applies the same properties to all functions
Globals:
  Function:
    Timeout: 3

# Creates a lambda function from local code, also creates an API gateway endpoint, mappings and permissions
Resources:
  HelloWorldFunction:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: hello_world/
      Handler: app.lambda_handler
      Runtime: python3.8
      Events:
        HellowWorld:
          Type: Api
          Properties:
            Path: /hello
            Method: get

# Output relevant information
Outputs:
  HelloWorldApi:
    Description: "API Gateway endpoint URL for Prod stage for Hello World function"
    Value: !Sub "https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prode/hello"
  HelloWorldFunction:
    Description: "Hello World Lambda Function ARN"
    Value: !GetAtt HelloWorldFunction.Arn
  HelloWorldIamRoe:
    Description: "Implicit IAM Role create for Hello World Function"
    Value: !GetAtt HelloWorldFunctionRole.Arn
Outputs:
```

we can intgrate other service as event sources for our lambda, not just an API gateway.

in an EC2 machine with the sam cli

```sh
sam init
# choose quick start, then python
# give a name and choose a template
ll
cd sam-app
ls
cat template.yaml
cd hello_world
cat app.py
# lets see the python app
```

we modify the file a bit, we uncomment some codes, so that we get the ip address of the lambda.

```sh
sam build
sam deploy --guided
# stack name
# aws region
# confirm changes
# allows iam role creation
# confirm that our lambda doesn't have authorization
# save deployment configuration.
```

we can see the api gateway in the aws console, with the resources and the requests, we also see the newly created lambda function.

in our terminal, we test this lambda.

```sh
curl <endpoint url>/Prod/hello
```

### Elastic Container Service (ECS) [SA A-C02]

containers and docker - a bundle of application and dependencies.

> - A **Container** is a package that contains an application, libraries, runtime, and tools required to run it.
> - Run on a container engine like **Docker**.
> - Provides the **isolation** benfits of virtualization with less overhead and faster starts than Virtual Machines.
> - Containerized applications are **Portable** and offer a consistent environment.

ECS:

> - Managed container orchestartion service
> - Create clusters to manage fleets of container deployments.
> - ECS manages EC2 or fargate instances.
> - Schedules containers for optimal placement
> - Defines rules for CPU and memory requirements
> - Monitors resource utilization
> - Deploy, update and roll back.
> - FREE - only pay for the EC2 machines.
> - Integrates with:
>   - VPC, Security groups, EBS Volumes
>   - ELB
>   - CloudTrail and CloudWatch
>
> **ECS Components**
>
> - Cluster: Logical collection of ECS resources, either ECS EC2 instances or Fargate instances.
> - Task Defintion: Defines your application. Similar to a DockerFile but for tunning containers in ECS. Can contain multiple containers.
> - Container Defintion: Inside a task defintion, it defines the individual containers a task uses. Controls CPU and memory allocation and port mappings.
> - Task: Single Running Copy of any containers defined by a task defintion. One work copy of an application (e.g. DB and web containers).
> - Service: Allows task definitions to be scaled by adding tass. Defines minimum and maximum values.
> - Registry: Storage for Container images. used to download images to create containers. examples are Docker Hub or Elastic Container Registry (ECR).
>
> **Fargate**
>
> - Serveless Container Engine.
> - Works with ECS and EKS.
> - Eliminates need to provision and mange servers.
> - Specify and pay for resources per application.
> - Each workload runs in its own kernel
> - Provides isolation and security

we can usually choose fargate over EC2 instance, unless we have a special reason which requires having a virtual machine, such as:

> - compliance requirement
> - require broader customization
> - require access to GPU

there is also EKS - Elastic Kubernetes Service

> - K8s is an open-source software that lets you deploy and manage containerized application at scale.
> - Same toolset for on-premises and cloud.
> - Containers are grouped in **pods**.
> - EKS support EC2 and fargate.

ECR

> - Managed Docker container Registry
> - Store, manage and deploy images.
> - Integrated with ECS and EKS
> - works with on-premises deployment
> - Highly avaliable
> - Integrated with IAM
> - Pat for storage and data transfer

The ECS and ELB (elastic load balancer) can work together

> - Distribute traffic evenly across tasks in your service.
> - Supports
>   - ALB - Application load balancer - layer 7, route HTTP/S traffic.
>   - NLB - Network load balancer - layer 4, route TCP traffic
>   - CLB - Classic load balancer - layer 4, route TCP traffic
> - ALB allows for
>   - Dynamic host port mapping
>   - Path-based routing
>   - Priority rules.

**Security**\
Instance roles and task Roles: the EC2 instance role applies a policy to all task running on that EC2 instance. a task role applies a policy per task. we should use the "least privilege approach" and use task roles to give permissins per task, rather than put them on the EC2 instance role.

**DEMO**\

in the AWS console. we choose the <kbd>ECS</kbd> service, we can see a diagram with main components of ECS; container definition inside task defintion inside a service inside a cluster.

we can use some pre-built options, such as "sample-app" for apache, we can use the default Task defintion as well, same as with the service (how many tasks, whether or not to use a load balancer), and for the cluster configurations, we give it a name, a VPC and subnets.

it first creates the ECS resources: cluster, task definition and service. After that, it creates the other AWS resources: log group, CloudFormation Stack, VPC, subnets and security groups. if we have a service, we can stop the task and a after a while we will get a new task. we can also update the service to control how many tasks we will have.

each task has a different public IP, which is why we would want a load balancer.

### Summary

**Traditional vs Serverless architecture**

> - Traditional
>   - Elastic Load Balancer
>   - EC2 instance
>   - Database
>   - Auto Scaling group
>   - Can be highly available - if we make it so
> - Serverless
>   - API Gateway
>   - Lambda Function
>   - Serverless Database (dynamo, aurora)

**Lambda**

> - Lambda scales out (not up) automatically
> - Lambda functions are independent - 1 event = 1 function
> - Lambda is serverless
> - Other serverless services
>   - RDS is not, only Aurora
> - Lambda function can trigger other lambda function, so 1 event can trigger many lambda function if they are configured to trigger each other.
> - Architectures can get extremely complicated, AWS X-ray allows you debug what is happening.
> - Lambda functions can do things globally, you can use it to back up S3 buckets to other S3 bucket, or to run jobs on a schedule.
> - Understand the differnet triggers, what can and what can't trigger lambdas.

### Quiz 9: Serverless Quiz

- "As a DevOps engineer you are told to prepare complete solution to run a piece of code that required multi-threaded processing. The code has been running on an old custom-built server based around a 4 core Intel Xeon processor. Which of these best describe the AWS compute services that could be used?" _ANSWER: (EC2,ECS, Lambda) The exact ratio of cores to memory has varied over time for Lambda instances, however Lambda like EC2 and ECS supports hyper-threading on one or more virtual CPUs (if your code supports hyper-threading)._
</details>

##

[main](README.md)
