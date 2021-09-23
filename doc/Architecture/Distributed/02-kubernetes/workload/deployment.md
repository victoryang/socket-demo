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

> A Deployment's rollout is triggered if and only if the Deployment's Pod template(that is, .spec.template) is changed, for example if the labels or container images of the template are updated. Other updates, such as scaling the Deployment, do not trigger a rollout.

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

It is generally discouraged to make label selector updates and it is suggested to plan your selectors up front. In any case, if you need to perform a label selector update, exercise great caution and make sure you have grasped all of the implications.

> In API version apps/v1, a Deployment's label selector is immutable after it gets created.

- Selector additions require the Pod template labels in the Deployment spec to be updated with new label too, otherwise a validation error is returned. This change is a non-overlapping one, meaning that the new selectors does not select ReplicaSets and Pods created with old selector, resulting in orphaning all old replicaSet and creating a new replicaSet.

- Selector updates changes the existing value in a selector key -- result in the same behaviour as addtions.

- Selector removals remove an existing key from the Deployment selector -- do not require any changes in the Pod template labels. Existing ReplicaSets are not orphaned, and a new ReplicaSet is not created, but note that the removal label still exists in any existing Pods and ReplicaSets.

## Rolling Back a Deployment

Sometimes, you may want to rollback a Deployment; for example, when the Deployment is not stable, such as crash looping. By default, all of the Deployment's rollout history is kept in the system so that you can rollback anytime you want(you can change that by modifying revision history limit).

> **Noted:** A Deployment's revision is created when a Deployment's rollout is triggered. This means that the new revision is created if and only if the Deployment's Pod template(.spec.template) is changed,for example if you update the labels or container images of the template. 

> **Noted:** The Deployment controller stops the bad rollout automatically, and stops scaling up the new ReplicaSet. This depends on the rollingUpdate parameters(maxUnavailable specifically) that you have specified. Kubernetes by default sets the value to 25%.

```bash
# checkout status
kubectl rollout status deployment.apps/nginx-deployment

# checkout history
kubectl rollout history deployment.apps/nginx-deployment --revision=2

# undo, rolling back to a previous revision
kubectl rollout undo deployment.apps/nginx-deployment

# or rolling back to a specified revision
kubectl rollout undo deployment.apps/nginx-deployment --to-revision=2
```

## Scaling a Deployment

You can scale a Deployment by using the following command:

```bash
# scale to 10
kubectl scale deployment.apps/nginx-deployment --replicas=10
```

Assuming [horizontal Pod autoscaling](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/) is enabled in your cluster, you can setup an autoscaler for your Deployment and choose the minimum and maximum number of Pods you want to run based on the utilization of your exising Pods.

```bash
# auto scale
kubectl autoscale deployment.v1.apps/nginx-deployment --min=10 --max=15 --cpu-percent=80
```

### Proportional scaling

RollingUpdate Deployment support running multiple version of an application at the same time. When you or an autoscaler scales a RollingUpdate Deployment that is in the middle of a rollout(either in progress or paused), the Deployment controller balances the additional replicas in the existing active ReplicaSet(ReplicaSets with Pods) in order to mitigate risk. This is called *proportional scaling*.

For example, you are running a Deployment with 10 replicas, maxSurge=3, and maxUnavailable=2.

## Pausing and Resuming a Deployment

You can pause a Deployment before triggering one or more updates and then resume it. This allows you to apply multiple fixes in between pausing and resuming without triggering unnecessary rollouts.

## Deployment Status

A Deployment enters various states during its lifecycle. It can be progressing while rolling out a new ReplicaSet, it can be complete, or it can fail to progress.

### Progressing Deployment

Kuberenetes marks a Deployment as *progressing* when one of the following tasks is performed:

- The Deployment creates a new ReplicaSet.
- The Deployment is scaling up its newest ReplicaSet.
- The Deployment is scaling down its older ReplicaSet.
- New Pods become ready or available(ready for at least MinReadySeconds)

You can monitor the progress for a Deployment by using `kubectl rollout status`

### Complete Deployment

Kubernetes marks a Deployment as `complete` when it has the following characteristics:

- All of the replicas associated with the Deployment have been updated to the latest version you've specified, meaning any updates you've requested have been completed.

- All of the replicas associated with the Deployment are available.

- No old replicas for the Deployment are running.

### Failed Deployment

### Operating on a failed deployment

All actions that apply to a complete Deployment also apply to a failed Deployment. You can scale it up/down, roll back to a previous revision, or even pause it if you need to apply multiple tweaks in the Deployment Pod template.

## Clean up Policy

You can set `.spec.revisionHistoryLimit` field in a Deployment to specify how many old ReplicaSets for this Deployment you want to retain. The rest will be garbage-collected in the background. By default, it is 10.

> **Note:** Explictly setting this field to 0, will result in cleaning up all the history of your Deployment thus that Deployment will not be able to roll back.

## Canary Deployment