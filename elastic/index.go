package elastic

import (
	"fmt"
	"time"
)

type IndexManager struct {
	BaseUrl string
}

func (m *IndexManager) BuildUrl() string {
	return m.BaseUrl + getAnimeIndexName()
}

func (m *IndexManager) BuildSettings() string {
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
