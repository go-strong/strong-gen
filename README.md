
普通代码生成工具，从 (kratos)[https://go-kratos.dev/en/docs/getting-started/start/] 改成 strong ，符合自己平时需要的生成文档格式

win
```
go get -u github.com/go-strong/strong-gen/tool/strong
go get -u github.com/go-strong/strong-gen/tool/strong@v0.0.5

go get -u github.com/go-strong/strong-gen/tool/strong-gen-project
go get -u github.com/go-strong/strong-gen/tool/strong-gen-project@v0.0.5

go get -u github.com/go-strong/strong-gen/tool/strong-gen-mc

go get -u github.com/go-strong/strong-gen/tool/testcli
go get -u github.com/go-strong/strong-gen/tool/testgen

go get -u github.com/go-strong/strong-gen/tool/strong-protoc

go get -u github.com/go-strong/strong-gen/tool/strong-gen-bts

go get -u github.com/golang/protobuf@v1.3.5

go get -u google.golang.org/protobuf@v1.3.5

go get -u github.com/google/wire/cmd/wire@v0.5.0  github.com/google/wire

log.Init(nil) // debug flag: log.dir={path}

go:generate strong



```