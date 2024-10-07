package controller

import (
	"flatbuffer-explore/server/entities/object"

	flatbuffers "github.com/google/flatbuffers/go"
)

type HistoryController interface {
	GetBuilder() *flatbuffers.Builder
	MakeHistoryData(data []object.HistoryItem) flatbuffers.UOffsetT
}

func NewHistoryController() HistoryController {
	impl := &historyController{
		builder: flatbuffers.NewBuilder(0),
	}

	return impl
}

type historyController struct {
	builder *flatbuffers.Builder
}

func (h *historyController) MakeHistoryData(data []object.HistoryItem) flatbuffers.UOffsetT {
	Offsetstore := make([]flatbuffers.UOffsetT, len(data))
	
	return 
}
