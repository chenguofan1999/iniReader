## func `Init`

```go
func Init()
```

Init : init 函数，使得 Unix 系统默认采用 # 作为注释行，Windows 系统默认采用 ; 作为注释行。

## type `Cfg`

```go
type Cfg struct {
	Map               map[string]*Sec
	LastDescription   string
	UnusedDescription bool
	Cur               *Sec
}
```

Cfg : 配置信息的数据结构


## func `Load`

```go
func Load(fileName string) Cfg
```

Load : 从 .ini 文件中读取配置信息，返回一个 Cfg 

## func `Watch`

```go
func Watch(ls Listener, filePath string) (Cfg, error)
```

Watch : 开始监听目标 .ini 文件，目标文件发生改动后停止监听，对修改后的文件进行解析，返回解析的配置信息。






## type `Listener`
```go 
type Listener interface {
	Listen(string) error
}
```
Listener 是一个接口


## func (FileListener) `Listen `
```go
func (fl FileListener) Listen(filePath string) error
```
Listen 是Listener的接口函数


## type `FileListener`
```go 
type FileListener struct {
}
```
FileListener 实现了 Listener 接口

## type `Sec`

```go
type Sec struct {
	Name         string
	Descriptions map[string]string
	Map          map[string]string
}
```
Sec : Section 的数据结构

### func (Sec) `Key`

```go
func (s Sec) Key(key string) string
```

Key : 根据键获取值

