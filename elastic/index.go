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

func getAnimeIndexName() string {
	now := time.Now()
	currentTime := fmt.Sprintf("%d-%02d-%02d", now.Day(), now.Month(), now.Year())
	return "anime-" + currentTime
}
