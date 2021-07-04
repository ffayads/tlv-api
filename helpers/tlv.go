package helpers

import (
	"strconv"

	httpmodels "github.com/api/tlv-api/httpmodels"
	utils "github.com/api/tlv-api/utils"
)

//Separacion del parametro obtenido
func GetTlv(param string) (error, []*httpmodels.TlvResponse, bool, int) {
	var position, count int
	var status bool = true
	var paramPosition string
	var err error
	response := []*httpmodels.TlvResponse{}
	for {
		count++
		if len(param) < position+utils.SizeLenght {
			if count == 1 {
				status = false
			}
			break
		}
		lenght, _ := strconv.Atoi(param[position : position+utils.SizeLenght])
		if len(param) < position+lenght+utils.SizeLenght+utils.SizeType {
			if count == 1 {
				status = false
			}
			break
		}
		paramPosition = param[position : position+lenght+utils.SizeLenght+utils.SizeType]
		err, tlv := utils.GetTlv(paramPosition, lenght)
		if err != nil {
			status = false
			break
		}
		position = position + lenght + utils.SizeLenght + utils.SizeType
		response = append(response, tlv)
	}
	return err, response, status, count
}
