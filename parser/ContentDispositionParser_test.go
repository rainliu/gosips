package parser

import (
	"testing"
)

func TestContentDispositionParser(t *testing.T) {
	var tvs = []string{
		"Content-Disposition: session\n",
		"Content-Disposition: render;handling=hand;optional=opt\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewContentDispositionParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Content-Disposition: session\n",
          "Content-Disposition: render;handling=hand;optional=opt \n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ContentDispositionParser parser =
          new ContentDispositionParser(r[i]);
          ContentDisposition cd= (ContentDisposition) parser.parse();
          System.out.println("encoded = " + cd.encode());
      }
  }
*/
