package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"strconv"
)

type DeviceData struct {
	Power float32 `json:"power,omitempty"`
	Tp    uint32 `json:"tp,omitempty"`
}

type DeviceMsg struct {
	Event string `json:"event"`
	Data  DeviceData `json:"data,omitempty"`
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	log := newLog(nil)
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	deviceId := vars["device-id"]
	log.Debug("Device-id = %v", deviceId)

	len1 := r.ContentLength
	data := make([]byte, len1)
	r.Body.Read(data)
	deviceMsg := DeviceMsg{}
	err := json.Unmarshal(data, &deviceMsg)
	if err != nil {
		return
	}
	fmt.Printf("deviceMsg: %v\n", deviceMsg)

	id, err := strconv.ParseUint(deviceId, 10, 32)
	oneNetMsg := date{Id: uint32(id)}
	switch deviceMsg.Event {
	case "login":
		oneNetMsg.Dp.MsgType = []MessageType{{V: LOGIN,}}
		err := ue.db.Set(deviceId, deviceMsg.Event, 0).Err()
		if err != nil {
			panic(err)
		}
	case "logout":
		oneNetMsg.Dp.MsgType = []MessageType{{V: LOGOUT,}}
		err := ue.db.Del(deviceId).Err()
		if err != nil {
			panic(err)
		}

	case "Data":
		oneNetMsg.Dp.MsgType = []MessageType{{V: DP,}}
		val, err := ue.db.Get(deviceId).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("key: %s, val: %s\n", deviceId, val)

	default:
		log.Error("Unknow Event: %s", deviceMsg.Event)
		return
	}
	ue.publishData(oneNetMsg)
	if err != nil {
		log.Error("Failed to Publish; Err: %v", err)
		return
	}
/*
	if ue.db != nil {
		err := ue.db.Set(deviceId, deviceMsg.Event, 0).Err()
		if err != nil {
			panic(err)
		}
		val, err := ue.db.Get(deviceId).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("key: %s, val: %s\n", deviceId, val)
	}
*/
}

func server(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/device/{device-id}/event", EventHandler).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
