package parser

import (
	"testing"
)

func TestRequireParser(t *testing.T) {
	var tvs = []string{
		"Require: 100rel \n",
		"Require: 100rel,200ok,389\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewRequireParser(tvs[i])
		testHeaderParser(t, shp, tvs[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Require: 100rel \n",
          "Require: 100rel, 200ok , 389\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          RequireParser parser =
          new RequireParser(r[i]);
          RequireList rl= (RequireList) parser.parse();
          System.out.println("encoded = " + rl.encode());
      }
  }
*/
