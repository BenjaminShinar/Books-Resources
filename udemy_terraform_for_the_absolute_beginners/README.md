<!--
// cSpell:ignore HashiCorp KodeKloud FIFA tfvars tfstate falshpoint Tsvg Flexit toset aone xlarge azurrerm untaint
 -->

# Terraform For The Absolute Beginners

udemy course [Terraform for the absolute beginners](https://www.udemy.com/course/terraform-for-the-absolute-beginners/).

[KodeKloud Lab](https://kodekloud.com/courses/udemy-labs-terraform-for-beginners/)

if copying from the lab doesn't work, try <kbd>shift + insert</kbd>

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
  - _Terraform_
  - _CloudFormation_

#### Configuration Managements Tools

> - Designed to install and manage Software on existing infrastructure
> - Maintain Standard Structure
> - Version Control
> - Idempotent (run the code many times, without messing things up)

examples:

- _Ansible_
- _SaltStack_
- _Puppet_

#### Server Templating

> - Pre-Installed Software and dependencies
> - Virtual Machine or Docker Images
> - Immutable Infrastructures - once deployed, replace rather than update.

examples:

- _Packer_
- _Docker_
- _Vagrant_

#### Provisioning Tools

> - Deploy Immutable Infrastructure resources
> - Multiple Providers

examples:

- _Terraform_ - works with many vendors
- _CloudFormation_ - aws specific

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

<details>
<summary>
First Steps.
</summary>

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

we also need to update and sometimes destroy infrastructure that we created.

to update, we first modify the terraform file. like changing the file permissions.

```hcl
resource "local_file" "pet" {
    filename = "/root/pets.txt"
    content = "We love pets!"
    file_permission = "0700"
}
```

we then run `terraform plan`, which informs us that the file needs to be replaced (not updated in place). this file as an **immutable infrastructure**. to move along with change, we run the `terraform apply` command.

if we wish to delete the infrastructure, we can run `terraform destroy`, which also requires confirmation. this will delete all the resources in the current directory.

### Lab Intro

each lab has some exercises for us to train with. there is a terminal, a vscode editor, and half a screen is dedicated to the question. we might need to perform queries in the terminal to inspect the configuration and the infrastrcure. there are also questions that require us to run some terraform command. in the aws sections there is a aws-test-account.

the vscode editor has some nice plug-ins installed, which makes writing easier. we can use code completion to see resource types.

Using the coupon to access the kodeKloud labs.

#### Lab: HCL Basics

main.tf example

```hcl
resource "local_file" "games" {
  filename     = "/root/favorite-games"
  content  = "FIFA 21"
}
```

`terraform plan` - won't work without `terraform init` (which create a hidden _.terrafrom_ folder).

_sensitive_content_ - hides the content from being printed on the screen! this is for _local_file_ resource, not a general thing.

</details>

## Terraform Basics

<details>
<summary>
Playing around with Terraform
</summary>

### Using Terraform Providers

a deeper look at providers.

the `terraform init` command downloads and installs plug-ins for the providers specified in the terrafrom files. these can be plugins for cloud vendors, databases, or even the local file provider.

all plugins are hosted by hashicorp at [terraform registry](registry.terraform.io).

there are three tiers of providers:

1. official providers - owned and maintained by hashicorp. this includes the big cloud providers such as AWS.
2. verified providers - owned and maintained by third party entities which are verified by hashicorp, services such as as bigip or heroku are verified providers.
3. community providers - plugins with no formal relationship to hashicorp.

the `init` command shows the version of the plugin installed, this command is safe to run, as many times as required. running the commnad creates a hidden folder.

> - hashicorp/local: version = "~>2.0.0"

[Organization Namespace]/[Type]

there can also a hostname, the name of the register where the plugin is contained. by default it uses the hashicorp registry. the newest version is used by default. we can choose to lock down a specific version, if we wish to.

### Configuration Directory

so far we used a single file,

local.tf

```hcl
resource "local_file" "pet" {
    filename ="/root/pets.txt"
    content = "We love pets!"
}
```

we can create more configuration file
cat.tf

```hcl
resource "local_file" "cat" {
    filename ="/root/cat.txt"
    content = "my cat name is danny!"
}
```

we can also put several configuration blocks inside a single file, which is commonly called "main.tf".

```hcl
resource "local_file" "pet" {
    filename ="/root/pets.txt"
    content = "We love pets!"
}

resource "local_file" "cat" {
    filename ="/root/cat.txt"
    content = "my cat name is danny!"
}
```

other common files are "variables.tf", "outputs.tf","provider.tf".

#### Lab: Terraform Providers

we can see the providers in the hidden folder.

`terraform init`\
`terraform apply`

```hcl
resource "local_file" "xbox" {
  filename     = "/root/xbox.txt"
  content  = "Wouldn't mind an XBox either!"
}
```

### Multiple Providers

using multiple providers and resources.

from the "random" provider, we use the "pet" resource with the name "my-pet".

```hcl
resource "local_file" "pet" {
    filename ="/root/pets.txt"
    content = "We love pets!"
}

resource "random_pet" "my-pet {
    prefix = "Mrs"
    seperator= "."
    length = "1"
}
```

when we run the `terraform init` command, we will install the required addition plugin for the random_pet resource.
we we apply the change, the output for of the random pet resource is displayed on the screen.

#### Lab: Multiple Providers

```hcl
resource "local_file" "my-pet" {
	    content = "My pet is called finnegan!!"
	    filename = "/root/pet-name"
}


resource "random_pet" "other-pet" {
	      prefix = "Mr"
	      separator = "."
	      length = "1"
}
```

### Using Input Variables

```hcl
resource "local_file" "my-pet" {
	    content = "My pet is called finnegan!!"
	    filename = "/root/pet-name"
}


resource "random_pet" "other-pet" {
	      prefix = "Mr"
	      separator = "."
	      length = "1"
}
```

the arguments and the values are hardcoded. we want a way to provide them during execution.

we do this with a new file. _variables.tf_

```hcl
variable "filename" {
    default = "/root/pets/txt"
}
variable "content" {
    default = "We love pets!"
}
variable "prefix" {
    default = "Mrs"
}
variable "separator" {
    default = "."
}
variable "length" {
    default = "1"
}
```

just as always, there are blocks, where the block type is **variable**, then the name, and a default value.

to use the variables. we simply reference them in the defintion block with the **var** preceding them.

```hcl
resource "local_file" "my-pet" {
	    content = var.content
	    filename = var.filename
}


resource "random_pet" "other-pet" {
	      prefix = var.prefix
	      separator = var.separator
	      length = var.length
}
```

now we can update the variables file, rather than the resource files.

heres an example:

_main.tf_

```hcl
resource "aws_instance" "webserver"{
    ami = var.ami
    instance_type = var.instance_type
}
```

_variables.tf_

```hcl
variable "ami" {
    default = "ami-0edab43b6fa892279"
}
variable "instance_type" {
    default = "t2.micro"
}
```

### Understanding the Variable Block

the variable block has three parts

- default value
- type (optional)
- description (optional)

```hcl
variable "filename" {
    default = "/root/pets/txt"
    type = string
    description = "the path of local file"
}
```

| type   | example                | notes                                 |
| ------ | ---------------------- | ------------------------------------- |
| string | "/root/pets/txt"       |
| number | 1                      |
| bool   | true / false           |
| list   | ["cat","dog"]          | numbered, index zero                  |
| set    | ["cat","dog"]          | numbered, index zero, no duplications |
| map    | {pet1=cat pet2=dog}    | key-value pairs                       |
| tuple  | complex data structure | list, but not the same type of values |
| object | complex data structure |
| any    | default value          |

lets start using them

_variable.tf_

```hcl
variable "prefix" {
    default = ["Mr","Mrs","Sir"]
    type = list
}
variable "file-contents"{
    type= map
    default = {
        "statement1" = "We love pets!"
        "statement2" = "We love animals!"
    }
}
```

_main.tf_

```hcl
resource "random_pet" "my-pet" {
	      prefix = var.prefix[0]
}

resource "local_file" "my-pet" {
	    content = var.file-contents["statement2"]

}
```

we can also combine type constaints

```hcl
variable "prefix" {
    default = ["Mr","Mrs","Sir"]
    type = list(string)
}
```

for maps, they key is always string, but the value can be constrained. if we have duplications in the set, things will fail. when the default elements and the type don't match, `terraform plan` will fail.

objects allow us to define complex strcuteres;

```hcl
variable "bella" {
    type = object({
        name = string
        color = string
        age = number
        food = list(string)
        favorite_pet = bool
    })
    default = {
        name = "bella"
        color = "brown"
        age = 7
        food =["fish","chicken", "turkey"]
        favorite_pet = true
    }
}
```

tuple looks like a list, but it requires a fixed amount of elements with a defined type for each.

```hcl
variable "kitty" {
    type = tuple([string, number, bool])
    default = ["cat",7,true]
}
```

#### Lab: Variables

_main.tg_

```hcl
resource "local_file" "jedi" {
     filename = var.jedi["filename"]
     content = var.jedi["content"]
}
```

### Using Variables in Terraform

different ways of using the input variables.

we aren't required to have a default value for each variable. if we run the `apply` command without them, then we will prompted to enter them.\
a diffrent way of using them is to pass the values in the command line with the `-var` flag. alternatively, we can set them as part of the terrafrom environment by exporting them with the **TF*VAR*** prefix. then they will picked up by the apply command.

```sh
export TF_VAR_prefix="Mrs"
export TF_VAR_length="2"
terraform apply -var "filename=/root/pets.txt" -var "content=We Love Pets!"
```

another way to pass variables is with a specific file, with the _.tfvars_ or _.tfvars.json_ extension

```
filename = "/root/pets.txt"
content = "We love pets!"
prefix = "Mrs"
separator = "."
length = "2"
```

we then pass them with the `-var-file` flag.

```sh
terrafrom apply -var-file variables.tfvars
```

if we name the files as one the following options, it will be loaded without us needing to specify it in the command line.

- terraform.tfvars
- terraform.tfvars.json
- \*.auto.tfvars
- \*.auto.tfvars.json

to understand the way in which terraform decides which value to use, let's have an example:

_main.tf_

```hcl
resource local_file pet{
    filename = var.filename
}
```

_variables.tf_

```hcl
variable filename{
    type=string
    description= "file path"
    //no default
}
```

we have files that should load automatically:\
_terraform.tfvars_

```hcl
filename = "/root/pets.txt"
```

_variable.auto.tfvars_

```hcl
filename = "/root/pets.txt"
```

and we export a variable

```sh
export TF_VAR_filename="/root/cats.txt"
```

and we use the `-var` flag in the command line

```sh
terraform apply -var "filename=/root/best-pet.txt"
```

the order, from weakest to strongest:

0. (default variables)
1. environment variables (`export TF_VAR_`)
2. automatically loaded files (_\*.auto.tfvars_), by lexical order
3. command line flags `-var` and `-var-file` at the same strength

#### Lab: Using Variables in terraform

don't forget! we must first declare the variable in a variable block!

```hcl
variable filename{
    type="string"
}
```

### Resource Attributes

linking resource together. so far we used separate variables for each resource, but in most real world scenarios, resources are dependent on one another, we would want to use the data from one resource as the value for another.

in our example, we would like to use the random pet name inside the contents of the file

this can be done with **attributes**. if we look at the documentation for the random pet resource, we will see that it has one attribute, _id_ of type string. so lets use it.

we use the `${}` string interpolation for this, with the resource type, resource name and the attribute.

```hcl
resource "local_file" "my-pet" {
	    content = "My pet is called ${random_pet.other-pet.id}!"
	    filename = "/root/pet-name"
}


resource "random_pet" "other-pet" {
	      prefix = "Mr"
	      separator = "."
	      length = "1"
}
```

#### Lab: Resource Attributes

[time_static](https://registry.terraform.io/providers/hashicorp/time/latest/docs/resources/static)

```hcl
resource "time_static" "time_update"{

}

resource local_file time {
  filename="/root/time.txt"
  content="Time stamp of this file is ${time_static.time_update.id}"
}

```

### Resource Dependencies

different types of resource dependencies. output from one resource to another. the order is set by terraform based on dependencies, and the resources are destroyed in the reverse order. this dependency is **implicit**.

we can also use **explicit dependency** and force a specific order, this is done with the `depends_on` argument. this argument takes a list of dependencies. we should use it when one resource uses another, but not in a direct way.

```hcl
resource "local_file" "my-pet" {
	    content = "My pet is called Rex!"
	    filename = "/root/pet-name"
        depends_on = [
            random_pet.other-pet
        ]
}


resource "random_pet" "other-pet" {
	      prefix = "Mr"
	      separator = "."
	      length = "1"
}
```

#### Lab: Resource Dependencies

[tls_private_ket](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key)

this key lives in the terraform state.

```hcl
resource "tls_private_key" "pvtkey" {
    algorithm = "RSA"
    rsa_bits=4096
}

resource "local_file" "key_details" {
  filename="/root/key.txt"
  content = "${tls_private_key.pvtkey.private_key_pem}"
}
```

explicit dependency

```hcl
resource "local_file" "whale" {
    filename="/root/whale"
  depends_on=[
      local_file.krill
  ]
}

resource "local_file" "krill" {
    filename="/root/krill"

}
```

### Output Variables

terraform also suppots output variables.

```hcl
resource "local_file" "my-pet" {
	    content = "My pet is called ${random_pet.other-pet.id}!"
	    filename = "/root/pet-name"
}


resource "random_pet" "other-pet" {
	      prefix = "Mr"
	      separator = "."
	      length = "1"
}

output pet-name
{
    value = random_pet.other-pet.id
    description = "Record the value of pet ID"
}
```

when we apply the change,the value of the output will be printed to the screen. we will also be able to use `terraform output` to display all the output variable, or `terraform output pet-name` to show a specific variable.

#### Lab: Output Variables

```hcl
resource "random_uuid" "uid" {

}

resource "random_integer" "number" {
    min = 1
    max = 15
}
```

```sh
terraform output id2
terraform output order1
```

</details>

## Terraform State

<details>
<summary>
Terraform state - single source of truth
</summary>

### Introduction to Terraform State

terraform state - what happens under the hood when we run commands.

when we run `terraform init`, we download the plugins. the `terraform plan` command tried to update the state, and if there is no state, it knows that it should create the resources. the same thing happens when we run `terraform apply`. the internal state is checked compared to the requested state.

we can see this in the _terraform.tfstate_ file. this file is created by the apply comand. the file itself is a json file, it has every detail about the infrastructure, and it is the single source of truth. every `apply` command is checked against the state file and because of that, we know if there are changes to the resources or not.

### Purpose of State

the terraform state tracks the dependencies between the resources. therefore, it also controls the order of creating resources. this also allows it to destroy resources, and the correct order of doing so. having a local file allows us to avoid requesting the state from external objects each time.

state is refreshed when we `plan` a deployment, but we can suppress this behavior.
`terraform plan --refresh=false`

the state file is usually located in the end-user folder, but it is also possible to store it remotely so that every member of the team has the same state. this is called remote state, and will be covered later0

#### Lab: Terraform State

```sh
terraform show
terraform apply

```

### Terraform State Consideration

the state file contains sensitive information, ips, memory, OS, even the SSH key. for databases, the state might store the initial passwords. when it's stored locally, the state file is plain text.

the configuration files can be stored in version control, and the state file should be stored in a dedicated location. we shouldn't manually edit the file, but in some cases, we would modify it using `terraform state` commands.

</details>

## Working With Terraform

<details>
<summary>
Additional ways to work with terraform.
</summary>

### Terraform Commands

lets get aquatinted with some other commands

`terraform validate` - determine if the configuration file is valid, and will try to help us fix errors if the are any.

`terraform fmt` - format the configuration files

`terraform show` - displays the terrafrom state

`terraform providers` - will show us the providers used in our configuration files. we can use `terraform providers mirror <path>` to copy the plugins to a different folder.

`terraform refresh` - sync with the state at the external world, this is done automatically when we run `plan` and `apply` commands.

`terraform graph` - will show us dependencies between our resources, this can be run even before running `init`, the default format (_dot_) is confusing. but we can pass it to a graphing software.

```sh
apt update
apt install graphviz -y
terraform graph | dot -Tsvg > graph.svg
```

#### Lab: Terraform Commands

```sh
terraform validate
terraform plan
terraform apply
terraform fmt
terraform show
terraform providers
```

### Mutable vs Immutable Infrastructure

infrastructure can be mutable or immutable. when updating an immutable infrastructure, the resource must first be destroyed and the re-created.

in-place update, mutable infrastructure, like updating software.

configuration drift - when infrastrcutes (servers) which began as identical slowly become different over time across changes and updates.

terraform uses the replacement approach, by default, it first deletes an existing resource before spinning up a new one, but this can changed by using lifecycle rules.

### LifeCycle Rules

if we have a local_file resource and we change the file permssions, running `apply` will first remove the old file, but we might want to change this behavior. this is done with inner **lifecycle blocks**.

```hcl
resource "local_file" pet{
    filename = ".root/pets.txt"
    content = "We love pets"
    file_permission="0700"

    lifecycle{
        create_before_destroy = true
    }
}
```

if we don't want the resource to be destroyed at all, we can control that. this might be relevent for databases and so on.

```hcl
resource "local_file" pet{
    filename = ".root/pets.txt"
    content = "We love pets"
    file_permission="0700"

    lifecycle{
        prevent_destroy = true
    }
}
```

we can also decide to ignroe changes, maybe we want to allow changes to the tags, even if they aren't coming from terraform.

```hcl

resource "aws_instance" "webserver" {
    ami = "ami-0edab43b6fa892279"
    instance_type = "t2.micro"
    tags = {
        Name = "ProjectA-Webserver"
    }
    lifecycle {
        ignore_changes = [
            tags
        ]
    }
}
```

- create_before_destroy. true / false
- prevent_destroy. true / false
- ignore_changes. list / all

#### Lab: Lifecycle Rules

```sh
terraform init
terraform plan
terraform apply
terraform state show local_file.file
```

```hcl
resource "random_string" "string" {
    length = var.length
    keepers = {
        length = var.length
    }
    lifecycle{
        create_before_destroy=true
    }

}
```

**issue with creating files before destroying**\
a file is a file is a file. it's an actual unique resource, we don't have instances of it. we can't create a file with the same name before destroying the previous one, so our new file overwrites the old one, and is then destroyed!

### Datasources

terraform is the only way to provision infrastructure, there are other IoC tools, and the GUI console from the provider itself. terraform can also interact with those resources, even if it didn't create them.

this is done with **data** source blocks. in this example,we read a local file which we didn't create, and use it as a source for another local file resource.

```hcl
data "local_file" "dog"{
    filename = "/root/dog.txt"
}

resource "local_file" "pet"{
    filename = "/root/pets.txt"
    content = data.local_file.dog.content
}
```

data blocks are similar to resource blocks, the exposed attributes are different

| \              | Resource                                   | Data source                 |
| -------------- | ------------------------------------------ | --------------------------- |
| keyword        | _resource_                                 | _data_                      |
| usage          | **create, update, destroy** infrastructure | only **read** infrastrcuter |
| alternate name | Managed resources                          | Data resources              |

#### Lab: Datasources

```hcl
output "os-version" {
  value = data.local_file.os.content
}
data "local_file" "os" {
  filename = "/etc/os-release"
}
```

### Meta-Arguments

so far we used single resource, but we might want multiple instaces of the same resource.

in a shell script, a for loop would look like this.

```sh
#!/bin/bash

for i in {1..3}
    do
        touch /root/pet${i}
    done
```

in Terraform, we can achieve a similar result, by using a **meta-argument**.
we already used two meta-arguments:

- depends_on
- lifecycle

#### Count

a meta argument to create multiple instances:

```hcl
resource "local_file" "pet"{
    filename = var.filename
    count = 3
}
```

now the resource is a list of elements, but because this is a file, the file created three times, so it doesn't work.

but we can work around it by working with a list variable. now the created resource itself is a list.

```hcl
variable "filename" {
    default = [
        "/root/pets.txt",
        "/root/dogs.txt",
        "/root/cats.txt"
    ]
}

resource "local_file" "pet" {
    filename = var.filename[count.index]
    count = 3
}
```

how there will be three files, but now we have the count as a static variable. we can make use of the `length` function to get the correct amount of instances

```hcl
variable "filename" {
    default = [
        "/root/pets.txt",
        "/root/dogs.txt",
        "/root/cats.txt",
        "/root/cows.txt"
        "/root/ducks.txt"
    ]
}

resource "local_file" "pet" {
    filename = var.filename[count.index]
    count = length(var.filename)
}
```

there is another drawback: if we remove the first value from the list, then all the values after it will be modified. in our example, we replace two resources and delete the third, even though we actually just wish to delete one.

#### for-each

`for-each` is another meta-argument, however, it only works with maps (or sets), and not with lists.

```hcl
variable "filename" {
    type=set(string)
    default = [
        "/root/pets.txt",
        "/root/dogs.txt",
        "/root/cats.txt",
        "/root/cows.txt"
        "/root/ducks.txt"
    ]
}

resource "local_file" "pet" {
    filename = each.value
    for_each = var.filename
}
```

we can keep the variables as a list, but use the `toset` function. this might

```hcl
variable "filename" {
    type=list(string)
    default = [
        "/root/pets.txt",
        "/root/dogs.txt",
        "/root/cats.txt",
        "/root/cows.txt"
        "/root/ducks.txt"
    ]
}

resource "local_file" "pet" {
    filename = each.value
    for_each = toset(var.filename)
}
```

lets take a look using output variables.

```hcl
output "pets"{
    value = local_file.pet
}
```

`terraform output.pets`

now the resource is stored a map/set, rather than a list.

#### Lab: Count and for each

```hcl
variable "users" {
    type = list(string)
    default = [
        "/root/user10",
        "/root/user11",
        "/root/user12",
        "/root/user10"
    ]
}

variable "content" {
    default = "password: S3cr3tP@ssw0rd"
}

resource "local_file" "name" {
    filename = each.value
    sensitive_content = var.content

    for_each = toset(var.users)
}
```

### Version Constraints

provider use plug-ins, the `init` command takes by default the latest version. if we want to use a specific version, we need to specify it.

for each provider, the default and latest version is shown in the doumentation.

now we introduce the _terraform_ block, which can control which version is used.

in this example, we set the source and version of the plugin for terraform to download.

```hcl
terraform {
    required_providers{
        local ={
            source = "hashicorp/local"
            version = "1.4.0"
        }
    }
}
```

there are also version constraints,

```hcl
terraform {
    required_providers{
        local ={
            source = "hashicorp/local"
            version = "!= 2.0.0"
        }
    }
}
```

we can also use multiple contrains, such as `version = "> 1.2.0, <2.0.0, !=1.4.0"`, the "~>1.2" allows us to take incrmental versions, so we can download any "1.x" version, but not "2.0"

#### Lab: Version Constraints

```hcl
terraform {
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "1.2.2"
    }
  }
}
```

</details>

## Terraform with AWS

<details>
<summary>
Focusing on AWS Cloud Vendor.
</summary>

### Getting Started with AWS

AWS is the leading cloud vendor, with hundereds of services, both general and specific. aws has infrastructure in many regions across the world.

AWS is a first tier terraform plugin, so it's managed by hashicorp itself.

### Demo Setup an AWS Account

learning how to set an aws account.

[aws homepage](www.aws.amazon.com)

creating an account, payment information (even for the free tier). multi factor authentication.

- Compute
- Storage
- Database
- Security, Identity & Compliance

### Introduction to IAM

<details>
<summary>
IAM - Identity Access management
</summary>

the root account can do anything, so it's not adviced to use it. the best practice is to create other users with the appropriate premissions.

there are two types of access methods: access to the management console with a user name and password, and programatic permissions, which use access Key Id and Secret Access Key.

permissions are described in aws Policies.

there are some default policies which are managed by aws. the policy is defined in a json file.

this policy is the administrator policy, it can do anything.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*"
    }
  ]
}
```

some common managed AWS polices:

| Job function           | Policy                |
| ---------------------- | --------------------- |
| Administrator          | AdministratorAccess   |
| Billing                | Billing               |
| Database Administrator | DatabaseAdministrator |
| Network Administrator  | NetworkAdministrator  |
| View Only User         | ViewOnlyAccess        |

there are also **IAM Groups**, which can help us manage policies across a group of users, instead of managing them individually.

Services also have permissions, we need to configure access between aws resources. this is done with **IAM Roles**.

IAM roles can also be used to provide access to user from other aws accounts, to software and other user management services.

here is another policy, which allows to create and delete tags from any ec2 resource.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["ec2:CreateTags", "ec2:DeleteTags"],
      "Resource": "*"
    }
  ]
}
```

#### Demo IAM

introduction the IAM with the console: groups, users, roles, policies.

the IAM region is always Global. in the dashboard:

**Create User**\
adding user, choosing access types (programatic access, aws management console acceses), passwords. skipping permissions and tags for now. at the final page we can download the access key. if we look at the user policies, we can see that it got the _IAMUserChangePassword_ policy. and we can attach other permissions for it.

**Create Group**\
a group of permissions, what can members of the group do, instead of manually setting each user permissions.

**Polices**\
the awsManagedPolicies are default, sensible policies that are available for use without configuration. we can also create a policy for some resources and for specific actions on those resources. Policies are described as json documents

**Roles**

- one aws service to another
- users from another aws account
- web Identity
- other user management systems.

lets create a role, we choose the trusted service, and give it a policy.

#### Programmatic Access

interacting with the aws services using the aws CLI (command line interface)

```sh
aws --version
aws s3api create-bucket -bucket my-bucket -region us-east-1
aws configure
#type the access key, secret access key, default region, default output format
cat ./aws/config/config
```

the base syntax is:\
`aws <command> <subcommand> [option and parameters]`

the top level command is usually the service we want to use.

```sh
aws iam create-user --user-name lucy
```

we can also get help for specific commands

```sh
aws help
aws iam help
aws iam create-user-help
```

#### Lab: AWS CLI and IAM

the lab uses an aws mocking service, so there is always a `--endpoint http://aws:4566` parameter added.

```sh
aws --version
aws iam help
aws iam list-users --endpoint http://aws:4566 #option 1
aws --endpoint http://aws:4566 iam list-users  #option2
aws iam create-user help
aws --endpoint http://aws:4566 iam create-user --user-name mary
aws configure list
cat ~/.aws/config
cat ~/.aws/credentials
aws iam attach-user-policy help
aws --endpoint http://aws:4566 iam attach-user-policy --user-name mary --policy-arn arn:aws:iam::aws:policy/AdministratorAccess
aws iam create-group help
aws --endpoint http://aws:4566 iam create-group --group-name project-sapphire-developers
aws iam add-user-to-group help
aws iam list-groups-for-user --user-name jack

aws --endpoint http://aws:4566 iam list-attached-user-policies --user-name jack
aws --endpoint http://aws:4566 iam list-attached-group-policies --group-name project-sapphire-developers
aws --endpoint http://aws:4566 iam attach-group-policy --group-name project-sapphire-developers --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess
```

#### AWS IAM with Terraform

the third way to work with IAM is through [terraform aws provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

```hcl
resource "aws_iam_user" "admin-user"{
    name = "lucy"
    tags = {
        Description = "Technical Team Leader"
    }
}
```

the provider is aws, the resource type is iam*user, the resource name is "admin-user", and we provide the \_name* required argument, and the optional tags map. we could also provide a _path_ argument, a _permissions_boundary_ arn and an _force_destroy_ option.

now when we run terrafrom init, we download the plugins as usual, but when we run `terraform plan`, we will get an error because we don't have valid permissions.

we need to decide on a region, and to either pass the access key and secret.

```hcl
provider "aws" {
    region = "us-west-2"
    access_key=<>
    secret_key=<>
}
```

now running `terraform plan` doesn't fail and we can see the execution plan.

we could also get the credentials from a shared credentials file or from the environment variables **AWS_ACCESS_KEY_ID** and **AWS_SECRET_ACCESS_KEY**

to configure the profile we can run `aws configure` with a profile.

```sh
export AWS_ACCESS_KEY_ID=<>
export AWS_SECRET_ACCESS_KEY=<>
```

#### IAM Policies with Terraform

now we want to give our user permissions.

| argument    | required                       | notes               |
| ----------- | ------------------------------ | ------------------- |
| policy      | required                       | a json object       |
| description | optional                       | forces new resource |
| name        | optional                       | forces new resource |
| name_prefix | optional - clashes with "name" | forces new resource |
| path        | optional                       |
| tags        | optional                       |

the problem is how to pass the policy document. we can use something called <kbd>heredoc Syntax</kbd>.

```hcl
resource aws_iam_policy "adminUser" {
    name = "AdminUsers"
    policy= <<EOF
    {
        "Version":"2012-10-17",
        "Statement":[
            {
                "Effect": "Allow",
                "Action": "*",
                "Resource": "*"
            }
        ]
    }
    EOF
}
```

now that we have a policy,we can attach it to our user. we can see how this resource block uses that reference syntax.

```hcl
resource "aws_iam_user_policy_attachment" "lucy-admin-access"{
    user = aws_iam_user.admin-user.name
    policy_arn = aws_iam_policy.adminUser.arn
}
```

rather than write the policy in the terraform file, we can grab it from an existing file in the folder.

admin-policy.json

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*"
    }
  ]
}
```

and we get it in the resource block by using the `file` function.

```hcl
resource aws_iam_policy "adminUser" {
    name = "AdminUsers"
    policy = file("admin-policy.json")
}
```

#### Lab: IAM with Terraform

the lab uses an aws mocking service, so there is always a `--endpoint http://aws:4566` parameter added.

```hcl
resource "aws_iam_user" "users" {
    name = "mary"
}
```

```sh
terraform init
terraform validate
# region is not set
```

```hcl
provider "aws" {
  region= "ca-central-1"

    # skip_credentials_validation = true
    # skip_requesting_account_id  = true
    #
    #  endpoints {
    #    iam = "http://aws:4566"
    #  }
}
```

```hcl
variable "project-sapphire-users" {
     type = list(string)
     default = [ "mary", "jack", "jill", "mack", "buzz", "mater"]
}

resource "aws_iam_user" "users" {
    name = var.project-sapphire-users[count.index]
    count = length(var.project-sapphire-users)

}
```

</details>

### Introduction to AWS S3

<details>
<summary>
S3 - Simple Storage Service
</summary>

Object based (flat file), with no hard limits. not suitable for operating systems or databases. data is stored inside bucket. a bucket is like a directory, but there can also be nested folders.

the management console provides a simple interface to create buckets, the name of the bucket must be unique (across the world), must be DNS compliant (lowercase, doesn't end with dash). the bucket will get a unique address and is (theoretically) globally accessable.

each object in S3 has data and metadata, the data includes the key (file name) and the value (actual data), the metadata contains information about the file. like other aws services, access to buckets is controlled through permissions, and also _access control lists_. permissions can be defined in bucket level or even at a file level.

this is an example to a bucket policy.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["s3:GetObject"],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::all-pets/*",
      "Principal": {
        "AWS": ["arn:aws:iam:::123456123457:user/Lucy"]
      }
    }
  ]
}
```

we can even write a bucket policy to grant access to users from other aws accounts.

#### S3 with Terraform

if we don't provide a name, it will be randomly generated.

```hcl
resource "aws_s3_bucket" "finance"{
    bucket = "finance-21092020"
    tags = {
        Description= "Finance and Payroll"
    }
}
```

now we wish to add a file to that bucket. we must provide the bucket onto which we want to upload the file, the key (the path in the bucket), and the file itself.

in this example we use the reference syntax.

```hcl
resource "aws_s3_bucket_object" "finance-2020"{
    content = "/root/finance/finance-2020.doc"
    key = "finance-2020.doc"
    bucket = aws_s3_bucket_finance.id
}
```

now we assume there is an existing users group, which wasn't created by Terraform, but we wish to give those users access to the bucket. we will use a data block. afterwards, we will also create a bucket policy resource.

```hcl
data "aws_iam_group" "finance-data"{
    group_name = "finance-analysts"

}

resource "aws_s3_bucket_policy" "finance-policy"{
    bucket = aws_s3_bucket_finance.id
    policy= << EOF
    {
        "Version":"2012-10-17",
        "Statement": [
            {
                "Action":"*"
                "Effect":"Allow",
                "Resource": "arn:aws:s3:::${aws_s3_bucket.finance.id}/*",
                "Principal":{
                    "AWS":[
                        "${data.aws_iam_group.finance-data.arn}"
                    ]
                }
            }
        ]
    }
    EOF
}
```

#### Lab: S3

[aws_s3_bucket](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket)

playing with buckets, getting an error about incorrect DNS format, trying to use _acl = "public-read-write"_ and failing.

</details>

### Introduction to DynamoDB

<details>
<summary>
NoSQL database.
</summary>

highly scalable, fully managed, no server for the the user to manage. low latency data access. data is replicated across region, so it's highly available.

DynamoDB is a key-value and document database, unlike a relational database. each entry in the collection is an item, Dynamo db uses a primary key to uniquely identify elements. we aren't required to fill in attributes which aren't the primary key, they can be duplicated, empty or null.

#### Demo Dynamodb

in the management console. we go to the dynamoDB service and create a table, we give it a name, and choose the primary key and it's type. we can manually add item by clicking <kbd>Create item</kbd>. we can now start filling in values. we can search for items using the console and filters.

#### DynamoDB with Terraform

lets define a dynamoDB resource block. we provide the table name and the hash*key to definf the primary key, we must define an \_attribute* for the primary key, but we can also provide attributes for other fields.

```hcl
resource "aws_dynamodb_table" "cars"{
    name = "cars"
    hash_key = "VIN"
    billing_mode = "PAY_PER_REQUEST"
    attribute {
        name = "VIN"
        type ="S"
    }
}
```

to add items, we use another resource type, and the _heradoc_ syntax, but we need to define each element as a json with the type

```hcl
resource "aws_dynamodb_table_item" "car-items"{
    table_name = aws_dynamodb_table.cars.name
    hash_key = aws_dynamodb_table.cars.hash_key
    item = <<EOF
    {
        "Manufacturer": {"S": "Toyota"},
        "Make": {"S": "Corrolla"},
        "Year": {"N": 2004},
        "VIN": {"S": "4Y1SL65848Z411439"}
    }
    EOF
}
```

this is just an example of adding an item to the table, this isn't how it should be done in large scale database.

#### Lab: DynamoDB

resource "aws_dynamodb_table" "project_sapphire_inventory" {
name = "inventory"
billing_mode = "PAY_PER_REQUEST"
hash_key = "AssetID"

attribute {
name = "AssetID"
type = "N"
}
attribute {
name = "AssetName"
type = "S"
}
attribute {
name = "age"
type = "N"
}
attribute {
name = "Hardware"
type = "B"
}
global_secondary_index {
name = "AssetName"
hash_key = "AssetName"
projection_type = "ALL"

}
global_secondary_index {
name = "age"
hash_key = "age"
projection_type = "ALL"

}
global_secondary_index {
name = "Hardware"
hash_key = "Hardware"
projection_type = "ALL"

}
}

resource "aws_dynamodb_table_item" "upload" {
table_name = aws_dynamodb_table.project_sapphire_inventory.name
hash_key = aws_dynamodb_table.project_sapphire_inventory.hash_key
item = <<EOF
{
"AssetID": {"N": "1"},
"AssetName": {"S": "printer"},
"age": {"N": "5"},
"Hardware": {"B": "true" }
}

EOF
}

</details>

</details>

## Remote State

<details>
<summary>
Storing the State file in a remote location.
</summary>

### What is Remote State and State Locking?

we already saw the TF uses the state file to map resources. this file is created when we `terraform apply` for the first time.

> - Mapping configuration to real world
> - Tracking metadata
> - Performance
> - Colabaration

we also discussed the we shouldn't store this file in a version control tool as it contains passwords and other sensitive information.

imagine that we have a terraform state file that we manually store on github, each time we changed the configuration, we upload the file again. this creates merge conflicts, especially if there are many users.\
when using terraform locally, we can't change the state file once ome operation started. this is called **state locking**. we can see tis in action by running `terraform apply` from two terminals.

to overcome these problems, we can use a **remote state backend**. now the state file is moved to a shared storage. this will also enable state locking for the state file, so there won't be conflicts and always maintains the updated configuration without having to manually upload it.

### Remote Backends with S3

using S3 bucket and a dynamo db table as a remote State backend.

| Object         | Value                              |
| -------------- | ---------------------------------- |
| Bucket         | kodekloud-terraform-state-bucket01 |
| Key            | finance/terraform.tfstate          |
| Region         | us-west-1                          |
| DynamoDB Table | State-locking                      |

```hcl
resource "local_file" "pet"{
    filename = "/root/pets.txt"
    content = "We love pets!"
}
```

if we run `terraform apply`, then a local state file will be created. if we want you use a remote state file, we need to configure the terraform block. the dynamodb_table is used to control state locking.

this block should be in the **terraform.tf** file.

```hcl
terraform{
    backend "s3"{
        bucket = "kodekloud-terraform-state-bucket01"
        key = "finance/terraform.tfstate"
        region = "us-west-1"
        dynamodb_table = "state-locking"
    }
}
```

before being able to use the remote backend, we should run the `terrafrom init` command again. we can then copy our local file into the S3 bucket. we should now delete the local file `rm -rf terraform.tfstate`.

#### Lab: Remote State

```sh
terraform apply
terraform show
```

### Terraform State Commands

the `terraform state` commands. the state is stored in a json format, which we should directly edit. instead, we can use some cli commands:

`terraform state <sub commands> [options] [args]`

sub comands:

- `terrform state list` - list all the resources. we can pass a resource name to get a subset of results.
- `terraform state show` - prints the attributes of an resource
- `terraform state mv [options] [SOURCE] [DESTINATION]` - either rename a resource or move it to another state file.
  ```sh
  terraform state mv aws_dynamodb_table.state-locking aws_dynamodb_table.state-locking-db
  ```
  (we should then rename the resource in the configuration file)
- `terraform state pull [options] [SOURCE] [DESTINATION]` - get the remote state
  ```sh
  terraform state pull | jq '.resource[] | select (.name =="state-locing-dbb").instances[].attributes.hash_key'`
  ```
- `terraform state rm <ADDRESS>` - remove an address and not longer manage it, it isn't destroyed, simple stops being managed.

#### Lab: Terraform State Commands

```sh
terraform state list
terraform state show local_file.classics
terraform state show local_file.top10
terrafrom state rm local_file.hall_of_fame
terraform state mv random_pet.super_pet_1 random_pet.ultra_pet
```

</details>

## Terraform Provisioners

<details>
<summary>
Understanding EC2 (Elastic Compute Cloud) instances and running commands on them.
</summary>

### Introduction to AWS EC2 (optional)

EC2 (Elastic Compute Cloud) instances in AWS, virtual machines in the cloud, based on a OS (unix or windows).

> AMI: Amazon Machine Image - templates for virtual machine configurations.

the templates contain the the OS and additional software, each AMI has an id, which is different per region. there also different configurations for machine cpu and hardware, which are identified as _Instance Types_.

general purpose instance type, compute optimized, memory optimized.

the general purpose are divided into categories:

here are some configuration for the T2 type, but there are also T3, M5, etc...
InstanceType | vCPU | Memory (GB)
---|---|---
t2.nano|1|0.5
t2.micro|1|1
t2.small|1|2
t2.medium|2|4
t2.large|2|8
t2.xlarge|4|16
t2.x2large|8|32

> EBS - Elastic Block Storage - the storage attached to the EC2

| Name | Type | Description                                                            |
| ---- | ---- | ---------------------------------------------------------------------- |
| io1  | SSD  | for business-critical Apps                                             |
| io2  | SSD  | for latency sensitive transactional workloads                          |
| gp2  | SSD  | general purpose                                                        |
| st1  | HDD  | low cost HDD frequently accessed, throughput-intensive workloads       |
| sc1  | HDD  | lowest cost HDD volume designed for less frequently accessed workloads |

we can also pass User Data to the Ec2, so that commands will be run as soon as the machine starts.

```sh
#!/bin/bash
sudo apt update
sudo apt install nginx
systemctl enable nginx
systemctl start nginx
```

for windows machines we can pass batch files or power shell. access to EC2 machine is done with a SSH key-pair.

### Demo: Deploying an EC2 Instance (optional)

in the management console, select <kbd>EC2</kbd> from the services (under the **compute** group).

- <kbd>Launch Instance</kbd>
- choose an _ami_ and an _instance type_, such as ubuntu and t2.micro.
- configure the instance, using the default vpc for the network, using the default value of the subnet.
- in the _user data_ block (advanced), we pass in the shell script
  ```sh
  #!/bin/bash
  sudo apt update
  sudo apt install nginx
  systemctl enable nginx
  systemctl start nginx
  ```
- in the storage section, we can use the default value.
- we can add tags to the instances
- configure security group, and allow it SSH access (inbound rule). when the source is "0.0.0.0/0", it means that all access is allowed.
- we also create a key-pair and download them.
- <kbd>View instaces</kbd> to see the details of the virtual machine.

to access it from the terminal, we copy the public ip address (3.94.9.249). we might run into a problem and have to change the private key file permissions.

```sh
chmod 400 ~/Downloads/web.pem
ssh -i ~/Downloads/web.pem ubuntu@3.97.9.249
```

if we did it correctly, our shell is now configured to the machine

```sh
systemctl status nginx
```

### AWS EC2 with Terraform

at "main.tf" file,
the instance resource supports:

- ami (required)
- instance_type (required)
- tags (optional)
- user_date (optional)
- key_name (optional)
- vpc_security_groups_ids (optional)

```hcl
resource "aws_instance" "webserver"{
    ami = "ami-0edab43bdfa892279"
    instance_type= "t2.micro"
    tags = {
        Name = "webserver"
        Description = "An Nginx Web Server on Ubuntu"
    }
    user_date = <<EOF
        #!/bin/bash
        sudo apt update
        sudo apt install nginx -y
        systemctl enable nginx
        systemctl start nginx
        EOF
}
```

at the "provider.tf" file

```hcl
provider "aws"{
    region = "ws-west-1"
}
```

but now we need the ip of the machine, and keys to access it via ssh. we use another resource for that.

```hcl
resource "aws_key_pair" "web"{
    public_key=file("/root/.ssh/web/pub")
}
```

and now we update the instance configuration to tell it to make use of the the key resource.

```hcl
resource "aws_instance" "webserver"{
    ami = "ami-0edab43bdfa892279"
    instance_type= "t2.micro"
    # tags, shell

    key_name = aws_key_pair.web.id
}
```

the next issue is the networking, in the demo, we used the default vpc and used a new security group with inbound access rules.

```hcl
resource "aws_security_group" "ssh-access"{
    name = "ssh-access"
    description = "Allow SSH access from the internet"
    ingress = {
        from_port = 22
        to_port = 22
        protocol ="tcp"
        cider_blocks = ["0.0.0.0/0"]
    }
}
```

and we connect our instance to this security group.

```hcl
resource "aws_instance" "webserver"{
    ami = "ami-0edab43bdfa892279"
    instance_type= "t2.micro"
    # tags, shell

    key_name = aws_key_pair.web.id
    vpc_security_groups_ids = [aws_security_group.ssh-access.id]
}
```

lets also have an output variable to display the public ip address

```hcl
output public-ip {
    value = aws_instance.webserver.public_ip
}
```

and to test everything

```sh
terraform apply
terraform output public-ip #get ip
ssh -i /root/.ssh/web ubuntu@3.96.203.171 #the ip
systemctl status nginx #from inside the instance
```

### Terraform Provisioners

- remote-exec
- local-exec

Provisioners allows us to run scripts or commands on resources. we can specify the script in a provisioner block.

this requires a network connectivity

```hcl
resource "aws_instance" "webserver"{
    ami = "ami-0edab43bdfa892279"
    instance_type= "t2.micro"
    # tags

    key_name = aws_key_pair.web.id
    vpc_security_groups_ids = [aws_security_group.ssh-access.id]

    provisioner "remote-exec" {
        inline = ["sudo apt update",
        "sudo apt install nginx -y",
        "systemctl enable nginx",
        "systemctl start nginx"
        ]
    }
    provisioner "local-exec" {
        command = "echo {aws_instance.webserver.public_ip} >> /tmp/ips.txt"
    }
    connection {
        type = "ssh"
        host = self.public_ip
        user = "ubuntu"
        private_key = file ("/root/.ssh/web")
    }
}
```

the provisioners are run once the resource is created. we can also specify provisioners to run when the resource is destroyed by specifing the _when_ argument.

```hcl
resource "aws_instance" "webserver"{

    # ...

provisioner "local-exec" {
        command = "echo {aws_instance.webserver.public_ip} Created! > /tmp/ips.txt"
    }}
    provisioner "local-exec" {
        when = destroy
        command = "echo {aws_instance.webserver.public_ip} Destroyed! > /tmp/ips.txt"
    }
}
```

if the provisioner command fails, then the `terraform apply` command fails. but we can control this behavior with the _on_failure_ argument in the provisioner block.

```hcl
provisioner "local-exec" {
    on_failure = continue
    command = "echo {aws_instance.webserver.public_ip} Created! > /temp/ips.txt"
    }}
```

the best practice is to avoid using provisioners if possible, and to use the options from the provider instead

| Provider       | Resource                 | Option        |
| -------------- | ------------------------ | ------------- |
| AWS            | aws_instance             | user_data     |
| Azure          | azurrerm_virtual_machine | custom_data   |
| GCP            | google_compute_instance  | meta_data     |
| Vmware vSphere | vsphere_virtual_machine  | user_data.txt |

### Provisioner Behavior

as before the default behavior of provisioners is to run when the resource is created and to fail the apply command if there was as problem.

```hcl
provisioner "local-exec" {
    when = destroy
    on_failure = continue
    command = "echo Instance ${aws_instance.webserver.public_ip} Destroyed! > /tmp/instance_state.txt"
}
```

if we provisioner command fails, then the resource is considered to be **tainted**.

#### Lab: AWS EC2 and Provisioners

[aws_key_pair](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/key_pair), [aws_eip](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip)

creating an instance

```hcl
resource "aws_instance" "cerberus" {
  ami = var.ami
  instance_type = var.instance_type
}
variable "region" {
    default = "eu-west-2"
    type = string
}
variable "instance_type" {
    default = "m5.large"
    type = string
}
variable "ami" {
  default = "ami-06178cf087598769c"
}
```

```sh
terraform init
terraform validate
terraform plan
terraform apply
terraform show
```

creating a key_pair

```hcl
resource "aws_key_pair" "cerberus-key" {
    key_name = "cerberus"
    public_key = file("/root/terraform-projects/project-cerberus/.ssh/cerberus.pub")
}
```

using the key_pair on the instance

```hcl
resource "aws_instance" "cerberus" {
  ami = var.ami
  instance_type = var.instance_type
  key_name = aws_key_pair.cerberus-key.id
}
```

adding the scripts

```hcl
resource "aws_instance" "cerberus" {
  ami = var.ami
  instance_type = var.instance_type
  key_name = aws_key_pair.cerberus-key.id
  user_data = file("install-nginx.sh")
}
```

`terraform state show aws_instance.cerberus`

elastic ip resource for a consistent ip address

```hcl
resource "aws_eip" "eip" {
    vpc = true
    instance = aws_instance.cerberus.id
    provisioner "local-exec" {
        command = "echo ${self.public_dns} > /root/cerberus_public_dns.txt"
    }
}
```

`terraform state show aws_eip.eip`

### Considerations with Provisioners

Provisioners aren't best practice for TF, they are powerful tools, but carry some issues.

```hcl
resource "aws_instance" "webserver"{
    ami = "ami-0edab43b6fa892279"
    instance_type = "t2.micro"
    tags = {
        Name = "webserver"
        Description = "An NGINX WebServer on Ubuntu"
    }
    provisioner "remote-exec" {
        inline = ["echo $(hostname -i) >> tmp/ips.txt"]
    }
}
```

**No Provisioner Information in Plan**: we can run anything on the resource, so we don't have any way to parse it in the `terraform plan` stage.

**Network Connectivity and Authentication**: some provisioners require a connection block, which isn't alway possible. it's better to avoid provisioners which are native to the resource, like _user_date_

it's better to keep the provisioners work to the minimum, it's better to use an image (ami) which has what we want already installed. we can create custom ami with tools like _Packer_

</details>

## Terraform Import, Tainting Resources and Debugging

<details>
<summary>
Additional Terraform commands.
</summary>

### Terraform Taint

sometimes a resource creation can fail.

here, we try to write to an invalid path.

```hcl
resource "aws_instance" "webserver-3"{
    ami = "ami-0edab43b6fa892279"
    instance_type = "t2.micro"
    key_name = "ws"

    provisioner "local-exec" {
        command = "echo ${aws_instance.webserver-3.public_ip} > /temp/pub_ip.txt"
    }
}
```

`terraform apply`

now the resource is marked as tainted, and when we run `terraform plan`, we will see the resource marked as tainted. and at the next time we run `terraform apply`, it will be re-created.

if, for some reason, our resource was changed manually (and not by terraform), we need to recreate the resource. we could remove and re-apply the configuration block to force the creation, but it's easier to mark the resource as tainted, and then it'll happen in the next `terraform apply` run.

`terraform taint aws_instance.webserver`

if we want the resource to remain as it is, we can remove the this mark.

`terraform untaint aws_instance.webserver`

### Debugging

sometimes, looking at the output from the apply command isn't enough to understand the problem. in those cases, we might wish to increase the verbosity of the logs, this is done by changing the log_level.

`export TF_LOG=TRACE`

- INFO
- WARNING
- ERROR
- DEBUG
- TRACE - the most verbose

now when we run commands, we will see much more logs text, we can also store them in a file (if we want to send a bug report)

`export TF_LOG_PATH=/tmp/terraform.log`

to disable logging, we can unset the environment variable.

`unset TF_LOG_PATH`

#### Lab: Taint and Debugging

```sh
export TF_LOG=ERROR
export TF_LOG_PATH=/tmp/ProjectA.log
```

### Terraform Import

importing existing infrastructure into our configuration, we don't always have the luxury of doing everything with terrafrom. but we can bring them into our control.

the `data` block allows us to read from existing resources,

```hcl
data "aws_instance" "newserver"{
    instance_id = "i-026e13be10d5326f7"
}
output newserver{
    value = data.aws_instace.newserver.public_ip
}
```

we can't update or delete this resource. if we want it to be under our command, we need to import it.

`terrafrom import <resource_type>.<resource_name> <attribute>`

when we run this command, we are requested to create a corresponding resource file in our configuration. the command only updates the state file, it doesn't update the config block on it own.

```hcl
resource "aws_instance" "webserver-2"{

}
```

now when we run the import command\
`terraform import aws_instance.webserver-2 i-026e13be10d536f7`

we will see a msg that the import has succeeded, and we could fill in the missing stuff in the resource block. we can inspect the state file to find them, and run `terraform plan` to make sure our configuration is the same as what exists.

#### Lab: Terraform Import

```hcl
resource "aws_instance" "ruby" {
  ami           = var.ami
  instance_type = var.instance_type
  for_each      = var.name
  key_name      = var.key_name
  tags = {
    Name = each.value
  }
}
output "instances" {
  value = aws_instance.ruby
}

```

key was created

```sh
# example
aws ec2 create-key-pair --endpoint http://aws:4566 --key-name jade --query 'KeyMaterial' --output text > /root/terraform-projects/project-jade/jade.pem.
# example
```

describe instance

```sh
aws ec2 describe-instances --endpoint http://aws:4566
#or
aws ec2 describe-instances --endpoint http://aws:4566 --filters "Name=image-id,Values=ami-082b3eca746b12a89" |jq -r '.Reservations[].Instances[].InstanceId'
```

import

```sh
terraform import aws_instance.jade-mw i-f6f4f794b1577f607
terraform state show aws_instance.jade-mw
```

fill in the blanks in the configuration

```hcl
resource aws_instance jade-mw{
    ami = "ami-082b3eca746b12a89"
    instance_type = "t2.large"
    key_name = "jade"
    tags= {
        Name = "jade-mw"
    }
}
```

</details>

## Terraform Modules

<details>
<summary>
Using existing Terraform configuration folders as module resources.
</summary>

### What are Modules?

since we've started with aws resources, our configuration files have gotten quite large:

- aws_instance (ec2)
- aws_key_pair
- aws_iam_policy
- aws_s3_bucket
- aws_dynamodb_table

but if we have shared defintions (like several ec2 machine), wouldn't it be easier to share the configuration between them?

we can split a file into several smaller files.

Any configuration directory is already a terraform module. so far we run commands directly on a module, which made it into the root module. but if we want, we can reference this module from another one, by using a module block.

we simple provide the path to the other module, which is now the "child module" for us to use.

(in this example, we give a relative path)

```hcl
module "dev-webserver"{
    source = "../aws-instance"
}
```

### Creating and Using a Module

assume that we have a architecture that needs to be deployed in several countries.

- ec2 with custom ami
- dynamodb table
- s3 bucket
- default vpc - no special end points

so we start by designing this a module in folder "modules/payroll-app" with several files:

- app_server.tf
- dynamodb_table.tf
- s3_bucket.tf
- variables.tf

those files will look something like this:

```hcl
resource "aws_instance" "app_server"{
    ami = var.ami
    instance_type = "t2.medium"
    tags = {
        Name = "${var.app_region}-app-server"
    }
    depends_on =[
        aws_dynamodb_table.payroll_db,
        asw_s3_bucket.payroll_data
    ]
}

resource "aws_s3_bucket" "payroll_data" {
    bucket = "${var.app_region}-${var.bucket}
}

resource "aws_dynamodb_table" "payroll_db" {
    name = "user_data"
    billing_mode = "PAY_PER_REQUEST"
    hash_key = "EmployeeID"

    attribute{
        name = "EmployeeID"
        type = "N"
    }
}

variable "app_region"{
    description = "where this software is deployed"
    type = string
}

variable "bucket"
{
    default = "flexit-payroll-alpha-22001c"
    type = string
}

variable "ami" {
    description = "custom ami image"
    type = string
}

```

now, we want to deploy this configuration, first to the US region. so we create a different folder, with two configuration files in it:

- main.tf
- provider.tf

we now configure the module block in the root module configurations to connect to the child-module.

```hcl
module "us_payroll" {
    source = "../modules/payroll-app"
    app_region ="us-east-1"
    ami = "ami-24e140119877avm"
}
```

running `terraform init` will now tell us we are initilazing a module as well. all of there resources now have the \<module>.\<module_name> "module.us_payroll" prefix.

we can now create a different folder for another region:

```hcl
module "uk_payroll" {
    source = "../modules/payroll-app"
    app_region ="eu-west-2"
    ami = "ami-24e140119877avm"
}
provider "aws" {
    region = "eu-west-2"
}
```

the state file is in the root module, not in the child modules.

using modules is simpler, lowers the risk, and makes the re-usable.

### Using Modules from the Registry

so far, we use local modules, they are located in a folder in the local machine. but it's also possible to share modules, this is done by getting them from the terraform registry.

modules are grouped by the resource they use. there are official modules (managed by hashicorp) and community modules.

the module page in the registry should contain instructions on how to use the module

[example: aws security group](https://registry.terraform.io/modules/terraform-aws-modules/security-group/aws/latest)

```hcl
module "security-group" {
    source  = "terraform-aws-modules/       security-group/aws"
    version = "4.8.0"

    # insert the 3 required variables here
    # name
    # security_group_id
    # vpc_id
    # ingress_cider_blocks
}
```

there are also some submodules.

we can run the `terraform get` command to download modules and plugins, even after `terraform init` has been run.

#### Lab: Terraform Modules

[aws_iam_user](https://registry.terraform.io/modules/terraform-aws-modules/iam/aws/latest/submodules/iam-user)

```hcl
module "iam_iam-user" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-user"
  version = "3.4.0"
  # insert the 1 required variable here
  name ="max"
  # optional
  create_iam_access_key	= false
  create_iam_user_login_profile = false
}
```

</details>

## Terraform Functions and Conditional Expressions

<details>
<summary>
More functions, Terraform workspaces.
</summary>

### More Terraform Functions

terraform has an interactive console.

`terraform console`

this loads the terraform state, so we can run functions and see how they would interpolate.

```
file("/root/terraform-projects.main.tf")
length(var.region)
toset(var.region)
```

there are many function in terraform, so we will look only at some of them:

- numeric functions
- string functions
- collections functions
- type conversion functions

`max(-1,2,-10,200,-250)` -> 200\
`min(-1,2,-10,200,-250)` -> -250

```hcl
variable "num" {
    type = set(number)
    default = [250,1,11,5]
    description = "A set of numbers"
}
```

to get use the "num" variable as an argument, we need to use the expansion operator (`...`, three dots):

`max(var.num...)` -> 250

`ceil`, `floor` rounding function.

```hcl
variable "ami" {
    type= string
    description = "A string containing the ami ids"
    default = "ami-xyz,AMI-ABC,ami-efg"
}
```

we can use the `split` function to split a string into a list. we provide the seperator

`split(",","ami-xyz,AMI-ABC,ami-efg")` -> ["ami-xyz","AMI-ABC","ami-efg"]\
`split(",",var.ami)`

`lower("aBC")`->"abc"\
`upper("aBC")`->"ABC"\
`title("abc")`->"Abc"

taking a substring\
`substr(var.ami,0,7)` -> "ami-xyz"

creating a string from a list\
`join(",",["ab","cd","ef"])` -> "ab,cd,ef"

`length` - number of elements in a collection\
`index(var.ami, "AMI-ABC")` -> 1\
`element(var.ami,2)` -> "ami-efg"\
`contains(var.ami,"AMI-ABC")`->true
`contains(var.ami,"AMI-AbC")`->false

```hcl
variable "ami" {
    type= map
    default  {
        "us-east-1"="ami-xyz",
        "ca-central-1"="ami-efg",
        "ap-south-1"="ami-ABC"
    }
    description = "A map of AMI ids for specific regions"
}
```

`keys(var.ami)` -> ["us-east-1","ca-central-1","ap-south-1"]\
`values(var.ami)` ->["ami-xyz","ami-efg","ami-ABC"]\
`lookup(var.ami, "ca-central-1")` -> "ami-efg"\
`lookup(var.ami, "ca-central-2", "ami-pqr)` -> "ami-pqr"

### Conditional Expressions

common operator and conditional operators. we can use the terraform console to play with operators.

- arithmetical operators
- comparative operators `>`,`>=`,`==`,`<=`,`<`
- logical operators: `&&`,`||`,`!`

`var.a < var.b`

lets have some conditional statements: we want to use the value given, but no smaller than eight.

as a bash script

```sh
if [ $length -lt 8]
    then
        length=8
        echo $length
    else
        echo $length
    fi
```

in terraform

```hcl
variable password-length{
    type=number
    description = "the length of the password"
}
resource "random_password" "password-generator" {
    length =max(var.password-length,8)
}

output password{
    value = random_password.password-generator.result
}
```

we can also use a three way coitional expression
`condition ? true_val : false_value`

```hcl
resource "random_password" "password-generator" {
    length = var.password-length < 8? 8 : var.password-length
}
```

`terraform apply -var=password-length=5`

#### Lab: Functions and Conditional Expressions

```sh
terraform console
$floor(10.9)
$title("user-generated password file")

```

creating iam users by spliting a string

```hcl
resource "aws_iam_user" "users"
{
    name = split(":",var.cloud_users)[count.index]
    count = length(split(":",var.cloud_users))
}
```

back to the console

```sh
terraform console
$element(aws_iam_user.cloud,6)
$index(var.sf,"oni")
```

uploading elements to a bucket

```hcl
resource "aws_s3_bucket_object" "upload_sonice_media" {
    bucket= aws_s3_bucket.sonic_media.id
    key = substr(each.value,length("/media/"),length(each.value)-length("/media"))
    content = file(each.value)
    for_each = var.media
}
```

conditional statement

```hcl
resource "aws_instance" "mario_servers" {
  ami = var.ami
  tags ={
      "Name"=var.name
  }
  instance_type = var.name == "tiny"? var.small : var.large
}
```

### Terraform Workspaces (OSS)

The Terraform state is the mapping between the configuration and the provisioned resources. so far we had one terraform state file (either stored locally or remotely).

workspace are another way to re-use code. it's two state files that use the same configuration

`terraform workspace new <Workspacename>`\
`terraform workspace list`

we move everything into the variable resource

```hcl
variable region {
    default = "ca-central-1"
}
variable instance_type {
    default = "t2.micro"
}
variable ami {
    type=map(string)
    default = {
        "ProjectA" = "ami-0edab43b6fa892279"
        "ProjectB" = "ami-0c2f25c1f66a1ff4d"
    }
}
```

and use them in the configuration file

```hcl
resource "aws_instance" "project"{
    ami = lookup(var.ami,terraform.workspace)
    instance_type = var.instance_type
    tags = {
        Name = terraform.workspace
    }
}
```

lets play in the console

```sh
terraform console
$terraform.workspace
$lookup(var.ami,terraform.workspace)
```

and if we want to work on the other workspace, we can create another workspace and run the same command.

`terrafrom workspace select ProjectA`

when we use workspaces, the statefiles are stored in a special folder **terraform.tfstate.d** which has sub directories per workspace.

#### Lab: terraform Workspaces

```sh
terraform workspace new uk-payroll
terraform workspace new us-payroll
terraform workspace new india-payroll
terraform workspace list
terraform workspace select us-payroll
```

```hcl
module "payroll_app"{
  source = "/root/terraform-projects/modules/payroll-app"
    app_region = lookup(var.region,terraform.workspace)
    ami =  lookup(var.ami,terraform.workspace)
}
```

</details>

## Takeaways

<details>
<summary>
Things to remember
</summary>

AWS human users have **Users**, aws services have **Roles**, and they both use **Policies**. **ARN** - Amazon Resource Name.

### Terraform functions

- `file("file-path")` - read file contents.
- Numeric Functions
  - `max`
  - `min`
  - `ceil`
  - `floor`
- String Functions
  - `split(delimiter, string)`
  - `join(delimeter,list of strings)`
  - `upper`
  - `lower`
  - `title`
  - `substr`
- Collection Functions
  - `length(collection)`
  - `index(list, value)`
  - `element(list, index)`
  - `contains(collection, value)`
  - `toset(list)` - turn a list into a set
  - `keys(map)` - get keys
  - `values(map)` - get values
  - `lookup(map,value, <default>)`

### Cli Commands

- `terraform version`
- `terraform init`
- `terraform plan`. `--refresh=false`
- `terraform apply`. `-var`, `-var-file`
- `terraform show`. `-json`
- `terraform destroy`
- `terraform output`
- `terraform validate`
- `terraform fmt`
- `terraform providers`
  - `terraform providers mirror <path>`
- `terraform refresh`
- `terraform graph`
- `terraform taint`
- `terraform untaint`
- `terraform import`
- `terraform get`
- `terraform console`
- `terrafrom workspace`
  - `terraform workspace show`
  - `terraform workspace list`
  - `terraform workspace select`
  - `terraform workspace new`
  - `terraform workspace delete`

### Common File Structure

| File Name           | Purpose                                                    |
| ------------------- | ---------------------------------------------------------- |
| main.tf             | Main configuration files containing resource definitions   |
| variables.tf        | variables decelerations                                    |
| outputs.tf          | Outputs from resources                                     |
| provider.tf         | Providers defintions                                       |
| variables.tfvars    | environment variables                                      |
| terraform.tfstate   | state, single source of truth                              |
| terraform.tf        | terraform block, for plugin configuration and remote state |
| terraform.tfstate.d | folder that holds state files for workpaces                |

### Block Types

| block type | purpose                                                                                                |
| ---------- | ------------------------------------------------------------------------------------------------------ |
| resource   | provision a resource                                                                                   |
| variable   | define variables to use in `var.$`                                                                     |
| output     | displaying on screen, or to pass it forwad to other shell commands. `terraform output <variable_name>` |
| data       | using resources that weren't created by Terraform.                                                     |
| terraform  | controling versions and provider source                                                                |
| module     | reference and use a different tf module (folder or registry)                                           |

### Environment variables

| variable                | usage                 | possible values                    |
| ----------------------- | --------------------- | ---------------------------------- |
| TF*VAR*\<variable name> | provide variable name | any                                |
| TF_LOG_PATH             | where to store logs   | location of file                   |
| TF_LOG                  | log verbosity         | INFO, ERROR, WARNING, DEBUG, TRACE |

</details>

## Resources samples

<details>
<summary>
Samples of Resource blocks
</summary>

### Terraform

```hcl
terraform {
  backend "s3" {
    key = "terraform.tfstate"
    region = "us-east-1"
    bucket = "remote-state"
    endpoint = "http://172.16.238.105:9000"
    force_path_style = true


    skip_credentials_validation = true

    skip_metadata_api_check = true
    skip_region_validation = true
  }
}

```

### Local

```hcl
resource "local_file" "my-pet" {
	    content = "My pet is called ${random_pet.other-pet.id}!"
        #sensitive_content
	    filename = "/root/pet-name"
        file_permission = "0700"

}
```

### Time

```hcl
resource "time_static" "time_update"{

}
```

### Random

> All the resources for the random provider can be recreated by using a map type argument called **keepers**. A change in the value will force the resource to be recreated.

```hcl
resource "random_pet" "pet" {
    prefix = "Mr"
    separator = "."
    length = "2"
}

resource "random_uuid" "uid" {

}

resource "random_integer" "number" {
    min = 1
    max = 15
}

resource "random_string" "string" {
    length = var.length
    keepers = {
        length = var.length
    }
}
```

### AWS

provisioning resources:

```hcl

provider "aws" {
    region = "us-west-2"
    access_key=<>
    secret_key=<>
}

resource "aws_iam_user" "admin-user"{
    name = "lucy"
    tags = {
        Description = "Technical Team Leader"
    }
}

resource "aws_iam_user_policy_attachment" "lucy-admin-access"{
    user = aws_iam_user.admin-user.name
    policy_arn = aws_iam_policy.adminUser.arn
}

resource aws_iam_policy "adminUser" {
    name = "AdminUsers"
    policy = file("admin-policy.json")
}

resource "aws_instance" "dev-server" {
    instance_type = "t2.micro"
    ami         = "ami-02cff456777cd"
}

resource "aws_s3_bucket" "finance"{
    bucket = "finance-21092020"
    tags = {
        Description= "Finance and Payroll"
    }
}

resource "aws_s3_bucket_object" "finance-2020"{
    content = "/root/finance/finance-2020.doc"
    key = "finance-2020.doc"
    bucket = aws_s3_bucket_finance.id
}
resource "aws_s3_bucket_policy" "finance-policy"{
    bucket = aws_s3_bucket_finance.id
    policy = file("finance-policy.json")
}

resource "aws_security_group" "ssh-access"{
    name = "ssh-access"
    description = "Allow SSH access from the internet"
    ingress = {
        from_port = 22
        to_port = 22
        protocol ="tcp"
        cider_blocks = ["0.0.0.0/0"]
    }
}

```

data sources:

```hcl
data "aws_ebs_volume" "gp2_volume" {
  most_recent = true

  filter {
    name   = "volume-type"
    values = ["gp2"]
  }
}

data "aws_s3_bucket" "selected" {
  bucket_name = "bucket.test.com"
}

data "aws_iam_group" "finance-data"{
    group_name = "finance-analysts"

}
```

### TLS

```hcl
resource "tls_private_key" "private_key" {
  algorithm   = "RSA"
  rsa_bits  = 4096
}
```

### Google

```hcl
resource "google_compute_instance" "special" {
  name         = "aone"
  machine_type = "e2-micro"
  zone         = "us-west1-c"

}
```

</details>
