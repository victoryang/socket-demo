#!/bin/bash

export MAVEN_HOME=/usr/local/java/apache-maven-3.6.1
export JAVA_HOME=/usr/local/java/jdk1.8.0_241
export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
export PATH=$MAVEN_HOME/bin:$JAVA_HOME/bin:$PATH