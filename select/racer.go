package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrRacerTimeout = errors.New("both racers timed out")

func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(time.Second * 1):
		return "", ErrRacerTimeout
	}
}

func ping(url string) chan struct{} {
	c := make(chan struct{})
	go func() {
		http.Get(url)
		// this is more intuitive for me right now
		c <- struct{}{}
	}()
	return c
}
