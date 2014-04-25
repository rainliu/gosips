package parser

import (
	"testing"
)

func TestAddressParser(t *testing.T) {
	var tvi = []string{
		"<sip:user@example.com?Route=%3csip:sip.example.com%3e>",
		"\"M. Ranganathan\"   <sip:mranga@nist.gov>",
		"<sip:+1-650-555-2222@ss1.wcom.com;user=phone>",
		"M. Ranganathan <sip:mranga@nist.gov>",
	}

	for i := 0; i < len(tvi); i++ {
		addressParser := NewAddressParser(tvi[i])
		if addr, err := addressParser.Address(); err != nil {
			t.Log(err)
			t.Fail()
		} else {
			t.Log("encoded = " + addr.String())
		}
	}
}
