package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/BlackRoad-OS/LocalAI/core/config"
	"github.com/BlackRoad-OS/LocalAI/core/http/middleware"
	"github.com/BlackRoad-OS/LocalAI/core/services"
	"github.com/BlackRoad-OS/LocalAI/internal"
)

func registerBackendGalleryRoutes(app *echo.Echo, appConfig *config.ApplicationConfig, galleryService *services.GalleryService, opcache *services.OpCache) {
	// Show the Backends page (all backends are loaded client-side via Alpine.js)
	app.GET("/browse/backends", func(c echo.Context) error {
		summary := map[string]interface{}{
			"Title":        "LocalAI - Backends",
			"BaseURL":      middleware.BaseURL(c),
			"Version":      internal.PrintableVersion(),
			"Repositories": appConfig.BackendGalleries,
		}

		// Render index - backends are now loaded via Alpine.js from /api/backends
		return c.Render(200, "views/backends", summary)
	})
}
