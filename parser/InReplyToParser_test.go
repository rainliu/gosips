package parser

import (
	"testing"
)

func TestInReplyToParser(t *testing.T) {
	var inReplyTo = []string{
		"In-Reply-To: 70710@saturn.bell-tel.com, 17320@saturn.bell-tel.com\n",
		"In-Reply-To: 70710 \n",
	}

	for i := 0; i < len(inReplyTo); i++ {
		shp := NewInReplyToParser(inReplyTo[i])
		testParser(t, shp)
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String p[] = {
                "In-Reply-To: 70710@saturn.bell-tel.com, 17320@saturn.bell-tel.com\n",
                "In-Reply-To: 70710 \n"
                };

		for (int i = 0; i < p.length; i++ ) {
		    InReplyToParser parser =
			  new InReplyToParser(p[i]);
		    InReplyToList in= (InReplyToList) parser.parse();
		    System.out.println("encoded = " + in.encode());
		}
	}
*/
