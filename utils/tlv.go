package utils

import (
	"errors"

	httpmodels "github.com/api/tlv-api/httpmodels"
)

const (
	SizeLenght = 2
	SizeType   = 3
)

const (
	TypeString = "Alfanumberico"
	TypeNumber = "Numerico"
)

//Asignacion del tlv para generar response del endpoint
func GetTlv(param string, lenght int) (error, *httpmodels.TlvResponse) {
	if len(param) >= (lenght + SizeLenght + SizeType) {
		value := param[:lenght+SizeLenght+SizeType]
		var typeTlv string
		if value[lenght+SizeLenght:lenght+SizeType] == "A" {
			typeTlv = TypeString
		} else if value[lenght+SizeLenght:lenght+SizeType] == "N" {
			typeTlv = TypeNumber
		} else {
			return errors.New("Tipo de dato invalido"), nil
		}
		tlv := &httpmodels.TlvResponse{
			Type:  typeTlv + " - Numero de campo: " + value[lenght+SizeType:lenght+SizeType+2],
			Len:   lenght,
			Value: value[SizeLenght : lenght+SizeLenght],
		}
		return nil, tlv
	} else {
		return errors.New("Tamaño del tlv invalido"), nil
	}
}
