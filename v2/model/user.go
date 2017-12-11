package model

import (
	"github.com/KiranKanadeCubastion/xShowroomGo/v2/database"
)

type User struct {
	Id   int           `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
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
	database.SQL.Debug().Model(&User{}).Joins("inner join devices on devices.user_id = users.id").
		Where("devices.id = ?", deviceId).Scan(&data)
	return data
}