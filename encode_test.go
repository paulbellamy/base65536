package base65536

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func ExampleEncode() {
	for _, ip := range [][]byte{
		{127, 0, 0, 1},
	} {
		fmt.Println(net.IP(ip), strings.Join(Encode(ip), " "))
	}
	// Output: 127.0.0.1 wooden people different history
}

func TestEncode_OddLengths(t *testing.T) {
	in := []byte{0xFF}
	got := strings.Join(Encode(in), " ")
	expected := nouns[0xFF]
	if got != expected {
		t.Errorf("Expected %v, to encode to: %s, but got %s", in, expected, got)
	}
}
