package main

import (
	"context"
	"errors"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"time"

	"log"
	"net/http"
	"strings"
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

		token := c.Request().Header.Get("authorization")

		err := verifyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

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

func verifyToken(tokenString string) error {

	if tokenString == "" {
		return errors.New("bearer Token is must")
	}

	jwksURL := "https://dev-nuiltl2t5mki2uq1.us.auth0.com/.well-known/jwks.json"

	ctx, cancel := context.WithCancel(context.Background())

	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
		return err
	}

	splitToken := strings.Split(tokenString, "Bearer ")
	jwtB64 := splitToken[1]

	token, err := jwt.Parse(jwtB64, jwks.Keyfunc)
	if err != nil {
		log.Fatalf("Failed to parse the JWT.\nError: %s", err.Error())
		return err
	}

	if !token.Valid {
		log.Fatalf("The token is not valid.")
	}
	log.Println("The token is valid.")

	cancel()
	jwks.EndBackground()

	return nil
}
