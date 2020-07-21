/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udr/handler/message"
	"free5gc/src/udr/logger"

	"github.com/gin-gonic/gin"
)

// AmfContextNon3gpp - To modify the AMF context data of a UE using non 3gpp access in the UDR
func AmfContextNon3gpp(c *gin.Context) {
	var patchItemArray []models.PatchItem
	if err := c.ShouldBindJSON(&patchItemArray); err != nil {
		logger.DataRepoLog.Panic(err.Error())
	}

	req := http_wrapper.NewRequest(c.Request, patchItemArray)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := message.NewHandlerMessage(message.EventAmfContextNon3gpp, req)
	message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// CreateAmfContextNon3gpp - To store the AMF context data of a UE using non-3gpp access in the UDR
func CreateAmfContextNon3gpp(c *gin.Context) {
	var amfNon3GppAccessRegistration models.AmfNon3GppAccessRegistration
	if err := c.ShouldBindJSON(&amfNon3GppAccessRegistration); err != nil {
		logger.DataRepoLog.Panic(err.Error())
	}

	req := http_wrapper.NewRequest(c.Request, amfNon3GppAccessRegistration)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := message.NewHandlerMessage(message.EventCreateAmfContextNon3gpp, req)
	message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// QueryAmfContextNon3gpp - Retrieves the AMF context data of a UE using non-3gpp access
func QueryAmfContextNon3gpp(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := message.NewHandlerMessage(message.EventQueryAmfContextNon3gpp, req)
	message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}