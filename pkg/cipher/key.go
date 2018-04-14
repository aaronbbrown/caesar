package cipher

import (
	"fmt"
	"strconv"
	"unicode"
)

type direction int

const (
	base = 25

	dirLeft direction = iota
	dirRight
)

type CaesarKey struct {
	Key     int
	offsets []int
	letters []rune
}

func NewCaesarKey(key int) CaesarKey {
	offsets := []int{}

	str := strconv.FormatInt(int64(key), base)
	for _, r := range str {
		val, _ := strconv.ParseInt(string(r), base, 64)
		offsets = append(offsets, int(val))
	}

	return CaesarKey{
		offsets: offsets,
		Key:     key,
		letters: letters(),
	}
}

// shift the rune in the specified direction using the offset into the offsets
// slice as specified. Offset will be incremented if the key is in the letters
// slice.
func (k *CaesarKey) ShiftRune(r rune, dir direction, idx *int) rune {
	lower := unicode.ToLower(r)
	if lower >= k.letters[0] && lower <= k.letters[len(k.letters)-1] {
		offset := int(lower) - int(k.letters[0])
		if dir == dirRight {
			offset += k.offsets[saneModInt(*idx, len(k.offsets))]
		} else {
			offset -= k.offsets[saneModInt(*idx, len(k.offsets))]
		}
		shifted := rune(k.letters[saneModInt(offset, len(k.letters))])
		*idx++
		return shifted
	}
	return r
}

func (k *CaesarKey) String() string {
	return fmt.Sprintf("%d %v", k.Key, k.offsets)
}
