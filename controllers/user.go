package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Aryaman21/Instant-Insta/models"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// Get User
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("Instant-Insta").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// Create User
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	Bytes, pas_err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if pas_err != nil {
		panic(pas_err)
	}
	u.Password = string(Bytes)
	uc.session.DB("Instant-Insta").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// Delete User

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("Instant-Intsa").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	fmt.Fprint(w, "Deleted user", oid, "\n")

}

// Get Posts
func (uc UserController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid := bson.ObjectIdHex(id)
	u := models.Post{}

	if err := uc.session.DB("Instant-Insta").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// Create Post
func (uc UserController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Post{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	current_time := time.Now()
	u.Timestamp = current_time.String()
	uc.session.DB("Instant-Insta").C("posts").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// Get Posts of a User
func (uc UserController) GetPostsOfUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("Instant-Insta").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "-----------------USER INFORMATION-------------")
	fmt.Fprintf(w, "%s\n", uj)
	fmt.Fprintf(w, "-----------------USER POSTS-------------")
	for i := 0; i < len(u.User_Posts); i++ {
		if !bson.IsObjectIdHex(u.User_Posts[i]) {
			w.WriteHeader(http.StatusNotFound)
		}
		poid := bson.ObjectIdHex(u.User_Posts[i])
		pu := models.Post{}

		if perr := uc.session.DB("Instant-Insta").C("posts").FindId(poid).One(&pu); perr != nil {
			w.WriteHeader(404)
			return
		}

		pj, perr := json.Marshal(pu)
		if perr != nil {
			fmt.Println(perr)
		}
		fmt.Fprintf(w, "%s\n", pj)
	}

}
