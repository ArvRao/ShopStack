package models

type Admin struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"uniqueIndex;type:varchar(100);not null"`
}
