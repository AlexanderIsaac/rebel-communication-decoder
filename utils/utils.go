package utils

import (
	"encoding/base64"
	"encoding/json"

	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
)

func BodyDumpHandler(ctx echo.Context, reqBody, _ []byte) {
	str := base64.StdEncoding.EncodeToString([]byte(reqBody))
	data, _ := base64.StdEncoding.DecodeString(str)

	bodyString := string(data)

	var bodyJSON map[string]interface{}

	err := json.Unmarshal([]byte(bodyString), &bodyJSON)

	if err != nil {
		slogecho.AddCustomAttributes(ctx, slog.Any("body", bodyString))
	} else {
		slogecho.AddCustomAttributes(ctx, slog.Any("body", bodyJSON))
	}

}

func RequestHandler(ctx echo.Context, requestID string) {
	ctx.Set("RequestID", requestID)
}
