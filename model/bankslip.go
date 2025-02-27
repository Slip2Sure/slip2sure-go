package model

import "time"

type Slip2SureBankSlip struct {
	Slip2SureBase
	Info Slip2SureBankSlipInfo `json:"info"`
}

type Slip2SureBankSlipInfo struct {
	Payload           string                   `json:"payload"`
	TransRef          string                   `json:"transRef"`
	Ref1              *string                  `json:"ref1"`
	Ref2              *string                  `json:"ref2"`
	Ref3              *string                  `json:"ref3"`
	TransDate         string                   `json:"transDate"`
	TransTime         string                   `json:"transTime"`
	TransDateTime     time.Time                `json:"transDateTime"`
	Sender            Slip2SureBankSlipAccount `json:"sender"`
	Receiver          Slip2SureBankSlipAccount `json:"receiver"`
	Amount            int                      `json:"amount"`
	TransFeeAmount    int                      `json:"transFeeAmount"`
	PaidLocalAmount   int                      `json:"paidLocalAmount"`
	PaidLocalCurrency string                   `json:"paidLocalCurrency"`
	CountryCode       string                   `json:"countryCode"`
	ToMerchantID      string                   `json:"toMerchantId"`
}

type Slip2SureBankSlipAccount struct {
	DisplayName string                       `json:"displayName"`
	Proxy       Slip2SureBankSlipAccountInfo `json:"proxy"`
	Account     Slip2SureBankSlipAccountInfo `json:"account"`
}

type Slip2SureBankSlipAccountInfo struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}
