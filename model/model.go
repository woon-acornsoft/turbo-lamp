package model

type User struct {
	tableName struct{} `pg:"user,alias:user"`
	Id        string   `json:"id"`
	Name      string   `json:"name"`
}
