<!--
// cSpell:ignore Krausen datagram Netflow IPFIX HSRP VRRP nslookup NGFW Nord
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# Introduction to Computer Networking - Beginner Crash Course

Udemy course [Introduction to Computer Networking - Beginner Crash Course](https://www.udemy.com/course/networkingbasics) by *Rick Crisci*.

> Learn the basics of IP, Ethernet, VPNs, WANs, DHCP, DNS, with no prior knowledge required. Bonus IP addressing course!

## Introduction

<details>
<summary>
The basic concepts.
</summary>

local area network, for computers and other devices. we can use a physical ethernet switch and connect the device through physical cables. we could also have wifi connecting the devices in the local network. in the past we had hubs, today we usually use switches. when the switch is in the center of the network, we call this a "star topology".
Each of the devices has a MAC (media access control) address, every port on every network card has a unique MAC address, hardcoded onto it.

we can see the mac address under the "physical address" section.

```sh
ifconfig -a # show all
```

if our devices in the network want to communicate with other devices, they send frames across the network with *ethernet frames*.

an ethernet frame has the structure of:

- payload
- destination mac address
- source mac address
  
</details>


## Basics of Layers 2, 3, 4

<details>
<summary>
Layers 2,3,4 communication basics.
</summary>

### LAN: Local Area Networks, including Hubs, Bridges, and Switches

<details>
<summary>
short intro to hubs, bridges and switches.
</summary>

unicast - one device sending data to a another, single device.

**HUB**

when a device in the local network wants to communicate with another device, it sends a unicast onto the network (the hub). the hub then sends the message to all ports. this is not efficient. and a hub can't handle more than one message at a time (collisions).\
to handle this, we add a *Layer 2 Bridge*, this device has a MAC address table, and it tracks which addresses are reachable through which port. when the bridge gets a message, it can check the destination address on the message, and only forward the message to the port that is connected to that address. this reduces collisions and makes the network more efficient.\
A switch also has a MAC table, but it replaces the hub entirely. there are no more collisions.

a message can also be a broadcast, which is intended to go to all devices on the network (not just to a single destination). on the switch there is special mac address for broadcasting, when the switch gets a message to that address, it floods the message back to all ports. this becomes a problem when we have too many devices attached to the switch (even more switches), since now we have a *broadcast domain*. this will be fixed by using a *router*.

</details>

### Understanding the OSI Model

<details>
<summary>
introducing the 7 layer model
</summary>

the layers are, from the bottom to the top.

1. Physical
2. Data Link
3. Network
4. Transport
5. Session
6. Presentation
7. Application

or the mnemonic "All People Seem To Need Data Processing"

this course won't focus much on the upper layers (7-5: application, presentation, session).

TCP is a layer 4 protocol. http works on TCP, so while it's an layer 7 protocol, it uses layer 4. the network layer 3 uses IP addresses, so our computer sends the message to a default gateway, to send out the message onto the local network (which has the gateway), we use layer 2 mac addresses.

| Layer   | address type            |
| ------- | ----------------------- |
| layer 4 | TCP/ UDP port           |
| Layer 3 | Source/ Destination IP  |
| Layer 2 | Source/ Destination MAC |
| Layer 1 | physical data flow      |

each layer adds it's own headers. the top levels (5-7) write the payload (the data we want to transmit), layer 4 appends to the port number based on the protocol. layer 3 adds the ip addressing, routers and routing protocols, and layer 2 adds the local network protocol (ethernet, wifi, fiber channel etc...) like the mac addressing.

</details>

### Layer 2 - Data Link

<details>
<summary>
Layer 2, MAC address and special kinds of traffic
</summary>

#### What Happens When a Computer Is Connected to a Switch?

a computer has a network interface card (NIC) with a hard coded MAC address. this card is connected to the ethernet switch at one of the ports. traffic from the computer is received by the switch, which creates an entry in the Mac table, this entry maps the relation between the mac address and the port. the entry is valid for a few minute, as is then removed, each time traffic flows, the entry is renewed.\
multiple MAC addresses can have the same port, this happens when we have chained switches.

#### BUM Traffic: Broadcast, Unknown Unicast, Multicast

other than unicast messages, there are some unique types of traffic that can happen in the local ethernet network.

BUM
- Broadcast
- Unknown unicast
- Multicast

let's look at an example, in our local network, computer A wants to communicate with computer B. each computer has both the physical MAC address, but also is assigned an IP address inside the network. the IP is unique inside the network, and most traffic is used based on ip addresses.\
if the sender knows only the ip address of the destination, and not the MAC address, it needs to first send out an ARP broadcast request to discover the MAC associated with the IP address. the switch gets the ARP message, and floods it out to any devices connected to it. devices that aren't associated with the IP address simply ignore the request, but a device that has the IP sends back an ARP response (unicast), then the original sending device can update the internal ARP table (mapping ips to MAC addresses) and send the original message it wanted to send.

```sh
arp -a # display arp table
```

another example is for the *Unknown Unicast*, this is when a switch doesn't know which port a MAC address is on. this can happen when we have switches chained together. Computer A sends a unicast message to Computer B, but they aren't on the same switch, so switch A isn't aware of the MAC address for Computer B. in this case, Switch a floods an "unknown unicast" to all the ports, and when it gets a response, it knows to associate the mac address with a port. this is called "unicast", but is still a type of broadcast message.\
The third type of special traffic is *Multicast* - which are sent to members of a group, one message to multiple destinations, but not all. the switch has a special address for each group, so it can send the source frame message through all those ports.

</details>

### Layer 3: What are Routers?

<details>
<summary>
Routers break Layer 2 networks into segments.
</summary>

we can connect multiple switches together, this makes our broadcast domain larger and larger, and also increases the size of the mac table, since each switch needs to hold all the mac addresses.\
Routers come and help us here. rather than connect switches directly, we connect them into a router, which separates them into segments (subnets). devices will have ip address inside the subnet range of the segment.\
if we send an ARP request from one machine it's flooded from the Switch to all connected devices. the router can inspect the message and decide if it should flood it out to which segment based on the ip range. this also reduces the size of the mac tables in each switch.\
the router acts as a default gateway, and it has it's own mac address. messages sent with an IP address outside the local subnet are directed to the local gateway MAC address.

we can see this default gateway address in the command line. both the mac address and ip address.

```sh
ifconfig -a
```

the router gets the message, and knows how to make routing decisions and how to forward it to the correct switch.

</details>

### Layer 4: TCP and UDP

<details>
<summary>
TCP and UDP protocols.
</summary>

TCP - handshake, error checking. UDP - no error checking.

UDP is connection-less, just send data, values speed over accuracy. use cases: Voice over IP, DNS, DHCP, TFTP. we call the data we send through UDP *datagram*.

TCP values accuracy over speed, it has error-checking, it first makes a connection between the two sides of the communication. used for web browser, file downloads, and email. we call the data sent over this protocol *segments*.\
Segments have a sequence number as a header, and for each segment, the server expects to receive an acknowledgement message with the same sequence number. if it's not acknowledged, it re-sends the segment. in addition to the sequence number, the headers also contain a checksum, which is calculated from the payload and the headers, and is used to detect corrupted packets.

</details>

</details>

## IP Addressing, Routing and VLAN Basics

<details>
<summary>
Some basics.
</summary>

### Breaking Down IP Addresses and Subnets

<details>
<summary>
Subnets and Ip Ranges
</summary>

ipv4 address are four number between 0-255 separated by a dot. or four bytes (8 bits).

in our example, we have subnets: `10.1.1.0`, `10.1.2.0` and `10.0.0.0`, subnets are actually a range of ip address, which are denoted with a `/` after the base address. the number after the slash is how many bytes are the network address, the remaining address are the "free" address in the range.

this is CIDR notation for ip ranges.

- `10.1.1.0/24` - has 8 bytes of address, or 256 addresses in the range [10.1.1.0-10.1.1.255]
- `10.1.2.0/24` - has 8 bytes of address, or 256 addresses in the range [10.1.2.0-10.2.1.255]
- `10.1.0.0/16` - has 16 bytes of address, or $256^2 = 65,536$ addresses in the range [10.1.0.0-10.1.255.255]

there are some reserved addresses in each subnet.

- the first address is the network address - `10.1.1.0`
- the next address is the default gateway address `10.1.1.1`
- the last address in the range is the broadcast address `10.1.1.255`

some providers take additional address for themselves (<cloud>AWS</cloud> takes additional two address at the end of the range)

</details>


### Packet Walk - Follow A Packet Through The Network!

<details>
<summary>
Follow a packet route.
</summary>

we return to our example of sending packets between two computers.

On computer A we start with the payload, and add the source IP and Destination IP (layer 3). Since the destination address is outside the local subnet range, the MAC address of the default router is added as layer 2 headers (source and destination).\
The switch sees the mac addresses, and forwards the frame to the router. the switch doesn't care about layer 2 headers.\
the router sees the MAC address destination is his own address, so it unpacks the L2 headers and removes them, and now it looks at the destination IPs and and determines which network to send it through the network interface. it appends the new layer2 headers, with itself as the source, and it uses the internal ARP table to map the destination IP with the MAC address.\
The frame arrives at the switch, looks at the layer 2 headers, and uses the mac table to send the message to the correct port.\
The receiving computer reads the L2 headers and detects that it's the destination, so it reads the L3 headers, and it's also the correct IP, so it can read the payload.\
(we ignored some L4 stuff)

</details>


### VLANs: Virtual Local Area Networks

<details>
<summary>
Multiple subnets on the same switch
</summary>

VLAN create multiple logical partitions inside a switch. isolation between devices using the same switch. we could do this with multiple switches connected to the same router (and set rules on the router). we force the traffic through the switch into the router and have it manage the permissions. we configure VLANs on switch and associate ports to VLAN networks, they also act as a broadcast domains. we turn one physical switch into multiple virtual ones.

- access ports belong into a single VLAN
- the switch connects to the router with a "trunk port".
- the router uses vlan interfaces for each vlan, and treats the networks as if they were separate from one another.

we can take create security rules on the router.

</details>


</details>


## The Internet, WANs and VPNs

<details>
<summary>
Internet Basics
</summary>

### WAN: Wide Area Networks

<details>
<summary>
Connecting distant networks
</summary>

LAN networks are in the same physical location, connected through hubs, switches and routers. but if our devices are spread across and don't have a direct connection, we can still have them communicate. each location has an *Edge Router*, those are connected with a WAN connection. this could be a T1 circuit, a fiber optic connection, or something else we get from the telecom company.\
The challenge is setting up the route tables, the routers need to direct traffic to segments which it isn't connected to. for this, we set the tables to use the WAN interface, this is a static route, which we need to manually configure. they aren't discovered automatically. we could also use a dynamic routing protocol, such as **BGP**. the default route is `0.0.0.0/0`, which encompasses all ipV4 address, and we use it as the last resort. traffic to this route is usually sent to the public internet.

</details>

### Connecting Your Network to the Internet

<details>
<summary>
Border Router and connecting to the ISP.
</summary>

connecting LAN network to the network.

the router is connected to another ethernet switch, which connects it to the *Border Router* - which is then connected to the ISP and the public internet. the area between the two routers is *DMZ*. the border router also controls incoming traffic (simple rules, not as complicated as a firewall).

> - Bandwidth - How much traffic can a connection handle?
> - Speed - How much latency

</details>

### VPN - Virtual Private Networks

<details>
<summary>
VPNs instead of WAN.
</summary>

#### IPSEC - Layer 2 VPN

- IPSEC - IP Security
- VPN - Virtual Private Network

VPNs come to replace WAN as a connection between physical locations. instead of putting down cables between locations, we can create an VPN using IPSEC. this is an interface on the router, we put it on the router with a public IP address. A VPN tunnel has the public ip of the two routers, and they use the same secret. The VPN encrypts the data that's being sent to that IP with the key, and the VPN on the destination secret can see that the traffic came from the known ip, and uses the secret key to decrypt the data.

#### Layer 2 VPN

VPNs using IPSEC are layer 3, but we can also have layer 2 VPNs. this will maintain the ip address scheme at both sites. the subnet is shared on the two networks, even if they aren't physically connected. both networks have the same ip ranges.\
Layer 2 VPNs still use tunnel between routers with public ips. when the computer creates the frame, the mac address it has in the ARP table (mapping ips to MAC addresses) will be that of the router, and the router will encrypt and send the data over the internet.

this is called "stretched layer 2 network", since we "stretch" the network across different sites. this allows us to migrate workloads between sites without changing the ip address.

</details>

### LTE/5G and and Its Impact on Networking

<details>
<summary>
Mobile Internet Connections.
</summary>

we can use mobile internet (5G or LTE) connections as backup connections for sites, or even use mobile connections entirely, without attaching cables.

#### Software Defined WAN

we can use 5G connection as the basis for a SD-WAN. in our example, we have a branch office with some devices, and this site connects to the internet through broadband cable connection, and has a backup 5G connection. we also have the HQ office dataCenter, and we are also using cloud services.\
the branch needs to send data to the public internet, to the cloud and to the main office. the SD-WAN holds both the broadband connection and the mobile connection. it creates VPNs to both HQ and the Cloud provider. the SD-WAN device needs to direct traffic to the correct VPN, and it can choose which connection will send which traffic. it can also prioritize traffic and send important traffic to the faster connection. the SD-WAN can identify traffic going to the internet and send it to the internet directly, or it could send all traffic through the VPN - this gives us the advantage of using security at a central location, on the other hand, we get higher latency and we might be putting unnecessary load on the HQ.

</details>

</details>

## Network Management, Design and Troubleshooting

<details>
<summary>
ReRedundancy, High Availability and some more stuff.
</summary>

### Understanding Network Redundancy

<details>
<summary>
Redundancy and Failover Protocols.
</summary>

keep our network functioning, even if something fails. we want to avoid having "single points of failures", any time we have one of something (NIC, cable, switch) it becomes a point of failure that can take the network done.\
When something fails, we need to have a failover process that directs the traffic to the other option. for routers, this would mean updating the route table and adjusting it to not send traffic if the connection is down. this is done by dynamic routing protocol, protocols such HSRP or VRRP allow routers to "share" ip addresses and have different priories. this ip address can be the default gateway, so having one router fail doesn't affect the devices in the network.

</details>

### Understanding Load Balancing

<details>
<summary>
High Availability through a load balancer.
</summary>

Load Balancing distributes traffic across services which are interchangeable, this gives us High Availability and can allow us to scale work on many weak machines instead of having a single strong machine.\
in our example, we have a dns record which points traffic to a load balancer, the load balancer has listener, which is defined for a protocol and port, such as http on port 80. the traffic is then directed to one of several identical web servers, which can be spread across different locations and dataCenters.\
the load balancer also performs health check on the web-servers and make sure they are still operating, and can know to not send traffic to unhealthy web servers.\
some cloud providers can replace unhealthy servers, depending on the configuration.

</details>

### Basic Network Troubleshooting Methodology

<details>
<summary>
Common problems and solutions.
</summary>

isolating problems, eliminating potentials causes. we need to know the network diagram, and find where the problem exists. we reduce the area to search for the problem in.

#### Netflow and IPFIX

trouble shooting network issues.
Netflow is a network protocol developed by cisco. it gives viability about traffic. IPFIX is a non-cisco version of the same thing.\
we can set a special address on the router and send summaries of traffic to a netflow collector, this is the history of traffic in our network, and we can query it. it will also show us unusual traffic.\
netflow records
- source and destination ip
- source and destination ports
- protocol used
- number of bytes send and received
- timestamp

</details>

### Protocols

<details>
<summary>
Focusing on some protocols.
</summary>

some common protocols.

#### DHCP: Dynamic Host Control Protocol

DHCP is one way to assign IP addresses. we could configure static ip manually, we could put the ip address, the subnet mask and the default gateway, and also set up the DNS server.

To get an IP address, the computer sends a layer 2 `DHCP broadcast`, and the server responds with an `DHCP Offer`, which the computer might accept, which the server will acknowledge. we can have a dedicated dhcp server, or run it on the router.\
we can see the dhcp when looking at the output. we can release the ip address or renew it.

```sh
ifconfig -a
ifconfig release
ifconfig renew
```

#### DNS: Domain Name System

using human-readable names instead of public ip addresses. we give it FQDN - fully qualified domain name, which is the internet address, and it goes to the DNS server and resolves the address. this is done by a chain of DNS servers, starting from the locally configured dns and up to the root DNS server.

we can check the responses with the CLI, this will give us ip addresses.

```sh
nslookup # who is the dns server
nslookup www.example.com # lookup address
```

#### Network Time Protocol (NTP)

the network time protocol gives us an authoritative time source for all devices in the network, we use it to align the time in all devices. this is also important for digital certificate.

NTP traffic flows over UDP port 123. we can get the time from several different sources, and we want all the devices to use the same source. so we choose one device to act as the NTP server authority, it will be the only device to pull data from the external source, and all other devices will pull the time from it.

</details>

</details>

## Network Security

<!-- <details> -->
<summary>
Basic network security and firewall.
</summary>


### NAT, Public IP Addresses, and Private IP Addresses

<details>
<summary>
NAT - network address translation.
</summary>

public ips work for anywhere on the internet. private ips aren't globally unique, and they only make sense in the context of a private network.

there are some defined private network addresses, there are unusable on the internet, and can't be used on the public internet.

- `10.0.0.0` - `10.255.255.255`
- `172.16.0.0` - `172.31.255.255`
- `192.168.0.0` - `192.168.255.255`

instead, they only make sense in the confines of the private network, there are more devices in the world than ipv4 addresses.

we still get internet access, because our router has NAT capability, and can translate ip addresses. it can substitute the source ip of the private ip with that of the router, and will also do the same thing for the response from the internet.

</details>

### Using DMZ Networks to Protect Your Servers
### Basic Layer 3 and 4 Firewall
### Layer 7 Firewall (NGFW)
### Intrusion Detection and Prevention Systems (IDS/IPS)
### VPN Services for Internet Access (Nord, SurfShark, etc..)


</details>


## Takeaways

<!-- <details> -->
<summary>
Stuff worth remembering
</summary>

| Layer   | Name         | Headers                        | Moniker                    |
| ------- | ------------ | ------------------------------ | -------------------------- |
| Layer 1 | Physical     | NA                             | NA                         |
| Layer 2 | Data Link    | MAC address                    | frames                     |
| Layer 3 | Network      | IP address                     | packets                    |
| Layer 4 | Transport    | TCP sequence numbers, checksum | UDP Datagram, TCP Segments |
| Layer 5 | Session      |                                |
| Layer 6 | Presentation |                                |                            |
| Layer 7 | Application  |                                |


- NIC - Network Interface Card
- MAC address - Media Access Control Address
- Hub - stupid, broadcast all messages
- L2 Bridge - has mac address table, smarter the hub
- Switch - replaces the hub+l2-bridge, has the mac address table, eliminates more collisions.
- BUM - Broadcast, Unknown unicast, Multicast
- ARP - Address Resolution Protocol, a broadcast request to discover the MAC associated with an IP address.
- Unknown Unicast - discover mapping of MAC to Port (switch level)
- Multicast - send one message to multiple (but not all) destinations
- Routers break L2 networks into segments with distinct ip ranges.
- DNS
- DHCP - Dynamic Host Control Protocol
- TFTP
- network address - first address in range
- default gateway address - second address in range
- broadcast address - last address in range
- ICMP (ping) - Internet Control Message Protocol (L4 - network layer)
- BGP - border gateway protocol, dynamic routing protocol.
- NTP - Network Time Protocol
- HSRP - Hot Standby Router Protocol
- VRRP - Virtual Router Redundancy Protocol 
- IPFIX - network visibility
- NAT - network address translation
</details>
