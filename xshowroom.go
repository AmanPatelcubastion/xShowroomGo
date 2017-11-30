//Sample query on graphiql
//
//	{user(id:"101") {
//		id
//		name
//	}}
//
//	{user(id:"") {
//		id
//		name
//	}}
package xShowroom

import "github.com/neelance/graphql-go"

var Schema = `
	schema {
		query: Query
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		user(id: ID!) : [User]!
	}

	type User {
		id: ID!
		name: String!
		device: Device!
	}

	type Device {
		id: ID!
		device_uuid: String!
	}

`

type x_user struct {
	id     graphql.ID
	name   string
	device x_device
}

type x_device struct {
	id          graphql.ID
	device_uuid string
}

var users = []*x_user{
	{
		id:   "101",
		name: "Aatish",
		device: x_device{
			id:          "201",
			device_uuid: "SJBCVU273F83CGU3",
		},
	},
	{
		id:   "102",
		name: "Vibhanshu",
		device: x_device{
			id:          "202",
			device_uuid: "23FGHKJBVDJVNKDJNV",
		},
	},
	{
		id:   "103",
		name: "Sandeep",
	},
}

var userData = make(map[graphql.ID]*x_user)

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

//==================		User		===========================

func (r *userResolver) ID() graphql.ID {
	return r.user.id
}

func (r *userResolver) Name() string {
	return r.user.name
}

func (r *userResolver) Device() *deviceResolver {
	return &deviceResolver{&r.user.device}
}

//==================		Device		===========================
func (r *deviceResolver) ID() graphql.ID {
	return r.device.id
}

func (r *deviceResolver) DeviceUuid() string {
	return r.device.device_uuid
}
