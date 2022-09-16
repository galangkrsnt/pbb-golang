package main

import (
	"pajak-bumi-bangunan/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/inquiry", handler.InquiryHandler)
	v1.POST("/payment", handler.PaymentHandler)

	router.Run(":8888")
}
