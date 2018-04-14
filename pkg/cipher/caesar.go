package cipher

import (
	"bytes"
	"strconv"
	"unicode"
)

const (
	base = 25

	dirLeft direction = iota
	dirRight
)

type direction int

type Caesar struct {
	Key     int
	Offsets []int
	Msg     string
	letters []rune
}

func NewCaesar(msg string, key int) *Caesar {
	caesar := Caesar{
		Msg:     msg,
		Key:     key,
		letters: letters(),
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

// shift the rune in the specified direction using the offset into the offsets
// slice as specified. Offset will be incremented if the key is in the letters
// slice.
func (c *Caesar) shift(r rune, dir direction, idx *int) rune {
	lower := unicode.ToLower(r)
	if lower >= c.letters[0] && lower <= c.letters[len(c.letters)-1] {
		offset := int(lower) - int(c.letters[0])
		if dir == dirRight {
			offset += c.Offsets[saneModInt(*idx, len(c.Offsets))]
		} else {
			offset -= c.Offsets[saneModInt(*idx, len(c.Offsets))]
		}
		shifted := rune(c.letters[saneModInt(offset, len(c.letters))])
		*idx++
		return shifted
	}
	return r

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

		r := c.shift(r, dir, &idx)
		buffer.WriteRune(r)
	}

	return buffer.String()
}
