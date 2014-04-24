package parser

import (
	"testing"
)

func TestMimeVersionParser(t *testing.T) {
	var mimeVersion = []string{
		"MIME-Version: 1.0 \n",
	}

	for i := 0; i < len(mimeVersion); i++ {
		shp := NewMimeVersionParser(mimeVersion[i])
		testParser(t, shp)
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
