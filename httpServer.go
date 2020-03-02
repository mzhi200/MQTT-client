package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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
