package main

import (
	"bytes"
	"encoding/json"
	"flatbuffer-explore/client/entities/controller"
	"flatbuffer-explore/client/entities/object"
	"fmt"
	"net/http"

	"github.com/kpango/glg"
	net "github.com/randyardiansyah25/libpkg/net/http"
)

func main() {
	_ = glg.Log("hello client")

	client := net.NewSimpleClient("POST", "http://localhost:8800/request", 30)
	client.SetContentType("application/octet-stream")
	ctl := controller.NewRequestController()
	req := object.Request{
		DateStart: "2024-10-04",
		DateEnd:   "2024-10-04",
	}

	buf := ctl.MakeArchiveRequest(req)
	res := client.Do(bytes.NewBuffer(buf))

	if res.StatusCode() != http.StatusOK {
		fmt.Printf("code:%d, message : %s\n", res.StatusCode(), res.Message())
	} else {
		ctl.Reset()
		respData := ctl.ReadArchiveResponse(res.Message())
		// fmt.Println(res.Message())
		// fmt.Println(respData)
		j, _ := json.MarshalIndent(respData,"", "  ")
		fmt.Println(string(j))
	}
}
