package shared

import (
	"strings"
	"net/http"
	"net/url"
)

type SMSDeliverer interface {
	SendMessage(phone string, message string)
}

type TwillioSMSDeliverer struct {}


func NewTwillioSMSDeliverer(){

}

func (self *TwillioSMSDeliverer) SendMessage(phone string, message string) {
	accountSid := "ACb9a936a46ff54a17e2cc076229412c27"
	authToken := "4435c0b1dac67d59c605134ea7e8faab"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To",phone)
	v.Set("From","+12349014223")
	body := message
	v.Set("Body", body)

	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")


	client.Do(req)
}


