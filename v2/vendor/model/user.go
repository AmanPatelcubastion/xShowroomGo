package model

import (
	"database"
)

type User struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string        `gorm:"column:name" json:"name,omitempty"`
}

func GetUser(id int) User {
	data := User{}
	database.SQL.First(&data, id)
	return data
}

func GetUsers() []User {
	data := []User{}
	database.SQL.Find(&data)
	return data
}

func CreateUser(name string) User {
	data := User{Name: name}
	database.SQL.Create(&data)
	return data
}

func UpdateUser(id int, name string) User {
	oldData := User{Id: id}
	newData := User{Id: id, Name: name}
	database.SQL.Model(&oldData).Updates(newData)
	return newData
}

func GetUserOfDevice(deviceId int) User {
	data := User{}
    database.SQL.Model(&User{}).Select("users.id,users.name").Joins("left join devices on devices.user_id = users.id").Scan(&data)
	//data := User{Id: 1, Name: "join likhna h"}
	//database.SQL.First(&data, "user_id", userId)
	return data
}

func GetUserOfLead(leadlId int) User {
	data :=User{}

	database.SQL.Model(&User{}).Joins("inner join leads on leads.type_id = users.id").Where("leads.id=?",leadlId).Scan(&data)

	return data
}