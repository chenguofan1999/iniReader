package inireader

func Watch(ls Listener, filePath string) (Cfg, error) {
	err := ls.Listen(filePath)
	cfg := Load(filePath)
	if err != nil {
		return cfg, ErrReadingIniFile{}
	} else {
		return cfg, nil
	}
}
