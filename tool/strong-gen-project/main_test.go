package main

import (
	"bufio"
	"fmt"
	"github.com/go-strong/strong-gen/tool/pkg"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
	"github.com/huangbosbos/go-hutool/log"
	"github.com/huangbosbos/go-hutool/uuid"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestChina(t *testing.T) {
	genCn := fileHexGzip("api/http.go.tmpl")
	fmt.Println("###########==> ", genCn)
	//文件解密
	y, err := resolver.UnHexGzipString(genCn)
	fmt.Println("###########==> ", y)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------------------")
	genCn2 := fileHexGzip("api/main.go.tmpl")
	fmt.Println("###########==> ", genCn2)
	//文件解密
	y2, err := resolver.UnHexGzipString(genCn2)
	fmt.Println("###########==> ", y2)
	if err != nil {
		panic(err)
	}
}

func TestDaoBranchSpe(t *testing.T) {
	createFile()
	m := make(map[string]string)

	genCmdMain := fileHexGzip("cmd/main.go.tmpl")
	genCmdSaga := fileHexGzip("cmd/saga-admin-test.toml")
	genConfConf := fileHexGzip("conf/conf.go.tmpl")

	//新增biz 2022年1月9日14:00:57
	genDaoBusiness := fileHexGzip("dao/business.go.tmpl")
	genDaoBusinessTest := fileHexGzip("dao/business_test.go.tmpl")
	genDaoProject := fileHexGzip("dao/db_project.go.tmpl")
	genDaoProjectTest := fileHexGzip("dao/db_project_test.go.tmpl")
	genDaoMc := fileHexGzip("dao/mc.go.tmpl")
	genDaoMcTest := fileHexGzip("dao/mc_test.go.tmpl")
	genDaoRedis := fileHexGzip("dao/redis.go.tmpl")
	genDaoRedisTest := fileHexGzip("dao/redis_test.go.tmpl")
	//add http 2022年1月9日14:04:57
	genHttpBusiness := fileHexGzip("http/business.go.tmpl")
	genModelBusiness := fileHexGzip("model/business.go.tmpl")
	genModelModelTest := fileHexGzip("model/model_test.go.tmpl")
	genServiceBusiness := fileHexGzip("service/business.go.tmpl")
	genServiceBusinessTest := fileHexGzip("service/business_test.go.tmpl")
	genContributors := fileHexGzip("CONTRIBUTORS.md")
	//add http 2022年1月8日11:04:57
	genDaoDao := fileHexGzip("dao/dao.go.tmpl")
	genDaoTest := fileHexGzip("dao/dao_test.go.tmpl")
	genHttpHttp := fileHexGzip("http/http.go.tmpl")
	genModelModel := fileHexGzip("model/model.go.tmpl")
	genServiceService := fileHexGzip("service/service.go.tmpl")
	genServiceTest := fileHexGzip("service/service_test.go.tmpl")
	genGitignore := fileHexGzip(".gitignore.tmpl")
	genGoMod := fileHexGzip("go.mod.tmpl")
	genGoSum := fileHexGzip("go.sum.tmpl")
	genChangelog := fileHexGzip("CHANGELOG.md")
	genOwners := fileHexGzip("OWNERS")
	genReadme := fileHexGzip("README.md")

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
	m["2db2998574594dab822b64f6c78b099d"] = genGoSum
	m["21390808875e3972b5fb30ef533f7595"] = genOwners
	m["240e2e99283e50af1f153a929e3d1116"] = genReadme
	m["f3a4d5b93f5048c5871bff2641d2db8b"] = genGitignore
	//新增biz
	m["a893bbdc5e8b4a89b6de26f8e4523a40"] = genDaoBusiness
	m["9738f20ce1ec45b984af5f7b11d6853e"] = genDaoBusinessTest
	m["414b5c01cfc242e29a23deac0d866bd5"] = genDaoProject
	m["39e7595b0d324b89968d542f72386908"] = genDaoProjectTest
	m["8dd41504fc224ca2a4fc52fedc37cb81"] = genDaoMc
	m["eccd21683c9c4b6e8648af2317d1ba59"] = genDaoMcTest
	m["8d4fad2f0e5c4e1389db2f84577407b8"] = genDaoRedis
	m["a209b29b1ebd488083c064b49eb68ab5"] = genDaoRedisTest
	//add http
	m["0fcae10d97df4c22b110ca5ebbef88e0"] = genHttpBusiness
	m["822d864904ff454fb7d470dbb5a8421e"] = genModelBusiness
	m["57320f024ff14a3cb327ef178a647e23"] = genModelModelTest
	m["3fbfa2843f6d44858557fed167dbcb2c"] = genServiceBusiness
	m["783c6ee045574b13a72aa758b5f12509"] = genServiceBusinessTest
	m["07c7b07f56e846f2bb62f22aa112c67a"] = genContributors

	filePath := "./create/test.txt"
	WriteMapToFile(m, filePath)
	//log.Info(" ")

}

func TestUID(t *testing.T) {
	for i := 0; i < 6; i++ {
		//fmt.Println(getUuid())
		//fmt.Println("m[\"" + getUuid() + "\"] = genDemo_%d" ,i)
		fmt.Println("m[\"" + getUuid() + "\"] = genDemo_" + strconv.Itoa(i))
	}
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
	str, err := box.FindString(filename)

	isGbk := pkg.IsGBK(str)
	//fmt.Println(filename+"is GBK coding:", isGbk) //判断是否是gbk
	//fmt.Println("Is UTF8 coding:", pkg.IsUTF8(str)) //判断是否是utf8
	if isGbk {
		str = pkg.ConvertToString(str, "gbk", "utf-8")
	}
	//文件压缩
	x, err := resolver.HexGzipString(str)
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
