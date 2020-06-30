package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMessengerList(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("datasources/messenger.json")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(file))
}
