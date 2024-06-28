<!--
// cSpell:ignore
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

## Managin EC2 at scale - Systems Manager (SSM)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

### Systems Manager Overview
### Start EC2 Instances with SSM Agent
### AWS Tags & SSM Resource Groups
### SSM Documents & SSM Run Command
### SSM Automations
### [SAA/DVA] SSM Parameter Store Overview
### [SAA/DVA] SSM Parameter Store Hands On (CLI)
### SSM Inventory & State Manager
### SSM Patch Manager and Maintenance Windows
### SSM Patch Manager and Maintenance Windows - Hands On
47. SSM Session Manager Overview
48. SSM Session Manager Hands On
49. SSM Cleanup

</details>

## EC2 High Availability and Scalability
## Elastic BeanStalk for SysOps
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

- ENA - elastic network adapter
- EFA - Elastic fabric adapter, improved ENA

</details>


</details>

</details>