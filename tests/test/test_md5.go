package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
//给password:123456加密
func GetMd5(pwd string)  string{
	h :=md5.New()
	h.Write([]byte(pwd))

	return hex.EncodeToString(h.Sum(nil))
}
func main() {
  md5_pwd :=GetMd5("123456")
  fmt.Println(md5_pwd)
}