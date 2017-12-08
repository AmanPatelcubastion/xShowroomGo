package mygraphql

import (
	"github.com/neelance/graphql-go"
	"github.com/AmanPatelcubastion/xShowroomGo/v2/model"
	"strconv"
)

//struct for graphql
type product struct {
	id    graphql.ID
	name  string
	relatedproductgroups []*relatedproductgroup
}

//struct for upserting
type productInput struct {
	Id    *graphql.ID
	Name  string
	Relatedproductgroupids *[]graphql.ID
	//Relatedproductgroups *[]relatedproductgroupInput
}

//struct for response
type productResolver struct {
	product *product
}

func ResolveProduct(args struct{ ID graphql.ID }) (response []*productResolver) {

	if args.ID != "" {
		response = append(response, &productResolver{MapProduct(
			model.GetProduct(convertId(args.ID)),
		)})
		return response
	}
	for _, val := range model.GetProducts() {
		response = append(response, &productResolver{MapProduct(
			val,
		)})
	}

	return response
}

func ResolveCreateProduct(args *struct {
	Product *productInput
}) *productResolver {

	var product = &product{}

	if args.Product.Id == nil {
		//create product
		product = MapProduct(model.CreateProduct(args.Product.Name,args.Product.Relatedproductgroupids)) //since Name is required field in schema no need for null check
	} else {
		//update product
		product = MapProduct(model.UpdateProduct(convertId(*args.Product.Id), args.Product.Name,args.Product.Relatedproductgroupids))
	}

/*	if product != nil && args.Product.Relatedproductgroups != nil {

		for _, relprogroup := range *args.Product.Relatedproductgroups {
			if relprogroup.Id == nil {
				model.CreateRelatedProductGroup(relprogroup.GroupType, convertId(product.id))
			} else {
				model.UpdateRelatedProductGroup(convertId(*relprogroup.Id), relprogroup.GroupType, convertId(product.id))
			}
		}

	}*/

	return &productResolver{product}
}

//==================		fields resolvers		===========================

func (r *productResolver) ID() graphql.ID {
	return r.product.id
}

func (r *productResolver) Name() string {
	return r.product.name
}

//This method will run, if device is asked for
func (r *productResolver) Relatedproductgroups() []*relatedproductgroupResolver {

	var l []*relatedproductgroupResolver
	if r.product != nil {
		//if product not null get device of product from db and map
		relprogrp := model.GetRelatedProductGroupsOfProduct(convertId(r.product.id))
		for _, v := range relprogrp {
			l = append(l, &relatedproductgroupResolver{MapRelatedProductGroup(v)})
		}
		return l
	}

	for _, v := range r.product.relatedproductgroups {
		l = append(l, &relatedproductgroupResolver{v})
	}
	return l
}

//=================			mapper methods			==============================
func MapProduct(modelProduct model.Product) *product {

	//if modelProduct == (model.Product{}) {
	//	return &product{}
	//}

	//create graphql (product) from model (product)
	product := product{
		id:   graphql.ID(strconv.Itoa(modelProduct.Id)),
		name: modelProduct.Name,
	}
	return &product
}

//Graphql to Db
func MapProduct1(modelProduct *[]productInput) []model.Product {

	var RelproGrp []model.Product
	if modelProduct!=nil{
		for _,v:=range *modelProduct{
			id1,_:=strconv.Atoi(string(*v.Id))
			RelproGrp = append(RelproGrp,model.Product{
				Id:   id1,
				Name: v.Name,
			})
		}
	}

	return RelproGrp
}