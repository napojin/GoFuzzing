package main

import (
	"errors"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		wantErr  error
	}{
		{"empty string", "", "", nil},
		{"single character", "a", "a", nil},
		{"two characters", "ab", "ba", nil},
		{"sentence", "The quick brown fox jumped over the lazy dog", "god yzal eht revo depmuj xof nworb kciuq ehT", nil},
		{"word", "foobar", "raboof", nil},
		{"日本語", "ローマは一日にして成らず", "ずら成てしに日一はマーロ", nil},
		{"invalid utf-8", string([]byte{0x80}), string([]byte{0x80}), errors.New("input is not valid UTF-8")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rev, err := Reverse(tc.input)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.expected, rev)
		})
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345", "日本語", string([]byte{0x80})}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, data string) {
		orig := string(data)
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}

		assert.Equal(t, orig, doubleRev, "Before: %q, after: %q", orig, doubleRev)

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			assert.Fail(t, "Reverse produced invalid UTF-8 string", rev)
		}
	})
}
