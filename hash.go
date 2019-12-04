package php_go

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString MD5哈希值
func MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}

// SHA1Hash SHA1哈希值
func SHA1Hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString SHA1哈希值
func SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum(nil))
}

//MD5编码方式2
func Md5(data []byte) string {
	m := md5.Sum(data)
	return hex.EncodeToString(m[:])
}

//获取文件MD5值
func MD5File(file *os.File) string {
	_md5 := md5.New()
	_, _ = io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}
