package data_models

type History struct {
	ID       string `gorm:"primaryKey"`
	Distance string
	Emission int64
	UserID   string `gorm:""`
}
