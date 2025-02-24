<!--
// cSpell:ignore Krausen
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# HashiCorp Certified: Consul Associate (w Hands-On Labs)

Udemy course [HashiCorp Certified: Consul Associate (w Hands-On Labs)](https://www.udemy.com/course/hashicorp-consul) by *Bryan Krausen*. [kodeKloud Lab](https://kodekloud.com/courses/lab-hashicorp-certified-consul-associate-certification/), [related github repository](https://github.com/btkrausen/hashicorp).


<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

teaching the basic of hashicorp consul and getting ready for the associate certification.

> Course Goals and Topics:
>
> - The core concepts of HashiCorp Consul.
> - Key protocols used in Consul.
> - Creating and managing a Consul cluster.
> - Discuss service discovery and service mesh.
> - Using Consul's KV store.
> - Securing Consul with ACLs.

the exam has 9 objective, each will get a section of the course, there will a mindMap and a quiz and the end of each section.

1. Explain Consul Architecture
2. Deploy a Single DataCenter
3. Register Services and Use Service Discovery 
4. Access the Consul Key/Value (KV)
5. Back up and Restore
6. Use Consul Service Mesh
7. Secure Agent Communication
8. Secure Services with Basic ACLs
9. Use Gossip Encryption

## Objective 1: Explain Consul Architecture

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

> sub objectives:
>
> - Objective 1a: Identify the components of Consul dataCenter, including agents and communication protocols
> - Objective 1b: Prepare Consul for high availability and performance
> - Objective 1c: Identify Consul's core functionality
> - Objective 1d: Differentiate agent roles

Hashicorp has a suit of tools, all are open source, some also have enterprise editions.

- terraform
- vault
- consul
- nomad
- boundary
- packer
- vagrant
- waypoint

> Cloud networking automation for dynamic infrastructure

fits together in a microservice based architecture, providing for communication and discovery. Consul provides Service Discovery, Service Segmentation and Service Configuration. it is application agnostic, platform agnostic (kubernetes, openshift, vmware) and code agnostic (including on-premises and cross clouds).\
Traditionally, applications are deployed in monoliths, there would a load balancer at the entrance, and scaling is done on the application level, and not for each service. micro-services are deployed individually, so they can scale independently of one another, and if one fails, it doesn't affect the rest of the services. however, this introduces some operational issues, the services need to communicate with one another, and if they are containers, they don't have a consistent IP address, as they are mostly ephemeral. we also want to control which services are able to communicate and with which other services. this is where consul comes in.

the core features of consul are:

> 1. Dynamic Service Registration
> 1. Service Discovery
> 1. Distributed Health Checks
> 1. Centralized K/V Storage
> 1. Access Control Lists
> 1. Segmentation of Services
> 1. Cross Cloud/Data Center Availability
> 1. HTTP API, UI, and CLI Interfaces


### Service Discovery

<details>
<summary>
Centralized Service Registry
</summary>

a centralized service registry, a single point of truth to know the address of service from a different kind, including knowing which service is health and which is not. this is very helpful since workload no longer have consistent location and can scale up and down very quickly. this also replaces the need of having load balancers, both in terms of connecting services and for health-checks.\
each service has a consul agent, which registers itself with consul, that's how consul can direct the services where to find other services (either by dns or api calls). we can also register external service, even if they don't have the consul agent running. when a workload scales, the new copies of the service register to consul, so there are more options, when workload scales down, the health check fails for workloads that no longer exist, and consul stops serving them.\
We can also use consul for identity-based authorization, this replaces Ip-based and fire-wall based security. instead of having many copies of workloads with different ip address and having to run and fix the firewall rules to accommodate the changes. with consul, we can specify which services can communicate with which other targets.\
Service discovery also works across multiple data centers and provide **Mesh Gateways**, even without direct networking connection.

 
> - Centralized Service Registry
>   - Single point of contact for services to communicate to other services
>   - Important for dynamic workloads (such as containers)
>   - Especially important for a microservice architecture
> - Reduction or elimination of load balancers to front-end services
>   - Frequently referred to as east/west traffic
> - Real-time health monitoring
>   - Distributed responsibility throughout the cluster
>   - Local agent performs query on services
>     - Node-level health checks
>     - Application-level health checks
> - Automate networking and security using identity-based authorization
>   - no more IP-based or firewall-based security

</details>

### Service Mesh


<details>
<summary>
Allowing and Blocking Service Communication.
</summary>

Service mesh provides secure communication between services, and we can also deny communication between them. uses secure **Mutual TLS**. makes use of sidecar service (such as Envoy) next to the registered service to direct traffic (in and outbound), the application doesn't need to be aware of the sidecar.\
We can control the access control (which services can establish communication) and restrict connections without firewall rules. the rules in consul are called *Intentions*, and when a service attempt to connect to a service that it's not supposed to, the connection is denied. Consul itself becomes a Certificate Authority.

> - Enables secure communication between services
>   - Integrated mTLS secures communication
>   - Uses sidecar architecture that is placed alongside the registered  service
>   - Sidecar (Envoy, etc.) transparently handles inbound/outbound connections
> - Defined access control for services
>   - Defines which service can establish connections to other service
</details>


### Network Automation

<details>
<summary>
Dynamic Load Balancing
</summary>

dynamic load balancing, only sending traffic to healthy instances of the services, can distribute traffic based on traffic-shaping policies. can also be extended by dedicated networking services. this can help us with failover if one machine fails (for example, if we have a services across Availability Zones or cloud providers).\
Traffic can be managed on the L7 layer (path based), we can do weighted traffic distribution, this is done by placing a traffic splitting policy. we also get metrics for our traffic from the Envoy proxy, and send them to a centralized location.

> - Dynamic load balancing among services
>   - Consul will only send traffic to healthy nodes & services
>   - Use traffic-shaping to influence how traffic is sent
> - Extensible through networking partners
>   - F5, nginx, haproxy, Envoy
> - Reduce downtime by using multi-cloud and failover for services
> - L7 traffic management based on your workloads and environment
>   - service failover, path-based routing, and traffic shifting capabilities
> - Increased L7 visibility between services
>   - View metrics such as connections, timeouts, open circuits, etc.

</details>

### Service Configuration

<details>
<summary>
Consul As Service Configuration (Key-Value Store)
</summary>

Consul acts as a service configuration store. it is a distributed key-value store replicated across all service instances. this is not a full feature data-store. this data is also managed by the ACL for restrictions. objects can be of any size, but can't be larger than 512Kb.\
as an example, we might store connection strings, application versions, table names and other data, this data is also available for services outside the Network (such as jenkins).

> - Consul provides a distributed K/V store 
> - All data is replicated across all Consul servers
>   - Can be used to store configuration and parameters
>   - It is NOT a full featured datastore (like <cloud>DynamoDB</cloud>)
> - Can be accessed by any agent (client or server)
>   - Accessed using the CLI, API, or Consul UI
>   - Make sure to enable ACLs to restrict access (Objective 8)
> - No restrictions on the type of object stored
> - Primary restriction is the object size - capped at 512 KB
> - Doesn't use a directory structure, although you can use / to organize your data within the KV store
>   - `/` is treated like any other character
>   - This is different than Vault where `/` signifies a path

</details>

### Basic Consul Architecture

<details>
<summary>
Types of Consul Agent Modes
</summary>

Consul can be run as an agent, a long running daemon that manages the consul service itself. this is architecture, platform and cloud agnostic. it is not limited to running on a specific type of machine, and can run on a virtual machine, inside a container or directly on linux, MacOs or Windows.\
The consul agent can run either in server mode or a client mode. there is also dev mode, usually used locally for development and debugging. when running as a consul server mode, we sometimes call it "server agent", "consul server" or "consul node", when running the agent in client mode, we might call it "client agent". client agents run together with workload services, and they interact with the consul server agents. we should have a few server nodes, and a lot more of client agents.

The consul server manages the cluster state, the membership (which agents are in the cluster), it responds to Queries (dns and API) about services in the cluster. it also registers services (from the consul clients). consul servers maintain the state of the cluster quorum, and might act as a Gateway to other data centers.\
consul clients are deployed on workloads (a service for the application), they register the service with the consul server, and it performs health checks against the service or node that it's running on. all RPC calls are forwarded to servers. the clients are mostly stateless, and can be spun up quickly.
Dev mode shouldn't be used in production, it runs a non-secure (no tls) and non-scalable cluster, everything is stored in memory and nothing is preserved after closing.

> - Server
>   - Consul (cluster) State
>   - Membership
>   - Responds to Queries
>   - Registers Services
>   - Maintains Quorum
>   - Acts as Gateway to other DCs
> - Client
>   - Register Local Services
>   - Perform Health Checks
>   - Forwards RPC calls to Servers
>   - Takes Part in LAN Gossip Pool
>   - Relatively Stateless
> - Dev
>   - Used Only for Testing/Demo
>   - Runs as a Consul Server
>   - Not Secure or Scalable
>   - Runs Locally
>   - Stores Everything in Memory
>   - Does Not Write to Disk

Consul defines a DataCenter as a combination of consul servers and clients, inside a single physical or regional location (when on a cloud provider). Multi DataCenters are deployed across cloud providers, regions or physical locations. the use WAN gossip pools (each dataCenter has it's own LAN gossip pool). joined clusters are called "Federated". if there isn't a built-in connection and everything is going on through the public internet, we could deploy consul mesh gateways.

> What Is a DataCenter?
> 
> - single-cluster
> - private
> - low latency
> - high bandwidth
> - contained in a single location
> - multi-AZ is acceptable
> - uses the LAN gossip pool
> 
> **What a DataCenter Is Not!**
> 
> - multi-cloud or location
> - multiple Consul clusters
> - uses the WAN gossip pool
> - communicates via WAN or Internet
> 
> What Is Multi-DataCenter?
> 
> - multi-cloud, multi-region, location, or cluster
> - multiple Consul cluster federation
> - uses the WAN gossip pool
> - communicates via WAN or Internet
> - WAN federation through mesh gateways

consul uses two protocols:

- consensus protocol **Raft** - only servers agents, cluster operations, quorum.
- gossip protocol **Serf** - server and client agents, manage membership, broadcast messages. LAN and WAN pools.

</details>

### Consensus Protocol (Raft)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[RAFT protocol](https://raft.github.io/)

runs only on server agents. maintains strongly consistent data across all nodes. each cluster has a single leader node. for a consul cluster to function, it must have quorum.

> Based on Raft Protocol
> - Used on only Server nodes (cluster) - not clients
> - Strongly consistent
> 
> Responsible for:
>
> - Leadership elections
> - Maintaining committed log entries across server nodes
> - Establishing a quorum

A node (consul server agent) can be either the leader, a follower or a candidate.

> Raft nodes are always in one of three states:
> - Leader
> - Follower
> - Candidate
> 
>  - Leader is responsible for:
>   - Ingesting new log entries
>   - Processing all queries and transactions
>   - Replicating to followers
>   - Determining when an entry is considered committed
> - Follower is responsible for:
>   - Forwarding RPC request to the leader
>   - Accepting logs from the leader
>   - Casting votes for leader election


> Consensus Protocol - Leader Election
> - Leadership is based on randomized election timeouts
> - Leader sends out frequent heartbeats to follower nodes
> - Each server has a randomly assigned timeout (e.g., 150ms - 300ms)
> - If a heartbeat isn't received from the leader, an election takes place
> - The node changes its state to candidate, votes for itself, and issues a request for votes to establish majority

a Log is an ordered sequence of entries, a changeset of events (changes to the cluster or the key value store) from the onset of the cluster.\
A peerSet is all the members who maintain the replicated logs, which in our case, means all the consul server agents. a quorum is the number of nodes required for the cluster to function. 

> - Log
>   - Primary unit of work - an ordered sequence of entries
>   - Entries can be a cluster change, key/value changes, etc.
>   - All members must agree on the entries and their order to be considered a consistent log
> - Peer Set
>   - All members participating in log replication
>   - In Consul's case, all servers nodes in the local dataCenter
> - Quorum
>   - Majority of members of the peer set (servers)
>   - No quorum = no Consul
>   - A quorum requires at least $(n+1)/2$ members, with `n` being the number of nodes.
>       - Five-node cluster = $(5+1)/2 = 3$
>       - Three-node cluster = $(3+1)/2 = 2$


</details>

### Gossip Protocol (Serf)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Network Traffic and Ports
### Consul High Availability
### Scaling for Performance
### Voting vs. Non-Voting Servers
### Redundancy Zones
### Consul Autopilot
### Objective 1 - Section Recap
### Objective 1 - Practice Questions


</details>


## Takeaways

<details>
<summary>
Stuff worth remembering
</summary>


</details>

</details>
