package allocationfilterutil

import (
	"testing"
)

func TestLexer(t *testing.T) {
	cases := []struct {
		name string

		input       string
		expectError bool
		expected    []token
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []token{{kind: eof}},
		},
		{
			name:     "colon",
			input:    ":",
			expected: []token{{kind: colon, s: ":"}, {kind: eof}},
		},
		{
			name:     "comma",
			input:    ",",
			expected: []token{{kind: comma, s: ","}, {kind: eof}},
		},
		{
			name:     "plus",
			input:    "+",
			expected: []token{{kind: plus, s: "+"}, {kind: eof}},
		},
		{
			name:     "bangColon",
			input:    "!:",
			expected: []token{{kind: bangColon, s: "!:"}, {kind: eof}},
		},
		{
			name: "multiple symbols",
			// This is a valid string to lex but not to parse.
			input:    "!::,+",
			expected: []token{{kind: bangColon, s: "!:"}, {kind: colon, s: ":"}, {kind: comma, s: ","}, {kind: plus, s: "+"}, {kind: eof}},
		},
		{
			name:     "string",
			input:    `"test"`,
			expected: []token{{kind: str, s: `test`}, {kind: eof}},
		},
		{
			name:     "keyed access",
			input:    "[app]",
			expected: []token{{kind: keyedAccess, s: "app"}, {kind: eof}},
		},
		{
			name:     "identifier pure alpha",
			input:    "abc",
			expected: []token{{kind: identifier, s: "abc"}, {kind: eof}},
		},
		{
			name:     "label access",
			input:    "app[kubecost]",
			expected: []token{{kind: identifier, s: "app"}, {kind: keyedAccess, s: "kubecost"}, {kind: eof}},
		},
		// TODO: more cases
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("Input: '%s'", c.input)
			result, err := lexAllocationFilterV2(c.input)
			if c.expectError && err == nil {
				t.Errorf("expected error but got nil")
			} else if !c.expectError && err != nil {
				t.Errorf("unexpected error: %s", err)
			} else {
				if len(c.expected) != len(result) {
					t.Fatalf("Token slices don't match in length.\nExpected: %+v\nGot: %+v", c.expected, result)
				}
				for i := range c.expected {
					if c.expected[i] != result[i] {
						t.Fatalf("Incorrect token at position %d.\nExpected: %+v\nGot: %+v", i, c.expected, result)
					}
				}
			}
		})
	}
}
