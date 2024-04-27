package entity

type User struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
