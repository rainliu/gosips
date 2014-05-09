package parser

import (
	"testing"
)

func TestAcceptEncodingParser(t *testing.T) {
	var tvi = []string{
		"Accept-Encoding: compress,gzip\n",
		"Accept-Encoding:\n",
		"Accept-Encoding: *\n",
		"Accept-Encoding: compress;q=0.5,gzip;q=1.0\n",
		"Accept-Encoding: gzip;q=1.0,identity;q=0.5,*;q=0\n",
	}
	var tvo = []string{
		"Accept-Encoding: compress,gzip\n",
		"Accept-Encoding:\n",
		"Accept-Encoding: *\n",
		"Accept-Encoding: compress;q=0.5,gzip;q=1\n",
		"Accept-Encoding: gzip;q=1,identity;q=0.5,*;q=0\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewAcceptEncodingParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}
