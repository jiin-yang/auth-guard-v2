package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jiin-yang/auth-guard-v2/api-gateway/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

type Handler interface {
	CreateEmployee(echo.Context) error
	GetAllEmployees(echo.Context) error
	CreateEmployeeHandler() echo.HandlerFunc
}

type handler struct {
	e *echo.Echo
	c *http.Client
}

func NewHandler(e *echo.Echo, c *http.Client) *handler {
	h := &handler{e: e, c: c}
	h.registerRoutes()
	return h
}
func (h *handler) registerRoutes() {
	h.e.POST("/employees", h.CreateEmployee, middleware.CustomValidToken)
	h.e.GET("/employees", h.GetAllEmployees)
	h.e.GET("/token", h.token)
}

func (h *handler) token(context echo.Context) error {
	url := "https://dev-nuiltl2t5mki2uq1.us.auth0.com/oauth/token"

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	payload := strings.NewReader("{\"client_id\":\"" + clientID + "\",\"client_secret\":\"" + clientSecret + "\",\"audience\":\"https://employee-service/\",\"grant_type\":\"client_credentials\"}")

	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Close = true
	resp, err := h.c.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, token)
}

func (h *handler) CreateEmployee(context echo.Context) error {
	reqBody := CreateEmployeeRequest{}
	if err := context.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if reqBody.Name == "" || reqBody.Department == "" {
		err := errors.New("bad Request - name and department info must")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token := context.Request().Header.Get("authorization")

	req, err := http.NewRequest("POST", os.Getenv("EMPLOYEE_API")+"/api/v1/employees", bytes.NewReader(reqBodyBytes))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Close = true
	resp, err := h.c.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	var respBody CreateEmployeeResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, respBody)
}

func (h *handler) GetAllEmployees(context echo.Context) error {
	req, err := http.NewRequest("GET", os.Getenv("EMPLOYEE_API")+"/api/v1/employees", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Close = true
	resp, err := h.c.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	var employees []Employee
	err = json.NewDecoder(resp.Body).Decode(&employees)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, employees)
}
