package inireader

import "fmt"

// ErrReadingIniFile indicates the error type of reading ini file failed
type ErrReadingIniFile struct {
	Line string
}

func (err ErrReadingIniFile) Error() string {
	return fmt.Sprintf("key-value delimiter not found: %s", err.Line)
}
