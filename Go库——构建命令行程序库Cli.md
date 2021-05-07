# Go库——构建命令行程序库Cli

## 1.快速入门

```shell
$ mkdir cli && cd cli
$ go mod init github.com/darjun/go-daily-lib/cli
```

获取V2版本

```shell
$ go get -u github.com/urfave/cli/v2
```

使用：

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name:  "hello",
    Usage: "hello world example",
    Action: func(c *cli.Context) error {
      fmt.Println("hello world")
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

编译

```shell
$ go build 文件名
```

## 2.参数

通过`cli.Context`的相关方法我们可以获取传给命令行的参数信息：

- `NArg()`：返回参数个数；
- `Args()`：返回`cli.Args`对象，调用其`Get(i)`获取位置`i`上的参数。

示例：

```go
func main() {
  app := &cli.App{
    Name:  "arguments",
    Usage: "arguments example",
    Action: func(c *cli.Context) error {
      for i := 0; i < c.NArg(); i++ {
        fmt.Printf("%d: %s\n", i+1, c.Args().Get(i))
      }
      return nil
    },
  }
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

## 3.选项

`cli`设置和获取选项非常简单。在`cli.App{}`结构初始化时，设置字段`Flags`即可添加选项。`Flags`字段是`[]cli.Flag`类型，`cli.Flag`实际上是接口类型。`cli`为常见类型都实现了对应的`XxxFlag`，如`BoolFlag/DurationFlag/StringFlag`等。它们有一些共用的字段，`Name/Value/Usage`（名称/默认值/释义）。

带存入变量language的

```Go
func main() {
  var language string

  app := &cli.App{
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name:        "lang",
        Value:       "english",
        Usage:       "language for the greeting",
        Destination: &language,
      },
    },
    Action: func(c *cli.Context) error {
      name := "world"
      if c.NArg() > 0 {
        name = c.Args().Get(0)
      }

      if language == "english" {
        fmt.Println("hello", name)
      } else {
        fmt.Println("你好", name)
      }
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

**别名**：选项可以设置多个别名，设置对应选项的`Aliases`字段即可：

```yaml
 Aliases: []string{"language", "l"},
```

**环境变量**：只需要将环境变量的名字设置到选项对象的`EnvVars`字段即可。可以指定多个环境变量名字，`cli`会依次查找，第一个有值的环境变量会被使用。

```yaml
 EnvVar: "APP_LANG",
```

**文件**：`cli`还支持从文件中读取选项的值，设置选项对象的`FilePath`字段为文件路径。

```yaml
FilePath: "./lang.txt",
```

**组合短选项**：`cli`也支持短选项合写，只需要设置`cli.App`的`UseShortOptionHandling`字段为`true`即可。

```yaml
UseShortOptionHandling: true,
```

**必要选项**：如果将选项的`Required`字段设置为`true`，那么该选项就是必要选项。必要选项必须指定，否则会报错：

```yaml
Required: true,
```

**帮助文本中的默认值**：

默认情况下，帮助文本中选项的默认值显示为`Value`字段值。有些时候，`Value`并不是实际的默认值。这时，我们可以通过`DefaultText`设置：

```yaml
DefaultText :"random",
```

**子命令**

`cli`通过设置`cli.App`的`Commands`字段添加命令，设置各个命令的`SubCommands`字段，即可添加子命令。

```json
Subcommands: []*cli.Command{
	{
		Name:  "add",
		Usage: "add a new template",
		Action: func(c *cli.Context) error {
			fmt.Println("new task template: ")
			return nil
		},
	},
	{
		Name:  "remove",
		Usage: "remove an existing template",
		Action: func(c *cli.Context) error {
			fmt.Println("removed task template: ")
			return nil
		},
	},
},
```

**分类**：在子命令数量很多的时候，可以设置`Category`字段为它们分类，在帮助信息中会将相同分类的命令放在一起展示

```yaml
Category: "template",
```

## 参考资料

1.Go语言构建命令行程序的Cli库：https://zhuanlan.zhihu.com/p/150396264

