package parser

import (
	"testing"
)

func TestUnsupportedParser(t *testing.T) {
	var tvi = []string{
		"Unsupported: foo \n",
		"Unsupported: foo1, foo2 ,foo3 , foo4\n",
	}
	var tvo = []string{
		"Unsupported: foo \n",
		"Unsupported: foo1,foo2,foo3,foo4\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewUnsupportedParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String unsupported[] = {
                "Unsupported: foo \n",
                "Unsupported: foo1, foo2 ,foo3 , foo4\n"
                };

		for (int i = 0; i < unsupported.length; i++ ) {
		    UnsupportedParser parser =
			  new UnsupportedParser(unsupported[i]);
		    UnsupportedList u= (UnsupportedList) parser.parse();
		    System.out.println("encoded = " + u.encode());
		}

	}
*/
