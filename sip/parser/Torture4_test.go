/*%% -------------------------------------------------------------------
%%
%% torture_test: RFC4475 "Invalid" tests (3.1.2.1 to 3.1.2.19)
%%
%% Copyright (c) 2013 Carlos Gonzalez Florido.  All Rights Reserved.
%%
%% This file is provided to you under the Apache License,
%% Version 2.0 (the "License"); you may not use this file
%% except in compliance with the License.  You may obtain
%% a copy of the License at
%%
%%   http://www.apache.org/licenses/LICENSE-2.0
%%
%% Unless required by applicable law or agreed to in writing,
%% software distributed under the License is distributed on an
%% "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
%% KIND, either express or implied.  See the License for the
%% specific language governing permissions and limitations
%% under the License.
%%
%% -------------------------------------------------------------------*/

package parser

import (
	"strings"
	"testing"
)

func TestTorture4(t *testing.T) {
	tvi := torture4_i
	tvo := torture4_o

	for i := 0; i < 5; /*len(tvi);*/ i++ {
		hp := CreateParser(tvi[i])
		if sh, err := hp.Parse(); err != nil {
			t.Log(tvo[i])
			if strings.Contains(tvo[i], "Invalid:") {
				t.Log(err)
			} else {
				t.Fail()
			}
		} else {
			d := sh.String()
			s := tvo[i]

			if strings.TrimSpace(d) != strings.TrimSpace(s) {
				t.Log("golden = " + s)
				t.Log("failed = " + d)

				for j, k := 0, 0; j < len(s); j++ {
					if d[j] != s[j] {
						t.Logf("%d:%c vs %c", j, s[j], d[j])
						k++
						if k == 10 {
							break
						}
					}
				}

				t.Fail()
			}
		}
	}
}

var torture4_i = []string{
	"To: sip:j.user@example.com\n",
	"From: sip:caller@example.net;tag=134161461246\n",
	"Max-Forwards: 7\n",
	"Call-ID: badinv01.0ha0isndaksdjasdf3234nas\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.15;;,;,,\n",
	"Contact: \"Joe\" <sip:joe@example.org>;;;;\n",
	"Content-Length: 152\n",
	"Content-Type: application/sdp\n",

	"Max-Forwards: 80\n",
	"To: sip:j.user@example.com\n",
	"From: sip:caller@example.net;tag=93942939o2\n",
	"Contact: <sip:caller@hungry.example.net>\n",
	"Call-ID: clerr.0ha0isndaksdjweiafasdk3\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP host5.example.com;branch=z9hG4bK-39234-23523\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 9999\n",

	"Max-Forwards: 254\n",
	"To: sip:j.user@example.com\n",
	"From: sip:caller@example.net;tag=32394234\n",
	"Call-ID: ncl.0ha0isndaksdj2193423r542w35\n",
	"CSeq: 0 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.53;branch=z9hG4bKkdjuw\n",
	"Contact: <sip:caller@example53.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: -999\n",

	"Via: SIP/2.0/TCP host129.example.com;branch=z9hG4bK342sdfoi3\n",
	"To: <sip:user@example.com>\n",
	"From: <sip:user@example.com>;tag=239232jh3\n",
	"CSeq: 36893488147419103232 REGISTER\n",
	"Call-ID: scalar02.23o0pd9vanlq3wnrlnewofjas9ui32\n",
	"Max-Forwards: 300\n",
	"Expires: 1000000000000000000000000000000000000000000000000</repeat>\n",
	"Contact: <sip:user@host129.example.com>;expires=280297596632815\n",
	"Content-Length: 0\n",

	"Via: SIP/2.0/TCP host129.example.com;branch=z9hG4bKzzxdiwo34sw;received=192.0.2.129\n",
	"To: <sip:user@example.com>\n",
	"From: <sip:other@example.net>;tag=2easdjfejw\n",
	"CSeq: 9292394834772304023312 OPTIONS\n",
	"Call-ID: scalarlg.noase0of0234hn2qofoaf0232aewf2394r\n",
	"Retry-After: 949302838503028349304023988\n",
	"Warning: 1812 overture \"In Progress\"\n",
	"Content-Length: 0\n",

	"To: \"Mr. J. User <sip:j.user@example.com>\n",
	"From: sip:caller@example.net;tag=93334\n",
	"Max-Forwards: 10\n",
	"Call-ID: quotbal.aksdj\n",
	"Contact: <sip:caller@host59.example.net>\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.59:5050;branch=z9hG4bKkdjuw39234\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 152\n",

	"To: sip:user@example.com\n",
	"From: sip:caller@example.net;tag=39291\n",
	"Max-Forwards: 23\n",
	"Call-ID: ltgtruri.1@192.0.2.5\n",
	"CSeq: 1 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.5\n",
	"Contact: <sip:caller@host5.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 159\n",

	"To: sip:user@example.com;tag=3xfe-9921883-z9f\n",
	"From: sip:caller@example.net;tag=231413434\n",
	"Max-Forwards: 5\n",
	"Call-ID: lwsruri.asdfasdoeoi2323-asdfwrn23-asd834rk423\n",
	"CSeq: 2130706432 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.1:5060;branch=z9hG4bKkdjuw2395\n",
	"Contact: <sip:caller@host1.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 159\n",

	"Max-Forwards: 8\n",
	"To: sip:user@example.com\n",
	"From: sip:caller@example.net;tag=8814\n",
	"Call-ID: lwsstart.dfknq234oi243099adsdfnawe3@example.com\n",
	"CSeq: 1893884 INVITE\n",
	"Via: SIP/2.0/UDP host1.example.com;branch=z9hG4bKkdjuw3923\n",
	"Contact: <sip:caller@host1.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 150\n",

	"Via: SIP/2.0/TCP host1.example.com;branch=z9hG4bK299342093\n",
	"To: <sip:remote-target@example.com>\n",
	"From: <sip:local-resource@example.com>;tag=329429089\n",
	"Call-ID: trws.oicu34958239neffasdhr2345r\n",
	"Accept: application/sdp\n",
	"CSeq: 238923 OPTIONS\n",
	"Max-Forwards: 70\n",
	"Content-Length: 0\n",

	"To: sip:user@example.com\n",
	"From: sip:caller@example.net;tag=341518\n",
	"Max-Forwards: 7\n",
	"Contact: <sip:caller@host39923.example.net>\n",
	"Call-ID: escruri.23940-asdfhj-aje3br-234q098w-fawerh2q-h4n5\n",
	"CSeq: 149209342 INVITE\n",
	"Via: SIP/2.0/UDP host-of-the-hour.example.com;branch=z9hG4bKkdjuw\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 150\n",

	"To: sip:user@example.com\n",
	"From: sip:caller@example.net;tag=2234923\n",
	"Max-Forwards: 70\n",
	"Call-ID: baddate.239423mnsadf3j23lj42--sedfnm234\n",
	"CSeq: 1392934 INVITE\n",
	"Via: SIP/2.0/UDP host.example.com;branch=z9hG4bKkdjuw\n",
	"Date: Fri, 01 Jan 2010 16:00:00 EST\n",
	"Contact: <sip:caller@host5.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 150\n",

	"To: sip:user@example.com\n",
	"From: sip:user@example.com;tag=998332\n",
	"Max-Forwards: 70\n",
	"Call-ID: regbadct.k345asrl3fdbv@10.0.0.1\n",
	"CSeq: 1 REGISTER\n",
	"Via: SIP/2.0/UDP 135.180.130.133:5060;branch=z9hG4bKkdjuw\n",
	"Contact: sip:user@example.com?Route=%3Csip:sip.example.com%3E\n",
	"l: 0\n",

	"Via: SIP/2.0/UDP host4.example.com:5060;branch=z9hG4bKkdju43234\n",
	"Max-Forwards: 70\n",
	"From: \"Bell, Alexander\" <sip:a.g.bell@example.com>;tag=433423\n",
	"To: \"Watson, Thomas\" < sip:t.watson@example.org >\n",
	"Call-ID: badaspec.sdf0234n2nds0a099u23h3hnnw009cdkne3\n",
	"Accept: application/sdp\n",
	"CSeq: 3923239 OPTIONS\n",
	"l: 0\n",

	"Via:     SIP/2.0/UDP c.example.com:5060;branch=z9hG4bKkdjuw\n",
	"Max-Forwards:      70\n",
	"From:    Bell, Alexander <sip:a.g.bell@example.com>;tag=43\n",
	"To:      Watson, Thomas <sip:t.watson@example.org>\n",
	"Call-ID: baddn.31415@c.example.com\n",
	"Accept: application/sdp\n",
	"CSeq:    3923239 OPTIONS\n",
	"l: 0\n",

	"Via:     SIP/7.0/UDP c.example.com;branch=z9hG4bKkdjuw\n",
	"Max-Forwards:     70\n",
	"From:    A. Bell <sip:a.g.bell@example.com>;tag=qweoiqpe\n",
	"To:      T. Watson <sip:t.watson@example.org>\n",
	"Call-ID: badvers.31417@c.example.com\n",
	"CSeq:    1 OPTIONS\n",
	"l: 0\n",

	"To: sip:j.user@example.com\n",
	"From: sip:caller@example.net;tag=34525\n",
	"Max-Forwards: 6\n",
	"Call-ID: mismatch01.dj0234sxdfl3\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP host.example.com;branch=z9hG4bKkdjuw\n",
	"l: 0\n",

	"To: sip:j.user@example.com\n",
	"From: sip:caller@example.net;tag=34525\n",
	"Max-Forwards: 6\n",
	"Call-ID: mismatch02.dj0234sxdfl3\n",
	"CSeq: 8 INVITE\n",
	"Contact: <sip:caller@host.example.net>\n",
	"Via: SIP/2.0/UDP host.example.net;branch=z9hG4bKkdjuw\n",
	"Content-Type: application/sdp\n",
	"l: 138\n",

	"Via: SIP/2.0/UDP 192.0.2.105;branch=z9hG4bK2398ndaoe\n",
	"Call-ID: bigcode.asdof3uj203asdnf3429uasdhfas3ehjasdfas9i\n",
	"CSeq: 353494 INVITE\n",
	"From: <sip:user@example.com>;tag=39ansfi3\n",
	"To: <sip:user@example.edu>;tag=902jndnke3\n",
	"Content-Length: 0\n",
	"Contact: <sip:user@host105.example.com>\n",
}

var torture4_o = []string{
	"To: <sip:j.user@example.com>\n",
	"From: <sip:caller@example.net>;tag=134161461246\n",
	"Max-Forwards: 7\n",
	"Call-ID: badinv01.0ha0isndaksdjasdf3234nas\n",
	"CSeq: 8 INVITE\n",
	"Invalid: Via: SIP/2.0/UDP 192.0.2.15;;,;,,\n",
	"Invalid: Contact: \"Joe\" <sip:joe@example.org>;;;;\n",
	"Content-Length: 152\n",
	"Content-Type: application/sdp\n",

	"Max-Forwards: 80\n",
	"To: <sip:j.user@example.com>\n",
	"From: <sip:caller@example.net>;tag=93942939o2\n",
	"Contact: <sip:caller@hungry.example.net>\n",
	"Call-ID: clerr.0ha0isndaksdjweiafasdk3\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP host5.example.com;branch=z9hG4bK-39234-23523\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 9999\n",

	"Max-Forwards: 254\n",
	"To: <sip:j.user@example.com>\n",
	"From: <sip:caller@example.net>;tag=32394234\n",
	"Call-ID: ncl.0ha0isndaksdj2193423r542w35\n",
	"CSeq: 0 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.53;branch=z9hG4bKkdjuw\n",
	"Contact: <sip:caller@example53.example.net>\n",
	"Content-Type: application/sdp\n",
	"Invalid: Content-Length: -999\n",

	"Via: SIP/2.0/TCP host129.example.com;branch=z9hG4bK342sdfoi3\n",
	"To: <sip:user@example.com>\n",
	"From: <sip:user@example.com>;tag=239232jh3\n",
	"Invalid: CSeq: 36893488147419103232 REGISTER\n",
	"Call-ID: scalar02.23o0pd9vanlq3wnrlnewofjas9ui32\n",
	"Invalid: Max-Forwards: 300\n",
	"Invalid: Expires: 1000000000000000000000000000000000000000000000000</repeat>\n",
	"Invalid: Contact: <sip:user@host129.example.com>;expires=280297596632815\n",
	"Content-Length: 0\n",

	"Via: SIP/2.0/TCP host129.example.com;branch=z9hG4bKzzxdiwo34sw;received=192.0.2.129\n",
	"To: <sip:user@example.com>\n",
	"From: <sip:other@example.net>;tag=2easdjfejw\n",
	"Invalid: CSeq: 9292394834772304023312 OPTIONS\n",
	"Call-ID: scalarlg.noase0of0234hn2qofoaf0232aewf2394r\n",
	"Invalid: Retry-After: 949302838503028349304023988\n",
	"Invalid: Warning: 1812 overture \"In Progress\"\n",
	"Content-Length: 0\n",

	"Invalid: To: \"Mr. J. User <sip:j.user@example.com>\n",
	"From: <sip:caller@example.net>;tag=93334\n",
	"Max-Forwards: 10\n",
	"Call-ID: quotbal.aksdj\n",
	"Contact: <sip:caller@host59.example.net>\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.59:5050;branch=z9hG4bKkdjuw39234\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 152\n",

	"To: <sip:user@example.com>\n",
	"From: <sip:caller@example.net>;tag=39291\n",
	"Max-Forwards: 23\n",
	"Call-ID: ltgtruri.1@192.0.2.5\n",
	"CSeq: 1 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.5\n",
	"Contact: <sip:caller@host5.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 159\n",

	"To: <sip:user@example.com>;tag=3xfe-9921883-z9f\n",
	"From: <sip:caller@example.net>;tag=231413434\n",
	"Max-Forwards: 5\n",
	"Call-ID: lwsruri.asdfasdoeoi2323-asdfwrn23-asd834rk423\n",
	"CSeq: 2130706432 INVITE\n",
	"Via: SIP/2.0/UDP 192.0.2.1:5060;branch=z9hG4bKkdjuw2395\n",
	"Contact: <sip:caller@host1.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 159\n",

	"Max-Forwards: 8\n",
	"To: <sip:user@example.com>\n",
	"From: <sip:caller@example.net>;tag=8814\n",
	"Call-ID: lwsstart.dfknq234oi243099adsdfnawe3@example.com\n",
	"CSeq: 1893884 INVITE\n",
	"Via: SIP/2.0/UDP host1.example.com;branch=z9hG4bKkdjuw3923\n",
	"Contact: <sip:caller@host1.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 150\n",

	"Via: SIP/2.0/TCP host1.example.com;branch=z9hG4bK299342093\n",
	"To: <sip:remote-target@example.com>\n",
	"From: <sip:local-resource@example.com>;tag=329429089\n",
	"Call-ID: trws.oicu34958239neffasdhr2345r\n",
	"Accept: application/sdp\n",
	"CSeq: 238923 OPTIONS\n",
	"Max-Forwards: 70\n",
	"Content-Length: 0\n",

	"To: <sip:user@example.com>\n",
	"From: <sip:caller@example.net>;tag=341518\n",
	"Max-Forwards: 7\n",
	"Contact: <sip:caller@host39923.example.net>\n",
	"Call-ID: escruri.23940-asdfhj-aje3br-234q098w-fawerh2q-h4n5\n",
	"CSeq: 149209342 INVITE\n",
	"Via: SIP/2.0/UDP host-of-the-hour.example.com;branch=z9hG4bKkdjuw\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 150\n",

	"To: <sip:user@example.com>\n",
	"From: <sip:caller@example.net>;tag=2234923\n",
	"Max-Forwards: 70\n",
	"Call-ID: baddate.239423mnsadf3j23lj42--sedfnm234\n",
	"CSeq: 1392934 INVITE\n",
	"Via: SIP/2.0/UDP host.example.com;branch=z9hG4bKkdjuw\n",
	"Invalid: Date: Fri, 01 Jan 2010 16:00:00 EST\n",
	"Contact: <sip:caller@host5.example.net>\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 150\n",

	"To: <sip:user@example.com>\n",
	"From: <sip:user@example.com>;tag=998332\n",
	"Max-Forwards: 70\n",
	"Call-ID: regbadct.k345asrl3fdbv@10.0.0.1\n",
	"CSeq: 1 REGISTER\n",
	"Via: SIP/2.0/UDP 135.180.130.133:5060;branch=z9hG4bKkdjuw\n",
	"Invalid: Contact: sip:user@example.com?Route=%3Csip:sip.example.com%3E\n",
	"Content-Length: 0\n",

	"Via: SIP/2.0/UDP host4.example.com:5060;branch=z9hG4bKkdju43234\n",
	"Max-Forwards: 70\n",
	"From: \"Bell, Alexander\" <sip:a.g.bell@example.com>;tag=433423\n",
	"Invalid: To: \"Watson, Thomas\" < sip:t.watson@example.org >\n",
	"Call-ID: badaspec.sdf0234n2nds0a099u23h3hnnw009cdkne3\n",
	"Accept: application/sdp\n",
	"CSeq: 3923239 OPTIONS\n",
	"Content-Length: 0\n",

	"Via: SIP/2.0/UDP c.example.com:5060;branch=z9hG4bKkdjuw\n",
	"Max-Forwards: 70\n",
	"Invalid: From:    Bell, Alexander <sip:a.g.bell@example.com>;tag=43\n",
	"Invalid: To:      Watson, Thomas <sip:t.watson@example.org>\n",
	"Call-ID: baddn.31415@c.example.com\n",
	"Accept: application/sdp\n",
	"CSeq: 3923239 OPTIONS\n",
	"Content-Length: 0\n",

	"Invalid: Via: SIP/7.0/UDP c.example.com;branch=z9hG4bKkdjuw\n",
	"Max-Forwards: 70\n",
	"From: \"A. Bell\" <sip:a.g.bell@example.com>;tag=qweoiqpe\n",
	"To: \"T. Watson\" <sip:t.watson@example.org>\n",
	"Call-ID: badvers.31417@c.example.com\n",
	"CSeq: 1 OPTIONS\n",
	"Content-Length: 0\n",

	"To: <sip:j.user@example.com>\n",
	"From: <sip:caller@example.net>;tag=34525\n",
	"Max-Forwards: 6\n",
	"Call-ID: mismatch01.dj0234sxdfl3\n",
	"CSeq: 8 INVITE\n",
	"Via: SIP/2.0/UDP host.example.com;branch=z9hG4bKkdjuw\n",
	"Content-Length: 0\n",

	"To: <sip:j.user@example.com>\n",
	"From: <sip:caller@example.net>;tag=34525\n",
	"Max-Forwards: 6\n",
	"Call-ID: mismatch02.dj0234sxdfl3\n",
	"CSeq: 8 INVITE\n",
	"Contact: <sip:caller@host.example.net>\n",
	"Via: SIP/2.0/UDP host.example.net;branch=z9hG4bKkdjuw\n",
	"Content-Type: application/sdp\n",
	"Content-Length: 138\n",

	"Via: SIP/2.0/UDP 192.0.2.105;branch=z9hG4bK2398ndaoe\n",
	"Call-ID: bigcode.asdof3uj203asdnf3429uasdhfas3ehjasdfas9i\n",
	"CSeq: 353494 INVITE\n",
	"From: <sip:user@example.com>;tag=39ansfi3\n",
	"To: <sip:user@example.edu>;tag=902jndnke3\n",
	"Content-Length: 0\n",
	"Contact: <sip:user@host105.example.com>\n",
}
