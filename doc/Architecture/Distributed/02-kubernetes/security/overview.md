# Overview

## Cloud
### Infrastructure security

|Area of Concernt for Kubernetes Infrastructure|Recommendation|
|-|-|
|Network access to API Server (Control plane)|All access to the Kubernetes control plane is not allowed publicly on the internet and is controlled by network access control lists restricted to the set of IP addresses needed to administer the cluster|
|Network access to Nodes(nodes)|Nodes should be configured to *only* accept connections(via network access control lists) from the control plane on the specified ports, and accept connections for services in Kubernetes of type NodePort and LoadBalancer. If possible, these nodes should be exposed on the public internet entirely|
|Kubernetes access to Cloud Provider API|Each cloud provider needs to grant a different set of permissions to the Kubernetes control plane and nodes. It is best to provide the cluster with cloud provider access that follows the principle of least priviledge for the resources it needs to administer. The Kops documentation about IAM policies and roles|
|Access to etcd|Access to etcd(the datastore of Kubernetes| should be limited to the control plane only. Depending on your configuration, you should attempt to use etcd over TLS.
|etcd Encryption|Whenever possible it's a good practice to encrypt all storage at rest, and since etcd holds the state of the entire cluster(including Secrets) its disk should especially be encrypted at rest|

## Cluster

There are two areas of concern for securing Kubernetes:

- Securing the cluster components that are configurable
- Securing the applications which run in the cluster

### Components of Cluster

Depending on the attack surface of your application, you may want to focus on specific aspects of security. For example: If you are running a service(Service A) that is critical in a chain of other resources and a seperate workload(Service B) which is vulnerable to a resource exhaustion attack, then the risk of compromising Service A is high if you do not limit the resources of Service B. The following table lists areas of security concerns and recommendations for securing workloads running in Kubernetes:

- RBAC Authorization(Access to the Kubernetes API)
- Authentication
- Application secrets management(and encrypting them in etcd at rest)
- Ensuring that pods meet defined Pod Security Standards
- Quality of Service(and Cluster resource management)
- Network Policies
- TLS for Kubernetes Ingress

## Container

Container security is outside the scope of this guide.

- Container Vulnerablitiy Scanning and OS Dependency Security
    - As part of an image build step, you should scan your containers for known vulnerablities.
- Image Signing and Enforcement
    - Sign container images to maintain a system of trust for the content of your containers.
- Disallow priviledged users
    - When constructing containers, consult your documentation for how to create users inside of the containers that have the least level of operating system privilege necessary in order to carry out the goal of the container
- Use container runtime with stronger isloation
    - Select container runtime classes that provide stronger isloation

## Code

Application code is one of the primary attack surfaces over which you have the most control. While securing application code is outside of the Kubernetes security topic, here are recommendations to protect application code:

- Access over TLS only
    - If your code needs to communicate by TCP, perform a TLS handshake with client ahead of time. With the exception of a few cases, encrypt everything in transit. Going one step further, it's a good idea to encrypt network traffic between services. This can be done through a process known as mutual TLS authentication or mTLS which performs a two sided verification of communication between two certificate holding services.
- Limiting port ranges of communication
    - This recommendation may be a bit self-explanatory,but wherever possible you should only expose the ports of your service that are absolutely essential for communication or metric gathering.
- 3rd Party Dependency Sercurity
    - It is a good practice to regularly scan your application's third party libraries for known security vulnerabilities. Each programming language has a tool for performing this check automatically
- Static Code Analysis
    - Most languages provide a way for a snippet of code to be analyzed for any potential unsafe coding practices. Whenever possible you should perform checks using automated tooling that can scan codebases for common security errors.
- Dynamic probing attacks
    - There are a few automated tools that you can run against your service to try some of the well known service attacks. These include SQL injection, CSRF, and XSS. One of the most popular dynamic analysis tools is the OWASP Xed Attack proxy tool