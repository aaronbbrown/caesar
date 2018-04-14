package cipher

import (
	"reflect"
	"testing"
)

func TestNewCaesarKey(t *testing.T) {
	key := NewCaesarKey(1029814123)

	expected := []int{4, 5, 11, 8, 2, 14, 23}
	if !reflect.DeepEqual(key.offsets, expected) {
		t.Errorf("Got %+v, expected %+v", key.offsets, expected)
	}
}

func TestLetters(t *testing.T) {
	letters := letters()
	if letters[0] != 'a' {
		t.Errorf("Got '%c', expected 'a'", letters[0])
	}
}
