package cipher

import (
	"reflect"
	"testing"
)

func TestKeyToOffsets(t *testing.T) {
	c := Caesar{}
	c.SetKey(1029814123)
	expected := []int{4, 5, 11, 8, 2, 14, 23}
	if !reflect.DeepEqual(c.Offsets, expected) {
		t.Errorf("Got %+v, expected %+v", c.Offsets, expected)
	}
}

func TestEncrypt(t *testing.T) {
	c := NewCaesar("abc xyz", 678)
	encrypted := c.Encrypt()
	expected := "bdf yac"
	if encrypted != expected {
		t.Errorf("Got %s, expected %s", encrypted, expected)
	}
}

func TestDecrypt(t *testing.T) {
	c := NewCaesar("bdf yac", 678)
	decrypted := c.Decrypt()
	expected := "abc xyz"
	if decrypted != expected {
		t.Errorf("Got %s, expected %s", decrypted, expected)
	}
}

func TestLetters(t *testing.T) {
	letters := letters()
	if letters[0] != 'a' {
		t.Errorf("Got '%c', expected 'a'", letters[0])
	}
}
