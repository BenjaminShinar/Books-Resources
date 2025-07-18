<!--
// cSpell:ignore jobspec datcenters pkill
-->

<link rel="stylesheet" type="text/css" href="./markdown-style.css">

# HashiCorp stuff

## Nomad

<details>
<summary>
Workload orchestrator.
</summary>

[getting started repository](https://github.com/hashicorp-education/learn-nomad-getting-started)

> Nomad is a flexible workload orchestrator that enables an organization to easily deploy and manage any containerized or legacy application using a single, unified workflow. Nomad can run a diverse workload of Docker, non-containerized, microservice, and batch applications.\
> Nomad enables developers to use declarative infrastructure-as-code for deploying applications. Nomad uses bin packing to efficiently schedule jobs and optimize for resource utilization. Nomad is supported on macOS, Windows, and Linux.

### Quick Start

[nomad quick start](https://developer.hashicorp.com/nomad/tutorials/get-started)

> Nomad is a flexible scheduler and workload orchestrator that enables you to deploy and manage any application across on-premise and cloud infrastructure at scale.

(video)

application layer and the VM / OS layer. two different audiences that share a resource. nomad splits this, having separate workflows for the application and the OS.\
The developer can write a job file (declarative style), and nomad finds where to run this on the cluster. we can update the application configuration and control the deployment strategy.\
in the operational level, we also need some way to restart the application if it fails, and restart the machine (if it fails), and maybe even reschedule it to run somewhere else.

for the operator, the focus is on the machines and OS, with a different APIs and goals, such as patching the OS and replacing the machines.

in most cases, we have a strong machine with with very low utilization (2% percent), the target goal is usually to have around 30% utilization per machine. with nomad, we can run our applications as container images, and run several of them on the same machine. there are cases which the application isn't containerized yet, and that's also a use-case for nomad.

nomad has job files, but also has APIs for jobs, and a job queue. for example, CircleCI uses nomad to run builds (1000 jobs a minute), even if there isn't enough capacity at one moment, the queue will handle the requests eventually.\
A different use-case is high-performance computing (HPC), there is c C1m challenge (one million containers), and nomad can do this in less than five minutes, there are also even larger jobs.

> Some of Nomad's main features include:
>
> - **Efficient resource usage** - Nomad optimizes the available cluster resources by efficiently placing the workloads onto the client nodes of the cluster through a process known as *bin packing*.
> - **Self-healing** - Nomad constantly monitors and detects if tasks stop responding and takes appropriate actions to reschedule them for high uptime.
> - **Zero downtime deployments** - Nomad supports several update strategies including rolling, blue/green, and canary deployments to make sure your applications are updated with zero downtime to your users.
> - **Different workload types** - Nomad's flexibility comes from its use of task drivers and allows orchestration of Docker and other containers, as well as Java Jar files, QEMU virtual machines, raw commands with the exec driver, and more. Additionally, users can create their own task driver plugins for customized workloads.
> - **Cross platform support** - Nomad runs as a single binary and allows you to orchestrate your application across macOS, Windows, and Linux clients running on-premises, in the cloud, or on the edge.
> - **Single unified and declarative workflow** - Regardless of the workload type, the workflow for deploying and maintaining applications on Nomad is unified within a declarative job specification that outlines important attributes like workload type and configuration, service definitions for communication between components, and location values such as region and dataCenter.

nomad is run on agents, which can be of different modes (very similar to consul):

- client agent, runs tasks (workloads) assigned to it. registers itself with servers and waits for work. also called *Node*
- server agent - manages jobs, monitors clients, controls jobs placement on client nodes. server agent can be replicated across multiple machines for High Availability.
- dev agent - development mode. runs both, doesn't write to disk, local, no persistent storage.

nomad operations:

- task - unit of work, executed by `task drivers` such as docker or "exec", tasks specify requirement, configuration, constraints and resources.
- group (task group) - a series of tasks running on the same nomad client
- job - core unit of control, defines application and configuration, contains one or more tasks.
- job specification (*jobspec*) - schema for nomad jobs.
- allocation - mapping between task group in a job and the client.

an application is a *jobspec* with tasks, when submitted to nomad, this becomes a *job* and *allocations* for the task group defined in it.

jobs can be *services* - which run until stopped explicitly, or *batch* - which run until completed.

commands used:

- `nomad agent`
- `nomad node status`
- `nomad ui -authenticate`
- `nomad job run <file-name>`
- `nomad job allocs <job-name>`
- `nomad job dispatch -meta <key-value pairs> <job-name>`
- `nomad job stop -purge <job-name>`

### jobspec file

uses the hashicorp configuration format, like terraform.

```hcl
task "example" {
  driver = "docker"
}

<BLOCK TYPE> "<BLOCK LABEL>" {
  # Block body
  <IDENTIFIER> = <EXPRESSION> # Argument
}
```

each file is a single job, but can contain multiple groups, which can contain multiple tasks. tasks inside the same group run on the same node.

jobs have types, which tell the scheduler how to run them, the important ones are "service" and "batch". we can specify where the job runs with "region" and "datcenters", control the update policy with an `update` block and do other stuff. a batch job can be parameterized, which will load it onto the cluster, and can be *dispatched* (invoked) with different arguments. this is like storing function on the cluster and firing it. jobs also specify vault and consul connections (which can be overridden in group and task blocks).\
the next level is *Group*, a set of tasks running in the same node. we can control the number of copies with `count` value, and register it to HashiCorp's Consul using `service`. many other stuff.\
the *Task* is what does the actual work, it specifies a `driver` (docker, java, exec, qemu) (like a runtime platform) and a `config` block (what to do), the task can be a docker image, a bash command, etc.. there are many thing to specify here.
</details>

## Vault

<!-- <details> -->
<summary>
Manage Secrets & Protect Sensitive Data
</summary>

> What is Vault?\
> Secure, store, and tightly control access to tokens, passwords, certificates, encryption keys for protecting secrets, and other sensitive data using a UI, CLI, or HTTP API.

[get started tutorial](https://developer.hashicorp.com/vault/tutorials/get-started)

(video)\
managing secrets and credentials: user names and passwords, database credentials, api token or TLS certificates. access control and rotation.\
all secrets in vault are encrypted, secrets have fine grained access control, and are audited and we log whenever someone accesses the secret (audit trails). the next problem is that applications might leak and expose the secrets, like sending them to a log file, an error report, or something like that. the solution is to provide the application short-lived credentials, which stops being valid after a period of time. this also allows us to create unique secrets, so if a secret is exposed, it's easy to track where the leak was, rather than just know it was one of the services who use it. we can also revoke access based on unique secrets.\
Vault also provides an "encrypt-as-a-service" feature, it provides good cryptographic implementation and key management as high-level APIs.

- Storing secrets and avoid secret sprawl
- protect again leaks with dynamic secrets
- protect data at rest with good cryptographic implementations

vault also has plug-ins and extensions, one such plug-in is for authentication, which can work with different identity providers and connect those identities with Vault. vault connects with different logging systems to store the trail logs (auditing data), the secret data itself are stored in highly available durable backend storage. plug-ins also allow us to do dynamic secrets, a plug-in can work against a database and through it generate the short-lived credentials.\
High Availability is achieved with consult, shared network service. exposes REST api.

vault has community and enterprise editions, enterprise edition has more features, can deployed both as a service and self-managed (community can only be self-managed).

plugins for client authentications and plugins for communicating with databases and other things which have passwords/secrets/credentials.

Vault is packed as a binary, it needs to be added to the PATH. then we can run it from the cli. we can install it from package managers.

### Set up Vault

we can create a vault server in development mode on the local machine. it does not persist data, and it unseals data automatically,

```sh
vault
vault -help
vault server -dev
# in other terminal
# grab the unseal key and token values, and save them somewhere
echo "<unsealed.key>" > unseal.key
export VAULT_DEV_ROOT_TOKEN_ID=<token>
export VAULT_ADDR='http://127.0.0.1:8200'
# export VAULT_CACERT= # not for vault-dev mode
vault status
# authenticate as root
vault login root

# clean up
pkill vault
unset VAULT_ADDR VAULT_CACERT
```

### Token and Secrets

to authenticate against vault, we need a token, this is what we grabbed and saved in the previous example. that was the root token, which has the root policy. tokens usually have a TTL, after which they must be renewed. tokens can have a limit on usage, even before the TTL expires. a token is identified by the accessor, which is the unique name.

```sh
vault token lookup -accessor <token accessor>
```

we can't use token accessors to login to Vault.

there are different types of tokens - service tokens, batch tokens and recovery tokens (not relevant for now). we usually use service tokens, unless the situation calls for batch tokens - which are easier to replicate, but aren't flexible and have less features. there are also periodic tokens and orphan token.

```sh
vault token create -policy="default" -period=24h
vault token create -policy=default -orphan
vault token renew -accessor <token accessor>
```

we can revoke a token manually

```sh
vault token create -policy=default -period=10m
# grab token
vault token revoke <token>
```

</details>
