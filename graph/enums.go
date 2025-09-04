package graph

import "github.com/graphql-go/graphql"

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
			"datetime_local": &graphql.EnumValueConfig{Value: "datetime-local"},
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
)
