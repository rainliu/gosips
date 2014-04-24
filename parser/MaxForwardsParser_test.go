package parser

import (
	"testing"
)

func TestMaxForwardsParser(t *testing.T) {
	var maxforwards = []string{
		"Max-Forwards: 35\n",
		"Max-Forwards: 3495\n",
		"Max-Forwards: 0 \n",
	}

	for i := 0; i < len(maxforwards); i++ {
		shp := NewMaxForwardsParser(maxforwards[i])
		testHeaderParser(t, shp)
	}
}
