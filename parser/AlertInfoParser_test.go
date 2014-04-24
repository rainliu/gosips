package parser

import (
	"testing"
)

func TestAlertInfoParser(t *testing.T) {
	var alertInfo = []string{
		"Alert-Info: <http://www.example.com/sounds/moo.wav>\n",
		"Alert-Info			:	 <http://www.example.com/sounds/moo.wav>\n",
		"Alert-Info		      :      http://www.example.com/sounds/moo.wav    \n",
		"Alert-Info  :    <http://www.example.com/sounds/moo.wav>\n",
	}

	for i := 0; i < len(alertInfo); i++ {
		shp := NewAlertInfoParser(alertInfo[i])
		testHeaderParser(t, shp)
	}
}
