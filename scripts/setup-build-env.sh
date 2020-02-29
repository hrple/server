#!/bin/bash

golangci_content=$(golangci-lint --version)
golangci_res=$?
if [ $golangci_res -gt "0" ]
then
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.7	    
fi