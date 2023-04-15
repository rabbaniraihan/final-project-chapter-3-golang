package model

type User struct {
	Id          string        `gorm:"primaryKey" json:"id"`
	Username    string        `gorm:"not null;unique;type:varchar(255)" json:"username"`
	Email       string        `gorm:"not null;unique;type:varchar(255)" json:"email"`
	Password    string        `gorm:"not null;type:varchar(255)" json:"password"`
	Age         int           `gorm:"not null" json:"age"`
	Photo       []Photo       `gorm:"foreignKey:UserId" json:"user_photos"`
	Comment     []Comment     `gorm:"foreignKey:UserId" json:"user_comments"`
	SocialMedia []SocialMedia `gorm:"foreignKey:UserId" json:"user_social_media"`
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Id string `json:"id"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
