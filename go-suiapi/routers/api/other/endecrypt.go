package other

import (
	"fmt"
	"gotest20230626/pkg/e"
	"gotest20230626/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 加密
func Getencrypt(c *gin.Context) {
	id := com.StrTo(c.Param("id")).String()
	cipherText, err := util.Encrypt([]byte(id))
	code := e.SUCCESS
	if err != nil {
		fmt.Println("Error encrypting text:", err)
		code = e.INVALID_PARAMS
	}
	fmt.Println(cipherText)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": cipherText,
	})
}

// 解密
func Getdecrypt(c *gin.Context) {
	id := com.StrTo(c.Param("id")).String()
	decryptedText, err := util.Decrypt(id)
	code := e.SUCCESS
	if err != nil {
		fmt.Println("Error encrypting text:", err)
		code = e.INVALID_PARAMS
	}
	fmt.Println(decryptedText)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": decryptedText,
	})
}
