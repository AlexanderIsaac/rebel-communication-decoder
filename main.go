package main

import (
	"app/internal/adapter/inbound/http"
	"app/internal/adapter/outbound/repository"

	"app/utils"
	"log/slog"
	"os"

	"app/internal/core/service"

	"github.com/go-playground/validator/v10"
	// "github.com/labstack/echo-contrib/echoprometheus"
	_ "app/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/swag"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	return logger
}

func configureEcho(logger *slog.Logger) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HideBanner = true
	e.HidePort = true

	config := slogecho.Config{
		WithTraceID:      true,
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
		WithRequestID:    true,
		WithUserAgent:    true,
		WithSpanID:       true,
		Filters: []slogecho.Filter{slogecho.IgnorePathContains("swagger"),
			slogecho.IgnorePathPrefix("/swagger"), slogecho.IgnorePathContains("favicon")},
	}

	e.Use(slogecho.NewWithConfig(logger, config))
	e.Use(middleware.Secure())
	e.Use(middleware.CSRF())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: utils.RequestHandler,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: utils.BodyDumpHandler,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
	// 	Namespace: "challenge",
	// })) // adds middleware to gather metrics
	// go func() {
	// 	metrics := echo.New()                                // this Echo will run on separate port 8081
	// 	metrics.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics
	// 	metrics.HideBanner = true
	// 	metrics.HidePort = true
	// 	metrics.Logger.Fatal(metrics.Start(":8081"))
	// }()

	return e
}

func setupServicesAndRoutes(e *echo.Echo) {
	repo := repository.NewSatelliteRepository()
	decipherService := service.NewDecipherService(repo)
	satelliteService := service.NewSatelliteService(repo)
	http.RegisterRoutes(e, decipherService, satelliteService)
}

// @title Challenge
// @version 1.0
// @description This is the Challenge API documentation.
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger := createLogger()
	e := configureEcho(logger)
	setupServicesAndRoutes(e)

	port := ":8080"
	logger.Info("Server starting on port " + port)
	e.Logger.Fatal(e.Start(port))
}
