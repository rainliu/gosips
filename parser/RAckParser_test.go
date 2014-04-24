package parser

import (
	"testing"
)

func TestRAckParser(t *testing.T) {
	var tvs = []string{
		"RAck: 776656 1 INVITE\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewRAckParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String r[] = {
                "RAck: 776656 1 INVITE\n"
                };

		for (int i = 0; i < r.length; i++ ) {
		    RAckParser parser =
			  new RAckParser(r[i]);
		    RAck ra= (RAck) parser.parse();
		    System.out.println("encoded = " + ra.encode());
		}
	}
*/
