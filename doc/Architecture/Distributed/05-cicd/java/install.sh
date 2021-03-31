#!/bin/bash

echo "Deploying jdk1.8.0_241..."
if [ ! -e "/usr/local/java/jdk1.8.0_241" ];then
	echo "jdk1.8.0_241 does not exist, ready to deploy"
	tar -C /usr/local/java/ -xvf jdk-8u241-linux-x64.tar.gz
fi
echo "jdk1.8.0_241 deployed"

sleep 3

echo "Deploying apache-maven-3.6.1..."
if [ ! -e "/usr/local/java/apache-maven-3.6.1" ];then
	echo "apache-maven-3.6.1 does not exist, ready to deploy"
	tar -C /usr/local/java/ -xvf apache-maven-3.6.1-bin.tar.gz
fi
echo "apache-maven-3.6.1 deployed"