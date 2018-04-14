package cipher

import (
	"bytes"
	"strconv"
	"unicode"
)

type Caesar struct {
	Key     int
	Offsets []int
	Msg     string
}

const base = 25

func NewCaesar(msg string, key int) *Caesar {
	caesar := Caesar{
		Msg: msg,
		Key: key,
	}

	caesar.SetKey(key)

	return &caesar
}

func (c *Caesar) SetKey(key int) {
	offsets := []int{}

	str := strconv.FormatInt(int64(key), base)
	for _, r := range str {
		val, _ := strconv.ParseInt(string(r), base, 64)
		offsets = append(offsets, int(val))
	}

	c.Offsets = offsets
	c.Key = key
}

func (c *Caesar) Encrypt() string {
	return c.crypt(true)
}

func (c *Caesar) Decrypt() string {
	return c.crypt(false)
}

func (c *Caesar) crypt(encrypt bool) string {
	var buffer bytes.Buffer
	var n int

	letters := letters()

	for _, r := range c.Msg {
		lower := unicode.ToLower(r)
		if lower >= letters[0] && lower <= letters[len(letters)-1] {
			offset := int(lower) - int(letters[0])
			if encrypt {
				offset += c.Offsets[saneModInt(n, len(c.Offsets))]
			} else {
				offset -= c.Offsets[saneModInt(n, len(c.Offsets))]
			}
			shifted := rune(letters[saneModInt(offset, len(letters))])
			buffer.WriteRune(shifted)
			n++
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}
