package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	newEmail := v1.Group("/email")
	newEmail.POST("", h.SendEmail)
}
