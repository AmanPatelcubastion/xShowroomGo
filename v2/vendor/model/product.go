package model

import (
	"database"
	"github.com/neelance/graphql-go"
	"strconv"
)

type Product struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string        `gorm:"column:name" json:"name,omitempty"`
	RelatedproductgroupId []int    `gorm:"-" json:"account_id,omitempty"`
}

func GetProduct(id int) Product {
	data := Product{}
	database.SQL.Debug().First(&data,"id=(?)",id)
	return data
}

func GetProducts() []Product {
	data := []Product{}
	database.SQL.Find(&data)
	return data
}

func CreateProduct(name string, relatedproductgroupId *[]graphql.ID) Product {

	data := Product{Name: name}
	database.SQL.Create(&data)

    database.SQL.Model(&Product{}).Last(&data)
	if *relatedproductgroupId != nil {
		for _, v := range *relatedproductgroupId {
			val, _ := strconv.Atoi(string(v))
			database.SQL.Create(Product_Group{ProductId:data.Id, RelatedProductGroupId: val})
		}
	}

	return data
}

func UpdateProduct(id int, name string, relatedproductgroupId *[]graphql.ID) Product {
	oldData := Product{Id: id}
	newData := Product{Id: id, Name: name}
	database.SQL.Model(&oldData).Updates(newData)

	if *relatedproductgroupId != nil {
		for _, v := range *relatedproductgroupId {
			val, _ := strconv.Atoi(string(v))
			database.SQL.Create(Product_Group{ProductId:id, RelatedProductGroupId: val})
			}

	}

	return newData
}

//
func GetProductOfRelatedProductGroups(relatedproductgroupId int) []Product {
	var data []Product
	var temp []Product_Group
	database.SQL.Debug().Select("product_id").Where("related_product_group_id=?",relatedproductgroupId).Find(&temp)

	var ids []int
	for _,v:=range temp{
		ids=append(ids,v.ProductId)
	}

	database.SQL.Debug().Where("id IN (?)",ids).Find(&data)
	return data
}
