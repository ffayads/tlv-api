package httpmodels_test

import (
	"fmt"
	"testing"

	"github.com/api/tlv-api/httpmodels"
)

func TestTlvRequest(t *testing.T) {
	param := []byte{49 ,49 ,65 ,66 ,51 ,57 ,56 ,55 ,54 ,53 ,85 ,74 ,49 ,65 ,48 ,53 ,48 ,50 ,48 ,48 ,78 ,50 ,51}
	tlv := &httpmodels.TlvRequest{
		Data: param,
	}
	fmt.Println(tlv)
}

func TestTlvResonse(t *testing.T) {
	tlv := &httpmodels.TlvResponse{
		Type: "Number",
		Len: 11,
		Value: "value",
	}
	fmt.Println(tlv)
}