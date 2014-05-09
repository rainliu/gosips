package parser

import (
	"testing"
)

func TestMinExpiresParser(t *testing.T) {
	var tvi = []string{
		"Min-Expires: 60 \n",
	}
	var tvo = []string{
		"Min-Expires: 60 \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewMinExpiresParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String r[] = {
                "Min-Expires: 60 \n"
                };

		for (int i = 0; i < r.length; i++ ) {
		    MinExpiresParser parser =
			  new MinExpiresParser(r[i]);
		    MinExpires m= (MinExpires) parser.parse();
		    System.out.println("encoded = " +m.encode());
		}
	}
*/
