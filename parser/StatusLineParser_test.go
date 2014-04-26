package parser

import (
	"strings"
	"testing"
)

func TestStatusLineParser(t *testing.T) {
	var tvi = []string{
		"SIP/2.0 200 OK\n",
		//"BOO 200 OK\n",
		"SIP/2.0 500 OK bad things happened \n",
	}
	var tvo = []string{
		"SIP/2.0 200 OK\n",
		//"BOO 200 OK\n",
		"SIP/2.0 500 OK bad things happened \n",
	}

	for i := 0; i < len(tvi); i++ {
		rlp := NewStatusLineParser(tvi[i])
		testStatusLineParser(t, rlp, tvo[i])
	}
}

func testStatusLineParser(t *testing.T, rlp *StatusLineParser, s string) {
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
public static void main(String[] args)  throws ParseException {
	String[] statusLines = {
	 "SIP/2.0 200 OK\n",
	 "BOO 200 OK\n",
	 "SIP/2.0 500 OK bad things happened \n"
	};
	for (int i = 0 ; i < statusLines.length; i++) {
	   try {
	   StatusLineParser slp = new StatusLineParser(statusLines[i]);
	   StatusLine sl = slp.parse();
	   System.out.println("encoded = " + sl.encode());
	   } catch (ParseException ex) {
		System.out.println("error message " + ex.getMessage());
	   }
	}
}
*/
