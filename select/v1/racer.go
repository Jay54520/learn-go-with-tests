package racer

import (
	"time"
	"net/http"
)

func Racer(a, b string) (winner string) {

	aDuration := getDurationOfURL(a)
	bDuration := getDurationOfURL(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

// getDurationOfURL takes an URL and return the duration of requesting it
func getDurationOfURL(URL string) (duration time.Duration) {
	start := time.Now()
	http.Get(URL)
	duration = time.Since(start)
	return
}
