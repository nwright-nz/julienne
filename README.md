# Julienne - slicing knife application info to CSV

After experiencing some pain with getting a clean export of node package data into CSV (JSON output from knife is hard to parse!), I put together this quick tool to do this for me. 

To get a list of all the packages installed on all your nodes registered to your Chef Infra server, you can run this fun command :   `knife search node 'name:*' -a name -a os -a os_version -a packages -F j  > nodeinfo.json`

There is probably a better way to run this, but this works for me and provides some additional info like OS and version. From here, to get a useful CSV representation of this (for filtering and doing whatever Excel black magic people do) you can use Julienne: 

`julienne -path nodeinfo.json -output nodeinfo.csv`

## Building
The makefile is pretty self explanatory, but the TLDR is :
`make build-windows` (builds for Windows)   
`make build-linux` (builds for Linux)   
`make` (builds for OSX)

This application serves a direct need for me, YMMV
