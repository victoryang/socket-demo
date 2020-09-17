#!/bin/bash

SecretName=`kubectl get secrets -n kubernetes-dashboard | grep kubernetes-dashboard-token | awk '{ print $1 }'`

kubectl describe secret kubernetes-dashboard-token-9k6f6 -n kubernetes-dashboard