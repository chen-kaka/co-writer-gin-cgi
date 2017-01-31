package controllers
import (
	"../models/mongodb"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
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
	fmt.Printf("repository List")
	db := c.MustGet("db").(*mgo.Database)
	repos := []mongodb.Repository{}
	err := db.C(mongodb.CollectionRepository).Find(nil).Sort("-create_at").All(&repos)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": repos})
}