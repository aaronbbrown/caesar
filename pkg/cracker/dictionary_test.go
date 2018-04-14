package cracker

import "testing"

func TestPercentMatches(t *testing.T) {
	d := Dictionary(map[string]bool{
		"foo": false,
		"bar": false,
	})

	percent := d.PercentMatches("foo baz foo bar blah asdf fdsa bar")
	if percent != 50 {
		t.Errorf("Got %d expected 50", percent)
	}
}
