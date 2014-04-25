package parser

import (
	"testing"
)

func TestReplyToParser(t *testing.T) {
	var tvs = []string{
		"Reply-To: Bob <sip:bob@biloxi.com>\n",
	}
	var tvs_o = []string{
		"Reply-To: \"Bob\" <sip:bob@biloxi.com>\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewReplyToParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
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
