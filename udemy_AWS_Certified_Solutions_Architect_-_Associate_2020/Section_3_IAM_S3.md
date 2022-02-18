<!--
// cSpell:ignore
 -->

[main](README.md)

## Section 3 - Identity Access Management & S3

<!-- <details> -->
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
- **These are not the same as a password**. You cannot use the Access Key Id & Secret Access Key to log0in to the the console. you can usee this to access aws via the APIs and the command line.
- **You only get to view these once**. if you lose them, you have to regenerate them, So save them in a secure location.
- \*_Always setup Multifacto authentication on your account_.
- **You can create and customize your own password rotation policies**.

</details>

### Create A Billing Alarm

we want an alarm to prevent our account from going over the free-tier amount.

we start by logging into the AWS console (website), and we go to the service **Cloud Watch**. then we click _billing_ and we create a billing alarm (there are two buttons with the same name), we choose "total_estimated_charge", we choose a 'static threshold' type, and create the alarm.\
this requires us to create an SNS topic, we create a new topic and enter our email address. we then define the alarm text and hit <kbd>create alarm</kbd>. which we then need to confirm with the mail that were sent.

### S3 - Simple Storage Service

<!-- <details> -->
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

### AWS Organizations [SAA-C02]

### Sharing S3 Buckets Between Accounts [SAA-C02]

### Cross Region Replication

### Transfer Acceleration

### DataSync Overview [SAA-C02]

### CloudFront Overview

### CloudFront Lab

### CloudFront Signed URL's and Cookies [SAA-C02]

### Snowball Overview

### Snowball Lab

### Storage Gateway

### Athena vs Macie [SAA-C02]

</details>

### Identity Access Management & S3 Summary

### Quiz 2: Identity Access Management & S3 Quiz

</details>

##

[next]()
[main](README.md)
