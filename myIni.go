package inireader

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strings"
)

// Init : Determine commentSymbol by GOOS
func Init() {
	if runtime.GOOS == "windows" {
		commentSymbol = ';'
	} else {
		commentSymbol = '#'
	}
}

// InitConfig : Get configure struct from .init file
func InitConfig(fileName string) Cfg {

	Init()

	var initSec Sec = Sec{Name: "", Map: map[string]string{}}
	var cfg Cfg = Cfg{Map: map[string]*Sec{"": &initSec}, Cur: &initSec}

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := string(b)

		// empty line
		if len(s) == 0 {
			continue
		}

		if s[0] == '[' {
			// A section
			index := strings.Index(s, "]")
			secName := strings.TrimSpace(s[1:index])
			cfg.Cur = cfg.Section(secName)
		} else if s[0] == commentSymbol {
			// A description for sec
			desc := s[1:]
			cfg.Cur.Description = desc
		} else {
			// A key - value pair
			index := strings.Index(s, "=")
			if index < 0 {
				continue
			}

			key := strings.TrimSpace(s[:index])
			if len(key) == 0 {
				continue
			}

			val := strings.TrimSpace(s[index+1:])
			if len(val) == 0 {
				continue
			}

			cfg.Cur.Map[key] = val
		}

	}

	return cfg

}
