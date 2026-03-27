package models

import "time"

type Client struct {
	ID          int       `json:"id"`
	ClientCode  string    `json:"clientCode"`
	CompanyName string    `json:"companyName"`
	Industry    string    `json:"industry"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Website     string    `json:"webPage"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	IsActive    bool      `json:"isActive"`
	CreditLimit float64   `json:"creditLimit"`
	CreatedAt   time.Time `json:"createdAt"`
}
