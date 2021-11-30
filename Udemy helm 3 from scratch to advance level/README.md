<!--
ignore these words in spell check for this file
// cSpell:ignore nindent kubectl udemy
-->

# Helm 3 From Scratch To Advance Level

[udemy: Helm 3 From Scratch To Advance Level](https://www.udemy.com/course/helm-3-from-scratch-to-advance-leves/)\
[helm documentation](https://helm.sh/)\
[Chart museum](https://chartmuseum.com/)

- [Helm 3 From Scratch To Advance Level](#helm-3-from-scratch-to-advance-level)
  - [Helm Basics](#helm-basics)
  - [Errors Troubleshooting](#errors-troubleshooting)
  - [Chart Deep Dive](#chart-deep-dive)
  - [Working with Multiple Values.yaml](#working-with-multiple-valuesyaml)
  - [Creating and Accessing Template](#creating-and-accessing-template)
  - [Advance Templates and Flow Control](#advance-templates-and-flow-control)
  - [Chart Museum](#chart-museum)
  - [Grafana Example](#grafana-example)
  - [Extra TakeAways](#extra-takeaways)
    - [Helm cli](#helm-cli)
    - [Files and Fields](#files-and-fields)

## Helm Basics

- creating the chart
- chart structure (files and folders)
- configuring yaml files
- deploying the chart
- viewing the chart
- deleting the chart

we create a chart with the `helm create <chart name>` command. when we see a line that uses the "{{ }}" syntax and includes the chart name with a dot notation, that means it fills in the value from the chart.

like in "templates/deployment.yaml" the "application-1.fullname".
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "application-1.fullname" . }}
  labels:
    {{- include "application-1.labels" . | nindent 4 }}
spec:
```

we start by clearing the templates folder, we delete everything except the deployment and service files.

now we will have a simple nginx deployment, with three replicas, exposed with NodePort. we clear the files from unnecessary stuff, leaving us with the most basic form of the files. we clear the values.yaml file, and start working from scratch, we add values that we will use in the deployment and service files.

values.yaml
```yaml
# Default values for application-1.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

deployment:
  replicaCount: 1
  name: my-deployment
  image:
    app: nginx
    version: latest

service:
  name: my-service
  type: NodePort
  port: 80
  targetPort: 80
  nodePort: 32036
```



templates/deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.deployment.name }}
  labels:
    app: nginx
spec:
  replicas: {{ .Values.deployment.replicaCount }}
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: "{{ .Values.deployment.image.app }}:{{ .Values.deployment.image.version }}"
```
templates/service.yaml
```yaml
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.service.name }}
  labels:
    app: nginx
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: {{ .Values.service.targetPort }}
      nodePort: {{ .Values.service.nodePort }}
      protocol: TCP
      name: http
  selector:
    app: nginx
```

now we can deploy the chart.

```sh
kubectl get deployment
helm install chart-1 --dry-run .\application-1 \
helm install chart-1 .\application-1 \
kubectl get deployment,svc
```
now we want the ip address
```sh
kubectl get nodes -o wide
```

and finally, we remove it

```sh
helm ls
helm uninstall chart-1
helm list
```

## Errors Troubleshooting

Trouble shooting, if we missed something or have an indentation problem.

we can create an error by deleting parts of the text inside *{{ }}* in the deployment file, then we can run `helm lint .\application-1` and see that it failed. we can also run `helm template .\application-1` to see the template, and even try to pipe into into kubectl `helm template .\application-1 | kubectl apply -f -` and see the errors from kubectl.

## Chart Deep Dive

now we will create a new chart to learn about the structure

```sh
helm create new-chart
tree new-chart
```
or in powershell
```ps
tree /f
```

we get something like this
```
│   .helmignore
│   Chart.yaml
│   values.yaml
│
├───charts
└───templates
    │   deployment.yaml
    │   hpa.yaml
    │   ingress.yaml
    │   NOTES.txt
    │   service.yaml
    │   serviceaccount.yaml
    │   _helpers.tpl
    │
    └───tests
            test-connection.yaml
```


the "notes.txt" provides data to the user.

the "_helpers.tpl" file, and any other file starting with underscore inside the templates folder, are files which don't create resources.  it just creates variable which are used by resources.\
the format for comments is block comments inside double curly braces. in yaml files we use the pound symbol.

```yaml
{{ /*
comment!
{{- define "some.variable" "value" }}
*/}}

{{/*
Create the name of the service account to use
*/}}
{{- define "new-chart.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "new-chart.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
```




in the chart.yaml, there is a key "version", which we must have, and they appear in the terminal when we list our charts `helm ls`

## Working with Multiple Values.yaml

now we want to work with more value files, so we go back to the "application-1" folder. we first update the version number, and then we duplicate the "values.yaml" file into "new-value.yaml" and we change some values

```yaml
deployment:
  replicaCount: 4
  name: new-deployment
  image:
    app: nginx
    version: latest

service:
  name: new-service
  type: NodePort
  port: 80
  targetPort: 80
  nodePort: 32046
  ```

now we can run the command and see what happens, if we run the default command, we still use the old values. but we can tell to also use another file that will override the commands.
```sh
helm install chart-1 --dry-run .\application-1\
helm install chart-2 --dry-run .\application-1\ -f .\application-1\new-values
```

## Creating and Accessing Template

we now want some way to generate multiple labels. this is why we have templates.

we create a file "templates/_my-template.tpl"

```
{{- define "labels" -}}
app: nginx
version: v1
team: production
{{- end -}}
```

so we got back to the deployment file and update it. wherever we see the labels, we replace them with the values from the template with the `include` command, provide the context (the dot, current context) and pipe into the `nindent` function to drop a line and indent by four space.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.deployment.name }}
  labels: {{- include "labels" . | nindent 4 }}
```

we can now run a dry-run, view the template or check with lint.

## Advance Templates and Flow Control

we now move on to more advanced templates, we start by creating a new file "templates/_my-container.tpl". we add some data to the "values.yaml" file.

```tpl
{{- define "container1" -}}
- name: newcontainer
  image: "{{ .Values.deployment.image.app }}:{{ .Values.deployment.image.version }}"
  ports:
    - name: http
      containerPort: 80
      protocol: TCP
{{- end -}}
```

and we modify the deployment.yaml file to use a conditional statement

```yaml
    spec:
      containers:
        - name: nginx
          image: "{{ .Values.deployment.image.app }}:{{ .Values.deployment.image.version }}"
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
        {{- if eq .Values.container1.enabled true -}}
        {{ include "container1" . | nindent 8 }}
        {{- end -}}
```
and now we can change the key:value pair in the values.file (or set it with a flag or an additional file) to control whether the container is created or not.

we can also create another container defintion in our, file, and use the `{{- else -}}` part to instansiate it instead.


## Chart Museum

Chart museum is a repository for helm charts, like github for code, dockerhub for docker images, etc...

- setup the chart museum repo
- pull the charts
- edit the values.yaml file
- install chart and verify it
- add the repository
- packaging the chart
- pushing the chart
- how to search our charts
- insalling a chart from chart museum
- deleting the chart and the repo


we add repository and check it, we can install directly from the repo, or pull the charts to our local machine.

```sh
helm repo list
helm repo add chartmuseum https://chartmuseum.github.io/charts
helm repo list
helm repo update
helm install --generate-name chartmuseum/chartmuseum
helm uninstall <name>
helm list
helm pull chartmuseum/chartmuseum
ls
tar -xvf .\chartmuseum-3.4.0.tgz
del .\chartmuseum-3.4.0.tgz
ls chartmuseum
```


now we go around and modify "chartmuseum/values.yaml". we change the service.type to "NodePort" and we add the "service.nodePort" key with some value, and we comment out the "service.externalTrafficPolicy: local" pair.

now we try running it, and getting the address. this is actually the address of a repository, our very own chart museum

```sh
helm install mus --dry-run .\chartmuseum\
kubectl get all
kubectl get nodes -o wide
```

and now we want to add the chart into the personal chart museum, so we first add the repo to our list, and then we package our application into an archive.
we also need to change the values.yaml file again and re-enable the api routing.

```yaml
    # disable all routes prefixed with /api
    DISABLE_API: false
```

now lets start packaging our chart and pushing it.
```sh
helm repo add our-repo "ip"
helm repo update
helm repo list
helm package ./application-1
ls

curl --data-binary "@application-1-0.1.1.tgz" <ip adderess>:<port>/api/charts
#helm push
```

to find charts in a repository, we can search all our repositories.
```sh
helm repo update
helm search repo "museum"
```

## Grafana Example

(note: actually, use the [grafana github](https://github.com/grafana/helm-charts) instructions instead)

now we will try to use grafana

we first need some repositories

```sh
helm repo add stable https://charts.helm.sh/stable
helm repo update
helm repo list
hel search repo grafana
helm install grafana stable/grafana
```
we can see some messages coming from the "NOTES.txt" file

```sh
kubectl get all
kubectl edit service grafana
#change to NodePort
```


## Extra TakeAways

### Helm cli

- *helm version*
- *helm create <chart name>*
- *helm install <chart name> <chart folder>*
  - *--dry-run*
  - *--debug*
  - *--set <key1>=<value1>*
  - *--generate-name*
- *helm upgrade <chart name> <chart folder>*
- *helm uninstall <chart name>*
- *helm list, ls*
- *helm lint*
- *helm template*
- **helm repo**
  - *helm repo add <repo name> <url address>*
  - *helm repo index*
  - *helm repo list*
  - *helm repo remove*
  - *helm repo update*
- *helm search repo <optional chart name>*
- *helm pull <repo>/<chart name>*
- *helm push*
- *helm package*


### Files and Fields

[The Chart.yaml file](https://helm.sh/docs/topics/charts/#the-chartyaml-file)

key | required |values | notes
---|---|---|---
apiVersion | true | v2 | ???
version| true | 0.0.0 | chart version `helm ls`
name | true | "some name" |
icon | reccomended | |
