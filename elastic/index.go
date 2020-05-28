package elastic

import (
	"fmt"
	"time"
)

func AnimeIndexName() string {
	now := time.Now()
	currentTime := fmt.Sprintf("%d-%02d-%02d", now.Day(), now.Month(), now.Year())
	return "anime-" + currentTime
}
