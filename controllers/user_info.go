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

/**
    http://localhost:9000/app/user_info/create
    {"username": "chen-kaka","name":"kakachan","nickname":"kaka","email":"chen-kaka@163.com",
    "avatar":"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1486632351469&di=5aed748c89a6252365bf52c9ff652224&imgtype=0&src=http%3A%2F%2Fwww.qqzhi.com%2Fuploadpic%2F2014-09-26%2F215403763.jpg"
    ,"summary":"work hard."}
    POST
 */
func (ctrl UserInfoController)CreateUser(c *gin.Context)  {
	user := mongodb.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusExpectationFailed, gin.H{"data": bson.M{"msg":"params not enough."}})
		return
	}
	if user.Username == "" || user.Name == "" || user.Nickname == "" || user.Email == "" {
		c.Error(errors.New("params not enough."))
		c.JSON(http.StatusExpectationFailed, gin.H{"data":bson.M{"msg": "params not enough."}})
		return
	}

	user.Id = bson.NewObjectId()
	user.CreateAt = time.Now()
	user.LastUpdate = time.Now()
	user.LastLogin = time.Now()

	fmt.Println("insert user is:", user)
	db := c.MustGet("db").(*mgo.Database)
	err = db.C(mongodb.CollectionUser).Insert(user)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusExpectationFailed, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":bson.M{"succ":true}})
}