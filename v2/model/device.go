package model

import (
	"github.com/AmanPatelcubastion/xShowroomGo/v2/database"
)

type Device struct {
	Id     int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	UUID   string        `gorm:"column:uuid" json:"uuid,omitempty"`
	UserId int         `gorm:"column:user_id" json:"user_id,omitempty"`
}

func GetDevice(id int) Device {
	data := Device{}
	database.SQL.Debug().First(&data, id)
	return data
}

func GetDevices() []Device {
	data := []Device{}
	database.SQL.Debug().Find(&data)
	return data
}

func CreateDevice(uuid string, userId int) Device {

	data := Device{UUID: uuid}

	if userId != -1 {
		data = Device{UUID: uuid, UserId: userId}
	}

	database.SQL.Create(&data)
	return data
}

func UpdateDevice(id int, uuid string, userId int) Device {
	oldData := Device{Id: id}
	newData := Device{Id: id, UUID: uuid}

	if userId != -1 {
		newData = Device{Id: id, UUID: uuid, UserId: userId}
	}

	database.SQL.Model(&oldData).Updates(newData)
	return newData
}


func GetDeviceOfUser(userId int) Device {
	data := Device{}
	database.SQL.Debug().First(&data, "user_id = (?)", userId)
	return data
}
