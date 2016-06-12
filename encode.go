package base65536

// Encode returns the base65536 encoding of src (a list of words).
func Encode(src []byte) []string {
	dstSize := len(src)
	dst := make([]string, dstSize, dstSize)
	for i := 0; i < len(src)-1; i += 2 {
		dst[i] = adjectives[src[i]]
		dst[i+1] = nouns[src[i+1]]
	}
	if len(src)%2 == 1 {
		// len(src) is odd
		dst[dstSize-1] = nouns[src[len(src)-1]]
	}
	return dst
}
