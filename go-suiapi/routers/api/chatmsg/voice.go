package chatmsg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Getvoiceinfo(c *gin.Context) {
	filePathname := c.Query("filename")
	if filePathname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'path' query parameter"})
		return
	}

	// 获取音频文件路径
	//filePath := "/path/to/your/audio/file.mp3" // 你的音频文件路径
	filePath := "files/voice/" + filePathname
	//fmt.Println(filePath)
	//// 检查文件是否存在
	//_, err := os.Stat(filePath)
	//if os.IsNotExist(err) {
	//	fmt.Println(err)
	//	c.JSON(http.StatusNotFound, gin.H{"error": "Audio file not found"})
	//	return
	//}

	// 构建返回的URL
	//baseURL := "http://192.168.1.7:2023" // 你的域名或服务器地址
	//fileURL := filepath.Join(baseURL, "audio/file.mp3")
	//fileURL := fmt.Sprintf("http://%s/%s", c.Request.Host, filePath)
	// 返回URL
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read audio file")
		return
	}

	c.Data(http.StatusOK, "audio/mpeg", data)

	//c.JSON(http.StatusOK, gin.H{"url": fileURL})
}

// 上传文件
func Uploadvoicefile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败："+err.Error())
		return
	}
	//var uploadir string
	uploadir := "./files/voice/" //"./../vuesui/src/files/" //"./files/"  "../../vscodepj2023/testvue/test20230625b/src/files/"
	_, err = os.Stat(uploadir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadir, os.ModePerm)
	}

	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := uploadir + newFileName
	uplouderr := c.SaveUploadedFile(file, dst)

	if uplouderr != nil {
		fmt.Println(uplouderr)
		//rsp := response.FailWithCodeAndData(response.Error)
		c.JSON(500, gin.H{"err": "file saving failed:" + uplouderr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "suceess",
		"name":      newFileName, //file.Filename,
		"size":      file.Size,
		"uploadRes": uploadir,
	})

}
