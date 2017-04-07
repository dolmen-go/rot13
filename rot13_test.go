package rot13_test

import (
	"testing"

	"github.com/dolmen-go/rot13"
)

func TestRot13(t *testing.T) {
	enc := rot13.Encoding{}
	tests := []string{
		"", "",
		" ", " ",
		"a", "n",
		" AaMmNn ", " NnZzAa ",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		"NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm",
	}
	for i := 0; i < len(tests); i += 2 {
		in := tests[i]
		out := tests[i+1]
		t.Logf("%q", in)

		dst := make([]byte, enc.EncodedLen(len(in)))
		enc.Encode(dst, []byte(in))
		if string(dst) != out {
			t.Errorf("%q -> %q", in, dst)
			continue
		}

		src := dst
		dst = make([]byte, enc.DecodedLen(len(src)))
		n, err := enc.Decode(dst, src)
		if err != nil {
			t.Error("error:", err)
			continue
		}
		if n != len(dst) {
			t.Error("n:", n)
			continue
		}
		if string(dst) != in {
			t.Errorf("%q -> %q", out, dst)
		}
	}
}
