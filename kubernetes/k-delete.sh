#!/bin/bash

# Delete Ingress resources
kubectl delete ingress "gojo-http-ingress" --force=true
kubectl delete ingress "gojo-grpc-ingress" --force=true

delete_resource() {
    local resource="$1"
    kubectl delete "$resource" --force=true 
    if [ $? -ne 0 ]; then
        echo "Error deleting $resource" 
    fi
}

# Delete gojo-migration
delete_resource "job.batch/gojo-migration"

# Delete gojo-postgres
delete_resource "service/gojo-postgres-service"
delete_resource "deployment/gojo-postgres-deployment"
delete_resource "configmap/postgres-config"

# Delete gojo-queue
delete_resource "service/gojo-queue-service"
delete_resource "deployment/gojo-queue-deployment"

# Delete gojo-cache
delete_resource "service/gojo-cache-service"
delete_resource "deployment/gojo-cache-deployment"

# Delete gojo-meili
delete_resource "service/gojo-meili-service"
delete_resource "deployment/gojo-meili-deployment"

# Delete gojo-api
delete_resource "service/gojo-api-service"
delete_resource "deployment/gojo-api-deployment"
delete_resource "configmap/api-config"

echo "Deletion completed."
