<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# Kubernetes Training Course

personal document for the [internal k8s course](https://fluffy-couscous-v3rem2o.pages.github.io/).


In <k8s>Kubernetes</k8s> we set the desired state, we use declerative style (not imperative), and the cluster takes care of moving the parts around to get to that state.

<k8s>Control Plane</k8s>:

- <k8s>API Server</k8s> - the "front door" - communication goes through it, the kubectl commands and internal .communication inside the clusters. writes to etcd.
- <k8s>etcd</k8s> - key-value store that holds ALL cluster state.
- <k8s>Scheduler</k8s> - decides which node runs which pod.
- <k8s>Controller Manager</k8s> - runs control loops that reconcile desired vs actual state (e.g., "keep 3 replicas running").    

Worker Node:

- <k8s>Kubelet</k8s> - agent the manages pods on this node.
- <k8s>Kubeproxy</k8s> - handles service networking.
- <k8s>Container runtime</k8s> - run the actual containers.


the controller manager runs the _reconciliation loop_ to bring the current state to the desired state.

The <k8s>Pod</k8s> is the smallest 'unit' in k8s, it wraps around one or more containers which share network namespace and storage volumes.\
<k8s>ReplicaSet</k8s> - a set of identical pods.\
A <k8s>Deployment</k8s> is one step above the ReplicaSet, it handles updates, rolling updates and rollbacks.

```sh
# Apply and inspect
kubectl apply -f deployment.yaml
kubectl annotate deploy/nginx-deploy kubernetes.io/change-cause="Initial deployment of nginx:1.25"
kubectl get deploy,rs,pods    # See the Deployment, ReplicaSet, and Pods

# Scale up
kubectl scale deploy nginx-deploy --replicas=5

# Rolling update - change the image version
kubectl set image deploy/nginx-deploy nginx=nginx:1.26
kubectl annotate deploy/nginx-deploy kubernetes.io/change-cause="Updated nginx to 1.26" --overwrite

# Watch the rollout progress
kubectl rollout status deploy/nginx-deploy

# Rollback to the previous version
kubectl rollout undo deploy/nginx-deploy

# View rollout history
kubectl rollout history deploy/nginx-deploy
```

<k8s>Services</k8s> exist to make communication inside cluster possible, since everytime a pod restarts it gets a new IP. the Service resource provides a stable IP and routes traffic to the current set of matching resources (pods).\
The default <k8s>ClusterIP</k8s> service allows internal communication only.\
There are cases where a client needs to connect to a specific pod, such as databases which need to have the same connection to a replica pod, or a redis database that each pod holds a subset of the data. in these cases, the client needs to match to a specific IP. a headless service exposes all the internal ip addresses, without load balancing them. headless services are often used in <k8s>StatefulSet</k8s> workloads.


| Type                    | Description                                              | Use Case                                                                   |
|-------------------------|----------------------------------------------------------|----------------------------------------------------------------------------|
| <k8s>ClusterIP</k8s>    | Internal cluster IP only (default)                       | Service-to-service communication                                           |
| <k8s>NodePort</k8s>     | Exposes on each node's IP at a static port (30000-32767) | Dev/testing external access                                                |
| <k8s>LoadBalancer</k8s> | Provisions external load balancer (cloud providers)      | Production external access                                                 |
| <k8s>ExternalName</k8s> | Maps to a DNS name (CNAME record)                        | Aliasing external services                                                 |
| None                    | Headless Service, no virtual IP                          | when the clients need to connect to specific pods (statefull applications) |

```sh
kubectl apply -f service.yaml
kubectl get svc

# Test from inside the cluster - launch a temporary curl pod
kubectl run curl --image=curlimages/curl --rm -it -- sh
# Inside: curl http://nginx-svc.default.svc.cluster.local
# The format is: <service-name>.<namespace>.svc.cluster.local
```

## Ingress

<details>
<summary>
Kubernetes Ingress for external access. 
</summary>

<k8s>NodePort</k8s> services make internal addresses reachble from outside, but it exposes the address through ports, rather than hostname and paths:

- have: `http://node-ip:31234`
- want: `https://myapp.example.com`

The solution is to use Ingress:

- Host based routing (`api.example.com` vs `app.example.com`)
- Path based routing (`/api` vs `/web`)
- TLS termination

The Ingress itself has three components (resources) that work together:
- <k8s>Ingress Resource</k8s> (`kind: Ingress`) - a yaml manifest that declares the routing, which path to which service on which port. this is then stored into the etcd.
- <k8s>Ingress Controller</k8s> (`kind: IngressController`) - the active pod that's running in the cluster, it configures the actual proxy routing. it creates the IngressClass and watches the Ingress resources.
- <k8s>Ingress Class</k8s> (`kind: IngressClass`) - a mapping between the controller and the ingress resource. it glues a specific behavior to the controller to handle it.

A cluster can run multiple controllers at the same time, so the ingress class determines which controller handles which route. we can mark one class as the default class so all traffic not explicitly claimed by another controller goes through the default one.

Controllers:

- [nginx](https://docs.nginx.com/nginx-ingress-controller/) - for simplicity.
- [traefik](https://traefik.io/solutions/kubernetes-ingress) - for automatic HTTPS.
- [istio](https://istio.io/latest/docs/tasks/traffic-management/ingress/kubernetes-ingress/) - for service mesh.

The analog is thinking of the ingress like DNS: the controller is the DNS server, and the IngressClass is how we choose which DNS server to ask.

we don't need to write the <k8s>IngressController</k8s> ourselves, but it looks like this.

```yaml
# This is created automatically when you install an Ingress Controller -
# you don't need to write it yourself.
apiVersion: networking.k8s.io/v1
kind: IngressClass
metadata:
  name: nginx     # The name you reference in ingressClassName
spec:
  controller: k8s.io/ingress-nginx    # Identifies which controller pod watches for this class
```

the important parts of the document:

> - `metadata.name` - this is the value you put in `ingressClassName: nginx` in your Ingress resources.
> - `spec.controller `- a unique identifier string that the controller pod uses to claim "I handle this class"" Each controller installation registers itself with a specific controller string (name).

<k8s>Ingress</k8s> and LoadBalancers services both expose access for external requestes. In general, we use the Ingress to expose HTTP routing with hostnames and paths, like most web apps use. we use LoadBalncers for non-HTTP traffic (databases, gRPC, custom TCP protocols).

| Operation         | Ingress (L7)                                        | Service: LoadBalancer (L4)  |
|-------------------|-----------------------------------------------------|-----------------------------|
| Layer             | HTTP / HTTPS                                        | TCP / UDP                   |
| Routes by         | Hostname, path, headers                             | IP + port only              |
| Entry points      | One for many services                               | One external IP per service |
| TLS termination   | Yes                                                 | No (passthrough)            |
| Advanced features | Weighted routing, canary deployments, rate limiting | None                        |
| Protocol support  | HTTP/HTTPS only                                     | Any TCP/UDP traffic         |


for our example, we will use the NGINX controller.

```sh
# Install NGINX Ingress Controller (designed for Kind)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for the controller to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=120s

# Verify the IngressClass was created by the install
kubectl get ingressclass

# Switch the controller Service from LoadBalancer to NodePort
# (Kind has no cloud provider to provision a real load balancer)
kubectl patch svc ingress-nginx-controller -n ingress-nginx \
  -p '{"spec":{"type":"NodePort"}}'

# Force the controller onto a worker node
kubectl patch deployment ingress-nginx-controller -n ingress-nginx \
  --type='json' \
  -p='[{"op":"add","path":"/spec/template/spec/nodeSelector","value":{"kubernetes.io/hostname":"training-worker"}}]'

kubectl apply -f ingress.yaml

# Add the hostname to /etc/hosts (requires sudo on macOS)
sudo sh -c 'echo "127.0.0.1 myapp.local" >> /etc/hosts'

# Flush DNS cache on macOS
sudo dscacheutil -flushcache; sudo killall -HUP mDNSResponder

# Test
curl http://myapp.local
```

now we have a challenge fo fix the curl command which doesn't work, need to figure out why. we get "Recv failure: Connection reset by peer" error.
1. is the problem ingress or is it in the services? could I curl the underlying service with the low-level path?

```sh
kubectl run curl --image=curlimages/curl --rm -it -- sh
# Inside: curl http://nginx-svc.default.svc.cluster.local
# The format is: <service-name>.<namespace>.svc.cluster.local
```
this seems to work.\
maybe this is a namepsace issue? probably not...\
why do we get this error?\
looking at headlamp, we have an event "Scheduled for sync", maybe this is important?\
but if the trainning is about ingresses, won't the problem be in the ingress resource? does the ingress resource refernce the correct ingressClass name?

maybe it's because our service is clusterIP and not note port?

```sh
kubectl patch svc nginx-svc -p '{"spec":{"type":"NodePort"}}'
kubectl get svc nginx-svc
curl http://myapp.local 
```
still no match. revert back to clusterIP?

maybe we need to reload the ingress controller somehow?

```sh
kubectl rollout restart deployment -n ingress-nginx ingress-nginx-controller
```
still nothing, now we see another `SYNC` in the ingress, but no matches.

maybe it's the re-write target that does something stupid?
```sh
kubectl annotate ingress nginx-ingress nginx.ingress.kubernetes.io/rewrite-target-
```

nope, we apply the ingress again.\
maybe it's because we are missing the default controller somehow?

gave up, asked claude for hints:

> (ME) give me a hint for figuring out the challenge in module 5. do not give me a full answer, just suggest a few directions. be vauge.
>
> (CLAUDE) The challenge is that traffic enters the cluster at one specific node, but the thing that handles it is running somewhere else.
>
> A few directions to explore:
>
> - Where does curl http://myapp.local actually enter the  cluster? Which node?
> - Where is the Ingress Controller pod currently running? Which node?
> - Look at the Kind config again - what's special about  the control-plane node that the worker nodes don't have?


maybe the the problem is that we force the contoller on the trainning worker, but we set the control plane node to take all ingress traffic?

```yaml
nodes:
  - role: control-plane
    kubeadmConfigPatches:
 - |
   kind: InitConfiguration
   nodeRegistration:
     kubeletExtraArgs:
  # Label this node so the NGINX Ingress Controller knows
  # to bind to it (used in Module 2)
  node-labels: "ingress-ready=true"
```
and then we did this
```sh
# Force the controller onto a worker node
kubectl patch deployment ingress-nginx-controller -n ingress-nginx \
  --type='json' \
  -p='[{"op":"add","path":"/spec/template/spec/nodeSelector","value":{"kubernetes.io/hostname":"training-worker"}}]'
```

so maybe we need to replace it to run the controller on the control-plane-node?

```sh
kubectl patch deployment ingress-nginx-controller -n  ingress-nginx \
    --type='json' \
    -p='[{"op":"replace","path":"/spec/template/spec/nodeSelector","value":{"kubernetes.io/hostname":"training-control-plane"}}]'
```
ok, this worked. after claude gave me the answer.

the answer the document wanted was to use the `ingress-ready` nodeSelector value, which makes more sense.

```sh
kubectl patch deployment ingress-nginx-controller -n ingress-nginx \
  --type='json' \
  -p='[{"op":"replace","path":"/spec/template/spec/nodeSelector","value":{"ingress-ready":"true"}}]'

# Wait for the controller to restart on the correct node
kubectl rollout status deployment ingress-nginx-controller -n ingress-nginx

# Verify the controller is now on the control-plane
kubectl get pods -n ingress-nginx -o wide

# Now it works
curl http://myapp.local
```

this is because only the control-plane node exposes the ports, but we told the controller to start on a worker node.

</details>

We next move to other resources, like how a deployment manages pods through replicaSets, so do <k8s>StatefulSets</k8s> and <k8s>DaemonSets</k8s>\
A StatefulSet creates pods with stable, consistent and predictble names, rather than just random names. this allows us to have a stable name that database replicas can look at, and we can configure consistent storage for each pod that is unique to the pod, even after restarts (by using <k8s>PersistentVolumeClaim</k8s>).\
This works great with headless services (which hand out all the address), so each pod name becomes a stable dns entry. in this scenario, database replicas always know where to point at, since the name is consistent and known. 

```sh
kubectl apply -f headless-service.yaml
kubectl apply -f statefulset.yaml
kubectl get pods -w
# Watch: pods are created in ORDER: redis-0, then redis-1, then redis-2
# Each has a stable DNS name: redis-0.redis.default.svc.cluster.local

# Resolve each pod's DNS name - each should return a different IP
kubectl run dns-test --image=busybox:1.36 --rm -it --restart=Never -- sh -c \
  "nslookup redis-0.redis.default.svc.cluster.local && \
   nslookup redis-1.redis.default.svc.cluster.local && \
   nslookup redis-2.redis.default.svc.cluster.local"

# Cross-check: the IPs from nslookup should match the pod IPs here
kubectl get pods -l app=redis -o wide
```

<k8s>DaemonSets</k8s> offer something else, they are used for 'infrastrucure' workloads that must run on every node, this might be log collection, metrics exporting, network plugin, etc...\
Deployment have a set number of instances, and there's no guarntee that every node will run a pod and that no node will have more than one pod running. The DaemonSet is designed to run one pod on every node, adding and removing them as needed when nodes join and leave the cluster.

```sh
kubectl apply -f daemonset.yaml

# One pod per worker node (control-plane has a NoSchedule taint)
kubectl get pods -o wide -l app=log-collector
```

| Feature      | Deployment                         | StatefulSet                   | DaemonSet                         |
|--------------|------------------------------------|-------------------------------|-----------------------------------|
| Pod Identity | Random names                       | Stable ordinal (pod-0, pod-1) | One per node                      |
| Storage      | Shared / ephemeral                 | Per-pod PVC                   | Typically hostPath                |
| Scaling      | Any replica count                  | Ordered scale up/down         | Follows node count                |
| DNS          | Via Service only                   | Per-pod DNS (pod-N.svc)       | Via Service only                  |
| Use Case     | Stateless apps (web servers, APIs) | Databases, queues, clusters   | Node agents, log collectors, CNIs |


> What about <k8s>Jobs</k8s> and <k8s>CronJobs</k8s>?\
> Kubernetes also provides Jobs and CronJobs for finite, run-to-completion workloads. While Deployments, StatefulSets, and DaemonSets keep Pods running indefinitely, a Job creates one or more Pods that run a task until it succeeds and then stops - think database migrations, batch processing, or one-off scripts.
>
> A Job adds retry logic, completion tracking, and optional parallelism on top of a bare Pod.\
> A CronJob is simply a Job on a schedule (e.g., "run this backup every night at 2 AM").\
> Jobs and CronJobs are out of scope for this course, but you may see them listed in tools like Headlamp and k9s.

some clean up

```sh
kubectl delete -f statefulset.yaml
kubectl delete -f headless-service.yaml
kubectl delete -f daemonset.yaml
# StatefulSet PVCs are not automatically deleted - clean them up manually
kubectl delete pvc -l app=redis
```



<k8s>ConfigMaps</k8s> store non-sensative configuration data as key:value maps, they aren't coupled to any container image, and the data can be mounted as environment variables or as files.\
we have an example of a pod we both forms. using `envFrom` for the config as environment variables, and `volumeMounts` for the config as files.

```sh
kubectl apply -f app-env-config.yaml
kubectl apply -f app-file-config.yaml
kubectl apply -f configmap-deploy.yaml

# Check env vars injected from app-env-config
kubectl exec deploy/configmap-demo -- env | grep -E "APP_ENV|LOG_LEVEL"

# Check the mounted config file from app-file-config
kubectl exec deploy/configmap-demo -- cat /etc/app/config.json
```

(we can also check them through headLamp and k9s).

for sensitive data, we don't use configMaps, we use <k8s>secrets</k8s>. they are similiar, but are base64-encoded in etcd. there are some basic secret storage types, with `Opaque` being the type for arbitrary user-defined data, and some other types for common use cases:

- docker config file, either in cfg or json form.
- user/password basic auth credentials
- ssh authentication credentials
- tls credentials
- bootstrap token

we get the secrets onto the container the same way as we did with the configMap, either as env variables (`envFrom` and `secretRef`) or as volume mounts. when we mount secrets inside the container, the data is stored in memory only, never written to disk on the node.

```sh
kubectl apply -f secret.yaml
kubectl apply -f secret-deploy.yaml

# Verify Secret env vars are present (plain text inside the pod)
kubectl exec deploy/secret-demo -- env | grep -E "DB_PASSWORD|API_KEY"

# Verify the mounted secret files
kubectl exec deploy/secret-demo -- ls /etc/secrets
kubectl exec deploy/secret-demo -- cat /etc/secrets/DB_PASSWORD
```

so far, we used Volume and mounts for configurations, but we can use them for so much more.

| Volume Type                              | Description                                              | Persists After Pod Delete?  |
|------------------------------------------|----------------------------------------------------------|-----------------------------|
| <k8s>emptyDir</k8s>                      | Temp dir shared between containers in a pod              | No                          |
| <k8s>hostPath</k8s>                      | Mounts a path from the host node's filesystem            | Yes (on that specific node) |
| <k8s>persistentVolumeClaim</k8s>         | Claims a PersistentVolume (decoupled from pod lifecycle) | Yes                         |
| <k8s>configMap</k8s> / <k8s>secret</k8s> | Mounts config/secret data as files inside the container  | N/A (data lives in etcd)    |

usually, the data inside the pod 'resets' once the pod dies down. there are cases where we want the data to persist across pod sessions. for these cases, we can set up persistent storage, which matches a <k8s>persistentVolume</k8s> and a <k8s>PersistentVolumeClaim</k8s>. thd data lives separatly from the workload which uses it.

we define the <k8s>PersistentVolume</k8s> resource and the accompanying <k8s>PersistentVolumeClaim</k8s>. we set the accessmode to control how many nodes can read or write to the volume, and we set the storage amount and path.\
For the workload, we set the volume to the claim that we made (not to the storage directly). the data will persist across pod restarts and appear in each pod file system, it is no longer ephemeral.

```sh
kubectl apply -f persistent-volume.yaml
kubectl apply -f volume-deploy.yaml
kubectl get pv,pvc

# Write data to the persistent volume
kubectl exec deploy/volume-demo -- sh -c "echo 'hello from k8s' > /app/data/test.txt"

# Read it back
kubectl exec deploy/volume-demo -- cat /app/data/test.txt

# Restart the pod and verify data survives
kubectl rollout restart deploy/volume-demo
kubectl rollout status deploy/volume-demo
kubectl exec deploy/volume-demo -- cat /app/data/test.txt
```

cleanups:

```sh
# ConfigMaps section
kubectl delete -f configmap-deploy.yaml
kubectl delete -f app-env-config.yaml
kubectl delete -f app-file-config.yaml

# Secrets section
kubectl delete -f secret-deploy.yaml
kubectl delete -f secret.yaml

# Volumes section
kubectl delete -f volume-deploy.yaml
kubectl delete -f persistent-volume.yaml
```

## Demo App

<details>
<summary>
A demo Golang app.
</summary>

This section will walk through building, containerizing and deploying a demo golang application. we will also expose it via <k8s>Ingress</k8s>.

for most (if not all) applications, we will need to set up a few things:

1. Dockerfile to package our our app into a container.
1. Readyness probe (`/readyz`) - an endpoint that must be hit before the cluster starts routing traffic to the application, making sure we don't start work until the application is ready.
1. Liveness probe (`/healthz`) - an endpoint that detects if the app is healthy or stuck in a deadlock, so that the cluster could restart the pod.
1. Gracefull shutdown - when K8s kills the pod, it send a `SIGTERM` signal, we want the application to know what to do in this case, and not just shutdown immediatly. this usually means that we stop accepting new requests, complete all the in-flight requests, and then shutdown.
1. Resource request and limits - we set the resources we allow the workload to use so the <k8s>scheduler</k8s> can place it correctly, and so that one bad pod can't starve the entire cluster.


let's start building the app

```sh
mkdir -p demo-app && cd demo-app

# Build the Docker image and load it into Kind
# (Kind nodes can't pull from Docker Hub by default for local images)
docker build -t demo-app:1.0 .
# verify the image is built locally
docker image ls demo-app

# load onto kind
kind load docker-image demo-app:1.0 --name training
# verfiy it's been loaded
docker exec training-control-plane crictl images

# make sure we still have an ingress controller
kubectl get ingressclass nginx
```

next, we create the workload to run the image (<k8s>Deployment</k8s>), a <k8s>ConfigMap</k8s>, a <k8s>Service</k8s> expoisng port 80 and directing the traffic to port 8080 on the pod, and an <k8s>Ingress</k8s> resource to send traffic via the `demo.local` path.\
The deployment sets the `readinessProbe` and `livenessProbe` to the application paths, and we set the imagePullPolicy to never download the image from registry (since it's a local image, and we loaded it directly with `kind load`).

now that we are ready, we can test it.

```sh
kubectl apply -f demo-app.yaml

# Add the hostname to /etc/hosts
sudo sh -c 'echo "127.0.0.1 demo.local" >> /etc/hosts'
sudo dscacheutil -flushcache; sudo killall -HUP mDNSResponder

# Test - notice different hostnames (load balanced across pods)
curl http://demo.local
curl http://demo.local
curl http://demo.local
```

now we patch the <k8s>configMap</k8s> and switch the version, then we rollout a new version of the deployment to see the change take effect.

```sh
# check before
kubectl get configmap demo-config -o=jsonpath='{.data}'
# path
kubectl patch configmap demo-config -p '{"data":{"APP_VERSION":"2.0.0"}}'
#check after
kubectl get configmap demo-config -o=jsonpath='{.data}'
# check to see the pods still return the previous value
curl http://demo.local
# restart deployment
kubectl rollout restart deploy/demo-app
# check pods status after running deployment
kubectl get pods -w
# check pods
kubectl describe pod <pod-name>
# check to see the pods now return the new value
curl http://demo.local
```

breaking challange, set the liveness path to something else (from `/readyz` to `/nonexistent`). the document said to modify the file, but i'll try using `patch` instead.

```sh
# before
kubectl get deploy demo-app -o jsonpath='{.spec.template.spec.containers[0].readinessProbe.httpGet.path}'
# check the replacement syntax
kubectl patch deploy demo-app\
  --dry-run=server \
  --type='json' \
  -p='[{"op":"replace","path":"/spec/template/spec/containers/0/readinessProbe/httpGet/path","value":"/nonexistent"}]'

# perform the replacements
kubectl patch deploy demo-app\
  --type='json' \
  -p='[{"op":"replace","path":"/spec/template/spec/containers/0/readinessProbe/httpGet/path","value":"/nonexistent"}]'
# check after
kubectl get deploy demo-app -o jsonpath='{.spec.template.spec.containers[0].readinessProbe.httpGet.path}'
# check pods
kubectl get pods -w
# describe events
kubectl describe pod <pod-name>
```

because I used patch directly, the new pod exists alongside the previous pods, kubernetes waits for the new pod to reach the healthy status before removing the older pods. that's why curl commands still worked for me. the instructions are to remove the deployment and then modify the file, but I didn't do it like that.

we can now remove the demo app
```sh
kubectl delete -f demo-app.yaml
```

1. if the healthProb fails, the cluster sends a killTerm command. if the readiness probe fails, the cluster waits and doesn't send traffic yet, but it's ok for the readiness to be false.
1. we don't want to get the images from a registry, we load the directly onto the nodes (containers). we can't load images from the local docker daemon, so we pre-load them.
1. probably each nodes (which is dokcer) has it's own docker-runtime, so we we load the image onto that docker daemon, so it will have them there.\
_Answer: "It exports the image as a tar archive from your local Docker daemon, then imports it into the containerd runtime running inside each Kind node container."_

we didn't do any graceful shutdown or limits in our demo app. maybe later.

if would wait for signal, capture it on a channel, cancel running contexts, remove resources if the application has taken any (files, connections, etc...). we have some examples in our code base.

</details>

## Kustomize & Helm

<details>
<summary>
Managing Kubernetes at Scale.
</summary>

So far, we used raw yaml files, applied them directly, removed them in reverse order, and modified them in the files or patched them inplace. this isn't feasible in production grade clusters. for real development, we have the same applications running on different layers (Dev, Staging, Production) or in different data residencies (US, EU, etc..), each of these clusters needs to be identical in terms of workloads, but differes in the configurations or settings (resources, number of replicas).\
Creating a copy for each cluster is 100% going to lead to configuration drift, small mistakes and omissions will cause the clusters to end up being different, and the behavior will change between them.\
The solution `or this problem is to add another layer, some tool that manages all the files for the cluster, and allows for replicating the same strucutre without replicating the files. <k8s>Kustomize</k8s> and <k8s>Helm</k8s> are two tools for doing just that.

<k8s>Kustomize</k8s> is built directly into the <k8s>kubectl</k8s> cli tool (no need to install anything new), and it uses a _base/overlay_ pattern. the base is the manifest yaml files, and the overlays are the per environment patches.\
Each directory has a `kustomization.yaml` file with the resources and what patches apply for them. when we run `kubectly apply --kustomize <dir>` the changes are merged with the base manifests to produce the final yaml files.

- base - the directory with the shared manifests yaml files and the kustomization file listing them.
- overlay - a directory that references the base and adds patches for the specific environment.
- startegic merge patch - a partial yaml file that only has the fields we want to change, with the rest being inherited from the base.
- <k8s>ConfigMap</k8s>, <k8s>Secret</k8s> generators - <k8s>kustomize</k8s> can automatically genereate configurations and secrets with hash suffixes, when they change, it triggers a rolling update (no need for manual deployment rollout command).

we will start a new application, it has a base layers and two overlays (dev and production).

folder structure:

> - my-app/
>   - base/
>     - deployment.yaml
>     - service.yaml
>     - kustomization.yaml
>   - overlays/
>     - dev/
>       - kustomization.yaml
>       - replica-patch.yaml
>     - prod/
>       - kustomization.yaml
>       - replica-patch.yaml

we have a <k8s>Deployment</k8s> running nginx image (rwo replicas, port 80, setting reource requests and limits), a <k8s>Service</k8s> for internal traffic (`type: ClusterIP`) direcrting traffic to port 80 on the target pods using a selector to target resources with `app:nginx`. finally, we have a <k8s>Kustomization</k8s> resource which "exposes" the two files for patching.\
In the overlays directories, we have partial yaml files, one for the deployment which sets the number of replicas abd the image (we can set a different tag, or even a debug image). and the kustomization file for the overlay, with resources being the same 'base' layer, patches directing to the local files containing the patches, and a namespace field, which sets it on all resources (allows for partition).

let's see this in action.

```sh
# Create namespaces
kubectl create namespace dev
kubectl create namespace prod

kubectl get namespaces

# Preview what Kustomize produces for each environment (no changes applied yet)
kubectl kustomize my-app/overlays/dev/
kubectl kustomize my-app/overlays/prod/

# Compare the two outputs side by side
diff <(kubectl kustomize my-app/overlays/dev/) \
     <(kubectl kustomize my-app/overlays/prod/)

# Apply both environments
kubectl apply -k my-app/overlays/dev/
kubectl apply -k my-app/overlays/prod/

# Verify dev: 1 replica, nginx:latest
kubectl get deploy,pods -n dev

# Verify prod: 3 replicas, nginx:1.25.3
kubectl get deploy,pods -n prod
```

if we modify that patch file `metadata.name` to something wrong ("nginx-typo") - the patch will fail with the following error:

> error: no resource matches strategic merge patch "Deployment.v1.apps/nginx-typo.[noNs]": no matches for Id Deployment.v1.apps/nginx-typo.[noNs]; failed to find unique target for patch Deployment.v1.apps/nginx-typo.[noNs]

this indicates the kustomize couldn't find any matching targets with that name to apply the changes on. there is not deployment with that name.

let's clean this up:

```sh
kubectl delete -k my-app/overlays/dev/ 2>/dev/null
kubectl delete -k my-app/overlays/prod/ 2>/dev/null
kubectl delete namespace dev 2>/dev/null
kubectl delete namespace prod 2>/dev/null
```

<k8s>Helm</k8s> offers a different Solution to the same problem. instead of patching YAML files, helm uses templating. it uses go-templates to generate manifests and fills-in the required values from variables.

> - Kustomize = "here's my YAML, patch these fields".
> - Helm = "here's a template, fill in these variables".

Helm is an external tool, it uses <k8s>Charts</k8s> to generate manifests from templates (yaml files) based on values, it then uses <k8s>releases</k8s> to track all the changes. The chart is a package of the manifests and the templates, a blueprint. a release is a running instance of a chart. charts can be stored on registries.

> - demo-app/
>   - Chart.yaml            <-- Chart metadata (name, version, description)
>   - values.yaml           <-- Default configuration values
>   - templates/
>     - deployment.yaml     <-- Deployment template (with Go templating)
>     - service.yaml        <-- Service template
>   - charts/               <-- Subdirectory for dependency charts


the helm cli tool must be installed separtly, and it uses the same kubeconfig as kubectl to point itself at a cluster.

we will start by adding a chart repository, updating it, and then installing the chart onto our cluster.

```sh
# Register the metrics-server repository under the alias "metrics-server"
# (stored locally in ~/.config/helm/repositories.yaml)
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/

# Download the latest chart list from all registered repositories
helm repo update

# Search for charts in your registered repositories
helm search repo metrics-server

# Install metrics-server with overrides
helm install my-metrics metrics-server/metrics-server \
  --set replicas=2 \
  --set args='{--kubelet-insecure-tls}'

# Verify the release
helm list
kubectl get pods -l app.kubernetes.io/instance=my-metrics
kubectl get svc -l app.kubernetes.io/instance=my-metrics
```

now we will use helm to upgrade the chart release with new values from a file

```sh
# Upgrade the release with the new values file
helm upgrade my-metrics metrics-server/metrics-server \
  --set args='{--kubelet-insecure-tls}' \
  -f my-prod-values.yaml

# Check revision history
helm history my-metrics

# Rollback to the original installation (revision 1)
helm rollback my-metrics 1
kubectl get pods -l app.kubernetes.io/instance=my-metrics   # Back to 2 replicas

# Clean up
helm uninstall my-metrics
```

now we create our own chart. a "Chart.yaml" file with metadata about the chart, a "values.yaml" with the replacement values, and inside the "templates" directory, yaml files represnting the resources with templates to fill up values.\
we have templates for `{{ .Release.Name }}` which takes from the helm release, and templates like `{{ .Values.app.name }}` which take from the passed values. there are also other [built-in helm objects](https://helm.sh/docs/chart_template_guide/builtin_objects/) to use, but we won't be looking into them for this course.

- Release
- Values
- Chart
- Subcharts
- Files
- Capabilites
- Template

helm templaes use golang style, and we can use condition `{{ if <condition>}} {{ else }} {{ end }}`, and all sorts of other functionalities like `toYaml` (which is very much used) and `nindent` to control indentation. it's easy to make mistakes with indentation of yaml files.\
lucklily for us, we can preview the generated files and run them in dry-run mode to make sure we are still formatted correctly.

```sh
# Preview the rendered templates (nothing is applied yet)
helm template my-release ./demo-app

# Dry-run against the cluster API (validates resource schemas)
helm install my-release ./demo-app --dry-run=client --debug

# Install for real
helm install my-release ./demo-app

# Verify
helm list
kubectl get deploy,pods,svc

# Upgrade: change replica count
helm upgrade my-release ./demo-app --set replicaCount=3
kubectl get pods    # Expect 3 pods

# Check revision history
helm history my-release

# Rollback to revision 1
helm rollback my-release 1
kubectl get pods    # Back to 2 pods
```

now, we break the release by using a bad values file.

```sh
helm install test ./demo-app -f demo-app/bad-values.yaml
```

there was no match for "replicacount", so it was ignored, it fell back to using "values.yaml" for the values (such as app name), so the error was silent. we can use a "values.schema.json" file to enforec constratints on the values before rendering, this will happen during `helm install`, `helm upgrade` and `helm lint`.

cleanup

```sh
helm uninstall my-metrics 2>/dev/null
helm uninstall my-release 2>/dev/null
helm uninstall test 2>/dev/null
```

| Aspect                | Kustomize                                         | Helm                              |
|-----------------------|---------------------------------------------------|-----------------------------------|
| Approach              | Patch plain YAML                                  | Go template engine                |
| Built into kubectl    | Yes (`kubectl apply -k`)                          | No (separate helm CLI)            |
| Learning curve        | Low - just YAML                                   | Medium - Go templates             |
| Best for              | Environment-specific tweaks to existing manifests | Reusable, distributable packages  |
| Dependency management | No                                                | Yes (subcharts, Chart.lock)       |
| Rollback              | Manual (kubectl / GitOps)                         | Built-in (`helm rollback`)        |
| Release tracking      | No                                                | Yes (revisions stored in cluster) |

many teams use both together, helm for outside packages (databases, monitoring, ingress controllers), and kustomize to manage the internal application manifests.

1. Kustomize base includes the complete (shared, env agnoistic) yaml manifest resources, with just an extra kustomization file. the overlay includes the partial yaml files for specific resources patching, with a kustomization adding data for all matching files (such as namespace, or any other labels).
1. helm chart is a packaged collection of manifests, which can be used to generate releases throught templates.
1. `helm-template` just renders the template, `helm install --dry-run` tries checking it against the server and the scehma, which might have constraints or if we don't have the correct namespace, permissions, or other reasons.
1. helm offers release tracking, distribution and strong templateing capabilites. kustomize is built-in, fast and simple.

</details>

## Controllers And Operators

<summary>
Extending the cluster
</summary>

some advnace kubernetes features.

> A controller is a control loop that continuously watches the state of Kubernetes resources and takes action to move the current state toward the desired state.\
> Controllers are the backbone of Kubernetes - almost everything in the cluster works via controllers.

<k8s>Controllers</k8s> are what handles the _reconsilliation loop_, so when we create a deployment with 3 replicas, the controller is what ensures our cluster matches the required state, and if we ever drift from it (one pod goes down), the controller is the one which creates the replacment.


> Examples of built-in controllers:
>
> - <k8s>ReplicaSet</k8s> controller - ensures the correct number of Pod replicas are running.
> - <k8s>Deployment</k8s> controller - manages ReplicaSets and handles rolling updates.
> - <k8s>Service</k8s> controller - allocates cloud load balancers for LoadBalancer-type Services.

controllers react to changes in the current state (level-triggered), they compare the current state to some desired state, and act accordingly. this is in contrast to "edge-triggered" systems, which reacts to events. this means that even if a controller crashes and needs to restart, it can always check the current state against the desired state. the controller doesn't need to exist before the resource it manages on.\

> An <k8s>Operator</k8s> is a controller that manages Custom Resources and encodes application-specific operational logic. While built-in controllers handle generic K8s operations, operators capture the expertise needed to run a specific application - how to deploy, configure, scale, backup, and restore it.\
> Operator = Custom Controller + CRD. Every operator IS a controller, but not every controller is an operator. The term "operator" specifically means a controller that manages a Custom Resource and encodes operational knowledge (how to deploy, scale, backup, restore, etc.).

| Aspect          | Controller                                   | Operator                                                |
|-----------------|----------------------------------------------|---------------------------------------------------------|
| What it watches | Built-in resources (Pods, Deployments)       | Custom Resources (CRDs you define)                      |
| Knowledge       | Generic K8s operations                       | Application-specific logic (e.g., "how to backup a DB") |
| Examples        | ReplicaSet controller, Deployment controller | Prometheus Operator, Cert-Manager, Strimzi (Kafka)      |

> Real-world examples:
>
> - Prometheus Operator - you declare a Prometheus CR, the operator deploys and configures a full monitoring stack.
> - Cert-Manager - you declare a Certificate CR, the operator obtains and renews TLS certificates from Let's Encrypt.
> - Strimzi - you declare a Kafka CR, the operator deploys and manages an entire Kafka cluster with topics, users, and rebalancing.

Custom Resources and resource which aren't built in to kubernetes. the built-in resources are Pods, Services, Deployments, Configmaps, Secrets, but custom resources can be websites, database clusters or certificates. we use <k8s>Custom Resource Defintions</k8s> to define them on our kubernetes cluster, and the we use <k8s>Operators</k8s> to act as custom controllers. the Custom Resource Defintion is the 'blueprint' that defines the resource, with the custom resource simply being an instance of the definition.

Custom resources allow us to extend kubernetes without forking the source code and breaking apart from it, we can use them to make the cluster managae multiple kinds of resources, and not just containers.\
The promethous operator stores alert rules as custom resources, the cert-manager operator stores certificates as custom resources, so we too, could store our unique resources and manage them directly with kubernetes.\
Custrom resources are used when the resource has domain-specific lifecycle management operations, when we want to have a declerative style api for our infrstructure, and for operations which go beyond the initial deployment and upgrade flows (health checks, montioring, self healing, rolling updates). 

> Rule of thumb: If you can express your workload entirely with built-in resources (<k8s>Deployments, Services, ConfigMaps</k8s>), a controller or even <k8s>Helm/Kustomize</k8s> is enough.\
> If your application needs custom reconciliation logic - decisions that depend on the application's internal state - that's when an operator adds value.

### Website Operator

a minimal excersize for build a CRD and an operator.

```sh
mkdir -p website-operator && cd website-operator
go mod init website-operator
```

we create a <k8s>Custom Resource Definition</k8s> (`kind: CustomResourceDefinition`) and apply it to our cluster. teaching our cluster what this is. there are many fields here, including names (what `kubectl get <>` will recognize), display formats (what columns will be displayed), versioning, schema validation, and properties.

```sh
# Apply the CRD
kubectl apply -f crd.yaml

# Verify it exists
kubectl get crd websites.training.example.com

# Apply the Website custom resource
# This creates a Website object in the cluster, but nothing will happen yet
# because we haven't built the operator that watches for it.
kubectl apply -f website.yaml

# View it
kubectl get websites
kubectl get ws          # shortname works too
kubectl describe ws my-site

# Clean up before building the operator - this CR was just for exploration
kubectl delete -f website.yaml
```

now we build the actual operator, this is the <golang>go</golang> file.

> `WebsiteSpec` struct — defines the fields our operator reads from the CR.

the websiteSpec matches what we defined in the crd.
```go
// WebsiteSpec matches our CRD's spec fields
type WebsiteSpec struct {
    Image    string `json:"image"`
    Replicas int32  `json:"replicas"`
}
```
> `informer.AddEventHandler(...)` — how the operator watches for Website resource changes.

```go
 // Register event handlers - these are called when a Website is
 // created, updated, or deleted
 informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
  AddFunc: func(obj interface{}) {
   u := obj.(*unstructured.Unstructured)
   log.Printf("Website ADDED: %s/%s", u.GetNamespace(), u.GetName())
   reconcile(clientset, dynClient, u)
  },
  UpdateFunc: func(oldObj, newObj interface{}) {
   u := newObj.(*unstructured.Unstructured)
   log.Printf("Website UPDATED: %s/%s", u.GetNamespace(), u.GetName())
   reconcile(clientset, dynClient, u)
  },
  DeleteFunc: func(obj interface{}) {
   u := obj.(*unstructured.Unstructured)
   log.Printf("Website DELETED: %s/%s", u.GetNamespace(), u.GetName())
   // No manual cleanup needed - OwnerReferences handle garbage collection
  },
 })
```

> - `reconcile` function — the core logic that ensures a Deployment + Service exist for each Website.
> - `OwnerReferences` — how child resources get automatically cleaned up when the parent Website is deleted.

(big methods)

it seems that we create an expected deployment and service, and then try and find them based on the name. if they don't exist, we create them from scratch, otherwise, we update them. there's also some behavior to make the operator either from the pod or on the local machine.


lets create the operator and run it.
```sh
# Install dependencies and run the operator
go mod tidy
go run main.go

# different terminal
kubectl apply -f hello-site.yaml
kubectl get deploy,svc,ws

kubectl delete -f hello-site.yaml
kubectl get deploy,svc  # The Deployment and Service are gone!
```

in this demo, we run the operator on the local machine, but we want it to run inside the cluster. that means we need to containerize it somehow. so let's create the dokcerfile for it, and load it on the the kind cluster. we will also need a <k8s>Resource Based Access Control</k8s> (RBAC) to give the operator permissions to watch the cluster, and we will run the operator as a <k8s>Deployment</k8s>


the RBAC resource creates a service account with permissions, in our case, it mnages the <k8s>Deployments</k8s>, <k8s>Services</k8s> and <k8s>websites</k8s> resources, and it can read, list, watch, create, update and delete them. we then use the <k8s>ClusterRoleBinding</k8s> to match the role and the user (website operator).
```sh
docker build -t website-operator:latest .
kind load docker-image website-operator:latest --name training

# Apply RBAC and deploy the operator
kubectl apply -f rbac.yaml
kubectl apply -f operator-deployment.yaml

# Verify the operator Pod is running
kubectl get pods -l app=website-operator

kubectl logs -l app=website-operator -f
```
and we can test is like before, loading the resources from a different template.

```sh
kubectl apply -f hello-site.yaml
kubectl get deploy,svc,ws

# Check operator logs to see the reconciliation
kubectl logs -l app=website-operator
```

1. controller is the general name, operator is a controller for custom resources (has a custrom resource defintion).
1. OwnerReferences tell resources who owns them, so if they no longer exist, they also get deleted. this way the operator only deletes it's own stuff, and everything that is in a lower level gets deleted in the next reconsilliation loop (since the owner no longer exists).\
_Answer: actually, this is done tthe k8s garbage collector_.
1. level triggered checks current state vs desired state, edge trigger checks events. for kubernetes it means that we are ok with missing events or having the operators come up after the resources they manage, since they will just handle them in the next iteration.
1. no, when the operator comes back up it will check the state and try to fix it to reach the desired state.


cleanup

```sh
# Delete custom resources (child Deployments/Services are garbage collected)
kubectl delete ws --all

# Delete the in-cluster operator and RBAC
kubectl delete -f operator-deployment.yaml
kubectl delete -f rbac.yaml

# Delete the CRD
kubectl delete -f crd.yaml
```

</details>

## Container Network Interface (CNI)

<details>
<summary>
Networking 
</summary>

The <k8s>Container Network Interface</k8s> is a specification about how the container runtime confgures network interfaces in linux containers, all the networking of the cluster is handled by the CNI, not by kubernetes. this means we can swap in different implementations of CNIs without changing kubernetes itself.

> - create veth (virtual ethernet)
> - attach to bridge
> - assigne IP
> - setup routes

The Kubernetes model specifies regarding networking:

> - Every pod gets its own unique IP address
> - Pods can communicate with each other across nodes without NAT
> - Processes on the node can communicate with pods on that node

however, the kubernetes itself doesn't do any of that, it delegates all work to the CNI. The <k8s>Kubelet</k8s>, which is the primary kubernetes agent that runs on every node, is the bridge between the cluster and the networking.\
The kublet is the agent responisble for managing the pods on its' node. it tells the container runtime to start and stop them, it reports the status of the node and the pods, and handles health checks.\
When a new pod is created, the kubelet tell the container runtime to create a network namespace, and then it tells the CNI to set the pod network. when the pod is deleted, the kubelet tells the CNI to remove that network.

The CNI is configured via a CNI configuration file and plugin binaries, if they are missing, the node won't be marked as `READY` and won't accept pods.

CNI supports four types of operations:

| Operation | When                 | What Happens                                    |
|-----------|----------------------|-------------------------------------------------|
| `ADD`     | Pod is created       | Create network interface, assign IP, set routes |
| `DEL`     | Pod is deleted       | Remove network interface, release IP            |
| `CHECK`   | Periodic / on demand | Verify network setup is still correct           |
| `VERSION` | On init              | Report supported CNI spec versions              |

the cni itself is just an executable plugin which reads the netwrok configuration (via stdin or a JSON file) and the env variables, and run the commands for each operation, and returns the result.\
the network configuration json contains the CNI spec, which binary plugin to run, which bridge to use, and IPAM  configuration (how to allocate ip addresses to pods). 

when a pod is created, the first thing the CNI does is to create a __veth pair__, one side connected to the host node and one side inside the pod. the side inside the pod is moved into the pods' _network namesapce_ (netns) - the isolated network environment of the pod, with a seperate stack of network resources (interfaces, routes, ip addresses) which are all separate from those of the host and the other pods. the CNI then calls the IPAM plugin to allocate the uniqueIP for the pod and configure routing, so the pod could communicate with the rest of the cluster.\
When the pod is deleted, the kubelet calls the container runtime to stop the containers, the CNI is called with the `DEL` operation (same parameters as `ADD`), which removes the veth pair, releases the IPs via the IPAM, and destroys the network namespace.

since the CNI is a specification, there are many implementations for it:

- Kindnet - default CNI, uses point-to-point plugin with host-local. simple and lightweight.
- Calico - supports BGP routing, VXLAN overlayy, EBPF, VPP dataplanes. rich networking policy enforcement.
- Cilium - eBPF native CNI which replaces iptables entirely, has L3-L7 networking policy, built-observability, encryption via wireguard and service mesh capabilites.
 

| CNI     | Data Plane               | Network Policy | Best For                                |
|---------|--------------------------|----------------|-----------------------------------------|
| Kindnet | ptp + bridge             | No             | Local dev (Kind)                        |
| Calico  | BGP / VXLAN / eBPF / VPP | Yes            | Production, enterprise                  |
| Cilium  | eBPF                     | Yes (L3-L7)    | Production, observability, service mesh |

in our excersize, we will build a custom logging CNI.

we install Multus, a daemonset wrapper the wrapps the existing kindnet config and becomes the entry point for the CNI, it will first run our custom CNI, and then delegate to kindnet internally.

```sh
# Install the Multus thick plugin DaemonSet
kubectl apply -f https://raw.githubusercontent.com/k8snetworkplumbingwg/multus-cni/v4.1.0/deployments/multus-daemonset-thick.yml

# Wait for Multus to be ready on all nodes
kubectl -n kube-system rollout status daemonset/kube-multus-ds

# Verify Multus installed its CRD
kubectl get crd network-attachment-definitions.k8s.cni.cncf.io

# Verify Multus inserted itself as the primary CNI config
docker exec training-control-plane ls /etc/cni/net.d/
# You should see 00-multus.conf listed first (before kindnet)
```

now we build our custom logger, it only uses standard go packages.

```sh
mkdir -p cni-logger && cd cni-logger
go mod init cni-logger

# No external dependencies needed - just the Go standard library!
go mod tidy

# Cross-compile for Linux (Kind nodes are Linux containers)
ARCH=$(uname -m); [ "$ARCH" = "arm64" ] && GOARCH=arm64 || GOARCH=amd64
GOOS=linux GOARCH=$GOARCH CGO_ENABLED=0 go build -o cni-logger main.go
```

the next step is to copy the binary from the local machine onto each node.

```sh
# Copy the cni-logger binary to each node
for NODE in training-control-plane training-worker training-worker2; do
  docker cp cni-logger ${NODE}:/opt/cni/bin/cni-logger
  docker exec ${NODE} chmod +x /opt/cni/bin/cni-logger
done

# verify
docker exec training-control-plane ls /opt/cni/bin/

```

but even if the binary is there, we need to tell multusus about it, we register it with a custom resource <k8s>NetworkAttachmentDefinition</k8s>, which maps a name to a cni config type.


```sh
# Apply the NetworkAttachmentDefinition
kubectl apply -f logging-net.yaml

# Verify it exists
kubectl get network-attachment-definitions
```

we next deploy a pod and anotate it with the metadata so it would match the cni.

```sh
# Create the pod
kubectl apply -f multus-test.yaml

# Wait for it to be running
kubectl get pods -w

# Verify the pod is running and check its network status annotation
kubectl get pod multus-test -o jsonpath='{.metadata.annotations.k8s\.v1\.cni\.cncf\.io/network-status}' | python3 -m json.tool

# You should see two entries:
#   "kindnet" as the default network (with a real IP like 10.244.x.x)
#   "logging-net" as the secondary (with the dummy IP 198.51.100.1)
# Note: cni-logger doesn't create a real interface — it only logs.
# The dummy IP appears in the status because our ADD returned it, but no net1 interface exists.

# Read the log file — this is where the magic happens!
# Check whichever node the pod was scheduled to:
NODE=$(kubectl get pod multus-test -o jsonpath='{.spec.nodeName}')
docker exec ${NODE} cat /var/log/cni-logger.log

# now deleting

# Capture the node before deleting
NODE=$(kubectl get pod multus-test -o jsonpath='{.spec.nodeName}' 2>/dev/null || echo "training-worker")

# Delete the pod and verify the DEL log entry
kubectl delete pod multus-test

docker exec ${NODE} cat /var/log/cni-logger.log | tail -15
```

in this example, we copied the cni to the nodes manually, this works in the <k8s>Kind</k8s> cluster, where we have complete control over all the nodes. but in real clusters, like those managed by cloud providers (<cloud>AWS EKS</cloud>, <cloud>GCP GKE</cloud> and <cloud>Azure AKS</cloud>) we don't have access to the nodes themselves.\
this is also a problem when we have autosacling, and nodes are added and removed constantly. we also don't get versionning for rollbacks or other stuff.

The solution is to package the CNI into a <k8s>DaemonSet</k8s> and have the <k8s>init container</k8s> copy the cni to the local host via an <k8s>HostPath</k8s> volume mount (copy the binary to `/opt/cpi/bin` like we did with the `docker cp` command). this is how the production grade CLIs do this.\
Both the init container and the running container use the same image. the init container copies the `multus-shim` onto the host, and the running container starts with the `usr/src/multus-cni/bin/multus-daemon` command.


(copied from headlamp, github seems more advanced as of september 2025, but it doesn't really matter).

```yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: kube-multus-ds
  namespace: kube-system
  uid: f43d5274-39d8-4d9b-96cd-8a3760885309
  resourceVersion: '111362'
  generation: 1
  creationTimestamp: '2026-06-08T09:14:38Z'
  labels:
    app: multus
    name: multus
    tier: node
  annotations:
    deprecated.daemonset.template.generation: '1'
    kubectl.kubernetes.io/last-applied-configuration: >
      {"apiVersion":"apps/v1","kind":"DaemonSet","metadata":{"annotations":{},"labels":{"app":"multus","name":"multus","tier":"node"},"name":"kube-multus-ds","namespace":"kube-system"},"spec":{"selector":{"matchLabels":{"name":"multus"}},"template":{"metadata":{"labels":{"app":"multus","name":"multus","tier":"node"}},"spec":{"containers":[{"command":["/usr/src/multus-cni/bin/multus-daemon"],"env":[{"name":"MULTUS_NODE_NAME","valueFrom":{"fieldRef":{"fieldPath":"spec.nodeName"}}}],"image":"ghcr.io/k8snetworkplumbingwg/multus-cni:snapshot-thick","name":"kube-multus","resources":{"limits":{"cpu":"100m","memory":"50Mi"},"requests":{"cpu":"100m","memory":"50Mi"}},"securityContext":{"privileged":true},"terminationMessagePolicy":"FallbackToLogsOnError","volumeMounts":[{"mountPath":"/host/etc/cni/net.d","name":"cni"},{"mountPath":"/opt/cni/bin","name":"cnibin"},{"mountPath":"/host/run","name":"host-run"},{"mountPath":"/var/lib/cni/multus","name":"host-var-lib-cni-multus"},{"mountPath":"/var/lib/kubelet","mountPropagation":"HostToContainer","name":"host-var-lib-kubelet"},{"mountPath":"/run/k8s.cni.cncf.io","name":"host-run-k8s-cni-cncf-io"},{"mountPath":"/run/netns","mountPropagation":"HostToContainer","name":"host-run-netns"},{"mountPath":"/etc/cni/net.d/multus.d","name":"multus-daemon-config","readOnly":true},{"mountPath":"/hostroot","mountPropagation":"HostToContainer","name":"hostroot"}]}],"hostNetwork":true,"hostPID":true,"initContainers":[{"command":["cp","/usr/src/multus-cni/bin/multus-shim","/host/opt/cni/bin/multus-shim"],"image":"ghcr.io/k8snetworkplumbingwg/multus-cni:snapshot-thick","name":"install-multus-binary","resources":{"requests":{"cpu":"10m","memory":"15Mi"}},"securityContext":{"privileged":true},"terminationMessagePolicy":"FallbackToLogsOnError","volumeMounts":[{"mountPath":"/host/opt/cni/bin","mountPropagation":"Bidirectional","name":"cnibin"}]}],"serviceAccountName":"multus","terminationGracePeriodSeconds":10,"tolerations":[{"effect":"NoSchedule","operator":"Exists"},{"effect":"NoExecute","operator":"Exists"}],"volumes":[{"hostPath":{"path":"/etc/cni/net.d"},"name":"cni"},{"hostPath":{"path":"/opt/cni/bin"},"name":"cnibin"},{"hostPath":{"path":"/"},"name":"hostroot"},{"configMap":{"items":[{"key":"daemon-config.json","path":"daemon-config.json"}],"name":"multus-daemon-config"},"name":"multus-daemon-config"},{"hostPath":{"path":"/run"},"name":"host-run"},{"hostPath":{"path":"/var/lib/cni/multus"},"name":"host-var-lib-cni-multus"},{"hostPath":{"path":"/var/lib/kubelet"},"name":"host-var-lib-kubelet"},{"hostPath":{"path":"/run/k8s.cni.cncf.io"},"name":"host-run-k8s-cni-cncf-io"},{"hostPath":{"path":"/run/netns/"},"name":"host-run-netns"}]}},"updateStrategy":{"type":"RollingUpdate"}}}
spec:
  selector:
    matchLabels:
      name: multus
  template:
    metadata:
      labels:
        app: multus
        name: multus
        tier: node
    spec:
      volumes:
        - name: cni
          hostPath:
            path: /etc/cni/net.d
            type: ''
        - name: cnibin
          hostPath:
            path: /opt/cni/bin
            type: ''
        - name: hostroot
          hostPath:
            path: /
            type: ''
        - name: multus-daemon-config
          configMap:
            name: multus-daemon-config
            items:
              - key: daemon-config.json
                path: daemon-config.json
            defaultMode: 420
        - name: host-run
          hostPath:
            path: /run
            type: ''
        - name: host-var-lib-cni-multus
          hostPath:
            path: /var/lib/cni/multus
            type: ''
        - name: host-var-lib-kubelet
          hostPath:
            path: /var/lib/kubelet
            type: ''
        - name: host-run-k8s-cni-cncf-io
          hostPath:
            path: /run/k8s.cni.cncf.io
            type: ''
        - name: host-run-netns
          hostPath:
            path: /run/netns/
            type: ''
      initContainers:
        - name: install-multus-binary
          image: ghcr.io/k8snetworkplumbingwg/multus-cni:snapshot-thick
          command:
            - cp
            - /usr/src/multus-cni/bin/multus-shim
            - /host/opt/cni/bin/multus-shim
          resources:
            requests:
              cpu: 10m
              memory: 15Mi
          volumeMounts:
            - name: cnibin
              mountPath: /host/opt/cni/bin
              mountPropagation: Bidirectional
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: FallbackToLogsOnError
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: true
      containers:
        - name: kube-multus
          image: ghcr.io/k8snetworkplumbingwg/multus-cni:snapshot-thick
          command:
            - /usr/src/multus-cni/bin/multus-daemon
          env:
            - name: MULTUS_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          resources:
            limits:
              cpu: 100m
              memory: 50Mi
            requests:
              cpu: 100m
              memory: 50Mi
          volumeMounts:
            - name: cni
              mountPath: /host/etc/cni/net.d
            - name: cnibin
              mountPath: /opt/cni/bin
            - name: host-run
              mountPath: /host/run
            - name: host-var-lib-cni-multus
              mountPath: /var/lib/cni/multus
            - name: host-var-lib-kubelet
              mountPath: /var/lib/kubelet
              mountPropagation: HostToContainer
            - name: host-run-k8s-cni-cncf-io
              mountPath: /run/k8s.cni.cncf.io
            - name: host-run-netns
              mountPath: /run/netns
              mountPropagation: HostToContainer
            - name: multus-daemon-config
              readOnly: true
              mountPath: /etc/cni/net.d/multus.d
            - name: hostroot
              mountPath: /hostroot
              mountPropagation: HostToContainer
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: FallbackToLogsOnError
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: true
      restartPolicy: Always
      terminationGracePeriodSeconds: 10
      dnsPolicy: ClusterFirst
      serviceAccountName: multus
      serviceAccount: multus
      hostNetwork: true
      hostPID: true
      securityContext: {}
      schedulerName: default-scheduler
      tolerations:
        - operator: Exists
          effect: NoSchedule
        - operator: Exists
          effect: NoExecute
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 0
  revisionHistoryLimit: 10
status:
  currentNumberScheduled: 3
  numberMisscheduled: 0
  desiredNumberScheduled: 3
  numberReady: 3
  observedGeneration: 1
  updatedNumberScheduled: 3
  numberAvailable: 3

```

before try copying them, lets remove the manual stuff and test pods.

```sh
# Remove the manually-copied binary from all nodes
for NODE in training-control-plane training-worker training-worker2; do
  docker exec ${NODE} rm -f /opt/cni/bin/cni-logger
done

# Verify it's gone
docker exec training-worker ls /opt/cni/bin/ | grep cni-logger
# Should return nothing

kubectl delete pod multus-test --ignore-not-found
kubectl delete net-attach-def logging-net --ignore-not-found
```

now we create a Dockerfile again (using multi stage build), and load the image onto the kind containers.

```sh
# Build the image (run from inside the cni-logger/ directory)
docker build -t cni-logger:daemonset .

# Load it into the Kind cluster (Kind nodes can't pull from your local Docker daemon)
kind load docker-image cni-logger:daemonset --name training
```

next, we create the deamonset resource for the logger. it has an init container that copies the file into the node, and then the container itself which just keeps itself busy. it has some extra stuff like tolerations to install itself on all nodes (and not just worker nodes) and `hostNetwork:true` to use the host networking, since we are a CNI and the node might not have a CNI yet.

```sh
# Deploy the DaemonSet
kubectl apply -f cni-logger-daemonset.yaml

# Wait for it to roll out on all nodes
kubectl -n kube-system rollout status daemonset/cni-logger-installer

# Verify the binary was installed on every node
for NODE in training-control-plane training-worker training-worker2; do
  echo "--- ${NODE} ---"
  docker exec ${NODE} ls -la /opt/cni/bin/cni-logger
done

# You should see the binary on all three nodes
```

now we re-create the <k8s>NetworkAttachmentDefinition</k8s> and a pod, and see if things still work like before.

```sh
# Re-apply the same NetworkAttachmentDefinition from Step 3
kubectl apply -f logging-net.yaml

# Re-create the test pod from Step 4
kubectl apply -f multus-test.yaml

# Wait for it to be running
kubectl get pods -w

# Verify the CNI log was written (proving the DaemonSet-installed binary works)
NODE=$(kubectl get pod multus-test -o jsonpath='{.spec.nodeName}')
docker exec ${NODE} cat /var/log/cni-logger.log

# You should see the same ADD log output as in Step 4
```

and that's all. we are ready to cleanup the entire cluster.

```sh
# Remove test pod and network definition
kubectl delete pod multus-test --ignore-not-found
kubectl delete net-attach-def logging-net --ignore-not-found

# Remove the CNI installer DaemonSet
kubectl -n kube-system delete daemonset cni-logger-installer --ignore-not-found
```

1. four operations, `ADD`, `DELETE` - when pod is created or removed, `CHECK`, and `VERSION`\
_Answer: `VERSION` is called on startup to dusciver CNI spec version and plugin support. `CHECK` is called periodically to make sure the network setup is still correct._
1. the <k8s>kubelet</k8s> is the main agent on the node, it is the bridge from the kubernetes container to the node container runtime, it also delgates work to the CNI as needed. it's also the thing that watchs pods being created or destroyed.
1. Multus enables multiple CNIs by having an ordered list of cnis to ue?\
_Answer: it's actually a metaplugin, it delegates the original CNI to eth0 interface, and then reads all the <k8s>NetworkAttachmentDefinition</k8s> based on the pods annotations to create additional interfaces for each CNI_
1. scaling, rollback, version control, access to the underlying nodes...


</details>

final deleting of course resources:

```sh
# Delete the training cluster
kind delete cluster --name training

# Verify no training containers remain
docker ps --format "table {{.Names}}\t{{.Status}}" | grep training

# Verify the kubectl context was removed
kubectl config get-contexts
```

## Tools and Commands

<details>
<summary>
Tools which are used in training and kubernetes commands.
</summary>

- [Kind](https://kind.sigs.k8s.io/) - running local kubernetes using docker containers as nodes.
- [k9s](https://k9scli.io/) - interactive terminal for managing and observing clusters.
- [headlamp](https://headlamp.dev/) - user friendly kubernetes UI

- `kubectl apply -k <dir>` - apply and kustomize directory.
- `kubectl kustomize <dir>` - view the merged result of the base and overlay before applying patch
- `kubectl delete pod <resource name> --ignore-not-found=false` - the `--ignore-not-found` flag acts like `2>/dev/null` and ignore missing resources when trying to delete.

### Kind

<details>
<summary>
Running kubernetes clusters on docker.
</summary>

Kind stands for "Kubernetes in Docker" (like Dind - "Docker in Dlcoker"). it runs kubernetes clusters by having each Node as a docker container. the nodes can either be control plane nodes or worker nodes.\
we use declerative style for this (just like other kubernetes resources)

```sh
# Create the cluster
kind create cluster --name training --config kind-config.yaml

# Verify all 3 nodes are Ready
kubectl get nodes -o wide

# See the Docker containers backing each node
docker ps --format "table {{.Names}}\t{{.Image}}\t{{.Status}}"


# load image onto cluster
kind load docker-image demo-app:1.0 --name training

# check images on nodes
docker exec training-control-plane crictl images
kubectl get node training-worker -o  jsonpath='{.status.images}'
```

</details>

### k9s

<details>
<summary>
Interactive terminal for managing kubernetes.
</summary>

setting context on the running cluster

```sh
k9s --context kind-training
:pods
:events
```

essential K9s Shortcuts

| Key               | Action                                                    |
|-------------------|-----------------------------------------------------------|
| <kbd>:</kbd>      | Command mode - type resource names: deploy, svc, pods, ns |
| <kbd>/</kbd>      | Filter / search within the current view                   |
| <kbd>Enter</kbd>  | Drill into a resource (e.g., Deployment -> its pods)      |
| <kbd>d</kbd>      | Describe the selected resource                            |
| <kbd>l</kbd>      | View logs of the selected pod                             |
| <kbd>s</kbd>      | Shell into the selected pod                               |
| <kbd>e</kbd>      | Edit the resource YAML in your editor                     |
| <kbd>ctrl-d</kbd> | Delete the selected resource                              |
| <kbd>y</kbd>      | View the full YAML of the selected resource               |
| <kbd>Esc</kbd>    | Go back / cancel                                          |
| <kbd>ctrl-c</kbd> | Exit K9s                                                  |


finding a crashloop pod guide:

> Type <kbd>:pods</kbd> to list all pods, then use <kbd>CrashLoop</kbd> to filter by status.\
> You can also sort by restarts column or type <kbd>:events</kbd> to see BackOff events.\
> Press <kbd>d</kbd> on a pod to describe it and see the crash reason.


view all resources with `:workloads`
</details>

### Headlamp

<details>
<summary>
Kubernetes web UI and a CNCF Sandbox project.
</summary>

installation: 

```sh
brew install --cask headlamp
```

we run it as an application, not a terminal command.


</details>

### Helm

<details>
<summary>
Manage Kuberntes at scale.
</summary>


```sh
# Register the metrics-server repository under the alias "metrics-server"
# (stored locally in ~/.config/helm/repositories.yaml)
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/

# Download the latest chart list from all registered repositories
helm repo update

# Search for charts in your registered repositories
helm search repo metrics-server

helm lint
helm template
helm install my-metrics metrics-server/metrics-server --dry-run

# Install metrics-server with overrides
helm install my-metrics metrics-server/metrics-server \
  --set replicas=2 \
  --set args='{--kubelet-insecure-tls}'

# Verify the release
helm list

# Upgrade: change replica count
helm upgrade my-release ./demo-app --set replicaCount=3


# Check revision history
helm history my-release

# Rollback to revision 1
helm rollback my-release 1

helm uninstall
```

we can use <k8s>Kustomize</k8s> on the output of `helm template`.

</details>


</details>



## Takeaways
<details>
<summary>
stuff worth remembering
</summary>

which cluster am I working against?

```sh
kubectl config current-context
```

running `kubectl top` requires a [metric server](https://github.com/kubernetes-sigs/metrics-server/blob/master/FAQ.md)

```sh
kubectl top nodes
```

operators and controlles resources:

- [CertManager](https://github.com/cert-manager/cert-manager/tree/master)
- [Promethous](https://github.com/prometheus-operator/prometheus-operator)
- [Grafana](https://github.com/grafana/grafana-operator)
- [KodeBuilder Sample project](https://book.kubebuilder.io/introduction.html)
- [Kopf](https://docs.kopf.dev/en/stable/) - python operator framework


[Owner references and garbage collection](https://kubernetes.io/docs/concepts/architecture/garbage-collection/).
- the owner is stored in the `metadata.ownerReferences` field.
- resources must be in the same namespace is their owner.
- cascading deletion (foreground, background)

IPAM - IP Address Management

</details>
