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
			t.Log("golden = " + s)
			t.Log("failed = " + d)

			for j, k := 0, 0; j < len(s); j++ {
				if d[j] != s[j] {
					t.Logf("%d:%c vs %c", j, s[j], d[j])
					k++
					if k == 10 {
						break
					}
				}
			}

			t.Fail()
		}
	}
}
