package htmlparser

import (
	"bufio"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name      string
		in        string
		expTags   []Tag
		wantError bool
	}{
		{
			name:    "no input no tags",
			in:      "",
			expTags: []Tag(nil),
		},
		{
			name: "h1 only",
			in:   `<h1>Go Proverbs</h1>`,
			expTags: []Tag{
				newTag(TagTypeHeaderLvl1, "Go Proverbs"),
			},
		},
		{
			name: "p only",
			in:   `<p>Don't panic!</p>`,
			expTags: []Tag{
				newTag(TagTypeParagraph, "Don't panic!"),
			},
		},
		{
			name: "h1 p",
			in:   `<h1>Go Proverbs</h1><p>Don't panic!</p>`,
			expTags: []Tag{
				newTag(TagTypeHeaderLvl1, "Go Proverbs"),
				newTag(TagTypeParagraph, "Don't panic!"),
			},
		},
		{
			name: "h1 p p",
			in:   `<h1>Go Proverbs</h1><p>Don't panic!</p><p>Errors are values</p>`,
			expTags: []Tag{
				newTag(TagTypeHeaderLvl1, "Go Proverbs"),
				newTag(TagTypeParagraph, "Don't panic!"),
				newTag(TagTypeParagraph, "Errors are values"),
			},
		},
		{
			name: "h1 p p h1 p",
			in: `
<h1>Go Proverbs</h1>
<p>Don't panic!</p>
<p>Errors are values</p>

<h1>Need to remember</h1>
<p>Premature optimization is the root of all evil</p>`,
			expTags: []Tag{
				newTag(TagTypeHeaderLvl1, "Go Proverbs"),
				newTag(TagTypeParagraph, "Don't panic!"),
				newTag(TagTypeParagraph, "Errors are values"),

				newTag(TagTypeHeaderLvl1, "Need to remember"),
				newTag(TagTypeParagraph, "Premature optimization is the root of all evil"),
			},
		},
		{
			name: "p h1 p",
			in:   `<p>Welcome </p><h1>https://golang-ninja.ru/</h1><p> to site!</p>`,
			expTags: []Tag{
				newTag(TagTypeParagraph, "Welcome "),
				newTag(TagTypeHeaderLvl1, "https://golang-ninja.ru/"),
				newTag(TagTypeParagraph, " to site!"),
			},
		},
		{
			name: "with unicode",
			in:   `<h1>Труд в СССР есть дело чести, славы, доблести и геройства</h1>`,
			expTags: []Tag{
				newTag(TagTypeHeaderLvl1, "Труд в СССР есть дело чести, славы, доблести и геройства"),
			},
		},

		// Negative scenario.
		{
			name:      "unknown tag",
			in:        `<h2>Go Proverbs</h2>`,
			wantError: true,
		},
		{
			name:      "malformed opening tag",
			in:        `<h1`,
			wantError: true,
		},
		{
			name:      "malformed opening tag",
			in:        `<h1 Go Proverbs</h1>`,
			wantError: true,
		},
		{
			name:      "no closing tag",
			in:        `<h1>Go Proverbs`,
			wantError: true,
		},
		{
			name:      "incorrect closing tag",
			in:        `<h1>Go Proverbs</p>`,
			wantError: true,
		},
		{
			name:      "opening tag as closing tag",
			in:        `<h1>Go Proverbs<h1>`,
			wantError: true,
		},
		{
			name:      "nested tags",
			in:        `<h1>Go Proverbs <p>Don't panic!</p></h1>`,
			wantError: true,
		},
		{
			name:      "outside data",
			in:        "<h1>Go</h1>Proverbs",
			wantError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tags, err := Parse([]byte(tt.in))
			if tt.wantError {
				t.Log(err)
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.expTags, tags)
		})
	}
}

func Test_catchErr(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		tags := []Tag{
			newTag(TagTypeHeaderLvl1, "Go Proverbs"),
			newTag(TagTypeParagraph, "Don't panic!"),
		}
		var err error

		defer catchErr(&tags, &err)

		assert.NoError(t, err)
		assert.Equal(t, []Tag{
			newTag(TagTypeHeaderLvl1, "Go Proverbs"),
			newTag(TagTypeParagraph, "Don't panic!"),
		}, tags)
	})

	t.Run("known error", func(t *testing.T) {
		tags := []Tag{
			newTag(TagTypeHeaderLvl1, "Go Proverbs"),
			newTag(TagTypeParagraph, "Don't panic!"),
		}
		var err error

		errForPanic := parseError{error: bufio.ErrInvalidUnreadByte}
		defer func() {
			assert.Equal(t, errForPanic, err) // Not ErrorIs!
			assert.Nil(t, tags)
		}()

		defer catchErr(&tags, &err)
		panic(errForPanic)
	})

	t.Run("unknown error", func(t *testing.T) {
		var tags []Tag
		var err error

		valueForPanic := new(runtime.TypeAssertionError)
		assert.PanicsWithValue(t, valueForPanic, func() {
			defer catchErr(&tags, &err)
			panic(valueForPanic)
		})
	})
}

func newTag(t TagType, v string) Tag {
	return Tag{Type: t, Val: v}
}
