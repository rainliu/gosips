package parser

import (
	"testing"
)

func TestAcceptParser(t *testing.T) {
	var accept = []string{
		"Accept: application/sdp;level=1,application/x-private, text/html\n",
		"Accept       :    application/sdp       ;      level=1,application/x-private, text/html\n",
		"Accept		:		 application/sdp;		level=1,	application/x-private, 	text/html\n",
	}

	for i := 0; i < len(accept); i++ {
		shp := NewAcceptParser(accept[i])
		testHeaderParser(t, shp)
	}
}
