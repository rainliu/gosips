package parser

import (
	"testing"
)

func TestOrganizationParser(t *testing.T) {
	var tvs = []string{
		"Organization: Boxes by Bob\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewOrganizationParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
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
