package controller

import (
	"flatbuffer-explore/client/entities/fb"
	"flatbuffer-explore/client/entities/object"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ArchiveRequestController interface {
	MakeArchiveRequest(req object.Request) []byte
	ReadArchiveResponse(buf []byte) object.Response[[]object.ArchiveItem]
	ReadArchiveItemResponse(buf []byte) object.Response[object.ArchiveItem]
	Reset()
}

func NewArchiveRequestController() ArchiveRequestController {
	impl := &archiveRrequestControllerImpl{
		builder: flatbuffers.NewBuilder(0),
	}

	return impl
}

type archiveRrequestControllerImpl struct {
	builder *flatbuffers.Builder
}

func (rc *archiveRrequestControllerImpl) Reset() {
	rc.builder.Reset()
}

func (rc *archiveRrequestControllerImpl) MakeArchiveRequest(req object.Request) []byte {
	sDt := rc.builder.CreateString(req.DateStart)
	eDt := rc.builder.CreateString(req.DateEnd)

	fb.RequestStart(rc.builder)
	fb.RequestAddDateStart(rc.builder, sDt)
	fb.RequestAddDateEnd(rc.builder, eDt)
	fbReq := fb.RequestEnd(rc.builder)
	rc.builder.Finish(fbReq)
	data := rc.builder.FinishedBytes()
	return data
}

func (rc *archiveRrequestControllerImpl) ReadArchiveResponse(buf []byte) object.Response[[]object.ArchiveItem] {
	resp := object.Response[[]object.ArchiveItem]{}

	fbRootResp := fb.GetRootAsResponseArray(buf, 0)

	fbStatus := new(fb.Status)
	fbRootResp.Response(fbStatus)
	resp.Response.Code = string(fbStatus.Code())
	resp.Response.Message = string(fbStatus.Message())

	items := make([]object.ArchiveItem, fbRootResp.DataLength())
	itemWrapper := &fb.ItemUnionWrapper{}
	for i := 0; i < fbRootResp.DataLength(); i++ {
		if fbRootResp.Data(itemWrapper, i) {
			unionTable := new(flatbuffers.Table)
			if itemWrapper.Item(unionTable) {
				fbItem := &fb.ArchiveItem{}
				fbItem.Init(unionTable.Bytes, unionTable.Pos)
				item := object.ArchiveItem{
					Id:                fbItem.Id(),
					DateTrans:         string(fbItem.DateTrans()),
					TransactionAmount: fbItem.TransactionAmount(),
					Description:       string(fbItem.Description()),
					Status:            fbItem.Status(),
				}
				items[i] = item
			}
		}
	}
	resp.Data = items
	return resp
}

func (rc *archiveRrequestControllerImpl) ReadArchiveItemResponse(buf []byte) (resp object.Response[object.ArchiveItem]) {
	fbdata := fb.GetRootAsResponseObject(buf, 0)
	fbstatus := &fb.Status{}
	fbdata.Response(fbstatus)

	resp.Response.Code = string(fbstatus.Code())
	resp.Response.Message = string(fbstatus.Message())

	itemWrapper := &fb.ItemUnionWrapper{}
	fbdata.Data(itemWrapper)
	unionTable := new(flatbuffers.Table)
	if itemWrapper.Item(unionTable) {
		item := fb.ArchiveItem{}
		item.Init(unionTable.Bytes, unionTable.Pos)
		resp.Data.Id = item.Id()
		resp.Data.DateTrans = string(item.DateTrans())
		resp.Data.Description = string(item.Description())
		resp.Data.TransactionAmount = item.TransactionAmount()
		resp.Data.Status = item.Status()
	}

	return
}
