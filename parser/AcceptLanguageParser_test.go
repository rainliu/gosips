package parser

import (
	"testing"
)

func TestAcceptLanguageParser(t *testing.T) {
	var tvi = []string{
		"Accept-Language: da    \n",
		"Accept-Language: 	\n",
		"Accept-Language: da,en-gb;q=0.8\n",
		"Accept-Language: *		\n"}
	var tvo = []string{
		"Accept-Language: da    \n",
		"Accept-Language: 	\n",
		"Accept-Language: da,en-gb;q=0.8\n",
		"Accept-Language: *		\n"}

	for i := 0; i < len(tvi); i++ {
		shp := NewAcceptLanguageParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String acceptLanguage[] = {
        		"Accept-Language: da    \n",
        		"Accept-Language:\n",
        		"Accept-Language: da, en-gb;q=0.8\n",
        	        "Accept-Language: *\n" };

		for (int i =0 ; i < acceptLanguage.length; i++) {
			AcceptLanguageParser alp = new AcceptLanguageParser
				(acceptLanguage[i]);
			AcceptLanguageList all = (AcceptLanguageList)
				alp.parse();
			System.out.println(all.toString());
		}
	}
**/
