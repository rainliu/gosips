package parser

import (
	"testing"
)

func TestErrorInfoParser(t *testing.T) {
	var tvs = []string{
		"Error-Info: <sip:not-in-service-recording@atlanta.com>\n",
		"Error-Info: <sip:not-in-service-recording@atlanta.com>;param1=oli\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewErrorInfoParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Error-Info: <sip:not-in-service-recording@atlanta.com>\n",
          "Error-Info: <sip:not-in-service-recording@atlanta.com>;param1=oli\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ErrorInfoParser parser =
          new ErrorInfoParser(r[i]);
          ErrorInfoList e= (ErrorInfoList) parser.parse();
          System.out.println("encoded = " + e.encode());
      }
  }
*/
