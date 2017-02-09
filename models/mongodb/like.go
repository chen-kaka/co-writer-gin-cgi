package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionLike = "like"
)

//用户对仓库节点点赞
type Like struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"`
	NodeId 		bson.ObjectId	`json:"node_id,omitempty" form: "node_id" bson:"node_id"`
	Like 		int 		`json:"like,omitempty" form:"like" bson:"like"`
}