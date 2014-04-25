package parser

import (
	"testing"
)

func TestReferToParser(t *testing.T) {
	var tvi = []string{
		"Refer-To: <sip:+1-650-555-2222@ss1.wcom.com;user=phone>;tag=5617\n",
		"Refer-To: T. A. Watson <sip:watson@bell-telephone.com>\n",
		"Refer-To: LittleGuy <sip:UserB@there.com>\n",
		"Refer-To: sip:mranga@120.6.55.9\n",
		"Refer-To: sip:mranga@129.6.55.9 ; tag=696928473514.129.6.55.9\n",
	}
	var tvo = []string{
		"Refer-To: <sip:+1-650-555-2222@ss1.wcom.com;user=phone>;tag=5617\n",
		"Refer-To: \"T. A. Watson\" <sip:watson@bell-telephone.com>\n",
		"Refer-To: \"LittleGuy\" <sip:UserB@there.com>\n",
		"Refer-To: <sip:mranga@120.6.55.9>\n",
		"Refer-To: <sip:mranga@129.6.55.9>;tag=696928473514.129.6.55.9\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewReferToParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/*
   public static void main(String args[]) throws ParseException {
       String to[] = {
          "Refer-To: <sip:+1-650-555-2222@ss1.wcom.com;user=phone>;tag=5617\n",
          "Refer-To: T. A. Watson <sip:watson@bell-telephone.com>\n",
          "Refer-To: LittleGuy <sip:UserB@there.com>\n",
          "Refer-To: sip:mranga@120.6.55.9\n",
          "Refer-To: sip:mranga@129.6.55.9 ; tag=696928473514.129.6.55.9\n"
       };

       for (int i = 0; i < to.length; i++ ) {
           ReferToParser tp =
           new ReferToParser(to[i]);
           ReferTo t = (ReferTo) tp.parse();
           System.out.println("encoded = " + t.encode());
       }

   }*/
