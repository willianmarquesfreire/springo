package domain

type User struct {
	GenericDomain      `bson:",inline"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}