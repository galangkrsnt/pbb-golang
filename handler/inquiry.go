package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"pajak-bumi-bangunan/model"

	"github.com/gin-gonic/gin"
)

func InquiryHandler(c *gin.Context) {

	var inquiryRequest model.InquiryRequest
	var inquiryRequestFromClient model.InquiryRequest
	var inquiryResponse model.InquiryResponse

	err := c.ShouldBindJSON(&inquiryRequestFromClient)

	if err != nil {
		log.Fatalln(err)
	}

	inquiryRequest.TrxId = inquiryRequestFromClient.TrxId
	inquiryRequest.Account = inquiryRequestFromClient.Account
	inquiryRequest.AccountType = inquiryRequestFromClient.AccountType
	inquiryRequest.InstitutionCode = inquiryRequestFromClient.InstitutionCode
	inquiryRequest.Product = inquiryRequestFromClient.Product
	inquiryRequest.Retrieval = inquiryRequestFromClient.Retrieval
	inquiryRequest.BillNumber = inquiryRequestFromClient.BillNumber
	inquiryRequest.Type = inquiryRequestFromClient.Type

	jsonReq, err := json.Marshal(inquiryRequest)

	response, err := http.Post("https://apidev.banksinarmas.com/internal/payment/pbb/v1.0/inquiry", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	response.Header.Set("x-Gateway-APIKey", "59db964c-172c-463d-b745-8a46c140bb0d")
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(bodyBytes, &inquiryResponse)
	c.JSON(http.StatusOK, inquiryResponse)
}
