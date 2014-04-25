package parser

import (
	"testing"
)

func TestCallInfoParser(t *testing.T) {
	var tvs = []string{
		"Call-Info: <http://wwww.example.com/alice/photo.jpg>;purpose=icon," +
			"<http://www.example.com/alice/>;purpose=info\n",
		"Call-Info: <http://wwww.example.com/alice/photo1.jpg>\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewCallInfoParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
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
