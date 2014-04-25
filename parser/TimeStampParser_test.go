package parser

import (
	"testing"
)

func TestTimeStampParser(t *testing.T) {
	var tvi = []string{
		"Timestamp: 54 \n",
		"Timestamp: 52.34 34.5 \n",
	}
	var tvo = []string{
		"Timestamp: 54 \n",
		"Timestamp: 52.34 34.5 \n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewTimeStampParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
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
