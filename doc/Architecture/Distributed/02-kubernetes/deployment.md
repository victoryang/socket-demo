# Deployments

## Creating a Deployment

```yaml
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
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

### Create
```bash
kubectl apply -f nginx-deployment.yaml
```

### Check
```bash
kubectl get deployment
```

## Update a Deployment

```bash
# update
kubectl edit deployment.apps/nginx-deployment

# see rollout status
kubectl rollout status deployment.apps/nginx-deployment
```

Deployment ensures that only a certain number of Pods are down while they are being updated. By default, it ensures that at least 75% of desired number of Pods are up(25% max unavailable).

Deployment also ensures that only a certain number of Pods are created above the disired number of Pods. By default, it ensures that at most 125% of the desired number of Pods are up(25% max surge).

For example, if you look at the above Deployment closely, you will see that it first created a new Pod, then delete some old Pods, and created new ones. It does not kill old Pods until a sufficient number of new Pods have come up, and does not create new Pods until a sufficient number of old Pods have been killed. It makes sure that at least 2 Pods are available and that at max 4 Pods in total are available.

### Rollover (aka multiple updates in-flight)

Each time a new Deployment is observed by the Deployment controller, a ReplicaSet is created to bring up the desired Pods. If the Deployment is updated, the existing ReplicaSet that controls Pods whose labels match `.spec.selector` but whose template does not match `.spec.template` are scaled down. Eventually, the new ReplicaSet is scaled to `.spec.replicas` and all old ReplicaSet is scaled to 0.

If you update a Deployment while an existing rollout is in progress, the Deployment creates a new ReplicaSet as per the update and start scaling that up, and rolls over the ReplicaSet that it was scaling up previously -- it will add it to its list of old ReplicaSets and start scaling it down.

For example, suppose you create a Deployment to create 5 replicas of `nginx:1.14.2`, but then update the Deployment to create 5 replicas of `nginx:1.16.1`, when only 3 replicas of `nginx:1.14.2` had been created. In that case, the Deployment immediately starts killing the 3 `nginx:1.14.2` Pods that it had created, and starts creating `nginx:1.16.1` Pods. It does not wait for the 5 replicas of `nginx:1.14.2` to be created before changing course.

### Label selector updates

It is generally discouraged to make label selector updates and it is suggested to plan your selectors up front.

## Rolling Back a Deployment

```bash
# checkout status
kubectl rollout status deployment.apps/nginx-deployment

# checkout history
kubectl rollout history deployment.apps/nginx-deployment --revision=2

# undo
kubectl rollout undo deployment.apps/nginx-deployment

# or
kubectl rollout undo deployment.apps/nginx-deployment --to-revision=2
```

## Scaling a Deployment

```bash
# scale to 10
kubectl scale deployment.apps/nginx-deployment --replicas=10

# auto scale
kubectl autoscale deployment.v1.apps/nginx-deployment --min=10 --max=15 --cpu-percent=80
```

## Pausing and Resuming a Deployment

## Deployment Status

## Clean up Policy

## 