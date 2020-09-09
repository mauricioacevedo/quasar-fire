package middlewares

import "github.com/labstack/echo"

func HeadersCheck(next echo.HandlerFunc) echo.HandlerFunc {
	//TODO: we need to check headers.. ,
	return func(c echo.Context) error {
		return next(c)

	}
}
