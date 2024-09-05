package http

import (
	"app/internal/core/service"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, decipher *service.DecipherService, satellite *service.SatelliteService) {
	handler := NewHTTPHandler(decipher, satellite)

	prefix := "api/v1"

	// GET
	healthy := prefix + "/healthy"
	topsecretSplit := prefix + "/topsecret_split"
	slog.Debug("Mapped {" + healthy + ", GET}")
	e.GET(healthy, handler.Healthy)
	slog.Debug("Mapped {" + topsecretSplit + ", GET}")
	e.GET(topsecretSplit, handler.TopSecretSplitData)

	// POST
	topsecret := prefix + "/topsecret"
	topsecretSplitSatellite := prefix + "/topsecret_split/:satellite_name"
	e.POST(topsecret, handler.TopSecret)
	slog.Debug("Mapped {" + topsecret + ", POST}")
	e.POST(topsecretSplitSatellite, handler.TopSecretSplit)
	slog.Debug("Mapped {" + topsecretSplitSatellite + ", POST}")

}
