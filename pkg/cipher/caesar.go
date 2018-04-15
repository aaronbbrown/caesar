package cipher

import (
	"bytes"
)

type encryptMode int

const (
	modeEncrypt encryptMode = iota
	modeDecrypt
)

type Caesar struct {
	Key CaesarKey
	Msg string
}

func NewCaesar(msg string, key CaesarKey) *Caesar {
	return &Caesar{
		Msg: msg,
		Key: key,
	}
}

func (c *Caesar) Encrypt() string {
	return c.crypt(modeEncrypt)
}

func (c *Caesar) Decrypt() string {
	return c.crypt(modeDecrypt)
}

func (c *Caesar) crypt(mode encryptMode) string {
	var buffer bytes.Buffer
	var idx int

	for _, r := range c.Msg {
		// shift right if encrypting, left otherwise
		dir := dirRight
		if mode == modeDecrypt {
			dir = dirLeft
		}

		shifted := c.Key.ShiftRune(r, dir, &idx)
		buffer.WriteRune(shifted)
	}

	return buffer.String()
}
