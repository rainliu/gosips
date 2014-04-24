package parser

import (
	"testing"
)

func TestCSeqParser(t *testing.T) {
	var cseq = []string{
		"CSeq: 17 INVITE\n",
		"CSeq: 17 ACK\n",
		"CSeq : 18   BYE\n",
		"CSeq:1 CANCEL\n",
		"CSeq: 3 BYE\n",
	}

	for i := 0; i < len(cseq); i++ {
		shp := NewCSeqParser(cseq[i])
		testParser(t, shp)
	}
}
