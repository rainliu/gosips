package parser

import (
	"testing"
)

func TestPriorityParser(t *testing.T) {
	var tvi = []string{
		"Priority: emergency\n",
	}
	var tvo = []string{
		"Priority: emergency\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewPriorityParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
