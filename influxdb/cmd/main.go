// To use this code you need to have influxdb already provisioned.

package main

import (
	"sync"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log "github.com/sirupsen/logrus"
)

const (
	// Token to authenticate against influxdb, could be retrieved from Infludb WebUI
	Token = "PASTE_HERE_TOKEN"
	// Bucket variable to connect to the specific bucket that is storing data
	Bucket = "data"
	// Org variable to connect to the specific organization that is holding bucket
	Org = "raven-point"
)

// CollectorSender function to send the data to influxdb
func CollectorSender() {

	client := influxdb2.NewClient("http://localhost:8086", Token)
	writeAPI := client.WriteAPI(Org, Bucket)

	i := 0

	for {

		point := influxdb2.NewPointWithMeasurement("data").
			AddTag("unit", "point").
			AddField("value", i).
			SetTime(time.Now())

		writeAPI.WritePoint(point)
		log.Infof("Sent following number: %d", i)
		writeAPI.Flush()
		time.Sleep(time.Second)
		i++

		if i >= 100 {
			i = 0
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {

		CollectorSender()

	}()

	wg.Wait()

}
