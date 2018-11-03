package dto

type GoogleUserInfo struct {
	ID             string `json:"sub"`
	Email          string `json:"email"`
	ProfilePicture string `json:"picture"`
}
