package parser

import (
	"testing"
)

func TestAllowEventsParser(t *testing.T) {
	var tvi = []string{
		"Allow-Events: pack1.pack2,pack3,pack4\n",
		"Allow-Events: pack1\n",
	}
	var tvo = []string{
		"Allow-Events: pack1.pack2,pack3,pack4\n",
		"Allow-Events: pack1\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewAllowEventsParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/**
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Allow-Events: pack1.pack2, pack3 , pack4\n",
          "Allow-Events: pack1\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          AllowEventsParser parser =
          new AllowEventsParser(r[i]);
          AllowEventsList a= (AllowEventsList) parser.parse();
          System.out.println("encoded = " + a.encode());
      }
  }
*/
