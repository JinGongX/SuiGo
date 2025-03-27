package models

type Softinfo struct {
	Model

	TagID    int    `json:"tag_id" gorm:"index"`
	Tag      Tag    `json:"tag"`
	Title    string `json:"title"`
	Photoimg string `json:"photoimg"`
	Filesimg string `json:"filesimg"`
	Label    string `json:"label"`
	Desc     string `json:"desc"`

	ModifiedBy string `json:"modified_by"`
	CreatedBy  string `json:"created_by"`
	State      int    `json:"state"`
}

func ExistSoftinfoByID(id int) bool {
	var softinfo Softinfo
	db.Select("id").Where("id = ?", id).First(&softinfo)

	if softinfo.ID > 0 {
		return true
	}

	return false
}

func GetSoftinfoTotal(maps interface{}) (count int) {
	db.Model(&Softinfo{}).Where(maps).Count(&count)

	return
}

func GetSoftinfos(pageNum int, pageSize int, maps interface{}) (softinfo []Softinfo) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&softinfo)
	//
	return
}

func GetSoftinfosOrorby(pageNum int, pageSize int, maps interface{}, orderby string) (softinfo []Softinfo) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Order(orderby).Find(&softinfo)
	//.Preload("Tag").Preload("Account")
	return
}

func GetSoftinfo(id int) (softinfo Softinfo) {
	db.Where("id = ?", id).First(&softinfo)
	db.Model(&softinfo).Related(&softinfo.Tag)
	//.Preload("Account")
	return
}

func GetSoftinfodown(id int) (fileurl string) {
	//db.Where("id = ?", id).First(&softinfo.filesimg)
	var softinfo Softinfo
	db.Select("filesimg").Where("Id=?", id).First(&softinfo)
	fileurl = softinfo.Filesimg
	return
}
