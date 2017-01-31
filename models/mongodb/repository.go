package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionRepository = "repository"
)

type Repository struct {
	Id        	bson.ObjectId 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string 		`json:"name" bson:"name"`
	Creator		string 		`json:"creator" form: creator bson:"creator"`
	CreateAt	int64		`json:"create_at" form: create_at bson:"create_at"`
	LastUpdate 	time.Time		`json:"last_update" form: last_update bson:"last_update"`
	UId 		bson.ObjectId	`json:"u_id" form: u_id bson:"u_id"`
	Description 	string		`json:"description" form: description bson:"description"`
	Type 		int 		`json: type form: type bson:"type"`
	Status 		int		`json:"status" form: status bson:"status"`
}