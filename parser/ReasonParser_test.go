package parser

import (
	"testing"
)

func TestReasonParser(t *testing.T) {
	var tvi = []string{
		"Reason: SIP ;cause=200 ;text=\"Call completed elsewhere\"\n",
		"Reason: Q.850 ;cause=16 ;text=\"Terminated\"\n",
		"Reason: SIP ;cause=600 ;text=\"Busy Everywhere\"\n",
		"Reason: SIP ;cause=580 ;text=\"Precondition Failure\"," +
			"SIP ;cause=530 ;text=\"Pre Failure\"\n",
		"Reason: SIP \n",
	}
	var tvo = []string{
		"Reason: SIP;cause=200;text=\"Call completed elsewhere\"\n",
		"Reason: Q.850;cause=16;text=\"Terminated\"\n",
		"Reason: SIP;cause=600;text=\"Busy Everywhere\"\n",
		"Reason: SIP;cause=580;text=\"Precondition Failure\"," +
			"SIP;cause=530;text=\"Pre Failure\"\n",
		"Reason: SIP \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewReasonParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Reason: SIP ;cause=200 ;text=\"Call completed elsewhere\"\n",
          "Reason: Q.850 ;cause=16 ;text=\"Terminated\"\n",
          "Reason: SIP ;cause=600 ;text=\"Busy Everywhere\"\n",
          "Reason: SIP ;cause=580 ;text=\"Precondition Failure\","+
          "SIP ;cause=530 ;text=\"Pre Failure\"\n",
          "Reason: SIP \n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ReasonParser parser =
          new ReasonParser(r[i]);
          ReasonList rl= (ReasonList) parser.parse();
          System.out.println("encoded = " + rl.encode());
      }
  }
*/
