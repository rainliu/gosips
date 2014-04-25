package parser

import (
	"testing"
)

func TestToParser(t *testing.T) {
	var tvs = []string{
		"To: <sip:+1-650-555-2222@ss1.wcom.com;user=phone>;tag=5617\n",
		"To: T. A. Watson <sip:watson@bell-telephone.com>\n",
		"To: LittleGuy <sip:UserB@there.com>\n",
		"To: sip:mranga@120.6.55.9\n",
		"To: sip:mranga@129.6.55.9 ; tag=696928473514.129.6.55.9\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewToParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}
