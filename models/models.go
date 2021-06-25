package models

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  []byte `json:"-" `
}

//Product struct

type Product struct {
	ProductID          string `json:"product_id"`
	ImageUrlPublicID   string `json:"image_url_public_id"`
	ImageUrlSecureID   string `json:"image_url_secure_id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	Price              string `json:"price"`
}

type Products struct {
	Products []Product `json:"product"`
}

//Review struct

type Review struct {
	ReviewID  string `json:"review_id"`
	ProductID string `json:"product_id"`
	UserID    string `json:"user_id"`
	Reviews   string `json:"reviews"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type Reviews struct {
	Reviews []Review `json:"review"`
}

type StripePayment struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receipt_email"`
	ProductName  string `json:"product_name"`
	Street       string `json:"street"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	State        string `json:"state"`
	City         string `json:"city"`
	ZipCode      string `json:"zip_code"`
}
