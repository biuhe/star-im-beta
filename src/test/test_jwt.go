package main

import "star-im/src/utils"

func main() {
	// 加密
	genToken, err := utils.GenToken("张三")
	if err != nil {
		println("芜湖:", err)
	}

	println(genToken)

	// 解密
	token, _ := utils.ParseToken(genToken)
	println(token.Username)

}
