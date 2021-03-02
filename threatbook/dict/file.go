package dict

import (
	"bytes"
)

type FileType struct {
	Key      string
	Filename string
	Reader   *bytes.Reader
}
