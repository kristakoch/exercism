#!/usr/bin/env bash

usage="Usage: error_handling.sh <person>"
if [ $# -ne 1 ]; then 
  echo $usage && exit 1
fi 

echo "Hello, $1" 