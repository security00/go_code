package models

import (
	"fmt"
	"mygo/Configs"
)

type GoodsForOmoCourses struct {
	Id               int64  `json:"id" xorm:"pk autoincr comment('id') BIGINT(20)"`
	ZxOnlineCourseId int64  `json:"zx_online_course_id" xorm:"not null default 0 comment('【在线】在线课程id') unique(idx_snj2_key) BIGINT(20)"`
	ZxOmoCourseId    int64  `json:"zx_omo_course_id" xorm:"not null default 0 comment('【在线】omo课程id') index BIGINT(20)"`
	Snj2OmoId        int64  `json:"snj2_omo_id" xorm:"not null default 0 comment('【神2】套餐id') BIGINT(20)"`
	Snj2OmoName      string `json:"snj2_omo_name" xorm:"comment('【神2】套餐名称') VARCHAR(100)"`
	Snj2ProductId    int64  `json:"snj2_product_id" xorm:"not null default 0 comment('【神2】教育线下产品ID') unique(idx_snj2_key) BIGINT(20)"`
	Snj2ProductName  string `json:"snj2_product_name" xorm:"comment('【神2】教育线下产品名称') VARCHAR(100)"`
	Snj2CourseId     int64  `json:"snj2_course_id" xorm:"not null default 0 comment('【神2】教育线下课程ID') unique(idx_snj2_key) BIGINT(20)"`
	Snj2CourseName   string `json:"snj2_course_name" xorm:"comment('【神2】教育线下课程名称') VARCHAR(100)"`
	Snj2GoodId       int64  `json:"snj2_good_id" xorm:"not null default 0 comment('【神2】教育线下商品ID') unique(idx_snj2_key) BIGINT(20)"`
	Snj2GoodName     string `json:"snj2_good_name" xorm:"comment('【神2】教育线下商品名称') VARCHAR(100)"`
	Snj2ProvinceId   int    `json:"snj2_province_id" xorm:"not null default 0 comment('神2省份id') TINYINT(1)"`
	Snj2SubSchoolId  int    `json:"snj2_sub_school_id" xorm:"not null default 0 comment('【神2】分校id') index INT(11)"`
	Snj2IsSubPackage int    `json:"snj2_is_sub_package" xorm:"not null default 1 comment('【神2】是否子套餐课程，0:否，1:是') TINYINT(1)"`
	DeletedAt        string `json:"deleted_at" xorm:"comment('删除时间') DATETIME"`
	UpdatedAt        string `json:"updated_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
	CreatedAt        string `json:"created_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
}

func (g *GoodsForOmoCourses) Insert() (r int64, err error) {
	fmt.Println("<================ to do insert relation ================>")
	r, err = Configs.Eg.Insert(*g)
	if err != nil {
		return 0, err
	}
	fmt.Sprintf("insert return:%d", r)
	return r, err
}
