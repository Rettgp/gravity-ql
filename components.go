package main

import (
	"github.com/graphql-go/graphql"
)

var recipeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recipe",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.Int},
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
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Thumbnail        string   `json:"thumbnail"`
	Ingredients      []string `json:"ingredients"`
	DescriptionImage string   `json:"descriptionImage"`
	CookTime         int      `json:"cookTime"`
	PrepTime         int      `json:"prepTime"`
	Created          string   `json:"created"`
}
