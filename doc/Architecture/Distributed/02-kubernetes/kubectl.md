# Kubectl

[kubectl](http://dockone.io/article/9134)

[kubectl management](http://dockone.io/article/9087)

[docs](https://kubectl.docs.kubernetes.io)

## Introduction

Kubectl is the Kubernetes cli version of a swiss army knife, and can do many things.

### Command Families

|Type|Used For|Description|
|-|-|-|
|Declarative Resource Management|Deployment and Operations|Decalaritively manage Kubenetes Workloads using Resource Config|
|Imperarive Resource Management|Development Only|Run commands to manage Kubernetes Workloads using Command Line arguements and flags|
|Printing Workload State|Debugging|Print information about Workloads|
|Interacting with Containers|Debugging|Exec, Attach,Cp,Logs|
|Cluster Management|Cluster Ops|Drain and Cordon Nodes|

### Declarative Application Management

The preferred approach for managing Resource is through declarative files called Resource Config used the Kubectl *Apply* command. This command reads a local (or remote) file structure and modifies cluster state to reflect the declared state.

Apply is the preffered mechanism for managing Resources in a Kubernetes cluster.

### Print state about Workloads

Users will need to view Workload state

- Printing summarize state and information about Resources
- Printing complete state and information about Resources
- Printing specific fields from Resourves
- Query Resources matching labels

### Debugging Workloads

Kubectl supports debugging by providing commands for: 

- Printing Container logs
- Printing cluster events
- Exec or attach to a Container
- Copying files from Containers in the cluster to a user's filesystem

### Cluster Management

On occasion, users may need to perform operations to the Nodes of cluster. Kubectl supports commands to drain Workloads from a Node so that it can be decommission or debugged.

### Porcelain

Users may find using Resource Config overly verbose for Development and prefer to work with the cluster imperatively with a shell-like workflow. Kubectl offers porcelain commands for generating and modifying Resources.

- Generating + creating Resources such as Deployments, StatefulSets, Services, ConfigMaps, etc
- Setting fields on Resources
- Edting (live) Resources in a text editor

## Resources + Controllers Overview

- A Kubernetes API has 2 parts - a Resource Type and a Controller
- Resources are objects declared as json or yaml and written to a cluster
- Controllers asynchronously actuate Resources after they are stored

### Kubernetes Resources and Controllers Overview

#### Resources

Instances of Kubernetes objects (e.g. Deployment, Services, Namespaces, etc) are called Resources

Resources which run containers are referred to as Workloads

Examples of Workloads:

- Deployments
- StatefulSets
- Jobs
- CronJobs
- DaemonSets

**Users work with Resource APIs by declaring them in files which are then applied to Kubernetes cluster. These declarative files are called Resource Config**

Resource Config is *Applied* (declarative Create/Update/Delete) to a Kubernetes cluster using tools such as kubectl, and then actuated by a Controller.

Resources are uniquely identified

- apiVersion (API Type Group and Version)
- kind (API Type Name)
- metadata.namespace (Instance namespace)
- metadata.name (Instance name)

If namespace is ommited from the Resource Config, the *default* namespace is used. Users should almost always explicitly specify the namespace for their Application using a `kustomization.yaml`.

##### Resource Structure

Resource have following components

**TypeMeta**: Resource Type **apiVersion** and **kind**

**ObjectMeta**: Resource **name** and **namespace** + other metadata (label, annotations, etc)

**Spec**: the desired state of the Resource - intended state the user provides to the cluster.

**Status**: the observed state of the object - recorded state the cluster provides to users.

Resource Config written by the user omits the Status fields.

**Example Deployment Resource Config**

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
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
        image: nginx:1.15.4

```

#### Controllers

Controllers actuate Kubernetes APIs. They observe the state of the system and look for changes either to desired state of Resources (create, update, delete) or the system (Pod or Node dies).

Controllers then make changes to the cluster to fulfill the intent specified by the user (e.g. in Resource Config) or automation (e.g. changes from Autoscalers)

Example: After a user creates a Deployment, the Deployment Controller will see that the Deployment exists and verify the corresponding ReplicaSet it expects to find exists. The Controller will see that the ReplicaSet does not exist and will create one.

`
Because Controllers run asynchronously, issues such as a bad Container Image or unschedulable Pods will not be present in the CRUD response. Tooling must facilitate processes for watching the state of the system until changes are completely actuated by Controllers. Once the changes have been fully actuated such that the desired state matches the observed state, the Resource is considered *Settled*.
`

##### Controller Structure

**Reconcile**