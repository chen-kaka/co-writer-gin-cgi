package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionRepository = "repository"
)

type Repository struct {
	Id        	bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	Name 		string 		`json:"name,omitempty" bson:"name"` // binding:"required"
	Creator		string 		`json:"creator,omitempty" form: creator bson:"creator"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: create_at bson:"create_at"`
	LastUpdate 	time.Time	`json:"last_update,omitempty" form: last_update bson:"last_update"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"` // binding:"required"
	Description 	string		`json:"description,omitempty" form: description bson:"description"`
	Type 		int 		`json: "type,omitempty" form:"type" bson:"type"`
	Status 		int		`json:"status,omitempty" form: "status" bson:"status"`
}
