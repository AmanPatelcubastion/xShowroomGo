package model

import (
	"database"
	"github.com/neelance/graphql-go"
	"strconv"
)

type Product struct {
	Id   int        `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string        `gorm:"column:name" json:"name,omitempty"`
//	Relatedproductgroup  []Relatedproductgroup                  `gorm:"many2many:products_relatedproductgroups;"`
	RelatedproductgroupId []int    `gorm:"-" json:"account_id,omitempty"`
}

func GetProduct(id int) Product {
	data := Product{}
	database.SQL.Debug().First(&data)
//	database.SQL.Debug().Select("products.id,products.name,relatedproductgroups.id,relatedproductgroups.name").Find(&data)
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
   // var ids []int
   // newID:=
	if *relatedproductgroupId != nil {
		for _, v := range *relatedproductgroupId {
			val, _ := strconv.Atoi(string(v))
			database.SQL.Create(Product_Group{ProductId:data.Id, RelatedProductGroupId: val})
			//ids = append(ids, val)
		}

		//for _,v:=range ids{
		//}
	}

	return data
}

func UpdateProduct(id int, name string, relatedproductgroupId *[]graphql.ID) Product {
	oldData := Product{Id: id}
	newData := Product{Id: id, Name: name}
	database.SQL.Model(&oldData).Updates(newData)

	//var ids []int

	//newID:=*relatedproductgroupId
	if *relatedproductgroupId != nil {
		for _, v := range *relatedproductgroupId {
			val, _ := strconv.Atoi(string(v))
			database.SQL.Create(Product_Group{ProductId:id, RelatedProductGroupId: val})

			//ids = append(ids, val)
		}

		//for _,v:=range ids{
		//}
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
