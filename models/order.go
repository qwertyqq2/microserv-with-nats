package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID                uuid.UUID `gorm:"primary_key; unique;type:uuid;" json:"order_uid"`
	Delivery          Delivery  `gorm:"foreignKey:DeliveryID" json:"delivery"`
	Payment           Payment   `gorm:"foreignKey:PaymentID" json:"payment"`
	Items             []Item    `gorm:"many2many:order_items;" json:"items"`
	DeliveryID        int
	PaymentID         int
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shard_key"`
	SmID              int    `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OffShard          string `json:"off_shard"`
}

type Delivery struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	gorm.Model
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	gorm.Model
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type OrderRequest struct {
	ID uuid.UUID `json:"order_uid"`
}
