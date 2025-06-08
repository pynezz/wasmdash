package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

/*
ThemePerformance tracks theme-related performance

- returns a middleware function that logs slow responses
*/
func ThemePerformance() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			if c.Request().URL.Path == "/static/css/styles.css" {
				c.Response().Header().Set("Cache-Control", "public, max-age=31536000, immutable")
				c.Response().Header().Set("Vary", "Accept-Encoding")
			}
			err := next(c)

			// If the request was successful and took longer than 100ms
			if err == nil && time.Since(start) > 100*time.Millisecond {
				c.Logger().Warnf("Slow response: %s took %v", c.Request().URL.Path, time.Since(start))
			}

			return err
		}
	}
}
