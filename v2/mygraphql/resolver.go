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

//===========		helper methods		===================
func convertId(id graphql.ID) int {
	val, err := strconv.Atoi(string(id))
	if err != nil {
		fmt.Println("Id conversion failed")
		return 0
	}
	return val
}
