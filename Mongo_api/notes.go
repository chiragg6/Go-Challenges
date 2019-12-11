package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	Name      string   `bson:"name"`
	Year      string   `bson:"year"`
	Directors []string `bson:"directors"`
	Writers   []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017") // making a connection with db using mgo.Dial()
	if err != nil {
		// panic(err)
		fmt.Println("connection not formed", err)
	}

	defer session.Close()

	c := session.DB("appdb").C("movies") // Fetch a collection using the DB and C functions in a chained manner

	// create a movie
	darkNight := &Movie{
		Name:      "The dark knight",
		Year:      "2008",
		Directors: []string{"ABC"},
		Writers:   []string{"BAC", "KBS"},
		BoxOffice: BoxOffice{
			Budget: 101093,
			Gross:  53321352,
		},
	}

	err = c.Insert(darkNight)
	if err != nil {
		log.Fatal(err)

	}

	var result Movie
	// bson.M is used for nested fields

	err = c.Find(bson.M{"BoxOffice.Budget": bson.M{"$gt": 101093}}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie:", result.Name)
}
