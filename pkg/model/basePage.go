package model

type BasePage struct {
	Page    int    `json:"page"`    // 页码
	Size    int    `json:"size"`    // 每页大小
	OrderBy string `json:"orderBy"` //排序字段，多个逗号隔开  username,password
	Sort    string `json:"sort"`    // desc 倒序
}
