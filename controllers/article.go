package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gipsh/golang-rest-api-boilerplate/models"
	"github.com/morkid/paginate"
	gorm "gorm.io/gorm"
)

type Article struct {
}

type CreateArticleInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

// @Summary Create a new Article
// @Description Crete a new Article based on json data
// @Accept  json
// @Produce  json
// @Tags Article
// @Param message body controllers.CreateArticleInput true "Article"
// @Header 200 {string} Token "qwerty"
// @success 200 {object} controllers.JSONResult{data=models.Article} "desc"
// @Router /api/v1/article [post]
func (controller *Article) Create(c *gin.Context) {

	var input CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)

	article := models.Article{Title: input.Title, Body: input.Body}

	err := db.Create(&article).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": article})

}

// @Summary Show an Article
// @Description get Article by ID
// @Accept  json
// @Produce  json
// @Tags Article
// @Param id path int true "Article ID"
// @success 200 {object} controllers.JSONResult{data=models.Article} "desc"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/article/{id} [get]
func (controller *Article) Get(c *gin.Context) {

	db := c.MustGet("DB").(*gorm.DB)

	articleID := c.Params.ByName("id")

	var article models.Article
	err := db.Where("id = ?", articleID).First(&article).Error
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": article})
	}

}

// @Summary Deletes a Article
// @Description Removes a Article from DB
// @Accept  json
// @Produce  json
// @Tags Article
// @Param  some_id path int true "ArticleID"
// @Header 200 {string} Token "qwerty"
// @success 200 {object} controllers.JSONResult{data=controllers.Deleted} "desc"
// @Router /api/v1/article/{id} [delete]
func (controller *Article) Delete(c *gin.Context) {

	db := c.MustGet("DB").(*gorm.DB)

	articleID := c.Params.ByName("id")

	var article models.Article
	err := db.Where("id = ?", articleID).First(&article).Error
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		db.Delete(&article)
		c.JSON(200, articleID)

		//		c.JSON(200, RenderDeletedResponse(articleID))
	}

}

// @Summary Update a Article
// @Description updates a Article by id
// @Accept  json
// @Produce  json
// @Tags Article
// @Param  some_id path int true "ArticleID"
// @Param message body controllers.CreateArticleInput true "Article"
// @Header 200 {string} Token "qwerty"
// @success 200 {object} controllers.JSONResult{data=models.Article} "desc"
// @Router /api/v1/article/{id} [put]
func (controller *Article) Update(c *gin.Context) {

	db := c.MustGet("DB").(*gorm.DB)

	articleID := c.Params.ByName("id")

	var article models.Article
	err := db.Where("id = ?", articleID).First(&article).Error
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		if err = c.ShouldBind(&article); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&article)
		c.JSON(http.StatusOK, gin.H{"data": article})
	}

}

// @Summary Lists all Articles
// @Description Lists all Articles
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Tags Article
// @Header 200 {string} Token "qwerty"
// @Success 200 {string} []models.Article
// @success 200 {object} controllers.JSONResult{data=[]models.Article} "desc"
// @Router /api/v1/articles [get]
func (controller *Article) List(c *gin.Context) {

	db := c.MustGet("DB").(*gorm.DB)

	paginate := c.MustGet("PAGINATE").(*paginate.Pagination)

	model := db.Model(&models.Article{})

	c.JSON(200, paginate.With(model).Request(c.Request).Response(&[]models.Article{}))

}
