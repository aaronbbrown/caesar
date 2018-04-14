package cipher

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	c := NewCaesar("abc xyz", NewCaesarKey(678))
	encrypted := c.Encrypt()
	expected := "bdf yac"
	if encrypted != expected {
		t.Errorf("Got %s, expected %s", encrypted, expected)
	}
}

func TestDecrypt(t *testing.T) {
	c := NewCaesar("bdf yac", NewCaesarKey(678))
	decrypted := c.Decrypt()
	expected := "abc xyz"
	if decrypted != expected {
		t.Errorf("Got %s, expected %s", decrypted, expected)
	}
}
