package main

import (
	"net/http"

	"github.com/Aryaman21/Instant-Insta/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.GET("/posts/:id", uc.GetPost)
	r.POST("/posts", uc.CreatePost)
	r.GET("/post/user/:id", uc.GetPostsOfUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017/")
	if err != nil {
		panic(err)
	}
	return s
}
