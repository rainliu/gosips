package parser

import (
	"testing"
)

func TestAlertInfoParser(t *testing.T) {
	var tvi = []string{
		"Alert-Info: <http://www.example.com/sounds/moo.wav>\n",
	}
	var tvo = []string{
		"Alert-Info: <http://www.example.com/sounds/moo.wav>\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewAlertInfoParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}
