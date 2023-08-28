package store

import (
	"L0task/models"
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Cache() OrdersCache
	OrdersDatabase
}

type OrdersDatabase interface {
	CreateOrder(ctx context.Context, order models.Order) error
	GetOrder(ctx context.Context, orderUID uuid.UUID) (models.Order, error)
	GetLastOrders(ctx context.Context) ([]models.Order, error)
}

type OrdersCache interface {
	Add(key uuid.UUID, order models.Order, expiryTime time.Duration)
	Get(key uuid.UUID) (models.Order, bool)
}

type simpleStore struct {
	db    OrdersDatabase
	cache OrdersCache
	lk    sync.RWMutex
}

func NewStore(ctx context.Context) (Store, error) {
	db, err := newDatabase()
	if err != nil {
		return nil, err
	}

	tcxt, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	orders, err := db.GetLastOrders(tcxt)
	if err != nil {
		return nil, err
	}

	log.Println("Cache loading...")

	cache, err := newOrdersCache(tcxt, orders...)
	if err != nil {
		return nil, err
	}

	return &simpleStore{db: db, cache: cache}, nil
}

func (s *simpleStore) CreateOrder(ctx context.Context, order models.Order) error {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.db.CreateOrder(ctx, order)
}

func (s *simpleStore) GetOrder(ctx context.Context, orderUID uuid.UUID) (models.Order, error) {
	s.lk.RLock()
	defer s.lk.RUnlock()

	return s.db.GetOrder(ctx, orderUID)
}

func (s *simpleStore) GetLastOrders(ctx context.Context) ([]models.Order, error) {
	s.lk.RLock()
	defer s.lk.RUnlock()

	return s.db.GetLastOrders(ctx)
}

func (s *simpleStore) Cache() OrdersCache { return s.cache }
