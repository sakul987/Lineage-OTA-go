package handler

import (
	"errors"
	"net/http"
	"github.com/sakul987/Lineage-OTA-go/internal"
)

type OTAHandler struct{
  otaService internal.OTAService
}

func NewOTAHandler(otaService internal.OTAService) OTAHandler{
  return OTAHandler{otaService: otaService}
}

func (o OTAHandler) HandleX(w http.ResponseWriter, r *http.Request) {
	err := errors.New("test")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}