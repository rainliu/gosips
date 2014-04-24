package parser

import (
	"testing"
)

func TestAllowParser(t *testing.T) {
	var tvs = []string{
		"Allow		: INVITE, 		ACK,	   OPTIONS, CANCEL, BYE\n",
		"Allow	 :	  INVITE\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewAllowParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Allow: INVITE, ACK, OPTIONS, CANCEL, BYE\n",
          "Allow: INVITE\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          AllowParser parser =
          new AllowParser(r[i]);
          AllowList a= (AllowList) parser.parse();
          System.out.println("encoded = " + a.encode());
      }
  }
*/
