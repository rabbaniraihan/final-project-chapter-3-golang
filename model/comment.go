package model

type Comment struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	UserId  int    `json:"user_id"`
	User    User   `gorm:"foreignKey:UserId" json:"user"`
	PhotoId int    `json:"photo_id"`
	Photo   Photo  `gorm:"foreignKey:PhotoId" json:"photo"`
	Message string `gorm:"not null;type:varchar(255)" json:"message"`
}
