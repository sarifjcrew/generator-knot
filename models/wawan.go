package models

import(
	"github.com/eaciit/orm/v1"
)

type WawanModel struct {
	orm.ModelBase `bson:"-",json:"-"`
	ID string `json:"ID",bson:"ID"`
}

func (u *WawanModel) TableName() string {
	return "wawan_model"
}

func (u *WawanModel) RecordID() interface{} {
	return "u.ID"
}

