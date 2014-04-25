package parser

import (
	"testing"
)

func TestMaxForwardsParser(t *testing.T) {
	var tvi = []string{
		"Max-Forwards: 35\n",
		"Max-Forwards: 3495\n",
		"Max-Forwards: 0 \n",
	}
	var tvo = []string{
		"Max-Forwards: 35\n",
		"Max-Forwards: 0\n",
		"Max-Forwards: 0 \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewMaxForwardsParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}
