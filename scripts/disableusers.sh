#!/bin/bash

currentuser=$(whoami)
mapfile -t usersArray < <(awk -v curuser="$currentuser" -F":" '((($7=="/bin/bash")||($7=="/bin/sh"))&&(($1!="root")&&($1!=curuser))){print $1}' /etc/passwd)
echo "Found Users: " "${usersArray[@]}"
# shellcheck disable=SC2068
for elm in ${usersArray[@]};
do
  echo "Disabling: " "$elm"
  usermod -s /sbin/nologon $elm
  killall -u $elm
  skill -kill -u $elm
done
