package parser

import (
	"testing"
)

func TestMaxForwardsParser(t *testing.T) {
	var tvs = []string{
		"Max-Forwards: 35\n",
		"Max-Forwards: 3495\n",
		"Max-Forwards: 0 \n",
	}
	var tvs_o = []string{
		"Max-Forwards: 35\n",
		"Max-Forwards: 0\n",
		"Max-Forwards: 0 \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewMaxForwardsParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}
