package router

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	log := v1.Group("/logs")
	log.POST("", h.WriteLogs)
}
