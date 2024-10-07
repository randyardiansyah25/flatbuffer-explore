package router

import (
	"flatbuffer-explore/server/delivery/handler"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func Start() error {
	gin.SetMode(gin.ReleaseMode)

	//Discard semua output yang dicatat oleh gin karena print out akan dicetak sesuai kebutuhan programmer
	gin.DefaultWriter = io.Discard

	router := gin.Default() //create router engine by default
	router.Use(gin.Recovery())

	RegisterHandler(router)
	listenerPort := "8800"
	_ = glg.Logf("[HTTP] Listening at : %s", listenerPort)
	return router.Run(":" + listenerPort)

}

func RegisterHandler(router *gin.Engine) {
	router.POST("/request/archive", handler.ArchiveHandler)
	router.POST("/request/history", handler.HistoryHandler)
	router.GET("/request/archive/:id", handler.ArchiveItemHandler)
}
