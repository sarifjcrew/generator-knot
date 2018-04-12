package models

import(
	"github.com/eaciit/orm/v1"
)

type DashboardModel struct {
	orm.ModelBase `bson:"-",json:"-"`
	ID string `json:"ID",bson:"ID"`
}

func (u *DashboardModel) TableName() string {
	return "dashboard_model"
}

func (u *DashboardModel) RecordID() interface{} {
	return "u.ID"
}

