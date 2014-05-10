package core

import (
	"testing"
)

func TestHostNameParser(t *testing.T) {
	var hostNames = []string{"foo.bar.com:1234",
		"proxima.chaplin.bt.co.uk",
		"129.6.55.181:2345",
		":2345",
	}

	for i := 0; i < len(hostNames); i++ {
		hnp := NewHostNameParser(hostNames[i])
		if hnp == nil {
			t.Fail()
		} else {
			hp, err := hnp.GetHostPort()
			if err != nil {
				t.Fail()
			} else {
				t.Log(hp.String())
			}
		}
	}
}
