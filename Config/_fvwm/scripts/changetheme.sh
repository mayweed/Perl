#!/bin/bash
# $1 = r�pertoire des th�mes
#for i in "$1"/*; do
for i in $(find $1/ -mindepth 1 -maxdepth 1 -type d); do
  echo " + \"%appearance.png%$(basename ${i})\" ChangeTheme \"$(basename ${i})\"";
done
