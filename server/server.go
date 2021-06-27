package server

import (
	"challenge/controllers"
	"challenge/models"
	"challenge/pkg/config"
	"challenge/repositories"
	"challenge/services"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	AppConfig   *config.Config
	DB          *mongo.Database
	serverReady chan bool
}

func (s Server) Start() {
	log.Info("Init app backbone")
	factRepository := repositories.NewFactRepository(s.DB)
	factService := services.NewFactService(factRepository)
	factController := controllers.NewFactController(factService)

	keyRepository := repositories.NewKeyRepository(s.DB)
	keyService := services.NewKeyService(keyRepository)

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = JSONHTTPErrorHandler

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("/api/v1")

	// Key Authintication Validator
	v1.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			return keyService.IsValid(c.Request().Context(), key)
		},
	}))

	v1.GET("/facts", factController.GetFacts)

	// Run server on Goroutine
	go e.Logger.Fatal(e.StartTLS(fmt.Sprintf(":%d", s.AppConfig.AppPort), s.AppConfig.Certificate.CertFile, s.AppConfig.Certificate.KeyFile))

	if s.serverReady != nil {
		s.serverReady <- true
	}

	// Start Listing on termination Signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}

// JSONHTTPErrorHandler return error as models.Response to unificate response Schema
func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	var errorI interface{} = err
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		errorI = he.Message
	}

	c.JSON(code, models.Response{
		Success: false,
		Error:   errorI,
	})

	c.Logger().Error(err)
}
