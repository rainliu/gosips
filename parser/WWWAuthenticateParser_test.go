package parser

import (
	"testing"
)

func TestWWWAuthenticateParser(t *testing.T) {
	var tvi = []string{
		"WWW-Authenticate: Digest realm=\"MCI WorldCom SIP\"," +
			"domain=\"sip:ss2.wcom.com\",nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\"," +
			"opaque=\"\",stale=FALSE,algorithm=\"MD5\"\n",
	}
	var tvo = []string{
		"WWW-Authenticate: Digest realm=\"MCI WorldCom SIP\"," +
			"domain=\"sip:ss2.wcom.com\",nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\"," +
			"opaque=\"\",stale=FALSE,algorithm=\"MD5\"\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewWWWAuthenticateParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}
