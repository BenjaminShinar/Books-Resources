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

## Database System Concepts

<!-- <details> -->
<summary>
Database System Concepts book.
</summary>

I'm reading a god-damn book about SQL?

Authors: Abraham Silberschatz, Henry F. Korth, S. Sudarshan.

### Chapter 3 - Introduction To SQL

- relations
- basic sql query structure
  - `select`
  - `from`
  - `where`
  - `order`
- Natural Join
  - `as`
- Set operations
  - `union`
  - `intersect`
  - `except`
  - set comparisons
- Null value (truthiness)
- Aggregate functions
  - `avg`, `min`, `max`, `sum`, `count`
  - `group by`
  - `having`
- sub queries

### Chapter 4 - Intermediate SQL

- Joins
- Views
  - Materialized views
- Transactions
  - `Commit`
  - `Rollback`
- Constraints
- `Check` Clause (arbitrary predicate)
- Referential Integrity
- Data Types
  - user defined types
  - domains
- Default values
- Catalogs and Schemas
- Authorization
  - privileges
  - roles

### Chapter 5 - Advanced SQL

- Accessing SQL from software
  - Dynamic SQL - queries are strings and are processed by the database
  - Embedded SQL - queries are compiled by the client
  - Prepared Statements - faster and safer
  - Callable Statements
- Metadata Features
- ODBC and JDBC
- SQL functions and procedures
- Exception conditions and handlers
- Triggers
  - event, condition, action
  - before or after event
  - per row or per table
- Recursive Queries
  - transitive closures
- Advanced aggregations
  - `rank`, `dense_rank`,
  - `percent_rank`, `ntile`
  - `cume_dist`
  - `row_number`
  - `over`
  - `partition by`
- Windowing
  - `preceding`, `following`
  - `unbounded`
  - `current row`
  - `range between`
- OLAP - online analytical processing - analytics solutions and tools
  - `pivot`(cross tabs)
  - `cube`, `rollup` (generalizations of `group by`)
  - `decode`

### Chapter 6 - Formal Relational Query Languages

- operations - can act on one relation and produce one relation back (unary), or run on pairs of relations (binary)
  - select (unary) - select/filter tuples that satisfy a predicate
  - project (unary) - project relation into new form with some of the attributes. will remove duplicates.
  - union (binary) - combine two relations together. will remove duplicates. relations must be the same arity (number of attributes) and the attributes must be the same domain.
  - set difference (binary) - tuples in one relation but not another. order matters.
  - Cartesian product (binary) - create a new relation with a schema that is the union of the two relations schema. each tuple gets paired with all the tuples in the other relations.
  - rename (unary) - give temporary name to a relation as part of the formula (????)
  - set intersection - appearing in both
  - natural join - Cartesian product which is filtered on some natural common attribute
  - assignment - give temporary name (???)
  - outer join - dealing with missing information (right outer join, left outer join, full outer join)
- aggregations
- MultiSets (may contain duplicates)
- Tuple relational calculus - another type of syntax/form to write algebra
- Domain relation calculus.

the format for unary operator is the symbol, and then the predicate in subscript and finally the relation name in parentheses.

[symbols](https://www.cs.uleth.ca/~rice/latex/worksheet.pdf)

| Symbol name | LATEXcode |symbol|
|---|---|---|
| leftarrow | `\leftarrow` | $\leftarrow$ |
| select | `\sigma` | $\sigma$ |
| project | `\Pi` | $\Pi$ |
| inner join | `\bowtie` | $\bowtie$ |
| left outer join | `\sqsubpet\bowtie` | $\sqsupset\bowtie$ |
| right outer join | `\bowtie\sqsuspet` | $\bowtie\sqsubset$ |
| full outer join | `\sqsupbset\bowtie\sqsuspet` | $\sqsupset\bowtie\sqsubset$ |
| cross product | `\times` | $\times$ |
| rename | `\rho` | $\rho$ |
| less than | `<` | $<$ |
| greater than |`>` | $>$ |
| less than or equal | `\leq`  | $\leq$ |
| greater than or equal | `\geq` | $\geq$ |
| equal | `=`  | $=$ |
| not equal | `\neq` | $\neq$ |
| and | `\wedge` | $\wedge$ |
| or | `\vee` | $\vee$ |
| not | `\neg` | $\neg$ |

$\sigma_{department\_name="Physics"}(instructor)$
$\Pi_{ID, Name, Salary}(instructor)$

combining together to find the names of all instructors from the physics department. since the result of all operations on a relation is a relation, we can use them inside the parentheses.

$\Pi_{name}(\sigma_{department\_name="Physics"}(instructor))$

the `set intersection`, `natural join` and `assignment` operators don't add power over the fundamental operators, but make our life easier.

Extended relation algebra operations:

have the format of symbol, list of arguments which are attributes or constants, and then the relations.

we use them for aggregation, with the calligraphic G denoting aggregation `\mathcal{G}` $\mathcal{G}$. the name of the aggregation is written as it (sum, max, min, count, etc...). we can denote grouping by prefixing the symbol with the attribute we want to group on.

### Chapter 7 - Database Design and the E-R Model

E-R -> Entity Relationship, focusing on how to choose a schema. which entities are present, what are the main attributes and what relations they have to each other. specifying the required operations on the data.\
entities are the "things" the system deals with - people, places, products, etc...

_Redundancy_ is bad for schema design. data should only appear once, and should not be replicated across entities. if it is replicated, the different representations of the data may become inconsistent across time. _Incompleteness_ is another product of bad design. we might end up creating workaround items that belong to one entity and have missing data, because the entity actually covers more than one distinct "thing".

the entity-relationship model has three concepts:

- entity sets
- relationship sets
- attributes

if an entity is a distinct "thing", the entity set is a collection of those different things together, while they still belong to same category of "things". entities have attributes, so while all the entities in set have the same attributes, the value of those attributes differs.

a relationship is an association between entities. the entities are "participants" in the relationship, and we can define them as having a "role", relationships can also have attributes.

attributes themselves have domains (value sets), or possible allowed values. this can numbers, dates, text, etc.\
we can have simple attributes such as name or address, or composite attributes - first and last name, street, city, state and zip code. in some cases this makes sense, and in some cases it doesn't. (it makes sense to query for all students from a certain state, but less so to query for all student with the same first name). attributes can also be _single-valued_ or _multi-valued_. an id is single value, any entity can only have one. but a phone number is a multi-valued attribute, since a single person can have multiple numbers. a derived attribute is something that we can infer from from other relationships or entities, such as deriving someone's age from the date of birth. we don't store the derived values, we compute them instead.

a null value can mean the data is missing for the attribute (unknown) or doesn't exist at all.

schemas have constrains, such as relationships mapping cardinality - how many times can an entity be in the role in the relationship set.

- one to one
- one to many
- many to one
- many to many

participation can be _total_ if every entity from the entity set participates in a relationship (every person has at least one person as a parent), or _partial_ if some entities don't participate in it (not every person owns a car).

entities are different from one another because of their "keys", special attributes that must uniquely identify the entity from all other entities. the key can be one or more attributes, as long as the set of the attributes is unique. this is also true for relationships, they also have some unique way to identify them from one another.

- super key
- candidate key
- primary key

when we define the schema, we remove all the attributes which aren't unique, that means that we also remove attributes which refer to other entities, since those are conceptually relations, rather than true attributes.
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

[cspell dictionaries](https://cspell.org/docs/dictionaries/)

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
