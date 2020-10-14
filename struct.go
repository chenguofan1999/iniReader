package inireader

var commentSymbol byte

// Cfg : The struct of configuration
type Cfg struct {
	Map               map[string]*Sec
	LastDescription   string
	UnusedDescription bool
	Cur               *Sec
}

// Sec : The struct of configuration Section
type Sec struct {
	Name         string
	Descriptions map[string]string
	Map          map[string]string
}

// Section : Get section by section name
func (c Cfg) Section(name string) *Sec {
	if sec, ok := c.Map[name]; ok {
		return sec
	}

	//fmt.Println("Creating Sec: ", name)
	c.Map[name] = &Sec{Name: name, Map: map[string]string{}, Descriptions: map[string]string{}}
	return c.Map[name]
}

// Key : Get value by key
func (s Sec) Key(key string) string {
	return s.Map[key]
}
