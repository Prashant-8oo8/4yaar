package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func main(){
	mux:= http.NewServeMux()

	mux.HandleFunc("GET /api/health",func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":"ok",
			"app":"4Yaar",
		})
	})

	log.Println("4yaar is running... on : 8000")
	
	err:=http.ListenAndServe(":8000",mux)
	if err!=nil{
		log.Fatalf("server is not running...")
	}
}