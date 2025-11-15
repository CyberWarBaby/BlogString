package routes

import (
	"net/http"
	"strconv"

	"example.com/blog-api/models"
	"github.com/gin-gonic/gin"
)

func mainPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This BlogString"})
}

func getBlogs(context *gin.Context) {
	blogs, err := models.GetAllBlogs()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch blogs"})
		return
	}

	// if len(blogs) == 0 {
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"message": "no blogs found",
	// 		"blogs":   []any{},
	// 	})
	// 	return
	// }

	context.JSON(http.StatusOK, blogs)
}

func createBlogs(context *gin.Context) {
	var blog models.Blog

	// bind only title & content from JSON
	if err := context.ShouldBindJSON(&blog); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	// assign author_id inside Save(), so client cannot override
	err := blog.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create blog"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "successfully created blog",
		"blog":    blog,
	})
}

func getBlog(context *gin.Context) {
	blogId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse blog id."})
		return
	}

	blog, err := models.GetBlogById(blogId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch blog"})
		return
	}

	context.JSON(http.StatusOK, blog)
}
