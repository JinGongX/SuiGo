package models

type Bookinfo struct {
	Model

	TagID     int    `json:"tag_id" gorm:"index"`
	Tag       Tag    `json:"tag"`
	Title     string `json:"title"`
	Photoimg  string `json:"photoimg"`
	Filesbook string `json:"filesbook"`
	Label     string `json:"label"`
	Desc      string `json:"desc"`

	ModifiedBy string `json:"modified_by"`
	CreatedBy  string `json:"created_by"`
	State      int    `json:"state"`
	Author     string `json:"author"`
}

func ExistBookinfoByID(id int) bool {
	var bookinfo Bookinfo
	db.Select("id").Where("id = ?", id).First(&bookinfo)

	if bookinfo.ID > 0 {
		return true
	}

	return false
}

func GetBookinfoTotal(maps interface{}) (count int) {
	db.Model(&Bookinfo{}).Where(maps).Count(&count)

	return
}

func GetBookinfos(pageNum int, pageSize int, maps interface{}) (bookinfo []Bookinfo) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&bookinfo)
	return
}

func GetBookinfosOrorby(pageNum int, pageSize int, maps interface{}, orderby string) (bookinfo []Bookinfo) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Order(orderby).Find(&bookinfo)

	return
}

func GetBookinfo(id int) (bookinfo Bookinfo) {
	db.Where("id = ?", id).First(&bookinfo)
	db.Model(&bookinfo).Related(&bookinfo.Tag)
	return
}
