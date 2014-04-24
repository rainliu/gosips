package parser

import (
	"testing"
)

func TestAlertInfoParser(t *testing.T) {
	var tvs = []string{
		"Alert-Info: <http://www.example.com/sounds/moo.wav>\n",
		"Alert-Info			:	 <http://www.example.com/sounds/moo.wav>\n",
		"Alert-Info		      :      http://www.example.com/sounds/moo.wav    \n",
		"Alert-Info  :    <http://www.example.com/sounds/moo.wav>\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewAlertInfoParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}
