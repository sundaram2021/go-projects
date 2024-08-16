package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Id string `json:id`
	Name string `json:name`
	Tag []string `json:tag`
	Ingredients []string `json:ingredients`
	Instructions []string `json:instructions`
	PublishedAt string `json:publishedAt`
}

var Recipes []Recipe

func init(){
	file, _ := ioutil.ReadFile("recipes.json")

	_ = json.Unmarshal([]byte(file), &Recipes)
}

func GetAllRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, Recipes)
}


func GetRecipeById(c *gin.Context){
	var recipe Recipe
	id := c.Param("Id")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "id is empty",
		})
		return
	}

	for _, ele := range Recipes {
		if ele.Id == id {
			recipe = ele
			c.JSON(http.StatusOK, recipe)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "id is not matched",
	})
}

func CreateRecipe(c *gin.Context){
	var recipe Recipe

	if err := c.BindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "json binding error while creating recipes",
		})
	}

	Recipes = append(Recipes, recipe)

	c.JSON(http.StatusOK, Recipes)
}

func UpdateRecipe(c *gin.Context){
	var recipe Recipe
	id := c.Param("Id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is empty",
		})
		return
	}

	if err := c.BindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "json encoding error while updating",
		})
	}

	fmt.Println("id", id)

	for idx, ele := range Recipes {
		if ele.Id == id {
			Recipes[idx] = recipe
			c.JSON(http.StatusOK, recipe)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("recipe of %v not found", id),
	})
	// return
}

func DeleteRecipe(c *gin.Context) {
	// var recipe Recipe
	id := c.Param("Id")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": "id is empty",
		})
		return
	}

	for i , ele := range Recipes {
		if ele.Id == id {
			// recipe = ele
			Recipes = append(Recipes[:i],Recipes[i+1:]... )
			c.JSON(http.StatusOK, gin.H{
				"message" : "deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "id is not matched",
	})
}