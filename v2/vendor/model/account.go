package model

import (
	"database"
)

type Account struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string        `gorm:"column:name" json:"name,omitempty"`
}

func GetAccount(id int) Account {
	data := Account{}
	database.SQL.First(&data, id)
	return data
}

func GetAccounts() []Account {
	data := []Account{}
	database.SQL.Find(&data)
	return data
}

func CreateAccount(name string) Account {
	data := Account{Name: name}
	database.SQL.Create(&data)
	return data
}

func UpdateAccount(id int, name string) Account {
	oldData := Account{Id: id}
	newData := Account{Id: id, Name: name}
	database.SQL.Model(&oldData).Updates(newData)
	return newData
}

func GetAccountOfLead(leadlId int,leadType string) Account {
	data :=Account{}

	database.SQL.Model(&Account{}).Select("accounts.id,accounts.name").Joins("inner join leads on" +
		" leads.type_id = accounts.id").Where("leads.id=? && leads.lead_type=?",leadlId,leadType).Scan(&data)
	return data
}
