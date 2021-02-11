Communication with `MQTT` requires a broker to relay the messages across subscribers. There are a lot of possible ways to provision such broker. In this development stage for simpilcity of usage and testing `Eclipse Mosquitto` was chosen:

> Eclipse Mosquitto is an open source (EPL/EDL licensed) message broker that implements the MQTT protocol versions 5.0, 3.1.1 and 3.1. Mosquitto is lightweight and is suitable for use on all devices from low power single board computers to full servers.
> 
> -- *[Mosquitto.org](https://mosquitto.org/)*


###  General idea 

You can provision your `Mosquitto` broker on Kubernetes by using following manifests: 

* `configmap.yaml` 
* `deployment.yaml` 
* `service.yaml`

> Side notes!
> 
> * `configmap.yaml` is used to pass the `mosquitto.conf` to the `Pod` so that `Mosquitto` will listen on `0.0.0.0`
> * `service.yaml` will expose your `Deployment` with `Service` of type `LoadBalancer`. Please review the documentation of used Kubernetes solution on the premise of handling the `External IP` assignment. 

You can download this repository and use this example with following commands: 

* `$ kubectl apply -f manifests/configmap.yaml`
* `$ kubectl apply -f manifests/deployment.yaml`
* `$ kubectl apply -f manifests/service.yaml`

---

### Current tasks: 

> A side note!
> >
> The tasks importance is placed in a descending fashion. 

* `broker` - secure the broker (authorization and authentication of subscribers)
* modify the `broker` documentation to support multiple Kubernetes solutions (`GKE`, `Minikube`, `Docker Desktop`)

--- 

### Mac OS

Alternatively you can use the `brew` package manager on `Mac OS` and install `mosquitto` by following commands: 

* `$ brew install mosquitto`
* `$ brew services run mosquitto`

---

