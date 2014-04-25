package parser

import (
	"testing"
)

func TestOrganizationParser(t *testing.T) {
	var tvi = []string{
		"Organization: Boxes by Bob\n",
	}
	var tvo = []string{
		"Organization: Boxes by Bob\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewOrganizationParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String o[] = {
                "Organization: Boxes by Bob\n"
                };

		for (int i = 0; i <o.length; i++ ) {
		    OrganizationParser parser =
			  new OrganizationParser(o[i]);
		    Organization org= (Organization) parser.parse();
		    System.out.println("encoded = " + org.encode());
		}
	}
*/
