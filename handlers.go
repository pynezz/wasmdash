package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pynezz/wasmdash/pkg/server/handlers"
)

// Legacy handler functions for backward compatibility
// These delegate to the new structured handlers

func homeHandler(c echo.Context) error {
	return handlers.HomeHandler(c)
}

func aboutHandler(c echo.Context) error {
	return handlers.AboutHandler(c)
}

func serviceWorker(c echo.Context) error {
	return handlers.ServiceWorkerHandler(c)
}
