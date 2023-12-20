package data_models

type Course struct {
	ID      uint `gorm:"primaryKey"`
	Url     string
	Teacher string
	SubName string
}
