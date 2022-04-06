#!/bin/bash

mapfile -t usersArray < <(awk -F":" '($7=="/bin/bash"||$7=="/bin/sh"){print $1}' /etc/passwd)
echo "Found Users:" "${usersArray[@]}"