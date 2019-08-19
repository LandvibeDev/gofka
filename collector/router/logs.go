package router

import (
	"net/http"

	"github.com/LandvibeDev/gofka/collector/kafka/message"
	"github.com/LandvibeDev/gofka/collector/service"
	"github.com/labstack/echo/v4"
)

// e.Post("/logs", WriteLogs)
func (h *Handler) WriteLogs(c echo.Context) error {
	l := new(message.LogMessage)
	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.logService.Send(service.LogTopic, *l)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, l)
}
