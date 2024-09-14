package entities

type {{Capitalize .ResourceName}} struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
}