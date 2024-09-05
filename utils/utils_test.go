package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TTestBodyDumpHandler_Success(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{"foo":"bar"}`)))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	BodyDumpHandler(ctx, []byte(`{"foo":"bar"}`), nil)
	assert.NotNil(t, ctx)
}

func TestBodyDumpHandler_Error(t *testing.T) {

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`foobar`)))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	BodyDumpHandler(ctx, []byte(`foobar`), nil)
	assert.NotNil(t, ctx)
}

func TestRequestHandler(t *testing.T) {

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	RequestHandler(ctx, "challenge")

	requestID := ctx.Get("RequestID")
	assert.Equal(t, "challenge", requestID)
}
