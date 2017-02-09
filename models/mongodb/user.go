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
	Username 	string 		`json:"username,omitempty" bson:"username"`
	Name 		string 		`json:"name,omitempty" bson:"name"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	LastUpdate 	time.Time	`json:"last_update,omitempty" form: "last_update" bson:"last_update"`
	Sex 		string		`json:"sex,omitempty" form: "sex" bson:"sex"`
	Nickname 	string 		`json:"nickname,omitempty" form:"nickname" bson:"nickname"`
	Birthday 	time.Time	`json:"birthday,omitempty" form: "birthday" bson:"birthday"`
	Email 		string		`json:"email,omitempty" form: "email" bson:"email"`
	LastLogin 	time.Time 	`json:"last_login,omitempty" form: "last_login" bson:"last_login"`
	Avatar 		string		`json:"avatar,omitempty" form: "avatar" bson:"avatar"`
	Summary 	string		`json:"summary,omitempty" form: "summary" bson:"summary"`
}