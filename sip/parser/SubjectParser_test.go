package parser

import (
	"testing"
)

func TestSubjectParser(t *testing.T) {
	var tvi = []string{
		"Subject: Where is the Moscone?\n",
		"Subject: Need more boxes\n",
	}
	var tvo = []string{
		"Subject: Where is the Moscone?\n",
		"Subject: Need more boxes\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewSubjectParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
