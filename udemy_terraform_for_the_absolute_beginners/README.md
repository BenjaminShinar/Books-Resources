<!--
// cSpell:ignore HashiCorp KodeKloud FIFA tfvars tfstate falshpoint Tsvg Flexit toset aone
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
`terraform plan` - won't work without `terraform init` (which create a hidden *.terrafrom* folder). 

*sensitive_content* - hides the content from being printed on the screen! this is for *local_file* resource, not a general thing.
</details>


## Terraform Basics

<details>
<summary>
Playing around with Terraform
</summary>

### Using Terraform Providers

a deeper look at providers.

the `terraform init` command downloads and installs plug-ins for the providers specified in the terrafrom files. these  can be plugins for cloud vendors, databases, or even the local file provider. 

all plugins are hosted by hashicorp at [terraform registry](registry.terraform.io).

there are three tiers of providers:
1. official providers - owned and maintained by hashicorp. this includes the big cloud providers such as AWS.
2. verified providers - owned and maintained by third party entities which are verified by hashicorp, services such as as bigip or heroku are verified providers.
3. community providers - plugins with no formal relationship to hashicorp.

the `init` command shows the version of the plugin installed, this command is safe to run, as many times as required. running the commnad creates a hidden folder.

> * hashicorp/local: version = "~>2.0.0"

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

we do this with a new file. *variables.tf*

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

*main.tf*
```hcl
resource "aws_instance" "webserver"{
    ami = var.ami
    instance_type = var.instance_type
}
```
*variables.tf*
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

*variable.tf*
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
*main.tf*
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

*main.tg*
```hcl
resource "local_file" "jedi" {
     filename = var.jedi["filename"]
     content = var.jedi["content"]
}
```

### Using Variables in Terraform

different ways of using the input variables.

we aren't required to have a default value for each variable. if we run the `apply` command without them, then we will prompted to enter them.\
a diffrent way of using them is to pass the values in the command line with the `-var` flag. alternatively, we can set them as part of the terrafrom environment by exporting them with the **TF_VAR_** prefix. then they will picked up by the apply command. 

```sh
export TF_VAR_prefix="Mrs"
export TF_VAR_length="2"
terraform apply -var "filename=/root/pets.txt" -var "content=We Love Pets!"
```

another way to pass variables is with a specific file, with the *.tfvars* or *.tfvars.json*  extension

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
- *.auto.tfvars
- *.auto.tfvars.json

to understand the way in which terraform decides which value to use, let's have an example:

*main.tf*
```hcl
resource local_file pet{
    filename = var.filename
}
```
*variables.tf*
```hcl
variable filename{
    type=string
    description= "file path"
    //no default
}
```
we have files that should load automatically:\
*terraform.tfvars*
```hcl
filename = "/root/pets.txt"
```

*variable.auto.tfvars*
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
2. automatically loaded files (*\*.auto.tfvars*), by lexical order
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

this can be done with **attributes**. if we look at the documentation for the random pet resource, we will see that it has one attribute, *id* of type string. so lets use it.


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

we can see this in the *terraform.tfstate* file. this file is created by the apply comand. the file itself is a json file, it has every detail about the infrastructure, and it is the single source of truth. every `apply` command is checked against the state file and because of that, we know if there are changes to the resources or not.

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

`terraform graph` - will show us dependencies between our resources, this can be run even before running `init`, the default format (*dot*) is confusing. but we can pass it to a graphing software.

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

* create_before_destroy. true / false
* prevent_destroy. true / false
* ignore_changes. list / all


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
| keyword        | *resource*                                 | *data*                      |
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

now we introduce the *terraform* block, which can control which version is used.

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
            "Action": [
                "ec2:CreateTags",
                "ec2:DeleteTags"
            ],
            "Resource": "*"
        }
    ]
}
```

#### Demo IAM

introduction the IAM with the console: groups, users, roles, policies.

the IAM region is always Global. in the dashboard:

**Create User**\
adding user, choosing access types (programatic access, aws management console acceses), passwords. skipping permissions and tags for now. at the final page we can download the access key. if we look at the user policies, we can see that it got the *IAMUserChangePassword* policy. and we can attach other permissions for it.

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
the provider is aws, the resource type is iam_user, the resource name is "admin-user", and we provide the *name* required argument, and the optional tags map. we could also provide a *path* argument, a *permissions_boundary* arn and an *force_destroy* option.

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

argument| required | notes
---|---|---
policy| required|a json object
description|optional | forces new resource
name|optional| forces new resource
name_prefix|optional - clashes with "name"|forces new resource
path|optional
tags| optional

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
    "Version":"2012-10-17",
    "Statement":[
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

each object in S3 has data and metadata, the data includes the key (file name) and the value (actual data), the metadata contains information about the file. like other aws services, access to buckets is controlled through permissions, and also *access control lists*. permissions can be defined in bucket level or even at a file level.

this is an example to a bucket policy.
```json
{
    "Version":"2012-10-17",
    "Statement": [
        {
            "Action":[
                "s3:GetObject"
            ],
            "Effect":"Allow",
            "Resource": "arn:aws:s3:::all-pets/*",
            "Principal":{
                "AWS":[
                    "arn:aws:iam:::123456123457:user/Lucy"
                ]
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

playing with buckets, getting an error about incorrect DNS format, trying to use *acl = "public-read-write"*  and failing.

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

lets define a dynamoDB resource block. we provide the table name and the hash_key to definf the primary key, we must define an *attribute* for the primary key, but we can also provide attributes for other fields.

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

to add items, we use another resource type, and the *heradoc* syntax, but we need to define each element as a json with the type 
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
  name           = "inventory"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "AssetID"

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
    name             = "AssetName"
    hash_key         = "AssetName"
    projection_type    = "ALL"
    
  }
  global_secondary_index {
    name             = "age"
    hash_key         = "age"
    projection_type    = "ALL"
    
  }
  global_secondary_index {
    name             = "Hardware"
    hash_key         = "Hardware"
    projection_type    = "ALL"
    
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

## Terraform Provisioners

## Terraform Import, Tainting Resources and Debugging

## Terraform Modules

## Terraform Functions and Conditional Expressions


## Takeaways

<details>
<summary>
Things to remember
</summary>

AWS human users have **Users**, aws services have **Roles**, and they both use **Policies**. **ARN** - Amazon Resource Name.

`file("file-path")` - read file function.



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

### Common File Structure

| File Name         | Purpose                                                  |
| ----------------- | -------------------------------------------------------- |
| main.tf           | Main configuration files containing resource definitions |
| variables.tf      | variables decelerations                                  |
| outputs.tf        | Outputs from resources                                   |
| provider.tf       | Providers defintions                                     |
| variables.tfvars  | environment variables                                    |
| terraform.tfstate | state, single source of truth                            |


### Block Types

| block type | purpose                                                                                                |
| ---------- | ------------------------------------------------------------------------------------------------------ |
| resource   | provision a resource
| variable   | define variables to use in `var.$`                                                                     |
| output     | displaying on screen, or to pass it forwad to other shell commands. `terraform output <variable_name>` |
| data       | using resources that weren't created by Terraform.                                                     |
| terraform  | controling versions and provider source                                                                |

</details>

## Resources samples

<details>
<summary>
Samples of Resource blocks
</summary>


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