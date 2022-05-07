package main

import (
	"io/ioutil"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/udonetsm/help/helper"
)

func ConnectAndPublish(clientid, clusterid, url, sub string) {
	sc, err := stan.Connect(clusterid, clientid, stan.NatsURL(url))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connetced successfully")
	data := GetData()
	Pub(data, sub, sc)
}

func Pub(data []byte, sub string, sc stan.Conn) {
	err := sc.Publish(sub, data)
	if err != nil {
		log.Println("Fail publish with ", err)
		return
	}
}

func main() {
	ConnectAndPublish("-", "test-cluster", "nats://127.0.0.1:4222", "orders")
}

func GetData() []byte {
	bytesdata, err := ioutil.ReadFile(helper.Home() + "/Documents/L0/model.json")
	helper.Errors(err, "ioutilreadfile")
	return bytesdata
}
