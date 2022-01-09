package pkg

import (
	"fmt"
	"testing"
)

func TestChina(t *testing.T) {
	str := "月色真美，风也温柔"                  //go字符串编码为utf-8
	fmt.Println("before convert:", str) //打印转换前的字符串

	fmt.Println("coding:", IsGBK(str)) //判断是否是utf-8

}
