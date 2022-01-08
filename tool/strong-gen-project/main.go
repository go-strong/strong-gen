package main

import (
	"bytes"
	"fmt"
	_ "github.com/go-strong/strong-gen/tool/strong-gen-project/packrd"
	"github.com/gobuffalo/packr/v2"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
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

// project project config
type project struct {
	// project name
	Name string
	// mod prefix
	ModPrefix string
	// project dir
	path     string
	none     bool
	onlyGRPC bool
	onlyHTTP bool
}

var p project

//go:generate packr2
func create() (err error) {
	box := packr.New("http", "./templates/http")
	if err = os.MkdirAll(p.path, 0755); err != nil {
		return
	}
	for _, name := range box.List() {
		if p.ModPrefix != "" && name == "go.mod.tmpl" {
			//TODO continue
		}
		tmpl, _ := box.FindString(name)
		i := strings.LastIndex(name, string(os.PathSeparator))
		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(p.path, dir), 0755); err != nil {
				return
			}
		}
		if strings.HasSuffix(name, ".tmpl") {
			name = strings.TrimSuffix(name, ".tmpl")
		}
		if err = write(filepath.Join(p.path, name), tmpl); err != nil {
			return
		}
	}

	if err = generate("./..."); err != nil {
		return
	}
	/*if err = generate("./internal/dao/wire.go"); err != nil {
		return
	}*/
	return
}

func generate(path string) error {
	cmd := exec.Command("go", "generate", "-x", path)
	cmd.Dir = p.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func write(path, tpl string) (err error) {
	data, err := parse(tpl)
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, data, 0644)
}

func parse(s string) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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

//var _ = func() error {
//	const gk = "23f256f8ae67626d12a89a2afb9fd55c"
//	g := packr.New(gk, "")
//	hgr, err := resolver.NewHexGzip(map[string]string{
//		"eb20d555281a4332b042ab743d98d738": "1f8b08000000000000ff9452418bdb3c103ddbe0ff309f6141fe486cda63da9c966d37d086b449ce45516459acad091a39db25e4bf979152d8ec6ea1bdd8d2cc7ba3376fe620d583341a06695d9117b91d0ee8038822cfcab697a6e483b1a11b77b5c2a1e946e9cc0e6987d4189c766340ec1b7a2225fb3e81714ad2c8a9d1ae51e8da57c12e84c3ab20697fb44ac738d2e5d790354ea6b2c10e9cfd2b393db2ee8ad10a1da5767eec472f8345f709fd6d8f649d596b7fd41ee6f01e9a06482b74fb0baf1d9d8aa6880a4e459eb117f54a7ad2a22af2ccb6a0bd87d91cb8c57ae16c10d58718fb6f0ecef69194f568ea3befd18bf2198e61e8c5cdb12a277ce682d9413aabc4e5768e8d323b3222f7963f5fd0707eaf5bed81f3dc499294d02d8af274aa9772d0e73350903e94b1a18c58edc5e67aa91f23896791de20be2ac60cf2410bd5490748f53a4e6002ef389dc6512f31d8f649a8095ce65eaf179fefb7ababfbb7ed627315d8dc7dff7a15582c375cb4459fcc22cbaf7f9caadfcebdecc6e800129208b8a1720264eb75f0d619514513e9d106d501d9545149d2ffac69c6cc37dfd73f6d3433cb327a667c96f16ed6eb5eeb83f8f396fd0f0916b72cf1bc0ea3776f29bddfaea28ebd6ee5d887d90bf839edc8b9c87f050000ffff93d955a7c0030000",
//	})
//	if err != nil {
//		panic(err)
//	}
//	g.DefaultResolver = hgr
//
//	func() {
//		b := packr.New("http", "./templates/http")
//		b.SetResolver("cmd/main.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "eb20d555281a4332b042ab743d98d738"})
//	}()
//
//	return nil
//}()
