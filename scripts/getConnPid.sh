#!/bin/bash

processname=$1
mapfile -t usersArray < <(ss -tulpn | awk -F"\"$processname\"" '{print $2}' | awk -F"," '{print $2}' | awk -F"=" '{print $2}' )
# shellcheck disable=SC2068
for elm in ${usersArray[@]};
do
  echo $elm
done

