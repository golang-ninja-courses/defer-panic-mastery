package htmlparser

type TagType string

const (
	TagTypeHeaderLvl1 TagType = "H1"
	TagTypeParagraph  TagType = "P"
)

type Tag struct {
	Type TagType
	Val  string
}
