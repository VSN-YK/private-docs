package routers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitRouter(t *testing.T) {
	var b_body *bytes.Buffer
	router := InitRouters()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/test-send", nil)
	router.ServeHTTP(w, req)
	//checking Response in  Body
	b_body = w.Body
	assert.Equal(t, "{\"Status\":true,\"message\":\"Ping Test\"}", b_body.String())
	assert.Equal(t, 200, w.Code)

}
