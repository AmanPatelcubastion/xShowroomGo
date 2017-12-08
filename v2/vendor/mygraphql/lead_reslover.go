package mygraphql

import (
	"github.com/neelance/graphql-go"
	"model"
	"strconv"
)

//struct for graphql
type lead struct {
	id     graphql.ID
	name   string
	accounts *account
}

//struct for upserting
type leadInput struct {
	Id     *graphql.ID
	Name   string
}

//struct for response
type leadResolver struct {
	lead *lead
}

func ResolveLead(args struct{ ID graphql.ID }) (response []*leadResolver) {

	if args.ID != "" {
		response = append(response, &leadResolver{MapLead(
			model.GetLead(convertId(args.ID)),
		)})
		return response
	}
	for _, val := range model.GetLeads() {
		response = append(response, &leadResolver{MapLead(
			val,
		)})
	}

	return response
}

func ResolveCreateLead(args *struct {
	Lead *leadInput
}) *leadResolver {

	var lead = &lead{}

	if args.Lead.Id == nil {
		//create device
		lead = MapLead(model.CreateLead(args.Lead.Name, -1)) //new device created set userId null
	} else {
		//update device
		lead = MapLead(model.UpdateLead(convertId(*args.Lead.Id), args.Lead.Name, -1)) //device updated keep userId whatever it was
	}

	return &leadResolver{lead}
}

//==================		fields resolvers		===========================

func (r *leadResolver) ID() graphql.ID {
	return r.lead.id
}

func (r *leadResolver) Name() string {
	return r.lead.name
}

//This method will run, if device is asked for
func (r *leadResolver) Accounts() *accountResolver {

	if r.lead != nil {
		//if device not null get user of device from db and map
		account := model.GetAccountOfLead(convertId(r.lead.id))
		return &accountResolver{MapAccount(account)}
	}
	return &accountResolver{r.lead.accounts}
}

//=================			mapper methods			==============================
func MapLead(modelLead model.Lead) *lead {

	if modelLead == (model.Lead{}) {
		return &lead{}
	}

	//create graphql (lead) from model (lead)
	lead := lead{
		id:   graphql.ID(strconv.Itoa(modelLead.Id)),
		name: modelLead.Name,
	}
	return &lead
}
