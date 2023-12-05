package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func main() {

	http.HandleFunc("/msg", service.MsgHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
