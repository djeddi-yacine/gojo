#!/bin/bash

apply_manifest() {
    local manifest="$1"
    kubectl apply -f "$manifest"
    if [ $? -ne 0 ]; then
        echo "Error applying $manifest"
        exit 1
    fi
}

kubectl apply -f "https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml"

# Apply gojo-api-cert
apply_manifest "cluster-issuer.yaml"

# Apply gojo-ingrees
apply_manifest "ingress-nginx.yaml"
apply_manifest "ingress-http.yaml"
apply_manifest "ingress-grpc.yaml"

# Apply gojo-postgres
apply_manifest "postgres-configmap.yaml"
apply_manifest "postgres-deployment.yaml"
apply_manifest "postgres-service.yaml"

# Wait for gojo-postgres to become healthy (adjust the timeout as needed)
kubectl wait --for=condition=ready pod -l app=gojo-postgres --timeout=300s

# Apply gojo-meili
apply_manifest "meili-configmap.yaml"
apply_manifest "meili-deployment.yaml"
apply_manifest "meili-service.yaml"

# Wait for gojo-meili to become healthy (adjust the timeout as needed)
kubectl wait --for=condition=ready pod -l app=gojo-meili --timeout=300s

# Apply gojo-migration
apply_manifest "migration-job.yaml"

# Apply gojo-queue
apply_manifest "queue-deployment.yaml"
apply_manifest "queue-service.yaml"

# Wait for gojo-queue to become healthy (adjust the timeout as needed)
kubectl wait --for=condition=ready pod -l app=gojo-queue --timeout=300s

# Apply gojo-cache
apply_manifest "cache-deployment.yaml"
apply_manifest "cache-service.yaml"

# Wait for gojo-cache to become healthy (adjust the timeout as needed)
kubectl wait --for=condition=ready pod -l app=gojo-cache --timeout=300s

# Apply gojo-api
apply_manifest "api-configmap.yaml"
apply_manifest "api-deployment.yaml"
apply_manifest "api-service.yaml"

echo "Application deployment completed successfully."
