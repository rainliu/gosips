package parser

import (
	"testing"
)

func TestUnsupportedParser(t *testing.T) {
	var tvs = []string{
		"Unsupported: foo \n",
		"Unsupported: foo1, foo2 ,foo3 , foo4\n",
	}
	var tvs_o = []string{
		"Unsupported: foo \n",
		"Unsupported: foo1,foo2,foo3,foo4\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewUnsupportedParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
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
