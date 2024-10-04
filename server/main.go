package main

import (
	"flatbuffer-explore/server/delivery/router"

	"github.com/kpango/glg"
)

func main() {
	_ = glg.Log("Server Running")
	router.Start()
}
