package model

type SocialMedia struct {
	Id             int    `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null;type:varchar(255)" json:"name"`
	SocialMediaUrl string `gorm:"not null;type:varchar(255)" json:"social_media_url"`
	UserId         int    `json:"user_id"`
	User           User   `gorm:"foreignKey:UserId" json:"user"`
}
