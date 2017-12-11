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
        product(id: ID!) : [Product]!
        relatedproductgroup(id: ID!) : [RelatedProductGroup]!
        unrelatedproductgroup(id: ID!) : [UnRelatedProductGroup]!
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createDevice(device: DeviceInput!): Device
		createUser(user: UserInput!): User
        createLead(lead: LeadInput!): Lead
        createAccount(account: AccountInput!): Account
        createProduct(product: ProductInput): Product
        createRelatedProductGroup(relatedproductgroup: RelatedProductGroupInput!): RelatedProductGroup
        createUnRelatedProductGroup(unrelatedproductgroup: UnRelatedProductGroupInput!): UnRelatedProductGroup
	}

	# The individual using xShowroom application
	type User {
		id: ID!
		name: String!
		device: Device!
        leads: [Lead!]!
        product: [Product!]!
	}
	input UserInput {
		id: ID!
		name: String!
		device: DeviceInput
        leads: [LeadInput!]
        product: [ProductInput!]
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
      users: User!
      accounts: Account!
   }

    input LeadInput {
      id: ID!
      name: String!
      location: String!
   }

    type Product {
      id: ID!
      name: String!
      group: String!
      user: User!
      relprogroups: [RelatedProductGroup!]!
      unrelprogroups: [UnRelatedProductGroup!]!
   }

    input ProductInput {
      id: ID!
      name: String!
      group: String!
      relprogroups: [RelatedProductGroupInput!]
      unrelprogroups: [UnRelatedProductGroupInput!]
   }

    type RelatedProductGroup {
      id: ID!
      grouptype: String!
      product: [Product!]!
   }

    input RelatedProductGroupInput {
      id: ID!
      grouptype: String!
      product: [ProductInput!]
   }

    type UnRelatedProductGroup {
      id: ID!
      grouptype: String!
      product: [Product!]!
   }

   input UnRelatedProductGroupInput {
      id: ID!
      grouptype: String!
      product: [ProductInput!]
   }

`
//=======================		Types 		=====================================
type x_user struct {
	id      graphql.ID
	name    string
	device  *x_device
	leads   []*x_lead
	product []*x_product
}

type userInput struct {
	Id      graphql.ID
	Name    string
	Device  *deviceInput
	Leads   *[]leadInput
	Product *[]productInput
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
	users    *x_user
	accounts *x_account
}

type leadInput struct {
	Id       graphql.ID
	Name     string
	Location string
}

type x_product struct {
	id             graphql.ID
	name           string
	group          string
	user           *x_user
	relprogroups   []*x_related_product_group
	unrelprogroups []*x_unrelated_product_group
}

type productInput struct {
	ID             graphql.ID
	Name           string
	Group          string
	Relprogroups   *[]relatedproductgroupInput
	UnRelprogroups *[]unrelatedproductgroupInput
}

type x_related_product_group struct {
	id        graphql.ID
	grouptype string
	product   []*x_product
}

type relatedproductgroupInput struct {
	ID        graphql.ID
	Grouptype string
	Product   *[]productInput
}

type x_unrelated_product_group struct {
	id        graphql.ID
	grouptype string
	product   []*x_product
}

type unrelatedproductgroupInput struct {
	ID        graphql.ID
	Grouptype string
	Product   *[]productInput
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

var products = []*x_product{
	{
		id:    "90",
		name:  "Air Conditioner",
		group: "Electronics",
	},
	{
		id:    "91",
		name:  "Washing Machine",
		group: "Electronics",
	},
	{
		id:    "92",
		name:  "Cricket Bat",
		group: "Sports",
	},
	{
		id:    "94",
		name:  "Volleyball",
		group: "sports",
	},
	{
		id:    "95",
		name:  "Dining Table",
		group: "Home",
	},
	{
		id:    "96",
		name:  "Handy Cam",
		group: "Battery",
	},
	{
		id:    "97",
		name:  "Video Camera",
		group: "Miscellaneous",
	},
	{
		id:    "98",
		name:  "Spikes",
		group: "Track",
	},
}

var relatedProductGroup = []*x_related_product_group{
	{
		id:        "190",
		grouptype: "Electronics",
	},
	{
		id:        "192",
		grouptype: "Sports",
	},
}

var unrelatedProductGroup = []*x_unrelated_product_group{
	{
		id:        "290",
		grouptype: "Miscellaneous",
	},
	{
		id:        "291",
		grouptype: "Heavy",
	},
}

var userData = make(map[graphql.ID]*x_user)

var deviceData = make(map[graphql.ID]*x_device)

var accountData = make(map[graphql.ID]*x_account)

var leadData = make(map[graphql.ID]*x_lead)

var productData = make(map[graphql.ID]*x_product)

var relatedproductgroupData = make(map[graphql.ID]*x_related_product_group)

var unrelatedproductgroupData = make(map[graphql.ID]*x_unrelated_product_group)

func init() {

	// create sample data
	for i, user := range users {
		userData[user.id] = user

		if len(devices) > i {
			userData[user.id].device = devices[i]
		}

		if len(leads) > i {
			userData[user.id].leads = []*x_lead{leads[i], leads[i+1]}
		}

		if len(products) > i {
			userData[user.id].product = []*x_product{products[i], products[i+1]}
		}

	}

	for i, device := range devices {
		deviceData[device.id] = device

		if len(users) > i {
			deviceData[device.id].user = users[i]
		}
	}

	for i, account := range accounts {
		accountData[account.id] = account

		if len(leads) > i {
			accountData[account.id].leads = []*x_lead{leads[i], leads[i+1]}
		}
	}

	for i, lead := range leads {
		leadData[lead.id] = lead

		if len(users) > i && i%2 == 0 {
			leadData[lead.id].users = users[i]
		}
		if len(accounts) > i && i%2 != 0 {
			leadData[lead.id].accounts = accounts[i]
		}

	}

	for i, product := range products {
		productData[product.id] = product

		if len(relatedProductGroup) > i {
			productData[product.id].relprogroups = []*x_related_product_group{relatedProductGroup[i]}
		}

		if len(users) > i {
			productData[product.id].user = users[i]
		}

		if len(unrelatedProductGroup) > i {
			productData[product.id].unrelprogroups = []*x_unrelated_product_group{unrelatedProductGroup[i]}
		}

	}

	for i, relatedproductgroup := range relatedProductGroup {
		relatedproductgroupData[relatedproductgroup.id] = relatedproductgroup

		if len(products) > i {
			relatedproductgroupData[relatedproductgroup.id].product = []*x_product{products[i]}
		}
	}

	for i, unrelatedproductgroup := range unrelatedProductGroup {
		unrelatedproductgroupData[unrelatedproductgroup.id] = unrelatedproductgroup

		if len(products) > i {
			unrelatedproductgroupData[unrelatedproductgroup.id].product = []*x_product{products[i]}
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

type productResolver struct {
	product *x_product
}

type relatedproductgroupResolver struct {
	relatedproductgroup *x_related_product_group
}

type unrelatedproductgroupResolver struct {
	unrelatedproductgroup *x_unrelated_product_group
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

func (r *Resolver) Product(args struct{ ID graphql.ID }) []*productResolver {
	var l []*productResolver
	if args.ID != "" {
		l = append(l, &productResolver{productData[args.ID]})
		return l
	}
	for _, val := range productData {
		l = append(l, &productResolver{val})
	}
	return l
}

func (r *Resolver) RelatedProductGroup(args struct{ ID graphql.ID }) []*relatedproductgroupResolver {
	var d []*relatedproductgroupResolver
	if args.ID != "" {
		d = append(d, &relatedproductgroupResolver{relatedproductgroupData[args.ID]})
		return d
	}
	for _, val := range relatedproductgroupData {
		d = append(d, &relatedproductgroupResolver{val})
	}
	return d
}

func (r *Resolver) UnRelatedProductGroup(args struct{ ID graphql.ID }) []*unrelatedproductgroupResolver {
	var d []*unrelatedproductgroupResolver
	if args.ID != "" {
		d = append(d, &unrelatedproductgroupResolver{unrelatedproductgroupData[args.ID]})
		return d
	}
	for _, val := range unrelatedproductgroupData {
		d = append(d, &unrelatedproductgroupResolver{val})
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

	var createdlead []*x_lead
	if args.User.Leads != nil {
		for _, v := range *args.User.Leads {
			createdlead = append(createdlead, &x_lead{
				id:       v.Id,
				name:     v.Name,
				location: v.Location,
			}, )
		}
	}

	var createdproduct []*x_product
	if args.User.Product != nil {
		for _, v := range *args.User.Product {
			createdproduct = append(createdproduct, &x_product{
				id:    v.ID,
				name:  v.Name,
				group: v.Group,
			}, )
		}
	}

	user := &x_user{
		id:   args.User.Id,
		name: args.User.Name,
		device: &x_device{
			id:          deviceId,
			device_uuid: deviceUuid,
		},
		leads:   createdlead,
		product: createdproduct,
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

func (r *Resolver) CreateProduct(args *struct {
	Product *productInput
}) *productResolver {

	var createdRelatedproductgroup []*x_related_product_group
	if args.Product.Relprogroups != nil {
		for _, v := range *args.Product.Relprogroups {
			createdRelatedproductgroup = append(createdRelatedproductgroup, &x_related_product_group{
				id:        v.ID,
				grouptype: v.Grouptype,
			}, )
		}
	}

	var createdUnRelatedproductgroup []*x_unrelated_product_group
	if args.Product.UnRelprogroups != nil {
		for _, v := range *args.Product.UnRelprogroups {
			createdUnRelatedproductgroup = append(createdUnRelatedproductgroup, &x_unrelated_product_group{
				id:        v.ID,
				grouptype: v.Grouptype,
			}, )
		}
	}

	product := &x_product{
		id:             args.Product.ID,
		name:           args.Product.Name,
		group:          args.Product.Group,
		relprogroups:   createdRelatedproductgroup,
		unrelprogroups: createdUnRelatedproductgroup,
	}

	productData[product.id] = product
	return &productResolver{productData[product.id]}
}

func (r *Resolver) CreateRelatedProductGroup(args *struct {
	RelatedProductGroup *relatedproductgroupInput
}) *relatedproductgroupResolver {

	var createdProduct []*x_product
	if args.RelatedProductGroup.Product != nil {
		for _, v := range *args.RelatedProductGroup.Product {
			createdProduct = append(createdProduct, &x_product{
				id:    v.ID,
				name:  v.Name,
				group: v.Group,
			}, )
		}
	}
	relatedproductgroup := &x_related_product_group{
		id:        args.RelatedProductGroup.ID,
		grouptype: args.RelatedProductGroup.Grouptype,
		product:   createdProduct,
	}

	relatedproductgroupData[relatedproductgroup.id] = relatedproductgroup
	return &relatedproductgroupResolver{relatedproductgroupData[relatedproductgroup.id]}
}

func (r *Resolver) CreateUnRelatedProductGroup(args *struct {
	UnRelatedProductGroup *unrelatedproductgroupInput
}) *unrelatedproductgroupResolver {

	var createdProduct []*x_product
	if args.UnRelatedProductGroup.Product != nil {
		for _, v := range *args.UnRelatedProductGroup.Product {
			createdProduct = append(createdProduct, &x_product{
				id:    v.ID,
				name:  v.Name,
				group: v.Group,
			}, )
		}
	}
	unrelatedproductgroup := &x_unrelated_product_group{
		id:        args.UnRelatedProductGroup.ID,
		grouptype: args.UnRelatedProductGroup.Grouptype,
		product:   createdProduct,
	}

	unrelatedproductgroupData[unrelatedproductgroup.id] = unrelatedproductgroup
	return &unrelatedproductgroupResolver{unrelatedproductgroupData[unrelatedproductgroup.id]}
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

func (r *userResolver) Leads() []*leadResolver {
	//return r.user.device

	var l []*leadResolver
	for _, v := range r.user.leads {
		l = append(l, &leadResolver{v})
	}
	return l
}

func (r *userResolver) Product() []*productResolver {
	//return r.user.device

	var l []*productResolver
	for _, v := range r.user.product {
		l = append(l, &productResolver{v})
	}
	return l
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

func (r *leadResolver) Users() *userResolver {
	if r.lead.users == nil {
		return &userResolver{&x_user{}}
	}
	return &userResolver{r.lead.users}
}

func (r *leadResolver) Accounts() *accountResolver {
	if r.lead.accounts == nil {
		return &accountResolver{&x_account{}}
	}
	return &accountResolver{r.lead.accounts}
}

//==================      Product       ===========================

func (r *productResolver) ID() graphql.ID {
	return r.product.id
}

func (r *productResolver) Name() string {
	return r.product.name
}

func (r *productResolver) Group() string {
	return r.product.group
}

func (r *productResolver) User() *userResolver {
	if r.product.user == nil {
		return &userResolver{&x_user{}}
	}
	return &userResolver{r.product.user}
}

func (r *productResolver) Relprogroups() []*relatedproductgroupResolver {
	//return r.user.device

	var l []*relatedproductgroupResolver
	for _, v := range r.product.relprogroups {
		l = append(l, &relatedproductgroupResolver{v})
	}
	return l
}

func (r *productResolver) Unrelprogroups() []*unrelatedproductgroupResolver {
	//return r.user.device

	var l []*unrelatedproductgroupResolver
	for _, v := range r.product.unrelprogroups {
		l = append(l, &unrelatedproductgroupResolver{v})
	}
	return l
}

//==================      RelatedProductGroup       ===========================

func (r *relatedproductgroupResolver) ID() graphql.ID {
	return r.relatedproductgroup.id
}

func (r *relatedproductgroupResolver) Grouptype() string {
	return r.relatedproductgroup.grouptype
}

func (r *relatedproductgroupResolver) Product() []*productResolver {
	//return r.user.device

	var l []*productResolver
	for _, v := range r.relatedproductgroup.product {
		l = append(l, &productResolver{v})
	}
	return l
}

//==================      UnRelatedProductGroup       ===========================

func (r *unrelatedproductgroupResolver) ID() graphql.ID {
	return r.unrelatedproductgroup.id
}

func (r *unrelatedproductgroupResolver) Grouptype() string {
	return r.unrelatedproductgroup.grouptype
}

func (r *unrelatedproductgroupResolver) Product() []*productResolver {
	//return r.user.device

	var l []*productResolver
	for _, v := range r.unrelatedproductgroup.product {
		l = append(l, &productResolver{v})
	}
	return l
}
