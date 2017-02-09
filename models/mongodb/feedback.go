package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionFeedback = "feedback"
)

//用户反馈评论
type Feedback struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"`
	Title 		string 		`json:"title,omitempty" form:"title" bson:"title"`
	Content 	string		`json:"content,omitempty" form: "content" bson:"content"`
}