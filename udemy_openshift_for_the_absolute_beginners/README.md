<!--
// cSpell:ignore Mesos LACP
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

> "**Openshift Origin** is based on top of Docker containers and the Kubernetes cluster manager, with added developer and operational centric tools that enable rapid application development, deployment and lifecycle management."

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

we click <kbd>next</kbd> in the wizard to install it. now we open the manager interface. we get an image from [OsBoxes website](https://www.osboxes.org/), we choose ubuntu and download the image to our computer. we open the zipped file and extract a _.vdi_ file.

in the management window, we click <kbd>new</kbd> and give the machine a name, and then we decide the type and version (linux release,32 and 64 bit options).

if we don't see the 64 bit option, we might need to enable the virtualization in the bios.

we choose the memory resource and give the machine some diskspace. then <kbd>browse</kbd> to choose the image we downloaded. before powering it up, we click <kbd>settings</kbd> and under _network_ we set the network type from **NAT** to **bridge**. we can also take a snapshot of the machine before starting.

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

<details>
<summary>
Openshift concepts - projects, users, builds, deployments.
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

we want to add an application to our project, we integrate source code management services such as github into Openshift.

when we add an application, we must specify the repository location, the build job is created automatically (clones the repository and builds an image), and it's pushed into a openshift repository.

a successful build also triggers an Openshit deployment, which is similar, but not the same as kubernetes deployment.

kubernetes deployment:

```yaml
apiVersion: apps/v1
kind: Deployment
```

openshift deployment:

```yaml
apiVersion: apps.openshift.io/v1
kind: DeploymentConfig
```

adding an application does all the steps:

> 1. Create Build
> 2. Download Source
> 3. Build Image
> 4. Push to Repository
> 5. Deploy

#### Demo - Deploy an Application

an example of deploying a simple application on the Openshift origin. we start with the openshift console, and a gitlab repository with two files: "app.py" and "requirements.txt".

we create a new project, select it, click <kbd>Browse Catalog</kbd>, choose python, and fill in the details in the wizard, the name of the application and the source code path.

now we move to the **overview** tab, and see that the deployment is in process. we can then move to the **builds** tab, to see the build pipeline in process. once it finished, we can see the deployment again and get the public ip address.

### Builds

build strategies and build configurations.

strategies:

- Docker Build
- Source To Image (S2I)

imagine we have this simple python app, built with the Flask framework.

```py
import os
from flask import Flask
app = Flask(__name__)

@app.route('/')
def main():
  return "Welcome"

@app.route('/how are you')
def hello():
  return "I am good, how about you?"

if __name__ == "__main__":
  app.run(host="0.0.0.0", port=8080)
```

kubernetes expects a docker image, so we need the configration to build it. manually, the dockerfile would look like this. this is the **Docker Build Strategy**

```docker
FROM ubuntu:16.04

RUN apt-get update && apt-get install -y python python-pip

RUN pip install flask

COPY app.py /opt/

ENTRYPOINT FLASK_APP=/opt/app.py flask run --host=0.0.0.0
```

the other strategy is **Source To Image Strategy**, which skips the manual docker file. and uses pre-built configurations.

we can also build artifacts, like library files, with the **Custom Build** option.

the internal registry is located at the address of "172.30.1.1:5000", we can use **Image Streams** to map images at registries to images used by the applications, this uses the image sha, so even if the image at the source changes, the original image is still used.

in the web console, we can see the Build configuration, and under the <kbd>Actions</kbd> button, we can view the build configuration yaml and see how it relates to the other fields/

```yaml
kind: "BuildConfig"
apiVersion: "v1"
metadate:
  name: "simple-webapp"
spec:
  runPolicy: "Serial"
  triggers:
    - type: "GitHub"
      github:
        secret: "<>"
    - type: "Generic"
      generic:
        secret: "<>"
    - type: "ImageChange"
  source:
    git:
      uri: "https://github.com/mmushad/simple-webapp-flask.git"
  strategy:
    type: Source
    sourceStrategy:
      from:
        kind: "ImageStreamTag"
        name: "python:3.6"
  output:
    to:
      kind: "ImageStreamTag"
      name: "simple-webapp:latest"
```

and to a docker file build strategy, we modify this file

```yaml
kind: "BuildConfig"
apiVersion: "v1"
metadate:
  name: "simple-webapp-docker" #changed
spec:
  runPolicy: "Serial"
  triggers:
    - type: "GitHub"
      github:
        secret: "<>"
    - type: "Generic"
      generic:
        secret: "<>"
    - type: "ImageChange"
  source:
    git:
      uri: "https://github.com/mmushad/simple-webapp-docker.git" #dockerfile
  strategy:
    type: Docker
    dockerStrategy:
      from:
        kind: "DockerImage"
        name: "ubuntu:16.04"
  output:
    to:
      kind: "ImageStreamTag"
      name: "simple-webapp:latest"
```

to add this configuration, we click <kbd>Add To Project</kbd> and then <kbd>Import</kbd>, and now we will have two configurations.

#### Demo - Builds

in the web console, we choose a project and an application, and view the build configuration. we can also see the web hooks/triggers, add environment variables and view events.

now we do the same process as before, in this example we create the docker file. we also create a new ImageStream tag, we create it under the <kbd>Images</kbd> tab, and duplicating the existing json and changing the name.

#### Coding Exercises: Builds

```yaml
apiVersion: build.openshift.io/v1
kind: BuildConfig
metadata:
  name: my-build-config
spec:
  runPolicy: Serial
  source:
    git:
      ref: master
      uri: https://github.com/mmumshad/simple-webapp-docker
    type: Git
  strategy:
    dockerStrategy:
    type: Docker
  output:
    to:
      kind: ImageStreamTag
      name: simple-webapp-docker:latest
  triggers:
    - type: ConfigChange
```

### Build Triggers

so far, we've initiated builds in a manual way, but in CI-CD cycles, we want to have automated builds. this means that when the source code is changed, then a new version of the application is built using the new code.

this is done by using webhooks, which are triggered when an evert (code change) happens, and they notify a diffrent service.

in github, we can look the <kbd>settings</kbd> tab, choose <kbd>webhooks</kbd> and then paste the url from the build configuration details.

we just need to make sure that github and openshift can talk to one another, they need to belong to the same network!

#### Demo - Build Triggers

adding a webhook, there are some defaults, but we can create more, under <kbd>Actions</kbd>, in **triggers** and we select **gitlab** as the type, and we create a new secret. once created, we get a url that we can paste into gitlab.

we might need to allow requests to local network if our openshift isn't public.

### Deployments

deployment in OS are similar to kubernetes deployments.

in kubernetes, the smallest unit is a pod (which contains one or more containers), next up is the replica-set, which is a set of pods. another level up is the deployment, which controls life time and updating.

in openshift, we can view the deployments under the applications tab in our project. the deployment uses the build configuration. it has the number of replicas and the deployment strategy (such as rolling update). just like builds, we can set deployment manually or by a trigger, and it is configured out of the box to run when the source code changes.

we can configure the deployment at a yaml file or with a wizard. we can rollback to previous deployments.

strategies:

- Recreate - destroy all, deploy all new.
- Rolling Update - destroy one, deploy one.
- Blue Green Update - deploy all, route to new, destroy old.
- AB deployment - splitting traffic between versions.

```sh
oc rollout latest dc/simple-webapp-docker
oc rollout history dc/simple-webapp-docker
oc rollout undo dc/simple-webapp-docker
```

#### Demo - Deployments

in the web management, we add a deployment, we click the <kbd>learn more</kbd> link to get a template for a deploymentConfig yaml.

we modify is and update the triggers. we click <kbd>add to project</kbd> and import the configuration file and fix any issues with the file.

we update the source code in the version control, which triggers a build, which triggers a deployment.

</details>

## Networks, Services, Routes and Scaling

<details>
<summary>
Communication Between Pods
</summary>

Kubernetes cluster are composed of nodes (master and workers), which run pods. the pods might need to communicate with one another, which means they must be on the same network, and have a unique IP address which they can access.

openshift uses SDN - Software Defined Network, which is a virtual network, an overlay network.

Open vSwitch:

- VLAN Tagging
- Trunking
- LACP
- Port Mirroring

the default id for this overlay network is _10.128.0.0/14_, where each node has a unique subnet

- _10.128.0.0/23_
- _10.128.2.0/23_
- _10.128.4.0/23_

if we want to see the ip addresses.
`oc get pods -o wide`

but those ips aren't set int stone, if a pod goes down, it will come back with a differnet ip. Openshift comes with a builtin DNS, based on SkyDns.

SDN plugins:

- **ovs-subnet** - connects all pods in the cluster (across all projects)
- **ovs-multitenant** - virtual network for each project.

there is also support for other plugins. if we want external connectivity, we use services and routes.

### Services and Routes

services are a way to connect pods with one another, rather than using ip addresses or dns name, we have the service which acts as a load balancer. we can use services to connect between pods internally, between the outside world and our pods, and our pods can connect to other sites/services.

each service has an ip address and DNS entry, there is an internal ip assigned, which is called **cluster IP**. services are linked to pods by using **selectors**. there are also the service ports and the target port.

to connect the service to the outside world, we use **Routes**, which act as a proxy. routes uses load balancing strategies, such as

- source - (default strategy) - always matching the same external ip to the same target, making it a _sticky session_.
- roundrobin - each connection is routed to the next target.
- leastconn - a connection is routed to the least used backend target.

we can configure SSL connection in the **security section**, we can configure http/https access, certificates, and configure a/b routing.

#### Demo - Services and Routes

in the web console, we take an example service file from the _learn more_ link, and we'll start editing it.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: simple-webapp-docker
spec:
  selector:
    deployment: simple-webapp-docker
  ports:
    - name: 8080-tcp
      port: 8080
      targetPort: 8080
      protocol: TCP
```

now we import the file, and create the service, and we see the ip address. we can use the internal ip (ClusterIP) to access the pod from inside the cluster, but if we want to access it from outside, we need to create a **route**.

we fill in the details, if we want to use an exiting DNS, we need to have it set up to forward to the openshift route.

we can make a change to the application code and have the end-to-end cycle start again.

### Scaling

deployment controller -> replication controller.

the number of pods (replications) is controlled by the replication contoller. we can change it in the yaml file or in the web console.

#### Demo - Scaling

in the web management console.

we start a new application, which has a random colored background. we create a new projects, a new application, a new deployment. we can increase the number of pods by clicking the up arrow in the application page. this creates a new pod, which is managed by the same service, each access to the website is done by the service controller.

by default, the load balancing strategy uses the stickySession strategy, which means that the same client (browser) will always be directed to the same instance of the application.

we can change this by using a different load balancing strategy, such as round robin, and by disabling the cookies.

</details>

## Storage, Templates and Catalog

<details>
<summary>
Storage (volumes, claims). Templates, Catalog.
</summary>

### Storage

docker containers are meant to be ephemeral, and aren't supposed to hold data over time. if we wish the data to remain, we need to use persistent volumes.

openshift uses the same Kubernetes plugins to persist data and add storage to the cluster. we have persistant volumes and persistant volumes claims.

<kbd>Create Storage</kbd>

access modes:

- Single User (RWO)- read and write to a single user.
- Shared Access (RWX) - read and write for multiple users.
- Read Only (ROX) - read only for multiple users.

we give the volume a name and a size. and then we can add the storage to a pod.

#### Demo - Storage

we have a program that reads a file from a storage, so we need to provide a volume claim to make the pods able to access this file.

<kbd>Create Storage</kbd>, and then in the deployment configuration, we select <kbd>add Storage</kbd> and select our new volume.

in the example, the data is shared across multiple instances of the application. they can all see the changes done by the others.

### Example Voting Application Introduction

microservice architecture using a simple application, the voting application sample.

- voting-app
- in-memory DB (redis)
- worker (.Net)
- db (Postgres SQL)
- results-app (node JS)

we have different services and different programs in our deployment.

#### Demo - Deploy Example Voting Application on Openshift

now in openshift.

we create a new project. and now start deploying appliication. we take the template from github, copy it to a file and import it, so now we can use the redis template. we use the deafult values and set up a password. we next move to the python frontend application, we again use the default option. we need to pass the REDIS password as an environment variable.

we now move to the results side, a postgres SQL database and a node JS results page. for now we use environment variables, but in real life scenarios, we should use secrets. we also change the default port to 8080.

the final part is to deploy the worker application. this time we will use a docker-build strategy, we edit the yaml file and change the strategy. we can see the changes in the log of the application.

### Templates and Catalog

the catalog consists of the known applications that openshift knows how to handle, if we look at the **Django + Postgres SQL** option:

- image stream
- Build
- Deployment (application)
- Deployment (database)
- service (port 5432)
- secret (database credentials)
- Route
- Additional User Parameters

this is a template, it contains data about deploying a stack.

```yaml
apiVersion: v1
kind: Template
metadata:
  name: custom-app
objects:
  - apiVersion: v1
    kind: Secret
    #more fields
  - apiVersion: v1
    kind: Service
    #more fields
  - apiVersion: v1
    kind: Service
    #more fields
  - apiVersion: v1
    kind: Route
    #more fields
  - apiVersion: v1
    kind: BuildConfig
    #more fields
  - apiVersion: v1
    kind: DeploymentConfig
    #more fields
  - apiVersion: v1
    kind: DeploymentConfig
    #more fields
  - apiVersion: v1
    kind: ImageStream
    #more fields
parameters:
  - displayName: "Namespace"
    name: "NAMESPACE"
```

we can now create this template as a catalog item by running the openshift command
`oc create -f template-config.yml`.

#### Demo - Create a custom Catalog

we want to take the voting application and make it into a template / catalog item.

_example-coting-app-template.yml_
now we start copying the templates for all of the app components, we clear the unneeded parts. we start with the secrets, which are the passwords fot the postgres and redis databases.

```yml
apiVersion: v1
kind: Template
metadata:
  creationTimeStamp: null
  name: example-voting-app-template
objects:
  - apiVersion: v1
    kind: Secret
    data:
      database-name: db
      database-password: <> #postgres password
      database-user: postgres_user
    metadata:
      name: db
  - apiVersion: v1
    kind: Secret
    data:
      database-password: <> #redis_password
    metadata:
      name: redis
```

we do the same with the build configuration, and then the image streams configurations, we have 5 deployment configurations which we need to include. the next part is the services yaml files. and finally the routes.

we need to provide the namespace, because we can't add templates to the default namespace.

`oc create -f example-coting-app-template.yml -n someNameSpace`

we could also make things more complex by requesting parameters from the user in the wizard.

</details>

## Misc

<details>
<summary>
Additional stuff
</summary>

### Yaml

yaml is a file format, like xml and json. it is used to represent data, these three files represent the same data.

xml

```xml
<Servers>
  <Server>
    <name>Server1</name>
    <owner>John</owner>
    <created>12232012</created>
    <status>active</status>
  </Server>
</Servers>
```

json

```json
{
  "Servers": [
    {
      "name": "Server1",
      "owner": "John",
      "created": 12232012,
      "status": "active"
    }
  ]
}
```

yaml

```yaml
Servers:
  - name: Server1
    owner: John
    created: 12232012
    status: active
```

yaml uses key-value pairs, the format is\
`<key>: <value>`

where the space after `:` matters.

we can have arrays/list with the `-` to indicate elements, or inner dictionaries/maps. the indentation matters, elements at the same level must algin.

```yaml
Fruits:
  - Oranage
  - Apple
  - Banana

Banana:
  Calories: 105
  Fat: 0.4 g
  Carbs: 27 g
```

we can either have a value or a inner object (list/map), not both.

we can have nested elements, lists inside dictionaries, and so one.

**Dictionary vs List vs List of Dictionaries**

to store information on a single object, we use a dictionary. for many elements of the same type, we use a list.

```yaml
listExample:
  - element1
  - element2
  - element3

listOfDictionaries:
  - Key1: Value1
    Key2: Value2
  - Key1: Value3
    Key2: Value4
  - Key1: Value5
    Key2: Value6
```

dictionaries are unordered, while lists are ordered. dictionaries with the same proprties are identical if the values are the same, but lists aren't identical if the elements are identical, but not in the same values.

#### Yaml Exercise

```yaml
Fruit: Apple
Drink: Water
Dessert: Cake
Vegetable: Carrot
```

```yaml
Fruits:
  - Apple
  - Banana
  - Orange
Vegetables:
  - Carrot
  - Tomato
  - Cucumber
```

```yaml
Fruits:
  - Apple:
      Calories: 95
      Fat: 0.3
  - Banana:
      Calories: 105
      Fat: 0.4
  - Orange:
      Calories: 45
      Fat: 0.1

Vegetables:
  - Carrot:
      Calories: 25
      Fat: 0.1
  - Tomato:
      Calories: 22
      Fat: 0.2
  - Cucumber:
      Calories: 8
      Fat: 0.1
```

```yaml
Employee:
  Name: Jacob
  Sex: Male
  Age: 30
  Title: Systems Engineer
  Projects:
    - Automation
    - Support
  Payslips:
    - Month: June
      Wage: 4000
    - Month: July
      Wage: 4500
    - Month: August
      Wage: 4000
```

### GitLab - Setup

setting up a local gitlab, virtual machine. we start with the orcale virtual box, so now we use it to deploy a **CentOS** machine.

we go to the osBoxes website and pull the CentOS image. we extract it to a folder. we create a new virtual machine, the os system is _other (linux-64)_, we then take the image as the base for the hard drive. we need to set up a consistent, static ip address, so we need to add an adaptor, and disable the **dhcp** option.

the default username and password is mentioned in the site where we took the image.

we close the machine and then choose <kbd>clone virtual machine</kbd> (linked clone) to use the machine as a template.

we need to run a script to get the static ip for the cloned machine.

we get the gitlab to image from the gitlab website, we can use a docker image if we want.

now we can use the image to host a gitlab web console.

</details>

## Takeways

<details>
<summary>
Things to remember
</summary>

| Acronym | Full Name                    | Notes |
| ------- | ---------------------------- | ----- |
| OCR     | Openshift Container Registry |       |
| SCM     | Source Code Management       |       |

- `oc whoami` - get user
  - `oc whoami -t` get token
- `oc status`
- `oc login`
- `oc logout`
- `oc get <resource type>` - list resource
  - `oc get projects`
  - `oc get users`
- `oc adm add-cluster-role-to-user <cluster role> <user name>`
- `oc rollout latest dc/<deployment-name>`
- `oc rollout history dc/<deployment-name>`
- `oc rollout undo dc/<deployment-name>`
- `oc get <resource-type>`
  - `oc get <resource-type> -o wide`
- `oc create -f <file>`
- `oc export service db`
</details>
