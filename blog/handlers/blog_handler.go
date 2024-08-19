package handlers

import (
	"blog/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Id     string `json:id`
	Author string `json:author`
	Body   string `json:body`
}

var Blogs []Blog

func GetAllBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, Blogs)
}

func GetBlogById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "id not found",
		})
		return
	}

	middlewares.AdminMiddleware(c)

	for _, ele := range Blogs {
		if ele.Id == id {
			c.JSON(http.StatusOK, ele)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("Blog with %s not found", id),
	})
	// return
}

func CreateBlog(c *gin.Context){
	var blog Blog;

	if err:= c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "json binding error while creating blog",
		})
		return
	}

	Blogs = append(Blogs, blog)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v created", blog),
	})
}

func UpadateBlog(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "id not found",
		})
		return
	}

	var blog Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "json binding error while updating",
		})
		return
	}

	for i, ele := range Blogs {
		if ele.Id == id {
			Blogs[i] = blog
			c.JSON(http.StatusOK, gin.H{
				"message" : fmt.Sprintf("%v updated", blog),
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error" : fmt.Sprintf("Blog of id %s not found", id),
	})
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "id not found",
		})
		return
	}

	for i, ele := range Blogs {
		if ele.Id == id {
			Blogs = append(Blogs[:i],Blogs[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message" : fmt.Sprintf("%v Deleted", ele),
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error" : fmt.Sprintf("Blog of id %s not found", id),
	})
}
