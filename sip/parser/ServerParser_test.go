package parser

import (
	"testing"
)

func TestServerParser(t *testing.T) {
	var tvi = []string{
		"Server: Softphone/Beta1.5 \n",
		"Server: HomeServer v2\n",
		"Server: Nist/Beta1 (beta version) \n",
		"Server: Nist proxy (beta version)\n",
		"Server: Nist1.0/Beta2 UbiServer/vers.1.0 (new stuff) (Cool) \n",
	}
	var tvo = []string{
		"Server: Softphone/Beta1.5 \n",
		"Server: HomeServer v2\n",
		"Server: Nist/Beta1 (beta version) \n",
		"Server: Nist proxy (beta version)\n",
		"Server: Nist1.0/Beta2 UbiServer/vers.1.0 (new stuff) (Cool) \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewServerParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String server[] = {
                "Server: Softphone/Beta1.5 \n",
                "Server: HomeServer v2\n",
                "Server: Nist/Beta1 (beta version) \n",
                "Server: Nist proxy (beta version)\n",
                "Server: Nist1.0/Beta2 UbiServer/vers.1.0 (new stuff) (Cool) \n"
                };

		for (int i = 0; i < server.length; i++ ) {
		    ServerParser parser =
			  new ServerParser(server[i]);
		    Server s= (Server) parser.parse();
		    System.out.println("encoded = " + s.encode());
		}

	}
*/
