package main

import (
	"net/http"
	"github.com/sakul987/Lineage-OTA-go/handler"
	"github.com/sakul987/Lineage-OTA-go/internal"
)

func main(){
  //services
  otaService := internal.NewOTAService()

  //handlers
  otaHandler := handler.NewOTAHandler(otaService)

  //register handlers
  http.HandleFunc("/", otaHandler.HandleX))

  http.ListenAndServe(":6666", nil)
}