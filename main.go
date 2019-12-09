package main

import (
	"musigo/actions"
	"os"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	version = "1.0.0"
	appName = "musigo"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Compiled = time.Now()

	app.Commands = []cli.Command{
		{
			Name:   "input",
			Usage:  "入力",
			Action: actions.Input,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "出力するファイル名",
				},
			},
		},
		{
			Name:   "read",
			Usage:  "読み込み",
			Action: actions.Read,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "入力するファイル名",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "書き出し",
			Action: actions.Create,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Usage: "出力するファイル名",
				},
			},
		},
	}
	app.Run(os.Args)
}
