package mygraphql

import (
	"github.com/neelance/graphql-go"
	"fmt"
	"strconv"
)

type Resolver struct{}

//============		root	query 	methods			=============================
func (r *Resolver) User(args struct{ ID graphql.ID }) []*userResolver {
	return ResolveUser(args)
}

func (r *Resolver) Device(args struct{ ID graphql.ID }) []*deviceResolver {
	return ResolveDevice(args)
}

func (r *Resolver) Account(args struct{ ID graphql.ID }) []*accountResolver {
	return ResolveAccount(args)
}

func (r *Resolver) Lead(args struct{ ID graphql.ID }) []*leadResolver {
	return ResolveLead(args)
}

func (r *Resolver) Product(args struct{ ID graphql.ID }) []*productResolver {
	return ResolveProduct(args)
}

func (r *Resolver) Relatedproductgroup(args struct{ ID graphql.ID }) []*relatedproductgroupResolver {
	return ResolveRelatedproductgroup(args)
}

//=============		root	mutation	methods		===============================

func (r *Resolver) CreateUser(args *struct {
	User *userInput
}) *userResolver {
	return ResolveCreateUser(args)
}

func (r *Resolver) CreateDevice(args *struct {
	Device *deviceInput
}) *deviceResolver {
	return ResolveCreateDevice(args)
}

func (r *Resolver) CreateAccount(args *struct {
	Account *accountInput
}) *accountResolver {
	return ResolveCreateAccount(args)
}

func (r *Resolver) CreateLead(args *struct {
	Lead *leadInput
}) *leadResolver {
	return ResolveCreateLead(args)
}

func (r *Resolver) CreateProduct(args *struct {
	Product *productInput
}) *productResolver {
	return ResolveCreateProduct(args)
}

func (r *Resolver) CreateRelatedproductgroup(args *struct {
	RelatedProductGroup *relatedproductgroupInput
}) *relatedproductgroupResolver {
	return ResolveCreateRelatedproductgroup(args)
}

//===========		helper methods		===================
func convertId(id graphql.ID) int {
	val, err := strconv.Atoi(string(id))
	if err != nil {
		fmt.Println("Id conversion failed")
		return 0
	}
	return val
}
