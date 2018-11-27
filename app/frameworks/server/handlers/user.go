package handlers

import "github.com/labstack/echo"

func GetUserProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(204)
	}
}

func PostUserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(204)
	}
}
