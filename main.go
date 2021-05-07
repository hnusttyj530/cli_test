/**
  @author:tyj
  @date:2021/5/6
  @note:1670171244@qq.com
  @TODO:
  @Param:
  @return:
**/
package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"path/filepath"
	)

var GetCPU = cli.Command{
	Name:  "getCPU",
	Usage: "get CPU information from Linux OS",
	Category: "get",
	Action: func(c *cli.Context) {
		fmt.Println("get cpu!")
	},
}
var GetMac = cli.Command{
	Name:  "getMac",
	Usage: "get Mac address from Linux OS",
	Category: "get",
	Action: func(c *cli.Context) {
		fmt.Println("get mac!")
	},
}
var GetDisk = cli.Command{
	Name:  "getDisk",
	Usage: "get disk information from Linux OS",
	Category: "get",
	//Flags: []cli.Flag{
	//	cli.StringFlag{
	//		Name:  "file,f",
	//		Value: "../conf/toolconfig.ini",
	//		Usage: "zookeeper config file",
	//	},
	//},
	Action: func(c *cli.Context) {
		fmt.Println("get disk!")
	},
}

var GetMainBoard = cli.Command{
	Name:  "getBoard",
	Usage: "get main board information from Linux OS",
	Category: "get",
	Action: func(c *cli.Context) {
		fmt.Println("get main board sequnce!")
	},
}

var GetArgs=cli.Command{
	Name:"getArgs",
	Usage: "arguments example.",
	Category: "get",
	Action: func(c *cli.Context){
		for i := 0; i < c.NArg(); i++ {
			fmt.Printf("%d: %s\n", i+1, c.Args().Get(i))
		}
	},
}

var GetLanguage=cli.Command{
	UseShortOptionHandling: true,
	Name:"getLanguage",
	Usage: "get language.",
	Category: "set",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
			EnvVar: "SYSTEM_LANG",
			FilePath: "./lang.txt",
		},
	},
	Action: func(c *cli.Context) error {
		name := "world"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}

		if c.String("lang") == "english" {
			fmt.Println("hello", name)
		} else {
			name="世界"
			fmt.Println("你好", name)
		}
		return nil
	},
}

func main() {
	changeDirToBinPath()
	app := cli.NewApp()
	app.Name = "Encrypt"
	app.Author = "tyj"
	app.Version = "1.0"
	app.Usage = "Get the machine information from Linux Operation System."

	app.Commands = []cli.Command{
		GetCPU,
		GetDisk,
		GetMac,
		GetMainBoard,
		GetArgs,
		GetLanguage,
	}
	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stdout)
		return nil
	}

	//运行app
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return
}

func changeDirToBinPath() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Errorf("can't get my exe pwd,%s", err.Error())
		os.Exit(1)
	}

	binPath := filepath.Dir(file)
	err = os.Chdir(binPath)
	if err != nil {
		log.Errorf("can't chdir to bin path %s", binPath)
		os.Exit(1)
	}
}
