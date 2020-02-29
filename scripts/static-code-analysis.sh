#!/bin/bash

#move linting to another script...
#add red and stuff to messages
#init 
PATH=${PATH}:$(go env GOPATH)/bin
export PATH

echo 
gofmt -l -s -w ./

golint ./...

golangci-lint run \
    --enable=bodyclose \
    --enable=gofmt \
    --enable=golint \
    --enable=rowserrcheck \
    --enable=gosec \
    --enable=unconvert \
    --enable=dupl \
    --enable=goconst \
    --enable=gocyclo \
    --enable=gocognit\
    --enable=goimports \
    --enable=maligned \
    --enable=depguard \
    --enable=misspell \
    --enable=unparam \
    --enable=scopelint \
    --enable=gocritic \
    --enable=whitespace \
    --enable=goprintffuncname \
    --enable=gomnd
    #--enable=funlen \ -- TODO : Implement
    # --enable=godox \ -- TODO : Implement
    # --enable=gochecknoglobals \  -- TODO : Convert the errors outputted here to objects?

sql_content=$(safesql -v ./)
sql_res=$?

PASSED="\e[1;32mPASSED\e[0m"
FAILED="\e[1;31mFAILED\e[0m"
WARNING="\e[1;33mWARNING\e[0m"

# PASSED_COLOR="\e[32m"
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


# go get github.com/stripe/safesql

# PASSED="\e[1;32mPASSED\e[0m"
# FAILED="\e[1;31mFAILED\e[0m"
# WARNING="\e[1;33mWARNING\e[0m"

# PASSED_COLOR="\e[32m"
# FAILED_COLOR="\e[31m"
# WARNING_COLOR="\e[33m"
# RESET="\e[0m"


# # go get -u github.com/golang/lint/golint
# # go get -u github.com/opennota/check/cmd/defercheck

# # go get -u github.com/opennota/check/cmd/varcheck

# echo "Static Code Analysis........................"

# fmt_content=$(gofmt -l -s -w ./)
# fmt_res=$?

# if [ $fmt_res -eq "0" ]
# then
#     echo -e "  Go with style............................."$PASSED
# else
#     echo -e "  Go with style............................."$FAILED
#     echo -e $FAILED_COLOR $fmt_content $RESET
# fi

# vet_content=$(go vet ./...)
# vet_res=$?

# if [ $vet_res -eq "0" ]
# then
#     echo -e "  Vetted code is better code................"$PASSED
# else
#     echo -e "  Vetted code is better code................"$FAILED
#     echo -e $FAILED_COLOR $vet_content $RESET
# fi

# err_content=$(errcheck ./...)
# err_res=$?

# if [ $err_res -eq "0" ]
# then
#     echo -e "  Get a handle on them errors..............."$PASSED
# else
#     echo -e "  Get a handle on them errors..............."$FAILED
#     echo -e $FAILED_COLOR $err_content $RESET
# fi

# sql_content=$(safesql -v ./)
# sql_res=$?

# if [ $sql_res -eq "0" ]
# then
#     echo -e "  SQL - Keep it secret, keep it safe........"$PASSED
# else
#     if [[ $sql_content == *"supported database driver"* ]]
#     then
#         echo -e "  SQL - Keep it secret, keep it safe........"$WARNING
#         echo -e $WARNING_COLOR $sql_content $RESET
#     else
#         echo -e "  SQL - Keep it secret, keep it safe........"$FAILED
#         echo -e $FAILED_COLOR  $sql_content $RESET
#     fi
# fi

# goreport_content=$(goreportcard-cli)
# goreport_res=$?
# if [ $goreport_res -eq "0" ]
# then
#     echo -e "  Go report card, hope you been studying...."$PASSED
#     echo -e $PASSED_COLOR $err_content $RESET
# else
#     echo -e "  Go report card, hope you been studying...."$FAILED
#     echo -e $FAILED_COLOR $err_content $RESET
# fi


# if [ $vet_res -gt "0" ] || [ $fmt_res -gt "0" ] || [ $err_res -gt "0" ]
# then
#     echo -e "Static Code Analysis........................"$FAILED
#     exit 1
# else
#     echo "Static Code Analysis........................"$PASSED
# fi