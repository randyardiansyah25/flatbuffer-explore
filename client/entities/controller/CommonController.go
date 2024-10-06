package controller

import (
	"flatbuffer-explore/server/entities/fb"
	"flatbuffer-explore/server/entities/object"

	flatbuffers "github.com/google/flatbuffers/go"
)

type RequestController interface {
	MakeArchiveRequest(req object.Request) []byte
}

func NewRequestController() RequestController {
	impl := &requestControllerImpl{
		builder: flatbuffers.NewBuilder(0),
	}

	return impl
}

type requestControllerImpl struct {
	builder *flatbuffers.Builder
}

func (rc *requestControllerImpl) MakeArchiveRequest(req object.Request) []byte {
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
