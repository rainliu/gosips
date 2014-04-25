package parser

import (
	"testing"
)

func TestRAckParser(t *testing.T) {
	var tvi = []string{
		"RAck: 776656 1 INVITE\n",
	}
	var tvo = []string{
		"RAck: 776656 1 INVITE\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewRAckParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
