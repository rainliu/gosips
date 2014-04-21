package parser

import (
	"testing"
)

func TestContentLengthParser(t *testing.T) {
	var contentLength = []string{
		"l: 345\n",
		"Content-Length: 3495\n",
		"Content-Length: 0 \n",
	}

	for i := 0; i < len(contentLength); i++ {
		shp := NewContentLengthParser(contentLength[i])
		testHeaderParser(t, shp)
	}
}

/**
        public static void main(String args[]) throws ParseException {
		String content[] = {
			"l: 345\n",
			"Content-Length: 3495\n",
			"Content-Length: 0 \n"
                };

		for (int i = 0; i < content.length; i++ ) {
		    ContentLengthParser cp =
			  new ContentLengthParser(content[i]);
		    ContentLength c = (ContentLength) cp.parse();
		    System.out.println("encoded = " + c.encode());
		}

	}
**/
