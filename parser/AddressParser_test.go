package parser

import (
	"strings"
	"testing"
)

func TestAddressParser(t *testing.T) {
	var tvi = []string{
		"<sip:user@example.com?Route=%3csip:sip.example.com%3e>",
		"\"M. Ranganathan\"   <sip:mranga@nist.gov>",
		"<sip:+1-650-555-2222@ss1.wcom.com;user=phone>",
		"M. Ranganathan <sip:mranga@nist.gov>",
	}
	var tvo = []string{
		"<sip:user@example.com?Route=%3csip:sip.example.com%3e>",
		"\"M. Ranganathan\" <sip:mranga@nist.gov>",
		"<sip:+1-650-555-2222@ss1.wcom.com;user=phone>",
		"\"M. Ranganathan\" <sip:mranga@nist.gov>",
	}

	for i := 0; i < len(tvi); i++ {
		addressParser := NewAddressParser(tvi[i])
		if addr, err := addressParser.Address(); err != nil {
			t.Log(err)
			t.Fail()
		} else {
			d := addr.String()
			s := tvo[i]

			if strings.TrimSpace(d) != strings.TrimSpace(s) {
				t.Log("origin = " + s)
				t.Log("failed = " + d)
				t.Fail()
			}
		}
	}
}
