package controller

import (
	"net/http"
	"payment/helper"
	"payment/service"
	"payment/web"

	"github.com/gin-gonic/gin"
)

type MidtransControllerImpl struct {
	MidtransService service.MidtransService
}

func NewMidtransControllerImpl(midtransService service.MidtransService) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		MidtransService: midtransService,
	}
}

func (controller *MidtransControllerImpl) Create(c *gin.Context) {
	var request web.MidtransRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.PanicIfError(err)
	}

	// Call the Midtrans service to create a transaction
	midtransResponse := controller.MidtransService.Create(c, request)
	var webResponse web.WebResponse

	// Check if the response contains a token (indicating success)
	if midtransResponse.Token != "" {
		webResponse = web.WebResponse{
			Code:   http.StatusOK,
			Status: "Payment Success",
			Data:   midtransResponse,
		}
	} else {
		// If there's no token, consider it a failure
		webResponse = web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Payment Failed",
			Data:   nil,
		}
	}

	c.JSON(webResponse.Code, webResponse)
}
