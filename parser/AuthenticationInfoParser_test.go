package parser

import (
	"testing"
)

func TestAuthenticationInfoParser(t *testing.T) {
	var tvs = []string{
		"Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\"\n",
		"Authentication-Info  :   nextnonce   =     \"47364c23432d2e131a5fb210812c\"  ,       rspauth=\"hello\"\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewAuthenticationInfoParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/**
    public static void main(String args[]) throws ParseException {
        String r[] = {
            "Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\"\n",
            "Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\",rspauth=\"hello\"\n"
        };

        for (int i = 0; i < r.length; i++ ) {
            AuthenticationInfoParser parser =
            new AuthenticationInfoParser(r[i]);
            AuthenticationInfo a= (AuthenticationInfo) parser.parse();
            System.out.println("encoded = " + a.encode());
        }
    }
**/
