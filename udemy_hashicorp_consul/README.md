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

<details>
<summary>
Consul Architecture Introduction.
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

<details>
<summary>
Node states and elections
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

A node (consul server agent) can be either the leader, a follower or a candidate. a node starts as a follower, it can vote for itself or another node to be a leader. when in election stage, nodes can become candidates. so this is a very short phase.\
Only the leader can ingest and write logs (updates), it is responsible for processing queries and transactions, and it replicates them to the followers, and after sending them to the table, it can determine that a log was "committed".\
Followers accept logs from the leader, participate in leader election, and forward the RPC request to the leader.

once a leader is established, it sends out heartbeat to all follower nodes. if a follower node doesn't receive a heartbeat during their randomly assigned timeout, it assumes the leader is dead and starts an election by changing it's state to be a candidate. the candidate votes for itself, and issues a request to establish the majority.

> Raft nodes are always in one of three states:
> - Leader
> - Follower
> - Candidate (short time)
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
>
> Consensus Protocol - Leader Election
> - Leadership is based on randomized election timeouts
> - Leader sends out frequent heartbeats to follower nodes
> - Each server has a randomly assigned timeout (e.g., 150ms - 300ms)
> - If a heartbeat isn't received from the leader, an election takes place
> - The node changes its state to candidate, votes for itself, and issues a request for votes to establish majority

when we make an API calls, it gets to the leader, which performs the request (writing the log) and tells the followers to replicate the change.

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

<details>
<summary>
Messaging, Service Discovery through gossip pools.
</summary>

The gossip protocol is used for communication for both consul servers and clients. it's manages the cluster membership, it handles broadcast messages, and and it can work across dataCenters, traffic inside the dataCenter uses the LAN gossip pool, and traffic between dataCenters uses the WAN gossip pool.

> - Based on Serf
>   - Used cluster wide – including multi-cluster
>   - Used by clients and servers
> - Responsible for:
>   - Manage membership of the cluster (clients and servers)
>   - Broadcast messages to the cluster such as connectivity failures
>   - Allows reliable and fast broadcasts across datacenters
>   - Makes use of two different gossip pools
>       - LAN
>       - WAN

each dataCenter has it's own LAN gossip pool, containing all the members (client and servers), and it allows clients to quickly discover servers. it also offloads failure detection duties from the server to the clients, which reduces the server load. members of the same LAN gossip pool can communicate with one another quickly.

> LAN Gossip Pool
> 
> - Each datacenter has its own LAN gossip pool
> - Contains all members of the datacenter (clients & servers)
> - Purpose
>   - Membership information allows clients to discover servers
>   - Failure detection duties are shared by members of the entire cluster
> - Reliable and fast event broadcasts gossip Protocol

The WAN Gossip pool is a global, unique pool, it includes all the servers, regardless of which dataCenter the belong to. server nodes can perform cross dataCenter requests, so we can handle failures at the dataCenter level.

> WAN Gossip Pool
> 
> - Separate, globally unique pool
> - All servers participate in the WAN pool regardless of datacenter
> - Purpose
>   - Allows servers to perform cross datacenter requests
>   - Assists with handling single server or entire datacenter failures


</details>

### Network Traffic and Ports

<details>
<summary>
Traffic protocols and ports.
</summary>

communication between consul agents (servers and client) goes over http and https protocols, this is secured by TLS certificate and the *gossip key*.

we need to open the ports and configure them correctly. DNS usually uses port 53, so this might cause problems, especially with windows machine.

> - All communication happens over http and https.
> - Network communication protected by TLS and gossip key.
> - Ports (assumes default):
>   - HTTP API and UI – tcp/8500
>   - LAN Gossip – tcp & udp/8301
>   - WAN Gossip – tcp & udp/8302
>   - RPC – tcp/8300
>   - DNS – tcp/8600
>   - Sidecar Proxy – 21000 - 21255

we can access (send requests) to consul from just about anywhere in the network. the CLI can be run from any server node (since it forwards to the leader). we can also enable a UI interface and manage through it, but it needs to be in our network. 

> - Consul API can be accessed by any machine (assuming network/firewall)
> - Consul CLI can be accessed and configured from any server node
> - UI can be enabled in the configuration file and accessed from anywhere

</details>

### Consul High Availability

<details>
<summary>
Configuring Highly available cluster, and some advanced features.
</summary>

Consul should be clustered, and provide High Availability. we should strive to have either 3 or five consul servers in the cluster, seven max. running a cluster with one server is not for production environments. it's better to have an odd number of servers, which allows to lose one extra server before we no longer have quorum.

> High availability is achieved using clustering
> 
> - HashiCorp recommends 3-5 servers in a Consul cluster
> - Uses the Consensus protocol to establish a cluster leader
> - If a leader becomes unavailable, a new leader is elected
> 
> General recommendation is to not exceed (7) server nodes
> 
> - Consul generates a lot of traffic for replication
> - More than 7 servers may be negatively impacted by the network or negatively impact the network

fault tolerance is the number of server we can lose before we don't have a quorum anymore. it increases in odd numbers, so that's why they are preferred over even numbered clusters.\
(there are some more stuff we can do with non-voting nodes).\
A single node should only be used for development, two nodes are practically useless, since losing one means the cluster is unusable. the more servers we have, the slower replication becomes, so going above 7 might effect performance.

> | Consul Server Nodes | Quorum Size | Fault Tolerance |
> | ------------------- | ----------- | --------------- |
> | 1                   | 1           | 0               |
> | 2                   | 2           | 0               |
> | 3                   | 2           | 1               |
> | 4                   | 3           | 1               |
> | 5                   | 3           | 2               |
> | 6                   | 4           | 2               |
> | 7                   | 4           | 3               |

#### Scaling for Performance

for enterprise edition only - with enhanced read scalability we can have read replicas that only handle reads from clients, which takes some work off the leader node. these nodes participate in data replication, but not in quorum elections (they are non-voting members).

> Consul Enterprise supports Enhanced Read Scalability with Read Replicas.
> 
> - Scale your cluster to include read replicas to scale reads
> - Read replicas participate in cluster replication
> - They do NOT take part in quorum election operations (non-voting)

#### Voting vs. Non-Voting Servers

when we used read-replicas, we created non-voting members,

```sh
consul operator raft list-peers
```

> - Consul servers can be provisioned to provide read scalability.
> - Non-voting do not participate in the raft quorum (voting).
> - Generally used in conjunction with redundancy zones.
>
> Configured using:
> - `read_replica` (previously `non_voting_member`) setting in the config file.
> - the `-read-replica`(previously`–non-voting-member`) flag using the CLI.

#### Redundancy Zones

another way to get scaling and resilience, we can deploy only one voting member in the Redundancy Zones, so if the voting member fails, the non-voting member is promoted and we maintain the quorum. if the entire zone fails, a non-voter from a different zone is promoted (to maintain the quorum). this is a a more managed way of having read-replicas, with an eye out for resiliency

> - Provides both scaling and resiliency benefits by using non-voting servers
> - Each fault zone only has (1) voting member
>    - All others are non-voting members
>
> - If a voting member fails, a non-voting member in the same fault zone is promoted in order to maintain resiliency and maintain a quorum
> - If an entire availability zone fails, a non-voting member in a surviving fault zone is promoted to maintain a quorum
</details>

### Consul Autopilot

<details>
<summary>
Automatic Server management features.
</summary>

also an enterprise edition feature, helps with managing the cluster.

```sh
consul operator autopilot get-config
consul operator autopilot set-config -cleanup-dead-servers=false
```

> Built-in solution to assist with managing Consul nodes
>
> - Dead Server Cleanup
> - Server Stabilization
> - Redundancy Zone Tags
> - Automated Upgrades
> 
> Autopilot is on by default – disable features you don't want

remove dead server from the cluster once the replacement is up.

> Dead Server Cleanup
> 
> - Dead server cleanup will remove failed servers from the cluster once the  replacement comes online based on configurable threshold
> - Cleanup will also be initialized anytime a new server joins the cluster
> - Previously, it would take 72 hours to reap a failed server or it had to be done manually using consul force-leave. 

A new server must be healthy for some amount of time before it can act as fully pledged node. this protects against having nodes that are shaky and unstable mess up the election protocol.

> Server Stabilization
> 
> New Consul server nodes must be healthy for x amount of time 
before being promoted to a full, voting member.
> - Configurable time – default is 10 seconds

spread cluster members across tags, good for High Availability.

> Redundancy Zone Tags
> 
> - Ensure that Consul voting members will be spread across fault zones to always ensure high availability
> - Example: In AWS, you can create fault zones based upon Availability Zones

manage which servers can vote during consul server migration, nodes with the new version won't be able to vote until they are the majority, and after that point, nodes with the old version won't be able to vote until they are updated to the new version.

> Automated Upgrade Migrations
> 
> - New Consul Server version > current Consul Server version
> - Consul won't immediately promote newer servers as voting members
> - Number of 'new' nodes must match the number of 'old' nodes

</details>

</details>

## Objective 2: Deploy a Single DataCenter

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

## Takeaways

<details>
<summary>
Stuff worth remembering
</summary>


</details>

</details>
