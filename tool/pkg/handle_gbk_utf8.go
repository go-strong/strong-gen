package pkg

import (
	"github.com/axgle/mahonia"
	"unicode/utf8"
)

// GO代码实现判断字符编码格式及编码格式转换（utf-8、gbk）
// https://studygolang.com/articles/24814

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// IsUTF8 是否utf8
// golang学习之判断字符串编码是否是UTF8或GBK
// https://yoby123.cn/2021/03/11/2021-3-11-2.html
func IsUTF8(s string) bool {
	return utf8.ValidString(s)
}

func IsGBK(s string) bool {
	if IsUTF8(s) {
		return false
	}
	data := []byte(s)
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
