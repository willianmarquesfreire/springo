package domain

type User struct {
	GenericDomain      `bson:",inline"`
	Login       string `json:"login" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"newPassword"`
}
