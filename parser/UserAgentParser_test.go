package parser

import (
	"testing"
)

func TestUserAgentParser(t *testing.T) {
	var tvs = []string{
		"User-Agent: Softphone/Beta1.5 \n",
		"User-Agent: Nist/Beta1 (beta version) \n",
		"User-Agent: Nist UA (beta version)\n",
		"User-Agent: Nist1.0/Beta2 Ubi/vers.1.0 (very cool) \n",
	}
	var tvs_o = []string{
		"User-Agent: Softphone/Beta1.5 \n",
		"User-Agent: Nist/Beta1 (beta version) \n",
		"User-Agent: Nist UA (beta version)\n",
		"User-Agent: Nist1.0/Beta2 Ubi/vers.1.0 (very cool) \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewUserAgentParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/*
        public static void main(String args[]) throws ParseException {
		String userAgent[] = {
                "User-Agent: Softphone/Beta1.5 \n",
                "User-Agent: Nist/Beta1 (beta version) \n",
                "User-Agent: Nist UA (beta version)\n",
                "User-Agent: Nist1.0/Beta2 Ubi/vers.1.0 (very cool) \n"
                };

		for (int i = 0; i < userAgent.length; i++ ) {
		    UserAgentParser parser =
			  new UserAgentParser(userAgent[i]);
		    UserAgent ua= (UserAgent) parser.parse();
		    System.out.println("encoded = " + ua.encode());
		}

	}*/
