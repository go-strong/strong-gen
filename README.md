
普通代码生成工具，从 [kratos](https://go-kratos.dev/en/docs/getting-started/start) 改成 strong ，符合自己平时需要的生成文档格式

![image](https://user-images.githubusercontent.com/17606795/148645370-4a45ceac-1a38-420f-a6c6-f493e5e9936d.png)

![image](https://user-images.githubusercontent.com/17606795/148645395-eb9ed4b3-3e70-45ea-a80d-cbffa37013b1.png)

![image](https://user-images.githubusercontent.com/17606795/148645454-26dad89a-ed01-4974-8092-5ce83986d55c.png)

![image](https://user-images.githubusercontent.com/17606795/148645471-295d87c1-8363-49b2-ac22-673c87f11ca9.png)



win
```
go install github.com/go-strong/strong-gen/tool/strong@v0.0.12
strong tool install all

cd /e/Dream/Plan2021_blockgo/src

$ strong new strong-saga-demo
$ cd strong-saga-demo/cmd/
$ go mod tidy
$ go run main.go -conf saga-admin-test.toml 

```
