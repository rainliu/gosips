package parser

import (
	"testing"
)

func TestAllowParser(t *testing.T) {
	var allow = []string{
		"Allow		: INVITE, 		ACK,	   OPTIONS, CANCEL, BYE\n",
		"Allow	 :	  INVITE\n",
	}

	for i := 0; i < len(allow); i++ {
		shp := NewAllowParser(allow[i])
		testHeaderParser(t, shp)
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
