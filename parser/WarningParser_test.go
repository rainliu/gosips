package parser

import (
	"testing"
)

func TestWarningParser(t *testing.T) {
	var tvs = []string{
		"Warning: 307 isi.edu \"Session parameter 'foo' not understood\"\n",
		"Warning: 301 isi.edu \"Incompatible network address type 'E.164'\"\n",
		"Warning: 312 ii.edu \"Soda\", " +
			" 351 i.edu \"I network address 'E.164'\" , 323 ii.edu \"Sodwea\"\n",
	}
	var tvs_o = []string{
		"Warning: 307 isi.edu \"Session parameter 'foo' not understood\"\n",
		"Warning: 301 isi.edu \"Incompatible network address type 'E.164'\"\n",
		"Warning: 312 ii.edu \"Soda\"," +
			"351 i.edu \"I network address 'E.164'\",323 ii.edu \"Sodwea\"\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewWarningParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String warning[] = {
                "Warning: 307 isi.edu \"Session parameter 'foo' not understood\"\n",
                "Warning: 301 isi.edu \"Incompatible network address type 'E.164'\"\n",
                "Warning: 312 ii.edu \"Soda\", "+
                " 351 i.edu \"I network address 'E.164'\" , 323 ii.edu \"Sodwea\"\n"
                };

		for (int i = 0; i < warning.length; i++ ) {
		    WarningParser parser =
			  new WarningParser(warning[i]);
		    WarningList warningList= (WarningList) parser.parse();
		    System.out.println("encoded = " + warningList.encode());
		}

	}
*/
