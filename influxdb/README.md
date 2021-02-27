### InfluxDB integration 

This is the first iteration of integration between `raven-point` and `InfluxDB`. 

Code is in the ALPHA stage. 

#### Goals: 
* Introduce the data structure sent between services that will be used for data collection and further analysis. 
* Create a `Go` code that will be sending specific payload to `InfluxDB` instance running on Kubernetes. 
* Integrate code from above point with the `data-collector`

### Running `InfluxDB2` on Kubernetes: 

You can run `InfluxDB` by following the official documentation: 

* *[Docs.Influxdata.com: Influxdb: v2.0: Get started](https://docs.influxdata.com/influxdb/v2.0/get-started/)*

Quickstart: 

* `$ kubectl apply -f https://raw.githubusercontent.com/influxdata/docs-v2/master/static/downloads/influxdb-k8-minikube.yaml`
* `$ kubectl port-forward -n influxdb service/influxdb 8086:8086`