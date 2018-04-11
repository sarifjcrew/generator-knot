package models

import(
	"github.com/eaciit/orm/v1"
)

type TestModel struct {
	orm.ModelBase `bson:"-",json:"-"`
	ID string `json:"ID",bson:"ID"`
}

func (u *TestModel) TableName() string {
	return "test_model"
}

func (u *TestModel) RecordID() interface{} {
	return "u.ID"
}

