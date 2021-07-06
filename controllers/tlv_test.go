package controllers_test

import (
    "bytes"
    "net/http"
    "testing"
    "github.com/api/tlv-api/controllers"
    "github.com/gin-gonic/gin"
    "net/http/httptest"
)

func TestGetTlv(t *testing.T) {
    gin.SetMode(gin.TestMode)
    resp := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(resp)
    var param = []byte(`{"data": [49 ,49 ,65 ,66 ,51 ,57 ,56 ,55 ,54 ,53 ,85 ,74 ,49 ,65 ,48 ,53 ,48 ,50 ,48 ,48 ,78 ,50 ,51]}`)
    c.Request, _ = http.NewRequest(http.MethodPost, "/getTlv", bytes.NewBuffer(param))
    c.Request.Header.Add("Content-Type", "application/json")
    c.GetRawData()
    controllers.GetTlv(c)
}