package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	c "github.com/quasar-fire/controllers"
	"github.com/quasar-fire/middlewares"
	"github.com/quasar-fire/utils"
)

var (
	port = fmt.Sprintf(":%s", utils.GetEnv("PORT", "3000"))
)

func main() {

	e := echo.New()

	locationGroup := e.Group("", middlewares.HeadersCheck)
	locationGroup.POST("/topsecret", c.PostTopSecret)
	e.POST("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Never tell me the odds!")
	})
	e.Logger.Fatal(e.Start(port))
}
