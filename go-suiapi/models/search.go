package models

// func GetArticleTotal(maps interface{}) (count int) {
// 	db.Model(&Article{}).Where(maps).Count(&count)

// 	return
// }

func GetSearchs(pageNum int, pageSize int, query string) (articles []Article) {
	db.Where("title LIKE ?", ""+query+"%").Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetSearchsOrorby(pageNum int, pageSize int, orderby string, query string) (articles []Article) {
	db.Where("title LIKE ?", ""+query+"%").Offset(pageNum).Limit(pageSize).Order(orderby).Find(&articles)
	//.Preload("Tag").Preload("Account")
	return
}

func GetSearch(id int) (article Article) {
	db.Preload("Account").Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}
