package main

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type UserdData struct {
	logLevel logLevelFlag
	client   mqtt.Client
	event    chan mqtt.Message
}

func (ue UserdData) subscribeEven(ty string) (err error){
	topic := fmt.Sprintf("$sys/%s/%s/%s", config.OneNet.Username, config.OneNet.ClientId, ty)
	if token := ue.client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("%+v", token.Error()))
		return
	}
	return
}

func (ue UserdData) unSubscribeEven(ty string) (err error){
	topic := fmt.Sprintf("$sys/%s/%s/%s", config.OneNet.Username, config.OneNet.ClientId, ty)
	if token := ue.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("%+v", token.Error()))
		return
	}
	return
}

func (ue UserdData) publishData(ty string, payload interface{}) (err error) {
	topic := fmt.Sprintf("$sys/%s/%s/%s", config.OneNet.Username, config.OneNet.ClientId, ty)
	//token := ue.client.Publish(topic, 1, false, payload)
	//token.Wait()
	if token := ue.client.Publish(topic, 1, false, payload); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("%+v", token.Error()))
		return
	}
	return
}