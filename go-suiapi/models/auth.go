package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

func CheckAuth(username, password string) (int, bool, string, string) {
	var auth Auth
	var account Account
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		db.Select("name,photo").Where("Id=?", auth.ID).First(&account)
		return auth.ID, true, account.Name, account.Photo
	} else {
		return auth.ID, false, account.Name, account.Photo
	}
}
