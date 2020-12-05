# Montool
 Montool spins up a monitoring environment that include Prometheus, Node Exporter, cAdvisor, and Grafana.
 
 Montool accepts parameters to set different versions of the components. 
 

# Set Up
1. Make sure you have docker-compose and Golang version 1.13 or above installed on your machine.

2. Run:
```
go install montool
```
And now you are good to go.

# Usage
You can run "--help" \ "-h" for the montool command and subcommands.

## Create

To create your monitoring environment just simply run 
```
montool create
```
That spins up all components with the latest versions, and with a Prometheus retention of 10 days.

Flags:

"-p" for Prometheus version.
"-r" for Prometheus retention.
"-n" for Node Exporter version.
"-g" for Grafana version.

## Show
Use show to see all montool containers. 
```
montool show
```

## Remove 
To stop and remove all montool containers:
```
montool remove
```
