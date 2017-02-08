package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionUser = "user"
)

type User struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	Username 	string 		`json:"username" bson:"username"`
	Name 		string 		`json:"name" bson:"name"`
	CreateAt	time.Time	`json:"create_at" form: "create_at" bson:"create_at"`
	LastUpdate 	time.Time	`json:"last_update" form: "last_update" bson:"last_update"`
	Sex 		string		`json:"sex" form: "sex" bson:"sex"`
	Nickname 	string 		`json:"nickname" form:"nickname" bson:"nickname"`
	Birthday 	time.Time	`json:"birthday" form: "birthday" bson:"birthday"`
	Email 		string		`json:"email" form: "email" bson:"email"`
	LastLogin 	time.Time 	`json:"last_login" form: "last_login" bson:"last_login"`
	Avatar 		string		`json:"avatar" form: "avatar" bson:"avatar"`
	Summary 	string		`json:"summary" form: "summary" bson:"summary"`
}