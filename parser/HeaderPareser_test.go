package parser

import (
	"strings"
	"testing"
)

func testHeaderParser(t *testing.T, hp IHeaderParser, s string) {
	if sh, err := hp.Parse(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		d := sh.String()

		if strings.TrimSpace(d) != strings.TrimSpace(s) {
			t.Log("origin = " + s)
			t.Log("failed = " + d)
			t.Fail()
		} /*else {
			t.Log("passed = " + d)
		}*/
	}
}
