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
		device(id: ID!) : [Device]!
        account(id: ID!) : [Account]!
		lead(id: ID!) : [Lead]!
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createDevice(device: DeviceInput!): Device
		createUser(user: UserInput!): User
        createLead(lead: LeadInput!): Lead
        createAccount(account: AccountInput!): Account
	}

	# The individual using xShowroom application
	type User {
		id: ID!
		name: String!
		device: Device!
	}
	input UserInput {
		id: ID!
		name: String!
		device: DeviceInput
	}

	# The Android or I-pad device used by a user
	type Device {
		id: ID!
		device_uuid: String!
		user: User!
	}
	input DeviceInput {
		id: ID!
		device_uuid: String!
	}

    type Account {
      id: ID!
      name: String!
      acctype: String!
        leads: [Lead!]!
   }

    input AccountInput {
      id: ID!
      name: String!
      acctype: String!
      leads: [LeadInput!]
   }

    type Lead {
      id: ID!
      name: String!
      location: String!
   }

    input LeadInput {
      id: ID!
      name: String!
      location: String!
   }

`
//=======================		Types 		=====================================
type x_user struct {
	id     graphql.ID
	name   string
	device *x_device
}

type userInput struct {
	Id     graphql.ID
	Name   string
	Device *deviceInput
}

type x_device struct {
	id          graphql.ID
	device_uuid string
	user        *x_user
}

type deviceInput struct {
	Id          graphql.ID
	Device_uuid string
}

type x_account struct {
	id      graphql.ID
	name    string
	acctype string
	leads   []*x_lead
}

type accountInput struct {
	Id      graphql.ID
	Name    string
	Acctype string
	Leads   *[]leadInput
}

type x_lead struct {
	id       graphql.ID
	name     string
	location string
}

type leadInput struct {
	Id       graphql.ID
	Name     string
	Location string
}

//========================		Sample Data		======================================
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
		id:   "101",
		name: "Aatish",
	},
	{
		id:   "102",
		name: "Vibhanshu",
	},
	{
		id:   "103",
		name: "Sandeep",
	},
}

var leads = []*x_lead{
	{
		id:       "201",
		name:     "Ford",
		location: "Delhi",
	},
	{
		id:       "202",
		name:     "Maruti",
		location: "Noida",
	},
	{
		id:       "203",
		name:     "Hyundai",
		location: "Gurgaon",
	},
	{
		id:       "204",
		name:     "Honda",
		location: "Mumbai",
	},
}

var accounts = []*x_account{
	{
		id:      "11",
		name:    "Aatish",
		acctype: "admin",
		leads:   []*x_lead{leads[0], leads[2]},
	},
	{
		id:      "102",
		name:    "Vibhanshu",
		acctype: "admin",
		leads:   []*x_lead{leads[1], leads[3]},
	},
	{
		id:   "103",
		name: "Sandeep",
	},
}

var userData = make(map[graphql.ID]*x_user)

var deviceData = make(map[graphql.ID]*x_device)

var accountData = make(map[graphql.ID]*x_account)

var leadData = make(map[graphql.ID]*x_lead)

func init() {

	// create sample data
	for i, user := range users {
		userData[user.id] = user

		if len(devices) > i {
			userData[user.id].device = devices[i] //add devices to users (temp, db joins will generate this)
		}
	}

	for i, device := range devices {
		deviceData[device.id] = device

		if len(users) > i {
			deviceData[device.id].user = users[i] //add users to device (temp, db joins will generate this)
		}
	}

	for i, account := range accounts {
		accountData[account.id] = account

		if len(leads) > i {
			accountData[account.id].leads = []*x_lead{leads[i], leads[i+1]} //add devices to users (temp, db joins will generate this)
		}
	}

}

//============================		Resolver types		===================================
type Resolver struct{}

type userResolver struct {
	user *x_user
}

type deviceResolver struct {
	device *x_device
}

type accountResolver struct {
	account *x_account
}

type leadResolver struct {
	lead *x_lead
}

//======================		Query	methods		===============================

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

func (r *Resolver) Device(args struct{ ID graphql.ID }) []*deviceResolver {
	var d []*deviceResolver
	if args.ID != "" {
		d = append(d, &deviceResolver{deviceData[args.ID]})
		return d
	}
	for _, val := range deviceData {
		d = append(d, &deviceResolver{val})
	}
	return d
}

func (r *Resolver) Account(args struct{ ID graphql.ID }) []*accountResolver {
	var l []*accountResolver
	if args.ID != "" {
		l = append(l, &accountResolver{accountData[args.ID]})
		return l
	}
	for _, val := range accountData {
		l = append(l, &accountResolver{val})
	}
	return l
}

func (r *Resolver) Lead(args struct{ ID graphql.ID }) []*leadResolver {
	var d []*leadResolver
	if args.ID != "" {
		d = append(d, &leadResolver{leadData[args.ID]})
		return d
	}
	for _, val := range leadData {
		d = append(d, &leadResolver{val})
	}
	return d
}

//======================		Mutation	methods		===============================

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

	var deviceId graphql.ID
	var deviceUuid string

	if args.User.Device != nil {
		deviceId = args.User.Device.Id
		deviceUuid = args.User.Device.Device_uuid
	}

	user := &x_user{
		id:   args.User.Id,
		name: args.User.Name,
		device: &x_device{
			id:          deviceId,
			device_uuid: deviceUuid,
		},
	}

	userData[user.id] = user
	return &userResolver{userData[user.id]}
}

func (r *Resolver) CreateLead(args *struct {
	Lead *leadInput
}) *leadResolver {

	lead := &x_lead{
		id:       args.Lead.Id,
		name:     args.Lead.Name,
		location: args.Lead.Location,
	}

	leadData[lead.id] = lead
	return &leadResolver{leadData[lead.id]}
}

func (r *Resolver) CreateAccount(args *struct {
	Account *accountInput
}) *accountResolver {

	var createdlead []*x_lead
	if args.Account.Leads != nil {
		for _, v := range *args.Account.Leads {
			createdlead = append(createdlead, &x_lead{
				id:       v.Id,
				name:     v.Name,
				location: v.Location,
			}, )
		}
	}

	account := &x_account{
		id:      args.Account.Id,
		name:    args.Account.Name,
		acctype: args.Account.Acctype,
		leads:   createdlead,
	}

	accountData[account.id] = account
	return &accountResolver{accountData[account.id]}
}

//==================		User fields resolvers		===========================

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

//==================		Device	fields resolvers	===========================
func (r *deviceResolver) ID() graphql.ID {
	return r.device.id
}

func (r *deviceResolver) DeviceUuid() string {
	return r.device.device_uuid
}

func (r *deviceResolver) User() *userResolver {
	if r.device.user == nil {
		return &userResolver{&x_user{}}
	}
	return &userResolver{r.device.user}
}

//==================      Account       ===========================

func (r *accountResolver) ID() graphql.ID {
	return r.account.id
}

func (r *accountResolver) Name() string {
	return r.account.name
}

func (r *accountResolver) Acctype() string {
	return r.account.acctype
}

func (r *accountResolver) Leads() []*leadResolver {
	//return r.user.device

	var l []*leadResolver
	for _, v := range r.account.leads {
		l = append(l, &leadResolver{v})
	}
	return l
}

//==================      Lead      ===========================
func (r *leadResolver) ID() graphql.ID {
	return r.lead.id
}

func (r *leadResolver) Name() string {
	return r.lead.name
}

func (r *leadResolver) Location() string {
	return r.lead.location
}
