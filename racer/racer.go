package racer

import (
	"net/http"
	"time"
)

func Racer(slowURL, fastURL string) string {
	slowDuration := meassureResponsetime(slowURL)
	fastDuration := meassureResponsetime(fastURL)

	if slowDuration < fastDuration {
		return slowURL
	}

	return fastURL
}

func meassureResponsetime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}
