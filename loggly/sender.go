package loggly

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type LogglySender struct {
	Client http.Client
	Uri    string
}

func NewSender(token string) *LogglySender {

	interped := fmt.Sprintf("https://logs-01.loggly.com/inputs/%s", token)

	return &LogglySender{
		Uri:    interped,
		Client: http.Client{},
	}
}

func SendEvent(entry string, sender LogglySender) {
	req, err := http.NewRequest("POST", sender.Uri, ioutil.NopCloser(strings.NewReader(entry)))
	if err != nil {
		log.Fatalf("Could not set up http post: %v", err)
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := sender.Client.Do(req)
	if err != nil {
		log.Println("Unable to send data: %v", err)
		log.Println(entry)
	}
	if resp != nil {
	}
}
