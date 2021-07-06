package controllers

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	helpers "github.com/api/tlv-api/helpers"
	httpmodels "github.com/api/tlv-api/httpmodels"
	utils "github.com/api/tlv-api/utils"
)

/* Ejemplo de datos usados:
- String: 11AB398765UJ1A050200N23
- String: 04HOLAA050200N2306PRUEBAA06
*/
//Acceso a endpoint con validacion de datos recibidos
func GetTlv(c *gin.Context) {
	var count int
	var status bool = true
	var err error
	params := &httpmodels.TlvRequest{}
	if err = c.Bind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{"err": err.Error()}})
		return
	}
	if params == nil || params.Data == nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Datos Incompletos", "data": gin.H{"err": ""}})
		return
	}
	param := bytes.NewBuffer(params.Data).String()
	if len(param) <= utils.SizeLenght {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Formato invalido", "data": gin.H{"err": ""}})
		return
	}
	err, response, status, count := helpers.GetTlv(param)
	if status == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Formato invalido en el lenght: " + strconv.Itoa(count), "data": gin.H{"err": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "",
		"data":    response,
	})
}
