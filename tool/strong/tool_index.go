package main

import "time"

var toolIndexs = []*Tool{
	{
		Name:      "strong",
		Alias:     "strong",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		//Install:   "go get -u github.com/go-strong/strong-gen/tool/strong@" + Version,
		Install:  "go get -u github.com/go-strong/strong-gen/tool/strong",
		Summary:  "Strong工具集本体",
		Platform: []string{"darwin", "linux", "windows"},
		Author:   "lucky",
		Hidden:   true,
	},
	//{
	//	Name:      "protoc",
	//	Alias:     "strong-protoc",
	//	BuildTime: time.Date(2022, 1 8, 0, 0, 0, 0, time.Local),
	//	Install:   "go get -u github.com/go-strong/strong-gen/tool/strong-protoc@" + Version,
	//	Summary:   "快速方便生成pb.go的protoc封装，windows、Linux请先安装protoc工具",
	//	Platform:  []string{"darwin", "linux", "windows"},
	//	Author:    "lucky",
	//},
	{
		Name:      "genbts",
		Alias:     "strong-gen-bts",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		//Install:   "go get -u github.com/go-strong/strong-gen/tool/strong-gen-bts@" + Version,
		Install:  "go get -u github.com/go-strong/strong-gen/tool/strong-gen-bts",
		Summary:  "缓存回源逻辑代码生成器",
		Platform: []string{"darwin", "linux", "windows"},
		Author:   "lucky",
	},
	{
		Name:      "genmc",
		Alias:     "strong-gen-mc",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		//Install:   "go get -u github.com/go-strong/strong-gen/tool/strong-gen-mc@" + Version,
		Install:  "go get -u github.com/go-strong/strong-gen/tool/strong-gen-mc",
		Summary:  "mc缓存代码生成",
		Platform: []string{"darwin", "linux", "windows"},
		Author:   "lucky",
	},
	{
		Name:      "genproject",
		Alias:     "strong-gen-project",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		//Install:      "go get -u github.com/go-strong/strong-gen/tool/strong-gen-project@" + Version,
		Install:      "go get -u github.com/go-strong/strong-gen/tool/strong-gen-project",
		Platform:     []string{"darwin", "linux", "windows"},
		Hidden:       true,
		Requirements: []string{"wire"},
	},
	{
		Name:      "testgen",
		Alias:     "testgen",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-strong/strong-gen/tool/testgen@" + Version,
		//Install:  "go get -u github.com/go-strong/strong-gen/tool/testgen",
		Summary:  "测试代码生成",
		Platform: []string{"darwin", "linux", "windows"},
		Author:   "lucky",
	},
	{
		Name:      "testcli",
		Alias:     "testcli",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-strong/strong-gen/tool/testcli@" + Version,
		//Install:  "go get -u github.com/go-strong/strong-gen/tool/testcli",
		Summary:  "测试代码运行",
		Platform: []string{"darwin", "linux", "windows"},
		Author:   "lucky",
	},
	//  third party
	{
		Name:      "wire",
		Alias:     "wire",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/google/wire/cmd/wire@v0.5.0",
		Platform:  []string{"darwin", "linux", "windows"},
		Hidden:    true,
	},
	{
		Name:      "swagger",
		Alias:     "swagger",
		BuildTime: time.Date(2022, 1, 8, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-swagger/go-swagger/cmd/swagger@v0.28.0",
		Summary:   "swagger api文档",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "goswagger.io",
	},
}
