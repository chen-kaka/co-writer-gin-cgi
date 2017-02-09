package controllers

import (
	"../models/mongodb"
	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"errors"
	//"strings"
	"time"
)

type CommentController struct{}

/**
	http://localhost:9000/app/comment/query_repo_comment?repo_id=585a1c1a36af130950b1d5be
 */
func (ctrl CommentController)QueryRepoComment(c *gin.Context)  {
	repoId := c.Query("repo_id")
	if repoId == "" {
		c.Error(errors.New("params not enough."))
		c.JSON(http.StatusExpectationFailed, gin.H{"data": bson.M{"msg":"请输入repoId参数。"}})
		return;
	}

	db := c.MustGet("db").(*mgo.Database)
	repoComment := []mongodb.RepoComment{}
	//queryId = strings.TrimSpace(queryId)
	fmt.Println("repoId is:", repoId, ",length is:",len(repoId) )
	err := db.C(mongodb.CollectionRepoComment).Find(bson.M{"repo_id":bson.ObjectIdHex(repoId)}).All(&repoComment)
	if err != nil {
		c.Error(err)
	}
	fmt.Println("ret repoComment is: ", repoComment)
	c.JSON(http.StatusOK, gin.H{"data": repoComment})
}

/**
    http://localhost:9000/app/comment/create_repo_comment
    {"u_id": "58538dcc9822d109091c1d51","repo_id": "585a1c1a36af130950b1d5be","title":"title~", "comment": "repo u sucks"}
    POST
 */
func (ctrl CommentController)CreateRepoComment(c *gin.Context)  {
	repoComment := mongodb.RepoComment{}
	err := c.BindJSON(&repoComment)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusExpectationFailed, gin.H{"data": bson.M{"msg":"params not enough."}})
		return
	}
	if repoComment.UId.Hex() == "" || repoComment.RepoId.Hex() == "" || repoComment.Title == "" || repoComment.Comment == "" {
		c.Error(errors.New("params not enough."))
		c.JSON(http.StatusExpectationFailed, gin.H{"data":bson.M{"msg": "params not enough."}})
		return
	}

	repoComment.Id = bson.NewObjectId()
	repoComment.CreateAt = time.Now()

	fmt.Println("insert repoComment is:", repoComment)
	db := c.MustGet("db").(*mgo.Database)
	err = db.C(mongodb.CollectionRepoComment).Insert(repoComment)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusExpectationFailed, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":bson.M{"succ":true}})
}