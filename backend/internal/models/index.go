package models

import (
	"time"
)

type Room struct {
	ID          int     `json:"id"`
	RoomNo      string  `json:"roomNo"`
	RoomWidth   float32 `json:"roomWidth"`
	RoomLength  float32 `json:"roomLength"`
	DoubleBeds  int16   `json:"doubleBeds"`
	SingleBeds  int16   `json:"singleBeds"`
	Windows     int16   `json:"windows"`
	AC          bool    `json:"aC"`
	Wifi        bool    `json:"wifi"`
	HotWater    bool    `json:"hotWater"`
	Balcony     bool    `json:"balcony"`
	Location    string  `json:"location"`
	RoomCharges float64 `json:"roomCharges"`
	RoomImage   *string `json:"roomImage"`
	Remarks     string  `json:"remarks"`
	Status      string  `json:"status"`
}

type Client struct {
	ID         int     `json:"id"`
	ClientName string  `json:"clientName"`
	CNIC       string  `json:"cnic"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Discount   float32 `json:"discount"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	UserRole string `json:"userRole"`
}

type Reservation struct {
	ID           int        `json:"id"`
	IDCustomer   int        `json:"idCustomer"`
	IDRoom       int        `json:"idRoom"`
	BookingStart *time.Time `json:"bookingStart"`
	BookingEnd   *time.Time `json:"bookingEnd"`
	ExtraCharges float64    `json:"extraCharges"`
	AmountPaid   float64    `json:"amountPaid"`
	ReservedAt   *time.Time `json:"reservedAt"`
	IDReservedBy int        `json:"idReservedBy"`
	Status       string     `json:"status"`
}

type BookingDetails struct {
	ID           int     `json:"id"`
	CustomerName string  `json:"customerName"`
	RoomNo       string  `json:"roomNo"`
	BookingStart *string `json:"bookingStart"`
	BookingEnd   *string `json:"bookingEnd"`
	ExtraCharges float64 `json:"extraCharges"`
	AmountPaid   float64 `json:"amountPaid"`
	ReservedAt   *string `json:"reservedAt"`
	ReservedBy   string  `json:"reservedBy"`
	Status       string  `json:"status"`
	IDCustomer   int     `json:"idCustomer"`
	IDRoom       int     `json:"idRoom"`
	IDReservedBy int     `json:"idReservedBy"`
}
