package models

type UserBaseModel struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	// Roles    []RoleBaseModel `json:"roles"`
}
type UserAuthModel struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserTokenModel struct {
	Token string `json:"token"`
}

type UserListBaseModel struct {
	Users []*UserBaseModel `json:"users"`
}
