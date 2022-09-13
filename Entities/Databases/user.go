package Databases

type User struct {
	Id   int    `json:"id" xorm:"not null pk autoincr INT"`
	Name string `json:"name" xorm:"not null default '' VARCHAR(255)"`
}
