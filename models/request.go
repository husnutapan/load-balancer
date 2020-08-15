package models

import (
	"math/rand"
	"time"
)

type Request struct {
	data int
	resp chan float64
}

func CreateAndRequest(req chan Request) {
	resp := make(chan float64)

	for {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Millisecond))))
		req <- Request{int(rand.Int31n(90)), resp}
		<-resp
	}
}
