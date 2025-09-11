package mongodb

import (
	"log"

	"github.com/squishedfox/fictional-fiesta/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func mapFieldSetInputModelToDocument(model *db.FormInputModel) bson.M {
	return bson.M{
		"label":     model.Label,
		"inputType": model.Type,
		"minLength": model.MinLength,
		"maxLength": model.MaxLength,
		"min":       model.Min,
		"max":       model.Max,
		"required":  model.Required,
		"multiple":  model.Multiple,
	}
}

func mapFieldsetModelToDocument(model *db.FieldSetModel) bson.M {
	inputs := []bson.M{}
	for _, input := range model.Inputs {
		mapped := mapFieldSetInputModelToDocument(&input)
		inputs = append(inputs, mapped)
	}
	return bson.M{
		"legend": model.Legend,
		"inputs": inputs,
	}
}

func mapFormModelToDocument(model *db.CreateFormModel) bson.M {
	fieldsets := []bson.M{}

	log.Println(model.Fieldsets)

	for _, fieldset := range model.Fieldsets {
		mapped := mapFieldsetModelToDocument(&fieldset)
		fieldsets = append(fieldsets, mapped)
	}

	return bson.M{
		"name":      model.Name,
		"active":    model.Active,
		"fieldsets": fieldsets,
	}
}
