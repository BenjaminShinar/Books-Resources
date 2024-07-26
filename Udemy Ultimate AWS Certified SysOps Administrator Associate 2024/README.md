<!--
// cSpell:ignore proto deregisteration_delay
 -->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# Ultimate AWS Certified SysOps Administrator Associate 2024

<!-- <details> -->
<summary>
Practice towards AWS Certified SysOps Administrator Associate exam.
</summary>

udemy course [Ultimate AWS Certified SysOps Administrator Associate 2024](ultimate-aws-certified-sysops-administrator-associate). by [_Stephane Maarek_](https://www.stephanemaarek.com/).

## EC2 for SysOps

<details>
<summary>
Stuff about EC2 machines.
</summary>

<cloud>EC2</cloud> - <cloud>Elastic Cloud Compute</cloud> - virtual machines.

### EC2 Instance Types
<details>

#### Launching an EC2 Instance

from the console, launch a new machine, with a ssh key pai, create new security group allowing SSH access (port 22) from anywhere, and ssh into the machine using the public ip.

```sh
chmod 0400 DemoKeyPair.pem
ssh -i DemoKeyPair.pem ec2-user@<public_ip>
```

or by using <kbd>EC2 Instance Connect</kbd>.

#### Changing EC2 Instance Type

changing EC2 instance types works only for <cloud>EBS</cloud>-backed instances.

1. Stop the instance
2. <kbd>Instance Settings</kbd> -> <kbd>Change Instance Type</kbd>
3. Start instance

for some instances, we can have EBS-optimization available.

we can see this in action by running the `free -m` shell command before and after changing the machine type. going from "t2.micro" to "t2.small" changes the amount of "total" memory from ~983 to ~1991.

</details>

### Enhanced Networking

<details>
<summary>
Get better networking performance for EC2 instances.
</summary>


> <cloud>EC2 Enhanced Networking (SR-IOV)</cloud>
> - Higher bandwidth, higher PPS (packet per second), lower latency.
> - Option 1: Elastic Network Adapter (ENA) up to 100 Gbps.
> - Option 2: Intel 82599 VF up to 10 Gbps - LEGACY.
> - Works for newer generation EC2 Instances.
> 
><cloud>Elastic Fabric Adapter (EFA)</cloud>
>
> - Improved ENA for <cloud>HPC</cloud> (high performance computing), only works for Linux.
> - Great for inter-node communications, tightly coupled workloads.
> - Leverages Message Passing Interface (MPI) standard.
> - Bypasses the underlying Linux OS to provide low-latency, reliable transport.

as a demo, we create a new machine from the "t3.micro" type, and we can check if the <cloud>ENA</cloud> model is installed by running `modinfo ena` and `ethtool -i eth0` (replace `eth` with `enX0` or `ens5` in amazon linux 2023 AMI) to see the driver used - it will show "ena" when the ENA is used or "vif" when not (like in the t2 machine).

</details>


### EC2 Placement Groups


<details>
<summary>
Controling how the EC2 instances are placed in the data center.
</summary>

> - *Cluster* - clusters instances into a low-latency group in a single Availability Zone.
> - *Spread* - spreads instances across underlying hardware (max 7 instances per group per AZ) - critical applications.
> - *Partition* - spreads instances across many different partitions (which rely on different sets of racks) within an AZ. Scales to 100s of EC2 instances per group (Hadoop, Cassandra, Kafka)

In the *cluster* placement group, we stick all the instances in the same AZ, which gives us better networking, and we can use enhanced networking. the risk is that if the Availability Zone fails, all the instances fail. we use this strategy when we have big jobs that need to complete fast, or when our applications need extremely low latency and high network throughput.\
For the *spread* placement groups, we span across multiple Availability Zone and hardware, this reduces the risk of failure (better Availability). however, we ate limited to having 7 instances per Availability Zone per placement group.
The *Partition* group uses partitions to separate machines. instances in a partition don't share physical racks with instances from other partitions, so the partition represents a rack inside the data center. EC2 instances have the partiton information as part of the metadata. this is usually used in big-data application which are partition-aware

under the <cloud>EC2</cloud> service, we can choose the<kbd>Placement Group</kbd> tab and <kbd>Create Placement Group</kbd>, we give it a name, tags, and select the placement strategy. there are some options for the "spread" strategy (racks and outhosts) and the "partition" (number of partitions).\
If we launch an instance, we look under the "advanced details" configuration and select which placement group to use.

</details>

### EC2 Shutdown Behavior & Termination Protection
<details>
<summary>
Shutdown from inside the machine.
</summary>

> How should the instance react when shutdown is done using the OS?
>
> - Stop (default)
> - Terminate

we can set the shutdown behavior to either stop the instance or terminate. this will still terminate the instance even if we had "terminate protection" available (it only effects termination from the console or the cli, not from inside the instance OS).

</details>

### EC2 TroubleShooting

<details>
<summary>
Common troubleShooting methods.
</summary>

#### Troubleshooting EC2 Launch Issues

possible reasons that we fail to launch <cloud>EC2</cloud> instances:

> `#InstanceLimitExceeded` if you get this error, it means that you have reached your limit of max number of vCPUs per region.

this applies to "On-demand" and "Spot" instances. we can fix this by moving to another region, shutting down other machines, or requesting a quota Increase from AWS. we can see the limit in the <cloud>Service Quotes</cloud> page.

> `#InsufficientInstanceCapacity` if you get this error, it means AWS does not have that enough On-Demand capacity in the particular AZ where the instance is launched.

for this error, there isn't much we can do, we can wait or change the instance type or the Availability Zone.

> Instance Terminates Immediately: (goes from pending to terminated)
>
> - You've reached your EBS volume limit.
> - An EBS snapshot is corrupt.
> - The root EBS volume is encrypted and you do not have permissions to access the KMS key for decryption.
> - The instance store-backed AMI that you used to launch the instance is missing a required part (an image.part.xx file).

#### Troubleshooting EC2 SSH Issues

Common issues with SSH access:

1. "Unprotected private key file" error - the ".pem" key needs to have the permissions set to 400 (use `chmod 0400`).
2. incorrect username - "Host key not found", "permission denied", "Connection closed by \<instance\> port 22". - can happen when we use `ec2-user` for ubuntu AMIs, and vice versa.
3. "Connection timed out" - can be from the security group, the Network Access Control List, the route table or if the instance doesn't have a public IPv4 address or the instance CPU load is too high.

EC2 Instance Connect is an alternative to SSH connection. rather than having the security group permitting an inbound rule access from a source, we have it allow a specific range of ips which are reserved for instance connect.

we can see the ranges in this [address](https://ip-ranges.amazonaws.com/ip-ranges.json):

```json
{
    "ip_prefix": "18.206.107.24/29",
    "region": "us-east-1",
    "service": "EC2_INSTANCE_CONNECT",
    "network_border_group": "us-east-1"
}
```

so when we set the security group, we need to add the cidr range of the region.

instance connect creates a one-time ssh public key for each connection.

</details>

### EC2 Instance Purchasing Options

<details>
<summary>
How much we pay for our EC2 instances.
</summary>

> - On-Demand Instances - short workload, predictable pricing, pay by second.
> - Reserved (1 & 3 years)
>   - Reserved Instances - long workloads.
>   - Convertible Reserved Instances - long workloads with flexible instances.
> - Savings Plans (1 & 3 years) - commitment to an amount of usage, long workload.
> - Spot Instances - short workloads, cheap, can lose instances (less reliable).
> - Dedicated Hosts - book an entire physical server, control instance placement.
> - Dedicated Instances - no other customers will share your hardware.
> - Capacity Reservations - reserve capacity in a specific AZ for any duration.

On-Demand instances have pay-by-second for Linux and Windows machines, but pay-by-hour for other OS.\
Reserved instances can be bought and sold in the Reserved Instance Marketplace, convertable instances allows us to change the type and still get a discount.\
Savings plans apply to usage, and are locked to instance family (e.g. M5, T3) per region, but allow for flexibility in instance size (micro, large, 2xLarge), the OS and the tenancy (host, dedicated, default).\
Spot instances can be the cheapest, but you can "lose" them at any time so they should be used for workloads that can run at anytime and stop and start.\
Dedicated Hosts gives us a physical server, this is important when we have compliance requirements and we need to use existing software licensees which are bound to the hardware itself. this can be on-demand or reserved. this is the most expensive option in AWS.\
Dedicate instances - no other accounts can run on the same hardware as these EC2 instances. but it can share the hardware with instances from the same account, and we have no control over instance placement.\
Capacity reservations - always have capacity available, but no discounts, always pay. for critical workloads that must run at a specific Availability Zone. should be combined with some other saving plan to get billing discounts.

#### Spot Instances & Spot Fleet

this is an auction - we define a max spot price which are willing to pay, and as long as the current spot price is below that price then we get the instance. if the price raises above what we are willing to pay, our instances will either stop or terminate (within a two minutes grace period). we can stop if we care about the state of the machine, or terminate if we can restart from an empty state at any time.

there is an old option of *Spot Blocks*, which allowed to get spot instances that won't be interrupted once they started (it wasn't 100% guranteed either).

to start using a spot instance, we need a **spot request**.

- maximum price
- desired number of instances
- launch specifications
- request type: One-time or Persistent
- Valid from, Valid until

we can only cancel requests that are either in the 'open', 'active' or 'disabled' state. canceling a spot request won't terminate the instances that it's running. to terminate them we need to first cancel the request and then manually terminate them.

A Spot Fleet is a way to have a set of Spot Instances with additional (optional) On-Demand Instances. ir will launch from possible launch pools (based on instance type, OS, Availability Zone), and the fleet will choose from the possible launch pools until reaching cpacity or maximum cost.

Allocation strategies
- lowestPrice - from the lowest priced pool
- diversifed - distributed across all pools
- capacityOptimized - pool with the optimal cpacity for the number of instances.
- priceCapacityOptimized (recommended) - first the pool with the highest cpacity available. then select the lowest priced pool.

#### EC2 Instances Launch Types Hands On

in the <cloud>EC2</cloud> service page, we select the <kbd>Spot Requests</kbd> tab, and we can view how the price changed over time. we can click <kbd>Request Spot Instances</kbd>, choose the launch template or define one manually.\
Then we define the request details, setting the maximal price, the valid from and valid until options, use load balancers, decide if we terminate the instances when the request expires. we set the target cpacities (number of instaces, vCPU, memory), networking, matching possible instances, and select the allocation strategy.

we can also launch a spot instance directly, when we launch an EC2 machine we can ask to use a spot instance instead, and then we set the termination behavior to control how the instance will behave if it gets reclaimed.

#### Burstable Instances

> AWS has the concept of burstable instances (T2/T3 machines). Burst means that overall, the instance has OK CPU performance.
>
> - When the machine needs to process something unexpected (a spike in load for example), it can burst, and CPU can be VERY good.
> - If the machine bursts, it utilizes "burst credits"
> - If all the credits are gone, the CPU becomes BAD
> - If the machine stops bursting, credits are  accumulated over time
> - Burstable instances can be amazing to handle unexpected traffic and getting the insurance that it will be handled correctly.
> - If your instance consistently runs low on credit, you need to move to a different kind of non-burstable instance
> 
> T2, T3 "Unlimited" - It is possible to have an "unlimited burst credit balance"
> 
> - You pay extra money if you go over your credit balance, but you don't lose in performance
> - If average CPU usage over a 24-hour period exceeds the baseline, the instance is billed for additional usage per vCPU/hour
> - Be careful, costs could go high if you're not monitoring the CPU health of your instances

</details>

### Elastic IPs

<details>
<summary>
Fixed Ips
</summary>

each time we start and stop an EC2 instance, the public IP changes. if we want it to have a fixed public IP address, then we need to use an Elastic IP.

it's an IPv4 address, which we own and remains ours until we delete it. we can remap the address as we wish. we pay for the address if it's not being used (mapped to a server), but if we attach it to some server, then it doesn't have additional charges.\
Elastic IPs can be used to mask failures of instances by rolling over to another instance. there is a limit of 5 elastic IPs per account - they shouldn't be used much.

- use a random ip address and register a DNS name to it.
- use a load balancer with a static hostname.

(demo of creating an elastic IP and attaching it to one instance, detaching it and then attaching it to another one)

</details>

### CloudWatch Metrics for EC2

<details>
<summary>
EC2 machine monitoring.
</summary>

we have metrics that AWS automatically pushes for EC2, either at 5 minutes interval or 1 minute (extra pay). we can also define custom metrics at 1 minute resolution, all the way to 1 second.

> AWS Provided metrics (AWS pushes them):
>
> - Basic Monitoring (default): metrics are collected at a 5 minute internal
> - Detailed Monitoring (paid): metrics are collected at a 1 minute interval
> - Includes:
>   - CPU - utilization, credit usage and balance
>   - Network - in/out
>   - Status Checks 
>       - Instance statues for the EC2 VM
>       - System status for the underlying hardware
>   - Disk (instance store only) - read/write for ops/bytes. (doesn't apply to <cloud>EBS</cloud>-based instances).
> - **Doesn't Include RAM**
> 
> Custom metric (yours to push):
> - Basic Resolution: 1 minute resolution
> - High Resolution: all the way to 1 second resolution
> - Include RAM, application level metrics
> - Make sure the IAM permissions on the EC2 instance role are correct!

for each EC2 machine, we click the <kbd>Monitoring</kbd> tab and see the metrics, we can also add them to <cloud>CloudWatch</cloud> dashboards for a centralized location. we can enabled detailed monitoring for the machine if we want, but it will cost us more.

#### Unified CloudWatch Agent 

This agent allows us to gather metrics from EC2 instances or from other servers (such as on-premises machines). this can collect system level metrics (RAM, processes, used disk space), and can send the to <cloud>AWS CloudWatch</cloud>.

we can configure the agent with <cloud>SSM parameter store</cloud> or a configuration file, the machine needs to have the correct <cloud>IAM role</cloud> and permissions. the metrics from the unified agent all have the default "CWAgent" namespace. there is a *procstat* plugin that can collect metrics for individual processes running on the machine. we can select which processes are monitored with the pid identification number or based on regex (process name, command line which started it). mertics from the plugin will have the "procstat" prefix.

(demo of installing the Unified CloudWatch Agent)

we first create an IAM role for the EC2, and give it the "CloudWatchAgenetServer" policy, which allows is to push log events and read fom the SSM store.\
We next create an EC2 instance, we will make it a webserver (so it will need HTTP access). we connect to it with the EC2 Instance Connect.

```sh
sudo su # elevate permissions
yum install httpd # install webserver
echo "hello world" > /var/www/html/index.html
sudo systemctl start httpd
sudo systemctl enable httpd
# use browser to access the server and see the hello world message
cat /var/log/httpd/access_log # see logs
cat /var/log/httpd/error_log # see errors
```

if we want to install the agent, we can run the following commands. when we run the configuration wizard once, we can also tell the agent to monitor other files (like httpd/access_logs). after we finish we get the option to store it in the <cloud>SSM parameter Store</cloud>. but for this we need even more permissions, since we didn't allow the newly created role to push to the parameter store (the admin policy has this permission).

when we next create other instances we can have them run the command to grab the configuration directly from SSM.

```sh
#!/bin/bash
# install the agent on Amazon Linux 2
sudo yum install amazon-cloudwatch-agent

# run the wizard
sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-config-wizard
# Ec2, root, turn on statsD daemon, port, interval, aggregates, collectD, host metrics, metrics at code levels, Ec2 dimensions, aggregations, resolutions, default metrics config level
# which files we monitor /var/log/httpd/access_logs, retention rate, other log files
# store in ssm - we need permission

# get the configuration from parameter store 
sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl -a fetch-config -m ec2 -c ssm:AmazonCloudWatch-linux -s

# get the configuration from local file
sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl -a fetch-config -m ec2 -c file:/opt/aws/amazon-cloudwatch-agent/bin/config.json -s
```
now our agents push the logs from inside the machine directly to cloudWatch, and we can see the metrics being sent from the cloudWatch Agent.


#### EC2 Instance Status Checks


Automatic Checks to Identify hardware and sotware issues.

- System Status Checks - problems with AWS system. we might need to restart the machine (move it to another host) to fix the issue.
- Instance Status Checks - invalid network configuration, exhausted memory, we usually reboot the machine to solve

we have cloudWatch metrics for these failures

- StatusCheckFailed_System
- StatusCheckFailed_Instance
- StatusCheckFailed (both)

We can set an cloudWatch alarm to recover the instance (and send a notification) when we get the alarms. this will mantain the private/public and elstic IP addresses, and use the same metadata and placement group.\
Another option is to use an autoscaling group and have it target the health check, and it will recover the instance, but without keeping the same private and elastic Ip.

we can the status checks for each machine going to the <kbd>Status Checks</kbd> tab. and we can also set an alarm based on the metrics (statusCheckFailed) and decide which action should be triggered, and we can test it by forcibly moving the alarm state.

```sh
aws cloudwatch set-alarm-state --alarm-name <alarm-name> --state-vale ALARM --state-reason "testing recovery action"
```

> SYSTEM status checks - System status checks monitor the AWS systems on which your instance runs.\
> Problem with the underlying host. Example: 
> 
> - Loss of network connectivity
> - Loss of system power
> - Software issues on the physical host
> - Hardware issues on the physical host that impact network reachability
> 
> Either wait for AWS to fix the host, OR Move the EC2 instance to a new host = STOP & START the instance (if EBS backed).
>
> INSTANCE status checks - Instance status checks monitor the software and network configuration of your individual instance.\
> Example of issues
> 
> - Incorrect networking or startup configuration
> - Exhausted memory
> - Corrupted file system
> - Incompatible kernel
> 
> Requires your involvement to fix Restart the EC2 instance, OR Change the EC2 instance configuration.
> 
</details>

### EC2 Hibernate

<details>
<summary>
Preserving Data in RAM
</summary>

we can stop or terminate instances.
- stop - the data on disk (EBS) is kept intact
- terminate - any EBS volumes (root) also set-up to be destroyed is lost.

the first time a machine starts it boots the OS and runs the EC2 user data script, for later times, it just boots the OS without running the user script.

Hibernate allows us to store and preserve the data in RAM, we don't stop and restart the operation system. instead, we dump the data in RAM inside a special file in the root EBS volume. when we start the machine after hibernating, the data is loaded back into RAM, without requireing the OS to start up.

it is supported in many instance families, and for most popular AMI images. the root volume must be EBS based, and enctyped (not instance store) and has enough storage space. there is a limit for the instance RAM size (no more than 150GB), and it's not supported on *bare metal* instances. an instance can not be hibernated for more than 60 days.

we can see this in action by running the `uptime` command after hibernating compared to running the same command after "stopping" it.

</details>

</details>

## AMI - Amazon Machine Image

<details>
<summary>
Amazon Machine Image
</summary>

<cloud>AMI</cloud> - Amazon Machine Image. a customization of an EC2 image, we can add our own software, configuration, monitoring, OS... and then it's faster to boot and faster to configure.

AMIs are built for a specific region, and can then be copied across regions. there are public AMIs, but we can also use our own AMIs or get and AMI from the AMI marketplace (which is created and maintained by someone else).

we first create an EC2 machine and customize it. stop it, and then build the AMI (EBS snapshot), and we can then launch a new instance from the AMI we created.

we choose the <kbd>Create Image</kbd> option from the EC2 and follow the option.

### AMI No Reboot Option

<details>
<summary>
Option to create AMI without shutting the instance
</summary>

allows to create an AMI without shutting down the instance. we need to enable the option, this is faster. it's also part of the back-up plan to create AMIs without bringing down the EBS volume, but we don't have data integrity, and we can't be sure the OS data buffer will be flushed properly.

</details>

### EC2 Instance Migration using AMIs

<details>
<summary>
Migrating and Sharing AMI between Regions and Accounts.
</summary>

to migrate an EC2, we create an AMI from the instance, and then copy it to another region, where we can launch the EC2 from the AMI.

we can also share AMI's across accounts. this doesn't change the ownership of the AMI.\
For sharing to work, the AMI must either be un-encrypted, or be encrypted with a customer managed key. if the AMI has encrypted volumes, we also need to share the key used to encrypt the volumes.

sharing the key is down with IAM permissions.

we can also copy the AMI o a different account, this again requires sharing keys if the the image or the underlying snapshots are encrypted.

we can see the process under the "AMI permissions" tab.
</details>

### EC2 Image Builder

<details>
<summary>
Automatically build Custom AMIs
</summary>

this service is used to automate the creation of virtual machines or container images. we can also set validation tests to run on the newly created AMI and have the image be distributed across regions.  this is useful when we have dependencies which update and we want to re-build our AMI each time with the new versions.

we can create AMIs or Docker Images, and then we apply components as part of the recipe. there are pre-built components and we can make them ourselves. we can also use pre-built tests of define them. we need a <cloud>IAM</cloud> role wih some policies.
</details>


### AMI In Production

<details>
<summary>
Restrict Which AMIs can be used
</summary>

> You can force users to only launch EC2 instances from pre-approved AMIs (AMIs tagged with specific tags) using IAM policies.\
> Combine with AWS Config to find noncompliant EC2 instance (instances launched with non-approved AMIs)

</details>

</details>

## Systems Manager (SSM) - Managing EC2 at Scale

<details>
<summary>
Built-in Utilities for managing EC2 instances.
</summary>

> Helps you manage your <cloud>EC2</cloud> and On-Premises systems at scale.
> 
> - Get operational insights about the state of your infrastructure
> - Easily detect problems
> - Patching automation for enhanced compliance
> - Works for both Windows and Linux OS
> - Integrated with <cloud>CloudWatch</cloud> metrics / dashboards
> - Integrated with <cloud>AWS Config</cloud>
> - Free service

there are several sub systems in the service.

> - Resource Groups
> - Operations Management
>   - Explorer
>   - OpsCenter
>   - CloudWatch Dashboard
>   - PHD
>   - Incident Manager
> - Shared Resources
>   - Documents
> - Change Management
>   - Change Manager
>   - Automation
>   - Change Calendar
>   - Maintenance Windows
> - Application Management
>   - Application Manager
>   - AppConfig
>   - Parameter Store
> - Node Management
>   - Fleet Manager
>   - Compliance
>   - Inventory
>   - Hybrid Activations
>   - Session Manager
>   - Run Command
>   - State Manager
>   - Patch Manager
>   - Distributer

we need to install the <cloud>SSM</cloud> agent on the machine, it's pre-installed on the Amazon Linux AMI and some ubuntu AMIs. we also need a proper <cloud>IAM Role</cloud> attached.

We start with an example of EC2 machine, under the <cloud>SSM</cloud> service, we choose the <kbd>Fleet Manager</kbd> option. we need some virtual machines running, so we create a machine, it needs a role with the "AmazonSSMManagedInstanceCore" policy. in the security group option, we don't need to have it accessable from outside. we will start three instances.

Once they start running, we can see them in the fleet manager screen.

### AWS Tags & SSM Resource Groups

<details>
<summary>
Logical Groups of resources.
</summary>

we can add tags to many AWS resources. we use them for resource grouping, cost allocation and automation. we can group resources together based on tags.

we add tags to our EC2 machines from before.

we go the <kbd>Resource Group</kbd> service and there we can create a new resource group, a resource can belong to multiple groups at the same time.
</details>

### SSM Documents & SSM Run Command

<details>
<summary>
Running Scripts and commands on resources.
</summary>

SSM Documents can be Json or yaml, and can have parameters and actions, and services can execute them. documents can retrieve data from the <cloud>parameter store</cloud> service.

(like github actions)

we can use documents in
- run commands
- state manager
- patch manager
- automation

we can look at the amazon created documents, such as AWS-ApplyPatchBaseline.

we can execute a document as a script (the "Run Command" feature), and have it run on multiple instances (using the resource groups we created). fully integrated with <cloud>IAM</cloud> and <cloud>CloudTrail</cloud>, no need for SSH, can send notification to <cloud>SNS</cloud> and output the logs to <cloud>S3</cloud> or <cloud>CloudWatch</cloud> Logs. and can be triggered from <cloud>Event Bridge</cloud>.

for this to work, we need to open the security group to incoming data on port 80.

we create a simple document, make it target EC2 instances, and have it as a "Command document" (not session document"). this command will install the apache web server.

```yaml
---
schemaVersion: '2.2'
description: Sample YAML template to install Apache
parameters: 
  Message:
    type: "String"
    description: "Welcome Message"
    default: "Hello World"
mainSteps:
- action: aws:runShellScript
  name: configureApache
  inputs:
    runCommand:
    - 'sudo yum update -y'
    - 'sudo yum install -y httpd'
    - 'sudo systemctl start httpd'
    - 'sudo systemctl enable httpd'
    - 'echo "{{Message}} from $(hostname -f)" > /var/www/html/index.html'
```

in the <cloud>Run Command</cloud> service, we can choose this document, and then filter for the resource based on resource groups or other tags.

we can add a comment, control timeout, and add rate control - how many concurrent executions will happen, and set the error threshold (if the rate is above the threshold, we stop the operation). here we can also choose what happens to the logs and set notifications.

#### SSM Automations

we can create documents for repeated tasks, this is a "runbook" of common things we might want to do, such as taking an <cloud>EBS</cloud> snapshot, restarting an instance, creating <cloud>AMI</cloud>, etc...

unlike run commands, they perform action **ON** resources, rather than **INSIDE** them.\
There are predefined runbooks from AWS, or we can create our own. they can be triggered manually (by the web console, cli or sdk), by <cloud>EventBridge</cloud> rule, on a schedule using SSM Maintenance Windows, or with <cloud>AWS Config</cloud> for rules remediation.

in the console, we navigate to <kbd>Change Management</kbd> section and then <cloud>Automation</cloud> page. we choose <cloud>Execute Automation</cloud>, we select the document we want, such as "AWS-RestartEC2Instance", we can choose the target based on resource groups, and choose how the process will happen (cross region and accounts, step by step, rate control). we need an IAM role for the automation script.

</details>

### SSM Parameter Store

<details>
<summary>
Shared Storage of parameters.
</summary>

> - Secure storage for configuration and secrets.
> - Optional Seamless Encryption using <cloud>KMS</cloud>.
> - Serverless, scalable, durable, easy SDK.
> - Version tracking of configurations / secrets.
> - Security through <cloud>IAM</cloud>.
> - Notifications with Amazon <cloud>EventBridge</cloud>.
> - Integration with <cloud>CloudFormation</cloud>.

parameters are stored in hierarchal order, we define paths and store the parameters like files in a folder system, this makes controling access with IAM easier. we can use the parameter store to access secret from the secret store, or use public parameters from AWS (like finding out the most recent AMI in the region).

There are two tiers, standard free tier and advanced tier.

| Value                                                           | Standard Tier        | Advance Tier                           |
| --------------------------------------------------------------- | -------------------- | -------------------------------------- |
| Total number of parameters allowed (per AWS account and Region) | 10,000               | 100,000                                |
| Maximum size of a parameter value                               | 4 KB                 | 8 KB                                   |
| Parameter policies available                                    | No                   | Yes                                    |
| Cost                                                            | No additional charge | Charges apply                          |
| Storage Pricing                                                 | Free                 | $0.05 per advanced parameter per month |

Parameter policies allow us to set TTL to a parameter (expiration policy), and we can an eventBridge notification before the expiration date, or after it.

#### SSM Parameter Store Hands On

we navigate to the <kbd>Shared Resources</kbd> section and choose the <kbd>Parameter Store</kbd> option. now we can <kbd>Create Parameter</kbd>, give it a name (path), description, type (string, list of stings, secure string with <cloud>KMS</cloud>), choose the tier (standard or advanced) and the value of the parameter.


we can get the parameter value in the cli, and we can also ask to decrypt the secure strings if we have the permissions to access the key.

```sh
aws ssm get-parameter --names <parameter-name1> <parameter-name2>
aws ssm get-parameter --names <parameter-name1> <parameter-name2> --with-decryption
aws ssm get-parameters-by-path --path <parameter-path> --recursive
```

</details>

### SSM Inventory & State Manager

<details>
<summary>
Collect metadata from your managed instances.
</summary>

SSM Inventory 

> Collect metadata from your managed instances (EC2/On-premises)
>
> - Metadata includes installed software, OS drivers, configurations, installed Updates, running services...
> - View data in AWS Console or store in S3 and query and analyze using Athena and QuickSight.
> - Specify metadata collection interval (minutes, hours, days).
> - Query data from multiple AWS accounts and regions.
> - Create Custom Inventory for your custom metadata (e.g., rack location of each managed instance).

in the console, under <kbd>Node Management</kbd> section, we choose <kbd>Inventory</kbd> and enable the invtory on all instances. we can now navigate to the <kbd>State Manager</kbd> and make sure all our instances are in the proper state.

State Manager

> automate the process of keeping your managed instances (EC2/On-premises) in a state that you define
> - Use cases: bootstrap instances with software, patch OS/software Updates on a schedule...
> - State Manager Association:
>   - Defines the state that you want to maintain to your managed instances.
>   - Specify a schedule when this configuration is applied.
>   - Example: port 22 must be closed, antivirus must be installed...
> - Uses SSM Documents to create an Association (e.g., SSM Document to Configure CW Agent).

in the inventory page, we can see the coverage, check the most used OS, check the most used applications in the instances, and get detailed data in the resource group level. this data is sent to a <cloud>S3</cloud> bucket. we need to edit the bucket policy.\

</details>

### SSM Patch Manager and Maintenance Windows

<details>
<summary>
Automates the process of patching managed instances.
</summary>

> - OS updates, applications updates, security updates, etc...
> - Supports both EC2 instances and on-premises servers.
> - Supports Linux, macOS, and Windows.
> - Patch on-demand or on a schedule using Maintenance Windows.
> - Scan instances and generate patch compliance report (missing patches).
> - Patch compliance report can be sent to S3.
> - Patch Baseline
>   - Defines which patches should and shouldn't be installed on your instances.
>   - Ability to create custom Patch Baselines (specify approved/rejected patches).
>   - Patches can be auto-approved within days of their release.
>   - By default, install only critical patches and patches related to security.
> - Patch Group
>   - Associate a set of instances with a specific Patch Baseline.
>   - Example: create Patch Groups for different environments (dev, test, prod)
>   - Instances should be defined with the tag key Patch Group
>   - An instance can only be in one Patch Group
>   - Patch Group can be registered with only one Patch Baseline

find which patches are installed on instances, and which instances are missing patches (compliance report), have deny lists of patches which shouldn't be installed. we associate patch groups with a patch baseline. Patch groups are defined by a tag with the key "Patch Group", so they can only belong to one group.

there are predefined patch baselines managed by aws, and we can usually run a SSM document to apply them. we can also define our own custom patch baselines, with operating systems, allowed and rejected patches, and have external patch repositories.


we can define Maintenance windows, specifying when the actions are perfromed:
> - schedule
> - duration
> - set of registered instances
> - set of registered tasks

under the <kbd>Node Management</kbd> section, we navigate to the <kbd>Patch Management</kbd> page, here we can do a scan or scan-and-install, and we can define targets as usual. we can also view the patch baselines and see if an instance is compliant. we can also create the maintenance windows as Cron jobs, and set which documents to run.
</details>

### SSM Session Manager

<details>
<summary>
Run Shell Commands in Instances through a managed service.
</summary>

> Allows you to start a secure shell on your EC2 and on-premises servers
>
> - Access through AWS Console, AWS CLI, or Session manager SDK.
> - Does not need SSH access, bastion hosts, or SSH keys.
> - Supports Linux, macOS, and Windows.
> - Log connections to your instances and executed commands.
> - Session log data can be sent to S3 or CloudWatch Logs.
> - CloudTrail can intercept StartSession events.

we can run commands in our machines, without having to keep SSH keys, and we can log and audit all of our connections and commands.

we need IAM permissions, and we create polices to define which instances we can shell into, and even restrict which commands can be run in a session.

under the <kbd>Node Management</kbd> section, we navigate to the <kbd>Session Manager</kbd> page,
we click <kbd>Start Session</kbd>, and we can connect to any EC2 instance that has the SSM agent installed, even if it didn't open the 22 port for SSH access. we can view the open sessions and control timeout and manage logging.
</details>

</details>

## EC2 High Availability and Scalability

<details>
<summary>
Load Balancers and Auto-Scaling.
</summary>

> What is High Availability and Scalability?

> Scalability means that an application / system can handle greater loads by adapting. There are two kinds of scalability:\
> 
> - Vertical Scalability - scale up/down.
> - Horizontal Scalability (= elasticity) - scale out/in.
>
> Scalability is linked but different to High Availability.

Vertical scalability usually means increasing the compute power of each instance, such as moving it to a stronger instance type, it's commonly used for non-distributed services, such as some databases.\
Horizontal scalability usually means increasing the number of instances handling the workload, this happens in distributed systems, unlike vertical scaling, there is no hard limit on horizontal scaling.

High Availability usually goes with horizontal scaling, the goal is to have the application survive the loss of some of the instances. in AWS, this usually means surviving the loss of a data center (Availability Zone), it can also be a passive configuration, such as having a multi AZ RDS.

### Elastic Load Balancing (ELB)

<details>
<summary>
Managed Load Balancers.
</summary>

> Load Balances are servers that forward traffic to multiple servers (e.g., EC2 instances) downstream.
>
> - Spread load across multiple downstream instances.
> - Expose a single point of access (DNS) to your application.
> - Seamlessly handle failures of downstream instances.
> - Do regular health checks to your instances.
> - Provide SSL termination (HTTPS) for your websites.
> - Enforce stickiness with cookies.
> - High availability across zones.
> - Separate public traffic from private traffic.

<cloud>ELB</cloud> is a managed load balancer service, so it's managed by AWS, which takes care of the upgrades, patching and High Availability requirements. it is also integrated with many other services:
- <cloud>EC2</cloud> AutoScaling Groups
- <cloud>ECS</cloud> - elastic container service
- <cloud>ACM</cloud> - aws certificate manager
- <cloud>Route53</cloud> - dns
- <cloud>AWS WAF</cloud> - firewall
- <cloud>AWS Global Accelerator</cloud>

Load Balancers work with health checks, which tell the upstream load balancer if the downstream instance can handle traffic at the current moment.

there are four types of elastic load balancers, but one of them is already deprecated:
- Classic Load Balancer - http, https, TCP, SSL - not used anymore.
- Application Load Balancer - http, https, websocket
- Network Load Balancer - TCP, TLS (secure TCP), UDP
- Gateway Load Balancer - network layer (layer 3) - ip protocol

load balancers can be internal or external. the security group should allow connections from everywhere, but the the downstream instances should only allow access from the load balancers.

#### Application Load Balancer (ALB)

<details>
<summary>
Basic Layer 7 load balancer.
</summary>

Layer 7 load balancer (HTTP layer), routes to multiple http applications across machine (target groups), and even multiple applications on the same machine (like different containers). supports HTTP/2 and websocket, and allows for redirects.\
Routing options:
- url path (example.com/**users** and example.com/**posts**)
- url hostname (**one**.example.com and **two**.example.com)
- query string headers (example.com?**order=false**)

ALB are a great fit for micro services and container based applications (docker, kubernetes, ECS), and it has port mapping feature to redirect to a dynamic port in ECS.

target groups can be:
- <cloud>EC2</cloud> instances
- <cloud>ECS</cloud> tasks
- <cloud>Lambda</cloud> functions
- private IP addresses

> - health check are done at the target group level.
> - Fixed hostname (XXX.region.elb.amazonaws.com)
> - The application servers don't see the IP of the client directly
>   - The true IP of the client is inserted in the header X-Forwarded-For
>   - We can also get Port (X-Forwarded-Port) and proto (X-Forwarded-Proto)

in the demo, we create two instances of EC2 machines, we use the scirpt to install the httpd web server on the machines. we have the security group allow traffic from outside, (we will change this later).

now we <kbd>Create Load Balancer</kbd>, and choose the Application Load Balancer type. we give it a name, make it external facing (not internal), and use IPv4. we need to deploy it in one or more Availability Zones, and we create a new security group with all http traffic allowed. next, at the <kbd>Listeners and Routing</kbd>, we need to create the target group for our instances, we set the health check path, and register the instances to it.  we can link the target group to the listener on port 80. when we create a load balancer, we get a DNS name that we can use in the browser, and it will direct us to one of the registered instances.\
if we stop one of the instances, the load balancer will see it's inactive (based on the health check), and it won't direct the traffic to it anymore.

we can go to the instances and and fix the security group, and remove the access from the 0.0.0.0/0 cidr block, and only allow access from the security group the ALB belongs to.\
we can also set more complex listener rules, we can modify the routing by adding rules to match requests and set different targets. rules have conditions (url path, url hostname, query parameters, request method, ip source) and then we can set the target group, redirect or return a fixed response. rules also have priority, with the last one being the default rule - what happens when no rule matches the request.
</details>


#### Network Load Balancer (NLB)

<details>
<summary>
High performance Layer 4 load balancer.
</summary>

forwards TCP and UDP traffic to instances. handles millions of requests per second with low latency. it is not part of the AWS free tier.\
it has only one static IP per Availability Zone, and it supports assinging elastic IP (usefull  when whitelisting a specific IP).

target groups can be:
- <cloud>EC2</cloud> instances
- private IP addresses
- application load balancer

health checks support TCP, HTTP and HTTPS protocols.

like before, we can create a network load balancer, we get a fixed IP4 address from AWS for each Availability Zone, but we can also provide our own elastic IP. we can also attach a security group. we need to modify the security group of the downstream instances to allow the health checks to pass.

</details>

#### Gateway Load Balancer (GWLB)

<details>
<summary>
Layer 3 Load Balancer - manages all requests before forwarding them.
</summary>

a way to have all traffic go through a single point of entry (the gateway load balancer), which is first directed at some virtual appliance, and if that appliance confirms it, it is then directed to the target application.

> - Deploy, scale, and manage a fleet of 3rd party network virtual appliances in AWS
> - Example: Firewalls, Intrusion Detection and Prevention Systems, Deep Packet Inspection Systems, payload manipulation, etc...
> - Operates at Layer 3 (Network Layer) – IP packets
> - Combines the following functions:
>   - Transparent Network Gateway – single entry/exit for all traffic
>   - Load Balancer – distributes traffic to your virtual appliances
> - Uses the **GENEVE** protocol on port **6081**

target groups (which analyze the requests):
- <cloud>EC2</cloud> instances
- private Ip addresses

</details>


#### Sticky Sessions

It is possible to implement stickiness so that the
same client is always redirected to the same
instance behind a load balancer.

> - This works for Classic Load Balancer, Application Load Balancer, and Network Load Balancer
> - For both CLB & ALB, the "cookie" used for stickiness has an expiration date you control
> - Use case: make sure the user doesn’t lose his session data
> - Enabling stickiness may bring imbalance to the load over the backend EC2 instances
> - NLB doesn't use cookies

cookies can be application based (custom or load balancer generated), or be duration based (always created by the load balancer). we can set the cookie name ourselves, but there are some reserved names.

#### Cross Zone Load Balancing

if we have an imbalance in the number of instances inside each Availability Zone, the traffic will be directed evenly across the Availability Zones, so some instances will get larger load of the traffic.\
if we set "cross zone load balancing" the traffic will be directed evenly across all instances.

| ELB                               | Default  | Cross-Zone Traffic Charges |
| --------------------------------- | -------- | -------------------------- |
| Classic Load Balancer             | disabled | no charges                 |
| Application Load Balancer         | enabled  | no charges                 |
| Network and Gateway Load Balancer | disabled | charges apply              |

#### SSL Certificates

> An SSL Certificate allows traffic between your clients and your load balancer to be encrypted in transit (in-flight encryption).
>
> - SSL refers to Secure Sockets Layer, used to encrypt connections.
> - TLS refers to Transport Layer Security, which is a newer version.
> - Nowadays, TLS certificates are mainly used, but people still refer as SSL.
> - Public SSL certificates are issued by Certificate Authorities (CA).
> - Comodo, Symantec, GoDaddy, GlobalSign, Digicert, Letsencrypt, etc...
> - SSL certificates have an expiration date (you set) and must be renewed.
>
> The load balancer uses an X.509 certificate (SSL/TLS server certificate)
> - You can manage certificates using <cloud>ACM</cloud> (AWS Certificate Manager).
> - You can create and upload your own certificates alternatively.
> - HTTPS listener:
>   - You must specify a default certificate.
>   - You can add an optional list of certs to support multiple domains.
>   - Clients can use **SNI** (Server Name Indication) to specify the hostname they reach.
>   - Ability to specify a security policy to support older versions of SSL/TLS (legacy clients).

SNI allow loading multiple certificates onto one websever (the load balancer), so it can serve multiple websites (downstream targets). the clients needs to indicate the target hostname for the initial SSL handshake. the server will find the correct certificate, or use the default one.

| ELB                       | Version | SNI support | SSL Certificates      |
| ------------------------- | ------- | ----------- | --------------------- |
| Classic Load Balancer     | V1      | No          | One Certificate       |
| Application Load Balancer | V2      | Yes         | Multiple Certificates |
| Network Load Balancer     | V2      | Yes         | Multiple Certificates |

SNI is also supported with <cloud>CloudFront</cloud>.

when we go to our load balancers <cloud>ALB</cloud>, we can add a listener, and set the "Secure listener settings" with a security policy, and we can add certificates. for <cloud>ELB</cloud>, we can also add ALPN (Application Level Protocol Negotiation).

#### De-Registration Delay & Connection Draining

With Classic Load Balancers, it's called connection draining, but with newer types, it's called De-Registration Delay. it gives the instance time to complete "in-flight" requests, it won't send new requests to those instances, but it will give a grace period for requests to complete. default value is 300 seconds, but it can start from zero second (no grace period) to 3600 seconds - which is an hour. we decide the period based on the time our requests take to complete.

#### Health Checks

some settings that we can configure for health checks.

| Setting                    | Value | Description                                                     |
| -------------------------- | ----- | --------------------------------------------------------------- |
| HealthCheckProtocol        | HTTP  | Protocol used to perform health checks                          |
| HealthCheckPort            | 80    | Port used to perform health checks                              |
| HealthCheckPath            | `/`   | Destination for health checks on targets                        |
| HealthCheckTimeoutSeconds  | 5     | Consider the health check failed if no response after 5 seconds |
| HealthCheckIntervalSeconds | 30    | Send health check every 30 seconds                              |
| HealthyThresholdCount      | 3     | Consider the target healthy after 3 successful health checks    |
| UnhealthyThresholdCount    | 5     | Consider the target unhealthy after 5 failed health checks      |

> Target Health status options:
>
>- Initial: registering the target
>- Healthy
>- Unhealthy
>- Unused: target is not registered
>- Draining: de-registering the target
>- Unavailable: health checks disabled

if all the targets are unhealthy, the ELB will send them the requests anyway. health checks are best effort.

#### Elastic Load Balancer - Monitoring, Troubleshooting, Logging and Tracing

Response codes are standard for HTTP requests:

- Successful request: Code 200.
- Unsuccessful at client side: 4XX code.
    - 400: Bad Request    
    - 401: Unauthorized   
    - 403: Forbidden  
    - 460: Client closed connection.  
    - 463: X-Forwarded For header with >30 IP (Similar to malformed request). 
- Unsuccessful at server side: 5xx code.
    - 500: Internal server error would mean some error on the ELB itself.
    - 502: Bad Gateway
    - 503: Service Unavailable
    - 504: Gateway timeout: probably an issue within the server.
    - 561: Unauthorized

there are metrics for the load balancers which are directly pushed into <cloud>CloudWatch</cloud>

> Metrics:
>
> - BackendConnectionErrors
> - HealthyHostCount / UnHealthyHostCount
> - HTTPCode_Backend_2XX: Successful request.
> - HTTPCode_Backend_3XX: redirected request
> - HTTPCode_ELB_4XX: Client error codes
> - HTTPCode_ELB_5XX: Server error codes generated by the load balancer.
> - Latency
> - RequestCount
> - RequestCountPerTarget
> - SurgeQueueLength: The total number of requests (HTTP listener) or connections (TCP listener) that are pending routing to a healthy instance. Help to scale out ASG. Max value is 1024
> - SpillOverCount: The total number of requests that were rejected because the surge queue is full.
>
> Load Balancer troubleshooting using metrics
> -  HTTP 400: BAD_REQUEST => The client sent a malformed request that does not meet HTTP specifications.
> -  HTTP 503: Service Unavailable => Ensure that you have healthy instances in every Availability Zone that your load balancer is configured to respond in. look for HealthyHostCount in CloudWatch.
> -  HTTP 504: Gateway Timeout => Check if keep-alive settings on your EC2 instances are enabled and make sure that the keep-alive timeout is greater than the idle timeout settings of load balancer.

Access Logs can be stored in S3, we only pay for hte S3 storage, we might need them for compliance or debugging.\
There is also a built-in request tracing custom header (`X-Amzn-Trace-Id`), but it's ELB isn't integrated with <cloud>X-Ray</cloud> yet.

#### Target Group Attributes

Target Groups Settings

> - deregisteration_delay.timeout_seconds: time the load balancer waits before
deregistering a target
> - slow_start.duration_seconds: 0 for disabled, duration in which the target recives a smaller share of the traffic
> - load_balancing.algorithm.type: how the load balancer selects targets when routing requests (Round Robin, Least Outstanding Requests)
> - stickness - Sticky Session
>   - stickiness.enabled: true or false
>   - stickiness.type: application-based or duration-based cookie
>   - stickiness.app_cookie.cookie_name: name of the application cookie
>   - stickiness.app_cookie.duration_seconds: application-based cookie expiration period
>   - stickiness.lb_cookie.duration_seconds: duration-based cookie expiration period

slow start mode is a way to slowly ramp up new instances, during the slow start duration, the instance receive less traffic, gradually increasing towards the weighted amount.

the request routing is done based on algorithms, which determine which request should go to which instance.
- Round Robin - equally choose instances from the target group, cycle in order. works with application load balancers and classic load balancers.
- Least Outstanding Requests - send the request to the instance with the lowest number of pending or unfinished requests. works with application load balancers and classic load balancers. **doesn't work with slow start**.
- Flow Hash - selects the target based on hasing of the request (protocol, source and destination ip and port, TCP sequence number), same request from the same user will reach the same target. works with network load balancers.

### ALB Rules - Deep Dive

Rules are set on a listener, they are processed in order of priority.

supported actions:
- forward
- redirect
- fixed-response

rule conditions:
- host-header
- http-request-method
- path-pattern
- source-ip
- http-header
- query-string

a single rule can have multiple target groups, and each target group can have different weights. this allows us for gradual deployment of new version (blue/green deployment), and for better control of how the traffic is distributed.
</details>

### Auto Scaling Groups (ASG)

<details>
<summary>
Automatically Add and Remove EC2 Instances.
</summary>

the load on on our websites can change over time, and in the cloud, it's easy to add and remove instances. therefore, we can use Auto Scaling Groups to scale out and scale in based on demand, adding and removing instances. we can also pair it with a load balancer to automatically register new instances to it. unhealthy EC2 instances are automatically removed and replaced with new ones.

There are no costs for using autoscaling groups, just the cost of the underlying instances.

when we create an auto-scaling group, we choose the minimum and maximum capacity, and the desired capacity to start with. when combined with ELB, it can remove and replace unhealthy instances based on the health check.

when we create an auto scaling group, we need:
- Launch Template (older "launch configurations" are deprecated) - how to launch instances (ami, storage, networking and security groups)
- Min/Max/Desired Size
- auto scaling strategy - can be based on <cloud>CloudWatch</cloud> alarms and metrics (like Average CPU).

under the <cloud>EC2</cloud> Service, in <kbd>Auto Scaling Group</kbd>, we can create a new one. we first need the launch template, which we fill with the AMI, instance type, storage, ssh keys and user data script. we can prefill each section, or have it open for changes later on. we can still override the values in the launch template if needed. we can attach the auto scaling group to one of our load balancers, and then set the health check. we need to set the capacity (min, max, desired) and set the scaling policy, and add notifications if we wish to be informed when something changes.

once this is created, we can play with the desired capacity and see how instances are created and removed.

#### Auto Scaling Groups - Scaling Policies

Scaling Policies (or strategies) control when new instances are created or removed.

> - **Dynamic Scaling**
>   - **Target Tracking Scaling**
>     - Simple to set-up
>     - Example: I want the average ASG CPU to stay at around 40%
>   - **Simple / Step Scaling**
>     - Example: When a CloudWatch alarm is triggered (example CPU > 70%), then add 2 units
>     - Example: When a CloudWatch alarm is triggered (example CPU < 30%), then remove 1
> - **Scheduled Scaling**
>   - Anticipate a scaling based on known usage patterns
>   - Example: increase the min capacity to 10 at 5 pm on Fridays
> - **Predictive Scaling**: continuously forecast load and schedule scaling ahead.

We can scale based on metrics, we usually scale on CPU utilization, number of request, network in/out load, or on any of our own custom metrics.

after a scaling activity happens, there is a cooldown period when we wait for the metrics to stabilize, and during this time there won't be additional ASG activies.

#### Auto Scaling Groups - Additional Actions

ways to run scripts on instances during their life cycle.

> Lifecycle Hooks
> - By default, as soon as an instance is launched in an ASG it's in service
> - You can perform extra steps before the instance goes in service (Pending state)
> - Define a script to run on the instances as they start
> - You can perform some actions before the instance is terminated (Terminating state)
> - Pause the instances before they're terminated for troubleshooting
> - Use cases: cleanup, log extraction, special health checks
> - Integration with EventBridge, SNS, and SQS

Launch Configuration (legacy) vs. Launch Template (new)
Both allow us to set the ID of the Amazon Machine Image (AMI), the instance type, a key pair, security groups, and the other parameters that you use to launch EC2 instances (tags, EC2 user-data...). we can't edit either Launch Configurations or Launch Templates.

> Launch Configuration (legacy):
> - Must be re-created every time
> 
> Launch Template (newer):
> - Can have multiple versions
> - Create parameters subsets (partial configuration for re-use and inheritance)
> - Provision using both On-Demand and Spot instances (or a mix)
> - Supports Placement Groups, Capacity Reservations, Dedicated hosts, multiple instance types
> - Can use T2 unlimited burst feature


we can scale our AGG based on <cloud>SQS</cloud>, using a <cloud>CloudWatch</cloud> alarm that looks at the number of the messages in the queue.

for High Availability, we need a least two instances in different Availability Zones (the ASG must be multi-AZ). we need to have health check (EC2, ELB, or custom). when the health check fails, we terminate the instance, we don't reboot it.

```sh
aws autoscaling set-instance-health
aws autoscaling terminate-instance-in-autoscaling-group
```

if a ASG fails to launch an instance for more than 24 hours, the process is suspended.

#### CloudWatch for ASG

> Metrics are collected every 1 minute
> - ASG-level metrics: (opt-in)
>   - You should enable metric collection to see these metrics
>   - GroupMinSize, GroupMaxSize, GroupDesiredCapacity
>   - GroupInServiceInstances, GroupPendingInstances, GroupStandbyInstances
>   - GroupTerminatingInstances, GroupTotalInstances
> - EC2-level metrics (enabled): 
>   - CPU Utilization, etc...
>   - Basic monitoring: 5 minutes granularity
>   - Detailed monitoring: 1 minute granularity

</details>

### Auto Scaling

<details>
<summary>
The Service at the back of auto scaling groups.
</summary>

A backbone service for scalable resources in AWS:

> - Amazon EC2 Auto Scaling groups: Launch or terminate EC2 instances
> - Amazon EC2 Spot Fleet requests: Launch or terminate instances from a Spot Fleet request, or automatically replace instances that get interrupted for price or capacity reasons.
> - Amazon <cloud>ECS</cloud>: Adjust the ECS service desired count up or down
> - Amazon <cloud>DynamoDB</cloud> (table or global secondary index):WCU & RCU
> - Amazon <cloud>Aurora</cloud>: Dynamic Read Replicas Auto Scaling

we can create scaling plans from the auto-scaling service, rather than from the auto-scaling group.

> Scaling Plans:
>
> - Dynamic scaling: creates a target tracking scaling policy
> - Optimize for availability => 40% of resource utilization
> - Balance availability and cost => 50% of resource utilization
> - Optimize for cost => 70% of resource utilization
> - Custom => choose own metric and target value
> - Options: Disable scale-in, cooldown period, warmup time (for ASG)
> - Predictive scaling: continuously forecast loadand schedule scaling ahead


</details>


</details>

## Elastic BeanStalk for SysOps

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

## CloudFormation for SysOps

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

### CloudFormation - Create Stack - Hands On
### CloudFormation - Update & Delete Stack - Hands On
### YAML Crash Course
### CloudFormation - Resources
### CloudFormation - Parameters
###  CloudFormation - Mappings
###  CloudFormation - Outputs & Exports
### CloudFormation - Conditions
### CloudFormation - Intrinsic Functions
###  CloudFormation - Rollbacks
### CloudFormation - Service Role
### CloudFormation - Capabilities
### CloudFormation - Deletion Policy
### CloudFormation - Stack Policy
###  CloudFormation - Termination Protection
### CloudFormation - Custom Resources
### CloudFormation - Dynamic References
### CloudFormation - User Data
### CloudFormation - cfn-init
### CloudFormation - cfn-signal & Wait Condition
### CloudFormation - cfn-signal Failures
### CloudFormation - Nested Stacks
### CloudFormation - Depends On
### StackSets - Warning
### CloudFormation - StackSets
### CloudFormation - Create StackSets - Hands On
### CloudFormation - Update StackSets - Hands On
### CloudFormation - Delete StackSets - Hands On
### CloudFormation - Troubleshooting

</details>

## Lambda for SysOps
## EC2 Storage and Data Management - EBS and EFS
## Amazon S3 Introduction
## Advanced Amazon S3 and Athena
## Amazon S3 Security
## Advanced Storage Section
## CloudFront
## Databases for SysOps
## Monitoring, Auditing and Performance
## Aws Account Management
## Disaster Recovery
## Security and Compliance for SysOps
## Identity
## Networking - Route53
## Networking - VPC
## Other Services


## Misc

<details>
<summary>
Stuff worth remembering.
</summary>

### Additional Terms

<details>
<summary>
Additional terms and acronyms to keep.
</summary>

- ENA - Elastic network adapter
- EFA - Elastic fabric adapter, improved ENA
- SSL - Secure Sockets Layer
- TLS - Transport Layer Security (new SSL)
- SNI - Server Name Indication (works with SSL)
- ALPN - Application Level Protocol Negotiation (Certificate)

</details>

</details>

</details>