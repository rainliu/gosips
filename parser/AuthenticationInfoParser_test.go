package parser

import (
	"testing"
)

func TestAuthenticationInfoParser(t *testing.T) {
	var tvi = []string{
		"Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\"\n",
		"Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\",rspauth=\"hello\"\n",
	}
	var tvo = []string{
		"Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\"\n",
		"Authentication-Info: nextnonce=\"47364c23432d2e131a5fb210812c\",rspauth=\"hello\"\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewAuthenticationInfoParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
