package helpers_test

import (
	"testing"

	"github.com/api/tlv-api/helpers"
)

func TestGetTlv(t *testing.T) {
	params := []string{"27AB398765UJ1A050200N23","11AB398765UJ1C050200N23","11AB398765UJ1A050200N23"}
	for _, param := range params {
		err, tlv, status, count := helpers.GetTlv(param)
		if err != nil {
			t.Fail()
		} else if tlv == nil && status == false {
			t.Fail()
		} else if count < 1 {
			t.Fail()
		}
	}
}
