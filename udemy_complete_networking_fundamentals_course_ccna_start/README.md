<!--
// cSpell:ignore Crisci datagram Netflow IPFIX HSRP VRRP nslookup NGFW Nord subnetting classful VLSM EIGRP OSRF IGPS EGPS OSPF
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# The Complete Networking Fundamentals Course. Your CCNA Start.

Udemy course [The Complete Networking Fundamentals Course. Your CCNA start](https://www.udemy.com//course/complete-networking-fundamentals-course-ccna-start) by *David Bombal*.

not doing the entire course, it's 70+ hours.

starting with focus on:

1. NAT - network address translation
2. Dynamic Routing Protocol (BGP - border gateway protocol)
3. lifetime of a package
4. WAN
5. VPN

## Welcome to the Course

<!-- <details> -->
<summary>
Introduction stuff.
</summary>

### Network Basics: What is a Network?

<details>
<summary>
Basic definitions.
</summary>

> What is a Network? (wikipedia definition)
>
> - A computer network is a set of computers sharing resources located on or provided by network nodes.
> - Computers use common communication protocols over digital interconnections to communicate with each other.
> - These interconnections are made up of telecommunication network technologies based on physically wired, optical, and wireless radio-frequency methods that may be arranged in a variety of network topologies.
> - The nodes of a computer network can include personal computers, servers, networking hardware, or other specialized or general-purpose hosts. 
> - They are identified by network addresses and may have hostnames. Hostnames serve as memorable labels for the nodes and are rarely changed after initial assignment.
> - Network addresses serve for locating and identifying the nodes by communication protocols such as the Internet Protocol.

resources can be files, videos, i/o. nodes is anything connected to the network. a host is a client connected to the network via some connection (wired, wifi, optical fiber, etc...). there are also "servers", which provide services to the network itself, and are providing resources for the network. a device can be both a server and a client.\
the server exposes the service on a port with a specific protocol, and clients connect to it by specifying the correct ip address and port.

a network is any two devices connected between them, it can be two laptops connected with ethernet cables, or two phones connected with bluetooth or wifi.

MAC addresses are 48 bit addresses. hard coded into the network card.
</details>

### What is a Switch? A Router? What network is this? And what are these?

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

physical and logical topologies,

</details>

</details>

## WAN - Wide Area Networks (Point to Point)

<details>
<summary>
Wide Area Networks.
</summary>

layer 2 encapsulations:
> - ppp - point to point protocol
> - HDLC -high level data link control

connecting devices separated geographically.

ppp is also called a serial link, or "leased lines", is a dedicated connection between two sites using the service provider network (ISP). devices are far apart from one another and can't be connected with a cable.

today many WANs were replaced by VPNs, which go over the public internet and don't require leasing connections from the provider.

### WAN Point-to-point link

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

a pc in one local area network sends data to a computer in a different local area network.\
when the packet reaches the router, it strips away the layer 2 headers, and wraps it  with HDLC headers and send it to other router, the receiving router re-wraps the package it with ethernet headers.

we care about the Layer 1 (physical) and Layer 2 (data link).

- CSU - channel service unit
- DCU - data service unit
- ethernet WAN connections

</details>

### CSU DSU

a physical device connected to the router on one side, and to the telecom provider on the other. today it can be part of the router.

### DTE / DCE

two devices synchronize the clocks between sending data, and also agree on the data interval between sends.

- DCE - Data Communication Equipment
- DTE - Data Terminal Equipment

the router is the DTE, and it gets the clocking from the CSU/DSU device - the DCE.

### Serial interfaces (WAN Interface Cards)

different serial cards, such as WIC 1T (wireless interface card) or WIC 2T. the number indicates how many WAN interfaces there are on the card.

### Serial Interface Speeds

we can get different speeds and bandwidth from the ISP. there are standards for which speeds can be taken. the cost is for the bandwidth, no matter if it's used or not.

</details>


## VPN - Virtual Private Networks

<details>
<summary>
Virtual Private Networks.
</summary>

secure access connection across the public internet, replace WAN for connecting to corporate resources. encrypted connection between private networks over the public internet. 

ip transmits in clear text, if captured, it can easily be read by a 3rd party agents. many protocols are clear text, including authentication and content.

- FTP
- Telnet
- SMTP
- HTTP
- SNMP v1

symmetric algorithms use the same key encryption and decryption, asymmetric algorithms have different keys (private and public).

we want the data to be confidential, we want to know it wasn't changed in the middle, that it comes from the real sender, and that packages can't be duplicated.

the algorithm has a key space /key-length property, which is the set of all possible values. or 2 to the power of the number of bits in the key (minus overhead).

AES is a symmetric key algorithm, the problem is making sure both sides have the same key, and the key must be transferred out-of-bound, because there still isn't a secure channel. they are fast, secure and easy to implement.

DES is 56bit length key, which is considered too small and susceptible to brute force attacks. 3 DES is another algorithm, first encrypt with key 1, decrypt with key 2, and encrypt again with key 3. so reading the data requires doing things in reverse, first decrypting with key 3, encrypt with key 2, and decrypt with key 1. AES is the recommended algorithm today, it comes in different variants (128, 192 and 256) of key-space length.\
RSA is an asymmetrical algorithm, the key to decrypt isn't the same as the key to encrypt. the key length is longer than symmetrical keys. one side generates a private key, and from that it generates a public key, the public key can be shared with others.

Diffie and Hellman discovered a way to share keys across public network. something about shared secrets. this allows us to create a symmetrical key for VPN. this also comes in different variant (key lengths).

we want data integrity, to be sure it wasn't tampered with and didn't change. there are some fixed length hashing algorithms (MD5, SHA), the algorithm is non-reversible, so if the sha in the message matches the sha calculated from the data, we can be sure it wasn't changed. the hash check is performed before decrypting, since we won't waster time doing an expensive decryption. there is an extra step of adding an a secret key for the hashing, to prevent someone completely replacing the message and adding a new sha.

the next part is authenticating the data is really coming from who it says it's coming from. there are some steps for this. either using pre-shared keys or digital signatures.

we still need a mechanism to know the public key really comes from the correct place. this requires a certificate of authority, a trusted signer will ensure the public key really belongs to the person who claims there are theirs.

IPSec is a network layer protocol to protect and authenticate VPN, it provides internet key exchange, authentication headers, and payload encapsulation. there are two modes: transport and tunnel mode.

VPN can be site-to-site (replacing point-to-point wan) or remote-access VPN (also something about SSL access).
</details>

## Routing

<details>
<summary>
Introduction to Routing
</summary>

### Introduction To Ip Routing And Routed Vs. Routing Protocols

routes vs routing protocols, static and dynamic protocols. distance vector and linked state protocols.

Routed protocols carry user information. every router in the path makes an independent router decision.

hop-by-hop routing paradigm. unicast packets based on destination address only.

```sh
tracert -w 50 www.facebook.com
```

routing protocols:

- EIGRP
- OSRF
- RIP
- ISIS
- BGP

each protocol makes decision differently to decide on routing, RIP uses hop count, OSPF uses bandwidth, EIGRP uses bandwidth + delay...

if a router doesn't know about an address (not in the routing protocol), then it will drop the unicast packet.

IPv4 and IPv6 are routed protocols, and are independent from one another. "ships in the night" - what one does isn't seen by others.

routed protocol -> routing protocol

static routes are added manually (no overhead on network), dynamic routing consume bandwidth (keep alive messages) but don't have operational overhead and are scalable.\
there usually is a default static route to the ISP provider (default gateway).

### Static Routes And Dynamic Routes

```sh
telnet route-server.ip.att.net
```

when we enable a dynamic routing protocols, the router will exchange information with other routers to update the routing table. this is how the can adjust to topology changes.
### How Do Routers Determine The Best Route?
determining the best path to a destination, RIP uses the hops count, OSPF would consider the bandwidth when determining paths.

### Terms: As, Igps, Egps
- AS - autonomous system
- IGP - internal gateway protocol, inside an autonomous system, such as RIP, EIGRP, OSPF
- EGP - external gateway protocol, between autonomous systems, like BGP

autonomous systems are registered number when connecting to the internet (no need to register when not connecting to the internet), we might get an AS number from the provider.

### Types Of Routing Protocols - Distance Vector, Link State

distance vector - determine the direction and distance to destination. but limited visibility, might not have all information. requires routers to publish their route table.

linked state routing protocols have the entire visibility of the network, and can calculate all routes, OSPF is one such protocol. each router has a complete copy of the topology. they use the "shortest path first" algorithm (by dijkstra). they are more difficult to configure and require more memory and processing power.

### EIGRP - Advanced Distance Vector Routing Protocol And Administrative Distance

Enhanced Interior gateway routing protocol

a hybrid (advanced) routing protocol, combines both distance vector and link state routing. easy to configure like distance vector, but also has neighbor relationships

cisco propriety protocol.

AD - administrative distance, the believability of a route, which route is more trustworthy. this acts as a tie-breaker between different routing protocols. (lower = more trust). the router internal network has the lowest value (1), static routes have the next value (1), unless they are directly connected, and then the also have AD of zero.


| Type                              | AD Value | Notes               |
| --------------------------------- | -------- | ------------------- |
| Connected interface               | 0        | max trust in self   |
| Static routes - direct connection | 0        | manually configured |
| Static routes                     | 1        | manually configured |
| External BGP                      | 20       |
| Internal EIGRP                    | 90       |
| OSRF                              | 110      |
| ISIS                              | 115      |
| RIP                               | 120      |
| Internal BGP                      | 200      |
| unknown                           | 255      |

### Classful Routing Protocols
classful don't advertize subnet mask, they say the subnet address, but not the the "/8" part, so it's hard to tell if two address are on the same subnet or not. classful protocols assumer the same subnet mask is used for all routers. this isn't viable today, so it's not used.

#### Auto Summarization
automatically summarize when moving across class boundaries, it thinks the neighboring routes use the same class definition.
### Classless Routing Protocols
classless routing protocols advertize the subnet mask, and don't need to assume anything. and they support VLSM. they also support manually configuration of summary routes.

### Administrative Distance Versus Mask Length
the length of the mask takes precedence over the administrative distance, a table can have multiple entries for subnets, each with different subnets that overlap one another, and each based on a different protocol with a different AD, but matching entry with the longest mask will be chosen.

for example, we want to ping 10.1.1.1, and we have 3 routes with AD values:

- 10.0.0.0/8 OSPF (110)
- 10.1.0.0/16 BGP (200)
- 10.1.1.0/27 RIPv2 (120)

our destination is inside all three subnets, but rather than choosing based on the AD value (OSPF is the most trust worthy), we first use the subnet mask (/27 is the longest-best match), with AD only acting as a tie-breaker.
#### Administrative Distance And Multiple Routes With The Same Mask

if there are multiple candidates for a route with the same subnet mask, then the AD value is used to choose which value goes into the routing table. since we can't have two entries for the same destination in the routing table.

### Link State Routing Protocols

Link state Protocols are usually better than AD protocols, since they have a complete visibility of the network, rather than just knowing a single route.

LSA - Link State Advertisement

They flood the network with LSA, either the entire network if inside a single area, or just within the area (if configured). all routers get the LSA, and use the data to populate a topological database. this database should be the same for all routers. this will contain the links, the state of the link, etc...

each router will run a SPF (shortest Path First) algorithm, with itself as the root, and then find the best path from them to the routes, and will then then populate the routing table.


#### OSPF Hierarchy - Multiple Areas

link state protocols allow 
for division of autonomous systems into areas. this allows for smaller routing tables at each router, less flooding, and route summarization.\
Routers can either be internal to the area, or be at the border of the area (ABR). The topmost router that connects the area itself to an external routing system is the Autonomous System Border Router.\
link changes in one area are hidden from routers in other areas, they don't need to run the algorithm again.

#### Benefits Of Link State Routing And Drawbacks Of Link State Routing

Benefits:

- Fast convergence - Changes are reported immediately
- Robustness again routing loops 
  - Routers know topology
  - Link state packets are sequenced and acknowledged
- Hierarchical Design enables optimization of resources
- Can scale to much larger environment than distance vector routing protocols

Drawbacks:

- Significant demand for resources
  - memory - Adjacency (neighbor table), topology, forwarding table
linked state routing protocols have the entire visibility of the network, and can calculate all routes, OSPF is one such protocol. each router has a complete copy of the topology. they use the "shortest path first" algorithm (by dijkstra). they are more difficult to configure and require more memory and processing power.
  - CPU - dijkstra's algorithm can be processor intensive
- requires strict design
- configuration and design can be complex

</Details>

### Static Routing

<details>
<summary>
Static routes configurations
</summary>

A routing table is the list of tables about networks, and how to reach to them (interface, next hop-router).

- directly connected networks - added automatically
- static routes - manually added to the router, don't adjust to changes, even if the destination router is down.
- default route - special kind of static route, or dynamic, used when no routes matches.

(some shell commands)

lots of static routes configuration, loop backs and all, bi-directional. needing to add so much stuff.

### Static Routing: Default Routes

gateway of last resort, static route to 0.0.0. with 0.0.0.0 mask, directed to the router at the next hop.

### Static Routing: Removing Static Routes

more shell commands.

</details>

## ACL - Access Control List

<details>
<summary>
Security using access control lists
</summary>

inbound and outbound ACLs, wildcard masks, sits on an router interface. used to permit or deny packets moving through the router, limit connections to private resources. also used for classification, encrypting traffic in an ipsec vpn tunnel. can be used for redistribution between routing protocols, together with NAT (which packets should be translated).

1. create access list
2. apply access list to an interface (inbound or outbound).

it's usually better to use inbound, since we won't have wasted work of route processing.

an acl is ordered, the first match determines if the packet is denied or permitted. uses implicit deny default.

- standard acl - only source ip address.
- extended acl - source and destination, ip, protocols, and ports (source and destination).

the number of the ACL determines the type, there are different ranges for standard, extended (more types). there are also named ACL.
ACLs have masks to ignore parts of the ip (like range), 

</details>

## NAT - Network Address Translation

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

PAT - Port Address Translation

NAT can be dynamic or static, translates between private and public ip addresses. 

the three blocks of private subnets:

- 10.0.0.0/8
- 17.16.0.0/12
- 192.168.0.0/16

we're basically out of ipv4 addresses, and we should have moved to ipv6 already

PAT allows us to translate one public ip to to 500 internal device.

- inside local - inside the network
- inside global - as seen by the internet
- outside local - how it's seen from inside
- outside global - how it's seen from outside.

static nat maps private to public ipv4,one to one mapping. dynamic maps private address from one pool to another pool, both networks have the same range, but dynamic NAT handles this. PAT maps multiple private into a single public ip address. uses port number, also called NAT overloading.
</details>


### Terms
- EIGRP - Enhanced Interior gateway routing protocol
- VLSM - Variable Length Subnet Mask
- BGP - Border Gateway Protocol
- LSA - Link State Advertisement
- SPF - Shortest Path First
- OSPF - Open Shortest Path First
- ABR - Area Border Routers
