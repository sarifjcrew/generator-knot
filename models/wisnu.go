package models

import(
	"github.com/eaciit/orm/v1"
)

type WisnuModel struct {
	orm.ModelBase `bson:"-",json:"-"`
	ID string `json:"ID",bson:"ID"`
}

func (u *WisnuModel) TableName() string {
	return "wisnu_model"
}

func (u *WisnuModel) RecordID() interface{} {
	return "u.ID"
}

