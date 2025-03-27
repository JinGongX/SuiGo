package v1

import (
	"gotest20230626/models"
	"gotest20230626/pkg/e"
	"gotest20230626/pkg/setting"
	"gotest20230626/pkg/util"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取单个软件信息
func GetSoftinfo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistSoftinfoByID(id) {
			data = models.GetSoftinfo(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 获取多个软件信息
func GetSoftinfos(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	var createdBy string = "0"
	if createdor := c.Query("created_by"); createdor != "" {
		createdBy = createdor
		maps["created_by"] = createdBy

		valid.Required(createdBy, "created_by").Message("创建人必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetSoftinfosOrorby(util.GetPage(c), setting.PageSize, maps, "created_on DESC")
		data["total"] = models.GetSoftinfoTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetdownSoftfile(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	//var data interface{}
	var files string = ""
	if !valid.HasErrors() {
		if models.ExistSoftinfoByID(id) {
			filePath := models.GetSoftinfodown(id)
			code = e.SUCCESS
			files = filePath
			// 从文件路径中提取文件名
			fileName := filepath.Base(filePath)
			// 设置响应头信息
			c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
			c.Writer.Header().Add("Content-Type", "application/octet-stream")
			// 打开文件
			file, err := os.Open(filePath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer file.Close()
			// 将文件内容复制到响应体
			_, err = io.Copy(c.Writer, file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": files,
	})
}
