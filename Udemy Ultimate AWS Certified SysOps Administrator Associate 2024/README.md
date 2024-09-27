<!--
// cSpell:ignore proto deregisteration_delay sysvinit Teradata xvda POSIX Apahce etag requesturi SERDE DSSE ONTAP NTFS PITR HDFS MITM fileb certutil FIPS ADFS OIDC ICANN NAPTR nslookup IMDSV NACL Mbps ECMP
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# Ultimate AWS Certified SysOps Administrator Associate 2024

<summary>
Practice towards AWS Certified SysOps Administrator Associate exam.
</summary>

Udemy course [Ultimate AWS Certified SysOps Administrator Associate 2024](ultimate-aws-certified-sysops-administrator-associate). by [_Stephane Maarek_](https://www.stephanemaarek.com/).

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
>
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
Controlling how the EC2 instances are placed in the data center.
</summary>

> - _Cluster_ - clusters instances into a low-latency group in a single Availability Zone.
> - _Spread_ - spreads instances across underlying hardware (max 7 instances per group per AZ) - critical applications.
> - _Partition_ - spreads instances across many different partitions (which rely on different sets of racks) within an AZ. Scales to 100s of EC2 instances per group (Hadoop, Cassandra, Kafka)

In the _cluster_ placement group, we stick all the instances in the same AZ, which gives us better networking, and we can use enhanced networking. the risk is that if the Availability Zone fails, all the instances fail. we use this strategy when we have big jobs that need to complete fast, or when our applications need extremely low latency and high network throughput.\
For the _spread_ placement groups, we span across multiple Availability Zone and hardware, this reduces the risk of failure (better Availability). however, we ate limited to having 7 instances per Availability Zone per placement group.
The _Partition_ group uses partitions to separate machines. instances in a partition don't share physical racks with instances from other partitions, so the partition represents a rack inside the data center. EC2 instances have the partiton information as part of the metadata. this is usually used in big-data application which are partition-aware

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

there is an old option of _Spot Blocks_, which allowed to get spot instances that won't be interrupted once they started (it wasn't 100% guranteed either).

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
- priceCapacityOptimized (recommended) - first the pool with the highest capacity available. then select the lowest priced pool.

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
> - If the machine stops bursting, credits are accumulated over time
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
>     - Instance statues for the EC2 VM
>     - System status for the underlying hardware
>   - Disk (instance store only) - read/write for ops/bytes. (doesn't apply to <cloud>EBS</cloud>-based instances).
> - **Doesn't Include RAM**
>
> Custom metric (yours to push):
>
> - Basic Resolution: 1 minute resolution
> - High Resolution: all the way to 1 second resolution
> - Include RAM, application level metrics
> - Make sure the IAM permissions on the EC2 instance role are correct!

for each EC2 machine, we click the <kbd>Monitoring</kbd> tab and see the metrics, we can also add them to <cloud>CloudWatch</cloud> dashboards for a centralized location. we can enabled detailed monitoring for the machine if we want, but it will cost us more.

#### Unified CloudWatch Agent

This agent allows us to gather metrics from EC2 instances or from other servers (such as on-premises machines). this can collect system level metrics (RAM, processes, used disk space), and can send the to <cloud>AWS CloudWatch</cloud>.

we can configure the agent with <cloud>SSM parameter store</cloud> or a configuration file, the machine needs to have the correct <cloud>IAM role</cloud> and permissions. the metrics from the unified agent all have the default "CWAgent" namespace. there is a _procstat_ plugin that can collect metrics for individual processes running on the machine. we can select which processes are monitored with the pid identification number or based on regex (process name, command line which started it). mertics from the plugin will have the "procstat" prefix.

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

it is supported in many instance families, and for most popular AMI images. the root volume must be EBS based, and enctyped (not instance store) and has enough storage space. there is a limit for the instance RAM size (no more than 150GB), and it's not supported on _bare metal_ instances. an instance can not be hibernated for more than 60 days.

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

this service is used to automate the creation of virtual machines or container images. we can also set validation tests to run on the newly created AMI and have the image be distributed across regions. this is useful when we have dependencies which update and we want to re-build our AMI each time with the new versions.

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
>
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
>
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
>
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
it has only one static IP per Availability Zone, and it supports assinging elastic IP (usefull when whitelisting a specific IP).

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
> - Operates at Layer 3 (Network Layer) - IP packets
> - Combines the following functions:
>   - Transparent Network Gateway - single entry/exit for all traffic
>   - Load Balancer - distributes traffic to your virtual appliances
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
> - Use case: make sure the user doesn't lose his session data
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
>
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
>
> - HTTP 400: BAD_REQUEST => The client sent a malformed request that does not meet HTTP specifications.
> - HTTP 503: Service Unavailable => Ensure that you have healthy instances in every Availability Zone that your load balancer is configured to respond in. look for HealthyHostCount in CloudWatch.
> - HTTP 504: Gateway Timeout => Check if keep-alive settings on your EC2 instances are enabled and make sure that the keep-alive timeout is greater than the idle timeout settings of load balancer.

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
>
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
>
> - Must be re-created every time
>
> Launch Template (newer):
>
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
>
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

<details>
<summary>
A IaaS service to create EC2 based applications.
</summary>

a way to deploy EC2 based applications, with auto-scaling groups and load balancers.

re-using the same configurations and storing them as code.

- capacity provisioning
- load balancing
- scaling
- health monitoring
- instance configuration

free service, only the underlying used instances count towards the cost.

> - Application: collection of Elastic Beanstalk components (environments, versions, configurations, ...)
> - Application Version: an iteration of your application code
> - Environment-  Collection of AWS resources running an application version (only one application version at a time)
>   - Tiers: Web Server Environment Tier & Worker Environment Tier
>   - You can create multiple environments (dev, test, prod, ...)

supports many languages, and also container based application, can be used in web-server tier or with worker-tier (based on <cloud>SQS</cloud>).

deployment modes:

- single instance with elastic IP- great for development
- high availability mode with load balancer

when we first create the environment, we need to choose between the two environment tiers - web server and worker. then we can give the environment and the application itself a name, add tags, choose a domain if we want (or use he auto-generated one), we next choose the platform, we used node.js for the demo. and we can add the application code, in our case we use the sample application code.\
there are several pre-sets, controlling the number of instances (single or high availability) and choosing to use on-demand <cloud>EC2</cloud> machines or spot instances. we need some <cloud>IAM</cloud> roles, one for the <cloud>Elastic Beanstalk</cloud> service and another for the virtual machines. we can set up networking, databases, and scaling configurations. and then click <kbd>submit</kbd> to create the environment.\
under the hood, a <cloud>CloudFormation</cloud> stack is created. and we can see the events, the resources, and the template. once the environment is created, we can navigate to the domain and see that it's operational. we could click <kbd>Upload and deploy</kbd> to release a new version of our application code, and use the tabs to view events, logs, health, alarms and scheduled updates.
</details>

## CloudFormation for SysOps

<details>
<summary>
IaaS for AWS resources.
</summary>

Declarative resource managed resource deployment, uses templates.

Benefits:
>
> - Infrastructure as code
>   - No resources are manually created, which is excellent for control
>   - The code can be version controlled for example using Git
>   - Changes to the infrastructure are reviewed through code
>
> - Cost
>   - Each resources within the stack is tagged with an identifier so you can easily see how much a stack costs you
>   - You can estimate the costs of your resources using the CloudFormation template
>   - Savings strategy: In Dev, you could automation deletion of templates at 5 PM and recreated at 8 AM, safely
>
> - Productivity
>   - Ability to destroy and re-create an infrastructure on the cloud on the fly
> - Automated generation of Diagram for your templates!
> - Declarative programming (no need to figure out ordering and orchestration)
>
> - Separation of concern: create many stacks for many apps, and many layers. Ex:
>   - VPC stacks
>   - Network stacks
>   - App stacks
>
> - Don't re-invent the wheel
>   - Leverage existing templates on the web!
>   - Leverage the documentation

templates are immutable, stored in S3, changes to cloudFormation templates create new templates. a template creates a stack which holds the resources. when a stack is removed, all the resources which it created are removed in reverse order.

the deployment can be done manually from the web console (editing a template, filling in any variable), or automated thorugh the cli or a CI-CD pipeline.

### CloudFormation Hands-On Demo

<details>
<summary>
Simple demo of creating, updating and removing a stack.
</summary>

we should do this demo in the us-east-1 region. at the <cloud>CloudFormation</cloud> service, we can see our stacks if we have them from the <cloud>Elastic Beanstalk</cloud> demo, but we should click <kbd>Create Stack</kbd>. here we can can choose to use a sample template, build one from the application composer, or to provide our own template. we can take a template from S3, git repository or upload directly. for now, we se;ect the wordpress-blog sample, we can click <kbd>View in Application Composer</kbd> for a visual view of the stack, and we see the componenets and how they relate to one another. we can see the template in yaml or json formats. we won't deploy that template, instead, we will use a simple template to just create an <cloud>EC2</cloud> machine.

```yaml
---
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-0a3c3a20c09d6f377
      InstanceType: t2.micro
```

(we can get a cloud formation plug-in to vscode)

we upload this file, give the stack a name, and for now we skip the other option. we can create the stack, and once it's completed, we can go to the "resources" tab and navigate to the resource we created, the machine will have tags tor the stack (stack id, stack name, resource name).

the only way to update a stack is with a new template, so we will upload the new template file:

```yaml
---
Parameters:
  SecurityGroupDescription:
    Description: Security Group Description
    Type: String

Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-0a3c3a20c09d6f377
      InstanceType: t2.micro
      SecurityGroups:
        - !Ref SSHSecurityGroup
        - !Ref ServerSecurityGroup

  # an elastic IP for our instance
  MyEIP:
    Type: AWS::EC2::EIP
    Properties:
      InstanceId: !Ref MyInstance

  # our EC2 security group
  SSHSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Enable SSH access via port 22
      SecurityGroupIngress:
      - CidrIp: 0.0.0.0/0
        FromPort: 22
        IpProtocol: tcp
        ToPort: 22

  # our second EC2 security group
  ServerSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: !Ref SecurityGroupDescription
      SecurityGroupIngress:
      - IpProtocol: tcp
        FromPort: 80
        ToPort: 80
        CidrIp: 0.0.0.0/0
      - IpProtocol: tcp
        FromPort: 22
        ToPort: 22
        CidrIp: 192.168.1.1/32
```

in this template we have more resources (ec2 machine, elastic ip, security groups), we have a parameter section for dynamic variables, and the `!Ref` keyword. we can upload the template and view it in the designer. now when we try to create the stack,we are are asked to enter a parameter. before submit the update, we are provided with a change set preview.\
This change set details what actions will be taken - we are adding some resources, but also modifying the already existing one (either in place or replace). other template updates could remove resources and modify them without replacing them. when we click <kbd>Submit</kbd>, we can wait for the resources to be created and removed as needed.

When we are ready to remove the resources, we could go to each resource and manually remove them, but that is hard manual work which needs to be done in a specific order and is error-prone (we might forget to remove the elasticIP instance after we disassociated it from the <cloud>EC2</cloud> machine). instead, we go to the stack page and click <kbd>Delete</kbd>, which will remove the resources in the correct order, and make sure we don't forget anything.

#### YAML Crash Course

a declarative format to store data, like Json or XML.

> - Key value Pairs
> - Nested objects
> - Support Arrays
> - Multi line strings
> - Can include comments!

```yaml
key: value # this is a comment
key2: value2
object:
  property: value3
  property2: value4
list:
  - objectNameKey: objectName
    objectDescription: objectDesc
  - objectNameKey: object2Name
    objectDescription: object2Desc
lines: |
  long line
  of text
  combined together
```

</details>

### CloudFormation Template Components

<details>
<summary>
The different sections in CloudFormation Template
</summary>

a <cloud>CloudFormation</cloud> template has different sections, some of which are mandatory and some are optional. all cloudFormation templates have the same structure.

> Template Components
>
> - AWSTemplateFormatVersion - identifies the capabilities of the template ("2010-09-09")
> - Description - comments about the template
> - Resources (MANDATORY) - your AWS resources declared in the template
> - Parameters - the dynamic inputs for your template
> - Mappings - the static variables for your template
> - Outputs - references to what has been created
> - Conditionals - list of conditions to perform resource creation

we also have template helpers: _references_ and _functions_.

#### CloudFormation Resources

The core of the template, this section defines which resources are used in the stack. this is the mandatory section of the template.

> - Resources are declared and can reference each other
> - AWS figures out creation, updates and deletes of resources for us
> - There are over 700 types of resources (!)
> - Resource types identifiers are of the form: `service-provider::service-name::data-type-name`

example:

```yaml
Type: AWS::EC2::Instance
```

[list of resource types](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)

we can use the documentation to see how to create each type of resource, what are the required and optional properties, and if that property change requires interrupting the resource (might require replacing the instance) or if it can updated in-place. most of the AWS resources are supported, and if we need to use one that isn't supported yet we can use a custom resource defintion. in theory, there can be non-AWS resources, but there aren't any.

It is possible to create a dynamic number of resources using Macros and transform, but that's outside the scope for now.

#### CloudFormation Parameters

> Parameters are a way to provide inputs to your
> AWS CloudFormation template.\
> Parameters are extremely powerful, controlled, and can prevent errors from happening in your templates, thanks to types.\
> They're important to know about if:
>
> - You want to reuse your templates across the company.
> - Some inputs can not be determined ahead of time.

parameters allow us to re-use the same template with small changes, without having to re-upload the template again and again with minor changes.

parameters have a type:

- string
- number
- comma delimited list
- list\<number>
- AWS specific parameter (matched against existing values in AWS)
- list of AWS specific parameters
- <cloud>SSM</cloud> Parameter (to get from the parameter store)

parameters also have a description with a possible constraint:

- min/max length
- min/max value
- default value
- allowed values array
- allowed values regex
- "NoEcho" (boolean option)

```yaml
Parameters:
  InstanceType: # parameter Name
    Description: Choose an EC2 instance type
    Type: String
    AllowedValues:
      - t2.micro
      - t2.small
      - t2.medium
    Default: t2.micro
  DBPassword:
    Description: the database admin password
    Type: string
    NoEcho: true #don't expose this value in logs or anywhere else
```

we can reference parameters using `Fn::Ref`, or the shorthand form `!Ref`.

there are also "Pseudo Parameters", which are provided by AWS and provide common functionality data. [Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/pseudo-parameter-reference.html).

| Reference Value         | Usage                                                                    | Example Returned Value                                                                          |
| ----------------------- | ------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------- |
| `AWS::AccountId`        | the accountId running the template                                       | 123456789012                                                                                    |
| `AWS::Region`           | the region the stack is in                                               | us-east-1                                                                                       |
| `AWS::StackId`          | The stackId being used                                                   | arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a11e3-aff7-50e2416294e0 |
| `AWS::StackName`        | the stack name being use                                                 | MyStack                                                                                         |
| `AWS::NotificationARNs` | list of notification Amazon Resource Names (ARNs) for the current stack. | arn:aws:sns:us-east-1:123456789012:MyTopic                                                      |
| `AWS::NoValue`          | no value                                                                 | Doesn't return a value                                                                          |

#### CloudFormation Mappings

> Mappings are **fixed** variables within your CloudFormation template. They're very handy to differentiate between different environments
> (dev vs prod), regions (AWS regions), AMI types...\
> All the values are hardcoded within the template

mappings are static, they aren't controlled by the user which runs the template. they are defined with a map name, then a toplevel key, and a secondary level key.

in this example, we set the ami to use based on the region and the machine type.

```yaml
Mappings: # section
  RegionMap: # map name
    us-east-1: # top level key
      HVM64: ami-1 # second level key and value
      HVM32: ami-2
    us-east-2:
      HVM64: ami-3
      HVM32: ami-4

Resources:
  MyEC2Instance:
    Type: AWS::EC2::Instance
    Properties:
      InstanceType: t2.micro
      ImageId: !FindInMap [RegionMap, !Ref "AWS::Region", HMV64] # map name, pseudo parameter, second level key
```

we can use the `Fn::FindInMap` function (or `!FindInMap`) to get the value.

mapping are used when we know all the options in advance, and we can determine which one to use ahead on time, they are safer to use, but don't allow as much freedom as parameters.

#### CloudFormation Outputs & Exports

> The Outputs section declares optional outputs values that we can import into other stacks (if you export them first)!
>
> - You can also view the outputs in the AWS Console or in using the AWS CLI
> - They're very useful for example if you define a network CloudFormation, and output the variables such as VPC ID and your Subnet IDs
> - It's the best way to perform some collaboration cross stack, as you let expert handle their own part of the stack

in this example, we have a Security group that we define once and export, and then we can import it and reUse it in any other stack, the exported name must be unique in the cloud account.

```yaml
Outputs:
  StackSSHSecurityGroup:
    Description: The SSH Security Group for our company
    Value: !Ref MyCompanyWideSSHSecurityGroup
    Export:
      Name: SSHSecurityGroup
```

Other stacks can reference this value with the `Fn::ImportValue` function (`!ImportValue`). the underlying stack which created the exported value can't be deleted until all stacks which reference it are removed.

```yaml
Resources:
  MyEC2Instance:
    Type: AWS::EC2::Instance
    Properties:
      InstanceType: t2.micro
      ImageId: ami-1
      AvailabilityZone: us-east-1
      SecurityGroups:
        - !ImportValue SSHSecurityGroup
```

#### CloudFormation Conditions

> Conditions are used to control the creation of
resources or outputs based on a condition
> Conditions can be whatever you want them to
be, but common ones are:
>
> - Environment (dev / test / prod)
> - AWS Region
> - Any parameter value
>
> Each condition can reference another condition,
parameter value or mapping.

```yaml
Parameters:
  EnvType:
    Type: string
    Description: environment type (sandbox, dev, prod)
    
Conditions:
  CreateProdResources: !Equals [!Ref EnvType, prod]

Resources:
  MountPoint:
    Type: AWS::EC2:VolumeAttachment
    Condition: CreateProdResources
```

#### CloudFormation Intrinsic Functions

[documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html)

| Function           | Shorthand Form  | Arguments                                   | Usage                                                                                | requires "AWS::LanguageExtensions" |
| ------------------ | --------------- | ------------------------------------------- | ------------------------------------------------------------------------------------ | ---------------------------------- |
| `Fn::Ref`          | `!Ref`          | resourceName                                | reference value in the template, either get the entire thing or the physical ID      | No                                 |
| `Fn::FindInMap`    | `!FindInMap`    | [MapName, TopLevelKey, SecondLevelKey]      | return a named value from a specific key                                             | No                                 |
| `Fn::GetAtt`       | `!GetAtt`       | resourceName.attributeName                  | get attribute from a resource                                                        | No                                 |
| `Fn::ImportValue`  | `!ImportValue`  | valueName                                   | reference an exported value from another stack                                       | No                                 |
| `Fn::And`          | `!And`          | [value1, value2]                            | logical operator                                                                     | No                                 |
| `Fn::Or`           | `!Or`           | [value1, value2]                            | logical operator                                                                     | No                                 |
| `Fn::Equals`       | `!Equals`       | [value1, value2]                            | logical operator                                                                     | No                                 |
| `Fn::Not`          | `!Not`          | value                                       | logical operator                                                                     | No                                 |
| `Fn::If`           | `!If`           | [condition, valueIfTrue, ValueIfFalse]      | ternary operator                                                                     | No                                 |
| `Fn::Join`         | `!Join`         | [delimiter, [comma limited list of values]] | join string elements together with a delimiter                                       | No                                 |
| `Fn::Split`        | `!Split`        | [delimiter, source-string]                  | split a string based on delimiter                                                    | No                                 |
| `Fn::Sub`          | `!Sub`          |                                             | substitutes                                                                          | No                                 |
| `Fn::Base64`       | `!Base64`       | stringValue                                 | transform string to base64, used in user data scripts to <cloud>EC2</cloud> machines | No                                 |
| `Fn::Cidr`         | `!Cidr`         | [ipBlock, count, cidrBits]                  | ?                                                                                    | No                                 |
| `Fn::GetAZs`       | `!GetAZs`       | region                                      | get a list of availability zones for the region                                      | No                                 |
| `Fn::Select`       | `!Select`       | [index, list]                               | select an element from a list                                                        | No                                 |
| `Fn::Length`       | `!Length`       | array                                       | number of elements in array                                                          | No                                 |
| `Fn::Transform`    |                 |                                             | apply a macro transformation                                                         | No                                 |
| `Fn::ToJsonString` | (no shortForm?) | object                                      | stringify an object or array to json form                                            | Yes                                |
| `Fn::ForEach`      |                 |                                             | Iteration Construct                                                                  | Yes                                |

</details>

### CloudFormation Options

<details>
<summary>
Additional options when creating and updating stacks.
</summary>

Other options such as stack rollbacks,

#### CloudFormation Rollbacks

When a stack creation fails, the default option is to delete all the created resources. but we can also disable the rollback and keep the resources.\
When a stack update fails, the default behavior is to rollback into the previous known working state.

A rollback can fail, this might happen if resources were manually changed. we need to find the issue and fix it, and then use <kbd>ContinueUpdateRollback</kbd> from the console or with `continue-update-rollback` api call from the cli.

we can try a demo with a bad template, this will fail since the AMI for the <cloud>EC2</cloud> resource is not a valid ami.\

```yaml
---
Parameters:
  SecurityGroupDescription:
    Description: Security Group Description
    Type: String

Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-1234
      InstanceType: t2.micro
      SecurityGroups:
        - !Ref SSHSecurityGroup
        - !Ref ServerSecurityGroup

  # an elastic IP for our instance
  MyEIP:
    Type: AWS::EC2::EIP
    Properties:
      InstanceId: !Ref MyInstance

  # our EC2 security group
  SSHSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Enable SSH access via port 22
      SecurityGroupIngress:
      - CidrIp: 0.0.0.0/0
        FromPort: 22
        IpProtocol: tcp
        ToPort: 22

  # our second EC2 security group
  ServerSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: !Ref SecurityGroupDescription
      SecurityGroupIngress:
      - IpProtocol: tcp
        FromPort: 80
        ToPort: 80
        CidrIp: 0.0.0.0/0
      - IpProtocol: tcp
        FromPort: 22
        ToPort: 22
        CidrIp: 192.168.1.1/32

```

When we update the stack, there is a section called "Stack failure Options", where we choose to perform rollback to last known state or to preserve the resources.\
if we set the option to preserve the resources, the security groups will be created, even though we failed to create the machine.

#### CloudFormation Service Role

> <cloud>IAM</cloud> role that allows CloudFormation to create/update/delete stack resources on your behalf.\
> Give ability to users to create/update/delete the stack resources even if they don't have permissions to work with the resources in the stack
>
> Use cases:
>
> - You want to achieve the least privilege principle
> - But you don't want to give the user all the required permissions to create the stack resources
>
> User (which creates the template) must have `iam:PassRole` permissions

in the <cloud>IAM</cloud> service, we create a role, choose the trusted entity as <cloud>CloudFormation</cloud>, and we give it the capabilities for the resources we want it to create (such as <cloud>S3</cloud>). and when we create a stack, we can define which role will create the resource (rather than using the user role).

#### CloudFormation Capabilities

capabilities that we need to give cloudFormation if we want it to create <cloud>IAM</cloud> resources.

> `CAPABILITY_NAMED_IAM` and `CAPABILITY_IAM`
>
> - Necessary to enable when you CloudFormation template is creating or updating IAM resources (IAM User, Role, Group, Policy, Access Keys, Instance Profile... )
> - Specify `CAPABILITY_NAMED_IAM` if the resources are named.
>
> `CAPABILITY_AUTO_EXPAND`
>
> - Necessary when your CloudFormation template includes Macros or Nested Stacks (stacks within stacks) to perform dynamic transformations
> - You're acknowledging that your template may change before deploying
>
> `InsufficientCapabilitiesException` Exception that will be thrown by CloudFormation if the capabilities haven't been acknowledged when deploying a template (security measure).

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Description: An example CloudFormation that requires CAPABILITY_IAM and CAPABILITY_NAMED_IAM

Resources:
  MyCustomNamedRole:
    Type: AWS::IAM::Role
    Properties:
      RoleName: MyCustomRoleName
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service: [ec2.amazonaws.com]
            Action: ['sts:AssumeRole']
      Path: "/"
      ManagedPolicyArns:
        - arn:aws:iam::aws:policy/AmazonEC2FullAccess
      Policies:
        - PolicyName: MyPolicy
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Effect: Allow
                Action: 's3:*'
                Resource: '*'

Outputs:
  RoleArn:
    Description: The ARN of the created IAM Role
    Value: !GetAtt MyCustomNamedRole.Arn
```

if we try creating a stack with this template, we will see a prompt the requires us to confirm that we want to have this stack create IAM resources

#### CloudFormation Deletion Policy

> DeletionPolicy Controls what happens when the
CloudFormation template is deleted or when a resource is removed from a CloudFormation template.\
> Extra safety measure to preserve and backup resources.

we attach this policy to resources in the template, we don't need to specify the default behavior.

```yaml
Resources:
  MyResource:
  Type: AWS::DynamoDB:Table
  Properties:
    TableName: someName
  DeletionPolicy: Retain # default is delete
```

the default option is "Delete", which removes the resource, this won't work on <cloud>S3</cloud> buckets unless they are empty!.\
The "Retain" option preserves the resource if the stack is deleted. this works for all objects.\
The "Snapshot" policy will take a final snapshot of the object before deleting it. it is supported in some resources:

- EBS Volume
- ElastiCache Cluster
- ElastiCache ReplicationGroup
- RDS DBInstance
- RDS DBCluster
- Redshift Cluster
- Neptune DBCluster
- DocumentDB DBCluster
- and others...

if we skip deleting resources, it's up to us to delete them manually.

#### CloudFormation Stack Policy

> During a CloudFormation Stack update, all update actions are allowed on all resources (default).\
> A Stack Policy is a JSON document that defines the update actions that are allowed on specific resources during Stack updates.\
> Protect resources from unintentional updates
> When you set a Stack Policy, all resources in the Stack are protected by default, Specify an explicit ALLOW for the resources you want to be allowed to be updated.

#### CloudFormation Termination Protection

an option to prevent deleteions of stacks, we can control the termination protection, it must be disable before deleting the stack, and the user must have the permissions to disable it.

</details>

### CloudFormation Custom Resources

<details>
<summary>
Defining custom resource or logic
</summary>

> - define resources not yet supported by CloudFormation.
> - define custom provisioning logic for resources can that be outside of CloudFormation (on-premises resources, 3rd party resources,...)
> - have custom scripts run during create / update / delete through Lambda functions (running a Lambda function to empty an S3 bucket before being deleted)
>
> Defined in the template using `AWS::CloudFormation::CustomResource` or
`Custom::MyCustomResourceTypeName` (recommended).\
> Backed by a Lambda function (most common) or an SNS topic.

we need a service token, either a <cloud>Lambda</cloud> function or a <cloud>SNS</cloud> arn

```yaml
Resources:
  MyCustomResourceUsingLambda:
    Type: Custom::MyLambdaResource
    Properties:
      ServiceToken: some arn
      # input values (optional)
      ExampleProperty: "example value"
```

a common example is to set a custom resource to empty a bucket by using a lambda.

</details>

### CloudFormation Dynamic References

<details>
<summary>
Get Values from additional source
</summary>

> Reference external values stored in Systems Manager Parameter Store and Secrets Manager within CloudFormation templates. CloudFormation retrieves the value of the specified reference during create/update/delete operations.\
> For example: retrieve RDS DB Instance master password from Secrets Manager.
>
> Supports:
>
> - ssm - for plaintext values stored in SSM Parameter Store
> - ssm-secure - for secure strings stored in SSM Parameter Store
> - secretsmanager - for secret values stored in Secrets Manager

the syntax to use in the template is `{{resolve:service-name:reference-key}}`.

- ssm - `{{resolve:ssm:parameter-name:version}}`
- ssm-secure - `{{resolve:ssm-secure:parameter-name:version}}`
- secretsmanager - `{{resolve:secretsmanager::secret-id:secret-sting:json-key:version-stage:version-id}}`

```yaml
Resources:
  MyResource:
    Type: AWS::S3:Bucket
    Properties:
      AccessControl: `{{resolve::ssm:S3AccessControl:2}}`
```

somethings to note:

if we create an <cloud>RDS</cloud> resource, we can set the "ManageMasterUserPassword" property to "true", and then we could output it as a secret.

the other option is to create the secret as a cloudFormation resource `AWS::SecretManager::Secret`, and then use the dynamic reference syntax (`{{resolve:}}`) to reference it. we also use `AWS::SecretManager::SecretTargetAttachment` to link the secret together with the RDS and control it for rotations.

</details>

### CloudFormation User Data

<details>
<summary>
Passing User Data Scripts and Helper Scripts
</summary>

passing user Data to <cloud>EC2</cloud> instance. the important thing to remember is to pass the script data in base64 using `Fn::Base64` (or `!Base64`). the script log will be store in "/var/log/cloud-init-output.log" file.

we start with this example

```yaml
---
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-0a3c3a20c09d6f377
      InstanceType: t2.micro
      SecurityGroups:
        - !Ref SSHSecurityGroup
      # we install our web server with user data
      UserData: 
        Fn::Base64: |
          #!/bin/bash -xe
          dnf update -y
          dnf install -y httpd
          systemctl start httpd
          systemctl enable httpd
          echo "<h1>Hello World from user data</h1>" > /var/www/html/index.html

  # our EC2 security group
  SSHSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: SSH and HTTP
      SecurityGroupIngress:
      - CidrIp: 0.0.0.0/0
        FromPort: 22
        IpProtocol: tcp
        ToPort: 22
      - CidrIp: 0.0.0.0/0
        FromPort: 80
        IpProtocol: tcp
        ToPort: 80
```

however, something important to know, is that even if we had problem with script running, the stack would still result in success.

also, this works for small scripts, but what if we had a large instance configuration? and what if we wanted to evolve the script without affecting the ec2 machine (terminating and restarting it).\
we cane use some helper scripts to solver those problems.

[documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-helper-scripts-reference.html)

- `cfn-init`
- `cfn-signal`
- `cfn-get-metadate`
- `cfn-hup`

#### `cfn-init` Helper Function

under the <cloud>EC2</cloud> resource, there is the `AWS::CloudForamation::Init` metdadate property.

> - Packages: used to download and install pre-packaged apps and components on Linux/Windows (ex. MySQL, PHP, etc...)
> - Groups: define user groups • Users: define users, and which group they belong to
> - Sources: download files and archives and place them on the EC2 instance
> - Files: create files on the EC2 instance, using inline or can be pulled from a URL
> - Commands: run a series of commands
> - Services: launch a list of sysvinit services

we run the `cfn-init` command in the user data script, but we keep the script to only the minimal.

> the `cfn-init` command is used to retrieve and interpret the resource metadata, installing packages, creating files and starting services.\
> With the cfn-init script, it helps make complex
EC2 configurations readable. The EC2 instance will query the CloudFormation service to get init data.
`AWS::CloudFormation::Init` must be in the Metadata of a resource.

`cfn-init` flags:

- `-s`,`--stack` - stack
- `-r`, `--resource` - resource name
- `--region` - AWS region

```yaml
---
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-0a3c3a20c09d6f377
      InstanceType: t2.micro
      SecurityGroups:
        - !Ref SSHSecurityGroup
      # we install our web server with user data
      UserData: 
        Fn::Base64:
          !Sub |
            #!/bin/bash -xe
            # Get the latest CloudFormation package
            dnf update -y aws-cfn-bootstrap
            # Start cfn-init
            /opt/aws/bin/cfn-init -s ${AWS::StackId} -r MyInstance --region ${AWS::Region} || error_exit 'Failed to run cfn-init'
    Metadata:
      Comment: Install a simple Apache HTTP page
      AWS::CloudFormation::Init:
        config:
          packages:
            yum:
              httpd: []
          files:
            "/var/www/html/index.html":
              content: |
                <h1>Hello World from EC2 instance!</h1>
                <p>This was created using cfn-init</p>
              mode: '000644'
          commands:
            hello:
              command: "echo 'hello world'"
          services:
            sysvinit:
              httpd:
                enabled: 'true'
                ensureRunning: 'true'

  # our EC2 security group
  SSHSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: SSH and HTTP
      SecurityGroupIngress:
      - CidrIp: 0.0.0.0/0
        FromPort: 22
        IpProtocol: tcp
        ToPort: 22
      - CidrIp: 0.0.0.0/0
        FromPort: 80
        IpProtocol: tcp
        ToPort: 80
```

#### `cfn-signal` & Wait Condition Helper Function

to make sure the script worked properly, we we use the `cfn-signal` function right after the `cfn-init`. this will send a signal that our template will wait on. so we need to define a "Wait Condition".

this is another resource, which waits for a signal before creation, and if it doesn't receive the signal, the resource will fail and the stack won't succeed.

a success code is zero, if it recives a code other than zero it's counted as a failure, or if it doesn't receive anything until the timeout expires.

```yaml
---
Resources:
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      AvailabilityZone: us-east-1a
      ImageId: ami-0a3c3a20c09d6f377
      InstanceType: t2.micro
      SecurityGroups:
        - !Ref SSHSecurityGroup
      # we install our web server with user data
      UserData: 
        Fn::Base64:
          !Sub |
            #!/bin/bash -x
            # Get the latest CloudFormation package
            dnf update -y aws-cfn-bootstrap
            # Initialize EC2 Instance
            /opt/aws/bin/cfn-init -v --stack ${AWS::StackName} --resource MyInstance --region ${AWS::Region}
            # Get result of last command
            INIT_STATUS=$?
            # send result back using cfn-signal
            /opt/aws/bin/cfn-signal -e $INIT_STATUS --stack ${AWS::StackName} --resource SampleWaitCondition --region ${AWS::Region}
            # exit the script
            exit $INIT_STATUS
    Metadata:
      Comment: Install a simple Apache HTTP page
      AWS::CloudFormation::Init:
        config:
          packages:
            yum:
              httpd: []
          files:
            "/var/www/html/index.html":
              content: |
                <h1>Hello World from EC2 instance!</h1>
                <p>This was created using cfn-init</p>
              mode: '000644'
          commands:
            hello:
              command: "echo 'hello world'"
          services:
            sysvinit:
              httpd:
                enabled: 'true'
                ensureRunning: 'true'

  SampleWaitCondition:
    CreationPolicy:
      ResourceSignal:
        Timeout: PT3M
        Count: 1
    Type: AWS::CloudFormation::WaitCondition
```

#### `cfn-signal` Helper Function Failures

common reasons for this not to work is if the AMI doesn't have the helper scripts installed, this can be verified by not rolling back the machine and checking the logs.

</details>

### Nested Stacks

<details>
<summary>
Reusable stack components
</summary>

Nested stacks are stacks inside other stacks.

> Nested stacks are stacks as part of other stacks.  They allow you to isolate repeated patterns/ common components in separate stacks and call them from other stacks.\
>
> - Load Balancer configuration that is re-used
> - Security Group that is re-used
>
> Nested stacks are considered best practice.- To update a nested stack, always update the parent (root stack).\
> Nested stacks can have nested stacks themselves!

They shouldn't be confused with Cross-Stack references, which are independent of the parent stack and export value to be used by many other stacks.

a nested stack needs the templateURL, and whatever parameters we need to pass to it.

```yaml
Resources:
  MyKeyPair:
    Type: AWS::EC2::KeyPair
    Properties:
      KeyName: DemoKeyPair
      KeyType: rsa

  myStack:
    Type: AWS::CloudFormation::Stack
    Properties:
      TemplateURL: https://s3.amazonaws.com/cloudformation-templates-us-east-1/LAMP_Single_Instance.template
      Parameters:
        KeyName: !Ref MyKeyPair
        DBName: "myDb"
        DBUser: "user"
        InstanceType: t2.micro
        SSHLocation: "0.0.0.0/0"

Outputs:
  StackRef:
    Value: !Ref myStack
  OutputFromNestedStack:
    Value: !GetAtt myStack.Outputs.WebsiteURL
```

#### "Depends On" Property

> Specify that the creation of a specific resource follows another.
>
> - When added to a resource, that resource is created only after the creation of the resource specified in the DependsOn attribute
> - Applied automatically when using `!Ref` and `!GetAtt`
> - Use with any resource

</details>

### CloudFormation - StackSets

<details>
<summary>
Deploying stacks across multiple accounts and regions.
</summary>

administrator only operation.

> Create, update, or delete stacks across multiple accounts and regions with a single operation/template
>
> - Administrator account to create StackSets
> - Target accounts to create, update, delete stack instances from StackSets
> - When you update a stack set, all associated stack instances are updated throughout all accounts and regions
> - Can be applied into all accounts of an AWS organizations

for normal, personal accounts, we use self-managed permessions, which requires manual work on both the source and the target accounts.

> Self-managed Permissions
>
> - Create the IAM roles (with established trusted relationship) in both administrator and target accounts.
> - Deploy to any target account in which you have permissions to create IAM role

if we are using AWS organization, we can use service-managed permission, this is a more powerful option.
> Service-managed Permissions
>
> - Deploy to accounts managed by AWS Organizations
> - StackSets create the IAM roles on your behalf (enable trusted access with AWS Organizations)
> - Must enable all features in AWS Organizations
> - Ability to deploy to accounts added to your organization in the future (Automatic Deployments)

#### StackSets Demo

we first create the two roles, one in each account, the roles will have the AWSCloudFormagonStackSetAdministrationRole and the AWSCloudFormationStackSetExecutionRole, both are assumable by cloud formation, and the administrator role can assume the target execution role.

we set the template to only allow the specific account to assume the execution role.

under <cloud>CloudFormation</cloud>, we can create stackSets, there are many parameters, conditions, etc...

we choose to deploy stacks in different accounts, regions, and across the organization, we can choose how many accounts are affected concurrently, and the failure tolerance.

we can detach stackSets from a target account, but without deleting the resources, or we can delete the stackSet and remove the created resources. then we can delete the parent stackSet.

</details>

### CloudFormation - Troubleshooting

<details>
<summary>
Failures and possible fixes.
</summary>

> **DELETE_FAILED**
>
> - Some resources must be emptied before deleting, such as S3 buckets
> - Use Custom Resources with Lambda functions to automate some actions
> - Security Groups cannot be deleted until all EC2 instances in the group are gone
> - Think about using DeletionPolicy=Retain to skip deletions
>
> **UPDATE_ROLLBACK_FAILED**
>
> - Can be caused by resources changed outside of CloudFormation, insufficient permissions, Auto Scaling Group that doesn't receive enough signals…
> - Manually fix the error and then `ContinueUpdateRollback`
>
> A stack operation failed, and the stack instance status is `OUTDATED`.
>
> - Insufficient permissions in a target account for creating resources that are specified in your template.
> - The template could be trying to create global resources that must be unique but aren't, such as S3 buckets
> - The administrator account does not have a trust relationship with the target account
> - Reached a limit or a quota in the target account (too many resources)

</details>

</details>

## Lambda for SysOps

<details>
<summary>
Serverless code execution.
</summary>

<cloud>EC2</cloud> machine have to be continuesly running, even when there's no work. we also need to manage the configuration and scaling. with <cloud>Lambda</cloud> functions, we don't need to manage the scaling, and when there's no work to do, we don't use any compute power (so we don't pay).

Lambda functions are also integrated with many services, and it has built-in runtimes for many languages, with other languages as extended runtime support, or by using container images (if they implement the lambda runtime API).

integrations:

- <cloud>API Gateway</cloud> - REST api to invoke lambda
- <cloud>Kinesis</cloud> - transform data
- <cloud>DynamoDB</cloud> - trigger lambda when something changes
- <cloud>S3</cloud> - trigger lambda when object happens
- <cloud>CloudFront</cloud> - lambda @ edge
- <cloud>CloudWatch EventBridge</cloud> - other kinds of events
- <cloud>CloudWatch Logs</cloud> - stream logs to somewhere else
- <cloud>SNS</cloud> - react to topic
- <cloud>SQS</cloud> - process messages in queues
- <cloud>Cognito</cloud> - react to log-in events

we can navigate to the lambda service page and see examples of different runtime langauges, different sources and how it scales.

we can can create a function starting with a blueprint, we need an <cloud>IAM</cloud> role for the lambda to use, and we can test the function by editing the event.

we can check the configuration (memory, timeout, environment variables), and monitor the lambda, we can see what our lambda can do by looking at the permission tab.

### Lambda Event Integrations

<details>
<summary>
Triggering Lambda function from Events
</summary>

#### Lambda & CloudWatch Events / EventBridge

intergration with <cloud>EventBridge</cloud> - we can either have a serverless cron rule which runs on a schedule, or a <cloud>CodePipeline</cloud> event that triggers when a state changes.

we create a demo function, and then navigate to the eventBridge service, select <cloud>Create Rule</cloud>. we can set a scheduled rule (cron), or a rule based on event pattern. there is a new way to set cron rules (<cloud>EventBridge Scheduler</cloud>), we need to set the target (Lambda in our case, but can be other services), and make sure the permissions fit.

#### Lambda & S3 Event Notifications

we can also integrate <cloud>S3</cloud> event notification with out lambda, such as object creation. we can set the event to create <cloud>SQS</cloud>, <cloud>SNS</cloud> or directly invoke <cloud>Lambda</cloud> asynchronously. we can fine tune which object trigger the events based on the prefix and suffix.

in the demo, we create a bucket, set the notification rule to invoke a lambda.

</details>

### Lambda Permissions - IAM Roles & Resource Policies

<details>
<summary>
Execution roles and permissions
</summary>

The lambda has an <cloud>IAM</cloud> execution role, which grants it permissions to other aws services and resources. there are some managed policies for common use cases.\
When we have an event source mapping which triggers the lambda function, this should be defined either in the role that invokes the lambda, or in the Lambda resource/access policy.

</details>

### Lambda Monitoring & X-Ray Tracing

<details>
<summary>
Different methods of monitoring Lambda functions
</summary>

- <cloud>CloudWatch Logs</cloud> - write to logs, part of the basic Lambda Role
- <cloud>CloudWatch Metrics</cloud> - information about innvocations, durations, errors, etc...
- <cloud>XRayTracing</cloud> - trace the flow of the lambda, requires using the x-ray daemon ("active tracing" option in the configuration), there are several needed environment variables.

> - Invocations - number of times your function is invoked (success/failure)
> - Duration - amount of time your function spends processing an event
> - Errors - number of invocations that result in a function error
> - Throttles - number of invocation requests that are throttled (no concurrency available)
> - DeadLetterErrors - number of times Lambda failed to send an event to a DLQ (async invocations)
> - IteratorAge - time between when a Stream receives a record and when the Event Source Mapping sends the event to the function (for Event Source Mapping that reads from Stream)
> - ConcurrentExecutions - number of function instances that are processing event

we can see them in the dashboard, and we can build alarms on top of those metrics. like having an alarm for number of invocations, errors or throtelling. we can use <cloud>CloudWatch</cloud> Logs Insight feature to query the logs, we can also get aggregated insights about the lambda at the system level. we need to set it as a separate lambda layer.
</details>

### Lambda Function Performance

<details>
<summary>
Getting more compute power
</summary>

RAM - memory can be as low 128MB and go to 10GB. CPU is linked to the memory, at about 1792Mb we get an additional vCPU unit. to get value of this compute power increase, we need to use multi-threading in our code.

default timeout is three seconds, we can set it up to 15 minutes maximum.

when we run the lambda, it first creates an execution context, which is a temporary runtime environment that initializes any external dependencies of the lambda code. this execution context outlives a single execution, so it can be re-used if we have another invocation. we can use this behavior to our advantage, by opening database connections, taking resources as needed, and doing other work that can persist across invocations.\
the "/tmp" folder persists across invocations, and we have 10GB of disk space to work with. if we want to encrypt it, we need to manually use <cloud>KMS</cloud> data keys

</details>

### Lambda Concurrency

<details>
<summary>
Cold and Warm starts
</summary>

there is a limit of 1000 concurrent executions for all lambda functions. this stops our lambda from scaling out of control. we can set "reserved concurrency" to limit the number to something lower, or request an increase from AWS if we need more. we need to be careful that one lambda doesn't hog all the concurrent executions and throttles our system. there is a different behavior for synchronous and asynchronous invocations.

the first time we run a lambda, the execution runtime is created, this can take longer - we have a "cold start", as long as this execution context exists, lambda invocations will be "warm starts", and will be faster.\
If we want to always have the better performance, we can set a number of "provisioned concurrency", which will ensure some execution contexts remain, even if no one has invoked the lambda in a while. we can set auto-scaling to manage this behavior.
</details>

</details>

## EC2 Storage and Data Management - EBS and EFS

<details>
<summary>
Block Storage and File Storage for EC2 instances.
</summary>

<cloud>EBS</cloud> - Elastic Block Store, <cloud>EFS</cloud> - Elastic File Storage.

> - An EBS (Elastic Block Store) Volume is a network drive you can attach to your instances while they run.
> - It allows your instances to persist data, even after their termination.
> - They can normally only be mounted to one instance at a time, unless using "Multi-attach" advanced feature.
> - They are bound to a specific availability zone.

diving a bit deeper:

> - It's a network drive (i.e. not a physical drive)
>   - It uses the network to communicate the instance, which means there might be a bit of latency
>   - It can be _detached_ from an EC2 instance and _attached_ to another one quickly
> - It's locked to an Availability Zone (AZ)
>   - An EBS Volume in us-east-1a cannot be attached to us-east-1b
>   - To move a volume across, you first need to _snapshot_ it
> - Have a provisioned capacity (size in GBs, and IOPS)
>   - You get billed for all the provisioned capacity
>   - You can increase the capacity of the drive over time

We can create EBS volumes without having them attached to a machine. an <cloud>EC2</cloud> machine can attach more than one volume.\
When we create the volumes directly from the machine, we can control if it's deleted when the machine is terminated ("delete on termination"). the default behavior is to delete the root volume, but not the others. there is also an option to encrypt the volume using <cloud>KMS</cloud>.

(hands on)\
in our <cloud>EC2</cloud> machine, we can look at the "Storage" tab, there we can see the root block device and check the details about it. we can also <kbd>Create volume</kbd> to create another one. we need to choose the volume type and the size in GB, we must select the Availability Zone for it. we can also choose to create the volume from an existing snapshot.\
The newly created volume is not attached to any instance, so we can click <kbd>Actions</kbd> and <kbd>Attach</kbd> to an instance in that Availability Zone. from the machine, we would need to run some commands to mount the volume so it would be usable.\
We can terminate our instance and observe how the root volume is deleted, but not the volume we created.

### EC2 Instance Store

<details>
<summary>
Better Performance with an attached disk space
</summary>

A non-network store option is <cloud>EC2</cloud> Instance store. it is not a network drive, and is ephemeral, can't be easily moved across machines (it's possible to do so, but requires some steps), and the data can be lost if the machine or hardware fails. but it does offer better performance.

a good use case for instance store is for buffer, cache, scracth data and temporary content, while using network volumes for data that needs to persist.

</details>

### EBS Volume Types Deep Dive

<details>
<summary>
Understanding The Differences Between Volume Types.
</summary>

EBS volume differ in type of disk (SSD vs HDD), size (measured in GB), throughput and IOPS (I/O operation per second).

> - gp2 / gp3 (SSD): General purpose SSD volume that balances price and performance for a wide variety of workloads.
> - io1 / io2 Block Express (SSD): Highest-performance SSD volume for mission-critical low-latency or high-throughput workloads.
> - also called "provisioned iops" volumes
> - st1 (HDD): Low cost HDD volume designed for frequently accessed, throughput- intensive workloads.
> - sc1 (HDD): Lowest cost HDD volume designed for less frequently accessed workloads.
>
> Only gp2/gp3 and io1/io2 Block Express can be used as boot (root) volumes.

We use the _gp2/gp3_ SSD volumes for most cases, when we have mission critical workloads we can use _io1/io2_ SSD devices for low latency. when we want lower cost we can switch to the HDD devices, _st1_ will give us good performance for frequently accessed data, and can have a very high max throughput. _sc1_ volumes have the lowest cost, and are most fit for storing data that is not frequently accessed.

| _                    | gp2                           | gp3                                  | io1                                        | io2                                | st1 (throughput optimized)           | sc1 (cold storage)    |
| -------------------- | ----------------------------- | ------------------------------------ | ------------------------------------------ | ---------------------------------- | ------------------------------------ | --------------------- |
| Use Case             | General                       | General                              | Critical workloads, Sustained IOPS         | Critical workloads, Sustained IOPS | Frequently Accessed, high throughput | Infrequently Accessed |
| Device               | SSD                           | SSD                                  | SSD                                        | SSD                                | HDD                                  | HDD                   |
| Boot Volume?         | Yes                           | Yes                                  | Yes                                        | Yes                                | No                                   | No                    |
| Cost                 | Normal                        | Normal                               | Higher                                     | Higher                             | Lower                                | Lowest                |
| Size                 | 1 GiB - 16 TiB                | 1 GiB - 16 TiB                       | 4 GiB - 16 TiB                             | 4 GiB - 64 TiB                     | 125 Gib - 16 TiB                     | 125 Gib - 16 TiB      |
| IOPS                 | 3000-16000 (3 Iops per 1 GiB) | 3000-16000                           | max is 32000 normally, 64000 for Nitro EC2 | max is 256000                      | max is 500                           | max is 250            |
| Throughput           | linked with IOps              | 125-1000 MiB/s (independent of IOPS) | (independent of IOPS)                      | linked with IOps                   | throughput optimized, 500 MiB/s      | nax is 250 MiB/s      |
| Multi Attach Support | No                            | No                                   | Yes                                        | Yes                                | No                                   | No                    |

</details>

### EBS Operations

<details>
<summary>
Different Operations For EBS Volumes.
</summary>

#### EBS Multi Attach

Attaching the same EBS volume to multiple machines in the same Availability Zone. all machines have full read/write permissions, the use case is when we must support concurrent write operations, or when we have clustered application (such as Teradata). Only in the provision iops family (io1/io2). up to 16 instances can be attach the volume at the same time, and they must use a file system that is cluster-aware (**XFS, EXT4 can't be used**).

#### Volume Resizing

Increasing the size of the EBS volume. this is available in all volume types. in io1 devices it's also possible to increase the IOPS. after this is done, we need to "re-partition" the drive to make the machine aware of the increased size. the volume might go into an "optimization" phase (it will still be usable).\
This is a one way operation, we **cannot decrease** the size, we would have to create a smaller volume and migrate the data.

we can launch an EC2 instance using the default settings, and now we can connect to it, and we verify the disk space we have and see we have 8Gb of disk space

```sh
lsblk
df -h
```

now we select the volume click <kbd>Actions</kbd> and <kbd>Modify Volume</kbd>, and we can increase the size. we connect back to the machine and run some commands

```sh
lsblk
df -h
sudo growpart /dev/xvda 1
lsblk
df -h
```

at this stage, we can either run more commands, but it's easier to simply restart the instance.

#### Snapshots

we use snapshots to have a persistent backup of our EBS volumes.

> - Make a backup (snapshot) of your EBS volume at a point in time
> - Not necessary to detach volume to do snapshot, but recommended
> - Can copy snapshots across AZ or Region

there is also an AWS service to manage snapshots for us - Amazon Data Lifecycle Manager (DLM)

> - Automate the creation, retention, and deletion of EBS snapshots and EBS-backed AMIs
> - Schedule backups, cross-account snapshot copies, delete outdated backups, etc...
> - Uses resource tags to identify the resources (EC2 instances, EBS volumes)
> - Can't be used to manage snapshots/AMIs created outside DLM
> - Can't be used to manage instance-store backed AMIs
>
> Fast Snapshot Restore (FSR)
>
> - EBS Snapshots are stored in S3
> - By default, there's a latency of I/O operations the first time each block is accessed (block must be pulled from S3)
> - Solution: force the initialization of the entire volume (using the `dd` or `fio` command), or you can enable FSR
> - FSR helps you to create a volume from a snapshot that is fully initialized at creation (no I/O latency)
> - Enabled for a snapshot in a particular AZ (billed per minute - very expensive $$$)
> - Can be enabled on snapshots created by Data Lifecycle Manager

(normal behavior is lazy getting blocks, FSR keeps a warm copy of it)

we can achieve snapshots by moving them to a cheaper storage, they will take more time to restore. we can also setup rules for a snapshot recycle bin, which can help recover from accidental deletion, we can set a retention period (minimum 1 day, maximum one year).

(hands on)

we navigate to one of our volumes, and click <kbd>Actions</kbd> and <kbd>Create Snapshot</kbd>. if the volume is encrypted, the snapshot will also be encrypted. with this snapshot created, we can create volumes out of it, and we are not limited to Availability Zone the volume was in, we can copy it to another region if we want. this is how we migrate data between Availability Zones or regions. we can also add encryption when we copy it to another region or create a volume from it.\
another option is <kbd>Actions</kbd> and <cloud>Manage fast snapshot restore</cloud>, and we enable it for each AZ. this will incur heavy charges. we can <kbd>Achieve</kbd> the snapshot to move it to a cheaper storage tier.\
Under the "lifecycle Manager" secion, we can crate a new lifecycle policy for either EBS snapshots, EBS-back AMI or Cross Account Copy events. we choose the snapshot policy. and we define the target (volume or instances), the target tags (which resources will be managed), the <cloud>IAM</cloud> role, and we set the schedule and how many snapshots are stored. we can also set additional tags for the created snapshots, set Fast Snapshot Restore (very pricy!), and manage cross-region copy and cross-account sharing. a policy can be enabled or disabled.\
we can also navigate to the "Recycle Bin" service and <kbd>Create Retention Rules</kbd>, we choose the target type, the target tags and set the retention period. we can add a rule lock, which adds a delay of time in which the rule can't be deleted, even after being unlocked. now when we delete a snapshot, it is moved to the bin and can be restored.

#### Volume Migration

> EBS Volumes are only locked to a specific AZ.To migrate it to a different AZ (or region):
>
> - Snapshot the volume
> - (optional) Copy the volume to a different region
> - Create a volume from the snapshot in the AZ of your choice

#### Volume Encryption

> - When you create an encrypted EBS volume, you get the following:
>   - Data at rest is encrypted inside the volume
>   - All the data in flight moving between the instance and the volume is encrypted
>   - All snapshots are encrypted
>   - All volumes created from the snapshot
> - Encryption and decryption are handled transparently (you have nothing to do)
> - Encryption has a minimal impact on latency
> - EBS Encryption leverages keys from <cloud>KMS</cloud> (AES-256)
> - Copying an unencrypted snapshot allows encryption
> - Snapshots of encrypted volumes are encrypted
>
> Encryption: encrypt an unencrypted EBS volume
>
> - Create an EBS snapshot of the volume
> - Encrypt the EBS snapshot (using copy)
> - Create new ebs volume from the snapshot ( the volume will also be encrypted)
> - Now you can attach the encrypted volume to the original instance
>
</details>

### Amazon EFS - Elastic File System

<details>
<summary>
Managed network file system for linux machines.
</summary>

<cloud>EFS</cloud> - Elastic File Storage

> Managed NFS (network file system) that can be mounted on many EC2 across different Availability Zones
>
> - Highly available, scalable, expensive (3x gp2), pay per use.
> - Use cases: content management, web serving, data sharing, Wordpress
> - Uses NFSv4.1 protocol
> - Uses security group to control access to EFS
> - Compatible with Linux based AMI (**not Windows**)
> - Encryption at rest using KMS
> - POSIX file system (~Linux) that has a standard file API
> - File system scales automatically, pay-per-use, no capacity planning

EFS is highly scalable, allows for thousends of concurrent NFS clients, and more than 10 GB/s throughput. it grows automatically in size, and can reach Petabyte scale of storage, without the user having to manage anything.

it has two performance modes, it must be set at creation.

- General Purpose (default) - latency-sensitive cases (web servers CMS)
- Max I/O - highter latency, but better parallelization for big data or media processing.

> Throughput Modes:
>
> - Bursting - start with (50MiB/s for 1 TB storage) + burst of up to 100MiB/s
> - Provisioned - set your throughput regardless of storage size (ex: 1 GiB/s for 1 TB storage)
> - Elastic - automatically scales throughput up or down based on your workloads
>   - Up to 3GiB/s for reads and 1GiB/s for writes
>   - Used for unpredictable workloads

supports storage classes and lifecycle management using lifecycle policies to move files between tiers.
>
> - Standard: for frequently accessed files
> - Infrequent access (EFS-IA): cost to retrieve files, lower price to store.
> - Archive: rarely accessed data (few times each year), 50% cheaper.

Default is multi-AZ mode, but can also be set to a single Availability Zone (backup enabled by default), works with EFS-IA tier and is cheaper.

(hands on)\
In the <cloud>EFS</cloud> Service, <kbd>Create File System</kbd>, choose the <cloud>VPC</cloud>, select the file system type (Regional multi-az or One Zone), enable or disable backup, control the file lifecycle (transition between tiers) and enable encryption at rest (with more settings). we can set the performance settings (elastic, provisoned, bursting).

- Bursting - scale with the amount of storage
- Elastic - scale with usage
- Provisioned - manually set throughput

if we choose bursing or provisioned, we can select max I/O to get better throughput at the cost of higher latency.

at the network settings, we choose the VPC, and we can mount the file system to different Availability Zones in it, we can choose the subnet and security group. we need a security group that allows for NFS connection (port 2049), but we can have the console manage that for us.\
We can also control the file system policies to prevent root access, enforce read-only as default, enforce in-transit encryption and other stuff.

now we create a new <cloud>EC2</cloud> machine in the same VPC and subnet, and under storage, we can add the <cloud>EFS</cloud> as a file system. we have EFS and FSx as options, we can set the mount point and have aws set the permssions and the correct user script to mount it.

We can connect to two instances, create a file in the shared file system from one machine, and see it from the other machine.

#### EFS vs EBS

- EFS - Network file system
- EBS - Network disk device
- Instance Store - Attached disk device

<cloud>EFS</cloud> is file storage, <cloud>EBS</cloud> is block storage. EFS can support 1000s of instances across Availability Zones by default, EBS usually has one machine, or up to 16 machines when using multi-attach (io1, io2) and even then, all machine must be in the same Availability Zone.

EFS only supports linux machine (unless using <cloud>FSx</cloud>), and it has a higher cost. we can use storage tiers for cost savings.\
EFS scales automatically, and can reach petabytes of data.

</details>

### EFS Access Points

<details>
<summary>
Manage Application Accesses
</summary>

Easily manage applications access to NFS environments, separate what each user can access, so not everybody can access every file

> - Enforce a POSIX user and group to use when accessing the file system
> - Restrict access to a directory within the file system and optionally specify a different root directory
> - Can restrict access from NFS clients using IAM policies

in the <cloud>EFS</cloud> service, we navigate to the "Access Point" section, <kbd>Create Access Point</kbd>, select the file system, choose the root directory for this point, and we set the Posix user identity (user id, group id), control what can be done in the root directory.\ Once this is created, we get an arn and the appropriate mounting command to be run on the EC2 machine.

```sh
sudo mount -t efs -o tls, accesspoint=<accesspoint identifier> <filesystem-identifier>:/ efs
```

</details>

### EFS - Operations

<details>
<summary>
What can we do with EFS.
</summary>

> Operations that can be done in place:
>
> - Lifecycle Policy (enable IA or change IA settings)
> - Throughput Mode and Provisioned Throughput Numbers
> - EFS Access Points
>  
> Operations that require a migration using **DataSync** (replicates all file attributes and metadata)
>
> - Migration to encrypted EFS
> - Performance Mode (e.g. Max IO)

</details>

### EFS - CloudWatch Metrics

<details>
<summary>
CloudWatch EFS Metrics.
</summary>

> - `PercentIOLimit` - How close the file system reaching the I/O limit (General Purpose)
>   - If at 100%, move to Max I/O (migration with **DataSync**)
> - `BurstCreditBalance` - The number of burst credits the file system can use to achieve higher throughput levels
> - `StorageBytes` - File system's size in bytes (15 minutes interval)
>   - Dimensions: Standard, IA, Total (Standard + IA)

</details>

</details>

## Amazon S3 And Athena

<details>
<summary>
Object Storage, infinity scaling.
</summary>

one of the core services of AWS, used for many things:

- Backup and storage
- Disaster Recovery
- Archive
- Hybrid Cloud Storage
- Application hosting
- Data lakes & big data analytics
- Software delivery
- Static website

<cloud>S3</cloud> uses buckets, which are top-level objects, they must have a globally unique name (across all aws accounts). they are defined in the region level, even if they sometimes behave as if they are global.\
There are some restrictions on how they can be named.

Inside buckets, there are objects (files), the objects have a key, which is the full path. we can define keys with prefix that look like directories (folder), but basic buckets don't truly support the concept of folders, it's just that the UI behaves as if they do.

the objects themselves also have values - the content of the files. max object size is 5TB (5000 GB), for files which are larger than 5GB, uploading requires using "multi-part upload".\
Objects also have metadata - such as custom tags and version.

in the <cloud>S3</cloud> service, we cna <kbd>Create a Bucket</kbd>. the bucket has a region. there is a general purpose bucket and directory bucket (which is for a specific use case). we need to use a globally unique bucket name, and match the allowed pattern.\
Under the ownershop section, we can allow other accounts to own objects in the bucket based on Access Control Lists (ACL). and we can also choose to block public access to the bucket. we can also choose to use versioning, configure the encryption for objects in the bucket, we can use the default setting for now.

we can click <kbd>Upload</kbd> and <kbd>Add File</kbd>, and once uploaded, we can view the objects properties. when we get object url, we can't view it directly, since we blocked public access to the bucket. we can create "folders" inside the bucket and add files to it, and now the files will have the folder prefix.

### S3 Security: Bucket Policy

<details>
<summary>
Controlling Access to Objects in the Bucket.
</summary>

- user-based: using <cloud>IAM</cloud> Policies
- resource-based: bucket policies, bucket-wide rules, allows for cross-account access
- Bucket Access Control List (ACL) - can be disabled
- Object Access Control List (ACL) - even more fine grained (can be disabled).

> Note: an IAM principal can access an S3 object if
> - The user IAM permissions ALLOW it OR the resource policy ALLOWS it.
> - AND there's no explicit DENY.

the common thing to use is bucket policies, they have a similar structure as <cloud>IAM</cloud> policies

> JSON based policies
> 
> - Resources: buckets and objects
> - Effect: Allow / Deny
> - Actions: Set of API to Allow or Deny
> - Principal: The account or user to apply the policy to
>
> Use S3 bucket for policy to:
>
> - Grant public access to the bucket
> - Force objects to be encrypted at upload
> - Grant access to another account (Cross Account)
> - Optional Conditions on:
>   - Public IP or Elastic IP (not on Private IP)
>   - Source VPC or Source VPC Endpoint - only works with VPC Endpoints
>   - CloudFront Origin Identity
>   - MFA

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicRead",
      "Effect": "Allow",
      "Principal": "*",
      "Action": [
        "s3:GetObject"
      ],
      "Resource":[
        "arn:aws:s3:::examplebucket/*"
      ]
    }
  ]
}
```

the bucket level seetings can stop public acccess, even with ACL and bucket policies, this is another layer of security.

under the "permission tab", we can allow public access, and now objects *CAN* be public. so we can create a bucket policy. we can use the sample policies or the policy generator.

</details>

### S3 Website Overview

<details>
<summary>
Static Website hosting
</summary>

> S3 can host static websites and have them accessible on the Internet.\
> The website URL will be (depending on the region):
> 
> - http://bucket-name.s3-website-aws-region.amazonaws.com
> - http://bucket-name.s3-website.aws-region.amazonaws.com
> 
> If you get a 403 Forbidden error, make sure the bucket
policy allows public reads!

websites on S3 don't have a "backend" server, they are static and use static html documents. we can specify the index page and error page, and set some redirection rules. all files should be publicly accessible.
</details>

### S3 Versioning

<details>
<summary>
Objects Version Control
</summary>

> You can version your files in Amazon S3
> 
> - It is enabled at the bucket level
> - Same key overwrite will change the "version": 1, 2, 3...
> - It is best practice to version your buckets
>   - Protect against unintended deletes (ability to restore a version)
>   - Easy roll back to previous version
> - Any file that is not versioned prior to enabling versioning will have version "null"
> - Suspending versioning does not delete the previous versions

when we "delete" a file, we actually add a delete marker. we can add lifecycle rules to control what happens with older versions and have them.
</details>

### S3 Replication

<details>
<summary>
Replicating Data from one Bucket to Another
</summary>

> - Must enable Versioning in source and destination buckets.
> - Cross-Region Replication (CRR)
> - Same-Region Replication (SRR)
> - Buckets can be in different AWS accounts
> - Copying is asynchronous
> - Must give proper IAM permissions to S3

we can use replication to move data between buckets, either in the same account and region, or across regions and accounts. replicating data across regions costs more, and replicating across AWS accounts reqires specific permissions.

> - After you enable Replication, only new objects are replicated
>   - Optionally, you can replicate existing objects using *S3 Batch Replication* - Replicates existing objects and objects that failed replication
> - For DELETE operations
>   - Can replicate delete markers from source to target (optional setting)
>   - Deletions with a version ID are not replicated (to avoid malicious deletes)
> - There is no "chaining" of replication - If bucket 1 has replication into bucket 2, which has replication into bucket 3. Then objects created in bucket 1 are **not** replicated to bucket 3.

in the origin bucket, under the "management" tab, we can <kbd>Create replication rule</kbd>, we set the scope (which objects will be replicated), and the target bucket. we can also control more settings like storage tiers and lifecycle, and also replicate the delete markers (not enabled by default).

</details>

### S3 Storage Classes Overview

<details>
<summary>
Storage Tiers For S3 Objects
</summary>

storage classes (tiers) control the durability, availability, retrieval time and pricing of storing an object.

all storage classes have 99.999999999% durability, and avalability from 99% (SLA) to 99.99%.

tiers:

- Amazon S3 Standard - General Purpose
- Amazon S3 Standard-Infrequent Access (IA)
- Amazon S3 One Zone-Infrequent Access
- Amazon S3 Glacier Instant Retrieval
- Amazon S3 Glacier Flexible Retrieval
- Amazon S3 Glacier Deep Archive
- Amazon S3 Intelligent Tiering

glacier storage is for data we don't access usually, we want to store it for long time. storing costs are low, and retrival costs differ based on how fast we want the data.\
Intelligent tier allows AWS to move the data between storage classes, there is a small fee for the tiering, but it manages the objects for us, and there is no retrival fee.

#### S3 Lifecycle Rules

object can be moved between storage classes automatically using lifecycle rules, from standard to standard-IA, to glacier, etc...

a lifecycle rules has the transition action (where to move), or to expire them (delete), expire old versions or remove incomplete multipart uploads. we can also set the scope of the objects by specifying the path prefix.

we can use S3 analytics on the bucket to get recomendation on our access patterns to help us decide what rules to write. 

</details>

### S3 Event Notifications

<details>
<summary>
Integrating S3 events with other services
</summary>

we can set actions to happen on different S3 events (created objects, removed objects), and have other services react to to them:

- <cloud>SNS</cloud> - directly
- <cloud>SQS</cloud> - directly
- <cloud>Lambda</cloud> - directly
- <cloud>EventBridge</cloud> - more advanced

for this to work, we need to attach <cloud>IAM</cloud> Access Policies. the "Principal" will be the S3 bucket service, with the resource being the topic/queue or lambda, and the condition specifying which specific bucket it is (by ARN).\
another option is to use <cloud>EventBridge</cloud> to send events to other servies, with advanced filtering options, sending to multiple targets, with extra features.

in our bucket, under the "properties", we can set up event notification or enable the <cloud>EvenBridge</cloud> notifications. for now, we <kbd>Create Event Notification</kbd> set the the event types, specify a destination (<cloud>SNS</cloud> topic, <cloud>SQS</cloud> queue or <cloud>Lambda</cloud> function).\
We need to allow our S3 bucket to send events to that target, so we first need to modify the target to allow the bucket to operate on it. this is the access policy. the web console will check that we have the permissions.\
We can now create a queue and set it up, and then upload a file to S3 and see that a message was created in the queue.

</details>

### S3 Performance

<details>
<summary>
Limits, multipart upload, S3 Transfer Acceleration and byte-range fetch
</summary>

> Amazon S3 automatically scales to high request rates, latency 100-200 ms.  Your application can achieve at least 3,500 `PUT`/`COPY`/`POST`/`DELETE` or 5,500 `GET`/`HEAD` requests per second per prefix in a bucket.\
> There are no limits to the number of prefixes in a bucket.\
> Examples (object path => prefix):
> 
> - bucket/folder1/sub1/file => /folder1/sub1/
> - bucket/folder1/sub2/file => /folder1/sub2/
> - bucket/1/file => /1/
> - bucket/2/file => /2/
>
> If you spread reads across all four prefixes evenly, you can achieve 22,000 requests per second for `GET` and `HEAD`

For large files, there is a "multi-part upload" option, which should be used for for files above 100Mb, and must be used for files larger than 5GB.\
another option to get better performance is "S3 global Transfer Acceleration", which routes the files through AWS edge locations, from which AWS can transfer them faster in it's own network. this works for both downloads and uploads.

when we read files, we can perform GET operations based on byte-ranges, this helps us in resiliency against failures and increse the download speed (acting like multi-part download). we can also use this to get a part of the file (like getting the header) and using that instead of getting the entirety of the file. this will only if we know how the file is formatted.
</details>

### S3 Batch Operations

<details>
<summary>
Bulk operation on multiple object, Inventory objects in the bucket.
</summary>

> Perform bulk operations on existing S3 objects with a
single request, example:
> -  Modify object metadata & properties
> -  Copy objects between S3 buckets
> -  Encrypt un-encrypted objects
> -  Modify ACLs, tags
> -  Restore objects from S3 Glacier
> -  Invoke Lambda function to perform custom action on each object
> 
>  A job consists of a list of objects, the action to perform, and optional parameters.\
> S3 Batch Operations manages retries, tracks progress, sends completion notifications, generate reports. You can use S3 Inventory to get object list and use S3 Select to filter your objects.

(demo)\
we create bucket upload five files into it, including the batch.csv file

```csv
s3-batch-demo-stephane-v2,beach.jpg
s3-batch-demo-stephane-v2,coffee.jpg
s3-batch-demo-stephane-v2,index.html
s3-batch-demo-stephane-v2,error.html
```

our batch operation will add tags to the files, so under "Batch Operations" we choose <kbd>Create Job</kbd>. we give it the key to the csv file as a manifest object (we can also use a S3 inventory report). we choose the operation type, we choose to "replace all object tags", and we give them new tags, we can generate a report for the job, set the job priority.\
we also need an <cloud>IAM</cloud> role with permssions to act on the objects in the bucket. we confirm the job <kbd>Run Job</kbd>, and when it's complete, we can see the status (successes and failures), and view the objects and the newly added tags.

#### S3 Inventory

Listing all the objects and the corresponding metadata - better the using the <cloud>S3</cloud> list API.\
We can run this in a scheduled job (daily, weekly, etc...)

> - Usage examples:
> 
> - Audit and report on the replication and encryption status of your objects
> - Get the number of objects in an S3 bucket
> - Identify the total storage of previous object versions

we can get the results in csv, orc or Apahce Parquet format. and then filter using S3 select, use some service to query the data, and also feed it into a batch operation job.

in out bucket, under the "management" tab, we can <kbd>Create inventory Configuration</kbd>. we set up the scope (prefix), where the report is created to (this bucket or another one), and we update the bucket policy. we choose the frequency of the report and the format (choose CSV if we want to use in the S3 batch operations). we can add some additional fields. the first report will be generated in about 48 hours, for each report a folder is created with the schema manifest and the manifest checksum.

#### S3 Select and Glacier Select

> - Retrieve less data using SQL by performing server-side filtering
> - Can filter by rows & columns (simple SQL statements)
> - Less network transfer, less CPU cost client-side

</details>

### S3 Glacier

<details>
<summary>
Long Term Storage Class
</summary>

> Low-cost object storage meant for archiving / backup
> - Data is retained for the longer term (10s of years)
> - Alternative to on-premises magnetic tape storage
> - Average annual durability is 99.999999999%
> - Cost per storage per month ($0.004 / GB - Standard vs $0.00099 / GB Deep Archive)
>
> Each item in Glacier is called "Archive" (up to 40TB)
> 
> - Archives are stored in "Vaults"
> - By default, data encrypted at rest using AES-256 - keys managed by AWS

<cloud>S3</cloud> has buckets and objects, Glacier hsa vaults and achieves.

Vaults can be created, but deleted only when they are empty. we can get the metadata of a vault containing the creation data, number of achieves (files) and total size. we can also download the Vault inventory, a list of all the achieves in it.\
We can upload files into glacier, and when we request to download from it- we actually create a retrival job, so we get a limited time link to download the data. we can choose how fast the retrival job will be, faster == more expensive.

> Retrieval Options:
>
> - Expedited (1 to 5 minutes retrieval) - $0.03 per GB and $10 per 1000 requests
> - Standard (3 to 5 hours) - $0.01 per GB and 0.03 per 1000 requests
> - Bulk (5 to 12 hours) - $0.0025 per GB and $0.025 per 1000 requests

when we want to use expedited retrival, we need provisioned retrival units, which we purchase from AWS.

#### Glacier Vault Lock and Policies

Glacier vaults can have policies and locks, each vault having one of each and that's all. they are written in json and are like access policies. Locks are immutable, and can't be changed (until they expire), even by the root account. they are used for complainace and regulatory requirements. they control what can be done with the files, like forbiding deleting achieves from vaults if they are younger than one year.\
We can set notifications on S3 Glacier jobs, like notifying us that a retrival job was completed, or we can use S3 event notifications.

</details>

### S3 Multi-Part Upload Deep Dive

<details>
<summary>
Upload large files in parts
</summary>

> Upload large objects in parts (any order)
> - Recommended for files > 100MB, must use for files > 5GB
> - Can help parallelize uploads (speed up transfers)
> - Max. parts: 10,000
> - Failures: restart uploading ONLY failed parts (improves performance)
> - Use Lifecycle Policy to automate old parts deletion of unfinished upload after x days (e.g., network outage)
> - Upload using AWS CLI or AWS SDK

we finish the upload by running a `Complete` request to concatenate all the files into a single object. we create the lifecycle rule through the "Management" tab.

(hands on)\

we use cloud shell to create a file, split it into 3 parts,create a multi-part upload operation,  and upload each part separately. then we take the Id we got and use it to send a "Complete" command - we need to use the part number instead of the E-TAGS.
The parts don't appear in the web console, we can only list them from the CLI.

```sh
BUCKET_NAME=yourbucketnamehere

# create a bucket
aws s3 mb s3://$BUCKET_NAME

# Generate a 100 MB file
dd if=/dev/zero of=100MB.txt bs=1MB count=100

# Split it into 3 parts of 35 MB
split -b 35m 100MB.txt 100MB_part_

# Initiate the multi-part upload
aws s3api create-multipart-upload --bucket $BUCKET_NAME --key 100MB.txt

# get back the upload_id and insert it below
UPLOAD_ID=<LO99SQpq7LSDaMiArIV0SWCUlc3fyCBXflTX8yfae8Ux62y2nkceBXProyeV54kN4SIajjOZvE7x8btrohRLXpJKcu0vQd7oomr2ATVu8iTC.lFc4nPDLvwZVpFMoqL9AXiXuQpYANbIw8jT.7q42A-->

# list existing multi part uploads
aws s3api list-multipart-uploads --bucket $BUCKET_NAME

# Upload the parts
aws s3api upload-part --bucket $BUCKET_NAME --key 100MB.txt --part-number 1 --body 100MB_part_aa --upload-id $UPLOAD_ID

aws s3api upload-part --bucket $BUCKET_NAME --key 100MB.txt --part-number 2 --body 100MB_part_ab --upload-id $UPLOAD_ID

aws s3api upload-part --bucket $BUCKET_NAME --key 100MB.txt --part-number 3 --body 100MB_part_ac --upload-id $UPLOAD_ID

# list the parts
aws s3api list-parts --upload-id $UPLOAD_ID --bucket $BUCKET_NAME --key 100MB.txt

# Complete multi-part upload
# replace the ETag with the part id from the previous command
aws s3api complete-multipart-upload --bucket $BUCKET_NAME --key 100MB.txt --upload-id $UPLOAD_ID --multipart-upload "{\"Parts\":[{\"ETag\":\"etag1\",\"PartNumber\":1},{\"ETag\":\"etag2\",\"PartNumber\":2},{\"ETag\":\"etag3\",\"PartNumber\":3}]}"

# list the parts
# nothing will show up here!
aws s3api list-parts --upload-id $UPLOAD_ID --bucket $BUCKET_NAME --key 100MB.txt
```

</details>

### Athena

<details>
<summary>
Serverless query service to analyze data stored in Amazon S3
</summary>

> Uses standard SQL language to query the files (built on Presto engine):
> 
> - Supports CSV, JSON, ORC, Avro, and Parquet
> - Pricing: $5.00 per TB of data scanned
> - Commonly used with <cloud>Amazon Quicksight</cloud> for reporting/dashboards
> - Use cases: Business intelligence / analytics /  reporting, analyze & query VPC Flow Logs, ELB Logs, CloudTrail trails, etc... 

analyzes data in place, without moving it outside the bucket.
We can get better performance by scanning less data, so we can use columnar data - either Apache Parquet or ORC. we can transform our data into those types with <cloud>AWS Glue</cloud> service. another option is to compress data. it's better to work on several large files than a large number of small files.

We can partition our data by storing some columns inside the path. such as "bucket/prefix/year=VALUE/month=VALUE/day=VALUE/file.csv", so we can set the job to only scan files from a specific time range. this is called virtual columns partition.

> Federated Queries Allows you to run SQL queries across data stored in relational, non-relational, object, and custom data sources (AWS or on-premises)
>
> - Uses Data Source Connectors that run on AWS Lambda to run Federated Queries (e.g., CloudWatch Logs, DynamoDB, RDS...)
> - Store the results back in Amazon S3

(hands on)\
In the <cloud>Athena</cloud> service, we choose the query result location, such as an S3 bucket. this is where the results will be saved.\
In the Query editor, we choose the Data Source, such as access logs. we create the database, and the external table and set the correct location and prefix.\
We can preview the data by running a `SELECT` command with ten items limit, and we can run other SQL commands,


```sql
create database s3_access_logs_db;

CREATE EXTERNAL TABLE IF NOT EXISTS s3_access_logs_db.my_bucket_logs(
  BucketOwner STRING,
  Bucket STRING,
  RequestDateTime STRING,
  RemoteIP STRING,
  Requester STRING,
  RequestID STRING,
  Operation STRING,
  Key STRING,
  RequestURI_operation STRING,
  RequestURI_key STRING,
  RequestURI_httpProtoversion STRING,
  HTTPstatus STRING,
  ErrorCode STRING,
  BytesSent BIGINT,
  ObjectSize BIGINT,
  TotalTime STRING,
  TurnAroundTime STRING,
  Referrer STRING,
  UserAgent STRING,
  VersionId STRING,
  HostId STRING,
  SigV STRING,
  CipherSuite STRING,
  AuthType STRING,
  EndPoint STRING,
  TLSVersion STRING
) 
ROW FORMAT SERDE 'org.apache.hadoop.hive.serde2.RegexSerDe'
WITH SERDEPROPERTIES(
  'serialization.format' = '1', 'input.regex' = '([^ ]*) ([^ ]*) \\[(.*?)\\] ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) \\\"([^ ]*) ([^ ]*) (- |[^ ]*)\\\" (-|[0-9]*) ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) (\"[^\"]*\") ([^ ]*)(?: ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*))?.*$'
)
LOCATION 's3://target-bucket-name/prefix/';


SELECT requesturi_operation, httpstatus, count(*) FROM "s3_access_logs_db"."my_bucket_logs" 
GROUP BY requesturi_operation, httpstatus;

SELECT * FROM "s3_access_logs_db"."my_bucket_logs"
where httpstatus='403';
```

There are some other things we can do, like saving queries, seeing previous queries, etc...

</details>

### S3 Encryption

<details>
<summary>
Encrypting Objects in the bucket and during transport.
</summary>

SSE - server side encryption
- SSE-S3 - <cloud>S3</cloud> managed keys, this is the default
- SSE-KMS - <cloud>KMS</cloud> keys, stored in AWS service
- DSSE-KMS - <cloud>KMS</cloud> keys, dual-layer encryption
- SSE-C - Customer provided keys

There is also client side encryption.

SSE-S3 uses AES-256 encryption, so the header should be `"x-amz-server-side-encryption": "AES256"`. SSE-KMS gets the keys from <cloud>KMS</cloud>, so the keys can be audited with <cloud>CloudTrail</cloud>,  and the header is `"x-amz-server-side-encryption": "aws:kms"`, using SSE-KMS counts toward the keys quotas limit of the Decrypt KMS Api.\
For customer provided keys (SSE-C), the https protocol must be used, the keys aren't stored in AWS, and must be part of the request. to read the file the same key ust be provided.\
Client side encryption means that the client encrypt and decrypt the data before sending and after receiving it.

In-flight encryption (SSL/TLS) - we have two S3 endpoints, one for HTTP (non-encrypted) and one for HTTPS (encrypted in flight), HTTPS is the recommended option, and must be used for SSE-C. we can set a bucket policy to force in-transit encryption

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "InFlightOnly",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::examplebucket/*",
      "Condition": {
        "Bool": {
          "aws:SecureTransport": "false"
        }
      }
    }
  ]
}
```

when we create a bucket, we choose the default encryption method, if we choose KMS key, we can use the default <cloud>S3-KMS</cloud> key, which saves us a bit of money, or use a different KMS key. we can't use SSE-C (customer provided keys) from the console, only from the command line.

even if we have a default encryption method, we can still override it by providing the encryption headers, but bucket policies can stop that from happening
</details>

### S3 CORS

<details>
<summary>
Cross Origin Resource Sharing
</summary>

Origin = scheme (protocol) + host (domain) + port. CORS is a web browser based mechanism to allow requests to other origins while visiting the main origin.

"http://example.com/app1" and "http://example.com/app2" have the same origin, scheme is http, host is "example.com", and the port it implicit 80 (since this is http). but "http://www.example.com" and "http://other.example.com" don't share the same origin.

when the browser access the first origin (A), it gets a response telling it needs resources from a different origin (B) are needed. so the browser performs a `OPTION` pre-flight request on origin B, asking it if origin A is approved (CORS), and if it's approved, then the browser performs the requests on origin B.

> For <cloud>S3</cloud>, If a client makes a cross-origin request on our S3 bucket, we need to enable the correct CORS headers. You can allow for a specific origin or for * (all origins).

this will come up if we have a website that uses resources from the bucket (such as images). we can see this in the demo by using two buckets (both hosting static website). we can the CORS request being blocked in developer tools in the browser.

to allow CORS, we go to the target bucket (B), and add the policy to allow the `GET` method from the first bucket static website (A).

```json
[
    {
        "AllowedHeaders": [
            "Authorization"
        ],
        "AllowedMethods": [
            "GET"
        ],
        "AllowedOrigins": [
            "<url of first bucket with http://...without slash at the end>",
            "http://<bucket-name>.s3-website.eu-west-1.amazonaws.com"
        ],
        "ExposeHeaders": [],
        "MaxAgeSeconds": 3000
    }
]
```

</details>

### S3 MFA Delete

<details>
<summary>
Multi-Factor Authentication, protection from accidental delete.
</summary>

> MFA (Multi-Factor Authentication) - force users to generate a code on a device (usually a mobile phone or hardware) before doing important operations on S3.\
> MFA will be required to:
> 
> - Permanently delete an object version
> - Suspend Versioning on the bucket
> 
> MFA won't be required to:
> 
> - Enable Versioning
> - List deleted versions
> 
> To use MFA Delete, Versioning must be enabled on the bucket. Only the bucket owner (root account) can enable/disable MFA Delete.

we need to have an MFA device configured (with an ARN), and MFA-delete protection can only be done using the api, not the web console.

```sh
PROFILE_NAME=root-mfa-delete-demo
MFA_ARN=arn-of-mfa-device

# generate root access keys
aws configure --profile $PROFILE_NAME

# list all buckets
aws s3 ls --profile $PROFILE_NAME
BUCKET_NAME=mfa-demo-stephane


# enable mfa delete - get code from device
aws s3api put-bucket-versioning --bucket $BUCKET_NAME --versioning-configuration Status=Enabled,MFADelete=Enabled --mfa "$MFA_ARN <mfa-code>" --profile $PROFILE_NAME

# disable mfa delete - get code from device
aws s3api put-bucket-versioning --bucket $BUCKET_NAME --versioning-configuration Status=Enabled,MFADelete=Disabled --mfa "$MFA_ARN <mfa-code>" --profile $PROFILE_NAME

# delete the root credentials in the IAM console!!!
```

objects with MFA delete can't be deleted in the console, only through the API.
</details>

### S3 Access Logs

<details>
<summary>
Log any access to S3 bucket.
</summary>

> For audit purpose, you may want to log all access to S3 buckets.
>
> - Any request made to S3, from any account, authorized or denied, will be logged into another S3 bucket
> - That data can be analyzed using data analysis tools (like <cloud>Athena</cloud>)
> - The target logging bucket must be in the same AWS region
> 
> **Warning - Do not set your logging bucket to be the monitored bucket, It will create a logging loop, and your bucket will grow exponentially.**

we can see the log format in the [documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/LogFormat.html), this is also what we used in the above example for Athena.

we need to change the bucket policy to allow the logged bucket to write access logs.

</details>

### S3 Pre-signed URLs

<details>
<summary>
Pre-signed URL to allow temporary access and permissions to resources.
</summary>

> Users given a pre-signed URL inherit the permissions of the user that generated the URL for GET / PUT.\
> Examples:
> 
> - Allow only logged-in users to download a premium video from your S3 bucket
> - Allow an ever-changing list of users to download files by generating URLs dynamically
> - Allow temporarily a user to upload a file to a precise location in your S3 bucket

can be generated using the web console, cli or an SDK, this effects the maximum expiration time.

- Web console - from 1 minute to 720 (12 hours).
- API and SDK - default is 3600 seconds (1 hour), maximum 604800 seconds (168 hours = 7 days = 1 week).

in the console, we can choose an object, <kbd>Object Actions</kbd>, <cloud>Share with a presigned url</cloud>. under the hood, this is what happens when we click <kbd>Open</kbd> on the object.
</details>

### Glacier Vault Lock & S3 Object Lock

<details>
<summary>
Protecting Objects from deletion for legal reasons
</summary>

S3 glacier vault lock - Adopt a WORM (Write Once ReadMany) model for the vault (analogue to bucket).

> - Create a Vault Lock Policy
> - Lock the policy for future edits (can no longer be changed or deleted)
> - Helpful for compliance and data retention

S3 object lock prevents object version deletion for a specific period of time (even if it's not in <cloud>Glacier</cloud>). works on a single object version. has two modes:

- governance mode - some users have permession to change retention time or delete the object - more permissive.
- complinace mode - can't be deleted, retention perios can't be changed by anyone (including root) - more strict.

in both cases we have a retention period, which can be extended as needed.

another thing is **Legal Hold**, which protects the object indefinitely, regardless of other projects, requires the `s3::PutObjectLegalHold` IAM permission.
</details>

### S3 Access Points

<details>
<summary>
Simplify security management for S3 buckets with Access Points.
</summary>

granting access points to specific prefixes (paths) in the bucket, with specialized permissions. fine grained permissions, defined not through <cloud>IAM</cloud> roles, but through the bucket and the access points.

Each access point has it's own DNS name (internet origin or VPC origin for private access), and it's own access point policy (like the bucket policy).\
<cloud>VPC</cloud> origin access points are only accessible from within the VPC. we create a *VPC Endpooint* (gateway or interface) which can connect to the access point. the endpoint policy must allow access to the bucket and the access policy.

(VPC has endPoint, S3 has accessPoint. the endPoint connects to the accessPoint. both have policies)


handsOn:\
Under the <cloud>S3</cloud> console, in the "Access Points" secion, we click <kbd>Create access point</kbd>, we choose the bucket, and we choose the network origin as internet or <cloud>VPC</cloud> (select which one). we choose to block all public access, and we can define the policy for fine grained control. we can look at examples,

in this example, we only allow the user "Alice" to access the bucket through the accessPoint, and only on resources in the specified path.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AccessPointExample",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::<account-number>:user/Alice"
      },
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource":[
        "arn:aws:s3:us-west2:<account-number>:accesspoint/my-access-point/objects/Alice/*"
      ]
    }
  ]
}
```

we can also define a bucket policy to only allow the bucket to be access through accessPoints in the current account
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "OnlyAllowAccessFromAccessPoint",
      "Effect": "Allow",
      "Principal": {
        "AWS": "**"
      },
      "Action": "*",
      "Resource":[
        "<bucket-arn>",
        "<bucket-arn>/*"
      ],
      "Condition":{
        "StringEquals": {"s3:DataAccessPointAccount": "<account-number>"}
      }
    }
  ]
}
```

#### S3 Multi-Region Access Points

> - Provide a global endpoint that span S3 buckets in multiple AWS regions
> - Dynamically route requests to the nearest S3 bucket (lowest latency)
> - Bi-directional S3 bucket replication rules are created to keep data in sync across regions
> - Failover Controls - allows you to shift requests across S3 buckets in
different AWS regions within minutes (Active-Active or Active-Passive)

hands On:\
we create two buckets in two regions, and in the "Multi-Region Access Points" section, we choose <kbd>Create Multi-Region Access Point</kbd>. next we add the two bucket we created, we set the security policies as before, and keep the default to block all public access. this might take up to 24 hours to create. under the "Replication and failover" tab, we can see where our buckets are on the map, how they are replicated and the failover.\
We click <kbd>Create Replication Rules</kbd> to choose bi-directional replication (we need bucket versioning), and we can filter the replication for specific prefixes.

#### S3 VPC Endpoints

usually the S3 bucket is access through the public internet, VPC endpoints allow us to connect to S3 bucket using th private IP of our <cloud>EC2</cloud> machines, using a *VPC EndPoint Gateway* and not the Internet Gateway. the endPoint gateway connects to the bucket, and we can define bucket policies to only allow access from that source

- "AWS:SourceVpce" for one or more endPoints
- "AWS:SourceVpc" for all possible endPoints
</details>

</details>

## Advanced Storage Section

<details>
<summary>
Additional Storage Stuff
</summary>

### AWS Snow

<details>
<summary>
Offline data transfer Device
</summary>

Physical Devices to store and migrate data to and from AWS.

> Highly-secure, portable devices to collect and process data at the
edge, and migrate data into and out of AWS

| Device         | SnowCone             | SnowBall Edge  |
| -------------- | -------------------- | -------------- |
| Capacity       | 8 TB hdd / 14 TB SSD | 80 TB / 210 TB |
| Migration Size | terabyes             | petabyes       |

there used to be something called snowmobile.

the devices allow us to transfer data from our machines (on-prem) to the <cloud>snowball</cloud> device using the local connection, and then we ship the device back to AWS, which uploads the data to the cloud. we use them when we have limited connectivity, limited bandwidth, high network costs, when we are sharing the bandwidth with others, and when the connection is not stable.

we use the snowball client (or AWS OpsHub) to manage the data transfer, the device is completely wiped by AWS after the transfer is complete, so there is no risk of the next customer seeing your data.

SnowBall Edge devices act as an edge location, they can run <cloud>EC2</cloud> or <cloud>Lambda</cloud> functions directly on the device, without going to the global AWS. this acts as an "local-aws" datacenter. they come with a large number of Virtual CPUs. we can change the AMI that comes with the device, and configure some other stuff.
</details>

### Amazon FSx

<details>
<summary>
Additional managed file systems.
</summary>

similar to <cloud>EFS</cloud> - Elastic file system, but using third-party file systems. a fully managed service

- Lustre
- Windows File Server
- NetApp ONTAP
- OpenZFS

#### Amazon FSx for Windows (File Server)
 
> FSx for Windows is a fully managed Windows file system share drive:
> 
> - Supports SMB protocol & Windows NTFS
> - Microsoft Active Directory integration, ACLs, user quotas
> - Can be mounted on Linux EC2 instances
> - Supports Microsoft's Distributed File System (DFS) Namespaces (group files across multiple FS)
> - Scale up to 10s of GB/s, millions of IOPS, 100s PB of data
> - Storage Options:
>   - SSD - latency sensitive workloads (databases, media processing, data analytics, ...)
>   - HDD - broad spectrum of workloads (home directory, CMS, ...)
> - Can be accessed from your on-premises infrastructure (VPN or <cloud>Direct Connect</cloud>)
> - Can be configured to be Multi-AZ (high availability)
> - Data is backed-up daily to <cloud>S3</cloud>

#### Lustre - Linux Cluster
> Lustre is a type of parallel distributed file system, for large-scale computing.\
> The name Lustre is derived from "Linux" and "cluster
> 
> - Machine Learning, High Performance Computing (HPC)
> - Video Processing, Financial Modeling, Electronic Design Automation
> - Scales up to 100s GB/s, millions of IOPS, sub-ms latencies
> 
> Storage Options:
> - SSD - low-latency, IOPS intensive workloads, small & random file operations
> - HDD - throughput-intensive workloads, large & sequential file operations
> 
> Seamless integration with S3:
> - Can "read S3" as a file system (through FSx)
> - Can write the output of the computations back to S3 (through FSx)
> 
> Can be used from on-premises servers (VPN or <cloud>Direct Connect</cloud>)

Two files system options:

> Scratch File System
> - Temporary storage
> - Data is not replicated (doesn't persist if file server fails)
> - High burst (6x faster, 200MBps per TiB)
> - Usage: short-term processing, optimize costs
> 
> Persistent File System
> 
> - Long-term storage
> - Data is replicated within same AZ
> - Replace failed files within minutes
> - Usage: long-term processing,sensitive data

#### Amazon FSx for NetApp ONTAP

> Managed NetApp ONTAP on AWS, File System compatible with NFS, SMB, iSCSI protocol. Move workloads running on ONTAP or NAS to AWS.
> 
> - Works with:
>   - Linux
>   - Windows
>   - MacOS
>   - VMware Cloud on AWS
>   - Amazon Workspaces & AppStream 2.0
>   - Amazon EC2, ECS and EKS
>
> - Storage shrinks or grows automatically
> - Snapshots, replication, low-cost, compression and data de-duplication
> - Point-in-time instantaneous cloning (helpful for testing new workloads)

#### Amazon FSx for OpenZFS
>  Managed OpenZFS file system on AWS. File System compatible with NFS (v3, v4, v4.1, v4.2), Move workloads running on ZFS to AWS.
> - Works with:
>   - Linux
>   - Windows
>   - MacOS
>   - VMware Cloud on AWS
>   - Amazon Workspaces & AppStream 2.0
>   - Amazon EC2, ECS and EKS
> 
> - Up to 1,000,000 IOPS with < 0.5ms latency
> - Snapshots, compression and low-cost
> - Point-in-time instantaneous cloning (helpful for testing new workloads)

#### Hands On
in the amazon <cloud>Fsx</cloud> server, we can <kbd>Create File System</kbd>, we choose the type, the configuration, where it's deployed, persistent and scratch storage.\
We can also look at the option for windows file server, enable Active Directory authentication, adn other stuff.

Windows fileserver has single or multi Availability Zone option.
</details>

### Storage Gateway Overview

<details>
<summary>
Bridging between local server and the cloud.
</summary>

Hybrid Cloud - part of infrastructure is on-premises, and part of it is on the cloud.\
there can be many reasons, such as being in the middle of a long migration, having to store some workloads locally for security or compliance reasons, or for other performance and costs reasons.

for the on-premises data to interact cleanly with S3 buckets, we can use <cloud>Storage Gateway</cloud> as a bridge.

- Block storage: <cloud>EBS</cloud>
- File storage: <cloud>EFS</cloud>, <cloud>FSX</cloud>
- Object storage: <cloud>S3</cloud>, <cloud>Glacier</cloud>

use cases:
- disaster recovery
- backup and restore
- tiered storage
- on-premises cache and low-latency file access

there are different types of storage gateways:

- <cloud>S3</cloud> file gateway
- <cloud>Fsx</cloud> file gateway
- Volume Gateway
- Tape Gateway

The <cloud>S3</cloud> file gateway (not glacier) uses the NFS or SMB protocol to expose the data on the bucket as files. the on-premises machines will see the bucket as another disk, and the gateway will translate the commands to HTTP requests, and will store the recently used data cached. the file gateway needs <cloud>IAM</cloud> role and for SMB protocol requires Active Directory authentication.

The <cloud>Fsx</cloud> file gateway provides native access for Windows File Server, even if it's already possible without the gateway. but it adds the local cache and low latency.

The Volume Gateway is for block storage, using the iSCSI protocol (backed by <cloud>S3</cloud>). the data is backed into <cloud>EBS</cloud> snapshots.
- Cached Volume: Low latency access to most recent data
- Stored Volume: entire dataset is on-premies, scheduled backup to <cloud>S3</cloud>

The Tape gateway is a a way to store Tape Data into the cloud.

> - Some companies have backup processes using physical tapes (!).
> - With Tape Gateway, companies use the same processes but, in the cloud.
> - Virtual Tape Library (VTL) backed by Amazon S3 and Glacier.
> - Back up data using existing tape-based processes (and iSCSI interface).
> - Works with leading backup software vendors.

For all gateways, it needs to be installed locally, either on a virtual server (virtual machine), or by buying a hardware appliance and setting it as one of the gateway options.


The file gateway is posix compliant (linux file system), there are some considerations for rebooting the storage gateways (especially for volume and tape gateways, stop the service, reboot, start the service).\
The Storage Gateway needs to be activated with an activation key. activation can fail if port 80 isn't opened or if the time isn't synchronized.

for volume gateway where we ise it as a cache, we want the cache to be efficient, we need to look at some metrics (CacheHitPercent, CachePercentUsed), we can use a larger disk.
</details>

</details>

## CloudFront

<details>
<summary>
Content Delivery Network Using Edge Locations.
</summary>

CDN - Content Delivery Network,  improves read performance by caching data at edge locations around the world. also protects against DDOS attacks, integrates with <cloud>Shield</cloud> and <cloud>WAF</cloud> (Web Application Firewall) services.

Origins:
- <cloud>S3</cloud> buckets, using Origin Access Control for enhanced security (only <cloud>CloudFront</cloud> will access the files)
  - can also be used as an ingress to upload files into S3
- Custom Origin HTTP
  - Application Load Balancer
  - <cloud>EC2</cloud> instace
  - <cloud>S3</cloud> website (enable bucket as static website)
  - Any HTTP backend

the client connects to the edge location, and will return a cached result, or use the quicker AWS network (faster than public internet) to grab the value.

### CloudFront with S3

<details>
<summary>
Using S3 as the origin
</summary>

Origin Access Control and the Bucket Policy make the objects more secure.

this is not the same as <cloud>S3</cloud> Cross-Region Replication.

> CloudFront:
> 
> - Global Edge network
> - Files are cached for a TTL (maybe a day)
> - Great for static content that must be available everywhere
> 
> S3 Cross Region Replication:
> 
> - Must be setup for each region you want replication to happen
> - Files are updated in near real-time
> - Read only
> - Great for dynamic content that needs to be available at low-latency in few regions

hands on:\
we need a bucket with our files, and in the <cloud>CloudFront</cloud> service, we choose the bucket as the origin domain, we can set the path in it, and we can control the origin access, we choose Origin Access Control, <kbd>Create Control Setting</kbd>, and we update the <cloud>S3</cloud> bucket policy. we disable <cloud>WAF</cloud> for now, and set the default root object for "index.html".\
the bucket policy needs to be updated to allow he cloudFront service distribution to access the files, we can always regenerate the policy from the console.\
When everything finishes, we get a domain name, which we can put in the address bar.

</details>

### CloudFront - ALB as an Origin

<details>
<summary>
EC2 machines and Application Load Balancer as origin.
</summary>

now we want to make our origin an <cloud>EC2</cloud> instance or the Application Load Balancer. for this we need to allow connection form the edge location to the security group of the resource.\
The other option is to have the <cloud>EC2</cloud> machine in a private subnet, and just expose the load balancer in the public subnet.

</details>

### CloudFront Geo Restriction

<details>
<summary>
Limiting Who Can Access.
</summary>

we can restrict the access based on countries. anyone accessing the edge location from outside the allow list won't be approved and won't get the data.\
we can use allow or block lists, this is done usually for copyright restriction reasons.

under the security tab, we can edit the list of countries, and allow or block countries

</details>

### CloudFront Reports, Logs and Troubleshooting

<details>
<summary>
Other Stuff
</summary>

> Logs every request made to CloudFront into a logging S3 bucket.
Like <cloud>S3</cloud> access logs, we can maintain a log of every request done through the distribution.

even without the access logs
> It's possible to generate reports on:
> 
> - Cache Statistics Report 
> - Popular Objects Report 
> - Top Referrers Report 
> - Usage Reports 
> - Viewers Report

under the settings tab, we can set the S3 bucket as the target.

> CloudFront caches HTTP 4xx and 5xx status codes returned by S3 ( or the origin server)
> - 4xx error code indicates that user doesn't have access to the underlying bucket (403) or the object user is requesting is not found (404)
> - 5xx error codes indicates Gateway issues
>
> 
</details>

### CloudFront Caching - Deep Dive

<details>
<summary>
Understanding Cache Performance
</summary>

> Cache based on
> - Headers
> - Session Cookies
> - Query String Parameters
> 
> The cache lives at each CloudFront Edge Location. You want to maximize the cache hit rate to minimize requests on the origin.
> 
> - Control the TTL (0 seconds to 1 year), can be set by the origin using the "Cache-Control" header, "Expires" header
> - You can invalidate part of the cache using the `CreateInvalidation` API

#### Header Caching
we can configure CloudFront to forward all headers to the origin, which would mean there is no caching. the TTL must be set zero.\
We can set a "white-list" set of headers, so the caching will be done based on only those values.\
We can also choose to only forward the default headers, or have no caching based on request headers at all, this will give us the best caching performance.

#### CloudFront Headers
when we use cloudFront, we can add constant headers to all requests which we send to the origin, they are constant values. a use case for this is to define custom behavior in the origin for requests coming from cloudFront.\
We can also set behavior settings, which are used for caching

#### Caching TTL
TTL - Time To live

"Cache-Control: max-age" replaces the "Expires" header. we can use this header in the response from the origin to control the TTL. we can also choose minimum, maximum and default values for the cache, to augment the response from the origin

#### Cookies and Query Parameters Caching

cookies are a special kind of headers, with additional data headers.

we have different settings:

> - Default: do not process the cookies
  > - Caching is not based on cookies
  > - Cookies are not forwarded
> - Forward a whitelist of cookies
>   - caching based on values in all the specified cookies
> - Forward all cookies
>   - Worst caching performance

for Query Parameters, there are similar options:

> - Default: do not process the query strings
>   - Caching is not based on query strings
>   - Parameters are not forwarded
> - Forward a whitelist of query strings
>   - Caching based on the parameter whitelist
> - Forward all query strings
>   - Caching based on all parameters

#### Dynamic and Static Caching

for static content, we don't need to forward headers, we can cache everything. for dynamic data, we need to forward some headers/cookies/parameter strings, so our cache will be worse. we should have two distributions, one for static content and one for dynamic (with two origins).

</details>

### CloudFront with ALB Sticky Sessions

<details>
<summary>
Sticky Sessions with CloudFront
</summary>

Sticky Sessions is a feature of <cloud>ELB</cloud> (Elastic Load Balancer) that tries to always pair the same request from the same user to the same instance that handled it before.\
For this to work, we need to forward the cookie that controls the seession affinity to the backend origin, so the sticky session behavior will remain.

</details>


</details>

## Databases for SysOps

<details>
<summary>
Relation Database Service
</summary>

<cloud>RDS</cloud> - Relation Database Service

> It's a managed DB service for DB use SQL as a query language.\
> It allows you to create databases in the cloud that are managed by AWS:
> 
> - Postgres
> - MySQL
> - MariaDB
> - Oracle
> - Microsoft SQL Server
> - IBM DB2
> - <cloud>Aurora</cloud> (AWS Proprietary database)

it's easier to use RDS than hand rolling our own DB on machines.

> RDS is a managed service:
> 
> - Automated provisioning, OS patching
> - Continuous backups and restore to specific timestamp (Point in Time Restore)!
> - Monitoring dashboards
> - Read replicas for improved read performance
> - Multi AZ setup for DR (Disaster Recovery)
> - Maintenance windows for upgrades
> - Scaling capability (vertical and horizontal)
> - Storage backed by EBS
> - BUT you can't SSH into your instances

Storage Auto Scaling - if we enable this option, then AWS can auto-scale the database storage, up to a maximum threshold. we can auto-incrase the storage, depending on some conditions.

handsOn:\
in the <cloud>RDS</cloud> service, we click <kbd>Create Database</kbd>. we need to select the database engine:

- <cloud>Aurora</cloud>
  - MySQL compatible edition
  - PostgreSQL compatible edition
- MySQL
- MariaDB
- PostgreSQL
- Oracle
- Microsoft SQL Server

we will use MySQL for the demo, and we can choose a pre-set configuration for production, dev/test or free-tier. in the availability and durability section, we can choose a single instance, multi-az DB instances for Disaster recovery, or multi-AZ DB cluster for disaster recovery and read performance.\
We need a user name and password for the database, we can integrate this with other services for password rotation, but we won't do this for the demo.\
We next choose the instance configuration - which <cloud>EC2</cloud> machine will run the database, we also select the storage type and size, for real cases we will want the io2 type, but for free-tier we will keep with gp2. we can enable auto-scaling for storage and set the limits.\
Under the connectivity section, we can set a connection between an existing <cloud>EC2</cloud> machine and the new database. this will set up the security group and subnets to match the instance. otherwise, we create the new VPC, subnets and security groups.\
The next section is database authentication (not the same as the database master password), and we can use passwords authentication or combine it with <cloud>IAM</cloud> credentials.\
We can enable enhanced monitoring, control back-up configurations (retention period, mainline window), select the option to export logs, and enable/disable deletion protection.

we can download a SQL client (sql-electron) to connect to the database. this will allow us to connect to the database.\
when the DB is created, we get an endpoint address, we put it in the database connection screen in the client, and we can run queries.

```sql
CREATE TABLE films (
    code        char(5) CONSTRAINT firstkey PRIMARY KEY,
    title       varchar(40) NOT NULL,
    did         integer NOT NULL,
    date_prod   date,
    kind        varchar(10),
    len         interval hour to minute
);

INSERT INTO films VALUES
    ('UA502', 'Bananas', 105, '1971-07-13', 'Comedy', '82 minutes');
```

we can click <kbd>Create Read Replica</kbd> to add a instance as a read-replica, we simply choose the instance type and storage, and either use a single or multi AZ deployment. we can <kbd>Take snapshots</kbd>, <kbd>Restore to a Point in Time</kbd>

under the monitoring tab, we can see CPU Utilization, Database connection count, and other options.

### RDS Multi AZ vs Read Replicas

<details>
<summary>
High availability and Disaster recovery
</summary>

Read-Replicas help for scaling read operations, we can create up to 15 read replicas, which can be in the same Availability Zone, same region or even another region. they are replicated asynchronously, so there is eventual consistency.\
Each of the read-replicas can be promoted to a read/write Database.\
a common use-case is to add read-performance when we have an analytics job. so the additional queries won't effect the production.

Unlike most AWS services, when using read replicas in the same Region (cross Availability Zone), there **AREN'T** network costs. this behavior usually happens only with managed services. and this is another reason to use <cloud>RDS</cloud> over manually managing the DB on <cloud>EC2</cloud> machines. there are still network costs for cross-region replication.

Multi-AZ is used for disaster recovery. in this case, the replication is done synchronously. there is one master instance that is currently taking requests, and one stand-by that is replicating it.\
In case of a problem, a failover happens, and the stand-by instance is promoted to being the master. there is a single DNS name, no need for manual intervention in the application.

A read replica instance can be used as a secondary-instance for disaster recovery. there is no down-time moving from a single-AZ to multi-AZ. AWS takes a snapshot of the database, restores it in another Availability Zone, and then synchronizes them again for any recent changes.

failover from the primary to the standby instance will happen when one of the following has happend to the primary instance.

> - Failed
> - OS is undergoing software patches
> - Unreachable due to loss of network connectivity
> - Modified (e.g., DB instance type changed)
> - Busy and unresponsive
> - Underlying storage failure
> - An Availability Zone outage
> - A manual failover of the DB instance was initiated using Reboot with failover.

</details>


### RDS Proxy

<details>
<summary>
Having Lambda functions connect to RDS without creating too many connections.
</summary>

> By default, your <cloud>Lambda</cloud> function is
launched outside your own <cloud>VPC</cloud> (in
an AWS-owned VPC). Therefore, it cannot access resources in your VPC (<cloud>RDS</cloud>, <cloud>ElastiCache</cloud>,
internal <cloud>ELB</cloud>...).\
> to run a lambda in a VPC,  You must define the VPC ID, the Subnets and the Security Groups, so  Lambda will create an ENI (Elastic Network Interface) in your subnets.
>
> this requires the IAM AWSLambdaVPCAccessExecutionRole.

however, this only works for a small number of connections. since each lambda crates a connection to the database, it's easy to reach a "TooManyConnections" error state.\
The RDS Proxy manages the connection pool and cleaning up idle connections. we no longer to write code for this. we can deploy it an the same subnet or a different one, the lambdas will connect to the Proxy and all share the same connection pool. the proxy itself can auto-scale up.\
The lambda function must have a connection to the subnet in which the proxy resides. if the RDS subnet allows public connections, there is no need for special configuration. but if the proxy is in a subnet with restrictions on incoming traffic, then the lambdas should also be deployed in a subnet that can connect to it.

in the <cloud>RDS</cloud> service, we can select <kbd>Create Proxy</kbd>, select the engine (MySql, PostgreSQL), and we can set timeouts for idle connections. we need a secret from the secret manager to allow the proxy to connect to the database, this also requires an IAM role. we can use the connection URL in our application client.

</details>

### RDS Parameter Groups

<details>
<summary>
Customize Database configuration with Parameter groups.
</summary>

> You can configure the DB engine using Parameter Groups
> - Dynamic parameters are applied immediately
> - Static parameters are applied after instance reboot
> - You can modify parameter group associated with a DB (must reboot)
> - See documentation for list of parameters for a DB technology
> 
> Must-know parameter:
> - PostgreSQL / SQL Server: rds.`force_ssl=1` => force SSL connections
> - MySQL / MariaDB: `require_secure_transport=1` => force SSL connections

in the RDS service, under "Parameter group", we can view the existing groups or <kbd>Create parameter group</kbd>. a parameter group applies to a DB engine and a version, there is a pre-defined list of parameters which can be modified to suit our needs. each parameter can be dynamic (applied immediately) or static (applied on instance reboot).\
in the database options, we can assign the parameter group to the database, so it will use the defaults from there.

</details>

### RDS Backups and Snapshots

<details>
<summary>
Backup and Snapshots
</summary>

restoring from an automated backup or a snapshot creates a new DB instance. Backups can't be shared, but snapshots can be. backups happen during maintenance windows, and don't interrupt the performance of the database. snapshots take I/O and effect the performace. snapshots can be done manually or in a schedule.

> RDS Backup:
> 
> - Backups are "continuous" and allow point in time recovery
> - Backups happen during maintenance windows
> - When you delete a DB instance, you can retain automated backups
> - Backups have a retention period you set between 0 and 35 days
> - To disable backups, set retention period to 0
>
> RDS Snapshot:
> 
> - Snapshots takes IO operations and can stop the database from seconds to minutes
> - Snapshots taken on Multi AZ DB don't impact the master - just the standby
> - Snapshots are incremental after the first snapshot (which is full)
> - You can copy & share DB Snapshots
> - Sharing:
>   - Manual snapshots: can be shared with AWS accounts
>   - Automated snapshots: can't be shared, copy first
>   - You can only share unencrypted snapshots and snapshots encrypted with a customer managed key
>   - If you share an encrypted snapshots, you must also share any customer managed keys used to encrypt them
> - Manual Snapshots don't expire
> - You can take a "final snapshot" when you delete your DB

</details>

### RDS Monitoring

<details>
<summary>
Events, Logs, Metrics and insights
</summary>

#### RDSEvents and Logs

> RDS keeps record of events related to:
> 
> - DB instances (state)
> - Snapshots
> - Parameter groups, security groups, etc..

we can subscribe to those events using an <cloud>SNS</cloud> topic filtered on the event source (which instance or resource) and category (what kind of event). the other option is to use <cloud>EventBridge</cloud> and set rules about them.\
for example, we can set a rule about database snapshots, and have some actions associated with it.

The DB instance creates some logs while running:

- general
- audit
- error
- slow queries

we can view the events and logs from the "events and logs" tab of our instances. if we want all the logs, we need to enable <cloud>Log exports</cloud> to have them stored in <cloud>CloudWatch</cloud>.\
we can send them into <cloud>cloudWatch Logs</cloud> and apply filters (like for errors) and add <cloud>CloudWatch Alarms</cloud> to be notified about them.

#### RDS & CloudWatch Metrics

there are two levels of metrics that RDS monitors, we can use them for monitoring, trouble-shooting and for alerts. enhanced tier uses an agent on the instance.

> Basic CloudWatch metrics associated with RDS (gathered from the hypervisor):
> 
> - DatabaseConnections
> - SwapUsage
> - ReadIOPS / WriteIOPS
> - ReadLatency / WriteLatency
> - ReadThroughPut / WriteThroughPut
> - DiskQueueDepth
> - FreeStorageSpace
> 
> Enhanced Monitoring (gathered from an agent on the DB instance). Useful when you need to see how different processes or threads use the CPU.
> - Access to over 50 new CPU, memory, file system, and disk I/O metrics

we enable the enhanced monitoring for the instance, and we choose the granularity (what frequency the metrics are sent). to view them we need to select "enhanced monitoring" on the monitoring tab of the database.

#### RDS Performance Insights

we can enable performance insights for our database, and then get a visualized view of what affects it.

> Visualize your database performance and analyze any issues that affect it.\
> DBLoad = the number of active sessions for the DB engine.
> With the Performance Insights dashboard, you can visualize the database load and filter it:
> 
> - By Waits => find the resource that is the bottleneck (CPU, IO, lock, etc...)
> - By SQL statements => find the SQL statement that is the problem
> - By Hosts => find the server that is using the most our DB
> - By Users => find the user that is using the most our DB
>
> You can view the SQL queries that are putting load on your database.

</details>

###  Amazon Aurora

<details>
<summary>
AWS Cloud-Optimized Database.
</summary>

> <cloud>Aurora</cloud> is a proprietary technology from AWS (not open sourced).
> - Postgres and MySQL are both supported as Aurora DB (that means your drivers will work as if Aurora was a Postgres or MySQL database).
> - Aurora is "AWS cloud optimized" and claims 5x performance improvement over MySQL on RDS, over 3x the performance of Postgres on RDS.
> - Aurora storage automatically grows in increments of 10GB, up to 128 TB.
> - Aurora can have up to 15 replicas and the replication process is faster than MySQL (sub 10 ms replica lag).
> - Failover in Aurora is instantaneous. It's HA (High Availability) native.
> - Aurora costs more than RDS (20% more) - but is more efficient.

It stores the data in six copies accross 3 Availability Zones. and it requires only 4 for writes and 3 copies for read (so it can handle losing some instances). it has peer-to-peer replication for self-healing, and storage is stripped across 100's of volumes.\
there is a single instance that accepts writes, and if it doesn't respond, there is an automatic failover that takes up to 30 seconds. for read performance, Aurora supports up to 15 read replica instances (any of them can be promoted), with auto scaling to increase or decrease the number. cross-region replication is supported natively.

Aurora works as a cluster, it has a shared storage volume, and only one writer at a time, the writer is accessed through a *write EndPoint*. since there are so many read replicas, there is also a read *EndPoint*, which acts as a load balancer for read requests.

> The features of <cloud>Aurora</cloud> are:
>
> - Automatic fail-over
> - Backup and Recovery
> - Isolation and security
> - Industry compliance
> - Push-button scaling
> - Automated Patching with Zero Downtime
> - Advanced Monitoring
> - Routine Maintenance
> - Backtrack: restore data at any point of time without using backups

hands on:\

we <kbd>Create new Database</kbd>, and choose the Aurora option with MySQL compatibility. we can filter for versions that have support for aurora features. we can choose a template like before, and modify the configuration. we can use standard storage of I/O optimized, like before, we select which instance type we need, or use the "serverless v2" option to allow AWS to manage the instances as well. we can choose to create replicas if we want, and like before, decide on connectivity we want (which VPC, security group, etc...). we have monitoring options, parameter groups, deletion protection, and the option to allow write forwarding (forward write requests from the reader instances). we have the writer and reader endpoints, but also an endpoint for each instance.\
We can create a policy for read replica and have them scale based on target metric (scale in and scale out). we can add Regions to our cluster (if we used a compatible instance) to make it truly global.

to delete the database, we first need to the delete the reader replicas, and only then the database itself.

#### Aurora Backups, Restore and Backtracking

we can do backups like other databases, but also back-track the existing cluster to a previous place in time without creating new instances. and we can clone the database and continue using the same shared storage for data that wasn't changed in the new environment yet.

> Automatic Backups
> - Retention period 1-35 days (can't be disabled)
> - PITR, restore your DB cluster within 5 minutes of the current time
> - Restore to a new DB cluster
> 
> Aurora Backtracking
> - Rewind the DB cluster back and forth in time (up to 72 hours)
> - Doesn't create a new DB cluster (in-place restore)
> - Supports Aurora MySQL only
> 
> Aurora Database Cloning
> - Creates a new DB cluster that uses the same DB cluster volume as the original cluster
> - Uses copy-on-write protocol (use the original/single copy of the data and allocate storage only when changes made to the data)
>   - Example: create a test environment using your production data

#### Aurora Special Features

Each read replica can have a priority rating (0-15), this controls the failover priority (which replica will be promoted), if two instances have the same priority, the instance with the larger size is promoted.\
We can migrate RDS MySql into RDS Aurora by taking a snapshot and restoring it.

there are some <cloud>CloudWatch</cloud> metrics regarding aurora:

> - `AuroraReplicaLag`: amount of lag when replicating updates from the primary instance.\
> If replica lag is high, that means the users willhave a different experience based on which replica they get the data from (eventual consistency)
>   - `AuroraReplicaLagMaximum`: max. amount of lag across all DB instances in the cluster
>   - `AuroraReplicaLagMinimum`: min. amount of lag across all DB instances in the cluster
> - `DatabaseConnections`: current number of connections to a DB instance
> - `InsertLatency`: average duration of insert operations
</details>

### RDS Security Options

<details>
<summary>
Security options for datbase
</summary>

> - At-rest encryption:
> - Database master & replicas encryption using AWS <cloud>KMS</cloud> - must be defined as launch time
> - If the master is not encrypted, the read replicas cannot be encrypted
> - To encrypt an un-encrypted database, go through a DB snapshot & restore as encrypted
> - In-flight encryption: TLS-ready by default, use the AWS TLS root certificates client-side
> 
> - IAM Authentication: IAM roles to connect to your database (instead of username/pw)
> - Security Groups: Control Network access to your RDS / Aurora DB
> - No SSH available except on RDS Custom
> - Audit Logs can be enabled and sent to CloudWatch Logs for longer retention

</details>

### ElastiCache Overview

<details>
<summary>
Managed Cache Services: Redis and Memcached.
</summary>

> The same way RDS is to get managed Relational Databases - <cloud>ElastiCache</cloud> is to get managed Redis or Memcached.\
> 
> - Caches are in-memory databases with really high performance and low latency
> - Helps reduce load off of databases for read intensive workloads
> - Helps make your application stateless
> - AWS takes care of OS maintenance / patching, optimizations, setup, configuration, monitoring, failure recovery and backups
> - Using ElastiCache involves heavy application code changes

when using cache, the application first queries the cache, and retrives the data from it if exists (cache hit), or from the database itself if not (cache miss). the cache needs to be invalidated after a time to ensure the most recent data is being used.\
Caching is also used to store user session data and make the application stateless - no matter which instance handles the request, it will grab the data from the cache and behave as if it has a state.

There are two optios, **Redis** for high availability, and **Memcached** for multi-node partitoning and multi-threading

> REDIS:
> 
> - Multi AZ with Auto-Failover
> - Read Replicas to scale reads and have high availability
> - Data Durability using AOF persistence
> - Backup and restore features
> - Supports Sets and Sorted Sets
> 
> MEMCACHED
> 
> - Multi-node for partitioning of data (sharding)
> - No high availability (replication)
> - Non persistent
> - No backup and restore
> - Multi-threaded architecture

HandsOn:\
In the <cloud>ElasticCache</cloud> service, wee can create a cache (either Redis or Memcached), we will use Redis. we choose the configuration, either using serverless of defining our own instance. we can create the cache from zero or restore it from a backup (redis only).\
we can set our cache to be replicated across multiple shards, or use on primary node with read replicas. we can also have our cache be hosted on-premises with <cloud>AWS Outposts</cloud>. we can use multi-AZ configuration and set automatic failover. like before, we set the vpc, subnet and security groups, we can add encryption (at rest and in-flight), and control backup (again, redis only) and maintenance windows.

#### ElastiCache Redis

there are two modes for <cloud>ElasticCache Redis</cloud>, with and without cluster-mode (Sharding).

when the cluster mode is disabled, there is a single primary node for read and writes, and replica nodes for reads.

> - One primary node, up to 5 replicas
> - Asynchronous Replication
> - The primary node is used for read/write
> - The other nodes are read-only
> - One shard, all nodes have all the data
> - Guard against data loss if node failure
> - Multi-AZ enabled by default for failover
> - Helpful to scale read performance
> - Vertical and horizontal scaling
>   - vertical scaling creates a new node group

when cluster mode is enabled, we replicate the above behavior across shards (partitions of the data)

> - Data is partitioned across shards (helpful to scale writes)
> - Each shard has a primary and up to 5 replica nodes (same concept as before)
> - Multi-AZ capability
> - Up to 500 nodes per cluster
> - supports autoscaling for shards and replicas
>   - Target Tracking
>   - Scheduled Scaling

the application must use the correct EndPoint.

> Standalone Node:
> - One endpoint for read and write operations
> 
> Cluster Mode Disabled Cluster:
> 
> - Primary Endpoint - for all write operations
> - Reader Endpoint - evenly split read operations across all read replicas
> - Node Endpoint - for read operations
> 
> Cluster Mode Enabled Cluster:
> 
> - Configuration Endpoint - for all read/write operations that support Cluster Mode Enabled Commands
> - Node Endpoint - for read operations

when using Redis with cluster mode disabled (single shard), horizontal scaling adds on removes replica nodes, up to a maximum of 5. vertical scaling creates a new node group with the required instance type, and when all the data is replicated, the DNS is updated and the old instances are terminated.

with cluster mode enabled (multiple shards), there is online scaling - which continues to serve requests during the scaling process, with some performance hit. and offline scaling, which has downtime, but allows for more configuration changes (node type, engine version, etc..,).\
Horizontal scaling can involve adding and removing shards, or with *Shard Rebalancing*, which re-distributes the keyspaces among the shards for better load distribution, horizontal scaling supports both online and offline scaling.\
Vertical Scaling (change read/write capacity) supports online scaling as well.

there are metrics we can monitor:

> - `Evictions`: the number of non-expired items the cache evicted to allow space for new writes (memory is overfilled).\
> Solution:
>   - Choose an eviction policy to evict expired items (e.g., evict least recently used (LRU) items)
>   - Scale up to larger node type (more memory) or scale out by adding more nodes
> - `CPUUtilization`: monitor CPU utilization for the entire host.\
> Solution: scale up to larger node type (more memory) or scale out by adding more nodes
> - `SwapUsage`: should not exceed 50 MB.\
>  Solution: verify that you have configured enough reserved memory
> - `CurrentConnections`: the number of concurrent and active connections.
> Solution: investigate application behavior to address the issue
> - `DatabaseMemoryUsagePercentage`: the percentage of memory utilization
> - `NetworkBytesIn/Out` & `NetworkPacketsIn/Out`
> - `ReplicationBytes`: the volume of data being replicated
> - `ReplicationLag`: how far behind the replica is from the primary node

#### ElastiCache Memcached

Memcached Cluster can have between 1 to 40 nodes (soft limit), horizontal scaling will add and remove nodes, and the *auto-discovery* mechanism will allow the app to find them. vertical scaling means that we create new cluster with the new node type, but our application needs to update the endpoint to use the new one (not automatically), since there is no backup mechanism, the new cluster and the nodes start empty.

> Memcached Auto Discovery
>
> - Typically you need to manually connect to individual cache nodes in the cluster using its DNS endpoints
> - Auto Discovery automatically identifies all of the nodes
> - All the cache nodes in the cluster maintain a list of metadata about all other nodes
> - This is seamless from a client perspective

metrics to monitor:

> - `Evictions`: the number of non-expired items the cache evicted to allow space for new writes (memory is overfilled).\
> Solution:
>   - Choose an eviction policy to evict expired items (e.g., LRU items)
>   - Scale up to larger node type (more memory) or scale out by adding more nodes
> - `CPUUtilization`- Solution: scale up to larger node type or scale out by adding more nodes
> - `SwapUsage`: should not exceed 50 MB
> - `CurrentConnections`: the number of concurrent and active connections. Solution: investigate application behavior to address the issue
> - `FreeableMemory`: amount of free memory on the host
</details>

</details>

## Monitoring, Auditing and Performance

<details>
<summary>
Monitoring Services in AWS
</summary>

<cloud>CloudWatch</cloud> for logs, metrics and alarms. <cloud>EventBridge</cloud>, <cloud>CloudTrail</cloud> to audit API calls into AWS services

### CloudWatch Service

<details>
<summary>
Metrics, Logs, Alarms and Dashboards.
</summary>

The most common service that holds metrics and logs for other services, and can define alarms and dashboards.

#### CloudWatch Metrics

in the CloudWatch service, we can see all the types of the metrics by service, and look at all the metrics over a range of time.

> CloudWatch provides metrics for every services in AWS:
> 
> - Metric is a variable to monitor (CPUUtilization, NetworkIn…)
> - Metrics belong to namespaces
> - Dimension is an attribute of a metric (instance id, environment, etc…).
> - Up to 30 dimensions per metric
> - Metrics have timestamps
> - Can create CloudWatch dashboards of metrics

for example, <cloud>EC2</cloud> instances create metrics every five minutes, but we can use "detailed" monitoring to get a higher frequency of metrics. we can use this toegether with AutoScaling groups and <cloud>CloudWatch</cloud> alarms to react faster and scale more quickly.\
EC2 memory usage (RAM) isn't one of the default metrics, it must be pushed from inside the instace as a custom metric.

we can define custom metrics (RAM, diskspace, number of logged-in users) and push them to <cloud>CloudWatch</cloud> with the `PutMetricData` API call. metrics have dimensions (key-value pairs that identify the metric), and the resolution, either the standard resolution of once every 60 seconds, or higher resolution: [1,5,10,30], this has an increased cost.\
Custom metrics can be pushed with a delay of two weeks in the past, and up to two hours in the future, so we can push recent backlog data, and we need to be aware of any time differences and how our machine is configured. we can use the unified cloudWatch agent to manage pushing the metrics from our machines.

#### CloudWatch Dashboards

we can take the metrics and display them in a dashboard.

> - Great way to setup custom dashboards for quick access to key metrics and alarms
> - Dashboards are global
> - Dashboards can include graphs from different AWS accounts and regions
> - You can change the time zone & time range of the dashboards
> - You can setup automatic refresh (10s, 1m, 2m, 5m, 15m)
> - Dashboards can be shared with people who don't have an AWS account (public, email address, 3rd party SSO provider through <cloud>Amazon Cognito</cloud>)

we can click <kbd>CreateDashboard</kbd> or use the automatic dashboard templates based on the service. for example, we can use the auto-scaling dahsboard and all the metrics will be shown in widgets (panel). we can add widget for metric data or logs. we can choose metrics from any region and from multiple accounts.

#### CloudWatch Logs

when our instances write logs, they usually end up in cloudWatch Logs service.

> - Log groups: arbitrary name, usually representing an application
> - Log stream: instances within application / log files / containers
> - Can define log expiration policies (never expire, 1 day to 10 years...)
> - CloudWatch Logs can send logs to:
>   - Amazon <cloud>S3 bucket</cloud> - exports
>   - <cloud>Kinesis Data Streams</cloud>
>   - <cloud>Kinesis Data Firehose</cloud>
>   - AWS <cloud>Lambda</cloud>
>   - <cloud>OpenSearch</cloud>
>
> - Logs are encrypted by default,
> - Can setup KMS-based encryption with your own keys

common sources for logs:

> - SDK, <cloud>CloudWatch Logs Agent</cloud>, <cloud>CloudWatch Unified Agent</cloud>
> - <cloud>Elastic Beanstalk</cloud>: collection of logs from application
> - <cloud>ECS</cloud>: collection from containers
> - AWS <cloud>Lambda</cloud>: collection from function logs
> - VPC <cloud>Flow Logs</cloud>: VPC specific logs
> - <cloud>API Gateway</cloud>
> - <cloud>CloudTrail</cloud> based on filter
> - <cloud>Route53</cloud>: Log DNS queries

we can use Logs Insight to query the logs, we can use this same query as part of a dashboard. we can filter, group, aggregate and do all the usual stuff. it's a query engine, but not real-time.

One of the options is to export Logs into S3 buckets, this is done via the `CreateExportTask` API call, and can take up to 12 hours. if we want real-time (or near real-time), we use Logs Subscriptions, the logs are sent to a downstream service for further processing.

possible targets are:

- <cloud>Kinesis Data Streams</cloud>
- <cloud>Kinesis Data Firehose</cloud>
- <cloud>Lambda function</cloud>

we can also add filters on the subscriptions, to select which logs are delievered for processing. Log subscriptions can be aggregated across multiple regions and multiple accounts, so all logs (even from different sources) will reach a single location. for this we need to use cross account subscriptions, a filter in the source account (sender), and a subscription destination in the target account (receiver). the recipient account has a destination access policy which allows the sender account to send logs to it. 

IAM Role in the Recipient stream (receiver), needs to be assumed by the sender.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "kinesis:PutRecord",
      "Resource":"arn:aws:kinesis:us-eas-1:<accountId>:stream/RecipientStream"
    }
  ]
}
```

destination access policy:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal":
      {
        "AWS":"<senderAccount>"
      },
      "Action": [
        "logs:PutSubscriptionFilter"
      ],
      "Resource":"arn:aws:logs:us-east-1:<receiverAccount>:destination:testDestination"
    }
  ]
}
```

when we look at the logs, we can <kbd>Create Metric Filter</kbd> to have predefined searches that will create metrics from the logs. then we can create alarms on the metric, meaning that we managed to set an alarm on textual data.

we can set the retention period of a log group, create new log groups (which we could send data to), and use the logs insight to better understand how to query logs.

there is also an option called "live-tail", which gives us real-time view of logs. this is helpful for debugging, but has additional costs.

#### CloudWatch Alarms

we can set alarms to respond to metrics, common use cases are acting on <cloud>EC2</cloud> machines (rebooting, terminating, etc), auto scaling actions (removing and adding instances), or send a <cloud>SNS</cloud> notification, which can then trigger a custom lambda.

> - Alarms are used to trigger notifications for any metric
> - Various options (sampling, %, max, min, etc...)
> - Alarm States:
>   - `OK`
>   - `INSUFFICIENT_DATA`
>   - `ALARM`
> 
> - Period:
>   - Length of time in seconds to evaluate the metric
>   - High resolution custom metrics: 10 sec, 30 sec or multiples of 60 sec

Alarms are usually on a single metric, but we can also have composite alarms, which apply on top of other alarms, using boolean conditions (AND, OR) to create complex rules to reduce noise.

> - Alarms can be created based on CloudWatch Logs Metrics Filters
> - To test alarms and notifications, set the alarm state to Alarm using CLI

```sh
aws cloudwatch set-alarm-state --alarm-name "myAlarm" --state-value ALARM --state-reason "testing purposes"
```

in the <cloud>CloudWatch</cloud> service, we select <kbd>Create Alarm</kbd>, we choose the EC2 instance, and the `CPU Utilization` metric. we select the 'average' statictics, choose 5 minutes period as data point, and create the condition for the alarm - if this average is larger than 95 for more than 3 consecutive data points.\
we choose the EC2 action to terminate the instance on the alarm state. we can trigger the behavior with a custom script to push the cpu, or use the API call to set the alarm state, so we would see the alarm behavior in action.

#### CloudWatch Synthetics Canary

a script (like automation), that tests behavior of our websites or APIs and can trigger an alert if something fails.

> - Configurable script that monitor your APIs, URLs, Websites, etc..
> - Reproduce what your customers do programmatically to find issues before customers are impacted
> - Checks the availability and latency of your endpoints and can store load time data and screenshots of the UI
> - Integration with CloudWatch Alarms
> - Scripts written in Node.js or Python
> - Programmatic access to a headless Google Chrome browser
> - Can run once or on a regular schedule

we can use some bluprints, which are templates for common tests:

> - Heartbeat Monitor - load URL, store screenshot and an HTTP archive file
> - API Canary - test basic read and write functions of REST APIs
> - Broken Link Checker - check all links inside the URL that you are testing
> - Visual Monitoring - compare a screenshot taken during a canary run with a baseline screenshot
> - Canary Recorder - used with CloudWatch Synthetics Recorder (record your Actions on a website and automatically generates a script for that)
> - GUI Workflow Builder - verifies that actions can be taken on your webpage (e.g., test a webpage with a login form)
</details>

### Amazon EventBridge

<details>
<summary>
React to AWS events from a centralized location.
</summary>

Formerly known as <cloud>CloudWatch Events</cloud>.

> - Schedule: Cron jobs (scheduled scripts)
> - Event Pattern: Event rules to react to a service doing something
> 
> Trigger Lambda functions, send SQS/SNS messages...

sources are events from services, including <cloud>CloudTrail</cloud>, which effectively means we can intercept any kind of event based on the api call.

we can filter events based on conditions, and then send the event as json to other destinations:

- <cloud>Lambda</cloud>
- <cloud>AWS Batch</cloud>
- <cloud>ECS</cloud> Task
- <cloud>SQS</cloud>
- <cloud>SNS</cloud>
- <cloud>Kinesis Data Streams</cloud>
- <cloud>Step Function</cloud>
- <cloud>CodePipeline</cloud>
- <cloud>CodeBuild</cloud>
- <cloud>SSM</cloud> (documents, run commands)
- <cloud>EC2 Action</cloud>

<cloud>Event Bridge</cloud> can also integrate with AWS SaaS Parteners, such as:

- Zendesk
- DataDog
- Saleforce
- others...

those are referred to as **Partner Event Bus** (opposed to the **Default Event Bus** for aws services). they can also send events to EventBridge and trigger events. we can also create **Custom Event Bus** for our own applications.

Events are in json format, and have a schema, from the schema, we can generate code for our applications to properly handle the data. the schemas are versioned.\
events can be archived (all, filtered), and we can replay the achieved events.

the event buses can be accessed by other AWS account using resource-based policies. for example, we can aggregate all the events from many AWS accounts into a single event bus.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "EventBridge",
      "Effect": "Allow",
      "Principal": {"AWS": "<otherAccount>"},
      "Action": [
        "events::PutEvents"
      ],
      "Resource": "arn:aws:events:us-east-1:123456789012:event-bus/central-event-bus"
    }
  ]
}
```


Hands On:\

in the <cloud>Event Bridge</cloud> service, we can click on the "Event Bus" section and see the default event bus, or <kbd>Create Event Bus</kbd> if we want a cross-account central event bus.\
We can see the event sources parterns and follow the instructions to send them to one of our buses.

We can <kbd>Create Rule</kbd> and either set it to run it as cron job (scheduled) or to respond to events in one of our event buses. the source can be an aws-service events, custom events, or all events.

there is a sandbox to test a patterns against events, we can use the sample events to create rules with event patterns.

```json
{
  "source": ["aws.ec2"],
  "detail-type": ["EC2 Instance State-Change Notification"],
  "detail":{
    "state": ["stopped", "terminated"]
  }
}
```

we then set up the target (destination), which is either an aws service, an API destination, or another event bus (which can be in a different account). a common destination is an <cloud>SNS</cloud> Topic, which many services can subscribe to.

under the "schemas" section, we can see the schema for many events (what fields exist) and <kbd>Download Code Bindings</kbd> to generate code for the specific event.

#### EventBridge Content Filtering

we can filter events to only react to some events which match the event pattern. for example, we can set the rule to only respond to events from a specific bucket, either exactly or with a prefix/suffix.


```json
{
  "source":["aws.s3"],
  "detail-type": "Object Created",
  "detail":{
    "bucket":{
      "name":[{"prefix":"myBucket"}]
    }
  },
  "Object":{
    "key":[{"suffix":".png"}]
  },
  "source-ip-address":[{"cidr":"10.0.0.0/24"}]
}
```

#### EventBridge Input Transformation

We can also transform the input to modify the event before sending it to the service. for example, we can choose to send a log to a specific <cloud>CloudWatch</cloud> log group to keep a record when an <cloud>EC2</cloud> machine starts running. we can choose <kbd>Configure input Transformer</kbd> and set a transfomer which takes the input path and converts it into a template.

this input path defines 4 variables from the original events.

```json
{
  "timestamp" : "$.time",
  "instance" : "$.detail.instance-id", 
  "state" : "$.detail.state",
  "resource" : "$.resources[0]"
}
```

and we use those variables to create a new message through our template:

```json
{
  "timestamp": <timestamp>,
  "message": "<instance> is in state <state> at <timestamp>, with arn <resource>."
}
```

now the thing we send forward to the target log group is the new transformed object, and not the original thing.

</details>

### Service Quotas Overview

<details>
<summary>
Monitor AWS Service Quotas
</summary>

> Notify you when you're close to a service quota value threshold
> - Create CloudWatch Alarms on the Service Quotas console
> - Example: Lambda concurrent executions
> - Helps you know if you need to request a quota increase or shutdown resources before limit is reached
> 
> Alternative: Trusted Advisor + CW Alarms:
> 
> - Limited number of Service Limits checks in Trusted Advisor (~50)
> - Trusted Advisor publishes its check results to CloudWatch
> - You can create CloudWatch Alarms on service quota usage (Service Limits)

In the <cloud>Service Quota</cloud> service dashboard, we can find all the quota, see the default value, create alarms and monitoring, and request quota increase for resources which are adjustable.

in <cloud>Trusted Advisor</cloud>, under "Service Limits", we can use pre-configured checks for some of the services.
</details>



### CloudTrail

<details>
<summary>
Audit AWS API actions
</summary>

> Provides governance, compliance and audit for your AWS Account:
> 
> - CloudTrail is enabled by default!
> - Get an history of events / API calls made within your AWS Account by:
>   - Console
>   - SDK
>   - CLI
>   - AWS Services
> 
> - Can put logs from CloudTrail into CloudWatch Logs or S3
> - A trail can be applied to All Regions (default) or a single Region.
> - If a resource is deleted in AWS, investigate CloudTrail first!

there are "management events" - which are actions performed on resources, which are then divided into read events (which don't modify the resources) and write events (which do). management events are stored by default into the trail.\
in contrast, there are also "data events" - they aren't stored by default, and they represent high-volume events, like <cloud>S3</cloud> objects events or <cloud>Lambda</cloud> invocations. we can still choose to store the events into the trail.

finally, <cloud>CloudTrail</cloud> **Insights** events are a special feature (we need to enable it) that creates a baseline of operations in the account, and then analyzes write events to detect unusual patterns. 

> - Enable CloudTrail Insights to detect unusual activity in your account:
> 
> - inaccurate resource provisioning
> - hitting service limits
> - Bursts of AWS IAM actions
> - Gaps in periodic maintenance activity
> 
> CloudTrail Insights analyzes normal management events to create a baseline. And then continuously analyzes write events to detect unusual patterns:
> 
> - Anomalies appear in the CloudTrail console
> - Event is sent to Amazon S3
> - An EventBridge event is generated (for automation needs)

Events are stored for 90 days by default, but we can send them to <cloud>S3</cloud> to store them for a longer time, and then use <cloud>Athena</cloud> to query them.

Hands On:\
in the <cloud>CloudTrail</cloud> service, we can see the events history, if we terminate an <cloud>EC2</cloud> machine we can see the termination event in the console.

#### EventBridge Integration And Other Stuff

we can integrate <cloud>CloudTrail</cloud> with <cloud>EventBridge</cloud>, we can set up rules to listen to specific API calls and respond to them.

- when a <cloud>DynamoDB</cloud> table is deleted
- when a user assumes a different role
- when a user edits the <cloud>EC2</cloud> security groups inbound rules.

> CloudTrail is not "real-time":
> 
> - Delivers an event within 15 minutes of an API call
> - Delivers log files to an S3 bucket every 5 minutes

Organizations Trails can store all events from different aws accounts in the same AWS Organization.

> - A trail that will log all events for all AWS accounts in an AWS Organization
> - Log events for management and member accounts
> - Trail with the same name will be created in every AWS account (IAM permissions)
> - Member accounts can't remove or modify the organization trail (view only)

Log File Integrity Validation - Create a hash of all the events being stored in <cloud>S3</cloud> to make sure they weren't tempered with.

> Digest Files:
> - References the log files for the last hour and contains a hash of each
> - Stored in the same S3 bucket as log files (different folder)
> 
> Helps you determine whether a log file was modified/deleted after CloudTrail delivered it.
> - Hashing using SHA-256, Digital Signing using SHA- 256 with RSA
> 
> - Protect the S3 bucket using bucket policy, versioning, MFA Delete protection, encryption, object lock
> - Protect CloudTrail using IAM

</details>


### AWS Config

<details>
<summary>
Monitor and detect changes in AWS resource configurations.
</summary>

> Helps with auditing and recording compliance of your AWS resources. Helps record configurations and changes over time\
> - Questions that can be solved by AWS Config:
>   - Is there unrestricted SSH access to my security groups?
>   - Do my buckets have any public access?
>   - How has my ALB configuration changed over time?
> - You can receive alerts (SNS notifications) for any changes
> - AWS Config is a per-region service
> - Can be aggregated across regions and accounts
> - Possibility of storing the configuration data into S3 (analyzed by <cloud>Athena</cloud>)

<cloud>AWS Config</cloud> has predefined rules, or we can create our own rules using <cloud>AWS Lambda</cloud>. rules can be triggered at each config change or at regular time intervals.\

the cost is based on configuration item recorded (per region) and for each rule evaluation (per region). no free-tier.

we can view the configuration of a resource and how it changed over time, and we can link it to <cloud>CloudTrail</cloud> logs which effected it.

The rules are compliance only, they don't have prevent the changes from happening, they only detect and notify on them. but we can set up "remediation" action using <cloud>SSM Automation Document</cloud> that will be triggered. we can use the pre-configured documents or custom ones, and we can use a custom lambda.

we can also integrate <cloud>AWS Config</cloud> with <cloud>Event Bridge</cloud> and respond to non-compliant configuration events through them (send message to <cloud>SNS</cloud> topic and respond from there).

#### AWS Config Hands On

in the <cloud>AWS Config</cloud> service, we can choose to record all resource types or just some of them, and we can choose to include or exclude global resources. we need an <cloud>IAM</cloud> role and a <cloud>S3</cloud> bucket to store the configurations, and we can set a <cloud>SNS</cloud> target.\
There is a list of pre-defined rules which we can use to mark resources as non-compliant. once the operation is completed, we can see the list of all resources, and see the <kbd>Resource Timeline</kbd> to view all the linked <cloud>CloudTrail</cloud> events which effected the configuration.

we can <kbd>Add Rule</kbd>, either from the managed rules or a custom rule, for example, we can use the managed rule that controls which AMIs are allowed for <cloud>EC2</cloud> machines. we can also set the "restricted-ssh" rule that will mark security groups if they allow ssh access. if we delete the inbound access from the security group, the rule will be evaluated again and be marked as compliant.\
in the rule, we can <kbd>Action</kbd> and <kbd>Manage Remediation</kbd> to choose what happens when a non-compliant resource is detected. we can use automatic remediation, choose a pre-configured action with parameters, or use a custom action.

#### AWS Config Aggregators
we can use <kbd>AWS Config</kbd> aggregator to store a view of all the configurations from many AWS accounts. the single accounts is called the aggregator account, with the other accounts being the source accounts.\
The source accounts need to authorize the aggregator account and send the data to it, when using AWS organizations, there is no need for that.\
The configuration and rules are created in each of the source accounts, and are not shared from the aggregator. this is a "view-only" aggregation.\
If we want to deploy rules to multiple accounts, we should use <cloud>CloudFormation</cloud> StackSets.

</details>



### CloudWatch vs CloudTrail vs Config

<!-- <details> -->
<summary>
Comparison between the services.
</summary>

- <cloud>CloudWatch</cloud> - logs, mertics, events, alerts
- <cloud>CloudTrail</cloud> - AWS api and actions
- <cloud>AWS Config</cloud> - resource configurations

> <cloud>CloudWatch</cloud>
> 
> - Performance monitoring (metrics, CPU, network, etc...) & dashboards
> - Events & Alerting
> - Log Aggregation & Analysis
> 
> <cloud>CloudTrail</cloud>
> 
> - Record API calls made within your Account by everyone
> - Can define trails for specific resources
> - Global Service
>
> <cloud>AWS Config</cloud>
> 
> - Record configuration changes
> - Evaluate resources against compliance rules
> - Get timeline of changes and compliance

for example, with an <cloud>Elastic Load Balancer</cloud> resource, each of the services will give us different information:

> - CloudWatch:
>   - Monitoring Incoming connections metric
>   - Visualize error codes as a % over time
>   - Make a dashboard to get an idea of your load balancer performance
> - AWS Config:
>   - Track security group rules for the Load Balancer
>   - Track configuration changes for the Load Balancer
>   - Ensure an SSL certificate is always assigned to the Load Balancer (compliance)
> - CloudTrail:
>   - Track who made any changes to the Load Balancer with API calls

</details>


</details>

## Aws Account Management

<details>
<summary>
Service Health, Multiple Accounts Management and Monitoring Costs
</summary>

Health Dashboards, Billing Console, AWS Organizations

### AWS Health Dashboard

<details>
<summary>
Learn about issues effecting AWS
</summary>

Service health, either for AWS services or for your specific account.

AWS Service History is for tracking issues with AWS globally.

> - Shows all regions, all services health
> - Shows historical information for each day
> - Has an RSS feed you can subscribe to
> - Previously called AWS Service Health Dashboard

AWS Health Dashboard is for when aws has issues that affect resources in the account, it is a global service, and we can see what might effect us in the future and history.

> - AWS Account Health Dashboard provides alerts and remediation guidance when AWS is experiencing events that may impact you.
> - While the Service Health Dashboard displays the general status of AWS services, Account Health Dashboard gives you a **personalized** view into the performance and availability of the AWS services underlying your AWS resources.
> - The dashboard displays relevant and timely information to help you manage events in progress and provides proactive notification to help you plan for scheduled activities.
> - Can aggregate data from an entire AWS Organization
> - Previously called AWS Personal Health Dashboard (PHD)

we can search for issues effecting specific services in a specific region, and look at historical data.

#### AWS Health Dashboard - Events & Notifications

we can use <cloud>EventBridge</cloud> to react to Health Event Notifications, such as getting an email when an instance is set to a scheduled update, or when a resource is potentially effected by an outage in AWS.\
Another example is when there was a leak of <cloud>IAM</cloud> access keys, we can invoke a lambda to delete the keys from our account if they belong to us. or if an AMI or instance type is scheduled for retirement.

</details>

### AWS Organizations and Control Tower

<details>
<summary>
Managing multiple accounts at once
</summary>

Manage multiple accounts, one management and multiple member accounts (each account can only be a member of one organization), allows for consolidated billing across all account - a single payment method.Getting pricing benefits from aggregated usage and sharing reserved instances and saving plans discounts across member accounts.\
Supports an API to automate AWS account creation.

we have a root organization unit which holds the management account, and inside it we can have a hierarchy of organizational units, which can hold member accounts or other organizational units. we can separate the OUs by business unit, project, environment (prod, dev, sandbox).

> Advantages:
> 
> - Multi Account vs One Account Multi VPC
> - Use tagging standards for billing purposes
> - Enable CloudTrail on all accounts, send logs to central S3 account
> - Send CloudWatch Logs to central logging account
> - Establish Cross Account Roles for Admin purposes
> 
> Security: Service Control Policies (SCP)
> 
> - IAM policies applied to OU or Accounts to restrict Users and Roles
> - They do not apply to the management account (full admin power)
> - Must have an explicit allow from the root through each OU in the direct path to the target account (does not allow anything by default - like IAM)

#### AWS Organizations Service Control Policies

in the <cloud>AWS Organization</cloud> service, we <kbd>Create Organization</kbd> from the account we want to be the management account. we can then <kbd>Add Account</kbd> and either create a new account under the organization, or invite an existing account to join the organization. the management organization has full control over the member accounts, until the account decides to leave the AWS organization or the management account removes it.\
From the management account, we can <kbd>Create Organizational Unit</kbd> to partition our members accounts, we can create OUs inside other OUs and for each one we can add member accounts and set policies.

(we can move the management account into an OU, but it's best practice to keep it at the root).

there are different kinds of policies, they are disabled by default.

- AI Services
- Backup Policies
- Service Control Policies
- Tag Policies

Service Control Policies allow us to define what the member accounts can do. there is a default *FullAwsAccess* policy, and we can add explicit deny policies to prevent unwanted behavior. policies are inherited from the containing Organizational units.

#### AWS Organizations Misc

> For billing purposes, the consolidated billing feature of AWS Organizations treats all the accounts in the organization as one account.
> 
> - This means that all accounts in the organization can receive the hourly cost benefit of Reserved Instances that are purchased by any other account.
> - The payer account (master account) of an organization can turn off Reserved Instance (RI) discount and Savings Plans discount sharing for any accounts in that organization, including the payer account.
> - This means that RIs and Savings Plans discounts aren't shared between any accounts that have sharing turned off.
> - To share an RI or Savings Plans discount with an account, both accounts must have sharing turned on.

use the AWS organization id in policy condition instead of specific accounts

> Use `aws:PrincipalOrgID` condition key in your resource-based policies to restrict access to IAM principals from accounts in an AWS Organization.

Tag Policies are used to standardize tags across resources in an AWS Organization
> - Ensure consistent tags, audit tagged resources,maintain proper resources categorization, etc...
> - You define tag keys and their allowed values
> - Helps with AWS Cost Allocation Tags and Attribute based Access Control
> - Prevent any non-compliant tagging operations on specified services and resources (has no effect on resources without tags)
> - Generate a report that lists all tagged/on-compliant resources
> - Use CloudWatch Events to monitor non-compliant tags

#### AWS Control Tower

<cloud>Control Tower</cloud> is a higher layer than <cloud>AWS Organizations</cloud>, everything is configured inside a *Landing Zone*.

> Easy way to set up and govern a secure and compliant multi-account AWS environment based on best practices\
> Benefits:
> 
> - Automate the set up of your environment in a few clicks
> - Automate ongoing policy management using guardrails
> - Detect policy violations and remediate them
> - Monitor compliance through an interactive dashboard
> 
> AWS Control Tower runs on top of AWS Organizations:
> - It automatically sets up AWS Organizations

we can use <cloud>Control Tower</cloud> to deny creation of resources in specific resources, it creates the <cloud>AWS Organization</cloud> with pre-defined OUs and accounts (audit, logs), we can enable <cloud>CloudTrail</cloud> across all accounts, set up centeralized logs, and disable custom <cloud>KMS</cloud> keys use in accounts inside the landing zone.\
We can continue to define stuff in <cloud>Control Tower</cloud>, add users, modify guardrails on the accounts, control organizational units, manage user access, etc...

</details>

### AWS Service Catalog

<details>
<summary>
Simplified AWS Resource creation through a custom AWS portal.
</summary>

> Users that are new to AWS have too many options, and may create stacks that are not compliant / in line with the rest of the organization
> - Some users just want a quick self-service portal to launch a set of authorized products pre-defined by admins
> -  Includes: virtual machines, databases, storage options, etc...
> -  Enter AWS Service Catalog!

Admins create products - <cloud>CloudFormation</cloud> templates, and gather them inside "portfolios" and define <cloud>IAM</cloud> policies of who can access the portfolios.\
Users can view those products in portfolios, and choose which ones to launch, and those services will be properly configured with the correct tags.

catalogs (portfolios) can be shared to other accounts and be used by them.

> Catalog Sharing - Share portfolios with individual AWS accounts or AWS Organizations
> Sharing options:
> - Share a reference of the portfolio, then import the shared portfolio in the recipient account (stays in-sync with the original portfolio)
> - Deploy a copy of the portfolio into the recipient account (must re-deploy any updates)
> - Ability to add products from the imported portfolio to the local portfolio

the portfolios have `TagOptions`, which are inherited to products deployed from the catalog, and make sure the resources are properly tagged.

in the <cloud>Service Catalog</cloud>, we can submit a request to AWS to change the branding of the Service catalog that other accounts will see.\
Under the "administration" section, we can <kbd>Create Product</kbd>, we can use <cloud>Terraform</cloud> or <cloud>CloudFormation</cloud> stack. we can add the product into a portfolio, which we first <kbd>Create Portfolio</kbd>, we can add constraints to on the portfolio, and we need to modify the access to it.\

from a different account, we can navigate to <cloud>Service Catalog</cloud> and use the product to create the stack
</details>

### AWS Billing And Costs

<details>
<summary>
Costs and Pricing.
</summary>

billing data is stored in <cloud>CloudWatch</cloud> in one region (us-east-1), but it contains the costs of all regions.\
we first need to enable billing alerts in the account preferences (settings). in the <cloud>CloudWatch</cloud> service, we can expand the "Alarms" section and choose "billing alarms", and then <kbd>Create Alarm</kbd>. we select the metrics either by service or the total cost. we can set a condition when to trigger the alarm.

#### AWS Cost Explorer

we use the <cloud>Cost Explorer</cloud> to visualize the costs and usage, we can create reports based on region, service, tags and understand how much we spend on AWS services. we can also get forcast estimation based on past usage, and see how saving plans would effect our costs.

#### AWS Budgets

we can create a budget and send alarms based on more granular data than normal cost alarms.

> - Create budget and send alarms when costs exceeds the budget
> - 4 types of budgets: Usage, Cost, Reservation, Savings Plans
> - For Reserved Instances (RI)
>   -  Track utilization
>   -  Supports EC2, ElastiCache, RDS, Redshift
> - Up to 5 SNS notifications per budget
> - Can filter by: Service, Linked Account, Tag, Purchase Option, Instance type, Region, Availability Zone, API Operation, etc...
> - Same options as AWS Cost Explorer!
> - 2 budgets are free, then $0.02/day/budget

we can use pre-defined templates for common use-cases, or customize the budget for specific services, tags, regions, etc...\
We can set the alerts based on thresholds (more than one), and they can be based on actual cost or  forecasted spending.\
We can also have actions taken based on the budget alerts, so we need the proper <cloud>IAM</cloud> role. the role can apply actions such as attaching IAM policies (like deny permissions) or stopping EC2 instances.

#### AWS Cost Allocation Tags & Cost & Usage Reports

> Use cost allocation tags to track your AWS costs on a detailed level
> - AWS generated tags
>   - Automatically applied to the resource you create
>   - Starts with Prefix `aws:` (e.g. aws: createdBy)
> - User-defined tags
>   - Defined by the user
>   - Starts with Prefix `user:`

we can go the billing service and see the tags (aws and user created tags), and we can enable/disable tags to define which tags will be part of the cost allocation reports. we can create the report and have it exported to <cloud>S3</cloud> and then analyze it. we can set integration with services and control the compression of the data.

#### AWS Compute Optimizer 

> Reduce costs and improve performance by recommending optimal AWS resources for your workloads. Helps you choose optimal configurations and rightsize your workloads (over/under provisioned).\
> Uses Machine Learning to analyze your resources configurations and their utilization CloudWatch metrics.\
> Supports some services:
> 
>   - EC2 instances
>   - EC2 Auto Scaling Groups
>   - EBS volumes
>   - Lambda functions
> - Lower your costs by up to 25%
> - Recommendations can be exported to S3

</details>



</details>

## Disaster Recovery

<details>
<summary>
Synchronized data and backups.
</summary>

### AWS DataSync

<details>
<summary>
Synchronize data with AWS from other sources and between AWS services.
</summary>

> Move large amount of data to and from On-premises / other cloud to AWS (NFS, SMB, HDFS, S3 API...) - needs agent. Or move data from AWS to AWS (different storage services) - no agent needed.\
> 
> Can synchronize to:
> 
> - <cloud>Amazon S3</cloud> (any storage classes - including Glacier)
> - <cloud>Amazon EFS</cloud>
> - <cloud>Amazon FSx</cloud> (Windows, Lustre, NetApp, OpenZFS...)
> 
> - Replication tasks can be scheduled hourly, daily or weekly.
> - File permissions and metadata are preserved (NFS POSIX, SMB…)
> - One agent task can use 10 Gbps, can setup a bandwidth limit.

can synchronize two-ways (not just backup to the cloud), <cloud>Snowcone</cloud> devices come pre-installed with the agent.

</details>

### AWS Backup

<details>
<summary>
Central Management of Backups across services.
</summary>

> Fully managed service. Centrally manage and automate backups across AWS services, No need to create custom scripts and manual processes.\
> 
> Supported services:
> 
> - <cloud>Amazon EC2</cloud> / <cloud>Amazon EBS</cloud>
> - <cloud>Amazon S3</cloud>
> - <cloud>Amazon RDS</cloud> (all DBs engines) / <cloud>Amazon Aurora</cloud> / <cloud>Amazon DynamoDB</cloud>
> - <cloud>Amazon DocumentDB</cloud> / <cloud>Amazon Neptune</cloud>
> - <cloud>Amazon EFS</cloud> / <cloud>Amazon FSx</cloud> (Lustre & Windows File Server)
> - AWS <cloud>Storage Gateway</cloud> (Volume Gateway)
> 
> Supports cross-region and cross-account backups.

PITR - point in time recovery, on demand schedules, can backup resources based on tags, controls the backup frequency, retention period, transition to cold storage.\
<cloud>AWS Backup</cloud> can work with Vault Lock and the WORM model (write Once, read Many), which prevents changes and deletions to the data, locks the retention period, and even the root account user can not delete data when the Vault Lock policy is in-place.

hands On:\
in the <cloud>AWS Backup</cloud> service, we can look at the pre-defined template, we can see the backup rule - target, frequency, window for update, retention and transitions.\
once we create the backup plan, we need to assign resources to it, which is both the types of resources and filtering them by tags. we need to provide an <cloud>IAM</cloud> role for the backup task to use.
</details>


</details>

## Security and Compliance for SysOps

<details>
<summary>
Security Tools and Encryption Services
</summary>

AWS uses the Shared Responsibility Model - depending on the service, AWS handles some of the security, and the user handles other parts. The more managed the service, the more AWS is in charge of it.\
At all cases, AWS handles security of the cloud - physical security, infrastructure and networking. the user allways controls the data and access permessions.

The division of responsabilities is different depending on the service, RDS is more managed than EC2, for example.

### DDoS, AWS Shield and AWS WAF

<details>
<summary>
DDoS protection, firewall
</summary>

DDoS - distributed denial of service - an attack by sending a large volume of requests to the application.

<cloud>AWS Shield</cloud> is a service to protect against DDoS attacks. the standard version is enabled without costs. the Advanced version provides a premium protection around at all times.\
<cloud>AWS WAF</cloud> is a firewall service that filters requests based on rules.\
<cloud>CloudFront</cloud> (CDN) and <cloud>Route53</cloud> can be combined with the <cloud>AWS Shield</cloud> service to form a global edge network that is more resilient to attacks.\

(scaling might be required to support all the features and to provide high quality service when there is high demand)

The standard version of <cloud>Shield</cloud> protects against common types of attacks:

- SYN/UDP floods
- Reflection Attacks
- layer 3 and layer 4 common attacks

the advanced version of shield has higher costs:
- provides protection against more complicated attacks
- gives access to a DDoS Respone Team (DRP) for 24/7 support
- reduced fees during usage spike when they are caused by DDoS (if you are attacked, you pay less for for resources that were consumed during that attack).

#### AWS WAF

Web Application Firewall service, protection against layer 7 attacks (HTTP). can only be deployed on services which support HTTP protocol

- Application Load Balancer (<cloud>ELB</cloud>)
- <cloud>API gateway</cloud>
- <cloud>CloudFront</cloud>

we define web ACL rules (access control list):

> - Rules can include: IP addresses, HTTP headers, HTTP body, or URI strings
> - Protects from common attack - SQL injection and Cross-Site Scripting (XSS)
> - Size constraints, geo-match (block countries)
> - Rate-based rules (to count occurrences of events) - for DDoS protection

in the <cloud>WAF</cloud> service, we can <kbd>Create IP set</kbd>. the IP set is a list of ips (or cider), we can later use this set. then we <kbd>Create Web ACL</kbd>, and we attach either <cloud>CloudFront</cloud> distribution or some regional resource. each rule consumes capacity unit, which means there is a limit to how many rules we can apply to a single request.\
There are some predefined rules, each with a know capacity cost. but we can also <cloud>Create Rules</cloud>, and we can create a rule to only allow requests from inside our IpSet. we can also have rate-limit based rules, which can block requests if they exceed a threshold. we can filter them even more. another example is to block SQL injections commands in the request body. we can set the sensitivity level to low or high (which will have more false positives).\
each rule we create will have the capacity cost and priority, we can choose to allow or block requests that don't match any of the rules.\

Rule Groups are a combination of rules which we can attach to a resource together.


</details>


### Penetration Testing on AWS

<details>
<summary>
Attacking our own infrastructure and services
</summary>

Permitted penteration attempts (List can increase over time): 

> AWS customers are welcome to carry out security assessments or penetration tests against their AWS infrastructure without prior approval for 8 services:
>
> - Amazon EC2 instances, NAT Gateways, and Elastic Load Balancers
> - Amazon RDS
> - Amazon CloudFront
> - Amazon Aurora
> - Amazon API Gateways
> - AWS Lambda and Lambda Edge functions
> - Amazon Lightsail resources
> - Amazon Elastic Beanstalk environments

Prohibited penteration attempts

> - DNS zone walking via Amazon Route53 Hosted Zones
> - Denial of Service (DoS), Distributed Denial of Service (DDoS), Simulated DoS, Simulated DDoS
> - Port flooding
> - Protocol flooding
> - Request flooding (login request flooding, API request flooding)

</details>

### Amazon Inspector

<details>
<summary>
Automated Security Assessments
</summary>

automatic security checks - network access, known vunrabilites (package CVEs)

> For <cloud>EC2</cloud> instances
> 
> - Leveraging the <cloud>AWS System Manager</cloud> (SSM) agent
> - Analyze against unintended network accessibility
> - Analyze the running OS against known vulnerabilities
> 
> For Container Images push to Amazon <cloud>ECR</cloud>
> 
> - Assessment of Container Images as they are pushed
> 
> For <cloud>Lambda</cloud> Functions
> 
> - Identifies software vulnerabilities in function code and package dependencies
> - Assessment of functions as they are deployed

reports the findings to the <cloud>Security Hub</cloud> and to <cloud>EventBridge</cloud>.

hands on:\
in the <cloud>Amazon Inspector</cloud> service, we need to enable permissions to the service and <kbd>Enable inspector</kbd>. we can launch an <cloud>EC2</cloud> machine and wait for it to be detected and scanned. we can see the state of the instance, and it might have the "Unmanaged Ec2 Instance" message. we can see the trouble-shooting guide and it will tell us we didn't give the instance the correct permissions. we can fix is through the <cloud>System Manager</cloud> to apply the permissions to all machines.\
we can look at the findings, and see if there are issues like vulnerable OS or possible network access from the public internet.

</details>

### Logging in AWS

<details>
<summary>
Logs and data from security services
</summary>

> To help compliance requirements, AWS provides many service-specific security and audit logs.\
> Service Logs include:
> 
> - CloudTrail trails - trace all API calls
> - Config Rules - for config & compliance over time
> - CloudWatch Logs - for full data retention
> - VPC Flow Logs - IP traffic within your VPC
> - ELB Access Logs - metadata of requests made to your load balancers
> - CloudFront Logs - web distribution access logs
> - WAF Logs - full logging of all requests analyzed by the service
> 
> Logs can be analyzed using AWS Athena if they're stored in S3:
> - You should encrypt logs in S3, control access using IAM & Bucket Policies, MFA
> - Move Logs to Glacier for cost savings

</details>


### Additional Security Services

<details>
<summary>
anomaly detection, private data detection, best practices assessments
</summary>

- <cloud>GuardDuty</cloud> - machine learning to detect anomalies from event logs.
- <cloud>Amazon Macie</cloud> - detect personal data and private data that shouldn't be stored.
- <cloud>Trusted Advisor</cloud> - account assessments
#### Amazon GuardDuty

> Intelligent Threat discovery to protect your AWS Account. Uses Machine Learning algorithms, anomaly detection, 3rd party data.\
> One click to enable (30 days trial), no need to install software.\
> Input data includes:
> - <cloud>CloudTrail</cloud> Events Logs - unusual API calls, unauthorized deployments
>   - Management Events - create VPC subnet, create trail, etc...
>   - <cloud>S3</cloud> Data Events - get object, list objects, delete object, etc...
> - <cloud>VPC</cloud> Flow Logs - unusual internal traffic, unusual IP address
> - DNS Logs - compromised EC2 instances sending encoded data within DNS queries
> - Optional Features 
>   - EKS Audit Logs
>   - RDS & Aurora
>   - EBS
>   - Lambda
>   - S3 Data Events
> 
> Can setup EventBridge rules to be notified in case of findings. EventBridge rules can target AWS Lambda or SNS\
> Can protect against CryptoCurrency attacks (has a dedicated "finding" for it)


#### Amazon Macie

> Amazon Macie is a fully managed data security and data privacy service  that uses machine learning and pattern matching to discover and protect your sensitive data in AWS.\
> Macie helps identify and alert you to sensitive data, such as personally identifiable information (PII)

#### Trusted Advisor

> AWS account assessment - Analyze your AWS accounts and provides recommendation on 6 categories:
> - Cost optimization
> - Performance
> - Security
> - Fault tolerance
> - Service limits
> - Operational Excellence
> 
> Business & Enterprise Support plan
> 
> - Full Set of Checks
> - Programmatic Access using AWS Support API

</details>

### Key Management Service and Encryption

<details>
<summary>
Encryption and Keys.
</summary>

<cloud>KMS</cloud> - Key Management Service.

Encryption can happen inFlight - during transit, or atRest - server side, client-side - encrypted by the client (and not decrypted by the server).

> Encryption in flight (TLS / SSL)
> - Data is encrypted before sending and decrypted after receiving
> - TLS certificates help with encryption (HTTPS)
> - Encryption in flight ensures no MITM (man in the middle attack) can happen.
> 
> Server-side encryption at rest
> - Data is encrypted after being received by the server
> - Data is decrypted before being sent
> - It is stored in an encrypted form thanks to a key (usually a data key)
> - The encryption / decryption keys must be managed somewhere, and the server must have access to it
> 
> Client-side encryption
> - Data is encrypted by the client and never decrypted by the server
> - Data will be decrypted by a receiving client
> - The server should not be able to decrypt the data
> - Could leverage Envelope Encryption


KMS is a managed Key service

> - AWS manages encryption keys for us
> - Fully integrated with IAM for authorization
> - Easy way to control access to your data
> - Able to audit KMS Key usage using <cloud>CloudTrail</cloud>
> - Seamlessly integrated into most AWS services (<cloud>EBS</cloud>, <cloud>S3</cloud>, <cloud>RDS</cloud>, <cloud>SSM</cloud> and others)
> - Never ever store your secrets in plaintext, especially in your code!
> - KMS Key Encryption also available through API calls (SDK, CLI)
> - Encrypted secrets can be stored in the code / environment variables

Support symmetric (AES-256) and asymmetric (RSA, ECC key pairs) keys. symmetric keys use the same key to encrypt and decrypt the data, this is what intgrated services use. we never get access to the key itself. Asymmetric encryption uses two keys (private and public), we can download and share the public key, but not the private one. this is used for services outside of AWS which can't call the <cloud>KMS</cloud> api.

Keys can be owned by AWS - default keys for services, or be managed by the user - keys which are created or imported by the user

> - AWS Owned Keys (free): 
>   - SSE-S3
>   - SSE-SQS 
>   - SSE-DDB (default key)
> 
> - AWS Managed Key: free (aws/service-name, example: aws/rds or aws/ebs)
>
> Customer managed
> - Customer managed keys created in KMS: $1 / month
> - Customer managed keys imported: $1 / month
>   - + pay for API call to KMS ($0.03 / 10000 calls)

keys are scoped per region, so if we want to copy an encrypted <cloud>EBS</cloud> volume from one region to another, there are things we must do:
- make a snapshot with the same key
- copy the snapshot and reEncrypt it with a different key
- restore a volume from the copied snapshot

KMS has key policies, which control access to the keys, just like <cloud>S3</cloud> bucket policies,unlike them, key policies are mandatory and must be used.

> Default KMS Key Policy:
> 
> - Created if you don't provide a specific KMS Key Policy
> - Complete access to the key to the root user = entire AWS account
> 
> Custom KMS Key Policy:
> 
> - Define users, roles that can access the KMS key
> - Define who can administer the key
> - Useful for cross-account access of your KMS key

if we want to copy a snapshot between accounts, the snapshot is first enctypted using a KMS key. then we attach a key policy to authorize cross-account access,
we can share the snapshot with the target account, which can then create a copy with a different key or create a volume from it.

#### KMS Demo and CLI
in the <cloud>KMS</cloud> service, under "AWS Managed Keys" section, we can see the keys created for each service that we've used in the account. the key has an alias and a key policy that defines it can only be accessed from that service.\
Under the "Customer Managed Keys", we can <kbd>Create Key</kbd>, either symmetric or asymmetric, and we can choose if the key will be used for encryption and decryption, or for hash based authentication.\
Keys can be regional or multi-regional (which means the can be replicated across regions), and can be created in <cloud>KMS</cloud> or imported from a file. we can give the key an alias name, and then define who can manage the key and who can use it (including other accounts). this is what creates the key polices.\
The key can have automatic rotation, between 90 days and 2560 days (around seven years). we initate on-demand rotation if we want. keys can be disabled or be scheduled for deletion.

("customer key stores" are used for HSM - later)

We can use the CLI to encrypt and decrypt data. when we run the command we get a base64 coding on top of the encrypted binary file.

```sh
# 1) encryption
aws kms encrypt --key-id alias/tutorial --plaintext fileb://ExampleSecretFile.txt --output text --query CiphertextBlob  --region eu-west-2 > ExampleSecretFileEncrypted.base64

# base64 decode for Linux or Mac OS 
cat ExampleSecretFileEncrypted.base64 | base64 --decode > ExampleSecretFileEncrypted

# base64 decode for Windows
certutil -decode .\ExampleSecretFileEncrypted.base64 .\ExampleSecretFileEncrypted


# 2) decryption

aws kms decrypt --ciphertext-blob fileb://ExampleSecretFileEncrypted   --output text --query Plaintext > ExampleFileDecrypted.base64  --region eu-west-2

# base64 decode for Linux or Mac OS 
cat ExampleFileDecrypted.base64 | base64 --decode > ExampleFileDecrypted.txt


# base64 decode for Windows
certutil -decode .\ExampleFileDecrypted.base64 .\ExampleFileDecrypted.txt
```

#### KMS Key Rotation

> Automatic Key rotation:
> - AWS-managed KMS Key: automatic every 1 year
> - Customer-managed KMS Key: (must be enabled) automatic & on-demand
> - Imported KMS Key: only manual rotation possible using alias

when we rotate a key, the key id remains the same, but the backing key changes (the old one remains for decryption).

manual key rotation requires creating new keys and new key ids, so it's best to use aliases. the application only knows the alias, not the underlying key id.

#### KMS Operations

Encrypting and sharing Volumes and snapshots

> Changing The KMS Key For An Encrypted EBS Volume
> - You can't change the encryption keys used by an EBS volume
> - Create an EBS snapshot and create a new EBS volume and specify the new KMS key
> 
> Sharing KMS Encrypted RDS DB Snapshots
> - You can share RDS DB snapshots encrypted with KMS CMK with other accounts, but must first share the KMS CMK with the target account using Key Policy

Deleting KMS keys

> KMS Key Deletion Considerations
> - Schedule CMK for deletion with a waiting period of 7 to 30 days
>   - CMK's status is "Pending deletion" during the waiting period.
> - During the CMK's deletion waiting period:
>   - The CMK can't be used for cryptographic operations (e.g., can't decrypt KMSencrypted objects in S3 - SSE-KMS)
>   - The key is not rotated even if planned
> - You can cancel the key deletion during the waiting period
> - Consider disabling your key instead of deleting it if you're not sure!

we can create a <cloud>CloudWatch</cloud> alarm to monitor <cloud>CloudTrail</cloud> events regarding the keys in the "Pending Deletion" state to get notified if some user or service tries using the keys.

#### CloudHSM

HSM - Hardware Security Module
- KMS - encryption software
- HSM - encryption hardware

we manage the encryption keys, not AWS

> - HSM device is tamper resistant, FIPS 140-2 Level 3 compliance
> - Supports both symmetric and asymmetric encryption (SSL/TLS keys)
> - No free tier available
> - Must use the CloudHSM Client Software
> - <cloud>Redshift</cloud> supports CloudHSM for database encryption and key management
> - Good option to use with SSE-C encryption
> 
> - CloudHSM clusters are spread across Multi AZ (HA)
>   - Great for availability and durability

HSM has integration with KMS using the Custom Key Store. this gives us the benefit of getting <cloud>CloudTrail</cloud> audit logs when our HSM keys are used through the <cloud>KMS</cloud>

| Feature                    | AWS KMS                                                                                  | AWS CloudHSM                                                           |
| -------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| Tenancy                    | Multi-Tenant                                                                             | Single-Tenant                                                          |
| Standard                   | FIPS 140-2 Level 3                                                                       | FIPS 140-2 Level 3                                                     |
| Master Keys                | AWS Owned CMK, AWS Managed CMK, Customer Managed CMK                                     | Customer Managed CMK                                                   |
| Key Types                  | Symmetric, Asymmetric, Digital Signing                                                   | Symmetric, Asymmetric, Digital Signing & Hashing                       |
| Key Accessibility          | Accessible in multiple AWS regions (can'taccess keys outside the region it's created in) | Deployed and managed in a VPC, Can be shared across VPCs (VPC Peering) |
| Cryptographic Acceleration | None                                                                                     | SSL/TLS Acceleration, Oracle TDE Acceleration                          |
| Access & Authentication    | AWS IAM                                                                                  | You create users and manage their permissions                          |
| High Availability          | AWS Managed Service                                                                      | Add multiple HSMs over different AZs                                   |
| Audit Capability           | CloudTrail, CloudWatch,                                                                  | CloudTrail, CloudWatch, MFA support                                    |
| Free Tier                  | Yes                                                                                      | No                                                                     |

</details>

### AWS Artifact Overview

<details>
<summary>
AWS portal for Reports and Compliance agreements (AWS-wide and account specific).
</summary>

> AWS Artifact (not really a service)
> - Portal that provides customers with on-demand access to AWS compliance documentation and AWS agreements
> - Artifact Reports - Allows you to download AWS security and compliance documents from third-party auditors, like AWS ISO certifications, Payment Card Industry (PCI), and System and Organization Control (SOC) reports
> - Artifact Agreements - Allows you to review, accept, and track the status of AWS agreements such as the Business Associate Addendum (BAA) or the Health Insurance Portability and Accountability Act (HIPAA) for an individual account or in your organization
> - Can be used to support internal audit or compliance

</details>

### AWS Certificate Manager

<details>
<summary>
TLS Certificates for in-flight encryption
</summary>

ACM - <cloud>AWS Certificate Manager</cloud>

> Easily provision, manage, and deploy TLS Certificates
> - Provide in-flight encryption for websites (HTTPS)
> - Supports both public and private TLS certificates
> - Free of charge for public TLS certificates
> - Automatic TLS certificate renewal
> - Integrations with (load TLS certificates on):
>   - Elastic Load Balancers (CLB, ALB, NLB)
>   - CloudFront Distributions
>   - APIs on API Gateway
> 
> - Cannot use ACM with EC2 (can't be extracted)

expose an <cloud>ELB</cloud> as a secure endpoint with HTTPS protocol.

we request a Public certificate for a list of domain names, either a fully qualified name or a wild card. then we set validation mode: DNS or Email. if we use DNS validation we need to use a `CNAME` record.\
This process can take a few hours. certificates are renewed automatically 60 days before expiration.\

We can also import a public certificate, in this case there is no automatic renewal.

we can get notified about certificates expiration through the <cloud>EventBridge</cloud> service, or through <cloud>AWS Config</cloud> with the "acm-certificate-expiration-check" rule, and then we can get a non-complint resource configuration.

in the <cloud>ELB</cloud>, we can set a redirect rule from HTTP to HTTPS, and use the TLS certificates.

for <cloud>API Gateway</cloud>, <cloud>ACM</cloud> integrates with two of the three EndPoints types:

- edge-optimized (default, global clients) - possible
- regional (only on one regional) - possible
- private (from <cloud>VPC</cloud> using ENI) - meaningless

we first create a `Custom Domain Name` resource in the gateway.\
For edge-optimized, the TLS certificate must be created in us-east-1 region, which is where the <cloud>CloudFront</cloud> service operates from.\
For regional gateways, the certificate in <cloud>ACM</cloud> must be imported from the same region.\
for both cases we set the `CNAME` or `A-Alias` in <cloud>Route53</cloud>.

HandsOn:
in <cloud>Certificate Manager</cloud> service, we can <kbd>Request a Certificate</kbd>, either public or private, we select a public, and provide a fully qualified domain. the validation method will be DNS, and using default key. the request is now in pending state, and we are prompted to create a `CNAME` record in <cloud>Route53</cloud>.\
After a while, the certificate is up and validated, and we can create an application to use it.\
in the <cloud>Elastic BeanStalk</cloud> service, we <kbd>Create application</kbd> (web server), using the sample application and the custom configuration for availability. we use the default settings, and use a dedicated application load balancer. we then <kbd>Add Listener</kbd> on port 443 with the HTTPS protocol, and here we need to choose the SSL certificate to use. we also need to choose a SSL policy (how strong the HTTPS security is). we continue with the defaults and launch the environment.\
we navigate to <cloud>Route53</cloud>, and <kbd>Create A Record</kbd> of `CNAME` type, enter the load balancer address, and point our record to it.

</details>

### Secrets Manager

<details>
<summary>
Manage Secrets with rotation.
</summary>

> Newer service, meant for storing secrets. Capability to force rotation of secrets every X days.
> - Automate generation of secrets on rotation (uses <cloud>Lambda</cloud>)
> - Mostly meant for RDS integration
>   - MySQL
>   - PostgreSQL
>   - <cloud>Aurora</cloud>
> - Secrets are encrypted using <cloud>KMS</cloud>

There are also multi-region secrets, which are replicas of the data across other regions, all kept in-sync with the primary secret and automatically change when it is changed. Replica Secrets can be later promoted to StandAlone secrets.

Hands On:\
in the service console, <kbd>Store New Secret</kbd>, we are prompted with which kid of secret we want to use
- <cloud>RDS</cloud>
- <cloud>RedShift</cloud>
- <cloud>DocumentDB</cloud>
- Other database secrets
- other types of secrets (such as api key)

for databases, we get a user-password, and for other types, we can use a list of key-value pairs or plain text (a json). we select the encryption key (default or <cloud>KMS</cloud>) and we can configure the rotation period, and after that time a <cloud>Lambda</cloud> will be invoked and do the required actions to rotate the secret.

for RDS databases, the integration is a bit better, and the secret manager will handle setting the secret on the database directly.

#### Secrets Manager - Monitoring & Troubleshooting

> <cloud>CloudTrail</cloud> captures API calls to the Secrets Manager API. It captures other related events that might have a security or compliance impact on your AWS account or might help you troubleshoot operational problems.\
> CloudTrail records these events as non-API service events:
> 
> - `RotationStarted` event
> - `RotationSucceeded` event
> - `RotationFailed` event
> - `RotationAbandoned` event - a manual change to a secret instead of automated rotation
> - `StartSecretVersionDelete` event
> - `CancelSecretVersionDelete` event
> - `EndSecretVersionDelete` event
> 
> Combine with CloudWatch Logs and CloudWatch alarms for automations

we can also monitor the rotation through the <cloud>Lambda</cloud> Logs of the failed rotation.

#### SSM Parameter Store vs Secrets Manager

Secret Manager has rotation and management for secrets, unlike <cloud>SSM Parameter Store</cloud> with secure values.\
- Secret Manager is more expensive.
- Integration with Lambda for rotation (we have code templates to use).

Parameter Store uses a simpler API, doesn't require <cloud>KMS</cloud> encryption (it is optional). we can employ scheduled <cloud>EventBridge</cloud> rules to invoke a lambda which replaces the secret and stores the updated value.

</details>

</details>

## Identity

<details>
<summary>
Identity Security Tools, temporary identity access.
</summary>

> **IAM Credentials Report** (account-level)
> - a report that lists all your account's users and the status of their variouscredentials
> 
> **IAM Access Advisor (user-level)**
> - Access advisor shows the service permissions granted to a user and when those services were last accessed.
> - You can use this information to revise your policies.


in the <cloud>IAM</cloud> service page, we can choose "credentials report" and <kbd>download credentials report</kbd> to create a csv file, the file contains the users, password age, MFA, access keys and certificates.\
in the "Access advisor", we can see which users were accessed, and when. then we can limit the policy permissions we apply to the users.

### IAM Access Analyzer

<details>
<summary>
Find and Remediate public access to resources.
</summary>

> Find out which resources are shared externally
> - <cloud>S3</cloud> Buckets
> - <cloud>IAM</cloud> Roles
> - <cloud>KMS</cloud> Keys
> - <cloud>Lambda</cloud> Functions and Layers
> - <cloud>SQS</cloud> queues
> - Secrets Manager Secrets
> 
> Define Zone of Trust = AWS Account or AWS Organization. Access outside zone of trusts => findings

in the "Access Analyzer" section of <cloud>IAM</cloud>, we <kbd>Create Analyzer</kbd>, we need new role to execute the report, and we can add optional tags. the resulting report shows the users which can be accessed from outside the account, what kind of access is allowed (read, write), and how is the resource shared (like bucket access policy). for each finding we can either approve that this is the intended behavior, or take steps to mitigate it and limit the public access. then we can <kbd>rescan</kbd> the finding to make sure it's no longer the issue.\
We can create rules to auto achieve findings.
</details>

### Identity Federation with SAML & Cognito

<details>
<summary>
Accessing AWS resources through non-AWS identities.
</summary>

We use an identity provider (either corporate or public) to get temporary AWS access.

> Federation lets users outside of AWS to assume temporary role for accessing AWS resources. These users assume identity provided access role. Federation assumes a form of 3rd party authentication:
> 
> - LDAP
> - Microsoft Active Directory (~= SAML)
> - Single Sign On
> - Open ID
> - <cloud>Cognito</cloud>
>
> Using federation, you don't need to create <cloud>IAM</cloud> users (user management is outside of AWS).

SAML (Security Assertion Markup Language) for enterprises, this is what Active Directory and ADFS are. any SAML 2.0 compliant identity broker can 
integrated with AWS. users are given temporary credentials to assume a role, without creating an <cloud>IAM User</cloud>. we can use this to get programattic access and web console access.

programatic access flow:
1. User identifies itself against the identity provider.
2.  Idp (identity provider) checks the user credentials against the LDAP identity store.
3. If authenticated, IdP returns a SAML assertion to the user.
4. The user calls `AssumeRoleWithSAML` on the <cloud>STS</cloud> service with the SAML.
5. The service returns temporary security credentials.
6. The user uses those credentials to access AWS resources

web console (aws portal) access flow is similar, but instead of getting temporary roles credentials, <cloud>STS</cloud> returns an aws sign-in endpoint and the user is re-directed to it.

if we don't have a SAML 2.0 compatible Identity Provider (IdP), we can use the Custom Identity Broker Application instead. in this case we need to create the flow ourselves, and we write the identity broker to communicate with <cloud>STS</cloud> and return the appropriate credentials. there is more manual work required, since we don't have the SAML.
</details>

### STS & Cross Account Access

<details>
<summary>
The STS service - assume other identity for a limited time.
</summary>

<cloud>STS</cloud> - Security Token Service

> Allows to grant limited and temporary access to AWS resources. Token is valid for up to one hour (must be refreshed).
> - `AssumeRole` - Within your own account: for enhanced security, Cross Account Access: assume role in target account to perform actions there.
> - `AssumeRoleWithSAML` - return credentials for users logged with SAML
> - `AssumeRoleWithWebIdentity` - return credentials for users logged with an IdP (Facebook Login, Google Login, OIDC compatible...). AWS recommends against using this, and using Cognito instead
> - `GetSessionToken` - for MFA, from a user or AWS account root user
>
> Using STS to Assume a Role
> - Define an IAM Role within your account or cross-account
> - Define which principals can access this IAM Role
> - Use AWS STS (Security Token Service) to retrieve credentials and impersonate the IAM Role you have access to (`AssumeRole` API)
> - Temporary credentials can be valid between 15 minutes to 1 hour
</details>

### Cognito

<details>
<summary>
AWS public Identity service, either for AWS resources with Identity Pools or custom resources with User Pools.
</summary>

<cloud>AWS Cognito</cloud> is an option for non-corporate users. when we need to give the public some access to resources, but the identity provider is public, rather than corporate:
- Google
- Facebook
- CUP (cognito user pools)
- Twitter
- OpenID
- etc...

> - Goal:
>   - Provide direct access to AWS Resources from the Client Side
> - How:
>   - Log in to federated identity provider - or remain anonymous
>   - Get temporary AWS credentials back from the Federated Identity Pool
>   - These credentials come with a pre-defined IAM policy stating their permissions
> - Example:
>   - provide (temporary) access to write to S3 bucket using Facebook Login

There is also an alternative service, called <cloud>Web Identity Federation</cloud>, but AWS recommends using <cloud>Cognito</cloud> over it.

#### Cognito User Pools

CUP - Cognito User Pools - simple, serverless, managed authentication option for websites and applications.

> - Create a **serverless database** of user for your web & mobile apps
> - Simple login: Username (or email) / password combination
> - Password reset
> - Email & Phone Number Verification
> - Multi-factor authentication (MFA)
> - Federated Identities: users from Facebook, Google, SAML...
> - Feature: block users if their credentials are compromised elsewhere
> - Login sends back a JSON Web Token (JWT)
> 
> CUP integrates with <cloud>API Gateway</cloud> and Application Load Balancer (<cloud>ELB</cloud>).

#### Cognito Identity Pools

Cognito Identity Pools (Federated Identities). temporary access to AWS resources to a large number of guest users (not corporate), and even for guest access.

> - Get identities for "users" so they obtain temporary AWS credentials
> - Your identity pool (e.g identity source) can include:
>   - Public Providers (Login with Amazon, Facebook, Google, Apple)
>   - Users in an Amazon Cognito user pool
>   - OpenID Connect Providers & SAML Identity Providers
>   - Developer Authenticated Identities (custom login server)
>   - Cognito Identity Pools allow for unauthenticated (guest) access
> 
> - Users can then access AWS services directly or through API Gateway
>   - The IAM policies applied to the credentials are defined in Cognito
>   - They can be customized based on the user_id for fine grained control

Users log-in into an identity provider, get a token from them and trade that with the Cognito Identity Pool for temporary credentials (through the <cloud>STS</cloud> service). and now the users can access AWS resources.

Policies are applied bas on the IAM role, different roles for authenticated and guest users. we can give policies based on the user_id, and then limit the policies access to resources based on policy variables.

for example, the guest policy gives access to only a single resource.
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicReadGuestUser",
      "Effect": "Allow",
      "Action": [
        "s3:GetObject"
      ],
      "Resource":[
        "arn:aws:s3:::examplebucket/assets/my_picture/jpg*"
      ]
    }
  ]
}
```

while the authenticated policy will give more permessions, but limits the resource based on the users' cognito identity. we limit access to S3 based on prefix, and limit access to dynamoDB based on the keys. each user can only access their own resources, but they still use the same policy document.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListBucket"
      ],
      "Resource":[
        "arn:aws:s3:::examplebucket"
      ],
      "Condition" {
        "StringLike": {"s3:prefix":["${cognito-identity.amazon.aws:sub}/*"]}
      }
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:PutObject"
      ],
      "Resource":[
        "arn:aws:s3:::examplebucket/${cognito-identity.amazon.aws:sub}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "dynamoDb:GetItem",
        "dynamoDb:BatchGetItem",
        "dynamoDb:Query",
        "dynamoDb:PutItem",
        "dynamoDb:UpdateItem",
        "dynamoDb:BatchWriteItem",
        "dynamoDb:DeleteItem"
      ],
      "Resource":[
        "arn:aws:dynamodb:us-wes-2:<accountId>:table/MyTable"
      ],
      "Condition" {
        "ForAllValues:StringEquals":{
          "dynamodb:LeadingKeys": ["${cognito-identity.amazon.aws:sub}"]
        }
      }
    }
  ]
}
```

#### Cognito User Pools vs Cognito Identity Pools

> Cognito User Pools (for authentication = identity verification)
> - Database of users for your web and mobile application
> - Allows to federate logins through Public Social, OIDC, SAML…
> - Can customize the hosted UI for authentication (including the logo)
> - Has triggers with AWS Lambda during the authentication flow
> - Adapt the sign-in experience to different risk levels (MFA, adaptive authentication, etc...)
> 
> Cognito Identity Pools (for authorization = access control)
> - Obtain AWS credentials for your users
> - Users can login through Public Social, OIDC, SAML & Cognito User Pools
> - Users can be unauthenticated (guests)
> - Users are mapped to IAM roles & policies, can leverage policy variables
> 
> CUP + CIP = authentication + authorization
</details>


</details>

## Networking - Route53

<details>
<summary>
DNS Service, Route Types, Routing Policies, Hosting websites, Resolving DNS requests
</summary>

> - A highly available, scalable, fully managed and Authoritative DNS
> - Authoritative = the customer (you) can update the DNS records
> - Route53 is also a Domain Registrar
> - Ability to check the health of your resources
> - The only AWS service which provides 100% availability SLA
> - Why Route53? 53 is a reference to the traditional DNS port

<cloud>Route53</cloud> use record to control traffic for a domain, Each record contains:

> - Domain/subdomain Name - e.g., example.com
> - Record Type - e.g., `A` or `AAAA`
> - Value - e.g., 12.34.56.78
> - Routing Policy - how Route53 responds to queries
> - TTL - amount of time the record cached at DNS Resolvers

Route53 supports the following DNS record types ([wikipedia](https://en.wikipedia.org/wiki/List_of_DNS_record_types)):
> - must Know Types
>   - A - address record (IPv4)
>   - AAAA - IPv6 record
>   - CNAME - Canonical name record
>   - NS - Name server record
> - advanced types
>   - CAA - Certification Authority Authorization
>   - DS - Delegation signer
>   - MX - Mail exchange record
>   - NAPTR - Naming Authority Pointer
>   - PTR - PTR Resource Record - (similar to CNAME, but without further processing)
>   - SOA - Start of [a zone of] authority record
>   - SPF - Sender Policy framework
>   - SRV - Service Locator
>   - TXT - Text Record

the important Record types are:

> - `A` - maps a hostname to IPv4
> - `AAAA` - maps a hostname to IPv6
> - `CNAME` - maps a hostname to another hostname
>   - The target is a domain name which must have an `A` or `AAAA` record
>   - Can't create a `CNAME` record for the top node of a DNS namespace (Zone Apex)
>   - Example: you can't create for example.com, but you can create for www.example.com
> - `NS` - Name Servers for the Hosted Zone
>   - Control how traffic is routed for a domain

Hosted Zones are containers for records, and they define how to route traffic to the domain and the sub domain. there are public Hosted Zones for public domain names, and private hosted zone for records and routing traffic within <cloud>VPC</cloud>s.

- "application1.myPublicDomain.com" - public domain name - anyone from the internet
- "application1.company.internal" - private domain name - only from inside the <cloud>VPC</cloud>

<cloud>Route53</cloud> is not free, there is a cost for each hosted Zone, and a cost to register a domain name.

### What is a DNS?

<details>
<summary>
DNS Overview
</summary>

DNS - Domain Name System. it takes a human friendly hostname and translates them into machine ip addresses. www.google.com becomes "172.17.18.36". it uses hierarchical naming structure.

- "www.example.com"
- "api.example.com"

both domains are inside "example.com" Second Level domain, which itself is inside the "com" Top-level domain.


DNS Terminologies
> - Domain Registrar: Amazon <cloud>Route53</cloud>, GoDaddy, etc..
> - DNS Records: `A`, `AAAA`, `CNAME`, `NS`, etc..
> - Zone File: contains DNS records.
> - Name Server: resolves DNS queries (Authoritative or Non-Authoritative)
> - Top Level Domain (TLD): ".com", ".us", ".in", ".gov", ".org", etc..
> - Second Level Domain (SLD): "amazon.com", "google.com", etc..

the URL is a combination of the protocol (http, https), and the fully qualified domain name, which is the sub domain(more than one), second-level domain and the top-level domain and the root domain (a hidden dot).

to use the DNS, we register our public IP address in one of the domain Registrars. the web browser uses a series of DNS servers to match the FQDN to an ip address. the first DNS server is usually the local one, either managed by the company or the ISP (internet service provider), if they DNS server doesn't know the address, it can ask the root DNS server (which is managed by ICANN - The Internet Corporation for Assigned Names and Numbers) which either knows the IP address, or can direct us to the next DNS server, which manages the Top-level domain, and it can direct to the Domain Registar that we registered our IP with, at the next step, the Registar actually knows the ip address, so it can return it, and it is then cached in the local DNS server for future requests.
</details>


### Route53 - Demo Setup

<details>
<summary>
Using Route53 - setting up.
</summary>

In the <cloud>Route53</cloud> dashboard, we can see how many hostedZones wer have, how many registered domains, policy records and health checks.

#### Route53 - Registering a Domain
we can also choose the "Registered Domain" section and <kbd>Register Domains</kbd> if we want to, we could also <kbd>Transfer in</kbd> a domain from a different registar to manage it in <cloud>Route53</cloud>. if we create one, we need to find a domain that nobody else uses, we pay upfront for the entire year, and we can set ao renewal for it. we provide our contact information and confirm. once it's created, we will have new hosted ho with two records - `NS` record pointing traffic to route53 and the `SOA` record.

#### Route53 - Creating our first records

we can now <kbd>Create Record</kbd> inside our domain, this record is tha address, so the name must be compliant. we also choose:
- record type
- TTL
- Routing Policy
- Value (one or more)

since we don't have anything in that IP address, our browser doesn't show us anything. instead, we can use the <cloud>CloudShell</cloud> terminal to get a better understaing of what's going on there.

```sh
sudo yum install -y bind-utils # install nslookup, dig
nslookup test.stephaneTheTeacher.com
dig test.stephaneTheTeacher.com
```

#### Route53 - EC2 Setup

we will create three <cloud>EC2</cloud> machines to act as our webservers, each in a different Region, each in a security Group that allows public internet access, and each having the following userdate script, which will install the webserver, and create the index page saying where it is located.

```sh
#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
# updated script to make it work with Amazon Linux 2023
CHECK_IMDSV1_ENABLED=$(curl -s -o /dev/null -w "%{http_code}" http://169.254.169.254/latest/meta-data/)
if [[ "$CHECK_IMDSV1_ENABLED" -eq 200 ]]
then
    EC2_AVAIL_ZONE="$(curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone)"
else
    EC2_AVAIL_ZONE="$(TOKEN=`curl -s -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 21600"` && curl -s -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/placement/availability-zone)"
fi
echo "<h1>Hello world from $(hostname -f) in AZ $EC2_AVAIL_ZONE </h1>" > /var/www/html/index.html
```

we also launch a <cloud>ELB</cloud> of type application load balancer in one of the regions, make it internet facing and set the target group to include the <cloud>EC2</cloud> machine in it. we copy the public ips from the machines and the public dns name from the load balancer.
</details>

### Route53 - TTL

<details>
<summary>
Time To Live Cache
</summary>

each record has a TTL (time to live), which is the duration of time to cache the result. if the client requests the same address during that period, it will use the cached value and won't request it from <cloud>Route53</cloud>.\
High TTL values cost us less, since the amount of requests is reduced, but there is a longer period of time until all the records are updated. Low TTL values mean that we can change the address faster, but it will cost us more.

TTL values are mandory for each DNS record, except for `ALIAS` (which is an AWS type).

back in our demo, we can create a new RECORD, point it to one of our ips, and set a TTL with two minutes value. we can run the `nslookup` and `dig` commands on our record to see the TTL going down. we can next update the record to point to another machine, and until the cache expires, we will still go to the earlier machine. only after the cache period changes, we will reach the other machine.

</details>

### CNAME vs Alias

<details>
<summary>
Differences between CNAME and ALIAS records
</summary>

> AWS Resources (<cloud>Elastic Load Balancer</cloud>, <cloud>CloudFront</cloud>...) expose an AWS hostname. for example  lb1-1234.us-east-2.elb.amazonaws.com and you want myApp.myDomain.com.
> 
> `CNAME`:
> - Points a hostname to any other hostname. (app.myDomain.com => blaBla.anything.com)
> - ONLY FOR NON ROOT DOMAIN (something.myDomain.com)
> 
> `Alias`:
> - Points a hostname to an AWS Resource (app.myDomain.com => blaBla.amazonaws.com)
> - Works for ROOT DOMAIN and NON ROOT DOMAIN (myDomain.com and something.myDomain.com)
> - Free of charge
> - Native health check

ALIAS records can only map to AWS resources, this is an <cloud>Route53</cloud> extensions. the alias record must point to either `A` or `AAAA` records (not `CNAME`), and they don't have TTL values.
- <cloud>Elastic Load Balancers</cloud>
- <cloud>CloudFront</cloud> Distributions
- <cloud>API Gateway</cloud>
- <cloud>Elastic Beanstalk</cloud> environments
- <cloud>S3</cloud> Websites
- <cloud>VPC</cloud> Interface Endpoints
- <cloud>Global Accelerator accelerator</cloud>
- <cloud>Route53</cloud> record in the same hosted zone
- 
- CAN NOT POINT TO <cloud>EC2</cloud> DNS NAME

we can <kbd>Create Record</kbd> of `CNAME` type and give it the value of the load balancer DNS name. we can also create an `A` record, make it an `ALIAS` and in the dropdown, route the traffic to the load balancer. this will give us automatic health checks, and this is free to access.\
we can also set apex zone (empty record name) to point to the resource with an `ALIAS`, but not with a `CNAME`.

</details>


### Route53 Health Checks

<details>
<summary>
Determining which Records are healthy and can be used.
</summary>

> HTTP Health Checks are only for public resources. Health Check => Automated DNS Failover:
> 
> - Health checks that monitor an endpoint (application, server, other AWS resource)
> - Health checks that monitor other health checks (Calculated Health Checks)
> - Health checks that monitor <cloud>CloudWatch</cloud> Alarms (full control !!) - e.g., throttles of <cloud>DynamoDB</cloud>, alarms on <cloud>RDS</cloud>, custom metrics,etc... (helpful for private resources)
> 
> Health Checks are integrated with <cloud>CloudWatch</cloud>
metrics.

for example, if we have health check on an <cloud>ELB</cloud> endpoint, there will be about 15 health checkers from all around the world (or what we configure), we can control the frequency, interval, protocol of the check. a certain percentage of those health checks need to pass to consider the endpoint as healthy.\
Health checks pass if the status code is in the 2xx and 3xx ranges, and we can also check the first 5120 bytes of the response to determine pass/fail status.\
The router/firewall must allow access from the ip-ranges of the <cloud>Route53</cloud> health checkers.

Calculate Health Checks combine the results of multiple health checks into one. we can aggregate them with
- OR
- AND
- NOT

and monitor up to 256 child health checks, and we choose what percentage needs to pass to make the aggregated health check pass.

we can't check private resources with the default health checkers, since they live outside the VPC. if we want to monitor something private (like a private endpoint), we can create a <cloud>CloudWatch</cloud> metric, assign an alarm to it, and the have the health check monitor the alarm status.

in the <cloud>Route53</cloud> dashboard, we can <kbd>Create Health Check</kbd>, one for each of our EC2 instances. We will monitor an endpoint, we pass the ip address, the port and the path, and under the advance configuration section, we can control the request intervals, decide which health checkers to use, and even create an alarm. now we can modify the security group in one of the regions and remove inbound access on port 80, so the health checkers won't be able to communicate with it, and the health check will fail.\
We can configure a calculated health check to monitor all three health checks together, or create a health check on a cloudWatch alarm. 

now that we have health checks, our DNS records can use them.
</details>

### Routing Policies
<details>
<summary>
How Route53 chooses to return IPs
</summary>

> Define how Route53 responds to DNS queries. Don't get confused by the word "Routing", It's not the same as Load balancer routing which routes the traffic. DNS does not route any traffic, it only responds to the DNS queries.\
> Route53 Supports the following Routing Policies:
> 
> - Simple
> - Weighted
> - Failover
> - Latency based
> - Geolocation
> - Multi-Value Answer
> - Geoproximity (using Route53 Traffic Flow feature)

#### Simple Routing Policy

returns multiple values, client chooses randmoly, no health checks.

> - Typically, route traffic to a single resource
> - Can specify multiple values in the same record
> - If multiple values are returned, a random one is chosen by the **client**
> - When Alias enabled, specify only one AWS resource
> - Can't be associated with Health Checks

#### Weighted Routing Policy

define multiple records with the same name and all having the `weighted` routing policy type, each with a record id. a record (which can have multiple values) is chosen proportionally to the weight assigned to it. supports health checks.

> - Control the % of the requests that go to each specific resource
> - Assign each record a relative weight
> - Weights don't need to sum up to 100
> - DNS records must have the same name and type
> - Can be associated with Health Checks
> - Use cases: load balancing between regions, testing new application versions
> - Assign a weight of 0 to a record to stop sending traffic to a resource
> - If all records have weight of 0, then all records will be returned equally

#### Latency Routing Policy

multiple records with the same name, different record id. we give each record a region, and <cloud>Route53</cloud> calculates which region has the best latency for the calling user, and will return that record. supports health checks.

> Redirect to the resource that has the least latency close to us
> - Super helpful when latency for users is a priority
> - Latency is based on traffic between users and AWS Regions
> - Germany users may be directed to the US (if that's the lowest latency)
> - Can be associated with Health Checks (has a failover capability)

we can test this by using a VPN (virtual private network).

#### Failover Routing Policy

Active-Passive routing,two records with the same name, the primary record has must have a health check associated with it. if the health check passes, the primary record is returned, if it's not health, then the secondary record is returned.\
We can play with the health check (like before, modify the security group) to trigger the failover from the primary to the secondary record.

#### Geolocation Routing Policy

return records based on given location (not latency), multiple records with the same name, for each record we can choose the location (default, continent, country, US-state), health checks will 

> - Different from Latency-based!
> - This routing is based on user location
> - Specify location by Continent, Country or by US State (if there's overlapping, most precise location selected)
> - Should create a “Default” record (in case there's no match on location)
> - Use cases: website localization, restrict content distribution, load balancing, etc...
> - Can be associated with Health Checks

#### Geoproximity Routing Policy

(Routing based on location, can be weighted towards specific resources - imagine it a gravitational pull, with higher bias meaning more mass and more pulling power)

> Route traffic to your resources based on the geographic location of users and resources.\
> Ability to shift more traffic to resources based on the defined bias.\
> To change the size of the geographic region, specify bias values:
> 
> - To expand (1 to 99) - more traffic to the resource
> - To shrink (-1 to -99) - less traffic to the resource
> 
> Resources can be:
> - AWS resources (specify AWS region)
> - Non-AWS resources (specify Latitude and Longitude)
> - You must use Route53 Traffic Flow to use this feature

#### Traffic Flow

The Traffic Flow tool simplifies the creation and maintenance of DNS records, the records are saved into a traffic flow policy which we can version and use for other hosted zone (different domain names).

in the <cloud>Route53</cloud> service, we can choose "Traffic Policy" and then <kbd>Create Traffic Policy</kbd>, and then we can create a chain of records and routeing policies. for Geoproximity rules, we define the regions, and define the bias. then we have a visual map showing the entire world and where the requests will be directed to. when we create a record through the traffic policy, it can't be edited.

#### IP-based Routing Policy

fine grained routing based on ip addresses. 

> - Routing is based on clients' IP addresses
> - You provide a list of CIDRs for your clients and the corresponding endpoints/locations (user-IP-to-endpoint mappings)
> - Use cases: Optimize performance, reduce network costs, etc...
> - Example: route end users from a particular ISP to a specific endpoint

#### Multi Value Routing Policy

return multiple records, then the client chooses one of them. unlike simple routing policies, this can support health checks. multiple records with the same name, each with a different health check. the user gets back only the healthy records.

> - Use when routing traffic to multiple resources
> - Route53 return multiple values/resources
> - Can be associated with Health Checks (return only values for healthy resources)
> - Up to 8 healthy records are returned for each Multi-Value query
> - Multi-Value is not a substitute for having an ELB
</details>

### 3rd Party Domains & Route53

<details>
<summary>
Using Route53 with a different DNS Registar
</summary>

> Domain Registar vs. DNS Service
> - You buy or register your domain name with a Domain Registrar typically by paying annual charges (e.g., GoDaddy, Amazon Registrar Inc., etc...)
> - The Domain Registrar usually provides you with a DNS service to manage your DNS records
> - But you can use another DNS service to manage your DNS records
> - Example: purchase the domain from GoDaddy and use Route53 to manage
your DNS records

if we purchase the domain somewhere else (registar), and we want to manage the records in <cloud>Route53</cloud>, we change the nameservers on the registar to point to the nameservers in the <cloud>Route53</cloud> public hosted zone.

</details>

### S3 Website with Route53

<details>
<summary>
Serving a S3 static website through Route53
</summary>

We can host a static website from a <cloud>S3</cloud> bucket. but if we want to give it a DNS address, the bucket should have the exact same name as the target record (such as "acme.example.com"), it needs to be a website, with the objects being publicly available.\
Then we create a DNS record of type `A` and `ALIAS`, and point it to the S3 website endpoint. this only works for HTTP traffic, not HTTPS (we will need to use <cloud>CloudFront</cloud> to get over that).

we first <kbd>Create Bucket</kbd> in <cloud>S3</cloud>, with the name exactly like what we want in <cloud>Route53</cloud>. in our bucket, we upload some files (index.html), we grant public access, confirm. and at the bucket settings, we enable it as a static website.\
Next, in <cloud>Route53</cloud>, we <kbd>Create Record</kbd> in the appropriate hostname, the name is like what we gave the bucket. the record type is `A` and we toggle it as `ALIAS`. we choose the S3, the region of the bucket, and if we did everything correctly, it will automatically detect the bucket endpoint, and our website will be reachable through a simple address.

</details>

### Route53 Resolvers & Hybrid DNS

<details>
<summary>
Hybrid DNS Resolution
</summary>

Use Hybrid DNS and Route53 Resolvers to resolve DNS queries across networks, such as communicating the on-premises networks or across VPC boundaries.

> By default, Route53 Resolver automatically answers DNS queries for:
> - Local domain names for <cloud>EC2</cloud> instances
> - Records in Private Hosted Zones
> - Records in public Name Servers
> 
> Hybrid DNS - resolving DNS queries between VPC (Route53 Resolver) and
your networks (other DNS Resolvers).\
> Networks can be:
> - VPC itself / Peered VPC
> - On-premises Network (connected through <cloud>Direct Connect</cloud> or <cloud>AWS VPN</cloud>)

Resolver endpoints: inbound and outbound.

> Inbound Endpoint
>   - DNS Resolvers on your network can forward DNS queries to Route53 Resolver
>   - Allows your DNS Resolvers to resolve domain names for AWS resources (e.g., EC2 instances) and records in Route53 Private Hosted Zones
> 
> - Outbound Endpoint
>   - Route53 Resolver conditionally forwards DNS queries to your DNS Resolvers
>   - Use Resolver Rules to forward DNS queries to your DNS Resolvers
>
> - Associated with one or more VPCs in the same AWS Region
> - Create in two AZs for high availability
> - Each Endpoint supports 10,000 queries per second per IP address

example, we have the following resources
- on-premises server "web.onPremise.private"
- on-premises DNS resolvers "onPremise.private"
- VPC with two private subnets
- EC2 instace inside a private subnet "app.aws.private" in the VPC
- Each subnet has an ENI (elastic network interface) - "10.0.0.10" and "10.0.1.10"
- The VPC has inbound and outbound Resolver endpoints inside, connected to the ENIs
- 
- Route53 Resolver
- Private Hosted Zone "aws.private"
- a VPN or <cloud>DX</cloud> connection between the the on-premises network and AWS

if the on-premise server wants to connect to the EC2 instance, it issues a DNS query to the local DNS resolver, which knows that for the domain name "aws.private", it forwards requests to the ENI addresses ("10.0.0.10" and "10.0.1.10") through the VPN connection. the Resolver inbound Endpoint directs the request to the Route53 resolver, which hosts the private hosted zone that can do the lookup itself.

for outbound endpoints, the EC2 instance communicates with route53 resolver, which communicates with an outbound resolver endpoint (which maps the "onPremise.private" domain to the target ip of the local DNS resolver). the local DNS resolvers can respond with the ip address of the on-premises server, and the instance can connect with them.


> Route 53 - Resolver Rules\
> Control which DNS queries are forwarded to DNS Resolvers on your network.
> 
> - Conditional Forwarding Rules (Forwarding Rules)
>   - Forward DNS queries for a specified domain and all its subdomains to target IP addresses
> - System Rules
>   - Selectively overriding the behavior defined in Forwarding Rules (e.g.,don't forward DNS queries for a subdomain acme.example.com)
> - Auto-defined System Rules
>   - Defines how DNS queries for selected domains are resolved (e.g., AWS internal domain names, Private Hosted Zones)
> - If multiple rules matched, Route 53 Resolver chooses the most specific match
> 
> - Resolver Rules can be shared across accounts using <cloud>AWS RAM</cloud>
>   - Manage them centrally in one account
>   - Send DNS queries from multiple VPC to the target IP defined in the rule

System Rules override Conditional Rowarding Rules, if multiple rules match, the more specific rule (more sub domains) is chosen.
</details>


</details>

## Networking - VPC

<details>
<summary>
Virtual Private Cloud and networking.
</summary>

VPC - Virtual Private Cloud.

we can have multiple VPCs in each region, the soft limit is 5, but it can be increased. each VPC can have up to 5 CIDR ranges attached to it. with the CIDR minimal size of "/28" (16 ip addresses), and maximum size of "/16" (65,536 ip addresses). only private IPv4 ranges are allowed:
- 10.0.0.0/8
- 172.16.0.0/12
- 192.168.0.0/16

the range should not overlap with other networks (other vpcs, corporate network), so we could connect between them.

(demo)

in the <cloud>VPC</cloud> service page, we click <kbd>Create VPC</kbd>, we give it a name, provide an IPv4 cidr block, choose if we want to provide it with IPv6 cidr, the Tenancy option controls whether <cloud>EC2</cloud> instances launched from the VPC use shared or dedicated hardware. dedicated hardware is more expensive. \
Once the VPC is created, under <kbd>Actions</kbd>, we can <kbd>Edit CIDR</kbd> to add cidr ranges.

### CIDR, Private vs Public IP
<details>
<summary>
Understaind IP ranges
</summary>

CIDR - classless inter domain routing, a method for allocating IP addresses. it is composed of the base ip and a subnet mask (how many bits can change in the ip).

- WW.XX.YY.ZZ/32 -> one IP
- 0.0.0.0/0 -> all ips
- 192.168.0.0/26 => all ips between [192.168.0.0, 192.168.0.63]

subnet masks:
- /0  = 000.000.000.000
- /8  = 255.000.000.000
- /16 = 255.255.000.000
- /24 = 255.255.255.000
- /32 = 255.255.255.255

number of possible address: 


$$
\begin{align*}
2^{(32 - mask)}\\
2^{(32-24)}= 2^8 = 256 \\
2^{(32-26)}= 2^6 = 64 \\
\end{align*}
$$

>The Internet Assigned Numbers Authority (IANA) established certain blocks of IPv4 addresses for the use of private (LAN) and public (Internet) addresses.\
> Private IP can only allow certain values:
> - 10.0.0.0 - 10.255.255.255 (10.0.0.0/8) - in big networks
> - 172.16.0.0 - 172.31.255.255 (172.16.0.0/12) - AWS default VPC in that range
> - 192.168.0.0 - 192.168.255.255 (192.168.0.0/16) - e.g., home networks
> 
> All the rest of the IP addresses on the Internet are Public.

</details>

### Default VPC Overview

<details>
<summary>
The VPC AWS created for us
</summary>

all AWS accounts have a default VPC, and new instances are launched in this one if no subnet is specified. the default VPC has internet connectivity and all <cloud>EC2</cloud> instances in it have public IPv4 addresses (and public and private DNS names).

in the demo, the cidr block is 172.31.0.0/16 ([172.31.0.0, 172.31.255.255]), which is inside the 172.16.0.0/12 range ([172.16.0.0, 172.31.255.255]), which has 65,536 ips ($2^{16}$).

the VPC has three subnets, each in a different Availability Zone with a cidr range which is part of the VPC cidr.
- 172.31.0.0/20 - [172.31.0.0, 172.31.15.255]
- 172.31.16.0/20 - [172.31.16.0, 172.31.31.255]
- 172.31.32.0/20 - [172.31.32.0, 172.31.47.255]

each of the ranges covers 4096 ($2^{(32-20)}$) ip addresses, but there are 5 reserved addresses.

THE VPC has a route table and Network ACL (NACL) defined. the route table has routes defined in cidr terms, one with a destination range corresponding to the VPC cidr range (172.31.0.0/16) targeting local (inside the VPC), and one for the rest of the possible ip ranges (0.0.0.0/0), targeting at the internet gateway.\
subnets which aren't explicitly associated with a table use the main route table, so the default VPC subnets are implicitly associated with it.
The default NACL inbound and outbound rules allowing any source to access the VPC and allows resources inside the VPC to access any address outside it.\

</details>

### Basic VPC Componenets

<details>
<summary>
Components of the VPC: Subnets, Internet Gateways & Route Tables
</summary>

subnets are created within Availability Zone, they are ranges of IPs from the VPC cidr. we usually define subnets as either public or private, but the distinction is artificial.\
inside each subnet, AWS reserves 5 ip address: the first 4 address, and the last address.

> These 5 IP addresses are not available for use and can't be assigned to an <cloud>EC2</cloud> instance. Example: if CIDR block 10.0.0.0/24, then reserved IP addresses are:
> 
> - 10.0.0.0 - Network Address
> - 10.0.0.1 - reserved by AWS for the VPC router
> - 10.0.0.2 - reserved by AWS for mapping to Amazon-provided DNS
> - 10.0.0.3 - reserved by AWS for future use
> - 10.0.0.255 - Network Broadcast Address.

in our vpc, we can go to the subnet section, click <kbd>Create Subnet</kbd>, choose our new VPC, and give the subnet a name, choose the Availability Zone, and the cidr range. for public subnets, we want to keep the range relatively small. we can add more subnets as needed.\
we can define the default setting of auto-assinging public IP addresses to instances in the subnet.

#### Internet Gateway, Route Tables

IGW - Internet Gateway

Internet gateways allows resources (<cloud>EC2</cloud>) in a VPC to connect to the internet. the gateways are managed - scale horizontally for high availability and redundancy. they must be created separately from the VPC. each gateway can be connected to one VPC, and each VPC can connect to only one gateway.\
In order for internet gateways to provide access to the internet, they must be combined with a route table.

we can check the default behavior (not having internet access) by launching a virtual machine (with a security group allowing SSH access), and trying to connect ot it with "EC2 Instance Connect".

we next <kbd>Create Internet Gateway</kbd>, wi simply provide a name to it an other optional tags. we then <kbd>Actions</kbd> and <kbd>Attach to VPC</kbd>. we next <cloud>Create Route Table</cloud> in our VPC, one called public and one private. then we associate the subnets to the route tables.\
in the public route table, we click <kbd>Edit Routes</kbd> and add the 0.0.0.0/0 route with the internet gateway as the target. now we can finally connect to the virtual machine with instance connect.
</details>

### Bastion Hosts, NAT Instances and NAT Gateways

<details>
<summary>
Accessing Resource inside a private subnet, Outbound Internet access from private subnets.
</summary>

#### Bation Hosts
sometimes called bastion, sometimes called jumpbox.

a bastion Host resides in the public subnet, has it's own security group, and it does have access to the private subnet. so we first connect from the public internet to the bastion host, and then we connect to private instances.\
The bastion security group must allow SSH access from the internet, usually a restricted CIDR (just one cidr, or the corporate cidr range). the private security group must allow SSH access from either the Bastion security group, or from the ip of the bastion host.

demo:\
Since we already created an <cloud>EC2</cloud> machine and allowed for SSH access to it, we can rename it to be "Bastion Host". we will an SSH key pair, and we launch a new instance in a private subnet. the security group will allow ssh access from the public security group, rather than from the public internet.

we connect to the bastion host using instance connect, and from it we run the ssh command `ssh ec2-user@10.0.22.82 -i demoKeyPair.pem`, we can try to ping google from inside the private instance, but this won't work.

#### NAT Instances

NAT = Network Address Translation

> NAT Instances allow EC2 instances in private subnets to connect to the Internet.
> - Must be launched in a public subnet
> - Must disable EC2 setting: Source/destination Check
> - Must have Elastic IP attached to it
> - Route Tables must be configured to route traffic from private subnets to the NAT Instance

NAT instances are not considered best practice anymore, and NAT gateways should be used instead.

> - Pre-configured Amazon Linux AMI was available - Reached the end of standard support on December 31, 2020
> - Not highly available / resilient setup out of the box - You need to create an ASG in multi-AZ + resilient user-data script
> - Internet traffic bandwidth depends on EC2 instance type
>
> You must manage Security Groups & rules:
> - Inbound:
>   - Allow HTTP / HTTPS traffic coming from Private Subnets
>   - Allow SSH from your home network (access is provided through Internet Gateway)
> - Outbound:
>   - Allow HTTP / HTTPS traffic to the Internet

we can still create the NAT instance in our VPC, we create a <cloud>EC2</cloud> machine, name it "NAT Instance", and search for an AMI which was called NAT. we put it the public sub net, and give it it's own security group in, which allows HTTP and HTTPS access from the cidr block of our VPC.\
we click <kbd>Edit Connection Settings</kbd>, <kbd>Source Destination Check</kbd> and disable it.\
we next go to the private route table, and we add a route. the route cidr is 0.0.0.0/0 (everything), and the target is the NAT instance we created.
(we also need to allow the ping protocol in the security group).

#### NAT Gatways

we can remove the NAT instances (and the route table entry) and use a NAT gateway instead.

> - AWS-managed NAT, higher bandwidth, high availability, no administration
> - Pay per hour for usage and bandwidth
> - NAT-GW is created in a specific Availability Zone, uses an Elastic IP
> - Can't be used by EC2 instance in the same subnet (only from other subnets)
> - Requires an IGW (Private Subnet => NAT-GW => IGW)
> - 5 Gbps of bandwidth with automatic scaling up to 100 Gbps
> - No Security Groups to manage / required
> 
> High Availability
> - NAT Gateway is resilient within a single Availability Zone
> - Must create multiple NAT Gateways in multiple AZs for fault-tolerance
> - There is no cross-AZ failover needed because if an AZ goes down it doesn't need NAT

| Operation                     | NAT Gateway                                       | NAT Instance                                          |
| ----------------------------- | ------------------------------------------------- | ----------------------------------------------------- |
| Availability                  | Highly available within AZ (create in another AZ) | Use a script to manage failover between instances     |
| Bandwidth                     | Up to 100 Gbps                                    | Depends on EC2 instance type                          |
| Maintenance                   | Managed by AWS                                    | Managed by you (e.g., software, OS patches, etc...)   |
| Cost                          | Per hour & amount of data transferred             | Per hour, EC2 instance type and size, + network costs |
| Public                        | IPv4                                              | Yes                                                   | Yes |
| Private                       | IPv4                                              | Yes                                                   | Yes |
| Security Groups               | Not required                                      | required                                              |
| Can be used  as Bastion Host? | No                                                | Ye                                                    |

we <kbd>Create NAT Gateway</kbd> in the public subnet, choose the connectivity type to be public, and allocate an elastic IP to it.\
then we edit the route table to direct traffic from the private subnet to the gateway.

</details>


### DNS Resolution Options & Route 53 Private Zones

<details>
<summary>
DNS resolution
</summary>

> DNS Resolution (enableDnsSupport)
> 
> - Decides if DNS resolution from Route 53 Resolver server is supported for the VPC
> - True (default): it queries the Amazon Provider DNS Server at 169.254.169.253 or the reserved IP address at the base of the VPC IPv4 network range plus two (.2) (reserved address)
>
> DNS Hostnames (enableDnsHostnames)
> 
> - By default, True => default VPC, False => newly created VPCs
> - Won't do anything unless enableDnsSupport=true
> - If True, assigns public hostname to EC2 instance if it has a public IPv4

if we use custom DNS domains names in a private hosted zone in <cloud>Route53</cloud>, both attributes must be enabled.

in our demo VPC, we <kbd>Edit DNS resolution</kbd> and <kbd>Edit DNS Hostnames</kbd> and enable both. now the bastion host server, which has a public ipv4 address, will also have a public dns name.\
In <cloud>Route53</cloud>, we can create a private hosted zone, and associate it with our demo VPC.
</details>

### NACL & Security Groups

<details>
<summary>
Security Controls
</summary>

EC2 instances have security group, but the subnet itself can have an network access control list. the NACL is stateless (unlike security groups).

> NACL are like a firewall which control traffic from and to subnets. One NACL per subnet, new subnets are assigned the Default NACL.\
> You define NACL Rules:
> - Rules have a number (1-32766), higher precedence with a lower number
> - First rule match will drive the decision
> - Example: if you define #100 ALLOW 10.0.0.10/32 and #200 DENY 10.0.0.10/32, the IP Address will be allowed because 100 has a higher precedence over 200
> - The last rule is an asterisk (*) and denies a request in case of no rule match
> - AWS recommends adding rules by increment of 100
> 
> Newly created NACLs will deny everything.\
> NACL are a great way of blocking a specific IP address at the subnet level.

default NCL accepts everything inbound/outbound. this is used for the default VPC.

when two endpoint connect to one another, they use ports, the client connects to a defined port, and expects a response on an ephemeral port. different OS use different port ranges for ephemeral ports. the requests from the client to the server contains the ephemeral port, which is where the server will send it's response to.\
The NACL must allow outbound and inbound access on the ephemeral port.

| Metric          | Security Group                                                             | NACL                                                                                                      |
| --------------- | -------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| Effects         | Operates at the instance level                                             | Operates at the subnet level                                                                              |
| Allow/Deny      | Supports allow rules only                                                  | Supports allow rules and deny rules                                                                       |
| Statefullness   | Stateful: return traffic is automatically allowed, regardless of any rules | Stateless: return traffic must be explicitly allowed by rules (think of ephemeral ports)                  |
| Rule Evaluation | All rules are evaluated before deciding whether to allow traffic           | Rules are evaluated in order (lowest to highest) when deciding whether to allow traffic, first match wins |
| Applies to      | Applies to an EC2 instance when specified by someone                       | Automatically applies to all EC2 instances in the subnet that it's associated with                        |

demo:\
in our bastion instance, we can start an HTTP server, we download the "httpd" package and start it `sudo systemctl enable httpd` and `sudo systemctl start httpd`. we enable the HTTP access in the security group, and we can access it from the browser. we can edit the NACL to deny HTTP traffic with a lower rule number (higher priority), and now we won't be able to connect to it. Security groups and NACL work together, and both must pass.
we can remove outbound rules from the Security group, and then we can still connect to it, but it won't be able to initiate traffic from within. this is what being 'stateful' means for security groups.
</details>


### VPC Reachability Analyzer

<details>
<summary>
Network Diagnostics Tool
</summary>

> A network diagnostics tool that troubleshoots network connectivity between two endpoints in your VPC(s).\
> It builds a model of the network configuration, then checks the reachability based on these configurations (it doesn't send packets).\
> When the destination is:
> - Reachable - it produces hop-by-hop details of the virtual network path
> - Not reachable - it identifies the blocking component(s) (e.g., configuration issues in SGs, NACLs, Route Tables)
> - Use cases: troubleshoot connectivity issues, ensure network configuration is as intended

in the "Reachability Analyzer" page, we choose the source and destination, choose ports and protocol, and analyze the path. if the path is not reachable, we can see the reason why, like missing inbound rules in a security group. it also shows the entire path for the requests.
</details>


### VPC Peering

<details>
<summary>
Connect VPCs together
</summary>

> Privately connect two VPCs using AWS'
network. Make them behave as if they were in the
same network.
> - Must not have overlapping CIDRs.
> - VPC Peering connection is **NOT transitive** (must be established for each VPC that need to communicate with one another).
> - You must update route tables in each VPC's subnets to ensure EC2 instances can communicate with each other.
> - You can create VPC Peering connection between VPCs in different AWS accounts/regions
> - You can reference a security group in a peered VPC (works cross accounts - same region)

if VPC A is peered to VPC B and VPC is peered to VPC C, A and C are not peered.

in our demo, we can peer the new VPC with the default VPC, we can do this by creating an <cloud>EC2</cloud> instance and trying to use the private ip address of the webserver we set up in a different VPC.\
We <kbd>Create Peering Connection</kbd>, give the connection a name, choose the requester VPC (one of our VPCs), and which VPC to peer with. that VPC can be in our account or another account, in the same region and or another region. we can see the CIDR blocks and determine if they are overlapping. if not, great.\
Before the peering is established, the target VPC must accept the request. next, we need to modify the route tables in both VPCs. under routes, we add a route with destination of the cidr range from the peered VPC, and the target is the peered connection. we do the same with the other VPC.

</details>

### VPC Endpoints

<details>
<summary>
Use serverless resources from the VPC without going through the public internet
</summary>

> Every AWS service is publicly exposed (public URL).\
> VPC Endpoints (powered by AWS PrivateLink) allows you to connect to AWS services using a private network instead of using the public Internet
> 
> - They're redundant and scale horizontally
> - They remove the need of IGW, NAT-GW to access AWS Services
> - In case of issues:
>   - Check DNS Setting Resolution in your VPC
>   - Check Route Tables

A VPC endpoint is deployed in the VPC, and allows resources to communicate with AWS resources within the AWS network, rather than through the public internet.

there are two types of endpoints, interface endpoints (most services, costs money), and gateway endpoints (<cloud>S3</cloud> and <cloud>DynamoDb</cloud> only, free).

> Interface Endpoints (powered by PrivateLink)
> - Provisions an ENI (private IP address) as an entry point 
> - Must attach a Security Group
> - Supports most AWS services
> - cost: $ per hour + $ per GB of data processed
> 
> Gateway Endpoints
> - Provisions a gateway and must be used as a target in a route table
> - does not use security groups
> - Supports both <cloud>S3</cloud> and <cloud>DynamoDB</cloud>
> - Free

we could use interface endpoint to connect to <cloud>S3</cloud> and <cloud>DynamoDB</cloud>, but gateway endpoints are free, easier to use, and scale better. interface endpoints are only preferable when used from on-premises site, or from a different VPC or region.

demo:\
from our private <cloud>EC2</cloud> instance, we first give it role to perform <cloud>S3</cloud> operations, we can run `aws s3 ls` to list our buckets, and this operation used the public internet. we can remove the destination of the internet gateway from the private route table, and now the same operation fails. we can also run `curl example.com` command to check against the public internet.\
in the "Endpoints" section, we <kbd>Create Endpoint</kbd>, give it a name, choose the "AWS Services" category, and search for S3 as the service name, we will have either interface endpoint or gateway endpoint.\
For interface endpoints, we choose the VPC, ensure DNS names are enabled, and we choose which subnets to deploy the endpoint (ENI device) in, and select the security group.
For gateway endpoints, we select the VPC, and choose which routeTable to update. we can modify the endpoint access policy if we wish. now, in the route table, we will see an entry for the endpoint.

back in our instance, we can try running `curl` again, which will still fail, and we need to run `aws s3` command with the correct region flag (`--region eu-central-1`), but the command will work properly.


</details>

### VPC Flow Logs

<details>
<summary>
Monitor Traffic Data
</summary>

> Capture information about IP traffic going into your interfaces:
> 
> - VPC Flow Logs
> - Subnet Flow Logs
> - Elastic Network Interface (ENI) Flow Logs
> 
> Helps to monitor & troubleshoot connectivity issues.

We can send the data to different locations

- <cloud>S3</cloud>
- <cloud>CloudWatch Logs</cloud>
- <cloud>Kinesis Data Firehose</cloud>

There is also network information from specific services:
- <cloud>ELB</cloud> - Elastic Load Balance
- <cloud>RDS</cloud> - Relational Databases
- <cloud>ElastiCache</cloud> - Cache
- <cloud>Redshift</cloud> - Data warehouse (BI)
- <cloud>WorkSpaces</cloud> - VDI
- <cloud>NAT Gateways</cloud>
- <cloud>Transit Gateway</cloud>

the flow logs format:
- version
- accountId
- interfaceId
- srcAddr (source address)
- dstAddr (destination address)
- srcPort (source port)
- dstPort (destination port)
- protocol
- packets
- byres
- start timestamp
- end timestamp
- action (Accept or Reject)
- log-status

we can query VPC flow logs using <cloud>Athena</cloud> on <cloud>S3</cloud> or <cloud>CloudWatch</cloud> Logs Insight.

> - srcAddr & dstAddr - help identify problematic IP
> - srcPort & dstPort - help identity problematic ports
> - Action - success or failure of the request due to Security Group / NACL
> - Can be used for analytics on usage patterns, or malicious behavior

in our demo VPC, in the Flow Logs tab, <kbd>Create Flow Log</kbd>, we give it a name, a filter (Accept, Reject, All), choose aggregation interval (usually we use the duration) and the destination (<cloud>CloudWatch</cloud> Logs or <cloud>S3</cloud> bucket, <cloud>Kinesis Firehose</cloud>), it will need a proper <cloud>IAM</cloud> role. we can also choose to use the custom format or modify it. in the demo, we send it both to a bucket and to a log group.\

In <cloud>Athena</cloud>, we fist set up a query result location (a new bucket), and we create a database for VPC flow logs. we paste the proper `CREATE TABLE` and modify the locations, and then we can RUN SQL queries on the flow logs data.

</details>

### Acessing VPCs

<details>
<summary>
Different Ways to Connect on-premises resources to AWS
</summary>

connecting on-premises data center to AWS.

#### Site to Site VPN, Virtual Private Gateway & Customer Gateway

VPN uses the public internet. we have a VPN Gateway in the VPC, and a customer gateway on-premises. the customer gateway can be virtual or a dedicated device.

> Virtual Private Gateway (VGW)
> - VPN concentrator on the AWS side of the VPN connection
> - VGW is created and attached to the VPC from which you want to create the Site-to-Site VPN connection
> - Possibility to customize the ASN (Autonomous System Number)
> 
> Customer Gateway (CGW)
> - Software application or physical device on customer side of the VPN connection

if our customer gateway device (on-premises) has a public (unchanging) IP address, we should use that. if it is behind a NAT device, then we use the NAT device public ip. in either case, we need to enable *route propagation* for the virtual private gateway in the route table. we also need to make sure the security group allow for the ICMP (ping) protocol.

if we have multiple on-premises sites, we can use <cloud>AWS VPN CloudHub</cloud> and they can communicate with one another.

> - Provide secure communication between multiple sites, if you have multiple VPN connections
> - Low-cost hub-and-spoke model for primary or secondary network connectivity between different  locations (VPN only)
> - It's a VPN connection so it goes over the public Internet
> - To set it up, connect multiple VPN connections on the same VGW, setup dynamic routing and configure route tables

#### Direct Connect & Direct Connect Gateway

<cloud>Direct Connect</cloud> (DX)

> - Provides a dedicated private connection from a remote network to your VPC
> - Dedicated connection must be setup between your DC and AWS Direct Connect locations
> - You need to setup a Virtual Private Gateway on your VPC
> - Access public resources (<cloud>S3</cloud>) and private (<cloud>EC2</cloud>) on same connection
> - Use Cases:
>   - Increase bandwidth throughput - working with large data sets - lower cost
>   - More consistent network experience - applications using real-time data feeds
>   - Hybrid Environments (on prem + cloud)
> - Supports both IPv4 and IPv6

use private virtual interface for VPC resources, and public virtual interface for services outside it. uses Virtual private gateway, just like VPN connections..

if we want to connect to multiple VPCs (in different regions), we use a <cloud>Direct Connect</cloud> gateway.

> Dedicated Connections:
> - 1Gbps, 10 Gbps and 100 Gbps capacity
> - Physical ethernet port dedicated to a customer
> - Request made to AWS first, then completed by AWS Direct Connect Partners
> 
> Hosted Connections
> - 50Mbps, 500 Mbps, to 10 Gbps
> - Connection requests are made via AWS Direct Connect Partners
> - Capacity can be added or removed on demand
> - 1, 2, 5, 10 Gbps available at select AWS Direct Connect Partners
> 
> - Lead times are often longer than 1 month to establish a new connection
> - Data in transit is not encrypted but is private
> - AWS Direct Connect + VPN provides an IPsec-encrypted private connection
> - Good for an extra level of security, but slightly more complex to put in place
> - In case Direct Connect fails, you can set up a backup Direct Connect connection (expensive), or a Site-to-Site VPN connection

</details>

### VPC Endpoints

<details>
<summary>
Exposing A single Service
</summary>

#### AWS PrivateLink - VPC Endpoint Services

exposing VPCs to others VPCs. there are several options, first, they can be made public and accessed through the public internet. a second option is to use VPC peering - this requires multiple connections, and it opens up the entire VPC. but if we want just a single resource inside the network, we can use <cloud>PrivateLink</cloud>.

PrivateLink is a vpc endPoint service, it can securely expose services across VPCs (in the same account or across different accounts) witout peering, internet gateways, NAT and route tables.

the service VPC needs a Network Load Balancer (or gateway load balancer), and the customer VPC needs an ENI (elastic network interface) or a virtual private gateway (for on-premises), then we connect the two VPCs with privateLink.\
For example, we can have ECS clusters that we wish to expose as as service. we connect them to an application load balancer, which is then connected to a network load balancer. on the customer side, we can either have an ENI or a Virtual private gateway.

in the <cloud>VPC</cloud> service, under "Endpoint services", <kbd>Create EndPoint Service</kbd>, we choose the load balancer type. once we create this, we can go to "endPoints", <kbd>Create EndPoint</kbd>, and choose the service category as "other".

#### AWS ClassicLink

the older methods of connecting, either <cloud>EC2-classic</cloud> and <cloud>ClassicLink</cloud>. not relevant anymore.

</details>

### Transit Gateway

<details>
<summary>
Hub and Spoke Model for transitive connections.
</summary>

Hub and Spoke, connect VPCs, <cloud>Customer Gateway</cloud>, <cloud>Direct Connect</cloud> networks together.

> - For having transitive peering between thousands of VPC and on-premises, hub-and-spoke (star) connection
> - Regional resource, can work cross-region
> - Share cross-account using <cloud>Resource Access Manager</cloud> (RAM)
> - You can peer Transit Gateways across regions
> - Route Tables: limit which VPC can talk with other VPC
> - Works with Direct Connect Gateway, VPN connections
> - Supports IP Multicast (not supported by any other AWS service)

also allows for ECMP (Equal Cost MultiPath Routing), multiple site-to-site VPN connection to increase the bandwidth. it gives us more tunnels as compared to connecting to a Virtual Private Gateway directly.

we can also use transit gateway to share <cloud>Direct Connect</cloud> between multiple accounts, 
</details>

### VPC Traffic Mirroring

<details>
<summary>
Capture and inspect network traffic in your VPC
</summary>

> Allows you to capture and inspect network traffic in your VPC.
> Route the traffic to security appliances that you manage.\
> Capture the traffic:
> - From (Source) - ENIs
> - To (Targets) - an ENI or a Network Load Balancer
>
> Capture all packets or capture the packets of your interest (optionally, truncate packets).
> Source and Target can be in the same VPC or different VPCs (VPC Peering).\
> Use cases: content inspection, threat
monitoring, troubleshooting, etc...
</details>

### IPv6 for VPC

<details>
<summary>
IPv6 in AWS
</summary>

> IPv4 designed to provide 4.3 Billion addresses (they'll be exhausted soon).
> IPv6 is the successor of IPv4.
> IPv6 is designed to provide $3.4 * 10^{38}$ unique IP addresses.\
> Every IPv6 address in AWS is public and Internet-routable (no private range).\
> Format x:x:x:x:x:x:x:x (x is hexadecimal, range can be from 0000 to ffff).
> 
> Examples:
> - 2001:db8:3333:4444:5555:6666:7777:8888
> - 2001:db8:3333:4444:cccc:dddd:eeee:ffff
> - :: - all 8 segments are zero
> - 2001:db8:: - the last 6 segments are zero
> - ::1234:5678 - the first 6 segments are zero
> - 2001:db8::1234:5678 - the middle 4 segments are zero

we can't enable IPv4 for VPC subnets, but we can **enable IPv6** and operate in dual mode. then the <cloud>EC2</cloud> instances will have a private IPv4 address, and a public IPv6 address.

inside the <cloud>VPC</cloud> settings, we <kbd>Edit Cidr</kbd> and choose to <kbd>Add IPv6 Cidr</kbd>. we can either get one provided by AWS or use a cidr owned by us. we can then assign the cidr to the subnets, and make them assign IPv6 address automatically. for each <cloud>EC2</cloud> that already exists, we can assign an address from the range we have. then we could fix the security group to also allow traffic from IPv6 addresses.

we also see the IPv6 cidr in the route table.

#### Egress Only Internet Gateway

used only for IPv6. similar to NAT gateways. 

> Allows instances in your VPC outbound connections over IPv6 while preventing the internet to initiate an IPv6 connection to your instances.
> You must update the Route Tables.


public subnet route table

| Destination             | Target | Notes                       |
| ----------------------- | ------ | --------------------------- |
| 10.0.0.0/16             | local  | IPv4 address range          |
| 2001:db8:1234:1a00::/56 | local  | IPv6 address range          |
| 0.0.0.0/0               | intrent gateway-id | IPv4 access public internet |
| ::/0                    | intrent gateway-id | IPv6 access public internet |

private subnet route table

| Destination             | Target         | Notes                       |
| ----------------------- | -------------- | --------------------------- |
| 10.0.0.0/16             | local          | IPv4 address range          |
| 2001:db8:1234:1a00::/56 | local          | IPv6 address range          |
| 0.0.0.0/0               | nat-gateway-id | IPv4 access public internet |
| ::/0                    | egress-only intrent gateway-id  | IPv6 access public internet |

demo:\
inside the <cloud>VPC</cloud> section, we choose the "Egress Only Internet Gateway" section, <kbd>Create Egress Only Internet Gateway</kbd>, attach it to a VPC, and edit the route table in the private subnet.

</details>


### Networking Costs in AWS

<details>
<summary>
Understanding costs of networking.
</summary>

We usually pay for GB of network traffic.\
Inbound traffic is Free, traffic inside the same Availability Zone is free. it's cheaper to use private IPs to communicate between Availability Zones then using the public/elastic IP, (0.01$ per GB vs 0.02$ Per GB). traffic across regions is 0.02$ per GB.

> - Use Private IP instead of Public IP for good savings and better network performance 
> - Use same Availability Zone for maximum savings (at the cost of high availability)

> Minimizing egress traffic network cost
> - Egress traffic: outbound traffic (from AWS to outside)
> - Ingress traffic: inbound traffic - from outside to AWS (typically free)
> - Try to keep as much internet traffic within AWS to minimize costs
> - <cloud>Direct Connect</cloud> location that are co-located in the same AWS Region result in lower cost for egress network

pricing for <cloud>S3</cloud> traffic (USA costs)

> - S3 ingress: free
> - S3 to Internet: $0.09 per GB
> - S3 Transfer Acceleration:
>   - Faster transfer times (50% to 500% better)
>   - Additional cost on top of Data Transfer Pricing: +$0.04 to $0.08 per GB
> - S3 to <cloud>CloudFront</cloud>: $0.00 per GB (free!)
> - <cloud>CloudFront</cloud> to Internet: $0.085 per GB (slightly cheaper than S3)
>   - Caching capability (lower latency)
>   - Reduce costs associated with S3 Requests Pricing (7x cheaper with CloudFront)
> - S3 Cross Region Replication: $0.02 per GB


NAT Gateway vs Gateway VPC Endpoint, example of accessing a S3 bucket. either through the public internet (NAT gateway to internet gateway) or thorugh a VPC endpoint (don't forget the route table). it's much cheaper and easier to use an VPC endpoint.

> - $0.045 NAT Gateway / hour  (cost for existing)
> - $0.045 NAT Gateway data processed / GB
> - $0.09 Data transfer out to S3 (cross-region)
> - $0.00 Data transfer out to S3 (same-region)
> 
> - No cost for using Gateway Endpoint.
> - $0.01 Data transfer in/out (same-region)


</details>

### Network Firewall

<details>
<summary>
Protect the entire VPC
</summary>


> To protect network on AWS, we’ve seen
> - Network Access Control Lists (NACLs)
> - Amazon VPC security groups
> - <cloud>AWS WAF</cloud> (protect against malicious requests)
> - <cloud>AWS Shield</cloud> & AWS Shield Advanced
> - <cloud>AWS Firewall Manager</cloud> (to manage them across accounts)
>
> But what if we want to protect our entire VPC in a sophisticated way?
>
> AWS Network Firewall- Protect your entire Amazon VPC
> - From Layer 3 to Layer 7 protection
> - Any direction, you can inspect
>   - VPC to VPC traffic (VPC peering)
>   - Outbound to internet
>   - Inbound from internet
>   - To / from Direct Connect & Site-to-Site VPN
> 
> Internally, the AWS Network Firewall uses the AWS Gateway Load Balancer.
> Rules can be centrally managed cross-account by AWS Firewall Manager to apply to many VPCs.\
> - Supports 1000s of rules:
>   - IP & port - example: 10,000s of IPs filtering
>   - Protocol – example: block the SMB protocol for outbound communications
>   - Stateful domain list rule groups: only allow outbound traffic to *.myCorp.com or third-party software repo
>   - General pattern matching using regex
> - Traffic filtering: Allow, drop, or alert for the traffic that matches the rules.
> - Active flow inspection to protect against network threats with intrusion prevention capabilities (like <cloud>Gateway Load Balancer</cloud>, but all managed by AWS)
> - Send logs of rule matches to Amazon <cloud>S3</cloud>, <cloud>CloudWatch Logs</cloud>, <cloud>Kinesis Data Firehose</cloud>

</details>

Summary:

> - CIDR - IP Range
> - VPC - Virtual Private Cloud - we define a list of IPv4 & IPv6 CIDR
> - Subnets - tied to an AZ, we define a CIDR
> - Internet Gateway - at the VPC level, provide IPv4 & IPv6 Internet Access
> - Route Tables - must be edited to add routes from subnets to the IGW, VPC Peering Connections, VPC Endpoints, etc...
> - Bastion Host - public EC2 instance to SSH into, that has SSH connectivity to EC2 instances in private subnets
> - NAT Instances - gives Internet access to EC2 instances in private subnets. Old, must be setup in a public subnet, disable Source / Destination check flag
> - NAT Gateway - managed by AWS, provides scalable Internet access to private EC2 instances, when the target is an IPv4 address
> - Private DNS + Route 53 - enable DNS Resolution + DNS Hostnames (VPC)
> - NACL - stateless, subnet rules for inbound and outbound, don't forget Ephemeral Ports
> - Security Groups - stateful, operate at the EC2 instance level
> - Reachability Analyzer - perform network connectivity testing between AWS resources
> - VPC Peering - connect two VPCs with non overlapping CIDR, non-transitive
> - VPC Endpoints - provide private access to AWS Services (<cloud>S3</cloud>, <cloud>DynamoDB</cloud>, <cloud>CloudFormation</cloud>, <cloud>SSM</cloud>) within a VPC
> - VPC Flow Logs - can be setup at the VPC / Subnet / ENI Level, for `ACCEPT` and `REJECT` traffic, helps identifying attacks, analyze using <cloud>Athena</cloud> or <cloud>CloudWatch Logs Insights</cloud>
> - Site-to-Site VPN - setup a Customer Gateway on DC, a Virtual Private ateway on VPC, and site-to-site VPN over public Internet
> - AWS VPN CloudHub - hub-and-spoke VPN model to connect your sites
> - Direct Connect - setup a Virtual Private Gateway on VPC, and establish a direct private connection to an AWS Direct Connect Location
> - Direct Connect Gateway - setup a Direct Connect to many VPCs in different AWS regions
> - AWS PrivateLink / VPC Endpoint Services:
>   - Connect services privately from your service VPC to customers VPC
>   - Doesn't need VPC Peering, public Internet, NAT Gateway, Route Tables
>   - Must be used with Network Load Balancer & ENI
> - ClassicLink - connect EC2-Classic EC2 instances privately to your VPC (deprecated)
> - Transit Gateway - transitive peering connections for VPC, VPN & DX
> - Traffic Mirroring - copy network traffic from ENIs for further analysis
> - Egress-only Internet Gateway - like a NAT Gateway, but for IPv6
</details>

## Other Services

<details>
<summary>
Additional AWS Services
</summary>

### X-Ray

<details>
<summary>
Trace requests across services
</summary>

visual analysis of performance and dependencies.

> - Debugging in Production, the good old way:
>   - Test locally
>   - Add log statements everywhere
>   - Re-deploy in production
> - Log formats differ across applications and log analysis is hard.
> - Debugging: one big monolith "easy", distributed services "hard".
> - No common views of your entire architecture

needs to be enabled.

</details>

### AWS Amplify

<details>
<summary>
Web and Mobile Application Development
</summary>

create applications that use AWS as backend, and use frontend libraries to connect. then deploy directly to <cloud>CloudFront</cloud> or <cloud>Amplify Console</cloud>.

like <cloud>Elastic BeanStalk</cloud> for mobile applications.

> A set of tools and services that helps you develop and deploy scalable full stack web and mobile applications
> - Authentication
> - Storage
> - API (REST, GraphQL)
> - CI/CD
> - PubSub
> - Analytics
> - AI/ML Predictions
> - Monitoring
> 
> Connect your source code from GitHub, AWS CodeCommit, Bitbucket, GitLab, or upload directly.
</details>

</details>

## Misc

<details>
<summary>
Stuff worth remembering.
</summary>


arn format: `arn:<aws-partition>:<service>:<region>:<account-number>:<resource>`

partition is usually aws, but can also one of the following:
- aws
- aws-us-gov
- aws-cn

The resource suffix can be one of the following, and can include a parent resource, sub type, resource path, version, etc...

- `<resource-id>`
- `<resource-type>/<resource-id>`
- `<resource-type>:<resource-id>`

- IPv4 format is x.x.x.x (x is 0x00 to 0xff)
- IPv6 format is x:x:x:x:x:x:x:x, we can omit any part and have :: to indicate zeros. (x is 0x00 to 0xffff)
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
- DLM - Data Lifecycle Manager (backup EBS)
- FSR - Fast Snapshot Restore (restore fully initiliazed EBS snapshots to reduce latency)
- SSE - Server side encryption
- CORS - Cross Origin Resource Sharing
- WORM - Write Once, Read Many (Glacier)
- HPC - High Performance Computing
- PITR - Point in Time Recovery
- SCP - Service Control Policies (IAM organization)
- DRP - DDoS Respone Team - 24/7 support team when under DDoS attack (<cloud>AWS Shield</cloud> advance)
- SAML - Security Assertion Markup Language (IAM federation, Cognito)
- IdP - Identity Provider (IAM federation, Cognito)
- CUP - Cognito User Pools - users database for websites and applications managed by AWS
- CIP - Cognito Identity Pools (federated identities) - access to AWS resources.
- DNS - Domain Name System
- TLS - Top Level Domain (DNS) - ".com"
- SLD - Second Level Domain (DNS) - "amazon.com"
- FQDN - Fully Qualified Domain Name (DNS)
- ICANN - The Internet Corporation for Assigned Names and Numbers
- NACL - network Access Control List (<cloud>VPC</cloud>)
ECMP - Equal Cost MultiPath Routing (<cloud>Transit Gateway</cloud>)
</details>

<!-- misc end -->
</details>
