package store

import (
	"L0task/models"
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type simpleDatabase struct {
	db *gorm.DB
}

func newDatabase() (*simpleDatabase, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Item{}, &models.Payment{})

	return &simpleDatabase{db: db}, nil
}

func (s *simpleDatabase) CreateOrder(ctx context.Context, order models.Order) error {
	res := s.db.WithContext(ctx).Create(&order)

	if res.Error != nil {
		return fmt.Errorf("err store: %w", res.Error)
	}

	return nil
}

func (s *simpleDatabase) GetOrder(ctx context.Context, orderUID uuid.UUID) (models.Order, error) {
	var order models.Order
	res := s.db.WithContext(ctx).Where("id=", orderUID).First(&order)

	if res.Error != nil {
		return models.Order{}, fmt.Errorf("err store: %w", res.Error)
	}

	return order, nil
}

func (s *simpleDatabase) GetLastOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order

	res := s.db.WithContext(ctx).Find(&orders)

	if res.Error != nil {
		return nil, fmt.Errorf("err store: %w", res.Error)
	}

	return orders, nil

}
