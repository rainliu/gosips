package parser

import (
	"testing"
)

func TestRouteParser(t *testing.T) {
	var tvi = []string{
		"Route: <sip:alice@atlanta.com>\n",
		"Route: sip:bob@biloxi.com \n",
		"Route: sip:alice@atlanta.com, sip:bob@biloxi.com, sip:carol@chicago.com\n",
	}
	var tvo = []string{
		"Route: <sip:alice@atlanta.com>\n",
		"Route: sip:bob@biloxi.com\n",
		"Route: sip:alice@atlanta.com,sip:bob@biloxi.com,sip:carol@chicago.com\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewRouteParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
	String rou[] = {
     "Route: <sip:alice@atlanta.com>\n",
     "Route: sip:bob@biloxi.com \n",
     "Route: sip:alice@atlanta.com, sip:bob@biloxi.com, sip:carol@chicago.com\n"
         };

		for (int i = 0; i < rou.length; i++ ) {
		    RouteParser rp =
			  new RouteParser(rou[i]);
		    RouteList routeList = (RouteList) rp.parse();
		    System.out.println("encoded = " +routeList.encode());
		}

	}

*/
