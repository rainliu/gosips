package parser

import (
	"testing"
)

func TestAllowParser(t *testing.T) {
	var tvi = []string{
		"Allow: INVITE,ACK,OPTIONS,CANCEL,BYE\n",
		"Allow: INVITE\n",
	}
	var tvo = []string{
		"Allow: INVITE,ACK,OPTIONS,CANCEL,BYE\n",
		"Allow: INVITE\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewAllowParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
