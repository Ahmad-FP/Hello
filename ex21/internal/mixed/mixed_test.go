package mixed

import (
	"strings"
	"testing"
)


func TestTruncate(t *testing.T) {
	tests := []struct{
		s string
		max int
		need string
	}{
		{"abc",5,"abc"},
		{"abcdef",3, "abc\u2026"},
		{strings.Repeat("é",210), 200, strings.Repeat("é",200) + "\u2026" },
	}
	for _, val := range tests {
		if res := Truncate(val.s,val.max); res != val.need{
			t.Errorf("result: %q expected: %q", res, val.need)
		}
	}
}