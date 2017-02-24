#!/usr/bin/env bash

export SPARK_HOME=/usr/local/spark
export JAVA_HOME=/usr

export PATH=$PATH:$SPARK_HOME/bin

# Force IPv4
export SPARK_DAEMON_JAVA_OPTS="-Djava.net.preferIPv4Stack=true"

export SPARK_WORKER_CORES=2
export SPARK_STANDALONE_MASTER=master.spark.dev
export SPARK_MASTER_HOST=192.168.52.201
export SPARK_LOCAL_IP=$(ip addr show eth1 | awk '$1 == "inet" {gsub(/\/.*$/, "", $2); print $2}')
