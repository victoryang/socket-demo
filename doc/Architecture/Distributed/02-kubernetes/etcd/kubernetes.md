# Etcd in kubernetes

[etcdctl 访问 kubernetes](https://jimmysong.io/kubernetes-handbook/guide/using-etcdctl-to-access-kubernetes-data.html)


## Get keys from Kubernetes Cluster

```bash
#!/bin/bash

keys=`./etcdctl --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/peer.crt --key=/etc/kubernetes/pki/etcd/peer.key get /registry --prefix -w json|python -m json.tool|grep key|cut -d ":" -f2|tr -d '"'|tr -d ","`
for x in $keys;do
  echo $x|base64 -d|sort
done
```