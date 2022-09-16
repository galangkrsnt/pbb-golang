package model

type InquiryRequest struct {
	Type            string `json:"type" binding:"required"`
	AccountType     string `json:"accountType" binding:"required"`
	Account         string `json:"account" binding:"required"`
	InstitutionCode string `json:"institutionCode" binding:"required"`
	BillNumber      string `json:"billNumber" binding:"required"`
	Product         string `json:"product" binding:"required"`
	TrxId           string `json:"trxId" binding:"required"`
	Retrieval       string `json:"retrieval" binding:"required"`
}
