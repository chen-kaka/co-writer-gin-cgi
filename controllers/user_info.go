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
)

type UserInfoController struct{}

/**
	http://localhost:9000/app/user_info/info?id=58538dcc9822d109091c1d51
 */
func (ctrl UserInfoController)QueryInfo(c *gin.Context)  {
	queryId := c.Query("id")
	if queryId == "" {
		c.Error(errors.New("params not enough."))
		c.JSON(http.StatusExpectationFailed, gin.H{"data": bson.M{"msg":"请输入id参数。"}})
		return;
	}

	db := c.MustGet("db").(*mgo.Database)
	user := mongodb.User{}
	//queryId = strings.TrimSpace(queryId)
	fmt.Println("id is:", queryId, ",length is:",len(queryId) )
	err := db.C(mongodb.CollectionUser).FindId(bson.ObjectIdHex(queryId)).One(&user)
	if err != nil {
		c.Error(err)
	}
	fmt.Println("ret user is: ", user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}