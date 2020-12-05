# Montool
 Montool spins up a monitoring environment that include Prometheus, Node Exporter, cAdvisor, and Grafa.
 Montool takes 

# SetUp
1. Make sure you have docker-compose and Golang version 1.13 or above installed on your computer.

2. Run:
```
go install montool
```
And now you are good to go.

# Usage
You can run "--help" \ "-h" for the montool command and subcommands.
To create your monitoring environment just simply run 
```
montool create
```
That spins up all components with the latest versions, and with ten day as for the Prometheus retention.
