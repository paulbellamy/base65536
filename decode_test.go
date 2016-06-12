package base65536

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ExampleDecode() {
	input := "traditional bird able blossom"
	output, err := Decode(strings.Split(input, " "))
	if err != nil {
		fmt.Println("Error:", err)

	} else {
		fmt.Println(
			output,
			input,
		)
	}
	// Output:
	// [127 0 0 1] traditional bird able blossom
}

func TestDecode_OddLengths(t *testing.T) {
	in := []string{nouns[0xFF]}
	got, err := Decode(in)
	if err != nil {
		t.Error(err)
	}
	expected := []byte{0xFF}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, to decode to: %s, but got %s", in, expected, got)
	}
}

func TestDecode_CorruptInput(t *testing.T) {
	in := []string{"ancient", "corrupt"} // Not in the dictionary
	_, err := Decode(in)
	expected := CorruptInputError(1)
	if err != expected {
		t.Errorf("Expected %v, to error with: %s, but got %s", in, expected, err)
	}
}
