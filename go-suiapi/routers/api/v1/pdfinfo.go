package v1

import (
	"gotest20230626/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Getpdffile(c *gin.Context) {
	// 指定目录路径
	directoryPath := "./../vuesui/src/files/pdffile/"

	// 获取目录下的所有文件
	files, err := models.GetPDFFiles(directoryPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 返回文件列表
	c.JSON(http.StatusOK, gin.H{"pdfs": files})
}

// 上传文件
func UploadPdffile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败："+err.Error())
		return
	}
	var uploadir string
	uploadir = "./../vuesui/src/files/pdffile/" //"./files/"  "../../vscodepj2023/testvue/test20230625b/src/files/"
	_, err = os.Stat(uploadir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadir, os.ModePerm)
	}

	//fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	//newFileName := fileId + path.Ext(file.Filename)
	dst := uploadir + file.Filename //newFileName
	uplouderr := c.SaveUploadedFile(file, dst)

	if uplouderr != nil {
		//fmt.Println(uplouderr)
		//rsp := response.FailWithCodeAndData(response.Error)
		c.JSON(500, gin.H{"err": "file saving failed:" + uplouderr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "suceess",
		"name":      file.Filename, //file.Filename,
		"size":      file.Size,
		"uploadRes": uploadir,
	})

}
