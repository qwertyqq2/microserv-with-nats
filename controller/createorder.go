package controller

import (
	"L0task/models"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) createOrder(c *gin.Context) {
	var myOrderRequest models.Order

	jsonDataBytes, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(jsonDataBytes, &myOrderRequest); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "incorrect unmrsh")
		return
	}

	if err := h.publisher.Publish(jsonDataBytes); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "принято в работу")
}
