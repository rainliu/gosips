package parser

import (
	"testing"
)

func TestSubjectParser(t *testing.T) {
	var tvs = []string{
		"Subject: Where is the Moscone?\n",
		"Subject: Need more boxes\n",
	}
	var tvs_o = []string{
		"Subject: Where is the Moscone?\n",
		"Subject: Need more boxes\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewSubjectParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String subject[] = {
                "Subject: Where is the Moscone?\n",
                "Subject: Need more boxes\n"
                };

		for (int i = 0; i < subject.length; i++ ) {
		    SubjectParser parser =
			  new SubjectParser(subject[i]);
		    Subject s= (Subject) parser.parse();
		    System.out.println("encoded = " +s.encode());
		}

	}
*/
