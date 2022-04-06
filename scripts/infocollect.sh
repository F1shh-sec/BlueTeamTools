#!/bin/bash

echo "RUNNING PROCESSES:"
echo "--------------"
ps -aux
echo "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"

echo "RUNNING SERVICES:"
echo "--------------"
systemctl list-units --type=service --state=running
echo "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"

echo "CRONTAB:"
echo "--------------"
crontab -l
echo "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"

echo "ACTIVE TCP CONNECTIONS:"
echo "--------------"
ss -tulpn
echo "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"