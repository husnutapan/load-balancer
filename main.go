package main

import (
	"golang-lb/models"
)

func main() {
	work := make(chan models.Request)
	for i := 0; i < models.Requester; i++ {
		go models.CreateAndRequest(work)
	}
	models.InitBalancer().Balance(work)
}
