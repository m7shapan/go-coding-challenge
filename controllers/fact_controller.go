package controllers

import (
	"challenge/services"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FactController interface {
	GetFacts(c echo.Context) error
}

type factController struct {
	factService services.FactService
}

func NewFactController(s services.FactService) factController {
	return factController{
		factService: s,
	}
}

func (f factController) GetFacts(c echo.Context) error {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	facts, err := f.factService.GetFacts(ctx)
	if err != nil {
		c.Logger().Error(err)

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": true,
			"error":   "unable to process your Action",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"payload": facts,
	})
}
