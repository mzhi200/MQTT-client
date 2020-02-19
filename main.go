package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	//"log"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

type mqttUserdData struct {
	log mqttLog
}

func main() {
	ueData := mqttUserdData{}
	ueData.log = &mqttClientLog{}
	ueData.log.Init(os.Stdout, LogLevelDebug|LogLevelInfo|LogLevelError)
	ueData.log.Info("MQTT-Client start...")

	//mqttPackageTraceInit(os.Stdout, MqttTraceTypeDebug|MqttTraceTypeError, 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://183.230.40.96:1883").SetClientID("Test-Mqtt")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	opts.Username = "309606"
	opts.Password = "version=2018-10-31&res=products%2F309606%2Fdevices%2FTest-Mqtt&et=1672735919&method=md5&sign=w%2BZ1NmnqrT4B6MsEGbVxiA%3D%3D"

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connect error: ", token.Error())
		panic(token.Error())
	}

	if token := c.Subscribe("$sys/309606/Test-Mqtt/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for i := 0; i < 5; i++ {
		text := date{
			Id: 123,
			Dp: DpType{
				Tp: []TpData{
					{
						V: 30,
					},
				},
				Power: []PowerData{
					{
						V: 4.5,
					},
				},
			},
		}
		textString, err := json.Marshal(text)
		if err != nil {
			ueData.log.Error("Failed to Marshal(text); Err: %v",err)
			return
		}

		token := c.Publish("$sys/309606/Test-Mqtt/dp/post/json", 1, false, textString)
		token.Wait()
		time.Sleep(20 * time.Second)
	}

	time.Sleep(6 * time.Second)
	/*
		if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
					fmt.Println(token.Error())
							os.Exit(1)
								}
	*/
	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
