package controllers

import (
	"challenge/models"
	"challenge/services"
	"context"
	"math"
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
	factsRequest := models.NewFactsRequest()

	if err := c.Bind(factsRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	limit := int64(factsRequest.PerPage)
	skip := int64((factsRequest.Page - 1) * factsRequest.PerPage)

	facts, total, err := f.factService.GetFacts(ctx, &models.Filters{
		Search: factsRequest.Search,
		Skip:   skip,
		Limit:  limit,
	})

	if err != nil {
		c.Logger().Error(err)

		return c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Error:   "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Success:     true,
		Payload:     facts,
		CurrentPage: factsRequest.Page,
		LastPage:    int(math.Ceil(float64(total) / float64(factsRequest.PerPage))),
		PerPage:     factsRequest.PerPage,
		Total:       int(total),
	})
}
