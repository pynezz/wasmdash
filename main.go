package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezzentials/ansi"
	"github.com/pynezz/wasmdash/pkg/core"
	"github.com/pynezz/wasmdash/pkg/ui"
	"github.com/pynezz/wasmdash/pkg/ui/pages"
)

//go:embed static/*
var static embed.FS

//go:generate tailwindcss -i assets/css/base.css -o static/css/styles.css -m

func init() {
	log.Println("initializing application")
}

func main() {
	app := echo.New()
	ctx := app.AcquireContext()
	defer app.ReleaseContext(ctx)

	serve(app)
	log.Println("application started")
	log.Fatalln(http.ListenAndServe(":8080", app))
}

func serve(app *echo.Echo) {
	log.Println("serving on port 8080")
	app.HideBanner = true
	// set server header
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Response().Header().Set(echo.HeaderServer, "localhost:8080")
			return next(ctx)
		}
	})

	app.Static("/static", "static")
	app.GET("/", homeHandler)
	app.GET("/about", aboutHandler)
	app.GET("/service-worker.js", serviceWorker)
	app.GET("/robots.txt", func(c echo.Context) error {
		return c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	})

	app.GET("/404", func(c echo.Context) error {
		return Render(c, 404, pages.NotFound())
	})

}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	nonce, err := core.GenerateNonce()
	templCtx := templ.WithNonce(ctx.Request().Context(), nonce)
	ctx.Response().Header().Set("Content-Security-Policy", "default-src 'none'; script-src 'self' 'nonce-"+nonce+"'; style-src 'self'; img-src 'self' *.github.com; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; frame-src 'none'; base-uri 'self'; form-action 'self'; upgrade-insecure-requests;")

	if err != nil {
		ansi.PrintError("Error generating nonce: " + nonce)
	}

	if err := ui.Layout(t, nonce, ctx.Path()).Render(templCtx, buf); err != nil {
		return err
	}
	// if err := templates.Root(t, nonce, ctx.Path()).Render(ctx.Request().Context(), buf); err != nil {
	// 	return err
	// }

	ctx.Set("nonce", nonce)

	return ctx.HTML(statusCode, buf.String())
}
