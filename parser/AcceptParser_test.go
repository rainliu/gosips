package parser

import (
	"testing"
)

func TestAcceptParser(t *testing.T) {
	var tvs = []string{
		"Accept: application/sdp;level=1,application/x-private, text/html\n",
		"Accept       :    application/sdp       ;      level=1,application/x-private, text/html\n",
		"Accept		:		 application/sdp;		level=1,	application/x-private, 	text/html\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewAcceptParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}
