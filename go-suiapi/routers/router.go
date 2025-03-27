package routers

import (
	jwt "gotest20230626/middleware"
	"gotest20230626/middleware/gcors"
	"gotest20230626/pkg/setting"
	api "gotest20230626/routers/api"
	"gotest20230626/routers/api/chatmsg"
	"gotest20230626/routers/api/other"
	v1 "gotest20230626/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gcors.Cors())

	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)
	r.POST("/api/upload", v1.Uploadfile)

	//推送消息
	r.GET("/pushmsg", v1.GetwsConnections)
	//推送消息
	r.GET("/chatmsg", chatmsg.ClientConnections)
	r.GET("/downSoftfile", v1.GetdownSoftfile)
	r.GET("/getUserlist", v1.GetUserlist)
	r.GET("/getpdfinfo", v1.Getpdffile)
	r.POST("/uploadpdffile", v1.UploadPdffile)

	//获取指定文章
	r.GET("/Readarticles/:id", v1.GetBlogArticle)
	r.POST("/Chatlocalsui", v1.ChatlocalAI)
	//加解密
	r.GET("/Getencrypt/:id", other.Getencrypt)
	r.GET("/Getdecrypt/:id", other.Getdecrypt)

	go v1.GetwsMessages()
	go chatmsg.ClientMessages()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DelTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//获取软件信息
		apiv1.GET("/softinfo", v1.GetSoftinfos)
		apiv1.GET("/GetSeachs", v1.GetSeachs)
		//获取书本信息
		apiv1.GET("/Bookinfo", v1.GetBookinfos)
	}
	return r
}
