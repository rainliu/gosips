package parser

/*
func TestStringMsgParser(t *testing.T) {
	var tvi = []string{
		"SIP/2.0 180 Ringing\r\n" +
			"Via: SIP/2.0/UDP 172.18.1.29:5060;branch=z9hG4bK43fc10fb4446d55fc5c8f969607991f4\r\n" +
			"To: \"0440\" <sip:0440@212.209.220.131>;tag=2600\r\n" +
			"From: \"Andreas\" <sip:andreas@e-horizon.se>;tag=8524\r\n" +
			"Call-ID: f51a1851c5f570606140f14c8eb64fd3@172.18.1.29\r\n" +
			"CSeq: 1 INVITE\r\n" +
			"Max-Forwards: 70\r\n" +
			"Record-Route: <sip:212.209.220.131:5060>\r\n" +
			"Content-Length: 0\r\n\r\n",

		"REGISTER sip:nist.gov SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP 129.6.55.182:14826\r\n" +
			"Max-Forwards: 70\r\n" +
			"From: <sip:mranga@nist.gov>;tag=6fcd5c7ace8b4a45acf0f0cd539b168b;epid=0d4c418ddf\r\n" +
			"To: <sip:mranga@nist.gov>\r\n" +
			"Call-ID: c5679907eb954a8da9f9dceb282d7230@129.6.55.182\r\n" +
			"CSeq: 1 REGISTER\r\n" +
			"Contact: <sip:129.6.55.182:14826>;methods=\"INVITE, MESSAGE, INFO, SUBSCRIBE, OPTIONS, BYE, CANCEL, NOTIFY, ACK, REFER\"\r\n" +
			"User-Agent: RTC/(Microsoft RTC)\r\n" +
			"Event:  registration\r\n" +
			"Allow-Events: presence\r\n" +
			"Content-Length: 0\r\n\r\n",

		"INVITE sip:littleguy@there.com:5060 SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP 65.243.118.100:5050\r\n" +
			"From: M. Ranganathan  <sip:M.Ranganathan@sipbakeoff.com>;tag=1234\r\n" +
			"To: \"littleguy@there.com\" <sip:littleguy@there.com:5060> \r\n" +
			"Call-ID: Q2AboBsaGn9!?x6@sipbakeoff.com \r\n" +
			"CSeq: 1 INVITE \r\n" +
			"Content-Length: 247\r\n\r\n" +
			"v=0\r\n" +
			"o=4855 13760799956958020 13760799956958020 IN IP4  129.6.55.78\r\n" +
			"s=mysession session\r\n" +
			"p=+46 8 52018010\r\n" +
			"c=IN IP4  129.6.55.78\r\n" +
			"t=0 0\r\n" +
			"m=audio 6022 RTP/AVP 0 4 18\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:4 G723/8000\r\n" +
			"a=rtpmap:18 G729A/8000\r\n" +
			"a=ptime:20\r\n",
	}
	var tvo = []string{
		"SIP/2.0 180 Ringing\r\n" +
			"Via: SIP/2.0/UDP 172.18.1.29:5060;branch=z9hG4bK43fc10fb4446d55fc5c8f969607991f4\r\n" +
			"To: \"0440\" <sip:0440@212.209.220.131>;tag=2600\r\n" +
			"From: \"Andreas\" <sip:andreas@e-horizon.se>;tag=8524\r\n" +
			"Call-ID: f51a1851c5f570606140f14c8eb64fd3@172.18.1.29\r\n" +
			"CSeq: 1 INVITE\r\n" +
			"Max-Forwards: 70\r\n" +
			"Record-Route: <sip:212.209.220.131:5060>\r\n" +
			"Content-Length: 0\r\n\r\n",

		"REGISTER sip:nist.gov SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP 129.6.55.182:14826\r\n" +
			"Max-Forwards: 70\r\n" +
			"From: <sip:mranga@nist.gov>;tag=6fcd5c7ace8b4a45acf0f0cd539b168b;epid=0d4c418ddf\r\n" +
			"To: <sip:mranga@nist.gov>\r\n" +
			"Call-ID: c5679907eb954a8da9f9dceb282d7230@129.6.55.182\r\n" +
			"CSeq: 1 REGISTER\r\n" +
			"Contact: <sip:129.6.55.182:14826>;methods=\"INVITE, MESSAGE, INFO, SUBSCRIBE, OPTIONS, BYE, CANCEL, NOTIFY, ACK, REFER\"\r\n" +
			"User-Agent: RTC/(Microsoft RTC)\r\n" +
			"Event: registration\r\n" +
			"Allow-Events: presence\r\n" +
			"Content-Length: 0\r\n\r\n",

		"INVITE sip:littleguy@there.com:5060 SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP 65.243.118.100:5050\r\n" +
			"From: \"M. Ranganathan\" <sip:M.Ranganathan@sipbakeoff.com>;tag=1234\r\n" +
			"To: \"littleguy@there.com\" <sip:littleguy@there.com:5060>\r\n" +
			"Call-ID: Q2AboBsaGn9!?x6@sipbakeoff.com\r\n" +
			"CSeq: 1 INVITE\r\n" +
			"Content-Length: 247\r\n\r\n" +
			"v=0\r\n" +
			"o=4855 13760799956958020 13760799956958020 IN IP4  129.6.55.78\r\n" +
			"s=mysession session\r\n" +
			"p=+46 8 52018010\r\n" +
			"c=IN IP4  129.6.55.78\r\n" +
			"t=0 0\r\n" +
			"m=audio 6022 RTP/AVP 0 4 18\r\n" +
			"a=rtpmap:0 PCMU/8000\r\n" +
			"a=rtpmap:4 G723/8000\r\n" +
			"a=rtpmap:18 G729A/8000\r\n" +
			"a=ptime:20\r\n",
	}

	for i := 0; i < len(tvi); i++ {
		smp := NewStringMsgParser()
		if sm, err := smp.ParseSIPMessage(tvi[i]); err != nil {
			t.Log(err)
			t.Fail()
		} else {
			d := sm.String()
			s := tvo[i]

			if strings.TrimSpace(d) != strings.TrimSpace(s) {
				t.Log("origin = " + s)
				t.Log("failed = " + d)

				// for j := 0; j < len(s); j++ {
				// 	if d[j] != s[j] {
				// 		t.Logf("%d:%c vs %c", j, d[j], s[j])
				// 	}
				// }

				t.Fail()
			}
		}

		//println("dialog id = " + sipMessage.GetDialogId(false))
	}
}
*/

/*
    public static void main(String[] args) throws ParseException {
        String messages[] = {
"SIP/2.0 180 Ringing\r\n"+
"Via: SIP/2.0/UDP 172.18.1.29:5060;branch=z9hG4bK43fc10fb4446d55fc5c8f969607991f4\r\n"+
"To: \"0440\" <sip:0440@212.209.220.131>;tag=2600\r\n"+
"From: \"Andreas\" <sip:andreas@e-horizon.se>;tag=8524\r\n"+
"Call-ID: f51a1851c5f570606140f14c8eb64fd3@172.18.1.29\r\n"+
"CSeq: 1 INVITE\r\n" +
"Max-Forwards: 70\r\n"+
"Record-Route: <sip:212.209.220.131:5060>\r\n"+
"Content-Length: 0\r\n\r\n",

"REGISTER sip:nist.gov SIP/2.0\r\n"+
"Via: SIP/2.0/UDP 129.6.55.182:14826\r\n"+
"Max-Forwards: 70\r\n"+
"From: <sip:mranga@nist.gov>;tag=6fcd5c7ace8b4a45acf0f0cd539b168b;epid=0d4c418ddf\r\n"+
"To: <sip:mranga@nist.gov>\r\n"+
"Call-ID: c5679907eb954a8da9f9dceb282d7230@129.6.55.182\r\n"+
"CSeq: 1 REGISTER\r\n"+
"Contact: <sip:129.6.55.182:14826>;methods=\"INVITE, MESSAGE, INFO, SUBSCRIBE, OPTIONS, BYE, CANCEL, NOTIFY, ACK, REFER\"\r\n"+
"User-Agent: RTC/(Microsoft RTC)\r\n"+
"Event:  registration\r\n"+
"Allow-Events: presence\r\n"+
"Content-Length: 0\r\n\r\n"+

    "INVITE sip:littleguy@there.com:5060 SIP/2.0\r\n"+
    "Via: SIP/2.0/UDP 65.243.118.100:5050\r\n" +
    "From: M. Ranganathan  <sip:M.Ranganathan@sipbakeoff.com>;tag=1234\r\n"+
    "To: \"littleguy@there.com\" <sip:littleguy@there.com:5060> \r\n" +
    "Call-ID: Q2AboBsaGn9!?x6@sipbakeoff.com \r\n" +
    "CSeq: 1 INVITE \r\n" +
    "Content-Length: 247\r\n\r\n"+
    "v=0\r\n"+
    "o=4855 13760799956958020 13760799956958020 IN IP4  129.6.55.78\r\n"+
    "s=mysession session\r\n"+
    "p=+46 8 52018010\r\n"+
    "c=IN IP4  129.6.55.78\r\n"+
    "t=0 0\r\n"+
    "m=audio 6022 RTP/AVP 0 4 18\r\n"+
    "a=rtpmap:0 PCMU/8000\r\n"+
    "a=rtpmap:4 G723/8000\r\n"+
    "a=rtpmap:18 G729A/8000\r\n"+
    "a=ptime:20\r\n"

        };

        for (int i = 0; i < messages.length; i++) {
            StringMsgParser smp = new StringMsgParser();
            SIPMessage sipMessage = smp.parseSIPMessage(messages[i]);
            System.out.println("encoded " + sipMessage.toString());
	    System.out.println("dialog id = " + sipMessage.GetDialogId(false));
        }


    }*/
