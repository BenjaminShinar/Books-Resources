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

### Openshift - Architecture Overview
### Openshift - Setup
### Setup VirtualBox
#### Demo - OpenShift Setup with Minishift
### Student Preference
### Management - Web, CLI and REST API
#### Demo - Openshift Management



## Openshift Concepts - Projects andUsers

### Projects and Users
#### Demo - Projects and Users

## Concepts - Builds and Deployments

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
</details>
