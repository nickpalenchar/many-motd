package formatters

import (
	"testing"
)

func TestWrap(t *testing.T) {
	message := "The quick brown fox jumps over the lazy dog"
	width := 10
	want := "The quick\nbrown fox\njumps over\nthe lazy\ndog"

	have := wrap(message, width)

	if want != have {
		t.Fatalf("wrap = %#v, want %#v", have, want)
	}
}

func TestWrapLongWord(t *testing.T) {
	t.Skip()
	message := "this isaverylongwordandwillbesplit but the rest should go as usual"
	width := 18
	want := "this\nisaverylongwordand\nwillbesplit but\nthe rest should go\nas usual"

	have := wrap(message, width)

	if want != have {
		t.Fatalf("wrap (ling word) = %#v, want %#v", have, want)

	}
}
