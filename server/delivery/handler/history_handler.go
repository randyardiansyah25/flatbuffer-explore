package handler

import (
	"flatbuffer-explore/server/delivery/handler/httpio"
	"flatbuffer-explore/server/entities/controller"
	"flatbuffer-explore/server/entities/object"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
)

func HistoryHandler(ctx *gin.Context) {
	httpio := httpio.NewRequestIO(ctx)
	buf, er := io.ReadAll(ctx.Request.Body)
	if er != nil {
		httpio.ResponseString(http.StatusBadRequest, er.Error())
		return
	}

	ctl := controller.NewRequestController()
	reqObj := object.Request{}
	ctl.Read(buf, &reqObj)
	httpio.PrintRecv(reqObj)

	historyData := getHistory()
	resp := object.Response[[]object.HistoryItem]{
		Response: object.Status{
			Code:    "00",
			Message: "History Success",
		},
		Data: historyData,
	}

	historyController := controller.NewHistoryController()
	historyBuilder := historyController.GetBuilder()
	dataVector := historyController.BuildHistoryData(historyData)

	respCtl := controller.NewResponseController(historyBuilder)
	respdata := respCtl.BuildResponseArray(resp.Response.Code, resp.Response.Message, dataVector)
	httpio.ResponseData(http.StatusOK, respdata, resp, "application/octet-stream")
}

func getHistory() (record []object.HistoryItem) {
	for i := 0; i < 10; i++ {
		item := object.HistoryItem{
			Id:           int64(rand.Intn(1000000)),
			DateTrans:    faker.Timestamp(),
			DebetAmount:  float64(rand.Intn(1000000)) * 2.5,
			CreditAmount: float64(rand.Intn(1000000)) * 2.5,
			Description:  faker.Sentence(),
			Balance:      float64(rand.Intn(1000000)) * 2.5,
		}
		record = append(record, item)
	}
	return
}
