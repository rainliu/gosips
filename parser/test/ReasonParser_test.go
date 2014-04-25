package parser

import (
	"testing"
)

func TestReasonParser(t *testing.T) {
	var tvs = []string{
		"Reason: SIP ;cause=200 ;text=\"Call completed elsewhere\"\n",
		"Reason: Q.850 ;cause=16 ;text=\"Terminated\"\n",
		"Reason: SIP ;cause=600 ;text=\"Busy Everywhere\"\n",
		"Reason: SIP ;cause=580 ;text=\"Precondition Failure\"," +
			"SIP ;cause=530 ;text=\"Pre Failure\"\n",
		"Reason: SIP \n",
	}
	var tvs_o = []string{
		"Reason: SIP;cause=200;text=\"Call completed elsewhere\"\n",
		"Reason: Q.850;cause=16;text=\"Terminated\"\n",
		"Reason: SIP;cause=600;text=\"Busy Everywhere\"\n",
		"Reason: SIP;cause=580;text=\"Precondition Failure\"," +
			"SIP;cause=530;text=\"Pre Failure\"\n",
		"Reason: SIP \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewReasonParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
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
