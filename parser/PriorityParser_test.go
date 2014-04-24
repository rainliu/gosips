package parser

import (
	"testing"
)

func TestPriorityParser(t *testing.T) {
	var tvs = []string{
		"Priority: emergency\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewPriorityParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String p[] = {
                "Priority: emergency\n"
                };

		for (int i = 0; i < p.length; i++ ) {
		    PriorityParser parser =
			  new PriorityParser(p[i]);
		    Priority prio= (Priority) parser.parse();
		    System.out.println("encoded = " + prio.encode());
		}
	}
*/
