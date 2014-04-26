package parser

import (
	"strings"
	"testing"
)

func TestRequestLineParser(t *testing.T) {
	var tvi = []string{
		"REGISTER sip:company.com SIP/2.0\n",
		"INVITE sip:3660@166.35.231.140 SIP/2.0\n",
		"INVITE sip:user@company.com SIP/2.0\n",
		"OPTIONS sip:135.180.130.133 SIP/2.0\n",
	}
	var tvo = []string{
		"REGISTER sip:company.com SIP/2.0\n",
		"INVITE sip:3660@166.35.231.140 SIP/2.0\n",
		"INVITE sip:user@company.com SIP/2.0\n",
		"OPTIONS sip:135.180.130.133 SIP/2.0\n",
	}

	for i := 0; i < len(tvi); i++ {
		rlp := NewRequestLineParser(tvi[i])
		testRequestLineParser(t, rlp, tvo[i])
	}
}

func testRequestLineParser(t *testing.T, rlp *RequestLineParser, s string) {
	if rl, err := rlp.Parse(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		d := rl.String()

		if strings.TrimSpace(d) != strings.TrimSpace(s) {
			t.Log("origin = " + s)
			t.Log("failed = " + d)
			t.Fail()
		} /*else {
			t.Log("passed = " + d)
		}*/
	}
}

/**
public static void main(String args[]) throws ParseException {
	String requestLines[] = {
		"REGISTER sip:company.com SIP/2.0\n",
		"INVITE sip:3660@166.35.231.140 SIP/2.0\n",
		"INVITE sip:user@company.com SIP/2.0\n",
		"OPTIONS sip:135.180.130.133 SIP/2.0\n" };
	for (int i = 0; i < requestLines.length; i++ ) {
	    RequestLineParser rlp =
		  new RequestLineParser(requestLines[i]);
	    RequestLine rl = rlp.parse();
	    System.out.println("encoded = " + rl.encode());
	}

}
*/
