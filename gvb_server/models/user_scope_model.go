package models

type UserScopeModel struct {
	MODEL
	UserId uint `json:"user_id"`
	Scope  int  `json:"scope"`
	Status bool `json:"status"`
}
