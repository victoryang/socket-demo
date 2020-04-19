# Kubernetes Object

[Kuberntes Object](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)

[Kubernetes API Reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/)

[kubectl book](https://kubectl.docs.kubernetes.io/)

## Understanding Kubernetes Objects

Kuberntes objects are persistent entities in the Kubernetes system. Kubernetes uses these entities to represent the state of your cluster. Specifically, they can describe:

- What containerized application are running(and on which nodes)
- The resources available to those applications
- The policies around how those application behave, such as restart policies, upgrades, and fault-tolerance

A Kubernetes object is a "record of intent" - once you create the object, the Kubernetes system will constantly work to ensure that object exists. By creating an object, you're effectively telling the Kubernetes system what you want your cluster's workload to look like; this is your cluster's *desired* state.

To work with Kubernetes objects - whether to create, modify,or delete them - you'll need to use the Kubernetes API. When you use the `kubectl` command-line interface, for example, the CLI makes the necessary Kubernetes API calls for you. You can also use the Kubernetes API directly in your own programs using one of Client Libraries.

### Object Spec and Status

For Object that hava a `spec`,you have to set this when you create the object, providing a description of the characteristic you want the resource to have: its *desired* state.

The `state` describes the *current state* of the object, supplied and updated by the Kubernetes system and its components. The Kubernetes control plane continually and actively manages every object's actual state to match the desired state you supplied.

### Describing a Kubernetes Object

When you use the Kubernetes API to create the object, that API request must include that information as JSON in the request body. Most often, you provide the information to `kubectl` in a .yaml file. `kubectl` converts the information to JSON when making the API request.

### Required Fields

In the .yaml file for the Kubernetes object you want to create, you'll need to set values for the following fields:

- apiVersion - Which version of the Kubernetes API you're using to create this object
- kind - What kind of object you want to create
- metadata - Data that helps uniquely identify the object, including a `name` string, `UID` and optional `namespace`
- spec - What state you desir for the object

The precise format of object `spec` is different for every Kubernetes object, and contains nested fields specific to that object. The Kubernetes API reference can help you find the spec format for all of the objects you can create using Kubernetes. For example, the `spec` format for a Pod can be found in PodSpec v1 core, and the `spec` format for a Deployment can be found in DeploymentSpec v1 apps.

## Kubernetes Object Management

The `kubectl` command-line tool supports several different ways to create and manage Kubernetes objects.

- imperative commands
- imperative object configuration
- Declarative object configuration

A Kubernetes object should be managed using only one technique. Mixing and matching techniques for the same object results in undefined behavior

### Imperative commands

When using imperative commands, a user operates directly on live objects in a cluster. The user provides operations to the `kubectl` command as arguments or flags.

This is the the simplest way to get started or to run a one-off task in a cluster. Because this technique operates directly on live objects, it provides no history of previous configurations.

### Imperative object configuration

In imperative object configuration, the kubectl command specifies the operation(create, replace, etc), optional flags and at least one file name. The file specified must contain a full definition of the object in YAML or JSON format.

### Declarative object configuration

When using declarative object configuration, a user operates on object configuration files stored locally, however the user does not define the operations to be taken on the files.Create, update and delete opertaions are automatically detected per-object by kubectl. This enables working on directories, where different operations might be needed for different objects.

## Object Names and IDs

Each object in your cluster has a Name that unique for that type of resource. Every Kuberntes object also has a UID that is unique across your whole cluster.

For example, you can only have one Pod name myapp-1234 within the same namespace, but you can have one Pod and one Deployment that are each named myapp-1234.

For non-unique user-provided attributes, Kubernetes provides labels and annotations.

### Names

A client-provided string that refers to an object in a resource URL, such as /api/v1/pod/some-name.

Only one object of a given kind can have a given name at a time.However, if you delete the object, you can make a new object with the same name.

Below are three types of commonly used name constraints for resources.

#### DNS Subdomain Names

#### DNS Label Names

#### Path Segment Names

### UIDs

A Kubernetes system-generated string to uniquely identify objects.

Every object created over the whole lifetime of a Kubernetes cluster has a distinct UID. It is intended to distinguish between historical occurennces of similar entities.

Kubernetes UIDs are universally unique identifiers (also known as UUIDs)

### Namespaces

Kubernetes supports multiple virtual clusters backed by the same physical cluster. This virtual clusters are called namespaces.

#### When to use Multiple Namespaces

Namespace are intended for use in environments with many users spread across multiple teams, or projects. For cluster with a few to tens of users, you should not need to create or think about namespaces at all. Starting using namespaces when you need the features they provides.

Namespaces provide a scope of names. Names of resources need to be unqiue within a namespace, but not across namespaces. Namespaces can not be nested inside another and each Kubernetes resources can only be in one namespace.

Namespaces are a way to divide cluster resources between multiple users.

It is not necessary to use multiple namespaces just to seperate slightly different resources, such as different versions of the same software: use labels to distinguish resources within the same namespace.

#### Working with Namespaces

`kubectl get namespace`

Kubernetes starts with three initial namespaces;

- `default` The default namespace for objects with no other namespace
- `kube-system` The namespace for objects created by the Kubernets system
- `kube-public` This namespace is created automatically and is readable by all users. This namespace is mostly reserved for cluster usages, in case that some resources should be visible and readable publicly throughout the whole cluster. The public aspect of this namespace is only a convention, not a requirment.

#### Namespace and DNS

#### Not All Objects are in a Namespace

Most Kubernetes resources are in some namespaces.However namespace resources are not themselves in a namespace. And low-level resources, such as nodes and persistentVolumes, are not in any namespace.

## Lables and Selectors

## Annotations

You can use either labels or annotations to attach metadata to Kubernetes objects

## Field Selectors

## Recommanded Labels

