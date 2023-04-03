package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

const PORT = ":8081"

func main() {
	e := echo.New()

	e.GET("/api/v1/employees", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccept, echo.MIMEApplicationJSONCharsetUTF8)

		return c.JSON(http.StatusOK, employees)
	})
	e.POST("/api/v1/employees", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccept, echo.MIMEApplicationJSONCharsetUTF8)

		requestBody := CreateEmployeeRequest{}
		if err := c.Bind(&requestBody); err != nil {
			err := errors.New("bad Request employee API - name and department info must")
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		employee := Employee{
			ID:         len(employees) + 1,
			Name:       requestBody.Name,
			Department: requestBody.Department,
		}

		employees = append(employees, employee)

		return c.JSON(http.StatusOK, employee)
	})

	if err := e.Start(PORT); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
