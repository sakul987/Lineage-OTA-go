package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakul987/Lineage-OTA-go/internal"
)

type OTAHandler struct{
  otaService internal.OTAService
}

func NewOTAHandler(otaService internal.OTAService) OTAHandler{
  return OTAHandler{otaService: otaService}
}


func (o OTAHandler) Register(router *gin.Engine){
  router.GET("/", o.HandleX)
}

func (o OTAHandler) HandleX(ctx *gin.Context){
  err := errors.New("test")
  if err != nil{

  }
  ctx.JSON(http.StatusBadRequest, nil)
}