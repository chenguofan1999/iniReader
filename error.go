package inireader

// ErrReadingIniFile indicates the error type of reading ini file failed
type ErrReadingIniFile struct {
}

func (err ErrReadingIniFile) Error() string {
	return "Error in reading init file"
}
