package controller

import (
	"flatbuffer-explore/client/entities/object"
	"flatbuffer-explore/server/entities/fb"

	flatbuffers "github.com/google/flatbuffers/go"
)

type HistoryRequestController interface {
	MakeRequest(obj object.Request) []byte
	ReadHistoryResponse(buf []byte) object.Response[[]object.HistoryItem]
	Reset()
}

func NewHistoryRequestController() HistoryRequestController {
	impl := &historyControllerImpl{
		builder: flatbuffers.NewBuilder(0),
	}

	return impl
}

type historyControllerImpl struct {
	builder *flatbuffers.Builder
}

func (h *historyControllerImpl) MakeRequest(obj object.Request) []byte {
	dtStart := h.builder.CreateString(obj.DateStart)
	dtEnd := h.builder.CreateString(obj.DateEnd)

	fb.RequestStart(h.builder)
	fb.RequestAddDateStart(h.builder, dtStart)
	fb.RequestAddDateEnd(h.builder, dtEnd)
	fbreq := fb.RequestEnd(h.builder)
	h.builder.Finish(fbreq)
	data := h.builder.FinishedBytes()
	return data
}

func (h *historyControllerImpl) ReadHistoryResponse(buf []byte) (resp object.Response[[]object.HistoryItem]) {
	fbdata := fb.GetRootAsResponseArray(buf, 0)
	fbStatus := fb.Status{}
	fbdata.Response(&fbStatus)
	resp.Response.Code = string(fbStatus.Code())
	resp.Response.Message = string(fbStatus.Message())

	items := make([]object.HistoryItem, fbdata.DataLength())
	itemWrapper := &fb.ItemUnionWrapper{}
	for i := 0; i < fbdata.DataLength(); i++ {
		if fbdata.Data(itemWrapper, i) {
			unionTable := new(flatbuffers.Table)
			if itemWrapper.Item(unionTable) {
				fbitem := &fb.HistoryItem{}
				fbitem.Init(unionTable.Bytes, unionTable.Pos)
				item := object.HistoryItem{
					Id:           fbitem.Id(),
					DateTrans:    string(fbitem.DateTrans()),
					Description:  string(fbitem.Description()),
					DebetAmount:  fbitem.DebetAmount(),
					CreditAmount: fbitem.CreditAmount(),
					Balance:      fbitem.Balance(),
				}
				items[i] = item
			}
		}
	}
	resp.Data = items
	return
}

func (h *historyControllerImpl) Reset() {
	h.builder.Reset()
}
