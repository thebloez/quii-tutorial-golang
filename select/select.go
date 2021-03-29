package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	//aDuration := measureResponse(a)
	//bDuration := measureResponse(b)
	//
	//if aDuration < bDuration {
	//	return a
	//}
	//
	//return b
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
