package parser

import (
	"testing"
)

func TestMimeVersionParser(t *testing.T) {
	var tvi = []string{
		"MIME-Version: 1.0 \n",
	}
	var tvo = []string{
		"MIME-Version: 1.0 \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewMimeVersionParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String r[] = {
                "MIME-Version: 1.0 \n"
                };

		for (int i = 0; i < r.length; i++ ) {
		    MimeVersionParser parser =
			  new MimeVersionParser(r[i]);
		    MimeVersion m= (MimeVersion) parser.parse();
		    System.out.println("encoded = " +m.encode());
		}
	}
*/
