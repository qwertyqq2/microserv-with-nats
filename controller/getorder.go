package controller

import (
	"L0task/models"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) getOrder(c *gin.Context) {
	var newOrderRequest models.OrderRequest
	if err := c.ShouldBindJSON(&newOrderRequest); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "incorrect request")
		return
	}

	log.Println("Get from cache...")
	response, found := h.ordersStore.Cache().Get(newOrderRequest.ID)
	if found {
		c.JSON(http.StatusOK, response)
		return
	}

	log.Println("Get from db...")
	order, err := h.ordersStore.GetOrder(c, newOrderRequest.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	h.ordersStore.Cache().Add(order.ID, order, 0)

}

func (h *handler) getLastOrders(c *gin.Context) {
	orders, err := h.ordersStore.GetLastOrders(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	c.JSON(http.StatusOK, orders)
}
