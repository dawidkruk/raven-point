# Raven Point

Raven Point it's a project to build a flexible infrastructure provisioned on Kubernetes capable of sending, receiving and analyzing GPS data from embeeded devices (like `Raspberry Pi` or `Arduino`).

Due to being in early alpha stages there are some shortcuts taken that will be addressed in the further stages of development. 

--- 

Description of components (where each has it's own directory): 

* `broker` - `mosquitto` broker responsible for the communication between `sender` and `collector`
* `sender` - early implementation of a sender responsible for sending the data to a broker. As for now this is a Docker image with an `app` written in `Go` that sends the example payload to broker
* `collector` - early implementation of a collector responsible for collecting the data from a broker and it in `influxdb`.
* `influxdb` - database that will store the payload sent by a `collector`
* `analyzer` - frontend that will show and analyze acquired data (TBD)

---

Current tasks: 

* modify the whole documentation of this project to be more descriptive and informative

* `broker` - secure the broker (authorization and authentication of subscribers)
* `sender` - design the payload that will be sent and start working on an embeeded device 
* `collector` - refactor the code as in the `sender` and implement the part responsible for sending data to `influxdb`
* `influxdb` - research best practises and refactor the code responsible sending data to it 
* `analyzer` - task on hold, other tasks needs to be finished before engaging on this particular one. This should be only `POC` for the time being


> Disclaimer!
> 
> Please consider this documentation to be **work in progress**.

---

Feel free to point potential issues, report bugs and contribute to this project.



