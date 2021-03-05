package dict

import (
	"bytes"
)

// FileType type
type FileType struct {
	Key      string
	Filename string
	Reader   *bytes.Reader
}
