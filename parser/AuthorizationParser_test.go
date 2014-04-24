package parser

import (
	"testing"
)

func TestAuthorizationParser(t *testing.T) {
	var tvs = []string{
		"Authorization: Digest username=\"UserB\", realm=\"MCI WorldCom SIP\"," +
			" nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\", opaque=\"\"," +
			" uri=\"sip:ss2.wcom.com\", response=\"dfe56131d1958046689cd83306477ecc\"\n",
		"Authorization: Digest username=\"aprokop\",realm=\"Realm\",nonce=\"MTA1MDMzMjE5ODUzMjUwM2QyMzBhOTJlMTkxYjIxYWY1NDlhYzk4YzNiMGYz\",uri=\"sip:nortelnetworks.com:5060\",response=\"dbfba6c0e9664b45b7d224d2b52a1d01\",algorithm=\"MD5\",cnonce=\"VG05eWRHVnNJRTVsZEhkdmNtdHpNVEExTURNek16WTFOREUyTUE9PQ==\",qop=auth-int,nc=00000001\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewAuthorizationParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String auth[] = {

"Authorization: Digest username=\"UserB\", realm=\"MCI WorldCom SIP\","+
" nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\", opaque=\"\","+
" uri=\"sip:ss2.wcom.com\", response=\"dfe56131d1958046689cd83306477ecc\"\n",

"Authorization: Digest username=\"aprokop\",realm=\"Realm\",nonce=\"MTA1MDMzMjE5ODUzMjUwM2QyMzBhOTJlMTkxYjIxYWY1NDlhYzk4YzNiMGYz\",uri=\"sip:nortelnetworks.com:5060\",response=\"dbfba6c0e9664b45b7d224d2b52a1d01\",algorithm=\"MD5\",cnonce=\"VG05eWRHVnNJRTVsZEhkdmNtdHpNVEExTURNek16WTFOREUyTUE9PQ==\",qop=auth-int,nc=00000001\n"

    		};

		for (int i = 0; i <auth.length; i++ ) {
		    AuthorizationParser ap =
			  new AuthorizationParser(auth[i]);
		    Authorization a= (Authorization) ap.parse();
		    System.out.println("encoded = " + a.encode());
		}

	}

**/
