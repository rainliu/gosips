package parser

import (
	"testing"
)

func TestCallInfoParser(t *testing.T) {
	var call = []string{
		"Call-Info		: 	<http://wwww.example.com/alice/photo.jpg> ;purpose=icon," +
			"<http://www.example.com/alice/> ;purpose=info\n",
		"Call-Info		 :  <http://wwww.example.com/alice/photo1.jpg>\n",
	}

	for i := 0; i < len(call); i++ {
		shp := NewCallInfoParser(call[i])
		testHeaderParser(t, shp)
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Call-Info: <http://wwww.example.com/alice/photo.jpg> ;purpose=icon,"+
          "<http://www.example.com/alice/> ;purpose=info\n",
          "Call-Info: <http://wwww.example.com/alice/photo1.jpg>\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          CallInfoParser parser =
          new CallInfoParser(r[i]);
          CallInfoList e= (CallInfoList) parser.parse();
          System.out.println("encoded = " + e.encode());
      }
  }
*/
