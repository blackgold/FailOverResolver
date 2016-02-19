FailOverResolver : Health checks the remote service and returns one host as per the selection criteria


BUILD:
make all


RUN:
./bin/for


INTEGRATION:
Spin up service in go routine (look at for.go). Use Resolve function in your code to find healthy server.




CONFIGURATION:

Each service to health check must be placed in a different config file in json format. Look for sample config file in "config" directory. Servicename must be unique and this string is used to query for healthy instance. ttl feild in Algorithm determines how often to health check the service. uri in Servers is used to health check. Onlly http is supported.




CHECK STATUS:

List all services: curl localhost/services

List the configuration of specific service: curl localhost/services/${servicename}

List the healthcheck stats of specific service: curl localhost/services/${servicename}/stats

Get Hostname for a specific service: curl localhost/services/${servicename}/resolve



TODO:

Extend the functionality to record the rtt time to each server.

Extend the functionality to get load feedback information from each server.

Extend the supported algorithms. Right now simple randomizer is used to pick a healthy server.
