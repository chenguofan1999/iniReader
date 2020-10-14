package inireader

import (
	"fmt"
)

var commentSymbol byte

// Cfg : The struct of configuration
type Cfg struct {
	Map map[string]*Sec
	Cur *Sec
}

// Sec : The struct of configuration Section
type Sec struct {
	Name        string
	Description string
	Map         map[string]string
}

// Section : Get section by section name
func (c Cfg) Section(name string) *Sec {
	if sec, ok := c.Map[name]; ok {
		return sec
	}

	fmt.Println("Did not find sec:", name)
	c.Map[name] = &Sec{Name: name, Map: map[string]string{}}
	return c.Map[name]
}

// Key : Get value by key
func (s Sec) Key(key string) string {
	return s.Map[key]
}
