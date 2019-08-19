package router

import "github.com/LandvibeDev/gofka/collector/service"

type Handler struct {
	logService service.LogServiceInterface
}

func NewHandler(l service.LogServiceInterface) *Handler {
	return &Handler{
		logService: l,
	}
}
