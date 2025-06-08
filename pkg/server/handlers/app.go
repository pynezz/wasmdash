package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezzentials/ansi"
	"github.com/pynezz/wasmdash/pkg/core"
	"github.com/pynezz/wasmdash/pkg/server/middleware"
	"github.com/pynezz/wasmdash/pkg/ui"
	"github.com/pynezz/wasmdash/pkg/ui/pages"
)

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

func AboutHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.About(c.Path()))
}

func ServiceWorkerHandler(c echo.Context) error {
	if err := middleware.Log(c); err != nil {
		ansi.PrintError("Error logging request: " + err.Error())
	}
	return Render(c, http.StatusOK, ui.ServiceWorker())
}

func RobotsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
}

func NotFoundHandler(c echo.Context) error {
	return Render(c, http.StatusNotFound, pages.NotFound())
}

func HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status":    "healthy",
		"timestamp": http.StatusOK,
		"version":   "1.0.1",
		"services": map[string]string{
			"css":    "available",
			"static": "available",
			"templ":  "available",
		},
	})
}

func prodcsp(ctx echo.Context) string {
	if ctx.Get("environment") == "prod" {
		return "upgrade-insecure-requests; block-all-mixed-content"
	}
	return ""
}

func getcsp(ctx echo.Context) {
	ctx.Response().Header().Set("Content-Security-Policy",
		fmt.Sprintf("default-src 'none'; script-src 'self' 'nonce-%s'; style-src 'self'; img-src 'self' *.github.com; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; frame-src 'none'; base-uri 'self'; form-action 'self'; %s",
			ctx.Get("nonce"), prodcsp(ctx)))
}

// Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render()
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	nonce, err := core.GenerateNonce()
	if err != nil {
		ansi.PrintError("Error generating nonce: " + err.Error())

		nonce = "fallback-nonce"
	}

	templCtx := templ.WithNonce(ctx.Request().Context(), nonce)

	// Set Content Security Policy with proper nonce
	ctx.Set("nonce", nonce)
	getcsp(ctx)

	if err := ui.Layout(t, nonce, ctx.Path()).Render(templCtx, buf); err != nil {
		return err
	}

	ctx.Set("nonce", nonce)
	return ctx.HTML(statusCode, buf.String())
}
