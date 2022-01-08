
普通代码生成工具，从 (kratos)[https://go-kratos.dev/en/docs/getting-started/start/] 改成 strong ，符合自己平时需要的生成文档格式

![image](https://user-images.githubusercontent.com/17606795/148645370-4a45ceac-1a38-420f-a6c6-f493e5e9936d.png)


![image](https://user-images.githubusercontent.com/17606795/148645409-f1bc4be0-5a6d-4029-b635-b77e9894be3d.png)


![image](https://user-images.githubusercontent.com/17606795/148645395-eb9ed4b3-3e70-45ea-a80d-cbffa37013b1.png)


win
```

git tag -a v0.0.7 -m "版本 0.0.7" && git push origin --tags

go install github.com/go-strong/strong-gen/tool/strong@v0.0.7

cd /e/Dream/Plan2021_blockgo/src

strong new haha

$ cd haha/cmd/

go mod tidy

go run main.go -conf saga-admin-test.toml 

```
