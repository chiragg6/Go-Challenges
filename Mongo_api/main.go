// RESTful API with Gorilla MUX and MongoDB

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

type Movie struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitemoty"`
	Name      string        `json:"name" bson:"name"`
	Year      string        `json:"year" bson:"year"`
	Directors []string
	Writers   []string
	BoxOffice BoxOffice
}

type BoxOffice struct {
	Budget uint64
	Gross  uint64
}

func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.WriteHeader(http.StatusOK)
	var movie Movie
	err := db.collection.Find(bson.M{"_id": id}).One(&movie)

	if err != nil {
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}

}
