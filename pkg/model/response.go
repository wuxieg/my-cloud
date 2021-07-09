package model

type Response struct {
	Code int    `json:"code"`
	Msg  string `msg:""`
	// - 只用于调试的，前端不应使用该值，因为其可能会不存在
	Errors interface{} `json:"errors,omitempty"`
	// 具体响应数据
	// - 无数据时，不返回
	// - 如果是列表数据建议使用 ListData 类型
	// - 不需分页的列表类型时，也建议 data.list 这样响应
	Data interface{} `json:"data"`
}
