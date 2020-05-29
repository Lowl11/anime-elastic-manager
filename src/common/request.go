package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

// кидает реквест
func MakeRequest(w *http.ResponseWriter, url, data, requestType string) (string, error) {
	log.Printf("Sending request by URL: %s [%s]", url, requestType)
	transport := getTransport()

	// хрен его знает в чем разница но нужно понять зачем еще и клиент создавать
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	// видимо форматируем raw (кароч полезная штука)
	body := bytes.NewBufferString(data)

	// задаем метод запроса, урл и отправляемые данные
	request, _ := http.NewRequest(requestType, url, body)
	turnRequestToJson(request, strconv.Itoa(len(data)))

	response, err := client.Do(request) // непорседственно сам запрос

	if err != nil {
		fmt.Println("Request error:", err.Error())
		return "", err
	}
	defer response.Body.Close() // обязательно нужно закрывать

	// если все ок то вытаскиваем ответ
	responseBody, err := ioutil.ReadAll(response.Body)

	return string(responseBody), nil
}

func GetSimpleResult(success bool, message string) JsonResult {
	var status string
	var successMessage string
	var errorMessage string
	if success {
		status = "success"
		successMessage = message
	} else {
		status = "error"
		errorMessage = message
	}
	result := &JsonResult{
		Status:       status,
		Message:      successMessage,
		ErrorMessage: errorMessage,
	}
	return *result
}

// Создаем объект транспорта чтобы у нас были таймауты
func getTransport() *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 39 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// задает контент отправляемому запросу (куда-либо) - json
func turnRequestToJson(request *http.Request, dataLength string) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", dataLength)
}
