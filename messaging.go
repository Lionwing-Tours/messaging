package main

import (
	"log"
	"mes"
	"messaging/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/sendEmail", SendEmail)
	http.HandleFunc("/sendSMS", SendSMS)
	log.Fatal(endless.ListenAndServe(":3000", nil))
}
