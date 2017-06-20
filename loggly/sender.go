package loggly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Response string `json:"response"`
}

func GenerateUri(token string, tag string) string {

	uri := ""
	switch tag {
	case "":
		uri = fmt.Sprintf("https://logs-01.loggly.com/inputs/%s", token)
	default:
		uri = fmt.Sprintf("https://logs-01.loggly.com/inputs/%s/tag/%s/", token, tag)
	}

	log.Println("Will send data to", uri)

	return uri
}

func SendEvent(entry string, uri string) {
	resp, err := http.Post(uri, "application/x-www-form-urlencoded", ioutil.NopCloser(strings.NewReader(entry)))
	if err != nil {
		log.Println("Unable to send data:", err)
		log.Println(entry)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to get a response:", err)
		return
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Invalid response:", err)
		return
	}
	if response.Response != "ok" {
		log.Println("Response was:", response.Response)
	}

}
