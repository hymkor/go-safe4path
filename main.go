package safe4path

import (
	"errors"
	"fmt"
	"strings"
)

// Filename limitations for Windows.
// See also https://www.fileside.app/blog/2023-03-17_windows-file-paths
const (
	hexTable = "0123456789ABCDEF"
)

func findInvalidChar(s string) int {
	for i, r := range s {
		switch r {
		case '<', '>', '"', '/', '\\', '|', '?', '*', '.', ':':
			return i
		default:
			if r < ' ' {
				return i
			}
		}
	}
	return -1
}

func ToSafe(name string, prefix byte) string {
	var buffer strings.Builder
	for {
		index := findInvalidChar(name)
		if index < 0 {
			buffer.WriteString(name)
			return buffer.String()
		}
		buffer.WriteString(name[:index])
		buffer.WriteByte(prefix)
		c := name[index]
		buffer.WriteByte(hexTable[c>>4])
		buffer.WriteByte(hexTable[c&15])
		name = name[index+1:]
	}
}

var (
	ErrTooShort          = errors.New("too short string error")
	ErrFirstByteInvalid  = errors.New("first byte is invalid")
	ErrSecondByteInvalid = errors.New("first byte is invalid")
)

func FromSafe(name string, prefix byte) (string, error) {
	var buffer strings.Builder
	for {
		index := strings.IndexByte(name, prefix)
		if index < 0 {
			buffer.WriteString(name)
			return buffer.String(), nil
		}
		if len(name) < 3 {
			return "", ErrTooShort
		}
		c1 := strings.IndexByte(hexTable, name[index+1])
		if c1 < 0 {
			return "", fmt.Errorf("%w: %c", ErrFirstByteInvalid, name[index+1])
		}
		c2 := strings.IndexByte(hexTable, name[index+2])
		if c2 < 0 {
			return "", fmt.Errorf("%w: %c", ErrSecondByteInvalid, name[index+2])
		}
		buffer.WriteString(name[:index])
		buffer.WriteByte(byte((c1 << 4) | c2))
		name = name[index+3:]
	}
}
