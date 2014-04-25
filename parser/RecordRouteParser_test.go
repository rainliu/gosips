package parser

import (
	"testing"
)

func TestRecordRouteParser(t *testing.T) {
	var tvs = []string{
		"Record-Route: <sip:bob@biloxi.com;maddr=10.1.1.1>," +
			"<sip:bob@biloxi.com;maddr=10.2.1.1>\n",
		"Record-Route: <sip:UserB@there.com;maddr=ss2.wcom.com>\n",
		"Record-Route: <sip:+1-650-555-2222@iftgw.there.com;" +
			"maddr=ss1.wcom.com>\n",
		"Record-Route: <sip:UserB@there.com;maddr=ss2.wcom.com>," +
			"<sip:UserB@there.com;maddr=ss1.wcom.com>\n",
	}
	var tvs_o = []string{
		"Record-Route: <sip:bob@biloxi.com;maddr=10.1.1.1>," +
			"<sip:bob@biloxi.com;maddr=10.2.1.1>\n",
		"Record-Route: <sip:UserB@there.com;maddr=ss2.wcom.com>\n",
		"Record-Route: <sip:+1-650-555-2222@iftgw.there.com;" +
			"maddr=ss1.wcom.com>\n",
		"Record-Route: <sip:UserB@there.com;maddr=ss2.wcom.com>," +
			"<sip:UserB@there.com;maddr=ss1.wcom.com>\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewRecordRouteParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String rou[] = {
			"Record-Route: <sip:bob@biloxi.com;maddr=10.1.1.1>,"+
                        "<sip:bob@biloxi.com;maddr=10.2.1.1>\n",

			"Record-Route: <sip:UserB@there.com;maddr=ss2.wcom.com>\n",

                        "Record-Route: <sip:+1-650-555-2222@iftgw.there.com;"+
                        "maddr=ss1.wcom.com>\n",

                        "Record-Route: <sip:UserB@there.com;maddr=ss2.wcom.com>,"+
                        "<sip:UserB@there.com;maddr=ss1.wcom.com>\n"
                };

		for (int i = 0; i < rou.length; i++ ) {
		    RecordRouteParser rp =
			  new RecordRouteParser(rou[i]);
		    RecordRouteList recordRouteList = (RecordRouteList) rp.parse();
		    System.out.println("encoded = " +recordRouteList.encode());
		}

	}
*/
