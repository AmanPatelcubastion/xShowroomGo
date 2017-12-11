package mygraphql

import (
	"github.com/neelance/graphql-go"
	"model"
	"strconv"
	"fmt"
)

//struct for graphql
type user struct {
	id     graphql.ID
	name   string
	device *device
	leads   []*lead
}

//struct for upserting
type userInput struct {
	Id     *graphql.ID
	Name   string
	Device *deviceInput
	Leads *[]leadInput
}

//struct for response
type userResolver struct {
	user *user
}

func ResolveUser(args struct{ ID graphql.ID }) (response []*userResolver) {

	if args.ID != "" {
		response = append(response, &userResolver{MapUser(
			model.GetUser(convertId(args.ID)),
		)})
		return response
	}
	for _, val := range model.GetUsers() {
		response = append(response, &userResolver{MapUser(
			val,
		)})
	}

	return response
}

func ResolveCreateUser(args *struct {
	User *userInput
}) *userResolver {

	var user = &user{}

	if args.User.Id == nil {
		//create user
		user = MapUser(model.CreateUser(args.User.Name)) //since Name is required field in schema no need for null check
	} else {
		//update user
		user = MapUser(model.UpdateUser(convertId(*args.User.Id), args.User.Name))
	}

	if user != nil && args.User.Device != nil {

		if args.User.Device.Id == nil {
			user.device = MapDevice(model.CreateDevice(
				args.User.Device.Device_uuid,
				convertId(user.id),
			))
		} else {
			user.device = MapDevice(model.UpdateDevice(
				convertId(*args.User.Device.Id),
				args.User.Device.Device_uuid,
				convertId(user.id),
			))
		}
	}

	if user != nil && args.User.Leads != nil {

		//	user.leads = [len(*args.Account.Leads)]lead

		fmt.Println(len(*args.User.Leads))

		for _, dev := range *args.User.Leads {
			if dev.Id == nil {
				model.CreateLead(dev.Name, "user",convertId(user.id))
			} else {
				model.UpdateLead(convertId(*dev.Id), dev.Name,"user" ,convertId(user.id))
			}
		}

	}

	return &userResolver{user}
}

func ResolveDeleteUser(args struct{ ID graphql.ID }) (response *userResolver) {

	if args.ID != "" {
		response = &userResolver{MapUser(
			model.DeleteUser(convertId(args.ID)),
		)}
		return response
	}

	return response
}


//==================		fields resolvers		===========================

func (r *userResolver) ID() graphql.ID {
	return r.user.id
}

func (r *userResolver) Name() string {
	return r.user.name
}

//This method will run, if device is asked for
func (r *userResolver) Device() *deviceResolver {

	if r.user != nil {
		//if user not null get device of user from db and map
		device := model.GetDeviceOfUser(convertId(r.user.id))
		return &deviceResolver{MapDevice(device)}
	}
	return &deviceResolver{r.user.device}
}

func (r *userResolver) Leads() []*leadResolver {

	var l []*leadResolver
	if r.user != nil {
		//if account not null get device of account from db and map
		lead := model.GetLeadsofType(convertId(r.user.id))
		for _, v := range lead {
			l = append(l, &leadResolver{MapLead(v)})
		}
		return l
	}

	for _, v := range r.user.leads {
		l = append(l, &leadResolver{v})
	}
	return l
}

//=================			mapper methods			==============================
func MapUser(modelUser model.User) *user {

	if modelUser == (model.User{}) {
		return &user{}
	}

	//create graphql (user) from model (user)
	user := user{
		id:   graphql.ID(strconv.Itoa(modelUser.Id)),
		name: modelUser.Name,
	}
	return &user
}
