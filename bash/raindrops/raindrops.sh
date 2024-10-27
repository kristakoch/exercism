#!/usr/bin/env bash

raindrops="" 

if [[ $1%3 -eq 0 ]]; then
  raindrops+="Pling"
fi 

if [[ $1%5 -eq 0 ]]; then 
  raindrops+="Plang"
fi 

if [[ $1%7 -eq 0 ]]; then 
  raindrops+="Plong"
fi 

echo ${raindrops:-$1}