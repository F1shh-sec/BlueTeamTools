#!/bin/bash

newPassword="Sup3rS3curep@ss"
mapfile -t usersArray < <(awk -F":" '((($7=="/bin/bash")||($7=="/bin/sh"))&&($1!="root")){print $1}' /etc/passwd)
# shellcheck disable=SC2068
for elm in ${usersArray[@]};
do
  echo "Changed Password for $elm"
  echo -e $newPassword'\n'$newPassword'\n' | passwd $elm;

done
