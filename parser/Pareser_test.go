package parser

import (
	"testing"
)

func testParser(t *testing.T, hp Parser) {
	if sh, err := hp.Parse(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log("encoded = " + sh.String())
	}
}
