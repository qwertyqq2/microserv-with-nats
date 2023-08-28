package controller

import "fmt"

const (
	CreateOrderSubject = "create_order"

	SenderID         = "sender"
	SubcriberOrderID = "subscriber_order_id"

	DefualtSubject = "topek"
)

func generateSubID() func() string {
	var i int
	return func() string {
		i++
		return fmt.Sprintf("%s%d", SubcriberOrderID, i)
	}
}
