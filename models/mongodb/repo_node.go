package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionRepoNode = "repo_node"
)

//仓库节点
type RepoNode struct {
	Id        	bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	Name 		string 		`json:"name,omitempty" bson:"name"` // binding:"required"
	Creator		string 		`json:"creator,omitempty" form: creator bson:"creator"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: create_at bson:"create_at"`
	LastUpdate 	time.Time	`json:"last_update,omitempty" form: last_update bson:"last_update"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"` // binding:"required"
	RepoId 		bson.ObjectId	`json:"repo_id,omitempty" form: repo_id bson:"repo_id"`
	FrontNode 	bson.ObjectId	`json:"front_node,omitempty" form: front_node bson:"front_node"`
	Description 	string		`json:"description,omitempty" form: description bson:"description"`
	Type 		int 		`json: "type,omitempty" form:"type" bson:"type"`
	Status 		int		`json:"status,omitempty" form: "status" bson:"status"`
	Level 		int		`json:"level,omitempty" form: "level" bson:"level"`
	NodeList	bson.M		`json:"node_list,omitempty" form: "node_list" bson:"node_list"`
	Content		bson.M		`json:"content,omitempty" form: "content" bson:"content"`
}
