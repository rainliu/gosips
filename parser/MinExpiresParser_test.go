package parser

import (
	"testing"
)

func TestMinExpiresParser(t *testing.T) {
	var minexpires = []string{
		"Min-Expires: 60 \n",
	}

	for i := 0; i < len(minexpires); i++ {
		shp := NewMinExpiresParser(minexpires[i])
		testParser(t, shp)
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
