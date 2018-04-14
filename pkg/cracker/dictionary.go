package cracker

import (
	"io/ioutil"
	"strings"
)

type Dictionary map[string]bool

func NewDictionary(fn string) (Dictionary, error) {
	dict := make(map[string]bool)
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	words := strings.Fields(string(buf))
	for _, word := range words {
		if len(word) > 2 {
			dict[word] = false
		}
	}

	return dict, nil
}

func (d Dictionary) PercentMatches(msg string) int {
	var count int

	words := strings.Fields(msg)

	for _, word := range words {
		// word present in dictionary
		if _, ok := d[word]; ok {
			count++
		}
	}

	return count * 100 / len(words)
}
