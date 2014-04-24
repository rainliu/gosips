package parser

import (
	"testing"
)

func TestAllowEventsParser(t *testing.T) {
	var allowEvents = []string{
		"Allow-Events  : pack1.pack2, 		pack3 , pack4\n",
		"Allow-Events			: 		pack1\n",
	}

	for i := 0; i < len(allowEvents); i++ {
		shp := NewAllowEventsParser(allowEvents[i])
		testParser(t, shp)
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
