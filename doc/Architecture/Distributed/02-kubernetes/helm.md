# Helm

## Conception

**Helm**
Package Manager for Kubernetes
To package YAML Files and distribute them in public and private repositories

**Helm Charts**

- Bundle of YAML files
- Create your own Helm Charts with Helm
- Push them to Helm Repository
- Download and use existing ones

## Get Started

[Install](https://helm.sh/docs/intro/install/)

### Repo source
```bash
# add stable repo
helm repo remove stable
helm repo add stable https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts

# update
helm repo update

# list repo
helm search repo stable
```