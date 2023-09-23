package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sakul987/Lineage-OTA-go/handler"
	"github.com/sakul987/Lineage-OTA-go/internal"
)

func main(){
  //router config
  router := gin.Default()

  //services
  otaService := internal.NewOTAService()

  //handlers
  otaHandler := handler.NewOTAHandler(otaService)

  //register handlers
  otaHandler.Register(router)

  router.Run(":6666")
}