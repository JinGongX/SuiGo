package v1

import (
	"gotest20230626/models"
	"gotest20230626/pkg/e"
	"gotest20230626/pkg/setting"
	"gotest20230626/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 获取搜索多个文章
func GetSeachs(c *gin.Context) {
	data := make(map[string]interface{})
	//maps := make(map[string]interface{})
	valid := validation.Validation{}

	//var state int = -1
	// if arg := c.Query("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
	// 	maps["state"] = state

	// 	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	// }

	// var tagId int = -1
	// if arg := c.Query("tag_id"); arg != "" {
	// 	tagId = com.StrTo(arg).MustInt()
	// 	maps["tag_id"] = tagId

	// 	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	// }

	var title string = "0"
	if titledesc := c.Query("title"); titledesc != "" {
		title = titledesc
		//maps["created_by"] = title

		//valid.Required(title, "title").Message("创建人必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetSearchsOrorby(util.GetPage(c), setting.PageSize, "created_on DESC", title)
		//data["total"] = models.GetArticleTotal(maps)

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
