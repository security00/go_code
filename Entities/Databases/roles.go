package Databases

type Roles struct {
	Id   int `json:"id" xorm:"not null pk autoincr INT"`
	Role int `json:"role" xorm:"not null default 0 INT"`
}
