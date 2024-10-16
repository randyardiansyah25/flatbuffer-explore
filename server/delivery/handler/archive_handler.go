package handler

import (
	"flatbuffer-explore/server/delivery/handler/httpio"
	"flatbuffer-explore/server/entities/controller"
	"flatbuffer-explore/server/entities/object"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
)

func ArchiveHandler(ctx *gin.Context) {
	httpio := httpio.NewRequestIO(ctx)
	data, er := io.ReadAll(ctx.Request.Body)
	if er != nil {
		httpio.ResponseString(http.StatusBadGateway, er.Error())
		return
	}

	ctl := controller.NewRequestController()
	req := object.Request{}
	ctl.Read(data, &req)
	httpio.PrintRecv(req)

	items := getArchive()
	status := object.Status{
		Code:    "00",
		Message: "success",
	}

	//buat object utk log
	response := object.Response[[]object.ArchiveItem]{
		Response: status,
		Data:     items,
	}

	archiveCtl := controller.NewArchiveController()
	fbData := archiveCtl.BuildArchiveData(items)

	rctl := controller.NewResponseController(archiveCtl.GetBuilder())
	respData := rctl.BuildResponseArray(status.Code, status.Message, fbData)

	httpio.ResponseData(http.StatusOK, respData, response, "application/octet-stream")
}

func ArchiveItemHandler(ctx *gin.Context) {
	httpio := httpio.NewRequestIO(ctx)
	sId := ctx.Param("id")
	nId, _ := strconv.Atoi(sId)
	archiveId := int64(nId)
	httpio.Recv()

	item := getArchiveItem(archiveId)
	logObj := object.Response[object.ArchiveItem]{
		Response: object.Status{
			Code:    "00",
			Message: "Archive Item Success",
		},
		Data: item,
	}

	ctl := controller.NewArchiveController()
	fbItem := ctl.BuildArchiveItem(item)

	respctl := controller.NewResponseController(ctl.GetBuilder())
	respData := respctl.BuildResponseObject(logObj.Response.Code, logObj.Response.Message, fbItem)

	httpio.ResponseData(http.StatusOK, respData, logObj, "application/octet-stream")
}

func getArchiveItem(id int64) object.ArchiveItem {
	return object.ArchiveItem{
		Id:                id,
		DateTrans:         faker.Timestamp(),
		TransactionAmount: float64(rand.Intn(1000000) + 100),
		Description:       faker.Sentence(),
		Status:            rand.Intn(2),
	}
}

func getArchive() (items []object.ArchiveItem) {
	for i := 0; i < 10; i++ {
		item := object.ArchiveItem{
			Id:                int64(rand.Intn(1000000)),
			DateTrans:         faker.Timestamp(),
			TransactionAmount: float64(rand.Intn(140) + 10),
			Description:       faker.Sentence(),
			Status:            rand.Intn(2),
		}

		items = append(items, item)
	}
	return
}
