package model

import (
	"database"
)

type Lead struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string        `gorm:"column:name" json:"name,omitempty"`
	AccountId int         `gorm:"column:account_id" json:"account_id,omitempty"`
	UserId int         `gorm:"column:user_id" json:"user_id,omitempty"`
}

func GetLead(id int) Lead {
	data := Lead{}
	database.SQL.First(&data, id)
	return data
}

func GetLeads() []Lead {
	data := []Lead{}
	database.SQL.Find(&data)
	return data
}

func CreateLead(name string, accountId int) Lead {

	data := Lead{Name: name}

	if accountId != -1 {
		data = Lead{Name: name, AccountId: accountId}
	}

	database.SQL.Create(&data)
	return data
}

func UpdateLead(id int, name string, accountId int) Lead {
	oldData := Lead{Id: id}
	newData := Lead{Id: id, Name: name}

	if accountId != -1 {
		newData = Lead{Id: id, Name: name, AccountId: accountId}
	}

	database.SQL.Model(&oldData).Updates(newData)
	return newData
}

func GetLeadsOfAccount(accountId int) []Lead {
	var data []Lead
	database.SQL.Debug().Find(&data, "account_id = (?)", accountId)
	return data
}
