package mygraphql

import (
	"github.com/neelance/graphql-go"
	"model"
	"strconv"
)

//struct for graphql
type device struct {
	id          graphql.ID
	device_uuid string
	user        *user
}

//struct for upserting
type deviceInput struct {
	Id          *graphql.ID
	Device_uuid string
}

//struct for response
type deviceResolver struct {
	device *device
}

func ResolveDevice(args struct{ ID graphql.ID }) []*deviceResolver {
	var d []*deviceResolver
	if args.ID != "" {
		d = append(d, &deviceResolver{MapDevice(
			model.GetDevice(convertId(args.ID)),
		)})
		return d
	}
	for _, val := range model.GetDevices() {
		d = append(d, &deviceResolver{MapDevice(
			val,
		)})
	}
	return d
}

func ResolveCreateDevice(args *struct {
	Device *deviceInput
}) *deviceResolver {

	var device = &device{}

	if args.Device.Id == nil {
		//create device
		device = MapDevice(model.CreateDevice(args.Device.Device_uuid, -1)) //new device created set userId null
	} else {
		//update device
		device = MapDevice(model.UpdateDevice(convertId(*args.Device.Id), args.Device.Device_uuid, -1)) //device updated keep userId whatever it was
	}

	return &deviceResolver{device}
}

//==================		fields resolvers		===========================

func (r *deviceResolver) ID() graphql.ID {
	return r.device.id
}

func (r *deviceResolver) DeviceUuid() string {
	return r.device.device_uuid
}

//This method will run, if user is asked for
func (r *deviceResolver) User() *userResolver {

	if r.device != nil {
		//if device not null get user of device from db and map
		user := model.GetUserOfDevice(convertId(r.device.id))
		return &userResolver{MapUser(user)}
	}
	return &userResolver{r.device.user}
}

//=================			mapper methods			==============================
func MapDevice(modelDevice model.Device) *device {

	if modelDevice == (model.Device{}) {
		return &device{}
	}

	//create graphql (device) from model (device)
	device := device{
		id:          graphql.ID(strconv.Itoa(modelDevice.Id)),
		device_uuid: modelDevice.UUID,
	}
	return &device
}
