package parser

import (
	"testing"
)

func TestRouteParser(t *testing.T) {
	var tvs = []string{
		"Route: <sip:alice@atlanta.com>\n",
		"Route: sip:bob@biloxi.com \n",
		"Route: sip:alice@atlanta.com, sip:bob@biloxi.com, sip:carol@chicago.com\n",
	}
	var tvs_o = []string{
		"Route: <sip:alice@atlanta.com>\n",
		"Route: sip:bob@biloxi.com\n",
		"Route: sip:alice@atlanta.com,sip:bob@biloxi.com,sip:carol@chicago.com\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewRouteParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
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
