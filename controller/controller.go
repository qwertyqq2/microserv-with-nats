package controller

import (
	"L0task/broker"
	"L0task/models"
	"L0task/store"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

type handler struct {
	ordersStore store.Store
	publisher   *broker.Publisher
	subs        []stan.Subscription
}

func Registration(r *gin.Engine, s store.Store, conf broker.Config) error {
	h := handler{ordersStore: s}

	log.Println("Connection to nats...")

	if err := broker.Connect(conf.URL); err != nil {
		return fmt.Errorf("err connect nat-server: %w", err)
	}

	log.Println("Creation publishers...")

	publisher, err := broker.NewPublisher(SenderID, conf.Subject, conf)
	if err != nil {
		return fmt.Errorf("err pulisher create: %w", err)
	}

	h.publisher = publisher

	log.Println("Creation subscribers...")

	genIds := generateSubID()

	for i := 0; i < 2; i++ {
		sub, err := broker.NewSubscriber(
			genIds(),
			conf.Subject,
			conf,
			h.handlerSubcriber(),
		)

		if err != nil {
			return err
		}

		h.subs = append(h.subs, sub)
	}

	r.POST("/create_order", h.createOrder)
	r.GET("/get_order", h.getOrder)
	r.GET("/all", h.getLastOrders)

	return nil
}

func (h *handler) handlerSubcriber() func(msg *stan.Msg) {
	return func(msg *stan.Msg) {
		var order models.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Println(fmt.Errorf("err unmarshal: %w", err))
			return
		}
		if err := h.ordersStore.CreateOrder(context.Background(), order); err != nil {
			log.Println(fmt.Errorf("err create order : %w", err))
			return
		}
		h.ordersStore.Cache().Add(order.ID, order, 0)

		time.Sleep(3 * time.Second)
	}
}
