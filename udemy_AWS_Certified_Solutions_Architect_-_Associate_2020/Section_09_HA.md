<!--
// cSpell:ignore
 -->

[main](README.md)

## Section 9 - High Avalability Architecture

<details>
<summary>

</summary>

### Load Balancers Theory

Elastic load Balancers

a physical or virtual device designed to help balance the load across different devices and applications in the network.

aws provides three types of load balancers.

> - **Application Load Balancer**\
>   They are best suited for load balancing of HTTP and HTTPS traffic. they operate at Layer 7 and are _application aware_. They are intelligent, and you can create advanced requests routing, sending specified requests to specific web servers.
> - **Network Load Balancer**\
>   They are best suited for load balancing of TCP traffic where extreme performance is required. they operate at the connection level (Layer 4), network load balancers are capable of handling millions of requests per second, while maintaining ultra-low latencies.
>
> - **Classic Load Balancer**\
>    Those are the legacy Elastic Load Balancers. you can balance HTTP/HTTPS application and use Layer 7 specific features, such as X-forewarded and stick sessions. you can also use strict Layer 4 load balancing for application that rely purely on the TCP protocol.\
>   The Classic Load balancer isn't application aware.

when the application stops responding, the Classic ELB will respond with a 504 (gateway timeout) error, this could be from the web server layer or the database layer, so we should identify the failing point, and scale up or out if needed.

X-Forwarded-For Header, if we use the ELB, the source id at the target device will be that of the ELB, but if we want to know the enduser ipv4 address, it can be stored at the _X-forwarded-for_ field in the header.

#### Load Balancers And Health Checks Lab

in the aws console, we launch two EC2 instances, in two different avalability zones. we also pass them a bootstrap script to make them both webservers. we put them into different subnets and make sure they have a public Ip

```sh
#!/bin/bash
yum update -y
yum install httpd -y
service httpd start
chkconfig httpd on
cd /var/www/html
echo "<html><h1>This is WebServer 01</h1></html>" > index.html
```

next we create the load balancer, under <kbd>Load Balancers</kbd>, we create a **classic load balancer**, give it a name, choose the VPC, and use the default port for http 80.

we press next until we reach the _Configure Health Check_ step, we select the response timeout and the interval. we can also modify how many results are required to determine if a webserver is health or not.

we next add the instances to the ELB, and we can <kbd>Enable Cross-Zone Load balancing</kbd> and <kbd>Enable Connection Draining</kbd>.

(note: load balancers aren't part of the free tier)

when the classic ELB is created, we get a DNS name, but not an ip address. we wait for the instances to be in service (pass the health checks). now if we go to the dns address, we will get one of the web servers, if we refresh the page we might get the other instance.

we can stop one of our EC2 instances, and after the health check fails, the ELB will only direct us to the web server that is operational.

now we will create another load balancers, this time into target groups, we will <kbd>Create target group</kbd>, a target group is where the load balancers routes requests and performs health checks. we give it a name, a type (we choose <kbd>Instance</kbd>), protocol, port and VPC. we modify the health check path again, we can change the health check successfull result as well. now we click <kbd>Edit</kbd> to add instances to the target group.\
next we <kbd>Create Load Balancer</kbd> of the application load balancer type, we decide the avalability zones and security group, and use the target group which we set up before. we don't have registered targets for now. but we choose the target group, <kbd>edit</kbd> and add them as registed target.

under the <kbd>listener</kbd> tab of the Application load balancers, we can set rules for how the requests are forwarded.

### Advanced Load Balancer Theory [SAA-CO2]

> Sticky Sessions\
> Classic load Balancers routes each request independently to the registered EC2 instance with the smallest load. Sticky Sessions allow you to bind a users's session o a specific EC2 instance. This ensures that all requests from the user during the session are sent to the same instance.\
> You can enable Sticky Sessions for application load balancers as well, but the traffic will be sent at the Target Group level.

we will use them if we have data written on the ec2 machine, and we want to always have the user at the same EC2 instance.

Cross Zone Load Balancing - allows us to have load balancers in different Avaliability zones, so they can share ec2 instances.

> Path Patterns\
> You can create a listener with rules to forward request based on the URL path. this is known as path-based routing. if you are running microservices, you can route traffic to multiple back-end services using path-based routing. For example, you can route general requests to one target group and request to render images to another target group.
> (this required Application load balancer)

### Autoscaling Theory [SAA-C02]

Auto Sacling has three components:

1. Groups: Logical components. Webserver group, application group, database group, etc...
2. Configuration Templates: Groups uses a launch template or a launch configuration as a configuration template for its EC2 instances. you can specify information such as the AMI ID, instance type, key pair, security groups and block device mapping for your instances.
3. Scaling Options: proving several ways for you to scale your Auto Scaling groups. for example, you can configure a group to scale on the occurrences of specified conditions o(dynamic scaling) or on a schedule.

Sacling options

- maintain current instance levels at all times
- scale manually
- scale based on a schedule
- scale based on demand
- use Predictive scaling

when we choose to maintain instance amount, aws will perform health checks on instances, and if one is down, it will launch another.

when we scale manually, we specify the minimum, maximum or desired capacity of the scaling group, and amazon maintains the level.

if we know about the peak hours of our application, we can use scaling by schedule, and bring up more instances during the rush hours, and then take them down.

scaling based on demand defines proprties and thresholds, and once the threshold is passed, additional instances are sprung up or brought down. an example of a property can be CPU utilization. this is the most popular type of scaling, when demnad is high, we respond to the raise in requests.

predictive scaling is a proactive approach, based on the past demand to the service.

#### Autoscaling Groups Lab

in the AWS console, we choose<EC2> and then <kbd>Load Balancers</kbd>, we delete our old load balancers and target groups, and we will want an auto scaling group.\
for that, we first need to <kbd>Create launch configuration</kbd>, select the AMI and just like EC2 instances, but now we are creating a launch configurations, we will use the same bootstrap script, choose the security group, all as usual.

however, this doesn't create any EC2 instance, so we now <kbd>Create Auto Scaling Group</kbd>, we give it a name, the newly created launch configuration, and set the group size (number of instances), as well as which vpc and subnet will be used. we can also choose an elastic load balancer, and at the next step we can start managing the scaling process.

we select the minimum and maximum, and which metric to use. we can also get a notification when things are changed. we can terminate instances so the health check fails, and then we will see that the scaling groups starts running new instances until we get back to 3 instances. when we are done we can click <kbd>Action</kbd>, <kbd>Delete</kbd> and this will also delete the EC2 instances.

### HA Architecture

### HA Word Press Site

### Setting Up EC2

### Adding Resilience And Autoscaling

### Cleaning Up

### CloudFormation

### Elastic Beanstalk

### Highly Available Bastions [SAA-C02]

### On Premise Strategies [SAA-C02]

### Summary

### Quiz 7: HA Architecture Quiz

</details>

##

[next](Section_10_Applications.md)\
[main](README.md)
