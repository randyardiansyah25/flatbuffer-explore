package controller

import (
	"flatbuffer-explore/client/entities/object"
	"flatbuffer-explore/server/entities/fb"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ArchiveRequestController interface {
	MakeArchiveRequest(req object.Request) []byte
	ReadArchiveResponse(buf []byte) object.Response[[]object.ArchiveItem]
	Reset()
}

func NewRequestController() ArchiveRequestController {
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
					Status:            int(fbItem.Status()),
				}
				items[i] = item
			}
		}
	}
	resp.Data = items
	return resp
}
