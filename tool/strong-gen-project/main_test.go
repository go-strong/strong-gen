package main

import (
	"bufio"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
	"github.com/huangbosbos/go-hutool/log"
	"github.com/huangbosbos/go-hutool/uuid"
	"os"
	"strings"
	"testing"
)

func TestDaoBranchSpe(t *testing.T) {
	createFile()
	genCmdMain := fileHexGzip("cmd/main.go.tmpl")
	genCmdSaga := fileHexGzip("cmd/saga-admin-test.toml")
	genConfConf := fileHexGzip("conf/conf.go.tmpl")
	genDaoDao := fileHexGzip("dao/dao.go.tmpl")
	genDaoTest := fileHexGzip("dao/dao_test.go.tmpl")
	genHttpHttp := fileHexGzip("http/http.go.tmpl")
	genModelModel := fileHexGzip("model/model.go.tmpl")
	genServiceService := fileHexGzip("service/service.go.tmpl")
	genServiceTest := fileHexGzip("service/service_test.go.tmpl")
	genGitignore := fileHexGzip(".gitignore.tmpl")
	genChangelog := fileHexGzip("CHANGELOG.md")
	genGoMod := fileHexGzip("go.mod.tmpl")
	genOwners := fileHexGzip("OWNERS")
	genReadme := fileHexGzip("README.md")

	m := make(map[string]string)
	m["eb20d555281a4332b042ab743d98d738"] = genCmdMain
	m["07b6987724e01017288e376e48d717a4"] = genCmdSaga
	m["0800ba3c38cf5e9bf163d71215211839"] = genConfConf
	m["09569e3bf810225e525835d677053b1b"] = genDaoDao
	m["0a6bc3e72bc7017e18eb65cbe3145be8"] = genDaoTest
	m["0a985fe02d21f9d997ee7c01e1f60f82"] = genHttpHttp
	m["0babbb2bdef8b27731b8c6e569f74aa9"] = genModelModel
	m["0c5ba5fa8837f2def63c05cc79c53394"] = genServiceService
	m["17e45e0f3497be1d7b615abf452cafa3"] = genServiceTest
	m["1a41024e60169dfd892e57d36264a7e3"] = genChangelog
	m["1dc54a258d1b22873bd8a065ef5f5136"] = genGoMod
	m["21390808875e3972b5fb30ef533f7595"] = genOwners
	m["240e2e99283e50af1f153a929e3d1116"] = genReadme
	m["f3a4d5b93f5048c5871bff2641d2db8b"] = genGitignore

	filePath := "./create/test.txt"
	WriteMapToFile(m, filePath)

	//log.Info("\"eb20d555281a4332b042ab743d98d738\":\"" + genCmdMain + "\",")
	//log.Info("\"07b6987724e01017288e376e48d717a4\":\"" + genCmdSaga + "\",")
	//log.Info("\"0800ba3c38cf5e9bf163d71215211839\":\"" + genConfConf + "\",")
	//log.Info("\"09569e3bf810225e525835d677053b1b\":\"" + genDaoDao + "\",")
	//log.Info("\"0a6bc3e72bc7017e18eb65cbe3145be8\":\"" + genDaoTest + "\",")
	//log.Info("\"0a985fe02d21f9d997ee7c01e1f60f82\":\"" + genHttpHttp + "\",")
	//log.Info("\"0babbb2bdef8b27731b8c6e569f74aa9\":\"" + genModelModel + "\",")
	//log.Info("\"0c5ba5fa8837f2def63c05cc79c53394\":\"" + genServiceService + "\",")
	//log.Info("\"17e45e0f3497be1d7b615abf452cafa3\":\"" + genServiceTest + "\",")
	//log.Info("\"1a41024e60169dfd892e57d36264a7e3\":\"" + genChangelog + "\",")
	//log.Info("\"1dc54a258d1b22873bd8a065ef5f5136\":\"" + genGoMod + "\",")
	//log.Info("\"21390808875e3972b5fb30ef533f7595\":\"" + genOwners + "\",")
	//log.Info("\"240e2e99283e50af1f153a929e3d1116\":\"" + genReadme + "\",")
	//log.Info("\"f3a4d5b93f5048c5871bff2641d2db8b\":\"" + genReadme + "\",")
	//log.Info("结束")

}
func createFile() {
	log.Init(nil)
	fmt.Println("Hello, world")
	path := "./create"

	// check
	if _, err := os.Stat(path); err == nil {
		fmt.Println("path exists 1", path)
	} else {
		fmt.Println("path not exists ", path)
		err := os.MkdirAll(path, 0711)
		if err != nil {
			log.Error("Error creating directory")
			panic(err)
			return
		}
	}

	// check again
	if _, err := os.Stat(path); err == nil {
		fmt.Println("path exists 2", path)
	}

	_, err := os.Create("./create/test.txt")
	if err != nil {
		panic(err)
	}
}

func TestWriteMaptoFile(t *testing.T) {
	m := make(map[string]string)
	m["21390808875e3972b5fb30ef533f7595"] = "1231321313"
	m["240e2e99283e50af1f153a929e3d1116"] = "5436546546"
	filePath := "./create/test.txt"

	WriteMapToFile(m, filePath)
}

func WriteMapToFile(m map[string]string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k, v := range m {
		lineStr := fmt.Sprintf("\"%s\":\"%s\",", k, v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}

func getUuid() string {
	str, _ := uuid.New()
	//fmt.Println("n=-1: ", strings.Replace(str, "-", "", -1))
	//fmt.Println("n=-1: ", len(strings.Replace(str, "-", "", -1)))
	return strings.Replace(str, "-", "", -1)
}
func TestUID(t *testing.T) {

	fmt.Println(getUuid())

}
func writeFile(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(content))
		panic(err)
	}

}

func fileHexGzip(filename string) string {
	box := packr.New("myBox", "./templates/http/")
	//读取文件
	s, err := box.FindString(filename)
	//文件压缩
	x, err := resolver.HexGzipString(s)
	//fmt.Println(x)
	//文件解密
	//y, err := resolver.UnHexGzipString(x)
	//fmt.Println("###########==> ", y)
	if err != nil {
		panic(err)
	}
	// fmt.Println(s)
	return x
}
