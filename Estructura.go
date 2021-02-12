package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type persona struct {
	edad   int    `json: edad`
	nombre string `json:nombre`
}

func CreateHuman(w http.ResponseWriter, r *http.Request) {
	var newpersona persona

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newpersona)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newpersona)
}

var per persona
