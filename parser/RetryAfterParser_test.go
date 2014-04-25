package parser

import (
	"testing"
)

func TestRetryAfterParser(t *testing.T) {
	var tvi = []string{
		"Retry-After: 18000;duration=3600\n",
		"Retry-After: 120;duration=3600;ra=oli\n",
		"Retry-After: 1220 (I'm in a meeting)\n",
		"Retry-After: 1230 (I'm in a meeting);fg=der;duration=23\n",
	}
	var tvo = []string{
		"Retry-After: 18000;duration=3600\n",
		"Retry-After: 120;duration=3600;ra=oli\n",
		"Retry-After: 1220 (I'm in a meeting)\n",
		"Retry-After: 1230 (I'm in a meeting);fg=der;duration=23\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewRetryAfterParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String rr[] = {
          "Retry-After: 18000;duration=3600\n",
          "Retry-After: 120;duration=3600;ra=oli\n",
          "Retry-After: 1220 (I'm in a meeting)\n",
          "Retry-After: 1230 (I'm in a meeting);fg=der;duration=23\n"
      };

      for (int i = 0; i < rr.length; i++ ) {
          RetryAfterParser parser =
          new RetryAfterParser(rr[i]);
          RetryAfter r= (RetryAfter) parser.parse();
          System.out.println("encoded = " + r.encode());
      }

  }
*/
