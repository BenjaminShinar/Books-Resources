<!--
ignore these words in spell check for this file
// cSpell:ignore mgmt consec mychart fullname mynamespace mypod relname untitle substr myvalue openapi helmignore subcharts mysubchart dockerignore fnmatch globbing Cappucino
-->

##

[eksctl](https://eksctl.io/), [eksctl github](https://github.com/weaveworks/eksctl), [helm](https://helm.sh/)

## CLI commands

### kubectl

```sh
kubectl config view
```

### aws, aws eks

```sh
aws configure list
aws configure list-profiles

aws eks --region <region-code> update-kubeconfig --name <cluster_name>
```

### eksctl

```sh
eksctl get cluster
eksctl create cluster --dry-run > cluster.yaml
eksctl delete cluster --name=<name> --wait
```

### helm

```sh
helm version
helm list -A -a
helm status <name> -n <namespace>

helm lint
helm install <release name> ./mychart
helm get manifest <release name>
helm uninstall <release name>

helm install --debug --dry-run <release name> ./mychart
helm install --debug --disable-openapi-validation --dry-run <release name> ./mychart
helm install --debug --dry-run <release name> ./mychart --set <key>=<value>

```

###

actual running

```sh
aws configure
eksctl create cluster -f cluster_definition.yaml
# stack will be created, will take time. about 20 minutes


# verify that things work
eksctl get cluster
aws eks list-clusters
kubectl get all

helm install chartname --repo <https://raw.githubusercontent.com/CheckPointSW/charts/ea/repository/> --set credentials.user=<> --namespace checkpoint --create-namespace

kubectl get all -A

```

## Yaml

cluster,yaml

```yaml
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

metadata:
  name: basic-cluster
  region: eu-north-1

nodeGroups:
  - name: ng-1
    instanceType: m5.large
    desiredCapacity: 10
    volumeSize: 80
    ssh:
      allow: true # will use ~/.ssh/id_rsa.pub as the default ssh key
  - name: ng-2
    instanceType: m5.xlarge
    desiredCapacity: 2
    volumeSize: 100
    ssh:
      publicKeyPath: ~/.ssh/ec2_id_rsa.pub
```

yaml multi line begins with "|" - pipe operator. this can cause indentation errors, so we can put a fake first line comment, and we can control how the multiline and white space is preserved. we can also use ">" to fold the multiline.

```yaml
coffee: |
  Latte
  Cappuccino
  Espresso

#protect ourselves from weirdness in helm templates
moreCoffee: |
  # Commented first line
         Latte
  Cappuccino
  Espresso

#no trailing new ling
extraCoffee: |-
  Latte
  Cappuccino
  Espresso

# preserve trailing white spaces
moreCoffee: |+
  Latte
  Cappuccino
  Espresso


another: value

oneLongCoffee: >
  Latte
  Cappuccino
  Espresso
```

we can force values to be strings or numbers

```yaml
age: !!str 21
port: !!int "80"
```

we can stick some different yaml documents into a single file, we separate them with a leading "---" and a finishing "...", it's sometimes possible to omit one of them.

**anchoring** means giving a value a name, like a variable. we declare it with "&" and use it with "\*".

```yaml
coffee: "yes, please"
favorite: &favoriteCoffee "Cappucino"
coffees:
  - Latte
  - *favoriteCoffee
  - Espresso
```

## Helm

### Top Level Objects

- Release
  - Release.Name - the name of the release
  - Release.Namespace - the namespace the release is in (unless overridden by the manifest)
  - Release.IsUpgrade - true when upgrade or rollback
  - Release.IsInstall - true when install
  - Release.Revision - current version, starts with 1
  - Release.Service - always helm
- Values (everything in the values.yaml file and other files)
- Chart - the charts.yaml file
- Files - non special files, can't access template
  - Files.Get - function for getting file by name (`.Files.get config.ini`)
  - Files.GetBytes - function for getting file contents as byte stream and not string
  - Files.Glob - (function) pattern matching for file names
  - Files.Lines - (function) read file line by line
  - Files.AsSecrets - (function) read the file as base64 encoding
  - Files.AsConfig - (function) return file bodies as yaml
- Capabilities - the kubernetes cluster
  - Capabilities.APIVersions
    - Capabilities.APIVersions.Has - (function) determines if a version or a resource is available on this cluster.
  - Capabilities.KubeVersion - kubernetes version
    - Capabilities.KubeVersion.Version - version
    - Capabilities.KubeVersion.Major - major version
    - Capabilities.KubeVersion.Minor - minor version
  - Capabilities.HelmVersion - version, same as `helm version`
    - Capabilities.HelmVersion.Version - semver format
    - Capabilities.HelmVersion.GitCommit - git sha1
    - Capabilities.HelmVersion.GitTreeState -
    - Capabilities.HelmVersion.GoVersion - go compiler version
- Template - information about the template
  - Template.Name
  - Template.BasePath

### Structure

the helm **.Values** object is composed from the following. this list is ordered, so the parent file can override the child file, and the individually passed values have the top priority.

1. the values.yaml file of the chart
2. if this is a sub-chart, the values.yaml file of the parent chart
3. any file passed into `helm install` or `helm upgrade` with the _-f_ flag
4. parameters passed with the _--set_ flag (e.g. `helm install --set foo=bar ./mychart`)

if we wish to delete a key, we can set it to null. this is important for objects which allows only one value (like livenessProbe Handler)

### Functions

[many functions!](https://helm.sh/docs/chart_template_guide/function_list/)

we can call function either with the function name before the argument or with the pipe operator

- `quote` - add quotes `{{ quote .Values.someValue}}` or `{{ .Values.someValue | quote }}`
- `upper` - uppercase
- `lower` - lowercase
- `title` - title case
- `untitle` - remove title case
- `substr` - get substring
- `repeat N` - repeat arguments N times. `{{ .Values.someValue | repeat 5}}`
- `default` - default value if none given.
  - if the value is known, it should be part of the values.yaml file, but this works for computed values. e.g. `default (printf "%s-tea" (include "fullname" .))`
- `lookup` - find resources. format is apiVersion, kind, namespace(optional), name(optional).

  - works like getting resources with _kubectl_.

  | kubectl equivalent                     | look up function                           |
  | -------------------------------------- | ------------------------------------------ |
  | `kubectl get pod mypod -n mynamespace` | `lookup "v1" "Pod" "mynamespace" "mypod"`  |
  | `kubectl get pods -n mynamespace`      | `lookup "v1" "Pod" "mynamespace" ""`       |
  | `kubectl get pods --all-namespaces`    | `lookup "v1" "Pod" "" ""`                  |
  | `kubectl get namespace mynamespace`    | `lookup "v1" "Namespace" "" "mynamespace"` |
  | `kubectl get namespaces`               | `lookup "v1" "Namespace" "" ""`            |

  - the value returned is a dictionary, which we can drill into. e.g: `(lookup "v1" "Namespace" "" "mynamespace").metadata.annotations`

- range looping
  - we can loop over a range of items, which we get from a lookup.
  ```go
  {{ range $index, $service := (lookup "v1" "Service" "mynamespace" "").items }}
   {{/* do something with each service */}}
   {{ end }}
  ```
- logical operators are functions
  - `eq,ne,lt,le,gt,ge,and,or,not` and also `default, coalesce, empty, fail`
- `fail` - unconditionally fail with the specified error and text.
- `print` - e.g `print "Matt has " .Dogs " dogs"`
- `println`
- `printf` - e.g `printf "%s has %d dogs." .Name .NumberDogs`
- `trim` - e.g `trim " hello "`
  - `trimAll` - remove a character from string,
    - e.g. `trimAll "$" "$5.00"`
  - `trimPrefix`
  - `trimSuffix`
- `nospace` - remove all whitespace

- grouping `((, and ))`

empty definitions:

| type       | empty value     |
| ---------- | --------------- |
| numeric    | 0               |
| String     | ""              |
| Lists      | []              |
| Dicts      | {}              |
| Boolean    | false           |
| _anything_ | nil (like null) |
| _struct_   | no empty value  |

### Flow control

- `if/else` - conditional
- `with` - specify scope
- `range` - used in _for-each_ loops
- `define` - new named template inside a template
- `template` - import a named template
- `block` -
  > declares a special kind of fillable template area

if else: we can put either values inside the the condition or an entire pipeline. values are evaluated to false if they are boolean false, numeric zero, empty string, empty collection, or _nil_

```yaml
{{ if PIPELINE }}
  # Do something
{{ else if OTHER PIPELINE }}
  # Do something else
{{ else }}
  # Default case
{{ end }}
```

by default, lines with `{{ */.../* }}` remain in the output as whitespace lines, but if we add a dash, then they will be removed `{{- */.../* }}`.
**be very carful about spacing**

scoping: we can lif an object into scope, making it the current scope ( the **"."** ) instead of the global scope. this lasts until the `{{ end }}` line.

```yaml
{{ with PIPELINE }}
  # restricted scope
{{ end }}
```

we lose the simple access to the parent scope (usually the root scope), but we can always access the global scope with `$`. like this

```yaml
  {{- with .Values.favorite }}
  drink: {{ .drink | default "tea" | quote }}
  food: {{ .food | upper | quote }}
  release: {{ $.Release.Name }}
  {{- end }}
```

range: creating go slice out of a collection
a combination of pipe and dash means a multi-line string `|-`.

```yaml
sizes: |-
  {{- range tuple "small" "medium" "large" }}
  - {{ . }}
  {{- end }}
```

we can create a variable with the $ symbol.we declare a name and then assign a value to it. then we can access it again. the scoping rules are the usual (lexical scoping?)

```yaml
  {{- $relname := .Release.Name -}}

release: {{ $relname }}
```

this also works great with ranges

```yaml
toppings: |-
  {{- range $index, $topping := .Values.pizzaToppings }}
    {{ $index }}: {{ $topping }}
  {{- end }}
```

the global root **$** is always available, it's the root global scope.

### Named templates

names templates (partial templates, sub-templates) are templates which were defined inside a file and given a name. templates names are global, and the last loaded one is used.

the files inside the "./template" folder are assumed to be definitions of kubernetes resources. the exceptions are the "NOTES.txt" file, and files starting with an underscore. these files, like "\_helper.tpl" contain partial template and helpers.

we declare a partial with `define` and close the declaration with `end`, we can chomp off the empty lines with a dash, if needed.

```yaml
{{ define "mychart.labels"}}

{{ end }}
```

here is an example of using a partial template inside the same file it was defined in. we first define it and then call it with `template`, to bring it over.

```yaml
{{- define "mychart.labels" }}
  labels:
    generator: helm
    date: {{ now | htmlDate }}
{{- end }}

apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
  {{- template "mychart.labels" }}
data:
  myvalue: "Hello World"
  {{- range $key, $val := .Values.favorite }}
  {{ $key }}: {{ $val | quote }}
  {{- end }}
```

by convention, we usually stick the partial templates inside the "\_helpers.tpl" file. and we usually want to give them unique names.

when we simply try to install the chart, we get an error. if we add the _--disable-openapi-validation_ flag, it works, but without the parts that use _.Chart_.

```sh
helm install --dry-run moldy-jaguar ./mychart
helm install --dry-run --disable-openapi-validation moldy-jaguar ./mychart
```

the problem is the scope. the partial template is defined outside of the current scope, so we need to pass the scope to it, which is done by providing the current scope (the dot **"."**) as an argument to the template call. we can pass any scope

```Smarty
metadata:
  name: {{ .Release.Name }}-configmap
  {{- template "mychart.labels" . }}
```

using `template` is fine, but it has a problem. it doesn't indent itself well into it's position, it's aligned left, so it only uses the indention of the partial template, and not the indentation in the calling file.

we can go around this by using `include`, and then piping the result into the `indent` function.

```yaml
{ { include "mychart.app" . | indent 2 } }
```

we should use `include` and not `template`, for this reason.

### Files

we can also access files directly, (not as templates) with the _.Files_ object. however, we can't use files which are inside the "/templates" folder or files which were excluded with ".helmignore". also, there is a size limit for charts to be less than 1mb. also, some issue with permissions.

lets create some files and use them

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  {{- $files := .Files }}
  {{- range tuple "config1.toml" "config2.toml" "config3.toml" }}
  {{ . }}: |-
        {{ $files.Get . }}
  {{- end }}
```

when we are working with files, we get go's **path** package functions:

- `Base`
- `Dir`
- `Ext`
- `IsAbs`
- `Clean`

we can use **Glob** to find files, like all files with the yaml extensions.

```yaml
{{ $currentScope := .}}
{{ range $path, $_ :=  .Files.Glob  "**.yaml" }}
    {{- with $currentScope}}
        {{ .Files.Get $path }}
    {{- end }}
{{ end }}
```

was can also use `.AsConfig` and `.AsSecret` if we need to read files in a certain way

- `text | b64enc`
- `.Lines` to iterate over each line in the file.

we can't pass files to helm during `helm install`,

### The "Notes.txt" file

the "templates/NOTES.txt" file defines a message, this is printed after a install, and gives extra data to the user, it has the same access as other charts and templates.

### Sub Charts

we can have more than one chart, as charts can have dependencies (sub-charts), each with their own values and templates

> 1. A subchart is considered "stand-alone", which means a subchart can never explicitly depend on its parent chart.
> 2. For that reason, a subchart cannot access the values of its parent.
> 3. A parent chart can override values for subcharts.
> 4. Helm has a concept of global values that can be accessed by all charts.

we start by creating a subchart, from the root folder we move into the 'charts' folder, and then we run `helm create mysubchart`. this is just like creating the regular chart.

we populate the values.yaml and create a configmap.yaml resource, and now we can test it

```sh
helm install --generate-name --dry-run --debug mychart/charts/mysubchart
```

because the subchart is defined inside the "charts" folder of another chart, then it's a subchart, so we can define the same value in the parent chart, it will then override the value in sub chart

```yaml
mysubchart:
  dessert: ice cream
```

if we install the sub chart, it will use it's own values. but if we install the parent chart, the outer value will be used instead. this is because each chart is a **stand-alone**, and it can be generated on it's own or as a subchart.

```sh
helm install --generate-name --dry-run --debug mychart/charts/mysubchart
helm install --generate-name --dry-run --debug mychart
```

there is another section in the values.yaml file: **global**, globals are automatically shared with subcharts, without explicitly requiring the value to exist in both, but we do get an error if we try to create the sub chart, because there is no global key defined in the nested "values.yaml", so we need to be smart

[workaround solution](https://stackoverflow.com/questions/59795596/helm-optional-nested-variables)

```yaml
salad: { { default "yuck" ((.Values.global).salad) } }
```

we can share templates between charts and subcharts. any defined block in any chart is available to any other chart. (**not clear!**)

there is also the `block` statement, but it has some un predictable behavior, so we should avoid it.

### The ".helmignore" file

like ".gitignore", ".dockerignore", this file controls what won't be taken inside the package.

> Some notable differences from .gitignore:
>
> - The '\*\*' syntax is not supported.
> - The globbing library is Go's 'filepath.Match', not fnmatch(3)
> - Trailing spaces are always ignored (there is no supported escape sequence)
> - There is no support for '!' as a special leading sequence

### Extras

#### Operators

| operator | example                                      | notes                                                       |
| -------- | -------------------------------------------- | ----------------------------------------------------------- |
| and      | `and .Arg1 .Arg2 `                           | arg1 && arg2                                                |
| or       | `or .Arg1 .Arg2`                             | arg1 \|\| arg2                                              |
| not      | `not .Arg`                                   | arg1 != arg2                                                |
| eq       | `eq .Arg1 .Arg2`                             | arg1 == arg2                                                |
| ne       | `ne .Arg1 .Arg2`                             | arg1 &ne; arg2                                              |
| lt       | `lt .Arg1 .Arg2`                             | arg1 &lt; arg2                                              |
| le       | `le .Arg1 .Arg2`                             | arg1 &le; arg2                                              |
| gt       | `gt .Arg1 .Arg2`                             | arg1 &gt; arg2                                              |
| ge       | `ge .Arg1 .Arg2`                             | arg1 &ge; arg2                                              |
| empty    | `empty .Foo`                                 | determine if .foo is empty, see table for what empty means. |
| default  | `default "foo" .Bar`                         | if bar is empty, use "foo"                                  |
| coalesce | `coalesce .name .parent.name "Matt"`         | first non-empty value, like a chained default command       |
| ternary  | `ternary "foo" "bar" {{ .Values.someValue}}` | like c ` value ? "foo" : "bar"`                             |
|  has | `if has "security.openshift.io/v1" .Capabilities.APIVersions` | list contains, string contains

#### The "values.yaml" file

a values.yaml file cannot have separate sections with the "---" "..." format. only the first one will be used.

while we can use anchoring (declaring variables with "&") in yaml, helm usually ignores them

- key:value pairs, can be nested as needed.
- subchart as key, and then key:value pairs to override the values in the subchart.
- global section, with key value pairs.

#### RBAC

role based access control

- ServiceAccount (namespaced)
- Role (namespaced)
- ClusterRole
- RoleBinding (namespaced)
- ClusterRoleBinding
