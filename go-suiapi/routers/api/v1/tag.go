package v1

import (
	"fmt"
	"gotest20230626/models"
	"gotest20230626/pkg/e"
	"gotest20230626/pkg/setting"
	"gotest20230626/pkg/util"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// add标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DelTag(c *gin.Context) {

}

// 上传文件
func Uploadfile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败："+err.Error())
		return
	}
	var uploadir string
	uploadir = "./../vuesui/src/files/" //"./files/"  "../../vscodepj2023/testvue/test20230625b/src/files/"
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
