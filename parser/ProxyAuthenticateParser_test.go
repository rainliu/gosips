package parser

import (
	"testing"
)

func TestProxyAuthenticateParser(t *testing.T) {
	var tvs = []string{
		"Proxy-Authenticate: Digest realm=\"MCI WorldCom SIP\"," +
			"domain=\"sip:ss2.wcom.com\", nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\"," +
			"opaque=\"\", stale=FALSE, algorithm=MD5\n",

		"Proxy-Authenticate: Digest realm=\"MCI WorldCom SIP\"," +
			"qop=\"auth\" , nonce-value=\"oli\"\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewProxyAuthenticateParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String paAuth[] = {
     "Proxy-Authenticate: Digest realm=\"MCI WorldCom SIP\","+
     "domain=\"sip:ss2.wcom.com\", nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\","+
     "opaque=\"\", stale=FALSE, algorithm=MD5\n",

     "Proxy-Authenticate: Digest realm=\"MCI WorldCom SIP\","+
	"qop=\"auth\" , nonce-value=\"oli\"\n"
                };

		for (int i = 0; i < paAuth.length; i++ ) {
		    ProxyAuthenticateParser pap =
			  new ProxyAuthenticateParser(paAuth[i]);
		    ProxyAuthenticate pa= (ProxyAuthenticate) pap.parse();
		    System.out.println("encoded = " + pa.encode());
		}

	}
*/
