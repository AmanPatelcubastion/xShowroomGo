package mygraphql

import (
	"github.com/neelance/graphql-go"
	"model"
	"strconv"
)

//struct for graphql
type relatedproductgroup struct {
	id    graphql.ID
	grouptype  string
	products []*product
}

//struct for upserting
type relatedproductgroupInput struct {
	Id    *graphql.ID
	GroupType  string
	Productids *[]graphql.ID
	//Products *[]productInput
}

//struct for response
type relatedproductgroupResolver struct {
	relatedproductgroup *relatedproductgroup
}

func ResolveRelatedproductgroup(args struct{ ID graphql.ID }) (response []*relatedproductgroupResolver) {

	if args.ID != "" {
		response = append(response, &relatedproductgroupResolver{MapRelatedProductGroup(
			model.GetRelatedProductGroup(convertId(args.ID)),
		)})
		return response
	}
	for _, val := range model.GetRelatedProductGroups() {
		response = append(response, &relatedproductgroupResolver{MapRelatedProductGroup(
			val,
		)})
	}

	return response
}

func ResolveCreateRelatedproductgroup(args *struct {
	RelatedProductGroup *relatedproductgroupInput
}) *relatedproductgroupResolver {

	var relatedproductgroup = &relatedproductgroup{}

	if args.RelatedProductGroup.Id == nil {
		//create relatedproductgroup
		relatedproductgroup = MapRelatedProductGroup(model.CreateRelatedProductGroup(args.RelatedProductGroup.GroupType,args.RelatedProductGroup.Productids)) //since Name is required field in schema no need for null check
	} else {
		//update relatedproductgroup
		relatedproductgroup = MapRelatedProductGroup(model.UpdateRelatedProductGroup(convertId(*args.RelatedProductGroup.Id), args.RelatedProductGroup.GroupType,args.RelatedProductGroup.Productids))
	}

	/*if relatedproductgroup != nil && args.RelatedProductGroup.Products != nil {

		//	relatedproductgroup.leads = [len(*args.Account.Leads)]lead

		fmt.Println(len(*args.RelatedProductGroup.Products))

		for _, dev := range *args.RelatedProductGroup.Products {
			if dev.Id == nil {
				model.CreateProduct(dev.Name, convertId(relatedproductgroup.id))
			} else {
				model.UpdateRelatedProductGroup(convertId(*dev.Id), dev.Name, convertId(relatedproductgroup.id))
			}
		}

	}*/

	return &relatedproductgroupResolver{relatedproductgroup}
}

//==================		fields resolvers		===========================

func (r *relatedproductgroupResolver) ID() graphql.ID {
	return r.relatedproductgroup.id
}

func (r *relatedproductgroupResolver) GroupType() string {
	return r.relatedproductgroup.grouptype
}

//This method will run, if device is asked for
func (r *relatedproductgroupResolver) Products() []*productResolver {

	var l []*productResolver
	if r.relatedproductgroup != nil {
		//if product not null get device of product from db and map
		prod := model.GetProductOfRelatedProductGroups(convertId(r.relatedproductgroup.id))
		for _, v := range prod {
			l = append(l, &productResolver{MapProduct(v)})
		}
		return l
	}

	for _, v := range r.relatedproductgroup.products {
		l = append(l, &productResolver{v})
	}
	return l
}

//=================			mapper methods			==============================
func MapRelatedProductGroup(modelRelatedProductGroup model.Relatedproductgroup) *relatedproductgroup {

	//if modelRelatedProductGroup == (model.Relatedproductgroup{}) {
	//	return &relatedproductgroup{}
	//}

	//create graphql (relatedproductgroup) from model (relatedproductgroup)
	relatedproductgroup := relatedproductgroup{
		id:   graphql.ID(strconv.Itoa(modelRelatedProductGroup.Id)),
		grouptype: modelRelatedProductGroup.GroupType,
	}
	return &relatedproductgroup
}

//Graphql to Db
func MapRelatedProductGroup1(modelRelatedProductGroup *[]relatedproductgroupInput) []model.Relatedproductgroup {

	var RelproGrp []model.Relatedproductgroup
    if modelRelatedProductGroup!=nil{
    	for _,v:=range *modelRelatedProductGroup{
			id1,_:=strconv.Atoi(string(*v.Id))
		//	fmt.Printf("dfsf : %T",*v.Id)
			RelproGrp = append(RelproGrp,model.Relatedproductgroup{
				Id:   id1,
				GroupType: v.GroupType,
			})
		}
	}

	return RelproGrp
}

