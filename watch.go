package inireader

// Watch func watches the changes of the target file
func Watch(ls Listener, filePath string) (Cfg, error) {
	err := ls.Listen(filePath)
	cfg := Load(filePath)
	if err != nil {
		return cfg, ErrReadingIniFile{}
	}
	return cfg, nil

}
