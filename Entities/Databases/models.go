package Databases

type Roles struct {
	Id   string `xorm:"not null pk autoincr UNSIGNED INT"`
	Role int    `xorm:"not null default 0 INT"`
}

type User struct {
	Id   string `xorm:"not null pk autoincr UNSIGNED INT"`
	Name string `xorm:"not null default '' VARCHAR(255)"`
}
