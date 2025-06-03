package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/wasmdash/pkg/ui/pages"
)

func homeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}
