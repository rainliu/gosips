package parser

import (
	"testing"
)

func TestSupportedParser(t *testing.T) {
	var tvs = []string{
		"Supported: 100rel \n",
		"Supported: foo1, foo2 ,foo3 , foo4 \n",
	}
	var tvs_o = []string{
		"Supported: 100rel \n",
		"Supported: foo1,foo2,foo3,foo4 \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewSupportedParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String supported[] = {
          "Supported: 100rel \n",
          "Supported: foo1, foo2 ,foo3 , foo4 \n"
      };

      for (int i = 0; i < supported.length; i++ ) {
          SupportedParser parser =
          new SupportedParser(supported[i]);
          SupportedList s= (SupportedList) parser.parse();
          System.out.println("encoded = " + s.encode());
      }

  }
*/
