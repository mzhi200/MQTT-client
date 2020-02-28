package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type clienData struct {
	Id uint32 `json:"id"`
}

func CmdHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	len1 := r.ContentLength
	data := make([]byte, len1)
	r.Body.Read(data)
	clData := clienData{}
	err := json.Unmarshal(data, &clData)
	if err != nil {
		return
	}
	fmt.Printf("clData: %v\n", clData)
}

func server(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/cmd", CmdHandler).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
