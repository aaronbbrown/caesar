package cracker

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/aaronbbrown/caesar/pkg/cipher"
)

func TestCrack(t *testing.T) {
	dictionary, err := NewDictionary("../../wordlist.txt")
	if err != nil {
		t.Fatal("Could not load wordlist.txt")
	}

	data, err := ioutil.ReadFile("../../text.txt")
	if err != nil {
		t.Fatal("Could not load text.txt")
	}

	caesar := cipher.NewCaesar(string(data), 3)
	encrypted := caesar.Encrypt()
	caesar.Msg = encrypted
	cracker := NewCracker(encrypted, dictionary)
	t.Run("should decrypt successfully", func(t *testing.T) {
		decrypted, err := cracker.Crack(1, 8, 30, 5)
		if err != nil {
			t.Fatal("Crack failed")
		}
		inFields := strings.Fields(string(data))
		outFields := strings.Fields(decrypted.Msg)
		if strings.ToLower(inFields[0]) != strings.ToLower(outFields[0]) {
			t.Errorf("Crack failed, got: %q expected %q", decrypted, string(data))
		}
	})

	t.Run("should not decrypt successfully", func(t *testing.T) {
		_, err := cracker.Crack(1, 2, 30, 5)
		if err == nil {
			t.Fatal("Crack succeeded when it should have failed")
		}
	})

}
