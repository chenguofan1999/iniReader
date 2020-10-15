## 设计说明

本项目的主要目的是实现一个简单的 *.ini* 读取函数 `Load` ，该函数对目标 *.ini* 文件逐行读取，生成一个包含所有配置信息的数据结构。同时按要求需要设计一个接口 `listener` ，实现对文件修改的监听。将以上两者打包放进 `watch` 函数，便构成了这个包的基本骨架。

### 数据结构设计

.ini文件中有明显的层级结构，因此很自然能想到用两种自定义结构，分别为表示整体配置信息的 Cfg 和代表 Section 的 Sec:
- Cgf中主要的数据结构是一个map，以secName为key，Sec的指针为val
- sec中主要的数据结构也为map，具体有两个，均以属性名为key，分别查询值和注释。


```go
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
```

学习了 INI 包，也对接口做了简化：

```go
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
```

由此，可以用以下方法获取一个属性的值：

```go
cfg.Section(SecName).Key(KeyName)
```

也可以用以下方法获取一个属性的描述（如果没有就是空串）：
```go
cfg.Section(SecName).Descriptions[KeyName]
```

**自定义错误类型**

根据要求，也在项目中自定义了一种错误类型，该错误类型在主要的函数Watch中返回。

```go
type ErrReadingIniFile struct {
}

func (err ErrReadingIniFile) Error() string {
	return "Error in reading init file"
}

```

### Load 的实现


Load 在读取 *.ini* 文件的时候实现方式为按行读取，.ini文件中可能的行有四种：

| 种类 | 识别 |
|---|---|
|空行|长度为0|
|注释行|以#开头|
|切换Section|以 [ 开头，以 ] 结尾|
|键值对|非上述三种中任意一种且在中间有一个 = |

读取逻辑：

- 按空行、section、注释行、键值对的顺序进行判断
- Sec与属性的对应：通过当前Sec指针 *Cur* 来实现
- 属性与注释的对应：由于是先有注释后出现属性，由一个flag值 *UnusedDescription* 控制，预存注释为 *LastDescription* ，在紧接的一行遇到键值对时配对，否则丢弃。

```go

if len(s) == 0 {
    // empty line
    cfg.UnusedDescription = false

} else if s[0] == '[' {
    // A section
    cfg.UnusedDescription = false
    index := strings.Index(s, "]")
    secName := strings.TrimSpace(s[1:index])
    cfg.Cur = cfg.Section(secName)
    
} else if s[0] == commentSymbol {
    // A description for sec
    desc := strings.TrimSpace(s[1:])

    if cfg.UnusedDescription {
        cfg.LastDescription += "\n"
        cfg.LastDescription += desc
    } else {
        cfg.LastDescription = desc
    }
    cfg.UnusedDescription = true
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
    if cfg.UnusedDescription {
        cfg.Cur.Descriptions[key] = cfg.LastDescription
        cfg.UnusedDescription = false
    }
}   
```

### listener 接口设计

熟悉接口设计是这次项目的目的之一，需要完成是设计一个全新的接口，并用一个函数 / 结构来实现它。

接口的定义是简洁的，只有一个必须实现的函数。

```go
// Listener is the interface of listener
type Listener interface {
	Listen(string) error
}
```

在golang中，实现接口是隐式的，也就是说不需要显式地声明一个类要实现某个具体的接口，而是一个结构只要实现了一个接口要求的所有函数，便自动地实现了该接口。比如这里的`FileListener`实现了`Listener`接口所要求的所有函数 (仅一个`Listen`函数),便实现了该接口。

```go


// FileListener is the interface of FileListener
type FileListener struct {
}

// Listen implemented by FileListener
func (fl FileListener) Listen(filePath string) error {
	initialStat, err := os.Stat(filePath)
	initialStatSize := initialStat.Size()
	initialStatModTime := initialStat.ModTime()
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStatSize || stat.ModTime() != initialStatModTime {
			// fmt.Println("A file change detected, how do you want to deal load it again? y/Y for Yes, n/N for No")

			// // var choose string
			// fmt.Scanln(&choose)

			// for choose != "y" && choose != "Y" && choose != "N" && choose != "n" && choose != "exit" {
			// 	fmt.Println("Invalid option! Do you want to load it again? y/Y for Yes, n/N for No")
			// 	fmt.Scanln(&choose)
			// }

			// switch choose {
			// case "y", "Y":
			// 	Load(filePath)
			// 	fmt.Println("Reload!")
			// case "n", "N":
			// 	fmt.Println("Fine")
			// case "exit":
			// 	return nil
			// }

			break

		}
		initialStatSize = stat.Size()
		initialStatModTime = stat.ModTime()
		time.Sleep(1 * time.Second)
	}
	return nil
}
```

这里的 FileListener 调用 Listen 后开始对文件进行监听，在检测到文件发生修改后便退出。

（ 如果将这里的 break 换为注释掉的代码，实现的便是当检测到文件修改时询问用户是否重新读取ini文件信息，这里为了watch函数能正常返回并有同时符合函数签名和基本逻辑的使用方法，选择在检测到文件改动时直接退出 ）





### Watch 函数

最后是整合到一个 Watch 函数，

```go
func Watch(ls Listener, filePath string) (Cfg, error) {
	err := ls.Listen(filePath)
	cfg := Load(filePath)
	if err != nil {
		return cfg, ErrReadingIniFile{}
	} else {
		return cfg, nil
	}
}
```



## 单元测试

项目中对每个go文件都编写了一个测试文件用于单元测试。具体可见源代码：

[listener_test.go](listener_test.go)  
[load_test.go](load_test.go)  
[struct_test.go](struct_test.go)  
[watch_test.go](watch_test.go)


执行go test：

(TestFileListener 和 TestWatch 均需要监听等待一次文件修改，因此测试时需要手动保存一下testData/my.ini)

```sh
# chen @ ChenSurface in ~/go/src/github.com/chenguofan1999/inireader on git:main x [19:55:47] 
$ go test -v
=== RUN   TestFileListener
--- PASS: TestFileListener (13.01s)
=== RUN   TestLoadDescription
--- PASS: TestLoadDescription (0.00s)
=== RUN   TestSection
--- PASS: TestSection (0.00s)
=== RUN   TestStruct1
--- PASS: TestStruct1 (0.00s)
=== RUN   TestWatch
--- PASS: TestWatch (2.00s)
PASS
ok      github.com/chenguofan1999/inireader     15.010s
```

测试均能通过

## 功能测试

功能测试中主要测试 watch 函数是否能按预期执行：

- 传入一个 listener 和 文件路径
- 开始执行 watch 函数后程序对该文件进行监听，监听其是否发生更改
- 此时对该文件中的配置信息进行修改并保存
- watch 函数应该在此时返回修改后的配置信息

测试目录：
```
$ tree
.
├── login.ini
└── main.go
```


一开始，`login.ini`：

```ini
[User1]

# user_name
name = Luojun

# user_email
email = president@mail.sysu.edu.cn

# user_password
password = HowDoIKnow
```

`main.go`：

```go
package main

import (
	"fmt"

	"github.com/chenguofan1999/inireader"
)

func main() {

	// Watch for changes of a file and then load ini info
	// After running main.go please edit login.ini and save it
	var fl inireader.FileListener
	cfgLogin, err := inireader.Watch(fl, "login.ini")

	if err != nil {
		panic(err)
	}

	fmt.Println("User1.name")
	fmt.Println("	Description: ", cfgLogin.Section("User1").Descriptions["name"])
	fmt.Println("	Value      : ", cfgLogin.Section("User1").Key("name"))

	fmt.Println("User1.email")
	fmt.Println("	Description: ", cfgLogin.Section("User1").Descriptions["email"])
	fmt.Println("	Value      : ", cfgLogin.Section("User1").Key("email"))

	fmt.Println("User1.password")
	fmt.Println("	Description: ", cfgLogin.Section("User1").Descriptions["password"])
	fmt.Println("	Value      : ", cfgLogin.Section("User1").Key("password"))
}
```

开始运行，

```sh
go run main.go
```

开始运行后可见程序进入等待，此时编辑 `login.ini` 为：

```
[User1]

# user_name
name = Luojun

# user_email
email = president@mail.sysu.edu.cn

# user_password
password = HowDoIKnow
```

按下保存后观察到终端有输出：

```sh
$ go run main.go
User1.name
        Description:  user_name
        Value      :  Luojun
User1.email
        Description:  user_email
        Value      :  president@mail.sysu.edu.cn
User1.password
        Description:  user_password
        Value      :  HowDoIKnow
```

此时程序结束，程序的功能符合预期。