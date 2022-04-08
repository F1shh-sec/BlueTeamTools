#!/bin/bash
mapfile -t usersArray < <(ss -tulpn | awk -F"users:" '{print $2}' | awk -F"\"" '{print $2}')
# shellcheck disable=SC2068
for elm in ${usersArray[@]};
do
  echo $elm
done

