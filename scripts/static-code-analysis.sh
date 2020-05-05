#!/bin/bash

#move linting to another script...
#add red and stuff to messages
#init 
PATH=${PATH}:$(go env GOPATH)/bin
export PATH

go mod tidy
gofmt -l -s -w ./

golangci-lint run -c ./.golangci.yml  

sql_content=$(safesql -v ./)
sql_res=$?

PASSED="\e[1;32mPASSED\e[0m"
FAILED="\e[1;31mFAILED\e[0m"
WARNING="\e[1;33mWARNING\e[0m"

FAILED_COLOR="\e[31m"
WARNING_COLOR="\e[33m"
RESET="\e[0m"

if [ $sql_res -eq "0" ]
then
    echo -e "  SQL - Keep it secret, keep it safe........$PASSED"
else
    if [[ $sql_content == *"supported database driver"* ]]
    then
        echo -e "  SQL - Keep it secret, keep it safe........$WARNING"
        echo -e "$WARNING_COLOR $sql_content $RESET"
    else
        echo -e "  SQL - Keep it secret, keep it safe........$FAILED"
        echo -e "$FAILED_COLOR  $sql_content $RESET"
    fi
fi