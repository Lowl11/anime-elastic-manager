package elastic

import (
	"elastic-manager/common"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Index common.Index

// Возвращает все индексы находящиеся в эластике
func GetIndices(w *http.ResponseWriter, baseUrl string) []Index {
	fullUrl := baseUrl + "_cat/indices"
	data := ``
	response, err := common.MakeRequest(w, fullUrl, data, http.MethodGet)
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
func CreateIndex(w *http.ResponseWriter, baseUrl string) common.JsonResult {
	result := DeleteIndex(w, baseUrl)
	if result.Status != "success" {
		panic("Удаление индекса прошло безуспешно")
	}

	fullUrl := buildUrl(baseUrl)
	data := buildSettings()

	response, err := common.MakeRequest(w, fullUrl, data, http.MethodPut)
	if err != nil {
		return common.GetSimpleResult(false, err.Error())
	}

	return common.GetSimpleResult(true, response)
}

// удаление индекса
func DeleteIndex(w *http.ResponseWriter, baseUrl string) common.JsonResult {
	fullUrl := buildUrl(baseUrl)
	data := ``

	response, err := common.MakeRequest(w, fullUrl, data, http.MethodDelete)
	if err != nil {
		return common.GetSimpleResult(false, err.Error())
	}

	return common.GetSimpleResult(true, response)
}

func buildUrl(baseUrl string) string {
	return baseUrl + getAnimeIndexName()
}

func buildSettings() string {
	settings :=
		`{
			"settings": {
				"index": {
					"number_of_shards": 3,
					"number_of_replicas": 2
				},
				"analysis": {
					"analyzer": {
						"anime_analyzer": {
							"type": "custom",
							"tokenizer": "standard",
							"char_filter": [
								"html_strip"
							],
							"filter": ["lowercase", "asciifolding", "english_stop", "russian_stop"]
						}
					},
					"filter": {
						"english_stop": {
							"type": "stop",
							"stopwords": "_english_"
						},
						"russian_stop": {
							"type": "stop",
							"stopwords": "_russian_"
						}
					}
				}
			}
		}`
	return settings
}

func getAnimeIndexName() string {
	now := time.Now()
	currentTime := fmt.Sprintf("%d.%02d.%02d", now.Day(), now.Month(), now.Year())
	return "anime-" + currentTime
}
