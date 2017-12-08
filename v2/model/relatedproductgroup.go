package model

import (
	"github.com/AmanPatelcubastion/xShowroomGo/v2/database"
	"github.com/neelance/graphql-go"
	"strconv"
)

type Relatedproductgroup struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	GroupType string        `gorm:"column:group_type" json:"name,omitempty"`
//	Product  []Product                           `gorm:"many2many:products_relatedproductgroups;"`
	ProductId []int         `gorm:"-" json:"account_id,omitempty"`
}

func GetRelatedProductGroup(id int) Relatedproductgroup {
	data := Relatedproductgroup{}
	database.SQL.First(&data, id)
	return data
}

func GetRelatedProductGroups() []Relatedproductgroup {
	data := []Relatedproductgroup{}
	database.SQL.Find(&data)
	return data
}

func CreateRelatedProductGroup(grouptype string, productId *[]graphql.ID) Relatedproductgroup {

	data := Relatedproductgroup{GroupType: grouptype}
	database.SQL.Create(&data)

	database.SQL.Model(&Relatedproductgroup{}).Last(&data)
	var ids []int

	newID:=*productId
	if newID != nil {
		for _, v := range newID {
			val, _ := strconv.Atoi(string(v))
			ids = append(ids, val)
		}
	}

	for _,v:=range ids{
		database.SQL.Create(Product_Group{RelatedProductGroupId:data.Id, ProductId: v})
	}


	return data
}

func UpdateRelatedProductGroup(id int, grouptype string, productId *[]graphql.ID) Relatedproductgroup {
	oldData := Relatedproductgroup{Id: id}
	newData := Relatedproductgroup{Id: id, GroupType: grouptype}

	database.SQL.Model(&oldData).Updates(newData)

	var ids []int

	newID:=*productId
	if newID != nil {
		for _, v := range newID {
			val, _ := strconv.Atoi(string(v))
			ids = append(ids, val)
		}
	}

	for _,v:=range ids{
		database.SQL.Create(Product_Group{RelatedProductGroupId:id, ProductId: v})
	}

	return newData
}

func GetRelatedProductGroupsOfProduct(productId int) []Relatedproductgroup {
	var data []Relatedproductgroup
	var temp []Product_Group
	database.SQL.Debug().Select("related_product_group_id").Where("product_id=?",productId).Find(&temp)

	var ids []int
	for _,v:=range temp{
		ids=append(ids,v.RelatedProductGroupId)
	}

	database.SQL.Debug().Where("id IN (?)",ids).Find(&data)

	return data
}
