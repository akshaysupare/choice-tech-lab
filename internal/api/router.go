package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/import", handler.ImportExcel)
	r.GET("/records", handler.GetRecords)
	r.GET("/records/:id", handler.GetRecordByID)
	r.PUT("/records/:id", handler.UpdateRecord)
	r.DELETE("/records/:id", handler.DeleteRecord)

	return r
}
