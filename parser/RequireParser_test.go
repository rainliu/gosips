package parser

import (
	"testing"
)

func TestRequireParser(t *testing.T) {
	var tvi = []string{
		"Require: 100rel \n",
		"Require: 100rel,200ok,389\n",
	}
	var tvo = []string{
		"Require: 100rel \n",
		"Require: 100rel,200ok,389\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewRequireParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
