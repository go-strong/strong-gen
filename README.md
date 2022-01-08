
普通代码生成工具，从 (kratos)[https://go-kratos.dev/en/docs/getting-started/start/] 改成 strong ，符合自己平时需要的生成文档格式

win
```

git tag -a v0.0.5 -m "版本 0.0.5" && git push origin --tags

go get -u github.com/go-strong/strong-gen/tool/strong

cd /e/Dream/Plan2021_blockgo/src

strong new haha

cd haha

go mod tidy

cd cmd

go run main.go -conf saga-admin-test.toml 

```