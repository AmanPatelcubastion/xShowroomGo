package model

import (
	"database"
)

type Lead struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string        `gorm:"column:name" json:"name,omitempty"`
	LeadType  string     `gorm:"column:lead_type" json:"lead_type,omitempty"`
	TypeId int         `gorm:"column:type_id" json:"account_id,omitempty"`
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

func CreateLead(name string, leadType string ,typeId int) Lead {

	data := Lead{Name: name}

	if typeId != -1{
		data = Lead{Name: name, LeadType:leadType ,TypeId: typeId}
	}

	database.SQL.Create(&data)
	return data
}

func UpdateLead(id int, name string,leadType string ,typeId int) Lead {
	oldData := Lead{Id: id}
	newData := Lead{Id: id, Name: name}

	if typeId != -1 {
		newData = Lead{Id: id, Name: name, LeadType:leadType , TypeId: typeId}
	}

	database.SQL.Model(&oldData).Updates(newData)
	return newData
}


/*

func GetLeadsOfAccount(accountId int) []Lead {
	var data []Lead
	database.SQL.Debug().Find(&data, "account_id = (?)", accountId)
	return data
}

func GetLeadsOfUser(userId int) []Lead {
	var data []Lead
	database.SQL.Debug().Find(&data, "user_id = (?)", userId)
	return data
}*/


func GetLeadsofType (typeId int, leadType string) []Lead{

	var data []Lead
	database.SQL.Debug().Where("lead_type=?",leadType).Find(&data, "type_id = (?)", typeId)

	return data
}