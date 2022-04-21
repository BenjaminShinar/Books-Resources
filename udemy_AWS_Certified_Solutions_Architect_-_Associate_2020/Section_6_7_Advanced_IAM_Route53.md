<!--
// cSpell:ignore undecillion
 -->

[main](README.md)

## Section 6 - Advanced IAM

<details>
<summary>
More IAM features
</summary>

### AWS Directory Service [SAA-C02]

- AWS Managed Microsoft AD
- Simple AD
- AD connector
- Cloud Directory
- Amazon Cognito User Pool

AWS directory services is a family of managed services which allows us to connect AWS resources to on-premises Active Directory. it allows users to access AWS resources with existing corporate credentials, and we can use SSO (single sign on) to any domain-joined EC2 instance.

Active directory is an on-premises directory service. an hierarchical database of users, groups and computers (**trees** and **forests**). we can have group policies.

Active directory is based on two protocols: LDAP (Lightweight Directory Access Protocol) and DNS. this comes with the cost of management overhead, so we might prefer a managed service, such as AWS Managed Microsoft AD.

it uses Domain Controllers (DC) which run winnows Server, each on a different avaliability zone. each running on a separate VPC. the DC are exclusive, we can add more DCs for better availability and performance.\
we can also extended existing AD to on-premises by using **AD Trust**.

separation of responsibilities

AWS:

- multi-AZ deployment
- patch, monitor, recover
- instance rotation
- snapshot and restore

Customer:

- Users, groups, group polices
- standard AD tools
- scaling out the DC
- Trust (resource forest)
- Certificate authorities (LDAPs)
- AD federation

Simple AD is the simpler version of managed active directory. it's a standalone manged directory with basic AD features. it has two versions: small (for 500 and less users), and large (for 5000 and less). using Simple AD makes it easier to manage EC2 machines and keys.

Simple AD doesn't support Trusts, so it can't connect to the existing corporate Active directory.

**AD Connector**

- Directory gateway (proxy) for on-premises AD.
- Avoid caching information in the cloud.
- allow on-premises users to log in AWS using AD.
- join EC2 instance to your existing AD domain.
- scale actors multiple AD connetctors

**Cloud Directory**

- directory based store for developers
- multiple hierarchies with hundred of millions of objects
- fully manges service
- use cases:
  - organization charts
  - course catalog
  - device registries

**Amazon Cognito User Pool**

managed users directory for SaaS applications, sign up and sign-in for web or mobile, works with social media identities.

AD comptabile:

- Managed Microsoft AD
- Simple AD
- AD Connector

  Not AD compatible:

- Cloud Directory
- Cognito user pools

### IAM Policies [SAA-C02]

permission boundaries and IAM policies

ARN - Amazon Resource Name,a unique identifier to all amazon resource.

all arn begin with this format

> `arn:<partition>:<service>:<region>:<account_id>:`

- partition: the infrastructure, usually _aws_ but can also be _aws-cn_.
- service: the service itself, _s3_,_ec2_,_rds_
- region: aws region _us-east-1_
- account_id: the account id (12 digits)

and after that comes one of the following

> `resource`\
> `resource_type/resource`\
> `resource_type/resource/qualifer`\
> `resource_type/resource:qualifer`\
> `resource_type:resource`\
> `resource_type:resource:qualifer`\

if there is no region, like for IAM, then the field is skipped, so we'll see `::`. for an s3 object, theres no region needed or even a user (buckets are globally unique), so both are skipped, and we see `:::`.

we can use the asterisks `*` wild card to match all resources of a specific type.

#### Policies

IAM policies are json documents that define permissions.

- identity policies - to users
- resource policies - to resources, what actions are allowed.

policies must be attached after they are created, they aren't used on their own. a policy is a list of statement.

```json
{
  "Version": "2012-10-17",
  "Statement": []
}
```

each statement matches an AWS Api request.

- Sid - human readable name
- Effect - Allow or Deny.
- Action - `<service>:<verb>`, can use wildcard.
- Resource - which resource does the statement refer to? (can be a list as well)

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "SpecificTable",
      "Effect": "Allow",
      "Action": [
        "dynamodb:BatchGet*",
        "dynamodb:DescribeStream*",
        "dynamodb:DescribeTable*",
        "dynamodb:Get*",
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:BatchWrite*",
        "dynamodb:CreateTable",
        "dynamodb:Delete*",
        "dynamodb:Update*",
        "dynamodb:PutItem"
      ],
      "Resource": "arns:aws:dynamodb:*:*:table/MyTable"
    }
  ]
}
```

in aws Console, <kbd>IAM</kbd>, <kbd>Policies</kbd>, we have AWS managed policies (which aren't editable), and customer managed policies, which we create.

here is a policy that works against an S3 bucket.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["s3:ListBucket"],
      "Resource": "arns:aws:s3:::test"
    },
    {
      "Effect": "Allow",
      "Action": ["s3:PutObject", "s3:GetObject", "s3:DeleteObject"],
      "Resource": ["arns:aws:s3:::test/*"]
    }
  ]
}
```

to attach a policy, <kbd>Roles</kbd>, <kbd>Create Role</kbd>, trusted entity of AWS service, Ec2 and then we choose the policy which we created. we can also create an inline policy (a policy which is limited to a role), this is for ad-hok policies, not best practice.

- if an action isn't allowed, it's an implicitly denied.
- an Explicit deny policy overweights anything else. if one policy allows an action and another policy explicitly denies it, the action is prohibited.
- only attached policies have effect

#### Perission Boundaries

- used to **delgate** adminitsration to other users.
- Prevent **privilege escalation** or **unnecessarily broad permissions**.

it controls the maximum permissions an IAM policy can grant

- developers creating roles for lambda functions
- application owners creating roles for EC2 instances
- Admins creating ad hoc users

in the console, <kbd>IAM</kbd>, <kbd>Users</kbd>, <kbd>Set Boundary</kbd> choose a policy. this is stronger than any other permission of the role or the user.

### Resource Access Manager (RAM) [SAA-C02]

Account Isolation, multiAccount strategy. RAM allows resource sharing between accounts, not all resources can be shared, only a small subset

resources which can be shared:

- App Mesh
- Aurora
- CodeBuild
- EC2
- EC2 Image Builder
- License Manager
- Resource Groups
- Route53

demo using RDS, <kbd>RAM</kbd>, <kbd>Create Resource Share</kbd>, select the aurora resource type, the resource itself, and in the <kbd>principle</kbd> , we add the other account.

in the 2nd account, in <kbd>RAM</kbd>, we need to accept the resource sharing invitation.

### AWS Single Sign-On [SAA-C02]

SSO - single Sign On

managing user permssions is complicated, but we can simply it by using SSO to centrally manage access. we can the existing identities to access other accounts. it integrates with AD or any other SAML identities.

SAML - Security Assertion Markup Langauge

all sign-on activities are recorded in CloudTrail

### Advanced IAM Summary [SAA-C02]

- Active Directory
- Connect Aws Resource with on-premises AD
- SSO to any domain joined EC2 instance
- AWS manage microsoft AD
- AD Trust
- Division of responsability between Amazon and the customer
- Simple AD - doesn't support trust
- AD connector (directory gateway) - supports trusts
- CloudDirectory - hierarchical data (not AD compatible)
- Cognito user pool - works with social media (not AD compatible)
- ARN syntax
- IAM policy structure
  - effect
  - action
  - resource
  - sid (human readable name)
- identity vs. resource policy
- Policy evaluation logic (explicit deny is the strongest)
- Permission Boundaries - maximum permission
- Resource Access Manger
- Resource sharing between account
- Single Sign-on
- centrally manage accsss
- using existing identities
- Account-level permissions
- SAML

</details>

## Section 7 - Route53

<!-- <details> -->
<summary>
Domain Name Server
</summary>

### DNS 101

the source of the name is because port 53 is the dns port. it's like an interstate highway name.

this is not free tier.

dns is used to convert a human readable address into an ip address (IPv4 and IPv6). IPv4 is 32 bit field (over 4 billion addresses), but now that we have so many devices (phones, computers, smart television, sensors) the number doesn't look so large, so IPv6 is 128 bit, so it has about 340 undecillion addresses (3.4\*10^38 or something)

top level domain: the last word in the domain after the dot. the second to last word is "second level domain name" (optional).

| domain   | second level | top level |
| -------- | ------------ | --------- |
| ".com"   | NA           | com       |
| ".edu"   | NA           | edu       |
| ".gov"   | NA           | gov       |
| "co.uk"  | co           | uk        |
| "gov.uk" | gov          | uk        |
| "com.au" | com          | au        |

this is controlled by the Internet Assign Number Authority which manages the all the domains, each register (an address) must be unique in that region, so all domain registar must be part of the international organization.

a registar can assign registries in a domain, amazon is a domain registar, GoDaddy is a domain registar, etc..

SOA - start of authority record:

- The name of the server that supplied the data for the zone
- the administrator of the zone
- the current version of the data file
- the default number of second for the time-to-live file on resource records

NS - name Server Records

> used by top level domain servers to direct traffic to the content DNS server which contains the authoritative DNS records.

address -> top level domian -> NS records -> SOA

**A** record: from name to an Address. the most simple.

TTL - time to liver, the duration of time that a DNS record is cached on either the resolving server or the local computer. the shorter it is, the faster it is for changes to propagate. so if I access an address and the dns links me to an IP address, and then someone buys the domain name (and now it points to another IP address), it can take up to 48 hours for this to be updated.

**CName (canonical name)** record

> A CName can be used to resolve one domain name to another. for example, you may have a mobile website with the name "http://m.acloud.guru" that is used for when users browse to your domain name on their mobile devices. You may also want the name "http://mobile.acloud.guru" to resolve to the same address.

we map one domain record to another record.

**Alias** records

> Alias records are used to map resource record sets in your hoster zone to Elastic Load Balancers, CloudFront distributions or S3 buckets that are configured as websites.\
> Alias records work like a CNAME record in that you can map one DNS name "www.example.com" to another target DNS name "elb1234.elb.amazonaws.com"

A CName can't be used for a naked domain (zone apex record, no www). you can't have a CName for "http://acloud.guru", it must be either an A record of an Alias.

- Elastic Load Balancers do not have pre-defined IPv4 addresses. they are resolved by using a DNS name.
- differences between Alias Record and CNAME
  - given the choice, always choose alias over CNAME
- DNS types
  - SOA record
  - NS record
  - A record
  - CNAME record
  - MX record (mail)
  - PTR record (reverse lookup from ip to address)

### Route53 - Register A Domain Name Lab

demo in the AWS manager console. under <kbd>services</kbd>, choose <kbd>Route53</kbd>. we first register a domain (for a price), we need to fill the personal details and pay up. it can take up to three days to complete. once successful, the domain will appear as a **hosted zone**.

we will now launch EC2 instances, we choose three regions to deploy ec2 instances in. we use the same bootstrap issue to host an serve an html webpage (with the region name). we might need to set up new security group if we never used those regions. we will use them in the later portions.

> - you can buy domain names directly with AWS.
> - it can take up to 3 days to register a domain.

### Routing Policies

<!-- <details> -->
<summary>

//TODO: add Summary

</summary>

#### Route53 Routing Policies Available On AWS

#### Simple Routing Policy Lab

#### Weighted Routing Policy Lab

#### Latency Routing Policy

#### Failover Routing Policy

#### Geolocation Routing Policy

#### Geoproximity Routing Policy (Traffic Flow Only)

</details>

### Multivalue Answer

### Route53 Summary

### Quiz 5: Route53 Quiz

##

[next](Section_8_Route53.md)\
[main](README.md)
