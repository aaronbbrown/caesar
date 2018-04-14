package cipher

import (
	"bytes"
)

type Caesar struct {
	Key CaesarKey
	Msg string
}

func NewCaesar(msg string, key CaesarKey) *Caesar {
	caesar := Caesar{
		Msg: msg,
		Key: key,
	}

	return &caesar
}

func (c *Caesar) Encrypt() string {
	return c.crypt(true)
}

func (c *Caesar) Decrypt() string {
	return c.crypt(false)
}

func (c *Caesar) crypt(encrypt bool) string {
	var buffer bytes.Buffer
	var idx int

	for _, r := range c.Msg {
		// shift right if encrypting, left otherwise
		dir := dirRight
		if !encrypt {
			dir = dirLeft
		}

		r := c.Key.ShiftRune(r, dir, &idx)
		buffer.WriteRune(r)
	}

	return buffer.String()
}
