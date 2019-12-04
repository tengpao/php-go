package php_go

// https://github.com/henrylee2cn/pholcus/blob/master/common/util/util.go

import (
	"strconv"
	"strings"
	"unsafe"
)

// Bytes2String直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
// 转换之后若没做其他操作直接改变里面的字符，则程序会崩溃。
// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic。
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Atoa(str interface{}) string {
	if str == nil {
		return ""
	}
	return strings.Trim(str.(string), " ")
}

func Atoi(str interface{}) int {
	if str == nil {
		return 0
	}
	i, _ := strconv.Atoi(strings.Trim(str.(string), " "))
	return i
}

func Atoui(str interface{}) uint {
	if str == nil {
		return 0
	}
	u, _ := strconv.Atoi(strings.Trim(str.(string), " "))
	return uint(u)
}
