package controller

import (
	"flatbuffer-explore/server/entities/fb"
	"flatbuffer-explore/server/entities/object"

	flatbuffers "github.com/google/flatbuffers/go"
)

type RequestController interface {
	Read(buf []byte, reqobj *object.Request)
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

func (rc *requestControllerImpl) Read(buf []byte, reqobj *object.Request) {
	reqfb := fb.GetRootAsRequest(buf, 0)

	reqobj.DateStart = string(reqfb.DateStart())
	reqobj.DateEnd = string(reqfb.DateEnd())
}
