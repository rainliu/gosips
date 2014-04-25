package parser

import (
	"testing"
)

func TestContentDispositionParser(t *testing.T) {
	var tvi = []string{
		"Content-Disposition: session\n",
		"Content-Disposition: render;handling=hand;optional=opt\n",
	}
	var tvo = []string{
		"Content-Disposition: session\n",
		"Content-Disposition: render;handling=hand;optional=opt\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewContentDispositionParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
