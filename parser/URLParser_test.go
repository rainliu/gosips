package parser

import (
	"strings"
	"testing"
)

func TestURLParser(t *testing.T) {
	var tvi = []string{
		"sip:conference=1234@sip.convedia.com;xyz=pqd",
		"sip:herbivore.ncsl.nist.gov:5070;maddr=129.6.55.251;lc",
		"sip:1-301-975-3664@foo.bar.com;user=phone",
		"sip:129.6.55.181",
		"sip:herbivore.ncsl.nist.gov:5070;maddr=129.6.55.251?method=INVITE&contact=sip:foo.bar.com",
		"sip:j.doe@big.com",
		"sip:j.doe:secret@big.com;transport=tcp",
		"sip:j.doe@big.com?subject=project",
		"sip:+1-212-555-1212:1234@gateway.com;user=phone",
		"sip:1212@gateway.com",
		"sip:alice@10.1.2.3",
		"sip:alice@example.com",
		"sip:alice",
		"sip:alice@registrar.com;method=REGISTER",
		"sip:annc@10.10.30.186:6666;early=no;play=http://10.10.30.186:8080/examples/pin.vxml",
		"tel:+463-1701-4291",
		"tel:46317014291",
		"http://10.10.30.186:8080/examples/pin.vxml",
	}
	var tvo = []string{
		"sip:conference=1234@sip.convedia.com;xyz=pqd",
		"sip:herbivore.ncsl.nist.gov:5070;maddr=129.6.55.251;lc",
		"sip:1-301-975-3664@foo.bar.com;user=phone",
		"sip:129.6.55.181",
		"sip:herbivore.ncsl.nist.gov:5070;maddr=129.6.55.251?method=INVITE&contact=sip:foo.bar.com",
		"sip:j.doe@big.com",
		"sip:j.doe:secret@big.com;transport=tcp",
		"sip:j.doe@big.com?subject=project",
		"sip:+1-212-555-1212:1234@gateway.com;user=phone",
		"sip:1212@gateway.com",
		"sip:alice@10.1.2.3",
		"sip:alice@example.com",
		"sip:alice",
		"sip:alice@registrar.com;method=REGISTER",
		"sip:annc@10.10.30.186:6666;early=no;play=http://10.10.30.186:8080/examples/pin.vxml",
		"tel:+463-1701-4291",
		"tel:46317014291",
		"http://10.10.30.186:8080/examples/pin.vxml",
	}

	for i := 0; i < len(tvi); i++ {
		hp := NewURLParser(tvi[i])
		if sh, err := hp.Parse(); err != nil {
			t.Log(err)
			t.Fail()
		} else {
			d := sh.String()
			s := tvo[i]

			if strings.TrimSpace(d) != strings.TrimSpace(s) {
				t.Log("origin = " + s)
				t.Log("failed = " + d)
				t.Fail()
			}
		}
	}
}
