package parser

import (
	"testing"
)

func TestContentEncodingParserParser(t *testing.T) {
	var contentEncodingParser = []string{
		"Content-Encoding: gzip \n",
		"Content-Encoding: gzip, tar \n",
	}

	for i := 0; i < len(contentEncodingParser); i++ {
		shp := NewContentEncodingParser(contentEncodingParser[i])
		testHeaderParser(t, shp)
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Content-Encoding: gzip \n",
          "Content-Encoding: gzip, tar \n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ContentEncodingParser parser =
          new ContentEncodingParser(r[i]);
          ContentEncodingList e= (ContentEncodingList) parser.parse();
          System.out.println("encoded = " + e.encode());
      }
  }
*/
