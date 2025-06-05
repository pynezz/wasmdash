package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/wasmdash/pkg/ui"
	"github.com/pynezz/wasmdash/pkg/ui/pages"
)

func homeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

func aboutHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.About(c.Path()))
}

func serviceWorker(c echo.Context) error {
	return Render(c, http.StatusOK, ui.ServiceWorker())
}
