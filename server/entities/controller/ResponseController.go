package controller

import (
	"flatbuffer-explore/server/entities/fb"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ResponseController interface {
	BuildResponseArray(code string, message string, vec flatbuffers.UOffsetT) []byte
}

func NewResponseController(builder *flatbuffers.Builder) ResponseController {
	impl := &responseControllerImpl{
		builder: builder,
	}

	return impl
}

type responseControllerImpl struct {
	builder *flatbuffers.Builder
}

func (rc *responseControllerImpl) makeStatus(code string, message string) flatbuffers.UOffsetT {
	sCode := rc.builder.CreateString(code)
	sMsg := rc.builder.CreateString(message)

	fb.StatusStart(rc.builder)
	fb.StatusAddCode(rc.builder, sCode)
	fb.StatusAddMessage(rc.builder, sMsg)
	status := fb.StatusEnd(rc.builder)
	return status
}

func (rc *responseControllerImpl) BuildResponseArray(code string, message string, vec flatbuffers.UOffsetT) []byte {

	// Prepare dulu untuk status,
	status := rc.makeStatus(code, message)

	fb.ResponseArrayStart(rc.builder)
	fb.ResponseArrayAddResponse(rc.builder, status)
	fb.ResponseArrayAddData(rc.builder, vec)
	resp := fb.ResponseArrayEnd(rc.builder)

	rc.builder.Finish(resp)
	return rc.builder.FinishedBytes()
}
