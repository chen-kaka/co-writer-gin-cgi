package controllers
import (
	"../models/mongodb"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	"errors"
	"runtime/debug"
)

type RepositoryController struct{}

/**
    http://localhost:9000/app/repository/create
    u_id: 58538dcc9822d109091c1d51
    POST
 */
func (ctrl RepositoryController)CreateRepo(c *gin.Context)  {
	repository := mongodb.Repository{}
	err := c.BindJSON(&repository)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusExpectationFailed, gin.H{"data": bson.M{"msg":"params not enough."}})
		return
	}

	if repository.Name == "" || repository.UId == "" || repository.Description == "" {
		c.Error(errors.New("params not enough."))
		c.JSON(http.StatusExpectationFailed, gin.H{"data": bson.M{"msg":"params not enough."}})
		return
	}
	fmt.Println("desc: ", repository.Description)

	repository.Id = bson.NewObjectId()
	repository.CreateAt = time.Now()
	repository.LastUpdate = time.Now()
	repository.Status = 0
	repository.Type = 0

	fmt.Println("insert repo: ", repository)
	db := c.MustGet("db").(*mgo.Database)
	err = db.C(mongodb.CollectionRepository).Insert(repository)
	if err != nil {
		c.Error(err)
		debug.PrintStack()
		c.JSON(http.StatusExpectationFailed, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bson.M{"succ":true}})
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
