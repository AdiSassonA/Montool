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

"-n" for Node Exporter version.

"-g" for Grafana version.

*Note: you need to specify the full tag version's name, for exaple:
```
montool create -p v1.1.1
```

"-r" for Prometheus retention
*Notice that you need to mention the unit of time, "h" for hours \ "d" for days, for example:
```
montool create -r 100h
```

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
