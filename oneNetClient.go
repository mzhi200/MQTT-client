package main

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var receiveMqttMsgFromPeer mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	ue.event <- msg
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MSG: %s\n", msg.Payload())
}

func oneNetClientConfig(config Configuration) (*mqtt.ClientOptions, error) {
	//mqttPackageTraceInit(os.Stdout, MqttTraceTypeDebug|MqttTraceTypeError, 0)
	ipPort := fmt.Sprintf("tcp://%s:%d", config.OneNet.Server.Host, config.OneNet.Server.Port)
	opts := mqtt.NewClientOptions().AddBroker(ipPort).SetClientID(config.OneNet.EquipName)
	opts.SetKeepAlive(time.Duration(config.OneNet.KeepAlive) * time.Second)
	opts.SetDefaultPublishHandler(receiveMqttMsgFromPeer)
	opts.SetPingTimeout(time.Duration(config.OneNet.PingTimeout) * time.Second)
	opts.Username = config.OneNet.ProductId

	tk := &Token{}
	tokData, err := tk.TokenGenerateFun()
	if err != nil {
		return nil, err
	}
	opts.Password = tokData
	//opts.Password = "version=2018-10-31&res=products%2F309606%2Fdevices%2FTest-Mqtt&et=1672735919&method=md5&sign=w%2BZ1NmnqrT4B6MsEGbVxiA%3D%3D"

	return opts, err
}

func oneNetConnect(config Configuration) (c mqtt.Client, err error) {
	opts, err := oneNetClientConfig(config)
	if err != nil {
		return nil, err
	}
	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("Connect error: %+v", token.Error()))
		return
	}
	return
}

