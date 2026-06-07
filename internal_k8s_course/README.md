<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

# Kubernetes Training Course

personal document for the [internal k8s course](https://fluffy-couscous-v3rem2o.pages.github.io/).


In Kubernetes we set the desired state, we use declerative style (not imperative), and the cluster takes care of moving the parts around to get to that state.

Control Plane:

- <k8s>API Server</k8s> - the "front door" - communication goes through it, the kubectl commands and internal .communication inside the clusters. writes to etcd.
- <k8s>etcd</k8s> - key-value store that holds ALL cluster state.
- <k8s>Scheduler</k8s> - decides which node runs which pod.
- <k8s>Controller Manager</k8s> - runs control loops that reconcile desired vs actual state (e.g., "keep 3 replicas running").    

Worker Node:

- <k8s>Kubelet</k8s> - agent the manages pods on this node.
- <k8s>Kubeproxy</k8s> - handles service networking.
- <k8s>Container runtime</k8s> - run the actual containers.


the controller manager runs the _reconciliation loop_ to bring the current state to the desired state.

The Pod is the smallest 'unit' in k8s, it wraps around one or more containers which share network namespace and storage volumes.\
ReplicaSet - a set of identical pods.\
A Deployment is one step above the ReplicaSet, it handles updates, rolling updates and rollbacks.

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

Services exist to communication inside cluster possible, since everytime a pod restarts it gets a new IP. the Service resource provides a stable IP and routes traffic to the current set of matching resources (pods).\
The default ClusterIP service allows internal communication only.\
There are cases where a client needs to connect to a specific pod, such as databases which need to have the same connection to a replica pod, or a redis database that each pod holds a subset of the data. in these cases, the client needs to match to a specific IP. a headless service exposes all the internal ip addresses, without load balancing them. headless services are often used in StatefulSet workloads.


| Type         | Description                                              | Use Case                                                                   |
|--------------|----------------------------------------------------------|----------------------------------------------------------------------------|
| ClusterIP    | Internal cluster IP only (default)                       | Service-to-service communication                                           |
| NodePort     | Exposes on each node's IP at a static port (30000-32767) | Dev/testing external access                                                |
| LoadBalancer | Provisions external load balancer (cloud providers)      | Production external access                                                 |
| ExternalName | Maps to a DNS name (CNAME record)                        | Aliasing external services                                                 |
| None         | Headless Service, no virtual IP                          | when the clients need to connect to specific pods (statefull applications) |

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

NodePort services make internal addresses reachble from outside, but it exposes the address through ports, rather than hostname and paths:

- have: `http://node-ip:31234`
- want: `https://myapp.example.com`

The solution is to use Ingress:

- Host based routing (`api.example.com` vs `app.example.com`)
- Path based routing (`/api` vs `/web`)
- TLS termination

The Ingress itself has three components (resources) that work together:
- Ingress Resource (`kind: Ingress`) - a yaml manifest that declares the routing, which path to which service on which port. this is then stored into the etcd.
- Ingress Controller (`kind: IngressController`) - the active pod that's running in the cluster, it configures the actual proxy routing. it creates the IngressClass and watches the Ingress resources.
- Ingress Class (`kind: IngressClass`) - a mapping between the controller and the ingress resource. it glues a specific behavior to the controller to handle it.

A cluster can run multiple controllers at the same time, so the ingress class determines which controller handles which route. we can mark one class as the default class so all traffic not explicitly claimed by another controller goes through the default one.

Controllers:

- [nginx](https://docs.nginx.com/nginx-ingress-controller/) - for simplicity.
- [traefik](https://traefik.io/solutions/kubernetes-ingress) - for automatic HTTPS.
- [istio](https://istio.io/latest/docs/tasks/traffic-management/ingress/kubernetes-ingress/) - for service mesh.

The analog is thinking of the ingress like DNS: the controller is the DNS server, and the IngressClass is how we choose which DNS server to ask.

we don't need to write the IngressController ourselves, but it looks like this.

```yaml
# This is created automatically when you install an Ingress Controller —
# you don't need to write it yourself.
apiVersion: networking.k8s.io/v1
kind: IngressClass
metadata:
  name: nginx     # The name you reference in ingressClassName
spec:
  controller: k8s.io/ingress-nginx    # Identifies which controller pod watches for this class
```

the important parts of the document:

> - `metadata.name` — this is the value you put in `ingressClassName: nginx` in your Ingress resources.
> - `spec.controller `— a unique identifier string that the controller pod uses to claim "I handle this class"" Each controller installation registers itself with a specific controller string (name).

Ingress and LoadBalancers services both expose access for external requestes. In general, we use the Ingress to expose HTTP routing with hostnames and paths, like most web apps use. we use LoadBalncers for non-HTTP traffic (databases, gRPC, custom TCP protocols).

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
> - Look at the Kind config again — what's special about  the control-plane node that the worker nodes don't have?


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

this is because only the control-plane node exposes the ports, but we told the controller to start on 

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

# Resolve each pod's DNS name — each should return a different IP
kubectl run dns-test --image=busybox:1.36 --rm -it --restart=Never -- sh -c \
  "nslookup redis-0.redis.default.svc.cluster.local && \
   nslookup redis-1.redis.default.svc.cluster.local && \
   nslookup redis-2.redis.default.svc.cluster.local"

# Cross-check: the IPs from nslookup should match the pod IPs here
kubectl get pods -l app=redis -o wide
```

DaemonSets offer something else, they are used for 'infrastrucure' workloads that must run on every node, this might be log collection, metrics exporting, network plugin, etc...\
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


> What about Jobs and CronJobs?\
> Kubernetes also provides Jobs and CronJobs for finite, run-to-completion workloads. While Deployments, StatefulSets, and DaemonSets keep Pods running indefinitely, a Job creates one or more Pods that run a task until it succeeds and then stops — think database migrations, batch processing, or one-off scripts.
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



ConfigMaps store non-sensative configuration data as key:value maps, they aren't coupled to any container image, and the data can be mounted as environment variables or as files.\
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

for sensitive data, we don't use configMaps, we use secrets. they are similiar, but are base64-encoded in etcd. there are some basic secret storage types, with `Opaque` being the type for arbitrary user-defined data, and some other types for common use cases:

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

| Volume Type             | Description                                              | Persists After Pod Delete?  |
|-------------------------|----------------------------------------------------------|-----------------------------|
| <k8s>emptyDir</k8s>              | Temp dir shared between containers in a pod              | No                          |
| <k8s>hostPath</k8s>              | Mounts a path from the host node's filesystem            | Yes (on that specific node) |
| <k8s>persistentVolumeClaim</k8s> | Claims a PersistentVolume (decoupled from pod lifecycle) | Yes                         |
| <k8s>configMap</k8s> / <k8s>secret</k8s>  | Mounts config/secret data as files inside the container  | N/A (data lives in etcd)    |

usually, the data inside the pod 'resets' once the pod dies down. there are cases where we want the data to persist across pod sessions. for these cases, we can set up persistent storage, which matches a <k8s>persistentVolume</k8s> and a <k8s>PersistentVolumeClaim</k8s>. thd data lives separatly from the workload which uses it.\


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

## Tools and Commands

<details>
<summary>
Tools which are used in training and kubernetes commands.
</summary>

- [Kind](https://kind.sigs.k8s.io/) - running local kubernetes using docker containers as nodes.
- [k9s](https://k9scli.io/) - interactive terminal for managing and observing clusters.
- [headlamp](https://headlamp.dev/) - user friendly kubernetes UI

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
| <kbd>:</kbd>      | Command mode — type resource names: deploy, svc, pods, ns |
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

</details>



## Takeaways
<details>
<summary>
stuff worth remembering
</summary>

</details>
