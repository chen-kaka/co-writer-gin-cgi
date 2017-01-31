package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionRepository = "repository"
)

type Repository struct {
	Id        	bson.ObjectId 	`json:"_id,omitempty" bson:"_id,omitempty"`
	name 		string 		`json:"name" form: name`
	creator		string 		`json:"creator" form: creator`
	create_at	string		`json:"create_at" form: create_at`
	last_update 	int64		`json:"last_update" form: last_update`
	u_id 		bson.ObjectId	`json:"u_id" form: u_id`
	description 	string		`json:"creator" form: creator`
	_type 		int 		`json: type form: type`
	status 		int		`json:"status" form: status`
}
