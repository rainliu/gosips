package parser

import (
	"strings"
	"testing"
)

func TestRequestLineParser(t *testing.T) {
	var tvi = []string{
		"!interesting-Method0123456789_*+`.%indeed'~ sip:1_unusual.URI~(to-be!sure)&isn't+it$/crazy?,/;;*:&it+has=1,weird!*pas$wo~d_too.(doesn't-it)@example.com SIP/2.0\n",
		"REGISTER sip:company.com SIP/2.0\n",
		"INVITE sip:3660@166.35.231.140 SIP/2.0\n",
		"INVITE sip:user@company.com SIP/2.0\n",
		"OPTIONS sip:135.180.130.133 SIP/2.0\n",
	}
	var tvo = []string{
		"!interesting-Method0123456789_*+`.%indeed'~ sip:1_unusual.URI~(to-be!sure)&isn't+it$/crazy?,/;;*:&it+has=1,weird!*pas$wo~d_too.(doesn't-it)@example.com SIP/2.0\n",
		"REGISTER sip:company.com SIP/2.0\n",
		"INVITE sip:3660@166.35.231.140 SIP/2.0\n",
		"INVITE sip:user@company.com SIP/2.0\n",
		"OPTIONS sip:135.180.130.133 SIP/2.0\n",
	}

	for i := 0; i < 1; /*len(tvi)*/ i++ {
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
			for j, k := 0, 0; j < len(s); j++ {
				if d[j] != s[j] {
					t.Logf("%d:%c vs %c", j, d[j], s[j])
					k++
					if k == 10 {
						break
					}
				}
			}
		}
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
