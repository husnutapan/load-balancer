package models

import (
	"container/heap"
	"fmt"
)

type Pool []*Work



type Balancer struct {
	pool Pool
	done chan *Work
}

func InitBalancer() *Balancer {
	done := make(chan *Work, Requester)
	b := &Balancer{make(Pool, 0, Worker), done}
	for i := 0; i < Worker; i++ {
		w := &Work{wok: make(chan Request, Requester)}
		// put them in heap
		heap.Push(&b.pool, w)
		go w.doWork(b.done)
	}
	return b
}

func (b *Balancer) Balance(req chan Request) {
	for {
		select {
		case request := <-req:
			b.dispatch(request)
		case w := <-b.done:
			b.completed(w)
		}
		b.print()
	}
}

func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Work)
	w.wok <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Work) {
	w.pending--
	heap.Remove(&b.pool, w.id)
	heap.Push(&b.pool, w)
}

func (b *Balancer) print() {
	sum := 0
	sumsq := 0
	// Print pending stats for each worker
	for _, w := range b.pool {
		fmt.Printf("%d ", w.pending)
		sum += w.pending
		sumsq += w.pending * w.pending
	}
	// Print avg for worker pool
	avg := float64(sum) / float64(len(b.pool))
	variance := float64(sumsq)/float64(len(b.pool)) - avg*avg
	fmt.Printf(" %.2f %.2f\n", avg, variance)
}
