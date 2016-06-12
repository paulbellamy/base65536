package base65536

import "strconv"

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base65536 data at input word " + strconv.FormatInt(int64(e), 10)
}

// Decode returns the bytes represented by the base65536 words in src.
//
// It returns at most len(src) bytes. If src contains invalid base65536 data,
// it will return CorruptInputError.
func Decode(src []string) ([]byte, error) {
	if len(src) == 0 {
		return nil, nil
	}

	adjectiveIndex := index(adjectives)
	nounIndex := index(nouns)

	dstSize := len(src)
	dst := make([]byte, dstSize, dstSize)
	var ok bool
	for i := 0; i < len(src)-1; i += 2 {
		dst[i], ok = adjectiveIndex[src[i]]
		if !ok {
			return nil, CorruptInputError(i)
		}
		dst[i+1], ok = nounIndex[src[i+1]]
		if !ok {
			return nil, CorruptInputError(i + 1)
		}
	}
	if len(src)%2 == 1 {
		// len(src) is odd
		dst[dstSize-1], ok = nounIndex[src[len(src)-1]]
		if !ok {
			return nil, CorruptInputError(dstSize - 1)
		}
	}
	return dst, nil
}

// build an index of the haystack, where values are the offset in the original array
func index(haystack []string) map[string]byte {
	m := map[string]byte{}
	for i, item := range haystack {
		m[item] = byte(i)
	}
	return m
}
