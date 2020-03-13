#!/bin/bash
if [ ! -d "./dist" ]; then
    mkdir dist
fi
currentProjectName=${PWD##*/} 
zip ./dist/${currentProjectName}-package.zip ./bin/*.*