package main

import "net/http"

func checkReadiness(w http.ResponseWriter, r *http.Request){
	respondwithJSON(w, 200, map[string]string{"ready": "OK"})
}
