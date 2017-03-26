package handlers

import (
	"log"
	"messaging/mail"
	"net/http"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	req.ParseForm()
	htmlContent := r.FormValue("html")
	textContent := r.FormValue("text")
	subject := r.FormValue("subject")
	toName := r.FormValue("toName")
	fromName := r.FormValue("fromName")
	toAddr := r.FormValue("toAddr")

	if len(fromName) == 0 {
		fromName = "LionWing Tours"
	}

	err := mail.SendEmail(toAddr, toName, fromName, subject, htmlContent, textContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func SendSMS(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	countryCode := r.FormValue("countryCode")
	number := r.FormValue("number")
	message := r.FormValue("message")
	log.Println(countryCode+number, message)
	//go sms.SendSMS(countryCode, number, message)
	w.WriteHeader(http.StatusOK)
}
