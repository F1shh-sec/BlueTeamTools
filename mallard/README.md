# Mallard
The sledge hammer for the fly

# Project Overview
The goal of this project is to create a general purpose blue teaming tool capable of locking down a box in the first few minutes of a competition. 

It is extremely hard to automate securing every attack vector. Rather then striving for 100% accuracy, the goal is to automate my five minute plan allowing me to get right into threat hunting. This tool also aims to prevent the creation of new reverse shells, Services, and users.

## What this tool is not
This tool is not a general purpose Anti-virus/Anti-malware tool. The rules this tool enforces are extremely strict and destructive. **If you were to deploy this on a production server, you would have a very bad time. Don't do that.** This is a competition tool only.

This tool comes with no warranty or support, use at your own risk.

# USAGE:
|Command| Description|
|-|-|
|passwd|Changes all users passwords|
|users|Gets a list of all users who have a usable shell|
|disable|Disables shell access for all users with a shell|
|info|Runs Info collection script|
|conn|Prints out all active TCP/UDP connections|


|Flag|Description|
|-|-|
|-a|Aggressive mode. Runs user, passwd, and disable once the program launches|