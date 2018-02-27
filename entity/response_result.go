package entity

type ResponseResult struct {
	Result 	bool		`json:"result"`
	Data 	interface{}	`json:"data"`
}
