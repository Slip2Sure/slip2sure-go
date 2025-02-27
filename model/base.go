package model

type Slip2SureCredit struct {
	Before float64 `json:"before"`
	Usage  float64 `json:"usage"`
	After  float64 `json:"after"`
}

type Slip2SureBase struct {
	Credit  Slip2SureCredit `json:"credit"`
	IsExist bool            `json:"is_exist"`
}
