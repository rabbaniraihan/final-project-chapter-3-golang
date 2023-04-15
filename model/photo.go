package model

type Photo struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"not null" json:"title"`
	Caption  string `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl string `gorm:"not null;type:varchar(255)" json:"photo_url"`
	UserId   int    `json:"user_id"`
	// User     User      `gorm:"foreignKey:UserId" json:"user"`
	Comment []Comment `gorm:"foreignKey:PhotoId" json:"photo_comments"`
}
