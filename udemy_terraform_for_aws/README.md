<!--
// cSpell:ignore HashiCorp choco tfstate
 -->

# Terraform for AWS - Beginers To Expert

udemy course [Terraform for AWS - Beginers To Expert](https://www.udemy.com/course/terraform-fast-track/) by *TJ Addams*. [Warp-9 channel](https://www.youtube.com/channel/UCqKggOZpJKvZSPxb9hrN37A). [Github repository](https://github.com/addamstj/Terraform-012)

[TfSwitch](https://warrensbox.github.io/terraform-switcher/).

[Terraform Registry](https://registry.terraform.io/)

## Terraform Setup
<details>
<summary>
Getting things set up.
</summary>

getting the required stuff.

- code editor and plugins.
- folder setup.
- installing terraform.
- aws setup.

installing terraform
- linux `wget`, `unzip <> -d /usr/local/bin/`
- windows:
  - instal chocolate with powershell
  - with chocolate `choco install terraform`
- mac: `brew install terrafrom`
  - we can also download **TfSwitch** to have more than one version of terraform.

changin the default terminal on vscode.

click <kbd>terminal</kbd>, select <kbd>default shell</kbd> and choose **cmd** if there are problems.

we will need an AWS Account.
- setting an IAM user account with programmatic access and administrator permissions. don't forget to save the access key id and the secret access key.

we should never put the access key and secret access key inside the terraform provider block. never.

```hcl
provider "aws" {
    region = "eu-west-2"
    access_key = "" #no!
    secret_key = "" #no!
}
```

instead we can use environment variables 
- windows: `set`
- unix `export`

there is also a product by hashicorp called **vault**.

but we can use the aws command line and simply run `aws configure` to store the keys locally.


</details>

## Terraform 101
<details>
<summary>
First steps with Terraform.
</summary>

Terraform is an Infrastructure as Code solution by hashicorp. we can integrate it with many providers, aws, google, azure, etc...

Infrastructure as code means managing resources by using tools and configuration files, and not by hand.

lets start with the first Terraform file, "main.tf"

```hcl
provider "aws" {
  region = "eu-west-2"
}

resource "aws_vpc" "my_vpc"{
  cidr_block="10.0.0.0/16"
}
```

### Terraform Basic Workflow Commands: `init`, `plan`, `apply`, `destroy`

now that we have the first resource,we can deploy it to the cloud.

the first step is running `terraform` init. this will download the plugins and create a "project" for terraform to use.

the `terraform plan` command checks the current state against the required state. this is how terraform know what it should do.

to actually create the resources, we run `terraform apply`, which asks if we are sure that we want to continue, and it will create and modify the resources.

to delete resources, we run the `terraform destroy` command, if we confirm it, the resource will be deleted.

### VPC setup

in the aws console, we might need to create a vpc.

<kbd>launch vpc wizard</kbd>, give it a name, the terraform application requires us having one.

### State

*terraform.tfstate*. The heart of terraform, the 'source of truth'. 

when we run the apply command, the state file is updated. we should never manually change it. if we need to modify the state, we can use some terrafrom commands.

### Variables: input and output

variables are declared in a *variable block*. we can re-use them.

```hcl
variable "vpc_name"{
  type = string
  default ="my_vpc"
  description = "the name of the vpc"
}
```

types:

| type       | usage                                             |
| ---------- | ------------------------------------------------- |
| string     | text                                              |
| number     | numeric                                           |
| boolean    | true\false                                        |
| list(type) | list of the same type                             |
| map(type)  | key-value pairs. string key and value of another type. |
| set(type)  | list without duplications                         |
tuple([type1,...])| multiple data types
object({name1=type1,...}) | named arguments of diffrent types

variables are used by referencing them with the `var` suffix.

```hcl

variable "vpcname"{
  type = string
  default = "my_vpc"
}

variable "my_list"{
  type=list(string)
  default =["a","b","c"]

}

variable "my_map"{
  type = map(string)
  default = {
    Key1 = "value1"
    Key2 = "value2"
  }
}

resource "aws_vpc" "my_vpc"{
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = var.vpcname
    Name2 = var.my_list[1]
    Name3 = var.my_map["Key2]
  }
}

variable "my_tuple"{
  type=tuple([string, number, string])
  default = ["cat",1,"dog"]
}

variable "my_object"{
  type = object({name = string, port = list(number)})
  default = {
    name = "ab"
    port =[22,80]
  }
}
```
variables can be set from the user, if we have a variable without a default value, terraform will require us to pass the value when we run `terraform plan` or `terraform apply`.

output variables are a way to expose values. we declare a **output** block.

```hcl
output "my_output"{
  value = aws_vpc.my_vpc.id
}
```

we can also use string interpolation

```hcl
output "my_output"{
  value = $"the id is {aws_vpc.my_vpc.id}"
}
```


this will show us the value when we run the terraform commands. 

### Challenge

> 1. Create a folder called "challenge1"
> 2. Create a VPC resource named "TerraformVPC"
> 3. the CIDR range is "192.168.0.0/24"

[documentation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc)

```hcl
provider "aws" {
  region = "eu-west-2"
}
resource "aws_vpc" "challenge_vpc"{
  cidr_block = "192.168.0.0/24"
  tags = {
    Name = "TerraformVPC"
  }
}
```

and in the shell

```sh
terraform init
terraform plan
terraform apply
terraform destroy
```


</details>

## EC2
<details>
<summary>
EC2 resource.
</summary>

> - EC2 - Elastic Cloud Compute
> - AMI - Amazon Machine Instance
> - EIP - Elastic IP

### EC2 - Creating the Instance

EC2 instances are region specific, unlike IAM users. The required fields are "ami" and "instance_type"

```hcl
provier "aws" {
  region = "eu-west-2"
}

resource "aws_instance" "my_ec2"{
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  tags = {
    Name = "my_ec2"
  }
}
```

### Elastic IP (EIP)

if we want a static (non changing) ip for the instance, we can request a public, static ip. we can create an elastic IP resource and attach it to our instance.


```hcl
provier "aws" {
  region = "eu-west-2"
}

resource "aws_instance" "my_ec2"{
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  tags = {
    Name = "my_ec2"
  }
}

resource "aws_eip" "elastic_eip"{
  instance = aws_instance.my_ec2.id
}

output "EIP" {
  value = aws_eip.elastic_eip.public_ip
}
```

### Security Groups

a security group is similar to a **statefull** firewall, it allows traffic in and out, and every inbound traffic is allowed to get out (no need for separeate inbound and outbound rules).


the *from_port* and *to_port* describe a range of ports. the "0.0.0.0/0" security groups allows all.

```hcl
resource "aws_security_group" "webtraffic" {
  name = "Allow HTTPS"

  ingress {
    from_port = 443
    to_port = 443
    protocol = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 443
    to_port = 443
    protocol = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "my_ec2"{
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  tags = {
    Name = "my_ec2"
  }
  security_groups = [aws_security_group.webtraffic.name]
}
```

### Dynamic Blocks

a dynamic block is a terraform concepts that provides us the ability to turn lists or collections into terraform code.

we defind the block as `dynamic`, and then we set the iterator, we set what it iterators over, and how each repeated block is created.


lets see an example

```hcl
variable "ingressRules" {
  type = list(number)
  default = [80,443]
}

variable "egressRules" {
  type = list(number)
  default = [80,443,25,3306,53,8080]
}

resource "aws_security_group" "webtraffic" {
  name = "Allow HTTPS"

  dynamic ingress {
    iterator = port
    for_each = var.ingressRules
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  egress {
    iterator = port
    for_each = var.egressRules
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }
}
```

when we run the terraform commands, the contents of the "ingress" and "egress" resources are now lists where each element was created based on the variables.

### EC2 Challenge

> 1. Create a DB server and out the private IP.
> 2. Create a web server and ensure it has a fixed public IP.
> 3. Create a security group for the web server opening ports 80 and 443.
> 4. Run the provided script on the EC2 machine.

```hcl
provider "aws"{
  region ="us-east-1"
}

resource "aws_instance" "db_server"{
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  tags = {
    Name = "db_server"
  }
}

output "db_private_ip" {
  value = aws_instance.db_server.private_ip
}

variable "ports" {
  type = list(number)
  default = [80,443]
}

resource "aws_security_group" "web_traffic" {

  name = "Allow Web Traffic"

  dynamic ingress {
    iterator = port
    for_each = var.ports
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  dynamic egress {
    iterator = port
    for_each = var.ports
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }
}

resource "aws_instance" "web_server" {
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  user_date = file("script.sh")
  security_groups = [aws_security_group.web_traffic.name]
  tags = {
    Name = "web_server"
  }
}

resource "aws_eip" "elastic_eip" {
  instance = aws_instance.web_server.id
}

output "web_server_public_ip" {
  value = aws_eip.elastic_eip.public_ip
}
```

to see the outputs

```sh
terraform output web_server_public_ip
```

</details>

## Modules
<details>
<summary>
Reusable Terraform configurations.
</summary>

### Modules Deep-Dive
Reusable terraform "folders" that together compose a functionality. we can have terraform resources grouped together in order to achieve a goal.

we create new folder and put a new *.tf* file in it. this will be our module.

```hcl
variable "ec2name" {
  type = string
}

resource "aws_instance" "ec2"{
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  tags = {
    Name = var.ec2name
  }
}
```
in a different folder, we put the *main.tf* file and reference the other module from it. we can pass the paramters to the module.

```hcl
module "ec2module" {
  source = "./ec2"
  ec2name = "name from module"
}
```

### Handling Outputs

we want to get the values from the modules that we use. 

output variables in the module which we create
```hcl
output "instance_id" {
  value = aws_instance.ec2.id
}
```

to get the output variable, we simple reference it with the `module` prefix.

```hcl
output "module_output" {
  value = module.ec2module.instance_id
}
```


### Remote Modules

terraform registry is a remote repository for all kinds of modules, it's like a resource block, but with a complete configuration. modules can have input parameters and output parameters.

there are modules for many providers and configurations.

### Module Challenge

> Modularise Challenge 2 (EC2 challenge)
>
> (push as many things into modules.) 

*module file.tf*
```hcl
variable "instance_name" {
  type = string
}

variable "ingress_ports" {
  type = list(number)
}

variable "egress_ports" {
  type = list(number)
}

variable "user_data" {
  type = string
  default =""
}

variable "sg_name" {
  type = string
  default ="security_group"
}

resource "aws_security_group" "sg" {
  name = var.sg_name

  dynamic ingress {
    iterator = port
    for_each = var.ingress_ports
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  dynamic egress {
    iterator = port
    for_each = var.egress_ports
    content {
      from_port = port.value
      to_port = port.value
      protocol = "TCP"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }
}

resource "aws_instance" "ec2"{
  ami = "ami-032598fcc7e9d1c7a"
  instance_type = "t2.micro"
  security_groups = [aws_security_group.sg.name]
  tags = {
    Name = var.instance_name
  }
  userData = var.user_data
}

resource "aws_eip" "elastic_eip" {
  instance = aws_instance.ec2.id
}

output "ec2_private_ip" {
  value = aws_instance.ec2.private_ip
}

output "ec2_public_ip" {
  value = aws_eip.elastic_eip.public_ip
}

```

and the using code
```hcl
module "web_server" {
  source = "./ec2"
  instance_name = "web_server"
  ingress_ports = [80,443]
  egress_ports = [80,443]
  user_date = file("script.sh")
  sg_name = "web_traffic"
}

output "web_server_public_ip" {
value = module.web_server.public_ip
}

module "db_server" {
  source = "./ec2"
  instance_name = "db_server"
  ingress_ports = [80,443]
  egress_ports = [80,443]
  sg_name = "web_traffic"
}

output "db_server_private_ip" {
value = module.db_server.private_ip
}
```

#### Challenge Solution

1. he pushed the two aws instances into folders.
2. and then he pushed the eip and the security group into modules inside those folders.
3. not as much as reusable modules taking things out.

> - web_server:
>   - uses sg
>   - uses eip
> - db
>   - uses nothing

**there is an issue with the file address.**
the path to a module is static and depends on the current location (the path is relative to the file itself). but the path inside the `file("web_server_script.sh")` is relative to the caller context, which is the *main.tf* in challenge3 folder. we need to be aware of this.

one way to get around it is to use absolute paths, such as `user_data = file("${path.module}/server-script.sh")`

</details>

## IAM Masterclass
<details>
<summary>
Advanced IAM.
</summary>

creating iam policies and using them in terraform.

### How to Create and use IAM Policies
using the web management console.

Service <kbd>IAM</kbd>, selecting <kbd>policies</kbd>, then <kbd>Create Policy</kbd> and select the *glacier* policy, choose "List" and "write" actions, now we can add more services and policies.

we do all this to get a valid json example for the policies.

### IAM Users and Working WIth Policies
now we move to terraform

we create a user, a policy document using the HERDOC syntax (EOF,mind the indentations), and a policy attachment.

```hcl
provider "aws"{
  region = "us-east-1"
}

resource "aws_iam_user" "my_user"{
  name = "TJ"

}

resource "aws_iam_policy" "customPolicy"{
  name = "myCustomPolicy"
#indentation matters!
  policy = << EOF
{
  #the json content
}
EOF
}

resource "aws_iam_policy_attachment" "policy_bind"{
  name = "attachment"
  users = [aws_iam_user.my_user.name]
  #groups
  #roles
  policy_arn = aws_iam_policy.customPolicy.arn
}
```

we could then see the new user and the policy in the aws web console.

</details>

## RDS
<details>
<summary>
Relation Database
</summary>

### Planning Phase

we want to start a relational database instance.

in the web management console, services <kbd>RDS</kbd>, then <kbd>Create Database</kbd>, we choose an engine (mysql, aurora, maria, etc...). we can fill in the fields, set the storage. we can use storage types, choose the connectivity (the VPC) and which instance type we want to use.

### Creating an RDS Instance in Terraform


there is a difference between the identifier, which is the aws FQDN name, and the "name", which is the name of the database instance.

```hcl
resource aws_db_instance "my_RDS"{
  name= "myDB"
  identifier = "my-first-rds"
  
  instance_class = "db.t2.micro"
  engine = "mariadb"
  engine_version = "10.2.21"
  #should use secret
  username = "bob"
  password = <>

  port = 3306
  allocated_storage = 20
  #important for terraform
  skip_final_snapshot = true
}
```

</details>

## Advanced Terraform
<details>
<summary>
More Terraform concepts.
</summary>

### Remote Backend

the state file is the source of truth for the configuration. but if we work in a team, we can have a shared state file which is stored remotely, and this gives us version control and helps us avoid bugs.

we can host the remote state file on an S3 bucket.

to connect to a remote backend, we define the special terraform block.

we provide the bucket name, and the object key, which is the path to the state file inside the bucket.
it's better to provide a name. we also need to provide the region and the access key and the secret key.


```hcl
terraform {
  backend "s3" {
    bucket = "bucket_name"
    key = "terraform/tfstate.tfstate"
    region = "eu-west-2"
    access_key =""
    secret_key =""
  }
}
```

now the state file is stored in the s3 bucket.

### Dependencies

when we need resources to be created in a specified folder. there are direct implicit dependencies, like how values can depend on variables, but if we have dependencies which aren't directly expressed by the configuration (such as script that connects to a resource) , or those which are costly to create, we can set explicit dependencies order with the `depends_on` argument.

```hcl
resource "aws_instance" "db"{
  instance_type = "t2.micro"
  ami = "ami-032598fcc7e9d1c7a"
}

resource "aws_instance" "web_server"{
  instance_type = "t2.micro"
  ami = "ami-032598fcc7e9d1c7a"
  user_date = file("script_to_connect_db.sh")
  depends_on = [aws_instance.db]
}
```

### Count

`count` allows us to create multiple resources with slight changes between them.

this will create 3 instance of the ec2 instances. 
```hcl
resource "aws_instance" "ec2"{
  instance_type = "t2.micro"
  ami = "ami-032598fcc7e9d1c7a"
  count = 3
}
```
we can also use variables and get the element with `count.index`. and we can use the `*` to get all the elements from a list/enumerable.

```hcl

variable  "server_names"{
  type = list(string)
}

resource "aws_instance" "ec2"{
  instance_type = "t2.micro"
  ami = "ami-032598fcc7e9d1c7a"
  count = length(var.server_names)
  tags={
    Name = var.server_names[count.index]
  }
}

output "PrivateIP"{
  value = [aws_instance.ec2.*.private_ip]
}
```



### Multiple Variable Files

```hcl
variable  "number_servers"{
  type = number
}

resource "aws_instance" "ec2"{
  instance_type = "t2.micro"
  ami = "ami-032598fcc7e9d1c7a"
  count = var.number_servers
}
```

we can have different variable files and pass them with `-var-file=<>` when we run the terraform commands. `-var 

### Terraform Import

we can also import already existing resources and take control over them. this is called importing.

```hcl
resource "aws_vpc" "my_vpc"
{
  cidr_block = "10.0.0.16/17"
}
```

we create a resource on the aws web console, and then we grab the ip.

we run the import command with the ip
`terraform import aws_vpc.my_vpc" <id>`

which will connect the terraform block with existing resource on the cloud. after that, the resource belongs to the terraform configuration.

### Data Sources

data resources are a way to interact with resources which are already existing, but not under terraform control. just like the ones which we imported, but without taking over the control.

there are many ways to connect to existing resource, like a filter for some attribute or a direct connection with the id.

```hcl
date "aws_instance" "db_search"{
  filter {
    name = "tag:Name"
    values = ["DB Server"]
  }
}

output "db_servers"{
  value = data.aws_instance.db_search.avalability_zone
}
```

</details>


## Takeaways
<details>
<summary>
Things worth remembering
</summary>

- `terraform version`
- `terraform init`
- `terraform plan`
- `terraform apply`
  - `terraform apply -auto-approve`
- `terraform destroy`
  - `terraform destroy -auto-approve`
- `terraform output`

Dynamic content blocks - create a list inside a resource. allows us to repeat blocks of contents with a variable into a list.

"aws_instance" has the `user-date` which allows us to run a script at machine start.

`{user_data = file("${path.module}/server-script.sh")}` - use `${path.module}` to get the location of the current module.

policy documents have a size limit.

we can connect to a remote backend with the `terraform` configuration block. [documentation](https://www.terraform.io/language/settings/backends/s3)

iterating with `count`, `count.index`.

### Aws Resources

Resource | Terraform Block Name | Usage | Documentation
---|---|---|---|
EC2 | "aws_instace" | virtual machine on the cloud | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance
network | "aws_vpc" | virtual network | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc
Elastic IP | "aws_eip" | static public ip address | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip)
Security group | "aws_security_group" | statefull firewall | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group
IAM User | "aws_iam_user" | aws user | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user
IAM Policy | "aws_iam_policy" | aws policy | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy
IAM Policy Attachment | "aws_iam_policy_attachment" | attach policy to users/ groups/roles | https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy_attachment
AWS Database | "aws_db_instance" | Relational Database |https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance
</details>

