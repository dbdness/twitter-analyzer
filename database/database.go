package database

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

const hostURL = "mongodb://localhost:27017"
const dbName = "social_net"
const collName = "tweets"

var db *mgo.Database
var coll *mgo.Collection
var session *mgo.Session

type tweet struct {
	ID       int
	Polarity int
	Date     string
	Query    string
	User     string
	Text     string
}

func init() {
	var err error
	session, err = GetSession()
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dbName)
	coll = db.C(collName)

}

//GetSession returns the mgo MongoDB session.
func GetSession() (*mgo.Session, error) {
	var err error
	session, err = mgo.Dial(hostURL)
	if err != nil {
		log.Fatal(err)
	}

	return session, err
}

//GetUserCount returns the amount of users in the database.
func GetUserCount() int {
	var result []string
	err := coll.Find(nil).Distinct("user", &result)
	if err != nil {
		log.Fatal(err)
	}

	return len(result)

}
