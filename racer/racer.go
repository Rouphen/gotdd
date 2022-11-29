package racer

import (
	"net/http"
	"time"
)

func Racer(slowURL, fastURL string) string {
	startSlow := time.Now()
	http.Get(slowURL)
	slowDuration := time.Since(startSlow)

	startFast := time.Now()
	http.Get(fastURL)
	fastDuration := time.Since(startFast)

	if slowDuration < fastDuration {
		return slowURL
	}

	return fastURL
}
