#!/bin/bash

kubeadm init --kubernetes-version v1.18.3 --image-repository registry.aliyuncs.com/google_containers --pod-network-cidr=10.24.0.0/16 --v=5
