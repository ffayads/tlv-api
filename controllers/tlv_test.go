package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/api/tlv-api/controllers"
	"github.com/gin-gonic/gin"
)

func TestGetTlv(t *testing.T) {
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	params := []string{
		`{"data": [49 ,49 ,65 ,66 ,51 ,57 ,56 ,55 ,54 ,53 ,85 ,74 ,49 ,65 ,48 ,53 ,48 ,50 ,48 ,48 ,78 ,50 ,51]}`,
		"",
		`{"data": 2}`,
		`{"data": [50]}`,
		`{"data": [53 ,48 ,65 ,66 ,51 ,57 ,56 ,55 ,54 ,53 ,85 ,74 ,49 ,65 ,48 ,53 ,48 ,50 ,48 ,48 ,78 ,50 ,51]}`,
		`{"data": [49 ,49 ,65 ,66 ,51 ,57 ,56 ,55 ,54 ,53 ,85 ,74 ,49 ,66 ,48 ,53 ,48 ,50 ,48 ,48 ,78 ,50 ,51]}`,
	}
	for _, jsonData := range params {
		param := []byte(jsonData)
		c.Request, _ = http.NewRequest(http.MethodPost, "/getTlv", bytes.NewBuffer(param))
		c.Writer.Header().Set("Content-Type", "application/json")
		controllers.GetTlv(c)
	}
}
