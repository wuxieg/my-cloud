package model

type BasePage struct {
	Page    int    `json:"page" `    // 页码
	Size    int    `json:"size" `    // 每页大小
	OrderBy string `json:"orderBy" ` //排序字段，多个逗号隔开  username,password
	Sort    string `json:"sort" `    // desc 倒序
}

var (
	DEFAULT_ORDERBY_COLUMN = "created_at"
	EMPTY                  = ""
	DEFAULT_SORT           = "desc"
)

func (o *BasePage) DefaultPage() *BasePage {
	if o == nil {
		o = &BasePage{
			Size:    10,
			OrderBy: "created_at",
			Sort:    "desc",
		}
		return o
	}
	if o.Page < 0 {
		o.Page = 0
	}
	if o.OrderBy == EMPTY {
		o.OrderBy = DEFAULT_ORDERBY_COLUMN
	}
	if o.Sort == EMPTY {
		o.Sort = DEFAULT_SORT
	}
	return o
}
