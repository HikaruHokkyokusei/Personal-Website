package models

import (
	"time"
)

type Listing struct {
	ListingId  string    `json:"listing_id"`
	ItemName   string    `json:"item_name"`
	Category   string    `json:"category"`
	Price      float64   `json:"price"`
	Currency   string    `json:"currency"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
	SellDate   time.Time `json:"sell_date"`
	UserId     string    `json:"user_id"`
}
