<!--
// cSpell:ignore Upserts xattr
 -->

# Elasticsearch Complete Guide

udemy course [Elasticsearch Complete Guide](https://www.udemy.com/course/elasticsearch-complete-guide/). [github](https://github.com/codingexplained/complete-guide-to-elasticsearch). by *Bo Andersen*.

> Learn Elasticsearch from scratch and begin learning the ELK stack (Elasticsearch, Logstash & Kibana) and Elastic Stack.

1. Introduction
2. Getting Started
3. Managing Documents
4. Mapping & Analysis
5. Introduction To Searching
6. Term Level Queries
7. Full Text Queries
8. Adding Boolean Logic to Queries
9. Joining Queries
10. Controlling Query Results
11. Aggregations
12. Improving Search Results

## Section - 1 - Introduction
<details>
<summary>
Understanding Elastic Search and the realted stack.
</summary>


### Introduction to Elasticsearch

Analtics and full text search tool. open source. used in many websites, auto completion, relevancy, filtering, sorting...

it can also be used as an analtical tool, for structuring queries and creating reports,

APM - application Performance Management

we can also send events to elasticSearch, and then aggregate the results. Elasticsearch also works well with machine learning and forecasting. we can also use Elasticsearch as an anomaly detection tool, which we get for free without complicated manual set-up.

Elasticsearch stores data as documents, each document has fields. documents are json files. querying Elasticsearch is done via the rest API.

it itself is built on Java, on top of apache server, it's scalable and highly available, and is used by many companies.

### Overview of the Elastic Stack

some technologies related to Elasticsearch:

- Kibana
- Logstash
- X-pack
- Beats

**Kibana** is an analytics and visualization platform, it acts like a dashboard for creating reports and visualizations, it provides an interface to Elasticsearch for authentication, and can be used as a web interface. it also allows us to set the machine learning features.

**Logstash** is a data processing pipeline, it receives data as events, and handles them. each of the three stages makes use of plugins:
- input
- filter
- output

Logstach uses a markup language to define pipelines.

**X-Pack** - addition functionality to Elasticsearch and Kibana
- Security - authentication and authorization, roles, permissions
- Monitoring - how the stack is perfroming
- Alerting - notification services
- Reporting - exporting Kibana visualizations and data
- Machine Learning - abnormality detection, forecasting
- Graph - analyzing relationship in the data
- Elasticsearch SQL - a way to work with Elasticsearch with SQL instead of the internal DSL query.

**Beats** - a collection of lightweight agents that collect data and send to Elasticsearch. they can contain specialized modules for popular configurations.
- filebeat - collect log files
- metricbeat - system and service metrics
- packetbeat - network data
- winlogbeat - windows event logs
- auditbeat - audit data from linux
- heartbeat - monitor service uptime


the four services together form the elastic stack, and the term ELK - elasticsearch, Logstash and Kibana is commonly used to describe the collection of features.

### Walkthrough of common architectures

a theoretical use case:

E-Commerce websites, databases aren't great at full-search text. so we can use elastic search instead. the data is stored in both the Database and elasticsearch.

we next add Kibana to provide a dashboard, so we run it on a dedicated machine.

we next add monitoring, to know if we need to scale up. we then add metricbeat agent to the web servers, and we use this data to create alerts and notifications. metricbeat has a default kibana configuration.

now want to monitor the access and error logs, so we add a filebeat agent to get the logs into elasticsearch. and we want to enrich our logs and events, so we introduce logstash as a pipeline between our application server and elasticsearch.

### Guidelines for the course Q&A


</details>

## Section - 2 - Getting Started
<details>
<summary>

</summary>

### Overview of installation options
we would like to install elasticSearch and Kibana. we can either it directly or use a container. there is also a trial version.

### Running Elasticsearch & Kibana in Elastic Cloud


in [elasticCloud](https://www.elastic.co/cloud/cloud-trial-overview?medium=email&campaign=marketo-fallback), there is an option for a trial version, and we get a managed solution for 14 days.

the trial period begins we when launch a deployment, we choose the basic stack. we don't care which cloud provider we use (Azure, GCP, AWS).

when we create the deployment, we get a username (elastic), and the password.

we need the elasticSearch end point links, for both the engine and the kibana dashboard.

### Setting up Elasticsearch & Kibana on macOS & Linux

for macOS, we need to download and extract two archives. elasticSearch uses openJVM, while kibana uses nodeJS as a runtime.

```sh
tar -zxf archive.tar.gz
```

once extracted, we navigate to thefolders and run the commands. when we run the elasticsearch command, a superuser is created, as well as TLS certificates, we need the enrollment token for kibana.

```sh
cd elasticSearch
bin/elasticsearch
```

in a new working terminal, we need to start kibana, there is an extra command for macos. when we follow the kibana link, we are prompted to paste the enrollment token.

```sh
xattr -d  -r come.apple.quarantine path/to/kibana # MacOs only
cd kibana
bin/kibana
```

we authenticate ourselves using the super user (elastic) and the generated password.

### Setting up Elasticsearch & Kibana on Windows

we extract the elasticSearch and Kibana zipped files into folers.\
we navigate to the folders and run the scripts. we take the password and the enrollment token.
```ps
cd elasticSearch
bin\elasticSearch.nat
```

and for kibana, we run the script, go to the web browesr and enter the enrollment token, then fill in the password for the super users (user name is "elastic" by default).

```ps
cd kibana
bin\kibana.bat
```
### Understanding the basic architecture
### Inspecting the cluster
### Sending queries with cURL
### Sharding and scalability
#### Quiz 2: Sharding
### Understanding replication
### Quiz 3: Replication
### Adding more nodes to the cluster
### Overview of node roles
### Wrap up

</details>

## Section - 3 - Managing Documents
<details>
<summary>

</summary>

###  Creating & deleting indices
###  Indexing documents
###  Retrieving documents by ID
###  Updating documents
###  Scripted updates
###  Upserts
###  Replacing documents
###  Deleting documents
###  Understanding routing
###  How Elasticsearch reads data
###  How Elasticsearch writes data
###  Understanding document versioning
###  Optimistic concurrency control
###  Update by query
###  Delete by query
###  Batch processing
###  Importing data with cURL
###  Wrap up

</details>

## Section - 4 - Mapping & Analysis
## Section - 5 - Introduction To Searching
## Section - 6 - Term Level Queries
## Section - 7 - Full Text Queries
## Section - 8 - Adding Boolean Logic to Queries
## Section - 9 - Joining Queries
## Section - 10 - Controlling Query Results
## Section - 11 - Aggregations
## Section - 12 - Improving Search Results

## Takeaways
<!-- <details> -->
<summary>
Stuff worth remembering.
</summary>


### Docker setup

[docker instructions](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)


</details>
