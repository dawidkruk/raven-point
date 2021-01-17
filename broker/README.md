# WIP 

Communication with `MQTT` requires a broker to relay the messages across subscribers. There are a lot of possible ways to provision such broker. In this development stage for simpilcity of usage and testing `Eclipse Mosquitto` was chosen:

> Eclipse Mosquitto is an open source (EPL/EDL licensed) message broker that implements the MQTT protocol versions 5.0, 3.1.1 and 3.1. Mosquitto is lightweight and is suitable for use on all devices from low power single board computers to full servers.
> 
> -- *[Mosquitto.org](https://mosquitto.org/)*

Supported options are: 

* Kubernetes environment: 
  * `GKE` cluster 
  * `Minikube` - WIP
  * `Docker Desktop` - WIP 
  * `Docker` - WIP 
* Other ways: 
  * MacOS with `brew` package manager - WIP 

---

###  Kubernetes environment

**`GKE`**: 

You can provision your `Mosquitto` broker on Kubernetes (`GKE`) to be available externally by using: 

* `configmap.yaml` 
* `deployment.yaml` 
* `service.yaml`

> Side notes!
> 
> * `configmap.yaml` is used to pass the `mosquitto.conf` to the `Pod` so that `Mosquitto` will listen on `0.0.0.0`. 
> * `service.yaml` will expose your `Deployment` with `Service` of type `LoadBalancer`! Have it in mind as it will be externally available on that point. 

You can download this repository and use this example on a `GKE` cluster with following commands: 

* `gcloud container clusters get-credentials CLUSTER-NAME --zone=ZONE`

* `$ kubectl apply -f manifests/configmap.yaml`
* `$ kubectl apply -f manifests/deployment.yaml`
* `$ kubectl apply -f manifests/service.yaml`