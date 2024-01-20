#!/bin/sh

kubectl delete pods,service,deployment,ingress,clusterrole,clusterrolebinding,role --all --force=true
