package parser

import (
	"testing"
)

func TestProxyRequireParser(t *testing.T) {
	var tvi = []string{
		"Proxy-Require: foo \n",
		"Proxy-Require: foo1,foo2,389\n",
	}
	var tvo = []string{
		"Proxy-Require: foo \n",
		"Proxy-Require: foo1,foo2,389\n",
	}

	for i := 0; i < len(tvi); i++ {
		shp := NewProxyRequireParser(tvi[i])
		testHeaderParser(t, shp, tvo[i])
	}
}

/** Test program
  public static void main(String args[]) throws ParseException {
      String r[] = {
          "Proxy-Require: foo \n",
          "Proxy-Require: foo1, foo2 , 389\n"
      };

      for (int i = 0; i < r.length; i++ ) {
          ProxyRequireParser parser =
          new ProxyRequireParser(r[i]);
          ProxyRequireList rl= (ProxyRequireList) parser.parse();
          System.out.println("encoded = " + rl.encode());
      }
  }
*/
