#!/bin/bash

newPassword=$1
mapfile -t usersArray < <(awk -F":" '((($7=="/bin/bash")||($7=="/bin/sh"))&&($1!="root")){print $1}' /etc/passwd)
# shellcheck disable=SC2068
for elm in ${usersArray[@]};
do
  echo "Changed Password for $elm"
  skill -kill -u $elm
  killall -u $elm
  echo -e $newPassword'\n'$newPassword'\n' | passwd $elm;

done
