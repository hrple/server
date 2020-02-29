#!/bin/bash

DIR="bin"
if [ -d "$DIR" ]; then
  ### Take action if $DIR exists ###
  echo "bin folder exists..."
else
  ###  Control will jump here if $DIR does NOT exists ###
  mkdir bin
fi