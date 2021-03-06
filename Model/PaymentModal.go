package Model

import (
	"time"
)

type Payments struct {
	Id             int       `gorm:"primary_key; column:id" json:"id"`
	UserId         string    `gorm:"column:user_id" json:"user_id"`
	RazorpayLinkId string    `gorm:"column:razorpay_link_id" json:"razorpay_link_id"`
	RazorpayLink   string    `gorm:"column:razorpay_link" json:"razorpay_link"`
	Amount         int64   `gorm:"column:amount" json:"amount"`
	PaymentType    string    `gorm:"column:payment_type" json:"payment_type"`
	CurrentBalance int64  `gorm:"column:current_balance" json:"current_balance"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}


func (b *Payments) TableName() string {
	return "payments"
}