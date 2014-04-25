package parser

import (
	"testing"
)

func TestTimeStampParser(t *testing.T) {
	var tvs = []string{
		"Timestamp: 54 \n",
		"Timestamp: 52.34 34.5 \n",
	}
	var tvs_o = []string{
		"Timestamp: 54 \n",
		"Timestamp: 52.34 34.5 \n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewTimeStampParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String timeStamp[] = {
          "Timestamp: 54 \n",
          "Timestamp: 52.34 34.5 \n"
      };

      for (int i = 0; i < timeStamp.length; i++ ) {
          TimeStampParser parser =
          new TimeStampParser(timeStamp[i]);
          TimeStamp ts= (TimeStamp) parser.parse();
          System.out.println("encoded = " + ts.encode());
      }

  }
*/
