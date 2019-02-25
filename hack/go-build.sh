#!/bin/sh

REGISTRY=quay.io/kiegroup
IMAGE=kie-cloud-operator
TAG=1.0

go generate ./...
if [[ -z ${CI} ]]; then
    source hack/go-test.sh
    operator-sdk build ${REGISTRY}/${IMAGE}:${TAG}
    if [[ ${1} == "rhel" ]]; then
        if [[ ${2} == "release" ]]; then
            CFLAGS="--build-osbs-release"
        fi
        cekit build ${CFLAGS} \
            --redhat \
            --build-tech-preview \
            --package-manager=microdnf \
            --build-engine=osbs \
            --build-osbs-target=rhba-7.3-openshift-containers-candidate # rhpam-7-rhel-7-containers-candidate
    fi
else
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o build/_output/bin/kie-cloud-operator github.com/kiegroup/kie-cloud-operator/cmd/manager
fi
