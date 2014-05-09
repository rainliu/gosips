package parser

import (
	"testing"
)

func TestRSeqParser(t *testing.T) {
	var tvi = []string{
		"RSeq: 988789 \n",
	}
	var tvo = []string{
		"RSeq: 988789 \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewRSeqParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String r[] = {
                "RSeq: 988789 \n"
                };

		for (int i = 0; i < r.length; i++ ) {
		    RSeqParser parser =
			  new RSeqParser(r[i]);
		    RSeq rs= (RSeq) parser.parse();
		    System.out.println("encoded = " + rs.encode());
		}
	}
*/
