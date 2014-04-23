package parser

import (
	"testing"
)

func TestAcceptEncodingParser(t *testing.T) {
	var acceptEncodings = []string{
		"Accept-Encoding: compress, gzip\n",
		"Accept-Encoding   :\n",
		"Accept-Encoding	 : *\n",
		"Accept-Encoding	: compress;q=0.5, gzip;q=1.0\n",
		"Accept-Encoding  : gzip;q=1.0, identity; q=0.5, *;q=0\n",
	}

	for i := 0; i < len(acceptEncodings); i++ {
		shp := NewAcceptEncodingParser(acceptEncodings[i])
		testHeaderParser(t, shp)
	}
}
