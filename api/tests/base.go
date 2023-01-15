package testcases

import (
	"book-api/api/router"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeRequest(t *testing.T, method string, url string, data []byte, expectedStatusCode int) {
	routerInstance := router.InitRouter(false)
	w := httptest.NewRecorder()
	if data != nil {
		bodyReader := bytes.NewReader(data)
		req, _ := http.NewRequest(method, url, bodyReader)
		routerInstance.ServeHTTP(w, req)
	} else {
		req, _ := http.NewRequest(method, url, nil)
		routerInstance.ServeHTTP(w, req)
	}

	assert.Equal(t, expectedStatusCode, w.Code)
}
