<!--
// cSpell:ignore HashiCorp
 -->

# Terraform For The Absolute Beginners

udemy course [Terraform for the absolute beginners](https://www.udemy.com/course/terraform-for-the-absolute-beginners/).

## Infrastructure as Code

<details>
<summary>
Iac - Infrastructure as Code
</summary>

### Challenges with Traditional IT Infrastructure

in the traditional model of deploying applications, we have a solution architect that specifies which hardware is needed, and it all needs to belong to the company and reside in the data center.
once the hardware is available, it still needs to pass by many teams before the application can be deployed.

- field engineers to install the physical machines
- system administrators to set them up
- storage admins to allocate space on the server
- backup admins
- and in the end, the application team.

this whole process can take weeks, and it's hard to scale up and down when demand fluctuates. this all requires manual human labor, so there are many errors.

moving to cloud can reduce this problem, as the company doesn't need to own the hardware, and we use a virtual machine instead, this makes deployment much faster.
cloud providers also have APIs rather than human labor, which makes automation easier.

automating infrastructure provisioning was the basis for infrastructure as code.

### Types of IAC Tools
rather than using the management UI console of the cloud provider, its easier to write code that does it for us. which is faster, easier, and easier to maintain.

this shell script
```sh
#!/bin/bash
IP_ADDRESS="10.2.2.1"

EC2_INSTANCE=$(ec2-run-instance --instance-type t2.micro ami-0edab43b6fa892279)

INSTANCE=$(echo ${EC2_INSTANCE} | sed 's/*INSTANCE //' | sed 's/ .*//')

# Wait for instance to be ready
while !ec2-describe-instances $INSTANCE | grep | -q "running"
do
	echo Waiting for $INSTANCE to be ready...
done

# Check if instance is not provisioned and exit
if [! $(ec2-describe-instances $INSTANCE | grep | -q "running")]; then
	echo Instance $INSTANCE is stopped
	exit
fi

ec2-associate-address $IP_ADDRESS -i $INSTANCE
echo Instance $INSTANCE was created successfully!
```
can be written as a terraform configuration file, which is easier to read.

```hcl
resource "aws_instance" "webserver"{
	ami = "ami-0edab43b6fa892279"
	instance_type = "t2.micro"
}
```

this ansible yaml also provisions aws resources.

```yaml
- amazon.aws.ec2:
    key_name: my-key
    instance_type: t2.micro
    image: ami-123456
    wait: yes
    group: webserver
    count: 3
    vpc_subnet_id: subnet-29e63245
    assign_public_ip: yes
```

there are all sorts of IaC tools, each of them has some uses.
- Configuration Management
- Server Templating
- Provisioning Tools
  - *Terraform*
  - *CloudFormation*

#### Configuration Managements Tools

> - Designed to install and manage Software on existing infrastructure
> - Maintain Standard Structure
> - Version Control
> - Idempotent (run the code many times, without messing things up)

examples:
- *Ansible*
- *SaltStack*
- *Puppet*

#### Server Templating

> - Pre-Installed Software and dependencies
> - Virtual Machine or Docker Images
> - Immutable Infrastructures - once deployed, replace rather than update.

examples:
- *Packer*
- *Docker*
- *Vagrant*

#### Provisioning Tools

> - Deploy Immutable Infrastructure resources
> - Multiple Providers

examples:
- *Terraform* - works with many vendors
- *CloudFormation* - aws specific


### Why Terraform?

a tool by HashiCorp, can work with multiple cloud vendors, both public and private. this is done with providers, which supply an api to a specific resource. this can be a cloud vendor, a network provider, databases or any external tool, even version control tools!


it uses HCL - hashicorp configuration language

this sample code declares an instance on the cloud.

```hcl
resource "aws_instance" "webserver"{
    ami= "ami-0edab43b6fa892279"
    instance_type="t2.micro"
}

resource "aws_s3_bucket" "finance" {
    bucket "finance-21092020"
    tags= {
        Description = "Finance and Payroll"
    }
}

resource "aws_iam_user" "admin-user"{
    name="lucy"
    tags= {
        Description = "Team Leader"
    }
}
```

It uses declarative style. it defines the desired state, and terraform takes care of getting us from the current state to the desired state.
phases:
- Init
- Plan
- Apply
  
any object managed by terraform is called a "resource", it can be a cloud resource, database or credentials. terraform also controls the lifetime of those objects.

terraform can also take care of resources that were created from other sources.

</details>

## Getting Started


### Installing Terraform

installing terraform from cli
```sh
wget https://releases.hashicorp.com/terraform/<ver>/<release>.zip
unzip <release>.zip
mv terraform /usr/local/bin
terraform version
```

lets start with a simple file "aws.tf"

```hcl
resource "aws_instance" "webserver"{
    ami= "ami-0c22f25c1f66a1ff4d"
    instance_type ="t2.micro"
}
```
a resource is something that terrafrom manages, such databases, roles, cloud resources and others. we will begin with a simple resource type: a local file and a resource called "pet".


### HashiCorp Configuration Language (HCL) Basics

the hcl syntax consistent of block and arguments.

```hcl
<block> <parameters> {
    key1 = value1
    key2 = value2
}
```
a block contains information about the infrastructure and resources inside the platfrom.
to create a file,

```sh
mkdir /root/terraform-local-file
cd /root/terraform-local-file
touch local.tf
```
and lets edit the new file

```hcl
resource "local_file" "pet" {
    filename = "/root/pets.txt"
    content = "We love pets!"
}
```
the type of the block is "resource", and we then provide the type of the resource, "local_file",this is actually a combination of the provider "local", underscore, and the resource type "file". then is the resource name, "pet". inside the block we start providing values (argument and parameters).\
These fields are specific to the resource type. each type expects different fields.

other resources can be, block type, resource type (provider+type), name, and then the needed arguments.

```hcl
resource "aws_instance" "webserver"{
    ami= "ami-0c22f25c1f66a1ff4d"
    instance_type ="t2.micro"
}

resource "aws_s3_bucket" "data"{
    bucket = "webserver-bucket-org-2207"
    acl = "private"
}
```

a terraform workflow has four steps:
- writing the configuration file
- run `init` to install plugins and create the plan
- review the exectuition plan
- execute the plan

```sh
terraform init
terraform plan
terraform apply
<confirm>
terraform show
cat /root/pets.txt
```

terraform supports many providers, the local providers is one of them. each provider has resources, and each resource can accept any number of arguments.

### Update and Destroy Infrastructure
### Lab Intro
### Demo: Accessing Labs
### Accessing the Labs
### Lab: HCL Basics

## Terraform Basic

## Terraform State

## Working With Terraform

## Remote State

## Terraform Provisioners

## Terraform Import, Tainting Resources and Debugging

## Terraform Modules

## Terraform Functions and Conditional Expressions



## Cli Commands

- `terraform version`
- `terraform init`
- `terraform plan`
- `terraform apply`
- `terraform show`