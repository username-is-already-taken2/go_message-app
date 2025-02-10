package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestServer() *gin.Engine {
	s := &Server{}
	r := gin.New()
	r.GET("/", s.MessageHandler)
	r.GET("/_/healthz", s.HealthzHandler)
	return r
}

func testEndpoint(t *testing.T, method, path, expectedBody string, expectedStatus int) {
	r := setupTestServer()
	req, err := http.NewRequestWithContext(context.Background(), method, path, http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("Handler returned wrong status code: got %v want %v", rr.Code, expectedStatus)
	}

	if rr.Body.String() != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}

func TestMessageHandler(t *testing.T) {
	testEndpoint(t, "GET", "/", "{\"message\":\"Hello from Mondoo Engineer!\"}", http.StatusOK)
}

func TestHealthzHandler(t *testing.T) {
	testEndpoint(t, "GET", "/_/healthz", "{\"status\":\"ok\"}", http.StatusOK)
}
