package models

type ProductImage struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"type:varchar(255);not null"`
	ProductID uint   `gorm:"not null"`
}
