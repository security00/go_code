package Databases

import "mygo/Configs"

type Roles struct {
	Id   string `xorm:"not null pk autoincr UNSIGNED INT", json:"id"`
	Role int    `xorm:"not null default 0 INT", json:"role"`
}

type User struct {
	Id   string `xorm:"not null pk autoincr UNSIGNED INT"`
	Name string `xorm:"not null default '' VARCHAR(255)"`
}

func (r *Roles) InsertData(data interface{}) (int64, error) {
	ro := new(Roles)
	ro.Id = "1"
	ro.Role = 100
	res, err := Configs.Eg.Insert(ro)
	return res, err
}
