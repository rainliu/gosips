package parser

import (
	"testing"
)

func TestAcceptParser(t *testing.T) {
	var tvi = []string{
		"Accept: application/sdp;level=1,application/x-private,text/html\n",
	}
	var tvo = []string{
		"Accept: application/sdp;level=1,application/x-private,text/html\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewAcceptParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}
