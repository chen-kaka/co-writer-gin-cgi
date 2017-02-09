package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionUserTimeline = "user_timeline"
)

//用户时间线,用于接收更新消息
type UserTimeline struct {
	Id 		bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
	CreateAt	time.Time	`json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	LastUpdate 	time.Time	`json:"last_update,omitempty" form: last_update bson:"last_update"`
	UId 		bson.ObjectId	`json:"u_id,omitempty" form: u_id bson:"u_id"`
	Username 	string 		`json:"username,omitempty" bson:"username"`
	Timeline 	bson.M		`json:"timeline,omitempty" form: "timeline" bson:"timeline"`
}