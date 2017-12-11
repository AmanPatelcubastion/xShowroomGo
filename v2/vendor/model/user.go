package model

import (
	"fmt"
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

func DeleteUser(id int) User {
	data:=User{}
	//dataDevice:=Device{UserId:id}
	database.SQL.Where("id=(?)",id).First(&data)
	database.SQL.Model(Device{}).Where("user_id = ?", id).Update("user_id", 0)
	database.SQL.Model(Lead{}).Where("type_id = ?", id).Where("lead_type=(?)","user").Update("type_id", 0)

	fmt.Print("data :",data)
	if data!=(User{}){
		database.SQL.Debug().Delete(&data)
	}

	return data
}

func GetUserOfDevice(deviceId int) User {
	data := User{}
    database.SQL.Model(&User{}).Select("users.id,users.name").Joins("inner join devices on devices.user_id = users.id").Where("devices.id=(?)",deviceId).Scan(&data)


    //data := User{Id: 1, Name: "join likhna h"}
	//database.SQL.First(&data, "user_id", userId)
	return data
}

func GetUserOfLead(leadlId int) User {
	data :=User{}

	database.SQL.Model(&User{}).Joins("inner join leads on leads.type_id = users.id").Where("leads.id=?",leadlId).Scan(&data)

	return data
}