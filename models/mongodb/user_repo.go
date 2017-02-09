package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionUserRepo = "user_repo"
)

//用户 -- 仓库 关系表
type UserRepo struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"`
	RepoId 		bson.ObjectId	`json:"repo_id,omitempty" form: "repo_id" bson:"repo_id"`
	Type 		int 		`json: "type,omitempty" form:"type" bson:"type"`
	Status 		int		`json:"status,omitempty" form: "status" bson:"status"`
}