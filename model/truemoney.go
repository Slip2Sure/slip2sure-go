package model

import "time"

type Slip2SureTruemoney struct {
	Slip2SureBase
	Info Slip2SureTruemoneyInfo `json:"info"`
}

type Slip2SureTruemoneyInfo struct {
	TransactionID     string    `json:"transaction_id"`
	PaidAt            time.Time `json:"paid_at"`
	Amount            float32   `json:"amount"`
	AccountFromMobile *string   `json:"account_from_mobile"`
	AccountToMobile   string    `json:"account_to_mobile"`
}
