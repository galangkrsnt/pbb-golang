package model

type InquiryResponse struct {
	Type            string `json:"type" binding:"required"`
	Service         string `json:"service" binding:"required"`
	AccountType     string `json:"accountType" binding:"required"`
	Account         string `json:"account" binding:"required"`
	InstitutionCode string `json:"institutionCode" binding:"required"`
	Product         string `json:"product" binding:"required"`
	BillNumber      string `json:"billNumber" binding:"required"`
	ResultCode      string `json:"resultCode" binding:"required"`
	ResultDesc      string `json:"resultDesc" binding:"required"`
	Nominal         string `json:"nominal" binding:"required"`
	Admin           string `json:"admin" binding:"required"`
	TrxId           string `json:"trxId" binding:"required"`
	Retrieval       string `json:"retrieval" binding:"required"`
	PaymentCode     string `json:"paymentCode" binding:"required"`
	Info1           string `json:"info1" binding:"required"`
	Info2           string `json:"info2" binding:"required"`
	Info3           string `json:"info3" binding:"required"`
	JmlLembar       string `json:"jmlLembar" binding:"required"`
	Saldo           string `json:"saldo" binding:"required"`
}
