<!--
ignore these words in spell check for this file
// cSpell:ignore Okun kube fooa SVennam aaaabc kubelet kubernetes nginx openshift linux redhat Silberschatz Korth Sudarshan ntile cume_dist arity ODBC JDBC OLAP
-->

# Books-Resources

where I store my learning from books, YouTube videos, and other sources which don't fit elsewhere.

## Geneal stuff

### FireShip - WebSockets in 100 Seconds & Beyond with Socket.io

<details>
<summary>
A bidirectional connection over tcp/ip.
</summary>

[WebSockets in 100 Seconds & Beyond with Socket.io](https://youtu.be/1BfCnjr_Vjg)

good for realtime data. high speed, low latency._event based_ (callbacks).
we start with an tcp/ip connection, and then move to the websocket protocol (_ws_).

Socket.io is library that extends the abilities of the websocket and provides common solutions like broadcasting (messaging all other clients on this url)

_**cors** : cross origin resource sharing_

there're also **webRTC** for video/voice, and **webTransport** as a possible upgrade for the webSocket protocol.

</details>

### IBM Kubernetes Essentials

<details>
<summary>
IBM Kubernetes
</summary>

[Kubernetes Essentials](https://youtube.com/playlist?list=PLOspHqNVtKABAVX4azqPIu6UfsPzSu2YN)

#### Kubernetes Explained

Kubernetes is an orchestration tool to manage containerized applications

Kubernetes master has the Kubernetes API, the customer worker nodes are also kubernetes managed, and each has a kubelet.

Kubernetes use yaml to define the resources needed, we start with a small example

```yml
kind: pod
image: SVennam/frontend:latest
labels:
  app: frontend
```

we can deploy that manifest with kubectl, the kubernetes command line tool. this sends the configuration to a worker node.

we can also define a template in the manifest for how pods look

```yml
kind: deployment
selector: { app:frontend }
replicas: 3
template:
  image: SVennam/frontend:latest
  labels:
    app: frontend
```

now we use kubernetes to manage that deployment and ensure that state, so it will create more pods like that.

each pod has an ip address, but when a pod dies it gets a new ip address, if we want to talk to a pod, we need something to manage that. this is done with a service definition. this is also in the yaml file, this creates a reliable way for the KubeDNS to communicate.

```yml
kind: service
selector: { app:frontend }
type: LoadBalancer
```

the LoadBalancer type makes our nodes exposed to the outside world.

we have pods, deployments and services as resources of kubernetes.

#### How does Kubernetes create a Pod?

how are pods created by kubernetes?

nodes are the machines, we have control nodes and compute nodes.
if we want to make a pod, we use kubectl to talk with the kubeAPI server. the first thing that happens is authentication, and next the request is written to etcd (a key-value datastore), that is the source-of-truth for kubernetes.

the etcd defines the desired state for Kubernetes.

the scheduler controls which workloads needs to be assigned to a worker node, it talks to the kubeAPI server.

the compute nodes have a kubelet, which talks with the control plane and the kubeAPI, it's registered in the cluster. the compute nodes also have a container runtime engine (CRI - container runtime initiative,not only Docker), and a kube-proxy for communication purposes.

the scheduler chooses to which compute node the work should be passed, it tells the kubeAPI which, and then that is written to the etcd. then we have a desired state.

the last part of the Control node is the controller manager. it has the replication controller and it also monitor the state of the world against the desired state. if it sees that a pod is missing, it knows how to request a new one to be created.

- control node
  - kubeAPI server
  - etcd
  - scheduler
  - controller manager
- compute node
  - kubelet
  - kube-proxy
  - CRI - container runtime engine

#### Kubernetes Ingress in 5 Minutes

assume we have a service with three pods,

service types: Cluster-IP, NodePort, LoadBalancer

NLB: node work balancer.

ingress is resource, not a service type.\
it's set of rules, like nginx reverse proxy, it uses a load balancer, has path based routing and more stuff.

#### What Is Helm

Helm is a Kubernetes package manager that makes deployment easier.

let's take an example of a e-commerce site, which has a JS frontend application, with a mongo database and a node-port service.

in Kubernetes, we will have files for deployment that define the configuration.

```yaml
#deployment
image: node/mongo
replicas: 2

#service
type: NodePort
port: 8080
```

helm can help us separate the template of the configuration from the files themselves. we will have values.yml file which acts as our **chart**.

```yml
deployment:
  image: node/mongo1
  replicas: 2
service:
  type: NodePort
  port: 8080
```

the chart makes it possible to pull values from an external source.

```yaml
#deployment
image: { { values.deployment.image } }
replicas: { { values.deployment.replicas } }

#service
type: { { values.service.type } }
port: { { values.service.port } }
```

the command that we run is, helm will inject the parameters from the chart and send them into a **Tiller** component on the kubernetes side.

```sh
helm install <myApp>
```

when we want to change the values, we simply update the chart, we can also rollback changes, and save the changes to a repository for future use.

```sh
helm upgrade <myApp>
helm rollback
helm package
```

#### Kubernetes vs. Docker: It's Not an Either/Or Question

we still use all the knowledge we got when we used docker, we build on top-of it to get a better deployment.

it helps us with scaling up,orchestration replaces scripts that we would have written otherwise. deployment is easier, development is easier, and monitoring is done for us built-in.

a deployment is always alive, it's the desired state, no matter what happens, kubernetes will try to get back to that state.

#### Kubernetes Deployments: Get Started Fast

kubernetes resources.

deployment -> replica set. rolling update.

debugging

- kubectl logs
  - _--previous_ - from a crushed container
- kubectl describe pod
- kube exec -it sample-pod -- /bin/bash

#### Advantages of Managed Kubernetes

#### Using IBM CloudLabs for Hands-on Kubernetes Training on IBM Cloud

#### Kubernetes and OpenShift: What's the Difference?

openshift by redhat (not open source), OKD - origin kubernetes deployment.

kubernetes:source code, image registry, ci-cd cycle.

openshift is opinionated, it has defaults, takes less time. doesn't run everything as root.

#### Containerization Explained

not only docker. vms vs containers. how much resources are used.

#### Container Orchestration Explained

applications, orchestrator.
the development team cares about the applications, the operation teams cares about a whole lot more. deployment, scaling, networking (load balances, service discovery), insight, maintenance, plug-in configurations.

service mesh.

</details>

## Regex Stuff

<details>
<summary>
Things to remember about regex
</summary>

patterns:

- Email: `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`
- URI: `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`
- IPv4: `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`

commands:

```ps
# regex match
Select-String -Path .\regex22.txt -Pattern "^fooa+bar$"
# simple match
"aaaabc" -match "a+bc"
# simple replace
"aaaabc" -replace "a+","x"
# then single quote matters
"aaaabc" -replace "(a+).([cd])",'$2 then $1'
# from file
(Get-Content -Path .\udemy_regular_expressions_mastery\input_files\regex19.txt) -replace "^(l)(.*)",'Z$2'
```

- [About Regex](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_regular_expressions?view=powershell-7.2)
- [PowerShell](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/select-string?view=powershell-7.2)

</details>

## Markdown Formatting Tips

<details>
<summary>
Markdown formatting tricks, tips, etc
</summary>

[cspell dictionaries](https://cspell.org/docs/dictionaries/), [words in dictionaries](https://github.com/streetsidesoftware/cspell-dicts/tree/main/dictionaries).

[supported languages in code fences](https://github.com/highlightjs/highlight.js/blob/main/SUPPORTED_LANGUAGES.md)

### Tags

Dialog Box - don't use

<dialog open>This is an open dialog window</dialog>

\
\
\
\.

definition and tooltip

<dfn><abbr title="HyperText Markup Language">HTML</abbr></dfn> is the standard markup language for creating web pages.

### keyboard tricks

- Emojis <kbd>Windows</kbd> + <kbd>.</kbd>: ❄
- Unicode: <kbd>Alt</kbd> + four digits of decimal number.
- HTML Entity:
  - <<kbd>&#</kbd>>\<Decimal Number><<kbd>;</kbd>> &#945;
  - <<kbd>&#x</kbd>>\<Hexadecimal Number><<kbd>;</kbd>> &#x3B2;
  - <<kbd>&</kbd>>\<symbol name><<kbd>;</kbd>> &gamma;

[some html symbols](https://www.w3schools.com/charsets/ref_utf_punctuation.asp)

| symbol      | code     |
| ----------- | -------- |
| left arrow  | &larr;   |
| right arrow | &rarr;   |
| two arrows  | &#8644;  |
| alpha       | &alpha;  |
| Weird A     | &#x0041; |
| plus minus  | &pm;     |
| empty       | &empty;  |

### Latex

[latex symbols](https://www.cmor-faculty.rice.edu/~heinken/latex/symbols.pdf)

$$
\begin{align*}
state_m = state_n + path_{(n,m)} \\
state_m = state_n + delta_{(n,n+1)} + ... + delta_{(m-1, m)} + metadata_m
\end{align*}
$$
$$
\begin{align*}
\int_0^1 \frac{4.0}{1+x^2}dx = \pi
\end{align*}
$$
$$
\begin{align*}
\sum\limits_{i=0}^{n}F(X_i)\Delta \approx \pi
\end{align*}
$$

</details>

## Other stuff

<details>
<summary>
MISC
</summary>

set environment variables

- linux: `export AWS_DEFAULT_REGION=us-east-1`
- windows cmd: `set AWS_DEFAULT_REGION=us-east-1`
- powershell: `$Env:AWS_DEFAULT_REGION='us-east-1'`

</details>
