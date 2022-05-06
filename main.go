package main

import (
	"io/ioutil"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/udonetsm/help/helper"
)

var Sc stan.Conn

func Connect(clientid string) {
	cluster := "test-cluster"
	url := "nats://127.0.0.1:4222"
	sc, err := stan.Connect(cluster, clientid, stan.NatsURL(url))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connetced successfully")
	Sc = sc
}

func Pub(data []byte, sub string) {
	err := Sc.Publish(sub, data)
	if err != nil {
		log.Println("Fail publish with ", err)
		return
	}
}

func main() {
	Pub([]byte(Data_json), "sub")
}

var Data_json string

func init() {
	bytesdata, err := ioutil.ReadFile(helper.Home() + "/Documents/L0/model.json")
	helper.Errors(err, "ioutilreadfile(init)")
	Data_json = string(bytesdata)
	Connect("-")
}
