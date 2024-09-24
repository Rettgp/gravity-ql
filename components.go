package main

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var recipeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recipe",
	Fields: graphql.Fields{
		"uuid":             &graphql.Field{Type: graphql.String},
		"name":             &graphql.Field{Type: graphql.String},
		"thumbnail":        &graphql.Field{Type: graphql.String},
		"ingredients":      &graphql.Field{Type: graphql.NewList(graphql.String)},
		"descriptionImage": &graphql.Field{Type: graphql.String},
		"cookTime":         &graphql.Field{Type: graphql.Int},
		"prepTime":         &graphql.Field{Type: graphql.Int},
		"created":          &graphql.Field{Type: graphql.String},
	},
})

type Recipe struct {
	UUID             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Thumbnail        string    `json:"thumbnail"`
	Ingredients      []string  `json:"ingredients"`
	DescriptionImage string    `json:"descriptionImage"`
	CookTime         int       `json:"cookTime"`
	PrepTime         int       `json:"prepTime"`
	Created          string    `json:"created"`
}
