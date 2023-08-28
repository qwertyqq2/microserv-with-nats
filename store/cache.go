package store

import (
	"L0task/models"
	"context"
	"time"

	"github.com/google/uuid"
)

func newOrdersCache(ctx context.Context, orders ...models.Order) (OrdersCache, error) {
	cache := NewTTLCache[uuid.UUID, models.Order](200*time.Second, 5*time.Second)

	for _, order := range orders {
		cache.Add(order.ID, order, 0)
	}

	return cache, nil

}
