package rot13

var rot13tab [256]byte

func init() {
	for i := 0; i < 256; i++ {
		rot13tab[i] = byte(i)
	}
	const lower = 'a' - 'A'
	for i := 'A'; i < 'A'+13; i++ {
		rot13tab[i] = byte(i + 13)
		rot13tab[i+13] = byte(i)
		rot13tab[i+lower] = byte(i + lower + 13)
		rot13tab[i+lower+13] = byte(i + lower)
	}
}

type Encoding struct{}

func (Encoding) EncodedLen(n int) int { return n }

func (Encoding) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}
	_ = dst[len(src)-1]
	for i := 0; i < len(src); i++ {
		dst[i] = rot13tab[src[i]]
	}
}

// EncodeToString returns the rot13 encoding of src.
func (Encoding) EncodeToString(src []byte) string {
	dst := make([]byte, len(src))
	Encoding{}.Encode(dst, src)
	return string(dst)
}

func (Encoding) DecodedLen(n int) int { return n }

func (Encoding) Decode(dst, src []byte) (int, error) {
	Encoding{}.Encode(dst, src)
	return len(src), nil
}

func (Encoding) DecodeString(src string) ([]byte, error) {
	dst := make([]byte, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = rot13tab[src[i]]
	}
	return dst, nil
}
