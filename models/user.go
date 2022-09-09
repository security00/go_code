package models

type User struct {
	Id   int64  `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	Name string `json:"name" xorm:"VARCHAR(255)"`
}

func (u *User) GetName() (name string, err error) {
	//Configs.Eg
	return name, err
}
