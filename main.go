package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	config Configuration
	ue     UserdData
)
func main() {
	var err error
	ue.logLevel = LogLevelInfo
	log := newLog(nil)
	log.Info("MQTT-Client start...")

	err = InitConfig()
	if err != nil {
		log.Error("Failed to InitConfig; Err: %v", err)
		return
	}
	log.Info("Got Configuration %+v.", config)
	ue.logLevel, err = GetLogLevelFromConfig()
	if err != nil {
		log.Error("Failed to GetLogLevelFromConfig; Err: %v", err)
		return
	}
	//update logLevel
	log = newLog(nil)

	ue.client, err = oneNetConnect(config)
	if err != nil {
		log.Error("Failed to Connect; Err: %v", err)
		return
	}
	//receive chan
	ue.event = make(chan mqtt.Message, 1)
	defer close(ue.event)

	//cmd
	go triggerCommandFun()

	go server("8000")

	//Handle Event
	for {
		select {
		case msg := <-ue.event:
			log.Debug("TOPIC: %s", msg.Topic())
			log.Debug("MSG: %s", msg.Payload())
		}
	}
}
