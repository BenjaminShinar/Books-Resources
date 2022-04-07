<!--
// cSpell:ignore veltz eksctl mythicaleks moreutils envsubst mikefarah kubeconfig
 -->

# Container Services Day – EKS

2022-04-06 webinar, by [Floor28](https://aws-experience.com/emea/tel-aviv/)

> In this workshop you will learn how to run Kubernetes on AWS leveraging Amazons Elastic Kubernetes Service (EKS). 
> 
> The is an introductory level workshop that will start with fundamental concepts of containers and get you to the point where you can run applications on Amazon EKS cluster. 
> 
> During this hands-on workshop, you will get to deploy applications on Amazon EKS after understanding the basics of containers and be able to ask questions from AWS container experts who will be delivering and moderating the workshop.
> 
> Who Should attend: This workshop level is introductory and intended for those who are on their evaluating/using containers stage.
>
> *The workshop will be delivered in Hebrew*

Avi, Veltz - software solution architect

## Module 1 - Introduction

introduction to containers, comparison between containers and VM. 

application components:
- enging
- code
- dependencies
- configurations

container and container image. container image - singular, container - process, multiple instances, runtime. the layers model.

introduction to Docker, orchastration, ECS.

AWS container services landscape:
1. Management - ecs, eks
1. Hosting - ec2, aws fargate
1. Image registry - ecr

ecr: fully managed, secure, highly available, simplified workflow

fargate is managed by aws, so we don't need to configure the ec2 machine and manage it (packages, security, etc).

(scalability)

reason to use continers:
- accelerate software development (microservices)
- as opposed to monolith applications
- scalability to inner componenets
- cloud-native applications, the modern approach. we can even take an existing application and package it to a container image without refactoring (incremental approach).
- remove dependencies between development teams / components.
- take advantage of cloud orchestration tools to automate scaling, deployments, etc...

## EKS

Understaing eks architecture, kubelet-s, VPC (virtual private cloud), control plane

instace type, rightsizing. OnDemand, Reserved, Spot (auto scaling groups)

control plane - api-servers, etcd (key value database, source of truth). single tennants. eks cluster sits on more than one availability zone, NLB (network load balancer).

Kubernetes concepts:
- Pod - containers that share ip, namespace, storage volume.
- ReplicaSet - manages lifecycle of pods, ensure number of pods.
- Deployment - manage replicate set, provide declarative updates to pods.
- Label - organize and select groups of objects.
- Namespace - logical partition, isolate resources withing a cluster.

Network load balancers (services):
- ClusterIP - cluster internal ip
- NodePort - exposes the service on each node IP at static ip.
- LoadBalances - expose the service externally using a cloud providers load balancer.

ingres objects
- expose HTTP/HTTPS routes to services withing the cluster
- default service type: **ClusterIP**
- examples: AWS ALB (application load balancer),NGINX,F5, HAProxy

load balancer controller, by aws:
- NLB: service type load balancer
- ALB - ingress

kubernetes uses declarative style, `kubectl apply -f`, we modify the configuration files, and don't interact directly.

**Networking**

private eks API endpoints:
- **control plane vpc** (aws account), the manageing nodes
- **worker vpc** (your account). worker pods interact with each other with through this vpc.connection through *eni (elastic network interface)* to the control plane vpc with *private hosted zone*. 

**Security**

RBAC - roles based Access Control

kubernetes roles are defined in a namespace,we can define *ClusterRoles* accross all namespace. resources (pods,nodes) and verbs (get, put, delete).

authentication with IAM (aws identity) against kubernetes RBAC. access can be allowed or denied. we can't see the user who creates the cluster in the IAM console.

> Service account provides aws permissions to the containers in any pod that uses that service account.

with this pods can have identity and interact with AWS resources (s3 bucket)

logging and auditing, CloudWatch and CloudTrail, we can enable this from the console. we can also collect metrics for the data plane with 3rd party service (prometheus, grafane), or use Cloud Watch insight by AWS.


## Hands on Lab

[hands on Lab](https://catalog.us-east-1.prod.workshops.aws/workshops/ed1a8610-c721-43be-b8e7-0f300f74684e/en-US/contdock/dockerbasics)

### Docker
1. edit the dockerfile
2. build the image
3. run it (pass the environment variables)
4. push it to the ecr repository



```sh
TABLE_NAME=$(aws dynamodb list-tables | jq -r .TableNames[0]) #get table name
```
### Launch EKS
Starting with deploying the cluster before the lecture. getting and installing all sorts of stuff, enabling auto completion

```sh
 #install awscli
sudo pip install --upgrade awscli && hash -r

#install eksctl
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv -v /tmp/eksctl /usr/local/bin
#confirm 
eksctl version

#install kubectl
sudo curl --silent --location -o /usr/local/bin/kubectl \
https://amazon-eks.s3.us-west-2.amazonaws.com/1.19.6/2021-01-05/bin/linux/amd64/kubectl

sudo chmod +x /usr/local/bin/kubectl
#confirm, will have a warning about ports
kubectl version

#install helm
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

#instal stuff
sudo yum -y install gettext bash-completion moreutils

# enable auto completion
kubectl completion bash >>  ~/.bash_completion
. ~/.bash_completion

eksctl completion bash >> ~/.bash_completion
. ~/.bash_completion

# install yq
echo 'yq() {
docker run --rm -i -v "${PWD}":/workdir mikefarah/yq yq "$@"
}' | tee -a ~/.bashrc && source ~/.bashrc


# verify
for command in kubectl jq envsubst aws eksctl kubectl helm
do
    which $command &>/dev/null && echo "$command in path" || echo "$command NOT FOUND"
done

```

create cluster config yaml
```yaml
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

metadata:
  name: mythicaleks-eksctl
  region: ${AWS_REGION}
  version: "1.19"

availabilityZones: ["${AWS_REGION}a", "${AWS_REGION}b", "${AWS_REGION}c"]

managedNodeGroups:
- name: nodegroup
  desiredCapacity: 3
  ssh:
    allow: true
    publicKeyName: mythicaleks

# To enable all of the control plane logs, uncomment below:
# cloudWatch:
#  clusterLogging:
#    enableTypes: ["*"]
```
if we see this
> Unable to connect to the server: getting credentials: exec plugin is configured to use API version client.authentication.k8s.io/v1alpha1, plugin returned version client.authentication.k8s.io/v1beta1Launching

we should run this:
```sh
aws eks update-kubeconfig --name mythicaleks-eksctl
```


### Preparing EKS for the Mysfits

```sh
# create IAM OIDC
eksctl utils associate-iam-oidc-provider --cluster=mythicaleks-eksctl --approve

#download IAM policy
curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json

#create the policy
aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy
    --policy-document file://iam-policy.json

#get the policy ARN
export PolicyARN=$(aws iam list-policies --query 'Policies[?PolicyName==`AWSLoadBalancerControllerIAMPolicy`].Arn' --output text)
echo $PolicyARN


# create iamServiceAccount
eksctl create iamserviceaccount  --cluster=mythicaleks-eksctl --namespace=kube-system  --name=aws-load-balancer-controller --attach-policy-arn=$PolicyARN --override-existing-serviceaccounts --approve

# add eks helm repository
helm repo add eks https://aws.github.io/eks-charts

# target group binding CRD
kubectl apply -k "github.com/aws/eks-charts/stable/aws-load-balancer-controller//crds?ref=master"
# installing load balancer
helm upgrade -i aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n kube-system \
  --set clusterName=mythicaleks-eksctl \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller
#verify
kubectl logs -n kube-system deployments/aws-load-balancer-controller
kubectl -n kube-system get deployments
```

(we skip the dashboard part)

### Deploying the Monolith

```sh
# move to directory
cd /home/ec2-user/environment/amazon-ecs-mythicalmysfits-workshop/workshop-1/app/monolith-service
# create dynamoDB Policy file
TABLE_NAME=$(jq < ../../cfn-output.json -r '.DynamoTable')
cat << EOF > iam-dynamodb-policy.json
{
"Version": "2012-10-17",
"Statement": [
	{
		"Effect": "Allow",
		"Action": [
			"dynamodb:*"
		],
		"Resource": "arn:aws:dynamodb:${AWS_REGION}:${ACCOUNT_ID}:table/${TABLE_NAME}"
	}
]}
EOF

#create the policy
aws iam create-policy \
				--policy-name MythicalMisfitDynamoDBTablePolicy \
				--policy-document file://iam-dynamodb-policy.json

#get the policy ARN
export PolicyARNDynamoDB=$(aws iam list-policies --query 'Policies[?PolicyName==`MythicalMisfitDynamoDBTablePolicy`].Arn' --output text)
echo $PolicyARNDynamoDB

#Create the Kubernetes Service Account
eksctl create iamserviceaccount \
    --cluster=mythicaleks-eksctl \
    --namespace=default \
    --name=mythical-misfit \
    --attach-policy-arn=$PolicyARNDynamoDB \
    --override-existing-serviceaccounts \
    --approve


# creating a manifest - the variables with $ will be replaced, we create the replica set under the deployment, the service of the load balancer type

cat << EOF > monolith-app.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
 name: mythical-mysfits-eks
 namespace: default
 labels:
   app: mythical-mysfits-eks
spec:
 replicas: 2
 selector:
   matchLabels:
     app: mythical-mysfits-eks
 template:
   metadata:
     labels:
       app: mythical-mysfits-eks
   spec:
     serviceAccount: mythical-misfit
     containers:
       - name: mythical-mysfits-eks
         image: $MONO_ECR_REPOSITORY_URI:latest
         imagePullPolicy: Always
         ports:
           - containerPort: 80
             protocol: TCP
         env:
           - name: DDB_TABLE_NAME
             value: ${TABLE_NAME}
           - name: AWS_DEFAULT_REGION
             value: ${AWS_REGION}
---
apiVersion: v1
kind: Service
metadata:
 name: mythical-mysfits-eks
 namespace: default
spec:
 type: LoadBalancer
 selector:
   app: mythical-mysfits-eks
 ports:
 -  protocol: TCP
    port: 80
    targetPort: 80
EOF

# deploy the manifest
kubectl apply -f monolith-app.yaml

# check stuff
kubectl describe pods
kubectl get pods --watch

kubectl get services -o wide

# get ELB ip
ELB=$(kubectl get service mythical-mysfits-eks -o json | jq -r '.status.loadBalancer.ingress[].hostname')
curl -m3 -v $ELB # 'nothing to see here'
curl -m3 -v $ELB/mysfits # dynamodb table returned

cd /home/ec2-user/environment/amazon-ecs-mythicalmysfits-workshop/workshop-1/web
echo $ELB

# modify the index.html file

#upload to S3
BUCKET_NAME=$(jq < ../cfn-output.json -r '.SiteBucket')
aws s3 ls ${BUCKET_NAME}

aws s3 cp index.html s3://${BUCKET_NAME}/ --grants read=uri=http://acs.amazonaws.com/groups/global/AllUsers full=emailaddress=user@example.com


curl $BUCKET_NAME.s3-website.$AWS_REGION.amazonaws.com/

#delete stuff
cd /home/ec2-user/environment/amazon-ecs-mythicalmysfits-workshop/workshop-1/app/monolith-service
kubectl delete -f monolith-app.yaml
```
## Takeaways

https://aws.amazon.com/blogs/containers/automated-software-delivery-using-docker-compose-and-amazon-ecs/

Get curret IAM Role: `aws sts get-caller-identity`

set environment variables
- linux: `export AWS_DEFAULT_REGION=us-east-1`
- windows cmd: `set AWS_DEFAULT_REGION=us-east-1`
- powershell: `$Env:AWS_DEFAULT_REGION='us-east-1'`