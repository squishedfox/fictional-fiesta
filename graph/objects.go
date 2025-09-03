package graph

import (
	"github.com/graphql-go/graphql"
)

var (
	LabelPositionEnum = graphql.NewEnum(graphql.EnumConfig{
		Name: "LabelPositionEnum",
		Values: graphql.EnumValueConfigMap{
			"left":   &graphql.EnumValueConfig{Value: "left"},
			"top":    &graphql.EnumValueConfig{Value: "top"},
			"right":  &graphql.EnumValueConfig{Value: "right"},
			"bottom": &graphql.EnumValueConfig{Value: "bottom"},
		},
	})
	InputTypeEnum = graphql.NewEnum(graphql.EnumConfig{
		Name: "FieldType",
		Values: graphql.EnumValueConfigMap{
			"button":         &graphql.EnumValueConfig{Value: "button"},
			"checkbox":       &graphql.EnumValueConfig{Value: "checkbox"},
			"color":          &graphql.EnumValueConfig{Value: "color"},
			"date":           &graphql.EnumValueConfig{Value: "date"},
			"datetime-local": &graphql.EnumValueConfig{Value: "datetime-local"},
			"email":          &graphql.EnumValueConfig{Value: "email"},
			"file":           &graphql.EnumValueConfig{Value: "file"},
			"hidden":         &graphql.EnumValueConfig{Value: "hidden"},
			"image":          &graphql.EnumValueConfig{Value: "image"},
			"month":          &graphql.EnumValueConfig{Value: "month"},
			"number":         &graphql.EnumValueConfig{Value: "number"},
			"password":       &graphql.EnumValueConfig{Value: "password"},
			"radio":          &graphql.EnumValueConfig{Value: "radio"},
			"range":          &graphql.EnumValueConfig{Value: "range"},
			"reset":          &graphql.EnumValueConfig{Value: "reset"},
			"search":         &graphql.EnumValueConfig{Value: "search"},
			"submit":         &graphql.EnumValueConfig{Value: "submit"},
			"tel":            &graphql.EnumValueConfig{Value: "tel"},
			"text":           &graphql.EnumValueConfig{Value: "text"},
			"time":           &graphql.EnumValueConfig{Value: "time"},
			"url":            &graphql.EnumValueConfig{Value: "url"},
			"week":           &graphql.EnumValueConfig{Value: "week"},
		},
	})
	IDObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "ID",
		Description: "Hex string identifier of a record",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	FieldObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Field",
		Description: "A generic form field",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Unique field name for a form",
			},
			"type": &graphql.Field{
				Type: InputTypeEnum,
			},
			"ordinal": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Order for 0 to N where 0 is showing first, and N is showing last",
			},
		},
	})
	TextFieldObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "SimpleInputTextField",
		Description: "A simple text field for users to type in single line text",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Unique field name for a form",
			},
			"type": &graphql.Field{
				Type: InputTypeEnum,
			},
			"ordinal": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Order for 0 to N where 0 is showing first, and N is showing last",
			},
			"placeholder": &graphql.Field{
				Type:        graphql.String,
				Description: "Placeholder text to display to user",
			},
			"label": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Label to display next to input",
			},
			"labelPosition": &graphql.Field{
				Type:        LabelPositionEnum,
				Description: "Placement of where the label should be relative to the input field. Default left",
			},
		},
	})
	FieldSetObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "FieldSet",
		Description: "A grouping of fields within a form",
		Fields: graphql.Fields{
			"legend": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Grouping heading that should be unique on the form",
			},
			"fields": &graphql.Field{
				Type: graphql.NewList(FieldObject),
			},
		},
	})
	FormObject = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Form",
		Description: "Dynamic Form with groups or fields",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: IDObject,
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Unique form name that should be human readable",
			},
			"fieldsets": &graphql.Field{
				Type: graphql.NewList(FieldSetObject),
			},
		},
	})
)
