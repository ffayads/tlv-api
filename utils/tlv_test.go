package utils_test

import (
    "strconv"
	"testing"

	"github.com/api/tlv-api/utils"
)

func TestGetTlv(t *testing.T) {
    params := []string{"11AB398765UJ1A05","26AB398765UJ1A05","11AB398765UJ1B05"}
    for _, param := range params {
        lenght, _ := strconv.Atoi(param[0 : utils.SizeLenght])
        var paramPosition string
        if len(param) <= lenght+utils.SizeType{
            paramPosition = param[0:]
        } else {
            paramPosition = param[0 : lenght+utils.SizeLenght+utils.SizeType]
        }
        utils.GetTlv(paramPosition, lenght)
    }
}
