package main

import (
	"flatbuffer-explore/server/delivery"
	"flatbuffer-explore/server/delivery/router"

	"github.com/kpango/glg"
)

func main() {
	go delivery.PrintoutObserver()
	_ = glg.Log("Server Running")
	router.Start()
}
