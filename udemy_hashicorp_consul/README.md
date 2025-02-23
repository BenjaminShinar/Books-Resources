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
### Network Automation
### Service Configuration
### Basic Consul Architecture
### Consensus Protocol (Raft)
### Gossip Protocol (Serf)
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
