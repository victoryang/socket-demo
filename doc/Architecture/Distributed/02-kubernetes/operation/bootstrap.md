# Bootstrap a cluster

## Kubeadm

## Kops

This shows you how to easily install a Kubernetes cluster on AWS. It uses a tool called `kops`.

kops is an automated provisioning system:

- Fully automated installation
- Uses DNS to identify clusters
- Self-healing: everything runs in Auto-Scaling Groups
- Multiple OS support
- High-Availability support
- Can directly provision, or generate terraform manifests

## Kubespray

This quickstart helps to install a Kubernetes cluster hosted on GCE, Azure, OpenStack, AWS, vSphere, Packet(bare metal), Oracle Cloud Infrastructure or Baremetal with Kubespray.

Kubespray is a composition of Ansible playbooks, inventrory, provisioning tools, and domain knowledge for generic OS/Kubernetes clusters configuration management tasks.