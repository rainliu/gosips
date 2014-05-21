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

	for i := 0; i < 13; /*len(tvi)*/ i++ {
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
	"Contact: \"Joe\" <sip:joe@example.org>;;;;\n",
	"Content-Length: -999\n",
	"CSeq: 36893488147419103232 REGISTER\n",
	"CSeq: 9292394834772304023312 OPTIONS\n",
	"Contact: <sip:user@host129.example.com>;expires=280297596632815\n",
	"Date: Fri, 01 Jan 2010 16:00:00 EST\n",
	"Expires: 1000000000000000000000000000000000000000000000000</repeat>\n",
	"Max-Forwards: 300\n",
	"Retry-After: 949302838503028349304023988\n",
	"To: \"Mr. J. User <sip:j.user@example.com>\n",
	"Via: SIP/2.0/UDP 192.0.2.15;;,;,,\n",
	"Via: SIP/7.0/UDP c.example.com;branch=z9hG4bKkdjuw\n",

	"Warning: 1812 overture \"In Progress\"\n",

	"Contact: sip:user@example.com?Route=%3Csip:sip.example.com%3E\n",
	"From:    Bell, Alexander <sip:a.g.bell@example.com>;tag=43\n",
	"To:      Watson, Thomas <sip:t.watson@example.org>\n",
	"To: \"Watson, Thomas\" < sip:t.watson@example.org >\n",
}

var torture4_o = []string{
	"Invalid: Contact: \"Joe\" <sip:joe@example.org>;;;;\n",
	"Invalid: Content-Length: -999\n",
	"Invalid: CSeq: 36893488147419103232 REGISTER\n",
	"Invalid: CSeq: 9292394834772304023312 OPTIONS\n",
	"Invalid: Contact: <sip:user@host129.example.com>;expires=280297596632815\n",
	"Invalid: Date: Fri, 01 Jan 2010 16:00:00 EST\n",
	"Invalid: Expires: 1000000000000000000000000000000000000000000000000</repeat>\n",
	"Invalid: Max-Forwards: 300\n",
	"Invalid: Retry-After: 949302838503028349304023988\n",
	"Invalid: To: \"Mr. J. User <sip:j.user@example.com>\n",
	"Invalid: Via: SIP/2.0/UDP 192.0.2.15;;,;,,\n",
	"Invalid: Via: SIP/7.0/UDP c.example.com;branch=z9hG4bKkdjuw\n",

	"Invalid: Warning: 1812 overture \"In Progress\"\n",

	"Invalid: Contact: sip:user@example.com?Route=%3Csip:sip.example.com%3E\n",
	"Invalid: From:    Bell, Alexander <sip:a.g.bell@example.com>;tag=43\n",
	"Invalid: To:      Watson, Thomas <sip:t.watson@example.org>\n",
	"Invalid: To: \"Watson, Thomas\" < sip:t.watson@example.org >\n",
}
