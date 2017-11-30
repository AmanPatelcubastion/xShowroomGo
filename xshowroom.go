//Sample query on graphiql
//
//	1)	Query users with id
//	{user(id:"101") {
//		id
//		name
//	}}
//
//	2)	Query all users
//	{user(id:"") {
//		id
//		name
//	}}
//
//	3)  Upsert user without device
//	mutation{
//		createUser(user:{
//		id:"104"
//		name:"Sample"
//		}){
//			id
//			}
//		}
//
//  4) Upsert user with device
//	mutation{
//		createUser(user:{
//			id:"104"
//			name:"Sample"
//			device:{
//				id:"204"
//				device_uuid:"DUVBSJDKNVJU3874VBHJHDFK"
//			}
//		}){
//			id
//			name
//		   }
//		}


package xShowroom

import (
	"github.com/neelance/graphql-go"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		user(id: ID!) : [User]!
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createDevice(device: DeviceInput!): Device
		createUser(user: UserInput!): User
	}

	type User {
		id: ID!
		name: String!
		device: Device!
	}
	input UserInput {
		id: ID!
		name: String!
		device: DeviceInput!
	}

	type Device {
		id: ID!
		device_uuid: String!
	}
	input DeviceInput {
		id: ID!
		device_uuid: String!
	}

`

type x_user struct {
	id     graphql.ID
	name   string
	device *x_device
}

type userInput struct {
	Id     graphql.ID
	Name   string
	Device deviceInput
}

type x_device struct {
	id          graphql.ID
	device_uuid string
}

type deviceInput struct {
	Id          graphql.ID
	Device_uuid string
}

var devices = []*x_device{
	{
		id:          "201",
		device_uuid: "SJBCVU273F83CGU3",
	},
	{
		id:          "202",
		device_uuid: "FBEUVIWU3784HFBV",
	},
}

var users = []*x_user{
	{
		id:     "101",
		name:   "Aatish",
		device: devices[0],
	},
	{
		id:     "102",
		name:   "Vibhanshu",
		device: devices[1],
	},
	{
		id:   "103",
		name: "Sandeep",
	},
}

var userData = make(map[graphql.ID]*x_user)

var deviceData = make(map[graphql.ID]*x_device)

func init() {
	for _, user := range users {
		userData[user.id] = user
	}
}

type Resolver struct{}

type userResolver struct {
	user *x_user
}

type deviceResolver struct {
	device *x_device
}

//======================		query		===============================

func (r *Resolver) User(args struct{ ID graphql.ID }) []*userResolver {
	var l []*userResolver
	if args.ID != "" {
		l = append(l, &userResolver{userData[args.ID]})
		return l
	}
	for _, val := range userData {
		l = append(l, &userResolver{val})
	}
	return l
}

//======================		mutation		===============================

func (r *Resolver) CreateDevice(args *struct {
	Device *deviceInput
}) *deviceResolver {

	device := &x_device{
		id:          args.Device.Id,
		device_uuid: args.Device.Device_uuid,
	}

	deviceData[device.id] = device
	return &deviceResolver{deviceData[device.id]}
}

func (r *Resolver) CreateUser(args *struct {
	User *userInput
}) *userResolver {

	user := &x_user{
		id:   args.User.Id,
		name: args.User.Name,
		device: &x_device{
			id:          args.User.Device.Id,
			device_uuid: args.User.Device.Device_uuid,
		},
	}

	userData[user.id] = user
	return &userResolver{userData[user.id]}
}

//==================		User		===========================

func (r *userResolver) ID() graphql.ID {
	return r.user.id
}

func (r *userResolver) Name() string {
	return r.user.name
}

func (r *userResolver) Device() *deviceResolver {
	if r.user.device == nil {
		return &deviceResolver{&x_device{}}
	}
	return &deviceResolver{r.user.device}
}

//==================		Device		===========================
func (r *deviceResolver) ID() graphql.ID {
	return r.device.id
}

func (r *deviceResolver) DeviceUuid() string {
	return r.device.device_uuid
}
