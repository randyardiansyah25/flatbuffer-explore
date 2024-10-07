package controller

import (
	"flatbuffer-explore/server/entities/fb"
	"flatbuffer-explore/server/entities/object"

	flatbuffers "github.com/google/flatbuffers/go"
)

type HistoryController interface {
	GetBuilder() *flatbuffers.Builder
	BuildHistoryData(data []object.HistoryItem) flatbuffers.UOffsetT
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

func (h *historyController) GetBuilder() *flatbuffers.Builder {
	return h.builder
}

func (h *historyController) BuildHistoryData(data []object.HistoryItem) flatbuffers.UOffsetT {
	OffsetStore := make([]flatbuffers.UOffsetT, len(data))

	for i, v := range data {
		dt := h.builder.CreateString(v.DateTrans)
		desc := h.builder.CreateString(v.Description)

		fb.HistoryItemStart(h.builder)
		fb.HistoryItemAddId(h.builder, v.Id)
		fb.HistoryItemAddDateTrans(h.builder, dt)
		fb.HistoryItemAddDescription(h.builder, desc)
		fb.HistoryItemAddBalance(h.builder, v.Balance)
		fb.HistoryItemAddCreditAmount(h.builder, v.CreditAmount)
		fb.HistoryItemAddDebetAmount(h.builder, v.DebetAmount)
		fbrec := fb.HistoryItemEnd(h.builder)

		fb.ItemUnionWrapperStart(h.builder)
		fb.ItemUnionWrapperAddItem(h.builder, fbrec)
		fb.ItemUnionWrapperAddItemType(h.builder, fb.ItemUnionHistoryItem)
		fbItem := fb.ItemUnionWrapperEnd(h.builder)
		OffsetStore[i] = fbItem
	}

	fb.ResponseArrayStartDataVector(h.builder, len(OffsetStore))
	for _, v := range OffsetStore {
		h.builder.PrependUOffsetT(v)
	}

	dataVec := h.builder.EndVector(len(OffsetStore))
	return dataVec
}
