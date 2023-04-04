package order_api

import "time"

type OrderCreateRequest struct {
	UserId         string `json:"userId" bson:"userId"`
	Status         string `json:"status" bson:"status"`
	Address        string `json:"address" bson:"address"`
	InvoiceAddress string `json:"invoiceAddress" bson:"invoiceAddress"`
	Product        []struct {
		Name     string  `json:"name" bson:"name"`
		Quantity int     `json:"quantity" bson:"quantity"`
		Price    float64 `json:"price" bson:"price"`
	} `json:"product" bson:"product"`
}

// TODO: Dto da address id al ve model tarafında address response olarak ekle.

type OrderUpdateRequest struct {
	ID             string `json:"id" bson:"_id"`
	UserId         string `json:"userId" bson:"userId"`
	Status         string `json:"status" bson:"status"`
	Address        string `json:"address" bson:"address"`
	InvoiceAddress string `json:"invoiceAddress" bson:"invoiceAddress"`
	Product        []struct {
		Name     string  `json:"name" bson:"name"`
		Quantity int     `json:"quantity" bson:"quantity"`
		Price    float64 `json:"price" bson:"price"`
	} `json:"product" bson:"product"`
}

type OrderResponse struct {
	ID             string          `json:"id" bson:"_id"`
	UserId         string          `json:"userId" bson:"userId"`
	Status         string          `json:"status" bson:"status"`
	Address        AddressResponse `json:"address" bson:"address"`
	InvoiceAddress AddressResponse `json:"invoiceAddress" bson:"invoiceAddress"`
	Product        []struct {
		Name     string  `json:"name" bson:"name"`
		Quantity int     `json:"quantity" bson:"quantity"`
		Price    float64 `json:"price" bson:"price"`
	} `json:"product" bson:"product"`
	Total     float64   `json:"total" bson:"total"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type AddressResponse struct {
	ID       string   `json:"id"`
	Address  string   `json:"address" bson:"address"`
	City     string   `json:"city" bson:"city"`
	District string   `json:"district" bson:"district"`
	Type     []string `json:"type" bson:"type"`
	Default  struct {
		IsDefaultInvoiceAddress bool `json:"isDefaultInvoiceAddress" bson:"isDefaultInvoiceAddress"`
		IsDefaultRegularAddress bool `json:"isDefaultRegularAddress" bson:"isDefaultRegularAddress"`
	} `json:"default" bson:"default"`
}

type UserResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Addresses []AddressResponse `json:"address"`
}

type OrderResponseForElastic struct {
	OrderID string `json:"orderID" bson:"orderID"`
	Status  string `json:"status" bson:"status"`
}

// TODO: Bir kullanıcının her zaman max ve min sayı olarak 1 tane default adresi olması lazım (1 tane regular 1 tane invoice)
//
