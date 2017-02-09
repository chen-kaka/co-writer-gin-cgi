package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionRepoComment = "repo_comment"
)

//用户评论 -- 仓库
type RepoComment struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"`
	RepoId 		bson.ObjectId	`json:"repo_id,omitempty" form: "repo_id" bson:"repo_id"`
	Title 		string 		`json:"title,omitempty" form:"title" bson:"title"`
	Comment 	string		`json:"comment,omitempty" form: "comment" bson:"comment"`
}