package controller

import (
	"flatbuffer-explore/server/entities/fb"
	"flatbuffer-explore/server/entities/object"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ArchiveController interface {
	BuildArchiveData(data []object.ArchiveItem) flatbuffers.UOffsetT
	BuildArchiveItem(data object.ArchiveItem) flatbuffers.UOffsetT
	GetBuilder() *flatbuffers.Builder
}

func NewArchiveController() ArchiveController {
	impl := &archiveController{
		builder: flatbuffers.NewBuilder(0),
	}
	return impl
}

type archiveController struct {
	builder *flatbuffers.Builder
}

func (a *archiveController) GetBuilder() *flatbuffers.Builder {
	return a.builder
}

func (a *archiveController) BuildArchiveData(data []object.ArchiveItem) flatbuffers.UOffsetT {
	OffsetStore := make([]flatbuffers.UOffsetT, 0)
	for _, v := range data {
		// Create String dulu sebelum memanggil start
		dt := a.builder.CreateString(v.DateTrans)
		desc := a.builder.CreateString(v.Description)

		// Buat Item dari Archive (ArchiveItem Table)
		fb.ArchiveItemStart(a.builder)
		fb.ArchiveItemAddId(a.builder, v.Id)
		fb.ArchiveItemAddDateTrans(a.builder, dt)
		fb.ArchiveItemAddDescription(a.builder, desc)
		fb.ArchiveItemAddTransactionAmount(a.builder, v.TransactionAmount)
		fb.ArchiveItemAddStatus(a.builder, int32(v.Status))
		item := fb.ArchiveItemEnd(a.builder)

		// Buat dulu offset Item Union Wrapper (ItemUnionWrapper Table)
		fb.ItemUnionWrapperStart(a.builder)
		fb.ItemUnionWrapperAddItem(a.builder, item)
		fb.ItemUnionWrapperAddItemType(a.builder, fb.ItemUnionArchiveItem)
		itemWrapper := fb.ItemUnionWrapperEnd(a.builder)
		OffsetStore = append(OffsetStore, itemWrapper)
	}

	fb.ResponseArrayStartDataVector(a.builder, len(OffsetStore))
	for _, v := range OffsetStore {
		// menambahkan kedalam vector (array)
		a.builder.PrependUOffsetT(v)
	}

	dataVec := a.builder.EndVector(len(OffsetStore))
	return dataVec
}

func (a *archiveController) BuildArchiveItem(data object.ArchiveItem) flatbuffers.UOffsetT {
	dt := a.builder.CreateString(data.DateTrans)
	desc := a.builder.CreateString(data.Description)

	fb.ArchiveItemStart(a.builder)
	fb.ArchiveItemAddId(a.builder, data.Id)
	fb.ArchiveItemAddDateTrans(a.builder, dt)
	fb.ArchiveItemAddDescription(a.builder, desc)
	fb.ArchiveItemAddTransactionAmount(a.builder, data.TransactionAmount)
	fb.ArchiveItemAddStatus(a.builder, int32(data.Status))
	fbitem := fb.ArchiveItemEnd(a.builder)

	fb.ItemUnionWrapperStart(a.builder)
	fb.ItemUnionWrapperAddItem(a.builder, fbitem)
	fb.ItemUnionWrapperAddItemType(a.builder, fb.ItemUnionArchiveItem)
	fbwraperitem := fb.ItemUnionWrapperEnd(a.builder)
	return fbwraperitem
}
