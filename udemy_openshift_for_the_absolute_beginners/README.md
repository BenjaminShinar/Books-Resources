<!--
// cSpell:ignore Mesos
 -->

# Openshift For The Absolute Beginners

udemy course [Openshift for the absolute beginners](https://www.udemy.com/course/learn-openshift/).


## Introduction
<details>
<summary>
Course Introduction.
</summary>

### Introduction to Openshift
high level overview

Openshift is **Redhat**'s open source container platform, platform as a service.
- IaaS - Infrastructure as a service
- PaaS - Platform as a service
- SaaS - Software as a service

four different flavours:
1. Openshift Origin - open source application container platform.
2. Openshift Online - pulbic application development hosting service.
3. Openshift Dedicated - managed private cluster on AWS/Google clouds.
4. Openshift Enterprise - On-Premise private PaaS.

this course will mostly work with openshift origin.

>"**Openshift Origin** is based on top of Docker containers and the Kubernetes cluster manager, with added developer and operational centric tools that enable rapid application development, deployment and lifecycle management."

Openshift adds tools and support on top of Kubernetes clusters, with built in integrations for easier manager.
- source code manager (like github)
- pipeline
- registry
- software defined networking
- api-centric
- governance (access management)


</details>

## Pre-Requisite - Docker and Kubernetes

<details>
<summary>
Docker and Kubernetes high level overview. Understanding orchestration.
</summary>

Kubernetes = Containers + Orchestration

### Docker Overview

Docker is the most popular container technology.

- Hardware
- OS
- Libraries and Dependencies

complex software architecture requires a lot of diffrent services, such as a webserver, databases, messaging and so on. each service might need different versions, and those dependencies might conflict. and depending on how the software will be deployed, there might be other issues, this is called **the matrix from hell**.

there are problems with onboarding new developers, and each environment (development, testing, production) can also add complexity.

Containers allow us to package the dependencies for each part of the service, and run it independently, no matter the machine.

containers are separated environments that run isolated process, similar to a VM, but lightweight, as they share the underlying kernel.

containers are layered, each layer has the additional parts of the image. we can't run an windows container on a linux engine, but it doesn't matter so much, because the point of containers is to run applications, not to virtualize a machine. this is why containers are light weight compared to virtual machines. starting a container takes seconds, as opposed to minutes in virtual machines.

containers share more resources between one another as opposed to virtual machines.

docker registry - docker hub, which stores images for applications.

we can run many instances of the same image.
```sh
docker container run ansible
docker container run mongodb
docker container run redis
docker container run nodejs
```
Containers vs Image:

Images are used to create containers. an image is a blueprint, while the container is the instance of the application, each container with it's own "file system" layer.

images can be run on any container engine,and it will run the same no matter if it's done on the developer machine or by the operation teams.

### Kubernetes Overview

once we have images and containers, we can move on the orchastration, this matters when we have containers which are dependent on one another and must interact with one another, or when we need to scale up the amount of services to handle higher volume of requests.

Orchestration technologies:
- Docker Swarm
- Kubernetes
- Mesos

docker swarm is easy to set up and get started, Mesos is advanced and hard to learn, while kubernetes is very popular. all cloud providers support kubernetes.

kubernetes allows durability (highly available), scalability, ease of usage (load balancing). we can use kubernetes to run many nodes and even more containers with a declarative approach and configuration files.

</details>

## Getting Started with Openshift

<details>
<summary>
First steps.
</summary>

### Openshift - Architecture Overview

Openshift uses kubernetes as the base layer, with containers and images. it leverages the deployments,pods and services.

openshift can be configured to pull images from a public registry like docker hub, or use the Openshift Container Registry (OCR). 

openshift comes with a management web console that allows users to manage the cluster.

the etcd key-value datastore.

nodes - worker and master nodes.

### Openshift - Setup

setting up openshift:
- on premise
- using public/private cloud

- all in one - master and worker node on the same machine, used for development.
- single master, multiple worker nodes
- highly available with many master and many workers.

minishift is an easy way to get started with openshift on a local machine. it bundles the kubernetes, etcd and openshift into a single ISO image that we can use. we can run it as an image on a virtual machine such as virtual box or vmware.

minishift uses the openshift origin configuration.

### Setup VirtualBox
setting up virtual box.deploying an ubuntu machine on it. we go to the [virtualBox website](https://www.virtualbox.org/), download the file and run it.

we click <kbd>next</kbd> in the wizard to install it. now we open the manager interface. we get an image from [OsBoxes website](https://www.osboxes.org/), we choose ubuntu and download the image to our computer. we open the zipped file and extract a *.vdi* file.

in the management window, we click <kbd>new</kbd> and give the machine a name, and then we decide the type and version (linux release,32 and 64 bit options).

if we don't see the 64 bit option, we might need to enable the virtualization in the bios.

we choose the memory resource and give the machine some diskspace. then <kbd>browse</kbd> to choose the image we downloaded. before powering it up, we click <kbd>settings</kbd> and under *network* we set the network type from **NAT** to **bridge**. we can also take a snapshot of the machine before starting.

the username and the password are provided for us when we download the image in osBoxes website.
in the machine we open the terminal and check our ip and make sure we can ssh into the virtual machine.

```sh
ipconfig
service ssh status
sudo apt-get update
apt-get install openssh-server
service ssh status
```
now we try to shh into the virtual machine, so from the host machine.

#### Demo - OpenShift Setup with Minishift

we download the minishift utility, we choose a proper version,place it in the same drive tha has the virtual box installed, and run the executable with the correct driver. this causes the whole process to start a vm with the minishift image.
```sh
minishift.exe start --vm-driver virtualbox
```
once it's done downloading and setting up, we will see the ip address for the web management console, and some credentials which explain how to login. 

### Management - Web, CLI and REST API

we will start by looking at the ways to interact with the openshift cluster
- web console - the public ip address
- command line interface - with the openshift client
- rest api - used for integrations

we can start with the web console, we can start services and container, and set up projects.

the openshift client cli comes pre-packaged with the minishift.
```sh
oc login -u <username> -p <password>
oc logout
```

we can do rest api calls, but we need to provide an authorization header, the token is a unique key which we can get and is valid for 24 hours.

```sh
oc whoami -t #token
curl https://localhost:8443/oapi/v1/users \
    -H "Authorization: Bearer <Token>"
```

#### Demo - Openshift Management

a bit more about the webconsole and the cli.

the first page is a catalog, which allows us to quickly deploy applications wit a wizard, or to create a ci/cd pipeline. we can also create projects and configure the build/deployment and monitoring tools.

for the cli, we have the `oc` client.
```sh
minishift oc-env # get a command to set the path
SET PATH=C:\Users\usr\.minishift\cache\oc\v3.9.0\windows;%PATH% # add oc to the path
os login -u system:adming #login

oc status #
oc project <project_name> #move between projects
```

curl
```sh
oc whoami -t #token
curl https://192.168.99.102:8443/oapi/v1/projects \
    -H "Authorization: Bearer <Token>"
```


</details>

## Openshift Concepts
<!-- <details> -->
<summary>
Openshift concepts
</summary>

### Projects and Users
a cluster can hold hundred of pods and containers, with many services and endpoints.

projects are a way to separate the resources, it's an isolated part inside the same environment, projects are built on top of kubernetes namespaces, and adds to it more isolation and grouping.

three types of users
- regular users (developers)
- system users - system:admin, system:master,
- service accounts - enabling communications inside the project, have the prefix system:service prefix.

openshift comes with an OAuth server. the configuration for it can be found at "/etc/openshift/master/master-config.yaml" file.

#### Demo - Projects and Users

```sh
oc login -u system:admin
oc get projects
oc get users
```



we can't log into the web console with the system admin users right away, we first need to create a user and give him the proper permissions.

```sh
oc adm add-cluster-role-to-user cluster-admin administrator user
```

now this user can see all the projects (including the ones that were configured by minishift).

### Builds and Deployments
#### Demo - Deploy an Application
### Builds
#### Demo - Builds

### Coding Exercise
#### Coding Exercise 1: Builds - 1
#### Coding Exercise 2: Builds - 2
#### Coding Exercise 3: Builds - 3
#### Coding Exercise 4: Builds - 4
### Build Triggers
### GitLab - Setup
#### Demo - Build Triggers
### Deployments
#### Demo - Deployments

</details>

## Networks, Services, Routes and Scaling


### Networking Overview
### Services and Routes
#### Demo - Services and Routes
### Scaling
#### Demo - Scaling


## Storage, Templates and Catalog

### Storage
#### Demo - Storage
### Example Voting Application Introduction
#### Demo - Deploy Example Voting Application on Openshift
### Templates and Catalog
#### Demo - Create a custom Catalog


## Conclusion

## Takeways

<details>
<summary>
Things to remember
</summary>

| Acronym | Full Name | Notes |
|---|---|---|
|OCR | Openshift Container Registry| | 
|SCM | Source Code Management| |

- `oc whoami` - get user
  - `oc whoami -t` get token
- `oc status`
- `oc login`
- `oc logout`
- `oc get <resource type>` - list resource
  - `oc get projects`
  - `oc get users`
- `oc adm add-cluster-role-to-user <cluster role> <user name>`

</details>
