package parser

import (
	"testing"
)

func TestCallIdParser(t *testing.T) {
	var tvs = []string{
		"Call-ID: f0b40bcc-3485-49e7-ad1a-f1dfad2e39c9@10.5.0.53\n",
		"Call-ID: f81d4fae-7dec-11d0-a765-00a0c91e6bf6@foo.bar.com\n",
		"i:f81d4fae-7dec-11d0-a765-00a0c91e6bf6@foo.bar.com\n",
		"Call-ID: 1@10.0.0.1\n",
		"Call-ID: kl24ahsd546folnyt2vbak9sad98u23naodiunzds09a3bqw0sdfbsk34poouymnae0043nsed09mfkvc74bd0cuwnms05dknw87hjpobd76f\n",
		"Call-ID: 281794\n",
	}
	var tvs_o = []string{
		"Call-ID: f0b40bcc-3485-49e7-ad1a-f1dfad2e39c9@10.5.0.53\n",
		"Call-ID: f81d4fae-7dec-11d0-a765-00a0c91e6bf6@foo.bar.com\n",
		"Call-ID: f81d4fae-7dec-11d0-a765-00a0c91e6bf6@foo.bar.com\n",
		"Call-ID: 1@10.0.0.1\n",
		"Call-ID: kl24ahsd546folnyt2vbak9sad98u23naodiunzds09a3bqw0sdfbsk34poouymnae0043nsed09mfkvc74bd0cuwnms05dknw87hjpobd76f\n",
		"Call-ID: 281794\n",
	}

	for i := 0; i < len(tvs); i++ {
		shp := NewCallIDParser(tvs[i])
		testHeaderParser(t, shp, tvs_o[i])
	}
}
