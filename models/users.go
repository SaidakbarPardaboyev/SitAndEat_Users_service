package models

type UserInfo struct {
	Id           string
	Username     string
	Password     string
	Phone_number string
}

type RefreshToken struct {
	UserId     string
	Token      string
	ExpiresAt int64
}
