package mygraphql

import (
	"github.com/neelance/graphql-go"
	"model"
	"strconv"
	"fmt"
)

//struct for graphql
type account struct {
	id    graphql.ID
	name  string
	leads []*lead
}

//struct for upserting
type accountInput struct {
	Id    *graphql.ID
	Name  string
	Leads *[]leadInput
}

//struct for response
type accountResolver struct {
	account *account
}

func ResolveAccount(args struct{ ID graphql.ID }) (response []*accountResolver) {

	if args.ID != "" {
		response = append(response, &accountResolver{MapAccount(
			model.GetAccount(convertId(args.ID)),
		)})
		return response
	}
	for _, val := range model.GetAccounts() {
		response = append(response, &accountResolver{MapAccount(
			val,
		)})
	}

	return response
}

func ResolveCreateAccount(args *struct {
	Account *accountInput
}) *accountResolver {

	var account = &account{}

	if args.Account.Id == nil {
		//create account
		account = MapAccount(model.CreateAccount(args.Account.Name)) //since Name is required field in schema no need for null check
	} else {
		//update account
		account = MapAccount(model.UpdateAccount(convertId(*args.Account.Id), args.Account.Name))
	}

	if account != nil && args.Account.Leads != nil {

		//	account.leads = [len(*args.Account.Leads)]lead

		fmt.Println(len(*args.Account.Leads))

		for _, dev := range *args.Account.Leads {
			if dev.Id == nil {
				model.CreateLead(dev.Name, "account",convertId(account.id))
			} else {
				model.UpdateLead(convertId(*dev.Id), dev.Name, "account",convertId(account.id))
			}
		}

	}

	return &accountResolver{account}
}

//==================		fields resolvers		===========================

func (r *accountResolver) ID() graphql.ID {
	return r.account.id
}

func (r *accountResolver) Name() string {
	return r.account.name
}

//This method will run, if leads is asked for
func (r *accountResolver) Leads() []*leadResolver {

	var l []*leadResolver
	if r.account != nil {
		//if account not null get device of account from db and map
		lead := model.GetLeadsofType(convertId(r.account.id))
		for _, v := range lead {
			l = append(l, &leadResolver{MapLead(v)})
		}
		return l
	}

	for _, v := range r.account.leads {
		l = append(l, &leadResolver{v})
	}
	return l
}

//=================			mapper methods			==============================
func MapAccount(modelAccount model.Account) *account {

	if modelAccount == (model.Account{}) {
		return &account{}
	}

	//create graphql (account) from model (account)
	account := account{
		id:   graphql.ID(strconv.Itoa(modelAccount.Id)),
		name: modelAccount.Name,
	}
	return &account
}
