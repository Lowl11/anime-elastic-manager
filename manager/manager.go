package manager

import (
	"bytes"
	"elastic-manager/elastic"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ElasticManager struct {
	Url string
}

// Возвращает все индексы находящиеся в эластике
func (m *ElasticManager) GetIndices(w *http.ResponseWriter) []Index {
	fullUrl := m.Url + "_cat/indices"
	data := ``
	response, err := m.makeRequest(w, fullUrl, data, http.MethodGet)
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(strings.TrimSpace(response), "\n", -1)
	indices := make([]Index, 0, len(lines))
	for _, line := range lines {
		words := strings.Split(line, ` `)

		index := &Index{
			Status:         words[0],
			Name:           words[2],
			HashCode:       words[3],
			Shards:         words[4],
			DocumentsCount: words[6],
			Size:           words[8],
		}
		indices = append(indices, *index)
	}
	return indices
}

// Создание индекса
func (m *ElasticManager) CreateIndex(w *http.ResponseWriter) JsonResult {
	fullUrl := m.Url + elastic.AnimeIndexName()
	data := ``
	response, err := m.makeRequest(w, fullUrl, data, http.MethodPut)
	if err != nil {
		return getSimpleResult(false, err.Error())
	}
	return getSimpleResult(true, response)
}

// удаление индекса
func (m *ElasticManager) DeleteIndex(w *http.ResponseWriter) JsonResult {
	fullUrl := m.Url + elastic.AnimeIndexName()
	data := ``
	response, err := m.makeRequest(w, fullUrl, data, http.MethodDelete)
	if err != nil {
		return getSimpleResult(false, err.Error())
	}
	return getSimpleResult(true, response)
}

// кидает реквест
func (m *ElasticManager) makeRequest(w *http.ResponseWriter, url string, data string, requestType string) (string, error) {
	transport := getTransport()

	// хрен его знает в чем разница но нужно понять зачем еще и клиент создавать
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	// raw реквеста
	body := bytes.NewBufferString(data) // видимо форматируем дату (кароч полезная штука)

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

// задает контент отправляемому запросу (куда-либо) - json
func turnRequestToJson(request *http.Request, dataLength string) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", dataLength)
}

func getSimpleResult(success bool, message string) JsonResult {
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
