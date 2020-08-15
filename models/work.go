package models

import "math"

type Work struct {
	id int
	wok chan Request
	pending int
}

func (w *Work) doWork(done chan *Work)  {
	for {
		req := <-w.wok
		req.resp <- math.Sin(float64(req.data))
		done <- w
	}
}
