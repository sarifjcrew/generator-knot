package models

import(
	"github.com/eaciit/orm/v1"
)

type DashboardtestModel struct {
	orm.ModelBase `bson:"-",json:"-"`
	ID string `json:"ID",bson:"ID"`
}

func (u *DashboardtestModel) TableName() string {
	return "dashboardtest_model"
}

func (u *DashboardtestModel) RecordID() interface{} {
	return "u.ID"
}

