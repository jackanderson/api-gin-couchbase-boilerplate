package controllers

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	models "github.com/plagiari-sm/api-gin-couchbase-boilerplate/models"
)

// ArticleCTRL : Articles Contoller
type ArticleCTRL struct {
	Bucket *gocb.Bucket
}

// Create : Create an Article
func (ctrl *ArticleCTRL) Create(c *gin.Context) {
	var article models.Article
	if err := c.Bind(&article); err != nil {
		ResponseError(c, 404, err.Error())
	} else {
		article.ID = uuid.Must(uuid.NewUUID())
		article.Type = "article"
		article.CreatedAt = time.Now()
		article.UpdatedAt = time.Now()
		if res, err := ctrl.Bucket.Upsert(article.ID.String(), &article, 0); err != nil {
			ResponseError(c, 500, err.Error())
		} else {
			ResponseJSON(c, res)
		}
	}
}

// Read : Get Articles
func (ctrl *ArticleCTRL) Read(c *gin.Context) {
	query := gocb.NewN1qlQuery("Select * From `" + ctrl.Bucket.Name() + "` As data")
	rows, err := ctrl.Bucket.ExecuteN1qlQuery(query, nil)

	if err != nil {
		ResponseError(c, 404, err.Error())
		return
	}
	type response struct {
		Data models.Article `json:"data"`
	}
	var row response
	var data []models.Article
	for rows.Next(&row) {
		data = append(data, row.Data)
	}

	fmt.Println(data)
	ResponseJSON(c, data)
}

// ReadOne : Get Article by ID
func (ctrl *ArticleCTRL) ReadOne(c *gin.Context) {
	var article models.Article
	_, err := ctrl.Bucket.Get(c.Param("id"), &article)
	if err != nil {
		ResponseError(c, 404, err.Error())
		return
	}
	ResponseJSON(c, &article)
}

// Update : Update an Article by ID
func (ctrl *ArticleCTRL) Update(c *gin.Context) {
	var article models.Article
	if err := c.Bind(&article); err != nil {
		ResponseError(c, 404, err.Error())
	} else {
		article.UpdatedAt = time.Now()
		if res, err := ctrl.Bucket.Upsert(c.Param("id"), &article, 0); err != nil {
			ResponseError(c, 500, err.Error())
		} else {
			ResponseJSON(c, res)
		}
	}
}

// Delete : Delete an Article
func (ctrl *ArticleCTRL) Delete(c *gin.Context) {
	if _, err := ctrl.Bucket.Remove(c.Param("id"), 0); err != nil {
		ResponseError(c, 500, err.Error())
	} else {
		ResponseJSON(c, c.Param("id"))
	}
}
