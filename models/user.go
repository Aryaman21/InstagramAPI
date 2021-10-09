package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	User_id    string        `json:"user_id" bson:"user_id"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	User_Posts []string      `json:"user_posts" bson:"user_posts"`
}

type Post struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Post_id   string        `json:"post_id" bson:"post_id"`
	Caption   string        `json:"caption" bson:"caption"`
	Image_url string        `json:"image_url" bson:"image_url"`
	Timestamp string        `json:"timestamp" bson:"timestamp"`
}
