<!--
// cSpell:ignore crond
 -->

[main](README.md)

## Section 9 - High Avalability Architecture

<details>
<summary>
Designing For Failure.
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

we always plan for failure, everything fails, no matter what, even the entire data center can fail. so we should always be ready for failures of our machines and services.

Netflix actually has a system to introduce failure into their system to test how it handles it. they might shutdown an ec2 machine.

one example of high avaliability is having failover between regions and between availability zones in the same region.

- Always design for failure.
- Use multiple AZ's and multiple Region whereever you can.
- Know the difference between multi-AZ and Read Replicas for RDS.
  - Multi AZ- disaster recovery.
  - Read Replicas - performance.
- Know the difference between sacling out and scaling up.
  - scaling out: more ec2 machine, using auto scaling group.
  - scaling up: increasing the power of the machine.
- Read the question carefully and always consider the cost element.
- Know the different S3 storage classes.

### HA WordPress Site Demo

<details>
<summary>
Fault tolerant wordpress site
</summary>

we will have the user connect a **Route53** domain, connect into an **ELB**, we will have **EC2** machine in an **autoscaling group**, in different Avalability zones, also two **RDS** in a different **security group**. in addition, we will have two **S3 buckets**, one with the code for the site, and one with the media, and we will serve this through **CloudFront**.\
we will then create failure in different points.

#### Set up

we start with creating the S3 buckets. we have two buckets in the same region.

now we create a CloudFront distribution, and <kbd>Create Distribution</kbd>, we set the origin domain name to the s3-media bucket, and the rest is default.

we next want to create **security groups** in the VPC service, we create a new security group for port 3306 that is open for the Web DMZ sg.

we want a **RDS** database, we use the dev/test template, we also check the <kbd>Multi-AZ deployment</kbd> option, and we set the SG to the newly created one. we don't open it to the public. we also need to give the Database a name. the rest is deafult.

in **IAM**, we make sure we have a role that can connect to S3 bucket, so we create a new role with S3 full access permissions.

in **EC2**, we provision an ec2 machine, give it the IAM role which we created, and we pass in a script that installs apache for webserver, and gets the files required for wordpres. also does some permissions, and we also need something called _htaccess_.\
we use default storage, and we choose the WebDMZ security group.

#### Setting Up EC2

We ssh into the ec2 instance we created, using the public ip

```sh
ssh ec2-user@34.254.227.110 -i myKP.pem
#in the ec2 machine
sudo su
cd /var/www/html
ls
cat .htaccess

service httpd start #this should be in the bootstrap script
service httpd status
```

we now navigate to the same public ip with our browser, and we see the wordpress installation screen. we start by filling up the fields:

- database name: same as what we set up in RDS
- username: same as what we set up in RDS
- password: same as what we set up in RDS
- database host: the **rds** endpoint.
- table prefix:

we also need to create a certain file in the EC2 machine, so we copy the text and paste it in.

```sh
nano wp-config.php
```

then we click <kbd>Run the installation</kbd> in the browser, if it hangs, it might mean we haven't set up the security group properly. we now finish up the installation and set up the user name and password.

inside the wordpress platfrom, we can create a post. we want to add an image, so we upload some images and publish the post.

back in the ec2 terminal

```sh
ls
cd wp-content
ls # see years
cd 2019 #the current year
ls # see months
cd 02
ls # see images
```

now we want to make all files that are uploaded to be copied to S3, and eventually, we want to serve the image from cloud front.

```sh
# make sure the role is working properly
aws s3 ls # list buckets
aws s3 cp --recursive /var/www/html/wp-content/uploads s3://<media-cloud-name> # copy media to s3
aws s3 cp --recursive /var/www/html s3://<code-cloud-name> # copy all website to to s3
aws s3 ls s3://<code-cloud-name> #check that it was copied
ct healthy.html #the health check html file
cat .htaccess
```

this is a url rewrite rule, so it can serve the images from somewhere else, in this case, the cloud front distribution.

we take the domain name from the cloudFront Service

```
Options +FollowSymlinks
RewriteEngine On
rewriterule ^wp-content/uploads/(.*)$  http://<domain>.cloudfront.net/$1 [r=301, nc]

# BEGIN WordPress
# END WordPress
```

now we sync ourselves with s3, and update the s3 with the modifed file.

```sh
aws s3 sync /var/www/html s3://<>
```

we now need to tell apache to allow URL rewrites

```sh
cd /etc/httpd/conf
ls
cp httpd.conf httpd-backup.conf # create backup
nano httpd.conf
# change AllowOverride from 'NONE' to 'ALL'

service httpd restart #just to make sure
```

now we need to make the s3 bucket public, so in the <kbd>Permissions</kbd> tab, we paste the policy into the <kbd>Bucket Policy</kbd>. we just need to replace the resource with the current bucket ARN.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicReadGetObject",
      "Effect": "Allow",
      "Principal": "*",
      "Action": ["s3:GetObject"],
      "Resource": ["arn:aws:s3:::BUCKET_NAME/*"]
    }
  ]
}
```

now the bucket is public, if we have an error, we might need to set the option differently.

if we look at the site and take the address of the image, we see the image is served from cloud front, rather than the static website.

next stage is to create the **Elastic Load Balancer** (application load balancer), and move the EC2 behind it. <kbd>Create Load Balancer</kbd>, make it "internet facing", put it inside the webDMZ Sg, set up the Target group, and edit the health check. we register our EC2 to it.

we now go into **Route53** and point a domain name into the Elastic Node Balancer. <kbd>Create Record Set</kbd>, the type will be A (for address), we will use an alias, and the <kbd>Alias Target</kbd> is the ELB.

in **EC2**, we place the instance into the security group,under <kbd>Target</kbd>,<kbd>Add To Register</kbd>, and now we can use the dns address to navigate to the wordpress website.

#### Adding Resilience And Autoscaling

in the ec2 instance, we add a command to the crontab to scan the s3 site and always be synched with it.

```sh
cd /etc
nano crontab
#echo "*/1 * * * * root aws s3 sync --delete s3://<bucket name> /var/www/html" >> crontab
service crond restart #run all the commands
```

to test this, we can add a file to s3 and then check if it's in the ec2

```sh
cd /var/www/html
ls
```

now we want to make this ec2 machine an AMI, so our changes will be part of a template that we will use if the machine fails and needs to restart.

so in the EC2 Service, we choose <kbd>Instances</kbd>, <kbd>Actions</kbd>, <kbd>Image</kbd> and then <kbd>Create Image</kbd>.

we set the fields and call it a "default ec2 wordpress reader machine". now this is an AMI which we can use.

now we create a writer ami.

```sh
sudo su
cd etc
echo "*/1 *  * * * root aws s3 sync --delete /var/www/html s3://<code-bucket-name>" >> crontab
echo "*/1 *  * * * root aws s3 sync --delete /var/www/html/wp-content/uploads s3://<media-bucket-name>" >> crontab
cat crontab # check that it worked
cd /var/www/html
echo "This is a Test" > text.txt
service httpd status # check that the service is on.
```

we can test and see that the bucket are getting updates and the bucket has the new file. we might have issues of eventual consistency.

now this instance is the 'writer' node, it is what we use to update the buckets, which the other instance use to serve the website.

so next step is to create the auto scaling group with the "reader" instances. under EC2: <kbd>Auto Scaling</kbd>, <kbd>Auto scaling Group</kbd>, <kbd>Create Auto Scaling Group</kbd>, <kbd>Launch Configuration</kbd>, <kbd>new launch configuration</kbd> and we select the AMI which we created. we set the machine type, give the IAM permissions, and provide a bootsrap script

```sh
#!/bin/bash
yum update -y
aws s3 sync --delete s3://<code-bucket-name> /var/www/html
```

we use the webDMZ security group.

for the auto scaling group, we set a name, the group size to start with, and we check the <kbd>Load Balancing</kbd> box to receive traffic from the load balancer, we use the target group we created, we set the health check to <kbd>ELB</kbd> type. under the subnet section, we give all the Avalability zones. the rest is default.

for the scaling group, we can configure a scaling policy (based on some metric) or keep the group at the initial size.

now we remove the writer Node from the Target Group and deregister it from the target group, so it won't receive traffic from the LoadBalancer. we could also see now the new instances being created.

in the browser, we can log into the admin panel in the writer node and add a post with an image. we might have some wait time until the image is propagated to cloud front, so we won't see it right away.

now we have all the pieces needed, so we can start testing our site and see it it's really highly available.

we can start by going to the EC2 service, <kbd>Instances</kbd>, select one of the reader nodes, and then <kbd>Actions</kbd>, <kbd>Terminate</kbd>, so we crush one availability zone. now under <kbd>Target Group</kbd> we can see that one instance was lost and that health checks are failing.

under <kbd>Auto Scaling</kbd>, we can see that a new EC2 machine is being launched. and it will be added to the target group.

#### Cleaning Up

now we want to failover from one AZ to another with RDS. in RDS, we can select the database, <kbd>Actions</kbd>, <kbd>Reboot</kbd>, select <kbd>reboot With failover</kbd>. this will cause a failover, and the website will be down for a short while.

now we can start deleing all of our assets:

- auto scaling group
- application load balancer (remove registered target)
- target group
- remove the write node
- (remove the AMI we created)
- s3 buckets
- cloud front distributions (disable then delete)

</details>

### CloudFormation

with cloud formation, we can automate the same steps,
under <kbd>Management and Governance</kbd> services, we select **CloudFormation**, <kbd>Create Stack</kbd>.

we will use a sample template, and we choose the "WordPress blog" template. we can view the template by pressing <kbd>View in Designer</kbd>. we edit the template with the required fields, names, passwords, and in the ssh location we use the default `0.0.0.0/0` cidr. this will take few minutes.

**(this might not work because of the PHP version, but it will probably be fixed in the future)**

we can see an ec2 instance was created for our wordpress machine, and we can delete the stack.

there are many templates in aws [quick start](https://aws.amazon.com/quickstart/) page which can be used.

> - CloudFormation is a way of completely scripting your cloud environment.
> - Quick Start is a bunch of cloud Formation Templates already built by AWS Solutions Architects allowing you to create complex environment very quickly.

### Elastic Beanstalk

Elastic Beanstalk is aimed at developers without extensive knowledge of AWS services, it's less complicated than cloud formation.

under <kbd>Compute</kbd> services, we select **Elastic Beanstalk**, click <kbd>Get Started</kbd> and we choose pre-configured platform, such as PHP. we can see what resources are provisoned. we can see the revcent events, and we can then change configurations (adding load balancers, change capacity) to customize our environment.

> With **Elastic Beanstalk**, you can quickly deploy and manage application in the AWS Cloud without worrying about the infrastructure that runs those applications.\
> You simply upload your application, and elastic beanstalk automatically handles the details of capcity provisioning, load balancing, scaling, and application health monitoring.

### Highly Available Bastions

Bastions are a way to connect into a private subnet by establishing a host inside a public subnet.

if we want to make them highly available, we need some changes.

Scenario 1: more expensive, use in production environment

- EC2 instance in a private subnet
- 2 EC2 Bastion instances in public subnets
- 2 AZ
- network load balancer
- we can add an auto scaling group

scenario 2: cheapr, use for development

- EC2 instance in a private subnet
- an EC2 instances in public subnet
- 2 AZ
- auto scaling group that recreates the bastion ec2, and we take the Elastic IP address.

> High Avalability with Bastion Hosts
>
> - Two hosts in two separate avalability zones. Use a network Load balance with static IP address and health check to failover from one host to the other.
>   - Can't use an application load balancer, as it is layer 7 and you need to use layer 4 for ssh.
> - One host in one avalability zone behind an Auto Scaling group with health checks and a fixed EIP. if the host fails the health check will fail and the auto scaling group will provision a new EC2 instances in a separate AZ. You can use a user data script to provision the same EIP to the new host. This is the cheapest option, but it is not 100% fault tolerant.

### On Premise Strategies

High level AWS services that can be used on-premises

> - Database Migration service (DMS)
>
>   - allows you to move databases to and fromAWS
>   - might have your DR (disaster recovery) environment in AWS and you on-premises environment as you primary.
>   - Works with most popular database technologies, such as Orcale, MySQL, DynamoDB, etc...
>     - supports homogenous migration (orcale to orcale)
>     - supports heterogenous migration (SQL Server to Amazon Aurora)
>
> - Server Migration service (SMS)
>
>   - Support incremental replication your on-premises server in to AWS.
>   - Can be used as a backup tol, multi-site strategy (on-premises and off-premises), and as a DR tool.
>
> - AWS Application Discovery Service
>
>   - AWS application discovery Service helps enterprises customers plan migration projects by gathering information about their on premises data centers.
>   - You can install the AWS application Discovery Agentless Connector as a virtual appliance on VMware vCenter.
>   - It will then build a server utilization map and dependency map of you on-premises environment.
>   - The collected data is retained in encrypted format in an AWS application Discovery Service data store. you can export this data as a csv file and use it to estimate the Total Cost of Ownership (TCO) of running on AWS and to plan you migration to AWS.
>   - This data is also avalabile in AWS migration Hub, were you can migrate the discovered servers and track their progress as they get migrated to AWS.
>
> - VM Import/Export
>
>   - Migrate existing application in to EC2.
>   - Can be used to create a DR strategy on AWS or use AWS as a second site.
>   - Can also be used to export your AWS Vms to your on-premises data center.
>
> - Download Amazon Linux 2 as an ISO
>   - Works with all major virtualization providers, such as Vmware, Hyper-V, Vm, Virtual box

### Summary

> Load Balancers
>
> - 3 Different types of load balancers
>   - Application Load Balancers - layer 7 aware
>   - Network Load Balancers - layer 4 aware
>   - Classic Load balancers
> - 504 Error mean the gateway has timed out. This means that the application did not respond within the idle timeout period.
> - if you need the IPv4 address of your end user, look for the **X-Forwarded-For** header.
> - Instances monitored by ELV are repores as either _InService_ or _OutOfService_.
> - Health Checks check the instance health by talking to it.
> - Load Balancers have their own DNS name. You are never given an IP address.
>   - we can get an IP address for a network Load Balancer.
> - Sticky Session enable your users to stick to the same EC2 instance. Can be useful if you are storing information locally to that instance.
> - Cross Zone Load Balancing enables tou to load balance across multiple avaliability zones.
> - Path patterns allow you to direct traffic to different EC2 instance based on the URL contained in the request.
>
> Cloud Formation
>
> - A way of completely scripting your cloud Environment.
> - Quick Start is a bunch of CloudFormation tTemplates already built by AWS allowing you to create complex environments very quickly.
>
> Elastic Beanstalk
>
> - With Elastic Beanstalk, you can quickly deploy and manage application in the AWS Cloud without worrying about the infrastructure that runs those applications.
> - You simply upload your application, and elastic beanstalk automatically handles the details of capcity provisioning, load balancing, scaling, and application health monitoring.
>
> Highly Available Bastions
>
> - Using network Load Balancer with a static IP address and two AZ.
> - Using auto scaling group and one AZ. using elastic IP address.
>
> On Premises AWS services
>
> - Database Migration Service
> - Server Migration Service
> - AWs Application Discovery Service
> - VM Import/Export
> - Download Amazon Linux 2 as an ISO

### Quiz 7: HA Architecture Quiz

> - "You have a website with three distinct services, each hosted by different web server autoscaling groups. Which AWS service should you use?" _ANSWER: ALB. The ALB has functionality to distinguish traffic for different targets (mysite.co/accounts vs. mysite.co/sales vs. mysite.co/support) and distribute traffic based on rules for target group, condition, and priority._
> - "You manage a high-performance site that collects scientific data using a bespoke protocol over TCP port 1414. The data comes in at high speed and is distributed to an autoscaling group of EC2 compute services spread over three AZs. Which type of AWS load balancer would best meet this requirement?" _ANSWER: Network Load Balancer. The Network Load Balancer is specifically designed for high performance traffic that is not conventional web traffic. The Classic LB might also do the job, but would not offer the same performance._
> - "You have been tasked with creating a resilient website for your company. You create the Classic Load Balancer with a standard health check, a Route 53 alias pointing at the ELB, and a launch configuration based on a reliable Linux AMI. You have also checked all the security groups, NACLs, routes, gateways and NATs. You run the first test and cannot reach your web servers via the ELB or directly. What might be wrong?" _ANSWER: In a question like this you need to evaluate if all the necessary services are in place. The glaring omission is that you have not built an autoscaling group to invoke the launch configuration you specified. The instance count and health check depend on instances being created by the autoscaling group. Finally, key pairs have no relevance to services running on the instance._
> - "You work for a manufacturing company that operate a hybrid infrastructure with systems located both in a local data center and in AWS, connected via AWS Direct Connect. Currently, all on-premise servers are backed up to a local NAS, but your CTO wants you to decide on the best way to store copies of these backups in AWS. He has asked you to propose a solution which will provide access to the files within milliseconds should they be needed, but at the same time minimizes cost. As these files will be copies of backups stored on-premise, availability is not as critical as durability. Choose the best option from the following which meets the brief." _ANSWER: S3 OneZone-IA provides on-line access to files, while offering the same 11 9's of durability as all other storage classes. The trade-off is in the availability - 99.5% as opposed to 99.9%-99.99%. However in this brief as cost is more important than availability, S3 OneZone-IA is the logical choice . RRS is deprecated and new uses are strongly discouraged by AWS._
> - "You need to use an object-based storage solution to store your critical, non-replaceable data in a cost-effective way. This data will be frequently updated and will need some form of version control enabled on it. Which S3 storage solution should you use?" _ANSWER: The key point in the questions is that the data is non-replaceable and is frequently updated. The 1st excludes anything the has reduced durability, the second excluded anything with long recall, reduced availability, or billing based on infrequent access._
> - "In S3 the durability of my files is **\_\_\_\_**." _ANSWER: 99.99999999%._
> - "Placement groups can either be of the type 'cluster', 'spread', or 'partition'. Choose options from below which are only specific to Spread Placement Groups." _ANSWER: "There is only one answer that is specific to Spread Placement Groups, and that is the final option. Whilst some of these answers are correct for either Cluster Placement Groups only, or for both Cluster and Spread Placement Groups, the question stated that only options specific to Spread Placement Groups should be chosen. This would rule out two options as they are true for both Spread & Cluster type placement groups. The Logical grouping of instances within a single Availability Zone is only true of Cluster Placement Groups and is also incorrect."_

</details>

##

[next](Section_10_Applications.md)\
[main](README.md)
