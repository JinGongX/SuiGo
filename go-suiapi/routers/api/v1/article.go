package v1

import (
	"fmt"
	"gotest20230626/models"
	"gotest20230626/pkg/e"
	"gotest20230626/pkg/setting"
	"gotest20230626/pkg/util"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	userIP := c.ClientIP()
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			redisKey := fmt.Sprintf("article:%s:viewer:%s", id, userIP)
			if !models.RDBhasViewed(redisKey) {
				models.IncrementView(id)
				models.RDBmarkAsViewed(redisKey, 24*time.Hour)
			}

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

// 获取多个文章
func GetArticles(c *gin.Context) {
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

		data["lists"] = models.GetArticlesOrorby(util.GetPage(c), setting.PageSize, maps, "created_on DESC")
		data["total"] = models.GetArticleTotal(maps)

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

// 新增文章
func AddArticle(c *gin.Context) {
	//tagId := com.StrTo(c.Query("tag_id")).MustInt()
	//title := c.Query("title")
	//desc := c.Query("desc")
	//content := c.Query("content")
	//createdBy := c.Query("created_by")
	//state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	//imgcover := c.Query("imgcover")

	tagId := 1
	title := ""
	desc := ""
	content := ""
	createdBy := "ming"
	state := 1
	imgcover := ""
	//取data 的值
	form := &models.Article{}
	if c.Bind(form) == nil {
		//tagId = com.StrTo(rune(form.TagID)).MustInt()

		tagId = form.TagID
		content = form.Content
		title = form.Title
		desc = form.Desc
		createdBy = form.CreatedBy
		imgcover = form.Imgcover
		state = form.State
	}

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 2, "state").Message("状态只允许0或1、2")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		//if models.ExistTagByID(tagId) {
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state
		data["imgcover"] = imgcover

		models.AddArticle(data)
		code = e.SUCCESS
		//} else {
		//	code = e.ERROR_NOT_EXIST_TAG
		//}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改文章
func EditArticle(c *gin.Context) {

	// id := com.StrTo(c.Param("id")).MustInt()
	// tagId := com.StrTo(c.Query("tag_id")).MustInt()
	// title := c.Query("title")
	// desc := c.Query("desc")
	// content := c.Query("content")
	// modifiedBy := c.Query("modified_by")
	id := 1
	title := ""
	desc := ""
	content := ""
	modifiedBy := "1"
	imgcover := ""
	//state := 1

	form := &models.Article{}

	if c.Bind(form) == nil {
		//tagId = com.StrTo(rune(form.TagID)).MustInt()
		id = form.ID
		content = form.Content
		title = form.Title
		desc = form.Desc
		modifiedBy = form.ModifiedBy
		imgcover = form.Imgcover
	}

	// var state int = -1
	// if arg := c.Query("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
	// 	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	// }
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	//valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	// valid.Min(id, 1, "id").Message("ID必须大于0")
	// valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	// valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	// valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	// valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	// valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			//if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			// if tagId > 0 {
			// 	data["tag_id"] = tagId
			// }
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}

			data["modified_by"] = modifiedBy
			data["imgcover"] = imgcover
			models.EditArticle(id, data)
			code = e.SUCCESS
			//} else {
			//	code = e.ERROR_NOT_EXIST_TAG
			//}
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
		"data": make(map[string]string),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	//id := com.StrTo(c.Param("id")).MustInt()
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
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
		"data": make(map[string]string),
	})
}

func GetBlogArticle(c *gin.Context) {
	decryptedText, err := util.Decrypt(c.Param("id"))
	if err != nil {
		fmt.Println("Error encrypting text:", err)
	}
	id := com.StrTo(decryptedText).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	userIP := c.ClientIP()
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			redisKey := fmt.Sprintf("article:%s:viewer:%s", id, userIP)
			if !models.RDBhasViewed(redisKey) {
				models.IncrementView(id)
				models.RDBmarkAsViewed(redisKey, 24*time.Hour)
			}

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
