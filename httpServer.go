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

	dv := newDeviceCb()
	dv.deviceId = uint32(id)

	ev := event{}
	switch deviceMsg.Event {
	case "login":
		ev.eventID = LoginEv
	case "logout":
		ev.eventID = LogoutEv
	case "data":
		ev.eventID = DataEv
		ev.msg = deviceMsg.Data
	default:
		log.Error("Unknown Event: %s", deviceMsg.Event)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dv.chEvent <- ev
	w.WriteHeader(http.StatusOK)
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
