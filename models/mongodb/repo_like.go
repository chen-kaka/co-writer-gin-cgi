package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionRepoLike = "repo_like"
)
//用户对仓库点赞
type RepoLike struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"`
	RepoId 		bson.ObjectId	`json:"repo_id,omitempty" form: "repo_id" bson:"repo_id"`
	Like 		int 		`json:"like,omitempty" form:"like" bson:"like"`
}