package integration_test

import (
	"challenge/models"
	"challenge/pkg/config"
	"challenge/pkg/db"
	"challenge/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type e2eTestSuite struct {
	suite.Suite
	dbConnectionStr string
	appConfig       *config.Config
	db              *mongo.Database
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	var err error
	s.appConfig, err = config.NewWithPath("../env.test.yaml")
	s.Require().NoError(err)

	s.db, err = db.Connect(s.appConfig.DB)
	s.Require().NoError(err)

	serverReady := make(chan bool)
	server := server.Server{
		AppConfig:   s.appConfig,
		DB:          s.db,
		ServerReady: serverReady,
	}

	go server.Start()
	<-serverReady
}

func (s *e2eTestSuite) TestServerRunning() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("https://localhost:%d", s.appConfig.AppPort), nil)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, response.StatusCode)
}

func (s *e2eTestSuite) TestEndToEndFactsNoAuth() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("https://localhost:%d/api/v1/facts", s.appConfig.AppPort), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusBadRequest, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	var res models.Response
	err = json.Unmarshal(byteBody, &res)

	s.Equal(false, res.Success)
	response.Body.Close()
}

func (s *e2eTestSuite) TestEndToEndFactsWrongAuth() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("https://localhost:%d/api/v1/facts", s.appConfig.AppPort), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer xxx")

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusUnauthorized, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	var res models.Response
	err = json.Unmarshal(byteBody, &res)

	s.Equal(false, res.Success)

	response.Body.Close()
}

func (s *e2eTestSuite) TestEndToEndFactsOk() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("https://localhost:%d/api/v1/facts", s.appConfig.AppPort), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer G-KaPdSgVkYp3s6v9y/B?E(H+MbQeThWmZq4t7w!z%C&F)J@NcRfUjXn2r5u8x/A")

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	var res models.Response
	err = json.Unmarshal(byteBody, &res)

	s.Equal(true, res.Success)
	s.Equal(10, res.PerPage)

	response.Body.Close()
}
