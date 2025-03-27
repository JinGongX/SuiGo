package models

import "sort"

type Account struct {
	Model

	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Name_py string `json:"name_py"`
}
type Tapp_Account struct {
	Model

	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Name_py string `json:"name_py"`
}

// GroupedData 表示按字母分组后的数据结构
type GroupedData struct {
	Letter string         `json:"letter"`
	Data   []Tapp_Account `json:"data"`
}

func GetUserlist() []GroupedData {
	var tableData []Tapp_Account
	db.Find(&tableData) //db.Preload("Account")
	// 根据字母字段划分数据
	resultList := groupDataByLetter(tableData)

	return resultList
	//return new JsonResult(str);//await _context.Sys_User.ToListAsync();
}

// groupDataByLetter 根据字母字段划分数据
func groupDataByLetter(data []Tapp_Account) []GroupedData {
	groupedData := make(map[string][]Tapp_Account)
	for _, entry := range data {
		firstLetter := string(entry.Name_py[0])
		groupedData[firstLetter] = append(groupedData[firstLetter], entry)
	}

	var resultList []GroupedData
	for letter, entries := range groupedData {
		resultList = append(resultList, GroupedData{Letter: letter, Data: entries})
	}
	// 对字母进行排序
	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].Letter < resultList[j].Letter
	})
	return resultList
}

func GetFriendlist(userID int) []GroupedData {
	var tableData []Tapp_Account
	db.Joins("JOIN blog_friendship  ON user_id1 = id OR user_id2 = id").
		Where("(user_id1 = ? OR user_id2 = ?)", userID, userID).Select("DISTINCT blog_tapp_account.*").Find(&tableData) //db.Preload("Account")

	// 根据字母字段划分数据
	resultList := groupDataByLetter(tableData)

	return resultList
	//return new JsonResult(str);//await _context.Sys_User.ToListAsync();
}

func GetUserInfo(id int) (taccount Tapp_Account) {
	db.Where("id = ?", id).Find(&taccount) //db.Preload("Account")
	return
}

func GetGrouplist(userID int) []Conversationinfo {
	var tableData []Conversationinfo
	db.Joins("JOIN blog_chats  ON groupid = conversationid").
		Where("user1id = ? ", userID).Select("DISTINCT blog_conversationinfo.*").Find(&tableData) //db.Preload("Account")

	// 根据字母字段划分数据
	//resultList := groupDataByLetter(tableData)

	return tableData
	//return new JsonResult(str);//await _context.Sys_User.ToListAsync();
}

func EditUserInfo(id int, data interface{}) (taccount Tapp_Account) {
	//db.Where("id = ?", id).Find(&taccount) //db.Preload("Account")
	db.Model(&Tapp_Account{}).Where("id = ?", id).Updates(data)
	return
}
