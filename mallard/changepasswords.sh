#!/bin/bash

newPassword="Sup3rS3curep@ss"
mapfile -t usersArray < <(awk -F":" '# ((($7=="/bin/bash")||($7=="")||($7=="/bin/sh"))&&($1!="root")){print $1}' /etc/passwd)
echo "Found Users:" "${usersArray[@]}"
# shellcheck disable=SC2068
for elm in ${usersArray[@]};
do
  echo $elm
  echo -e $newPassword'\n'$newPassword'\n' | passwd $elm;
done
