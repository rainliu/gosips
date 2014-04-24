package parser

import (
	"testing"
)

func TestPriorityParser(t *testing.T) {
	var inputs = []string{
		"Priority: emergency\n",
	}

	for i := 0; i < len(inputs); i++ {
		shp := NewPriorityParser(inputs[i])
		testParser(t, shp)
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
