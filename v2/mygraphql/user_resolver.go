package mygraphql

import (
	"github.com/neelance/graphql-go"
	"github.com/aatishrana/GraphQLTesting/xShowroom/v2/model"
	"strconv"
)

//struct for graphql
type user struct {
	id     graphql.ID
	name   string
	device *device
}

//struct for upserting
type userInput struct {
	Id     *graphql.ID
	Name   string
	Device *deviceInput
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

	return &userResolver{user}
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
