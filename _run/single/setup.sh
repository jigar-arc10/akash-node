#!/bin/bash

docker image ls --filter='reference=ovrclk-local:*' --format '{{.Repository}}:{{.Tag}}' | xargs -L 1 docker image rm

make clean
make kind-cluster-delete

set -xe
make kind-cluster-create

make init
make kustomize-init

make kustomize-init-docker-image
make kind-configure-image
make kind-upload-image

sleep 95 # waiting on the ingress-nginx to be ready. Need a better way to do this
# try kubectl wait pod

make kustomize-install-node

sleep 6

make provider-create

sleep 1
make kustomize-install-ip-operator
make kustomize-install-hostname-operator


make kustomize-install-provider
