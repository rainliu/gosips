package parser

import (
	"testing"
)

func TestMimeVersionParser(t *testing.T) {
	var tvs = []string{
		"MIME-Version: 1.0 \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewMimeVersionParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
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
