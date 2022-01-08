package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var appHelpTemplate = `{{if .Usage}}{{.Usage}}{{end}}

USAGE:
  strong new {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
  {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
  {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
  {{range $index, $author := .Authors}}{{if $index}}
  {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

OPTIONS:
  {{range $index, $option := .VisibleFlags}}{{if $index}}
  {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

COPYRIGHT:
  {{.Copyright}}{{end}}
`

func main() {
	app := cli.NewApp()
	app.Name = ""
	app.Usage = "strong 新项目创建工具"
	app.UsageText = "项目名 [options]"
	app.HideVersion = true
	app.CustomAppHelpTemplate = appHelpTemplate
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "d",
			Value:       "",
			Usage:       "指定项目所在目录",
			Destination: &p.path,
		},
		&cli.BoolFlag{
			Name:        "http",
			Usage:       "只使用http 不使用grpc",
			Destination: &p.onlyHTTP,
		},
	}
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		app.Run([]string{"-h"})
		return
	}
	p.Name = os.Args[1]
	app.Action = runNew
	args := append([]string{os.Args[0]}, os.Args[2:]...)
	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}

func runNew(ctx *cli.Context) (err error) {
	if p.onlyGRPC && p.onlyHTTP {
		p.onlyGRPC = false
		p.onlyHTTP = false
	}
	if p.path != "" {
		if p.path, err = filepath.Abs(p.path); err != nil {
			return
		}
		p.path = filepath.Join(p.path, p.Name)
	} else {
		pwd, _ := os.Getwd()
		p.path = filepath.Join(pwd, p.Name)
	}
	p.ModPrefix = modPath(p.path)
	// creata a project
	if err := create(); err != nil {
		return err
	}
	fmt.Printf("Project: %s\n", p.Name)
	fmt.Printf("OnlyGRPC: %t\n", p.onlyGRPC)
	fmt.Printf("OnlyHTTP: %t\n", p.onlyHTTP)
	fmt.Printf("Directory: %s\n\n", p.path)
	fmt.Println("项目创建成功.")
	return nil
}

func modPath(p string) string {
	dir := filepath.Dir(p)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			content, _ := ioutil.ReadFile(filepath.Join(dir, "go.mod"))
			mod := RegexpReplace(`module\s+(?P<name>[\S]+)`, string(content), "$name")
			name := strings.TrimPrefix(filepath.Dir(p), dir)
			name = strings.TrimPrefix(name, string(os.PathSeparator))
			if name == "" {
				return fmt.Sprintf("%s/", mod)
			}
			return fmt.Sprintf("%s/%s/", mod, name)
		}
		parent := filepath.Dir(dir)
		if dir == parent {
			return ""
		}
		dir = parent
	}
}

// RegexpReplace replace regexp
// common.RegexpReplace(`module\s+(?P<name>[\S]+)`, string(content), "$name")
func RegexpReplace(reg, src, temp string) string {
	result := []byte{}
	pattern := regexp.MustCompile(reg)
	for _, submatches := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, submatches)
	}
	return string(result)
}
