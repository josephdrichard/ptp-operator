#!/bin/bash

vendor/k8s.io/code-generator/generate-groups.sh client,lister,informer \
	github.com/josephdrichard/ptp-operator/pkg/client \
	github.com/josephdrichard/ptp-operator/pkg/apis \
	ptp:v1
