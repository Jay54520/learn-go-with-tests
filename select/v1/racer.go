package racer

import (
	"time"
	"net/http"
)

func Racer(a, b string) (winner string) {

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}


func measureResponseTime(URL string) (duration time.Duration) {
	start := time.Now()
	http.Get(URL)
	duration = time.Since(start)
	return
}
