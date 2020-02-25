package main

import (
	"errors"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type UserdData struct {
	logLevel logLevelFlag
	client   mqtt.Client
	event    chan mqtt.Message
}

func (ue UserdData) subscribeEven(ty string) (err error){
	topic := fmt.Sprintf(TopicPublicSection, config.OneNet.ProductId, config.OneNet.EquipName, ty)
	if token := ue.client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("%+v", token.Error()))
		return
	}
	return
}

func (ue UserdData) unSubscribeEven(ty string) (err error){
	topic := fmt.Sprintf(TopicPublicSection, config.OneNet.ProductId, config.OneNet.EquipName, ty)
	if token := ue.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("%+v", token.Error()))
		return
	}
	return
}

func (ue UserdData) publishData(payload interface{}) (err error) {
	pl, err := json.Marshal(payload)
	if err != nil {
		return
	}

	topic := fmt.Sprintf(TopicPublicSection, config.OneNet.ProductId, config.OneNet.EquipName, TopicDpUplink)
	//token := ue.client.Publish(topic, 1, false, payload)
	//token.Wait()
	if token := ue.client.Publish(topic, 1, false, pl); token.Wait() && token.Error() != nil {
		err = errors.New(fmt.Sprintf("%+v", token.Error()))
		return
	}
	return
}