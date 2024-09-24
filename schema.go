package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ImportJsonData(fileName string, result interface{}) (ok bool) {
	ok = true
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Print("Error: ", err)
		ok = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		ok = false
		fmt.Print("Error: ", err)
	}

	return
}

func ExportJsonData(fileName string, data interface{}) (ok bool) {
	ok = true
	content, err := json.Marshal(data)

	if err != nil {
		fmt.Print("Error: ", err)
		ok = false
	}
	err = os.WriteFile(fileName, content, fs.FileMode(os.O_WRONLY))
	if err != nil {
		ok = false
		fmt.Print("Error: ", err)
	}

	return
}

var RecipeList []Recipe
var _ = ImportJsonData("./recipeData.json", &RecipeList)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"recipe": &graphql.Field{
			Type:        recipeType,
			Description: "Get a single recipe",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				idQuery, ok := params.Args["uuid"].(string)
				uuid, okUUID := uuid.Parse(idQuery)
				if ok && okUUID != nil {
					for _, recipe := range RecipeList {
						if recipe.UUID == uuid {
							return recipe, nil
						}
					}
				}

				return nil, nil
			},
		},
		"recipeList": &graphql.Field{
			Type:        graphql.NewList(recipeType),
			Description: "List of recipes",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return RecipeList, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addRecipe": &graphql.Field{
			Type:        recipeType,
			Description: "Add a new recipe",
			Args: graphql.FieldConfigArgument{
				"name":             &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"thumbnail":        &graphql.ArgumentConfig{Type: graphql.String},
				"ingredients":      &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				"descriptionImage": &graphql.ArgumentConfig{Type: graphql.String},
				"cookTime":         &graphql.ArgumentConfig{Type: graphql.Int},
				"prepTime":         &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
				thumbnail, _ := params.Args["thumbnail"].(string)
				ingredients, _ := params.Args["ingredients"].([]string)
				descriptionImage, _ := params.Args["descriptionImage"].(string)
				cookTime, _ := params.Args["cookTime"].(int)
				prepTime, _ := params.Args["prepTime"].(int)

				currentTime := time.Now()

				newRecipe := Recipe{
					Name:             name,
					Thumbnail:        thumbnail,
					Ingredients:      ingredients,
					DescriptionImage: descriptionImage,
					CookTime:         cookTime,
					PrepTime:         prepTime,
					UUID:             uuid.New(),
					Created:          currentTime.Format("2006-01-02"),
				}

				RecipeList = append(RecipeList, newRecipe)
				ExportJsonData("./recipeData.json", &RecipeList)

				return newRecipe, nil
			},
		},
		"updateRecipe": &graphql.Field{
			Type:        recipeType,
			Description: "Update existing recipe",
			Args: graphql.FieldConfigArgument{
				"uuid":             &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"name":             &graphql.ArgumentConfig{Type: graphql.String},
				"thumbnail":        &graphql.ArgumentConfig{Type: graphql.String},
				"ingredients":      &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				"descriptionImage": &graphql.ArgumentConfig{Type: graphql.String},
				"cookTime":         &graphql.ArgumentConfig{Type: graphql.Int},
				"prepTime":         &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				uuidString, _ := params.Args["uuid"].(string)
				affectedRecipe := Recipe{}
				uuid, _ := uuid.Parse(uuidString)

				for i := 0; i < len(RecipeList); i++ {
					if RecipeList[i].UUID == uuid {
						if _, ok := params.Args["name"]; ok {
							RecipeList[i].Name = params.Args["name"].(string)
						}
						if _, ok := params.Args["thumbnail"]; ok {
							RecipeList[i].Thumbnail = params.Args["thumbnail"].(string)
						}
						if _, ok := params.Args["ingredients"]; ok {
							RecipeList[i].Ingredients = params.Args["ingredients"].([]string)
						}
						if _, ok := params.Args["descriptionImage"]; ok {
							RecipeList[i].DescriptionImage = params.Args["descriptionImage"].(string)
						}
						if _, ok := params.Args["cookTime"]; ok {
							RecipeList[i].CookTime = params.Args["cookTime"].(int)
						}
						if _, ok := params.Args["prepTime"]; ok {
							RecipeList[i].PrepTime = params.Args["prepTime"].(int)
						}

						ExportJsonData("./recipeData.json", &RecipeList)
						affectedRecipe = RecipeList[i]
						break
					}
				}
				return affectedRecipe, nil
			},
		},
	},
})

var RecipeSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
