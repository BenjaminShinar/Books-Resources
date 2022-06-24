<!--
// cSpell:ignore xhttp
-->

[main](README.md)

## Section 12 - Serverless

<!-- <details> -->
<summary>
//TODO: add Summary
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

### Serverless Application Model (SAM) [SAA-C02]

### Elastic Container Service (ECS) [SAA-C02]

### Summary

### Quiz 9: Serverless Quiz

</details>

##

[main](README.md)
