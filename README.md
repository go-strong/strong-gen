
普通代码生成工具，从 [kratos](https://go-kratos.dev/en/docs/getting-started/start) 改成 strong ，符合自己平时需要的生成文档格式

win
```

git tag -a v0.0.7 -m "版本 0.0.7" && git push origin --tags

go install github.com/go-strong/strong-gen/tool/strong@v0.0.7

cd /e/Dream/Plan2021_blockgo/src

$ strong new strong-saga-demo
$ cd strong-saga-demo/cmd/
$ go mod tidy
$ go run main.go -conf saga-admin-test.toml 

```