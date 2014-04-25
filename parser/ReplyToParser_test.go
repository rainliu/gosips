package parser

import (
	"testing"
)

func TestReplyToParser(t *testing.T) {
	var tvi = []string{
		"Reply-To: Bob <sip:bob@biloxi.com>\n",
	}
	var tvo = []string{
		"Reply-To: \"Bob\" <sip:bob@biloxi.com>\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewReplyToParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/**
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Reply-To: Bob <sip:bob@biloxi.com>\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ReplyToParser rt =
          new ReplyToParser(r[i]);
          ReplyTo re = (ReplyTo) rt.parse();
          System.out.println("encoded = " +re.encode());
      }

  }
*/
