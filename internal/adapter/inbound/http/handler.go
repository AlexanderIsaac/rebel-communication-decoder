package http

import (
	"app/internal/adapter/inbound/http/dto"
	"app/internal/adapter/inbound/http/model"
	coremodel "app/internal/core/model"
	"app/internal/core/service"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	decipher  *service.DecipherService
	satellite *service.SatelliteService
}

func NewHTTPHandler(decipher *service.DecipherService, satellite *service.SatelliteService) *Handler {
	return &Handler{
		decipher:  decipher,
		satellite: satellite,
	}
}

// Healthy godoc
// @Summary Check health status
// @Description Returns the health status of the service
// @Tags health
// @Produce json
// @Success 200 {object} model.Healthy
// @Router /healthy [get]
func (h *Handler) Healthy(c echo.Context) error {
	result := model.Healthy{
		Sucess: true,
	}
	return c.JSON(http.StatusOK, result)
}

// TopSecret godoc
// @Summary Decode message and determine location
// @Description Decodes a message and calculates the location of the sender
// @Tags topsecret
// @Accept json
// @Produce json
// @Param topSecretDTO body dto.TopSecretDTO true "Top Secret DTO"
// @Success 200 {object} model.TopsecretResponse
// @Router /topsecret [post]
func (h *Handler) TopSecret(c echo.Context) error {

	// Declare a variable to hold the request data.
	var topSecretDTO dto.TopSecretDTO

	// Bind the incoming JSON request body to the topSecretDTO object.
	if err := c.Bind(&topSecretDTO); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	// Validate the bound topSecretDTO object.
	if err := c.Validate(&topSecretDTO); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	// Create a slice to hold the distance information for each satellite.
	var distances []coremodel.Distance
	for _, sat := range topSecretDTO.Satellites {
		distances = append(distances, coremodel.Distance{Name: sat.Name, Distance: sat.Distance})
	}

	// Call the decipher service to calculate the location from the sender
	position, err := h.decipher.GetLocation(distances)
	if err != nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, err.Error())
	}
	slog.Info("Location calculated", slog.Any("position", position), slog.Any("id", c.Get("RequestID")))

	// Create a slice to hold the messages received for each satellite.
	var messages [][]string
	for _, sat := range topSecretDTO.Satellites {
		messages = append(messages, sat.Message)
	}

	// Call the decipher service to decode the message
	message, err := h.decipher.GetMessage(messages)
	if err != nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, err.Error())
	}
	slog.Info("Message decoded", slog.Any("message", message), slog.Any("id", c.Get("RequestID")))

	response := model.TopsecretResponse{
		Position: model.Position(position),
		Message:  message,
	}

	return c.JSON(http.StatusOK, response)
}

// TopSecretSplit godoc
// @Summary Save satellite data
// @Description Saves the message and distance data for a specific satellite
// @Tags topsecret
// @Accept json
// @Produce json
// @Param satellite_name path string true "Satellite Name"
// @Param satelliteSplit body dto.SatelliteSplit true "Satellite Split DTO"
// @Success 201 {object} model.TopSecretSplitResponse
// @Router /topsecret_split/{satellite_name} [post]
func (h *Handler) TopSecretSplit(c echo.Context) error {
	// Retrieve the satellite name from the URL path parameters.
	satelliteName := c.Param("satellite_name")

	// Declare a variable to hold the request data.
	var satelliteSplit dto.SatelliteSplit

	// Bind the incoming JSON request body to the SatelliteSplit object.
	if err := c.Bind(&satelliteSplit); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	// Validate the bound SatelliteSplit object.
	if err := c.Validate(&satelliteSplit); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err)
	}

	// Call the satellite service to save the message and distance from the sender
	savedReceivedMessage, err := h.satellite.SaveReceivedMessage(satelliteName, satelliteSplit.Distance, satelliteSplit.Message)
	if err != nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, err.Error())
	}
	slog.Info("Satellite mesage received", slog.Any("SavedMessage", savedReceivedMessage), slog.Any("id", c.Get("RequestID")))

	return c.JSON(http.StatusCreated, model.TopSecretSplitResponse{SavedReceivedMessage: savedReceivedMessage})
}

// TopSecretSplitData godoc
// @Summary Retrieve split data
// @Description Retrieves the most recent calculated position and decoded message from split data
// @Tags topsecret
// @Produce json
// @Success 200 {object} model.TopsecretResponse
// @Router /topsecret_split [get]
func (h *Handler) TopSecretSplitData(c echo.Context) error {

	// Call the decipher service to get the most recent calculated position.
	position, err := h.decipher.GetSplitLocation()
	if err != nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, err.Error())
	}
	slog.Info("Split location calculated", slog.Any("position", position), slog.Any("id", c.Get("RequestID")))

	// Call the decipher service to get the most recent decoded message.
	message, err := h.decipher.GetSplitMessage()
	if err != nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, err.Error())
	}

	response := model.TopsecretResponse{
		Position: model.Position(position),
		Message:  message,
	}
	slog.Info("Split message decoded", slog.Any("message", message), slog.Any("id", c.Get("RequestID")))

	return c.JSON(http.StatusOK, response)
}
