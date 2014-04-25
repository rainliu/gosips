package parser

import (
	"testing"
)

func TestRSeqParser(t *testing.T) {
	var tvs = []string{
		"RSeq: 988789 \n",
	}
	var tvs_o = []string{
		"RSeq: 988789 \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewRSeqParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
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
