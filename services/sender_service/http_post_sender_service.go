package sender_service

import (
	"bytes"
	"net/http"
)

type HttpPostSenderService struct {
	ISenderService
}

func (HttpPostSenderService) Send(uri string, data string) bool {
	requestData := []byte(data)
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(requestData))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false
	}
	response.Body.Close()
	if response.StatusCode != 200 {
		return false
	}
	return true
}
