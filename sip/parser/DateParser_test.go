package parser

import (
	"testing"
)

func TestDateParser(t *testing.T) {
	var tvi = []string{
		"Date: Sun, 07 Jan 2001 19:05:06 GMT\n",
		"Date: Mon, 08 Jan 2001 19:05:06 GMT\n",
	}
	var tvo = []string{
		"Date: Sun, 07 Jan 2001 19:05:06 GMT\n",
		"Date: Mon, 08 Jan 2001 19:05:06 GMT\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewDateParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String date[] = {
			"Date: Sun, 07 Jan 2001 19:05:06 GMT\n",
			"Date: Mon, 08 Jan 2001 19:05:06 GMT\n" };

		for (int i = 0; i < date.length; i++ ) {
		    System.out.println("Parsing " + date[i]);
		    DateParser dp =
			  new DateParser(date[i]);
		    SIPDateHeader d = (SIPDateHeader) dp.parse();
		    System.out.println("encoded = " +d.encode());
		}

	}
**/
