package base58check

import (
	"reflect"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	encoded, err := Encode("00", "44D00F6EB2E5491CD7AB7E7185D81B67A23C4980F62B2ED0914D32B7EB1C5581")
	if err != nil {
		t.Error(err.Error())
	}
	actual := "1XJjHG4gLiJfxrx82yPFWC8tu8cxKvaQjZNvVfSrsfiX4mbUsw"
	if !reflect.DeepEqual(encoded, actual) {
		t.Errorf("Expected %s, got %s", actual, encoded)
	}
}

func TestDecode(t *testing.T) {
	decoded, err := Decode("5JLbJxi9koHHvyFEAERHLYwG7VxYATnf8YdA9fiC6kXMghkYXpk")
	if err != nil {
		t.Error(err.Error())
	}
	actual := strings.ToLower("8044D00F6EB2E5491CD7AB7E7185D81B67A23C4980F62B2ED0914D32B7EB1C5581")
	if !reflect.DeepEqual(decoded, actual) {
		t.Errorf("Expected %s, got %s", actual, decoded)
	}
}
