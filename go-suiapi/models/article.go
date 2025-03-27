package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	CreatedBy string  `json:"created_by"`
	Account   Account `json:"account" gorm:"foreignkey:Id;association_foreignkey:created_by"`

	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	//CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
	Imgcover   string `json:"imgcover"`
	Views      int    `json:"views" gorm:"default:0"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticlesOrorby(pageNum int, pageSize int, maps interface{}, orderby string) (articles []Article) {
	db.Preload("Tag").Preload("Account").Where(maps).Offset(pageNum).Limit(pageSize).Order(orderby).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Preload("Account").Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
		Imgcover:  data["imgcover"].(string),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func IncrementView(articleID int) {
	db.Model(&Article{}).Where("id = ?", articleID).Update("views", gorm.Expr("views + ?", 1))
}
