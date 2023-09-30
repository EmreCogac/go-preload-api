package models

type Section struct {
	ID    uint8  `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"type:varchar(10);not null"`
	Desc  string `gorm:"type:varchar(20);not null"`
}

type Level struct {
	Id      uint8  `gorm:"primary_key"`
	Title   string `gorm:"type:varchar(20);not null"`
	Desc    string `gorm:"type:varchar(50);not null"`
	SecID   uint8
	Section Section `gorm:"foreignKey:SecID"`
}
