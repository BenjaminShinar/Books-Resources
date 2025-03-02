<!--
// cSpell:ignore Crisci datagram Netflow IPFIX HSRP VRRP nslookup NGFW Nord subnetting classful VLSM
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
