package parser

import (
	"testing"
)

func TestErrorInfoParser(t *testing.T) {
	var tvi = []string{
		"Error-Info: <sip:not-in-service-recording@atlanta.com>\n",
		"Error-Info: <sip:not-in-service-recording@atlanta.com>;param1=oli\n",
	}
	var tvo = []string{
		"Error-Info: <sip:not-in-service-recording@atlanta.com>\n",
		"Error-Info: <sip:not-in-service-recording@atlanta.com>;param1=oli\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewErrorInfoParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
