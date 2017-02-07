package controllers
import (
	"../models/mongodb"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type RepositoryController struct{}

func (ctrl RepositoryController)CreateRepo(c *gin.Context)  {
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

/**
    http://localhost:9000/app/repository/list?queryStr=ttt
 */
func (ctrl RepositoryController)QueryMyRepos(c *gin.Context)  {
	var offset int = 0
	u_id := c.Query("u_id")
	limit := 10
	queryStr := c.Query("queryStr")

	if c.Query("limit") != "" {
		limit,_ = strconv.Atoi(c.Query("limit"))
	}
	if c.Query("offset") != "" {
		offset,_ = strconv.Atoi(c.Query("offset"))
	}

	db := c.MustGet("db").(*mgo.Database)
	repos := []mongodb.Repository{}
	queryRepo := bson.M{"name": bson.RegEx{queryStr,""}}
	if u_id != "" {
		queryRepo["u_id"] = bson.ObjectId(u_id)
	}
	fmt.Println("query my repos: ", queryRepo)
	err := db.C(mongodb.CollectionRepository).Find(&queryRepo).Skip(offset).Limit(limit).Sort("-create_at").All(&repos)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": repos})
}