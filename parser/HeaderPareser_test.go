package parser

import (
	"testing"
)
	func testHeaderParser(t *testing.T, hp HeaderParser){
			if sh, err := hp.Parse(); err!=nil{
            	t.Log(err);
            	t.Fail();
            }else{
            	t.Log("encoded = " + sh.String());
            }
	}

	func TestCallIdParser(t *testing.T) {
        var call = []string{
	    	"Call-ID :  f0b40bcc-3485-49e7-ad1a-f1dfad2e39c9@10.5.0.53\n",
            "Call-ID	: 	f81d4fae-7dec-11d0-a765-00a0c91e6bf6@foo.bar.com\n",
            "i:f81d4fae-7dec-11d0-a765-00a0c91e6bf6@foo.bar.com\n",
            "Call-ID:	1@10.0.0.1\n",
            "Call-ID	:kl24ahsd546folnyt2vbak9sad98u23naodiunzds09a3bqw0sdfbsk34poouymnae0043nsed09mfkvc74bd0cuwnms05dknw87hjpobd76f\n",
            "Call-Id: 281794\n",
        }
        
        for i := 0; i < len(call); i++  {
            shp := NewCallIDParser(call[i]);
            testHeaderParser(t, shp)
        }
    }

     func TestCSeqParser(t *testing.T) {
		var cseq =[]string{
			"CSeq: 17 INVITE\n",
			"CSeq: 17 ACK\n",
			"CSeq : 18   BYE\n",
            "CSeq:1 CANCEL\n",
            "CSeq: 3 BYE\n",
        };
			
		for i := 0; i < len(cseq); i++  {
		    shp := NewCSeqParser(cseq[i]);
		    testHeaderParser(t, shp);
		}	
	}
	
	func TestAddressParser(t *testing.T) {
	     var addresses =[]string { 
		"<sip:user@example.com?Route=%3csip:sip.example.com%3e>",
		"\"M. Ranganathan\"   <sip:mranga@nist.gov>",
		"<sip:+1-650-555-2222@ss1.wcom.com;user=phone>",
		"M. Ranganathan <sip:mranga@nist.gov>" };

	    for i := 0; i < len(addresses); i++ {
        	addressParser := NewAddressParser(addresses[i]);
			if addr, err := addressParser.Address(); err!=nil{
            	t.Log(err);
            	t.Fail();
            }else{
            	t.Log("encoded = " + addr.String());
            }
	    }
	}


/**
static private String urls[] = 
{ 
   "sip:conference=1234@sip.convedia.com;xyz=pqd",
   "sip:herbivore.ncsl.nist.gov:5070;maddr=129.6.55.251;lc",
  "sip:1-301-975-3664@foo.bar.com;user=phone", "sip:129.6.55.181",
  "sip:herbivore.ncsl.nist.gov:5070;maddr=129.6.55.251?method=INVITE&contact=sip:foo.bar.com",
  "sip:j.doe@big.com", 
  "sip:j.doe:secret@big.com;transport=tcp",
  "sip:j.doe@big.com?subject=project",  
  "sip:+1-212-555-1212:1234@gateway.com;user=phone" ,
  "sip:1212@gateway.com",
  "sip:alice@10.1.2.3",
  "sip:alice@example.com",
  "sip:alice",
  "sip:alice@registrar.com;method=REGISTER",
  "sip:annc@10.10.30.186:6666;early=no;play=http://10.10.30.186:8080/examples/pin.vxml",
"tel:+463-1701-4291" ,
"tel:46317014291" ,
"http://10.10.30.186:8080/examples/pin.vxml" 
};

    public static void main(String[] args) {
        
        try {
            for (int i = 0; i < urls.length; i++) {
                
                String url = urls[i];
                System.out.println("URI = " + url);
                URLParser urlParser = new URLParser(url);
                GenericURI uri = urlParser.parse();
		System.out.println("class = " + uri.getClass());
                System.out.println("encoded URI = " + uri.toString());
                System.out.println("cloned encoded URI = " + 
				uri.clone().toString());
            }
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        
    }
**/
