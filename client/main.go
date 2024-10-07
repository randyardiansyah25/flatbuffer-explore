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

	// GetArchive()
	// GetHistory()
	GetArchiveItem()
}

func GetArchive() {
	client := net.NewSimpleClient("POST", "http://localhost:8800/request/archive", 30)
	client.SetContentType("application/octet-stream")
	ctl := controller.NewArchiveRequestController()
	req := object.Request{
		DateStart: "2024-10-04",
		DateEnd:   "2024-10-04",
	}

	buf := ctl.MakeArchiveRequest(req)
	res := client.Do(bytes.NewBuffer(buf))

	if res.StatusCode() != http.StatusOK {
		fmt.Printf("code:%d, message : %s\n", res.StatusCode(), res.Message())
	} else {
		respData := ctl.ReadArchiveResponse(res.Message())
		j, _ := json.MarshalIndent(respData, "", "  ")
		fmt.Println(string(j))
	}
}

func GetArchiveItem() {
	client := net.NewSimpleClient("GET", "http://localhost:8800/request/archive/110029", 30)
	client.SetContentType("application/octet-stream")
	ctl := controller.NewArchiveRequestController()

	buf := make([]byte, 0)
	res := client.Do(bytes.NewBuffer(buf))

	if res.StatusCode() != http.StatusOK {
		fmt.Printf("code:%d, message : %s\n", res.StatusCode(), res.Message())
	} else {
		respData := ctl.ReadArchiveItemResponse(res.Message())
		j, _ := json.MarshalIndent(respData, "", "  ")
		fmt.Println(string(j))
	}
}

func GetHistory() {
	client := net.NewSimpleClient("POST", "http://localhost:8800/request/history", 30)
	client.SetContentType("application/octet-stream")
	ctl := controller.NewHistoryRequestController()
	req := object.Request{
		DateStart: "2024-10-04",
		DateEnd:   "2024-10-04",
	}

	buf := ctl.MakeRequest(req)
	res := client.Do(bytes.NewBuffer(buf))

	if res.StatusCode() != http.StatusOK {
		fmt.Printf("code:%d, message : %s\n", res.StatusCode(), res.Message())
	} else {
		respData := ctl.ReadHistoryResponse(res.Message())
		// fmt.Println(res.Message())
		// fmt.Println(respData)
		j, _ := json.MarshalIndent(respData, "", "  ")
		fmt.Println(string(j))
	}
}
