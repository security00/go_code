package models

import (
	"mygo/Configs"
)

type User struct {
	Id   int64  `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	Name string `json:"name" xorm:"VARCHAR(255)"`
}

func (u *User) GetName(id int64) (name string, err error) {
	u1 := new(User)
	_, err = Configs.Eg.Where("id = ?", id).Get(u1)
	if err != nil {
		return "", err
	}
	return u1.Name, err
}
