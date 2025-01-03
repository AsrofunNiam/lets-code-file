package domain

type Product struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Price    int    `gorm:"not null" json:"price"`
	Quantity int    `gorm:"not null" json:"quantity"`
}
