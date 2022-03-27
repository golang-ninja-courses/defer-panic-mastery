package htmlparser

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode"
)

// Parse разбирает HTML из идущих подряд заголовков первого уровня и/или параграфов.
func Parse(html []byte) (tags []Tag, err error) {
	defer catchErr(&tags, &err)

	var (
		currTagType  TagType
		currTagValue bytes.Buffer
		insideTag    bool
	)

	r := bufio.NewReader(bytes.NewReader(html))
	for {
		ch, err := r.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			throwErrf("read byte: %w", err)
		}

		if ch == '<' {
			if err := r.UnreadByte(); err != nil {
				throwErrf("unread byte: %w", err)
			}

			if !insideTag {
				currTagType = parseOpeningTag(r)
				insideTag = true
			} else {
				tt := parseClosingTag(r)
				if tt != currTagType {
					throwErrf("closing tag does not match opening tag")
				}

				tags = append(tags, Tag{
					Type: currTagType,
					Val:  currTagValue.String(),
				})
				currTagValue.Reset()
				insideTag = false
			}
			continue
		}

		if !insideTag {
			if unicode.IsSpace(rune(ch)) {
				continue
			}
			throwErrf("data is outside tag")
		}
		currTagValue.WriteByte(ch)
	}

	if insideTag {
		throwErrf("no closing tag for %q", currTagType)
	}
	return tags, nil
}

func parseOpeningTag(r io.ByteReader) TagType {
	// Реализуй меня.
	return ""
}

func parseClosingTag(r io.ByteReader) TagType {
	// Реализуй меня.
	return ""
}

type parseError struct {
	error
}

// catchErr ловит панику и, если её значение имеет тип parseError, то
// присваивает его ошибке err, зануляя tags. Иначе паника продолжается.
func catchErr(tags *[]Tag, err *error) {
	// Реализуй меня.
}

func throwErr(err error) {
	panic(parseError{error: err})
}

func throwErrf(format string, args ...interface{}) {
	throwErr(fmt.Errorf(format, args...))
}
