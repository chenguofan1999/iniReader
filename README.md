# INIreader


[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/chenguofan1999/inireader)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=for-the-badge&logo=sourcegraph)](https://sourcegraph.com/github.com/chenguofan1999/iniReader)

![](icon.jpg)

本包提供了 Go 语言中读 INI 文件的功能。

```
Q : 写 INI 文件呢？  
A : 不能。
```

### 看起来很眼熟？

是的，这个包参考了 [INI](https://github.com/go-ini/ini)。  


## 功能特性

- 支持基本的 INI 文件读取
- 支持分区
- 支持属性注释，包括多行注释
- 支持**监听** INI 文件改动, 并重新加载
- 轻松操作分区、键值和注释

## 安装

安装所需的最低 Go 语言版本不得而知，祝你好运。

```sh
go get github.com/chenguofan1999/inireader
```

**更新：**

更新请自便，反正我们也不会增加新的功能。

```sh
go get -u github.com/chenguofan1999/inireader
```

## 使用示例

目录结构：

```
$ tree
.
├── login.ini
├── main.go
└── my.ini
```

*my.ini*

```ini
# possible values : production, development
app_mode = development

[paths]
# Path to where grafana can store temp files
data = /home/git/grafana



[server]
# Protocol (http or https)
protocol = http

# The http port  to use
http_port = 9999

# Redirect to correct domain if host header does not match domain
# Prevents DNS rebinding attacks
enforce_domain = 0
```

*login.ini*

```ini
[User1]

# user_name
name = yourName

# user_email
email = yourEmail

# user_password
password = yourPassword
```


*main.go*

```go
package main

import (
	"fmt"

	"github.com/chenguofan1999/inireader"
)

func main() {
	// Read .ini file directly by calling inireader.Load
	cfg := inireader.Load("my.ini")
	fmt.Println(".app_mode : ")
	fmt.Println("	Description: ", cfg.Section("").Descriptions["app_mode"])
	fmt.Println("	Value      : ", cfg.Section("").Key("app_mode"))

	fmt.Println("paths.data: ")
	fmt.Println("	Description: ", cfg.Section("paths").Descriptions["data"])
	fmt.Println("	Value      : ", cfg.Section("paths").Key("data"))

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

运行它！

```sh
go run main.go
```


结果如下，注意程序并没有运行结束，而是在执行 watch 函数，等待 *login.ini* 被修改。

```
$ go run main.go
.app_mode : 
        Description:  possible values : production, development
        Value      :  development
paths.data: 
        Description:  Path to where grafana can store temp files
        Value      :  /home/git/grafana
```

现在编辑 *login.ini* 并保存

```
[User1]

# user_name
name = Luojun

# user_email
email = president@mail.sysu.edu.cn

# user_password
password = HowDoIKnow
```

终端中出现了新的输出，程序终止。

```
User1.name
        Description:  user_name
        Value      :  Luojun
User1.email
        Description:  user_email
        Value      :  president@mail.sysu.edu.cn
User1.password
        Description:  user_password
        Value      :  HowDoIKnow
$
```



## License

这个项目用的是 wtfpl License.
点击 [LICENSE](https://github.com/anak10thn/WTFPL) 查看具体内容。

## API 文档

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](doc_zh_CN.md)