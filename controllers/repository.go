package controllers
import (
	"../models/mongodb"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"net/http"
)

type RepositoryController struct{}

func (ctrl RepositoryController)createRepo(c *gin.Context)  {
	db := c.MustGet("db").(*mgo.Database)

	repository := mongodb.Repository{}
	err := c.Bind(&repository)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(mongodb.CollectionRepository).Insert(repository)
	if err != nil {
		c.Error(err)
	}
}

// List all repository
func (ctrl RepositoryController)List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	articles := []mongodb.Repository{}
	err := db.C(mongodb.CollectionRepository).Find(nil).Sort("-updated_on").All(&articles)
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "articles/list", gin.H{
		"title":    "Articles",
		"articles": articles,
	})
}