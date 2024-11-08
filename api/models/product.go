package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
	Price       float64
	CategoryID  uint
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
}
