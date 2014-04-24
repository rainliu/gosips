package parser

import (
	"testing"
)

func TestRequireParser(t *testing.T) {
	var inputs = []string{
		"Require: 100rel \n",
		"Require: 100rel, 200ok , 389\n",
	}

	for i := 0; i < len(inputs); i++ {
		shp := NewRequireParser(inputs[i])
		testParser(t, shp)
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
