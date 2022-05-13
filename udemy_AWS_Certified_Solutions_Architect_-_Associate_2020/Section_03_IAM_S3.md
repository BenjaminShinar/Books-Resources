<!--
// cSpell:ignore
 -->

[main](README.md)

## Section 3 - Identity Access Management & S3

<details>
<summary>
The IAM service controls permissions and access. The S3 service provides object (files) storage.
</summary>

### IAM - Identity Access Managemet

<details>
<summary>
IAM allows you to manage users and their level of access to AWS Console, programmatic access, and control which services can interact with which other services and resource.
</summary>

[IAM faq](https://aws.amazon.com/iam/faqs/)

#### IAM 101

users, groups, permissions, roles.

- Centeralized control of the AWS account
- Shared Access to the account
- Granular permissions
- Identity Federation (Active Directory, Facebook, linkedin, etc)
- Multifactor authentication
- Temporary access for users/devices when necessary
- Password rotation policy
- Integrates with many AWS services
- Support PCI DSS Compliance (a compliance framework)

key terms:

- Users - end users
- Groups - collection of users, each user inherits groups permissions
- Policies - json documents that describe permissions
- Roles - allowing one aws service to work with other resource.

#### IAM Lab

getting our hands dirty with IAM. permissions, access, users, groups, policies, roles. we should choose one region and stick with it.

services, IAM.

we can have a sign-in link, or customize it. we should first activate MFA (multi factor authentication) on our root account. the root account has control over anything, so we should protect it. we can use a virtual MFA device like google authenticator or something else.

anything we do in the IAM service is done globally, a user applies to all regions.

the next step is to create a user. we can decide on which access types are available. we can have _programmatic access_ for applications and working from command line tools, and the _aws management console access_ which is using the aws website. next, we add permissions to the users.

we start with a new group, and we give it the **AdministratorAccess** policy. this is one of the aws managed policies, and some of them are job function. this user has admin control, but it isn't the root user. this is the user we should use in most cases (and not the root user).

a policy is defined as a json object.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*"
    }
  ]
}
```

the next thing is to apply an _IAM Password Policy_, we can define rules to enforce password strength. if we happen to lose the password for one user, we can regenerate it.

Roles are a way for one AWS service to access another service (also, allow access to other aws accounts, other identity management softwares and software). we need to choose the service the gets the role (trusted entity), and give it a policy _AmazonS3FullAccess_.

- **Iam is universal**, it does not apply to regions (currently).
- the "**root account**" is the account created when first setup for your AWS account. it has complete admin Access.
- New Users **have no permission** when first created.
- New Users are assigned \*Access Key Id & Secret Access Keys\*\* when First created
- **These are not the same as a password**. You cannot use the Access Key Id & Secret Access Key to login to the the console. you can usee this to access aws via the APIs and the command line.
- **You only get to view these once**. if you lose them, you have to regenerate them, So save them in a secure location.
- \*_Always setup Multifacto authentication on your account_.
- **You can create and customize your own password rotation policies**.

</details>

### Create A Billing Alarm

we want an alarm to prevent our account from going over the free-tier amount.

we start by logging into the AWS console (website), and we go to the service **Cloud Watch**. then we click _billing_ and we create a billing alarm (there are two buttons with the same name), we choose "total_estimated_charge", we choose a 'static threshold' type, and create the alarm.\
this requires us to create an SNS topic, we create a new topic and enter our email address. we then define the alarm text and hit <kbd>create alarm</kbd>. which we then need to confirm with the mail that were sent.

### S3 - Simple Storage Service

<details>
<summary>
S3 is an Object based Storage (files) which is durable, available, and comes with different tiers.
</summary>

[S3 FAQ](https://aws.amazon.com/s3/faqs/)

#### S3 101

S3 is one of the oldest and most used AWS services.

> S3 providers developers and IT teams ith secure, durable, highly scalable object storage. Amazon S3 is easy to use, with a simple web service interface to store and retrive any amount of data from anywhere on the web.

- S3 is a safe place to store your files
- it is object-based storage
- the data is spread across multiple devices and facilities

files can from 0 bytes to 5 TB, there is unlimited storage. files are stored in buckets. a bucket is like a folder. the bucket name must be unique, as it becomes a globally addressable web address.

if we upload a file successfully, we get an http 200 code.

the data in S3 is object based, an object can be thought of like a file, and it should have:

- key (the file name)
- value (the data)
- version id (version control of the same file)
- metadate (data about the data)
- subresources
  - access control lists
  - torrent

the access control lists are an additional way to control access (permissions) to a specific objects.

data consitency in S3:

- Read after write consistency for PUT requests of new objects
- Eventual consistency for overwrite PUT and DELETE requests (can take some time to propagate)

when we upload a file, we can immediately read it. if we overwrite a file or delete it, then the effect won't be immediate. it might take a few seconds. **changes take more time**.

Amazon guarantees for S3

- Built for 99.99% avaliability for the S3 platform
- amazon guarantees 99.9% availability
- amazon guarantees 99.999999999% durability for s3 information(11 times 9). data won't be lost.

Features:

- Tiered Storage Availabl
- Lifecycle Management
- Versioning
- Encryption
- MFA Delete (multi factor authentication for delete)
- Secure your data using **Access Control Lists** and **Bucket Policies**

S3 Storage Classes:

- Standard: 99.99% avalability, 99.999999999% durability. stored redundantly across multiple devices in multiple facilities, and is designed to the sustain the loss of 2 facilities concurrently.
- IA (Infrequently Accessed): For data that is accessed less frequently, but requires rapid access when needed. lower fee than S3, but you are chared a retrieval fee.
- One Zone IA: a lower-cost option for infrequently accessed data, without the multiple availability zone data resilience.
- Intellgent Tiering: designed to optimize costs by automatically moving data o the most cost-effective access tier, without performace impact or operational overhead. (this uses machine learning, since 2018).
- Glacier: secure, durable, low cost storage class for data archiving, you can reliably store any amount of data at costs that are competitive or cheaper than on-premise solution. retrival times are configurable, and ranges from minutes to hours.
- Glacier Deep Archive: the lowest cost solution, data storage where a retrival time of 12 hours is acceptable

in a table form:

| Specification                             | Standard               | Intelgent Tiering | IA              | one Zone IA     | Glacier          | Glacier Deep Archive |
| ----------------------------------------- | ---------------------- | ----------------- | --------------- | --------------- | ---------------- | -------------------- |
| Durability                                | 99.999999999% (11 9's) | (11 9's)          | (11 9's)        | (11 9's)        | (11 9's)         | (11 9's)             |
| Avalability                               | 99.99%                 | 99.9%             | 99.9%           | 99.5%           | N/A              | N/A                  |
| Avalability SLA (Service Level Agreement) | 99.9%                  | 99%               | 99%             | 99%             | N/A              | N/A                  |
| Avalability Zones                         | &ge;3                  | &ge;3             | &ge;3           | 1               | &ge;3            | &ge;3                |
| Minimum capacity charge per object        | N/A                    | N/A               | 128kb           | 128kb           | 40kb             | 40kb                 |
| Minimum storage duration charge           | N/A                    | 30 days           | 30 days         | 30 days         | 90 days          | 180 days             |
| Retrival fee                              | N/A                    | N/A               | per GB retrived | per GB retrived | per GB retrived  | per GB retrived      |
| First bye latency                         | milliseconds           | milliseconds      | milliseconds    | milliseconds    | minutes or hours | hours                |

Billing:

- Storage
- Requests
- Storage management pricing
- Data transfer pricing
- Transfer accelration
- Cross Region Replication Pricing

cross region replication means that we replicate a bucket in one region, they are mirrored in another region.

> Amazon S3 Transfer acceleration enables fast, easy and secure transfers of files over long distances between you end users and an S3 bucket.\
> Transfer accelration takes advantage of Amazon CloudFront globally distributed ede locations. As the data arrives at an edge location, it is routed to Amazon S3 over an optimized network path.

this means that users uploads use the edge location datacenters to get better speed.

- S3 is **Object-based:** i.e. allows you to upload files.
- Files can be from 0 bytes to 5 TB.
- There is **unlimited storage**.
- Files are stored in **Buckets**.
- S3 is a universal namespace,** names must be unique globally**.
- Buckets get a **DNS name (web address)**.
- **Not suitable** to instal an operating system on or a database.
- successfull uploads generate **http 200 status** code.
- you can turn on **MFA delete**.
- S3 objects have:
  - key (name)
  - value (data)
  - version
  - metadata
  - sub resources
    - access control lists
    - torrent
- **Read after Write Consistency** for PUT.
- **Eventual Concitency** for overwrite PUT and DELETE
- **Storage classes**
  - Standard
  - IA
  - One Zone IA (previously RR - reduced redundancy)
  - Intellgent Tiering
  - Glacier
  - Glacier Deep Archive

The S3 topic makes a big portion of the exam, so it's worth reading the FAQ

#### Let's Create An S3 Bucket

creating an S3 bucket in the console. S3 is also a global service. the buckets have regions

<kbd>Create Bucket</kbd>

we can copy the settings of a different bucket (if we have one).
theres an option called _Bucket settings for Block Public Access_. by default, all files are private and cannot be accessed from outside. we can change this setting.

we can choose if we want to have version control in the bucker, we add tags to the bucket and choose if we want to have **server side encryption**.

tabs:

- Objects
- Properties
- Permissions
- Metrics
- Management
- Access points

the objects are the files, we can upload them with the UI, or use an S3 rest API. a file has a unique ARN and a unique (publicly accessable) URL (if we have the permission).

object Lock prevents object from being modifed or deleted. we can also click <kbd>Edit Storage Class</kbd> and choose a storage class for this object. we can change the storage file for specific files or the entire bucket. objects themselves can have tags.

to make an object publicly accessible, we select it, and then in the actions list we choose <kbd>Make Public</kbd>. the bucket configuration suppress the ability to make object public. so we can go to the _permissions_ tab and change the settings. now we can share this object with the world with the **object URL**.

Bucket names share a common name space. You cannot have the same bucket name as someone else. buckets are viewed globally, but are created in a region. we can have mirrored buckets using _cross region replication_. storage classes can be changed for specific files or the entire bucket.

_transfer acceleration_ allows us to speed up uploads by using edge locations.

Restricing bucket access

- Bucket policies - applies across the whole bucket
- Object policies - applies to individual files
- IAM Policies - apply to users and groups.

#### S3 Pricing Tiers [SAA-C02]

the costs of S3

- Storage
- Requests and data retrivals
- Data transfer
- Management and replication

the tiers:

- standard
- IA
- one zone IA
- intelgent tiering
- glacier
- glacier deep archive

the highest cost is S3 standard, then S3 IA, S3 intelligent tiering, S3 One Zone IA, Glacier, and the cheapest is S3 Glacier deep archive.

exam questions might require us to identify the best storage for a given scenario.

#### S3 Security And Encryption

by default, all buckets are private, wc can set up access control either via _bucket polices_ (per bucket) or with _access control lists_ (per object).\
We can configure buckets to create access logs (which record all request made to the s3 bucket), and then store the logs on another bucket, and even a bucket in a different AWS account.

two types of encryption:

Encryption in transit, like https. using SSL or TLS.

Encryption at rest (server side).

- S3 managed keys --SSE-S3: aws manages the keys
- AWS KMS (key management Service) - managed keys (SSE-KMS)
- Customer Provided keys, customer Key.

we can also encrypt the files on the client side before uploading it.

we select a bucket. choose an object, and then find the _server-side encryption_ settings. we enable it and choose key to use. the we encrypt the entire bucket by choosing the <kbd>Properties</kbd> tab.

#### S3 Version Control

Versioning:

- all versions of the object are stored (all writes, also deletes)
- great for backup
- once enabled, **versioning cant be disabled**, only suspended. this will last until we delete the bucket
- integrates with **Lifecycle** rules.
- we can set MFA Delete, which prevents deletion.

we create a simple file

```
version v1.0
```

we create a bucket, allow public access (to make viewing the file easy), and we enable versioning. now we upload the file and make it public.

we now edit the local file (only the content, not the name) and upload it again. we can toggle <kbd>List versions</kbd> to see all versions.

however, even if we made the first version public, it doesn't mean the new versions will be public. we need to choose which versions we want to make public.

if we delete an object, they will still show when we toggle the versions, they simply have a delete marker. to restore an object we delete the delete marker.

we can still delete specific versions manually, this is done one by one. deleting the file simply creates a delete marker, while deleting a version really deletes the object.

> - All versions are stored (even deletions)
> - Once enabled, can't be disbled, only suspended
> - integrates with lifecycle rules
> - MFA for deleteion can be enabled as well

#### S3 Lifecycle Management and Glacier

controling what happens to objects and control moving objects between tiers. doesn't require versioning, but is integrated with it.

in the s3 console, we look at a bucket, click the <kbd>Management</kbd> tab and <kbd>create a lifecycle rule</kbd>. we can determine a scope or effect the entire bucket. we can look at different lifecycle rules actions:

- transition _current_ versions of objects between storage classes
- transition _previous_ versions of objects between storage classes
- expire _current_ versions of objects
- permanently delete _previous_ versions of objects
- delete expired delete markers or incomplete multipart uploads

Transition allows us to move versions to a different storage tier.we can delete previous versions after a specific time.

#### S3 Lock Policies & Glacier Vault Lock [SAA-C02]

making an object unmodifiable for a set time (retention period).

> You can use S3 object Lock to store object using a **Write Once, Read Many** (WORM) model. it can help you prevent objects from being deleted or modified for a fixed amount of time or indefinitely.
>
> You can use S3 object lock to meet regulatory requirements that require WORM storage, or to add an extra layer of protection against object changes and deleteion.

- _Governace mode_: prevent changes by users based on permissions.
- _Compliance mode_: prevent changes by any user, including the root account user.

the retention period is the timestamp in the metadata that controls the object lock, there is also a **legal hold** on a object version. this isn't tied to a retention period, instead it's a type of AWS permssion.

> S3 Glacier Vault allows you to easily deploy and enforce compliance controls for individual S3 glacier vaults with a **vault lock** policy. you can specify controls, such as **WORM** in a vault lock policy and lock the policy from future changes. **Once locked, the policy can no longer be changed**.

#### S3 Performance [SAA-C02]

S3 prefix are the 'path' inside the bucket to a folder.

bucketname/folder1/folderA/a.txt
bucketname/folder2/folderB/b.txt
bucketname/folder3/folderC/c.txt

so anything between the bucket name and the file name is the **prefix**. because of how low the latency is, we can get better performance by spreading requests across different prefixes.

if we use KMS (server side encryption), there are also limits for that, the limit (quota) is region specific. so this can slow down performance.

multiparts uploads are recomendded for files over 100MB and required for files over 5GB. this increases efficiency by parallelizing the uploads. this can also be done for download with **S3 Byte-Range Fetches**, this makes downloads faster, and if there's a failure, it only effects a specific range, which makes things easier. we can also use this to download partial amounts, like if we design our files to have a header of a specific size, then we can only download that byte range and get some data about the file.

- the more prefixes, the better performance
- we can spread our requests across prefixes.
- KMS server side enctypton has limits.
- multipart uploads and downloads for better performace.

#### S3 Select & Glacier Select [SAA-C02]

> **S3 Select** enables application to retrieve only a subset of daa from an object by using simple SQL expressions. by using S3 select to retrive only the data needed by your applications, you can achieve drastic performance increase.

rather than download the entire file and parse it, we can get only what we want from the object.Glacier Select is similar, we can use Glacier Select to query S3 glacier directly.

(the example is querying a zip file with many csv files inside it)

we get better speed and save money on data transfer.

#### AWS Organizations [SAA-C02]

AWS organization and consoladated billing.

> "AWS Organizations is an accout managemend service that enables you to consolidate multiple AWS account in an organization that you create and centerally manage."

OU - Organization Unit

we can apply policies to organization units, just like user groups.

consolidated billing takes the aggregate usage of all the linked accounts in the organization, so the total price gets the advantages of volume pricing, and is easier to track charges.

demo, in the aws console (website).\
services: <kbd>Aws Organization</kbd>, click <kbd>Create Organization</kbd>, and now our account is the root user of this organization. we can now invite other aws accounts into this organization.

in the other account, we will see the invitation to join the organization.

back in the root account, we can create organizational units and apply service control policies and even enforce rules on how tags.

Some Best Practices with AWS Organizations

- Always enable multi-factor authentication on root account.
- Always use a strong and complex password on root account
- Paying account should be used for billing purposes only. Do not Deploy resources into the paying account.
- Enable/Disable AWS services using Service Control Policies either on the Organizational unit or on individual accounts.

#### Sharing S3 Buckets Between Accounts [SAA-C02]

Roles allow us to give access (temporary or not) to an aws service or other aws accounts.

in our account:\
services, IAM,<kbd>Create Role</kbd>, and we choose _"Another AWS account"_ as the trusted service.\
we paster the account id, and then move to <kbd>Permissions</kbd>. we attach a policy of **AmazonS3FullAccess**. we add tags as wanted, and name the role we want.

when we view this role summary, there is a field called "Give this Link to users who can switch roles in the console". we should copy it, and sign into the other account.

services, AIM, Users, and we add a user <kbd>Create User</kbd>, <kbd>Create Group</kbd>, give permissions. log out. we have to do this with the non-root user.

log in using the new user, and click on the user name in the top level frame, and then click <kbd>Switch Role</kbd>. we either fill in the details or use the link from before. and now we are using the other role, and we can go to the S3 buckets from the other User. if we try any other action we will see that we don't have permissions:

> Three different ways to share S3 buckets across accounts:
>
> 1. Using Bucket Policies and IAM (applies across the entire bucket). programmatic access only.
> 2. Using Bucket Access Control Lists and IAM (individual objects in the bucket). programmatic access only.
> 3. Cross accout IAM roles. programtic and console access.

#### Cross Region Replication

we want to replicate a bucket into a differnet region.

login the cosole. services, storage, S3. we create a new bucket <kbd>Create Bucket</kbd>, give it a unique name, and choose a region for it. we then un-tick the checkbox "block _all_ public access" and confirm this. we continue with the default settings for the bucket. this will be our replicated bucket (destination).

we go to a diffrent bucket, the <kbd>Management</kbd> and click <kbd>Create replication rule</kbd>. we give it a name, provide a role, and choose if we apply the replication to all objects in the bucket or individual objects. we choose a destination bucket (in our account or in a different account), and we get the warning that replication requires versioning to be enabled for the destination bucket.\
we can change the storage class for replicated objects (like S3 Standard-IA), add replication metrics and request Replication Time Control, which ensures speedy replication of objects (99.99% in 15 minutes time frame) but has additional costs. we can add encryption using KMS and decide if we wish to replicate deletion markers.

replicaion works only for new objects, not for existing objects. so we will only see the destination bucket filling up when we upload objects to the source bucket.

the replication doesn't replicate earlier versions and doesn't copy over the permssions to an object, so if the object was public in the source bucket, it won't be carried over to the destination object.

> - Versioning must be enabled on **both** the source and destination buckets.
> - Existing files in bucket are not replicated automatically.
> - All subsequent updated files wil be replicated automatically.
> - Delete markers are not replicated.
> - Deleting individual versions or delete markers will not be replicated.
> - Understand wat Cross Region Replication is at a high level.

#### Transfer Acceleration

> "S3 Transfer Acceleration utilizes the CloudFront Edge Network to accelerate your uploads to S3. Instead of uploading directly to your S3 bucket, you can use a distinct URL to upload directly to an edge location which will then transfer that file to S3. You will get a distinct URL to upload to".

a bucket sits in a region, the users will upload to then edge location.

[S3 Transfer acceleration speed test](http://s3-accelerate-speedtest.s3-accelerate.amazonaws.com/en/accelerate-speed-comparsion.html?) we can see how much we speed gain there is for each region.

</details>

### DataSync Overview [SAA-C02]

Datasync allows us to move large amounts of data to and from AWS, it has encryption, data check, etc...\
its a way of syncing data from an on premise location to the cloud.

- Used to move **large amounts** of data from on-premises to AWS.
- Used with **NFS** and **SMB** compatible file systems.
- **Replication** can be done hourly, daily or weekly.
- Install the **DataSync agent** to start the replication.
- Can be Used to Replicate **EFS** to **EFS** (elastic file system, from EC2 machine).

### CloudFront

<details>
<summary>
CloudFront is Amazon's Content Delivery System service.
</summary>

> "A content delivery system (CDN) is a system of distributed servers (a network) that deliever webpages and other web content to a user based on the geographical locations of the user, the origin of the webpage, and a content delivery server."

imagine if we don't have a CDN, all of the users in the world need to access the main server, which might be in a differnet continent.

if we have a CloudFront enabled, then the users will first access the edge location with the request, and at the first time, the edge location will access the origin and store the result. at the next request, this data will be served from the local location, and will be much faster.

> "Amazon CloudFront can be used to deliver your entire website, including dynamic, static, streaming and interactive content using a global network of edge locations. Request for your content are automatically routed to the nearest edge location, so content is delivered with the best possible performance."

distribution types:

> - Edge Loocation - this is the location where content will be cached. this is separate to an AWS Region/AZ.
>   - not READ only, we can write to edge locations also (like S3 transfer accelration).
> - Origin - this is the origin of all the files that the CDN will distribute, it can be an S3 bucket, an EC2 instance, an Elastic Load Balancer or Route53.
> - Distribution - this is the name given the CDN which consists of a collection of Edge Locations.
>   - Web Distribution - Typically used for Websites.
>   - RTMP - Used for Media Streaming.
> - Objects are cached for the life of the TTL (**Time To Live**).
> - it's possible to clear cached objects (invalidate it), but there is a cost.

#### CloudFront Lab

at our aws management console. we will use a bucket as an origin.

services, networking, <kbd>CloudFront</kbd>. this is a global service.

we click <kbd>Create Distribution</kbd> and choose a web distribution. we fill in the origin domain with S3 Cloud, we can specify a path inside the bucket. we can restrict bucket access, and leave the rest as default. here we can also set the TTL, and we can restrict access using signed URL (cookies, etc). Here we can also put up a WAF (web Application Firewall) to increase security.

creation and deletion take time, once done, we can copy the domain name, and use it as a web address (add the file path).

in the <kbd>settings</kbd> tab, we can see <kbd>invalidation</kbd> and that way we can invalidate objects on the cache. we must disable the distribution before deleting it.

#### CloudFront Signed URL's and Cookies [SAA-C02]

difference between CloudFront Signed Url and cookies and S3 signed URLs

how do we restrict access to a website?

- A signed url is used for individual files. 1 file = 1 url
- A signed cookie is for multiple files. 1 cookie= N files.

signed URLs and cookies are created with an attached policy. this policy can include:

- URL expiration
- IP ranges
- Trusted Signers (which AWS account can created signed URLs)

OAI - Origin Access Identification authentication

users log in to the application, and the application generates a signed URL fot the user to visit.

Cloud front signed URLs:

> - Can have different origins, EC2, S3, etc..
> - Can utilize **caching** features
> - Key-pair Account is wide and managed by the root user
> - Can filter by data, path, IP address, expiration, etc..

on the other hand, S3 signed URLs:

> - Issues are request as the **IAM user** who creates the presigned URL
> - limited **lifetime**

> Summary:
>
> - use signed **URLs/cookies** when you want to secure content so that only the people you authorize are able to access it.
> - a signed URL is for individual files. 1 file = 1 URL.
> - a signed cookie is for multiple files. 1 cookie = multiple files.
> - if you origin is EC2, then use CloudFront.
> - if the origin is S3 and we want a single file, then we use S3 signed URL.

 </details>

### Snowball

<details>
<summary>
Physical Data Transport Solutions.
</summary>

> "Snowball is a petabyte-scale data transport solution that uses secure applications to transfer large amounts of data into and out of AWS. Using Snowball addresses common challenges with large-scale data transfers,including high network costs, long transfer times and security concerns. Transferring data with Snowball is simple, fast, secure, and can be as little as one-fifth of the cost of high-speed internet."

it's big disk on key in a briefcase.

there are 50TB and 80TB disks, has multiple laters of security, TPM (trusted platform module). once used, AWS performs software wipe to prevent the possibility of recovering the data afterwards.

There's also Snowball Edge, which has 100TB data, which also has storage and compute capabilities, so it can be used not only for storing, but also for actual work. it's like having an aws cloud on premises.

AWS Snowmobile is a exabyte scale data-transfer solution.uses 100PB per snowmobile. its a shipping container, pulled by a semi-trailer, which makes transporting data secure, fast, and cost effective.

table of when to use a snowball

| available internet connection | theoretical time to transfer 100TB | when to consider AWS Snowball |
| ----------------------------- | ---------------------------------- | ----------------------------- |
| T3 (44.736 Mbps)              | 269 days                           | 2TB or more                   |
| 100 Mbps                      | 120 days                           | 5TB or more                   |
| 1000 Mbps                     | 12 days                            | 60TB or more                  |

my comparisons
| type | storage | physical size | notes |
| ------------- | ----------- | ------------------ | ---------------------- |
| Snowball | 50 or 80 TB | mini fridge |
| Snowball Edge | 100TB | mini fridge | also has compute power |
| Snowmobile | 100PB | shipping container on a truck |

> Snowball can import and export from S3.

#### Snowball Lab

in services, under "services and migrations", we can see the **Snowball** stuff. we request Amazon to send us a package.

to interact with it, we need to download the snowball client. we connect it to our local network, we get credentials from AWS (client unlock code, and the manifest fold file)

```sh
./snowball start -i 192.168.1.116 -m maifest.bin -u <unlock-code>

#copy files into S3 bucket
/.snowball cp source.file s3:://cloud-guru-snowball

/.snowball stop
```

</details>

### Storage Gateway

> Storage gateway is a service that connects an on-premises software appliance with cloud-based storage to provide seamless and secure integration between an organization's on-premises IT environment and AWS's storage infrastructure. The service enables you to securely store data to the AWS cloud for scalable and cost effective storage.

either a physical device or a virtual machine

- File gatways(FNS, SMB)
- Volume Gateways (iSCSI)
  - stored volumes
  - cached volumes
- Tapes Gateway

file gateways allow us to store our local files on in a amazon S3 bucket. the upload is done asynchronously

volume gateway means storing images (virtual harddisk), and store OS snapshots. Cached volume is for the frequently accessed data, so it's not everything. Tape gateway is a way to store tape gateway and move the data to cloud.

> - File Gateway: for flat files, stored directly on S3
> - Volume Gateway:
>   - Stored volumes: entire dataset is stored on site and is asynchronously backed up to S3
>   - Cached Volumes: entire data set is stored on S3 and the most frequently accessed data is cached on site.
> - Gateway Virtual Tape Library

### Athena vs Macie [SAA-C02]

other services that interact with S3.

**Athena**:\
Interactive query service which enables you analyse and query data located in S3 using standard SQL.

- Serverless. nothing to provision. pay per query / per TB scanned.
- No need to set up complex Extract/ Transfrom/ Load (ETL) process
- Works directly with data stored in S3

can be used to query log files stored in S3, generate business reports on data stored in S3, analyse AWS costs and usage

**Macie**:\

> PII - Personally Identifiable Information

data that can be used to establish a person identity. name, email, credit card number, social security number, address.

Macie is a security service which uses Machine Learning and NLP (natural language processing) to discover, classify and protect sensitive data stored in S3.

- Uses AI to recognize if your S3 objects contain sensitive data such as PII.
- Dashboards, reporting and alerts
- Works directly with data stored in S3.
- Can also analyze CloudTrail logs
- Great for PCI-DSS and preventing identity theft.

### Identity Access Management & S3 Summary

**IAM**

- User
- Group
- Roles
- Policies
  - how Policies look as Json documents
- Globally - across all regions
- The root account
- Least Privilege
- AccessKey Id and Secret Access Key
- Console password
- Multi factor authentication
- Password rotation policy

**S3**

- Object based
- Files are stored in buckets
- Global namespace, buckets must be uniquely named
- Buckets have policies
- Objects can have Access Control lists
- Can be configured to create access logs
- Private by default
- Objects are
  - key
  - value
  - version Id
  - metadata
  - sub resources
    - access control lists
    - torrents
- Consistency models
  - Read after write consistency for new objects
  - Eventual consistency for update and deletion
- Storage tiers:
  - S3 standard
  - S3 Infrequently Access
  - S3 one Zone IA
  - S3 intelligent tiering
  - S3 glacier
  - S3 glacier deep archive
- How to het the best value of S3. by cost:
  - S3 standard
  - S3 Infrequently Access
  - S3 intelligent tiering
  - S3 one Zone IA
  - S3 glacier
  - S3 glacier deep archive
- Encryption
  - In transit: SSL/TLS
  - At rest:
    - Server Side:
      - S3 managed Keys - SSE-S3
      - AWS key management service,SSE-KMS
      - Customer provided Keys - SSE-C
    - Client Side - enctyped before upload
- S3 Object Lock
  - Write Once, Read many Model
  - Applied across the bucket or per object
  - Governance mode: requires special permssion
  - Compliance mode: can't be modifed, even by root user
- S3 glacier Vault Lock
- S3 prefixes for performance gains.
- SSE-KMS limits
- Multipart uploads and downloads (byte range fetch) for performance gains
- S3 Select
- AWS organizations
  - Service Control Policies
- Accessing buckets across accounts
- Sharing buckets across accounts
  - Bucket Policies
  - Bucket Access control lists
  - Cross account Roles
- Replicating S3 buckets across regions
  - must enable versioning
  - only new files
  - delete markers aren't replicated
- lifecycle policies
  - moving objects across storage tiers
  - can work with versioning
- Transfer Acceleration

**CloudFront**

- Edge Location
- Origin
- Distribution
  - Web distribution
  - RTMP - media
- Time To live (TTL) caching
- invalidating cached object
- Signed Url, 1 per file
- Signed Cookie, multiple files
- if origin is EC2, use cloud front
- Signed S3 urls

**AWS datasync**

- data sync replication and synchronization

**Snowball**

- physical data transfer
- import and export to S3

**Storage Gateway**

- File Gateway
- Volume Gateway
  - Stored volume
  - Cached volume
- Tape Gateway

- Athena: query S3 using SQL
- Macie: analyze data in S3 to discover personal data.

it's worth reading the S3 FAQ before the exam.

### Quiz 2: Identity Access Management & S3 Quiz

> - Power User Access allows **Access to all AWS service except the management of groups and users within IAM**.
> - You have been asked to advise on a scaling concern. The client has an elegant solution that works well. As the information base grows they use CloudFormation to spin up another stack made up of an S3 bucket and supporting compute instances. The trigger for creating a new stack is when the PUT rate approaches 100 PUTs per second. The problem is that as the business grows that number of buckets is growing into the hundreds and will soon be in the thousands. You have been asked what can be done to reduce the number of buckets without changing the basic architecture.
>   - _ANSWER: Until 2018 there was a hard limit on S3 puts of 100 PUTs per second. To achieve this care needed to be taken with the structure of the name Key to ensure parallel processing. As of July 2018 the limit was raised to 3500 and the need for the Key design was basically eliminated. Disk IOPS is not the issue with the problem. The account limit is not the issue with the problem._
> - What is the availability of S3 – OneZone-IA?
>
>   - 99.5%
>   - _ANSWER: OneZone-IA is only stored in one Zone. While it has the same Durability, it may be less Available than normal S3 or S3-IA._
>
> - AWS S3 has four different URLs styles that it can be used to access content in S3. The Virtual Hosted Style URL, the Path-Style Access URL, the Static web site URL, and the Legacy Global Endpoint URL. Which of these represents a correct formatting of the Virtual Hosted Style URL style
>   - https://my-bucket.s3.us-west-2.amazonaws.com/fastpuppy.csv
>   - _ANSWER: Virtual style puts your bucket name 1st, s3 2nd, and the region 3rd. Path style puts s3 1st and your bucket as a sub domain. Legacy Global endpoint has no region. S3 static hosting can be your own domain or your bucket name 1st, s3-website 2nd, followed by the region. AWS are in the process of phasing out Path style, and support for Legacy Global Endpoint format is limited and discouraged. However it is still useful to be able to recognize them should they show up in logs. https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html_
> - How many S3 buckets can I have per account by default?
>   - 100 buckets
> - You work for a busy digital marketing company who currently store their data on-premise. They are looking to migrate to AWS S3 and to store their data in buckets. Each bucket will be named after their individual customers, followed by a random series of letters and numbers. Once written to S3 the data is rarely changed, as it has already been sent to the end customer for them to use as they see fit. However, on some occasions, customers may need certain files updated quickly, and this may be for work that has been done months or even years ago. You would need to be able to access this data immediately to make changes in that case, but you must also keep your storage costs extremely low. The data is not easily reproducible if lost. Which S3 storage class should you choose to minimize costs and to maximize retrieval times?
>   - S3 IA
>   - _ANSWER: The need to immediate access is an important requirement along with cost. Glacier has a long recovery time at a low cost or a shorter recovery time at a high cost, and 1Zone-IA has a lower Availability level which means that it may not be available when needed._

</details>

##

[next](Section_04_EC2.md)\
[main](README.md)
