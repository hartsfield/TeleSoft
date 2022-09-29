package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hartsfield/gmailer"
)

type newMail struct {
	From string `json:"from"`
	Body string `json:"body"`
}

// ajaxResponse is used to respond to ajax requests with arbitrary data in the
// format of map[string]string
// func ajaxResponse(w http.ResponseWriter, res map[string]string) {
// 	w.Header().Set("Content-Type", "application/json")
// 	err := json.NewEncoder(w).Encode(res)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func marshalMailData(r *http.Request) (*newMail, error) {
	t := &newMail{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.tmpl", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	m, _ := marshalMailData(r)
	msg := gmailer.Message{
		Recipient: "johnathanhartsfield@gmail.com",
		Subject:   "ALERT! NEW JOB REQUEST!",
		Body:      m.From + "::" + m.Body,
	}
	msg.Send(onMessageSent)
}

func onMessageSent() {
	fmt.Println("new email")
}
